# SAMLSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Source&#39;s display Name. | 
**Slug** | **string** | Internal source name, used in URLs. | 
**Enabled** | Pointer to **bool** |  | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** | Flow to use when authenticating existing users. | [optional] 
**EnrollmentFlow** | Pointer to **NullableString** | Flow to use when enrolling new users. | [optional] 
**UserPropertyMappings** | Pointer to **[]string** |  | [optional] 
**GroupPropertyMappings** | Pointer to **[]string** |  | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**UserMatchingMode** | Pointer to [**UserMatchingModeEnum**](UserMatchingModeEnum.md) | How the source determines if an existing user should be authenticated or a new user enrolled. | [optional] 
**UserPathTemplate** | Pointer to **string** |  | [optional] 
**PreAuthenticationFlow** | **string** | Flow used before authentication. | 
**Issuer** | Pointer to **string** | Also known as Entity ID. Defaults the Metadata URL. | [optional] 
**SsoUrl** | **string** | URL that the initial Login request is sent to. | 
**SloUrl** | Pointer to **NullableString** | Optional URL if your IDP supports Single-Logout. | [optional] 
**AllowIdpInitiated** | Pointer to **bool** | Allows authentication flows initiated by the IdP. This can be a security risk, as no validation of the request ID is done. | [optional] 
**NameIdPolicy** | Pointer to [**NameIdPolicyEnum**](NameIdPolicyEnum.md) | NameID Policy sent to the IdP. Can be unset, in which case no Policy is sent. | [optional] 
**BindingType** | Pointer to [**BindingTypeEnum**](BindingTypeEnum.md) |  | [optional] 
**VerificationKp** | Pointer to **NullableString** | When selected, incoming assertion&#39;s Signatures will be validated against this certificate. To allow unsigned Requests, leave on default. | [optional] 
**SigningKp** | Pointer to **NullableString** | Keypair used to sign outgoing Responses going to the Identity Provider. | [optional] 
**DigestAlgorithm** | Pointer to [**DigestAlgorithmEnum**](DigestAlgorithmEnum.md) |  | [optional] 
**SignatureAlgorithm** | Pointer to [**SignatureAlgorithmEnum**](SignatureAlgorithmEnum.md) |  | [optional] 
**TemporaryUserDeleteAfter** | Pointer to **string** | Time offset when temporary users should be deleted. This only applies if your IDP uses the NameID Format &#39;transient&#39;, and the user doesn&#39;t log out manually. (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 

## Methods

### NewSAMLSourceRequest

`func NewSAMLSourceRequest(name string, slug string, preAuthenticationFlow string, ssoUrl string, ) *SAMLSourceRequest`

NewSAMLSourceRequest instantiates a new SAMLSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLSourceRequestWithDefaults

`func NewSAMLSourceRequestWithDefaults() *SAMLSourceRequest`

NewSAMLSourceRequestWithDefaults instantiates a new SAMLSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SAMLSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SAMLSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SAMLSourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *SAMLSourceRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *SAMLSourceRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *SAMLSourceRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetEnabled

`func (o *SAMLSourceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SAMLSourceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SAMLSourceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SAMLSourceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *SAMLSourceRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *SAMLSourceRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *SAMLSourceRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *SAMLSourceRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *SAMLSourceRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *SAMLSourceRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetEnrollmentFlow

`func (o *SAMLSourceRequest) GetEnrollmentFlow() string`

GetEnrollmentFlow returns the EnrollmentFlow field if non-nil, zero value otherwise.

### GetEnrollmentFlowOk

`func (o *SAMLSourceRequest) GetEnrollmentFlowOk() (*string, bool)`

GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlow

`func (o *SAMLSourceRequest) SetEnrollmentFlow(v string)`

SetEnrollmentFlow sets EnrollmentFlow field to given value.

### HasEnrollmentFlow

`func (o *SAMLSourceRequest) HasEnrollmentFlow() bool`

HasEnrollmentFlow returns a boolean if a field has been set.

### SetEnrollmentFlowNil

`func (o *SAMLSourceRequest) SetEnrollmentFlowNil(b bool)`

 SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil

### UnsetEnrollmentFlow
`func (o *SAMLSourceRequest) UnsetEnrollmentFlow()`

UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
### GetUserPropertyMappings

`func (o *SAMLSourceRequest) GetUserPropertyMappings() []string`

GetUserPropertyMappings returns the UserPropertyMappings field if non-nil, zero value otherwise.

### GetUserPropertyMappingsOk

`func (o *SAMLSourceRequest) GetUserPropertyMappingsOk() (*[]string, bool)`

GetUserPropertyMappingsOk returns a tuple with the UserPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPropertyMappings

`func (o *SAMLSourceRequest) SetUserPropertyMappings(v []string)`

SetUserPropertyMappings sets UserPropertyMappings field to given value.

### HasUserPropertyMappings

`func (o *SAMLSourceRequest) HasUserPropertyMappings() bool`

HasUserPropertyMappings returns a boolean if a field has been set.

### GetGroupPropertyMappings

`func (o *SAMLSourceRequest) GetGroupPropertyMappings() []string`

GetGroupPropertyMappings returns the GroupPropertyMappings field if non-nil, zero value otherwise.

### GetGroupPropertyMappingsOk

`func (o *SAMLSourceRequest) GetGroupPropertyMappingsOk() (*[]string, bool)`

GetGroupPropertyMappingsOk returns a tuple with the GroupPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPropertyMappings

`func (o *SAMLSourceRequest) SetGroupPropertyMappings(v []string)`

SetGroupPropertyMappings sets GroupPropertyMappings field to given value.

### HasGroupPropertyMappings

`func (o *SAMLSourceRequest) HasGroupPropertyMappings() bool`

HasGroupPropertyMappings returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *SAMLSourceRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *SAMLSourceRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *SAMLSourceRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *SAMLSourceRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetUserMatchingMode

`func (o *SAMLSourceRequest) GetUserMatchingMode() UserMatchingModeEnum`

GetUserMatchingMode returns the UserMatchingMode field if non-nil, zero value otherwise.

### GetUserMatchingModeOk

`func (o *SAMLSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool)`

GetUserMatchingModeOk returns a tuple with the UserMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMatchingMode

`func (o *SAMLSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum)`

SetUserMatchingMode sets UserMatchingMode field to given value.

### HasUserMatchingMode

`func (o *SAMLSourceRequest) HasUserMatchingMode() bool`

HasUserMatchingMode returns a boolean if a field has been set.

### GetUserPathTemplate

`func (o *SAMLSourceRequest) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *SAMLSourceRequest) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *SAMLSourceRequest) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *SAMLSourceRequest) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.

### GetPreAuthenticationFlow

`func (o *SAMLSourceRequest) GetPreAuthenticationFlow() string`

GetPreAuthenticationFlow returns the PreAuthenticationFlow field if non-nil, zero value otherwise.

### GetPreAuthenticationFlowOk

`func (o *SAMLSourceRequest) GetPreAuthenticationFlowOk() (*string, bool)`

GetPreAuthenticationFlowOk returns a tuple with the PreAuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthenticationFlow

`func (o *SAMLSourceRequest) SetPreAuthenticationFlow(v string)`

SetPreAuthenticationFlow sets PreAuthenticationFlow field to given value.


### GetIssuer

`func (o *SAMLSourceRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SAMLSourceRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SAMLSourceRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *SAMLSourceRequest) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSsoUrl

`func (o *SAMLSourceRequest) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *SAMLSourceRequest) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *SAMLSourceRequest) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.


### GetSloUrl

`func (o *SAMLSourceRequest) GetSloUrl() string`

GetSloUrl returns the SloUrl field if non-nil, zero value otherwise.

### GetSloUrlOk

`func (o *SAMLSourceRequest) GetSloUrlOk() (*string, bool)`

GetSloUrlOk returns a tuple with the SloUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloUrl

`func (o *SAMLSourceRequest) SetSloUrl(v string)`

SetSloUrl sets SloUrl field to given value.

### HasSloUrl

`func (o *SAMLSourceRequest) HasSloUrl() bool`

HasSloUrl returns a boolean if a field has been set.

### SetSloUrlNil

`func (o *SAMLSourceRequest) SetSloUrlNil(b bool)`

 SetSloUrlNil sets the value for SloUrl to be an explicit nil

### UnsetSloUrl
`func (o *SAMLSourceRequest) UnsetSloUrl()`

UnsetSloUrl ensures that no value is present for SloUrl, not even an explicit nil
### GetAllowIdpInitiated

`func (o *SAMLSourceRequest) GetAllowIdpInitiated() bool`

GetAllowIdpInitiated returns the AllowIdpInitiated field if non-nil, zero value otherwise.

### GetAllowIdpInitiatedOk

`func (o *SAMLSourceRequest) GetAllowIdpInitiatedOk() (*bool, bool)`

GetAllowIdpInitiatedOk returns a tuple with the AllowIdpInitiated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIdpInitiated

`func (o *SAMLSourceRequest) SetAllowIdpInitiated(v bool)`

SetAllowIdpInitiated sets AllowIdpInitiated field to given value.

### HasAllowIdpInitiated

`func (o *SAMLSourceRequest) HasAllowIdpInitiated() bool`

HasAllowIdpInitiated returns a boolean if a field has been set.

### GetNameIdPolicy

`func (o *SAMLSourceRequest) GetNameIdPolicy() NameIdPolicyEnum`

GetNameIdPolicy returns the NameIdPolicy field if non-nil, zero value otherwise.

### GetNameIdPolicyOk

`func (o *SAMLSourceRequest) GetNameIdPolicyOk() (*NameIdPolicyEnum, bool)`

GetNameIdPolicyOk returns a tuple with the NameIdPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdPolicy

`func (o *SAMLSourceRequest) SetNameIdPolicy(v NameIdPolicyEnum)`

SetNameIdPolicy sets NameIdPolicy field to given value.

### HasNameIdPolicy

`func (o *SAMLSourceRequest) HasNameIdPolicy() bool`

HasNameIdPolicy returns a boolean if a field has been set.

### GetBindingType

`func (o *SAMLSourceRequest) GetBindingType() BindingTypeEnum`

GetBindingType returns the BindingType field if non-nil, zero value otherwise.

### GetBindingTypeOk

`func (o *SAMLSourceRequest) GetBindingTypeOk() (*BindingTypeEnum, bool)`

GetBindingTypeOk returns a tuple with the BindingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingType

`func (o *SAMLSourceRequest) SetBindingType(v BindingTypeEnum)`

SetBindingType sets BindingType field to given value.

### HasBindingType

`func (o *SAMLSourceRequest) HasBindingType() bool`

HasBindingType returns a boolean if a field has been set.

### GetVerificationKp

`func (o *SAMLSourceRequest) GetVerificationKp() string`

GetVerificationKp returns the VerificationKp field if non-nil, zero value otherwise.

### GetVerificationKpOk

`func (o *SAMLSourceRequest) GetVerificationKpOk() (*string, bool)`

GetVerificationKpOk returns a tuple with the VerificationKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationKp

`func (o *SAMLSourceRequest) SetVerificationKp(v string)`

SetVerificationKp sets VerificationKp field to given value.

### HasVerificationKp

`func (o *SAMLSourceRequest) HasVerificationKp() bool`

HasVerificationKp returns a boolean if a field has been set.

### SetVerificationKpNil

`func (o *SAMLSourceRequest) SetVerificationKpNil(b bool)`

 SetVerificationKpNil sets the value for VerificationKp to be an explicit nil

### UnsetVerificationKp
`func (o *SAMLSourceRequest) UnsetVerificationKp()`

UnsetVerificationKp ensures that no value is present for VerificationKp, not even an explicit nil
### GetSigningKp

`func (o *SAMLSourceRequest) GetSigningKp() string`

GetSigningKp returns the SigningKp field if non-nil, zero value otherwise.

### GetSigningKpOk

`func (o *SAMLSourceRequest) GetSigningKpOk() (*string, bool)`

GetSigningKpOk returns a tuple with the SigningKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKp

`func (o *SAMLSourceRequest) SetSigningKp(v string)`

SetSigningKp sets SigningKp field to given value.

### HasSigningKp

`func (o *SAMLSourceRequest) HasSigningKp() bool`

HasSigningKp returns a boolean if a field has been set.

### SetSigningKpNil

`func (o *SAMLSourceRequest) SetSigningKpNil(b bool)`

 SetSigningKpNil sets the value for SigningKp to be an explicit nil

### UnsetSigningKp
`func (o *SAMLSourceRequest) UnsetSigningKp()`

UnsetSigningKp ensures that no value is present for SigningKp, not even an explicit nil
### GetDigestAlgorithm

`func (o *SAMLSourceRequest) GetDigestAlgorithm() DigestAlgorithmEnum`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *SAMLSourceRequest) GetDigestAlgorithmOk() (*DigestAlgorithmEnum, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *SAMLSourceRequest) SetDigestAlgorithm(v DigestAlgorithmEnum)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *SAMLSourceRequest) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *SAMLSourceRequest) GetSignatureAlgorithm() SignatureAlgorithmEnum`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *SAMLSourceRequest) GetSignatureAlgorithmOk() (*SignatureAlgorithmEnum, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *SAMLSourceRequest) SetSignatureAlgorithm(v SignatureAlgorithmEnum)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *SAMLSourceRequest) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetTemporaryUserDeleteAfter

`func (o *SAMLSourceRequest) GetTemporaryUserDeleteAfter() string`

GetTemporaryUserDeleteAfter returns the TemporaryUserDeleteAfter field if non-nil, zero value otherwise.

### GetTemporaryUserDeleteAfterOk

`func (o *SAMLSourceRequest) GetTemporaryUserDeleteAfterOk() (*string, bool)`

GetTemporaryUserDeleteAfterOk returns a tuple with the TemporaryUserDeleteAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryUserDeleteAfter

`func (o *SAMLSourceRequest) SetTemporaryUserDeleteAfter(v string)`

SetTemporaryUserDeleteAfter sets TemporaryUserDeleteAfter field to given value.

### HasTemporaryUserDeleteAfter

`func (o *SAMLSourceRequest) HasTemporaryUserDeleteAfter() bool`

HasTemporaryUserDeleteAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


