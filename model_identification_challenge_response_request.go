/*
authentik

Making authentication simple.

API version: 2022.11.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// IdentificationChallengeResponseRequest Identification challenge
type IdentificationChallengeResponseRequest struct {
	Component *string        `json:"component,omitempty"`
	UidField  string         `json:"uid_field"`
	Password  NullableString `json:"password,omitempty"`
}

// NewIdentificationChallengeResponseRequest instantiates a new IdentificationChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentificationChallengeResponseRequest(uidField string) *IdentificationChallengeResponseRequest {
	this := IdentificationChallengeResponseRequest{}
	var component string = "ak-stage-identification"
	this.Component = &component
	this.UidField = uidField
	return &this
}

// NewIdentificationChallengeResponseRequestWithDefaults instantiates a new IdentificationChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentificationChallengeResponseRequestWithDefaults() *IdentificationChallengeResponseRequest {
	this := IdentificationChallengeResponseRequest{}
	var component string = "ak-stage-identification"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *IdentificationChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *IdentificationChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *IdentificationChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

// GetUidField returns the UidField field value
func (o *IdentificationChallengeResponseRequest) GetUidField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UidField
}

// GetUidFieldOk returns a tuple with the UidField field value
// and a boolean to check if the value has been set.
func (o *IdentificationChallengeResponseRequest) GetUidFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UidField, true
}

// SetUidField sets field value
func (o *IdentificationChallengeResponseRequest) SetUidField(v string) {
	o.UidField = v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentificationChallengeResponseRequest) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentificationChallengeResponseRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *IdentificationChallengeResponseRequest) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *IdentificationChallengeResponseRequest) SetPassword(v string) {
	o.Password.Set(&v)
}

// SetPasswordNil sets the value for Password to be an explicit nil
func (o *IdentificationChallengeResponseRequest) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *IdentificationChallengeResponseRequest) UnsetPassword() {
	o.Password.Unset()
}

func (o IdentificationChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["uid_field"] = o.UidField
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIdentificationChallengeResponseRequest struct {
	value *IdentificationChallengeResponseRequest
	isSet bool
}

func (v NullableIdentificationChallengeResponseRequest) Get() *IdentificationChallengeResponseRequest {
	return v.value
}

func (v *NullableIdentificationChallengeResponseRequest) Set(val *IdentificationChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentificationChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentificationChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentificationChallengeResponseRequest(val *IdentificationChallengeResponseRequest) *NullableIdentificationChallengeResponseRequest {
	return &NullableIdentificationChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableIdentificationChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentificationChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
