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
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this ACLTemplate
func (mg *ACLTemplate) GetTerraformResourceType() string {
	return "equinix_network_acl_template"
}

// GetConnectionDetailsMapping for this ACLTemplate
func (tr *ACLTemplate) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ACLTemplate
func (tr *ACLTemplate) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ACLTemplate
func (tr *ACLTemplate) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ACLTemplate
func (tr *ACLTemplate) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ACLTemplate
func (tr *ACLTemplate) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ACLTemplate
func (tr *ACLTemplate) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ACLTemplate using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ACLTemplate) LateInitialize(attrs []byte) (bool, error) {
	params := &ACLTemplateParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ACLTemplate) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BGP
func (mg *BGP) GetTerraformResourceType() string {
	return "equinix_network_bgp"
}

// GetConnectionDetailsMapping for this BGP
func (tr *BGP) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"authentication_key": "spec.forProvider.authenticationKeySecretRef"}
}

// GetObservation of this BGP
func (tr *BGP) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BGP
func (tr *BGP) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BGP
func (tr *BGP) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BGP
func (tr *BGP) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BGP
func (tr *BGP) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BGP using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BGP) LateInitialize(attrs []byte) (bool, error) {
	params := &BGPParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BGP) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Device
func (mg *Device) GetTerraformResourceType() string {
	return "equinix_network_device"
}

// GetConnectionDetailsMapping for this Device
func (tr *Device) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"cluster_details[*].node0[*].license_file_id": "spec.forProvider.clusterDetails[*].node0[*].licenseFileIdSecretRef", "cluster_details[*].node0[*].license_token": "spec.forProvider.clusterDetails[*].node0[*].licenseTokenSecretRef", "cluster_details[*].node0[*].vendor_configuration[*].activation_key": "spec.forProvider.clusterDetails[*].node0[*].vendorConfiguration[*].activationKeySecretRef", "cluster_details[*].node0[*].vendor_configuration[*].admin_password": "spec.forProvider.clusterDetails[*].node0[*].vendorConfiguration[*].adminPasswordSecretRef", "cluster_details[*].node0[*].vendor_configuration[*].root_password": "spec.forProvider.clusterDetails[*].node0[*].vendorConfiguration[*].rootPasswordSecretRef", "cluster_details[*].node1[*].license_file_id": "spec.forProvider.clusterDetails[*].node1[*].licenseFileIdSecretRef", "cluster_details[*].node1[*].license_token": "spec.forProvider.clusterDetails[*].node1[*].licenseTokenSecretRef", "cluster_details[*].node1[*].vendor_configuration[*].activation_key": "spec.forProvider.clusterDetails[*].node1[*].vendorConfiguration[*].activationKeySecretRef", "cluster_details[*].node1[*].vendor_configuration[*].admin_password": "spec.forProvider.clusterDetails[*].node1[*].vendorConfiguration[*].adminPasswordSecretRef", "cluster_details[*].node1[*].vendor_configuration[*].root_password": "spec.forProvider.clusterDetails[*].node1[*].vendorConfiguration[*].rootPasswordSecretRef"}
}

// GetObservation of this Device
func (tr *Device) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Device
func (tr *Device) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Device
func (tr *Device) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Device
func (tr *Device) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Device
func (tr *Device) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Device using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Device) LateInitialize(attrs []byte) (bool, error) {
	params := &DeviceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Device) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this DeviceLink
func (mg *DeviceLink) GetTerraformResourceType() string {
	return "equinix_network_device_link"
}

// GetConnectionDetailsMapping for this DeviceLink
func (tr *DeviceLink) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this DeviceLink
func (tr *DeviceLink) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this DeviceLink
func (tr *DeviceLink) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this DeviceLink
func (tr *DeviceLink) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this DeviceLink
func (tr *DeviceLink) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this DeviceLink
func (tr *DeviceLink) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this DeviceLink using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *DeviceLink) LateInitialize(attrs []byte) (bool, error) {
	params := &DeviceLinkParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *DeviceLink) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SSHKey
func (mg *SSHKey) GetTerraformResourceType() string {
	return "equinix_network_ssh_key"
}

// GetConnectionDetailsMapping for this SSHKey
func (tr *SSHKey) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SSHKey
func (tr *SSHKey) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SSHKey
func (tr *SSHKey) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SSHKey
func (tr *SSHKey) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SSHKey
func (tr *SSHKey) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SSHKey
func (tr *SSHKey) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SSHKey using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SSHKey) LateInitialize(attrs []byte) (bool, error) {
	params := &SSHKeyParameters_2{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SSHKey) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SSHUser
func (mg *SSHUser) GetTerraformResourceType() string {
	return "equinix_network_ssh_user"
}

// GetConnectionDetailsMapping for this SSHUser
func (tr *SSHUser) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"password": "spec.forProvider.passwordSecretRef"}
}

// GetObservation of this SSHUser
func (tr *SSHUser) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SSHUser
func (tr *SSHUser) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SSHUser
func (tr *SSHUser) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SSHUser
func (tr *SSHUser) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SSHUser
func (tr *SSHUser) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SSHUser using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SSHUser) LateInitialize(attrs []byte) (bool, error) {
	params := &SSHUserParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SSHUser) GetTerraformSchemaVersion() int {
	return 0
}
