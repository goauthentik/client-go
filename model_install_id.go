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

// checks if the InstallID type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstallID{}

// InstallID struct for InstallID
type InstallID struct {
	InstallId string `json:"install_id"`
}

type _InstallID InstallID

// NewInstallID instantiates a new InstallID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallID(installId string) *InstallID {
	this := InstallID{}
	this.InstallId = installId
	return &this
}

// NewInstallIDWithDefaults instantiates a new InstallID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallIDWithDefaults() *InstallID {
	this := InstallID{}
	return &this
}

// GetInstallId returns the InstallId field value
func (o *InstallID) GetInstallId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstallId
}

// GetInstallIdOk returns a tuple with the InstallId field value
// and a boolean to check if the value has been set.
func (o *InstallID) GetInstallIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstallId, true
}

// SetInstallId sets field value
func (o *InstallID) SetInstallId(v string) {
	o.InstallId = v
}

func (o InstallID) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstallID) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["install_id"] = o.InstallId
	return toSerialize, nil
}

func (o *InstallID) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"install_id",
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

	varInstallID := _InstallID{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInstallID)

	if err != nil {
		return err
	}

	*o = InstallID(varInstallID)

	return err
}

type NullableInstallID struct {
	value *InstallID
	isSet bool
}

func (v NullableInstallID) Get() *InstallID {
	return v.value
}

func (v *NullableInstallID) Set(val *InstallID) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallID) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallID(val *InstallID) *NullableInstallID {
	return &NullableInstallID{value: val, isSet: true}
}

func (v NullableInstallID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
