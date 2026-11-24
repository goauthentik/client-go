# ApplicationEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PbmUuid** | **string** |  | [readonly] 
**Name** | **string** |  | 
**App** | **string** |  | 
**AppName** | **string** | Application&#39;s display Name. | [readonly] 
**AppSlug** | **string** | Internal application name, used in URLs. | [readonly] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApplicationEntitlement

`func NewApplicationEntitlement(pbmUuid string, name string, app string, appName string, appSlug string, ) *ApplicationEntitlement`

NewApplicationEntitlement instantiates a new ApplicationEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationEntitlementWithDefaults

`func NewApplicationEntitlementWithDefaults() *ApplicationEntitlement`

NewApplicationEntitlementWithDefaults instantiates a new ApplicationEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPbmUuid

`func (o *ApplicationEntitlement) GetPbmUuid() string`

GetPbmUuid returns the PbmUuid field if non-nil, zero value otherwise.

### GetPbmUuidOk

`func (o *ApplicationEntitlement) GetPbmUuidOk() (*string, bool)`

GetPbmUuidOk returns a tuple with the PbmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbmUuid

`func (o *ApplicationEntitlement) SetPbmUuid(v string)`

SetPbmUuid sets PbmUuid field to given value.


### GetName

`func (o *ApplicationEntitlement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationEntitlement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationEntitlement) SetName(v string)`

SetName sets Name field to given value.


### GetApp

`func (o *ApplicationEntitlement) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *ApplicationEntitlement) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *ApplicationEntitlement) SetApp(v string)`

SetApp sets App field to given value.


### GetAppName

`func (o *ApplicationEntitlement) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ApplicationEntitlement) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ApplicationEntitlement) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetAppSlug

`func (o *ApplicationEntitlement) GetAppSlug() string`

GetAppSlug returns the AppSlug field if non-nil, zero value otherwise.

### GetAppSlugOk

`func (o *ApplicationEntitlement) GetAppSlugOk() (*string, bool)`

GetAppSlugOk returns a tuple with the AppSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSlug

`func (o *ApplicationEntitlement) SetAppSlug(v string)`

SetAppSlug sets AppSlug field to given value.


### GetAttributes

`func (o *ApplicationEntitlement) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationEntitlement) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationEntitlement) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ApplicationEntitlement) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


