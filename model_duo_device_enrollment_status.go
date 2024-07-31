/*
authentik

Making authentication simple.

API version: 2024.6.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DuoDeviceEnrollmentStatus struct for DuoDeviceEnrollmentStatus
type DuoDeviceEnrollmentStatus struct {
	DuoResponse DuoResponseEnum `json:"duo_response"`
}

// NewDuoDeviceEnrollmentStatus instantiates a new DuoDeviceEnrollmentStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDuoDeviceEnrollmentStatus(duoResponse DuoResponseEnum) *DuoDeviceEnrollmentStatus {
	this := DuoDeviceEnrollmentStatus{}
	this.DuoResponse = duoResponse
	return &this
}

// NewDuoDeviceEnrollmentStatusWithDefaults instantiates a new DuoDeviceEnrollmentStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDuoDeviceEnrollmentStatusWithDefaults() *DuoDeviceEnrollmentStatus {
	this := DuoDeviceEnrollmentStatus{}
	return &this
}

// GetDuoResponse returns the DuoResponse field value
func (o *DuoDeviceEnrollmentStatus) GetDuoResponse() DuoResponseEnum {
	if o == nil {
		var ret DuoResponseEnum
		return ret
	}

	return o.DuoResponse
}

// GetDuoResponseOk returns a tuple with the DuoResponse field value
// and a boolean to check if the value has been set.
func (o *DuoDeviceEnrollmentStatus) GetDuoResponseOk() (*DuoResponseEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DuoResponse, true
}

// SetDuoResponse sets field value
func (o *DuoDeviceEnrollmentStatus) SetDuoResponse(v DuoResponseEnum) {
	o.DuoResponse = v
}

func (o DuoDeviceEnrollmentStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["duo_response"] = o.DuoResponse
	}
	return json.Marshal(toSerialize)
}

type NullableDuoDeviceEnrollmentStatus struct {
	value *DuoDeviceEnrollmentStatus
	isSet bool
}

func (v NullableDuoDeviceEnrollmentStatus) Get() *DuoDeviceEnrollmentStatus {
	return v.value
}

func (v *NullableDuoDeviceEnrollmentStatus) Set(val *DuoDeviceEnrollmentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDuoDeviceEnrollmentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDuoDeviceEnrollmentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDuoDeviceEnrollmentStatus(val *DuoDeviceEnrollmentStatus) *NullableDuoDeviceEnrollmentStatus {
	return &NullableDuoDeviceEnrollmentStatus{value: val, isSet: true}
}

func (v NullableDuoDeviceEnrollmentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDuoDeviceEnrollmentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
