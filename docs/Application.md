# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** | Application&#39;s display Name. | 
**Slug** | **string** | Internal application name, used in URLs. | 
**Provider** | Pointer to **NullableInt32** |  | [optional] 
**ProviderObj** | [**Provider**](Provider.md) |  | [readonly] 
**BackchannelProviders** | Pointer to **[]int32** |  | [optional] 
**BackchannelProvidersObj** | [**[]Provider**](Provider.md) |  | [readonly] 
**LaunchUrl** | **NullableString** | Allow formatting of launch URL | [readonly] 
**OpenInNewTab** | Pointer to **bool** | Open launch URL in a new browser tab or window. | [optional] 
**MetaLaunchUrl** | Pointer to **string** |  | [optional] 
**MetaIcon** | **NullableString** | Get the URL to the App Icon image. If the name is /static or starts with http it is returned as-is | [readonly] 
**MetaDescription** | Pointer to **string** |  | [optional] 
**MetaPublisher** | Pointer to **string** |  | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 

## Methods

### NewApplication

`func NewApplication(pk string, name string, slug string, providerObj Provider, backchannelProvidersObj []Provider, launchUrl NullableString, metaIcon NullableString, ) *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *Application) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *Application) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *Application) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *Application) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Application) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Application) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *Application) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Application) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Application) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetProvider

`func (o *Application) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Application) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Application) SetProvider(v int32)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Application) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *Application) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *Application) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetProviderObj

`func (o *Application) GetProviderObj() Provider`

GetProviderObj returns the ProviderObj field if non-nil, zero value otherwise.

### GetProviderObjOk

`func (o *Application) GetProviderObjOk() (*Provider, bool)`

GetProviderObjOk returns a tuple with the ProviderObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderObj

`func (o *Application) SetProviderObj(v Provider)`

SetProviderObj sets ProviderObj field to given value.


### GetBackchannelProviders

`func (o *Application) GetBackchannelProviders() []int32`

GetBackchannelProviders returns the BackchannelProviders field if non-nil, zero value otherwise.

### GetBackchannelProvidersOk

`func (o *Application) GetBackchannelProvidersOk() (*[]int32, bool)`

GetBackchannelProvidersOk returns a tuple with the BackchannelProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelProviders

`func (o *Application) SetBackchannelProviders(v []int32)`

SetBackchannelProviders sets BackchannelProviders field to given value.

### HasBackchannelProviders

`func (o *Application) HasBackchannelProviders() bool`

HasBackchannelProviders returns a boolean if a field has been set.

### GetBackchannelProvidersObj

`func (o *Application) GetBackchannelProvidersObj() []Provider`

GetBackchannelProvidersObj returns the BackchannelProvidersObj field if non-nil, zero value otherwise.

### GetBackchannelProvidersObjOk

`func (o *Application) GetBackchannelProvidersObjOk() (*[]Provider, bool)`

GetBackchannelProvidersObjOk returns a tuple with the BackchannelProvidersObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelProvidersObj

`func (o *Application) SetBackchannelProvidersObj(v []Provider)`

SetBackchannelProvidersObj sets BackchannelProvidersObj field to given value.


### GetLaunchUrl

`func (o *Application) GetLaunchUrl() string`

GetLaunchUrl returns the LaunchUrl field if non-nil, zero value otherwise.

### GetLaunchUrlOk

`func (o *Application) GetLaunchUrlOk() (*string, bool)`

GetLaunchUrlOk returns a tuple with the LaunchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchUrl

`func (o *Application) SetLaunchUrl(v string)`

SetLaunchUrl sets LaunchUrl field to given value.


### SetLaunchUrlNil

`func (o *Application) SetLaunchUrlNil(b bool)`

 SetLaunchUrlNil sets the value for LaunchUrl to be an explicit nil

### UnsetLaunchUrl
`func (o *Application) UnsetLaunchUrl()`

UnsetLaunchUrl ensures that no value is present for LaunchUrl, not even an explicit nil
### GetOpenInNewTab

`func (o *Application) GetOpenInNewTab() bool`

GetOpenInNewTab returns the OpenInNewTab field if non-nil, zero value otherwise.

### GetOpenInNewTabOk

`func (o *Application) GetOpenInNewTabOk() (*bool, bool)`

GetOpenInNewTabOk returns a tuple with the OpenInNewTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInNewTab

`func (o *Application) SetOpenInNewTab(v bool)`

SetOpenInNewTab sets OpenInNewTab field to given value.

### HasOpenInNewTab

`func (o *Application) HasOpenInNewTab() bool`

HasOpenInNewTab returns a boolean if a field has been set.

### GetMetaLaunchUrl

`func (o *Application) GetMetaLaunchUrl() string`

GetMetaLaunchUrl returns the MetaLaunchUrl field if non-nil, zero value otherwise.

### GetMetaLaunchUrlOk

`func (o *Application) GetMetaLaunchUrlOk() (*string, bool)`

GetMetaLaunchUrlOk returns a tuple with the MetaLaunchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaLaunchUrl

`func (o *Application) SetMetaLaunchUrl(v string)`

SetMetaLaunchUrl sets MetaLaunchUrl field to given value.

### HasMetaLaunchUrl

`func (o *Application) HasMetaLaunchUrl() bool`

HasMetaLaunchUrl returns a boolean if a field has been set.

### GetMetaIcon

`func (o *Application) GetMetaIcon() string`

GetMetaIcon returns the MetaIcon field if non-nil, zero value otherwise.

### GetMetaIconOk

`func (o *Application) GetMetaIconOk() (*string, bool)`

GetMetaIconOk returns a tuple with the MetaIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaIcon

`func (o *Application) SetMetaIcon(v string)`

SetMetaIcon sets MetaIcon field to given value.


### SetMetaIconNil

`func (o *Application) SetMetaIconNil(b bool)`

 SetMetaIconNil sets the value for MetaIcon to be an explicit nil

### UnsetMetaIcon
`func (o *Application) UnsetMetaIcon()`

UnsetMetaIcon ensures that no value is present for MetaIcon, not even an explicit nil
### GetMetaDescription

`func (o *Application) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *Application) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *Application) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *Application) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### GetMetaPublisher

`func (o *Application) GetMetaPublisher() string`

GetMetaPublisher returns the MetaPublisher field if non-nil, zero value otherwise.

### GetMetaPublisherOk

`func (o *Application) GetMetaPublisherOk() (*string, bool)`

GetMetaPublisherOk returns a tuple with the MetaPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaPublisher

`func (o *Application) SetMetaPublisher(v string)`

SetMetaPublisher sets MetaPublisher field to given value.

### HasMetaPublisher

`func (o *Application) HasMetaPublisher() bool`

HasMetaPublisher returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *Application) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *Application) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *Application) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *Application) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetGroup

`func (o *Application) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Application) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Application) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Application) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


