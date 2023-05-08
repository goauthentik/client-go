# PatchedApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Application&#39;s display Name. | [optional] 
**Slug** | Pointer to **string** | Internal application name, used in URLs. | [optional] 
**Provider** | Pointer to **NullableInt32** |  | [optional] 
**BackchannelProviders** | Pointer to **[]int32** |  | [optional] 
**OpenInNewTab** | Pointer to **bool** | Open launch URL in a new browser tab or window. | [optional] 
**MetaLaunchUrl** | Pointer to **string** |  | [optional] 
**MetaDescription** | Pointer to **string** |  | [optional] 
**MetaPublisher** | Pointer to **string** |  | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedApplicationRequest

`func NewPatchedApplicationRequest() *PatchedApplicationRequest`

NewPatchedApplicationRequest instantiates a new PatchedApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedApplicationRequestWithDefaults

`func NewPatchedApplicationRequestWithDefaults() *PatchedApplicationRequest`

NewPatchedApplicationRequestWithDefaults instantiates a new PatchedApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedApplicationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedApplicationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedApplicationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedApplicationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedApplicationRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedApplicationRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedApplicationRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedApplicationRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetProvider

`func (o *PatchedApplicationRequest) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PatchedApplicationRequest) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PatchedApplicationRequest) SetProvider(v int32)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PatchedApplicationRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *PatchedApplicationRequest) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *PatchedApplicationRequest) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetBackchannelProviders

`func (o *PatchedApplicationRequest) GetBackchannelProviders() []int32`

GetBackchannelProviders returns the BackchannelProviders field if non-nil, zero value otherwise.

### GetBackchannelProvidersOk

`func (o *PatchedApplicationRequest) GetBackchannelProvidersOk() (*[]int32, bool)`

GetBackchannelProvidersOk returns a tuple with the BackchannelProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelProviders

`func (o *PatchedApplicationRequest) SetBackchannelProviders(v []int32)`

SetBackchannelProviders sets BackchannelProviders field to given value.

### HasBackchannelProviders

`func (o *PatchedApplicationRequest) HasBackchannelProviders() bool`

HasBackchannelProviders returns a boolean if a field has been set.

### GetOpenInNewTab

`func (o *PatchedApplicationRequest) GetOpenInNewTab() bool`

GetOpenInNewTab returns the OpenInNewTab field if non-nil, zero value otherwise.

### GetOpenInNewTabOk

`func (o *PatchedApplicationRequest) GetOpenInNewTabOk() (*bool, bool)`

GetOpenInNewTabOk returns a tuple with the OpenInNewTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInNewTab

`func (o *PatchedApplicationRequest) SetOpenInNewTab(v bool)`

SetOpenInNewTab sets OpenInNewTab field to given value.

### HasOpenInNewTab

`func (o *PatchedApplicationRequest) HasOpenInNewTab() bool`

HasOpenInNewTab returns a boolean if a field has been set.

### GetMetaLaunchUrl

`func (o *PatchedApplicationRequest) GetMetaLaunchUrl() string`

GetMetaLaunchUrl returns the MetaLaunchUrl field if non-nil, zero value otherwise.

### GetMetaLaunchUrlOk

`func (o *PatchedApplicationRequest) GetMetaLaunchUrlOk() (*string, bool)`

GetMetaLaunchUrlOk returns a tuple with the MetaLaunchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaLaunchUrl

`func (o *PatchedApplicationRequest) SetMetaLaunchUrl(v string)`

SetMetaLaunchUrl sets MetaLaunchUrl field to given value.

### HasMetaLaunchUrl

`func (o *PatchedApplicationRequest) HasMetaLaunchUrl() bool`

HasMetaLaunchUrl returns a boolean if a field has been set.

### GetMetaDescription

`func (o *PatchedApplicationRequest) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *PatchedApplicationRequest) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *PatchedApplicationRequest) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *PatchedApplicationRequest) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### GetMetaPublisher

`func (o *PatchedApplicationRequest) GetMetaPublisher() string`

GetMetaPublisher returns the MetaPublisher field if non-nil, zero value otherwise.

### GetMetaPublisherOk

`func (o *PatchedApplicationRequest) GetMetaPublisherOk() (*string, bool)`

GetMetaPublisherOk returns a tuple with the MetaPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaPublisher

`func (o *PatchedApplicationRequest) SetMetaPublisher(v string)`

SetMetaPublisher sets MetaPublisher field to given value.

### HasMetaPublisher

`func (o *PatchedApplicationRequest) HasMetaPublisher() bool`

HasMetaPublisher returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *PatchedApplicationRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *PatchedApplicationRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *PatchedApplicationRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *PatchedApplicationRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetGroup

`func (o *PatchedApplicationRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedApplicationRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedApplicationRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedApplicationRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


