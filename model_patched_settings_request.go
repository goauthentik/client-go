/*
authentik

Making authentication simple.

API version: 2024.8.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedSettingsRequest Settings Serializer
type PatchedSettingsRequest struct {
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
	// Default token duration
	DefaultTokenDuration *string `json:"default_token_duration,omitempty"`
	// Default token length
	DefaultTokenLength *int32 `json:"default_token_length,omitempty"`
}

// NewPatchedSettingsRequest instantiates a new PatchedSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedSettingsRequest() *PatchedSettingsRequest {
	this := PatchedSettingsRequest{}
	return &this
}

// NewPatchedSettingsRequestWithDefaults instantiates a new PatchedSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedSettingsRequestWithDefaults() *PatchedSettingsRequest {
	this := PatchedSettingsRequest{}
	return &this
}

// GetAvatars returns the Avatars field value if set, zero value otherwise.
func (o *PatchedSettingsRequest) GetAvatars() string {
	if o == nil || o.Avatars == nil {
		var ret string
		return ret
	}
	return *o.Avatars
}

// GetAvatarsOk returns a tuple with the Avatars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSettingsRequest) GetAvatarsOk() (*string, bool) {
	if o == nil || o.Avatars == nil {
		return nil, false
	}
	return o.Avatars, true
}

// HasAvatars returns a boolean if a field has been set.
func (o *PatchedSettingsRequest) HasAvatars() bool {
	if o != nil && o.Avatars != nil {
		return true
	}

	return false
}

// SetAvatars gets a reference to the given string and assigns it to the Avatars field.
func (o *PatchedSettingsRequest) SetAvatars(v string) {
	o.Avatars = &v
}

// GetDefaultUserChangeName returns the DefaultUserChangeName field value if set, zero value otherwise.
func (o *PatchedSettingsRequest) GetDefaultUserChangeName() bool {
	if o == nil || o.DefaultUserChangeName == nil {
		var ret bool
		return ret
	}
	return *o.DefaultUserChangeName
}

// GetDefaultUserChangeNameOk returns a tuple with the DefaultUserChangeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSettingsRequest) GetDefaultUserChangeNameOk() (*bool, bool) {
	if o == nil || o.DefaultUserChangeName == nil {
		return nil, false
	}
	return o.DefaultUserChangeName, true
}

// HasDefaultUserChangeName returns a boolean if a field has been set.
func (o *PatchedSettingsRequest) HasDefaultUserChangeName() bool {
	if o != nil && o.DefaultUserChangeName != nil {
		return true
	}

	return false
}

// SetDefaultUserChangeName gets a reference to the given bool and assigns it to the DefaultUserChangeName field.
func (o *PatchedSettingsRequest) SetDefaultUserChangeName(v bool) {
	o.DefaultUserChangeName = &v
}

// GetDefaultUserChangeEmail returns the DefaultUserChangeEmail field value if set, zero value otherwise.
func (o *PatchedSettingsRequest) GetDefaultUserChangeEmail() bool {
	if o == nil || o.DefaultUserChangeEmail == nil {
		var ret bool
		return ret
	}
	return *o.DefaultUserChangeEmail
}

// GetDefaultUserChangeEmailOk returns a tuple with the DefaultUserChangeEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSettingsRequest) GetDefaultUserChangeEmailOk() (*bool, bool) {
	if o == nil || o.DefaultUserChangeEmail == nil {
		return nil, false
	}
	return o.DefaultUserChangeEmail, true
}

// HasDefaultUserChangeEmail returns a boolean if a field has been set.
func (o *PatchedSettingsRequest) HasDefaultUserChangeEmail() bool {
	if o != nil && o.DefaultUserChangeEmail != nil {
		return true
	}

	return false
}

// SetDefaultUserChangeEmail gets a reference to the given bool and assigns it to the DefaultUserChangeEmail field.
func (o *PatchedSettingsRequest) SetDefaultUserChangeEmail(v bool) {
	o.DefaultUserChangeEmail = &v
}

// GetDefaultUserChangeUsername returns the DefaultUserChangeUsername field value if set, zero value otherwise.
func (o *PatchedSettingsRequest) GetDefaultUserChangeUsername() bool {
	if o == nil || o.DefaultUserChangeUsername == nil {
		var ret bool
		return ret
	}
	return *o.DefaultUserChangeUsername
}

// GetDefaultUserChangeUsernameOk returns a tuple with the DefaultUserChangeUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSettingsRequest) GetDefaultUserChangeUsernameOk() (*bool, bool) {
	if o == nil || o.DefaultUserChangeUsername == nil {
		return nil, false
	}
	return o.DefaultUserChangeUsername, true
}

// HasDefaultUserChangeUsername returns a boolean if a field has been set.
func (o *PatchedSettingsRequest) HasDefaultUserChangeUsername() bool {
	if o != nil && o.DefaultUserChangeUsername != nil {
		return true
	}

	return false
}

// SetDefaultUserChangeUsername gets a reference to the given bool and assigns it to the DefaultUserChangeUsername field.
func (o *PatchedSettingsRequest) SetDefaultUserChangeUsername(v bool) {
	o.DefaultUserChangeUsername = &v
}

// GetEventRetention returns the EventRetention field value if set, zero value otherwise.
func (o *PatchedSettingsRequest) GetEventRetention() string {
	if o == nil || o.EventRetention == nil {
		var ret string
		return ret
	}
	return *o.EventRetention
}

// GetEventRetentionOk returns a tuple with the EventRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSettingsRequest) GetEventRetentionOk() (*string, bool) {
	if o == nil || o.EventRetention == nil {
		return nil, false
	}
	return o.EventRetention, true
}

// HasEventRetention returns a boolean if a field has been set.
func (o *PatchedSettingsRequest) HasEventRetention() bool {
	if o != nil && o.EventRetention != nil {
		return true
	}

	return false
}

// SetEventRetention gets a reference to the given string and assigns it to the EventRetention field.
func (o *PatchedSettingsRequest) SetEventRetention(v string) {
	o.EventRetention = &v
}

// GetFooterLinks returns the FooterLinks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedSettingsRequest) GetFooterLinks() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FooterLinks
}

// GetFooterLinksOk returns a tuple with the FooterLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedSettingsRequest) GetFooterLinksOk() (*interface{}, bool) {
	if o == nil || o.FooterLinks == nil {
		return nil, false
	}
	return &o.FooterLinks, true
}

// HasFooterLinks returns a boolean if a field has been set.
func (o *PatchedSettingsRequest) HasFooterLinks() bool {
	if o != nil && o.FooterLinks != nil {
		return true
	}

	return false
}

// SetFooterLinks gets a reference to the given interface{} and assigns it to the FooterLinks field.
func (o *PatchedSettingsRequest) SetFooterLinks(v interface{}) {
	o.FooterLinks = v
}

// GetGdprCompliance returns the GdprCompliance field value if set, zero value otherwise.
func (o *PatchedSettingsRequest) GetGdprCompliance() bool {
	if o == nil || o.GdprCompliance == nil {
		var ret bool
		return ret
	}
	return *o.GdprCompliance
}

// GetGdprComplianceOk returns a tuple with the GdprCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSettingsRequest) GetGdprComplianceOk() (*bool, bool) {
	if o == nil || o.GdprCompliance == nil {
		return nil, false
	}
	return o.GdprCompliance, true
}

// HasGdprCompliance returns a boolean if a field has been set.
func (o *PatchedSettingsRequest) HasGdprCompliance() bool {
	if o != nil && o.GdprCompliance != nil {
		return true
	}

	return false
}

// SetGdprCompliance gets a reference to the given bool and assigns it to the GdprCompliance field.
func (o *PatchedSettingsRequest) SetGdprCompliance(v bool) {
	o.GdprCompliance = &v
}

// GetImpersonation returns the Impersonation field value if set, zero value otherwise.
func (o *PatchedSettingsRequest) GetImpersonation() bool {
	if o == nil || o.Impersonation == nil {
		var ret bool
		return ret
	}
	return *o.Impersonation
}

// GetImpersonationOk returns a tuple with the Impersonation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSettingsRequest) GetImpersonationOk() (*bool, bool) {
	if o == nil || o.Impersonation == nil {
		return nil, false
	}
	return o.Impersonation, true
}

// HasImpersonation returns a boolean if a field has been set.
func (o *PatchedSettingsRequest) HasImpersonation() bool {
	if o != nil && o.Impersonation != nil {
		return true
	}

	return false
}

// SetImpersonation gets a reference to the given bool and assigns it to the Impersonation field.
func (o *PatchedSettingsRequest) SetImpersonation(v bool) {
	o.Impersonation = &v
}

// GetDefaultTokenDuration returns the DefaultTokenDuration field value if set, zero value otherwise.
func (o *PatchedSettingsRequest) GetDefaultTokenDuration() string {
	if o == nil || o.DefaultTokenDuration == nil {
		var ret string
		return ret
	}
	return *o.DefaultTokenDuration
}

// GetDefaultTokenDurationOk returns a tuple with the DefaultTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSettingsRequest) GetDefaultTokenDurationOk() (*string, bool) {
	if o == nil || o.DefaultTokenDuration == nil {
		return nil, false
	}
	return o.DefaultTokenDuration, true
}

// HasDefaultTokenDuration returns a boolean if a field has been set.
func (o *PatchedSettingsRequest) HasDefaultTokenDuration() bool {
	if o != nil && o.DefaultTokenDuration != nil {
		return true
	}

	return false
}

// SetDefaultTokenDuration gets a reference to the given string and assigns it to the DefaultTokenDuration field.
func (o *PatchedSettingsRequest) SetDefaultTokenDuration(v string) {
	o.DefaultTokenDuration = &v
}

// GetDefaultTokenLength returns the DefaultTokenLength field value if set, zero value otherwise.
func (o *PatchedSettingsRequest) GetDefaultTokenLength() int32 {
	if o == nil || o.DefaultTokenLength == nil {
		var ret int32
		return ret
	}
	return *o.DefaultTokenLength
}

// GetDefaultTokenLengthOk returns a tuple with the DefaultTokenLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSettingsRequest) GetDefaultTokenLengthOk() (*int32, bool) {
	if o == nil || o.DefaultTokenLength == nil {
		return nil, false
	}
	return o.DefaultTokenLength, true
}

// HasDefaultTokenLength returns a boolean if a field has been set.
func (o *PatchedSettingsRequest) HasDefaultTokenLength() bool {
	if o != nil && o.DefaultTokenLength != nil {
		return true
	}

	return false
}

// SetDefaultTokenLength gets a reference to the given int32 and assigns it to the DefaultTokenLength field.
func (o *PatchedSettingsRequest) SetDefaultTokenLength(v int32) {
	o.DefaultTokenLength = &v
}

func (o PatchedSettingsRequest) MarshalJSON() ([]byte, error) {
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
	if o.DefaultTokenDuration != nil {
		toSerialize["default_token_duration"] = o.DefaultTokenDuration
	}
	if o.DefaultTokenLength != nil {
		toSerialize["default_token_length"] = o.DefaultTokenLength
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedSettingsRequest struct {
	value *PatchedSettingsRequest
	isSet bool
}

func (v NullablePatchedSettingsRequest) Get() *PatchedSettingsRequest {
	return v.value
}

func (v *NullablePatchedSettingsRequest) Set(val *PatchedSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedSettingsRequest(val *PatchedSettingsRequest) *NullablePatchedSettingsRequest {
	return &NullablePatchedSettingsRequest{value: val, isSet: true}
}

func (v NullablePatchedSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
