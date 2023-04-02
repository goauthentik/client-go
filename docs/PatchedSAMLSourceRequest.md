# PatchedSAMLSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Source&#39;s display Name. | [optional] 
**Slug** | Pointer to **string** | Internal source name, used in URLs. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** | Flow to use when authenticating existing users. | [optional] 
**EnrollmentFlow** | Pointer to **NullableString** | Flow to use when enrolling new users. | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**UserMatchingMode** | Pointer to [**UserMatchingModeEnum**](UserMatchingModeEnum.md) | How the source determines if an existing user should be authenticated or a new user enrolled.  * &#x60;identifier&#x60; - Use the source-specific identifier * &#x60;email_link&#x60; - Link to a user with identical email address. Can have security implications when a source doesn&#39;t validate email addresses. * &#x60;email_deny&#x60; - Use the user&#39;s email address, but deny enrollment when the email address already exists. * &#x60;username_link&#x60; - Link to a user with identical username. Can have security implications when a username is used with another source. * &#x60;username_deny&#x60; - Use the user&#39;s username, but deny enrollment when the username already exists. | [optional] 
**UserPathTemplate** | Pointer to **string** |  | [optional] 
**PreAuthenticationFlow** | Pointer to **string** | Flow used before authentication. | [optional] 
**Issuer** | Pointer to **string** | Also known as Entity ID. Defaults the Metadata URL. | [optional] 
**SsoUrl** | Pointer to **string** | URL that the initial Login request is sent to. | [optional] 
**SloUrl** | Pointer to **NullableString** | Optional URL if your IDP supports Single-Logout. | [optional] 
**AllowIdpInitiated** | Pointer to **bool** | Allows authentication flows initiated by the IdP. This can be a security risk, as no validation of the request ID is done. | [optional] 
**NameIdPolicy** | Pointer to [**NameIdPolicyEnum**](NameIdPolicyEnum.md) | NameID Policy sent to the IdP. Can be unset, in which case no Policy is sent.  * &#x60;urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress&#x60; - Email * &#x60;urn:oasis:names:tc:SAML:2.0:nameid-format:persistent&#x60; - Persistent * &#x60;urn:oasis:names:tc:SAML:2.0:nameid-format:X509SubjectName&#x60; - X509 * &#x60;urn:oasis:names:tc:SAML:2.0:nameid-format:WindowsDomainQualifiedName&#x60; - Windows * &#x60;urn:oasis:names:tc:SAML:2.0:nameid-format:transient&#x60; - Transient | [optional] 
**BindingType** | Pointer to [**BindingTypeEnum**](BindingTypeEnum.md) |  | [optional] 
**SigningKp** | Pointer to **NullableString** | Keypair which is used to sign outgoing requests. Leave empty to disable signing. | [optional] 
**DigestAlgorithm** | Pointer to [**DigestAlgorithmEnum**](DigestAlgorithmEnum.md) |  | [optional] 
**SignatureAlgorithm** | Pointer to [**SignatureAlgorithmEnum**](SignatureAlgorithmEnum.md) |  | [optional] 
**TemporaryUserDeleteAfter** | Pointer to **string** | Time offset when temporary users should be deleted. This only applies if your IDP uses the NameID Format &#39;transient&#39;, and the user doesn&#39;t log out manually. (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 

## Methods

### NewPatchedSAMLSourceRequest

`func NewPatchedSAMLSourceRequest() *PatchedSAMLSourceRequest`

NewPatchedSAMLSourceRequest instantiates a new PatchedSAMLSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSAMLSourceRequestWithDefaults

`func NewPatchedSAMLSourceRequestWithDefaults() *PatchedSAMLSourceRequest`

NewPatchedSAMLSourceRequestWithDefaults instantiates a new PatchedSAMLSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedSAMLSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedSAMLSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedSAMLSourceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedSAMLSourceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedSAMLSourceRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedSAMLSourceRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedSAMLSourceRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedSAMLSourceRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedSAMLSourceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedSAMLSourceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedSAMLSourceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedSAMLSourceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *PatchedSAMLSourceRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *PatchedSAMLSourceRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *PatchedSAMLSourceRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *PatchedSAMLSourceRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *PatchedSAMLSourceRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *PatchedSAMLSourceRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetEnrollmentFlow

`func (o *PatchedSAMLSourceRequest) GetEnrollmentFlow() string`

GetEnrollmentFlow returns the EnrollmentFlow field if non-nil, zero value otherwise.

### GetEnrollmentFlowOk

`func (o *PatchedSAMLSourceRequest) GetEnrollmentFlowOk() (*string, bool)`

GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlow

`func (o *PatchedSAMLSourceRequest) SetEnrollmentFlow(v string)`

SetEnrollmentFlow sets EnrollmentFlow field to given value.

### HasEnrollmentFlow

`func (o *PatchedSAMLSourceRequest) HasEnrollmentFlow() bool`

HasEnrollmentFlow returns a boolean if a field has been set.

### SetEnrollmentFlowNil

`func (o *PatchedSAMLSourceRequest) SetEnrollmentFlowNil(b bool)`

 SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil

### UnsetEnrollmentFlow
`func (o *PatchedSAMLSourceRequest) UnsetEnrollmentFlow()`

UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
### GetPolicyEngineMode

`func (o *PatchedSAMLSourceRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *PatchedSAMLSourceRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *PatchedSAMLSourceRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *PatchedSAMLSourceRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetUserMatchingMode

`func (o *PatchedSAMLSourceRequest) GetUserMatchingMode() UserMatchingModeEnum`

GetUserMatchingMode returns the UserMatchingMode field if non-nil, zero value otherwise.

### GetUserMatchingModeOk

`func (o *PatchedSAMLSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool)`

GetUserMatchingModeOk returns a tuple with the UserMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMatchingMode

`func (o *PatchedSAMLSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum)`

SetUserMatchingMode sets UserMatchingMode field to given value.

### HasUserMatchingMode

`func (o *PatchedSAMLSourceRequest) HasUserMatchingMode() bool`

HasUserMatchingMode returns a boolean if a field has been set.

### GetUserPathTemplate

`func (o *PatchedSAMLSourceRequest) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *PatchedSAMLSourceRequest) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *PatchedSAMLSourceRequest) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *PatchedSAMLSourceRequest) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.

### GetPreAuthenticationFlow

`func (o *PatchedSAMLSourceRequest) GetPreAuthenticationFlow() string`

GetPreAuthenticationFlow returns the PreAuthenticationFlow field if non-nil, zero value otherwise.

### GetPreAuthenticationFlowOk

`func (o *PatchedSAMLSourceRequest) GetPreAuthenticationFlowOk() (*string, bool)`

GetPreAuthenticationFlowOk returns a tuple with the PreAuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthenticationFlow

`func (o *PatchedSAMLSourceRequest) SetPreAuthenticationFlow(v string)`

SetPreAuthenticationFlow sets PreAuthenticationFlow field to given value.

### HasPreAuthenticationFlow

`func (o *PatchedSAMLSourceRequest) HasPreAuthenticationFlow() bool`

HasPreAuthenticationFlow returns a boolean if a field has been set.

### GetIssuer

`func (o *PatchedSAMLSourceRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *PatchedSAMLSourceRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *PatchedSAMLSourceRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *PatchedSAMLSourceRequest) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSsoUrl

`func (o *PatchedSAMLSourceRequest) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *PatchedSAMLSourceRequest) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *PatchedSAMLSourceRequest) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *PatchedSAMLSourceRequest) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.

### GetSloUrl

`func (o *PatchedSAMLSourceRequest) GetSloUrl() string`

GetSloUrl returns the SloUrl field if non-nil, zero value otherwise.

### GetSloUrlOk

`func (o *PatchedSAMLSourceRequest) GetSloUrlOk() (*string, bool)`

GetSloUrlOk returns a tuple with the SloUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloUrl

`func (o *PatchedSAMLSourceRequest) SetSloUrl(v string)`

SetSloUrl sets SloUrl field to given value.

### HasSloUrl

`func (o *PatchedSAMLSourceRequest) HasSloUrl() bool`

HasSloUrl returns a boolean if a field has been set.

### SetSloUrlNil

`func (o *PatchedSAMLSourceRequest) SetSloUrlNil(b bool)`

 SetSloUrlNil sets the value for SloUrl to be an explicit nil

### UnsetSloUrl
`func (o *PatchedSAMLSourceRequest) UnsetSloUrl()`

UnsetSloUrl ensures that no value is present for SloUrl, not even an explicit nil
### GetAllowIdpInitiated

`func (o *PatchedSAMLSourceRequest) GetAllowIdpInitiated() bool`

GetAllowIdpInitiated returns the AllowIdpInitiated field if non-nil, zero value otherwise.

### GetAllowIdpInitiatedOk

`func (o *PatchedSAMLSourceRequest) GetAllowIdpInitiatedOk() (*bool, bool)`

GetAllowIdpInitiatedOk returns a tuple with the AllowIdpInitiated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIdpInitiated

`func (o *PatchedSAMLSourceRequest) SetAllowIdpInitiated(v bool)`

SetAllowIdpInitiated sets AllowIdpInitiated field to given value.

### HasAllowIdpInitiated

`func (o *PatchedSAMLSourceRequest) HasAllowIdpInitiated() bool`

HasAllowIdpInitiated returns a boolean if a field has been set.

### GetNameIdPolicy

`func (o *PatchedSAMLSourceRequest) GetNameIdPolicy() NameIdPolicyEnum`

GetNameIdPolicy returns the NameIdPolicy field if non-nil, zero value otherwise.

### GetNameIdPolicyOk

`func (o *PatchedSAMLSourceRequest) GetNameIdPolicyOk() (*NameIdPolicyEnum, bool)`

GetNameIdPolicyOk returns a tuple with the NameIdPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdPolicy

`func (o *PatchedSAMLSourceRequest) SetNameIdPolicy(v NameIdPolicyEnum)`

SetNameIdPolicy sets NameIdPolicy field to given value.

### HasNameIdPolicy

`func (o *PatchedSAMLSourceRequest) HasNameIdPolicy() bool`

HasNameIdPolicy returns a boolean if a field has been set.

### GetBindingType

`func (o *PatchedSAMLSourceRequest) GetBindingType() BindingTypeEnum`

GetBindingType returns the BindingType field if non-nil, zero value otherwise.

### GetBindingTypeOk

`func (o *PatchedSAMLSourceRequest) GetBindingTypeOk() (*BindingTypeEnum, bool)`

GetBindingTypeOk returns a tuple with the BindingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingType

`func (o *PatchedSAMLSourceRequest) SetBindingType(v BindingTypeEnum)`

SetBindingType sets BindingType field to given value.

### HasBindingType

`func (o *PatchedSAMLSourceRequest) HasBindingType() bool`

HasBindingType returns a boolean if a field has been set.

### GetSigningKp

`func (o *PatchedSAMLSourceRequest) GetSigningKp() string`

GetSigningKp returns the SigningKp field if non-nil, zero value otherwise.

### GetSigningKpOk

`func (o *PatchedSAMLSourceRequest) GetSigningKpOk() (*string, bool)`

GetSigningKpOk returns a tuple with the SigningKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKp

`func (o *PatchedSAMLSourceRequest) SetSigningKp(v string)`

SetSigningKp sets SigningKp field to given value.

### HasSigningKp

`func (o *PatchedSAMLSourceRequest) HasSigningKp() bool`

HasSigningKp returns a boolean if a field has been set.

### SetSigningKpNil

`func (o *PatchedSAMLSourceRequest) SetSigningKpNil(b bool)`

 SetSigningKpNil sets the value for SigningKp to be an explicit nil

### UnsetSigningKp
`func (o *PatchedSAMLSourceRequest) UnsetSigningKp()`

UnsetSigningKp ensures that no value is present for SigningKp, not even an explicit nil
### GetDigestAlgorithm

`func (o *PatchedSAMLSourceRequest) GetDigestAlgorithm() DigestAlgorithmEnum`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *PatchedSAMLSourceRequest) GetDigestAlgorithmOk() (*DigestAlgorithmEnum, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *PatchedSAMLSourceRequest) SetDigestAlgorithm(v DigestAlgorithmEnum)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *PatchedSAMLSourceRequest) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *PatchedSAMLSourceRequest) GetSignatureAlgorithm() SignatureAlgorithmEnum`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *PatchedSAMLSourceRequest) GetSignatureAlgorithmOk() (*SignatureAlgorithmEnum, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *PatchedSAMLSourceRequest) SetSignatureAlgorithm(v SignatureAlgorithmEnum)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *PatchedSAMLSourceRequest) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetTemporaryUserDeleteAfter

`func (o *PatchedSAMLSourceRequest) GetTemporaryUserDeleteAfter() string`

GetTemporaryUserDeleteAfter returns the TemporaryUserDeleteAfter field if non-nil, zero value otherwise.

### GetTemporaryUserDeleteAfterOk

`func (o *PatchedSAMLSourceRequest) GetTemporaryUserDeleteAfterOk() (*string, bool)`

GetTemporaryUserDeleteAfterOk returns a tuple with the TemporaryUserDeleteAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryUserDeleteAfter

`func (o *PatchedSAMLSourceRequest) SetTemporaryUserDeleteAfter(v string)`

SetTemporaryUserDeleteAfter sets TemporaryUserDeleteAfter field to given value.

### HasTemporaryUserDeleteAfter

`func (o *PatchedSAMLSourceRequest) HasTemporaryUserDeleteAfter() bool`

HasTemporaryUserDeleteAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


