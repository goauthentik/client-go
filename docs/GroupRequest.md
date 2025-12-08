# GroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**IsSuperuser** | Pointer to **bool** | Users added to this group will be superusers. | [optional] 
**Parents** | Pointer to **[]string** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGroupRequest

`func NewGroupRequest(name string, ) *GroupRequest`

NewGroupRequest instantiates a new GroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRequestWithDefaults

`func NewGroupRequestWithDefaults() *GroupRequest`

NewGroupRequestWithDefaults instantiates a new GroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsSuperuser

`func (o *GroupRequest) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *GroupRequest) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *GroupRequest) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.

### HasIsSuperuser

`func (o *GroupRequest) HasIsSuperuser() bool`

HasIsSuperuser returns a boolean if a field has been set.

### GetParents

`func (o *GroupRequest) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *GroupRequest) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *GroupRequest) SetParents(v []string)`

SetParents sets Parents field to given value.

### HasParents

`func (o *GroupRequest) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetUsers

`func (o *GroupRequest) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GroupRequest) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GroupRequest) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GroupRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetAttributes

`func (o *GroupRequest) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GroupRequest) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GroupRequest) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GroupRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRoles

`func (o *GroupRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GroupRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GroupRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GroupRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


