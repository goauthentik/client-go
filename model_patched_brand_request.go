/*
authentik

Making authentication simple.

API version: 2024.12.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedBrandRequest Brand Serializer
type PatchedBrandRequest struct {
	// Domain that activates this brand. Can be a superset, i.e. `a.b` for `aa.b` and `ba.b`
	Domain             *string        `json:"domain,omitempty"`
	Default            *bool          `json:"default,omitempty"`
	BrandingTitle      *string        `json:"branding_title,omitempty"`
	BrandingLogo       *string        `json:"branding_logo,omitempty"`
	BrandingFavicon    *string        `json:"branding_favicon,omitempty"`
	FlowAuthentication NullableString `json:"flow_authentication,omitempty"`
	FlowInvalidation   NullableString `json:"flow_invalidation,omitempty"`
	FlowRecovery       NullableString `json:"flow_recovery,omitempty"`
	FlowUnenrollment   NullableString `json:"flow_unenrollment,omitempty"`
	FlowUserSettings   NullableString `json:"flow_user_settings,omitempty"`
	FlowDeviceCode     NullableString `json:"flow_device_code,omitempty"`
	// When set, external users will be redirected to this application after authenticating.
	DefaultApplication NullableString `json:"default_application,omitempty"`
	// Web Certificate used by the authentik Core webserver.
	WebCertificate NullableString `json:"web_certificate,omitempty"`
	Attributes     interface{}    `json:"attributes,omitempty"`
}

// NewPatchedBrandRequest instantiates a new PatchedBrandRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBrandRequest() *PatchedBrandRequest {
	this := PatchedBrandRequest{}
	return &this
}

// NewPatchedBrandRequestWithDefaults instantiates a new PatchedBrandRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBrandRequestWithDefaults() *PatchedBrandRequest {
	this := PatchedBrandRequest{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *PatchedBrandRequest) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBrandRequest) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *PatchedBrandRequest) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *PatchedBrandRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *PatchedBrandRequest) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBrandRequest) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *PatchedBrandRequest) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *PatchedBrandRequest) SetDefault(v bool) {
	o.Default = &v
}

// GetBrandingTitle returns the BrandingTitle field value if set, zero value otherwise.
func (o *PatchedBrandRequest) GetBrandingTitle() string {
	if o == nil || o.BrandingTitle == nil {
		var ret string
		return ret
	}
	return *o.BrandingTitle
}

// GetBrandingTitleOk returns a tuple with the BrandingTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBrandRequest) GetBrandingTitleOk() (*string, bool) {
	if o == nil || o.BrandingTitle == nil {
		return nil, false
	}
	return o.BrandingTitle, true
}

// HasBrandingTitle returns a boolean if a field has been set.
func (o *PatchedBrandRequest) HasBrandingTitle() bool {
	if o != nil && o.BrandingTitle != nil {
		return true
	}

	return false
}

// SetBrandingTitle gets a reference to the given string and assigns it to the BrandingTitle field.
func (o *PatchedBrandRequest) SetBrandingTitle(v string) {
	o.BrandingTitle = &v
}

// GetBrandingLogo returns the BrandingLogo field value if set, zero value otherwise.
func (o *PatchedBrandRequest) GetBrandingLogo() string {
	if o == nil || o.BrandingLogo == nil {
		var ret string
		return ret
	}
	return *o.BrandingLogo
}

// GetBrandingLogoOk returns a tuple with the BrandingLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBrandRequest) GetBrandingLogoOk() (*string, bool) {
	if o == nil || o.BrandingLogo == nil {
		return nil, false
	}
	return o.BrandingLogo, true
}

// HasBrandingLogo returns a boolean if a field has been set.
func (o *PatchedBrandRequest) HasBrandingLogo() bool {
	if o != nil && o.BrandingLogo != nil {
		return true
	}

	return false
}

// SetBrandingLogo gets a reference to the given string and assigns it to the BrandingLogo field.
func (o *PatchedBrandRequest) SetBrandingLogo(v string) {
	o.BrandingLogo = &v
}

// GetBrandingFavicon returns the BrandingFavicon field value if set, zero value otherwise.
func (o *PatchedBrandRequest) GetBrandingFavicon() string {
	if o == nil || o.BrandingFavicon == nil {
		var ret string
		return ret
	}
	return *o.BrandingFavicon
}

// GetBrandingFaviconOk returns a tuple with the BrandingFavicon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBrandRequest) GetBrandingFaviconOk() (*string, bool) {
	if o == nil || o.BrandingFavicon == nil {
		return nil, false
	}
	return o.BrandingFavicon, true
}

// HasBrandingFavicon returns a boolean if a field has been set.
func (o *PatchedBrandRequest) HasBrandingFavicon() bool {
	if o != nil && o.BrandingFavicon != nil {
		return true
	}

	return false
}

// SetBrandingFavicon gets a reference to the given string and assigns it to the BrandingFavicon field.
func (o *PatchedBrandRequest) SetBrandingFavicon(v string) {
	o.BrandingFavicon = &v
}

// GetFlowAuthentication returns the FlowAuthentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBrandRequest) GetFlowAuthentication() string {
	if o == nil || o.FlowAuthentication.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowAuthentication.Get()
}

// GetFlowAuthenticationOk returns a tuple with the FlowAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBrandRequest) GetFlowAuthenticationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowAuthentication.Get(), o.FlowAuthentication.IsSet()
}

// HasFlowAuthentication returns a boolean if a field has been set.
func (o *PatchedBrandRequest) HasFlowAuthentication() bool {
	if o != nil && o.FlowAuthentication.IsSet() {
		return true
	}

	return false
}

// SetFlowAuthentication gets a reference to the given NullableString and assigns it to the FlowAuthentication field.
func (o *PatchedBrandRequest) SetFlowAuthentication(v string) {
	o.FlowAuthentication.Set(&v)
}

// SetFlowAuthenticationNil sets the value for FlowAuthentication to be an explicit nil
func (o *PatchedBrandRequest) SetFlowAuthenticationNil() {
	o.FlowAuthentication.Set(nil)
}

// UnsetFlowAuthentication ensures that no value is present for FlowAuthentication, not even an explicit nil
func (o *PatchedBrandRequest) UnsetFlowAuthentication() {
	o.FlowAuthentication.Unset()
}

// GetFlowInvalidation returns the FlowInvalidation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBrandRequest) GetFlowInvalidation() string {
	if o == nil || o.FlowInvalidation.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowInvalidation.Get()
}

// GetFlowInvalidationOk returns a tuple with the FlowInvalidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBrandRequest) GetFlowInvalidationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowInvalidation.Get(), o.FlowInvalidation.IsSet()
}

// HasFlowInvalidation returns a boolean if a field has been set.
func (o *PatchedBrandRequest) HasFlowInvalidation() bool {
	if o != nil && o.FlowInvalidation.IsSet() {
		return true
	}

	return false
}

// SetFlowInvalidation gets a reference to the given NullableString and assigns it to the FlowInvalidation field.
func (o *PatchedBrandRequest) SetFlowInvalidation(v string) {
	o.FlowInvalidation.Set(&v)
}

// SetFlowInvalidationNil sets the value for FlowInvalidation to be an explicit nil
func (o *PatchedBrandRequest) SetFlowInvalidationNil() {
	o.FlowInvalidation.Set(nil)
}

// UnsetFlowInvalidation ensures that no value is present for FlowInvalidation, not even an explicit nil
func (o *PatchedBrandRequest) UnsetFlowInvalidation() {
	o.FlowInvalidation.Unset()
}

// GetFlowRecovery returns the FlowRecovery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBrandRequest) GetFlowRecovery() string {
	if o == nil || o.FlowRecovery.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowRecovery.Get()
}

// GetFlowRecoveryOk returns a tuple with the FlowRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBrandRequest) GetFlowRecoveryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowRecovery.Get(), o.FlowRecovery.IsSet()
}

// HasFlowRecovery returns a boolean if a field has been set.
func (o *PatchedBrandRequest) HasFlowRecovery() bool {
	if o != nil && o.FlowRecovery.IsSet() {
		return true
	}

	return false
}

// SetFlowRecovery gets a reference to the given NullableString and assigns it to the FlowRecovery field.
func (o *PatchedBrandRequest) SetFlowRecovery(v string) {
	o.FlowRecovery.Set(&v)
}

// SetFlowRecoveryNil sets the value for FlowRecovery to be an explicit nil
func (o *PatchedBrandRequest) SetFlowRecoveryNil() {
	o.FlowRecovery.Set(nil)
}

// UnsetFlowRecovery ensures that no value is present for FlowRecovery, not even an explicit nil
func (o *PatchedBrandRequest) UnsetFlowRecovery() {
	o.FlowRecovery.Unset()
}

// GetFlowUnenrollment returns the FlowUnenrollment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBrandRequest) GetFlowUnenrollment() string {
	if o == nil || o.FlowUnenrollment.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowUnenrollment.Get()
}

// GetFlowUnenrollmentOk returns a tuple with the FlowUnenrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBrandRequest) GetFlowUnenrollmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowUnenrollment.Get(), o.FlowUnenrollment.IsSet()
}

// HasFlowUnenrollment returns a boolean if a field has been set.
func (o *PatchedBrandRequest) HasFlowUnenrollment() bool {
	if o != nil && o.FlowUnenrollment.IsSet() {
		return true
	}

	return false
}

// SetFlowUnenrollment gets a reference to the given NullableString and assigns it to the FlowUnenrollment field.
func (o *PatchedBrandRequest) SetFlowUnenrollment(v string) {
	o.FlowUnenrollment.Set(&v)
}

// SetFlowUnenrollmentNil sets the value for FlowUnenrollment to be an explicit nil
func (o *PatchedBrandRequest) SetFlowUnenrollmentNil() {
	o.FlowUnenrollment.Set(nil)
}

// UnsetFlowUnenrollment ensures that no value is present for FlowUnenrollment, not even an explicit nil
func (o *PatchedBrandRequest) UnsetFlowUnenrollment() {
	o.FlowUnenrollment.Unset()
}

// GetFlowUserSettings returns the FlowUserSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBrandRequest) GetFlowUserSettings() string {
	if o == nil || o.FlowUserSettings.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowUserSettings.Get()
}

// GetFlowUserSettingsOk returns a tuple with the FlowUserSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBrandRequest) GetFlowUserSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowUserSettings.Get(), o.FlowUserSettings.IsSet()
}

// HasFlowUserSettings returns a boolean if a field has been set.
func (o *PatchedBrandRequest) HasFlowUserSettings() bool {
	if o != nil && o.FlowUserSettings.IsSet() {
		return true
	}

	return false
}

// SetFlowUserSettings gets a reference to the given NullableString and assigns it to the FlowUserSettings field.
func (o *PatchedBrandRequest) SetFlowUserSettings(v string) {
	o.FlowUserSettings.Set(&v)
}

// SetFlowUserSettingsNil sets the value for FlowUserSettings to be an explicit nil
func (o *PatchedBrandRequest) SetFlowUserSettingsNil() {
	o.FlowUserSettings.Set(nil)
}

// UnsetFlowUserSettings ensures that no value is present for FlowUserSettings, not even an explicit nil
func (o *PatchedBrandRequest) UnsetFlowUserSettings() {
	o.FlowUserSettings.Unset()
}

// GetFlowDeviceCode returns the FlowDeviceCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBrandRequest) GetFlowDeviceCode() string {
	if o == nil || o.FlowDeviceCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowDeviceCode.Get()
}

// GetFlowDeviceCodeOk returns a tuple with the FlowDeviceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBrandRequest) GetFlowDeviceCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowDeviceCode.Get(), o.FlowDeviceCode.IsSet()
}

// HasFlowDeviceCode returns a boolean if a field has been set.
func (o *PatchedBrandRequest) HasFlowDeviceCode() bool {
	if o != nil && o.FlowDeviceCode.IsSet() {
		return true
	}

	return false
}

// SetFlowDeviceCode gets a reference to the given NullableString and assigns it to the FlowDeviceCode field.
func (o *PatchedBrandRequest) SetFlowDeviceCode(v string) {
	o.FlowDeviceCode.Set(&v)
}

// SetFlowDeviceCodeNil sets the value for FlowDeviceCode to be an explicit nil
func (o *PatchedBrandRequest) SetFlowDeviceCodeNil() {
	o.FlowDeviceCode.Set(nil)
}

// UnsetFlowDeviceCode ensures that no value is present for FlowDeviceCode, not even an explicit nil
func (o *PatchedBrandRequest) UnsetFlowDeviceCode() {
	o.FlowDeviceCode.Unset()
}

// GetDefaultApplication returns the DefaultApplication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBrandRequest) GetDefaultApplication() string {
	if o == nil || o.DefaultApplication.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultApplication.Get()
}

// GetDefaultApplicationOk returns a tuple with the DefaultApplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBrandRequest) GetDefaultApplicationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultApplication.Get(), o.DefaultApplication.IsSet()
}

// HasDefaultApplication returns a boolean if a field has been set.
func (o *PatchedBrandRequest) HasDefaultApplication() bool {
	if o != nil && o.DefaultApplication.IsSet() {
		return true
	}

	return false
}

// SetDefaultApplication gets a reference to the given NullableString and assigns it to the DefaultApplication field.
func (o *PatchedBrandRequest) SetDefaultApplication(v string) {
	o.DefaultApplication.Set(&v)
}

// SetDefaultApplicationNil sets the value for DefaultApplication to be an explicit nil
func (o *PatchedBrandRequest) SetDefaultApplicationNil() {
	o.DefaultApplication.Set(nil)
}

// UnsetDefaultApplication ensures that no value is present for DefaultApplication, not even an explicit nil
func (o *PatchedBrandRequest) UnsetDefaultApplication() {
	o.DefaultApplication.Unset()
}

// GetWebCertificate returns the WebCertificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBrandRequest) GetWebCertificate() string {
	if o == nil || o.WebCertificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebCertificate.Get()
}

// GetWebCertificateOk returns a tuple with the WebCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBrandRequest) GetWebCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebCertificate.Get(), o.WebCertificate.IsSet()
}

// HasWebCertificate returns a boolean if a field has been set.
func (o *PatchedBrandRequest) HasWebCertificate() bool {
	if o != nil && o.WebCertificate.IsSet() {
		return true
	}

	return false
}

// SetWebCertificate gets a reference to the given NullableString and assigns it to the WebCertificate field.
func (o *PatchedBrandRequest) SetWebCertificate(v string) {
	o.WebCertificate.Set(&v)
}

// SetWebCertificateNil sets the value for WebCertificate to be an explicit nil
func (o *PatchedBrandRequest) SetWebCertificateNil() {
	o.WebCertificate.Set(nil)
}

// UnsetWebCertificate ensures that no value is present for WebCertificate, not even an explicit nil
func (o *PatchedBrandRequest) UnsetWebCertificate() {
	o.WebCertificate.Unset()
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBrandRequest) GetAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBrandRequest) GetAttributesOk() (*interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PatchedBrandRequest) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *PatchedBrandRequest) SetAttributes(v interface{}) {
	o.Attributes = v
}

func (o PatchedBrandRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Domain != nil {
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

type NullablePatchedBrandRequest struct {
	value *PatchedBrandRequest
	isSet bool
}

func (v NullablePatchedBrandRequest) Get() *PatchedBrandRequest {
	return v.value
}

func (v *NullablePatchedBrandRequest) Set(val *PatchedBrandRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBrandRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBrandRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBrandRequest(val *PatchedBrandRequest) *NullablePatchedBrandRequest {
	return &NullablePatchedBrandRequest{value: val, isSet: true}
}

func (v NullablePatchedBrandRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBrandRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
