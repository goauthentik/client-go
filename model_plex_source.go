/*
authentik

Making authentication simple.

API version: 2025.2.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PlexSource Plex Source Serializer
type PlexSource struct {
	Pk string `json:"pk"`
	// Source's display Name.
	Name string `json:"name"`
	// Internal source name, used in URLs.
	Slug    string `json:"slug"`
	Enabled *bool  `json:"enabled,omitempty"`
	// Flow to use when authenticating existing users.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow to use when enrolling new users.
	EnrollmentFlow        NullableString `json:"enrollment_flow,omitempty"`
	UserPropertyMappings  []string       `json:"user_property_mappings,omitempty"`
	GroupPropertyMappings []string       `json:"group_property_mappings,omitempty"`
	// Get object component so that we know how to edit the object
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName    string            `json:"meta_model_name"`
	PolicyEngineMode *PolicyEngineMode `json:"policy_engine_mode,omitempty"`
	// How the source determines if an existing user should be authenticated or a new user enrolled.
	UserMatchingMode *UserMatchingModeEnum `json:"user_matching_mode,omitempty"`
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed          NullableString `json:"managed"`
	UserPathTemplate *string        `json:"user_path_template,omitempty"`
	Icon             string         `json:"icon"`
	// How the source determines if an existing group should be used or a new group created.
	GroupMatchingMode *GroupMatchingModeEnum `json:"group_matching_mode,omitempty"`
	// Client identifier used to talk to Plex.
	ClientId *string `json:"client_id,omitempty"`
	// Which servers a user has to be a member of to be granted access. Empty list allows every server.
	AllowedServers []string `json:"allowed_servers,omitempty"`
	// Allow friends to authenticate, even if you don't share a server.
	AllowFriends *bool `json:"allow_friends,omitempty"`
	// Plex token used to check friends
	PlexToken string `json:"plex_token"`
}

// NewPlexSource instantiates a new PlexSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlexSource(pk string, name string, slug string, component string, verboseName string, verboseNamePlural string, metaModelName string, managed NullableString, icon string, plexToken string) *PlexSource {
	this := PlexSource{}
	this.Pk = pk
	this.Name = name
	this.Slug = slug
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.Managed = managed
	this.Icon = icon
	this.PlexToken = plexToken
	return &this
}

// NewPlexSourceWithDefaults instantiates a new PlexSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlexSourceWithDefaults() *PlexSource {
	this := PlexSource{}
	return &this
}

// GetPk returns the Pk field value
func (o *PlexSource) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *PlexSource) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *PlexSource) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *PlexSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlexSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PlexSource) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *PlexSource) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *PlexSource) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *PlexSource) SetSlug(v string) {
	o.Slug = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PlexSource) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexSource) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PlexSource) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PlexSource) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlexSource) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlexSource) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *PlexSource) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *PlexSource) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *PlexSource) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *PlexSource) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetEnrollmentFlow returns the EnrollmentFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlexSource) GetEnrollmentFlow() string {
	if o == nil || o.EnrollmentFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentFlow.Get()
}

// GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlexSource) GetEnrollmentFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentFlow.Get(), o.EnrollmentFlow.IsSet()
}

// HasEnrollmentFlow returns a boolean if a field has been set.
func (o *PlexSource) HasEnrollmentFlow() bool {
	if o != nil && o.EnrollmentFlow.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentFlow gets a reference to the given NullableString and assigns it to the EnrollmentFlow field.
func (o *PlexSource) SetEnrollmentFlow(v string) {
	o.EnrollmentFlow.Set(&v)
}

// SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil
func (o *PlexSource) SetEnrollmentFlowNil() {
	o.EnrollmentFlow.Set(nil)
}

// UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
func (o *PlexSource) UnsetEnrollmentFlow() {
	o.EnrollmentFlow.Unset()
}

// GetUserPropertyMappings returns the UserPropertyMappings field value if set, zero value otherwise.
func (o *PlexSource) GetUserPropertyMappings() []string {
	if o == nil || o.UserPropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.UserPropertyMappings
}

// GetUserPropertyMappingsOk returns a tuple with the UserPropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexSource) GetUserPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.UserPropertyMappings == nil {
		return nil, false
	}
	return o.UserPropertyMappings, true
}

// HasUserPropertyMappings returns a boolean if a field has been set.
func (o *PlexSource) HasUserPropertyMappings() bool {
	if o != nil && o.UserPropertyMappings != nil {
		return true
	}

	return false
}

// SetUserPropertyMappings gets a reference to the given []string and assigns it to the UserPropertyMappings field.
func (o *PlexSource) SetUserPropertyMappings(v []string) {
	o.UserPropertyMappings = v
}

// GetGroupPropertyMappings returns the GroupPropertyMappings field value if set, zero value otherwise.
func (o *PlexSource) GetGroupPropertyMappings() []string {
	if o == nil || o.GroupPropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.GroupPropertyMappings
}

// GetGroupPropertyMappingsOk returns a tuple with the GroupPropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexSource) GetGroupPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.GroupPropertyMappings == nil {
		return nil, false
	}
	return o.GroupPropertyMappings, true
}

// HasGroupPropertyMappings returns a boolean if a field has been set.
func (o *PlexSource) HasGroupPropertyMappings() bool {
	if o != nil && o.GroupPropertyMappings != nil {
		return true
	}

	return false
}

// SetGroupPropertyMappings gets a reference to the given []string and assigns it to the GroupPropertyMappings field.
func (o *PlexSource) SetGroupPropertyMappings(v []string) {
	o.GroupPropertyMappings = v
}

// GetComponent returns the Component field value
func (o *PlexSource) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *PlexSource) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *PlexSource) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *PlexSource) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *PlexSource) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *PlexSource) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *PlexSource) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *PlexSource) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *PlexSource) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *PlexSource) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *PlexSource) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *PlexSource) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *PlexSource) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexSource) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *PlexSource) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *PlexSource) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetUserMatchingMode returns the UserMatchingMode field value if set, zero value otherwise.
func (o *PlexSource) GetUserMatchingMode() UserMatchingModeEnum {
	if o == nil || o.UserMatchingMode == nil {
		var ret UserMatchingModeEnum
		return ret
	}
	return *o.UserMatchingMode
}

// GetUserMatchingModeOk returns a tuple with the UserMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexSource) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool) {
	if o == nil || o.UserMatchingMode == nil {
		return nil, false
	}
	return o.UserMatchingMode, true
}

// HasUserMatchingMode returns a boolean if a field has been set.
func (o *PlexSource) HasUserMatchingMode() bool {
	if o != nil && o.UserMatchingMode != nil {
		return true
	}

	return false
}

// SetUserMatchingMode gets a reference to the given UserMatchingModeEnum and assigns it to the UserMatchingMode field.
func (o *PlexSource) SetUserMatchingMode(v UserMatchingModeEnum) {
	o.UserMatchingMode = &v
}

// GetManaged returns the Managed field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PlexSource) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}

	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlexSource) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// SetManaged sets field value
func (o *PlexSource) SetManaged(v string) {
	o.Managed.Set(&v)
}

// GetUserPathTemplate returns the UserPathTemplate field value if set, zero value otherwise.
func (o *PlexSource) GetUserPathTemplate() string {
	if o == nil || o.UserPathTemplate == nil {
		var ret string
		return ret
	}
	return *o.UserPathTemplate
}

// GetUserPathTemplateOk returns a tuple with the UserPathTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexSource) GetUserPathTemplateOk() (*string, bool) {
	if o == nil || o.UserPathTemplate == nil {
		return nil, false
	}
	return o.UserPathTemplate, true
}

// HasUserPathTemplate returns a boolean if a field has been set.
func (o *PlexSource) HasUserPathTemplate() bool {
	if o != nil && o.UserPathTemplate != nil {
		return true
	}

	return false
}

// SetUserPathTemplate gets a reference to the given string and assigns it to the UserPathTemplate field.
func (o *PlexSource) SetUserPathTemplate(v string) {
	o.UserPathTemplate = &v
}

// GetIcon returns the Icon field value
func (o *PlexSource) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *PlexSource) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *PlexSource) SetIcon(v string) {
	o.Icon = v
}

// GetGroupMatchingMode returns the GroupMatchingMode field value if set, zero value otherwise.
func (o *PlexSource) GetGroupMatchingMode() GroupMatchingModeEnum {
	if o == nil || o.GroupMatchingMode == nil {
		var ret GroupMatchingModeEnum
		return ret
	}
	return *o.GroupMatchingMode
}

// GetGroupMatchingModeOk returns a tuple with the GroupMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexSource) GetGroupMatchingModeOk() (*GroupMatchingModeEnum, bool) {
	if o == nil || o.GroupMatchingMode == nil {
		return nil, false
	}
	return o.GroupMatchingMode, true
}

// HasGroupMatchingMode returns a boolean if a field has been set.
func (o *PlexSource) HasGroupMatchingMode() bool {
	if o != nil && o.GroupMatchingMode != nil {
		return true
	}

	return false
}

// SetGroupMatchingMode gets a reference to the given GroupMatchingModeEnum and assigns it to the GroupMatchingMode field.
func (o *PlexSource) SetGroupMatchingMode(v GroupMatchingModeEnum) {
	o.GroupMatchingMode = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PlexSource) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexSource) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PlexSource) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PlexSource) SetClientId(v string) {
	o.ClientId = &v
}

// GetAllowedServers returns the AllowedServers field value if set, zero value otherwise.
func (o *PlexSource) GetAllowedServers() []string {
	if o == nil || o.AllowedServers == nil {
		var ret []string
		return ret
	}
	return o.AllowedServers
}

// GetAllowedServersOk returns a tuple with the AllowedServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexSource) GetAllowedServersOk() ([]string, bool) {
	if o == nil || o.AllowedServers == nil {
		return nil, false
	}
	return o.AllowedServers, true
}

// HasAllowedServers returns a boolean if a field has been set.
func (o *PlexSource) HasAllowedServers() bool {
	if o != nil && o.AllowedServers != nil {
		return true
	}

	return false
}

// SetAllowedServers gets a reference to the given []string and assigns it to the AllowedServers field.
func (o *PlexSource) SetAllowedServers(v []string) {
	o.AllowedServers = v
}

// GetAllowFriends returns the AllowFriends field value if set, zero value otherwise.
func (o *PlexSource) GetAllowFriends() bool {
	if o == nil || o.AllowFriends == nil {
		var ret bool
		return ret
	}
	return *o.AllowFriends
}

// GetAllowFriendsOk returns a tuple with the AllowFriends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexSource) GetAllowFriendsOk() (*bool, bool) {
	if o == nil || o.AllowFriends == nil {
		return nil, false
	}
	return o.AllowFriends, true
}

// HasAllowFriends returns a boolean if a field has been set.
func (o *PlexSource) HasAllowFriends() bool {
	if o != nil && o.AllowFriends != nil {
		return true
	}

	return false
}

// SetAllowFriends gets a reference to the given bool and assigns it to the AllowFriends field.
func (o *PlexSource) SetAllowFriends(v bool) {
	o.AllowFriends = &v
}

// GetPlexToken returns the PlexToken field value
func (o *PlexSource) GetPlexToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlexToken
}

// GetPlexTokenOk returns a tuple with the PlexToken field value
// and a boolean to check if the value has been set.
func (o *PlexSource) GetPlexTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlexToken, true
}

// SetPlexToken sets field value
func (o *PlexSource) SetPlexToken(v string) {
	o.PlexToken = v
}

func (o PlexSource) MarshalJSON() ([]byte, error) {
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
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.AuthenticationFlow.IsSet() {
		toSerialize["authentication_flow"] = o.AuthenticationFlow.Get()
	}
	if o.EnrollmentFlow.IsSet() {
		toSerialize["enrollment_flow"] = o.EnrollmentFlow.Get()
	}
	if o.UserPropertyMappings != nil {
		toSerialize["user_property_mappings"] = o.UserPropertyMappings
	}
	if o.GroupPropertyMappings != nil {
		toSerialize["group_property_mappings"] = o.GroupPropertyMappings
	}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["verbose_name"] = o.VerboseName
	}
	if true {
		toSerialize["verbose_name_plural"] = o.VerboseNamePlural
	}
	if true {
		toSerialize["meta_model_name"] = o.MetaModelName
	}
	if o.PolicyEngineMode != nil {
		toSerialize["policy_engine_mode"] = o.PolicyEngineMode
	}
	if o.UserMatchingMode != nil {
		toSerialize["user_matching_mode"] = o.UserMatchingMode
	}
	if true {
		toSerialize["managed"] = o.Managed.Get()
	}
	if o.UserPathTemplate != nil {
		toSerialize["user_path_template"] = o.UserPathTemplate
	}
	if true {
		toSerialize["icon"] = o.Icon
	}
	if o.GroupMatchingMode != nil {
		toSerialize["group_matching_mode"] = o.GroupMatchingMode
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
	if true {
		toSerialize["plex_token"] = o.PlexToken
	}
	return json.Marshal(toSerialize)
}

type NullablePlexSource struct {
	value *PlexSource
	isSet bool
}

func (v NullablePlexSource) Get() *PlexSource {
	return v.value
}

func (v *NullablePlexSource) Set(val *PlexSource) {
	v.value = val
	v.isSet = true
}

func (v NullablePlexSource) IsSet() bool {
	return v.isSet
}

func (v *NullablePlexSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlexSource(val *PlexSource) *NullablePlexSource {
	return &NullablePlexSource{value: val, isSet: true}
}

func (v NullablePlexSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlexSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
