# InitialPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Mode** | [**InitialPermissionsModeEnum**](InitialPermissionsModeEnum.md) |  | 
**Role** | **string** |  | 
**Permissions** | Pointer to **[]int32** |  | [optional] 
**PermissionsObj** | [**[]Permission**](Permission.md) |  | [readonly] 

## Methods

### NewInitialPermissions

`func NewInitialPermissions(pk int32, name string, mode InitialPermissionsModeEnum, role string, permissionsObj []Permission, ) *InitialPermissions`

NewInitialPermissions instantiates a new InitialPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitialPermissionsWithDefaults

`func NewInitialPermissionsWithDefaults() *InitialPermissions`

NewInitialPermissionsWithDefaults instantiates a new InitialPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *InitialPermissions) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *InitialPermissions) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *InitialPermissions) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetName

`func (o *InitialPermissions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InitialPermissions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InitialPermissions) SetName(v string)`

SetName sets Name field to given value.


### GetMode

`func (o *InitialPermissions) GetMode() InitialPermissionsModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InitialPermissions) GetModeOk() (*InitialPermissionsModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InitialPermissions) SetMode(v InitialPermissionsModeEnum)`

SetMode sets Mode field to given value.


### GetRole

`func (o *InitialPermissions) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InitialPermissions) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InitialPermissions) SetRole(v string)`

SetRole sets Role field to given value.


### GetPermissions

`func (o *InitialPermissions) GetPermissions() []int32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InitialPermissions) GetPermissionsOk() (*[]int32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InitialPermissions) SetPermissions(v []int32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *InitialPermissions) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPermissionsObj

`func (o *InitialPermissions) GetPermissionsObj() []Permission`

GetPermissionsObj returns the PermissionsObj field if non-nil, zero value otherwise.

### GetPermissionsObjOk

`func (o *InitialPermissions) GetPermissionsObjOk() (*[]Permission, bool)`

GetPermissionsObjOk returns a tuple with the PermissionsObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionsObj

`func (o *InitialPermissions) SetPermissionsObj(v []Permission)`

SetPermissionsObj sets PermissionsObj field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


