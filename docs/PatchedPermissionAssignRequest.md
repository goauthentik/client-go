# PatchedPermissionAssignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to **[]string** |  | [optional] 
**Model** | Pointer to [**ModelEnum**](ModelEnum.md) |  | [optional] 
**ObjectPk** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedPermissionAssignRequest

`func NewPatchedPermissionAssignRequest() *PatchedPermissionAssignRequest`

NewPatchedPermissionAssignRequest instantiates a new PatchedPermissionAssignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedPermissionAssignRequestWithDefaults

`func NewPatchedPermissionAssignRequestWithDefaults() *PatchedPermissionAssignRequest`

NewPatchedPermissionAssignRequestWithDefaults instantiates a new PatchedPermissionAssignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *PatchedPermissionAssignRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PatchedPermissionAssignRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PatchedPermissionAssignRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PatchedPermissionAssignRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetModel

`func (o *PatchedPermissionAssignRequest) GetModel() ModelEnum`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PatchedPermissionAssignRequest) GetModelOk() (*ModelEnum, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PatchedPermissionAssignRequest) SetModel(v ModelEnum)`

SetModel sets Model field to given value.

### HasModel

`func (o *PatchedPermissionAssignRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetObjectPk

`func (o *PatchedPermissionAssignRequest) GetObjectPk() string`

GetObjectPk returns the ObjectPk field if non-nil, zero value otherwise.

### GetObjectPkOk

`func (o *PatchedPermissionAssignRequest) GetObjectPkOk() (*string, bool)`

GetObjectPkOk returns a tuple with the ObjectPk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectPk

`func (o *PatchedPermissionAssignRequest) SetObjectPk(v string)`

SetObjectPk sets ObjectPk field to given value.

### HasObjectPk

`func (o *PatchedPermissionAssignRequest) HasObjectPk() bool`

HasObjectPk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


