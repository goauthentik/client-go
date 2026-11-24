# OutpostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | [**OutpostTypeEnum**](OutpostTypeEnum.md) |  | 
**Providers** | **[]int32** |  | 
**ServiceConnection** | Pointer to **NullableString** | Select Service-Connection authentik should use to manage this outpost. Leave empty if authentik should not handle the deployment. | [optional] 
**Config** | **map[string]interface{}** |  | 
**Managed** | Pointer to **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 

## Methods

### NewOutpostRequest

`func NewOutpostRequest(name string, type_ OutpostTypeEnum, providers []int32, config map[string]interface{}, ) *OutpostRequest`

NewOutpostRequest instantiates a new OutpostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutpostRequestWithDefaults

`func NewOutpostRequestWithDefaults() *OutpostRequest`

NewOutpostRequestWithDefaults instantiates a new OutpostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OutpostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutpostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutpostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *OutpostRequest) GetType() OutpostTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutpostRequest) GetTypeOk() (*OutpostTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutpostRequest) SetType(v OutpostTypeEnum)`

SetType sets Type field to given value.


### GetProviders

`func (o *OutpostRequest) GetProviders() []int32`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *OutpostRequest) GetProvidersOk() (*[]int32, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *OutpostRequest) SetProviders(v []int32)`

SetProviders sets Providers field to given value.


### GetServiceConnection

`func (o *OutpostRequest) GetServiceConnection() string`

GetServiceConnection returns the ServiceConnection field if non-nil, zero value otherwise.

### GetServiceConnectionOk

`func (o *OutpostRequest) GetServiceConnectionOk() (*string, bool)`

GetServiceConnectionOk returns a tuple with the ServiceConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConnection

`func (o *OutpostRequest) SetServiceConnection(v string)`

SetServiceConnection sets ServiceConnection field to given value.

### HasServiceConnection

`func (o *OutpostRequest) HasServiceConnection() bool`

HasServiceConnection returns a boolean if a field has been set.

### SetServiceConnectionNil

`func (o *OutpostRequest) SetServiceConnectionNil(b bool)`

 SetServiceConnectionNil sets the value for ServiceConnection to be an explicit nil

### UnsetServiceConnection
`func (o *OutpostRequest) UnsetServiceConnection()`

UnsetServiceConnection ensures that no value is present for ServiceConnection, not even an explicit nil
### GetConfig

`func (o *OutpostRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *OutpostRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *OutpostRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetManaged

`func (o *OutpostRequest) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *OutpostRequest) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *OutpostRequest) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *OutpostRequest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *OutpostRequest) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *OutpostRequest) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


