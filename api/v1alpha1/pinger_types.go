/*
Copyright 2022.

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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PingerSpec defines the desired state of Pinger
type PingerSpec struct {
	//+kubebuilder:validation:MinLength=0

	// Target address to ping
	TargetAddress string `json:"TargetAddress"`
}

// PingerStatus defines the observed state of Pinger
type PingerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// A list of pointers to currently running jobs
	// +optional
	Active []corev1.ObjectReference `json:"active,omitempty"`

	// Last ping time
	// +optional
	LastPingTimes []*metav1.Time `json:"lastPingTimes,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Pinger is the Schema for the pingers API
type Pinger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PingerSpec   `json:"spec,omitempty"`
	Status PingerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PingerList contains a list of Pinger
type PingerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pinger `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Pinger{}, &PingerList{})
}
