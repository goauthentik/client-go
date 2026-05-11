# WSFederationProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AuthenticationFlow** | Pointer to **NullableString** | Flow used for authentication when the associated application is accessed by an un-authenticated user. | [optional] 
**AuthorizationFlow** | **string** | Flow used when authorizing this provider. | 
**InvalidationFlow** | **string** | Flow used ending the session from a provider. | 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**ReplyUrl** | **string** |  | 
**Wtrealm** | **string** |  | 
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

### NewWSFederationProviderRequest

`func NewWSFederationProviderRequest(name string, authorizationFlow string, invalidationFlow string, replyUrl string, wtrealm string, ) *WSFederationProviderRequest`

NewWSFederationProviderRequest instantiates a new WSFederationProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWSFederationProviderRequestWithDefaults

`func NewWSFederationProviderRequestWithDefaults() *WSFederationProviderRequest`

NewWSFederationProviderRequestWithDefaults instantiates a new WSFederationProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WSFederationProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WSFederationProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WSFederationProviderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAuthenticationFlow

`func (o *WSFederationProviderRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *WSFederationProviderRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *WSFederationProviderRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *WSFederationProviderRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *WSFederationProviderRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *WSFederationProviderRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetAuthorizationFlow

`func (o *WSFederationProviderRequest) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *WSFederationProviderRequest) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *WSFederationProviderRequest) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.


### GetInvalidationFlow

`func (o *WSFederationProviderRequest) GetInvalidationFlow() string`

GetInvalidationFlow returns the InvalidationFlow field if non-nil, zero value otherwise.

### GetInvalidationFlowOk

`func (o *WSFederationProviderRequest) GetInvalidationFlowOk() (*string, bool)`

GetInvalidationFlowOk returns a tuple with the InvalidationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidationFlow

`func (o *WSFederationProviderRequest) SetInvalidationFlow(v string)`

SetInvalidationFlow sets InvalidationFlow field to given value.


### GetPropertyMappings

`func (o *WSFederationProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *WSFederationProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *WSFederationProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *WSFederationProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetReplyUrl

`func (o *WSFederationProviderRequest) GetReplyUrl() string`

GetReplyUrl returns the ReplyUrl field if non-nil, zero value otherwise.

### GetReplyUrlOk

`func (o *WSFederationProviderRequest) GetReplyUrlOk() (*string, bool)`

GetReplyUrlOk returns a tuple with the ReplyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyUrl

`func (o *WSFederationProviderRequest) SetReplyUrl(v string)`

SetReplyUrl sets ReplyUrl field to given value.


### GetWtrealm

`func (o *WSFederationProviderRequest) GetWtrealm() string`

GetWtrealm returns the Wtrealm field if non-nil, zero value otherwise.

### GetWtrealmOk

`func (o *WSFederationProviderRequest) GetWtrealmOk() (*string, bool)`

GetWtrealmOk returns a tuple with the Wtrealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWtrealm

`func (o *WSFederationProviderRequest) SetWtrealm(v string)`

SetWtrealm sets Wtrealm field to given value.


### GetAssertionValidNotBefore

`func (o *WSFederationProviderRequest) GetAssertionValidNotBefore() string`

GetAssertionValidNotBefore returns the AssertionValidNotBefore field if non-nil, zero value otherwise.

### GetAssertionValidNotBeforeOk

`func (o *WSFederationProviderRequest) GetAssertionValidNotBeforeOk() (*string, bool)`

GetAssertionValidNotBeforeOk returns a tuple with the AssertionValidNotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionValidNotBefore

`func (o *WSFederationProviderRequest) SetAssertionValidNotBefore(v string)`

SetAssertionValidNotBefore sets AssertionValidNotBefore field to given value.

### HasAssertionValidNotBefore

`func (o *WSFederationProviderRequest) HasAssertionValidNotBefore() bool`

HasAssertionValidNotBefore returns a boolean if a field has been set.

### GetAssertionValidNotOnOrAfter

`func (o *WSFederationProviderRequest) GetAssertionValidNotOnOrAfter() string`

GetAssertionValidNotOnOrAfter returns the AssertionValidNotOnOrAfter field if non-nil, zero value otherwise.

### GetAssertionValidNotOnOrAfterOk

`func (o *WSFederationProviderRequest) GetAssertionValidNotOnOrAfterOk() (*string, bool)`

GetAssertionValidNotOnOrAfterOk returns a tuple with the AssertionValidNotOnOrAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionValidNotOnOrAfter

`func (o *WSFederationProviderRequest) SetAssertionValidNotOnOrAfter(v string)`

SetAssertionValidNotOnOrAfter sets AssertionValidNotOnOrAfter field to given value.

### HasAssertionValidNotOnOrAfter

`func (o *WSFederationProviderRequest) HasAssertionValidNotOnOrAfter() bool`

HasAssertionValidNotOnOrAfter returns a boolean if a field has been set.

### GetSessionValidNotOnOrAfter

`func (o *WSFederationProviderRequest) GetSessionValidNotOnOrAfter() string`

GetSessionValidNotOnOrAfter returns the SessionValidNotOnOrAfter field if non-nil, zero value otherwise.

### GetSessionValidNotOnOrAfterOk

`func (o *WSFederationProviderRequest) GetSessionValidNotOnOrAfterOk() (*string, bool)`

GetSessionValidNotOnOrAfterOk returns a tuple with the SessionValidNotOnOrAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionValidNotOnOrAfter

`func (o *WSFederationProviderRequest) SetSessionValidNotOnOrAfter(v string)`

SetSessionValidNotOnOrAfter sets SessionValidNotOnOrAfter field to given value.

### HasSessionValidNotOnOrAfter

`func (o *WSFederationProviderRequest) HasSessionValidNotOnOrAfter() bool`

HasSessionValidNotOnOrAfter returns a boolean if a field has been set.

### GetNameIdMapping

`func (o *WSFederationProviderRequest) GetNameIdMapping() string`

GetNameIdMapping returns the NameIdMapping field if non-nil, zero value otherwise.

### GetNameIdMappingOk

`func (o *WSFederationProviderRequest) GetNameIdMappingOk() (*string, bool)`

GetNameIdMappingOk returns a tuple with the NameIdMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdMapping

`func (o *WSFederationProviderRequest) SetNameIdMapping(v string)`

SetNameIdMapping sets NameIdMapping field to given value.

### HasNameIdMapping

`func (o *WSFederationProviderRequest) HasNameIdMapping() bool`

HasNameIdMapping returns a boolean if a field has been set.

### SetNameIdMappingNil

`func (o *WSFederationProviderRequest) SetNameIdMappingNil(b bool)`

 SetNameIdMappingNil sets the value for NameIdMapping to be an explicit nil

### UnsetNameIdMapping
`func (o *WSFederationProviderRequest) UnsetNameIdMapping()`

UnsetNameIdMapping ensures that no value is present for NameIdMapping, not even an explicit nil
### GetAuthnContextClassRefMapping

`func (o *WSFederationProviderRequest) GetAuthnContextClassRefMapping() string`

GetAuthnContextClassRefMapping returns the AuthnContextClassRefMapping field if non-nil, zero value otherwise.

### GetAuthnContextClassRefMappingOk

`func (o *WSFederationProviderRequest) GetAuthnContextClassRefMappingOk() (*string, bool)`

GetAuthnContextClassRefMappingOk returns a tuple with the AuthnContextClassRefMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnContextClassRefMapping

`func (o *WSFederationProviderRequest) SetAuthnContextClassRefMapping(v string)`

SetAuthnContextClassRefMapping sets AuthnContextClassRefMapping field to given value.

### HasAuthnContextClassRefMapping

`func (o *WSFederationProviderRequest) HasAuthnContextClassRefMapping() bool`

HasAuthnContextClassRefMapping returns a boolean if a field has been set.

### SetAuthnContextClassRefMappingNil

`func (o *WSFederationProviderRequest) SetAuthnContextClassRefMappingNil(b bool)`

 SetAuthnContextClassRefMappingNil sets the value for AuthnContextClassRefMapping to be an explicit nil

### UnsetAuthnContextClassRefMapping
`func (o *WSFederationProviderRequest) UnsetAuthnContextClassRefMapping()`

UnsetAuthnContextClassRefMapping ensures that no value is present for AuthnContextClassRefMapping, not even an explicit nil
### GetDigestAlgorithm

`func (o *WSFederationProviderRequest) GetDigestAlgorithm() DigestAlgorithmEnum`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *WSFederationProviderRequest) GetDigestAlgorithmOk() (*DigestAlgorithmEnum, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *WSFederationProviderRequest) SetDigestAlgorithm(v DigestAlgorithmEnum)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *WSFederationProviderRequest) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *WSFederationProviderRequest) GetSignatureAlgorithm() SignatureAlgorithmEnum`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *WSFederationProviderRequest) GetSignatureAlgorithmOk() (*SignatureAlgorithmEnum, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *WSFederationProviderRequest) SetSignatureAlgorithm(v SignatureAlgorithmEnum)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *WSFederationProviderRequest) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetSigningKp

`func (o *WSFederationProviderRequest) GetSigningKp() string`

GetSigningKp returns the SigningKp field if non-nil, zero value otherwise.

### GetSigningKpOk

`func (o *WSFederationProviderRequest) GetSigningKpOk() (*string, bool)`

GetSigningKpOk returns a tuple with the SigningKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKp

`func (o *WSFederationProviderRequest) SetSigningKp(v string)`

SetSigningKp sets SigningKp field to given value.

### HasSigningKp

`func (o *WSFederationProviderRequest) HasSigningKp() bool`

HasSigningKp returns a boolean if a field has been set.

### SetSigningKpNil

`func (o *WSFederationProviderRequest) SetSigningKpNil(b bool)`

 SetSigningKpNil sets the value for SigningKp to be an explicit nil

### UnsetSigningKp
`func (o *WSFederationProviderRequest) UnsetSigningKp()`

UnsetSigningKp ensures that no value is present for SigningKp, not even an explicit nil
### GetEncryptionKp

`func (o *WSFederationProviderRequest) GetEncryptionKp() string`

GetEncryptionKp returns the EncryptionKp field if non-nil, zero value otherwise.

### GetEncryptionKpOk

`func (o *WSFederationProviderRequest) GetEncryptionKpOk() (*string, bool)`

GetEncryptionKpOk returns a tuple with the EncryptionKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKp

`func (o *WSFederationProviderRequest) SetEncryptionKp(v string)`

SetEncryptionKp sets EncryptionKp field to given value.

### HasEncryptionKp

`func (o *WSFederationProviderRequest) HasEncryptionKp() bool`

HasEncryptionKp returns a boolean if a field has been set.

### SetEncryptionKpNil

`func (o *WSFederationProviderRequest) SetEncryptionKpNil(b bool)`

 SetEncryptionKpNil sets the value for EncryptionKp to be an explicit nil

### UnsetEncryptionKp
`func (o *WSFederationProviderRequest) UnsetEncryptionKp()`

UnsetEncryptionKp ensures that no value is present for EncryptionKp, not even an explicit nil
### GetSignAssertion

`func (o *WSFederationProviderRequest) GetSignAssertion() bool`

GetSignAssertion returns the SignAssertion field if non-nil, zero value otherwise.

### GetSignAssertionOk

`func (o *WSFederationProviderRequest) GetSignAssertionOk() (*bool, bool)`

GetSignAssertionOk returns a tuple with the SignAssertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAssertion

`func (o *WSFederationProviderRequest) SetSignAssertion(v bool)`

SetSignAssertion sets SignAssertion field to given value.

### HasSignAssertion

`func (o *WSFederationProviderRequest) HasSignAssertion() bool`

HasSignAssertion returns a boolean if a field has been set.

### GetSignLogoutRequest

`func (o *WSFederationProviderRequest) GetSignLogoutRequest() bool`

GetSignLogoutRequest returns the SignLogoutRequest field if non-nil, zero value otherwise.

### GetSignLogoutRequestOk

`func (o *WSFederationProviderRequest) GetSignLogoutRequestOk() (*bool, bool)`

GetSignLogoutRequestOk returns a tuple with the SignLogoutRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignLogoutRequest

`func (o *WSFederationProviderRequest) SetSignLogoutRequest(v bool)`

SetSignLogoutRequest sets SignLogoutRequest field to given value.

### HasSignLogoutRequest

`func (o *WSFederationProviderRequest) HasSignLogoutRequest() bool`

HasSignLogoutRequest returns a boolean if a field has been set.

### GetDefaultNameIdPolicy

`func (o *WSFederationProviderRequest) GetDefaultNameIdPolicy() SAMLNameIDPolicyEnum`

GetDefaultNameIdPolicy returns the DefaultNameIdPolicy field if non-nil, zero value otherwise.

### GetDefaultNameIdPolicyOk

`func (o *WSFederationProviderRequest) GetDefaultNameIdPolicyOk() (*SAMLNameIDPolicyEnum, bool)`

GetDefaultNameIdPolicyOk returns a tuple with the DefaultNameIdPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNameIdPolicy

`func (o *WSFederationProviderRequest) SetDefaultNameIdPolicy(v SAMLNameIDPolicyEnum)`

SetDefaultNameIdPolicy sets DefaultNameIdPolicy field to given value.

### HasDefaultNameIdPolicy

`func (o *WSFederationProviderRequest) HasDefaultNameIdPolicy() bool`

HasDefaultNameIdPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


