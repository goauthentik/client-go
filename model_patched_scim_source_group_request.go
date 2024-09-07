/*
authentik

Making authentication simple.

API version: 2024.8.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedSCIMSourceGroupRequest SCIMSourceGroup Serializer
type PatchedSCIMSourceGroupRequest struct {
	Id         *string     `json:"id,omitempty"`
	Group      *string     `json:"group,omitempty"`
	Source     *string     `json:"source,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}

// NewPatchedSCIMSourceGroupRequest instantiates a new PatchedSCIMSourceGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedSCIMSourceGroupRequest() *PatchedSCIMSourceGroupRequest {
	this := PatchedSCIMSourceGroupRequest{}
	return &this
}

// NewPatchedSCIMSourceGroupRequestWithDefaults instantiates a new PatchedSCIMSourceGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedSCIMSourceGroupRequestWithDefaults() *PatchedSCIMSourceGroupRequest {
	this := PatchedSCIMSourceGroupRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedSCIMSourceGroupRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMSourceGroupRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedSCIMSourceGroupRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedSCIMSourceGroupRequest) SetId(v string) {
	o.Id = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *PatchedSCIMSourceGroupRequest) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMSourceGroupRequest) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *PatchedSCIMSourceGroupRequest) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *PatchedSCIMSourceGroupRequest) SetGroup(v string) {
	o.Group = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PatchedSCIMSourceGroupRequest) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMSourceGroupRequest) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PatchedSCIMSourceGroupRequest) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *PatchedSCIMSourceGroupRequest) SetSource(v string) {
	o.Source = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedSCIMSourceGroupRequest) GetAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedSCIMSourceGroupRequest) GetAttributesOk() (*interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PatchedSCIMSourceGroupRequest) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *PatchedSCIMSourceGroupRequest) SetAttributes(v interface{}) {
	o.Attributes = v
}

func (o PatchedSCIMSourceGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedSCIMSourceGroupRequest struct {
	value *PatchedSCIMSourceGroupRequest
	isSet bool
}

func (v NullablePatchedSCIMSourceGroupRequest) Get() *PatchedSCIMSourceGroupRequest {
	return v.value
}

func (v *NullablePatchedSCIMSourceGroupRequest) Set(val *PatchedSCIMSourceGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedSCIMSourceGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedSCIMSourceGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedSCIMSourceGroupRequest(val *PatchedSCIMSourceGroupRequest) *NullablePatchedSCIMSourceGroupRequest {
	return &NullablePatchedSCIMSourceGroupRequest{value: val, isSet: true}
}

func (v NullablePatchedSCIMSourceGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedSCIMSourceGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
