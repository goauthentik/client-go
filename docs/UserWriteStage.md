# UserWriteStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Component** | **string** | Get object type so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**FlowSet** | Pointer to [**[]FlowSet**](FlowSet.md) |  | [optional] 
**UserCreationMode** | Pointer to [**UserCreationModeEnum**](UserCreationModeEnum.md) |  | [optional] 
**CreateUsersAsInactive** | Pointer to **bool** | When set, newly created users are inactive and cannot login. | [optional] 
**CreateUsersGroup** | Pointer to **NullableString** | Optionally add newly created users to this group. | [optional] 
**UserType** | Pointer to [**UserTypeEnum**](UserTypeEnum.md) |  | [optional] 
**UserPathTemplate** | Pointer to **string** |  | [optional] 

## Methods

### NewUserWriteStage

`func NewUserWriteStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, ) *UserWriteStage`

NewUserWriteStage instantiates a new UserWriteStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWriteStageWithDefaults

`func NewUserWriteStageWithDefaults() *UserWriteStage`

NewUserWriteStageWithDefaults instantiates a new UserWriteStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *UserWriteStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *UserWriteStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *UserWriteStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *UserWriteStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserWriteStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserWriteStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *UserWriteStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *UserWriteStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *UserWriteStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *UserWriteStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *UserWriteStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *UserWriteStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *UserWriteStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *UserWriteStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *UserWriteStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *UserWriteStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *UserWriteStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *UserWriteStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *UserWriteStage) GetFlowSet() []FlowSet`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *UserWriteStage) GetFlowSetOk() (*[]FlowSet, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *UserWriteStage) SetFlowSet(v []FlowSet)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *UserWriteStage) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetUserCreationMode

`func (o *UserWriteStage) GetUserCreationMode() UserCreationModeEnum`

GetUserCreationMode returns the UserCreationMode field if non-nil, zero value otherwise.

### GetUserCreationModeOk

`func (o *UserWriteStage) GetUserCreationModeOk() (*UserCreationModeEnum, bool)`

GetUserCreationModeOk returns a tuple with the UserCreationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCreationMode

`func (o *UserWriteStage) SetUserCreationMode(v UserCreationModeEnum)`

SetUserCreationMode sets UserCreationMode field to given value.

### HasUserCreationMode

`func (o *UserWriteStage) HasUserCreationMode() bool`

HasUserCreationMode returns a boolean if a field has been set.

### GetCreateUsersAsInactive

`func (o *UserWriteStage) GetCreateUsersAsInactive() bool`

GetCreateUsersAsInactive returns the CreateUsersAsInactive field if non-nil, zero value otherwise.

### GetCreateUsersAsInactiveOk

`func (o *UserWriteStage) GetCreateUsersAsInactiveOk() (*bool, bool)`

GetCreateUsersAsInactiveOk returns a tuple with the CreateUsersAsInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUsersAsInactive

`func (o *UserWriteStage) SetCreateUsersAsInactive(v bool)`

SetCreateUsersAsInactive sets CreateUsersAsInactive field to given value.

### HasCreateUsersAsInactive

`func (o *UserWriteStage) HasCreateUsersAsInactive() bool`

HasCreateUsersAsInactive returns a boolean if a field has been set.

### GetCreateUsersGroup

`func (o *UserWriteStage) GetCreateUsersGroup() string`

GetCreateUsersGroup returns the CreateUsersGroup field if non-nil, zero value otherwise.

### GetCreateUsersGroupOk

`func (o *UserWriteStage) GetCreateUsersGroupOk() (*string, bool)`

GetCreateUsersGroupOk returns a tuple with the CreateUsersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUsersGroup

`func (o *UserWriteStage) SetCreateUsersGroup(v string)`

SetCreateUsersGroup sets CreateUsersGroup field to given value.

### HasCreateUsersGroup

`func (o *UserWriteStage) HasCreateUsersGroup() bool`

HasCreateUsersGroup returns a boolean if a field has been set.

### SetCreateUsersGroupNil

`func (o *UserWriteStage) SetCreateUsersGroupNil(b bool)`

 SetCreateUsersGroupNil sets the value for CreateUsersGroup to be an explicit nil

### UnsetCreateUsersGroup
`func (o *UserWriteStage) UnsetCreateUsersGroup()`

UnsetCreateUsersGroup ensures that no value is present for CreateUsersGroup, not even an explicit nil
### GetUserType

`func (o *UserWriteStage) GetUserType() UserTypeEnum`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserWriteStage) GetUserTypeOk() (*UserTypeEnum, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserWriteStage) SetUserType(v UserTypeEnum)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UserWriteStage) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetUserPathTemplate

`func (o *UserWriteStage) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *UserWriteStage) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *UserWriteStage) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *UserWriteStage) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


