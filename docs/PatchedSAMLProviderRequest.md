# PatchedSAMLProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** | Flow used for authentication when the associated application is accessed by an un-authenticated user. | [optional] 
**AuthorizationFlow** | Pointer to **string** | Flow used when authorizing this provider. | [optional] 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**AcsUrl** | Pointer to **string** |  | [optional] 
**Audience** | Pointer to **string** | Value of the audience restriction field of the assertion. When left empty, no audience restriction will be added. | [optional] 
**Issuer** | Pointer to **string** | Also known as EntityID | [optional] 
**AssertionValidNotBefore** | Pointer to **string** | Assertion valid not before current time + this value (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3). | [optional] 
**AssertionValidNotOnOrAfter** | Pointer to **string** | Assertion not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**SessionValidNotOnOrAfter** | Pointer to **string** | Session not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**NameIdMapping** | Pointer to **NullableString** | Configure how the NameID value will be created. When left empty, the NameIDPolicy of the incoming request will be considered | [optional] 
**DigestAlgorithm** | Pointer to [**DigestAlgorithmEnum**](DigestAlgorithmEnum.md) |  | [optional] 
**SignatureAlgorithm** | Pointer to [**SignatureAlgorithmEnum**](SignatureAlgorithmEnum.md) |  | [optional] 
**SigningKp** | Pointer to **NullableString** | Keypair used to sign outgoing Responses going to the Service Provider. | [optional] 
**VerificationKp** | Pointer to **NullableString** | When selected, incoming assertion&#39;s Signatures will be validated against this certificate. To allow unsigned Requests, leave on default. | [optional] 
**SpBinding** | Pointer to [**SpBindingEnum**](SpBindingEnum.md) | This determines how authentik sends the response back to the Service Provider.  * &#x60;redirect&#x60; - Redirect * &#x60;post&#x60; - Post | [optional] 

## Methods

### NewPatchedSAMLProviderRequest

`func NewPatchedSAMLProviderRequest() *PatchedSAMLProviderRequest`

NewPatchedSAMLProviderRequest instantiates a new PatchedSAMLProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSAMLProviderRequestWithDefaults

`func NewPatchedSAMLProviderRequestWithDefaults() *PatchedSAMLProviderRequest`

NewPatchedSAMLProviderRequestWithDefaults instantiates a new PatchedSAMLProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedSAMLProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedSAMLProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedSAMLProviderRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedSAMLProviderRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *PatchedSAMLProviderRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *PatchedSAMLProviderRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *PatchedSAMLProviderRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *PatchedSAMLProviderRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *PatchedSAMLProviderRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *PatchedSAMLProviderRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetAuthorizationFlow

`func (o *PatchedSAMLProviderRequest) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *PatchedSAMLProviderRequest) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *PatchedSAMLProviderRequest) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.

### HasAuthorizationFlow

`func (o *PatchedSAMLProviderRequest) HasAuthorizationFlow() bool`

HasAuthorizationFlow returns a boolean if a field has been set.

### GetPropertyMappings

`func (o *PatchedSAMLProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *PatchedSAMLProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *PatchedSAMLProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *PatchedSAMLProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetAcsUrl

`func (o *PatchedSAMLProviderRequest) GetAcsUrl() string`

GetAcsUrl returns the AcsUrl field if non-nil, zero value otherwise.

### GetAcsUrlOk

`func (o *PatchedSAMLProviderRequest) GetAcsUrlOk() (*string, bool)`

GetAcsUrlOk returns a tuple with the AcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsUrl

`func (o *PatchedSAMLProviderRequest) SetAcsUrl(v string)`

SetAcsUrl sets AcsUrl field to given value.

### HasAcsUrl

`func (o *PatchedSAMLProviderRequest) HasAcsUrl() bool`

HasAcsUrl returns a boolean if a field has been set.

### GetAudience

`func (o *PatchedSAMLProviderRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *PatchedSAMLProviderRequest) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *PatchedSAMLProviderRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *PatchedSAMLProviderRequest) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetIssuer

`func (o *PatchedSAMLProviderRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *PatchedSAMLProviderRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *PatchedSAMLProviderRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *PatchedSAMLProviderRequest) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetAssertionValidNotBefore

`func (o *PatchedSAMLProviderRequest) GetAssertionValidNotBefore() string`

GetAssertionValidNotBefore returns the AssertionValidNotBefore field if non-nil, zero value otherwise.

### GetAssertionValidNotBeforeOk

`func (o *PatchedSAMLProviderRequest) GetAssertionValidNotBeforeOk() (*string, bool)`

GetAssertionValidNotBeforeOk returns a tuple with the AssertionValidNotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionValidNotBefore

`func (o *PatchedSAMLProviderRequest) SetAssertionValidNotBefore(v string)`

SetAssertionValidNotBefore sets AssertionValidNotBefore field to given value.

### HasAssertionValidNotBefore

`func (o *PatchedSAMLProviderRequest) HasAssertionValidNotBefore() bool`

HasAssertionValidNotBefore returns a boolean if a field has been set.

### GetAssertionValidNotOnOrAfter

`func (o *PatchedSAMLProviderRequest) GetAssertionValidNotOnOrAfter() string`

GetAssertionValidNotOnOrAfter returns the AssertionValidNotOnOrAfter field if non-nil, zero value otherwise.

### GetAssertionValidNotOnOrAfterOk

`func (o *PatchedSAMLProviderRequest) GetAssertionValidNotOnOrAfterOk() (*string, bool)`

GetAssertionValidNotOnOrAfterOk returns a tuple with the AssertionValidNotOnOrAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionValidNotOnOrAfter

`func (o *PatchedSAMLProviderRequest) SetAssertionValidNotOnOrAfter(v string)`

SetAssertionValidNotOnOrAfter sets AssertionValidNotOnOrAfter field to given value.

### HasAssertionValidNotOnOrAfter

`func (o *PatchedSAMLProviderRequest) HasAssertionValidNotOnOrAfter() bool`

HasAssertionValidNotOnOrAfter returns a boolean if a field has been set.

### GetSessionValidNotOnOrAfter

`func (o *PatchedSAMLProviderRequest) GetSessionValidNotOnOrAfter() string`

GetSessionValidNotOnOrAfter returns the SessionValidNotOnOrAfter field if non-nil, zero value otherwise.

### GetSessionValidNotOnOrAfterOk

`func (o *PatchedSAMLProviderRequest) GetSessionValidNotOnOrAfterOk() (*string, bool)`

GetSessionValidNotOnOrAfterOk returns a tuple with the SessionValidNotOnOrAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionValidNotOnOrAfter

`func (o *PatchedSAMLProviderRequest) SetSessionValidNotOnOrAfter(v string)`

SetSessionValidNotOnOrAfter sets SessionValidNotOnOrAfter field to given value.

### HasSessionValidNotOnOrAfter

`func (o *PatchedSAMLProviderRequest) HasSessionValidNotOnOrAfter() bool`

HasSessionValidNotOnOrAfter returns a boolean if a field has been set.

### GetNameIdMapping

`func (o *PatchedSAMLProviderRequest) GetNameIdMapping() string`

GetNameIdMapping returns the NameIdMapping field if non-nil, zero value otherwise.

### GetNameIdMappingOk

`func (o *PatchedSAMLProviderRequest) GetNameIdMappingOk() (*string, bool)`

GetNameIdMappingOk returns a tuple with the NameIdMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdMapping

`func (o *PatchedSAMLProviderRequest) SetNameIdMapping(v string)`

SetNameIdMapping sets NameIdMapping field to given value.

### HasNameIdMapping

`func (o *PatchedSAMLProviderRequest) HasNameIdMapping() bool`

HasNameIdMapping returns a boolean if a field has been set.

### SetNameIdMappingNil

`func (o *PatchedSAMLProviderRequest) SetNameIdMappingNil(b bool)`

 SetNameIdMappingNil sets the value for NameIdMapping to be an explicit nil

### UnsetNameIdMapping
`func (o *PatchedSAMLProviderRequest) UnsetNameIdMapping()`

UnsetNameIdMapping ensures that no value is present for NameIdMapping, not even an explicit nil
### GetDigestAlgorithm

`func (o *PatchedSAMLProviderRequest) GetDigestAlgorithm() DigestAlgorithmEnum`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *PatchedSAMLProviderRequest) GetDigestAlgorithmOk() (*DigestAlgorithmEnum, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *PatchedSAMLProviderRequest) SetDigestAlgorithm(v DigestAlgorithmEnum)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *PatchedSAMLProviderRequest) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *PatchedSAMLProviderRequest) GetSignatureAlgorithm() SignatureAlgorithmEnum`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *PatchedSAMLProviderRequest) GetSignatureAlgorithmOk() (*SignatureAlgorithmEnum, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *PatchedSAMLProviderRequest) SetSignatureAlgorithm(v SignatureAlgorithmEnum)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *PatchedSAMLProviderRequest) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetSigningKp

`func (o *PatchedSAMLProviderRequest) GetSigningKp() string`

GetSigningKp returns the SigningKp field if non-nil, zero value otherwise.

### GetSigningKpOk

`func (o *PatchedSAMLProviderRequest) GetSigningKpOk() (*string, bool)`

GetSigningKpOk returns a tuple with the SigningKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKp

`func (o *PatchedSAMLProviderRequest) SetSigningKp(v string)`

SetSigningKp sets SigningKp field to given value.

### HasSigningKp

`func (o *PatchedSAMLProviderRequest) HasSigningKp() bool`

HasSigningKp returns a boolean if a field has been set.

### SetSigningKpNil

`func (o *PatchedSAMLProviderRequest) SetSigningKpNil(b bool)`

 SetSigningKpNil sets the value for SigningKp to be an explicit nil

### UnsetSigningKp
`func (o *PatchedSAMLProviderRequest) UnsetSigningKp()`

UnsetSigningKp ensures that no value is present for SigningKp, not even an explicit nil
### GetVerificationKp

`func (o *PatchedSAMLProviderRequest) GetVerificationKp() string`

GetVerificationKp returns the VerificationKp field if non-nil, zero value otherwise.

### GetVerificationKpOk

`func (o *PatchedSAMLProviderRequest) GetVerificationKpOk() (*string, bool)`

GetVerificationKpOk returns a tuple with the VerificationKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationKp

`func (o *PatchedSAMLProviderRequest) SetVerificationKp(v string)`

SetVerificationKp sets VerificationKp field to given value.

### HasVerificationKp

`func (o *PatchedSAMLProviderRequest) HasVerificationKp() bool`

HasVerificationKp returns a boolean if a field has been set.

### SetVerificationKpNil

`func (o *PatchedSAMLProviderRequest) SetVerificationKpNil(b bool)`

 SetVerificationKpNil sets the value for VerificationKp to be an explicit nil

### UnsetVerificationKp
`func (o *PatchedSAMLProviderRequest) UnsetVerificationKp()`

UnsetVerificationKp ensures that no value is present for VerificationKp, not even an explicit nil
### GetSpBinding

`func (o *PatchedSAMLProviderRequest) GetSpBinding() SpBindingEnum`

GetSpBinding returns the SpBinding field if non-nil, zero value otherwise.

### GetSpBindingOk

`func (o *PatchedSAMLProviderRequest) GetSpBindingOk() (*SpBindingEnum, bool)`

GetSpBindingOk returns a tuple with the SpBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpBinding

`func (o *PatchedSAMLProviderRequest) SetSpBinding(v SpBindingEnum)`

SetSpBinding sets SpBinding field to given value.

### HasSpBinding

`func (o *PatchedSAMLProviderRequest) HasSpBinding() bool`

HasSpBinding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


