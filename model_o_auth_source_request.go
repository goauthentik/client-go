/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the OAuthSourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuthSourceRequest{}

// OAuthSourceRequest OAuth Source Serializer
type OAuthSourceRequest struct {
	// Source's display Name.
	Name string `json:"name"`
	// Internal source name, used in URLs.
	Slug    string `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Enabled *bool  `json:"enabled,omitempty"`
	// Flow to use when authenticating existing users.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow to use when enrolling new users.
	EnrollmentFlow        NullableString    `json:"enrollment_flow,omitempty"`
	UserPropertyMappings  []string          `json:"user_property_mappings,omitempty"`
	GroupPropertyMappings []string          `json:"group_property_mappings,omitempty"`
	PolicyEngineMode      *PolicyEngineMode `json:"policy_engine_mode,omitempty"`
	// How the source determines if an existing user should be authenticated or a new user enrolled.
	UserMatchingMode *UserMatchingModeEnum `json:"user_matching_mode,omitempty"`
	UserPathTemplate *string               `json:"user_path_template,omitempty"`
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
	ProfileUrl       NullableString `json:"profile_url,omitempty"`
	ConsumerKey      string         `json:"consumer_key"`
	ConsumerSecret   string         `json:"consumer_secret"`
	AdditionalScopes *string        `json:"additional_scopes,omitempty"`
	OidcWellKnownUrl *string        `json:"oidc_well_known_url,omitempty"`
	OidcJwksUrl      *string        `json:"oidc_jwks_url,omitempty"`
	OidcJwks         interface{}    `json:"oidc_jwks,omitempty"`
}

type _OAuthSourceRequest OAuthSourceRequest

// NewOAuthSourceRequest instantiates a new OAuthSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthSourceRequest(name string, slug string, providerType ProviderTypeEnum, consumerKey string, consumerSecret string) *OAuthSourceRequest {
	this := OAuthSourceRequest{}
	this.Name = name
	this.Slug = slug
	this.ProviderType = providerType
	this.ConsumerKey = consumerKey
	this.ConsumerSecret = consumerSecret
	return &this
}

// NewOAuthSourceRequestWithDefaults instantiates a new OAuthSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthSourceRequestWithDefaults() *OAuthSourceRequest {
	this := OAuthSourceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *OAuthSourceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OAuthSourceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OAuthSourceRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *OAuthSourceRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *OAuthSourceRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *OAuthSourceRequest) SetSlug(v string) {
	o.Slug = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OAuthSourceRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSourceRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OAuthSourceRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OAuthSourceRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthSourceRequest) GetAuthenticationFlow() string {
	if o == nil || IsNil(o.AuthenticationFlow.Get()) {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSourceRequest) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *OAuthSourceRequest) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *OAuthSourceRequest) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *OAuthSourceRequest) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *OAuthSourceRequest) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetEnrollmentFlow returns the EnrollmentFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthSourceRequest) GetEnrollmentFlow() string {
	if o == nil || IsNil(o.EnrollmentFlow.Get()) {
		var ret string
		return ret
	}
	return *o.EnrollmentFlow.Get()
}

// GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSourceRequest) GetEnrollmentFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentFlow.Get(), o.EnrollmentFlow.IsSet()
}

// HasEnrollmentFlow returns a boolean if a field has been set.
func (o *OAuthSourceRequest) HasEnrollmentFlow() bool {
	if o != nil && o.EnrollmentFlow.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentFlow gets a reference to the given NullableString and assigns it to the EnrollmentFlow field.
func (o *OAuthSourceRequest) SetEnrollmentFlow(v string) {
	o.EnrollmentFlow.Set(&v)
}

// SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil
func (o *OAuthSourceRequest) SetEnrollmentFlowNil() {
	o.EnrollmentFlow.Set(nil)
}

// UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
func (o *OAuthSourceRequest) UnsetEnrollmentFlow() {
	o.EnrollmentFlow.Unset()
}

// GetUserPropertyMappings returns the UserPropertyMappings field value if set, zero value otherwise.
func (o *OAuthSourceRequest) GetUserPropertyMappings() []string {
	if o == nil || IsNil(o.UserPropertyMappings) {
		var ret []string
		return ret
	}
	return o.UserPropertyMappings
}

// GetUserPropertyMappingsOk returns a tuple with the UserPropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSourceRequest) GetUserPropertyMappingsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserPropertyMappings) {
		return nil, false
	}
	return o.UserPropertyMappings, true
}

// HasUserPropertyMappings returns a boolean if a field has been set.
func (o *OAuthSourceRequest) HasUserPropertyMappings() bool {
	if o != nil && !IsNil(o.UserPropertyMappings) {
		return true
	}

	return false
}

// SetUserPropertyMappings gets a reference to the given []string and assigns it to the UserPropertyMappings field.
func (o *OAuthSourceRequest) SetUserPropertyMappings(v []string) {
	o.UserPropertyMappings = v
}

// GetGroupPropertyMappings returns the GroupPropertyMappings field value if set, zero value otherwise.
func (o *OAuthSourceRequest) GetGroupPropertyMappings() []string {
	if o == nil || IsNil(o.GroupPropertyMappings) {
		var ret []string
		return ret
	}
	return o.GroupPropertyMappings
}

// GetGroupPropertyMappingsOk returns a tuple with the GroupPropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSourceRequest) GetGroupPropertyMappingsOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupPropertyMappings) {
		return nil, false
	}
	return o.GroupPropertyMappings, true
}

// HasGroupPropertyMappings returns a boolean if a field has been set.
func (o *OAuthSourceRequest) HasGroupPropertyMappings() bool {
	if o != nil && !IsNil(o.GroupPropertyMappings) {
		return true
	}

	return false
}

// SetGroupPropertyMappings gets a reference to the given []string and assigns it to the GroupPropertyMappings field.
func (o *OAuthSourceRequest) SetGroupPropertyMappings(v []string) {
	o.GroupPropertyMappings = v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *OAuthSourceRequest) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || IsNil(o.PolicyEngineMode) {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSourceRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || IsNil(o.PolicyEngineMode) {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *OAuthSourceRequest) HasPolicyEngineMode() bool {
	if o != nil && !IsNil(o.PolicyEngineMode) {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *OAuthSourceRequest) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetUserMatchingMode returns the UserMatchingMode field value if set, zero value otherwise.
func (o *OAuthSourceRequest) GetUserMatchingMode() UserMatchingModeEnum {
	if o == nil || IsNil(o.UserMatchingMode) {
		var ret UserMatchingModeEnum
		return ret
	}
	return *o.UserMatchingMode
}

// GetUserMatchingModeOk returns a tuple with the UserMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool) {
	if o == nil || IsNil(o.UserMatchingMode) {
		return nil, false
	}
	return o.UserMatchingMode, true
}

// HasUserMatchingMode returns a boolean if a field has been set.
func (o *OAuthSourceRequest) HasUserMatchingMode() bool {
	if o != nil && !IsNil(o.UserMatchingMode) {
		return true
	}

	return false
}

// SetUserMatchingMode gets a reference to the given UserMatchingModeEnum and assigns it to the UserMatchingMode field.
func (o *OAuthSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum) {
	o.UserMatchingMode = &v
}

// GetUserPathTemplate returns the UserPathTemplate field value if set, zero value otherwise.
func (o *OAuthSourceRequest) GetUserPathTemplate() string {
	if o == nil || IsNil(o.UserPathTemplate) {
		var ret string
		return ret
	}
	return *o.UserPathTemplate
}

// GetUserPathTemplateOk returns a tuple with the UserPathTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSourceRequest) GetUserPathTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.UserPathTemplate) {
		return nil, false
	}
	return o.UserPathTemplate, true
}

// HasUserPathTemplate returns a boolean if a field has been set.
func (o *OAuthSourceRequest) HasUserPathTemplate() bool {
	if o != nil && !IsNil(o.UserPathTemplate) {
		return true
	}

	return false
}

// SetUserPathTemplate gets a reference to the given string and assigns it to the UserPathTemplate field.
func (o *OAuthSourceRequest) SetUserPathTemplate(v string) {
	o.UserPathTemplate = &v
}

// GetGroupMatchingMode returns the GroupMatchingMode field value if set, zero value otherwise.
func (o *OAuthSourceRequest) GetGroupMatchingMode() GroupMatchingModeEnum {
	if o == nil || IsNil(o.GroupMatchingMode) {
		var ret GroupMatchingModeEnum
		return ret
	}
	return *o.GroupMatchingMode
}

// GetGroupMatchingModeOk returns a tuple with the GroupMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSourceRequest) GetGroupMatchingModeOk() (*GroupMatchingModeEnum, bool) {
	if o == nil || IsNil(o.GroupMatchingMode) {
		return nil, false
	}
	return o.GroupMatchingMode, true
}

// HasGroupMatchingMode returns a boolean if a field has been set.
func (o *OAuthSourceRequest) HasGroupMatchingMode() bool {
	if o != nil && !IsNil(o.GroupMatchingMode) {
		return true
	}

	return false
}

// SetGroupMatchingMode gets a reference to the given GroupMatchingModeEnum and assigns it to the GroupMatchingMode field.
func (o *OAuthSourceRequest) SetGroupMatchingMode(v GroupMatchingModeEnum) {
	o.GroupMatchingMode = &v
}

// GetProviderType returns the ProviderType field value
func (o *OAuthSourceRequest) GetProviderType() ProviderTypeEnum {
	if o == nil {
		var ret ProviderTypeEnum
		return ret
	}

	return o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value
// and a boolean to check if the value has been set.
func (o *OAuthSourceRequest) GetProviderTypeOk() (*ProviderTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderType, true
}

// SetProviderType sets field value
func (o *OAuthSourceRequest) SetProviderType(v ProviderTypeEnum) {
	o.ProviderType = v
}

// GetRequestTokenUrl returns the RequestTokenUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthSourceRequest) GetRequestTokenUrl() string {
	if o == nil || IsNil(o.RequestTokenUrl.Get()) {
		var ret string
		return ret
	}
	return *o.RequestTokenUrl.Get()
}

// GetRequestTokenUrlOk returns a tuple with the RequestTokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSourceRequest) GetRequestTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestTokenUrl.Get(), o.RequestTokenUrl.IsSet()
}

// HasRequestTokenUrl returns a boolean if a field has been set.
func (o *OAuthSourceRequest) HasRequestTokenUrl() bool {
	if o != nil && o.RequestTokenUrl.IsSet() {
		return true
	}

	return false
}

// SetRequestTokenUrl gets a reference to the given NullableString and assigns it to the RequestTokenUrl field.
func (o *OAuthSourceRequest) SetRequestTokenUrl(v string) {
	o.RequestTokenUrl.Set(&v)
}

// SetRequestTokenUrlNil sets the value for RequestTokenUrl to be an explicit nil
func (o *OAuthSourceRequest) SetRequestTokenUrlNil() {
	o.RequestTokenUrl.Set(nil)
}

// UnsetRequestTokenUrl ensures that no value is present for RequestTokenUrl, not even an explicit nil
func (o *OAuthSourceRequest) UnsetRequestTokenUrl() {
	o.RequestTokenUrl.Unset()
}

// GetAuthorizationUrl returns the AuthorizationUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthSourceRequest) GetAuthorizationUrl() string {
	if o == nil || IsNil(o.AuthorizationUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AuthorizationUrl.Get()
}

// GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSourceRequest) GetAuthorizationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationUrl.Get(), o.AuthorizationUrl.IsSet()
}

// HasAuthorizationUrl returns a boolean if a field has been set.
func (o *OAuthSourceRequest) HasAuthorizationUrl() bool {
	if o != nil && o.AuthorizationUrl.IsSet() {
		return true
	}

	return false
}

// SetAuthorizationUrl gets a reference to the given NullableString and assigns it to the AuthorizationUrl field.
func (o *OAuthSourceRequest) SetAuthorizationUrl(v string) {
	o.AuthorizationUrl.Set(&v)
}

// SetAuthorizationUrlNil sets the value for AuthorizationUrl to be an explicit nil
func (o *OAuthSourceRequest) SetAuthorizationUrlNil() {
	o.AuthorizationUrl.Set(nil)
}

// UnsetAuthorizationUrl ensures that no value is present for AuthorizationUrl, not even an explicit nil
func (o *OAuthSourceRequest) UnsetAuthorizationUrl() {
	o.AuthorizationUrl.Unset()
}

// GetAccessTokenUrl returns the AccessTokenUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthSourceRequest) GetAccessTokenUrl() string {
	if o == nil || IsNil(o.AccessTokenUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AccessTokenUrl.Get()
}

// GetAccessTokenUrlOk returns a tuple with the AccessTokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSourceRequest) GetAccessTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessTokenUrl.Get(), o.AccessTokenUrl.IsSet()
}

// HasAccessTokenUrl returns a boolean if a field has been set.
func (o *OAuthSourceRequest) HasAccessTokenUrl() bool {
	if o != nil && o.AccessTokenUrl.IsSet() {
		return true
	}

	return false
}

// SetAccessTokenUrl gets a reference to the given NullableString and assigns it to the AccessTokenUrl field.
func (o *OAuthSourceRequest) SetAccessTokenUrl(v string) {
	o.AccessTokenUrl.Set(&v)
}

// SetAccessTokenUrlNil sets the value for AccessTokenUrl to be an explicit nil
func (o *OAuthSourceRequest) SetAccessTokenUrlNil() {
	o.AccessTokenUrl.Set(nil)
}

// UnsetAccessTokenUrl ensures that no value is present for AccessTokenUrl, not even an explicit nil
func (o *OAuthSourceRequest) UnsetAccessTokenUrl() {
	o.AccessTokenUrl.Unset()
}

// GetProfileUrl returns the ProfileUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthSourceRequest) GetProfileUrl() string {
	if o == nil || IsNil(o.ProfileUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileUrl.Get()
}

// GetProfileUrlOk returns a tuple with the ProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSourceRequest) GetProfileUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileUrl.Get(), o.ProfileUrl.IsSet()
}

// HasProfileUrl returns a boolean if a field has been set.
func (o *OAuthSourceRequest) HasProfileUrl() bool {
	if o != nil && o.ProfileUrl.IsSet() {
		return true
	}

	return false
}

// SetProfileUrl gets a reference to the given NullableString and assigns it to the ProfileUrl field.
func (o *OAuthSourceRequest) SetProfileUrl(v string) {
	o.ProfileUrl.Set(&v)
}

// SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil
func (o *OAuthSourceRequest) SetProfileUrlNil() {
	o.ProfileUrl.Set(nil)
}

// UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
func (o *OAuthSourceRequest) UnsetProfileUrl() {
	o.ProfileUrl.Unset()
}

// GetConsumerKey returns the ConsumerKey field value
func (o *OAuthSourceRequest) GetConsumerKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsumerKey
}

// GetConsumerKeyOk returns a tuple with the ConsumerKey field value
// and a boolean to check if the value has been set.
func (o *OAuthSourceRequest) GetConsumerKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerKey, true
}

// SetConsumerKey sets field value
func (o *OAuthSourceRequest) SetConsumerKey(v string) {
	o.ConsumerKey = v
}

// GetConsumerSecret returns the ConsumerSecret field value
func (o *OAuthSourceRequest) GetConsumerSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsumerSecret
}

// GetConsumerSecretOk returns a tuple with the ConsumerSecret field value
// and a boolean to check if the value has been set.
func (o *OAuthSourceRequest) GetConsumerSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerSecret, true
}

// SetConsumerSecret sets field value
func (o *OAuthSourceRequest) SetConsumerSecret(v string) {
	o.ConsumerSecret = v
}

// GetAdditionalScopes returns the AdditionalScopes field value if set, zero value otherwise.
func (o *OAuthSourceRequest) GetAdditionalScopes() string {
	if o == nil || IsNil(o.AdditionalScopes) {
		var ret string
		return ret
	}
	return *o.AdditionalScopes
}

// GetAdditionalScopesOk returns a tuple with the AdditionalScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSourceRequest) GetAdditionalScopesOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalScopes) {
		return nil, false
	}
	return o.AdditionalScopes, true
}

// HasAdditionalScopes returns a boolean if a field has been set.
func (o *OAuthSourceRequest) HasAdditionalScopes() bool {
	if o != nil && !IsNil(o.AdditionalScopes) {
		return true
	}

	return false
}

// SetAdditionalScopes gets a reference to the given string and assigns it to the AdditionalScopes field.
func (o *OAuthSourceRequest) SetAdditionalScopes(v string) {
	o.AdditionalScopes = &v
}

// GetOidcWellKnownUrl returns the OidcWellKnownUrl field value if set, zero value otherwise.
func (o *OAuthSourceRequest) GetOidcWellKnownUrl() string {
	if o == nil || IsNil(o.OidcWellKnownUrl) {
		var ret string
		return ret
	}
	return *o.OidcWellKnownUrl
}

// GetOidcWellKnownUrlOk returns a tuple with the OidcWellKnownUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSourceRequest) GetOidcWellKnownUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OidcWellKnownUrl) {
		return nil, false
	}
	return o.OidcWellKnownUrl, true
}

// HasOidcWellKnownUrl returns a boolean if a field has been set.
func (o *OAuthSourceRequest) HasOidcWellKnownUrl() bool {
	if o != nil && !IsNil(o.OidcWellKnownUrl) {
		return true
	}

	return false
}

// SetOidcWellKnownUrl gets a reference to the given string and assigns it to the OidcWellKnownUrl field.
func (o *OAuthSourceRequest) SetOidcWellKnownUrl(v string) {
	o.OidcWellKnownUrl = &v
}

// GetOidcJwksUrl returns the OidcJwksUrl field value if set, zero value otherwise.
func (o *OAuthSourceRequest) GetOidcJwksUrl() string {
	if o == nil || IsNil(o.OidcJwksUrl) {
		var ret string
		return ret
	}
	return *o.OidcJwksUrl
}

// GetOidcJwksUrlOk returns a tuple with the OidcJwksUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSourceRequest) GetOidcJwksUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OidcJwksUrl) {
		return nil, false
	}
	return o.OidcJwksUrl, true
}

// HasOidcJwksUrl returns a boolean if a field has been set.
func (o *OAuthSourceRequest) HasOidcJwksUrl() bool {
	if o != nil && !IsNil(o.OidcJwksUrl) {
		return true
	}

	return false
}

// SetOidcJwksUrl gets a reference to the given string and assigns it to the OidcJwksUrl field.
func (o *OAuthSourceRequest) SetOidcJwksUrl(v string) {
	o.OidcJwksUrl = &v
}

// GetOidcJwks returns the OidcJwks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthSourceRequest) GetOidcJwks() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OidcJwks
}

// GetOidcJwksOk returns a tuple with the OidcJwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthSourceRequest) GetOidcJwksOk() (*interface{}, bool) {
	if o == nil || IsNil(o.OidcJwks) {
		return nil, false
	}
	return &o.OidcJwks, true
}

// HasOidcJwks returns a boolean if a field has been set.
func (o *OAuthSourceRequest) HasOidcJwks() bool {
	if o != nil && !IsNil(o.OidcJwks) {
		return true
	}

	return false
}

// SetOidcJwks gets a reference to the given interface{} and assigns it to the OidcJwks field.
func (o *OAuthSourceRequest) SetOidcJwks(v interface{}) {
	o.OidcJwks = v
}

func (o OAuthSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuthSourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.AuthenticationFlow.IsSet() {
		toSerialize["authentication_flow"] = o.AuthenticationFlow.Get()
	}
	if o.EnrollmentFlow.IsSet() {
		toSerialize["enrollment_flow"] = o.EnrollmentFlow.Get()
	}
	if !IsNil(o.UserPropertyMappings) {
		toSerialize["user_property_mappings"] = o.UserPropertyMappings
	}
	if !IsNil(o.GroupPropertyMappings) {
		toSerialize["group_property_mappings"] = o.GroupPropertyMappings
	}
	if !IsNil(o.PolicyEngineMode) {
		toSerialize["policy_engine_mode"] = o.PolicyEngineMode
	}
	if !IsNil(o.UserMatchingMode) {
		toSerialize["user_matching_mode"] = o.UserMatchingMode
	}
	if !IsNil(o.UserPathTemplate) {
		toSerialize["user_path_template"] = o.UserPathTemplate
	}
	if !IsNil(o.GroupMatchingMode) {
		toSerialize["group_matching_mode"] = o.GroupMatchingMode
	}
	toSerialize["provider_type"] = o.ProviderType
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
	toSerialize["consumer_key"] = o.ConsumerKey
	toSerialize["consumer_secret"] = o.ConsumerSecret
	if !IsNil(o.AdditionalScopes) {
		toSerialize["additional_scopes"] = o.AdditionalScopes
	}
	if !IsNil(o.OidcWellKnownUrl) {
		toSerialize["oidc_well_known_url"] = o.OidcWellKnownUrl
	}
	if !IsNil(o.OidcJwksUrl) {
		toSerialize["oidc_jwks_url"] = o.OidcJwksUrl
	}
	if o.OidcJwks != nil {
		toSerialize["oidc_jwks"] = o.OidcJwks
	}
	return toSerialize, nil
}

func (o *OAuthSourceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"slug",
		"provider_type",
		"consumer_key",
		"consumer_secret",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOAuthSourceRequest := _OAuthSourceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOAuthSourceRequest)

	if err != nil {
		return err
	}

	*o = OAuthSourceRequest(varOAuthSourceRequest)

	return err
}

type NullableOAuthSourceRequest struct {
	value *OAuthSourceRequest
	isSet bool
}

func (v NullableOAuthSourceRequest) Get() *OAuthSourceRequest {
	return v.value
}

func (v *NullableOAuthSourceRequest) Set(val *OAuthSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthSourceRequest(val *OAuthSourceRequest) *NullableOAuthSourceRequest {
	return &NullableOAuthSourceRequest{value: val, isSet: true}
}

func (v NullableOAuthSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
