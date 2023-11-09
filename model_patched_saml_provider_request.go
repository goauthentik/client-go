/*
authentik

Making authentication simple.

API version: 2023.10.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedSAMLProviderRequest SAMLProvider Serializer
type PatchedSAMLProviderRequest struct {
	Name *string `json:"name,omitempty"`
	// Flow used for authentication when the associated application is accessed by an un-authenticated user.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow used when authorizing this provider.
	AuthorizationFlow *string  `json:"authorization_flow,omitempty"`
	PropertyMappings  []string `json:"property_mappings,omitempty"`
	AcsUrl            *string  `json:"acs_url,omitempty"`
	// Value of the audience restriction field of the assertion. When left empty, no audience restriction will be added.
	Audience *string `json:"audience,omitempty"`
	// Also known as EntityID
	Issuer *string `json:"issuer,omitempty"`
	// Assertion valid not before current time + this value (Format: hours=-1;minutes=-2;seconds=-3).
	AssertionValidNotBefore *string `json:"assertion_valid_not_before,omitempty"`
	// Assertion not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	AssertionValidNotOnOrAfter *string `json:"assertion_valid_not_on_or_after,omitempty"`
	// Session not valid on or after current time + this value (Format: hours=1;minutes=2;seconds=3).
	SessionValidNotOnOrAfter *string `json:"session_valid_not_on_or_after,omitempty"`
	// Configure how the NameID value will be created. When left empty, the NameIDPolicy of the incoming request will be considered
	NameIdMapping      NullableString          `json:"name_id_mapping,omitempty"`
	DigestAlgorithm    *DigestAlgorithmEnum    `json:"digest_algorithm,omitempty"`
	SignatureAlgorithm *SignatureAlgorithmEnum `json:"signature_algorithm,omitempty"`
	// Keypair used to sign outgoing Responses going to the Service Provider.
	SigningKp NullableString `json:"signing_kp,omitempty"`
	// When selected, incoming assertion's Signatures will be validated against this certificate. To allow unsigned Requests, leave on default.
	VerificationKp NullableString `json:"verification_kp,omitempty"`
	// This determines how authentik sends the response back to the Service Provider.  * `redirect` - Redirect * `post` - Post
	SpBinding *SpBindingEnum `json:"sp_binding,omitempty"`
	// Default relay_state value for IDP-initiated logins
	DefaultRelayState *string `json:"default_relay_state,omitempty"`
}

// NewPatchedSAMLProviderRequest instantiates a new PatchedSAMLProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedSAMLProviderRequest() *PatchedSAMLProviderRequest {
	this := PatchedSAMLProviderRequest{}
	return &this
}

// NewPatchedSAMLProviderRequestWithDefaults instantiates a new PatchedSAMLProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedSAMLProviderRequestWithDefaults() *PatchedSAMLProviderRequest {
	this := PatchedSAMLProviderRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedSAMLProviderRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLProviderRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedSAMLProviderRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedSAMLProviderRequest) SetName(v string) {
	o.Name = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedSAMLProviderRequest) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedSAMLProviderRequest) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *PatchedSAMLProviderRequest) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *PatchedSAMLProviderRequest) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *PatchedSAMLProviderRequest) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *PatchedSAMLProviderRequest) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetAuthorizationFlow returns the AuthorizationFlow field value if set, zero value otherwise.
func (o *PatchedSAMLProviderRequest) GetAuthorizationFlow() string {
	if o == nil || o.AuthorizationFlow == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLProviderRequest) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil || o.AuthorizationFlow == nil {
		return nil, false
	}
	return o.AuthorizationFlow, true
}

// HasAuthorizationFlow returns a boolean if a field has been set.
func (o *PatchedSAMLProviderRequest) HasAuthorizationFlow() bool {
	if o != nil && o.AuthorizationFlow != nil {
		return true
	}

	return false
}

// SetAuthorizationFlow gets a reference to the given string and assigns it to the AuthorizationFlow field.
func (o *PatchedSAMLProviderRequest) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = &v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *PatchedSAMLProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *PatchedSAMLProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *PatchedSAMLProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetAcsUrl returns the AcsUrl field value if set, zero value otherwise.
func (o *PatchedSAMLProviderRequest) GetAcsUrl() string {
	if o == nil || o.AcsUrl == nil {
		var ret string
		return ret
	}
	return *o.AcsUrl
}

// GetAcsUrlOk returns a tuple with the AcsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLProviderRequest) GetAcsUrlOk() (*string, bool) {
	if o == nil || o.AcsUrl == nil {
		return nil, false
	}
	return o.AcsUrl, true
}

// HasAcsUrl returns a boolean if a field has been set.
func (o *PatchedSAMLProviderRequest) HasAcsUrl() bool {
	if o != nil && o.AcsUrl != nil {
		return true
	}

	return false
}

// SetAcsUrl gets a reference to the given string and assigns it to the AcsUrl field.
func (o *PatchedSAMLProviderRequest) SetAcsUrl(v string) {
	o.AcsUrl = &v
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *PatchedSAMLProviderRequest) GetAudience() string {
	if o == nil || o.Audience == nil {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLProviderRequest) GetAudienceOk() (*string, bool) {
	if o == nil || o.Audience == nil {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *PatchedSAMLProviderRequest) HasAudience() bool {
	if o != nil && o.Audience != nil {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *PatchedSAMLProviderRequest) SetAudience(v string) {
	o.Audience = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *PatchedSAMLProviderRequest) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLProviderRequest) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *PatchedSAMLProviderRequest) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *PatchedSAMLProviderRequest) SetIssuer(v string) {
	o.Issuer = &v
}

// GetAssertionValidNotBefore returns the AssertionValidNotBefore field value if set, zero value otherwise.
func (o *PatchedSAMLProviderRequest) GetAssertionValidNotBefore() string {
	if o == nil || o.AssertionValidNotBefore == nil {
		var ret string
		return ret
	}
	return *o.AssertionValidNotBefore
}

// GetAssertionValidNotBeforeOk returns a tuple with the AssertionValidNotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLProviderRequest) GetAssertionValidNotBeforeOk() (*string, bool) {
	if o == nil || o.AssertionValidNotBefore == nil {
		return nil, false
	}
	return o.AssertionValidNotBefore, true
}

// HasAssertionValidNotBefore returns a boolean if a field has been set.
func (o *PatchedSAMLProviderRequest) HasAssertionValidNotBefore() bool {
	if o != nil && o.AssertionValidNotBefore != nil {
		return true
	}

	return false
}

// SetAssertionValidNotBefore gets a reference to the given string and assigns it to the AssertionValidNotBefore field.
func (o *PatchedSAMLProviderRequest) SetAssertionValidNotBefore(v string) {
	o.AssertionValidNotBefore = &v
}

// GetAssertionValidNotOnOrAfter returns the AssertionValidNotOnOrAfter field value if set, zero value otherwise.
func (o *PatchedSAMLProviderRequest) GetAssertionValidNotOnOrAfter() string {
	if o == nil || o.AssertionValidNotOnOrAfter == nil {
		var ret string
		return ret
	}
	return *o.AssertionValidNotOnOrAfter
}

// GetAssertionValidNotOnOrAfterOk returns a tuple with the AssertionValidNotOnOrAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLProviderRequest) GetAssertionValidNotOnOrAfterOk() (*string, bool) {
	if o == nil || o.AssertionValidNotOnOrAfter == nil {
		return nil, false
	}
	return o.AssertionValidNotOnOrAfter, true
}

// HasAssertionValidNotOnOrAfter returns a boolean if a field has been set.
func (o *PatchedSAMLProviderRequest) HasAssertionValidNotOnOrAfter() bool {
	if o != nil && o.AssertionValidNotOnOrAfter != nil {
		return true
	}

	return false
}

// SetAssertionValidNotOnOrAfter gets a reference to the given string and assigns it to the AssertionValidNotOnOrAfter field.
func (o *PatchedSAMLProviderRequest) SetAssertionValidNotOnOrAfter(v string) {
	o.AssertionValidNotOnOrAfter = &v
}

// GetSessionValidNotOnOrAfter returns the SessionValidNotOnOrAfter field value if set, zero value otherwise.
func (o *PatchedSAMLProviderRequest) GetSessionValidNotOnOrAfter() string {
	if o == nil || o.SessionValidNotOnOrAfter == nil {
		var ret string
		return ret
	}
	return *o.SessionValidNotOnOrAfter
}

// GetSessionValidNotOnOrAfterOk returns a tuple with the SessionValidNotOnOrAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLProviderRequest) GetSessionValidNotOnOrAfterOk() (*string, bool) {
	if o == nil || o.SessionValidNotOnOrAfter == nil {
		return nil, false
	}
	return o.SessionValidNotOnOrAfter, true
}

// HasSessionValidNotOnOrAfter returns a boolean if a field has been set.
func (o *PatchedSAMLProviderRequest) HasSessionValidNotOnOrAfter() bool {
	if o != nil && o.SessionValidNotOnOrAfter != nil {
		return true
	}

	return false
}

// SetSessionValidNotOnOrAfter gets a reference to the given string and assigns it to the SessionValidNotOnOrAfter field.
func (o *PatchedSAMLProviderRequest) SetSessionValidNotOnOrAfter(v string) {
	o.SessionValidNotOnOrAfter = &v
}

// GetNameIdMapping returns the NameIdMapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedSAMLProviderRequest) GetNameIdMapping() string {
	if o == nil || o.NameIdMapping.Get() == nil {
		var ret string
		return ret
	}
	return *o.NameIdMapping.Get()
}

// GetNameIdMappingOk returns a tuple with the NameIdMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedSAMLProviderRequest) GetNameIdMappingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NameIdMapping.Get(), o.NameIdMapping.IsSet()
}

// HasNameIdMapping returns a boolean if a field has been set.
func (o *PatchedSAMLProviderRequest) HasNameIdMapping() bool {
	if o != nil && o.NameIdMapping.IsSet() {
		return true
	}

	return false
}

// SetNameIdMapping gets a reference to the given NullableString and assigns it to the NameIdMapping field.
func (o *PatchedSAMLProviderRequest) SetNameIdMapping(v string) {
	o.NameIdMapping.Set(&v)
}

// SetNameIdMappingNil sets the value for NameIdMapping to be an explicit nil
func (o *PatchedSAMLProviderRequest) SetNameIdMappingNil() {
	o.NameIdMapping.Set(nil)
}

// UnsetNameIdMapping ensures that no value is present for NameIdMapping, not even an explicit nil
func (o *PatchedSAMLProviderRequest) UnsetNameIdMapping() {
	o.NameIdMapping.Unset()
}

// GetDigestAlgorithm returns the DigestAlgorithm field value if set, zero value otherwise.
func (o *PatchedSAMLProviderRequest) GetDigestAlgorithm() DigestAlgorithmEnum {
	if o == nil || o.DigestAlgorithm == nil {
		var ret DigestAlgorithmEnum
		return ret
	}
	return *o.DigestAlgorithm
}

// GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLProviderRequest) GetDigestAlgorithmOk() (*DigestAlgorithmEnum, bool) {
	if o == nil || o.DigestAlgorithm == nil {
		return nil, false
	}
	return o.DigestAlgorithm, true
}

// HasDigestAlgorithm returns a boolean if a field has been set.
func (o *PatchedSAMLProviderRequest) HasDigestAlgorithm() bool {
	if o != nil && o.DigestAlgorithm != nil {
		return true
	}

	return false
}

// SetDigestAlgorithm gets a reference to the given DigestAlgorithmEnum and assigns it to the DigestAlgorithm field.
func (o *PatchedSAMLProviderRequest) SetDigestAlgorithm(v DigestAlgorithmEnum) {
	o.DigestAlgorithm = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *PatchedSAMLProviderRequest) GetSignatureAlgorithm() SignatureAlgorithmEnum {
	if o == nil || o.SignatureAlgorithm == nil {
		var ret SignatureAlgorithmEnum
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLProviderRequest) GetSignatureAlgorithmOk() (*SignatureAlgorithmEnum, bool) {
	if o == nil || o.SignatureAlgorithm == nil {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *PatchedSAMLProviderRequest) HasSignatureAlgorithm() bool {
	if o != nil && o.SignatureAlgorithm != nil {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given SignatureAlgorithmEnum and assigns it to the SignatureAlgorithm field.
func (o *PatchedSAMLProviderRequest) SetSignatureAlgorithm(v SignatureAlgorithmEnum) {
	o.SignatureAlgorithm = &v
}

// GetSigningKp returns the SigningKp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedSAMLProviderRequest) GetSigningKp() string {
	if o == nil || o.SigningKp.Get() == nil {
		var ret string
		return ret
	}
	return *o.SigningKp.Get()
}

// GetSigningKpOk returns a tuple with the SigningKp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedSAMLProviderRequest) GetSigningKpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SigningKp.Get(), o.SigningKp.IsSet()
}

// HasSigningKp returns a boolean if a field has been set.
func (o *PatchedSAMLProviderRequest) HasSigningKp() bool {
	if o != nil && o.SigningKp.IsSet() {
		return true
	}

	return false
}

// SetSigningKp gets a reference to the given NullableString and assigns it to the SigningKp field.
func (o *PatchedSAMLProviderRequest) SetSigningKp(v string) {
	o.SigningKp.Set(&v)
}

// SetSigningKpNil sets the value for SigningKp to be an explicit nil
func (o *PatchedSAMLProviderRequest) SetSigningKpNil() {
	o.SigningKp.Set(nil)
}

// UnsetSigningKp ensures that no value is present for SigningKp, not even an explicit nil
func (o *PatchedSAMLProviderRequest) UnsetSigningKp() {
	o.SigningKp.Unset()
}

// GetVerificationKp returns the VerificationKp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedSAMLProviderRequest) GetVerificationKp() string {
	if o == nil || o.VerificationKp.Get() == nil {
		var ret string
		return ret
	}
	return *o.VerificationKp.Get()
}

// GetVerificationKpOk returns a tuple with the VerificationKp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedSAMLProviderRequest) GetVerificationKpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VerificationKp.Get(), o.VerificationKp.IsSet()
}

// HasVerificationKp returns a boolean if a field has been set.
func (o *PatchedSAMLProviderRequest) HasVerificationKp() bool {
	if o != nil && o.VerificationKp.IsSet() {
		return true
	}

	return false
}

// SetVerificationKp gets a reference to the given NullableString and assigns it to the VerificationKp field.
func (o *PatchedSAMLProviderRequest) SetVerificationKp(v string) {
	o.VerificationKp.Set(&v)
}

// SetVerificationKpNil sets the value for VerificationKp to be an explicit nil
func (o *PatchedSAMLProviderRequest) SetVerificationKpNil() {
	o.VerificationKp.Set(nil)
}

// UnsetVerificationKp ensures that no value is present for VerificationKp, not even an explicit nil
func (o *PatchedSAMLProviderRequest) UnsetVerificationKp() {
	o.VerificationKp.Unset()
}

// GetSpBinding returns the SpBinding field value if set, zero value otherwise.
func (o *PatchedSAMLProviderRequest) GetSpBinding() SpBindingEnum {
	if o == nil || o.SpBinding == nil {
		var ret SpBindingEnum
		return ret
	}
	return *o.SpBinding
}

// GetSpBindingOk returns a tuple with the SpBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLProviderRequest) GetSpBindingOk() (*SpBindingEnum, bool) {
	if o == nil || o.SpBinding == nil {
		return nil, false
	}
	return o.SpBinding, true
}

// HasSpBinding returns a boolean if a field has been set.
func (o *PatchedSAMLProviderRequest) HasSpBinding() bool {
	if o != nil && o.SpBinding != nil {
		return true
	}

	return false
}

// SetSpBinding gets a reference to the given SpBindingEnum and assigns it to the SpBinding field.
func (o *PatchedSAMLProviderRequest) SetSpBinding(v SpBindingEnum) {
	o.SpBinding = &v
}

// GetDefaultRelayState returns the DefaultRelayState field value if set, zero value otherwise.
func (o *PatchedSAMLProviderRequest) GetDefaultRelayState() string {
	if o == nil || o.DefaultRelayState == nil {
		var ret string
		return ret
	}
	return *o.DefaultRelayState
}

// GetDefaultRelayStateOk returns a tuple with the DefaultRelayState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLProviderRequest) GetDefaultRelayStateOk() (*string, bool) {
	if o == nil || o.DefaultRelayState == nil {
		return nil, false
	}
	return o.DefaultRelayState, true
}

// HasDefaultRelayState returns a boolean if a field has been set.
func (o *PatchedSAMLProviderRequest) HasDefaultRelayState() bool {
	if o != nil && o.DefaultRelayState != nil {
		return true
	}

	return false
}

// SetDefaultRelayState gets a reference to the given string and assigns it to the DefaultRelayState field.
func (o *PatchedSAMLProviderRequest) SetDefaultRelayState(v string) {
	o.DefaultRelayState = &v
}

func (o PatchedSAMLProviderRequest) MarshalJSON() ([]byte, error) {
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
	if o.AcsUrl != nil {
		toSerialize["acs_url"] = o.AcsUrl
	}
	if o.Audience != nil {
		toSerialize["audience"] = o.Audience
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.AssertionValidNotBefore != nil {
		toSerialize["assertion_valid_not_before"] = o.AssertionValidNotBefore
	}
	if o.AssertionValidNotOnOrAfter != nil {
		toSerialize["assertion_valid_not_on_or_after"] = o.AssertionValidNotOnOrAfter
	}
	if o.SessionValidNotOnOrAfter != nil {
		toSerialize["session_valid_not_on_or_after"] = o.SessionValidNotOnOrAfter
	}
	if o.NameIdMapping.IsSet() {
		toSerialize["name_id_mapping"] = o.NameIdMapping.Get()
	}
	if o.DigestAlgorithm != nil {
		toSerialize["digest_algorithm"] = o.DigestAlgorithm
	}
	if o.SignatureAlgorithm != nil {
		toSerialize["signature_algorithm"] = o.SignatureAlgorithm
	}
	if o.SigningKp.IsSet() {
		toSerialize["signing_kp"] = o.SigningKp.Get()
	}
	if o.VerificationKp.IsSet() {
		toSerialize["verification_kp"] = o.VerificationKp.Get()
	}
	if o.SpBinding != nil {
		toSerialize["sp_binding"] = o.SpBinding
	}
	if o.DefaultRelayState != nil {
		toSerialize["default_relay_state"] = o.DefaultRelayState
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedSAMLProviderRequest struct {
	value *PatchedSAMLProviderRequest
	isSet bool
}

func (v NullablePatchedSAMLProviderRequest) Get() *PatchedSAMLProviderRequest {
	return v.value
}

func (v *NullablePatchedSAMLProviderRequest) Set(val *PatchedSAMLProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedSAMLProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedSAMLProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedSAMLProviderRequest(val *PatchedSAMLProviderRequest) *NullablePatchedSAMLProviderRequest {
	return &NullablePatchedSAMLProviderRequest{value: val, isSet: true}
}

func (v NullablePatchedSAMLProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedSAMLProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
