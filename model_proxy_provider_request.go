/*
authentik

Making authentication simple.

API version: 2023.3.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ProxyProviderRequest ProxyProvider Serializer
type ProxyProviderRequest struct {
	Name string `json:"name"`
	// Flow used when authorizing this provider.
	AuthorizationFlow NullableString `json:"authorization_flow,omitempty"`
	PropertyMappings  []string       `json:"property_mappings,omitempty"`
	InternalHost      *string        `json:"internal_host,omitempty"`
	ExternalHost      string         `json:"external_host"`
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
	// Enable support for forwardAuth in traefik and nginx auth_request. Exclusive with internal_host.  * `proxy` - Proxy * `forward_single` - Forward Single * `forward_domain` - Forward Domain
	Mode NullableProxyMode `json:"mode,omitempty"`
	// When enabled, this provider will intercept the authorization header and authenticate requests based on its value.
	InterceptHeaderAuth *bool    `json:"intercept_header_auth,omitempty"`
	CookieDomain        *string  `json:"cookie_domain,omitempty"`
	JwksSources         []string `json:"jwks_sources,omitempty"`
	// Tokens not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	AccessTokenValidity *string `json:"access_token_validity,omitempty"`
	// Tokens not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	RefreshTokenValidity *string `json:"refresh_token_validity,omitempty"`
}

// NewProxyProviderRequest instantiates a new ProxyProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyProviderRequest(name string, externalHost string) *ProxyProviderRequest {
	this := ProxyProviderRequest{}
	this.Name = name
	this.ExternalHost = externalHost
	return &this
}

// NewProxyProviderRequestWithDefaults instantiates a new ProxyProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyProviderRequestWithDefaults() *ProxyProviderRequest {
	this := ProxyProviderRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ProxyProviderRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProxyProviderRequest) SetName(v string) {
	o.Name = v
}

// GetAuthorizationFlow returns the AuthorizationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxyProviderRequest) GetAuthorizationFlow() string {
	if o == nil || o.AuthorizationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationFlow.Get()
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxyProviderRequest) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationFlow.Get(), o.AuthorizationFlow.IsSet()
}

// HasAuthorizationFlow returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasAuthorizationFlow() bool {
	if o != nil && o.AuthorizationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthorizationFlow gets a reference to the given NullableString and assigns it to the AuthorizationFlow field.
func (o *ProxyProviderRequest) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow.Set(&v)
}

// SetAuthorizationFlowNil sets the value for AuthorizationFlow to be an explicit nil
func (o *ProxyProviderRequest) SetAuthorizationFlowNil() {
	o.AuthorizationFlow.Set(nil)
}

// UnsetAuthorizationFlow ensures that no value is present for AuthorizationFlow, not even an explicit nil
func (o *ProxyProviderRequest) UnsetAuthorizationFlow() {
	o.AuthorizationFlow.Unset()
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *ProxyProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetInternalHost returns the InternalHost field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetInternalHost() string {
	if o == nil || o.InternalHost == nil {
		var ret string
		return ret
	}
	return *o.InternalHost
}

// GetInternalHostOk returns a tuple with the InternalHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetInternalHostOk() (*string, bool) {
	if o == nil || o.InternalHost == nil {
		return nil, false
	}
	return o.InternalHost, true
}

// HasInternalHost returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasInternalHost() bool {
	if o != nil && o.InternalHost != nil {
		return true
	}

	return false
}

// SetInternalHost gets a reference to the given string and assigns it to the InternalHost field.
func (o *ProxyProviderRequest) SetInternalHost(v string) {
	o.InternalHost = &v
}

// GetExternalHost returns the ExternalHost field value
func (o *ProxyProviderRequest) GetExternalHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalHost
}

// GetExternalHostOk returns a tuple with the ExternalHost field value
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetExternalHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalHost, true
}

// SetExternalHost sets field value
func (o *ProxyProviderRequest) SetExternalHost(v string) {
	o.ExternalHost = v
}

// GetInternalHostSslValidation returns the InternalHostSslValidation field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetInternalHostSslValidation() bool {
	if o == nil || o.InternalHostSslValidation == nil {
		var ret bool
		return ret
	}
	return *o.InternalHostSslValidation
}

// GetInternalHostSslValidationOk returns a tuple with the InternalHostSslValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetInternalHostSslValidationOk() (*bool, bool) {
	if o == nil || o.InternalHostSslValidation == nil {
		return nil, false
	}
	return o.InternalHostSslValidation, true
}

// HasInternalHostSslValidation returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasInternalHostSslValidation() bool {
	if o != nil && o.InternalHostSslValidation != nil {
		return true
	}

	return false
}

// SetInternalHostSslValidation gets a reference to the given bool and assigns it to the InternalHostSslValidation field.
func (o *ProxyProviderRequest) SetInternalHostSslValidation(v bool) {
	o.InternalHostSslValidation = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxyProviderRequest) GetCertificate() string {
	if o == nil || o.Certificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxyProviderRequest) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *ProxyProviderRequest) SetCertificate(v string) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *ProxyProviderRequest) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *ProxyProviderRequest) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetSkipPathRegex returns the SkipPathRegex field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetSkipPathRegex() string {
	if o == nil || o.SkipPathRegex == nil {
		var ret string
		return ret
	}
	return *o.SkipPathRegex
}

// GetSkipPathRegexOk returns a tuple with the SkipPathRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetSkipPathRegexOk() (*string, bool) {
	if o == nil || o.SkipPathRegex == nil {
		return nil, false
	}
	return o.SkipPathRegex, true
}

// HasSkipPathRegex returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasSkipPathRegex() bool {
	if o != nil && o.SkipPathRegex != nil {
		return true
	}

	return false
}

// SetSkipPathRegex gets a reference to the given string and assigns it to the SkipPathRegex field.
func (o *ProxyProviderRequest) SetSkipPathRegex(v string) {
	o.SkipPathRegex = &v
}

// GetBasicAuthEnabled returns the BasicAuthEnabled field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetBasicAuthEnabled() bool {
	if o == nil || o.BasicAuthEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BasicAuthEnabled
}

// GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetBasicAuthEnabledOk() (*bool, bool) {
	if o == nil || o.BasicAuthEnabled == nil {
		return nil, false
	}
	return o.BasicAuthEnabled, true
}

// HasBasicAuthEnabled returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasBasicAuthEnabled() bool {
	if o != nil && o.BasicAuthEnabled != nil {
		return true
	}

	return false
}

// SetBasicAuthEnabled gets a reference to the given bool and assigns it to the BasicAuthEnabled field.
func (o *ProxyProviderRequest) SetBasicAuthEnabled(v bool) {
	o.BasicAuthEnabled = &v
}

// GetBasicAuthPasswordAttribute returns the BasicAuthPasswordAttribute field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetBasicAuthPasswordAttribute() string {
	if o == nil || o.BasicAuthPasswordAttribute == nil {
		var ret string
		return ret
	}
	return *o.BasicAuthPasswordAttribute
}

// GetBasicAuthPasswordAttributeOk returns a tuple with the BasicAuthPasswordAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetBasicAuthPasswordAttributeOk() (*string, bool) {
	if o == nil || o.BasicAuthPasswordAttribute == nil {
		return nil, false
	}
	return o.BasicAuthPasswordAttribute, true
}

// HasBasicAuthPasswordAttribute returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasBasicAuthPasswordAttribute() bool {
	if o != nil && o.BasicAuthPasswordAttribute != nil {
		return true
	}

	return false
}

// SetBasicAuthPasswordAttribute gets a reference to the given string and assigns it to the BasicAuthPasswordAttribute field.
func (o *ProxyProviderRequest) SetBasicAuthPasswordAttribute(v string) {
	o.BasicAuthPasswordAttribute = &v
}

// GetBasicAuthUserAttribute returns the BasicAuthUserAttribute field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetBasicAuthUserAttribute() string {
	if o == nil || o.BasicAuthUserAttribute == nil {
		var ret string
		return ret
	}
	return *o.BasicAuthUserAttribute
}

// GetBasicAuthUserAttributeOk returns a tuple with the BasicAuthUserAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetBasicAuthUserAttributeOk() (*string, bool) {
	if o == nil || o.BasicAuthUserAttribute == nil {
		return nil, false
	}
	return o.BasicAuthUserAttribute, true
}

// HasBasicAuthUserAttribute returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasBasicAuthUserAttribute() bool {
	if o != nil && o.BasicAuthUserAttribute != nil {
		return true
	}

	return false
}

// SetBasicAuthUserAttribute gets a reference to the given string and assigns it to the BasicAuthUserAttribute field.
func (o *ProxyProviderRequest) SetBasicAuthUserAttribute(v string) {
	o.BasicAuthUserAttribute = &v
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxyProviderRequest) GetMode() ProxyMode {
	if o == nil || o.Mode.Get() == nil {
		var ret ProxyMode
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxyProviderRequest) GetModeOk() (*ProxyMode, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasMode() bool {
	if o != nil && o.Mode.IsSet() {
		return true
	}

	return false
}

// SetMode gets a reference to the given NullableProxyMode and assigns it to the Mode field.
func (o *ProxyProviderRequest) SetMode(v ProxyMode) {
	o.Mode.Set(&v)
}

// SetModeNil sets the value for Mode to be an explicit nil
func (o *ProxyProviderRequest) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil
func (o *ProxyProviderRequest) UnsetMode() {
	o.Mode.Unset()
}

// GetInterceptHeaderAuth returns the InterceptHeaderAuth field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetInterceptHeaderAuth() bool {
	if o == nil || o.InterceptHeaderAuth == nil {
		var ret bool
		return ret
	}
	return *o.InterceptHeaderAuth
}

// GetInterceptHeaderAuthOk returns a tuple with the InterceptHeaderAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetInterceptHeaderAuthOk() (*bool, bool) {
	if o == nil || o.InterceptHeaderAuth == nil {
		return nil, false
	}
	return o.InterceptHeaderAuth, true
}

// HasInterceptHeaderAuth returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasInterceptHeaderAuth() bool {
	if o != nil && o.InterceptHeaderAuth != nil {
		return true
	}

	return false
}

// SetInterceptHeaderAuth gets a reference to the given bool and assigns it to the InterceptHeaderAuth field.
func (o *ProxyProviderRequest) SetInterceptHeaderAuth(v bool) {
	o.InterceptHeaderAuth = &v
}

// GetCookieDomain returns the CookieDomain field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetCookieDomain() string {
	if o == nil || o.CookieDomain == nil {
		var ret string
		return ret
	}
	return *o.CookieDomain
}

// GetCookieDomainOk returns a tuple with the CookieDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetCookieDomainOk() (*string, bool) {
	if o == nil || o.CookieDomain == nil {
		return nil, false
	}
	return o.CookieDomain, true
}

// HasCookieDomain returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasCookieDomain() bool {
	if o != nil && o.CookieDomain != nil {
		return true
	}

	return false
}

// SetCookieDomain gets a reference to the given string and assigns it to the CookieDomain field.
func (o *ProxyProviderRequest) SetCookieDomain(v string) {
	o.CookieDomain = &v
}

// GetJwksSources returns the JwksSources field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetJwksSources() []string {
	if o == nil || o.JwksSources == nil {
		var ret []string
		return ret
	}
	return o.JwksSources
}

// GetJwksSourcesOk returns a tuple with the JwksSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetJwksSourcesOk() ([]string, bool) {
	if o == nil || o.JwksSources == nil {
		return nil, false
	}
	return o.JwksSources, true
}

// HasJwksSources returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasJwksSources() bool {
	if o != nil && o.JwksSources != nil {
		return true
	}

	return false
}

// SetJwksSources gets a reference to the given []string and assigns it to the JwksSources field.
func (o *ProxyProviderRequest) SetJwksSources(v []string) {
	o.JwksSources = v
}

// GetAccessTokenValidity returns the AccessTokenValidity field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetAccessTokenValidity() string {
	if o == nil || o.AccessTokenValidity == nil {
		var ret string
		return ret
	}
	return *o.AccessTokenValidity
}

// GetAccessTokenValidityOk returns a tuple with the AccessTokenValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetAccessTokenValidityOk() (*string, bool) {
	if o == nil || o.AccessTokenValidity == nil {
		return nil, false
	}
	return o.AccessTokenValidity, true
}

// HasAccessTokenValidity returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasAccessTokenValidity() bool {
	if o != nil && o.AccessTokenValidity != nil {
		return true
	}

	return false
}

// SetAccessTokenValidity gets a reference to the given string and assigns it to the AccessTokenValidity field.
func (o *ProxyProviderRequest) SetAccessTokenValidity(v string) {
	o.AccessTokenValidity = &v
}

// GetRefreshTokenValidity returns the RefreshTokenValidity field value if set, zero value otherwise.
func (o *ProxyProviderRequest) GetRefreshTokenValidity() string {
	if o == nil || o.RefreshTokenValidity == nil {
		var ret string
		return ret
	}
	return *o.RefreshTokenValidity
}

// GetRefreshTokenValidityOk returns a tuple with the RefreshTokenValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProviderRequest) GetRefreshTokenValidityOk() (*string, bool) {
	if o == nil || o.RefreshTokenValidity == nil {
		return nil, false
	}
	return o.RefreshTokenValidity, true
}

// HasRefreshTokenValidity returns a boolean if a field has been set.
func (o *ProxyProviderRequest) HasRefreshTokenValidity() bool {
	if o != nil && o.RefreshTokenValidity != nil {
		return true
	}

	return false
}

// SetRefreshTokenValidity gets a reference to the given string and assigns it to the RefreshTokenValidity field.
func (o *ProxyProviderRequest) SetRefreshTokenValidity(v string) {
	o.RefreshTokenValidity = &v
}

func (o ProxyProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.AuthorizationFlow.IsSet() {
		toSerialize["authorization_flow"] = o.AuthorizationFlow.Get()
	}
	if o.PropertyMappings != nil {
		toSerialize["property_mappings"] = o.PropertyMappings
	}
	if o.InternalHost != nil {
		toSerialize["internal_host"] = o.InternalHost
	}
	if true {
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
	if o.Mode.IsSet() {
		toSerialize["mode"] = o.Mode.Get()
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

type NullableProxyProviderRequest struct {
	value *ProxyProviderRequest
	isSet bool
}

func (v NullableProxyProviderRequest) Get() *ProxyProviderRequest {
	return v.value
}

func (v *NullableProxyProviderRequest) Set(val *ProxyProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyProviderRequest(val *ProxyProviderRequest) *NullableProxyProviderRequest {
	return &NullableProxyProviderRequest{value: val, isSet: true}
}

func (v NullableProxyProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
