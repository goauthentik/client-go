# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**IsSuperuser** | Pointer to **bool** | Users added to this group will be superusers. | [optional] 
**Parent** | **NullableString** |  | 
**Users** | **[]int32** |  | 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**UsersObj** | [**[]GroupMember**](GroupMember.md) |  | [readonly] 

## Methods

### NewGroup

`func NewGroup(pk string, name string, parent NullableString, users []int32, usersObj []GroupMember, ) *Group`

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


### SetParentNil

`func (o *Group) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *Group) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


