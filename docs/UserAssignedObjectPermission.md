# UserAssignedObjectPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Username** | **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | 
**Name** | **string** | User&#39;s display name. | 
**IsActive** | Pointer to **bool** | Designates whether this user should be treated as active. Unselect this instead of deleting accounts. | [optional] 
**LastLogin** | Pointer to **NullableTime** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Uid** | **string** |  | [readonly] 
**Permissions** | [**[]UserObjectPermission**](UserObjectPermission.md) |  | 
**IsSuperuser** | **bool** |  | 

## Methods

### NewUserAssignedObjectPermission

`func NewUserAssignedObjectPermission(pk int32, username string, name string, uid string, permissions []UserObjectPermission, isSuperuser bool, ) *UserAssignedObjectPermission`

NewUserAssignedObjectPermission instantiates a new UserAssignedObjectPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAssignedObjectPermissionWithDefaults

`func NewUserAssignedObjectPermissionWithDefaults() *UserAssignedObjectPermission`

NewUserAssignedObjectPermissionWithDefaults instantiates a new UserAssignedObjectPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *UserAssignedObjectPermission) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *UserAssignedObjectPermission) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *UserAssignedObjectPermission) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetUsername

`func (o *UserAssignedObjectPermission) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserAssignedObjectPermission) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserAssignedObjectPermission) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetName

`func (o *UserAssignedObjectPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserAssignedObjectPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserAssignedObjectPermission) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *UserAssignedObjectPermission) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UserAssignedObjectPermission) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UserAssignedObjectPermission) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UserAssignedObjectPermission) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastLogin

`func (o *UserAssignedObjectPermission) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *UserAssignedObjectPermission) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *UserAssignedObjectPermission) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *UserAssignedObjectPermission) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *UserAssignedObjectPermission) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *UserAssignedObjectPermission) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetEmail

`func (o *UserAssignedObjectPermission) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserAssignedObjectPermission) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserAssignedObjectPermission) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserAssignedObjectPermission) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAttributes

`func (o *UserAssignedObjectPermission) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserAssignedObjectPermission) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserAssignedObjectPermission) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UserAssignedObjectPermission) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetUid

`func (o *UserAssignedObjectPermission) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UserAssignedObjectPermission) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UserAssignedObjectPermission) SetUid(v string)`

SetUid sets Uid field to given value.


### GetPermissions

`func (o *UserAssignedObjectPermission) GetPermissions() []UserObjectPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserAssignedObjectPermission) GetPermissionsOk() (*[]UserObjectPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserAssignedObjectPermission) SetPermissions(v []UserObjectPermission)`

SetPermissions sets Permissions field to given value.


### GetIsSuperuser

`func (o *UserAssignedObjectPermission) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *UserAssignedObjectPermission) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *UserAssignedObjectPermission) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


