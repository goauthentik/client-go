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

// RadiusCheckAccess Base serializer class which doesn't implement create/update methods
type RadiusCheckAccess struct {
	Attributes *string          `json:"attributes,omitempty"`
	Access     PolicyTestResult `json:"access"`
}

// NewRadiusCheckAccess instantiates a new RadiusCheckAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRadiusCheckAccess(access PolicyTestResult) *RadiusCheckAccess {
	this := RadiusCheckAccess{}
	this.Access = access
	return &this
}

// NewRadiusCheckAccessWithDefaults instantiates a new RadiusCheckAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadiusCheckAccessWithDefaults() *RadiusCheckAccess {
	this := RadiusCheckAccess{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RadiusCheckAccess) GetAttributes() string {
	if o == nil || o.Attributes == nil {
		var ret string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusCheckAccess) GetAttributesOk() (*string, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RadiusCheckAccess) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given string and assigns it to the Attributes field.
func (o *RadiusCheckAccess) SetAttributes(v string) {
	o.Attributes = &v
}

// GetAccess returns the Access field value
func (o *RadiusCheckAccess) GetAccess() PolicyTestResult {
	if o == nil {
		var ret PolicyTestResult
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *RadiusCheckAccess) GetAccessOk() (*PolicyTestResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *RadiusCheckAccess) SetAccess(v PolicyTestResult) {
	o.Access = v
}

func (o RadiusCheckAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableRadiusCheckAccess struct {
	value *RadiusCheckAccess
	isSet bool
}

func (v NullableRadiusCheckAccess) Get() *RadiusCheckAccess {
	return v.value
}

func (v *NullableRadiusCheckAccess) Set(val *RadiusCheckAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableRadiusCheckAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableRadiusCheckAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadiusCheckAccess(val *RadiusCheckAccess) *NullableRadiusCheckAccess {
	return &NullableRadiusCheckAccess{value: val, isSet: true}
}

func (v NullableRadiusCheckAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadiusCheckAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
