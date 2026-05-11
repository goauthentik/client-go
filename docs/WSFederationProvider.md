# WSFederationProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**AuthenticationFlow** | Pointer to **NullableString** | Flow used for authentication when the associated application is accessed by an un-authenticated user. | [optional] 
**AuthorizationFlow** | **string** | Flow used when authorizing this provider. | 
**InvalidationFlow** | **string** | Flow used ending the session from a provider. | 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**Component** | **string** | Get object component so that we know how to edit the object | [readonly] 
**AssignedApplicationSlug** | **NullableString** | Internal application name, used in URLs. | [readonly] 
**AssignedApplicationName** | **NullableString** | Application&#39;s display Name. | [readonly] 
**AssignedBackchannelApplicationSlug** | **NullableString** | Internal application name, used in URLs. | [readonly] 
**AssignedBackchannelApplicationName** | **NullableString** | Application&#39;s display Name. | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
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
**UrlDownloadMetadata** | **string** | Get metadata download URL | [readonly] 
**UrlWsfed** | **string** | Get WS-Fed url | [readonly] 

## Methods

### NewWSFederationProvider

`func NewWSFederationProvider(pk int32, name string, authorizationFlow string, invalidationFlow string, component string, assignedApplicationSlug NullableString, assignedApplicationName NullableString, assignedBackchannelApplicationSlug NullableString, assignedBackchannelApplicationName NullableString, verboseName string, verboseNamePlural string, metaModelName string, replyUrl string, wtrealm string, urlDownloadMetadata string, urlWsfed string, ) *WSFederationProvider`

NewWSFederationProvider instantiates a new WSFederationProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWSFederationProviderWithDefaults

`func NewWSFederationProviderWithDefaults() *WSFederationProvider`

NewWSFederationProviderWithDefaults instantiates a new WSFederationProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *WSFederationProvider) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *WSFederationProvider) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *WSFederationProvider) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetName

`func (o *WSFederationProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WSFederationProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WSFederationProvider) SetName(v string)`

SetName sets Name field to given value.


### GetAuthenticationFlow

`func (o *WSFederationProvider) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *WSFederationProvider) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *WSFederationProvider) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *WSFederationProvider) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *WSFederationProvider) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *WSFederationProvider) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetAuthorizationFlow

`func (o *WSFederationProvider) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *WSFederationProvider) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *WSFederationProvider) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.


### GetInvalidationFlow

`func (o *WSFederationProvider) GetInvalidationFlow() string`

GetInvalidationFlow returns the InvalidationFlow field if non-nil, zero value otherwise.

### GetInvalidationFlowOk

`func (o *WSFederationProvider) GetInvalidationFlowOk() (*string, bool)`

GetInvalidationFlowOk returns a tuple with the InvalidationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidationFlow

`func (o *WSFederationProvider) SetInvalidationFlow(v string)`

SetInvalidationFlow sets InvalidationFlow field to given value.


### GetPropertyMappings

`func (o *WSFederationProvider) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *WSFederationProvider) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *WSFederationProvider) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *WSFederationProvider) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetComponent

`func (o *WSFederationProvider) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *WSFederationProvider) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *WSFederationProvider) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetAssignedApplicationSlug

`func (o *WSFederationProvider) GetAssignedApplicationSlug() string`

GetAssignedApplicationSlug returns the AssignedApplicationSlug field if non-nil, zero value otherwise.

### GetAssignedApplicationSlugOk

`func (o *WSFederationProvider) GetAssignedApplicationSlugOk() (*string, bool)`

GetAssignedApplicationSlugOk returns a tuple with the AssignedApplicationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedApplicationSlug

`func (o *WSFederationProvider) SetAssignedApplicationSlug(v string)`

SetAssignedApplicationSlug sets AssignedApplicationSlug field to given value.


### SetAssignedApplicationSlugNil

`func (o *WSFederationProvider) SetAssignedApplicationSlugNil(b bool)`

 SetAssignedApplicationSlugNil sets the value for AssignedApplicationSlug to be an explicit nil

### UnsetAssignedApplicationSlug
`func (o *WSFederationProvider) UnsetAssignedApplicationSlug()`

UnsetAssignedApplicationSlug ensures that no value is present for AssignedApplicationSlug, not even an explicit nil
### GetAssignedApplicationName

`func (o *WSFederationProvider) GetAssignedApplicationName() string`

GetAssignedApplicationName returns the AssignedApplicationName field if non-nil, zero value otherwise.

### GetAssignedApplicationNameOk

`func (o *WSFederationProvider) GetAssignedApplicationNameOk() (*string, bool)`

GetAssignedApplicationNameOk returns a tuple with the AssignedApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedApplicationName

`func (o *WSFederationProvider) SetAssignedApplicationName(v string)`

SetAssignedApplicationName sets AssignedApplicationName field to given value.


### SetAssignedApplicationNameNil

`func (o *WSFederationProvider) SetAssignedApplicationNameNil(b bool)`

 SetAssignedApplicationNameNil sets the value for AssignedApplicationName to be an explicit nil

### UnsetAssignedApplicationName
`func (o *WSFederationProvider) UnsetAssignedApplicationName()`

UnsetAssignedApplicationName ensures that no value is present for AssignedApplicationName, not even an explicit nil
### GetAssignedBackchannelApplicationSlug

`func (o *WSFederationProvider) GetAssignedBackchannelApplicationSlug() string`

GetAssignedBackchannelApplicationSlug returns the AssignedBackchannelApplicationSlug field if non-nil, zero value otherwise.

### GetAssignedBackchannelApplicationSlugOk

`func (o *WSFederationProvider) GetAssignedBackchannelApplicationSlugOk() (*string, bool)`

GetAssignedBackchannelApplicationSlugOk returns a tuple with the AssignedBackchannelApplicationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedBackchannelApplicationSlug

`func (o *WSFederationProvider) SetAssignedBackchannelApplicationSlug(v string)`

SetAssignedBackchannelApplicationSlug sets AssignedBackchannelApplicationSlug field to given value.


### SetAssignedBackchannelApplicationSlugNil

`func (o *WSFederationProvider) SetAssignedBackchannelApplicationSlugNil(b bool)`

 SetAssignedBackchannelApplicationSlugNil sets the value for AssignedBackchannelApplicationSlug to be an explicit nil

### UnsetAssignedBackchannelApplicationSlug
`func (o *WSFederationProvider) UnsetAssignedBackchannelApplicationSlug()`

UnsetAssignedBackchannelApplicationSlug ensures that no value is present for AssignedBackchannelApplicationSlug, not even an explicit nil
### GetAssignedBackchannelApplicationName

`func (o *WSFederationProvider) GetAssignedBackchannelApplicationName() string`

GetAssignedBackchannelApplicationName returns the AssignedBackchannelApplicationName field if non-nil, zero value otherwise.

### GetAssignedBackchannelApplicationNameOk

`func (o *WSFederationProvider) GetAssignedBackchannelApplicationNameOk() (*string, bool)`

GetAssignedBackchannelApplicationNameOk returns a tuple with the AssignedBackchannelApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedBackchannelApplicationName

`func (o *WSFederationProvider) SetAssignedBackchannelApplicationName(v string)`

SetAssignedBackchannelApplicationName sets AssignedBackchannelApplicationName field to given value.


### SetAssignedBackchannelApplicationNameNil

`func (o *WSFederationProvider) SetAssignedBackchannelApplicationNameNil(b bool)`

 SetAssignedBackchannelApplicationNameNil sets the value for AssignedBackchannelApplicationName to be an explicit nil

### UnsetAssignedBackchannelApplicationName
`func (o *WSFederationProvider) UnsetAssignedBackchannelApplicationName()`

UnsetAssignedBackchannelApplicationName ensures that no value is present for AssignedBackchannelApplicationName, not even an explicit nil
### GetVerboseName

`func (o *WSFederationProvider) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *WSFederationProvider) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *WSFederationProvider) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *WSFederationProvider) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *WSFederationProvider) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *WSFederationProvider) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *WSFederationProvider) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *WSFederationProvider) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *WSFederationProvider) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetReplyUrl

`func (o *WSFederationProvider) GetReplyUrl() string`

GetReplyUrl returns the ReplyUrl field if non-nil, zero value otherwise.

### GetReplyUrlOk

`func (o *WSFederationProvider) GetReplyUrlOk() (*string, bool)`

GetReplyUrlOk returns a tuple with the ReplyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyUrl

`func (o *WSFederationProvider) SetReplyUrl(v string)`

SetReplyUrl sets ReplyUrl field to given value.


### GetWtrealm

`func (o *WSFederationProvider) GetWtrealm() string`

GetWtrealm returns the Wtrealm field if non-nil, zero value otherwise.

### GetWtrealmOk

`func (o *WSFederationProvider) GetWtrealmOk() (*string, bool)`

GetWtrealmOk returns a tuple with the Wtrealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWtrealm

`func (o *WSFederationProvider) SetWtrealm(v string)`

SetWtrealm sets Wtrealm field to given value.


### GetAssertionValidNotBefore

`func (o *WSFederationProvider) GetAssertionValidNotBefore() string`

GetAssertionValidNotBefore returns the AssertionValidNotBefore field if non-nil, zero value otherwise.

### GetAssertionValidNotBeforeOk

`func (o *WSFederationProvider) GetAssertionValidNotBeforeOk() (*string, bool)`

GetAssertionValidNotBeforeOk returns a tuple with the AssertionValidNotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionValidNotBefore

`func (o *WSFederationProvider) SetAssertionValidNotBefore(v string)`

SetAssertionValidNotBefore sets AssertionValidNotBefore field to given value.

### HasAssertionValidNotBefore

`func (o *WSFederationProvider) HasAssertionValidNotBefore() bool`

HasAssertionValidNotBefore returns a boolean if a field has been set.

### GetAssertionValidNotOnOrAfter

`func (o *WSFederationProvider) GetAssertionValidNotOnOrAfter() string`

GetAssertionValidNotOnOrAfter returns the AssertionValidNotOnOrAfter field if non-nil, zero value otherwise.

### GetAssertionValidNotOnOrAfterOk

`func (o *WSFederationProvider) GetAssertionValidNotOnOrAfterOk() (*string, bool)`

GetAssertionValidNotOnOrAfterOk returns a tuple with the AssertionValidNotOnOrAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionValidNotOnOrAfter

`func (o *WSFederationProvider) SetAssertionValidNotOnOrAfter(v string)`

SetAssertionValidNotOnOrAfter sets AssertionValidNotOnOrAfter field to given value.

### HasAssertionValidNotOnOrAfter

`func (o *WSFederationProvider) HasAssertionValidNotOnOrAfter() bool`

HasAssertionValidNotOnOrAfter returns a boolean if a field has been set.

### GetSessionValidNotOnOrAfter

`func (o *WSFederationProvider) GetSessionValidNotOnOrAfter() string`

GetSessionValidNotOnOrAfter returns the SessionValidNotOnOrAfter field if non-nil, zero value otherwise.

### GetSessionValidNotOnOrAfterOk

`func (o *WSFederationProvider) GetSessionValidNotOnOrAfterOk() (*string, bool)`

GetSessionValidNotOnOrAfterOk returns a tuple with the SessionValidNotOnOrAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionValidNotOnOrAfter

`func (o *WSFederationProvider) SetSessionValidNotOnOrAfter(v string)`

SetSessionValidNotOnOrAfter sets SessionValidNotOnOrAfter field to given value.

### HasSessionValidNotOnOrAfter

`func (o *WSFederationProvider) HasSessionValidNotOnOrAfter() bool`

HasSessionValidNotOnOrAfter returns a boolean if a field has been set.

### GetNameIdMapping

`func (o *WSFederationProvider) GetNameIdMapping() string`

GetNameIdMapping returns the NameIdMapping field if non-nil, zero value otherwise.

### GetNameIdMappingOk

`func (o *WSFederationProvider) GetNameIdMappingOk() (*string, bool)`

GetNameIdMappingOk returns a tuple with the NameIdMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdMapping

`func (o *WSFederationProvider) SetNameIdMapping(v string)`

SetNameIdMapping sets NameIdMapping field to given value.

### HasNameIdMapping

`func (o *WSFederationProvider) HasNameIdMapping() bool`

HasNameIdMapping returns a boolean if a field has been set.

### SetNameIdMappingNil

`func (o *WSFederationProvider) SetNameIdMappingNil(b bool)`

 SetNameIdMappingNil sets the value for NameIdMapping to be an explicit nil

### UnsetNameIdMapping
`func (o *WSFederationProvider) UnsetNameIdMapping()`

UnsetNameIdMapping ensures that no value is present for NameIdMapping, not even an explicit nil
### GetAuthnContextClassRefMapping

`func (o *WSFederationProvider) GetAuthnContextClassRefMapping() string`

GetAuthnContextClassRefMapping returns the AuthnContextClassRefMapping field if non-nil, zero value otherwise.

### GetAuthnContextClassRefMappingOk

`func (o *WSFederationProvider) GetAuthnContextClassRefMappingOk() (*string, bool)`

GetAuthnContextClassRefMappingOk returns a tuple with the AuthnContextClassRefMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnContextClassRefMapping

`func (o *WSFederationProvider) SetAuthnContextClassRefMapping(v string)`

SetAuthnContextClassRefMapping sets AuthnContextClassRefMapping field to given value.

### HasAuthnContextClassRefMapping

`func (o *WSFederationProvider) HasAuthnContextClassRefMapping() bool`

HasAuthnContextClassRefMapping returns a boolean if a field has been set.

### SetAuthnContextClassRefMappingNil

`func (o *WSFederationProvider) SetAuthnContextClassRefMappingNil(b bool)`

 SetAuthnContextClassRefMappingNil sets the value for AuthnContextClassRefMapping to be an explicit nil

### UnsetAuthnContextClassRefMapping
`func (o *WSFederationProvider) UnsetAuthnContextClassRefMapping()`

UnsetAuthnContextClassRefMapping ensures that no value is present for AuthnContextClassRefMapping, not even an explicit nil
### GetDigestAlgorithm

`func (o *WSFederationProvider) GetDigestAlgorithm() DigestAlgorithmEnum`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *WSFederationProvider) GetDigestAlgorithmOk() (*DigestAlgorithmEnum, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *WSFederationProvider) SetDigestAlgorithm(v DigestAlgorithmEnum)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *WSFederationProvider) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *WSFederationProvider) GetSignatureAlgorithm() SignatureAlgorithmEnum`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *WSFederationProvider) GetSignatureAlgorithmOk() (*SignatureAlgorithmEnum, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *WSFederationProvider) SetSignatureAlgorithm(v SignatureAlgorithmEnum)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *WSFederationProvider) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetSigningKp

`func (o *WSFederationProvider) GetSigningKp() string`

GetSigningKp returns the SigningKp field if non-nil, zero value otherwise.

### GetSigningKpOk

`func (o *WSFederationProvider) GetSigningKpOk() (*string, bool)`

GetSigningKpOk returns a tuple with the SigningKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKp

`func (o *WSFederationProvider) SetSigningKp(v string)`

SetSigningKp sets SigningKp field to given value.

### HasSigningKp

`func (o *WSFederationProvider) HasSigningKp() bool`

HasSigningKp returns a boolean if a field has been set.

### SetSigningKpNil

`func (o *WSFederationProvider) SetSigningKpNil(b bool)`

 SetSigningKpNil sets the value for SigningKp to be an explicit nil

### UnsetSigningKp
`func (o *WSFederationProvider) UnsetSigningKp()`

UnsetSigningKp ensures that no value is present for SigningKp, not even an explicit nil
### GetEncryptionKp

`func (o *WSFederationProvider) GetEncryptionKp() string`

GetEncryptionKp returns the EncryptionKp field if non-nil, zero value otherwise.

### GetEncryptionKpOk

`func (o *WSFederationProvider) GetEncryptionKpOk() (*string, bool)`

GetEncryptionKpOk returns a tuple with the EncryptionKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKp

`func (o *WSFederationProvider) SetEncryptionKp(v string)`

SetEncryptionKp sets EncryptionKp field to given value.

### HasEncryptionKp

`func (o *WSFederationProvider) HasEncryptionKp() bool`

HasEncryptionKp returns a boolean if a field has been set.

### SetEncryptionKpNil

`func (o *WSFederationProvider) SetEncryptionKpNil(b bool)`

 SetEncryptionKpNil sets the value for EncryptionKp to be an explicit nil

### UnsetEncryptionKp
`func (o *WSFederationProvider) UnsetEncryptionKp()`

UnsetEncryptionKp ensures that no value is present for EncryptionKp, not even an explicit nil
### GetSignAssertion

`func (o *WSFederationProvider) GetSignAssertion() bool`

GetSignAssertion returns the SignAssertion field if non-nil, zero value otherwise.

### GetSignAssertionOk

`func (o *WSFederationProvider) GetSignAssertionOk() (*bool, bool)`

GetSignAssertionOk returns a tuple with the SignAssertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAssertion

`func (o *WSFederationProvider) SetSignAssertion(v bool)`

SetSignAssertion sets SignAssertion field to given value.

### HasSignAssertion

`func (o *WSFederationProvider) HasSignAssertion() bool`

HasSignAssertion returns a boolean if a field has been set.

### GetSignLogoutRequest

`func (o *WSFederationProvider) GetSignLogoutRequest() bool`

GetSignLogoutRequest returns the SignLogoutRequest field if non-nil, zero value otherwise.

### GetSignLogoutRequestOk

`func (o *WSFederationProvider) GetSignLogoutRequestOk() (*bool, bool)`

GetSignLogoutRequestOk returns a tuple with the SignLogoutRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignLogoutRequest

`func (o *WSFederationProvider) SetSignLogoutRequest(v bool)`

SetSignLogoutRequest sets SignLogoutRequest field to given value.

### HasSignLogoutRequest

`func (o *WSFederationProvider) HasSignLogoutRequest() bool`

HasSignLogoutRequest returns a boolean if a field has been set.

### GetDefaultNameIdPolicy

`func (o *WSFederationProvider) GetDefaultNameIdPolicy() SAMLNameIDPolicyEnum`

GetDefaultNameIdPolicy returns the DefaultNameIdPolicy field if non-nil, zero value otherwise.

### GetDefaultNameIdPolicyOk

`func (o *WSFederationProvider) GetDefaultNameIdPolicyOk() (*SAMLNameIDPolicyEnum, bool)`

GetDefaultNameIdPolicyOk returns a tuple with the DefaultNameIdPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNameIdPolicy

`func (o *WSFederationProvider) SetDefaultNameIdPolicy(v SAMLNameIDPolicyEnum)`

SetDefaultNameIdPolicy sets DefaultNameIdPolicy field to given value.

### HasDefaultNameIdPolicy

`func (o *WSFederationProvider) HasDefaultNameIdPolicy() bool`

HasDefaultNameIdPolicy returns a boolean if a field has been set.

### GetUrlDownloadMetadata

`func (o *WSFederationProvider) GetUrlDownloadMetadata() string`

GetUrlDownloadMetadata returns the UrlDownloadMetadata field if non-nil, zero value otherwise.

### GetUrlDownloadMetadataOk

`func (o *WSFederationProvider) GetUrlDownloadMetadataOk() (*string, bool)`

GetUrlDownloadMetadataOk returns a tuple with the UrlDownloadMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlDownloadMetadata

`func (o *WSFederationProvider) SetUrlDownloadMetadata(v string)`

SetUrlDownloadMetadata sets UrlDownloadMetadata field to given value.


### GetUrlWsfed

`func (o *WSFederationProvider) GetUrlWsfed() string`

GetUrlWsfed returns the UrlWsfed field if non-nil, zero value otherwise.

### GetUrlWsfedOk

`func (o *WSFederationProvider) GetUrlWsfedOk() (*string, bool)`

GetUrlWsfedOk returns a tuple with the UrlWsfed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlWsfed

`func (o *WSFederationProvider) SetUrlWsfed(v string)`

SetUrlWsfed sets UrlWsfed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


