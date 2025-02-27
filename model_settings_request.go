/*
authentik

Making authentication simple.

API version: 2025.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SettingsRequest Settings Serializer
type SettingsRequest struct {
	// Configure how authentik should show avatars for users.
	Avatars *string `json:"avatars,omitempty"`
	// Enable the ability for users to change their name.
	DefaultUserChangeName *bool `json:"default_user_change_name,omitempty"`
	// Enable the ability for users to change their email address.
	DefaultUserChangeEmail *bool `json:"default_user_change_email,omitempty"`
	// Enable the ability for users to change their username.
	DefaultUserChangeUsername *bool `json:"default_user_change_username,omitempty"`
	// Events will be deleted after this duration.(Format: weeks=3;days=2;hours=3,seconds=2).
	EventRetention *string `json:"event_retention,omitempty"`
	// The option configures the footer links on the flow executor pages.
	FooterLinks interface{} `json:"footer_links,omitempty"`
	// When enabled, all the events caused by a user will be deleted upon the user's deletion.
	GdprCompliance *bool `json:"gdpr_compliance,omitempty"`
	// Globally enable/disable impersonation.
	Impersonation *bool `json:"impersonation,omitempty"`
	// Require administrators to provide a reason for impersonating a user.
	ImpersonationRequireReason *bool `json:"impersonation_require_reason,omitempty"`
	// Default token duration
	DefaultTokenDuration *string `json:"default_token_duration,omitempty"`
	// Default token length
	DefaultTokenLength *int32 `json:"default_token_length,omitempty"`
}

// NewSettingsRequest instantiates a new SettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsRequest() *SettingsRequest {
	this := SettingsRequest{}
	return &this
}

// NewSettingsRequestWithDefaults instantiates a new SettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsRequestWithDefaults() *SettingsRequest {
	this := SettingsRequest{}
	return &this
}

// GetAvatars returns the Avatars field value if set, zero value otherwise.
func (o *SettingsRequest) GetAvatars() string {
	if o == nil || o.Avatars == nil {
		var ret string
		return ret
	}
	return *o.Avatars
}

// GetAvatarsOk returns a tuple with the Avatars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRequest) GetAvatarsOk() (*string, bool) {
	if o == nil || o.Avatars == nil {
		return nil, false
	}
	return o.Avatars, true
}

// HasAvatars returns a boolean if a field has been set.
func (o *SettingsRequest) HasAvatars() bool {
	if o != nil && o.Avatars != nil {
		return true
	}

	return false
}

// SetAvatars gets a reference to the given string and assigns it to the Avatars field.
func (o *SettingsRequest) SetAvatars(v string) {
	o.Avatars = &v
}

// GetDefaultUserChangeName returns the DefaultUserChangeName field value if set, zero value otherwise.
func (o *SettingsRequest) GetDefaultUserChangeName() bool {
	if o == nil || o.DefaultUserChangeName == nil {
		var ret bool
		return ret
	}
	return *o.DefaultUserChangeName
}

// GetDefaultUserChangeNameOk returns a tuple with the DefaultUserChangeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRequest) GetDefaultUserChangeNameOk() (*bool, bool) {
	if o == nil || o.DefaultUserChangeName == nil {
		return nil, false
	}
	return o.DefaultUserChangeName, true
}

// HasDefaultUserChangeName returns a boolean if a field has been set.
func (o *SettingsRequest) HasDefaultUserChangeName() bool {
	if o != nil && o.DefaultUserChangeName != nil {
		return true
	}

	return false
}

// SetDefaultUserChangeName gets a reference to the given bool and assigns it to the DefaultUserChangeName field.
func (o *SettingsRequest) SetDefaultUserChangeName(v bool) {
	o.DefaultUserChangeName = &v
}

// GetDefaultUserChangeEmail returns the DefaultUserChangeEmail field value if set, zero value otherwise.
func (o *SettingsRequest) GetDefaultUserChangeEmail() bool {
	if o == nil || o.DefaultUserChangeEmail == nil {
		var ret bool
		return ret
	}
	return *o.DefaultUserChangeEmail
}

// GetDefaultUserChangeEmailOk returns a tuple with the DefaultUserChangeEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRequest) GetDefaultUserChangeEmailOk() (*bool, bool) {
	if o == nil || o.DefaultUserChangeEmail == nil {
		return nil, false
	}
	return o.DefaultUserChangeEmail, true
}

// HasDefaultUserChangeEmail returns a boolean if a field has been set.
func (o *SettingsRequest) HasDefaultUserChangeEmail() bool {
	if o != nil && o.DefaultUserChangeEmail != nil {
		return true
	}

	return false
}

// SetDefaultUserChangeEmail gets a reference to the given bool and assigns it to the DefaultUserChangeEmail field.
func (o *SettingsRequest) SetDefaultUserChangeEmail(v bool) {
	o.DefaultUserChangeEmail = &v
}

// GetDefaultUserChangeUsername returns the DefaultUserChangeUsername field value if set, zero value otherwise.
func (o *SettingsRequest) GetDefaultUserChangeUsername() bool {
	if o == nil || o.DefaultUserChangeUsername == nil {
		var ret bool
		return ret
	}
	return *o.DefaultUserChangeUsername
}

// GetDefaultUserChangeUsernameOk returns a tuple with the DefaultUserChangeUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRequest) GetDefaultUserChangeUsernameOk() (*bool, bool) {
	if o == nil || o.DefaultUserChangeUsername == nil {
		return nil, false
	}
	return o.DefaultUserChangeUsername, true
}

// HasDefaultUserChangeUsername returns a boolean if a field has been set.
func (o *SettingsRequest) HasDefaultUserChangeUsername() bool {
	if o != nil && o.DefaultUserChangeUsername != nil {
		return true
	}

	return false
}

// SetDefaultUserChangeUsername gets a reference to the given bool and assigns it to the DefaultUserChangeUsername field.
func (o *SettingsRequest) SetDefaultUserChangeUsername(v bool) {
	o.DefaultUserChangeUsername = &v
}

// GetEventRetention returns the EventRetention field value if set, zero value otherwise.
func (o *SettingsRequest) GetEventRetention() string {
	if o == nil || o.EventRetention == nil {
		var ret string
		return ret
	}
	return *o.EventRetention
}

// GetEventRetentionOk returns a tuple with the EventRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRequest) GetEventRetentionOk() (*string, bool) {
	if o == nil || o.EventRetention == nil {
		return nil, false
	}
	return o.EventRetention, true
}

// HasEventRetention returns a boolean if a field has been set.
func (o *SettingsRequest) HasEventRetention() bool {
	if o != nil && o.EventRetention != nil {
		return true
	}

	return false
}

// SetEventRetention gets a reference to the given string and assigns it to the EventRetention field.
func (o *SettingsRequest) SetEventRetention(v string) {
	o.EventRetention = &v
}

// GetFooterLinks returns the FooterLinks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingsRequest) GetFooterLinks() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FooterLinks
}

// GetFooterLinksOk returns a tuple with the FooterLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsRequest) GetFooterLinksOk() (*interface{}, bool) {
	if o == nil || o.FooterLinks == nil {
		return nil, false
	}
	return &o.FooterLinks, true
}

// HasFooterLinks returns a boolean if a field has been set.
func (o *SettingsRequest) HasFooterLinks() bool {
	if o != nil && o.FooterLinks != nil {
		return true
	}

	return false
}

// SetFooterLinks gets a reference to the given interface{} and assigns it to the FooterLinks field.
func (o *SettingsRequest) SetFooterLinks(v interface{}) {
	o.FooterLinks = v
}

// GetGdprCompliance returns the GdprCompliance field value if set, zero value otherwise.
func (o *SettingsRequest) GetGdprCompliance() bool {
	if o == nil || o.GdprCompliance == nil {
		var ret bool
		return ret
	}
	return *o.GdprCompliance
}

// GetGdprComplianceOk returns a tuple with the GdprCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRequest) GetGdprComplianceOk() (*bool, bool) {
	if o == nil || o.GdprCompliance == nil {
		return nil, false
	}
	return o.GdprCompliance, true
}

// HasGdprCompliance returns a boolean if a field has been set.
func (o *SettingsRequest) HasGdprCompliance() bool {
	if o != nil && o.GdprCompliance != nil {
		return true
	}

	return false
}

// SetGdprCompliance gets a reference to the given bool and assigns it to the GdprCompliance field.
func (o *SettingsRequest) SetGdprCompliance(v bool) {
	o.GdprCompliance = &v
}

// GetImpersonation returns the Impersonation field value if set, zero value otherwise.
func (o *SettingsRequest) GetImpersonation() bool {
	if o == nil || o.Impersonation == nil {
		var ret bool
		return ret
	}
	return *o.Impersonation
}

// GetImpersonationOk returns a tuple with the Impersonation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRequest) GetImpersonationOk() (*bool, bool) {
	if o == nil || o.Impersonation == nil {
		return nil, false
	}
	return o.Impersonation, true
}

// HasImpersonation returns a boolean if a field has been set.
func (o *SettingsRequest) HasImpersonation() bool {
	if o != nil && o.Impersonation != nil {
		return true
	}

	return false
}

// SetImpersonation gets a reference to the given bool and assigns it to the Impersonation field.
func (o *SettingsRequest) SetImpersonation(v bool) {
	o.Impersonation = &v
}

// GetImpersonationRequireReason returns the ImpersonationRequireReason field value if set, zero value otherwise.
func (o *SettingsRequest) GetImpersonationRequireReason() bool {
	if o == nil || o.ImpersonationRequireReason == nil {
		var ret bool
		return ret
	}
	return *o.ImpersonationRequireReason
}

// GetImpersonationRequireReasonOk returns a tuple with the ImpersonationRequireReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRequest) GetImpersonationRequireReasonOk() (*bool, bool) {
	if o == nil || o.ImpersonationRequireReason == nil {
		return nil, false
	}
	return o.ImpersonationRequireReason, true
}

// HasImpersonationRequireReason returns a boolean if a field has been set.
func (o *SettingsRequest) HasImpersonationRequireReason() bool {
	if o != nil && o.ImpersonationRequireReason != nil {
		return true
	}

	return false
}

// SetImpersonationRequireReason gets a reference to the given bool and assigns it to the ImpersonationRequireReason field.
func (o *SettingsRequest) SetImpersonationRequireReason(v bool) {
	o.ImpersonationRequireReason = &v
}

// GetDefaultTokenDuration returns the DefaultTokenDuration field value if set, zero value otherwise.
func (o *SettingsRequest) GetDefaultTokenDuration() string {
	if o == nil || o.DefaultTokenDuration == nil {
		var ret string
		return ret
	}
	return *o.DefaultTokenDuration
}

// GetDefaultTokenDurationOk returns a tuple with the DefaultTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRequest) GetDefaultTokenDurationOk() (*string, bool) {
	if o == nil || o.DefaultTokenDuration == nil {
		return nil, false
	}
	return o.DefaultTokenDuration, true
}

// HasDefaultTokenDuration returns a boolean if a field has been set.
func (o *SettingsRequest) HasDefaultTokenDuration() bool {
	if o != nil && o.DefaultTokenDuration != nil {
		return true
	}

	return false
}

// SetDefaultTokenDuration gets a reference to the given string and assigns it to the DefaultTokenDuration field.
func (o *SettingsRequest) SetDefaultTokenDuration(v string) {
	o.DefaultTokenDuration = &v
}

// GetDefaultTokenLength returns the DefaultTokenLength field value if set, zero value otherwise.
func (o *SettingsRequest) GetDefaultTokenLength() int32 {
	if o == nil || o.DefaultTokenLength == nil {
		var ret int32
		return ret
	}
	return *o.DefaultTokenLength
}

// GetDefaultTokenLengthOk returns a tuple with the DefaultTokenLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRequest) GetDefaultTokenLengthOk() (*int32, bool) {
	if o == nil || o.DefaultTokenLength == nil {
		return nil, false
	}
	return o.DefaultTokenLength, true
}

// HasDefaultTokenLength returns a boolean if a field has been set.
func (o *SettingsRequest) HasDefaultTokenLength() bool {
	if o != nil && o.DefaultTokenLength != nil {
		return true
	}

	return false
}

// SetDefaultTokenLength gets a reference to the given int32 and assigns it to the DefaultTokenLength field.
func (o *SettingsRequest) SetDefaultTokenLength(v int32) {
	o.DefaultTokenLength = &v
}

func (o SettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Avatars != nil {
		toSerialize["avatars"] = o.Avatars
	}
	if o.DefaultUserChangeName != nil {
		toSerialize["default_user_change_name"] = o.DefaultUserChangeName
	}
	if o.DefaultUserChangeEmail != nil {
		toSerialize["default_user_change_email"] = o.DefaultUserChangeEmail
	}
	if o.DefaultUserChangeUsername != nil {
		toSerialize["default_user_change_username"] = o.DefaultUserChangeUsername
	}
	if o.EventRetention != nil {
		toSerialize["event_retention"] = o.EventRetention
	}
	if o.FooterLinks != nil {
		toSerialize["footer_links"] = o.FooterLinks
	}
	if o.GdprCompliance != nil {
		toSerialize["gdpr_compliance"] = o.GdprCompliance
	}
	if o.Impersonation != nil {
		toSerialize["impersonation"] = o.Impersonation
	}
	if o.ImpersonationRequireReason != nil {
		toSerialize["impersonation_require_reason"] = o.ImpersonationRequireReason
	}
	if o.DefaultTokenDuration != nil {
		toSerialize["default_token_duration"] = o.DefaultTokenDuration
	}
	if o.DefaultTokenLength != nil {
		toSerialize["default_token_length"] = o.DefaultTokenLength
	}
	return json.Marshal(toSerialize)
}

type NullableSettingsRequest struct {
	value *SettingsRequest
	isSet bool
}

func (v NullableSettingsRequest) Get() *SettingsRequest {
	return v.value
}

func (v *NullableSettingsRequest) Set(val *SettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsRequest(val *SettingsRequest) *NullableSettingsRequest {
	return &NullableSettingsRequest{value: val, isSet: true}
}

func (v NullableSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
