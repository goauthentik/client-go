# SSFProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Component** | **string** | Get object component so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**SigningKey** | **string** | Key used to sign the SSF Events. | 
**TokenObj** | [**Token**](Token.md) |  | [readonly] 
**OidcAuthProviders** | Pointer to **[]int32** |  | [optional] 
**SsfUrl** | **NullableString** |  | [readonly] 
**EventRetention** | Pointer to **string** |  | [optional] 

## Methods

### NewSSFProvider

`func NewSSFProvider(pk int32, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, signingKey string, tokenObj Token, ssfUrl NullableString, ) *SSFProvider`

NewSSFProvider instantiates a new SSFProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSFProviderWithDefaults

`func NewSSFProviderWithDefaults() *SSFProvider`

NewSSFProviderWithDefaults instantiates a new SSFProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *SSFProvider) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *SSFProvider) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *SSFProvider) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetName

`func (o *SSFProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SSFProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SSFProvider) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *SSFProvider) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *SSFProvider) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *SSFProvider) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *SSFProvider) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *SSFProvider) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *SSFProvider) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *SSFProvider) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *SSFProvider) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *SSFProvider) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *SSFProvider) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *SSFProvider) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *SSFProvider) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetSigningKey

`func (o *SSFProvider) GetSigningKey() string`

GetSigningKey returns the SigningKey field if non-nil, zero value otherwise.

### GetSigningKeyOk

`func (o *SSFProvider) GetSigningKeyOk() (*string, bool)`

GetSigningKeyOk returns a tuple with the SigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey

`func (o *SSFProvider) SetSigningKey(v string)`

SetSigningKey sets SigningKey field to given value.


### GetTokenObj

`func (o *SSFProvider) GetTokenObj() Token`

GetTokenObj returns the TokenObj field if non-nil, zero value otherwise.

### GetTokenObjOk

`func (o *SSFProvider) GetTokenObjOk() (*Token, bool)`

GetTokenObjOk returns a tuple with the TokenObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenObj

`func (o *SSFProvider) SetTokenObj(v Token)`

SetTokenObj sets TokenObj field to given value.


### GetOidcAuthProviders

`func (o *SSFProvider) GetOidcAuthProviders() []int32`

GetOidcAuthProviders returns the OidcAuthProviders field if non-nil, zero value otherwise.

### GetOidcAuthProvidersOk

`func (o *SSFProvider) GetOidcAuthProvidersOk() (*[]int32, bool)`

GetOidcAuthProvidersOk returns a tuple with the OidcAuthProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAuthProviders

`func (o *SSFProvider) SetOidcAuthProviders(v []int32)`

SetOidcAuthProviders sets OidcAuthProviders field to given value.

### HasOidcAuthProviders

`func (o *SSFProvider) HasOidcAuthProviders() bool`

HasOidcAuthProviders returns a boolean if a field has been set.

### GetSsfUrl

`func (o *SSFProvider) GetSsfUrl() string`

GetSsfUrl returns the SsfUrl field if non-nil, zero value otherwise.

### GetSsfUrlOk

`func (o *SSFProvider) GetSsfUrlOk() (*string, bool)`

GetSsfUrlOk returns a tuple with the SsfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsfUrl

`func (o *SSFProvider) SetSsfUrl(v string)`

SetSsfUrl sets SsfUrl field to given value.


### SetSsfUrlNil

`func (o *SSFProvider) SetSsfUrlNil(b bool)`

 SetSsfUrlNil sets the value for SsfUrl to be an explicit nil

### UnsetSsfUrl
`func (o *SSFProvider) UnsetSsfUrl()`

UnsetSsfUrl ensures that no value is present for SsfUrl, not even an explicit nil
### GetEventRetention

`func (o *SSFProvider) GetEventRetention() string`

GetEventRetention returns the EventRetention field if non-nil, zero value otherwise.

### GetEventRetentionOk

`func (o *SSFProvider) GetEventRetentionOk() (*string, bool)`

GetEventRetentionOk returns a tuple with the EventRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRetention

`func (o *SSFProvider) SetEventRetention(v string)`

SetEventRetention sets EventRetention field to given value.

### HasEventRetention

`func (o *SSFProvider) HasEventRetention() bool`

HasEventRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


