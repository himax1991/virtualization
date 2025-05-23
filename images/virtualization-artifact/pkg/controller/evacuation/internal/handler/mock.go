// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package handler

import (
	"context"
	"sync"
)

// Ensure, that EvacuateCancelerMock does implement EvacuateCanceler.
// If this is not the case, regenerate this file with moq.
var _ EvacuateCanceler = &EvacuateCancelerMock{}

// EvacuateCancelerMock is a mock implementation of EvacuateCanceler.
//
//	func TestSomethingThatUsesEvacuateCanceler(t *testing.T) {
//
//		// make and configure a mocked EvacuateCanceler
//		mockedEvacuateCanceler := &EvacuateCancelerMock{
//			CancelFunc: func(ctx context.Context, name string, namespace string) error {
//				panic("mock out the Cancel method")
//			},
//		}
//
//		// use mockedEvacuateCanceler in code that requires EvacuateCanceler
//		// and then make assertions.
//
//	}
type EvacuateCancelerMock struct {
	// CancelFunc mocks the Cancel method.
	CancelFunc func(ctx context.Context, name string, namespace string) error

	// calls tracks calls to the methods.
	calls struct {
		// Cancel holds details about calls to the Cancel method.
		Cancel []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
			// Namespace is the namespace argument value.
			Namespace string
		}
	}
	lockCancel sync.RWMutex
}

// Cancel calls CancelFunc.
func (mock *EvacuateCancelerMock) Cancel(ctx context.Context, name string, namespace string) error {
	if mock.CancelFunc == nil {
		panic("EvacuateCancelerMock.CancelFunc: method is nil but EvacuateCanceler.Cancel was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Name      string
		Namespace string
	}{
		Ctx:       ctx,
		Name:      name,
		Namespace: namespace,
	}
	mock.lockCancel.Lock()
	mock.calls.Cancel = append(mock.calls.Cancel, callInfo)
	mock.lockCancel.Unlock()
	return mock.CancelFunc(ctx, name, namespace)
}

// CancelCalls gets all the calls that were made to Cancel.
// Check the length with:
//
//	len(mockedEvacuateCanceler.CancelCalls())
func (mock *EvacuateCancelerMock) CancelCalls() []struct {
	Ctx       context.Context
	Name      string
	Namespace string
} {
	var calls []struct {
		Ctx       context.Context
		Name      string
		Namespace string
	}
	mock.lockCancel.RLock()
	calls = mock.calls.Cancel
	mock.lockCancel.RUnlock()
	return calls
}
