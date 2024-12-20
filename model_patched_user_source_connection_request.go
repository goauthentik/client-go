/*
authentik

Making authentication simple.

API version: 2024.12.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedUserSourceConnectionRequest User source connection
type PatchedUserSourceConnectionRequest struct {
	User   *int32  `json:"user,omitempty"`
	Source *string `json:"source,omitempty"`
}

// NewPatchedUserSourceConnectionRequest instantiates a new PatchedUserSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedUserSourceConnectionRequest() *PatchedUserSourceConnectionRequest {
	this := PatchedUserSourceConnectionRequest{}
	return &this
}

// NewPatchedUserSourceConnectionRequestWithDefaults instantiates a new PatchedUserSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedUserSourceConnectionRequestWithDefaults() *PatchedUserSourceConnectionRequest {
	this := PatchedUserSourceConnectionRequest{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *PatchedUserSourceConnectionRequest) GetUser() int32 {
	if o == nil || o.User == nil {
		var ret int32
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserSourceConnectionRequest) GetUserOk() (*int32, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedUserSourceConnectionRequest) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given int32 and assigns it to the User field.
func (o *PatchedUserSourceConnectionRequest) SetUser(v int32) {
	o.User = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PatchedUserSourceConnectionRequest) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserSourceConnectionRequest) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PatchedUserSourceConnectionRequest) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *PatchedUserSourceConnectionRequest) SetSource(v string) {
	o.Source = &v
}

func (o PatchedUserSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedUserSourceConnectionRequest struct {
	value *PatchedUserSourceConnectionRequest
	isSet bool
}

func (v NullablePatchedUserSourceConnectionRequest) Get() *PatchedUserSourceConnectionRequest {
	return v.value
}

func (v *NullablePatchedUserSourceConnectionRequest) Set(val *PatchedUserSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedUserSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedUserSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedUserSourceConnectionRequest(val *PatchedUserSourceConnectionRequest) *NullablePatchedUserSourceConnectionRequest {
	return &NullablePatchedUserSourceConnectionRequest{value: val, isSet: true}
}

func (v NullablePatchedUserSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedUserSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}