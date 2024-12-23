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

// ConsentPermission Permission used for consent
type ConsentPermission struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

// NewConsentPermission instantiates a new ConsentPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsentPermission(name string, id string) *ConsentPermission {
	this := ConsentPermission{}
	this.Name = name
	this.Id = id
	return &this
}

// NewConsentPermissionWithDefaults instantiates a new ConsentPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsentPermissionWithDefaults() *ConsentPermission {
	this := ConsentPermission{}
	return &this
}

// GetName returns the Name field value
func (o *ConsentPermission) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConsentPermission) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConsentPermission) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *ConsentPermission) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConsentPermission) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConsentPermission) SetId(v string) {
	o.Id = v
}

func (o ConsentPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableConsentPermission struct {
	value *ConsentPermission
	isSet bool
}

func (v NullableConsentPermission) Get() *ConsentPermission {
	return v.value
}

func (v *NullableConsentPermission) Set(val *ConsentPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentPermission(val *ConsentPermission) *NullableConsentPermission {
	return &NullableConsentPermission{value: val, isSet: true}
}

func (v NullableConsentPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
