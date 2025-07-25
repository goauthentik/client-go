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

// DomainRequest Domain Serializer
type DomainRequest struct {
	Domain    string `json:"domain"`
	IsPrimary *bool  `json:"is_primary,omitempty"`
	Tenant    string `json:"tenant"`
}

// NewDomainRequest instantiates a new DomainRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainRequest(domain string, tenant string) *DomainRequest {
	this := DomainRequest{}
	this.Domain = domain
	this.Tenant = tenant
	return &this
}

// NewDomainRequestWithDefaults instantiates a new DomainRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainRequestWithDefaults() *DomainRequest {
	this := DomainRequest{}
	return &this
}

// GetDomain returns the Domain field value
func (o *DomainRequest) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *DomainRequest) SetDomain(v string) {
	o.Domain = v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *DomainRequest) GetIsPrimary() bool {
	if o == nil || o.IsPrimary == nil {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || o.IsPrimary == nil {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *DomainRequest) HasIsPrimary() bool {
	if o != nil && o.IsPrimary != nil {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *DomainRequest) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetTenant returns the Tenant field value
func (o *DomainRequest) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *DomainRequest) SetTenant(v string) {
	o.Tenant = v
}

func (o DomainRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if o.IsPrimary != nil {
		toSerialize["is_primary"] = o.IsPrimary
	}
	if true {
		toSerialize["tenant"] = o.Tenant
	}
	return json.Marshal(toSerialize)
}

type NullableDomainRequest struct {
	value *DomainRequest
	isSet bool
}

func (v NullableDomainRequest) Get() *DomainRequest {
	return v.value
}

func (v *NullableDomainRequest) Set(val *DomainRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainRequest(val *DomainRequest) *NullableDomainRequest {
	return &NullableDomainRequest{value: val, isSet: true}
}

func (v NullableDomainRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
