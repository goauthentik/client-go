# CustomPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Codename** | **string** |  | 
**Name** | **string** |  | 
**Model** | **string** |  | 
**Permission** | **string** |  | [readonly] 

## Methods

### NewCustomPermission

`func NewCustomPermission(pk int32, codename string, name string, model string, permission string, ) *CustomPermission`

NewCustomPermission instantiates a new CustomPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomPermissionWithDefaults

`func NewCustomPermissionWithDefaults() *CustomPermission`

NewCustomPermissionWithDefaults instantiates a new CustomPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *CustomPermission) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *CustomPermission) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *CustomPermission) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetCodename

`func (o *CustomPermission) GetCodename() string`

GetCodename returns the Codename field if non-nil, zero value otherwise.

### GetCodenameOk

`func (o *CustomPermission) GetCodenameOk() (*string, bool)`

GetCodenameOk returns a tuple with the Codename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodename

`func (o *CustomPermission) SetCodename(v string)`

SetCodename sets Codename field to given value.


### GetName

`func (o *CustomPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomPermission) SetName(v string)`

SetName sets Name field to given value.


### GetModel

`func (o *CustomPermission) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CustomPermission) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CustomPermission) SetModel(v string)`

SetModel sets Model field to given value.


### GetPermission

`func (o *CustomPermission) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *CustomPermission) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *CustomPermission) SetPermission(v string)`

SetPermission sets Permission field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


