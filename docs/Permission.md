# Permission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Codename** | **string** |  | 
**Model** | **string** |  | [readonly] 
**AppLabel** | **string** |  | [readonly] 
**AppLabelVerbose** | **string** | Human-readable app label | [readonly] 
**ModelVerbose** | **string** | Human-readable model name | [readonly] 

## Methods

### NewPermission

`func NewPermission(id int32, name string, codename string, model string, appLabel string, appLabelVerbose string, modelVerbose string, ) *Permission`

NewPermission instantiates a new Permission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionWithDefaults

`func NewPermissionWithDefaults() *Permission`

NewPermissionWithDefaults instantiates a new Permission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Permission) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Permission) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Permission) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Permission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Permission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Permission) SetName(v string)`

SetName sets Name field to given value.


### GetCodename

`func (o *Permission) GetCodename() string`

GetCodename returns the Codename field if non-nil, zero value otherwise.

### GetCodenameOk

`func (o *Permission) GetCodenameOk() (*string, bool)`

GetCodenameOk returns a tuple with the Codename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodename

`func (o *Permission) SetCodename(v string)`

SetCodename sets Codename field to given value.


### GetModel

`func (o *Permission) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Permission) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Permission) SetModel(v string)`

SetModel sets Model field to given value.


### GetAppLabel

`func (o *Permission) GetAppLabel() string`

GetAppLabel returns the AppLabel field if non-nil, zero value otherwise.

### GetAppLabelOk

`func (o *Permission) GetAppLabelOk() (*string, bool)`

GetAppLabelOk returns a tuple with the AppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLabel

`func (o *Permission) SetAppLabel(v string)`

SetAppLabel sets AppLabel field to given value.


### GetAppLabelVerbose

`func (o *Permission) GetAppLabelVerbose() string`

GetAppLabelVerbose returns the AppLabelVerbose field if non-nil, zero value otherwise.

### GetAppLabelVerboseOk

`func (o *Permission) GetAppLabelVerboseOk() (*string, bool)`

GetAppLabelVerboseOk returns a tuple with the AppLabelVerbose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLabelVerbose

`func (o *Permission) SetAppLabelVerbose(v string)`

SetAppLabelVerbose sets AppLabelVerbose field to given value.


### GetModelVerbose

`func (o *Permission) GetModelVerbose() string`

GetModelVerbose returns the ModelVerbose field if non-nil, zero value otherwise.

### GetModelVerboseOk

`func (o *Permission) GetModelVerboseOk() (*string, bool)`

GetModelVerboseOk returns a tuple with the ModelVerbose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVerbose

`func (o *Permission) SetModelVerbose(v string)`

SetModelVerbose sets ModelVerbose field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


