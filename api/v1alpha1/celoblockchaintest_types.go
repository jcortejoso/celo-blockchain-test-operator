/*
Copyright 2021.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CeloBlockchainTestSpec defines the desired state of CeloBlockchainTest
type CeloBlockchainTestSpec struct {
	// RangeSize is the size of the range each job instance will test
	RangeSize int `json:"rangeSize,omitempty"`
	// Replicas is the number of jobs to run
	Replicas int `json:"replicas,omitempty"`
	// GethTag is the celo-blockchain image tag to use in the tests
	GethTag string `json:"gethTag,omitempty"`
}

// CeloBlockchainTestStatus defines the observed state of CeloBlockchainTest
type CeloBlockchainTestStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Jobs are the names of the test jobs
	Jobs []string `json:"jobs"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CeloBlockchainTest is the Schema for the celoblockchaintests API
type CeloBlockchainTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CeloBlockchainTestSpec   `json:"spec,omitempty"`
	Status CeloBlockchainTestStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CeloBlockchainTestList contains a list of CeloBlockchainTest
type CeloBlockchainTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CeloBlockchainTest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CeloBlockchainTest{}, &CeloBlockchainTestList{})
}
