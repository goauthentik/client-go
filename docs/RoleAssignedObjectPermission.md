# RoleAssignedObjectPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RolePk** | **string** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**Permissions** | [**[]RoleObjectPermission**](RoleObjectPermission.md) |  | 

## Methods

### NewRoleAssignedObjectPermission

`func NewRoleAssignedObjectPermission(rolePk string, name string, permissions []RoleObjectPermission, ) *RoleAssignedObjectPermission`

NewRoleAssignedObjectPermission instantiates a new RoleAssignedObjectPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAssignedObjectPermissionWithDefaults

`func NewRoleAssignedObjectPermissionWithDefaults() *RoleAssignedObjectPermission`

NewRoleAssignedObjectPermissionWithDefaults instantiates a new RoleAssignedObjectPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRolePk

`func (o *RoleAssignedObjectPermission) GetRolePk() string`

GetRolePk returns the RolePk field if non-nil, zero value otherwise.

### GetRolePkOk

`func (o *RoleAssignedObjectPermission) GetRolePkOk() (*string, bool)`

GetRolePkOk returns a tuple with the RolePk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolePk

`func (o *RoleAssignedObjectPermission) SetRolePk(v string)`

SetRolePk sets RolePk field to given value.


### GetName

`func (o *RoleAssignedObjectPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleAssignedObjectPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleAssignedObjectPermission) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *RoleAssignedObjectPermission) GetPermissions() []RoleObjectPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RoleAssignedObjectPermission) GetPermissionsOk() (*[]RoleObjectPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RoleAssignedObjectPermission) SetPermissions(v []RoleObjectPermission)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


