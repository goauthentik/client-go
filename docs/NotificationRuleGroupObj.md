# NotificationRuleGroupObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**NumPk** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**IsSuperuser** | Pointer to **bool** | Users added to this group will be superusers. | [optional] 
**Parent** | Pointer to **NullableString** |  | [optional] 
**ParentName** | **string** |  | [readonly] 
**Users** | Pointer to **[]int32** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**UsersObj** | [**[]GroupMember**](GroupMember.md) |  | [readonly] 

## Methods

### NewNotificationRuleGroupObj

`func NewNotificationRuleGroupObj(pk string, numPk int32, name string, parentName string, usersObj []GroupMember, ) *NotificationRuleGroupObj`

NewNotificationRuleGroupObj instantiates a new NotificationRuleGroupObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleGroupObjWithDefaults

`func NewNotificationRuleGroupObjWithDefaults() *NotificationRuleGroupObj`

NewNotificationRuleGroupObjWithDefaults instantiates a new NotificationRuleGroupObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *NotificationRuleGroupObj) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *NotificationRuleGroupObj) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *NotificationRuleGroupObj) SetPk(v string)`

SetPk sets Pk field to given value.


### GetNumPk

`func (o *NotificationRuleGroupObj) GetNumPk() int32`

GetNumPk returns the NumPk field if non-nil, zero value otherwise.

### GetNumPkOk

`func (o *NotificationRuleGroupObj) GetNumPkOk() (*int32, bool)`

GetNumPkOk returns a tuple with the NumPk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPk

`func (o *NotificationRuleGroupObj) SetNumPk(v int32)`

SetNumPk sets NumPk field to given value.


### GetName

`func (o *NotificationRuleGroupObj) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRuleGroupObj) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRuleGroupObj) SetName(v string)`

SetName sets Name field to given value.


### GetIsSuperuser

`func (o *NotificationRuleGroupObj) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *NotificationRuleGroupObj) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *NotificationRuleGroupObj) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.

### HasIsSuperuser

`func (o *NotificationRuleGroupObj) HasIsSuperuser() bool`

HasIsSuperuser returns a boolean if a field has been set.

### GetParent

`func (o *NotificationRuleGroupObj) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *NotificationRuleGroupObj) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *NotificationRuleGroupObj) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *NotificationRuleGroupObj) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *NotificationRuleGroupObj) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *NotificationRuleGroupObj) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetParentName

`func (o *NotificationRuleGroupObj) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *NotificationRuleGroupObj) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *NotificationRuleGroupObj) SetParentName(v string)`

SetParentName sets ParentName field to given value.


### GetUsers

`func (o *NotificationRuleGroupObj) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *NotificationRuleGroupObj) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *NotificationRuleGroupObj) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *NotificationRuleGroupObj) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetAttributes

`func (o *NotificationRuleGroupObj) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NotificationRuleGroupObj) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NotificationRuleGroupObj) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NotificationRuleGroupObj) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetUsersObj

`func (o *NotificationRuleGroupObj) GetUsersObj() []GroupMember`

GetUsersObj returns the UsersObj field if non-nil, zero value otherwise.

### GetUsersObjOk

`func (o *NotificationRuleGroupObj) GetUsersObjOk() (*[]GroupMember, bool)`

GetUsersObjOk returns a tuple with the UsersObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersObj

`func (o *NotificationRuleGroupObj) SetUsersObj(v []GroupMember)`

SetUsersObj sets UsersObj field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


