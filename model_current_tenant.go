/*
authentik

Making authentication simple.

API version: 2023.10.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CurrentTenant Partial tenant information for styling
type CurrentTenant struct {
	MatchedDomain      string       `json:"matched_domain"`
	BrandingTitle      string       `json:"branding_title"`
	BrandingLogo       string       `json:"branding_logo"`
	BrandingFavicon    string       `json:"branding_favicon"`
	UiFooterLinks      []FooterLink `json:"ui_footer_links"`
	UiTheme            UiThemeEnum  `json:"ui_theme"`
	FlowAuthentication *string      `json:"flow_authentication,omitempty"`
	FlowInvalidation   *string      `json:"flow_invalidation,omitempty"`
	FlowRecovery       *string      `json:"flow_recovery,omitempty"`
	FlowUnenrollment   *string      `json:"flow_unenrollment,omitempty"`
	FlowUserSettings   *string      `json:"flow_user_settings,omitempty"`
	FlowDeviceCode     *string      `json:"flow_device_code,omitempty"`
	DefaultLocale      string       `json:"default_locale"`
}

// NewCurrentTenant instantiates a new CurrentTenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentTenant(matchedDomain string, brandingTitle string, brandingLogo string, brandingFavicon string, uiFooterLinks []FooterLink, uiTheme UiThemeEnum, defaultLocale string) *CurrentTenant {
	this := CurrentTenant{}
	this.MatchedDomain = matchedDomain
	this.BrandingTitle = brandingTitle
	this.BrandingLogo = brandingLogo
	this.BrandingFavicon = brandingFavicon
	this.UiFooterLinks = uiFooterLinks
	this.UiTheme = uiTheme
	this.DefaultLocale = defaultLocale
	return &this
}

// NewCurrentTenantWithDefaults instantiates a new CurrentTenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentTenantWithDefaults() *CurrentTenant {
	this := CurrentTenant{}
	return &this
}

// GetMatchedDomain returns the MatchedDomain field value
func (o *CurrentTenant) GetMatchedDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MatchedDomain
}

// GetMatchedDomainOk returns a tuple with the MatchedDomain field value
// and a boolean to check if the value has been set.
func (o *CurrentTenant) GetMatchedDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchedDomain, true
}

// SetMatchedDomain sets field value
func (o *CurrentTenant) SetMatchedDomain(v string) {
	o.MatchedDomain = v
}

// GetBrandingTitle returns the BrandingTitle field value
func (o *CurrentTenant) GetBrandingTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandingTitle
}

// GetBrandingTitleOk returns a tuple with the BrandingTitle field value
// and a boolean to check if the value has been set.
func (o *CurrentTenant) GetBrandingTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandingTitle, true
}

// SetBrandingTitle sets field value
func (o *CurrentTenant) SetBrandingTitle(v string) {
	o.BrandingTitle = v
}

// GetBrandingLogo returns the BrandingLogo field value
func (o *CurrentTenant) GetBrandingLogo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandingLogo
}

// GetBrandingLogoOk returns a tuple with the BrandingLogo field value
// and a boolean to check if the value has been set.
func (o *CurrentTenant) GetBrandingLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandingLogo, true
}

// SetBrandingLogo sets field value
func (o *CurrentTenant) SetBrandingLogo(v string) {
	o.BrandingLogo = v
}

// GetBrandingFavicon returns the BrandingFavicon field value
func (o *CurrentTenant) GetBrandingFavicon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandingFavicon
}

// GetBrandingFaviconOk returns a tuple with the BrandingFavicon field value
// and a boolean to check if the value has been set.
func (o *CurrentTenant) GetBrandingFaviconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandingFavicon, true
}

// SetBrandingFavicon sets field value
func (o *CurrentTenant) SetBrandingFavicon(v string) {
	o.BrandingFavicon = v
}

// GetUiFooterLinks returns the UiFooterLinks field value
func (o *CurrentTenant) GetUiFooterLinks() []FooterLink {
	if o == nil {
		var ret []FooterLink
		return ret
	}

	return o.UiFooterLinks
}

// GetUiFooterLinksOk returns a tuple with the UiFooterLinks field value
// and a boolean to check if the value has been set.
func (o *CurrentTenant) GetUiFooterLinksOk() ([]FooterLink, bool) {
	if o == nil {
		return nil, false
	}
	return o.UiFooterLinks, true
}

// SetUiFooterLinks sets field value
func (o *CurrentTenant) SetUiFooterLinks(v []FooterLink) {
	o.UiFooterLinks = v
}

// GetUiTheme returns the UiTheme field value
func (o *CurrentTenant) GetUiTheme() UiThemeEnum {
	if o == nil {
		var ret UiThemeEnum
		return ret
	}

	return o.UiTheme
}

// GetUiThemeOk returns a tuple with the UiTheme field value
// and a boolean to check if the value has been set.
func (o *CurrentTenant) GetUiThemeOk() (*UiThemeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UiTheme, true
}

// SetUiTheme sets field value
func (o *CurrentTenant) SetUiTheme(v UiThemeEnum) {
	o.UiTheme = v
}

// GetFlowAuthentication returns the FlowAuthentication field value if set, zero value otherwise.
func (o *CurrentTenant) GetFlowAuthentication() string {
	if o == nil || o.FlowAuthentication == nil {
		var ret string
		return ret
	}
	return *o.FlowAuthentication
}

// GetFlowAuthenticationOk returns a tuple with the FlowAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentTenant) GetFlowAuthenticationOk() (*string, bool) {
	if o == nil || o.FlowAuthentication == nil {
		return nil, false
	}
	return o.FlowAuthentication, true
}

// HasFlowAuthentication returns a boolean if a field has been set.
func (o *CurrentTenant) HasFlowAuthentication() bool {
	if o != nil && o.FlowAuthentication != nil {
		return true
	}

	return false
}

// SetFlowAuthentication gets a reference to the given string and assigns it to the FlowAuthentication field.
func (o *CurrentTenant) SetFlowAuthentication(v string) {
	o.FlowAuthentication = &v
}

// GetFlowInvalidation returns the FlowInvalidation field value if set, zero value otherwise.
func (o *CurrentTenant) GetFlowInvalidation() string {
	if o == nil || o.FlowInvalidation == nil {
		var ret string
		return ret
	}
	return *o.FlowInvalidation
}

// GetFlowInvalidationOk returns a tuple with the FlowInvalidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentTenant) GetFlowInvalidationOk() (*string, bool) {
	if o == nil || o.FlowInvalidation == nil {
		return nil, false
	}
	return o.FlowInvalidation, true
}

// HasFlowInvalidation returns a boolean if a field has been set.
func (o *CurrentTenant) HasFlowInvalidation() bool {
	if o != nil && o.FlowInvalidation != nil {
		return true
	}

	return false
}

// SetFlowInvalidation gets a reference to the given string and assigns it to the FlowInvalidation field.
func (o *CurrentTenant) SetFlowInvalidation(v string) {
	o.FlowInvalidation = &v
}

// GetFlowRecovery returns the FlowRecovery field value if set, zero value otherwise.
func (o *CurrentTenant) GetFlowRecovery() string {
	if o == nil || o.FlowRecovery == nil {
		var ret string
		return ret
	}
	return *o.FlowRecovery
}

// GetFlowRecoveryOk returns a tuple with the FlowRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentTenant) GetFlowRecoveryOk() (*string, bool) {
	if o == nil || o.FlowRecovery == nil {
		return nil, false
	}
	return o.FlowRecovery, true
}

// HasFlowRecovery returns a boolean if a field has been set.
func (o *CurrentTenant) HasFlowRecovery() bool {
	if o != nil && o.FlowRecovery != nil {
		return true
	}

	return false
}

// SetFlowRecovery gets a reference to the given string and assigns it to the FlowRecovery field.
func (o *CurrentTenant) SetFlowRecovery(v string) {
	o.FlowRecovery = &v
}

// GetFlowUnenrollment returns the FlowUnenrollment field value if set, zero value otherwise.
func (o *CurrentTenant) GetFlowUnenrollment() string {
	if o == nil || o.FlowUnenrollment == nil {
		var ret string
		return ret
	}
	return *o.FlowUnenrollment
}

// GetFlowUnenrollmentOk returns a tuple with the FlowUnenrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentTenant) GetFlowUnenrollmentOk() (*string, bool) {
	if o == nil || o.FlowUnenrollment == nil {
		return nil, false
	}
	return o.FlowUnenrollment, true
}

// HasFlowUnenrollment returns a boolean if a field has been set.
func (o *CurrentTenant) HasFlowUnenrollment() bool {
	if o != nil && o.FlowUnenrollment != nil {
		return true
	}

	return false
}

// SetFlowUnenrollment gets a reference to the given string and assigns it to the FlowUnenrollment field.
func (o *CurrentTenant) SetFlowUnenrollment(v string) {
	o.FlowUnenrollment = &v
}

// GetFlowUserSettings returns the FlowUserSettings field value if set, zero value otherwise.
func (o *CurrentTenant) GetFlowUserSettings() string {
	if o == nil || o.FlowUserSettings == nil {
		var ret string
		return ret
	}
	return *o.FlowUserSettings
}

// GetFlowUserSettingsOk returns a tuple with the FlowUserSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentTenant) GetFlowUserSettingsOk() (*string, bool) {
	if o == nil || o.FlowUserSettings == nil {
		return nil, false
	}
	return o.FlowUserSettings, true
}

// HasFlowUserSettings returns a boolean if a field has been set.
func (o *CurrentTenant) HasFlowUserSettings() bool {
	if o != nil && o.FlowUserSettings != nil {
		return true
	}

	return false
}

// SetFlowUserSettings gets a reference to the given string and assigns it to the FlowUserSettings field.
func (o *CurrentTenant) SetFlowUserSettings(v string) {
	o.FlowUserSettings = &v
}

// GetFlowDeviceCode returns the FlowDeviceCode field value if set, zero value otherwise.
func (o *CurrentTenant) GetFlowDeviceCode() string {
	if o == nil || o.FlowDeviceCode == nil {
		var ret string
		return ret
	}
	return *o.FlowDeviceCode
}

// GetFlowDeviceCodeOk returns a tuple with the FlowDeviceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentTenant) GetFlowDeviceCodeOk() (*string, bool) {
	if o == nil || o.FlowDeviceCode == nil {
		return nil, false
	}
	return o.FlowDeviceCode, true
}

// HasFlowDeviceCode returns a boolean if a field has been set.
func (o *CurrentTenant) HasFlowDeviceCode() bool {
	if o != nil && o.FlowDeviceCode != nil {
		return true
	}

	return false
}

// SetFlowDeviceCode gets a reference to the given string and assigns it to the FlowDeviceCode field.
func (o *CurrentTenant) SetFlowDeviceCode(v string) {
	o.FlowDeviceCode = &v
}

// GetDefaultLocale returns the DefaultLocale field value
func (o *CurrentTenant) GetDefaultLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultLocale
}

// GetDefaultLocaleOk returns a tuple with the DefaultLocale field value
// and a boolean to check if the value has been set.
func (o *CurrentTenant) GetDefaultLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultLocale, true
}

// SetDefaultLocale sets field value
func (o *CurrentTenant) SetDefaultLocale(v string) {
	o.DefaultLocale = v
}

func (o CurrentTenant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["matched_domain"] = o.MatchedDomain
	}
	if true {
		toSerialize["branding_title"] = o.BrandingTitle
	}
	if true {
		toSerialize["branding_logo"] = o.BrandingLogo
	}
	if true {
		toSerialize["branding_favicon"] = o.BrandingFavicon
	}
	if true {
		toSerialize["ui_footer_links"] = o.UiFooterLinks
	}
	if true {
		toSerialize["ui_theme"] = o.UiTheme
	}
	if o.FlowAuthentication != nil {
		toSerialize["flow_authentication"] = o.FlowAuthentication
	}
	if o.FlowInvalidation != nil {
		toSerialize["flow_invalidation"] = o.FlowInvalidation
	}
	if o.FlowRecovery != nil {
		toSerialize["flow_recovery"] = o.FlowRecovery
	}
	if o.FlowUnenrollment != nil {
		toSerialize["flow_unenrollment"] = o.FlowUnenrollment
	}
	if o.FlowUserSettings != nil {
		toSerialize["flow_user_settings"] = o.FlowUserSettings
	}
	if o.FlowDeviceCode != nil {
		toSerialize["flow_device_code"] = o.FlowDeviceCode
	}
	if true {
		toSerialize["default_locale"] = o.DefaultLocale
	}
	return json.Marshal(toSerialize)
}

type NullableCurrentTenant struct {
	value *CurrentTenant
	isSet bool
}

func (v NullableCurrentTenant) Get() *CurrentTenant {
	return v.value
}

func (v *NullableCurrentTenant) Set(val *CurrentTenant) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentTenant) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentTenant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentTenant(val *CurrentTenant) *NullableCurrentTenant {
	return &NullableCurrentTenant{value: val, isSet: true}
}

func (v NullableCurrentTenant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentTenant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
