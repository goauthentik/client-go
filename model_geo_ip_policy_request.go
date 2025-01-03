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

// GeoIPPolicyRequest GeoIP Policy Serializer
type GeoIPPolicyRequest struct {
	Name string `json:"name"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool             `json:"execution_logging,omitempty"`
	Asns             []int32           `json:"asns,omitempty"`
	Countries        []CountryCodeEnum `json:"countries"`
}

// NewGeoIPPolicyRequest instantiates a new GeoIPPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoIPPolicyRequest(name string, countries []CountryCodeEnum) *GeoIPPolicyRequest {
	this := GeoIPPolicyRequest{}
	this.Name = name
	this.Countries = countries
	return &this
}

// NewGeoIPPolicyRequestWithDefaults instantiates a new GeoIPPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoIPPolicyRequestWithDefaults() *GeoIPPolicyRequest {
	this := GeoIPPolicyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *GeoIPPolicyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GeoIPPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GeoIPPolicyRequest) SetName(v string) {
	o.Name = v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *GeoIPPolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoIPPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *GeoIPPolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *GeoIPPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetAsns returns the Asns field value if set, zero value otherwise.
func (o *GeoIPPolicyRequest) GetAsns() []int32 {
	if o == nil || o.Asns == nil {
		var ret []int32
		return ret
	}
	return o.Asns
}

// GetAsnsOk returns a tuple with the Asns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoIPPolicyRequest) GetAsnsOk() ([]int32, bool) {
	if o == nil || o.Asns == nil {
		return nil, false
	}
	return o.Asns, true
}

// HasAsns returns a boolean if a field has been set.
func (o *GeoIPPolicyRequest) HasAsns() bool {
	if o != nil && o.Asns != nil {
		return true
	}

	return false
}

// SetAsns gets a reference to the given []int32 and assigns it to the Asns field.
func (o *GeoIPPolicyRequest) SetAsns(v []int32) {
	o.Asns = v
}

// GetCountries returns the Countries field value
func (o *GeoIPPolicyRequest) GetCountries() []CountryCodeEnum {
	if o == nil {
		var ret []CountryCodeEnum
		return ret
	}

	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value
// and a boolean to check if the value has been set.
func (o *GeoIPPolicyRequest) GetCountriesOk() ([]CountryCodeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Countries, true
}

// SetCountries sets field value
func (o *GeoIPPolicyRequest) SetCountries(v []CountryCodeEnum) {
	o.Countries = v
}

func (o GeoIPPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if o.Asns != nil {
		toSerialize["asns"] = o.Asns
	}
	if true {
		toSerialize["countries"] = o.Countries
	}
	return json.Marshal(toSerialize)
}

type NullableGeoIPPolicyRequest struct {
	value *GeoIPPolicyRequest
	isSet bool
}

func (v NullableGeoIPPolicyRequest) Get() *GeoIPPolicyRequest {
	return v.value
}

func (v *NullableGeoIPPolicyRequest) Set(val *GeoIPPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoIPPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoIPPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoIPPolicyRequest(val *GeoIPPolicyRequest) *NullableGeoIPPolicyRequest {
	return &NullableGeoIPPolicyRequest{value: val, isSet: true}
}

func (v NullableGeoIPPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoIPPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
