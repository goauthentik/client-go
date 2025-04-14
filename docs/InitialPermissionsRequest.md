# InitialPermissionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Mode** | [**InitialPermissionsModeEnum**](InitialPermissionsModeEnum.md) |  | 
**Role** | **string** |  | 
**Permissions** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewInitialPermissionsRequest

`func NewInitialPermissionsRequest(name string, mode InitialPermissionsModeEnum, role string, ) *InitialPermissionsRequest`

NewInitialPermissionsRequest instantiates a new InitialPermissionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitialPermissionsRequestWithDefaults

`func NewInitialPermissionsRequestWithDefaults() *InitialPermissionsRequest`

NewInitialPermissionsRequestWithDefaults instantiates a new InitialPermissionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InitialPermissionsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InitialPermissionsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InitialPermissionsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMode

`func (o *InitialPermissionsRequest) GetMode() InitialPermissionsModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InitialPermissionsRequest) GetModeOk() (*InitialPermissionsModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InitialPermissionsRequest) SetMode(v InitialPermissionsModeEnum)`

SetMode sets Mode field to given value.


### GetRole

`func (o *InitialPermissionsRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InitialPermissionsRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InitialPermissionsRequest) SetRole(v string)`

SetRole sets Role field to given value.


### GetPermissions

`func (o *InitialPermissionsRequest) GetPermissions() []int32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InitialPermissionsRequest) GetPermissionsOk() (*[]int32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InitialPermissionsRequest) SetPermissions(v []int32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *InitialPermissionsRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


