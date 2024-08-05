/*
authentik

Making authentication simple.

API version: 2024.6.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedSCIMSourceRequest SCIMSource Serializer
type PatchedSCIMSourceRequest struct {
	// Source's display Name.
	Name *string `json:"name,omitempty"`
	// Internal source name, used in URLs.
	Slug                  *string  `json:"slug,omitempty"`
	Enabled               *bool    `json:"enabled,omitempty"`
	UserPropertyMappings  []string `json:"user_property_mappings,omitempty"`
	GroupPropertyMappings []string `json:"group_property_mappings,omitempty"`
	UserPathTemplate      *string  `json:"user_path_template,omitempty"`
}

// NewPatchedSCIMSourceRequest instantiates a new PatchedSCIMSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedSCIMSourceRequest() *PatchedSCIMSourceRequest {
	this := PatchedSCIMSourceRequest{}
	return &this
}

// NewPatchedSCIMSourceRequestWithDefaults instantiates a new PatchedSCIMSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedSCIMSourceRequestWithDefaults() *PatchedSCIMSourceRequest {
	this := PatchedSCIMSourceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedSCIMSourceRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMSourceRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedSCIMSourceRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedSCIMSourceRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PatchedSCIMSourceRequest) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMSourceRequest) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PatchedSCIMSourceRequest) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PatchedSCIMSourceRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchedSCIMSourceRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMSourceRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchedSCIMSourceRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchedSCIMSourceRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUserPropertyMappings returns the UserPropertyMappings field value if set, zero value otherwise.
func (o *PatchedSCIMSourceRequest) GetUserPropertyMappings() []string {
	if o == nil || o.UserPropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.UserPropertyMappings
}

// GetUserPropertyMappingsOk returns a tuple with the UserPropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMSourceRequest) GetUserPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.UserPropertyMappings == nil {
		return nil, false
	}
	return o.UserPropertyMappings, true
}

// HasUserPropertyMappings returns a boolean if a field has been set.
func (o *PatchedSCIMSourceRequest) HasUserPropertyMappings() bool {
	if o != nil && o.UserPropertyMappings != nil {
		return true
	}

	return false
}

// SetUserPropertyMappings gets a reference to the given []string and assigns it to the UserPropertyMappings field.
func (o *PatchedSCIMSourceRequest) SetUserPropertyMappings(v []string) {
	o.UserPropertyMappings = v
}

// GetGroupPropertyMappings returns the GroupPropertyMappings field value if set, zero value otherwise.
func (o *PatchedSCIMSourceRequest) GetGroupPropertyMappings() []string {
	if o == nil || o.GroupPropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.GroupPropertyMappings
}

// GetGroupPropertyMappingsOk returns a tuple with the GroupPropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMSourceRequest) GetGroupPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.GroupPropertyMappings == nil {
		return nil, false
	}
	return o.GroupPropertyMappings, true
}

// HasGroupPropertyMappings returns a boolean if a field has been set.
func (o *PatchedSCIMSourceRequest) HasGroupPropertyMappings() bool {
	if o != nil && o.GroupPropertyMappings != nil {
		return true
	}

	return false
}

// SetGroupPropertyMappings gets a reference to the given []string and assigns it to the GroupPropertyMappings field.
func (o *PatchedSCIMSourceRequest) SetGroupPropertyMappings(v []string) {
	o.GroupPropertyMappings = v
}

// GetUserPathTemplate returns the UserPathTemplate field value if set, zero value otherwise.
func (o *PatchedSCIMSourceRequest) GetUserPathTemplate() string {
	if o == nil || o.UserPathTemplate == nil {
		var ret string
		return ret
	}
	return *o.UserPathTemplate
}

// GetUserPathTemplateOk returns a tuple with the UserPathTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMSourceRequest) GetUserPathTemplateOk() (*string, bool) {
	if o == nil || o.UserPathTemplate == nil {
		return nil, false
	}
	return o.UserPathTemplate, true
}

// HasUserPathTemplate returns a boolean if a field has been set.
func (o *PatchedSCIMSourceRequest) HasUserPathTemplate() bool {
	if o != nil && o.UserPathTemplate != nil {
		return true
	}

	return false
}

// SetUserPathTemplate gets a reference to the given string and assigns it to the UserPathTemplate field.
func (o *PatchedSCIMSourceRequest) SetUserPathTemplate(v string) {
	o.UserPathTemplate = &v
}

func (o PatchedSCIMSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.UserPropertyMappings != nil {
		toSerialize["user_property_mappings"] = o.UserPropertyMappings
	}
	if o.GroupPropertyMappings != nil {
		toSerialize["group_property_mappings"] = o.GroupPropertyMappings
	}
	if o.UserPathTemplate != nil {
		toSerialize["user_path_template"] = o.UserPathTemplate
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedSCIMSourceRequest struct {
	value *PatchedSCIMSourceRequest
	isSet bool
}

func (v NullablePatchedSCIMSourceRequest) Get() *PatchedSCIMSourceRequest {
	return v.value
}

func (v *NullablePatchedSCIMSourceRequest) Set(val *PatchedSCIMSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedSCIMSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedSCIMSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedSCIMSourceRequest(val *PatchedSCIMSourceRequest) *NullablePatchedSCIMSourceRequest {
	return &NullablePatchedSCIMSourceRequest{value: val, isSet: true}
}

func (v NullablePatchedSCIMSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedSCIMSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
