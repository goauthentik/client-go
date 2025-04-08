/*
authentik

Making authentication simple.

API version: 2025.2.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ImpersonationRequest struct for ImpersonationRequest
type ImpersonationRequest struct {
	Reason string `json:"reason"`
}

// NewImpersonationRequest instantiates a new ImpersonationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImpersonationRequest(reason string) *ImpersonationRequest {
	this := ImpersonationRequest{}
	this.Reason = reason
	return &this
}

// NewImpersonationRequestWithDefaults instantiates a new ImpersonationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImpersonationRequestWithDefaults() *ImpersonationRequest {
	this := ImpersonationRequest{}
	return &this
}

// GetReason returns the Reason field value
func (o *ImpersonationRequest) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ImpersonationRequest) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *ImpersonationRequest) SetReason(v string) {
	o.Reason = v
}

func (o ImpersonationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableImpersonationRequest struct {
	value *ImpersonationRequest
	isSet bool
}

func (v NullableImpersonationRequest) Get() *ImpersonationRequest {
	return v.value
}

func (v *NullableImpersonationRequest) Set(val *ImpersonationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableImpersonationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableImpersonationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImpersonationRequest(val *ImpersonationRequest) *NullableImpersonationRequest {
	return &NullableImpersonationRequest{value: val, isSet: true}
}

func (v NullableImpersonationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImpersonationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
