/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type UserAPIKeyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type UserAPIKeyParameters struct {

	// Description string for the API key
	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// Flag indicating whether the API key shoud be read-only
	// +kubebuilder:validation:Required
	ReadOnly *bool `json:"readOnly" tf:"read_only,omitempty"`
}

// UserAPIKeySpec defines the desired state of UserAPIKey
type UserAPIKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserAPIKeyParameters `json:"forProvider"`
}

// UserAPIKeyStatus defines the observed state of UserAPIKey.
type UserAPIKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserAPIKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserAPIKey is the Schema for the UserAPIKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinixjet}
type UserAPIKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserAPIKeySpec   `json:"spec"`
	Status            UserAPIKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserAPIKeyList contains a list of UserAPIKeys
type UserAPIKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserAPIKey `json:"items"`
}

// Repository type metadata.
var (
	UserAPIKey_Kind             = "UserAPIKey"
	UserAPIKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserAPIKey_Kind}.String()
	UserAPIKey_KindAPIVersion   = UserAPIKey_Kind + "." + CRDGroupVersion.String()
	UserAPIKey_GroupVersionKind = CRDGroupVersion.WithKind(UserAPIKey_Kind)
)

func init() {
	SchemeBuilder.Register(&UserAPIKey{}, &UserAPIKeyList{})
}
