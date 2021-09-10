# ApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Application&#39;s display Name. | 
**Slug** | **string** | Internal application name, used in URLs. | 
**Provider** | Pointer to **NullableInt32** |  | [optional] 
**ProviderObj** | Pointer to [**ProviderRequest**](ProviderRequest.md) |  | [optional] 
**MetaLaunchUrl** | Pointer to **string** |  | [optional] 
**MetaDescription** | Pointer to **string** |  | [optional] 
**MetaPublisher** | Pointer to **string** |  | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 

## Methods

### NewApplicationRequest

`func NewApplicationRequest(name string, slug string, ) *ApplicationRequest`

NewApplicationRequest instantiates a new ApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRequestWithDefaults

`func NewApplicationRequestWithDefaults() *ApplicationRequest`

NewApplicationRequestWithDefaults instantiates a new ApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *ApplicationRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ApplicationRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ApplicationRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetProvider

`func (o *ApplicationRequest) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ApplicationRequest) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ApplicationRequest) SetProvider(v int32)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ApplicationRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *ApplicationRequest) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *ApplicationRequest) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetProviderObj

`func (o *ApplicationRequest) GetProviderObj() ProviderRequest`

GetProviderObj returns the ProviderObj field if non-nil, zero value otherwise.

### GetProviderObjOk

`func (o *ApplicationRequest) GetProviderObjOk() (*ProviderRequest, bool)`

GetProviderObjOk returns a tuple with the ProviderObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderObj

`func (o *ApplicationRequest) SetProviderObj(v ProviderRequest)`

SetProviderObj sets ProviderObj field to given value.

### HasProviderObj

`func (o *ApplicationRequest) HasProviderObj() bool`

HasProviderObj returns a boolean if a field has been set.

### GetMetaLaunchUrl

`func (o *ApplicationRequest) GetMetaLaunchUrl() string`

GetMetaLaunchUrl returns the MetaLaunchUrl field if non-nil, zero value otherwise.

### GetMetaLaunchUrlOk

`func (o *ApplicationRequest) GetMetaLaunchUrlOk() (*string, bool)`

GetMetaLaunchUrlOk returns a tuple with the MetaLaunchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaLaunchUrl

`func (o *ApplicationRequest) SetMetaLaunchUrl(v string)`

SetMetaLaunchUrl sets MetaLaunchUrl field to given value.

### HasMetaLaunchUrl

`func (o *ApplicationRequest) HasMetaLaunchUrl() bool`

HasMetaLaunchUrl returns a boolean if a field has been set.

### GetMetaDescription

`func (o *ApplicationRequest) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *ApplicationRequest) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *ApplicationRequest) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *ApplicationRequest) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### GetMetaPublisher

`func (o *ApplicationRequest) GetMetaPublisher() string`

GetMetaPublisher returns the MetaPublisher field if non-nil, zero value otherwise.

### GetMetaPublisherOk

`func (o *ApplicationRequest) GetMetaPublisherOk() (*string, bool)`

GetMetaPublisherOk returns a tuple with the MetaPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaPublisher

`func (o *ApplicationRequest) SetMetaPublisher(v string)`

SetMetaPublisher sets MetaPublisher field to given value.

### HasMetaPublisher

`func (o *ApplicationRequest) HasMetaPublisher() bool`

HasMetaPublisher returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *ApplicationRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *ApplicationRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *ApplicationRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *ApplicationRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


