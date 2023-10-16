# UserSelf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Username** | **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | 
**Name** | **string** | User&#39;s display name. | 
**IsActive** | **bool** | Designates whether this user should be treated as active. Unselect this instead of deleting accounts. | [readonly] 
**IsSuperuser** | **bool** |  | [readonly] 
**Groups** | [**[]UserSelfGroups**](UserSelfGroups.md) |  | [readonly] 
**Email** | Pointer to **string** |  | [optional] 
**Avatar** | **string** |  | [readonly] 
**Uid** | **string** |  | [readonly] 
**Settings** | **map[string]interface{}** | Get user settings with tenant and group settings applied | [readonly] 
**Type** | Pointer to [**UserTypeEnum**](UserTypeEnum.md) |  | [optional] 
**SystemPermissions** | **[]string** | Get all system permissions assigned to the user | [readonly] 

## Methods

### NewUserSelf

`func NewUserSelf(pk int32, username string, name string, isActive bool, isSuperuser bool, groups []UserSelfGroups, avatar string, uid string, settings map[string]interface{}, systemPermissions []string, ) *UserSelf`

NewUserSelf instantiates a new UserSelf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSelfWithDefaults

`func NewUserSelfWithDefaults() *UserSelf`

NewUserSelfWithDefaults instantiates a new UserSelf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *UserSelf) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *UserSelf) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *UserSelf) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetUsername

`func (o *UserSelf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserSelf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserSelf) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetName

`func (o *UserSelf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSelf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSelf) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *UserSelf) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UserSelf) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UserSelf) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetIsSuperuser

`func (o *UserSelf) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *UserSelf) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *UserSelf) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.


### GetGroups

`func (o *UserSelf) GetGroups() []UserSelfGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserSelf) GetGroupsOk() (*[]UserSelfGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserSelf) SetGroups(v []UserSelfGroups)`

SetGroups sets Groups field to given value.


### GetEmail

`func (o *UserSelf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSelf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSelf) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserSelf) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAvatar

`func (o *UserSelf) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserSelf) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserSelf) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### GetUid

`func (o *UserSelf) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UserSelf) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UserSelf) SetUid(v string)`

SetUid sets Uid field to given value.


### GetSettings

`func (o *UserSelf) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UserSelf) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UserSelf) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.


### GetType

`func (o *UserSelf) GetType() UserTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSelf) GetTypeOk() (*UserTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSelf) SetType(v UserTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *UserSelf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSystemPermissions

`func (o *UserSelf) GetSystemPermissions() []string`

GetSystemPermissions returns the SystemPermissions field if non-nil, zero value otherwise.

### GetSystemPermissionsOk

`func (o *UserSelf) GetSystemPermissionsOk() (*[]string, bool)`

GetSystemPermissionsOk returns a tuple with the SystemPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPermissions

`func (o *UserSelf) SetSystemPermissions(v []string)`

SetSystemPermissions sets SystemPermissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


