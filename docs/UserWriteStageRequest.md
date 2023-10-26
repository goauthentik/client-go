# UserWriteStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**UserCreationMode** | Pointer to [**UserCreationModeEnum**](UserCreationModeEnum.md) |  | [optional] 
**CreateUsersAsInactive** | Pointer to **bool** | When set, newly created users are inactive and cannot login. | [optional] 
**CreateUsersGroup** | Pointer to **NullableString** | Optionally add newly created users to this group. | [optional] 
**UserType** | Pointer to [**UserTypeEnum**](UserTypeEnum.md) |  | [optional] 
**UserPathTemplate** | Pointer to **string** |  | [optional] 

## Methods

### NewUserWriteStageRequest

`func NewUserWriteStageRequest(name string, ) *UserWriteStageRequest`

NewUserWriteStageRequest instantiates a new UserWriteStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWriteStageRequestWithDefaults

`func NewUserWriteStageRequestWithDefaults() *UserWriteStageRequest`

NewUserWriteStageRequestWithDefaults instantiates a new UserWriteStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserWriteStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserWriteStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserWriteStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowSet

`func (o *UserWriteStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *UserWriteStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *UserWriteStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *UserWriteStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetUserCreationMode

`func (o *UserWriteStageRequest) GetUserCreationMode() UserCreationModeEnum`

GetUserCreationMode returns the UserCreationMode field if non-nil, zero value otherwise.

### GetUserCreationModeOk

`func (o *UserWriteStageRequest) GetUserCreationModeOk() (*UserCreationModeEnum, bool)`

GetUserCreationModeOk returns a tuple with the UserCreationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCreationMode

`func (o *UserWriteStageRequest) SetUserCreationMode(v UserCreationModeEnum)`

SetUserCreationMode sets UserCreationMode field to given value.

### HasUserCreationMode

`func (o *UserWriteStageRequest) HasUserCreationMode() bool`

HasUserCreationMode returns a boolean if a field has been set.

### GetCreateUsersAsInactive

`func (o *UserWriteStageRequest) GetCreateUsersAsInactive() bool`

GetCreateUsersAsInactive returns the CreateUsersAsInactive field if non-nil, zero value otherwise.

### GetCreateUsersAsInactiveOk

`func (o *UserWriteStageRequest) GetCreateUsersAsInactiveOk() (*bool, bool)`

GetCreateUsersAsInactiveOk returns a tuple with the CreateUsersAsInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUsersAsInactive

`func (o *UserWriteStageRequest) SetCreateUsersAsInactive(v bool)`

SetCreateUsersAsInactive sets CreateUsersAsInactive field to given value.

### HasCreateUsersAsInactive

`func (o *UserWriteStageRequest) HasCreateUsersAsInactive() bool`

HasCreateUsersAsInactive returns a boolean if a field has been set.

### GetCreateUsersGroup

`func (o *UserWriteStageRequest) GetCreateUsersGroup() string`

GetCreateUsersGroup returns the CreateUsersGroup field if non-nil, zero value otherwise.

### GetCreateUsersGroupOk

`func (o *UserWriteStageRequest) GetCreateUsersGroupOk() (*string, bool)`

GetCreateUsersGroupOk returns a tuple with the CreateUsersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUsersGroup

`func (o *UserWriteStageRequest) SetCreateUsersGroup(v string)`

SetCreateUsersGroup sets CreateUsersGroup field to given value.

### HasCreateUsersGroup

`func (o *UserWriteStageRequest) HasCreateUsersGroup() bool`

HasCreateUsersGroup returns a boolean if a field has been set.

### SetCreateUsersGroupNil

`func (o *UserWriteStageRequest) SetCreateUsersGroupNil(b bool)`

 SetCreateUsersGroupNil sets the value for CreateUsersGroup to be an explicit nil

### UnsetCreateUsersGroup
`func (o *UserWriteStageRequest) UnsetCreateUsersGroup()`

UnsetCreateUsersGroup ensures that no value is present for CreateUsersGroup, not even an explicit nil
### GetUserType

`func (o *UserWriteStageRequest) GetUserType() UserTypeEnum`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserWriteStageRequest) GetUserTypeOk() (*UserTypeEnum, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserWriteStageRequest) SetUserType(v UserTypeEnum)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UserWriteStageRequest) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetUserPathTemplate

`func (o *UserWriteStageRequest) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *UserWriteStageRequest) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *UserWriteStageRequest) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *UserWriteStageRequest) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


