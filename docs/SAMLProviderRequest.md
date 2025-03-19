# SAMLProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AuthenticationFlow** | Pointer to **NullableString** | Flow used for authentication when the associated application is accessed by an un-authenticated user. | [optional] 
**AuthorizationFlow** | **string** | Flow used when authorizing this provider. | 
**InvalidationFlow** | **string** | Flow used ending the session from a provider. | 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**AcsUrl** | **string** |  | 
**Audience** | Pointer to **string** | Value of the audience restriction field of the assertion. When left empty, no audience restriction will be added. | [optional] 
**Issuer** | Pointer to **string** | Also known as EntityID | [optional] 
**AssertionValidNotBefore** | Pointer to **string** | Assertion valid not before current time + this value (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3). | [optional] 
**AssertionValidNotOnOrAfter** | Pointer to **string** | Assertion not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**SessionValidNotOnOrAfter** | Pointer to **string** | Session not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**NameIdMapping** | Pointer to **NullableString** | Configure how the NameID value will be created. When left empty, the NameIDPolicy of the incoming request will be considered | [optional] 
**AuthnContextClassRefMapping** | Pointer to **NullableString** | Configure how the AuthnContextClassRef value will be created. When left empty, the AuthnContextClassRef will be set based on which authentication methods the user used to authenticate. | [optional] 
**DigestAlgorithm** | Pointer to [**DigestAlgorithmEnum**](DigestAlgorithmEnum.md) |  | [optional] 
**SignatureAlgorithm** | Pointer to [**SignatureAlgorithmEnum**](SignatureAlgorithmEnum.md) |  | [optional] 
**SigningKp** | Pointer to **NullableString** | Keypair used to sign outgoing Responses going to the Service Provider. | [optional] 
**VerificationKp** | Pointer to **NullableString** | When selected, incoming assertion&#39;s Signatures will be validated against this certificate. To allow unsigned Requests, leave on default. | [optional] 
**EncryptionKp** | Pointer to **NullableString** | When selected, incoming assertions are encrypted by the IdP using the public key of the encryption keypair. The assertion is decrypted by the SP using the the private key. | [optional] 
**SignAssertion** | Pointer to **bool** |  | [optional] 
**SignResponse** | Pointer to **bool** |  | [optional] 
**SpBinding** | Pointer to [**SpBindingEnum**](SpBindingEnum.md) | This determines how authentik sends the response back to the Service Provider. | [optional] 
**DefaultRelayState** | Pointer to **string** | Default relay_state value for IDP-initiated logins | [optional] 

## Methods

### NewSAMLProviderRequest

`func NewSAMLProviderRequest(name string, authorizationFlow string, invalidationFlow string, acsUrl string, ) *SAMLProviderRequest`

NewSAMLProviderRequest instantiates a new SAMLProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLProviderRequestWithDefaults

`func NewSAMLProviderRequestWithDefaults() *SAMLProviderRequest`

NewSAMLProviderRequestWithDefaults instantiates a new SAMLProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SAMLProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SAMLProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SAMLProviderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAuthenticationFlow

`func (o *SAMLProviderRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *SAMLProviderRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *SAMLProviderRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *SAMLProviderRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *SAMLProviderRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *SAMLProviderRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetAuthorizationFlow

`func (o *SAMLProviderRequest) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *SAMLProviderRequest) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *SAMLProviderRequest) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.


### GetInvalidationFlow

`func (o *SAMLProviderRequest) GetInvalidationFlow() string`

GetInvalidationFlow returns the InvalidationFlow field if non-nil, zero value otherwise.

### GetInvalidationFlowOk

`func (o *SAMLProviderRequest) GetInvalidationFlowOk() (*string, bool)`

GetInvalidationFlowOk returns a tuple with the InvalidationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidationFlow

`func (o *SAMLProviderRequest) SetInvalidationFlow(v string)`

SetInvalidationFlow sets InvalidationFlow field to given value.


### GetPropertyMappings

`func (o *SAMLProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *SAMLProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *SAMLProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *SAMLProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetAcsUrl

`func (o *SAMLProviderRequest) GetAcsUrl() string`

GetAcsUrl returns the AcsUrl field if non-nil, zero value otherwise.

### GetAcsUrlOk

`func (o *SAMLProviderRequest) GetAcsUrlOk() (*string, bool)`

GetAcsUrlOk returns a tuple with the AcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsUrl

`func (o *SAMLProviderRequest) SetAcsUrl(v string)`

SetAcsUrl sets AcsUrl field to given value.


### GetAudience

`func (o *SAMLProviderRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *SAMLProviderRequest) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *SAMLProviderRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *SAMLProviderRequest) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetIssuer

`func (o *SAMLProviderRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SAMLProviderRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SAMLProviderRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *SAMLProviderRequest) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetAssertionValidNotBefore

`func (o *SAMLProviderRequest) GetAssertionValidNotBefore() string`

GetAssertionValidNotBefore returns the AssertionValidNotBefore field if non-nil, zero value otherwise.

### GetAssertionValidNotBeforeOk

`func (o *SAMLProviderRequest) GetAssertionValidNotBeforeOk() (*string, bool)`

GetAssertionValidNotBeforeOk returns a tuple with the AssertionValidNotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionValidNotBefore

`func (o *SAMLProviderRequest) SetAssertionValidNotBefore(v string)`

SetAssertionValidNotBefore sets AssertionValidNotBefore field to given value.

### HasAssertionValidNotBefore

`func (o *SAMLProviderRequest) HasAssertionValidNotBefore() bool`

HasAssertionValidNotBefore returns a boolean if a field has been set.

### GetAssertionValidNotOnOrAfter

`func (o *SAMLProviderRequest) GetAssertionValidNotOnOrAfter() string`

GetAssertionValidNotOnOrAfter returns the AssertionValidNotOnOrAfter field if non-nil, zero value otherwise.

### GetAssertionValidNotOnOrAfterOk

`func (o *SAMLProviderRequest) GetAssertionValidNotOnOrAfterOk() (*string, bool)`

GetAssertionValidNotOnOrAfterOk returns a tuple with the AssertionValidNotOnOrAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionValidNotOnOrAfter

`func (o *SAMLProviderRequest) SetAssertionValidNotOnOrAfter(v string)`

SetAssertionValidNotOnOrAfter sets AssertionValidNotOnOrAfter field to given value.

### HasAssertionValidNotOnOrAfter

`func (o *SAMLProviderRequest) HasAssertionValidNotOnOrAfter() bool`

HasAssertionValidNotOnOrAfter returns a boolean if a field has been set.

### GetSessionValidNotOnOrAfter

`func (o *SAMLProviderRequest) GetSessionValidNotOnOrAfter() string`

GetSessionValidNotOnOrAfter returns the SessionValidNotOnOrAfter field if non-nil, zero value otherwise.

### GetSessionValidNotOnOrAfterOk

`func (o *SAMLProviderRequest) GetSessionValidNotOnOrAfterOk() (*string, bool)`

GetSessionValidNotOnOrAfterOk returns a tuple with the SessionValidNotOnOrAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionValidNotOnOrAfter

`func (o *SAMLProviderRequest) SetSessionValidNotOnOrAfter(v string)`

SetSessionValidNotOnOrAfter sets SessionValidNotOnOrAfter field to given value.

### HasSessionValidNotOnOrAfter

`func (o *SAMLProviderRequest) HasSessionValidNotOnOrAfter() bool`

HasSessionValidNotOnOrAfter returns a boolean if a field has been set.

### GetNameIdMapping

`func (o *SAMLProviderRequest) GetNameIdMapping() string`

GetNameIdMapping returns the NameIdMapping field if non-nil, zero value otherwise.

### GetNameIdMappingOk

`func (o *SAMLProviderRequest) GetNameIdMappingOk() (*string, bool)`

GetNameIdMappingOk returns a tuple with the NameIdMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdMapping

`func (o *SAMLProviderRequest) SetNameIdMapping(v string)`

SetNameIdMapping sets NameIdMapping field to given value.

### HasNameIdMapping

`func (o *SAMLProviderRequest) HasNameIdMapping() bool`

HasNameIdMapping returns a boolean if a field has been set.

### SetNameIdMappingNil

`func (o *SAMLProviderRequest) SetNameIdMappingNil(b bool)`

 SetNameIdMappingNil sets the value for NameIdMapping to be an explicit nil

### UnsetNameIdMapping
`func (o *SAMLProviderRequest) UnsetNameIdMapping()`

UnsetNameIdMapping ensures that no value is present for NameIdMapping, not even an explicit nil
### GetAuthnContextClassRefMapping

`func (o *SAMLProviderRequest) GetAuthnContextClassRefMapping() string`

GetAuthnContextClassRefMapping returns the AuthnContextClassRefMapping field if non-nil, zero value otherwise.

### GetAuthnContextClassRefMappingOk

`func (o *SAMLProviderRequest) GetAuthnContextClassRefMappingOk() (*string, bool)`

GetAuthnContextClassRefMappingOk returns a tuple with the AuthnContextClassRefMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnContextClassRefMapping

`func (o *SAMLProviderRequest) SetAuthnContextClassRefMapping(v string)`

SetAuthnContextClassRefMapping sets AuthnContextClassRefMapping field to given value.

### HasAuthnContextClassRefMapping

`func (o *SAMLProviderRequest) HasAuthnContextClassRefMapping() bool`

HasAuthnContextClassRefMapping returns a boolean if a field has been set.

### SetAuthnContextClassRefMappingNil

`func (o *SAMLProviderRequest) SetAuthnContextClassRefMappingNil(b bool)`

 SetAuthnContextClassRefMappingNil sets the value for AuthnContextClassRefMapping to be an explicit nil

### UnsetAuthnContextClassRefMapping
`func (o *SAMLProviderRequest) UnsetAuthnContextClassRefMapping()`

UnsetAuthnContextClassRefMapping ensures that no value is present for AuthnContextClassRefMapping, not even an explicit nil
### GetDigestAlgorithm

`func (o *SAMLProviderRequest) GetDigestAlgorithm() DigestAlgorithmEnum`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *SAMLProviderRequest) GetDigestAlgorithmOk() (*DigestAlgorithmEnum, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *SAMLProviderRequest) SetDigestAlgorithm(v DigestAlgorithmEnum)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *SAMLProviderRequest) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *SAMLProviderRequest) GetSignatureAlgorithm() SignatureAlgorithmEnum`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *SAMLProviderRequest) GetSignatureAlgorithmOk() (*SignatureAlgorithmEnum, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *SAMLProviderRequest) SetSignatureAlgorithm(v SignatureAlgorithmEnum)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *SAMLProviderRequest) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetSigningKp

`func (o *SAMLProviderRequest) GetSigningKp() string`

GetSigningKp returns the SigningKp field if non-nil, zero value otherwise.

### GetSigningKpOk

`func (o *SAMLProviderRequest) GetSigningKpOk() (*string, bool)`

GetSigningKpOk returns a tuple with the SigningKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKp

`func (o *SAMLProviderRequest) SetSigningKp(v string)`

SetSigningKp sets SigningKp field to given value.

### HasSigningKp

`func (o *SAMLProviderRequest) HasSigningKp() bool`

HasSigningKp returns a boolean if a field has been set.

### SetSigningKpNil

`func (o *SAMLProviderRequest) SetSigningKpNil(b bool)`

 SetSigningKpNil sets the value for SigningKp to be an explicit nil

### UnsetSigningKp
`func (o *SAMLProviderRequest) UnsetSigningKp()`

UnsetSigningKp ensures that no value is present for SigningKp, not even an explicit nil
### GetVerificationKp

`func (o *SAMLProviderRequest) GetVerificationKp() string`

GetVerificationKp returns the VerificationKp field if non-nil, zero value otherwise.

### GetVerificationKpOk

`func (o *SAMLProviderRequest) GetVerificationKpOk() (*string, bool)`

GetVerificationKpOk returns a tuple with the VerificationKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationKp

`func (o *SAMLProviderRequest) SetVerificationKp(v string)`

SetVerificationKp sets VerificationKp field to given value.

### HasVerificationKp

`func (o *SAMLProviderRequest) HasVerificationKp() bool`

HasVerificationKp returns a boolean if a field has been set.

### SetVerificationKpNil

`func (o *SAMLProviderRequest) SetVerificationKpNil(b bool)`

 SetVerificationKpNil sets the value for VerificationKp to be an explicit nil

### UnsetVerificationKp
`func (o *SAMLProviderRequest) UnsetVerificationKp()`

UnsetVerificationKp ensures that no value is present for VerificationKp, not even an explicit nil
### GetEncryptionKp

`func (o *SAMLProviderRequest) GetEncryptionKp() string`

GetEncryptionKp returns the EncryptionKp field if non-nil, zero value otherwise.

### GetEncryptionKpOk

`func (o *SAMLProviderRequest) GetEncryptionKpOk() (*string, bool)`

GetEncryptionKpOk returns a tuple with the EncryptionKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKp

`func (o *SAMLProviderRequest) SetEncryptionKp(v string)`

SetEncryptionKp sets EncryptionKp field to given value.

### HasEncryptionKp

`func (o *SAMLProviderRequest) HasEncryptionKp() bool`

HasEncryptionKp returns a boolean if a field has been set.

### SetEncryptionKpNil

`func (o *SAMLProviderRequest) SetEncryptionKpNil(b bool)`

 SetEncryptionKpNil sets the value for EncryptionKp to be an explicit nil

### UnsetEncryptionKp
`func (o *SAMLProviderRequest) UnsetEncryptionKp()`

UnsetEncryptionKp ensures that no value is present for EncryptionKp, not even an explicit nil
### GetSignAssertion

`func (o *SAMLProviderRequest) GetSignAssertion() bool`

GetSignAssertion returns the SignAssertion field if non-nil, zero value otherwise.

### GetSignAssertionOk

`func (o *SAMLProviderRequest) GetSignAssertionOk() (*bool, bool)`

GetSignAssertionOk returns a tuple with the SignAssertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAssertion

`func (o *SAMLProviderRequest) SetSignAssertion(v bool)`

SetSignAssertion sets SignAssertion field to given value.

### HasSignAssertion

`func (o *SAMLProviderRequest) HasSignAssertion() bool`

HasSignAssertion returns a boolean if a field has been set.

### GetSignResponse

`func (o *SAMLProviderRequest) GetSignResponse() bool`

GetSignResponse returns the SignResponse field if non-nil, zero value otherwise.

### GetSignResponseOk

`func (o *SAMLProviderRequest) GetSignResponseOk() (*bool, bool)`

GetSignResponseOk returns a tuple with the SignResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignResponse

`func (o *SAMLProviderRequest) SetSignResponse(v bool)`

SetSignResponse sets SignResponse field to given value.

### HasSignResponse

`func (o *SAMLProviderRequest) HasSignResponse() bool`

HasSignResponse returns a boolean if a field has been set.

### GetSpBinding

`func (o *SAMLProviderRequest) GetSpBinding() SpBindingEnum`

GetSpBinding returns the SpBinding field if non-nil, zero value otherwise.

### GetSpBindingOk

`func (o *SAMLProviderRequest) GetSpBindingOk() (*SpBindingEnum, bool)`

GetSpBindingOk returns a tuple with the SpBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpBinding

`func (o *SAMLProviderRequest) SetSpBinding(v SpBindingEnum)`

SetSpBinding sets SpBinding field to given value.

### HasSpBinding

`func (o *SAMLProviderRequest) HasSpBinding() bool`

HasSpBinding returns a boolean if a field has been set.

### GetDefaultRelayState

`func (o *SAMLProviderRequest) GetDefaultRelayState() string`

GetDefaultRelayState returns the DefaultRelayState field if non-nil, zero value otherwise.

### GetDefaultRelayStateOk

`func (o *SAMLProviderRequest) GetDefaultRelayStateOk() (*string, bool)`

GetDefaultRelayStateOk returns a tuple with the DefaultRelayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRelayState

`func (o *SAMLProviderRequest) SetDefaultRelayState(v string)`

SetDefaultRelayState sets DefaultRelayState field to given value.

### HasDefaultRelayState

`func (o *SAMLProviderRequest) HasDefaultRelayState() bool`

HasDefaultRelayState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


