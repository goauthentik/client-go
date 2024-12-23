/*
authentik

Making authentication simple.

API version: 2024.12.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedGroupKerberosSourceConnectionRequest OAuth Group-Source connection Serializer
type PatchedGroupKerberosSourceConnectionRequest struct {
	Group      *string `json:"group,omitempty"`
	Source     *string `json:"source,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
}

// NewPatchedGroupKerberosSourceConnectionRequest instantiates a new PatchedGroupKerberosSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedGroupKerberosSourceConnectionRequest() *PatchedGroupKerberosSourceConnectionRequest {
	this := PatchedGroupKerberosSourceConnectionRequest{}
	return &this
}

// NewPatchedGroupKerberosSourceConnectionRequestWithDefaults instantiates a new PatchedGroupKerberosSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedGroupKerberosSourceConnectionRequestWithDefaults() *PatchedGroupKerberosSourceConnectionRequest {
	this := PatchedGroupKerberosSourceConnectionRequest{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *PatchedGroupKerberosSourceConnectionRequest) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGroupKerberosSourceConnectionRequest) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *PatchedGroupKerberosSourceConnectionRequest) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *PatchedGroupKerberosSourceConnectionRequest) SetGroup(v string) {
	o.Group = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PatchedGroupKerberosSourceConnectionRequest) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGroupKerberosSourceConnectionRequest) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PatchedGroupKerberosSourceConnectionRequest) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *PatchedGroupKerberosSourceConnectionRequest) SetSource(v string) {
	o.Source = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *PatchedGroupKerberosSourceConnectionRequest) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGroupKerberosSourceConnectionRequest) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *PatchedGroupKerberosSourceConnectionRequest) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *PatchedGroupKerberosSourceConnectionRequest) SetIdentifier(v string) {
	o.Identifier = &v
}

func (o PatchedGroupKerberosSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedGroupKerberosSourceConnectionRequest struct {
	value *PatchedGroupKerberosSourceConnectionRequest
	isSet bool
}

func (v NullablePatchedGroupKerberosSourceConnectionRequest) Get() *PatchedGroupKerberosSourceConnectionRequest {
	return v.value
}

func (v *NullablePatchedGroupKerberosSourceConnectionRequest) Set(val *PatchedGroupKerberosSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedGroupKerberosSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedGroupKerberosSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedGroupKerberosSourceConnectionRequest(val *PatchedGroupKerberosSourceConnectionRequest) *NullablePatchedGroupKerberosSourceConnectionRequest {
	return &NullablePatchedGroupKerberosSourceConnectionRequest{value: val, isSet: true}
}

func (v NullablePatchedGroupKerberosSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedGroupKerberosSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
