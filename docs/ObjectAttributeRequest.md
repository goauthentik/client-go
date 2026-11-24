# ObjectAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**Key** | **string** |  | 
**Label** | **string** |  | 
**Regex** | Pointer to **string** |  | [optional] 
**Type** | [**ObjectAttributeTypeEnum**](ObjectAttributeTypeEnum.md) |  | 
**Group** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 
**IsUnique** | Pointer to **bool** |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewObjectAttributeRequest

`func NewObjectAttributeRequest(objectType string, key string, label string, type_ ObjectAttributeTypeEnum, ) *ObjectAttributeRequest`

NewObjectAttributeRequest instantiates a new ObjectAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectAttributeRequestWithDefaults

`func NewObjectAttributeRequestWithDefaults() *ObjectAttributeRequest`

NewObjectAttributeRequestWithDefaults instantiates a new ObjectAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *ObjectAttributeRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ObjectAttributeRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ObjectAttributeRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnabled

`func (o *ObjectAttributeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ObjectAttributeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ObjectAttributeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ObjectAttributeRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetKey

`func (o *ObjectAttributeRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ObjectAttributeRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ObjectAttributeRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *ObjectAttributeRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ObjectAttributeRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ObjectAttributeRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRegex

`func (o *ObjectAttributeRequest) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *ObjectAttributeRequest) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *ObjectAttributeRequest) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *ObjectAttributeRequest) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetType

`func (o *ObjectAttributeRequest) GetType() ObjectAttributeTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ObjectAttributeRequest) GetTypeOk() (*ObjectAttributeTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ObjectAttributeRequest) SetType(v ObjectAttributeTypeEnum)`

SetType sets Type field to given value.


### GetGroup

`func (o *ObjectAttributeRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ObjectAttributeRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ObjectAttributeRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ObjectAttributeRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetManaged

`func (o *ObjectAttributeRequest) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ObjectAttributeRequest) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ObjectAttributeRequest) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ObjectAttributeRequest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *ObjectAttributeRequest) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *ObjectAttributeRequest) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetIsUnique

`func (o *ObjectAttributeRequest) GetIsUnique() bool`

GetIsUnique returns the IsUnique field if non-nil, zero value otherwise.

### GetIsUniqueOk

`func (o *ObjectAttributeRequest) GetIsUniqueOk() (*bool, bool)`

GetIsUniqueOk returns a tuple with the IsUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnique

`func (o *ObjectAttributeRequest) SetIsUnique(v bool)`

SetIsUnique sets IsUnique field to given value.

### HasIsUnique

`func (o *ObjectAttributeRequest) HasIsUnique() bool`

HasIsUnique returns a boolean if a field has been set.

### GetIsRequired

`func (o *ObjectAttributeRequest) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *ObjectAttributeRequest) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *ObjectAttributeRequest) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *ObjectAttributeRequest) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


