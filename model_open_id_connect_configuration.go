/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// OpenIDConnectConfiguration rest_framework Serializer for OIDC Configuration
type OpenIDConnectConfiguration struct {
	Issuer                            string   `json:"issuer"`
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	TokenEndpoint                     string   `json:"token_endpoint"`
	UserinfoEndpoint                  string   `json:"userinfo_endpoint"`
	EndSessionEndpoint                string   `json:"end_session_endpoint"`
	IntrospectionEndpoint             string   `json:"introspection_endpoint"`
	JwksUri                           string   `json:"jwks_uri"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	IdTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"`
	SubjectTypesSupported             []string `json:"subject_types_supported"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
}

// NewOpenIDConnectConfiguration instantiates a new OpenIDConnectConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIDConnectConfiguration(issuer string, authorizationEndpoint string, tokenEndpoint string, userinfoEndpoint string, endSessionEndpoint string, introspectionEndpoint string, jwksUri string, responseTypesSupported []string, idTokenSigningAlgValuesSupported []string, subjectTypesSupported []string, tokenEndpointAuthMethodsSupported []string) *OpenIDConnectConfiguration {
	this := OpenIDConnectConfiguration{}
	this.Issuer = issuer
	this.AuthorizationEndpoint = authorizationEndpoint
	this.TokenEndpoint = tokenEndpoint
	this.UserinfoEndpoint = userinfoEndpoint
	this.EndSessionEndpoint = endSessionEndpoint
	this.IntrospectionEndpoint = introspectionEndpoint
	this.JwksUri = jwksUri
	this.ResponseTypesSupported = responseTypesSupported
	this.IdTokenSigningAlgValuesSupported = idTokenSigningAlgValuesSupported
	this.SubjectTypesSupported = subjectTypesSupported
	this.TokenEndpointAuthMethodsSupported = tokenEndpointAuthMethodsSupported
	return &this
}

// NewOpenIDConnectConfigurationWithDefaults instantiates a new OpenIDConnectConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIDConnectConfigurationWithDefaults() *OpenIDConnectConfiguration {
	this := OpenIDConnectConfiguration{}
	return &this
}

// GetIssuer returns the Issuer field value
func (o *OpenIDConnectConfiguration) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *OpenIDConnectConfiguration) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *OpenIDConnectConfiguration) SetIssuer(v string) {
	o.Issuer = v
}

// GetAuthorizationEndpoint returns the AuthorizationEndpoint field value
func (o *OpenIDConnectConfiguration) GetAuthorizationEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationEndpoint
}

// GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field value
// and a boolean to check if the value has been set.
func (o *OpenIDConnectConfiguration) GetAuthorizationEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationEndpoint, true
}

// SetAuthorizationEndpoint sets field value
func (o *OpenIDConnectConfiguration) SetAuthorizationEndpoint(v string) {
	o.AuthorizationEndpoint = v
}

// GetTokenEndpoint returns the TokenEndpoint field value
func (o *OpenIDConnectConfiguration) GetTokenEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value
// and a boolean to check if the value has been set.
func (o *OpenIDConnectConfiguration) GetTokenEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenEndpoint, true
}

// SetTokenEndpoint sets field value
func (o *OpenIDConnectConfiguration) SetTokenEndpoint(v string) {
	o.TokenEndpoint = v
}

// GetUserinfoEndpoint returns the UserinfoEndpoint field value
func (o *OpenIDConnectConfiguration) GetUserinfoEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserinfoEndpoint
}

// GetUserinfoEndpointOk returns a tuple with the UserinfoEndpoint field value
// and a boolean to check if the value has been set.
func (o *OpenIDConnectConfiguration) GetUserinfoEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserinfoEndpoint, true
}

// SetUserinfoEndpoint sets field value
func (o *OpenIDConnectConfiguration) SetUserinfoEndpoint(v string) {
	o.UserinfoEndpoint = v
}

// GetEndSessionEndpoint returns the EndSessionEndpoint field value
func (o *OpenIDConnectConfiguration) GetEndSessionEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndSessionEndpoint
}

// GetEndSessionEndpointOk returns a tuple with the EndSessionEndpoint field value
// and a boolean to check if the value has been set.
func (o *OpenIDConnectConfiguration) GetEndSessionEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndSessionEndpoint, true
}

// SetEndSessionEndpoint sets field value
func (o *OpenIDConnectConfiguration) SetEndSessionEndpoint(v string) {
	o.EndSessionEndpoint = v
}

// GetIntrospectionEndpoint returns the IntrospectionEndpoint field value
func (o *OpenIDConnectConfiguration) GetIntrospectionEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntrospectionEndpoint
}

// GetIntrospectionEndpointOk returns a tuple with the IntrospectionEndpoint field value
// and a boolean to check if the value has been set.
func (o *OpenIDConnectConfiguration) GetIntrospectionEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntrospectionEndpoint, true
}

// SetIntrospectionEndpoint sets field value
func (o *OpenIDConnectConfiguration) SetIntrospectionEndpoint(v string) {
	o.IntrospectionEndpoint = v
}

// GetJwksUri returns the JwksUri field value
func (o *OpenIDConnectConfiguration) GetJwksUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value
// and a boolean to check if the value has been set.
func (o *OpenIDConnectConfiguration) GetJwksUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JwksUri, true
}

// SetJwksUri sets field value
func (o *OpenIDConnectConfiguration) SetJwksUri(v string) {
	o.JwksUri = v
}

// GetResponseTypesSupported returns the ResponseTypesSupported field value
func (o *OpenIDConnectConfiguration) GetResponseTypesSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResponseTypesSupported
}

// GetResponseTypesSupportedOk returns a tuple with the ResponseTypesSupported field value
// and a boolean to check if the value has been set.
func (o *OpenIDConnectConfiguration) GetResponseTypesSupportedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseTypesSupported, true
}

// SetResponseTypesSupported sets field value
func (o *OpenIDConnectConfiguration) SetResponseTypesSupported(v []string) {
	o.ResponseTypesSupported = v
}

// GetIdTokenSigningAlgValuesSupported returns the IdTokenSigningAlgValuesSupported field value
func (o *OpenIDConnectConfiguration) GetIdTokenSigningAlgValuesSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IdTokenSigningAlgValuesSupported
}

// GetIdTokenSigningAlgValuesSupportedOk returns a tuple with the IdTokenSigningAlgValuesSupported field value
// and a boolean to check if the value has been set.
func (o *OpenIDConnectConfiguration) GetIdTokenSigningAlgValuesSupportedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdTokenSigningAlgValuesSupported, true
}

// SetIdTokenSigningAlgValuesSupported sets field value
func (o *OpenIDConnectConfiguration) SetIdTokenSigningAlgValuesSupported(v []string) {
	o.IdTokenSigningAlgValuesSupported = v
}

// GetSubjectTypesSupported returns the SubjectTypesSupported field value
func (o *OpenIDConnectConfiguration) GetSubjectTypesSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SubjectTypesSupported
}

// GetSubjectTypesSupportedOk returns a tuple with the SubjectTypesSupported field value
// and a boolean to check if the value has been set.
func (o *OpenIDConnectConfiguration) GetSubjectTypesSupportedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubjectTypesSupported, true
}

// SetSubjectTypesSupported sets field value
func (o *OpenIDConnectConfiguration) SetSubjectTypesSupported(v []string) {
	o.SubjectTypesSupported = v
}

// GetTokenEndpointAuthMethodsSupported returns the TokenEndpointAuthMethodsSupported field value
func (o *OpenIDConnectConfiguration) GetTokenEndpointAuthMethodsSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TokenEndpointAuthMethodsSupported
}

// GetTokenEndpointAuthMethodsSupportedOk returns a tuple with the TokenEndpointAuthMethodsSupported field value
// and a boolean to check if the value has been set.
func (o *OpenIDConnectConfiguration) GetTokenEndpointAuthMethodsSupportedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenEndpointAuthMethodsSupported, true
}

// SetTokenEndpointAuthMethodsSupported sets field value
func (o *OpenIDConnectConfiguration) SetTokenEndpointAuthMethodsSupported(v []string) {
	o.TokenEndpointAuthMethodsSupported = v
}

func (o OpenIDConnectConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["issuer"] = o.Issuer
	}
	if true {
		toSerialize["authorization_endpoint"] = o.AuthorizationEndpoint
	}
	if true {
		toSerialize["token_endpoint"] = o.TokenEndpoint
	}
	if true {
		toSerialize["userinfo_endpoint"] = o.UserinfoEndpoint
	}
	if true {
		toSerialize["end_session_endpoint"] = o.EndSessionEndpoint
	}
	if true {
		toSerialize["introspection_endpoint"] = o.IntrospectionEndpoint
	}
	if true {
		toSerialize["jwks_uri"] = o.JwksUri
	}
	if true {
		toSerialize["response_types_supported"] = o.ResponseTypesSupported
	}
	if true {
		toSerialize["id_token_signing_alg_values_supported"] = o.IdTokenSigningAlgValuesSupported
	}
	if true {
		toSerialize["subject_types_supported"] = o.SubjectTypesSupported
	}
	if true {
		toSerialize["token_endpoint_auth_methods_supported"] = o.TokenEndpointAuthMethodsSupported
	}
	return json.Marshal(toSerialize)
}

type NullableOpenIDConnectConfiguration struct {
	value *OpenIDConnectConfiguration
	isSet bool
}

func (v NullableOpenIDConnectConfiguration) Get() *OpenIDConnectConfiguration {
	return v.value
}

func (v *NullableOpenIDConnectConfiguration) Set(val *OpenIDConnectConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIDConnectConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIDConnectConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIDConnectConfiguration(val *OpenIDConnectConfiguration) *NullableOpenIDConnectConfiguration {
	return &NullableOpenIDConnectConfiguration{value: val, isSet: true}
}

func (v NullableOpenIDConnectConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIDConnectConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
