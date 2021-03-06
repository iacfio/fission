/*
Copyright 2016 The Fission Authors.

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

package crd

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/fission/fission"
)

//
// To add a Fission CRD type:
//   1. Create a "spec" type, for everything in the type except metadata
//   2. Create the type with metadata + the spec
//   3. Create a list type (for example see FunctionList and Function, below)
//   4. Add methods at the bottom of this file for satisfying Object and List interfaces
//   5. Add the type to configureClient in client.go
//   6. Add the type to EnsureFissionCRDs in crd.go
//   7. Add tests to crd_test.go
//   8. Add a CRUD Interface type (analogous to FunctionInterface in function.go)
//   9. Add a getter method for your interface type to FissionClient in client.go
//

type (
	// Packages. Think of these as function-level images.
	Package struct {
		metav1.TypeMeta `json:",inline"`
		Metadata        metav1.ObjectMeta   `json:"metadata"`
		Spec            fission.PackageSpec `json:"spec"`

		Status fission.PackageStatus `json:"status"`
	}
	PackageList struct {
		metav1.TypeMeta `json:",inline"`
		Metadata        metav1.ListMeta `json:"metadata"`

		Items []Package `json:"items"`
	}

	// Functions.
	Function struct {
		metav1.TypeMeta `json:",inline"`
		Metadata        metav1.ObjectMeta    `json:"metadata"`
		Spec            fission.FunctionSpec `json:"spec"`
	}
	FunctionList struct {
		metav1.TypeMeta `json:",inline"`
		Metadata        metav1.ListMeta `json:"metadata"`

		Items []Function `json:"items"`
	}

	// Environments.
	Environment struct {
		metav1.TypeMeta `json:",inline"`
		Metadata        metav1.ObjectMeta       `json:"metadata"`
		Spec            fission.EnvironmentSpec `json:"spec"`
	}
	EnvironmentList struct {
		metav1.TypeMeta `json:",inline"`
		Metadata        metav1.ListMeta `json:"metadata"`

		Items []Environment `json:"items"`
	}

	HTTPTrigger struct {
		metav1.TypeMeta `json:",inline"`
		Metadata        metav1.ObjectMeta       `json:"metadata"`
		Spec            fission.HTTPTriggerSpec `json:"spec"`
	}
	HTTPTriggerList struct {
		metav1.TypeMeta `json:",inline"`
		Metadata        metav1.ListMeta `json:"metadata"`

		Items []HTTPTrigger `json:"items"`
	}

	// Kubernetes Watches as triggers
	KubernetesWatchTrigger struct {
		metav1.TypeMeta `json:",inline"`
		Metadata        metav1.ObjectMeta                  `json:"metadata"`
		Spec            fission.KubernetesWatchTriggerSpec `json:"spec"`
	}
	KubernetesWatchTriggerList struct {
		metav1.TypeMeta `json:",inline"`
		Metadata        metav1.ListMeta `json:"metadata"`

		Items []KubernetesWatchTrigger `json:"items"`
	}

	// Time triggers
	TimeTrigger struct {
		metav1.TypeMeta `json:",inline"`
		Metadata        metav1.ObjectMeta       `json:"metadata"`
		Spec            fission.TimeTriggerSpec `json:"spec"`
	}
	TimeTriggerList struct {
		metav1.TypeMeta `json:",inline"`
		Metadata        metav1.ListMeta `json:"metadata"`

		Items []TimeTrigger `json:"items"`
	}

	// Message Queue triggers
	MessageQueueTrigger struct {
		metav1.TypeMeta `json:",inline"`
		Metadata        metav1.ObjectMeta               `json:"metadata"`
		Spec            fission.MessageQueueTriggerSpec `json:"spec"`
	}
	MessageQueueTriggerList struct {
		metav1.TypeMeta `json:",inline"`
		Metadata        metav1.ListMeta `json:"metadata"`

		Items []MessageQueueTrigger `json:"items"`
	}
)

// Each CRD type needs:
//   GetObjectKind (to satisfy the Object interface)
//
// In addition, each singular CRD type needs:
//   GetObjectMeta (to satisfy the ObjectMetaAccessor interface)
//
// And each list CRD type needs:
//   GetListMeta (to satisfy the ListMetaAccessor interface)

func (f *Function) GetObjectKind() schema.ObjectKind {
	return &f.TypeMeta
}
func (e *Environment) GetObjectKind() schema.ObjectKind {
	return &e.TypeMeta
}
func (ht *HTTPTrigger) GetObjectKind() schema.ObjectKind {
	return &ht.TypeMeta
}
func (w *KubernetesWatchTrigger) GetObjectKind() schema.ObjectKind {
	return &w.TypeMeta
}
func (w *TimeTrigger) GetObjectKind() schema.ObjectKind {
	return &w.TypeMeta
}
func (w *MessageQueueTrigger) GetObjectKind() schema.ObjectKind {
	return &w.TypeMeta
}
func (w *Package) GetObjectKind() schema.ObjectKind {
	return &w.TypeMeta
}

func (f *Function) GetObjectMeta() metav1.Object {
	return &f.Metadata
}
func (e *Environment) GetObjectMeta() metav1.Object {
	return &e.Metadata
}
func (ht *HTTPTrigger) GetObjectMeta() metav1.Object {
	return &ht.Metadata
}
func (w *KubernetesWatchTrigger) GetObjectMeta() metav1.Object {
	return &w.Metadata
}
func (w *TimeTrigger) GetObjectMeta() metav1.Object {
	return &w.Metadata
}
func (w *MessageQueueTrigger) GetObjectMeta() metav1.Object {
	return &w.Metadata
}
func (w *Package) GetObjectMeta() metav1.Object {
	return &w.Metadata
}

func (fl *FunctionList) GetObjectKind() schema.ObjectKind {
	return &fl.TypeMeta
}
func (el *EnvironmentList) GetObjectKind() schema.ObjectKind {
	return &el.TypeMeta
}
func (hl *HTTPTriggerList) GetObjectKind() schema.ObjectKind {
	return &hl.TypeMeta
}
func (wl *KubernetesWatchTriggerList) GetObjectKind() schema.ObjectKind {
	return &wl.TypeMeta
}
func (wl *TimeTriggerList) GetObjectKind() schema.ObjectKind {
	return &wl.TypeMeta
}
func (wl *MessageQueueTriggerList) GetObjectKind() schema.ObjectKind {
	return &wl.TypeMeta
}
func (wl *PackageList) GetObjectKind() schema.ObjectKind {
	return &wl.TypeMeta
}

func (fl *FunctionList) GetListMeta() metav1.List {
	return &fl.Metadata
}
func (el *EnvironmentList) GetListMeta() metav1.List {
	return &el.Metadata
}
func (hl *HTTPTriggerList) GetListMeta() metav1.List {
	return &hl.Metadata
}
func (wl *KubernetesWatchTriggerList) GetListMeta() metav1.List {
	return &wl.Metadata
}
func (wl *TimeTriggerList) GetListMeta() metav1.List {
	return &wl.Metadata
}
func (wl *MessageQueueTriggerList) GetListMeta() metav1.List {
	return &wl.Metadata
}
func (wl *PackageList) GetListMeta() metav1.List {
	return &wl.Metadata
}
