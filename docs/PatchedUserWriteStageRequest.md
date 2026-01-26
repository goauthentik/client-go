# PatchedUserWriteStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**UserCreationMode** | Pointer to [**UserCreationModeEnum**](UserCreationModeEnum.md) |  | [optional] 
**CreateUsersAsInactive** | Pointer to **bool** | When set, newly created users are inactive and cannot login. | [optional] 
**CreateUsersGroup** | Pointer to **NullableString** | Optionally add newly created users to this group. | [optional] 
**UserType** | Pointer to [**UserTypeEnum**](UserTypeEnum.md) |  | [optional] 
**UserPathTemplate** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedUserWriteStageRequest

`func NewPatchedUserWriteStageRequest() *PatchedUserWriteStageRequest`

NewPatchedUserWriteStageRequest instantiates a new PatchedUserWriteStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedUserWriteStageRequestWithDefaults

`func NewPatchedUserWriteStageRequestWithDefaults() *PatchedUserWriteStageRequest`

NewPatchedUserWriteStageRequestWithDefaults instantiates a new PatchedUserWriteStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedUserWriteStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedUserWriteStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedUserWriteStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedUserWriteStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlowSet

`func (o *PatchedUserWriteStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PatchedUserWriteStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PatchedUserWriteStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *PatchedUserWriteStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetUserCreationMode

`func (o *PatchedUserWriteStageRequest) GetUserCreationMode() UserCreationModeEnum`

GetUserCreationMode returns the UserCreationMode field if non-nil, zero value otherwise.

### GetUserCreationModeOk

`func (o *PatchedUserWriteStageRequest) GetUserCreationModeOk() (*UserCreationModeEnum, bool)`

GetUserCreationModeOk returns a tuple with the UserCreationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCreationMode

`func (o *PatchedUserWriteStageRequest) SetUserCreationMode(v UserCreationModeEnum)`

SetUserCreationMode sets UserCreationMode field to given value.

### HasUserCreationMode

`func (o *PatchedUserWriteStageRequest) HasUserCreationMode() bool`

HasUserCreationMode returns a boolean if a field has been set.

### GetCreateUsersAsInactive

`func (o *PatchedUserWriteStageRequest) GetCreateUsersAsInactive() bool`

GetCreateUsersAsInactive returns the CreateUsersAsInactive field if non-nil, zero value otherwise.

### GetCreateUsersAsInactiveOk

`func (o *PatchedUserWriteStageRequest) GetCreateUsersAsInactiveOk() (*bool, bool)`

GetCreateUsersAsInactiveOk returns a tuple with the CreateUsersAsInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUsersAsInactive

`func (o *PatchedUserWriteStageRequest) SetCreateUsersAsInactive(v bool)`

SetCreateUsersAsInactive sets CreateUsersAsInactive field to given value.

### HasCreateUsersAsInactive

`func (o *PatchedUserWriteStageRequest) HasCreateUsersAsInactive() bool`

HasCreateUsersAsInactive returns a boolean if a field has been set.

### GetCreateUsersGroup

`func (o *PatchedUserWriteStageRequest) GetCreateUsersGroup() string`

GetCreateUsersGroup returns the CreateUsersGroup field if non-nil, zero value otherwise.

### GetCreateUsersGroupOk

`func (o *PatchedUserWriteStageRequest) GetCreateUsersGroupOk() (*string, bool)`

GetCreateUsersGroupOk returns a tuple with the CreateUsersGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUsersGroup

`func (o *PatchedUserWriteStageRequest) SetCreateUsersGroup(v string)`

SetCreateUsersGroup sets CreateUsersGroup field to given value.

### HasCreateUsersGroup

`func (o *PatchedUserWriteStageRequest) HasCreateUsersGroup() bool`

HasCreateUsersGroup returns a boolean if a field has been set.

### SetCreateUsersGroupNil

`func (o *PatchedUserWriteStageRequest) SetCreateUsersGroupNil(b bool)`

 SetCreateUsersGroupNil sets the value for CreateUsersGroup to be an explicit nil

### UnsetCreateUsersGroup
`func (o *PatchedUserWriteStageRequest) UnsetCreateUsersGroup()`

UnsetCreateUsersGroup ensures that no value is present for CreateUsersGroup, not even an explicit nil
### GetUserType

`func (o *PatchedUserWriteStageRequest) GetUserType() UserTypeEnum`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *PatchedUserWriteStageRequest) GetUserTypeOk() (*UserTypeEnum, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *PatchedUserWriteStageRequest) SetUserType(v UserTypeEnum)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *PatchedUserWriteStageRequest) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetUserPathTemplate

`func (o *PatchedUserWriteStageRequest) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *PatchedUserWriteStageRequest) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *PatchedUserWriteStageRequest) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *PatchedUserWriteStageRequest) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


