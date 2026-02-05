# PatchedWSFederationProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** | Flow used for authentication when the associated application is accessed by an un-authenticated user. | [optional] 
**AuthorizationFlow** | Pointer to **string** | Flow used when authorizing this provider. | [optional] 
**InvalidationFlow** | Pointer to **string** | Flow used ending the session from a provider. | [optional] 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**ReplyUrl** | Pointer to **string** |  | [optional] 
**Wtrealm** | Pointer to **string** |  | [optional] 
**AssertionValidNotBefore** | Pointer to **string** | Assertion valid not before current time + this value (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3). | [optional] 
**AssertionValidNotOnOrAfter** | Pointer to **string** | Assertion not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**SessionValidNotOnOrAfter** | Pointer to **string** | Session not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**NameIdMapping** | Pointer to **NullableString** | Configure how the NameID value will be created. When left empty, the NameIDPolicy of the incoming request will be considered | [optional] 
**AuthnContextClassRefMapping** | Pointer to **NullableString** | Configure how the AuthnContextClassRef value will be created. When left empty, the AuthnContextClassRef will be set based on which authentication methods the user used to authenticate. | [optional] 
**DigestAlgorithm** | Pointer to [**DigestAlgorithmEnum**](DigestAlgorithmEnum.md) |  | [optional] 
**SignatureAlgorithm** | Pointer to [**SignatureAlgorithmEnum**](SignatureAlgorithmEnum.md) |  | [optional] 
**SigningKp** | Pointer to **NullableString** | Keypair used to sign outgoing Responses going to the Service Provider. | [optional] 
**EncryptionKp** | Pointer to **NullableString** | When selected, incoming assertions are encrypted by the IdP using the public key of the encryption keypair. The assertion is decrypted by the SP using the the private key. | [optional] 
**SignAssertion** | Pointer to **bool** |  | [optional] 
**SignLogoutRequest** | Pointer to **bool** |  | [optional] 
**DefaultNameIdPolicy** | Pointer to [**SAMLNameIDPolicyEnum**](SAMLNameIDPolicyEnum.md) |  | [optional] 

## Methods

### NewPatchedWSFederationProviderRequest

`func NewPatchedWSFederationProviderRequest() *PatchedWSFederationProviderRequest`

NewPatchedWSFederationProviderRequest instantiates a new PatchedWSFederationProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWSFederationProviderRequestWithDefaults

`func NewPatchedWSFederationProviderRequestWithDefaults() *PatchedWSFederationProviderRequest`

NewPatchedWSFederationProviderRequestWithDefaults instantiates a new PatchedWSFederationProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWSFederationProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWSFederationProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWSFederationProviderRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWSFederationProviderRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *PatchedWSFederationProviderRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *PatchedWSFederationProviderRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *PatchedWSFederationProviderRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *PatchedWSFederationProviderRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *PatchedWSFederationProviderRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *PatchedWSFederationProviderRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetAuthorizationFlow

`func (o *PatchedWSFederationProviderRequest) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *PatchedWSFederationProviderRequest) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *PatchedWSFederationProviderRequest) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.

### HasAuthorizationFlow

`func (o *PatchedWSFederationProviderRequest) HasAuthorizationFlow() bool`

HasAuthorizationFlow returns a boolean if a field has been set.

### GetInvalidationFlow

`func (o *PatchedWSFederationProviderRequest) GetInvalidationFlow() string`

GetInvalidationFlow returns the InvalidationFlow field if non-nil, zero value otherwise.

### GetInvalidationFlowOk

`func (o *PatchedWSFederationProviderRequest) GetInvalidationFlowOk() (*string, bool)`

GetInvalidationFlowOk returns a tuple with the InvalidationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidationFlow

`func (o *PatchedWSFederationProviderRequest) SetInvalidationFlow(v string)`

SetInvalidationFlow sets InvalidationFlow field to given value.

### HasInvalidationFlow

`func (o *PatchedWSFederationProviderRequest) HasInvalidationFlow() bool`

HasInvalidationFlow returns a boolean if a field has been set.

### GetPropertyMappings

`func (o *PatchedWSFederationProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *PatchedWSFederationProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *PatchedWSFederationProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *PatchedWSFederationProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetReplyUrl

`func (o *PatchedWSFederationProviderRequest) GetReplyUrl() string`

GetReplyUrl returns the ReplyUrl field if non-nil, zero value otherwise.

### GetReplyUrlOk

`func (o *PatchedWSFederationProviderRequest) GetReplyUrlOk() (*string, bool)`

GetReplyUrlOk returns a tuple with the ReplyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyUrl

`func (o *PatchedWSFederationProviderRequest) SetReplyUrl(v string)`

SetReplyUrl sets ReplyUrl field to given value.

### HasReplyUrl

`func (o *PatchedWSFederationProviderRequest) HasReplyUrl() bool`

HasReplyUrl returns a boolean if a field has been set.

### GetWtrealm

`func (o *PatchedWSFederationProviderRequest) GetWtrealm() string`

GetWtrealm returns the Wtrealm field if non-nil, zero value otherwise.

### GetWtrealmOk

`func (o *PatchedWSFederationProviderRequest) GetWtrealmOk() (*string, bool)`

GetWtrealmOk returns a tuple with the Wtrealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWtrealm

`func (o *PatchedWSFederationProviderRequest) SetWtrealm(v string)`

SetWtrealm sets Wtrealm field to given value.

### HasWtrealm

`func (o *PatchedWSFederationProviderRequest) HasWtrealm() bool`

HasWtrealm returns a boolean if a field has been set.

### GetAssertionValidNotBefore

`func (o *PatchedWSFederationProviderRequest) GetAssertionValidNotBefore() string`

GetAssertionValidNotBefore returns the AssertionValidNotBefore field if non-nil, zero value otherwise.

### GetAssertionValidNotBeforeOk

`func (o *PatchedWSFederationProviderRequest) GetAssertionValidNotBeforeOk() (*string, bool)`

GetAssertionValidNotBeforeOk returns a tuple with the AssertionValidNotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionValidNotBefore

`func (o *PatchedWSFederationProviderRequest) SetAssertionValidNotBefore(v string)`

SetAssertionValidNotBefore sets AssertionValidNotBefore field to given value.

### HasAssertionValidNotBefore

`func (o *PatchedWSFederationProviderRequest) HasAssertionValidNotBefore() bool`

HasAssertionValidNotBefore returns a boolean if a field has been set.

### GetAssertionValidNotOnOrAfter

`func (o *PatchedWSFederationProviderRequest) GetAssertionValidNotOnOrAfter() string`

GetAssertionValidNotOnOrAfter returns the AssertionValidNotOnOrAfter field if non-nil, zero value otherwise.

### GetAssertionValidNotOnOrAfterOk

`func (o *PatchedWSFederationProviderRequest) GetAssertionValidNotOnOrAfterOk() (*string, bool)`

GetAssertionValidNotOnOrAfterOk returns a tuple with the AssertionValidNotOnOrAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionValidNotOnOrAfter

`func (o *PatchedWSFederationProviderRequest) SetAssertionValidNotOnOrAfter(v string)`

SetAssertionValidNotOnOrAfter sets AssertionValidNotOnOrAfter field to given value.

### HasAssertionValidNotOnOrAfter

`func (o *PatchedWSFederationProviderRequest) HasAssertionValidNotOnOrAfter() bool`

HasAssertionValidNotOnOrAfter returns a boolean if a field has been set.

### GetSessionValidNotOnOrAfter

`func (o *PatchedWSFederationProviderRequest) GetSessionValidNotOnOrAfter() string`

GetSessionValidNotOnOrAfter returns the SessionValidNotOnOrAfter field if non-nil, zero value otherwise.

### GetSessionValidNotOnOrAfterOk

`func (o *PatchedWSFederationProviderRequest) GetSessionValidNotOnOrAfterOk() (*string, bool)`

GetSessionValidNotOnOrAfterOk returns a tuple with the SessionValidNotOnOrAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionValidNotOnOrAfter

`func (o *PatchedWSFederationProviderRequest) SetSessionValidNotOnOrAfter(v string)`

SetSessionValidNotOnOrAfter sets SessionValidNotOnOrAfter field to given value.

### HasSessionValidNotOnOrAfter

`func (o *PatchedWSFederationProviderRequest) HasSessionValidNotOnOrAfter() bool`

HasSessionValidNotOnOrAfter returns a boolean if a field has been set.

### GetNameIdMapping

`func (o *PatchedWSFederationProviderRequest) GetNameIdMapping() string`

GetNameIdMapping returns the NameIdMapping field if non-nil, zero value otherwise.

### GetNameIdMappingOk

`func (o *PatchedWSFederationProviderRequest) GetNameIdMappingOk() (*string, bool)`

GetNameIdMappingOk returns a tuple with the NameIdMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdMapping

`func (o *PatchedWSFederationProviderRequest) SetNameIdMapping(v string)`

SetNameIdMapping sets NameIdMapping field to given value.

### HasNameIdMapping

`func (o *PatchedWSFederationProviderRequest) HasNameIdMapping() bool`

HasNameIdMapping returns a boolean if a field has been set.

### SetNameIdMappingNil

`func (o *PatchedWSFederationProviderRequest) SetNameIdMappingNil(b bool)`

 SetNameIdMappingNil sets the value for NameIdMapping to be an explicit nil

### UnsetNameIdMapping
`func (o *PatchedWSFederationProviderRequest) UnsetNameIdMapping()`

UnsetNameIdMapping ensures that no value is present for NameIdMapping, not even an explicit nil
### GetAuthnContextClassRefMapping

`func (o *PatchedWSFederationProviderRequest) GetAuthnContextClassRefMapping() string`

GetAuthnContextClassRefMapping returns the AuthnContextClassRefMapping field if non-nil, zero value otherwise.

### GetAuthnContextClassRefMappingOk

`func (o *PatchedWSFederationProviderRequest) GetAuthnContextClassRefMappingOk() (*string, bool)`

GetAuthnContextClassRefMappingOk returns a tuple with the AuthnContextClassRefMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnContextClassRefMapping

`func (o *PatchedWSFederationProviderRequest) SetAuthnContextClassRefMapping(v string)`

SetAuthnContextClassRefMapping sets AuthnContextClassRefMapping field to given value.

### HasAuthnContextClassRefMapping

`func (o *PatchedWSFederationProviderRequest) HasAuthnContextClassRefMapping() bool`

HasAuthnContextClassRefMapping returns a boolean if a field has been set.

### SetAuthnContextClassRefMappingNil

`func (o *PatchedWSFederationProviderRequest) SetAuthnContextClassRefMappingNil(b bool)`

 SetAuthnContextClassRefMappingNil sets the value for AuthnContextClassRefMapping to be an explicit nil

### UnsetAuthnContextClassRefMapping
`func (o *PatchedWSFederationProviderRequest) UnsetAuthnContextClassRefMapping()`

UnsetAuthnContextClassRefMapping ensures that no value is present for AuthnContextClassRefMapping, not even an explicit nil
### GetDigestAlgorithm

`func (o *PatchedWSFederationProviderRequest) GetDigestAlgorithm() DigestAlgorithmEnum`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *PatchedWSFederationProviderRequest) GetDigestAlgorithmOk() (*DigestAlgorithmEnum, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *PatchedWSFederationProviderRequest) SetDigestAlgorithm(v DigestAlgorithmEnum)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *PatchedWSFederationProviderRequest) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *PatchedWSFederationProviderRequest) GetSignatureAlgorithm() SignatureAlgorithmEnum`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *PatchedWSFederationProviderRequest) GetSignatureAlgorithmOk() (*SignatureAlgorithmEnum, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *PatchedWSFederationProviderRequest) SetSignatureAlgorithm(v SignatureAlgorithmEnum)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *PatchedWSFederationProviderRequest) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetSigningKp

`func (o *PatchedWSFederationProviderRequest) GetSigningKp() string`

GetSigningKp returns the SigningKp field if non-nil, zero value otherwise.

### GetSigningKpOk

`func (o *PatchedWSFederationProviderRequest) GetSigningKpOk() (*string, bool)`

GetSigningKpOk returns a tuple with the SigningKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKp

`func (o *PatchedWSFederationProviderRequest) SetSigningKp(v string)`

SetSigningKp sets SigningKp field to given value.

### HasSigningKp

`func (o *PatchedWSFederationProviderRequest) HasSigningKp() bool`

HasSigningKp returns a boolean if a field has been set.

### SetSigningKpNil

`func (o *PatchedWSFederationProviderRequest) SetSigningKpNil(b bool)`

 SetSigningKpNil sets the value for SigningKp to be an explicit nil

### UnsetSigningKp
`func (o *PatchedWSFederationProviderRequest) UnsetSigningKp()`

UnsetSigningKp ensures that no value is present for SigningKp, not even an explicit nil
### GetEncryptionKp

`func (o *PatchedWSFederationProviderRequest) GetEncryptionKp() string`

GetEncryptionKp returns the EncryptionKp field if non-nil, zero value otherwise.

### GetEncryptionKpOk

`func (o *PatchedWSFederationProviderRequest) GetEncryptionKpOk() (*string, bool)`

GetEncryptionKpOk returns a tuple with the EncryptionKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKp

`func (o *PatchedWSFederationProviderRequest) SetEncryptionKp(v string)`

SetEncryptionKp sets EncryptionKp field to given value.

### HasEncryptionKp

`func (o *PatchedWSFederationProviderRequest) HasEncryptionKp() bool`

HasEncryptionKp returns a boolean if a field has been set.

### SetEncryptionKpNil

`func (o *PatchedWSFederationProviderRequest) SetEncryptionKpNil(b bool)`

 SetEncryptionKpNil sets the value for EncryptionKp to be an explicit nil

### UnsetEncryptionKp
`func (o *PatchedWSFederationProviderRequest) UnsetEncryptionKp()`

UnsetEncryptionKp ensures that no value is present for EncryptionKp, not even an explicit nil
### GetSignAssertion

`func (o *PatchedWSFederationProviderRequest) GetSignAssertion() bool`

GetSignAssertion returns the SignAssertion field if non-nil, zero value otherwise.

### GetSignAssertionOk

`func (o *PatchedWSFederationProviderRequest) GetSignAssertionOk() (*bool, bool)`

GetSignAssertionOk returns a tuple with the SignAssertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAssertion

`func (o *PatchedWSFederationProviderRequest) SetSignAssertion(v bool)`

SetSignAssertion sets SignAssertion field to given value.

### HasSignAssertion

`func (o *PatchedWSFederationProviderRequest) HasSignAssertion() bool`

HasSignAssertion returns a boolean if a field has been set.

### GetSignLogoutRequest

`func (o *PatchedWSFederationProviderRequest) GetSignLogoutRequest() bool`

GetSignLogoutRequest returns the SignLogoutRequest field if non-nil, zero value otherwise.

### GetSignLogoutRequestOk

`func (o *PatchedWSFederationProviderRequest) GetSignLogoutRequestOk() (*bool, bool)`

GetSignLogoutRequestOk returns a tuple with the SignLogoutRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignLogoutRequest

`func (o *PatchedWSFederationProviderRequest) SetSignLogoutRequest(v bool)`

SetSignLogoutRequest sets SignLogoutRequest field to given value.

### HasSignLogoutRequest

`func (o *PatchedWSFederationProviderRequest) HasSignLogoutRequest() bool`

HasSignLogoutRequest returns a boolean if a field has been set.

### GetDefaultNameIdPolicy

`func (o *PatchedWSFederationProviderRequest) GetDefaultNameIdPolicy() SAMLNameIDPolicyEnum`

GetDefaultNameIdPolicy returns the DefaultNameIdPolicy field if non-nil, zero value otherwise.

### GetDefaultNameIdPolicyOk

`func (o *PatchedWSFederationProviderRequest) GetDefaultNameIdPolicyOk() (*SAMLNameIDPolicyEnum, bool)`

GetDefaultNameIdPolicyOk returns a tuple with the DefaultNameIdPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNameIdPolicy

`func (o *PatchedWSFederationProviderRequest) SetDefaultNameIdPolicy(v SAMLNameIDPolicyEnum)`

SetDefaultNameIdPolicy sets DefaultNameIdPolicy field to given value.

### HasDefaultNameIdPolicy

`func (o *PatchedWSFederationProviderRequest) HasDefaultNameIdPolicy() bool`

HasDefaultNameIdPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


