/*
authentik

Making authentication simple.

API version: 2023.8.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ApplicationRequest Application Serializer
type ApplicationRequest struct {
	// Application's display Name.
	Name string `json:"name"`
	// Internal application name, used in URLs.
	Slug                 string        `json:"slug"`
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

// NewApplicationRequest instantiates a new ApplicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationRequest(name string, slug string) *ApplicationRequest {
	this := ApplicationRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewApplicationRequestWithDefaults instantiates a new ApplicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationRequestWithDefaults() *ApplicationRequest {
	this := ApplicationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ApplicationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *ApplicationRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *ApplicationRequest) SetSlug(v string) {
	o.Slug = v
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationRequest) GetProvider() int32 {
	if o == nil || o.Provider.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationRequest) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *ApplicationRequest) HasProvider() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableInt32 and assigns it to the Provider field.
func (o *ApplicationRequest) SetProvider(v int32) {
	o.Provider.Set(&v)
}

// SetProviderNil sets the value for Provider to be an explicit nil
func (o *ApplicationRequest) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *ApplicationRequest) UnsetProvider() {
	o.Provider.Unset()
}

// GetBackchannelProviders returns the BackchannelProviders field value if set, zero value otherwise.
func (o *ApplicationRequest) GetBackchannelProviders() []int32 {
	if o == nil || o.BackchannelProviders == nil {
		var ret []int32
		return ret
	}
	return o.BackchannelProviders
}

// GetBackchannelProvidersOk returns a tuple with the BackchannelProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetBackchannelProvidersOk() ([]int32, bool) {
	if o == nil || o.BackchannelProviders == nil {
		return nil, false
	}
	return o.BackchannelProviders, true
}

// HasBackchannelProviders returns a boolean if a field has been set.
func (o *ApplicationRequest) HasBackchannelProviders() bool {
	if o != nil && o.BackchannelProviders != nil {
		return true
	}

	return false
}

// SetBackchannelProviders gets a reference to the given []int32 and assigns it to the BackchannelProviders field.
func (o *ApplicationRequest) SetBackchannelProviders(v []int32) {
	o.BackchannelProviders = v
}

// GetOpenInNewTab returns the OpenInNewTab field value if set, zero value otherwise.
func (o *ApplicationRequest) GetOpenInNewTab() bool {
	if o == nil || o.OpenInNewTab == nil {
		var ret bool
		return ret
	}
	return *o.OpenInNewTab
}

// GetOpenInNewTabOk returns a tuple with the OpenInNewTab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetOpenInNewTabOk() (*bool, bool) {
	if o == nil || o.OpenInNewTab == nil {
		return nil, false
	}
	return o.OpenInNewTab, true
}

// HasOpenInNewTab returns a boolean if a field has been set.
func (o *ApplicationRequest) HasOpenInNewTab() bool {
	if o != nil && o.OpenInNewTab != nil {
		return true
	}

	return false
}

// SetOpenInNewTab gets a reference to the given bool and assigns it to the OpenInNewTab field.
func (o *ApplicationRequest) SetOpenInNewTab(v bool) {
	o.OpenInNewTab = &v
}

// GetMetaLaunchUrl returns the MetaLaunchUrl field value if set, zero value otherwise.
func (o *ApplicationRequest) GetMetaLaunchUrl() string {
	if o == nil || o.MetaLaunchUrl == nil {
		var ret string
		return ret
	}
	return *o.MetaLaunchUrl
}

// GetMetaLaunchUrlOk returns a tuple with the MetaLaunchUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetMetaLaunchUrlOk() (*string, bool) {
	if o == nil || o.MetaLaunchUrl == nil {
		return nil, false
	}
	return o.MetaLaunchUrl, true
}

// HasMetaLaunchUrl returns a boolean if a field has been set.
func (o *ApplicationRequest) HasMetaLaunchUrl() bool {
	if o != nil && o.MetaLaunchUrl != nil {
		return true
	}

	return false
}

// SetMetaLaunchUrl gets a reference to the given string and assigns it to the MetaLaunchUrl field.
func (o *ApplicationRequest) SetMetaLaunchUrl(v string) {
	o.MetaLaunchUrl = &v
}

// GetMetaDescription returns the MetaDescription field value if set, zero value otherwise.
func (o *ApplicationRequest) GetMetaDescription() string {
	if o == nil || o.MetaDescription == nil {
		var ret string
		return ret
	}
	return *o.MetaDescription
}

// GetMetaDescriptionOk returns a tuple with the MetaDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetMetaDescriptionOk() (*string, bool) {
	if o == nil || o.MetaDescription == nil {
		return nil, false
	}
	return o.MetaDescription, true
}

// HasMetaDescription returns a boolean if a field has been set.
func (o *ApplicationRequest) HasMetaDescription() bool {
	if o != nil && o.MetaDescription != nil {
		return true
	}

	return false
}

// SetMetaDescription gets a reference to the given string and assigns it to the MetaDescription field.
func (o *ApplicationRequest) SetMetaDescription(v string) {
	o.MetaDescription = &v
}

// GetMetaPublisher returns the MetaPublisher field value if set, zero value otherwise.
func (o *ApplicationRequest) GetMetaPublisher() string {
	if o == nil || o.MetaPublisher == nil {
		var ret string
		return ret
	}
	return *o.MetaPublisher
}

// GetMetaPublisherOk returns a tuple with the MetaPublisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetMetaPublisherOk() (*string, bool) {
	if o == nil || o.MetaPublisher == nil {
		return nil, false
	}
	return o.MetaPublisher, true
}

// HasMetaPublisher returns a boolean if a field has been set.
func (o *ApplicationRequest) HasMetaPublisher() bool {
	if o != nil && o.MetaPublisher != nil {
		return true
	}

	return false
}

// SetMetaPublisher gets a reference to the given string and assigns it to the MetaPublisher field.
func (o *ApplicationRequest) SetMetaPublisher(v string) {
	o.MetaPublisher = &v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *ApplicationRequest) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *ApplicationRequest) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *ApplicationRequest) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ApplicationRequest) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ApplicationRequest) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *ApplicationRequest) SetGroup(v string) {
	o.Group = &v
}

func (o ApplicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
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

type NullableApplicationRequest struct {
	value *ApplicationRequest
	isSet bool
}

func (v NullableApplicationRequest) Get() *ApplicationRequest {
	return v.value
}

func (v *NullableApplicationRequest) Set(val *ApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationRequest(val *ApplicationRequest) *NullableApplicationRequest {
	return &NullableApplicationRequest{value: val, isSet: true}
}

func (v NullableApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
