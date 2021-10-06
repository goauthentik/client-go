/*
authentik

Making authentication simple.

API version: 2021.9.7
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedTenantRequest Tenant Serializer
type PatchedTenantRequest struct {
	// Domain that activates this tenant. Can be a superset, i.e. `a.b` for `aa.b` and `ba.b`
	Domain             *string        `json:"domain,omitempty"`
	Default            *bool          `json:"default,omitempty"`
	BrandingTitle      *string        `json:"branding_title,omitempty"`
	BrandingLogo       *string        `json:"branding_logo,omitempty"`
	BrandingFavicon    *string        `json:"branding_favicon,omitempty"`
	FlowAuthentication NullableString `json:"flow_authentication,omitempty"`
	FlowInvalidation   NullableString `json:"flow_invalidation,omitempty"`
	FlowRecovery       NullableString `json:"flow_recovery,omitempty"`
	FlowUnenrollment   NullableString `json:"flow_unenrollment,omitempty"`
	// Events will be deleted after this duration.(Format: weeks=3;days=2;hours=3,seconds=2).
	EventRetention *string `json:"event_retention,omitempty"`
}

// NewPatchedTenantRequest instantiates a new PatchedTenantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedTenantRequest() *PatchedTenantRequest {
	this := PatchedTenantRequest{}
	return &this
}

// NewPatchedTenantRequestWithDefaults instantiates a new PatchedTenantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedTenantRequestWithDefaults() *PatchedTenantRequest {
	this := PatchedTenantRequest{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *PatchedTenantRequest) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTenantRequest) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *PatchedTenantRequest) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *PatchedTenantRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *PatchedTenantRequest) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTenantRequest) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *PatchedTenantRequest) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *PatchedTenantRequest) SetDefault(v bool) {
	o.Default = &v
}

// GetBrandingTitle returns the BrandingTitle field value if set, zero value otherwise.
func (o *PatchedTenantRequest) GetBrandingTitle() string {
	if o == nil || o.BrandingTitle == nil {
		var ret string
		return ret
	}
	return *o.BrandingTitle
}

// GetBrandingTitleOk returns a tuple with the BrandingTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTenantRequest) GetBrandingTitleOk() (*string, bool) {
	if o == nil || o.BrandingTitle == nil {
		return nil, false
	}
	return o.BrandingTitle, true
}

// HasBrandingTitle returns a boolean if a field has been set.
func (o *PatchedTenantRequest) HasBrandingTitle() bool {
	if o != nil && o.BrandingTitle != nil {
		return true
	}

	return false
}

// SetBrandingTitle gets a reference to the given string and assigns it to the BrandingTitle field.
func (o *PatchedTenantRequest) SetBrandingTitle(v string) {
	o.BrandingTitle = &v
}

// GetBrandingLogo returns the BrandingLogo field value if set, zero value otherwise.
func (o *PatchedTenantRequest) GetBrandingLogo() string {
	if o == nil || o.BrandingLogo == nil {
		var ret string
		return ret
	}
	return *o.BrandingLogo
}

// GetBrandingLogoOk returns a tuple with the BrandingLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTenantRequest) GetBrandingLogoOk() (*string, bool) {
	if o == nil || o.BrandingLogo == nil {
		return nil, false
	}
	return o.BrandingLogo, true
}

// HasBrandingLogo returns a boolean if a field has been set.
func (o *PatchedTenantRequest) HasBrandingLogo() bool {
	if o != nil && o.BrandingLogo != nil {
		return true
	}

	return false
}

// SetBrandingLogo gets a reference to the given string and assigns it to the BrandingLogo field.
func (o *PatchedTenantRequest) SetBrandingLogo(v string) {
	o.BrandingLogo = &v
}

// GetBrandingFavicon returns the BrandingFavicon field value if set, zero value otherwise.
func (o *PatchedTenantRequest) GetBrandingFavicon() string {
	if o == nil || o.BrandingFavicon == nil {
		var ret string
		return ret
	}
	return *o.BrandingFavicon
}

// GetBrandingFaviconOk returns a tuple with the BrandingFavicon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTenantRequest) GetBrandingFaviconOk() (*string, bool) {
	if o == nil || o.BrandingFavicon == nil {
		return nil, false
	}
	return o.BrandingFavicon, true
}

// HasBrandingFavicon returns a boolean if a field has been set.
func (o *PatchedTenantRequest) HasBrandingFavicon() bool {
	if o != nil && o.BrandingFavicon != nil {
		return true
	}

	return false
}

// SetBrandingFavicon gets a reference to the given string and assigns it to the BrandingFavicon field.
func (o *PatchedTenantRequest) SetBrandingFavicon(v string) {
	o.BrandingFavicon = &v
}

// GetFlowAuthentication returns the FlowAuthentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedTenantRequest) GetFlowAuthentication() string {
	if o == nil || o.FlowAuthentication.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowAuthentication.Get()
}

// GetFlowAuthenticationOk returns a tuple with the FlowAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedTenantRequest) GetFlowAuthenticationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowAuthentication.Get(), o.FlowAuthentication.IsSet()
}

// HasFlowAuthentication returns a boolean if a field has been set.
func (o *PatchedTenantRequest) HasFlowAuthentication() bool {
	if o != nil && o.FlowAuthentication.IsSet() {
		return true
	}

	return false
}

// SetFlowAuthentication gets a reference to the given NullableString and assigns it to the FlowAuthentication field.
func (o *PatchedTenantRequest) SetFlowAuthentication(v string) {
	o.FlowAuthentication.Set(&v)
}

// SetFlowAuthenticationNil sets the value for FlowAuthentication to be an explicit nil
func (o *PatchedTenantRequest) SetFlowAuthenticationNil() {
	o.FlowAuthentication.Set(nil)
}

// UnsetFlowAuthentication ensures that no value is present for FlowAuthentication, not even an explicit nil
func (o *PatchedTenantRequest) UnsetFlowAuthentication() {
	o.FlowAuthentication.Unset()
}

// GetFlowInvalidation returns the FlowInvalidation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedTenantRequest) GetFlowInvalidation() string {
	if o == nil || o.FlowInvalidation.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowInvalidation.Get()
}

// GetFlowInvalidationOk returns a tuple with the FlowInvalidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedTenantRequest) GetFlowInvalidationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowInvalidation.Get(), o.FlowInvalidation.IsSet()
}

// HasFlowInvalidation returns a boolean if a field has been set.
func (o *PatchedTenantRequest) HasFlowInvalidation() bool {
	if o != nil && o.FlowInvalidation.IsSet() {
		return true
	}

	return false
}

// SetFlowInvalidation gets a reference to the given NullableString and assigns it to the FlowInvalidation field.
func (o *PatchedTenantRequest) SetFlowInvalidation(v string) {
	o.FlowInvalidation.Set(&v)
}

// SetFlowInvalidationNil sets the value for FlowInvalidation to be an explicit nil
func (o *PatchedTenantRequest) SetFlowInvalidationNil() {
	o.FlowInvalidation.Set(nil)
}

// UnsetFlowInvalidation ensures that no value is present for FlowInvalidation, not even an explicit nil
func (o *PatchedTenantRequest) UnsetFlowInvalidation() {
	o.FlowInvalidation.Unset()
}

// GetFlowRecovery returns the FlowRecovery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedTenantRequest) GetFlowRecovery() string {
	if o == nil || o.FlowRecovery.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowRecovery.Get()
}

// GetFlowRecoveryOk returns a tuple with the FlowRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedTenantRequest) GetFlowRecoveryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowRecovery.Get(), o.FlowRecovery.IsSet()
}

// HasFlowRecovery returns a boolean if a field has been set.
func (o *PatchedTenantRequest) HasFlowRecovery() bool {
	if o != nil && o.FlowRecovery.IsSet() {
		return true
	}

	return false
}

// SetFlowRecovery gets a reference to the given NullableString and assigns it to the FlowRecovery field.
func (o *PatchedTenantRequest) SetFlowRecovery(v string) {
	o.FlowRecovery.Set(&v)
}

// SetFlowRecoveryNil sets the value for FlowRecovery to be an explicit nil
func (o *PatchedTenantRequest) SetFlowRecoveryNil() {
	o.FlowRecovery.Set(nil)
}

// UnsetFlowRecovery ensures that no value is present for FlowRecovery, not even an explicit nil
func (o *PatchedTenantRequest) UnsetFlowRecovery() {
	o.FlowRecovery.Unset()
}

// GetFlowUnenrollment returns the FlowUnenrollment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedTenantRequest) GetFlowUnenrollment() string {
	if o == nil || o.FlowUnenrollment.Get() == nil {
		var ret string
		return ret
	}
	return *o.FlowUnenrollment.Get()
}

// GetFlowUnenrollmentOk returns a tuple with the FlowUnenrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedTenantRequest) GetFlowUnenrollmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowUnenrollment.Get(), o.FlowUnenrollment.IsSet()
}

// HasFlowUnenrollment returns a boolean if a field has been set.
func (o *PatchedTenantRequest) HasFlowUnenrollment() bool {
	if o != nil && o.FlowUnenrollment.IsSet() {
		return true
	}

	return false
}

// SetFlowUnenrollment gets a reference to the given NullableString and assigns it to the FlowUnenrollment field.
func (o *PatchedTenantRequest) SetFlowUnenrollment(v string) {
	o.FlowUnenrollment.Set(&v)
}

// SetFlowUnenrollmentNil sets the value for FlowUnenrollment to be an explicit nil
func (o *PatchedTenantRequest) SetFlowUnenrollmentNil() {
	o.FlowUnenrollment.Set(nil)
}

// UnsetFlowUnenrollment ensures that no value is present for FlowUnenrollment, not even an explicit nil
func (o *PatchedTenantRequest) UnsetFlowUnenrollment() {
	o.FlowUnenrollment.Unset()
}

// GetEventRetention returns the EventRetention field value if set, zero value otherwise.
func (o *PatchedTenantRequest) GetEventRetention() string {
	if o == nil || o.EventRetention == nil {
		var ret string
		return ret
	}
	return *o.EventRetention
}

// GetEventRetentionOk returns a tuple with the EventRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTenantRequest) GetEventRetentionOk() (*string, bool) {
	if o == nil || o.EventRetention == nil {
		return nil, false
	}
	return o.EventRetention, true
}

// HasEventRetention returns a boolean if a field has been set.
func (o *PatchedTenantRequest) HasEventRetention() bool {
	if o != nil && o.EventRetention != nil {
		return true
	}

	return false
}

// SetEventRetention gets a reference to the given string and assigns it to the EventRetention field.
func (o *PatchedTenantRequest) SetEventRetention(v string) {
	o.EventRetention = &v
}

func (o PatchedTenantRequest) MarshalJSON() ([]byte, error) {
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
	if o.EventRetention != nil {
		toSerialize["event_retention"] = o.EventRetention
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedTenantRequest struct {
	value *PatchedTenantRequest
	isSet bool
}

func (v NullablePatchedTenantRequest) Get() *PatchedTenantRequest {
	return v.value
}

func (v *NullablePatchedTenantRequest) Set(val *PatchedTenantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedTenantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedTenantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedTenantRequest(val *PatchedTenantRequest) *NullablePatchedTenantRequest {
	return &NullablePatchedTenantRequest{value: val, isSet: true}
}

func (v NullablePatchedTenantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedTenantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
