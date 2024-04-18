# UserGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**IsSuperuser** | Pointer to **bool** | Users added to this group will be superusers. | [optional] 
**Parent** | Pointer to **NullableString** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUserGroupRequest

`func NewUserGroupRequest(name string, ) *UserGroupRequest`

NewUserGroupRequest instantiates a new UserGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupRequestWithDefaults

`func NewUserGroupRequestWithDefaults() *UserGroupRequest`

NewUserGroupRequestWithDefaults instantiates a new UserGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsSuperuser

`func (o *UserGroupRequest) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *UserGroupRequest) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *UserGroupRequest) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.

### HasIsSuperuser

`func (o *UserGroupRequest) HasIsSuperuser() bool`

HasIsSuperuser returns a boolean if a field has been set.

### GetParent

`func (o *UserGroupRequest) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *UserGroupRequest) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *UserGroupRequest) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *UserGroupRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *UserGroupRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *UserGroupRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetAttributes

`func (o *UserGroupRequest) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserGroupRequest) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserGroupRequest) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UserGroupRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


