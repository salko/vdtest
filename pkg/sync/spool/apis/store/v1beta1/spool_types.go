package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SpoolSpec defines the desired state of Spool
type SpoolSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Spool. Edit Spool_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// SpoolStatus defines the observed state of Spool
type SpoolStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Spool is the Schema for the spools API
type Spool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpoolSpec   `json:"spec,omitempty"`
	Status SpoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpoolList contains a list of Spool
type SpoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Spool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Spool{}, &SpoolList{})
}
