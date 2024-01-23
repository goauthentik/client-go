# TenantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaName** | **string** |  | 
**Name** | **string** |  | 
**Ready** | Pointer to **bool** |  | [optional] 

## Methods

### NewTenantRequest

`func NewTenantRequest(schemaName string, name string, ) *TenantRequest`

NewTenantRequest instantiates a new TenantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantRequestWithDefaults

`func NewTenantRequestWithDefaults() *TenantRequest`

NewTenantRequestWithDefaults instantiates a new TenantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaName

`func (o *TenantRequest) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *TenantRequest) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *TenantRequest) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.


### GetName

`func (o *TenantRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantRequest) SetName(v string)`

SetName sets Name field to given value.


### GetReady

`func (o *TenantRequest) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *TenantRequest) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *TenantRequest) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *TenantRequest) HasReady() bool`

HasReady returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


