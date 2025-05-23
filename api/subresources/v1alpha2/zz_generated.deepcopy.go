//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright Flant JSC

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha2

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineAddVolume) DeepCopyInto(out *VirtualMachineAddVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineAddVolume.
func (in *VirtualMachineAddVolume) DeepCopy() *VirtualMachineAddVolume {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineAddVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineAddVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineCancelEvacuation) DeepCopyInto(out *VirtualMachineCancelEvacuation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.DryRun != nil {
		in, out := &in.DryRun, &out.DryRun
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineCancelEvacuation.
func (in *VirtualMachineCancelEvacuation) DeepCopy() *VirtualMachineCancelEvacuation {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineCancelEvacuation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineCancelEvacuation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineConsole) DeepCopyInto(out *VirtualMachineConsole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineConsole.
func (in *VirtualMachineConsole) DeepCopy() *VirtualMachineConsole {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineConsole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineConsole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineFreeze) DeepCopyInto(out *VirtualMachineFreeze) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.UnfreezeTimeout != nil {
		in, out := &in.UnfreezeTimeout, &out.UnfreezeTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineFreeze.
func (in *VirtualMachineFreeze) DeepCopy() *VirtualMachineFreeze {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineFreeze)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineFreeze) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachinePortForward) DeepCopyInto(out *VirtualMachinePortForward) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachinePortForward.
func (in *VirtualMachinePortForward) DeepCopy() *VirtualMachinePortForward {
	if in == nil {
		return nil
	}
	out := new(VirtualMachinePortForward)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachinePortForward) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineRemoveVolume) DeepCopyInto(out *VirtualMachineRemoveVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineRemoveVolume.
func (in *VirtualMachineRemoveVolume) DeepCopy() *VirtualMachineRemoveVolume {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineRemoveVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineRemoveVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineUnfreeze) DeepCopyInto(out *VirtualMachineUnfreeze) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineUnfreeze.
func (in *VirtualMachineUnfreeze) DeepCopy() *VirtualMachineUnfreeze {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineUnfreeze)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineUnfreeze) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineVNC) DeepCopyInto(out *VirtualMachineVNC) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineVNC.
func (in *VirtualMachineVNC) DeepCopy() *VirtualMachineVNC {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineVNC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineVNC) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
