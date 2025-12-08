# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**NumPk** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**IsSuperuser** | Pointer to **bool** | Users added to this group will be superusers. | [optional] 
**Parents** | Pointer to **[]string** |  | [optional] 
**ParentsObj** | [**[]RelatedGroup**](RelatedGroup.md) |  | [readonly] 
**Users** | Pointer to **[]int32** |  | [optional] 
**UsersObj** | [**[]PartialUser**](PartialUser.md) |  | [readonly] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**RolesObj** | [**[]Role**](Role.md) |  | [readonly] 
**Children** | **[]string** |  | [readonly] 
**ChildrenObj** | [**[]RelatedGroup**](RelatedGroup.md) |  | [readonly] 

## Methods

### NewGroup

`func NewGroup(pk string, numPk int32, name string, parentsObj []RelatedGroup, usersObj []PartialUser, rolesObj []Role, children []string, childrenObj []RelatedGroup, ) *Group`

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

### GetParents

`func (o *Group) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *Group) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *Group) SetParents(v []string)`

SetParents sets Parents field to given value.

### HasParents

`func (o *Group) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetParentsObj

`func (o *Group) GetParentsObj() []RelatedGroup`

GetParentsObj returns the ParentsObj field if non-nil, zero value otherwise.

### GetParentsObjOk

`func (o *Group) GetParentsObjOk() (*[]RelatedGroup, bool)`

GetParentsObjOk returns a tuple with the ParentsObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentsObj

`func (o *Group) SetParentsObj(v []RelatedGroup)`

SetParentsObj sets ParentsObj field to given value.


### SetParentsObjNil

`func (o *Group) SetParentsObjNil(b bool)`

 SetParentsObjNil sets the value for ParentsObj to be an explicit nil

### UnsetParentsObj
`func (o *Group) UnsetParentsObj()`

UnsetParentsObj ensures that no value is present for ParentsObj, not even an explicit nil
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

`func (o *Group) GetUsersObj() []PartialUser`

GetUsersObj returns the UsersObj field if non-nil, zero value otherwise.

### GetUsersObjOk

`func (o *Group) GetUsersObjOk() (*[]PartialUser, bool)`

GetUsersObjOk returns a tuple with the UsersObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersObj

`func (o *Group) SetUsersObj(v []PartialUser)`

SetUsersObj sets UsersObj field to given value.


### SetUsersObjNil

`func (o *Group) SetUsersObjNil(b bool)`

 SetUsersObjNil sets the value for UsersObj to be an explicit nil

### UnsetUsersObj
`func (o *Group) UnsetUsersObj()`

UnsetUsersObj ensures that no value is present for UsersObj, not even an explicit nil
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


### GetChildren

`func (o *Group) GetChildren() []string`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Group) GetChildrenOk() (*[]string, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Group) SetChildren(v []string)`

SetChildren sets Children field to given value.


### GetChildrenObj

`func (o *Group) GetChildrenObj() []RelatedGroup`

GetChildrenObj returns the ChildrenObj field if non-nil, zero value otherwise.

### GetChildrenObjOk

`func (o *Group) GetChildrenObjOk() (*[]RelatedGroup, bool)`

GetChildrenObjOk returns a tuple with the ChildrenObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenObj

`func (o *Group) SetChildrenObj(v []RelatedGroup)`

SetChildrenObj sets ChildrenObj field to given value.


### SetChildrenObjNil

`func (o *Group) SetChildrenObjNil(b bool)`

 SetChildrenObjNil sets the value for ChildrenObj to be an explicit nil

### UnsetChildrenObj
`func (o *Group) UnsetChildrenObj()`

UnsetChildrenObj ensures that no value is present for ChildrenObj, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


