# UserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**NumPk** | **int32** | Get a numerical, int32 ID for the group | [readonly] 
**Name** | **string** |  | 
**IsSuperuser** | Pointer to **bool** | Users added to this group will be superusers. | [optional] 
**Parent** | Pointer to **NullableString** |  | [optional] 
**ParentName** | **string** |  | [readonly] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUserGroup

`func NewUserGroup(pk string, numPk int32, name string, parentName string, ) *UserGroup`

NewUserGroup instantiates a new UserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupWithDefaults

`func NewUserGroupWithDefaults() *UserGroup`

NewUserGroupWithDefaults instantiates a new UserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *UserGroup) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *UserGroup) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *UserGroup) SetPk(v string)`

SetPk sets Pk field to given value.


### GetNumPk

`func (o *UserGroup) GetNumPk() int32`

GetNumPk returns the NumPk field if non-nil, zero value otherwise.

### GetNumPkOk

`func (o *UserGroup) GetNumPkOk() (*int32, bool)`

GetNumPkOk returns a tuple with the NumPk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPk

`func (o *UserGroup) SetNumPk(v int32)`

SetNumPk sets NumPk field to given value.


### GetName

`func (o *UserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGroup) SetName(v string)`

SetName sets Name field to given value.


### GetIsSuperuser

`func (o *UserGroup) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *UserGroup) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *UserGroup) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.

### HasIsSuperuser

`func (o *UserGroup) HasIsSuperuser() bool`

HasIsSuperuser returns a boolean if a field has been set.

### GetParent

`func (o *UserGroup) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *UserGroup) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *UserGroup) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *UserGroup) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *UserGroup) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *UserGroup) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetParentName

`func (o *UserGroup) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *UserGroup) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *UserGroup) SetParentName(v string)`

SetParentName sets ParentName field to given value.


### GetAttributes

`func (o *UserGroup) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserGroup) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserGroup) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UserGroup) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


