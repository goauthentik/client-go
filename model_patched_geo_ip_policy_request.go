/*
authentik

Making authentication simple.

API version: 2024.10.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedGeoIPPolicyRequest GeoIP Policy Serializer
type PatchedGeoIPPolicyRequest struct {
	Name *string `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool             `json:"execution_logging,omitempty"`
	Asns             []int32           `json:"asns,omitempty"`
	Countries        []CountryCodeEnum `json:"countries,omitempty"`
}

// NewPatchedGeoIPPolicyRequest instantiates a new PatchedGeoIPPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedGeoIPPolicyRequest() *PatchedGeoIPPolicyRequest {
	this := PatchedGeoIPPolicyRequest{}
	return &this
}

// NewPatchedGeoIPPolicyRequestWithDefaults instantiates a new PatchedGeoIPPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedGeoIPPolicyRequestWithDefaults() *PatchedGeoIPPolicyRequest {
	this := PatchedGeoIPPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedGeoIPPolicyRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGeoIPPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedGeoIPPolicyRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedGeoIPPolicyRequest) SetName(v string) {
	o.Name = &v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *PatchedGeoIPPolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGeoIPPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *PatchedGeoIPPolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *PatchedGeoIPPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetAsns returns the Asns field value if set, zero value otherwise.
func (o *PatchedGeoIPPolicyRequest) GetAsns() []int32 {
	if o == nil || o.Asns == nil {
		var ret []int32
		return ret
	}
	return o.Asns
}

// GetAsnsOk returns a tuple with the Asns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGeoIPPolicyRequest) GetAsnsOk() ([]int32, bool) {
	if o == nil || o.Asns == nil {
		return nil, false
	}
	return o.Asns, true
}

// HasAsns returns a boolean if a field has been set.
func (o *PatchedGeoIPPolicyRequest) HasAsns() bool {
	if o != nil && o.Asns != nil {
		return true
	}

	return false
}

// SetAsns gets a reference to the given []int32 and assigns it to the Asns field.
func (o *PatchedGeoIPPolicyRequest) SetAsns(v []int32) {
	o.Asns = v
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *PatchedGeoIPPolicyRequest) GetCountries() []CountryCodeEnum {
	if o == nil || o.Countries == nil {
		var ret []CountryCodeEnum
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGeoIPPolicyRequest) GetCountriesOk() ([]CountryCodeEnum, bool) {
	if o == nil || o.Countries == nil {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *PatchedGeoIPPolicyRequest) HasCountries() bool {
	if o != nil && o.Countries != nil {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []CountryCodeEnum and assigns it to the Countries field.
func (o *PatchedGeoIPPolicyRequest) SetCountries(v []CountryCodeEnum) {
	o.Countries = v
}

func (o PatchedGeoIPPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if o.Asns != nil {
		toSerialize["asns"] = o.Asns
	}
	if o.Countries != nil {
		toSerialize["countries"] = o.Countries
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedGeoIPPolicyRequest struct {
	value *PatchedGeoIPPolicyRequest
	isSet bool
}

func (v NullablePatchedGeoIPPolicyRequest) Get() *PatchedGeoIPPolicyRequest {
	return v.value
}

func (v *NullablePatchedGeoIPPolicyRequest) Set(val *PatchedGeoIPPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedGeoIPPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedGeoIPPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedGeoIPPolicyRequest(val *PatchedGeoIPPolicyRequest) *NullablePatchedGeoIPPolicyRequest {
	return &NullablePatchedGeoIPPolicyRequest{value: val, isSet: true}
}

func (v NullablePatchedGeoIPPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedGeoIPPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
