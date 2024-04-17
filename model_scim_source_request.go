/*
authentik

Making authentication simple.

API version: 2024.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SCIMSourceRequest SCIMSource Serializer
type SCIMSourceRequest struct {
	// Source's display Name.
	Name string `json:"name"`
	// Internal source name, used in URLs.
	Slug    string `json:"slug"`
	Enabled *bool  `json:"enabled,omitempty"`
	// How the source determines if an existing user should be authenticated or a new user enrolled.
	UserMatchingMode *UserMatchingModeEnum `json:"user_matching_mode,omitempty"`
	UserPathTemplate *string               `json:"user_path_template,omitempty"`
}

// NewSCIMSourceRequest instantiates a new SCIMSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSCIMSourceRequest(name string, slug string) *SCIMSourceRequest {
	this := SCIMSourceRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewSCIMSourceRequestWithDefaults instantiates a new SCIMSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSCIMSourceRequestWithDefaults() *SCIMSourceRequest {
	this := SCIMSourceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *SCIMSourceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SCIMSourceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SCIMSourceRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *SCIMSourceRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *SCIMSourceRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *SCIMSourceRequest) SetSlug(v string) {
	o.Slug = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SCIMSourceRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSourceRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SCIMSourceRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SCIMSourceRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUserMatchingMode returns the UserMatchingMode field value if set, zero value otherwise.
func (o *SCIMSourceRequest) GetUserMatchingMode() UserMatchingModeEnum {
	if o == nil || o.UserMatchingMode == nil {
		var ret UserMatchingModeEnum
		return ret
	}
	return *o.UserMatchingMode
}

// GetUserMatchingModeOk returns a tuple with the UserMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool) {
	if o == nil || o.UserMatchingMode == nil {
		return nil, false
	}
	return o.UserMatchingMode, true
}

// HasUserMatchingMode returns a boolean if a field has been set.
func (o *SCIMSourceRequest) HasUserMatchingMode() bool {
	if o != nil && o.UserMatchingMode != nil {
		return true
	}

	return false
}

// SetUserMatchingMode gets a reference to the given UserMatchingModeEnum and assigns it to the UserMatchingMode field.
func (o *SCIMSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum) {
	o.UserMatchingMode = &v
}

// GetUserPathTemplate returns the UserPathTemplate field value if set, zero value otherwise.
func (o *SCIMSourceRequest) GetUserPathTemplate() string {
	if o == nil || o.UserPathTemplate == nil {
		var ret string
		return ret
	}
	return *o.UserPathTemplate
}

// GetUserPathTemplateOk returns a tuple with the UserPathTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSourceRequest) GetUserPathTemplateOk() (*string, bool) {
	if o == nil || o.UserPathTemplate == nil {
		return nil, false
	}
	return o.UserPathTemplate, true
}

// HasUserPathTemplate returns a boolean if a field has been set.
func (o *SCIMSourceRequest) HasUserPathTemplate() bool {
	if o != nil && o.UserPathTemplate != nil {
		return true
	}

	return false
}

// SetUserPathTemplate gets a reference to the given string and assigns it to the UserPathTemplate field.
func (o *SCIMSourceRequest) SetUserPathTemplate(v string) {
	o.UserPathTemplate = &v
}

func (o SCIMSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.UserMatchingMode != nil {
		toSerialize["user_matching_mode"] = o.UserMatchingMode
	}
	if o.UserPathTemplate != nil {
		toSerialize["user_path_template"] = o.UserPathTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableSCIMSourceRequest struct {
	value *SCIMSourceRequest
	isSet bool
}

func (v NullableSCIMSourceRequest) Get() *SCIMSourceRequest {
	return v.value
}

func (v *NullableSCIMSourceRequest) Set(val *SCIMSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSCIMSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSCIMSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSCIMSourceRequest(val *SCIMSourceRequest) *NullableSCIMSourceRequest {
	return &NullableSCIMSourceRequest{value: val, isSet: true}
}

func (v NullableSCIMSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSCIMSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
