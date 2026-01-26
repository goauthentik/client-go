# ApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Application&#39;s display Name. | 
**Slug** | **string** | Internal application name, used in URLs. | 
**Provider** | Pointer to **NullableInt32** |  | [optional] 
**BackchannelProviders** | Pointer to **[]int32** |  | [optional] 
**OpenInNewTab** | Pointer to **bool** | Open launch URL in a new browser tab or window. | [optional] 
**MetaLaunchUrl** | Pointer to **string** |  | [optional] 
**MetaDescription** | Pointer to **string** |  | [optional] 
**MetaPublisher** | Pointer to **string** |  | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 

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
### GetBackchannelProviders

`func (o *ApplicationRequest) GetBackchannelProviders() []int32`

GetBackchannelProviders returns the BackchannelProviders field if non-nil, zero value otherwise.

### GetBackchannelProvidersOk

`func (o *ApplicationRequest) GetBackchannelProvidersOk() (*[]int32, bool)`

GetBackchannelProvidersOk returns a tuple with the BackchannelProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelProviders

`func (o *ApplicationRequest) SetBackchannelProviders(v []int32)`

SetBackchannelProviders sets BackchannelProviders field to given value.

### HasBackchannelProviders

`func (o *ApplicationRequest) HasBackchannelProviders() bool`

HasBackchannelProviders returns a boolean if a field has been set.

### GetOpenInNewTab

`func (o *ApplicationRequest) GetOpenInNewTab() bool`

GetOpenInNewTab returns the OpenInNewTab field if non-nil, zero value otherwise.

### GetOpenInNewTabOk

`func (o *ApplicationRequest) GetOpenInNewTabOk() (*bool, bool)`

GetOpenInNewTabOk returns a tuple with the OpenInNewTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInNewTab

`func (o *ApplicationRequest) SetOpenInNewTab(v bool)`

SetOpenInNewTab sets OpenInNewTab field to given value.

### HasOpenInNewTab

`func (o *ApplicationRequest) HasOpenInNewTab() bool`

HasOpenInNewTab returns a boolean if a field has been set.

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

### GetGroup

`func (o *ApplicationRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ApplicationRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ApplicationRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ApplicationRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


