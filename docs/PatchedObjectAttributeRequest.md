# PatchedObjectAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**Key** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Regex** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ObjectAttributeTypeEnum**](ObjectAttributeTypeEnum.md) |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 
**IsUnique** | Pointer to **bool** |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedObjectAttributeRequest

`func NewPatchedObjectAttributeRequest() *PatchedObjectAttributeRequest`

NewPatchedObjectAttributeRequest instantiates a new PatchedObjectAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedObjectAttributeRequestWithDefaults

`func NewPatchedObjectAttributeRequestWithDefaults() *PatchedObjectAttributeRequest`

NewPatchedObjectAttributeRequestWithDefaults instantiates a new PatchedObjectAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *PatchedObjectAttributeRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PatchedObjectAttributeRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PatchedObjectAttributeRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PatchedObjectAttributeRequest) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedObjectAttributeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedObjectAttributeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedObjectAttributeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedObjectAttributeRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetKey

`func (o *PatchedObjectAttributeRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PatchedObjectAttributeRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PatchedObjectAttributeRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PatchedObjectAttributeRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedObjectAttributeRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedObjectAttributeRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedObjectAttributeRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedObjectAttributeRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRegex

`func (o *PatchedObjectAttributeRequest) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *PatchedObjectAttributeRequest) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *PatchedObjectAttributeRequest) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *PatchedObjectAttributeRequest) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetType

`func (o *PatchedObjectAttributeRequest) GetType() ObjectAttributeTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedObjectAttributeRequest) GetTypeOk() (*ObjectAttributeTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedObjectAttributeRequest) SetType(v ObjectAttributeTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedObjectAttributeRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGroup

`func (o *PatchedObjectAttributeRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedObjectAttributeRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedObjectAttributeRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedObjectAttributeRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetManaged

`func (o *PatchedObjectAttributeRequest) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *PatchedObjectAttributeRequest) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *PatchedObjectAttributeRequest) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *PatchedObjectAttributeRequest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *PatchedObjectAttributeRequest) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *PatchedObjectAttributeRequest) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetIsUnique

`func (o *PatchedObjectAttributeRequest) GetIsUnique() bool`

GetIsUnique returns the IsUnique field if non-nil, zero value otherwise.

### GetIsUniqueOk

`func (o *PatchedObjectAttributeRequest) GetIsUniqueOk() (*bool, bool)`

GetIsUniqueOk returns a tuple with the IsUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnique

`func (o *PatchedObjectAttributeRequest) SetIsUnique(v bool)`

SetIsUnique sets IsUnique field to given value.

### HasIsUnique

`func (o *PatchedObjectAttributeRequest) HasIsUnique() bool`

HasIsUnique returns a boolean if a field has been set.

### GetIsRequired

`func (o *PatchedObjectAttributeRequest) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *PatchedObjectAttributeRequest) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *PatchedObjectAttributeRequest) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *PatchedObjectAttributeRequest) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


