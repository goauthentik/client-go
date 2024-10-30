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

// PatchedDomainRequest Domain Serializer
type PatchedDomainRequest struct {
	Domain    *string `json:"domain,omitempty"`
	IsPrimary *bool   `json:"is_primary,omitempty"`
	Tenant    *string `json:"tenant,omitempty"`
}

// NewPatchedDomainRequest instantiates a new PatchedDomainRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedDomainRequest() *PatchedDomainRequest {
	this := PatchedDomainRequest{}
	return &this
}

// NewPatchedDomainRequestWithDefaults instantiates a new PatchedDomainRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedDomainRequestWithDefaults() *PatchedDomainRequest {
	this := PatchedDomainRequest{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *PatchedDomainRequest) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDomainRequest) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *PatchedDomainRequest) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *PatchedDomainRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *PatchedDomainRequest) GetIsPrimary() bool {
	if o == nil || o.IsPrimary == nil {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDomainRequest) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || o.IsPrimary == nil {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *PatchedDomainRequest) HasIsPrimary() bool {
	if o != nil && o.IsPrimary != nil {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *PatchedDomainRequest) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *PatchedDomainRequest) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDomainRequest) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedDomainRequest) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *PatchedDomainRequest) SetTenant(v string) {
	o.Tenant = &v
}

func (o PatchedDomainRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.IsPrimary != nil {
		toSerialize["is_primary"] = o.IsPrimary
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedDomainRequest struct {
	value *PatchedDomainRequest
	isSet bool
}

func (v NullablePatchedDomainRequest) Get() *PatchedDomainRequest {
	return v.value
}

func (v *NullablePatchedDomainRequest) Set(val *PatchedDomainRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedDomainRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedDomainRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedDomainRequest(val *PatchedDomainRequest) *NullablePatchedDomainRequest {
	return &NullablePatchedDomainRequest{value: val, isSet: true}
}

func (v NullablePatchedDomainRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedDomainRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
