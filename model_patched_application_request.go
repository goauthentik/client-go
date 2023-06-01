/*
authentik

Making authentication simple.

API version: 2023.5.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedApplicationRequest Application Serializer
type PatchedApplicationRequest struct {
	// Application's display Name.
	Name *string `json:"name,omitempty"`
	// Internal application name, used in URLs.
	Slug                 *string       `json:"slug,omitempty"`
	Provider             NullableInt32 `json:"provider,omitempty"`
	BackchannelProviders []int32       `json:"backchannel_providers,omitempty"`
	// Open launch URL in a new browser tab or window.
	OpenInNewTab     *bool             `json:"open_in_new_tab,omitempty"`
	MetaLaunchUrl    *string           `json:"meta_launch_url,omitempty"`
	MetaDescription  *string           `json:"meta_description,omitempty"`
	MetaPublisher    *string           `json:"meta_publisher,omitempty"`
	PolicyEngineMode *PolicyEngineMode `json:"policy_engine_mode,omitempty"`
	Group            *string           `json:"group,omitempty"`
}

// NewPatchedApplicationRequest instantiates a new PatchedApplicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedApplicationRequest() *PatchedApplicationRequest {
	this := PatchedApplicationRequest{}
	return &this
}

// NewPatchedApplicationRequestWithDefaults instantiates a new PatchedApplicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedApplicationRequestWithDefaults() *PatchedApplicationRequest {
	this := PatchedApplicationRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedApplicationRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedApplicationRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedApplicationRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedApplicationRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PatchedApplicationRequest) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedApplicationRequest) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PatchedApplicationRequest) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PatchedApplicationRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedApplicationRequest) GetProvider() int32 {
	if o == nil || o.Provider.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedApplicationRequest) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *PatchedApplicationRequest) HasProvider() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableInt32 and assigns it to the Provider field.
func (o *PatchedApplicationRequest) SetProvider(v int32) {
	o.Provider.Set(&v)
}

// SetProviderNil sets the value for Provider to be an explicit nil
func (o *PatchedApplicationRequest) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *PatchedApplicationRequest) UnsetProvider() {
	o.Provider.Unset()
}

// GetBackchannelProviders returns the BackchannelProviders field value if set, zero value otherwise.
func (o *PatchedApplicationRequest) GetBackchannelProviders() []int32 {
	if o == nil || o.BackchannelProviders == nil {
		var ret []int32
		return ret
	}
	return o.BackchannelProviders
}

// GetBackchannelProvidersOk returns a tuple with the BackchannelProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedApplicationRequest) GetBackchannelProvidersOk() ([]int32, bool) {
	if o == nil || o.BackchannelProviders == nil {
		return nil, false
	}
	return o.BackchannelProviders, true
}

// HasBackchannelProviders returns a boolean if a field has been set.
func (o *PatchedApplicationRequest) HasBackchannelProviders() bool {
	if o != nil && o.BackchannelProviders != nil {
		return true
	}

	return false
}

// SetBackchannelProviders gets a reference to the given []int32 and assigns it to the BackchannelProviders field.
func (o *PatchedApplicationRequest) SetBackchannelProviders(v []int32) {
	o.BackchannelProviders = v
}

// GetOpenInNewTab returns the OpenInNewTab field value if set, zero value otherwise.
func (o *PatchedApplicationRequest) GetOpenInNewTab() bool {
	if o == nil || o.OpenInNewTab == nil {
		var ret bool
		return ret
	}
	return *o.OpenInNewTab
}

// GetOpenInNewTabOk returns a tuple with the OpenInNewTab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedApplicationRequest) GetOpenInNewTabOk() (*bool, bool) {
	if o == nil || o.OpenInNewTab == nil {
		return nil, false
	}
	return o.OpenInNewTab, true
}

// HasOpenInNewTab returns a boolean if a field has been set.
func (o *PatchedApplicationRequest) HasOpenInNewTab() bool {
	if o != nil && o.OpenInNewTab != nil {
		return true
	}

	return false
}

// SetOpenInNewTab gets a reference to the given bool and assigns it to the OpenInNewTab field.
func (o *PatchedApplicationRequest) SetOpenInNewTab(v bool) {
	o.OpenInNewTab = &v
}

// GetMetaLaunchUrl returns the MetaLaunchUrl field value if set, zero value otherwise.
func (o *PatchedApplicationRequest) GetMetaLaunchUrl() string {
	if o == nil || o.MetaLaunchUrl == nil {
		var ret string
		return ret
	}
	return *o.MetaLaunchUrl
}

// GetMetaLaunchUrlOk returns a tuple with the MetaLaunchUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedApplicationRequest) GetMetaLaunchUrlOk() (*string, bool) {
	if o == nil || o.MetaLaunchUrl == nil {
		return nil, false
	}
	return o.MetaLaunchUrl, true
}

// HasMetaLaunchUrl returns a boolean if a field has been set.
func (o *PatchedApplicationRequest) HasMetaLaunchUrl() bool {
	if o != nil && o.MetaLaunchUrl != nil {
		return true
	}

	return false
}

// SetMetaLaunchUrl gets a reference to the given string and assigns it to the MetaLaunchUrl field.
func (o *PatchedApplicationRequest) SetMetaLaunchUrl(v string) {
	o.MetaLaunchUrl = &v
}

// GetMetaDescription returns the MetaDescription field value if set, zero value otherwise.
func (o *PatchedApplicationRequest) GetMetaDescription() string {
	if o == nil || o.MetaDescription == nil {
		var ret string
		return ret
	}
	return *o.MetaDescription
}

// GetMetaDescriptionOk returns a tuple with the MetaDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedApplicationRequest) GetMetaDescriptionOk() (*string, bool) {
	if o == nil || o.MetaDescription == nil {
		return nil, false
	}
	return o.MetaDescription, true
}

// HasMetaDescription returns a boolean if a field has been set.
func (o *PatchedApplicationRequest) HasMetaDescription() bool {
	if o != nil && o.MetaDescription != nil {
		return true
	}

	return false
}

// SetMetaDescription gets a reference to the given string and assigns it to the MetaDescription field.
func (o *PatchedApplicationRequest) SetMetaDescription(v string) {
	o.MetaDescription = &v
}

// GetMetaPublisher returns the MetaPublisher field value if set, zero value otherwise.
func (o *PatchedApplicationRequest) GetMetaPublisher() string {
	if o == nil || o.MetaPublisher == nil {
		var ret string
		return ret
	}
	return *o.MetaPublisher
}

// GetMetaPublisherOk returns a tuple with the MetaPublisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedApplicationRequest) GetMetaPublisherOk() (*string, bool) {
	if o == nil || o.MetaPublisher == nil {
		return nil, false
	}
	return o.MetaPublisher, true
}

// HasMetaPublisher returns a boolean if a field has been set.
func (o *PatchedApplicationRequest) HasMetaPublisher() bool {
	if o != nil && o.MetaPublisher != nil {
		return true
	}

	return false
}

// SetMetaPublisher gets a reference to the given string and assigns it to the MetaPublisher field.
func (o *PatchedApplicationRequest) SetMetaPublisher(v string) {
	o.MetaPublisher = &v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *PatchedApplicationRequest) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedApplicationRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *PatchedApplicationRequest) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *PatchedApplicationRequest) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *PatchedApplicationRequest) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedApplicationRequest) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *PatchedApplicationRequest) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *PatchedApplicationRequest) SetGroup(v string) {
	o.Group = &v
}

func (o PatchedApplicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}
	if o.BackchannelProviders != nil {
		toSerialize["backchannel_providers"] = o.BackchannelProviders
	}
	if o.OpenInNewTab != nil {
		toSerialize["open_in_new_tab"] = o.OpenInNewTab
	}
	if o.MetaLaunchUrl != nil {
		toSerialize["meta_launch_url"] = o.MetaLaunchUrl
	}
	if o.MetaDescription != nil {
		toSerialize["meta_description"] = o.MetaDescription
	}
	if o.MetaPublisher != nil {
		toSerialize["meta_publisher"] = o.MetaPublisher
	}
	if o.PolicyEngineMode != nil {
		toSerialize["policy_engine_mode"] = o.PolicyEngineMode
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedApplicationRequest struct {
	value *PatchedApplicationRequest
	isSet bool
}

func (v NullablePatchedApplicationRequest) Get() *PatchedApplicationRequest {
	return v.value
}

func (v *NullablePatchedApplicationRequest) Set(val *PatchedApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedApplicationRequest(val *PatchedApplicationRequest) *NullablePatchedApplicationRequest {
	return &NullablePatchedApplicationRequest{value: val, isSet: true}
}

func (v NullablePatchedApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
