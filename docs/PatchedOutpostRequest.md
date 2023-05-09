# PatchedOutpostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**OutpostTypeEnum**](OutpostTypeEnum.md) |  | [optional] 
**Providers** | Pointer to **[]int32** |  | [optional] 
**ServiceConnection** | Pointer to **NullableString** | Select Service-Connection authentik should use to manage this outpost. Leave empty if authentik should not handle the deployment. | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Managed** | Pointer to **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 

## Methods

### NewPatchedOutpostRequest

`func NewPatchedOutpostRequest() *PatchedOutpostRequest`

NewPatchedOutpostRequest instantiates a new PatchedOutpostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedOutpostRequestWithDefaults

`func NewPatchedOutpostRequestWithDefaults() *PatchedOutpostRequest`

NewPatchedOutpostRequestWithDefaults instantiates a new PatchedOutpostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedOutpostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedOutpostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedOutpostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedOutpostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PatchedOutpostRequest) GetType() OutpostTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedOutpostRequest) GetTypeOk() (*OutpostTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedOutpostRequest) SetType(v OutpostTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedOutpostRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProviders

`func (o *PatchedOutpostRequest) GetProviders() []int32`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *PatchedOutpostRequest) GetProvidersOk() (*[]int32, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *PatchedOutpostRequest) SetProviders(v []int32)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *PatchedOutpostRequest) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetServiceConnection

`func (o *PatchedOutpostRequest) GetServiceConnection() string`

GetServiceConnection returns the ServiceConnection field if non-nil, zero value otherwise.

### GetServiceConnectionOk

`func (o *PatchedOutpostRequest) GetServiceConnectionOk() (*string, bool)`

GetServiceConnectionOk returns a tuple with the ServiceConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConnection

`func (o *PatchedOutpostRequest) SetServiceConnection(v string)`

SetServiceConnection sets ServiceConnection field to given value.

### HasServiceConnection

`func (o *PatchedOutpostRequest) HasServiceConnection() bool`

HasServiceConnection returns a boolean if a field has been set.

### SetServiceConnectionNil

`func (o *PatchedOutpostRequest) SetServiceConnectionNil(b bool)`

 SetServiceConnectionNil sets the value for ServiceConnection to be an explicit nil

### UnsetServiceConnection
`func (o *PatchedOutpostRequest) UnsetServiceConnection()`

UnsetServiceConnection ensures that no value is present for ServiceConnection, not even an explicit nil
### GetConfig

`func (o *PatchedOutpostRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PatchedOutpostRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PatchedOutpostRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *PatchedOutpostRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetManaged

`func (o *PatchedOutpostRequest) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *PatchedOutpostRequest) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *PatchedOutpostRequest) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *PatchedOutpostRequest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *PatchedOutpostRequest) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *PatchedOutpostRequest) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


