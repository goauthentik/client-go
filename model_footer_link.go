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

// FooterLink Links returned in Config API
type FooterLink struct {
	Href NullableString `json:"href"`
	Name string         `json:"name"`
}

// NewFooterLink instantiates a new FooterLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFooterLink(href NullableString, name string) *FooterLink {
	this := FooterLink{}
	this.Href = href
	this.Name = name
	return &this
}

// NewFooterLinkWithDefaults instantiates a new FooterLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFooterLinkWithDefaults() *FooterLink {
	this := FooterLink{}
	return &this
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FooterLink) GetHref() string {
	if o == nil || o.Href.Get() == nil {
		var ret string
		return ret
	}

	return *o.Href.Get()
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FooterLink) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Href.Get(), o.Href.IsSet()
}

// SetHref sets field value
func (o *FooterLink) SetHref(v string) {
	o.Href.Set(&v)
}

// GetName returns the Name field value
func (o *FooterLink) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FooterLink) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FooterLink) SetName(v string) {
	o.Name = v
}

func (o FooterLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["href"] = o.Href.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableFooterLink struct {
	value *FooterLink
	isSet bool
}

func (v NullableFooterLink) Get() *FooterLink {
	return v.value
}

func (v *NullableFooterLink) Set(val *FooterLink) {
	v.value = val
	v.isSet = true
}

func (v NullableFooterLink) IsSet() bool {
	return v.isSet
}

func (v *NullableFooterLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFooterLink(val *FooterLink) *NullableFooterLink {
	return &NullableFooterLink{value: val, isSet: true}
}

func (v NullableFooterLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFooterLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
