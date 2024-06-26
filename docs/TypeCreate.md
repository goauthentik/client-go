# TypeCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Component** | **string** |  | 
**ModelName** | **string** |  | 
**IconUrl** | Pointer to **string** |  | [optional] 
**RequiresEnterprise** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewTypeCreate

`func NewTypeCreate(name string, description string, component string, modelName string, ) *TypeCreate`

NewTypeCreate instantiates a new TypeCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypeCreateWithDefaults

`func NewTypeCreateWithDefaults() *TypeCreate`

NewTypeCreateWithDefaults instantiates a new TypeCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TypeCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TypeCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TypeCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TypeCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TypeCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TypeCreate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetComponent

`func (o *TypeCreate) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *TypeCreate) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *TypeCreate) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetModelName

`func (o *TypeCreate) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *TypeCreate) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *TypeCreate) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetIconUrl

`func (o *TypeCreate) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *TypeCreate) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *TypeCreate) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *TypeCreate) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetRequiresEnterprise

`func (o *TypeCreate) GetRequiresEnterprise() bool`

GetRequiresEnterprise returns the RequiresEnterprise field if non-nil, zero value otherwise.

### GetRequiresEnterpriseOk

`func (o *TypeCreate) GetRequiresEnterpriseOk() (*bool, bool)`

GetRequiresEnterpriseOk returns a tuple with the RequiresEnterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresEnterprise

`func (o *TypeCreate) SetRequiresEnterprise(v bool)`

SetRequiresEnterprise sets RequiresEnterprise field to given value.

### HasRequiresEnterprise

`func (o *TypeCreate) HasRequiresEnterprise() bool`

HasRequiresEnterprise returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


