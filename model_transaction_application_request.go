/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TransactionApplicationRequest Serializer for creating a provider and an application in one transaction
type TransactionApplicationRequest struct {
	App            ApplicationRequest                `json:"app"`
	ProviderModel  ProviderModelEnum                 `json:"provider_model"`
	Provider       ModelRequest                      `json:"provider"`
	PolicyBindings []TransactionPolicyBindingRequest `json:"policy_bindings,omitempty"`
}

// NewTransactionApplicationRequest instantiates a new TransactionApplicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionApplicationRequest(app ApplicationRequest, providerModel ProviderModelEnum, provider ModelRequest) *TransactionApplicationRequest {
	this := TransactionApplicationRequest{}
	this.App = app
	this.ProviderModel = providerModel
	this.Provider = provider
	return &this
}

// NewTransactionApplicationRequestWithDefaults instantiates a new TransactionApplicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionApplicationRequestWithDefaults() *TransactionApplicationRequest {
	this := TransactionApplicationRequest{}
	return &this
}

// GetApp returns the App field value
func (o *TransactionApplicationRequest) GetApp() ApplicationRequest {
	if o == nil {
		var ret ApplicationRequest
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *TransactionApplicationRequest) GetAppOk() (*ApplicationRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *TransactionApplicationRequest) SetApp(v ApplicationRequest) {
	o.App = v
}

// GetProviderModel returns the ProviderModel field value
func (o *TransactionApplicationRequest) GetProviderModel() ProviderModelEnum {
	if o == nil {
		var ret ProviderModelEnum
		return ret
	}

	return o.ProviderModel
}

// GetProviderModelOk returns a tuple with the ProviderModel field value
// and a boolean to check if the value has been set.
func (o *TransactionApplicationRequest) GetProviderModelOk() (*ProviderModelEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderModel, true
}

// SetProviderModel sets field value
func (o *TransactionApplicationRequest) SetProviderModel(v ProviderModelEnum) {
	o.ProviderModel = v
}

// GetProvider returns the Provider field value
func (o *TransactionApplicationRequest) GetProvider() ModelRequest {
	if o == nil {
		var ret ModelRequest
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *TransactionApplicationRequest) GetProviderOk() (*ModelRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *TransactionApplicationRequest) SetProvider(v ModelRequest) {
	o.Provider = v
}

// GetPolicyBindings returns the PolicyBindings field value if set, zero value otherwise.
func (o *TransactionApplicationRequest) GetPolicyBindings() []TransactionPolicyBindingRequest {
	if o == nil || o.PolicyBindings == nil {
		var ret []TransactionPolicyBindingRequest
		return ret
	}
	return o.PolicyBindings
}

// GetPolicyBindingsOk returns a tuple with the PolicyBindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionApplicationRequest) GetPolicyBindingsOk() ([]TransactionPolicyBindingRequest, bool) {
	if o == nil || o.PolicyBindings == nil {
		return nil, false
	}
	return o.PolicyBindings, true
}

// HasPolicyBindings returns a boolean if a field has been set.
func (o *TransactionApplicationRequest) HasPolicyBindings() bool {
	if o != nil && o.PolicyBindings != nil {
		return true
	}

	return false
}

// SetPolicyBindings gets a reference to the given []TransactionPolicyBindingRequest and assigns it to the PolicyBindings field.
func (o *TransactionApplicationRequest) SetPolicyBindings(v []TransactionPolicyBindingRequest) {
	o.PolicyBindings = v
}

func (o TransactionApplicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["app"] = o.App
	}
	if true {
		toSerialize["provider_model"] = o.ProviderModel
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if o.PolicyBindings != nil {
		toSerialize["policy_bindings"] = o.PolicyBindings
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionApplicationRequest struct {
	value *TransactionApplicationRequest
	isSet bool
}

func (v NullableTransactionApplicationRequest) Get() *TransactionApplicationRequest {
	return v.value
}

func (v *NullableTransactionApplicationRequest) Set(val *TransactionApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionApplicationRequest(val *TransactionApplicationRequest) *NullableTransactionApplicationRequest {
	return &NullableTransactionApplicationRequest{value: val, isSet: true}
}

func (v NullableTransactionApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
