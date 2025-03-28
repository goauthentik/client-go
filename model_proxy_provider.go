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

// ProxyProvider ProxyProvider Serializer
type ProxyProvider struct {
	Pk   int32  `json:"pk"`
	Name string `json:"name"`
	// Flow used for authentication when the associated application is accessed by an un-authenticated user.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow used when authorizing this provider.
	AuthorizationFlow string `json:"authorization_flow"`
	// Flow used ending the session from a provider.
	InvalidationFlow string   `json:"invalidation_flow"`
	PropertyMappings []string `json:"property_mappings,omitempty"`
	// Get object component so that we know how to edit the object
	Component string `json:"component"`
	// Internal application name, used in URLs.
	AssignedApplicationSlug string `json:"assigned_application_slug"`
	// Application's display Name.
	AssignedApplicationName string `json:"assigned_application_name"`
	// Internal application name, used in URLs.
	AssignedBackchannelApplicationSlug string `json:"assigned_backchannel_application_slug"`
	// Application's display Name.
	AssignedBackchannelApplicationName string `json:"assigned_backchannel_application_name"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName string  `json:"meta_model_name"`
	ClientId      string  `json:"client_id"`
	InternalHost  *string `json:"internal_host,omitempty"`
	ExternalHost  string  `json:"external_host"`
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
	InterceptHeaderAuth    *bool         `json:"intercept_header_auth,omitempty"`
	RedirectUris           []RedirectURI `json:"redirect_uris"`
	CookieDomain           *string       `json:"cookie_domain,omitempty"`
	JwtFederationSources   []string      `json:"jwt_federation_sources,omitempty"`
	JwtFederationProviders []int32       `json:"jwt_federation_providers,omitempty"`
	// Tokens not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	AccessTokenValidity *string `json:"access_token_validity,omitempty"`
	// Tokens not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	RefreshTokenValidity *string  `json:"refresh_token_validity,omitempty"`
	OutpostSet           []string `json:"outpost_set"`
}

// NewProxyProvider instantiates a new ProxyProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyProvider(pk int32, name string, authorizationFlow string, invalidationFlow string, component string, assignedApplicationSlug string, assignedApplicationName string, assignedBackchannelApplicationSlug string, assignedBackchannelApplicationName string, verboseName string, verboseNamePlural string, metaModelName string, clientId string, externalHost string, redirectUris []RedirectURI, outpostSet []string) *ProxyProvider {
	this := ProxyProvider{}
	this.Pk = pk
	this.Name = name
	this.AuthorizationFlow = authorizationFlow
	this.InvalidationFlow = invalidationFlow
	this.Component = component
	this.AssignedApplicationSlug = assignedApplicationSlug
	this.AssignedApplicationName = assignedApplicationName
	this.AssignedBackchannelApplicationSlug = assignedBackchannelApplicationSlug
	this.AssignedBackchannelApplicationName = assignedBackchannelApplicationName
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.ClientId = clientId
	this.ExternalHost = externalHost
	this.RedirectUris = redirectUris
	this.OutpostSet = outpostSet
	return &this
}

// NewProxyProviderWithDefaults instantiates a new ProxyProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyProviderWithDefaults() *ProxyProvider {
	this := ProxyProvider{}
	return &this
}

// GetPk returns the Pk field value
func (o *ProxyProvider) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *ProxyProvider) SetPk(v int32) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *ProxyProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProxyProvider) SetName(v string) {
	o.Name = v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxyProvider) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxyProvider) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *ProxyProvider) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *ProxyProvider) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *ProxyProvider) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *ProxyProvider) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetAuthorizationFlow returns the AuthorizationFlow field value
func (o *ProxyProvider) GetAuthorizationFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationFlow, true
}

// SetAuthorizationFlow sets field value
func (o *ProxyProvider) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = v
}

// GetInvalidationFlow returns the InvalidationFlow field value
func (o *ProxyProvider) GetInvalidationFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvalidationFlow
}

// GetInvalidationFlowOk returns a tuple with the InvalidationFlow field value
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetInvalidationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvalidationFlow, true
}

// SetInvalidationFlow sets field value
func (o *ProxyProvider) SetInvalidationFlow(v string) {
	o.InvalidationFlow = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *ProxyProvider) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *ProxyProvider) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *ProxyProvider) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetComponent returns the Component field value
func (o *ProxyProvider) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *ProxyProvider) SetComponent(v string) {
	o.Component = v
}

// GetAssignedApplicationSlug returns the AssignedApplicationSlug field value
func (o *ProxyProvider) GetAssignedApplicationSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedApplicationSlug
}

// GetAssignedApplicationSlugOk returns a tuple with the AssignedApplicationSlug field value
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetAssignedApplicationSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedApplicationSlug, true
}

// SetAssignedApplicationSlug sets field value
func (o *ProxyProvider) SetAssignedApplicationSlug(v string) {
	o.AssignedApplicationSlug = v
}

// GetAssignedApplicationName returns the AssignedApplicationName field value
func (o *ProxyProvider) GetAssignedApplicationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedApplicationName
}

// GetAssignedApplicationNameOk returns a tuple with the AssignedApplicationName field value
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetAssignedApplicationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedApplicationName, true
}

// SetAssignedApplicationName sets field value
func (o *ProxyProvider) SetAssignedApplicationName(v string) {
	o.AssignedApplicationName = v
}

// GetAssignedBackchannelApplicationSlug returns the AssignedBackchannelApplicationSlug field value
func (o *ProxyProvider) GetAssignedBackchannelApplicationSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedBackchannelApplicationSlug
}

// GetAssignedBackchannelApplicationSlugOk returns a tuple with the AssignedBackchannelApplicationSlug field value
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetAssignedBackchannelApplicationSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedBackchannelApplicationSlug, true
}

// SetAssignedBackchannelApplicationSlug sets field value
func (o *ProxyProvider) SetAssignedBackchannelApplicationSlug(v string) {
	o.AssignedBackchannelApplicationSlug = v
}

// GetAssignedBackchannelApplicationName returns the AssignedBackchannelApplicationName field value
func (o *ProxyProvider) GetAssignedBackchannelApplicationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedBackchannelApplicationName
}

// GetAssignedBackchannelApplicationNameOk returns a tuple with the AssignedBackchannelApplicationName field value
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetAssignedBackchannelApplicationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedBackchannelApplicationName, true
}

// SetAssignedBackchannelApplicationName sets field value
func (o *ProxyProvider) SetAssignedBackchannelApplicationName(v string) {
	o.AssignedBackchannelApplicationName = v
}

// GetVerboseName returns the VerboseName field value
func (o *ProxyProvider) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *ProxyProvider) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *ProxyProvider) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *ProxyProvider) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *ProxyProvider) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *ProxyProvider) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetClientId returns the ClientId field value
func (o *ProxyProvider) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *ProxyProvider) SetClientId(v string) {
	o.ClientId = v
}

// GetInternalHost returns the InternalHost field value if set, zero value otherwise.
func (o *ProxyProvider) GetInternalHost() string {
	if o == nil || o.InternalHost == nil {
		var ret string
		return ret
	}
	return *o.InternalHost
}

// GetInternalHostOk returns a tuple with the InternalHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetInternalHostOk() (*string, bool) {
	if o == nil || o.InternalHost == nil {
		return nil, false
	}
	return o.InternalHost, true
}

// HasInternalHost returns a boolean if a field has been set.
func (o *ProxyProvider) HasInternalHost() bool {
	if o != nil && o.InternalHost != nil {
		return true
	}

	return false
}

// SetInternalHost gets a reference to the given string and assigns it to the InternalHost field.
func (o *ProxyProvider) SetInternalHost(v string) {
	o.InternalHost = &v
}

// GetExternalHost returns the ExternalHost field value
func (o *ProxyProvider) GetExternalHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalHost
}

// GetExternalHostOk returns a tuple with the ExternalHost field value
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetExternalHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalHost, true
}

// SetExternalHost sets field value
func (o *ProxyProvider) SetExternalHost(v string) {
	o.ExternalHost = v
}

// GetInternalHostSslValidation returns the InternalHostSslValidation field value if set, zero value otherwise.
func (o *ProxyProvider) GetInternalHostSslValidation() bool {
	if o == nil || o.InternalHostSslValidation == nil {
		var ret bool
		return ret
	}
	return *o.InternalHostSslValidation
}

// GetInternalHostSslValidationOk returns a tuple with the InternalHostSslValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetInternalHostSslValidationOk() (*bool, bool) {
	if o == nil || o.InternalHostSslValidation == nil {
		return nil, false
	}
	return o.InternalHostSslValidation, true
}

// HasInternalHostSslValidation returns a boolean if a field has been set.
func (o *ProxyProvider) HasInternalHostSslValidation() bool {
	if o != nil && o.InternalHostSslValidation != nil {
		return true
	}

	return false
}

// SetInternalHostSslValidation gets a reference to the given bool and assigns it to the InternalHostSslValidation field.
func (o *ProxyProvider) SetInternalHostSslValidation(v bool) {
	o.InternalHostSslValidation = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxyProvider) GetCertificate() string {
	if o == nil || o.Certificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxyProvider) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *ProxyProvider) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *ProxyProvider) SetCertificate(v string) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *ProxyProvider) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *ProxyProvider) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetSkipPathRegex returns the SkipPathRegex field value if set, zero value otherwise.
func (o *ProxyProvider) GetSkipPathRegex() string {
	if o == nil || o.SkipPathRegex == nil {
		var ret string
		return ret
	}
	return *o.SkipPathRegex
}

// GetSkipPathRegexOk returns a tuple with the SkipPathRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetSkipPathRegexOk() (*string, bool) {
	if o == nil || o.SkipPathRegex == nil {
		return nil, false
	}
	return o.SkipPathRegex, true
}

// HasSkipPathRegex returns a boolean if a field has been set.
func (o *ProxyProvider) HasSkipPathRegex() bool {
	if o != nil && o.SkipPathRegex != nil {
		return true
	}

	return false
}

// SetSkipPathRegex gets a reference to the given string and assigns it to the SkipPathRegex field.
func (o *ProxyProvider) SetSkipPathRegex(v string) {
	o.SkipPathRegex = &v
}

// GetBasicAuthEnabled returns the BasicAuthEnabled field value if set, zero value otherwise.
func (o *ProxyProvider) GetBasicAuthEnabled() bool {
	if o == nil || o.BasicAuthEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BasicAuthEnabled
}

// GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetBasicAuthEnabledOk() (*bool, bool) {
	if o == nil || o.BasicAuthEnabled == nil {
		return nil, false
	}
	return o.BasicAuthEnabled, true
}

// HasBasicAuthEnabled returns a boolean if a field has been set.
func (o *ProxyProvider) HasBasicAuthEnabled() bool {
	if o != nil && o.BasicAuthEnabled != nil {
		return true
	}

	return false
}

// SetBasicAuthEnabled gets a reference to the given bool and assigns it to the BasicAuthEnabled field.
func (o *ProxyProvider) SetBasicAuthEnabled(v bool) {
	o.BasicAuthEnabled = &v
}

// GetBasicAuthPasswordAttribute returns the BasicAuthPasswordAttribute field value if set, zero value otherwise.
func (o *ProxyProvider) GetBasicAuthPasswordAttribute() string {
	if o == nil || o.BasicAuthPasswordAttribute == nil {
		var ret string
		return ret
	}
	return *o.BasicAuthPasswordAttribute
}

// GetBasicAuthPasswordAttributeOk returns a tuple with the BasicAuthPasswordAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetBasicAuthPasswordAttributeOk() (*string, bool) {
	if o == nil || o.BasicAuthPasswordAttribute == nil {
		return nil, false
	}
	return o.BasicAuthPasswordAttribute, true
}

// HasBasicAuthPasswordAttribute returns a boolean if a field has been set.
func (o *ProxyProvider) HasBasicAuthPasswordAttribute() bool {
	if o != nil && o.BasicAuthPasswordAttribute != nil {
		return true
	}

	return false
}

// SetBasicAuthPasswordAttribute gets a reference to the given string and assigns it to the BasicAuthPasswordAttribute field.
func (o *ProxyProvider) SetBasicAuthPasswordAttribute(v string) {
	o.BasicAuthPasswordAttribute = &v
}

// GetBasicAuthUserAttribute returns the BasicAuthUserAttribute field value if set, zero value otherwise.
func (o *ProxyProvider) GetBasicAuthUserAttribute() string {
	if o == nil || o.BasicAuthUserAttribute == nil {
		var ret string
		return ret
	}
	return *o.BasicAuthUserAttribute
}

// GetBasicAuthUserAttributeOk returns a tuple with the BasicAuthUserAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetBasicAuthUserAttributeOk() (*string, bool) {
	if o == nil || o.BasicAuthUserAttribute == nil {
		return nil, false
	}
	return o.BasicAuthUserAttribute, true
}

// HasBasicAuthUserAttribute returns a boolean if a field has been set.
func (o *ProxyProvider) HasBasicAuthUserAttribute() bool {
	if o != nil && o.BasicAuthUserAttribute != nil {
		return true
	}

	return false
}

// SetBasicAuthUserAttribute gets a reference to the given string and assigns it to the BasicAuthUserAttribute field.
func (o *ProxyProvider) SetBasicAuthUserAttribute(v string) {
	o.BasicAuthUserAttribute = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ProxyProvider) GetMode() ProxyMode {
	if o == nil || o.Mode == nil {
		var ret ProxyMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetModeOk() (*ProxyMode, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ProxyProvider) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given ProxyMode and assigns it to the Mode field.
func (o *ProxyProvider) SetMode(v ProxyMode) {
	o.Mode = &v
}

// GetInterceptHeaderAuth returns the InterceptHeaderAuth field value if set, zero value otherwise.
func (o *ProxyProvider) GetInterceptHeaderAuth() bool {
	if o == nil || o.InterceptHeaderAuth == nil {
		var ret bool
		return ret
	}
	return *o.InterceptHeaderAuth
}

// GetInterceptHeaderAuthOk returns a tuple with the InterceptHeaderAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetInterceptHeaderAuthOk() (*bool, bool) {
	if o == nil || o.InterceptHeaderAuth == nil {
		return nil, false
	}
	return o.InterceptHeaderAuth, true
}

// HasInterceptHeaderAuth returns a boolean if a field has been set.
func (o *ProxyProvider) HasInterceptHeaderAuth() bool {
	if o != nil && o.InterceptHeaderAuth != nil {
		return true
	}

	return false
}

// SetInterceptHeaderAuth gets a reference to the given bool and assigns it to the InterceptHeaderAuth field.
func (o *ProxyProvider) SetInterceptHeaderAuth(v bool) {
	o.InterceptHeaderAuth = &v
}

// GetRedirectUris returns the RedirectUris field value
func (o *ProxyProvider) GetRedirectUris() []RedirectURI {
	if o == nil {
		var ret []RedirectURI
		return ret
	}

	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetRedirectUrisOk() ([]RedirectURI, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// SetRedirectUris sets field value
func (o *ProxyProvider) SetRedirectUris(v []RedirectURI) {
	o.RedirectUris = v
}

// GetCookieDomain returns the CookieDomain field value if set, zero value otherwise.
func (o *ProxyProvider) GetCookieDomain() string {
	if o == nil || o.CookieDomain == nil {
		var ret string
		return ret
	}
	return *o.CookieDomain
}

// GetCookieDomainOk returns a tuple with the CookieDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetCookieDomainOk() (*string, bool) {
	if o == nil || o.CookieDomain == nil {
		return nil, false
	}
	return o.CookieDomain, true
}

// HasCookieDomain returns a boolean if a field has been set.
func (o *ProxyProvider) HasCookieDomain() bool {
	if o != nil && o.CookieDomain != nil {
		return true
	}

	return false
}

// SetCookieDomain gets a reference to the given string and assigns it to the CookieDomain field.
func (o *ProxyProvider) SetCookieDomain(v string) {
	o.CookieDomain = &v
}

// GetJwtFederationSources returns the JwtFederationSources field value if set, zero value otherwise.
func (o *ProxyProvider) GetJwtFederationSources() []string {
	if o == nil || o.JwtFederationSources == nil {
		var ret []string
		return ret
	}
	return o.JwtFederationSources
}

// GetJwtFederationSourcesOk returns a tuple with the JwtFederationSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetJwtFederationSourcesOk() ([]string, bool) {
	if o == nil || o.JwtFederationSources == nil {
		return nil, false
	}
	return o.JwtFederationSources, true
}

// HasJwtFederationSources returns a boolean if a field has been set.
func (o *ProxyProvider) HasJwtFederationSources() bool {
	if o != nil && o.JwtFederationSources != nil {
		return true
	}

	return false
}

// SetJwtFederationSources gets a reference to the given []string and assigns it to the JwtFederationSources field.
func (o *ProxyProvider) SetJwtFederationSources(v []string) {
	o.JwtFederationSources = v
}

// GetJwtFederationProviders returns the JwtFederationProviders field value if set, zero value otherwise.
func (o *ProxyProvider) GetJwtFederationProviders() []int32 {
	if o == nil || o.JwtFederationProviders == nil {
		var ret []int32
		return ret
	}
	return o.JwtFederationProviders
}

// GetJwtFederationProvidersOk returns a tuple with the JwtFederationProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetJwtFederationProvidersOk() ([]int32, bool) {
	if o == nil || o.JwtFederationProviders == nil {
		return nil, false
	}
	return o.JwtFederationProviders, true
}

// HasJwtFederationProviders returns a boolean if a field has been set.
func (o *ProxyProvider) HasJwtFederationProviders() bool {
	if o != nil && o.JwtFederationProviders != nil {
		return true
	}

	return false
}

// SetJwtFederationProviders gets a reference to the given []int32 and assigns it to the JwtFederationProviders field.
func (o *ProxyProvider) SetJwtFederationProviders(v []int32) {
	o.JwtFederationProviders = v
}

// GetAccessTokenValidity returns the AccessTokenValidity field value if set, zero value otherwise.
func (o *ProxyProvider) GetAccessTokenValidity() string {
	if o == nil || o.AccessTokenValidity == nil {
		var ret string
		return ret
	}
	return *o.AccessTokenValidity
}

// GetAccessTokenValidityOk returns a tuple with the AccessTokenValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetAccessTokenValidityOk() (*string, bool) {
	if o == nil || o.AccessTokenValidity == nil {
		return nil, false
	}
	return o.AccessTokenValidity, true
}

// HasAccessTokenValidity returns a boolean if a field has been set.
func (o *ProxyProvider) HasAccessTokenValidity() bool {
	if o != nil && o.AccessTokenValidity != nil {
		return true
	}

	return false
}

// SetAccessTokenValidity gets a reference to the given string and assigns it to the AccessTokenValidity field.
func (o *ProxyProvider) SetAccessTokenValidity(v string) {
	o.AccessTokenValidity = &v
}

// GetRefreshTokenValidity returns the RefreshTokenValidity field value if set, zero value otherwise.
func (o *ProxyProvider) GetRefreshTokenValidity() string {
	if o == nil || o.RefreshTokenValidity == nil {
		var ret string
		return ret
	}
	return *o.RefreshTokenValidity
}

// GetRefreshTokenValidityOk returns a tuple with the RefreshTokenValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetRefreshTokenValidityOk() (*string, bool) {
	if o == nil || o.RefreshTokenValidity == nil {
		return nil, false
	}
	return o.RefreshTokenValidity, true
}

// HasRefreshTokenValidity returns a boolean if a field has been set.
func (o *ProxyProvider) HasRefreshTokenValidity() bool {
	if o != nil && o.RefreshTokenValidity != nil {
		return true
	}

	return false
}

// SetRefreshTokenValidity gets a reference to the given string and assigns it to the RefreshTokenValidity field.
func (o *ProxyProvider) SetRefreshTokenValidity(v string) {
	o.RefreshTokenValidity = &v
}

// GetOutpostSet returns the OutpostSet field value
func (o *ProxyProvider) GetOutpostSet() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OutpostSet
}

// GetOutpostSetOk returns a tuple with the OutpostSet field value
// and a boolean to check if the value has been set.
func (o *ProxyProvider) GetOutpostSetOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutpostSet, true
}

// SetOutpostSet sets field value
func (o *ProxyProvider) SetOutpostSet(v []string) {
	o.OutpostSet = v
}

func (o ProxyProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
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
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["assigned_application_slug"] = o.AssignedApplicationSlug
	}
	if true {
		toSerialize["assigned_application_name"] = o.AssignedApplicationName
	}
	if true {
		toSerialize["assigned_backchannel_application_slug"] = o.AssignedBackchannelApplicationSlug
	}
	if true {
		toSerialize["assigned_backchannel_application_name"] = o.AssignedBackchannelApplicationName
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
	if true {
		toSerialize["client_id"] = o.ClientId
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
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.InterceptHeaderAuth != nil {
		toSerialize["intercept_header_auth"] = o.InterceptHeaderAuth
	}
	if true {
		toSerialize["redirect_uris"] = o.RedirectUris
	}
	if o.CookieDomain != nil {
		toSerialize["cookie_domain"] = o.CookieDomain
	}
	if o.JwtFederationSources != nil {
		toSerialize["jwt_federation_sources"] = o.JwtFederationSources
	}
	if o.JwtFederationProviders != nil {
		toSerialize["jwt_federation_providers"] = o.JwtFederationProviders
	}
	if o.AccessTokenValidity != nil {
		toSerialize["access_token_validity"] = o.AccessTokenValidity
	}
	if o.RefreshTokenValidity != nil {
		toSerialize["refresh_token_validity"] = o.RefreshTokenValidity
	}
	if true {
		toSerialize["outpost_set"] = o.OutpostSet
	}
	return json.Marshal(toSerialize)
}

type NullableProxyProvider struct {
	value *ProxyProvider
	isSet bool
}

func (v NullableProxyProvider) Get() *ProxyProvider {
	return v.value
}

func (v *NullableProxyProvider) Set(val *ProxyProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyProvider(val *ProxyProvider) *NullableProxyProvider {
	return &NullableProxyProvider{value: val, isSet: true}
}

func (v NullableProxyProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
