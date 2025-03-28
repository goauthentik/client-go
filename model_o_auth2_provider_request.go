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

// OAuth2ProviderRequest OAuth2Provider Serializer
type OAuth2ProviderRequest struct {
	Name string `json:"name"`
	// Flow used for authentication when the associated application is accessed by an un-authenticated user.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow used when authorizing this provider.
	AuthorizationFlow string `json:"authorization_flow"`
	// Flow used ending the session from a provider.
	InvalidationFlow string   `json:"invalidation_flow"`
	PropertyMappings []string `json:"property_mappings,omitempty"`
	// Confidential clients are capable of maintaining the confidentiality of their credentials. Public clients are incapable
	ClientType   *ClientTypeEnum `json:"client_type,omitempty"`
	ClientId     *string         `json:"client_id,omitempty"`
	ClientSecret *string         `json:"client_secret,omitempty"`
	// Access codes not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	AccessCodeValidity *string `json:"access_code_validity,omitempty"`
	// Tokens not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	AccessTokenValidity *string `json:"access_token_validity,omitempty"`
	// Tokens not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	RefreshTokenValidity *string `json:"refresh_token_validity,omitempty"`
	// Include User claims from scopes in the id_token, for applications that don't access the userinfo endpoint.
	IncludeClaimsInIdToken *bool `json:"include_claims_in_id_token,omitempty"`
	// Key used to sign the tokens.
	SigningKey NullableString `json:"signing_key,omitempty"`
	// Key used to encrypt the tokens. When set, tokens will be encrypted and returned as JWEs.
	EncryptionKey NullableString       `json:"encryption_key,omitempty"`
	RedirectUris  []RedirectURIRequest `json:"redirect_uris"`
	// Configure what data should be used as unique User Identifier. For most cases, the default should be fine.
	SubMode *SubModeEnum `json:"sub_mode,omitempty"`
	// Configure how the issuer field of the ID Token should be filled.
	IssuerMode             *IssuerModeEnum `json:"issuer_mode,omitempty"`
	JwtFederationSources   []string        `json:"jwt_federation_sources,omitempty"`
	JwtFederationProviders []int32         `json:"jwt_federation_providers,omitempty"`
}

// NewOAuth2ProviderRequest instantiates a new OAuth2ProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ProviderRequest(name string, authorizationFlow string, invalidationFlow string, redirectUris []RedirectURIRequest) *OAuth2ProviderRequest {
	this := OAuth2ProviderRequest{}
	this.Name = name
	this.AuthorizationFlow = authorizationFlow
	this.InvalidationFlow = invalidationFlow
	this.RedirectUris = redirectUris
	return &this
}

// NewOAuth2ProviderRequestWithDefaults instantiates a new OAuth2ProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ProviderRequestWithDefaults() *OAuth2ProviderRequest {
	this := OAuth2ProviderRequest{}
	return &this
}

// GetName returns the Name field value
func (o *OAuth2ProviderRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OAuth2ProviderRequest) SetName(v string) {
	o.Name = v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ProviderRequest) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ProviderRequest) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *OAuth2ProviderRequest) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *OAuth2ProviderRequest) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *OAuth2ProviderRequest) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *OAuth2ProviderRequest) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetAuthorizationFlow returns the AuthorizationFlow field value
func (o *OAuth2ProviderRequest) GetAuthorizationFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderRequest) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationFlow, true
}

// SetAuthorizationFlow sets field value
func (o *OAuth2ProviderRequest) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = v
}

// GetInvalidationFlow returns the InvalidationFlow field value
func (o *OAuth2ProviderRequest) GetInvalidationFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvalidationFlow
}

// GetInvalidationFlowOk returns a tuple with the InvalidationFlow field value
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderRequest) GetInvalidationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvalidationFlow, true
}

// SetInvalidationFlow sets field value
func (o *OAuth2ProviderRequest) SetInvalidationFlow(v string) {
	o.InvalidationFlow = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *OAuth2ProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *OAuth2ProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *OAuth2ProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetClientType returns the ClientType field value if set, zero value otherwise.
func (o *OAuth2ProviderRequest) GetClientType() ClientTypeEnum {
	if o == nil || o.ClientType == nil {
		var ret ClientTypeEnum
		return ret
	}
	return *o.ClientType
}

// GetClientTypeOk returns a tuple with the ClientType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderRequest) GetClientTypeOk() (*ClientTypeEnum, bool) {
	if o == nil || o.ClientType == nil {
		return nil, false
	}
	return o.ClientType, true
}

// HasClientType returns a boolean if a field has been set.
func (o *OAuth2ProviderRequest) HasClientType() bool {
	if o != nil && o.ClientType != nil {
		return true
	}

	return false
}

// SetClientType gets a reference to the given ClientTypeEnum and assigns it to the ClientType field.
func (o *OAuth2ProviderRequest) SetClientType(v ClientTypeEnum) {
	o.ClientType = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OAuth2ProviderRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OAuth2ProviderRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OAuth2ProviderRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *OAuth2ProviderRequest) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *OAuth2ProviderRequest) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *OAuth2ProviderRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetAccessCodeValidity returns the AccessCodeValidity field value if set, zero value otherwise.
func (o *OAuth2ProviderRequest) GetAccessCodeValidity() string {
	if o == nil || o.AccessCodeValidity == nil {
		var ret string
		return ret
	}
	return *o.AccessCodeValidity
}

// GetAccessCodeValidityOk returns a tuple with the AccessCodeValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderRequest) GetAccessCodeValidityOk() (*string, bool) {
	if o == nil || o.AccessCodeValidity == nil {
		return nil, false
	}
	return o.AccessCodeValidity, true
}

// HasAccessCodeValidity returns a boolean if a field has been set.
func (o *OAuth2ProviderRequest) HasAccessCodeValidity() bool {
	if o != nil && o.AccessCodeValidity != nil {
		return true
	}

	return false
}

// SetAccessCodeValidity gets a reference to the given string and assigns it to the AccessCodeValidity field.
func (o *OAuth2ProviderRequest) SetAccessCodeValidity(v string) {
	o.AccessCodeValidity = &v
}

// GetAccessTokenValidity returns the AccessTokenValidity field value if set, zero value otherwise.
func (o *OAuth2ProviderRequest) GetAccessTokenValidity() string {
	if o == nil || o.AccessTokenValidity == nil {
		var ret string
		return ret
	}
	return *o.AccessTokenValidity
}

// GetAccessTokenValidityOk returns a tuple with the AccessTokenValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderRequest) GetAccessTokenValidityOk() (*string, bool) {
	if o == nil || o.AccessTokenValidity == nil {
		return nil, false
	}
	return o.AccessTokenValidity, true
}

// HasAccessTokenValidity returns a boolean if a field has been set.
func (o *OAuth2ProviderRequest) HasAccessTokenValidity() bool {
	if o != nil && o.AccessTokenValidity != nil {
		return true
	}

	return false
}

// SetAccessTokenValidity gets a reference to the given string and assigns it to the AccessTokenValidity field.
func (o *OAuth2ProviderRequest) SetAccessTokenValidity(v string) {
	o.AccessTokenValidity = &v
}

// GetRefreshTokenValidity returns the RefreshTokenValidity field value if set, zero value otherwise.
func (o *OAuth2ProviderRequest) GetRefreshTokenValidity() string {
	if o == nil || o.RefreshTokenValidity == nil {
		var ret string
		return ret
	}
	return *o.RefreshTokenValidity
}

// GetRefreshTokenValidityOk returns a tuple with the RefreshTokenValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderRequest) GetRefreshTokenValidityOk() (*string, bool) {
	if o == nil || o.RefreshTokenValidity == nil {
		return nil, false
	}
	return o.RefreshTokenValidity, true
}

// HasRefreshTokenValidity returns a boolean if a field has been set.
func (o *OAuth2ProviderRequest) HasRefreshTokenValidity() bool {
	if o != nil && o.RefreshTokenValidity != nil {
		return true
	}

	return false
}

// SetRefreshTokenValidity gets a reference to the given string and assigns it to the RefreshTokenValidity field.
func (o *OAuth2ProviderRequest) SetRefreshTokenValidity(v string) {
	o.RefreshTokenValidity = &v
}

// GetIncludeClaimsInIdToken returns the IncludeClaimsInIdToken field value if set, zero value otherwise.
func (o *OAuth2ProviderRequest) GetIncludeClaimsInIdToken() bool {
	if o == nil || o.IncludeClaimsInIdToken == nil {
		var ret bool
		return ret
	}
	return *o.IncludeClaimsInIdToken
}

// GetIncludeClaimsInIdTokenOk returns a tuple with the IncludeClaimsInIdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderRequest) GetIncludeClaimsInIdTokenOk() (*bool, bool) {
	if o == nil || o.IncludeClaimsInIdToken == nil {
		return nil, false
	}
	return o.IncludeClaimsInIdToken, true
}

// HasIncludeClaimsInIdToken returns a boolean if a field has been set.
func (o *OAuth2ProviderRequest) HasIncludeClaimsInIdToken() bool {
	if o != nil && o.IncludeClaimsInIdToken != nil {
		return true
	}

	return false
}

// SetIncludeClaimsInIdToken gets a reference to the given bool and assigns it to the IncludeClaimsInIdToken field.
func (o *OAuth2ProviderRequest) SetIncludeClaimsInIdToken(v bool) {
	o.IncludeClaimsInIdToken = &v
}

// GetSigningKey returns the SigningKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ProviderRequest) GetSigningKey() string {
	if o == nil || o.SigningKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SigningKey.Get()
}

// GetSigningKeyOk returns a tuple with the SigningKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ProviderRequest) GetSigningKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SigningKey.Get(), o.SigningKey.IsSet()
}

// HasSigningKey returns a boolean if a field has been set.
func (o *OAuth2ProviderRequest) HasSigningKey() bool {
	if o != nil && o.SigningKey.IsSet() {
		return true
	}

	return false
}

// SetSigningKey gets a reference to the given NullableString and assigns it to the SigningKey field.
func (o *OAuth2ProviderRequest) SetSigningKey(v string) {
	o.SigningKey.Set(&v)
}

// SetSigningKeyNil sets the value for SigningKey to be an explicit nil
func (o *OAuth2ProviderRequest) SetSigningKeyNil() {
	o.SigningKey.Set(nil)
}

// UnsetSigningKey ensures that no value is present for SigningKey, not even an explicit nil
func (o *OAuth2ProviderRequest) UnsetSigningKey() {
	o.SigningKey.Unset()
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2ProviderRequest) GetEncryptionKey() string {
	if o == nil || o.EncryptionKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.EncryptionKey.Get()
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2ProviderRequest) GetEncryptionKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncryptionKey.Get(), o.EncryptionKey.IsSet()
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *OAuth2ProviderRequest) HasEncryptionKey() bool {
	if o != nil && o.EncryptionKey.IsSet() {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given NullableString and assigns it to the EncryptionKey field.
func (o *OAuth2ProviderRequest) SetEncryptionKey(v string) {
	o.EncryptionKey.Set(&v)
}

// SetEncryptionKeyNil sets the value for EncryptionKey to be an explicit nil
func (o *OAuth2ProviderRequest) SetEncryptionKeyNil() {
	o.EncryptionKey.Set(nil)
}

// UnsetEncryptionKey ensures that no value is present for EncryptionKey, not even an explicit nil
func (o *OAuth2ProviderRequest) UnsetEncryptionKey() {
	o.EncryptionKey.Unset()
}

// GetRedirectUris returns the RedirectUris field value
func (o *OAuth2ProviderRequest) GetRedirectUris() []RedirectURIRequest {
	if o == nil {
		var ret []RedirectURIRequest
		return ret
	}

	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderRequest) GetRedirectUrisOk() ([]RedirectURIRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// SetRedirectUris sets field value
func (o *OAuth2ProviderRequest) SetRedirectUris(v []RedirectURIRequest) {
	o.RedirectUris = v
}

// GetSubMode returns the SubMode field value if set, zero value otherwise.
func (o *OAuth2ProviderRequest) GetSubMode() SubModeEnum {
	if o == nil || o.SubMode == nil {
		var ret SubModeEnum
		return ret
	}
	return *o.SubMode
}

// GetSubModeOk returns a tuple with the SubMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderRequest) GetSubModeOk() (*SubModeEnum, bool) {
	if o == nil || o.SubMode == nil {
		return nil, false
	}
	return o.SubMode, true
}

// HasSubMode returns a boolean if a field has been set.
func (o *OAuth2ProviderRequest) HasSubMode() bool {
	if o != nil && o.SubMode != nil {
		return true
	}

	return false
}

// SetSubMode gets a reference to the given SubModeEnum and assigns it to the SubMode field.
func (o *OAuth2ProviderRequest) SetSubMode(v SubModeEnum) {
	o.SubMode = &v
}

// GetIssuerMode returns the IssuerMode field value if set, zero value otherwise.
func (o *OAuth2ProviderRequest) GetIssuerMode() IssuerModeEnum {
	if o == nil || o.IssuerMode == nil {
		var ret IssuerModeEnum
		return ret
	}
	return *o.IssuerMode
}

// GetIssuerModeOk returns a tuple with the IssuerMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderRequest) GetIssuerModeOk() (*IssuerModeEnum, bool) {
	if o == nil || o.IssuerMode == nil {
		return nil, false
	}
	return o.IssuerMode, true
}

// HasIssuerMode returns a boolean if a field has been set.
func (o *OAuth2ProviderRequest) HasIssuerMode() bool {
	if o != nil && o.IssuerMode != nil {
		return true
	}

	return false
}

// SetIssuerMode gets a reference to the given IssuerModeEnum and assigns it to the IssuerMode field.
func (o *OAuth2ProviderRequest) SetIssuerMode(v IssuerModeEnum) {
	o.IssuerMode = &v
}

// GetJwtFederationSources returns the JwtFederationSources field value if set, zero value otherwise.
func (o *OAuth2ProviderRequest) GetJwtFederationSources() []string {
	if o == nil || o.JwtFederationSources == nil {
		var ret []string
		return ret
	}
	return o.JwtFederationSources
}

// GetJwtFederationSourcesOk returns a tuple with the JwtFederationSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderRequest) GetJwtFederationSourcesOk() ([]string, bool) {
	if o == nil || o.JwtFederationSources == nil {
		return nil, false
	}
	return o.JwtFederationSources, true
}

// HasJwtFederationSources returns a boolean if a field has been set.
func (o *OAuth2ProviderRequest) HasJwtFederationSources() bool {
	if o != nil && o.JwtFederationSources != nil {
		return true
	}

	return false
}

// SetJwtFederationSources gets a reference to the given []string and assigns it to the JwtFederationSources field.
func (o *OAuth2ProviderRequest) SetJwtFederationSources(v []string) {
	o.JwtFederationSources = v
}

// GetJwtFederationProviders returns the JwtFederationProviders field value if set, zero value otherwise.
func (o *OAuth2ProviderRequest) GetJwtFederationProviders() []int32 {
	if o == nil || o.JwtFederationProviders == nil {
		var ret []int32
		return ret
	}
	return o.JwtFederationProviders
}

// GetJwtFederationProvidersOk returns a tuple with the JwtFederationProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ProviderRequest) GetJwtFederationProvidersOk() ([]int32, bool) {
	if o == nil || o.JwtFederationProviders == nil {
		return nil, false
	}
	return o.JwtFederationProviders, true
}

// HasJwtFederationProviders returns a boolean if a field has been set.
func (o *OAuth2ProviderRequest) HasJwtFederationProviders() bool {
	if o != nil && o.JwtFederationProviders != nil {
		return true
	}

	return false
}

// SetJwtFederationProviders gets a reference to the given []int32 and assigns it to the JwtFederationProviders field.
func (o *OAuth2ProviderRequest) SetJwtFederationProviders(v []int32) {
	o.JwtFederationProviders = v
}

func (o OAuth2ProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.AuthenticationFlow.IsSet() {
		toSerialize["authentication_flow"] = o.AuthenticationFlow.Get()
	}
	if true {
		toSerialize["authorization_flow"] = o.AuthorizationFlow
	}
	if true {
		toSerialize["invalidation_flow"] = o.InvalidationFlow
	}
	if o.PropertyMappings != nil {
		toSerialize["property_mappings"] = o.PropertyMappings
	}
	if o.ClientType != nil {
		toSerialize["client_type"] = o.ClientType
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.AccessCodeValidity != nil {
		toSerialize["access_code_validity"] = o.AccessCodeValidity
	}
	if o.AccessTokenValidity != nil {
		toSerialize["access_token_validity"] = o.AccessTokenValidity
	}
	if o.RefreshTokenValidity != nil {
		toSerialize["refresh_token_validity"] = o.RefreshTokenValidity
	}
	if o.IncludeClaimsInIdToken != nil {
		toSerialize["include_claims_in_id_token"] = o.IncludeClaimsInIdToken
	}
	if o.SigningKey.IsSet() {
		toSerialize["signing_key"] = o.SigningKey.Get()
	}
	if o.EncryptionKey.IsSet() {
		toSerialize["encryption_key"] = o.EncryptionKey.Get()
	}
	if true {
		toSerialize["redirect_uris"] = o.RedirectUris
	}
	if o.SubMode != nil {
		toSerialize["sub_mode"] = o.SubMode
	}
	if o.IssuerMode != nil {
		toSerialize["issuer_mode"] = o.IssuerMode
	}
	if o.JwtFederationSources != nil {
		toSerialize["jwt_federation_sources"] = o.JwtFederationSources
	}
	if o.JwtFederationProviders != nil {
		toSerialize["jwt_federation_providers"] = o.JwtFederationProviders
	}
	return json.Marshal(toSerialize)
}

type NullableOAuth2ProviderRequest struct {
	value *OAuth2ProviderRequest
	isSet bool
}

func (v NullableOAuth2ProviderRequest) Get() *OAuth2ProviderRequest {
	return v.value
}

func (v *NullableOAuth2ProviderRequest) Set(val *OAuth2ProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ProviderRequest(val *OAuth2ProviderRequest) *NullableOAuth2ProviderRequest {
	return &NullableOAuth2ProviderRequest{value: val, isSet: true}
}

func (v NullableOAuth2ProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
