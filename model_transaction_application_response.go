/*
authentik

Making authentication simple.

API version: 2024.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TransactionApplicationResponse Transactional creation response
type TransactionApplicationResponse struct {
	Applied bool     `json:"applied"`
	Logs    []string `json:"logs"`
}

// NewTransactionApplicationResponse instantiates a new TransactionApplicationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionApplicationResponse(applied bool, logs []string) *TransactionApplicationResponse {
	this := TransactionApplicationResponse{}
	this.Applied = applied
	this.Logs = logs
	return &this
}

// NewTransactionApplicationResponseWithDefaults instantiates a new TransactionApplicationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionApplicationResponseWithDefaults() *TransactionApplicationResponse {
	this := TransactionApplicationResponse{}
	return &this
}

// GetApplied returns the Applied field value
func (o *TransactionApplicationResponse) GetApplied() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Applied
}

// GetAppliedOk returns a tuple with the Applied field value
// and a boolean to check if the value has been set.
func (o *TransactionApplicationResponse) GetAppliedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Applied, true
}

// SetApplied sets field value
func (o *TransactionApplicationResponse) SetApplied(v bool) {
	o.Applied = v
}

// GetLogs returns the Logs field value
func (o *TransactionApplicationResponse) GetLogs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value
// and a boolean to check if the value has been set.
func (o *TransactionApplicationResponse) GetLogsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logs, true
}

// SetLogs sets field value
func (o *TransactionApplicationResponse) SetLogs(v []string) {
	o.Logs = v
}

func (o TransactionApplicationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["applied"] = o.Applied
	}
	if true {
		toSerialize["logs"] = o.Logs
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionApplicationResponse struct {
	value *TransactionApplicationResponse
	isSet bool
}

func (v NullableTransactionApplicationResponse) Get() *TransactionApplicationResponse {
	return v.value
}

func (v *NullableTransactionApplicationResponse) Set(val *TransactionApplicationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionApplicationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionApplicationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionApplicationResponse(val *TransactionApplicationResponse) *NullableTransactionApplicationResponse {
	return &NullableTransactionApplicationResponse{value: val, isSet: true}
}

func (v NullableTransactionApplicationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionApplicationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
