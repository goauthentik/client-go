/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the StaticDeviceToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticDeviceToken{}

// StaticDeviceToken Serializer for static device's tokens
type StaticDeviceToken struct {
	Token string `json:"token"`
}

type _StaticDeviceToken StaticDeviceToken

// NewStaticDeviceToken instantiates a new StaticDeviceToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticDeviceToken(token string) *StaticDeviceToken {
	this := StaticDeviceToken{}
	this.Token = token
	return &this
}

// NewStaticDeviceTokenWithDefaults instantiates a new StaticDeviceToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticDeviceTokenWithDefaults() *StaticDeviceToken {
	this := StaticDeviceToken{}
	return &this
}

// GetToken returns the Token field value
func (o *StaticDeviceToken) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *StaticDeviceToken) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *StaticDeviceToken) SetToken(v string) {
	o.Token = v
}

func (o StaticDeviceToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticDeviceToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *StaticDeviceToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varStaticDeviceToken := _StaticDeviceToken{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStaticDeviceToken)

	if err != nil {
		return err
	}

	*o = StaticDeviceToken(varStaticDeviceToken)

	return err
}

type NullableStaticDeviceToken struct {
	value *StaticDeviceToken
	isSet bool
}

func (v NullableStaticDeviceToken) Get() *StaticDeviceToken {
	return v.value
}

func (v *NullableStaticDeviceToken) Set(val *StaticDeviceToken) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticDeviceToken) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticDeviceToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticDeviceToken(val *StaticDeviceToken) *NullableStaticDeviceToken {
	return &NullableStaticDeviceToken{value: val, isSet: true}
}

func (v NullableStaticDeviceToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticDeviceToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
