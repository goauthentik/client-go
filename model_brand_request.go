/*
authentik

Making authentication simple.

API version: 2023.10.6
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BrandRequest Brand Serializer
type BrandRequest struct {
	// Domain that activates this brand. Can be a superset, i.e. `a.b` for `aa.b` and `ba.b`
	Domain             string         `json:"domain"`
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
	// Web Certificate used by the authentik Core webserver.
	WebCertificate NullableString `json:"web_certificate,omitempty"`
	Attributes     interface{}    `json:"attributes,omitempty"`
}

// NewBrandRequest instantiates a new BrandRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandRequest(domain string) *BrandRequest {
	this := BrandRequest{}
	this.Domain = domain
	return &this
}

// NewBrandRequestWithDefaults instantiates a new BrandRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandRequestWithDefaults() *BrandRequest {
	this := BrandRequest{}
	return &this
}

// GetDomain returns the Domain field value
func (o *BrandRequest) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *BrandRequest) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *BrandRequest) SetDomain(v string) {
	o.Domain = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *BrandRequest) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandRequest) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *BrandRequest) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *BrandRequest) SetDefault(v bool) {
	o.Default = &v
}

// GetBrandingTitle returns the BrandingTitle field value if set, zero value otherwise.
func (o *BrandRequest) GetBrandingTitle() string {
	if o == nil || o.BrandingTitle == nil {
		var ret string
		return ret
	}
	return *o.BrandingTitle
}

// GetBrandingTitleOk returns a tuple with the BrandingTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandRequest) GetBrandingTitleOk() (*string, bool) {
	if o == nil || o.BrandingTitle == nil {
		return nil, false
	}
	return o.BrandingTitle, true
}

// HasBrandingTitle returns a boolean if a field has been set.
func (o *BrandRequest) HasBrandingTitle() bool {
	if o != nil && o.BrandingTitle != nil {
		return true
	}

	return false
}

// SetBrandingTitle gets a reference to the given string and assigns it to the BrandingTitle field.
func (o *BrandRequest) SetBrandingTitle(v string) {
	o.BrandingTitle = &v
}

// GetBrandingLogo returns the BrandingLogo field value if set, zero value otherwise.
func (o *BrandRequest) GetBrandingLogo() string {
	if o == nil || o.BrandingLogo == nil {
		var ret string
		return ret
	}
	return *o.BrandingLogo
}

// GetBrandingLogoOk returns a tuple with the BrandingLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandRequest) GetBrandingLogoOk() (*string, bool) {
	if o == nil || o.BrandingLogo == nil {
		return nil, false
	}
	return o.BrandingLogo, true
}

// HasBrandingLogo returns a boolean if a field has been set.
func (o *BrandRequest) HasBrandingLogo() bool {
	if o != nil && o.BrandingLogo != nil {
		return true
	}

	return false
}

// SetBrandingLogo gets a reference to the given string and assigns it to the BrandingLogo field.
func (o *BrandRequest) SetBrandingLogo(v string) {
	o.BrandingLogo = &v
}

// GetBrandingFavicon returns the BrandingFavicon field value if set, zero value otherwise.
func (o *BrandRequest) GetBrandingFavicon() string {
	if o == nil || o.BrandingFavicon == nil {
		var ret string
		return ret
	}
	return *o.BrandingFavicon
}

// GetBrandingFaviconOk returns a tuple with the BrandingFavicon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandRequest) GetBrandingFaviconOk() (*string, bool) {
	if o == nil || o.BrandingFavicon == nil {
		return nil, false
	}
	return o.BrandingFavicon, true
}

// HasBrandingFavicon returns a boolean if a field has been set.
func (o *BrandRequest) HasBrandingFavicon() bool {
	if o != nil && o.BrandingFavicon != nil {
		return true
	}

	return false
}

// SetBrandingFavicon gets a reference to the given string and assigns it to the BrandingFavicon field.
func (o *BrandRequest) SetBrandingFavicon(v string) {
	o.BrandingFavicon = &v
}

// GetFlowAuthentication returns the FlowAuthentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BrandRequest) GetFlowAuthentication() string {
	if o == nil || o.FlowAuthentication.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowAuthentication.Get()
}

// GetFlowAuthenticationOk returns a tuple with the FlowAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BrandRequest) GetFlowAuthenticationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowAuthentication.Get(), o.FlowAuthentication.IsSet()
}

// HasFlowAuthentication returns a boolean if a field has been set.
func (o *BrandRequest) HasFlowAuthentication() bool {
	if o != nil && o.FlowAuthentication.IsSet() {
		return true
	}

	return false
}

// SetFlowAuthentication gets a reference to the given NullableString and assigns it to the FlowAuthentication field.
func (o *BrandRequest) SetFlowAuthentication(v string) {
	o.FlowAuthentication.Set(&v)
}

// SetFlowAuthenticationNil sets the value for FlowAuthentication to be an explicit nil
func (o *BrandRequest) SetFlowAuthenticationNil() {
	o.FlowAuthentication.Set(nil)
}

// UnsetFlowAuthentication ensures that no value is present for FlowAuthentication, not even an explicit nil
func (o *BrandRequest) UnsetFlowAuthentication() {
	o.FlowAuthentication.Unset()
}

// GetFlowInvalidation returns the FlowInvalidation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BrandRequest) GetFlowInvalidation() string {
	if o == nil || o.FlowInvalidation.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowInvalidation.Get()
}

// GetFlowInvalidationOk returns a tuple with the FlowInvalidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BrandRequest) GetFlowInvalidationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowInvalidation.Get(), o.FlowInvalidation.IsSet()
}

// HasFlowInvalidation returns a boolean if a field has been set.
func (o *BrandRequest) HasFlowInvalidation() bool {
	if o != nil && o.FlowInvalidation.IsSet() {
		return true
	}

	return false
}

// SetFlowInvalidation gets a reference to the given NullableString and assigns it to the FlowInvalidation field.
func (o *BrandRequest) SetFlowInvalidation(v string) {
	o.FlowInvalidation.Set(&v)
}

// SetFlowInvalidationNil sets the value for FlowInvalidation to be an explicit nil
func (o *BrandRequest) SetFlowInvalidationNil() {
	o.FlowInvalidation.Set(nil)
}

// UnsetFlowInvalidation ensures that no value is present for FlowInvalidation, not even an explicit nil
func (o *BrandRequest) UnsetFlowInvalidation() {
	o.FlowInvalidation.Unset()
}

// GetFlowRecovery returns the FlowRecovery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BrandRequest) GetFlowRecovery() string {
	if o == nil || o.FlowRecovery.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowRecovery.Get()
}

// GetFlowRecoveryOk returns a tuple with the FlowRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BrandRequest) GetFlowRecoveryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowRecovery.Get(), o.FlowRecovery.IsSet()
}

// HasFlowRecovery returns a boolean if a field has been set.
func (o *BrandRequest) HasFlowRecovery() bool {
	if o != nil && o.FlowRecovery.IsSet() {
		return true
	}

	return false
}

// SetFlowRecovery gets a reference to the given NullableString and assigns it to the FlowRecovery field.
func (o *BrandRequest) SetFlowRecovery(v string) {
	o.FlowRecovery.Set(&v)
}

// SetFlowRecoveryNil sets the value for FlowRecovery to be an explicit nil
func (o *BrandRequest) SetFlowRecoveryNil() {
	o.FlowRecovery.Set(nil)
}

// UnsetFlowRecovery ensures that no value is present for FlowRecovery, not even an explicit nil
func (o *BrandRequest) UnsetFlowRecovery() {
	o.FlowRecovery.Unset()
}

// GetFlowUnenrollment returns the FlowUnenrollment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BrandRequest) GetFlowUnenrollment() string {
	if o == nil || o.FlowUnenrollment.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowUnenrollment.Get()
}

// GetFlowUnenrollmentOk returns a tuple with the FlowUnenrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BrandRequest) GetFlowUnenrollmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowUnenrollment.Get(), o.FlowUnenrollment.IsSet()
}

// HasFlowUnenrollment returns a boolean if a field has been set.
func (o *BrandRequest) HasFlowUnenrollment() bool {
	if o != nil && o.FlowUnenrollment.IsSet() {
		return true
	}

	return false
}

// SetFlowUnenrollment gets a reference to the given NullableString and assigns it to the FlowUnenrollment field.
func (o *BrandRequest) SetFlowUnenrollment(v string) {
	o.FlowUnenrollment.Set(&v)
}

// SetFlowUnenrollmentNil sets the value for FlowUnenrollment to be an explicit nil
func (o *BrandRequest) SetFlowUnenrollmentNil() {
	o.FlowUnenrollment.Set(nil)
}

// UnsetFlowUnenrollment ensures that no value is present for FlowUnenrollment, not even an explicit nil
func (o *BrandRequest) UnsetFlowUnenrollment() {
	o.FlowUnenrollment.Unset()
}

// GetFlowUserSettings returns the FlowUserSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BrandRequest) GetFlowUserSettings() string {
	if o == nil || o.FlowUserSettings.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowUserSettings.Get()
}

// GetFlowUserSettingsOk returns a tuple with the FlowUserSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BrandRequest) GetFlowUserSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowUserSettings.Get(), o.FlowUserSettings.IsSet()
}

// HasFlowUserSettings returns a boolean if a field has been set.
func (o *BrandRequest) HasFlowUserSettings() bool {
	if o != nil && o.FlowUserSettings.IsSet() {
		return true
	}

	return false
}

// SetFlowUserSettings gets a reference to the given NullableString and assigns it to the FlowUserSettings field.
func (o *BrandRequest) SetFlowUserSettings(v string) {
	o.FlowUserSettings.Set(&v)
}

// SetFlowUserSettingsNil sets the value for FlowUserSettings to be an explicit nil
func (o *BrandRequest) SetFlowUserSettingsNil() {
	o.FlowUserSettings.Set(nil)
}

// UnsetFlowUserSettings ensures that no value is present for FlowUserSettings, not even an explicit nil
func (o *BrandRequest) UnsetFlowUserSettings() {
	o.FlowUserSettings.Unset()
}

// GetFlowDeviceCode returns the FlowDeviceCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BrandRequest) GetFlowDeviceCode() string {
	if o == nil || o.FlowDeviceCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowDeviceCode.Get()
}

// GetFlowDeviceCodeOk returns a tuple with the FlowDeviceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BrandRequest) GetFlowDeviceCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowDeviceCode.Get(), o.FlowDeviceCode.IsSet()
}

// HasFlowDeviceCode returns a boolean if a field has been set.
func (o *BrandRequest) HasFlowDeviceCode() bool {
	if o != nil && o.FlowDeviceCode.IsSet() {
		return true
	}

	return false
}

// SetFlowDeviceCode gets a reference to the given NullableString and assigns it to the FlowDeviceCode field.
func (o *BrandRequest) SetFlowDeviceCode(v string) {
	o.FlowDeviceCode.Set(&v)
}

// SetFlowDeviceCodeNil sets the value for FlowDeviceCode to be an explicit nil
func (o *BrandRequest) SetFlowDeviceCodeNil() {
	o.FlowDeviceCode.Set(nil)
}

// UnsetFlowDeviceCode ensures that no value is present for FlowDeviceCode, not even an explicit nil
func (o *BrandRequest) UnsetFlowDeviceCode() {
	o.FlowDeviceCode.Unset()
}

// GetWebCertificate returns the WebCertificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BrandRequest) GetWebCertificate() string {
	if o == nil || o.WebCertificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebCertificate.Get()
}

// GetWebCertificateOk returns a tuple with the WebCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BrandRequest) GetWebCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebCertificate.Get(), o.WebCertificate.IsSet()
}

// HasWebCertificate returns a boolean if a field has been set.
func (o *BrandRequest) HasWebCertificate() bool {
	if o != nil && o.WebCertificate.IsSet() {
		return true
	}

	return false
}

// SetWebCertificate gets a reference to the given NullableString and assigns it to the WebCertificate field.
func (o *BrandRequest) SetWebCertificate(v string) {
	o.WebCertificate.Set(&v)
}

// SetWebCertificateNil sets the value for WebCertificate to be an explicit nil
func (o *BrandRequest) SetWebCertificateNil() {
	o.WebCertificate.Set(nil)
}

// UnsetWebCertificate ensures that no value is present for WebCertificate, not even an explicit nil
func (o *BrandRequest) UnsetWebCertificate() {
	o.WebCertificate.Unset()
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BrandRequest) GetAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BrandRequest) GetAttributesOk() (*interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BrandRequest) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *BrandRequest) SetAttributes(v interface{}) {
	o.Attributes = v
}

func (o BrandRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.WebCertificate.IsSet() {
		toSerialize["web_certificate"] = o.WebCertificate.Get()
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableBrandRequest struct {
	value *BrandRequest
	isSet bool
}

func (v NullableBrandRequest) Get() *BrandRequest {
	return v.value
}

func (v *NullableBrandRequest) Set(val *BrandRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandRequest(val *BrandRequest) *NullableBrandRequest {
	return &NullableBrandRequest{value: val, isSet: true}
}

func (v NullableBrandRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
