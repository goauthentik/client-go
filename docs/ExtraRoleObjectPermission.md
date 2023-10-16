# ExtraRoleObjectPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Codename** | **string** |  | [readonly] 
**Model** | **string** |  | [readonly] 
**AppLabel** | **string** |  | [readonly] 
**ObjectPk** | **string** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**AppLabelVerbose** | **string** | Get app label from permission&#39;s model | [readonly] 
**ModelVerbose** | **string** | Get model label from permission&#39;s model | [readonly] 
**ObjectDescription** | **NullableString** | Get model description from attached model. This operation takes at least one additional query, and the description is only shown if the user/role has the view_ permission on the object | [readonly] 

## Methods

### NewExtraRoleObjectPermission

`func NewExtraRoleObjectPermission(id int32, codename string, model string, appLabel string, objectPk string, name string, appLabelVerbose string, modelVerbose string, objectDescription NullableString, ) *ExtraRoleObjectPermission`

NewExtraRoleObjectPermission instantiates a new ExtraRoleObjectPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtraRoleObjectPermissionWithDefaults

`func NewExtraRoleObjectPermissionWithDefaults() *ExtraRoleObjectPermission`

NewExtraRoleObjectPermissionWithDefaults instantiates a new ExtraRoleObjectPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtraRoleObjectPermission) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtraRoleObjectPermission) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtraRoleObjectPermission) SetId(v int32)`

SetId sets Id field to given value.


### GetCodename

`func (o *ExtraRoleObjectPermission) GetCodename() string`

GetCodename returns the Codename field if non-nil, zero value otherwise.

### GetCodenameOk

`func (o *ExtraRoleObjectPermission) GetCodenameOk() (*string, bool)`

GetCodenameOk returns a tuple with the Codename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodename

`func (o *ExtraRoleObjectPermission) SetCodename(v string)`

SetCodename sets Codename field to given value.


### GetModel

`func (o *ExtraRoleObjectPermission) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ExtraRoleObjectPermission) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ExtraRoleObjectPermission) SetModel(v string)`

SetModel sets Model field to given value.


### GetAppLabel

`func (o *ExtraRoleObjectPermission) GetAppLabel() string`

GetAppLabel returns the AppLabel field if non-nil, zero value otherwise.

### GetAppLabelOk

`func (o *ExtraRoleObjectPermission) GetAppLabelOk() (*string, bool)`

GetAppLabelOk returns a tuple with the AppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLabel

`func (o *ExtraRoleObjectPermission) SetAppLabel(v string)`

SetAppLabel sets AppLabel field to given value.


### GetObjectPk

`func (o *ExtraRoleObjectPermission) GetObjectPk() string`

GetObjectPk returns the ObjectPk field if non-nil, zero value otherwise.

### GetObjectPkOk

`func (o *ExtraRoleObjectPermission) GetObjectPkOk() (*string, bool)`

GetObjectPkOk returns a tuple with the ObjectPk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectPk

`func (o *ExtraRoleObjectPermission) SetObjectPk(v string)`

SetObjectPk sets ObjectPk field to given value.


### GetName

`func (o *ExtraRoleObjectPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtraRoleObjectPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtraRoleObjectPermission) SetName(v string)`

SetName sets Name field to given value.


### GetAppLabelVerbose

`func (o *ExtraRoleObjectPermission) GetAppLabelVerbose() string`

GetAppLabelVerbose returns the AppLabelVerbose field if non-nil, zero value otherwise.

### GetAppLabelVerboseOk

`func (o *ExtraRoleObjectPermission) GetAppLabelVerboseOk() (*string, bool)`

GetAppLabelVerboseOk returns a tuple with the AppLabelVerbose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLabelVerbose

`func (o *ExtraRoleObjectPermission) SetAppLabelVerbose(v string)`

SetAppLabelVerbose sets AppLabelVerbose field to given value.


### GetModelVerbose

`func (o *ExtraRoleObjectPermission) GetModelVerbose() string`

GetModelVerbose returns the ModelVerbose field if non-nil, zero value otherwise.

### GetModelVerboseOk

`func (o *ExtraRoleObjectPermission) GetModelVerboseOk() (*string, bool)`

GetModelVerboseOk returns a tuple with the ModelVerbose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVerbose

`func (o *ExtraRoleObjectPermission) SetModelVerbose(v string)`

SetModelVerbose sets ModelVerbose field to given value.


### GetObjectDescription

`func (o *ExtraRoleObjectPermission) GetObjectDescription() string`

GetObjectDescription returns the ObjectDescription field if non-nil, zero value otherwise.

### GetObjectDescriptionOk

`func (o *ExtraRoleObjectPermission) GetObjectDescriptionOk() (*string, bool)`

GetObjectDescriptionOk returns a tuple with the ObjectDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDescription

`func (o *ExtraRoleObjectPermission) SetObjectDescription(v string)`

SetObjectDescription sets ObjectDescription field to given value.


### SetObjectDescriptionNil

`func (o *ExtraRoleObjectPermission) SetObjectDescriptionNil(b bool)`

 SetObjectDescriptionNil sets the value for ObjectDescription to be an explicit nil

### UnsetObjectDescription
`func (o *ExtraRoleObjectPermission) UnsetObjectDescription()`

UnsetObjectDescription ensures that no value is present for ObjectDescription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


