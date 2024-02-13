/*
authentik

Making authentication simple.

API version: 2023.10.7
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SMSDevice Serializer for sms authenticator devices
type SMSDevice struct {
	// The human-readable name of this device.
	Name        string `json:"name"`
	Pk          int32  `json:"pk"`
	PhoneNumber string `json:"phone_number"`
}

// NewSMSDevice instantiates a new SMSDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSMSDevice(name string, pk int32, phoneNumber string) *SMSDevice {
	this := SMSDevice{}
	this.Name = name
	this.Pk = pk
	this.PhoneNumber = phoneNumber
	return &this
}

// NewSMSDeviceWithDefaults instantiates a new SMSDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSMSDeviceWithDefaults() *SMSDevice {
	this := SMSDevice{}
	return &this
}

// GetName returns the Name field value
func (o *SMSDevice) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SMSDevice) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SMSDevice) SetName(v string) {
	o.Name = v
}

// GetPk returns the Pk field value
func (o *SMSDevice) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *SMSDevice) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *SMSDevice) SetPk(v int32) {
	o.Pk = v
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *SMSDevice) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *SMSDevice) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *SMSDevice) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

func (o SMSDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	return json.Marshal(toSerialize)
}

type NullableSMSDevice struct {
	value *SMSDevice
	isSet bool
}

func (v NullableSMSDevice) Get() *SMSDevice {
	return v.value
}

func (v *NullableSMSDevice) Set(val *SMSDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableSMSDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableSMSDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSMSDevice(val *SMSDevice) *NullableSMSDevice {
	return &NullableSMSDevice{value: val, isSet: true}
}

func (v NullableSMSDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSMSDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
