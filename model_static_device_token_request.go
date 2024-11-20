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

// checks if the StaticDeviceTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticDeviceTokenRequest{}

// StaticDeviceTokenRequest Serializer for static device's tokens
type StaticDeviceTokenRequest struct {
	Token string `json:"token"`
}

type _StaticDeviceTokenRequest StaticDeviceTokenRequest

// NewStaticDeviceTokenRequest instantiates a new StaticDeviceTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticDeviceTokenRequest(token string) *StaticDeviceTokenRequest {
	this := StaticDeviceTokenRequest{}
	this.Token = token
	return &this
}

// NewStaticDeviceTokenRequestWithDefaults instantiates a new StaticDeviceTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticDeviceTokenRequestWithDefaults() *StaticDeviceTokenRequest {
	this := StaticDeviceTokenRequest{}
	return &this
}

// GetToken returns the Token field value
func (o *StaticDeviceTokenRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *StaticDeviceTokenRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *StaticDeviceTokenRequest) SetToken(v string) {
	o.Token = v
}

func (o StaticDeviceTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticDeviceTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *StaticDeviceTokenRequest) UnmarshalJSON(data []byte) (err error) {
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

	varStaticDeviceTokenRequest := _StaticDeviceTokenRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStaticDeviceTokenRequest)

	if err != nil {
		return err
	}

	*o = StaticDeviceTokenRequest(varStaticDeviceTokenRequest)

	return err
}

type NullableStaticDeviceTokenRequest struct {
	value *StaticDeviceTokenRequest
	isSet bool
}

func (v NullableStaticDeviceTokenRequest) Get() *StaticDeviceTokenRequest {
	return v.value
}

func (v *NullableStaticDeviceTokenRequest) Set(val *StaticDeviceTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticDeviceTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticDeviceTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticDeviceTokenRequest(val *StaticDeviceTokenRequest) *NullableStaticDeviceTokenRequest {
	return &NullableStaticDeviceTokenRequest{value: val, isSet: true}
}

func (v NullableStaticDeviceTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticDeviceTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
