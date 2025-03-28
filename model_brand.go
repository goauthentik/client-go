/*
authentik

Making authentication simple.

API version: 2025.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Brand Brand Serializer
type Brand struct {
	BrandUuid string `json:"brand_uuid"`
	// Domain that activates this brand. Can be a superset, i.e. `a.b` for `aa.b` and `ba.b`
	Domain                        string         `json:"domain"`
	Default                       *bool          `json:"default,omitempty"`
	BrandingTitle                 *string        `json:"branding_title,omitempty"`
	BrandingLogo                  *string        `json:"branding_logo,omitempty"`
	BrandingFavicon               *string        `json:"branding_favicon,omitempty"`
	BrandingCustomCss             *string        `json:"branding_custom_css,omitempty"`
	BrandingDefaultFlowBackground *string        `json:"branding_default_flow_background,omitempty"`
	FlowAuthentication            NullableString `json:"flow_authentication,omitempty"`
	FlowInvalidation              NullableString `json:"flow_invalidation,omitempty"`
	FlowRecovery                  NullableString `json:"flow_recovery,omitempty"`
	FlowUnenrollment              NullableString `json:"flow_unenrollment,omitempty"`
	FlowUserSettings              NullableString `json:"flow_user_settings,omitempty"`
	FlowDeviceCode                NullableString `json:"flow_device_code,omitempty"`
	// When set, external users will be redirected to this application after authenticating.
	DefaultApplication NullableString `json:"default_application,omitempty"`
	// Web Certificate used by the authentik Core webserver.
	WebCertificate NullableString `json:"web_certificate,omitempty"`
	Attributes     interface{}    `json:"attributes,omitempty"`
}

// NewBrand instantiates a new Brand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrand(brandUuid string, domain string) *Brand {
	this := Brand{}
	this.BrandUuid = brandUuid
	this.Domain = domain
	return &this
}

// NewBrandWithDefaults instantiates a new Brand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandWithDefaults() *Brand {
	this := Brand{}
	return &this
}

// GetBrandUuid returns the BrandUuid field value
func (o *Brand) GetBrandUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandUuid
}

// GetBrandUuidOk returns a tuple with the BrandUuid field value
// and a boolean to check if the value has been set.
func (o *Brand) GetBrandUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandUuid, true
}

// SetBrandUuid sets field value
func (o *Brand) SetBrandUuid(v string) {
	o.BrandUuid = v
}

// GetDomain returns the Domain field value
func (o *Brand) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *Brand) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *Brand) SetDomain(v string) {
	o.Domain = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *Brand) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *Brand) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *Brand) SetDefault(v bool) {
	o.Default = &v
}

// GetBrandingTitle returns the BrandingTitle field value if set, zero value otherwise.
func (o *Brand) GetBrandingTitle() string {
	if o == nil || o.BrandingTitle == nil {
		var ret string
		return ret
	}
	return *o.BrandingTitle
}

// GetBrandingTitleOk returns a tuple with the BrandingTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetBrandingTitleOk() (*string, bool) {
	if o == nil || o.BrandingTitle == nil {
		return nil, false
	}
	return o.BrandingTitle, true
}

// HasBrandingTitle returns a boolean if a field has been set.
func (o *Brand) HasBrandingTitle() bool {
	if o != nil && o.BrandingTitle != nil {
		return true
	}

	return false
}

// SetBrandingTitle gets a reference to the given string and assigns it to the BrandingTitle field.
func (o *Brand) SetBrandingTitle(v string) {
	o.BrandingTitle = &v
}

// GetBrandingLogo returns the BrandingLogo field value if set, zero value otherwise.
func (o *Brand) GetBrandingLogo() string {
	if o == nil || o.BrandingLogo == nil {
		var ret string
		return ret
	}
	return *o.BrandingLogo
}

// GetBrandingLogoOk returns a tuple with the BrandingLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetBrandingLogoOk() (*string, bool) {
	if o == nil || o.BrandingLogo == nil {
		return nil, false
	}
	return o.BrandingLogo, true
}

// HasBrandingLogo returns a boolean if a field has been set.
func (o *Brand) HasBrandingLogo() bool {
	if o != nil && o.BrandingLogo != nil {
		return true
	}

	return false
}

// SetBrandingLogo gets a reference to the given string and assigns it to the BrandingLogo field.
func (o *Brand) SetBrandingLogo(v string) {
	o.BrandingLogo = &v
}

// GetBrandingFavicon returns the BrandingFavicon field value if set, zero value otherwise.
func (o *Brand) GetBrandingFavicon() string {
	if o == nil || o.BrandingFavicon == nil {
		var ret string
		return ret
	}
	return *o.BrandingFavicon
}

// GetBrandingFaviconOk returns a tuple with the BrandingFavicon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetBrandingFaviconOk() (*string, bool) {
	if o == nil || o.BrandingFavicon == nil {
		return nil, false
	}
	return o.BrandingFavicon, true
}

// HasBrandingFavicon returns a boolean if a field has been set.
func (o *Brand) HasBrandingFavicon() bool {
	if o != nil && o.BrandingFavicon != nil {
		return true
	}

	return false
}

// SetBrandingFavicon gets a reference to the given string and assigns it to the BrandingFavicon field.
func (o *Brand) SetBrandingFavicon(v string) {
	o.BrandingFavicon = &v
}

// GetBrandingCustomCss returns the BrandingCustomCss field value if set, zero value otherwise.
func (o *Brand) GetBrandingCustomCss() string {
	if o == nil || o.BrandingCustomCss == nil {
		var ret string
		return ret
	}
	return *o.BrandingCustomCss
}

// GetBrandingCustomCssOk returns a tuple with the BrandingCustomCss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetBrandingCustomCssOk() (*string, bool) {
	if o == nil || o.BrandingCustomCss == nil {
		return nil, false
	}
	return o.BrandingCustomCss, true
}

// HasBrandingCustomCss returns a boolean if a field has been set.
func (o *Brand) HasBrandingCustomCss() bool {
	if o != nil && o.BrandingCustomCss != nil {
		return true
	}

	return false
}

// SetBrandingCustomCss gets a reference to the given string and assigns it to the BrandingCustomCss field.
func (o *Brand) SetBrandingCustomCss(v string) {
	o.BrandingCustomCss = &v
}

// GetBrandingDefaultFlowBackground returns the BrandingDefaultFlowBackground field value if set, zero value otherwise.
func (o *Brand) GetBrandingDefaultFlowBackground() string {
	if o == nil || o.BrandingDefaultFlowBackground == nil {
		var ret string
		return ret
	}
	return *o.BrandingDefaultFlowBackground
}

// GetBrandingDefaultFlowBackgroundOk returns a tuple with the BrandingDefaultFlowBackground field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetBrandingDefaultFlowBackgroundOk() (*string, bool) {
	if o == nil || o.BrandingDefaultFlowBackground == nil {
		return nil, false
	}
	return o.BrandingDefaultFlowBackground, true
}

// HasBrandingDefaultFlowBackground returns a boolean if a field has been set.
func (o *Brand) HasBrandingDefaultFlowBackground() bool {
	if o != nil && o.BrandingDefaultFlowBackground != nil {
		return true
	}

	return false
}

// SetBrandingDefaultFlowBackground gets a reference to the given string and assigns it to the BrandingDefaultFlowBackground field.
func (o *Brand) SetBrandingDefaultFlowBackground(v string) {
	o.BrandingDefaultFlowBackground = &v
}

// GetFlowAuthentication returns the FlowAuthentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Brand) GetFlowAuthentication() string {
	if o == nil || o.FlowAuthentication.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowAuthentication.Get()
}

// GetFlowAuthenticationOk returns a tuple with the FlowAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Brand) GetFlowAuthenticationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowAuthentication.Get(), o.FlowAuthentication.IsSet()
}

// HasFlowAuthentication returns a boolean if a field has been set.
func (o *Brand) HasFlowAuthentication() bool {
	if o != nil && o.FlowAuthentication.IsSet() {
		return true
	}

	return false
}

// SetFlowAuthentication gets a reference to the given NullableString and assigns it to the FlowAuthentication field.
func (o *Brand) SetFlowAuthentication(v string) {
	o.FlowAuthentication.Set(&v)
}

// SetFlowAuthenticationNil sets the value for FlowAuthentication to be an explicit nil
func (o *Brand) SetFlowAuthenticationNil() {
	o.FlowAuthentication.Set(nil)
}

// UnsetFlowAuthentication ensures that no value is present for FlowAuthentication, not even an explicit nil
func (o *Brand) UnsetFlowAuthentication() {
	o.FlowAuthentication.Unset()
}

// GetFlowInvalidation returns the FlowInvalidation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Brand) GetFlowInvalidation() string {
	if o == nil || o.FlowInvalidation.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowInvalidation.Get()
}

// GetFlowInvalidationOk returns a tuple with the FlowInvalidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Brand) GetFlowInvalidationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowInvalidation.Get(), o.FlowInvalidation.IsSet()
}

// HasFlowInvalidation returns a boolean if a field has been set.
func (o *Brand) HasFlowInvalidation() bool {
	if o != nil && o.FlowInvalidation.IsSet() {
		return true
	}

	return false
}

// SetFlowInvalidation gets a reference to the given NullableString and assigns it to the FlowInvalidation field.
func (o *Brand) SetFlowInvalidation(v string) {
	o.FlowInvalidation.Set(&v)
}

// SetFlowInvalidationNil sets the value for FlowInvalidation to be an explicit nil
func (o *Brand) SetFlowInvalidationNil() {
	o.FlowInvalidation.Set(nil)
}

// UnsetFlowInvalidation ensures that no value is present for FlowInvalidation, not even an explicit nil
func (o *Brand) UnsetFlowInvalidation() {
	o.FlowInvalidation.Unset()
}

// GetFlowRecovery returns the FlowRecovery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Brand) GetFlowRecovery() string {
	if o == nil || o.FlowRecovery.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowRecovery.Get()
}

// GetFlowRecoveryOk returns a tuple with the FlowRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Brand) GetFlowRecoveryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowRecovery.Get(), o.FlowRecovery.IsSet()
}

// HasFlowRecovery returns a boolean if a field has been set.
func (o *Brand) HasFlowRecovery() bool {
	if o != nil && o.FlowRecovery.IsSet() {
		return true
	}

	return false
}

// SetFlowRecovery gets a reference to the given NullableString and assigns it to the FlowRecovery field.
func (o *Brand) SetFlowRecovery(v string) {
	o.FlowRecovery.Set(&v)
}

// SetFlowRecoveryNil sets the value for FlowRecovery to be an explicit nil
func (o *Brand) SetFlowRecoveryNil() {
	o.FlowRecovery.Set(nil)
}

// UnsetFlowRecovery ensures that no value is present for FlowRecovery, not even an explicit nil
func (o *Brand) UnsetFlowRecovery() {
	o.FlowRecovery.Unset()
}

// GetFlowUnenrollment returns the FlowUnenrollment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Brand) GetFlowUnenrollment() string {
	if o == nil || o.FlowUnenrollment.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowUnenrollment.Get()
}

// GetFlowUnenrollmentOk returns a tuple with the FlowUnenrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Brand) GetFlowUnenrollmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowUnenrollment.Get(), o.FlowUnenrollment.IsSet()
}

// HasFlowUnenrollment returns a boolean if a field has been set.
func (o *Brand) HasFlowUnenrollment() bool {
	if o != nil && o.FlowUnenrollment.IsSet() {
		return true
	}

	return false
}

// SetFlowUnenrollment gets a reference to the given NullableString and assigns it to the FlowUnenrollment field.
func (o *Brand) SetFlowUnenrollment(v string) {
	o.FlowUnenrollment.Set(&v)
}

// SetFlowUnenrollmentNil sets the value for FlowUnenrollment to be an explicit nil
func (o *Brand) SetFlowUnenrollmentNil() {
	o.FlowUnenrollment.Set(nil)
}

// UnsetFlowUnenrollment ensures that no value is present for FlowUnenrollment, not even an explicit nil
func (o *Brand) UnsetFlowUnenrollment() {
	o.FlowUnenrollment.Unset()
}

// GetFlowUserSettings returns the FlowUserSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Brand) GetFlowUserSettings() string {
	if o == nil || o.FlowUserSettings.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowUserSettings.Get()
}

// GetFlowUserSettingsOk returns a tuple with the FlowUserSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Brand) GetFlowUserSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowUserSettings.Get(), o.FlowUserSettings.IsSet()
}

// HasFlowUserSettings returns a boolean if a field has been set.
func (o *Brand) HasFlowUserSettings() bool {
	if o != nil && o.FlowUserSettings.IsSet() {
		return true
	}

	return false
}

// SetFlowUserSettings gets a reference to the given NullableString and assigns it to the FlowUserSettings field.
func (o *Brand) SetFlowUserSettings(v string) {
	o.FlowUserSettings.Set(&v)
}

// SetFlowUserSettingsNil sets the value for FlowUserSettings to be an explicit nil
func (o *Brand) SetFlowUserSettingsNil() {
	o.FlowUserSettings.Set(nil)
}

// UnsetFlowUserSettings ensures that no value is present for FlowUserSettings, not even an explicit nil
func (o *Brand) UnsetFlowUserSettings() {
	o.FlowUserSettings.Unset()
}

// GetFlowDeviceCode returns the FlowDeviceCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Brand) GetFlowDeviceCode() string {
	if o == nil || o.FlowDeviceCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowDeviceCode.Get()
}

// GetFlowDeviceCodeOk returns a tuple with the FlowDeviceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Brand) GetFlowDeviceCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowDeviceCode.Get(), o.FlowDeviceCode.IsSet()
}

// HasFlowDeviceCode returns a boolean if a field has been set.
func (o *Brand) HasFlowDeviceCode() bool {
	if o != nil && o.FlowDeviceCode.IsSet() {
		return true
	}

	return false
}

// SetFlowDeviceCode gets a reference to the given NullableString and assigns it to the FlowDeviceCode field.
func (o *Brand) SetFlowDeviceCode(v string) {
	o.FlowDeviceCode.Set(&v)
}

// SetFlowDeviceCodeNil sets the value for FlowDeviceCode to be an explicit nil
func (o *Brand) SetFlowDeviceCodeNil() {
	o.FlowDeviceCode.Set(nil)
}

// UnsetFlowDeviceCode ensures that no value is present for FlowDeviceCode, not even an explicit nil
func (o *Brand) UnsetFlowDeviceCode() {
	o.FlowDeviceCode.Unset()
}

// GetDefaultApplication returns the DefaultApplication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Brand) GetDefaultApplication() string {
	if o == nil || o.DefaultApplication.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultApplication.Get()
}

// GetDefaultApplicationOk returns a tuple with the DefaultApplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Brand) GetDefaultApplicationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultApplication.Get(), o.DefaultApplication.IsSet()
}

// HasDefaultApplication returns a boolean if a field has been set.
func (o *Brand) HasDefaultApplication() bool {
	if o != nil && o.DefaultApplication.IsSet() {
		return true
	}

	return false
}

// SetDefaultApplication gets a reference to the given NullableString and assigns it to the DefaultApplication field.
func (o *Brand) SetDefaultApplication(v string) {
	o.DefaultApplication.Set(&v)
}

// SetDefaultApplicationNil sets the value for DefaultApplication to be an explicit nil
func (o *Brand) SetDefaultApplicationNil() {
	o.DefaultApplication.Set(nil)
}

// UnsetDefaultApplication ensures that no value is present for DefaultApplication, not even an explicit nil
func (o *Brand) UnsetDefaultApplication() {
	o.DefaultApplication.Unset()
}

// GetWebCertificate returns the WebCertificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Brand) GetWebCertificate() string {
	if o == nil || o.WebCertificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebCertificate.Get()
}

// GetWebCertificateOk returns a tuple with the WebCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Brand) GetWebCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebCertificate.Get(), o.WebCertificate.IsSet()
}

// HasWebCertificate returns a boolean if a field has been set.
func (o *Brand) HasWebCertificate() bool {
	if o != nil && o.WebCertificate.IsSet() {
		return true
	}

	return false
}

// SetWebCertificate gets a reference to the given NullableString and assigns it to the WebCertificate field.
func (o *Brand) SetWebCertificate(v string) {
	o.WebCertificate.Set(&v)
}

// SetWebCertificateNil sets the value for WebCertificate to be an explicit nil
func (o *Brand) SetWebCertificateNil() {
	o.WebCertificate.Set(nil)
}

// UnsetWebCertificate ensures that no value is present for WebCertificate, not even an explicit nil
func (o *Brand) UnsetWebCertificate() {
	o.WebCertificate.Unset()
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Brand) GetAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Brand) GetAttributesOk() (*interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Brand) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *Brand) SetAttributes(v interface{}) {
	o.Attributes = v
}

func (o Brand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["brand_uuid"] = o.BrandUuid
	}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.BrandingTitle != nil {
		toSerialize["branding_title"] = o.BrandingTitle
	}
	if o.BrandingLogo != nil {
		toSerialize["branding_logo"] = o.BrandingLogo
	}
	if o.BrandingFavicon != nil {
		toSerialize["branding_favicon"] = o.BrandingFavicon
	}
	if o.BrandingCustomCss != nil {
		toSerialize["branding_custom_css"] = o.BrandingCustomCss
	}
	if o.BrandingDefaultFlowBackground != nil {
		toSerialize["branding_default_flow_background"] = o.BrandingDefaultFlowBackground
	}
	if o.FlowAuthentication.IsSet() {
		toSerialize["flow_authentication"] = o.FlowAuthentication.Get()
	}
	if o.FlowInvalidation.IsSet() {
		toSerialize["flow_invalidation"] = o.FlowInvalidation.Get()
	}
	if o.FlowRecovery.IsSet() {
		toSerialize["flow_recovery"] = o.FlowRecovery.Get()
	}
	if o.FlowUnenrollment.IsSet() {
		toSerialize["flow_unenrollment"] = o.FlowUnenrollment.Get()
	}
	if o.FlowUserSettings.IsSet() {
		toSerialize["flow_user_settings"] = o.FlowUserSettings.Get()
	}
	if o.FlowDeviceCode.IsSet() {
		toSerialize["flow_device_code"] = o.FlowDeviceCode.Get()
	}
	if o.DefaultApplication.IsSet() {
		toSerialize["default_application"] = o.DefaultApplication.Get()
	}
	if o.WebCertificate.IsSet() {
		toSerialize["web_certificate"] = o.WebCertificate.Get()
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableBrand struct {
	value *Brand
	isSet bool
}

func (v NullableBrand) Get() *Brand {
	return v.value
}

func (v *NullableBrand) Set(val *Brand) {
	v.value = val
	v.isSet = true
}

func (v NullableBrand) IsSet() bool {
	return v.isSet
}

func (v *NullableBrand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrand(val *Brand) *NullableBrand {
	return &NullableBrand{value: val, isSet: true}
}

func (v NullableBrand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
