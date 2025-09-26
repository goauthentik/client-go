# PartialGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**NumPk** | **int32** | Get a numerical, int32 ID for the group | [readonly] 
**Name** | **string** |  | 
**IsSuperuser** | Pointer to **bool** | Users added to this group will be superusers. | [optional] 
**Parent** | Pointer to **NullableString** |  | [optional] 
**ParentName** | **NullableString** |  | [readonly] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPartialGroup

`func NewPartialGroup(pk string, numPk int32, name string, parentName NullableString, ) *PartialGroup`

NewPartialGroup instantiates a new PartialGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialGroupWithDefaults

`func NewPartialGroupWithDefaults() *PartialGroup`

NewPartialGroupWithDefaults instantiates a new PartialGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *PartialGroup) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *PartialGroup) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *PartialGroup) SetPk(v string)`

SetPk sets Pk field to given value.


### GetNumPk

`func (o *PartialGroup) GetNumPk() int32`

GetNumPk returns the NumPk field if non-nil, zero value otherwise.

### GetNumPkOk

`func (o *PartialGroup) GetNumPkOk() (*int32, bool)`

GetNumPkOk returns a tuple with the NumPk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPk

`func (o *PartialGroup) SetNumPk(v int32)`

SetNumPk sets NumPk field to given value.


### GetName

`func (o *PartialGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartialGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartialGroup) SetName(v string)`

SetName sets Name field to given value.


### GetIsSuperuser

`func (o *PartialGroup) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *PartialGroup) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *PartialGroup) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.

### HasIsSuperuser

`func (o *PartialGroup) HasIsSuperuser() bool`

HasIsSuperuser returns a boolean if a field has been set.

### GetParent

`func (o *PartialGroup) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PartialGroup) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PartialGroup) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PartialGroup) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PartialGroup) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PartialGroup) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetParentName

`func (o *PartialGroup) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *PartialGroup) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *PartialGroup) SetParentName(v string)`

SetParentName sets ParentName field to given value.


### SetParentNameNil

`func (o *PartialGroup) SetParentNameNil(b bool)`

 SetParentNameNil sets the value for ParentName to be an explicit nil

### UnsetParentName
`func (o *PartialGroup) UnsetParentName()`

UnsetParentName ensures that no value is present for ParentName, not even an explicit nil
### GetAttributes

`func (o *PartialGroup) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PartialGroup) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PartialGroup) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PartialGroup) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


