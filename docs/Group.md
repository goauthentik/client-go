# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**NumPk** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**IsSuperuser** | Pointer to **bool** | Users added to this group will be superusers. | [optional] 
**Parent** | Pointer to **NullableString** |  | [optional] 
**ParentName** | **NullableString** |  | [readonly] 
**Users** | Pointer to **[]int32** |  | [optional] 
**UsersObj** | [**[]GroupMember**](GroupMember.md) |  | [readonly] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**RolesObj** | [**[]Role**](Role.md) |  | [readonly] 

## Methods

### NewGroup

`func NewGroup(pk string, numPk int32, name string, parentName NullableString, usersObj []GroupMember, rolesObj []Role, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *Group) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *Group) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *Group) SetPk(v string)`

SetPk sets Pk field to given value.


### GetNumPk

`func (o *Group) GetNumPk() int32`

GetNumPk returns the NumPk field if non-nil, zero value otherwise.

### GetNumPkOk

`func (o *Group) GetNumPkOk() (*int32, bool)`

GetNumPkOk returns a tuple with the NumPk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPk

`func (o *Group) SetNumPk(v int32)`

SetNumPk sets NumPk field to given value.


### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.


### GetIsSuperuser

`func (o *Group) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *Group) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *Group) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.

### HasIsSuperuser

`func (o *Group) HasIsSuperuser() bool`

HasIsSuperuser returns a boolean if a field has been set.

### GetParent

`func (o *Group) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Group) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Group) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Group) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *Group) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *Group) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetParentName

`func (o *Group) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *Group) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *Group) SetParentName(v string)`

SetParentName sets ParentName field to given value.


### SetParentNameNil

`func (o *Group) SetParentNameNil(b bool)`

 SetParentNameNil sets the value for ParentName to be an explicit nil

### UnsetParentName
`func (o *Group) UnsetParentName()`

UnsetParentName ensures that no value is present for ParentName, not even an explicit nil
### GetUsers

`func (o *Group) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Group) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Group) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Group) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetUsersObj

`func (o *Group) GetUsersObj() []GroupMember`

GetUsersObj returns the UsersObj field if non-nil, zero value otherwise.

### GetUsersObjOk

`func (o *Group) GetUsersObjOk() (*[]GroupMember, bool)`

GetUsersObjOk returns a tuple with the UsersObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersObj

`func (o *Group) SetUsersObj(v []GroupMember)`

SetUsersObj sets UsersObj field to given value.


### GetAttributes

`func (o *Group) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Group) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Group) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Group) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRoles

`func (o *Group) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Group) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Group) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Group) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRolesObj

`func (o *Group) GetRolesObj() []Role`

GetRolesObj returns the RolesObj field if non-nil, zero value otherwise.

### GetRolesObjOk

`func (o *Group) GetRolesObjOk() (*[]Role, bool)`

GetRolesObjOk returns a tuple with the RolesObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolesObj

`func (o *Group) SetRolesObj(v []Role)`

SetRolesObj sets RolesObj field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


