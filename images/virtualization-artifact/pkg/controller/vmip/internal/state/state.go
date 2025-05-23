/*
Copyright 2024 Flant JSC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package state

import (
	"context"
	"errors"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/deckhouse/virtualization-controller/pkg/common/annotations"
	"github.com/deckhouse/virtualization-controller/pkg/common/ip"
	"github.com/deckhouse/virtualization-controller/pkg/common/object"
	"github.com/deckhouse/virtualization-controller/pkg/controller/conditions"
	"github.com/deckhouse/virtualization-controller/pkg/controller/ipam"
	"github.com/deckhouse/virtualization-controller/pkg/controller/vmip/internal/util"
	"github.com/deckhouse/virtualization-controller/pkg/logger"
	"github.com/deckhouse/virtualization/api/client/kubeclient"
	virtv2 "github.com/deckhouse/virtualization/api/core/v1alpha2"
	"github.com/deckhouse/virtualization/api/core/v1alpha2/vmiplcondition"
)

type VMIPState interface {
	Reload(ctx context.Context) error
	VirtualMachineIP() *virtv2.VirtualMachineIPAddress
	VirtualMachineIPLease() *virtv2.VirtualMachineIPAddressLease
	VirtualMachine() *virtv2.VirtualMachine

	AllocatedIPs() ip.AllocatedIPs
}

type state struct {
	client       client.Client
	virtClient   kubeclient.Client
	vmip         *virtv2.VirtualMachineIPAddress
	lease        *virtv2.VirtualMachineIPAddressLease
	vm           *virtv2.VirtualMachine
	allocatedIPs ip.AllocatedIPs
}

func New(c client.Client, virtClient kubeclient.Client, vmip *virtv2.VirtualMachineIPAddress) VMIPState {
	return &state{client: c, virtClient: virtClient, vmip: vmip}
}

func (s *state) Reload(ctx context.Context) error {
	if err := s.reloadVirtualMachineIPLease(ctx); err != nil {
		return err
	}

	if err := s.reloadVirtualMachine(ctx); err != nil {
		return err
	}

	return nil
}

func (s *state) VirtualMachineIP() *virtv2.VirtualMachineIPAddress {
	return s.vmip
}

func (s *state) VirtualMachineIPLease() *virtv2.VirtualMachineIPAddressLease {
	return s.lease
}

func (s *state) VirtualMachine() *virtv2.VirtualMachine {
	return s.vm
}

func (s *state) AllocatedIPs() ip.AllocatedIPs {
	return s.allocatedIPs
}

func (s *state) reloadVirtualMachineIPLease(ctx context.Context) error {
	var err error
	leaseName := ip.IpToLeaseName(s.vmip.Spec.StaticIP)

	if leaseName == "" {
		leaseName = ip.IpToLeaseName(s.vmip.Status.Address)
	}

	if leaseName != "" {
		leaseKey := types.NamespacedName{Name: leaseName}
		s.lease, err = object.FetchObject(ctx, leaseKey, s.client, &virtv2.VirtualMachineIPAddressLease{})
		if err != nil {
			return fmt.Errorf("unable to get Lease %s: %w", leaseKey, err)
		}
	}

	log := logger.FromContext(ctx)
	if s.lease == nil {
		leases := &virtv2.VirtualMachineIPAddressLeaseList{}

		err = s.client.List(ctx, leases, &client.ListOptions{
			LabelSelector: labels.SelectorFromSet(map[string]string{annotations.LabelVirtualMachineIPAddressUID: string(s.vmip.GetUID())}),
		})

		if err != nil {
			return err
		}

		if len(leases.Items) > 1 {
			log.Error("More than one VirtualMachineIPAddressLease found", "count", len(leases.Items))
		}

		for i, lease := range leases.Items {
			boundCondition, exist := conditions.GetCondition(vmiplcondition.BoundType, lease.Status.Conditions)
			if exist && boundCondition.Status == metav1.ConditionTrue {
				s.lease = &leases.Items[i]
				break
			}
		}
	}

	if s.lease == nil {
		leases, err := s.virtClient.VirtualMachineIPAddressLeases().List(ctx, metav1.ListOptions{
			LabelSelector: fmt.Sprintf("%s=%s", annotations.LabelVirtualMachineIPAddressUID, string(s.vmip.GetUID())),
		})
		if err != nil {
			return err
		}

		if len(leases.Items) != 0 {
			if len(leases.Items) > 1 {
				log.Error("More than one VirtualMachineIPAddressLease found in kubeclient without cache", "count", len(leases.Items))
			}
			log.Warn("VirtualMachineIPAddressLease found in kubeclient without cache", "vmip", s.vmip.Name)
			return errors.New("VirtualMachineIPAddressLease found in kubeclient without cache")
		}
	}

	if s.lease == nil {
		s.allocatedIPs, err = util.GetAllocatedIPs(ctx, s.client, s.vmip.Spec.Type)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *state) reloadVirtualMachine(ctx context.Context) error {
	var err error
	if s.vmip.Status.VirtualMachine != "" {
		vmKey := types.NamespacedName{Name: s.vmip.Status.VirtualMachine, Namespace: s.vmip.Namespace}
		vm, err := object.FetchObject(ctx, vmKey, s.client, &virtv2.VirtualMachine{})
		if err != nil {
			return fmt.Errorf("unable to get VM %s: %w", vmKey, err)
		}

		if vm == nil {
			return nil
		}

		if vm.Status.VirtualMachineIPAddress == s.vmip.Name && vm.Status.IPAddress == s.vmip.Status.Address {
			s.vm = vm
		}
	}

	if s.vm == nil {
		var vms virtv2.VirtualMachineList
		err = s.client.List(ctx, &vms, &client.ListOptions{
			Namespace: s.vmip.Namespace,
		})
		if err != nil {
			return err
		}

		for i, vm := range vms.Items {
			if vm.Spec.VirtualMachineIPAddress == s.vmip.Name || vm.Spec.VirtualMachineIPAddress == "" && vm.Name == ipam.GetVirtualMachineName(s.vmip) {
				s.vm = &vms.Items[i]
				break
			}
		}
	}

	return nil
}
