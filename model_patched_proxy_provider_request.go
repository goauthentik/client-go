/*
authentik

Making authentication simple.

API version: 2024.6.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedProxyProviderRequest ProxyProvider Serializer
type PatchedProxyProviderRequest struct {
	Name *string `json:"name,omitempty"`
	// Flow used for authentication when the associated application is accessed by an un-authenticated user.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow used when authorizing this provider.
	AuthorizationFlow *string  `json:"authorization_flow,omitempty"`
	PropertyMappings  []string `json:"property_mappings,omitempty"`
	InternalHost      *string  `json:"internal_host,omitempty"`
	ExternalHost      *string  `json:"external_host,omitempty"`
	// Validate SSL Certificates of upstream servers
	InternalHostSslValidation *bool          `json:"internal_host_ssl_validation,omitempty"`
	Certificate               NullableString `json:"certificate,omitempty"`
	// Regular expressions for which authentication is not required. Each new line is interpreted as a new Regular Expression.
	SkipPathRegex *string `json:"skip_path_regex,omitempty"`
	// Set a custom HTTP-Basic Authentication header based on values from authentik.
	BasicAuthEnabled *bool `json:"basic_auth_enabled,omitempty"`
	// User/Group Attribute used for the password part of the HTTP-Basic Header.
	BasicAuthPasswordAttribute *string `json:"basic_auth_password_attribute,omitempty"`
	// User/Group Attribute used for the user part of the HTTP-Basic Header. If not set, the user's Email address is used.
	BasicAuthUserAttribute *string `json:"basic_auth_user_attribute,omitempty"`
	// Enable support for forwardAuth in traefik and nginx auth_request. Exclusive with internal_host.
	Mode *ProxyMode `json:"mode,omitempty"`
	// When enabled, this provider will intercept the authorization header and authenticate requests based on its value.
	InterceptHeaderAuth *bool    `json:"intercept_header_auth,omitempty"`
	CookieDomain        *string  `json:"cookie_domain,omitempty"`
	JwksSources         []string `json:"jwks_sources,omitempty"`
	// Tokens not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	AccessTokenValidity *string `json:"access_token_validity,omitempty"`
	// Tokens not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	RefreshTokenValidity *string `json:"refresh_token_validity,omitempty"`
}

// NewPatchedProxyProviderRequest instantiates a new PatchedProxyProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedProxyProviderRequest() *PatchedProxyProviderRequest {
	this := PatchedProxyProviderRequest{}
	return &this
}

// NewPatchedProxyProviderRequestWithDefaults instantiates a new PatchedProxyProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedProxyProviderRequestWithDefaults() *PatchedProxyProviderRequest {
	this := PatchedProxyProviderRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedProxyProviderRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProxyProviderRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedProxyProviderRequest) SetName(v string) {
	o.Name = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedProxyProviderRequest) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedProxyProviderRequest) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *PatchedProxyProviderRequest) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *PatchedProxyProviderRequest) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *PatchedProxyProviderRequest) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetAuthorizationFlow returns the AuthorizationFlow field value if set, zero value otherwise.
func (o *PatchedProxyProviderRequest) GetAuthorizationFlow() string {
	if o == nil || o.AuthorizationFlow == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProxyProviderRequest) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil || o.AuthorizationFlow == nil {
		return nil, false
	}
	return o.AuthorizationFlow, true
}

// HasAuthorizationFlow returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasAuthorizationFlow() bool {
	if o != nil && o.AuthorizationFlow != nil {
		return true
	}

	return false
}

// SetAuthorizationFlow gets a reference to the given string and assigns it to the AuthorizationFlow field.
func (o *PatchedProxyProviderRequest) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = &v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *PatchedProxyProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProxyProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *PatchedProxyProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetInternalHost returns the InternalHost field value if set, zero value otherwise.
func (o *PatchedProxyProviderRequest) GetInternalHost() string {
	if o == nil || o.InternalHost == nil {
		var ret string
		return ret
	}
	return *o.InternalHost
}

// GetInternalHostOk returns a tuple with the InternalHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProxyProviderRequest) GetInternalHostOk() (*string, bool) {
	if o == nil || o.InternalHost == nil {
		return nil, false
	}
	return o.InternalHost, true
}

// HasInternalHost returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasInternalHost() bool {
	if o != nil && o.InternalHost != nil {
		return true
	}

	return false
}

// SetInternalHost gets a reference to the given string and assigns it to the InternalHost field.
func (o *PatchedProxyProviderRequest) SetInternalHost(v string) {
	o.InternalHost = &v
}

// GetExternalHost returns the ExternalHost field value if set, zero value otherwise.
func (o *PatchedProxyProviderRequest) GetExternalHost() string {
	if o == nil || o.ExternalHost == nil {
		var ret string
		return ret
	}
	return *o.ExternalHost
}

// GetExternalHostOk returns a tuple with the ExternalHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProxyProviderRequest) GetExternalHostOk() (*string, bool) {
	if o == nil || o.ExternalHost == nil {
		return nil, false
	}
	return o.ExternalHost, true
}

// HasExternalHost returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasExternalHost() bool {
	if o != nil && o.ExternalHost != nil {
		return true
	}

	return false
}

// SetExternalHost gets a reference to the given string and assigns it to the ExternalHost field.
func (o *PatchedProxyProviderRequest) SetExternalHost(v string) {
	o.ExternalHost = &v
}

// GetInternalHostSslValidation returns the InternalHostSslValidation field value if set, zero value otherwise.
func (o *PatchedProxyProviderRequest) GetInternalHostSslValidation() bool {
	if o == nil || o.InternalHostSslValidation == nil {
		var ret bool
		return ret
	}
	return *o.InternalHostSslValidation
}

// GetInternalHostSslValidationOk returns a tuple with the InternalHostSslValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProxyProviderRequest) GetInternalHostSslValidationOk() (*bool, bool) {
	if o == nil || o.InternalHostSslValidation == nil {
		return nil, false
	}
	return o.InternalHostSslValidation, true
}

// HasInternalHostSslValidation returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasInternalHostSslValidation() bool {
	if o != nil && o.InternalHostSslValidation != nil {
		return true
	}

	return false
}

// SetInternalHostSslValidation gets a reference to the given bool and assigns it to the InternalHostSslValidation field.
func (o *PatchedProxyProviderRequest) SetInternalHostSslValidation(v bool) {
	o.InternalHostSslValidation = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedProxyProviderRequest) GetCertificate() string {
	if o == nil || o.Certificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedProxyProviderRequest) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *PatchedProxyProviderRequest) SetCertificate(v string) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *PatchedProxyProviderRequest) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *PatchedProxyProviderRequest) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetSkipPathRegex returns the SkipPathRegex field value if set, zero value otherwise.
func (o *PatchedProxyProviderRequest) GetSkipPathRegex() string {
	if o == nil || o.SkipPathRegex == nil {
		var ret string
		return ret
	}
	return *o.SkipPathRegex
}

// GetSkipPathRegexOk returns a tuple with the SkipPathRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProxyProviderRequest) GetSkipPathRegexOk() (*string, bool) {
	if o == nil || o.SkipPathRegex == nil {
		return nil, false
	}
	return o.SkipPathRegex, true
}

// HasSkipPathRegex returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasSkipPathRegex() bool {
	if o != nil && o.SkipPathRegex != nil {
		return true
	}

	return false
}

// SetSkipPathRegex gets a reference to the given string and assigns it to the SkipPathRegex field.
func (o *PatchedProxyProviderRequest) SetSkipPathRegex(v string) {
	o.SkipPathRegex = &v
}

// GetBasicAuthEnabled returns the BasicAuthEnabled field value if set, zero value otherwise.
func (o *PatchedProxyProviderRequest) GetBasicAuthEnabled() bool {
	if o == nil || o.BasicAuthEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BasicAuthEnabled
}

// GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProxyProviderRequest) GetBasicAuthEnabledOk() (*bool, bool) {
	if o == nil || o.BasicAuthEnabled == nil {
		return nil, false
	}
	return o.BasicAuthEnabled, true
}

// HasBasicAuthEnabled returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasBasicAuthEnabled() bool {
	if o != nil && o.BasicAuthEnabled != nil {
		return true
	}

	return false
}

// SetBasicAuthEnabled gets a reference to the given bool and assigns it to the BasicAuthEnabled field.
func (o *PatchedProxyProviderRequest) SetBasicAuthEnabled(v bool) {
	o.BasicAuthEnabled = &v
}

// GetBasicAuthPasswordAttribute returns the BasicAuthPasswordAttribute field value if set, zero value otherwise.
func (o *PatchedProxyProviderRequest) GetBasicAuthPasswordAttribute() string {
	if o == nil || o.BasicAuthPasswordAttribute == nil {
		var ret string
		return ret
	}
	return *o.BasicAuthPasswordAttribute
}

// GetBasicAuthPasswordAttributeOk returns a tuple with the BasicAuthPasswordAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProxyProviderRequest) GetBasicAuthPasswordAttributeOk() (*string, bool) {
	if o == nil || o.BasicAuthPasswordAttribute == nil {
		return nil, false
	}
	return o.BasicAuthPasswordAttribute, true
}

// HasBasicAuthPasswordAttribute returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasBasicAuthPasswordAttribute() bool {
	if o != nil && o.BasicAuthPasswordAttribute != nil {
		return true
	}

	return false
}

// SetBasicAuthPasswordAttribute gets a reference to the given string and assigns it to the BasicAuthPasswordAttribute field.
func (o *PatchedProxyProviderRequest) SetBasicAuthPasswordAttribute(v string) {
	o.BasicAuthPasswordAttribute = &v
}

// GetBasicAuthUserAttribute returns the BasicAuthUserAttribute field value if set, zero value otherwise.
func (o *PatchedProxyProviderRequest) GetBasicAuthUserAttribute() string {
	if o == nil || o.BasicAuthUserAttribute == nil {
		var ret string
		return ret
	}
	return *o.BasicAuthUserAttribute
}

// GetBasicAuthUserAttributeOk returns a tuple with the BasicAuthUserAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProxyProviderRequest) GetBasicAuthUserAttributeOk() (*string, bool) {
	if o == nil || o.BasicAuthUserAttribute == nil {
		return nil, false
	}
	return o.BasicAuthUserAttribute, true
}

// HasBasicAuthUserAttribute returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasBasicAuthUserAttribute() bool {
	if o != nil && o.BasicAuthUserAttribute != nil {
		return true
	}

	return false
}

// SetBasicAuthUserAttribute gets a reference to the given string and assigns it to the BasicAuthUserAttribute field.
func (o *PatchedProxyProviderRequest) SetBasicAuthUserAttribute(v string) {
	o.BasicAuthUserAttribute = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *PatchedProxyProviderRequest) GetMode() ProxyMode {
	if o == nil || o.Mode == nil {
		var ret ProxyMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProxyProviderRequest) GetModeOk() (*ProxyMode, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given ProxyMode and assigns it to the Mode field.
func (o *PatchedProxyProviderRequest) SetMode(v ProxyMode) {
	o.Mode = &v
}

// GetInterceptHeaderAuth returns the InterceptHeaderAuth field value if set, zero value otherwise.
func (o *PatchedProxyProviderRequest) GetInterceptHeaderAuth() bool {
	if o == nil || o.InterceptHeaderAuth == nil {
		var ret bool
		return ret
	}
	return *o.InterceptHeaderAuth
}

// GetInterceptHeaderAuthOk returns a tuple with the InterceptHeaderAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProxyProviderRequest) GetInterceptHeaderAuthOk() (*bool, bool) {
	if o == nil || o.InterceptHeaderAuth == nil {
		return nil, false
	}
	return o.InterceptHeaderAuth, true
}

// HasInterceptHeaderAuth returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasInterceptHeaderAuth() bool {
	if o != nil && o.InterceptHeaderAuth != nil {
		return true
	}

	return false
}

// SetInterceptHeaderAuth gets a reference to the given bool and assigns it to the InterceptHeaderAuth field.
func (o *PatchedProxyProviderRequest) SetInterceptHeaderAuth(v bool) {
	o.InterceptHeaderAuth = &v
}

// GetCookieDomain returns the CookieDomain field value if set, zero value otherwise.
func (o *PatchedProxyProviderRequest) GetCookieDomain() string {
	if o == nil || o.CookieDomain == nil {
		var ret string
		return ret
	}
	return *o.CookieDomain
}

// GetCookieDomainOk returns a tuple with the CookieDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProxyProviderRequest) GetCookieDomainOk() (*string, bool) {
	if o == nil || o.CookieDomain == nil {
		return nil, false
	}
	return o.CookieDomain, true
}

// HasCookieDomain returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasCookieDomain() bool {
	if o != nil && o.CookieDomain != nil {
		return true
	}

	return false
}

// SetCookieDomain gets a reference to the given string and assigns it to the CookieDomain field.
func (o *PatchedProxyProviderRequest) SetCookieDomain(v string) {
	o.CookieDomain = &v
}

// GetJwksSources returns the JwksSources field value if set, zero value otherwise.
func (o *PatchedProxyProviderRequest) GetJwksSources() []string {
	if o == nil || o.JwksSources == nil {
		var ret []string
		return ret
	}
	return o.JwksSources
}

// GetJwksSourcesOk returns a tuple with the JwksSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProxyProviderRequest) GetJwksSourcesOk() ([]string, bool) {
	if o == nil || o.JwksSources == nil {
		return nil, false
	}
	return o.JwksSources, true
}

// HasJwksSources returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasJwksSources() bool {
	if o != nil && o.JwksSources != nil {
		return true
	}

	return false
}

// SetJwksSources gets a reference to the given []string and assigns it to the JwksSources field.
func (o *PatchedProxyProviderRequest) SetJwksSources(v []string) {
	o.JwksSources = v
}

// GetAccessTokenValidity returns the AccessTokenValidity field value if set, zero value otherwise.
func (o *PatchedProxyProviderRequest) GetAccessTokenValidity() string {
	if o == nil || o.AccessTokenValidity == nil {
		var ret string
		return ret
	}
	return *o.AccessTokenValidity
}

// GetAccessTokenValidityOk returns a tuple with the AccessTokenValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProxyProviderRequest) GetAccessTokenValidityOk() (*string, bool) {
	if o == nil || o.AccessTokenValidity == nil {
		return nil, false
	}
	return o.AccessTokenValidity, true
}

// HasAccessTokenValidity returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasAccessTokenValidity() bool {
	if o != nil && o.AccessTokenValidity != nil {
		return true
	}

	return false
}

// SetAccessTokenValidity gets a reference to the given string and assigns it to the AccessTokenValidity field.
func (o *PatchedProxyProviderRequest) SetAccessTokenValidity(v string) {
	o.AccessTokenValidity = &v
}

// GetRefreshTokenValidity returns the RefreshTokenValidity field value if set, zero value otherwise.
func (o *PatchedProxyProviderRequest) GetRefreshTokenValidity() string {
	if o == nil || o.RefreshTokenValidity == nil {
		var ret string
		return ret
	}
	return *o.RefreshTokenValidity
}

// GetRefreshTokenValidityOk returns a tuple with the RefreshTokenValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProxyProviderRequest) GetRefreshTokenValidityOk() (*string, bool) {
	if o == nil || o.RefreshTokenValidity == nil {
		return nil, false
	}
	return o.RefreshTokenValidity, true
}

// HasRefreshTokenValidity returns a boolean if a field has been set.
func (o *PatchedProxyProviderRequest) HasRefreshTokenValidity() bool {
	if o != nil && o.RefreshTokenValidity != nil {
		return true
	}

	return false
}

// SetRefreshTokenValidity gets a reference to the given string and assigns it to the RefreshTokenValidity field.
func (o *PatchedProxyProviderRequest) SetRefreshTokenValidity(v string) {
	o.RefreshTokenValidity = &v
}

func (o PatchedProxyProviderRequest) MarshalJSON() ([]byte, error) {
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
	if o.PropertyMappings != nil {
		toSerialize["property_mappings"] = o.PropertyMappings
	}
	if o.InternalHost != nil {
		toSerialize["internal_host"] = o.InternalHost
	}
	if o.ExternalHost != nil {
		toSerialize["external_host"] = o.ExternalHost
	}
	if o.InternalHostSslValidation != nil {
		toSerialize["internal_host_ssl_validation"] = o.InternalHostSslValidation
	}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.SkipPathRegex != nil {
		toSerialize["skip_path_regex"] = o.SkipPathRegex
	}
	if o.BasicAuthEnabled != nil {
		toSerialize["basic_auth_enabled"] = o.BasicAuthEnabled
	}
	if o.BasicAuthPasswordAttribute != nil {
		toSerialize["basic_auth_password_attribute"] = o.BasicAuthPasswordAttribute
	}
	if o.BasicAuthUserAttribute != nil {
		toSerialize["basic_auth_user_attribute"] = o.BasicAuthUserAttribute
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.InterceptHeaderAuth != nil {
		toSerialize["intercept_header_auth"] = o.InterceptHeaderAuth
	}
	if o.CookieDomain != nil {
		toSerialize["cookie_domain"] = o.CookieDomain
	}
	if o.JwksSources != nil {
		toSerialize["jwks_sources"] = o.JwksSources
	}
	if o.AccessTokenValidity != nil {
		toSerialize["access_token_validity"] = o.AccessTokenValidity
	}
	if o.RefreshTokenValidity != nil {
		toSerialize["refresh_token_validity"] = o.RefreshTokenValidity
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedProxyProviderRequest struct {
	value *PatchedProxyProviderRequest
	isSet bool
}

func (v NullablePatchedProxyProviderRequest) Get() *PatchedProxyProviderRequest {
	return v.value
}

func (v *NullablePatchedProxyProviderRequest) Set(val *PatchedProxyProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedProxyProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedProxyProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedProxyProviderRequest(val *PatchedProxyProviderRequest) *NullablePatchedProxyProviderRequest {
	return &NullablePatchedProxyProviderRequest{value: val, isSet: true}
}

func (v NullablePatchedProxyProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedProxyProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
