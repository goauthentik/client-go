# PatchedInitialPermissionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**InitialPermissionsModeEnum**](InitialPermissionsModeEnum.md) |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewPatchedInitialPermissionsRequest

`func NewPatchedInitialPermissionsRequest() *PatchedInitialPermissionsRequest`

NewPatchedInitialPermissionsRequest instantiates a new PatchedInitialPermissionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedInitialPermissionsRequestWithDefaults

`func NewPatchedInitialPermissionsRequestWithDefaults() *PatchedInitialPermissionsRequest`

NewPatchedInitialPermissionsRequestWithDefaults instantiates a new PatchedInitialPermissionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedInitialPermissionsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedInitialPermissionsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedInitialPermissionsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedInitialPermissionsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMode

`func (o *PatchedInitialPermissionsRequest) GetMode() InitialPermissionsModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedInitialPermissionsRequest) GetModeOk() (*InitialPermissionsModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedInitialPermissionsRequest) SetMode(v InitialPermissionsModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedInitialPermissionsRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetRole

`func (o *PatchedInitialPermissionsRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedInitialPermissionsRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedInitialPermissionsRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedInitialPermissionsRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetPermissions

`func (o *PatchedInitialPermissionsRequest) GetPermissions() []int32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PatchedInitialPermissionsRequest) GetPermissionsOk() (*[]int32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PatchedInitialPermissionsRequest) SetPermissions(v []int32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PatchedInitialPermissionsRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


