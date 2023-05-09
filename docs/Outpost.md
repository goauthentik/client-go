# Outpost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Type** | [**OutpostTypeEnum**](OutpostTypeEnum.md) |  | 
**Providers** | **[]int32** |  | 
**ProvidersObj** | [**[]Provider**](Provider.md) |  | [readonly] 
**ServiceConnection** | Pointer to **NullableString** | Select Service-Connection authentik should use to manage this outpost. Leave empty if authentik should not handle the deployment. | [optional] 
**ServiceConnectionObj** | [**ServiceConnection**](ServiceConnection.md) |  | [readonly] 
**TokenIdentifier** | **string** | Get Token identifier | [readonly] 
**Config** | **map[string]interface{}** |  | 
**Managed** | Pointer to **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 

## Methods

### NewOutpost

`func NewOutpost(pk string, name string, type_ OutpostTypeEnum, providers []int32, providersObj []Provider, serviceConnectionObj ServiceConnection, tokenIdentifier string, config map[string]interface{}, ) *Outpost`

NewOutpost instantiates a new Outpost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutpostWithDefaults

`func NewOutpostWithDefaults() *Outpost`

NewOutpostWithDefaults instantiates a new Outpost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *Outpost) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *Outpost) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *Outpost) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *Outpost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Outpost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Outpost) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Outpost) GetType() OutpostTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Outpost) GetTypeOk() (*OutpostTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Outpost) SetType(v OutpostTypeEnum)`

SetType sets Type field to given value.


### GetProviders

`func (o *Outpost) GetProviders() []int32`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *Outpost) GetProvidersOk() (*[]int32, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *Outpost) SetProviders(v []int32)`

SetProviders sets Providers field to given value.


### GetProvidersObj

`func (o *Outpost) GetProvidersObj() []Provider`

GetProvidersObj returns the ProvidersObj field if non-nil, zero value otherwise.

### GetProvidersObjOk

`func (o *Outpost) GetProvidersObjOk() (*[]Provider, bool)`

GetProvidersObjOk returns a tuple with the ProvidersObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidersObj

`func (o *Outpost) SetProvidersObj(v []Provider)`

SetProvidersObj sets ProvidersObj field to given value.


### GetServiceConnection

`func (o *Outpost) GetServiceConnection() string`

GetServiceConnection returns the ServiceConnection field if non-nil, zero value otherwise.

### GetServiceConnectionOk

`func (o *Outpost) GetServiceConnectionOk() (*string, bool)`

GetServiceConnectionOk returns a tuple with the ServiceConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConnection

`func (o *Outpost) SetServiceConnection(v string)`

SetServiceConnection sets ServiceConnection field to given value.

### HasServiceConnection

`func (o *Outpost) HasServiceConnection() bool`

HasServiceConnection returns a boolean if a field has been set.

### SetServiceConnectionNil

`func (o *Outpost) SetServiceConnectionNil(b bool)`

 SetServiceConnectionNil sets the value for ServiceConnection to be an explicit nil

### UnsetServiceConnection
`func (o *Outpost) UnsetServiceConnection()`

UnsetServiceConnection ensures that no value is present for ServiceConnection, not even an explicit nil
### GetServiceConnectionObj

`func (o *Outpost) GetServiceConnectionObj() ServiceConnection`

GetServiceConnectionObj returns the ServiceConnectionObj field if non-nil, zero value otherwise.

### GetServiceConnectionObjOk

`func (o *Outpost) GetServiceConnectionObjOk() (*ServiceConnection, bool)`

GetServiceConnectionObjOk returns a tuple with the ServiceConnectionObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConnectionObj

`func (o *Outpost) SetServiceConnectionObj(v ServiceConnection)`

SetServiceConnectionObj sets ServiceConnectionObj field to given value.


### GetTokenIdentifier

`func (o *Outpost) GetTokenIdentifier() string`

GetTokenIdentifier returns the TokenIdentifier field if non-nil, zero value otherwise.

### GetTokenIdentifierOk

`func (o *Outpost) GetTokenIdentifierOk() (*string, bool)`

GetTokenIdentifierOk returns a tuple with the TokenIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIdentifier

`func (o *Outpost) SetTokenIdentifier(v string)`

SetTokenIdentifier sets TokenIdentifier field to given value.


### GetConfig

`func (o *Outpost) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Outpost) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Outpost) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetManaged

`func (o *Outpost) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *Outpost) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *Outpost) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *Outpost) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *Outpost) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *Outpost) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


