/*
authentik

Making authentication simple.

API version: 2023.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SAMLSourceRequest SAMLSource Serializer
type SAMLSourceRequest struct {
	// Source's display Name.
	Name string `json:"name"`
	// Internal source name, used in URLs.
	Slug    string `json:"slug"`
	Enabled *bool  `json:"enabled,omitempty"`
	// Flow to use when authenticating existing users.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow to use when enrolling new users.
	EnrollmentFlow   NullableString    `json:"enrollment_flow,omitempty"`
	PolicyEngineMode *PolicyEngineMode `json:"policy_engine_mode,omitempty"`
	// How the source determines if an existing user should be authenticated or a new user enrolled.  * `identifier` - Use the source-specific identifier * `email_link` - Link to a user with identical email address. Can have security implications when a source doesn't validate email addresses. * `email_deny` - Use the user's email address, but deny enrollment when the email address already exists. * `username_link` - Link to a user with identical username. Can have security implications when a username is used with another source. * `username_deny` - Use the user's username, but deny enrollment when the username already exists.
	UserMatchingMode *UserMatchingModeEnum `json:"user_matching_mode,omitempty"`
	UserPathTemplate *string               `json:"user_path_template,omitempty"`
	// Flow used before authentication.
	PreAuthenticationFlow string `json:"pre_authentication_flow"`
	// Also known as Entity ID. Defaults the Metadata URL.
	Issuer *string `json:"issuer,omitempty"`
	// URL that the initial Login request is sent to.
	SsoUrl string `json:"sso_url"`
	// Optional URL if your IDP supports Single-Logout.
	SloUrl NullableString `json:"slo_url,omitempty"`
	// Allows authentication flows initiated by the IdP. This can be a security risk, as no validation of the request ID is done.
	AllowIdpInitiated *bool `json:"allow_idp_initiated,omitempty"`
	// NameID Policy sent to the IdP. Can be unset, in which case no Policy is sent.  * `urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress` - Email * `urn:oasis:names:tc:SAML:2.0:nameid-format:persistent` - Persistent * `urn:oasis:names:tc:SAML:2.0:nameid-format:X509SubjectName` - X509 * `urn:oasis:names:tc:SAML:2.0:nameid-format:WindowsDomainQualifiedName` - Windows * `urn:oasis:names:tc:SAML:2.0:nameid-format:transient` - Transient
	NameIdPolicy *NameIdPolicyEnum `json:"name_id_policy,omitempty"`
	BindingType  *BindingTypeEnum  `json:"binding_type,omitempty"`
	// When selected, incoming assertion's Signatures will be validated against this certificate. To allow unsigned Requests, leave on default.
	VerificationKp NullableString `json:"verification_kp,omitempty"`
	// Keypair used to sign outgoing Responses going to the Identity Provider.
	SigningKp          NullableString          `json:"signing_kp,omitempty"`
	DigestAlgorithm    *DigestAlgorithmEnum    `json:"digest_algorithm,omitempty"`
	SignatureAlgorithm *SignatureAlgorithmEnum `json:"signature_algorithm,omitempty"`
	// Time offset when temporary users should be deleted. This only applies if your IDP uses the NameID Format 'transient', and the user doesn't log out manually. (Format: hours=1;minutes=2;seconds=3).
	TemporaryUserDeleteAfter *string `json:"temporary_user_delete_after,omitempty"`
}

// NewSAMLSourceRequest instantiates a new SAMLSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLSourceRequest(name string, slug string, preAuthenticationFlow string, ssoUrl string) *SAMLSourceRequest {
	this := SAMLSourceRequest{}
	this.Name = name
	this.Slug = slug
	this.PreAuthenticationFlow = preAuthenticationFlow
	this.SsoUrl = ssoUrl
	return &this
}

// NewSAMLSourceRequestWithDefaults instantiates a new SAMLSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLSourceRequestWithDefaults() *SAMLSourceRequest {
	this := SAMLSourceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *SAMLSourceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SAMLSourceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SAMLSourceRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *SAMLSourceRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *SAMLSourceRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *SAMLSourceRequest) SetSlug(v string) {
	o.Slug = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SAMLSourceRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSourceRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SAMLSourceRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SAMLSourceRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLSourceRequest) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SAMLSourceRequest) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *SAMLSourceRequest) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *SAMLSourceRequest) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *SAMLSourceRequest) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *SAMLSourceRequest) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetEnrollmentFlow returns the EnrollmentFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLSourceRequest) GetEnrollmentFlow() string {
	if o == nil || o.EnrollmentFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentFlow.Get()
}

// GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SAMLSourceRequest) GetEnrollmentFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentFlow.Get(), o.EnrollmentFlow.IsSet()
}

// HasEnrollmentFlow returns a boolean if a field has been set.
func (o *SAMLSourceRequest) HasEnrollmentFlow() bool {
	if o != nil && o.EnrollmentFlow.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentFlow gets a reference to the given NullableString and assigns it to the EnrollmentFlow field.
func (o *SAMLSourceRequest) SetEnrollmentFlow(v string) {
	o.EnrollmentFlow.Set(&v)
}

// SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil
func (o *SAMLSourceRequest) SetEnrollmentFlowNil() {
	o.EnrollmentFlow.Set(nil)
}

// UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
func (o *SAMLSourceRequest) UnsetEnrollmentFlow() {
	o.EnrollmentFlow.Unset()
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *SAMLSourceRequest) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSourceRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *SAMLSourceRequest) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *SAMLSourceRequest) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetUserMatchingMode returns the UserMatchingMode field value if set, zero value otherwise.
func (o *SAMLSourceRequest) GetUserMatchingMode() UserMatchingModeEnum {
	if o == nil || o.UserMatchingMode == nil {
		var ret UserMatchingModeEnum
		return ret
	}
	return *o.UserMatchingMode
}

// GetUserMatchingModeOk returns a tuple with the UserMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool) {
	if o == nil || o.UserMatchingMode == nil {
		return nil, false
	}
	return o.UserMatchingMode, true
}

// HasUserMatchingMode returns a boolean if a field has been set.
func (o *SAMLSourceRequest) HasUserMatchingMode() bool {
	if o != nil && o.UserMatchingMode != nil {
		return true
	}

	return false
}

// SetUserMatchingMode gets a reference to the given UserMatchingModeEnum and assigns it to the UserMatchingMode field.
func (o *SAMLSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum) {
	o.UserMatchingMode = &v
}

// GetUserPathTemplate returns the UserPathTemplate field value if set, zero value otherwise.
func (o *SAMLSourceRequest) GetUserPathTemplate() string {
	if o == nil || o.UserPathTemplate == nil {
		var ret string
		return ret
	}
	return *o.UserPathTemplate
}

// GetUserPathTemplateOk returns a tuple with the UserPathTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSourceRequest) GetUserPathTemplateOk() (*string, bool) {
	if o == nil || o.UserPathTemplate == nil {
		return nil, false
	}
	return o.UserPathTemplate, true
}

// HasUserPathTemplate returns a boolean if a field has been set.
func (o *SAMLSourceRequest) HasUserPathTemplate() bool {
	if o != nil && o.UserPathTemplate != nil {
		return true
	}

	return false
}

// SetUserPathTemplate gets a reference to the given string and assigns it to the UserPathTemplate field.
func (o *SAMLSourceRequest) SetUserPathTemplate(v string) {
	o.UserPathTemplate = &v
}

// GetPreAuthenticationFlow returns the PreAuthenticationFlow field value
func (o *SAMLSourceRequest) GetPreAuthenticationFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreAuthenticationFlow
}

// GetPreAuthenticationFlowOk returns a tuple with the PreAuthenticationFlow field value
// and a boolean to check if the value has been set.
func (o *SAMLSourceRequest) GetPreAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreAuthenticationFlow, true
}

// SetPreAuthenticationFlow sets field value
func (o *SAMLSourceRequest) SetPreAuthenticationFlow(v string) {
	o.PreAuthenticationFlow = v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *SAMLSourceRequest) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSourceRequest) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *SAMLSourceRequest) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *SAMLSourceRequest) SetIssuer(v string) {
	o.Issuer = &v
}

// GetSsoUrl returns the SsoUrl field value
func (o *SAMLSourceRequest) GetSsoUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SsoUrl
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value
// and a boolean to check if the value has been set.
func (o *SAMLSourceRequest) GetSsoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SsoUrl, true
}

// SetSsoUrl sets field value
func (o *SAMLSourceRequest) SetSsoUrl(v string) {
	o.SsoUrl = v
}

// GetSloUrl returns the SloUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLSourceRequest) GetSloUrl() string {
	if o == nil || o.SloUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.SloUrl.Get()
}

// GetSloUrlOk returns a tuple with the SloUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SAMLSourceRequest) GetSloUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SloUrl.Get(), o.SloUrl.IsSet()
}

// HasSloUrl returns a boolean if a field has been set.
func (o *SAMLSourceRequest) HasSloUrl() bool {
	if o != nil && o.SloUrl.IsSet() {
		return true
	}

	return false
}

// SetSloUrl gets a reference to the given NullableString and assigns it to the SloUrl field.
func (o *SAMLSourceRequest) SetSloUrl(v string) {
	o.SloUrl.Set(&v)
}

// SetSloUrlNil sets the value for SloUrl to be an explicit nil
func (o *SAMLSourceRequest) SetSloUrlNil() {
	o.SloUrl.Set(nil)
}

// UnsetSloUrl ensures that no value is present for SloUrl, not even an explicit nil
func (o *SAMLSourceRequest) UnsetSloUrl() {
	o.SloUrl.Unset()
}

// GetAllowIdpInitiated returns the AllowIdpInitiated field value if set, zero value otherwise.
func (o *SAMLSourceRequest) GetAllowIdpInitiated() bool {
	if o == nil || o.AllowIdpInitiated == nil {
		var ret bool
		return ret
	}
	return *o.AllowIdpInitiated
}

// GetAllowIdpInitiatedOk returns a tuple with the AllowIdpInitiated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSourceRequest) GetAllowIdpInitiatedOk() (*bool, bool) {
	if o == nil || o.AllowIdpInitiated == nil {
		return nil, false
	}
	return o.AllowIdpInitiated, true
}

// HasAllowIdpInitiated returns a boolean if a field has been set.
func (o *SAMLSourceRequest) HasAllowIdpInitiated() bool {
	if o != nil && o.AllowIdpInitiated != nil {
		return true
	}

	return false
}

// SetAllowIdpInitiated gets a reference to the given bool and assigns it to the AllowIdpInitiated field.
func (o *SAMLSourceRequest) SetAllowIdpInitiated(v bool) {
	o.AllowIdpInitiated = &v
}

// GetNameIdPolicy returns the NameIdPolicy field value if set, zero value otherwise.
func (o *SAMLSourceRequest) GetNameIdPolicy() NameIdPolicyEnum {
	if o == nil || o.NameIdPolicy == nil {
		var ret NameIdPolicyEnum
		return ret
	}
	return *o.NameIdPolicy
}

// GetNameIdPolicyOk returns a tuple with the NameIdPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSourceRequest) GetNameIdPolicyOk() (*NameIdPolicyEnum, bool) {
	if o == nil || o.NameIdPolicy == nil {
		return nil, false
	}
	return o.NameIdPolicy, true
}

// HasNameIdPolicy returns a boolean if a field has been set.
func (o *SAMLSourceRequest) HasNameIdPolicy() bool {
	if o != nil && o.NameIdPolicy != nil {
		return true
	}

	return false
}

// SetNameIdPolicy gets a reference to the given NameIdPolicyEnum and assigns it to the NameIdPolicy field.
func (o *SAMLSourceRequest) SetNameIdPolicy(v NameIdPolicyEnum) {
	o.NameIdPolicy = &v
}

// GetBindingType returns the BindingType field value if set, zero value otherwise.
func (o *SAMLSourceRequest) GetBindingType() BindingTypeEnum {
	if o == nil || o.BindingType == nil {
		var ret BindingTypeEnum
		return ret
	}
	return *o.BindingType
}

// GetBindingTypeOk returns a tuple with the BindingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSourceRequest) GetBindingTypeOk() (*BindingTypeEnum, bool) {
	if o == nil || o.BindingType == nil {
		return nil, false
	}
	return o.BindingType, true
}

// HasBindingType returns a boolean if a field has been set.
func (o *SAMLSourceRequest) HasBindingType() bool {
	if o != nil && o.BindingType != nil {
		return true
	}

	return false
}

// SetBindingType gets a reference to the given BindingTypeEnum and assigns it to the BindingType field.
func (o *SAMLSourceRequest) SetBindingType(v BindingTypeEnum) {
	o.BindingType = &v
}

// GetVerificationKp returns the VerificationKp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLSourceRequest) GetVerificationKp() string {
	if o == nil || o.VerificationKp.Get() == nil {
		var ret string
		return ret
	}
	return *o.VerificationKp.Get()
}

// GetVerificationKpOk returns a tuple with the VerificationKp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SAMLSourceRequest) GetVerificationKpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VerificationKp.Get(), o.VerificationKp.IsSet()
}

// HasVerificationKp returns a boolean if a field has been set.
func (o *SAMLSourceRequest) HasVerificationKp() bool {
	if o != nil && o.VerificationKp.IsSet() {
		return true
	}

	return false
}

// SetVerificationKp gets a reference to the given NullableString and assigns it to the VerificationKp field.
func (o *SAMLSourceRequest) SetVerificationKp(v string) {
	o.VerificationKp.Set(&v)
}

// SetVerificationKpNil sets the value for VerificationKp to be an explicit nil
func (o *SAMLSourceRequest) SetVerificationKpNil() {
	o.VerificationKp.Set(nil)
}

// UnsetVerificationKp ensures that no value is present for VerificationKp, not even an explicit nil
func (o *SAMLSourceRequest) UnsetVerificationKp() {
	o.VerificationKp.Unset()
}

// GetSigningKp returns the SigningKp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLSourceRequest) GetSigningKp() string {
	if o == nil || o.SigningKp.Get() == nil {
		var ret string
		return ret
	}
	return *o.SigningKp.Get()
}

// GetSigningKpOk returns a tuple with the SigningKp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SAMLSourceRequest) GetSigningKpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SigningKp.Get(), o.SigningKp.IsSet()
}

// HasSigningKp returns a boolean if a field has been set.
func (o *SAMLSourceRequest) HasSigningKp() bool {
	if o != nil && o.SigningKp.IsSet() {
		return true
	}

	return false
}

// SetSigningKp gets a reference to the given NullableString and assigns it to the SigningKp field.
func (o *SAMLSourceRequest) SetSigningKp(v string) {
	o.SigningKp.Set(&v)
}

// SetSigningKpNil sets the value for SigningKp to be an explicit nil
func (o *SAMLSourceRequest) SetSigningKpNil() {
	o.SigningKp.Set(nil)
}

// UnsetSigningKp ensures that no value is present for SigningKp, not even an explicit nil
func (o *SAMLSourceRequest) UnsetSigningKp() {
	o.SigningKp.Unset()
}

// GetDigestAlgorithm returns the DigestAlgorithm field value if set, zero value otherwise.
func (o *SAMLSourceRequest) GetDigestAlgorithm() DigestAlgorithmEnum {
	if o == nil || o.DigestAlgorithm == nil {
		var ret DigestAlgorithmEnum
		return ret
	}
	return *o.DigestAlgorithm
}

// GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSourceRequest) GetDigestAlgorithmOk() (*DigestAlgorithmEnum, bool) {
	if o == nil || o.DigestAlgorithm == nil {
		return nil, false
	}
	return o.DigestAlgorithm, true
}

// HasDigestAlgorithm returns a boolean if a field has been set.
func (o *SAMLSourceRequest) HasDigestAlgorithm() bool {
	if o != nil && o.DigestAlgorithm != nil {
		return true
	}

	return false
}

// SetDigestAlgorithm gets a reference to the given DigestAlgorithmEnum and assigns it to the DigestAlgorithm field.
func (o *SAMLSourceRequest) SetDigestAlgorithm(v DigestAlgorithmEnum) {
	o.DigestAlgorithm = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *SAMLSourceRequest) GetSignatureAlgorithm() SignatureAlgorithmEnum {
	if o == nil || o.SignatureAlgorithm == nil {
		var ret SignatureAlgorithmEnum
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSourceRequest) GetSignatureAlgorithmOk() (*SignatureAlgorithmEnum, bool) {
	if o == nil || o.SignatureAlgorithm == nil {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *SAMLSourceRequest) HasSignatureAlgorithm() bool {
	if o != nil && o.SignatureAlgorithm != nil {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given SignatureAlgorithmEnum and assigns it to the SignatureAlgorithm field.
func (o *SAMLSourceRequest) SetSignatureAlgorithm(v SignatureAlgorithmEnum) {
	o.SignatureAlgorithm = &v
}

// GetTemporaryUserDeleteAfter returns the TemporaryUserDeleteAfter field value if set, zero value otherwise.
func (o *SAMLSourceRequest) GetTemporaryUserDeleteAfter() string {
	if o == nil || o.TemporaryUserDeleteAfter == nil {
		var ret string
		return ret
	}
	return *o.TemporaryUserDeleteAfter
}

// GetTemporaryUserDeleteAfterOk returns a tuple with the TemporaryUserDeleteAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSourceRequest) GetTemporaryUserDeleteAfterOk() (*string, bool) {
	if o == nil || o.TemporaryUserDeleteAfter == nil {
		return nil, false
	}
	return o.TemporaryUserDeleteAfter, true
}

// HasTemporaryUserDeleteAfter returns a boolean if a field has been set.
func (o *SAMLSourceRequest) HasTemporaryUserDeleteAfter() bool {
	if o != nil && o.TemporaryUserDeleteAfter != nil {
		return true
	}

	return false
}

// SetTemporaryUserDeleteAfter gets a reference to the given string and assigns it to the TemporaryUserDeleteAfter field.
func (o *SAMLSourceRequest) SetTemporaryUserDeleteAfter(v string) {
	o.TemporaryUserDeleteAfter = &v
}

func (o SAMLSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.PolicyEngineMode != nil {
		toSerialize["policy_engine_mode"] = o.PolicyEngineMode
	}
	if o.UserMatchingMode != nil {
		toSerialize["user_matching_mode"] = o.UserMatchingMode
	}
	if o.UserPathTemplate != nil {
		toSerialize["user_path_template"] = o.UserPathTemplate
	}
	if true {
		toSerialize["pre_authentication_flow"] = o.PreAuthenticationFlow
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if true {
		toSerialize["sso_url"] = o.SsoUrl
	}
	if o.SloUrl.IsSet() {
		toSerialize["slo_url"] = o.SloUrl.Get()
	}
	if o.AllowIdpInitiated != nil {
		toSerialize["allow_idp_initiated"] = o.AllowIdpInitiated
	}
	if o.NameIdPolicy != nil {
		toSerialize["name_id_policy"] = o.NameIdPolicy
	}
	if o.BindingType != nil {
		toSerialize["binding_type"] = o.BindingType
	}
	if o.VerificationKp.IsSet() {
		toSerialize["verification_kp"] = o.VerificationKp.Get()
	}
	if o.SigningKp.IsSet() {
		toSerialize["signing_kp"] = o.SigningKp.Get()
	}
	if o.DigestAlgorithm != nil {
		toSerialize["digest_algorithm"] = o.DigestAlgorithm
	}
	if o.SignatureAlgorithm != nil {
		toSerialize["signature_algorithm"] = o.SignatureAlgorithm
	}
	if o.TemporaryUserDeleteAfter != nil {
		toSerialize["temporary_user_delete_after"] = o.TemporaryUserDeleteAfter
	}
	return json.Marshal(toSerialize)
}

type NullableSAMLSourceRequest struct {
	value *SAMLSourceRequest
	isSet bool
}

func (v NullableSAMLSourceRequest) Get() *SAMLSourceRequest {
	return v.value
}

func (v *NullableSAMLSourceRequest) Set(val *SAMLSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLSourceRequest(val *SAMLSourceRequest) *NullableSAMLSourceRequest {
	return &NullableSAMLSourceRequest{value: val, isSet: true}
}

func (v NullableSAMLSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
