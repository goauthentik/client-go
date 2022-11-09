# PolicyBindingUserObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Username** | **string** |  | 
**Name** | **string** | User&#39;s display name. | 
**IsActive** | Pointer to **bool** | Designates whether this user should be treated as active. Unselect this instead of deleting accounts. | [optional] 
**LastLogin** | Pointer to **NullableTime** |  | [optional] 
**IsSuperuser** | **bool** |  | [readonly] 
**Groups** | **[]string** |  | 
**GroupsObj** | [**[]UserGroup**](UserGroup.md) |  | [readonly] 
**Email** | Pointer to **string** |  | [optional] 
**Avatar** | **string** |  | [readonly] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Uid** | **string** |  | [readonly] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicyBindingUserObj

`func NewPolicyBindingUserObj(pk int32, username string, name string, isSuperuser bool, groups []string, groupsObj []UserGroup, avatar string, uid string, ) *PolicyBindingUserObj`

NewPolicyBindingUserObj instantiates a new PolicyBindingUserObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyBindingUserObjWithDefaults

`func NewPolicyBindingUserObjWithDefaults() *PolicyBindingUserObj`

NewPolicyBindingUserObjWithDefaults instantiates a new PolicyBindingUserObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *PolicyBindingUserObj) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *PolicyBindingUserObj) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *PolicyBindingUserObj) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetUsername

`func (o *PolicyBindingUserObj) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PolicyBindingUserObj) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PolicyBindingUserObj) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetName

`func (o *PolicyBindingUserObj) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyBindingUserObj) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyBindingUserObj) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *PolicyBindingUserObj) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PolicyBindingUserObj) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PolicyBindingUserObj) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PolicyBindingUserObj) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastLogin

`func (o *PolicyBindingUserObj) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *PolicyBindingUserObj) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *PolicyBindingUserObj) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *PolicyBindingUserObj) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *PolicyBindingUserObj) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *PolicyBindingUserObj) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetIsSuperuser

`func (o *PolicyBindingUserObj) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *PolicyBindingUserObj) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *PolicyBindingUserObj) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.


### GetGroups

`func (o *PolicyBindingUserObj) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PolicyBindingUserObj) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PolicyBindingUserObj) SetGroups(v []string)`

SetGroups sets Groups field to given value.


### GetGroupsObj

`func (o *PolicyBindingUserObj) GetGroupsObj() []UserGroup`

GetGroupsObj returns the GroupsObj field if non-nil, zero value otherwise.

### GetGroupsObjOk

`func (o *PolicyBindingUserObj) GetGroupsObjOk() (*[]UserGroup, bool)`

GetGroupsObjOk returns a tuple with the GroupsObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsObj

`func (o *PolicyBindingUserObj) SetGroupsObj(v []UserGroup)`

SetGroupsObj sets GroupsObj field to given value.


### GetEmail

`func (o *PolicyBindingUserObj) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PolicyBindingUserObj) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PolicyBindingUserObj) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PolicyBindingUserObj) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAvatar

`func (o *PolicyBindingUserObj) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *PolicyBindingUserObj) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *PolicyBindingUserObj) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### GetAttributes

`func (o *PolicyBindingUserObj) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PolicyBindingUserObj) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PolicyBindingUserObj) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PolicyBindingUserObj) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetUid

`func (o *PolicyBindingUserObj) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *PolicyBindingUserObj) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *PolicyBindingUserObj) SetUid(v string)`

SetUid sets Uid field to given value.


### GetPath

`func (o *PolicyBindingUserObj) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PolicyBindingUserObj) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PolicyBindingUserObj) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PolicyBindingUserObj) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


