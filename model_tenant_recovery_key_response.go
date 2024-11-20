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
	"time"
)

// checks if the TenantRecoveryKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantRecoveryKeyResponse{}

// TenantRecoveryKeyResponse Tenant recovery key creation response serializer
type TenantRecoveryKeyResponse struct {
	Expiry time.Time `json:"expiry"`
	Url    string    `json:"url"`
}

type _TenantRecoveryKeyResponse TenantRecoveryKeyResponse

// NewTenantRecoveryKeyResponse instantiates a new TenantRecoveryKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantRecoveryKeyResponse(expiry time.Time, url string) *TenantRecoveryKeyResponse {
	this := TenantRecoveryKeyResponse{}
	this.Expiry = expiry
	this.Url = url
	return &this
}

// NewTenantRecoveryKeyResponseWithDefaults instantiates a new TenantRecoveryKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantRecoveryKeyResponseWithDefaults() *TenantRecoveryKeyResponse {
	this := TenantRecoveryKeyResponse{}
	return &this
}

// GetExpiry returns the Expiry field value
func (o *TenantRecoveryKeyResponse) GetExpiry() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value
// and a boolean to check if the value has been set.
func (o *TenantRecoveryKeyResponse) GetExpiryOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiry, true
}

// SetExpiry sets field value
func (o *TenantRecoveryKeyResponse) SetExpiry(v time.Time) {
	o.Expiry = v
}

// GetUrl returns the Url field value
func (o *TenantRecoveryKeyResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TenantRecoveryKeyResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *TenantRecoveryKeyResponse) SetUrl(v string) {
	o.Url = v
}

func (o TenantRecoveryKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantRecoveryKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expiry"] = o.Expiry
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *TenantRecoveryKeyResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"expiry",
		"url",
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

	varTenantRecoveryKeyResponse := _TenantRecoveryKeyResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTenantRecoveryKeyResponse)

	if err != nil {
		return err
	}

	*o = TenantRecoveryKeyResponse(varTenantRecoveryKeyResponse)

	return err
}

type NullableTenantRecoveryKeyResponse struct {
	value *TenantRecoveryKeyResponse
	isSet bool
}

func (v NullableTenantRecoveryKeyResponse) Get() *TenantRecoveryKeyResponse {
	return v.value
}

func (v *NullableTenantRecoveryKeyResponse) Set(val *TenantRecoveryKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantRecoveryKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantRecoveryKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantRecoveryKeyResponse(val *TenantRecoveryKeyResponse) *NullableTenantRecoveryKeyResponse {
	return &NullableTenantRecoveryKeyResponse{value: val, isSet: true}
}

func (v NullableTenantRecoveryKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantRecoveryKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
