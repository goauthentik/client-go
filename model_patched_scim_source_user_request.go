/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedSCIMSourceUserRequest SCIMSourceUser Serializer
type PatchedSCIMSourceUserRequest struct {
	Id         *string     `json:"id,omitempty"`
	User       *int32      `json:"user,omitempty"`
	Source     *string     `json:"source,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}

// NewPatchedSCIMSourceUserRequest instantiates a new PatchedSCIMSourceUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedSCIMSourceUserRequest() *PatchedSCIMSourceUserRequest {
	this := PatchedSCIMSourceUserRequest{}
	return &this
}

// NewPatchedSCIMSourceUserRequestWithDefaults instantiates a new PatchedSCIMSourceUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedSCIMSourceUserRequestWithDefaults() *PatchedSCIMSourceUserRequest {
	this := PatchedSCIMSourceUserRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedSCIMSourceUserRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMSourceUserRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedSCIMSourceUserRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedSCIMSourceUserRequest) SetId(v string) {
	o.Id = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *PatchedSCIMSourceUserRequest) GetUser() int32 {
	if o == nil || o.User == nil {
		var ret int32
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMSourceUserRequest) GetUserOk() (*int32, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedSCIMSourceUserRequest) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given int32 and assigns it to the User field.
func (o *PatchedSCIMSourceUserRequest) SetUser(v int32) {
	o.User = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PatchedSCIMSourceUserRequest) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMSourceUserRequest) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PatchedSCIMSourceUserRequest) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *PatchedSCIMSourceUserRequest) SetSource(v string) {
	o.Source = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedSCIMSourceUserRequest) GetAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedSCIMSourceUserRequest) GetAttributesOk() (*interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PatchedSCIMSourceUserRequest) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *PatchedSCIMSourceUserRequest) SetAttributes(v interface{}) {
	o.Attributes = v
}

func (o PatchedSCIMSourceUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedSCIMSourceUserRequest struct {
	value *PatchedSCIMSourceUserRequest
	isSet bool
}

func (v NullablePatchedSCIMSourceUserRequest) Get() *PatchedSCIMSourceUserRequest {
	return v.value
}

func (v *NullablePatchedSCIMSourceUserRequest) Set(val *PatchedSCIMSourceUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedSCIMSourceUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedSCIMSourceUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedSCIMSourceUserRequest(val *PatchedSCIMSourceUserRequest) *NullablePatchedSCIMSourceUserRequest {
	return &NullablePatchedSCIMSourceUserRequest{value: val, isSet: true}
}

func (v NullablePatchedSCIMSourceUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedSCIMSourceUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
