# ObjectAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**ObjectType** | **string** |  | 
**ObjectTypeObj** | [**ContentType**](ContentType.md) |  | [readonly] 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**Created** | **time.Time** |  | [readonly] 
**Key** | **string** |  | 
**Label** | **string** |  | 
**LastUpdated** | **time.Time** |  | [readonly] 
**Regex** | Pointer to **string** |  | [optional] 
**Type** | [**ObjectAttributeTypeEnum**](ObjectAttributeTypeEnum.md) |  | 
**Group** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 
**IsUnique** | Pointer to **bool** |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewObjectAttribute

`func NewObjectAttribute(pk string, objectType string, objectTypeObj ContentType, created time.Time, key string, label string, lastUpdated time.Time, type_ ObjectAttributeTypeEnum, ) *ObjectAttribute`

NewObjectAttribute instantiates a new ObjectAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectAttributeWithDefaults

`func NewObjectAttributeWithDefaults() *ObjectAttribute`

NewObjectAttributeWithDefaults instantiates a new ObjectAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *ObjectAttribute) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *ObjectAttribute) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *ObjectAttribute) SetPk(v string)`

SetPk sets Pk field to given value.


### GetObjectType

`func (o *ObjectAttribute) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ObjectAttribute) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ObjectAttribute) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectTypeObj

`func (o *ObjectAttribute) GetObjectTypeObj() ContentType`

GetObjectTypeObj returns the ObjectTypeObj field if non-nil, zero value otherwise.

### GetObjectTypeObjOk

`func (o *ObjectAttribute) GetObjectTypeObjOk() (*ContentType, bool)`

GetObjectTypeObjOk returns a tuple with the ObjectTypeObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeObj

`func (o *ObjectAttribute) SetObjectTypeObj(v ContentType)`

SetObjectTypeObj sets ObjectTypeObj field to given value.


### GetEnabled

`func (o *ObjectAttribute) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ObjectAttribute) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ObjectAttribute) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ObjectAttribute) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreated

`func (o *ObjectAttribute) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ObjectAttribute) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ObjectAttribute) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetKey

`func (o *ObjectAttribute) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ObjectAttribute) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ObjectAttribute) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *ObjectAttribute) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ObjectAttribute) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ObjectAttribute) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLastUpdated

`func (o *ObjectAttribute) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ObjectAttribute) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ObjectAttribute) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetRegex

`func (o *ObjectAttribute) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *ObjectAttribute) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *ObjectAttribute) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *ObjectAttribute) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetType

`func (o *ObjectAttribute) GetType() ObjectAttributeTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ObjectAttribute) GetTypeOk() (*ObjectAttributeTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ObjectAttribute) SetType(v ObjectAttributeTypeEnum)`

SetType sets Type field to given value.


### GetGroup

`func (o *ObjectAttribute) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ObjectAttribute) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ObjectAttribute) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ObjectAttribute) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetManaged

`func (o *ObjectAttribute) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ObjectAttribute) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ObjectAttribute) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ObjectAttribute) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *ObjectAttribute) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *ObjectAttribute) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetIsUnique

`func (o *ObjectAttribute) GetIsUnique() bool`

GetIsUnique returns the IsUnique field if non-nil, zero value otherwise.

### GetIsUniqueOk

`func (o *ObjectAttribute) GetIsUniqueOk() (*bool, bool)`

GetIsUniqueOk returns a tuple with the IsUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnique

`func (o *ObjectAttribute) SetIsUnique(v bool)`

SetIsUnique sets IsUnique field to given value.

### HasIsUnique

`func (o *ObjectAttribute) HasIsUnique() bool`

HasIsUnique returns a boolean if a field has been set.

### GetIsRequired

`func (o *ObjectAttribute) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *ObjectAttribute) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *ObjectAttribute) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *ObjectAttribute) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


