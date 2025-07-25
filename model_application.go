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

// Application Application Serializer
type Application struct {
	Pk string `json:"pk"`
	// Application's display Name.
	Name string `json:"name"`
	// Internal application name, used in URLs.
	Slug                    string        `json:"slug"`
	Provider                NullableInt32 `json:"provider,omitempty"`
	ProviderObj             Provider      `json:"provider_obj"`
	BackchannelProviders    []int32       `json:"backchannel_providers,omitempty"`
	BackchannelProvidersObj []Provider    `json:"backchannel_providers_obj"`
	// Allow formatting of launch URL
	LaunchUrl NullableString `json:"launch_url"`
	// Open launch URL in a new browser tab or window.
	OpenInNewTab  *bool   `json:"open_in_new_tab,omitempty"`
	MetaLaunchUrl *string `json:"meta_launch_url,omitempty"`
	// Get the URL to the App Icon image. If the name is /static or starts with http it is returned as-is
	MetaIcon         NullableString    `json:"meta_icon"`
	MetaDescription  *string           `json:"meta_description,omitempty"`
	MetaPublisher    *string           `json:"meta_publisher,omitempty"`
	PolicyEngineMode *PolicyEngineMode `json:"policy_engine_mode,omitempty"`
	Group            *string           `json:"group,omitempty"`
}

// NewApplication instantiates a new Application object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplication(pk string, name string, slug string, providerObj Provider, backchannelProvidersObj []Provider, launchUrl NullableString, metaIcon NullableString) *Application {
	this := Application{}
	this.Pk = pk
	this.Name = name
	this.Slug = slug
	this.ProviderObj = providerObj
	this.BackchannelProvidersObj = backchannelProvidersObj
	this.LaunchUrl = launchUrl
	this.MetaIcon = metaIcon
	return &this
}

// NewApplicationWithDefaults instantiates a new Application object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationWithDefaults() *Application {
	this := Application{}
	return &this
}

// GetPk returns the Pk field value
func (o *Application) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *Application) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *Application) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *Application) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Application) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Application) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *Application) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Application) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Application) SetSlug(v string) {
	o.Slug = v
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Application) GetProvider() int32 {
	if o == nil || o.Provider.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *Application) HasProvider() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableInt32 and assigns it to the Provider field.
func (o *Application) SetProvider(v int32) {
	o.Provider.Set(&v)
}

// SetProviderNil sets the value for Provider to be an explicit nil
func (o *Application) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *Application) UnsetProvider() {
	o.Provider.Unset()
}

// GetProviderObj returns the ProviderObj field value
func (o *Application) GetProviderObj() Provider {
	if o == nil {
		var ret Provider
		return ret
	}

	return o.ProviderObj
}

// GetProviderObjOk returns a tuple with the ProviderObj field value
// and a boolean to check if the value has been set.
func (o *Application) GetProviderObjOk() (*Provider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderObj, true
}

// SetProviderObj sets field value
func (o *Application) SetProviderObj(v Provider) {
	o.ProviderObj = v
}

// GetBackchannelProviders returns the BackchannelProviders field value if set, zero value otherwise.
func (o *Application) GetBackchannelProviders() []int32 {
	if o == nil || o.BackchannelProviders == nil {
		var ret []int32
		return ret
	}
	return o.BackchannelProviders
}

// GetBackchannelProvidersOk returns a tuple with the BackchannelProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetBackchannelProvidersOk() ([]int32, bool) {
	if o == nil || o.BackchannelProviders == nil {
		return nil, false
	}
	return o.BackchannelProviders, true
}

// HasBackchannelProviders returns a boolean if a field has been set.
func (o *Application) HasBackchannelProviders() bool {
	if o != nil && o.BackchannelProviders != nil {
		return true
	}

	return false
}

// SetBackchannelProviders gets a reference to the given []int32 and assigns it to the BackchannelProviders field.
func (o *Application) SetBackchannelProviders(v []int32) {
	o.BackchannelProviders = v
}

// GetBackchannelProvidersObj returns the BackchannelProvidersObj field value
func (o *Application) GetBackchannelProvidersObj() []Provider {
	if o == nil {
		var ret []Provider
		return ret
	}

	return o.BackchannelProvidersObj
}

// GetBackchannelProvidersObjOk returns a tuple with the BackchannelProvidersObj field value
// and a boolean to check if the value has been set.
func (o *Application) GetBackchannelProvidersObjOk() ([]Provider, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackchannelProvidersObj, true
}

// SetBackchannelProvidersObj sets field value
func (o *Application) SetBackchannelProvidersObj(v []Provider) {
	o.BackchannelProvidersObj = v
}

// GetLaunchUrl returns the LaunchUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Application) GetLaunchUrl() string {
	if o == nil || o.LaunchUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.LaunchUrl.Get()
}

// GetLaunchUrlOk returns a tuple with the LaunchUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetLaunchUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LaunchUrl.Get(), o.LaunchUrl.IsSet()
}

// SetLaunchUrl sets field value
func (o *Application) SetLaunchUrl(v string) {
	o.LaunchUrl.Set(&v)
}

// GetOpenInNewTab returns the OpenInNewTab field value if set, zero value otherwise.
func (o *Application) GetOpenInNewTab() bool {
	if o == nil || o.OpenInNewTab == nil {
		var ret bool
		return ret
	}
	return *o.OpenInNewTab
}

// GetOpenInNewTabOk returns a tuple with the OpenInNewTab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetOpenInNewTabOk() (*bool, bool) {
	if o == nil || o.OpenInNewTab == nil {
		return nil, false
	}
	return o.OpenInNewTab, true
}

// HasOpenInNewTab returns a boolean if a field has been set.
func (o *Application) HasOpenInNewTab() bool {
	if o != nil && o.OpenInNewTab != nil {
		return true
	}

	return false
}

// SetOpenInNewTab gets a reference to the given bool and assigns it to the OpenInNewTab field.
func (o *Application) SetOpenInNewTab(v bool) {
	o.OpenInNewTab = &v
}

// GetMetaLaunchUrl returns the MetaLaunchUrl field value if set, zero value otherwise.
func (o *Application) GetMetaLaunchUrl() string {
	if o == nil || o.MetaLaunchUrl == nil {
		var ret string
		return ret
	}
	return *o.MetaLaunchUrl
}

// GetMetaLaunchUrlOk returns a tuple with the MetaLaunchUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetMetaLaunchUrlOk() (*string, bool) {
	if o == nil || o.MetaLaunchUrl == nil {
		return nil, false
	}
	return o.MetaLaunchUrl, true
}

// HasMetaLaunchUrl returns a boolean if a field has been set.
func (o *Application) HasMetaLaunchUrl() bool {
	if o != nil && o.MetaLaunchUrl != nil {
		return true
	}

	return false
}

// SetMetaLaunchUrl gets a reference to the given string and assigns it to the MetaLaunchUrl field.
func (o *Application) SetMetaLaunchUrl(v string) {
	o.MetaLaunchUrl = &v
}

// GetMetaIcon returns the MetaIcon field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Application) GetMetaIcon() string {
	if o == nil || o.MetaIcon.Get() == nil {
		var ret string
		return ret
	}

	return *o.MetaIcon.Get()
}

// GetMetaIconOk returns a tuple with the MetaIcon field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetMetaIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetaIcon.Get(), o.MetaIcon.IsSet()
}

// SetMetaIcon sets field value
func (o *Application) SetMetaIcon(v string) {
	o.MetaIcon.Set(&v)
}

// GetMetaDescription returns the MetaDescription field value if set, zero value otherwise.
func (o *Application) GetMetaDescription() string {
	if o == nil || o.MetaDescription == nil {
		var ret string
		return ret
	}
	return *o.MetaDescription
}

// GetMetaDescriptionOk returns a tuple with the MetaDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetMetaDescriptionOk() (*string, bool) {
	if o == nil || o.MetaDescription == nil {
		return nil, false
	}
	return o.MetaDescription, true
}

// HasMetaDescription returns a boolean if a field has been set.
func (o *Application) HasMetaDescription() bool {
	if o != nil && o.MetaDescription != nil {
		return true
	}

	return false
}

// SetMetaDescription gets a reference to the given string and assigns it to the MetaDescription field.
func (o *Application) SetMetaDescription(v string) {
	o.MetaDescription = &v
}

// GetMetaPublisher returns the MetaPublisher field value if set, zero value otherwise.
func (o *Application) GetMetaPublisher() string {
	if o == nil || o.MetaPublisher == nil {
		var ret string
		return ret
	}
	return *o.MetaPublisher
}

// GetMetaPublisherOk returns a tuple with the MetaPublisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetMetaPublisherOk() (*string, bool) {
	if o == nil || o.MetaPublisher == nil {
		return nil, false
	}
	return o.MetaPublisher, true
}

// HasMetaPublisher returns a boolean if a field has been set.
func (o *Application) HasMetaPublisher() bool {
	if o != nil && o.MetaPublisher != nil {
		return true
	}

	return false
}

// SetMetaPublisher gets a reference to the given string and assigns it to the MetaPublisher field.
func (o *Application) SetMetaPublisher(v string) {
	o.MetaPublisher = &v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *Application) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *Application) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *Application) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *Application) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *Application) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *Application) SetGroup(v string) {
	o.Group = &v
}

func (o Application) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}
	if true {
		toSerialize["provider_obj"] = o.ProviderObj
	}
	if o.BackchannelProviders != nil {
		toSerialize["backchannel_providers"] = o.BackchannelProviders
	}
	if true {
		toSerialize["backchannel_providers_obj"] = o.BackchannelProvidersObj
	}
	if true {
		toSerialize["launch_url"] = o.LaunchUrl.Get()
	}
	if o.OpenInNewTab != nil {
		toSerialize["open_in_new_tab"] = o.OpenInNewTab
	}
	if o.MetaLaunchUrl != nil {
		toSerialize["meta_launch_url"] = o.MetaLaunchUrl
	}
	if true {
		toSerialize["meta_icon"] = o.MetaIcon.Get()
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

type NullableApplication struct {
	value *Application
	isSet bool
}

func (v NullableApplication) Get() *Application {
	return v.value
}

func (v *NullableApplication) Set(val *Application) {
	v.value = val
	v.isSet = true
}

func (v NullableApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplication(val *Application) *NullableApplication {
	return &NullableApplication{value: val, isSet: true}
}

func (v NullableApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
