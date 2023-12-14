/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PodWatcherSpec defines the desired state of PodWatcher
type PodWatcherSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of PodWatcher. Edit podwatcher_types.go to remove/update
	//+kubebuilder:validation:minimum=0
	Size int32 `json:"size"`
}

// PodWatcherStatus defines the observed state of PodWatcher
type PodWatcherStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Pods []string `json:"pods"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PodWatcher is the Schema for the podwatchers API
type PodWatcher struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodWatcherSpec   `json:"spec,omitempty"`
	Status PodWatcherStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PodWatcherList contains a list of PodWatcher
type PodWatcherList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodWatcher `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PodWatcher{}, &PodWatcherList{})
}
