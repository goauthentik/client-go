/*
authentik

Making authentication simple.

API version: 2024.10.5
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// OAuthSource OAuth Source Serializer
type OAuthSource struct {
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
	Icon             NullableString `json:"icon"`
	// How the source determines if an existing group should be used or a new group created.
	GroupMatchingMode *GroupMatchingModeEnum `json:"group_matching_mode,omitempty"`
	ProviderType      ProviderTypeEnum       `json:"provider_type"`
	// URL used to request the initial token. This URL is only required for OAuth 1.
	RequestTokenUrl NullableString `json:"request_token_url,omitempty"`
	// URL the user is redirect to to conest the flow.
	AuthorizationUrl NullableString `json:"authorization_url,omitempty"`
	// URL used by authentik to retrieve tokens.
	AccessTokenUrl NullableString `json:"access_token_url,omitempty"`
	// URL used by authentik to get user information.
	ProfileUrl  NullableString `json:"profile_url,omitempty"`
	ConsumerKey string         `json:"consumer_key"`
	// Get OAuth Callback URL
	CallbackUrl      string      `json:"callback_url"`
	AdditionalScopes *string     `json:"additional_scopes,omitempty"`
	Type             SourceType  `json:"type"`
	OidcWellKnownUrl *string     `json:"oidc_well_known_url,omitempty"`
	OidcJwksUrl      *string     `json:"oidc_jwks_url,omitempty"`
	OidcJwks         interface{} `json:"oidc_jwks,omitempty"`
}

// NewOAuthSource instantiates a new OAuthSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthSource(pk string, name string, slug string, component string, verboseName string, verboseNamePlural string, metaModelName string, managed NullableString, icon NullableString, providerType ProviderTypeEnum, consumerKey string, callbackUrl string, type_ SourceType) *OAuthSource {
	this := OAuthSource{}
	this.Pk = pk
	this.Name = name
	this.Slug = slug
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.Managed = managed
	this.Icon = icon
	this.ProviderType = providerType
	this.ConsumerKey = consumerKey
	this.CallbackUrl = callbackUrl
	this.Type = type_
	return &this
}

// NewOAuthSourceWithDefaults instantiates a new OAuthSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthSourceWithDefaults() *OAuthSource {
	this := OAuthSource{}
	return &this
}

// GetPk returns the Pk field value
func (o *OAuthSource) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *OAuthSource) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *OAuthSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OAuthSource) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *OAuthSource) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *OAuthSource) SetSlug(v string) {
	o.Slug = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OAuthSource) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OAuthSource) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OAuthSource) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthSource) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSource) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *OAuthSource) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *OAuthSource) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *OAuthSource) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *OAuthSource) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetEnrollmentFlow returns the EnrollmentFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthSource) GetEnrollmentFlow() string {
	if o == nil || o.EnrollmentFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentFlow.Get()
}

// GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSource) GetEnrollmentFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentFlow.Get(), o.EnrollmentFlow.IsSet()
}

// HasEnrollmentFlow returns a boolean if a field has been set.
func (o *OAuthSource) HasEnrollmentFlow() bool {
	if o != nil && o.EnrollmentFlow.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentFlow gets a reference to the given NullableString and assigns it to the EnrollmentFlow field.
func (o *OAuthSource) SetEnrollmentFlow(v string) {
	o.EnrollmentFlow.Set(&v)
}

// SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil
func (o *OAuthSource) SetEnrollmentFlowNil() {
	o.EnrollmentFlow.Set(nil)
}

// UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
func (o *OAuthSource) UnsetEnrollmentFlow() {
	o.EnrollmentFlow.Unset()
}

// GetUserPropertyMappings returns the UserPropertyMappings field value if set, zero value otherwise.
func (o *OAuthSource) GetUserPropertyMappings() []string {
	if o == nil || o.UserPropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.UserPropertyMappings
}

// GetUserPropertyMappingsOk returns a tuple with the UserPropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetUserPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.UserPropertyMappings == nil {
		return nil, false
	}
	return o.UserPropertyMappings, true
}

// HasUserPropertyMappings returns a boolean if a field has been set.
func (o *OAuthSource) HasUserPropertyMappings() bool {
	if o != nil && o.UserPropertyMappings != nil {
		return true
	}

	return false
}

// SetUserPropertyMappings gets a reference to the given []string and assigns it to the UserPropertyMappings field.
func (o *OAuthSource) SetUserPropertyMappings(v []string) {
	o.UserPropertyMappings = v
}

// GetGroupPropertyMappings returns the GroupPropertyMappings field value if set, zero value otherwise.
func (o *OAuthSource) GetGroupPropertyMappings() []string {
	if o == nil || o.GroupPropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.GroupPropertyMappings
}

// GetGroupPropertyMappingsOk returns a tuple with the GroupPropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetGroupPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.GroupPropertyMappings == nil {
		return nil, false
	}
	return o.GroupPropertyMappings, true
}

// HasGroupPropertyMappings returns a boolean if a field has been set.
func (o *OAuthSource) HasGroupPropertyMappings() bool {
	if o != nil && o.GroupPropertyMappings != nil {
		return true
	}

	return false
}

// SetGroupPropertyMappings gets a reference to the given []string and assigns it to the GroupPropertyMappings field.
func (o *OAuthSource) SetGroupPropertyMappings(v []string) {
	o.GroupPropertyMappings = v
}

// GetComponent returns the Component field value
func (o *OAuthSource) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *OAuthSource) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *OAuthSource) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *OAuthSource) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *OAuthSource) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *OAuthSource) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *OAuthSource) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *OAuthSource) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *OAuthSource) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *OAuthSource) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *OAuthSource) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetUserMatchingMode returns the UserMatchingMode field value if set, zero value otherwise.
func (o *OAuthSource) GetUserMatchingMode() UserMatchingModeEnum {
	if o == nil || o.UserMatchingMode == nil {
		var ret UserMatchingModeEnum
		return ret
	}
	return *o.UserMatchingMode
}

// GetUserMatchingModeOk returns a tuple with the UserMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool) {
	if o == nil || o.UserMatchingMode == nil {
		return nil, false
	}
	return o.UserMatchingMode, true
}

// HasUserMatchingMode returns a boolean if a field has been set.
func (o *OAuthSource) HasUserMatchingMode() bool {
	if o != nil && o.UserMatchingMode != nil {
		return true
	}

	return false
}

// SetUserMatchingMode gets a reference to the given UserMatchingModeEnum and assigns it to the UserMatchingMode field.
func (o *OAuthSource) SetUserMatchingMode(v UserMatchingModeEnum) {
	o.UserMatchingMode = &v
}

// GetManaged returns the Managed field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OAuthSource) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}

	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSource) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// SetManaged sets field value
func (o *OAuthSource) SetManaged(v string) {
	o.Managed.Set(&v)
}

// GetUserPathTemplate returns the UserPathTemplate field value if set, zero value otherwise.
func (o *OAuthSource) GetUserPathTemplate() string {
	if o == nil || o.UserPathTemplate == nil {
		var ret string
		return ret
	}
	return *o.UserPathTemplate
}

// GetUserPathTemplateOk returns a tuple with the UserPathTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetUserPathTemplateOk() (*string, bool) {
	if o == nil || o.UserPathTemplate == nil {
		return nil, false
	}
	return o.UserPathTemplate, true
}

// HasUserPathTemplate returns a boolean if a field has been set.
func (o *OAuthSource) HasUserPathTemplate() bool {
	if o != nil && o.UserPathTemplate != nil {
		return true
	}

	return false
}

// SetUserPathTemplate gets a reference to the given string and assigns it to the UserPathTemplate field.
func (o *OAuthSource) SetUserPathTemplate(v string) {
	o.UserPathTemplate = &v
}

// GetIcon returns the Icon field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OAuthSource) GetIcon() string {
	if o == nil || o.Icon.Get() == nil {
		var ret string
		return ret
	}

	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSource) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// SetIcon sets field value
func (o *OAuthSource) SetIcon(v string) {
	o.Icon.Set(&v)
}

// GetGroupMatchingMode returns the GroupMatchingMode field value if set, zero value otherwise.
func (o *OAuthSource) GetGroupMatchingMode() GroupMatchingModeEnum {
	if o == nil || o.GroupMatchingMode == nil {
		var ret GroupMatchingModeEnum
		return ret
	}
	return *o.GroupMatchingMode
}

// GetGroupMatchingModeOk returns a tuple with the GroupMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetGroupMatchingModeOk() (*GroupMatchingModeEnum, bool) {
	if o == nil || o.GroupMatchingMode == nil {
		return nil, false
	}
	return o.GroupMatchingMode, true
}

// HasGroupMatchingMode returns a boolean if a field has been set.
func (o *OAuthSource) HasGroupMatchingMode() bool {
	if o != nil && o.GroupMatchingMode != nil {
		return true
	}

	return false
}

// SetGroupMatchingMode gets a reference to the given GroupMatchingModeEnum and assigns it to the GroupMatchingMode field.
func (o *OAuthSource) SetGroupMatchingMode(v GroupMatchingModeEnum) {
	o.GroupMatchingMode = &v
}

// GetProviderType returns the ProviderType field value
func (o *OAuthSource) GetProviderType() ProviderTypeEnum {
	if o == nil {
		var ret ProviderTypeEnum
		return ret
	}

	return o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetProviderTypeOk() (*ProviderTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderType, true
}

// SetProviderType sets field value
func (o *OAuthSource) SetProviderType(v ProviderTypeEnum) {
	o.ProviderType = v
}

// GetRequestTokenUrl returns the RequestTokenUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthSource) GetRequestTokenUrl() string {
	if o == nil || o.RequestTokenUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.RequestTokenUrl.Get()
}

// GetRequestTokenUrlOk returns a tuple with the RequestTokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSource) GetRequestTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestTokenUrl.Get(), o.RequestTokenUrl.IsSet()
}

// HasRequestTokenUrl returns a boolean if a field has been set.
func (o *OAuthSource) HasRequestTokenUrl() bool {
	if o != nil && o.RequestTokenUrl.IsSet() {
		return true
	}

	return false
}

// SetRequestTokenUrl gets a reference to the given NullableString and assigns it to the RequestTokenUrl field.
func (o *OAuthSource) SetRequestTokenUrl(v string) {
	o.RequestTokenUrl.Set(&v)
}

// SetRequestTokenUrlNil sets the value for RequestTokenUrl to be an explicit nil
func (o *OAuthSource) SetRequestTokenUrlNil() {
	o.RequestTokenUrl.Set(nil)
}

// UnsetRequestTokenUrl ensures that no value is present for RequestTokenUrl, not even an explicit nil
func (o *OAuthSource) UnsetRequestTokenUrl() {
	o.RequestTokenUrl.Unset()
}

// GetAuthorizationUrl returns the AuthorizationUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthSource) GetAuthorizationUrl() string {
	if o == nil || o.AuthorizationUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationUrl.Get()
}

// GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSource) GetAuthorizationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationUrl.Get(), o.AuthorizationUrl.IsSet()
}

// HasAuthorizationUrl returns a boolean if a field has been set.
func (o *OAuthSource) HasAuthorizationUrl() bool {
	if o != nil && o.AuthorizationUrl.IsSet() {
		return true
	}

	return false
}

// SetAuthorizationUrl gets a reference to the given NullableString and assigns it to the AuthorizationUrl field.
func (o *OAuthSource) SetAuthorizationUrl(v string) {
	o.AuthorizationUrl.Set(&v)
}

// SetAuthorizationUrlNil sets the value for AuthorizationUrl to be an explicit nil
func (o *OAuthSource) SetAuthorizationUrlNil() {
	o.AuthorizationUrl.Set(nil)
}

// UnsetAuthorizationUrl ensures that no value is present for AuthorizationUrl, not even an explicit nil
func (o *OAuthSource) UnsetAuthorizationUrl() {
	o.AuthorizationUrl.Unset()
}

// GetAccessTokenUrl returns the AccessTokenUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthSource) GetAccessTokenUrl() string {
	if o == nil || o.AccessTokenUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccessTokenUrl.Get()
}

// GetAccessTokenUrlOk returns a tuple with the AccessTokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSource) GetAccessTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessTokenUrl.Get(), o.AccessTokenUrl.IsSet()
}

// HasAccessTokenUrl returns a boolean if a field has been set.
func (o *OAuthSource) HasAccessTokenUrl() bool {
	if o != nil && o.AccessTokenUrl.IsSet() {
		return true
	}

	return false
}

// SetAccessTokenUrl gets a reference to the given NullableString and assigns it to the AccessTokenUrl field.
func (o *OAuthSource) SetAccessTokenUrl(v string) {
	o.AccessTokenUrl.Set(&v)
}

// SetAccessTokenUrlNil sets the value for AccessTokenUrl to be an explicit nil
func (o *OAuthSource) SetAccessTokenUrlNil() {
	o.AccessTokenUrl.Set(nil)
}

// UnsetAccessTokenUrl ensures that no value is present for AccessTokenUrl, not even an explicit nil
func (o *OAuthSource) UnsetAccessTokenUrl() {
	o.AccessTokenUrl.Unset()
}

// GetProfileUrl returns the ProfileUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthSource) GetProfileUrl() string {
	if o == nil || o.ProfileUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProfileUrl.Get()
}

// GetProfileUrlOk returns a tuple with the ProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSource) GetProfileUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileUrl.Get(), o.ProfileUrl.IsSet()
}

// HasProfileUrl returns a boolean if a field has been set.
func (o *OAuthSource) HasProfileUrl() bool {
	if o != nil && o.ProfileUrl.IsSet() {
		return true
	}

	return false
}

// SetProfileUrl gets a reference to the given NullableString and assigns it to the ProfileUrl field.
func (o *OAuthSource) SetProfileUrl(v string) {
	o.ProfileUrl.Set(&v)
}

// SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil
func (o *OAuthSource) SetProfileUrlNil() {
	o.ProfileUrl.Set(nil)
}

// UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
func (o *OAuthSource) UnsetProfileUrl() {
	o.ProfileUrl.Unset()
}

// GetConsumerKey returns the ConsumerKey field value
func (o *OAuthSource) GetConsumerKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsumerKey
}

// GetConsumerKeyOk returns a tuple with the ConsumerKey field value
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetConsumerKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerKey, true
}

// SetConsumerKey sets field value
func (o *OAuthSource) SetConsumerKey(v string) {
	o.ConsumerKey = v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *OAuthSource) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *OAuthSource) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetAdditionalScopes returns the AdditionalScopes field value if set, zero value otherwise.
func (o *OAuthSource) GetAdditionalScopes() string {
	if o == nil || o.AdditionalScopes == nil {
		var ret string
		return ret
	}
	return *o.AdditionalScopes
}

// GetAdditionalScopesOk returns a tuple with the AdditionalScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetAdditionalScopesOk() (*string, bool) {
	if o == nil || o.AdditionalScopes == nil {
		return nil, false
	}
	return o.AdditionalScopes, true
}

// HasAdditionalScopes returns a boolean if a field has been set.
func (o *OAuthSource) HasAdditionalScopes() bool {
	if o != nil && o.AdditionalScopes != nil {
		return true
	}

	return false
}

// SetAdditionalScopes gets a reference to the given string and assigns it to the AdditionalScopes field.
func (o *OAuthSource) SetAdditionalScopes(v string) {
	o.AdditionalScopes = &v
}

// GetType returns the Type field value
func (o *OAuthSource) GetType() SourceType {
	if o == nil {
		var ret SourceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetTypeOk() (*SourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OAuthSource) SetType(v SourceType) {
	o.Type = v
}

// GetOidcWellKnownUrl returns the OidcWellKnownUrl field value if set, zero value otherwise.
func (o *OAuthSource) GetOidcWellKnownUrl() string {
	if o == nil || o.OidcWellKnownUrl == nil {
		var ret string
		return ret
	}
	return *o.OidcWellKnownUrl
}

// GetOidcWellKnownUrlOk returns a tuple with the OidcWellKnownUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetOidcWellKnownUrlOk() (*string, bool) {
	if o == nil || o.OidcWellKnownUrl == nil {
		return nil, false
	}
	return o.OidcWellKnownUrl, true
}

// HasOidcWellKnownUrl returns a boolean if a field has been set.
func (o *OAuthSource) HasOidcWellKnownUrl() bool {
	if o != nil && o.OidcWellKnownUrl != nil {
		return true
	}

	return false
}

// SetOidcWellKnownUrl gets a reference to the given string and assigns it to the OidcWellKnownUrl field.
func (o *OAuthSource) SetOidcWellKnownUrl(v string) {
	o.OidcWellKnownUrl = &v
}

// GetOidcJwksUrl returns the OidcJwksUrl field value if set, zero value otherwise.
func (o *OAuthSource) GetOidcJwksUrl() string {
	if o == nil || o.OidcJwksUrl == nil {
		var ret string
		return ret
	}
	return *o.OidcJwksUrl
}

// GetOidcJwksUrlOk returns a tuple with the OidcJwksUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSource) GetOidcJwksUrlOk() (*string, bool) {
	if o == nil || o.OidcJwksUrl == nil {
		return nil, false
	}
	return o.OidcJwksUrl, true
}

// HasOidcJwksUrl returns a boolean if a field has been set.
func (o *OAuthSource) HasOidcJwksUrl() bool {
	if o != nil && o.OidcJwksUrl != nil {
		return true
	}

	return false
}

// SetOidcJwksUrl gets a reference to the given string and assigns it to the OidcJwksUrl field.
func (o *OAuthSource) SetOidcJwksUrl(v string) {
	o.OidcJwksUrl = &v
}

// GetOidcJwks returns the OidcJwks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthSource) GetOidcJwks() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OidcJwks
}

// GetOidcJwksOk returns a tuple with the OidcJwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSource) GetOidcJwksOk() (*interface{}, bool) {
	if o == nil || o.OidcJwks == nil {
		return nil, false
	}
	return &o.OidcJwks, true
}

// HasOidcJwks returns a boolean if a field has been set.
func (o *OAuthSource) HasOidcJwks() bool {
	if o != nil && o.OidcJwks != nil {
		return true
	}

	return false
}

// SetOidcJwks gets a reference to the given interface{} and assigns it to the OidcJwks field.
func (o *OAuthSource) SetOidcJwks(v interface{}) {
	o.OidcJwks = v
}

func (o OAuthSource) MarshalJSON() ([]byte, error) {
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
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.GroupMatchingMode != nil {
		toSerialize["group_matching_mode"] = o.GroupMatchingMode
	}
	if true {
		toSerialize["provider_type"] = o.ProviderType
	}
	if o.RequestTokenUrl.IsSet() {
		toSerialize["request_token_url"] = o.RequestTokenUrl.Get()
	}
	if o.AuthorizationUrl.IsSet() {
		toSerialize["authorization_url"] = o.AuthorizationUrl.Get()
	}
	if o.AccessTokenUrl.IsSet() {
		toSerialize["access_token_url"] = o.AccessTokenUrl.Get()
	}
	if o.ProfileUrl.IsSet() {
		toSerialize["profile_url"] = o.ProfileUrl.Get()
	}
	if true {
		toSerialize["consumer_key"] = o.ConsumerKey
	}
	if true {
		toSerialize["callback_url"] = o.CallbackUrl
	}
	if o.AdditionalScopes != nil {
		toSerialize["additional_scopes"] = o.AdditionalScopes
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.OidcWellKnownUrl != nil {
		toSerialize["oidc_well_known_url"] = o.OidcWellKnownUrl
	}
	if o.OidcJwksUrl != nil {
		toSerialize["oidc_jwks_url"] = o.OidcJwksUrl
	}
	if o.OidcJwks != nil {
		toSerialize["oidc_jwks"] = o.OidcJwks
	}
	return json.Marshal(toSerialize)
}

type NullableOAuthSource struct {
	value *OAuthSource
	isSet bool
}

func (v NullableOAuthSource) Get() *OAuthSource {
	return v.value
}

func (v *NullableOAuthSource) Set(val *OAuthSource) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthSource) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthSource(val *OAuthSource) *NullableOAuthSource {
	return &NullableOAuthSource{value: val, isSet: true}
}

func (v NullableOAuthSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
