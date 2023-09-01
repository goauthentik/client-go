/*
authentik

Making authentication simple.

API version: 2023.8.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedPlexSourceRequest Plex Source Serializer
type PatchedPlexSourceRequest struct {
	// Source's display Name.
	Name *string `json:"name,omitempty"`
	// Internal source name, used in URLs.
	Slug    *string `json:"slug,omitempty"`
	Enabled *bool   `json:"enabled,omitempty"`
	// Flow to use when authenticating existing users.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow to use when enrolling new users.
	EnrollmentFlow   NullableString    `json:"enrollment_flow,omitempty"`
	PolicyEngineMode *PolicyEngineMode `json:"policy_engine_mode,omitempty"`
	// How the source determines if an existing user should be authenticated or a new user enrolled.  * `identifier` - Use the source-specific identifier * `email_link` - Link to a user with identical email address. Can have security implications when a source doesn't validate email addresses. * `email_deny` - Use the user's email address, but deny enrollment when the email address already exists. * `username_link` - Link to a user with identical username. Can have security implications when a username is used with another source. * `username_deny` - Use the user's username, but deny enrollment when the username already exists.
	UserMatchingMode *UserMatchingModeEnum `json:"user_matching_mode,omitempty"`
	UserPathTemplate *string               `json:"user_path_template,omitempty"`
	// Client identifier used to talk to Plex.
	ClientId *string `json:"client_id,omitempty"`
	// Which servers a user has to be a member of to be granted access. Empty list allows every server.
	AllowedServers []string `json:"allowed_servers,omitempty"`
	// Allow friends to authenticate, even if you don't share a server.
	AllowFriends *bool `json:"allow_friends,omitempty"`
	// Plex token used to check friends
	PlexToken *string `json:"plex_token,omitempty"`
}

// NewPatchedPlexSourceRequest instantiates a new PatchedPlexSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedPlexSourceRequest() *PatchedPlexSourceRequest {
	this := PatchedPlexSourceRequest{}
	return &this
}

// NewPatchedPlexSourceRequestWithDefaults instantiates a new PatchedPlexSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedPlexSourceRequestWithDefaults() *PatchedPlexSourceRequest {
	this := PatchedPlexSourceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedPlexSourceRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlexSourceRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedPlexSourceRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedPlexSourceRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PatchedPlexSourceRequest) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlexSourceRequest) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PatchedPlexSourceRequest) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PatchedPlexSourceRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchedPlexSourceRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlexSourceRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchedPlexSourceRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchedPlexSourceRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedPlexSourceRequest) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedPlexSourceRequest) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *PatchedPlexSourceRequest) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *PatchedPlexSourceRequest) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *PatchedPlexSourceRequest) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *PatchedPlexSourceRequest) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetEnrollmentFlow returns the EnrollmentFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedPlexSourceRequest) GetEnrollmentFlow() string {
	if o == nil || o.EnrollmentFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentFlow.Get()
}

// GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedPlexSourceRequest) GetEnrollmentFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentFlow.Get(), o.EnrollmentFlow.IsSet()
}

// HasEnrollmentFlow returns a boolean if a field has been set.
func (o *PatchedPlexSourceRequest) HasEnrollmentFlow() bool {
	if o != nil && o.EnrollmentFlow.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentFlow gets a reference to the given NullableString and assigns it to the EnrollmentFlow field.
func (o *PatchedPlexSourceRequest) SetEnrollmentFlow(v string) {
	o.EnrollmentFlow.Set(&v)
}

// SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil
func (o *PatchedPlexSourceRequest) SetEnrollmentFlowNil() {
	o.EnrollmentFlow.Set(nil)
}

// UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
func (o *PatchedPlexSourceRequest) UnsetEnrollmentFlow() {
	o.EnrollmentFlow.Unset()
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *PatchedPlexSourceRequest) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlexSourceRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *PatchedPlexSourceRequest) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *PatchedPlexSourceRequest) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetUserMatchingMode returns the UserMatchingMode field value if set, zero value otherwise.
func (o *PatchedPlexSourceRequest) GetUserMatchingMode() UserMatchingModeEnum {
	if o == nil || o.UserMatchingMode == nil {
		var ret UserMatchingModeEnum
		return ret
	}
	return *o.UserMatchingMode
}

// GetUserMatchingModeOk returns a tuple with the UserMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlexSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool) {
	if o == nil || o.UserMatchingMode == nil {
		return nil, false
	}
	return o.UserMatchingMode, true
}

// HasUserMatchingMode returns a boolean if a field has been set.
func (o *PatchedPlexSourceRequest) HasUserMatchingMode() bool {
	if o != nil && o.UserMatchingMode != nil {
		return true
	}

	return false
}

// SetUserMatchingMode gets a reference to the given UserMatchingModeEnum and assigns it to the UserMatchingMode field.
func (o *PatchedPlexSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum) {
	o.UserMatchingMode = &v
}

// GetUserPathTemplate returns the UserPathTemplate field value if set, zero value otherwise.
func (o *PatchedPlexSourceRequest) GetUserPathTemplate() string {
	if o == nil || o.UserPathTemplate == nil {
		var ret string
		return ret
	}
	return *o.UserPathTemplate
}

// GetUserPathTemplateOk returns a tuple with the UserPathTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlexSourceRequest) GetUserPathTemplateOk() (*string, bool) {
	if o == nil || o.UserPathTemplate == nil {
		return nil, false
	}
	return o.UserPathTemplate, true
}

// HasUserPathTemplate returns a boolean if a field has been set.
func (o *PatchedPlexSourceRequest) HasUserPathTemplate() bool {
	if o != nil && o.UserPathTemplate != nil {
		return true
	}

	return false
}

// SetUserPathTemplate gets a reference to the given string and assigns it to the UserPathTemplate field.
func (o *PatchedPlexSourceRequest) SetUserPathTemplate(v string) {
	o.UserPathTemplate = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PatchedPlexSourceRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlexSourceRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PatchedPlexSourceRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PatchedPlexSourceRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetAllowedServers returns the AllowedServers field value if set, zero value otherwise.
func (o *PatchedPlexSourceRequest) GetAllowedServers() []string {
	if o == nil || o.AllowedServers == nil {
		var ret []string
		return ret
	}
	return o.AllowedServers
}

// GetAllowedServersOk returns a tuple with the AllowedServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlexSourceRequest) GetAllowedServersOk() ([]string, bool) {
	if o == nil || o.AllowedServers == nil {
		return nil, false
	}
	return o.AllowedServers, true
}

// HasAllowedServers returns a boolean if a field has been set.
func (o *PatchedPlexSourceRequest) HasAllowedServers() bool {
	if o != nil && o.AllowedServers != nil {
		return true
	}

	return false
}

// SetAllowedServers gets a reference to the given []string and assigns it to the AllowedServers field.
func (o *PatchedPlexSourceRequest) SetAllowedServers(v []string) {
	o.AllowedServers = v
}

// GetAllowFriends returns the AllowFriends field value if set, zero value otherwise.
func (o *PatchedPlexSourceRequest) GetAllowFriends() bool {
	if o == nil || o.AllowFriends == nil {
		var ret bool
		return ret
	}
	return *o.AllowFriends
}

// GetAllowFriendsOk returns a tuple with the AllowFriends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlexSourceRequest) GetAllowFriendsOk() (*bool, bool) {
	if o == nil || o.AllowFriends == nil {
		return nil, false
	}
	return o.AllowFriends, true
}

// HasAllowFriends returns a boolean if a field has been set.
func (o *PatchedPlexSourceRequest) HasAllowFriends() bool {
	if o != nil && o.AllowFriends != nil {
		return true
	}

	return false
}

// SetAllowFriends gets a reference to the given bool and assigns it to the AllowFriends field.
func (o *PatchedPlexSourceRequest) SetAllowFriends(v bool) {
	o.AllowFriends = &v
}

// GetPlexToken returns the PlexToken field value if set, zero value otherwise.
func (o *PatchedPlexSourceRequest) GetPlexToken() string {
	if o == nil || o.PlexToken == nil {
		var ret string
		return ret
	}
	return *o.PlexToken
}

// GetPlexTokenOk returns a tuple with the PlexToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPlexSourceRequest) GetPlexTokenOk() (*string, bool) {
	if o == nil || o.PlexToken == nil {
		return nil, false
	}
	return o.PlexToken, true
}

// HasPlexToken returns a boolean if a field has been set.
func (o *PatchedPlexSourceRequest) HasPlexToken() bool {
	if o != nil && o.PlexToken != nil {
		return true
	}

	return false
}

// SetPlexToken gets a reference to the given string and assigns it to the PlexToken field.
func (o *PatchedPlexSourceRequest) SetPlexToken(v string) {
	o.PlexToken = &v
}

func (o PatchedPlexSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.AuthenticationFlow.IsSet() {
		toSerialize["authentication_flow"] = o.AuthenticationFlow.Get()
	}
	if o.EnrollmentFlow.IsSet() {
		toSerialize["enrollment_flow"] = o.EnrollmentFlow.Get()
	}
	if o.PolicyEngineMode != nil {
		toSerialize["policy_engine_mode"] = o.PolicyEngineMode
	}
	if o.UserMatchingMode != nil {
		toSerialize["user_matching_mode"] = o.UserMatchingMode
	}
	if o.UserPathTemplate != nil {
		toSerialize["user_path_template"] = o.UserPathTemplate
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.AllowedServers != nil {
		toSerialize["allowed_servers"] = o.AllowedServers
	}
	if o.AllowFriends != nil {
		toSerialize["allow_friends"] = o.AllowFriends
	}
	if o.PlexToken != nil {
		toSerialize["plex_token"] = o.PlexToken
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedPlexSourceRequest struct {
	value *PatchedPlexSourceRequest
	isSet bool
}

func (v NullablePatchedPlexSourceRequest) Get() *PatchedPlexSourceRequest {
	return v.value
}

func (v *NullablePatchedPlexSourceRequest) Set(val *PatchedPlexSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedPlexSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedPlexSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedPlexSourceRequest(val *PatchedPlexSourceRequest) *NullablePatchedPlexSourceRequest {
	return &NullablePatchedPlexSourceRequest{value: val, isSet: true}
}

func (v NullablePatchedPlexSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedPlexSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
