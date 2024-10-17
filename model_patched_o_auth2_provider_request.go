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

// PatchedOAuth2ProviderRequest OAuth2Provider Serializer
type PatchedOAuth2ProviderRequest struct {
	Name *string `json:"name,omitempty"`
	// Flow used for authentication when the associated application is accessed by an un-authenticated user.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow used when authorizing this provider.
	AuthorizationFlow *string `json:"authorization_flow,omitempty"`
	// Flow used ending the session from a provider.
	InvalidationFlow *string  `json:"invalidation_flow,omitempty"`
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
	EncryptionKey NullableString `json:"encryption_key,omitempty"`
	// Enter each URI on a new line.
	RedirectUris *string `json:"redirect_uris,omitempty"`
	// Configure what data should be used as unique User Identifier. For most cases, the default should be fine.
	SubMode *SubModeEnum `json:"sub_mode,omitempty"`
	// Configure how the issuer field of the ID Token should be filled.
	IssuerMode  *IssuerModeEnum `json:"issuer_mode,omitempty"`
	JwksSources []string        `json:"jwks_sources,omitempty"`
}

// NewPatchedOAuth2ProviderRequest instantiates a new PatchedOAuth2ProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedOAuth2ProviderRequest() *PatchedOAuth2ProviderRequest {
	this := PatchedOAuth2ProviderRequest{}
	return &this
}

// NewPatchedOAuth2ProviderRequestWithDefaults instantiates a new PatchedOAuth2ProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedOAuth2ProviderRequestWithDefaults() *PatchedOAuth2ProviderRequest {
	this := PatchedOAuth2ProviderRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedOAuth2ProviderRequest) SetName(v string) {
	o.Name = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOAuth2ProviderRequest) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOAuth2ProviderRequest) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *PatchedOAuth2ProviderRequest) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *PatchedOAuth2ProviderRequest) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *PatchedOAuth2ProviderRequest) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetAuthorizationFlow returns the AuthorizationFlow field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetAuthorizationFlow() string {
	if o == nil || o.AuthorizationFlow == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil || o.AuthorizationFlow == nil {
		return nil, false
	}
	return o.AuthorizationFlow, true
}

// HasAuthorizationFlow returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasAuthorizationFlow() bool {
	if o != nil && o.AuthorizationFlow != nil {
		return true
	}

	return false
}

// SetAuthorizationFlow gets a reference to the given string and assigns it to the AuthorizationFlow field.
func (o *PatchedOAuth2ProviderRequest) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = &v
}

// GetInvalidationFlow returns the InvalidationFlow field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetInvalidationFlow() string {
	if o == nil || o.InvalidationFlow == nil {
		var ret string
		return ret
	}
	return *o.InvalidationFlow
}

// GetInvalidationFlowOk returns a tuple with the InvalidationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetInvalidationFlowOk() (*string, bool) {
	if o == nil || o.InvalidationFlow == nil {
		return nil, false
	}
	return o.InvalidationFlow, true
}

// HasInvalidationFlow returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasInvalidationFlow() bool {
	if o != nil && o.InvalidationFlow != nil {
		return true
	}

	return false
}

// SetInvalidationFlow gets a reference to the given string and assigns it to the InvalidationFlow field.
func (o *PatchedOAuth2ProviderRequest) SetInvalidationFlow(v string) {
	o.InvalidationFlow = &v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *PatchedOAuth2ProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetClientType returns the ClientType field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetClientType() ClientTypeEnum {
	if o == nil || o.ClientType == nil {
		var ret ClientTypeEnum
		return ret
	}
	return *o.ClientType
}

// GetClientTypeOk returns a tuple with the ClientType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetClientTypeOk() (*ClientTypeEnum, bool) {
	if o == nil || o.ClientType == nil {
		return nil, false
	}
	return o.ClientType, true
}

// HasClientType returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasClientType() bool {
	if o != nil && o.ClientType != nil {
		return true
	}

	return false
}

// SetClientType gets a reference to the given ClientTypeEnum and assigns it to the ClientType field.
func (o *PatchedOAuth2ProviderRequest) SetClientType(v ClientTypeEnum) {
	o.ClientType = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PatchedOAuth2ProviderRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *PatchedOAuth2ProviderRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetAccessCodeValidity returns the AccessCodeValidity field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetAccessCodeValidity() string {
	if o == nil || o.AccessCodeValidity == nil {
		var ret string
		return ret
	}
	return *o.AccessCodeValidity
}

// GetAccessCodeValidityOk returns a tuple with the AccessCodeValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetAccessCodeValidityOk() (*string, bool) {
	if o == nil || o.AccessCodeValidity == nil {
		return nil, false
	}
	return o.AccessCodeValidity, true
}

// HasAccessCodeValidity returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasAccessCodeValidity() bool {
	if o != nil && o.AccessCodeValidity != nil {
		return true
	}

	return false
}

// SetAccessCodeValidity gets a reference to the given string and assigns it to the AccessCodeValidity field.
func (o *PatchedOAuth2ProviderRequest) SetAccessCodeValidity(v string) {
	o.AccessCodeValidity = &v
}

// GetAccessTokenValidity returns the AccessTokenValidity field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetAccessTokenValidity() string {
	if o == nil || o.AccessTokenValidity == nil {
		var ret string
		return ret
	}
	return *o.AccessTokenValidity
}

// GetAccessTokenValidityOk returns a tuple with the AccessTokenValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetAccessTokenValidityOk() (*string, bool) {
	if o == nil || o.AccessTokenValidity == nil {
		return nil, false
	}
	return o.AccessTokenValidity, true
}

// HasAccessTokenValidity returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasAccessTokenValidity() bool {
	if o != nil && o.AccessTokenValidity != nil {
		return true
	}

	return false
}

// SetAccessTokenValidity gets a reference to the given string and assigns it to the AccessTokenValidity field.
func (o *PatchedOAuth2ProviderRequest) SetAccessTokenValidity(v string) {
	o.AccessTokenValidity = &v
}

// GetRefreshTokenValidity returns the RefreshTokenValidity field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetRefreshTokenValidity() string {
	if o == nil || o.RefreshTokenValidity == nil {
		var ret string
		return ret
	}
	return *o.RefreshTokenValidity
}

// GetRefreshTokenValidityOk returns a tuple with the RefreshTokenValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetRefreshTokenValidityOk() (*string, bool) {
	if o == nil || o.RefreshTokenValidity == nil {
		return nil, false
	}
	return o.RefreshTokenValidity, true
}

// HasRefreshTokenValidity returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasRefreshTokenValidity() bool {
	if o != nil && o.RefreshTokenValidity != nil {
		return true
	}

	return false
}

// SetRefreshTokenValidity gets a reference to the given string and assigns it to the RefreshTokenValidity field.
func (o *PatchedOAuth2ProviderRequest) SetRefreshTokenValidity(v string) {
	o.RefreshTokenValidity = &v
}

// GetIncludeClaimsInIdToken returns the IncludeClaimsInIdToken field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetIncludeClaimsInIdToken() bool {
	if o == nil || o.IncludeClaimsInIdToken == nil {
		var ret bool
		return ret
	}
	return *o.IncludeClaimsInIdToken
}

// GetIncludeClaimsInIdTokenOk returns a tuple with the IncludeClaimsInIdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetIncludeClaimsInIdTokenOk() (*bool, bool) {
	if o == nil || o.IncludeClaimsInIdToken == nil {
		return nil, false
	}
	return o.IncludeClaimsInIdToken, true
}

// HasIncludeClaimsInIdToken returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasIncludeClaimsInIdToken() bool {
	if o != nil && o.IncludeClaimsInIdToken != nil {
		return true
	}

	return false
}

// SetIncludeClaimsInIdToken gets a reference to the given bool and assigns it to the IncludeClaimsInIdToken field.
func (o *PatchedOAuth2ProviderRequest) SetIncludeClaimsInIdToken(v bool) {
	o.IncludeClaimsInIdToken = &v
}

// GetSigningKey returns the SigningKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOAuth2ProviderRequest) GetSigningKey() string {
	if o == nil || o.SigningKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SigningKey.Get()
}

// GetSigningKeyOk returns a tuple with the SigningKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOAuth2ProviderRequest) GetSigningKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SigningKey.Get(), o.SigningKey.IsSet()
}

// HasSigningKey returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasSigningKey() bool {
	if o != nil && o.SigningKey.IsSet() {
		return true
	}

	return false
}

// SetSigningKey gets a reference to the given NullableString and assigns it to the SigningKey field.
func (o *PatchedOAuth2ProviderRequest) SetSigningKey(v string) {
	o.SigningKey.Set(&v)
}

// SetSigningKeyNil sets the value for SigningKey to be an explicit nil
func (o *PatchedOAuth2ProviderRequest) SetSigningKeyNil() {
	o.SigningKey.Set(nil)
}

// UnsetSigningKey ensures that no value is present for SigningKey, not even an explicit nil
func (o *PatchedOAuth2ProviderRequest) UnsetSigningKey() {
	o.SigningKey.Unset()
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOAuth2ProviderRequest) GetEncryptionKey() string {
	if o == nil || o.EncryptionKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.EncryptionKey.Get()
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOAuth2ProviderRequest) GetEncryptionKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncryptionKey.Get(), o.EncryptionKey.IsSet()
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasEncryptionKey() bool {
	if o != nil && o.EncryptionKey.IsSet() {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given NullableString and assigns it to the EncryptionKey field.
func (o *PatchedOAuth2ProviderRequest) SetEncryptionKey(v string) {
	o.EncryptionKey.Set(&v)
}

// SetEncryptionKeyNil sets the value for EncryptionKey to be an explicit nil
func (o *PatchedOAuth2ProviderRequest) SetEncryptionKeyNil() {
	o.EncryptionKey.Set(nil)
}

// UnsetEncryptionKey ensures that no value is present for EncryptionKey, not even an explicit nil
func (o *PatchedOAuth2ProviderRequest) UnsetEncryptionKey() {
	o.EncryptionKey.Unset()
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetRedirectUris() string {
	if o == nil || o.RedirectUris == nil {
		var ret string
		return ret
	}
	return *o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetRedirectUrisOk() (*string, bool) {
	if o == nil || o.RedirectUris == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasRedirectUris() bool {
	if o != nil && o.RedirectUris != nil {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given string and assigns it to the RedirectUris field.
func (o *PatchedOAuth2ProviderRequest) SetRedirectUris(v string) {
	o.RedirectUris = &v
}

// GetSubMode returns the SubMode field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetSubMode() SubModeEnum {
	if o == nil || o.SubMode == nil {
		var ret SubModeEnum
		return ret
	}
	return *o.SubMode
}

// GetSubModeOk returns a tuple with the SubMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetSubModeOk() (*SubModeEnum, bool) {
	if o == nil || o.SubMode == nil {
		return nil, false
	}
	return o.SubMode, true
}

// HasSubMode returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasSubMode() bool {
	if o != nil && o.SubMode != nil {
		return true
	}

	return false
}

// SetSubMode gets a reference to the given SubModeEnum and assigns it to the SubMode field.
func (o *PatchedOAuth2ProviderRequest) SetSubMode(v SubModeEnum) {
	o.SubMode = &v
}

// GetIssuerMode returns the IssuerMode field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetIssuerMode() IssuerModeEnum {
	if o == nil || o.IssuerMode == nil {
		var ret IssuerModeEnum
		return ret
	}
	return *o.IssuerMode
}

// GetIssuerModeOk returns a tuple with the IssuerMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetIssuerModeOk() (*IssuerModeEnum, bool) {
	if o == nil || o.IssuerMode == nil {
		return nil, false
	}
	return o.IssuerMode, true
}

// HasIssuerMode returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasIssuerMode() bool {
	if o != nil && o.IssuerMode != nil {
		return true
	}

	return false
}

// SetIssuerMode gets a reference to the given IssuerModeEnum and assigns it to the IssuerMode field.
func (o *PatchedOAuth2ProviderRequest) SetIssuerMode(v IssuerModeEnum) {
	o.IssuerMode = &v
}

// GetJwksSources returns the JwksSources field value if set, zero value otherwise.
func (o *PatchedOAuth2ProviderRequest) GetJwksSources() []string {
	if o == nil || o.JwksSources == nil {
		var ret []string
		return ret
	}
	return o.JwksSources
}

// GetJwksSourcesOk returns a tuple with the JwksSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuth2ProviderRequest) GetJwksSourcesOk() ([]string, bool) {
	if o == nil || o.JwksSources == nil {
		return nil, false
	}
	return o.JwksSources, true
}

// HasJwksSources returns a boolean if a field has been set.
func (o *PatchedOAuth2ProviderRequest) HasJwksSources() bool {
	if o != nil && o.JwksSources != nil {
		return true
	}

	return false
}

// SetJwksSources gets a reference to the given []string and assigns it to the JwksSources field.
func (o *PatchedOAuth2ProviderRequest) SetJwksSources(v []string) {
	o.JwksSources = v
}

func (o PatchedOAuth2ProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.AuthenticationFlow.IsSet() {
		toSerialize["authentication_flow"] = o.AuthenticationFlow.Get()
	}
	if o.AuthorizationFlow != nil {
		toSerialize["authorization_flow"] = o.AuthorizationFlow
	}
	if o.InvalidationFlow != nil {
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
	if o.RedirectUris != nil {
		toSerialize["redirect_uris"] = o.RedirectUris
	}
	if o.SubMode != nil {
		toSerialize["sub_mode"] = o.SubMode
	}
	if o.IssuerMode != nil {
		toSerialize["issuer_mode"] = o.IssuerMode
	}
	if o.JwksSources != nil {
		toSerialize["jwks_sources"] = o.JwksSources
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedOAuth2ProviderRequest struct {
	value *PatchedOAuth2ProviderRequest
	isSet bool
}

func (v NullablePatchedOAuth2ProviderRequest) Get() *PatchedOAuth2ProviderRequest {
	return v.value
}

func (v *NullablePatchedOAuth2ProviderRequest) Set(val *PatchedOAuth2ProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedOAuth2ProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedOAuth2ProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedOAuth2ProviderRequest(val *PatchedOAuth2ProviderRequest) *NullablePatchedOAuth2ProviderRequest {
	return &NullablePatchedOAuth2ProviderRequest{value: val, isSet: true}
}

func (v NullablePatchedOAuth2ProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedOAuth2ProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
