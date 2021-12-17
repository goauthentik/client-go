# SAMLProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**AuthorizationFlow** | **string** | Flow used when authorizing this provider. | 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**Component** | **string** |  | [readonly] 
**AssignedApplicationSlug** | **string** | Internal application name, used in URLs. | [readonly] 
**AssignedApplicationName** | **string** | Application&#39;s display Name. | [readonly] 
**VerboseName** | **string** |  | [readonly] 
**VerboseNamePlural** | **string** |  | [readonly] 
**MetaModelName** | **string** |  | [readonly] 
**AcsUrl** | **string** |  | 
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
**SpBinding** | Pointer to [**SpBindingEnum**](SpBindingEnum.md) | This determines how authentik sends the response back to the Service Provider. | [optional] 
**MetadataDownloadUrl** | **string** |  | [readonly] 

## Methods

### NewSAMLProvider

`func NewSAMLProvider(pk int32, name string, authorizationFlow string, component string, assignedApplicationSlug string, assignedApplicationName string, verboseName string, verboseNamePlural string, metaModelName string, acsUrl string, metadataDownloadUrl string, ) *SAMLProvider`

NewSAMLProvider instantiates a new SAMLProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLProviderWithDefaults

`func NewSAMLProviderWithDefaults() *SAMLProvider`

NewSAMLProviderWithDefaults instantiates a new SAMLProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *SAMLProvider) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *SAMLProvider) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *SAMLProvider) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetName

`func (o *SAMLProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SAMLProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SAMLProvider) SetName(v string)`

SetName sets Name field to given value.


### GetAuthorizationFlow

`func (o *SAMLProvider) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *SAMLProvider) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *SAMLProvider) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.


### GetPropertyMappings

`func (o *SAMLProvider) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *SAMLProvider) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *SAMLProvider) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *SAMLProvider) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetComponent

`func (o *SAMLProvider) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *SAMLProvider) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *SAMLProvider) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetAssignedApplicationSlug

`func (o *SAMLProvider) GetAssignedApplicationSlug() string`

GetAssignedApplicationSlug returns the AssignedApplicationSlug field if non-nil, zero value otherwise.

### GetAssignedApplicationSlugOk

`func (o *SAMLProvider) GetAssignedApplicationSlugOk() (*string, bool)`

GetAssignedApplicationSlugOk returns a tuple with the AssignedApplicationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedApplicationSlug

`func (o *SAMLProvider) SetAssignedApplicationSlug(v string)`

SetAssignedApplicationSlug sets AssignedApplicationSlug field to given value.


### GetAssignedApplicationName

`func (o *SAMLProvider) GetAssignedApplicationName() string`

GetAssignedApplicationName returns the AssignedApplicationName field if non-nil, zero value otherwise.

### GetAssignedApplicationNameOk

`func (o *SAMLProvider) GetAssignedApplicationNameOk() (*string, bool)`

GetAssignedApplicationNameOk returns a tuple with the AssignedApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedApplicationName

`func (o *SAMLProvider) SetAssignedApplicationName(v string)`

SetAssignedApplicationName sets AssignedApplicationName field to given value.


### GetVerboseName

`func (o *SAMLProvider) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *SAMLProvider) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *SAMLProvider) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *SAMLProvider) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *SAMLProvider) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *SAMLProvider) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *SAMLProvider) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *SAMLProvider) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *SAMLProvider) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetAcsUrl

`func (o *SAMLProvider) GetAcsUrl() string`

GetAcsUrl returns the AcsUrl field if non-nil, zero value otherwise.

### GetAcsUrlOk

`func (o *SAMLProvider) GetAcsUrlOk() (*string, bool)`

GetAcsUrlOk returns a tuple with the AcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsUrl

`func (o *SAMLProvider) SetAcsUrl(v string)`

SetAcsUrl sets AcsUrl field to given value.


### GetAudience

`func (o *SAMLProvider) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *SAMLProvider) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *SAMLProvider) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *SAMLProvider) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetIssuer

`func (o *SAMLProvider) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SAMLProvider) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SAMLProvider) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *SAMLProvider) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetAssertionValidNotBefore

`func (o *SAMLProvider) GetAssertionValidNotBefore() string`

GetAssertionValidNotBefore returns the AssertionValidNotBefore field if non-nil, zero value otherwise.

### GetAssertionValidNotBeforeOk

`func (o *SAMLProvider) GetAssertionValidNotBeforeOk() (*string, bool)`

GetAssertionValidNotBeforeOk returns a tuple with the AssertionValidNotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionValidNotBefore

`func (o *SAMLProvider) SetAssertionValidNotBefore(v string)`

SetAssertionValidNotBefore sets AssertionValidNotBefore field to given value.

### HasAssertionValidNotBefore

`func (o *SAMLProvider) HasAssertionValidNotBefore() bool`

HasAssertionValidNotBefore returns a boolean if a field has been set.

### GetAssertionValidNotOnOrAfter

`func (o *SAMLProvider) GetAssertionValidNotOnOrAfter() string`

GetAssertionValidNotOnOrAfter returns the AssertionValidNotOnOrAfter field if non-nil, zero value otherwise.

### GetAssertionValidNotOnOrAfterOk

`func (o *SAMLProvider) GetAssertionValidNotOnOrAfterOk() (*string, bool)`

GetAssertionValidNotOnOrAfterOk returns a tuple with the AssertionValidNotOnOrAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionValidNotOnOrAfter

`func (o *SAMLProvider) SetAssertionValidNotOnOrAfter(v string)`

SetAssertionValidNotOnOrAfter sets AssertionValidNotOnOrAfter field to given value.

### HasAssertionValidNotOnOrAfter

`func (o *SAMLProvider) HasAssertionValidNotOnOrAfter() bool`

HasAssertionValidNotOnOrAfter returns a boolean if a field has been set.

### GetSessionValidNotOnOrAfter

`func (o *SAMLProvider) GetSessionValidNotOnOrAfter() string`

GetSessionValidNotOnOrAfter returns the SessionValidNotOnOrAfter field if non-nil, zero value otherwise.

### GetSessionValidNotOnOrAfterOk

`func (o *SAMLProvider) GetSessionValidNotOnOrAfterOk() (*string, bool)`

GetSessionValidNotOnOrAfterOk returns a tuple with the SessionValidNotOnOrAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionValidNotOnOrAfter

`func (o *SAMLProvider) SetSessionValidNotOnOrAfter(v string)`

SetSessionValidNotOnOrAfter sets SessionValidNotOnOrAfter field to given value.

### HasSessionValidNotOnOrAfter

`func (o *SAMLProvider) HasSessionValidNotOnOrAfter() bool`

HasSessionValidNotOnOrAfter returns a boolean if a field has been set.

### GetNameIdMapping

`func (o *SAMLProvider) GetNameIdMapping() string`

GetNameIdMapping returns the NameIdMapping field if non-nil, zero value otherwise.

### GetNameIdMappingOk

`func (o *SAMLProvider) GetNameIdMappingOk() (*string, bool)`

GetNameIdMappingOk returns a tuple with the NameIdMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdMapping

`func (o *SAMLProvider) SetNameIdMapping(v string)`

SetNameIdMapping sets NameIdMapping field to given value.

### HasNameIdMapping

`func (o *SAMLProvider) HasNameIdMapping() bool`

HasNameIdMapping returns a boolean if a field has been set.

### SetNameIdMappingNil

`func (o *SAMLProvider) SetNameIdMappingNil(b bool)`

 SetNameIdMappingNil sets the value for NameIdMapping to be an explicit nil

### UnsetNameIdMapping
`func (o *SAMLProvider) UnsetNameIdMapping()`

UnsetNameIdMapping ensures that no value is present for NameIdMapping, not even an explicit nil
### GetDigestAlgorithm

`func (o *SAMLProvider) GetDigestAlgorithm() DigestAlgorithmEnum`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *SAMLProvider) GetDigestAlgorithmOk() (*DigestAlgorithmEnum, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *SAMLProvider) SetDigestAlgorithm(v DigestAlgorithmEnum)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *SAMLProvider) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *SAMLProvider) GetSignatureAlgorithm() SignatureAlgorithmEnum`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *SAMLProvider) GetSignatureAlgorithmOk() (*SignatureAlgorithmEnum, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *SAMLProvider) SetSignatureAlgorithm(v SignatureAlgorithmEnum)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *SAMLProvider) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetSigningKp

`func (o *SAMLProvider) GetSigningKp() string`

GetSigningKp returns the SigningKp field if non-nil, zero value otherwise.

### GetSigningKpOk

`func (o *SAMLProvider) GetSigningKpOk() (*string, bool)`

GetSigningKpOk returns a tuple with the SigningKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKp

`func (o *SAMLProvider) SetSigningKp(v string)`

SetSigningKp sets SigningKp field to given value.

### HasSigningKp

`func (o *SAMLProvider) HasSigningKp() bool`

HasSigningKp returns a boolean if a field has been set.

### SetSigningKpNil

`func (o *SAMLProvider) SetSigningKpNil(b bool)`

 SetSigningKpNil sets the value for SigningKp to be an explicit nil

### UnsetSigningKp
`func (o *SAMLProvider) UnsetSigningKp()`

UnsetSigningKp ensures that no value is present for SigningKp, not even an explicit nil
### GetVerificationKp

`func (o *SAMLProvider) GetVerificationKp() string`

GetVerificationKp returns the VerificationKp field if non-nil, zero value otherwise.

### GetVerificationKpOk

`func (o *SAMLProvider) GetVerificationKpOk() (*string, bool)`

GetVerificationKpOk returns a tuple with the VerificationKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationKp

`func (o *SAMLProvider) SetVerificationKp(v string)`

SetVerificationKp sets VerificationKp field to given value.

### HasVerificationKp

`func (o *SAMLProvider) HasVerificationKp() bool`

HasVerificationKp returns a boolean if a field has been set.

### SetVerificationKpNil

`func (o *SAMLProvider) SetVerificationKpNil(b bool)`

 SetVerificationKpNil sets the value for VerificationKp to be an explicit nil

### UnsetVerificationKp
`func (o *SAMLProvider) UnsetVerificationKp()`

UnsetVerificationKp ensures that no value is present for VerificationKp, not even an explicit nil
### GetSpBinding

`func (o *SAMLProvider) GetSpBinding() SpBindingEnum`

GetSpBinding returns the SpBinding field if non-nil, zero value otherwise.

### GetSpBindingOk

`func (o *SAMLProvider) GetSpBindingOk() (*SpBindingEnum, bool)`

GetSpBindingOk returns a tuple with the SpBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpBinding

`func (o *SAMLProvider) SetSpBinding(v SpBindingEnum)`

SetSpBinding sets SpBinding field to given value.

### HasSpBinding

`func (o *SAMLProvider) HasSpBinding() bool`

HasSpBinding returns a boolean if a field has been set.

### GetMetadataDownloadUrl

`func (o *SAMLProvider) GetMetadataDownloadUrl() string`

GetMetadataDownloadUrl returns the MetadataDownloadUrl field if non-nil, zero value otherwise.

### GetMetadataDownloadUrlOk

`func (o *SAMLProvider) GetMetadataDownloadUrlOk() (*string, bool)`

GetMetadataDownloadUrlOk returns a tuple with the MetadataDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataDownloadUrl

`func (o *SAMLProvider) SetMetadataDownloadUrl(v string)`

SetMetadataDownloadUrl sets MetadataDownloadUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


