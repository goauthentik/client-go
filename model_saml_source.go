/*
authentik

Making authentication simple.

API version: 2021.9.1-rc3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SAMLSource SAMLSource Serializer
type SAMLSource struct {
	Pk string `json:"pk"`
	// Source's display Name.
	Name string `json:"name"`
	// Internal source name, used in URLs.
	Slug string `json:"slug"`
	Enabled *bool `json:"enabled,omitempty"`
	// Flow to use when authenticating existing users.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow to use when enrolling new users.
	EnrollmentFlow NullableString `json:"enrollment_flow,omitempty"`
	Component string `json:"component"`
	VerboseName string `json:"verbose_name"`
	VerboseNamePlural string `json:"verbose_name_plural"`
	PolicyEngineMode *PolicyEngineMode `json:"policy_engine_mode,omitempty"`
	// How the source determines if an existing user should be authenticated or a new user enrolled.
	UserMatchingMode *UserMatchingModeEnum `json:"user_matching_mode,omitempty"`
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
	// NameID Policy sent to the IdP. Can be unset, in which case no Policy is sent.
	NameIdPolicy *NameIdPolicyEnum `json:"name_id_policy,omitempty"`
	BindingType *BindingTypeEnum `json:"binding_type,omitempty"`
	// Keypair which is used to sign outgoing requests. Leave empty to disable signing.
	SigningKp NullableString `json:"signing_kp,omitempty"`
	DigestAlgorithm *DigestAlgorithmEnum `json:"digest_algorithm,omitempty"`
	SignatureAlgorithm *SignatureAlgorithmEnum `json:"signature_algorithm,omitempty"`
	// Time offset when temporary users should be deleted. This only applies if your IDP uses the NameID Format 'transient', and the user doesn't log out manually. (Format: hours=1;minutes=2;seconds=3).
	TemporaryUserDeleteAfter *string `json:"temporary_user_delete_after,omitempty"`
}

// NewSAMLSource instantiates a new SAMLSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLSource(pk string, name string, slug string, component string, verboseName string, verboseNamePlural string, preAuthenticationFlow string, ssoUrl string) *SAMLSource {
	this := SAMLSource{}
	this.Pk = pk
	this.Name = name
	this.Slug = slug
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.PreAuthenticationFlow = preAuthenticationFlow
	this.SsoUrl = ssoUrl
	return &this
}

// NewSAMLSourceWithDefaults instantiates a new SAMLSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLSourceWithDefaults() *SAMLSource {
	this := SAMLSource{}
	return &this
}

// GetPk returns the Pk field value
func (o *SAMLSource) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetPkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *SAMLSource) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *SAMLSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SAMLSource) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *SAMLSource) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetSlugOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *SAMLSource) SetSlug(v string) {
	o.Slug = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SAMLSource) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SAMLSource) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SAMLSource) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLSource) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SAMLSource) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *SAMLSource) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *SAMLSource) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}
// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *SAMLSource) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *SAMLSource) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetEnrollmentFlow returns the EnrollmentFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLSource) GetEnrollmentFlow() string {
	if o == nil || o.EnrollmentFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentFlow.Get()
}

// GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SAMLSource) GetEnrollmentFlowOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EnrollmentFlow.Get(), o.EnrollmentFlow.IsSet()
}

// HasEnrollmentFlow returns a boolean if a field has been set.
func (o *SAMLSource) HasEnrollmentFlow() bool {
	if o != nil && o.EnrollmentFlow.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentFlow gets a reference to the given NullableString and assigns it to the EnrollmentFlow field.
func (o *SAMLSource) SetEnrollmentFlow(v string) {
	o.EnrollmentFlow.Set(&v)
}
// SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil
func (o *SAMLSource) SetEnrollmentFlowNil() {
	o.EnrollmentFlow.Set(nil)
}

// UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
func (o *SAMLSource) UnsetEnrollmentFlow() {
	o.EnrollmentFlow.Unset()
}

// GetComponent returns the Component field value
func (o *SAMLSource) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetComponentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *SAMLSource) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *SAMLSource) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetVerboseNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *SAMLSource) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *SAMLSource) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *SAMLSource) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *SAMLSource) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *SAMLSource) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *SAMLSource) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetUserMatchingMode returns the UserMatchingMode field value if set, zero value otherwise.
func (o *SAMLSource) GetUserMatchingMode() UserMatchingModeEnum {
	if o == nil || o.UserMatchingMode == nil {
		var ret UserMatchingModeEnum
		return ret
	}
	return *o.UserMatchingMode
}

// GetUserMatchingModeOk returns a tuple with the UserMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool) {
	if o == nil || o.UserMatchingMode == nil {
		return nil, false
	}
	return o.UserMatchingMode, true
}

// HasUserMatchingMode returns a boolean if a field has been set.
func (o *SAMLSource) HasUserMatchingMode() bool {
	if o != nil && o.UserMatchingMode != nil {
		return true
	}

	return false
}

// SetUserMatchingMode gets a reference to the given UserMatchingModeEnum and assigns it to the UserMatchingMode field.
func (o *SAMLSource) SetUserMatchingMode(v UserMatchingModeEnum) {
	o.UserMatchingMode = &v
}

// GetPreAuthenticationFlow returns the PreAuthenticationFlow field value
func (o *SAMLSource) GetPreAuthenticationFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreAuthenticationFlow
}

// GetPreAuthenticationFlowOk returns a tuple with the PreAuthenticationFlow field value
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetPreAuthenticationFlowOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PreAuthenticationFlow, true
}

// SetPreAuthenticationFlow sets field value
func (o *SAMLSource) SetPreAuthenticationFlow(v string) {
	o.PreAuthenticationFlow = v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *SAMLSource) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *SAMLSource) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *SAMLSource) SetIssuer(v string) {
	o.Issuer = &v
}

// GetSsoUrl returns the SsoUrl field value
func (o *SAMLSource) GetSsoUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SsoUrl
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetSsoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SsoUrl, true
}

// SetSsoUrl sets field value
func (o *SAMLSource) SetSsoUrl(v string) {
	o.SsoUrl = v
}

// GetSloUrl returns the SloUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLSource) GetSloUrl() string {
	if o == nil || o.SloUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.SloUrl.Get()
}

// GetSloUrlOk returns a tuple with the SloUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SAMLSource) GetSloUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SloUrl.Get(), o.SloUrl.IsSet()
}

// HasSloUrl returns a boolean if a field has been set.
func (o *SAMLSource) HasSloUrl() bool {
	if o != nil && o.SloUrl.IsSet() {
		return true
	}

	return false
}

// SetSloUrl gets a reference to the given NullableString and assigns it to the SloUrl field.
func (o *SAMLSource) SetSloUrl(v string) {
	o.SloUrl.Set(&v)
}
// SetSloUrlNil sets the value for SloUrl to be an explicit nil
func (o *SAMLSource) SetSloUrlNil() {
	o.SloUrl.Set(nil)
}

// UnsetSloUrl ensures that no value is present for SloUrl, not even an explicit nil
func (o *SAMLSource) UnsetSloUrl() {
	o.SloUrl.Unset()
}

// GetAllowIdpInitiated returns the AllowIdpInitiated field value if set, zero value otherwise.
func (o *SAMLSource) GetAllowIdpInitiated() bool {
	if o == nil || o.AllowIdpInitiated == nil {
		var ret bool
		return ret
	}
	return *o.AllowIdpInitiated
}

// GetAllowIdpInitiatedOk returns a tuple with the AllowIdpInitiated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetAllowIdpInitiatedOk() (*bool, bool) {
	if o == nil || o.AllowIdpInitiated == nil {
		return nil, false
	}
	return o.AllowIdpInitiated, true
}

// HasAllowIdpInitiated returns a boolean if a field has been set.
func (o *SAMLSource) HasAllowIdpInitiated() bool {
	if o != nil && o.AllowIdpInitiated != nil {
		return true
	}

	return false
}

// SetAllowIdpInitiated gets a reference to the given bool and assigns it to the AllowIdpInitiated field.
func (o *SAMLSource) SetAllowIdpInitiated(v bool) {
	o.AllowIdpInitiated = &v
}

// GetNameIdPolicy returns the NameIdPolicy field value if set, zero value otherwise.
func (o *SAMLSource) GetNameIdPolicy() NameIdPolicyEnum {
	if o == nil || o.NameIdPolicy == nil {
		var ret NameIdPolicyEnum
		return ret
	}
	return *o.NameIdPolicy
}

// GetNameIdPolicyOk returns a tuple with the NameIdPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetNameIdPolicyOk() (*NameIdPolicyEnum, bool) {
	if o == nil || o.NameIdPolicy == nil {
		return nil, false
	}
	return o.NameIdPolicy, true
}

// HasNameIdPolicy returns a boolean if a field has been set.
func (o *SAMLSource) HasNameIdPolicy() bool {
	if o != nil && o.NameIdPolicy != nil {
		return true
	}

	return false
}

// SetNameIdPolicy gets a reference to the given NameIdPolicyEnum and assigns it to the NameIdPolicy field.
func (o *SAMLSource) SetNameIdPolicy(v NameIdPolicyEnum) {
	o.NameIdPolicy = &v
}

// GetBindingType returns the BindingType field value if set, zero value otherwise.
func (o *SAMLSource) GetBindingType() BindingTypeEnum {
	if o == nil || o.BindingType == nil {
		var ret BindingTypeEnum
		return ret
	}
	return *o.BindingType
}

// GetBindingTypeOk returns a tuple with the BindingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetBindingTypeOk() (*BindingTypeEnum, bool) {
	if o == nil || o.BindingType == nil {
		return nil, false
	}
	return o.BindingType, true
}

// HasBindingType returns a boolean if a field has been set.
func (o *SAMLSource) HasBindingType() bool {
	if o != nil && o.BindingType != nil {
		return true
	}

	return false
}

// SetBindingType gets a reference to the given BindingTypeEnum and assigns it to the BindingType field.
func (o *SAMLSource) SetBindingType(v BindingTypeEnum) {
	o.BindingType = &v
}

// GetSigningKp returns the SigningKp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLSource) GetSigningKp() string {
	if o == nil || o.SigningKp.Get() == nil {
		var ret string
		return ret
	}
	return *o.SigningKp.Get()
}

// GetSigningKpOk returns a tuple with the SigningKp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SAMLSource) GetSigningKpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SigningKp.Get(), o.SigningKp.IsSet()
}

// HasSigningKp returns a boolean if a field has been set.
func (o *SAMLSource) HasSigningKp() bool {
	if o != nil && o.SigningKp.IsSet() {
		return true
	}

	return false
}

// SetSigningKp gets a reference to the given NullableString and assigns it to the SigningKp field.
func (o *SAMLSource) SetSigningKp(v string) {
	o.SigningKp.Set(&v)
}
// SetSigningKpNil sets the value for SigningKp to be an explicit nil
func (o *SAMLSource) SetSigningKpNil() {
	o.SigningKp.Set(nil)
}

// UnsetSigningKp ensures that no value is present for SigningKp, not even an explicit nil
func (o *SAMLSource) UnsetSigningKp() {
	o.SigningKp.Unset()
}

// GetDigestAlgorithm returns the DigestAlgorithm field value if set, zero value otherwise.
func (o *SAMLSource) GetDigestAlgorithm() DigestAlgorithmEnum {
	if o == nil || o.DigestAlgorithm == nil {
		var ret DigestAlgorithmEnum
		return ret
	}
	return *o.DigestAlgorithm
}

// GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetDigestAlgorithmOk() (*DigestAlgorithmEnum, bool) {
	if o == nil || o.DigestAlgorithm == nil {
		return nil, false
	}
	return o.DigestAlgorithm, true
}

// HasDigestAlgorithm returns a boolean if a field has been set.
func (o *SAMLSource) HasDigestAlgorithm() bool {
	if o != nil && o.DigestAlgorithm != nil {
		return true
	}

	return false
}

// SetDigestAlgorithm gets a reference to the given DigestAlgorithmEnum and assigns it to the DigestAlgorithm field.
func (o *SAMLSource) SetDigestAlgorithm(v DigestAlgorithmEnum) {
	o.DigestAlgorithm = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *SAMLSource) GetSignatureAlgorithm() SignatureAlgorithmEnum {
	if o == nil || o.SignatureAlgorithm == nil {
		var ret SignatureAlgorithmEnum
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetSignatureAlgorithmOk() (*SignatureAlgorithmEnum, bool) {
	if o == nil || o.SignatureAlgorithm == nil {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *SAMLSource) HasSignatureAlgorithm() bool {
	if o != nil && o.SignatureAlgorithm != nil {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given SignatureAlgorithmEnum and assigns it to the SignatureAlgorithm field.
func (o *SAMLSource) SetSignatureAlgorithm(v SignatureAlgorithmEnum) {
	o.SignatureAlgorithm = &v
}

// GetTemporaryUserDeleteAfter returns the TemporaryUserDeleteAfter field value if set, zero value otherwise.
func (o *SAMLSource) GetTemporaryUserDeleteAfter() string {
	if o == nil || o.TemporaryUserDeleteAfter == nil {
		var ret string
		return ret
	}
	return *o.TemporaryUserDeleteAfter
}

// GetTemporaryUserDeleteAfterOk returns a tuple with the TemporaryUserDeleteAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLSource) GetTemporaryUserDeleteAfterOk() (*string, bool) {
	if o == nil || o.TemporaryUserDeleteAfter == nil {
		return nil, false
	}
	return o.TemporaryUserDeleteAfter, true
}

// HasTemporaryUserDeleteAfter returns a boolean if a field has been set.
func (o *SAMLSource) HasTemporaryUserDeleteAfter() bool {
	if o != nil && o.TemporaryUserDeleteAfter != nil {
		return true
	}

	return false
}

// SetTemporaryUserDeleteAfter gets a reference to the given string and assigns it to the TemporaryUserDeleteAfter field.
func (o *SAMLSource) SetTemporaryUserDeleteAfter(v string) {
	o.TemporaryUserDeleteAfter = &v
}

func (o SAMLSource) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["verbose_name"] = o.VerboseName
	}
	if true {
		toSerialize["verbose_name_plural"] = o.VerboseNamePlural
	}
	if o.PolicyEngineMode != nil {
		toSerialize["policy_engine_mode"] = o.PolicyEngineMode
	}
	if o.UserMatchingMode != nil {
		toSerialize["user_matching_mode"] = o.UserMatchingMode
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

type NullableSAMLSource struct {
	value *SAMLSource
	isSet bool
}

func (v NullableSAMLSource) Get() *SAMLSource {
	return v.value
}

func (v *NullableSAMLSource) Set(val *SAMLSource) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLSource(val *SAMLSource) *NullableSAMLSource {
	return &NullableSAMLSource{value: val, isSet: true}
}

func (v NullableSAMLSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


