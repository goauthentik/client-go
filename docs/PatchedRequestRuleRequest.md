# PatchedRequestRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NotificationTransports** | Pointer to **[]string** |  | [optional] 
**NotificationMode** | Pointer to [**NotificationModeEnum**](NotificationModeEnum.md) |  | [optional] 
**MinReviewers** | Pointer to **int32** |  | [optional] 
**MinReviewersIsPerGroup** | Pointer to **bool** |  | [optional] 
**RequestFlow** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPatchedRequestRuleRequest

`func NewPatchedRequestRuleRequest() *PatchedRequestRuleRequest`

NewPatchedRequestRuleRequest instantiates a new PatchedRequestRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedRequestRuleRequestWithDefaults

`func NewPatchedRequestRuleRequestWithDefaults() *PatchedRequestRuleRequest`

NewPatchedRequestRuleRequestWithDefaults instantiates a new PatchedRequestRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PatchedRequestRuleRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PatchedRequestRuleRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PatchedRequestRuleRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PatchedRequestRuleRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *PatchedRequestRuleRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *PatchedRequestRuleRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *PatchedRequestRuleRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *PatchedRequestRuleRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetName

`func (o *PatchedRequestRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedRequestRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedRequestRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedRequestRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotificationTransports

`func (o *PatchedRequestRuleRequest) GetNotificationTransports() []string`

GetNotificationTransports returns the NotificationTransports field if non-nil, zero value otherwise.

### GetNotificationTransportsOk

`func (o *PatchedRequestRuleRequest) GetNotificationTransportsOk() (*[]string, bool)`

GetNotificationTransportsOk returns a tuple with the NotificationTransports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTransports

`func (o *PatchedRequestRuleRequest) SetNotificationTransports(v []string)`

SetNotificationTransports sets NotificationTransports field to given value.

### HasNotificationTransports

`func (o *PatchedRequestRuleRequest) HasNotificationTransports() bool`

HasNotificationTransports returns a boolean if a field has been set.

### GetNotificationMode

`func (o *PatchedRequestRuleRequest) GetNotificationMode() NotificationModeEnum`

GetNotificationMode returns the NotificationMode field if non-nil, zero value otherwise.

### GetNotificationModeOk

`func (o *PatchedRequestRuleRequest) GetNotificationModeOk() (*NotificationModeEnum, bool)`

GetNotificationModeOk returns a tuple with the NotificationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMode

`func (o *PatchedRequestRuleRequest) SetNotificationMode(v NotificationModeEnum)`

SetNotificationMode sets NotificationMode field to given value.

### HasNotificationMode

`func (o *PatchedRequestRuleRequest) HasNotificationMode() bool`

HasNotificationMode returns a boolean if a field has been set.

### GetMinReviewers

`func (o *PatchedRequestRuleRequest) GetMinReviewers() int32`

GetMinReviewers returns the MinReviewers field if non-nil, zero value otherwise.

### GetMinReviewersOk

`func (o *PatchedRequestRuleRequest) GetMinReviewersOk() (*int32, bool)`

GetMinReviewersOk returns a tuple with the MinReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReviewers

`func (o *PatchedRequestRuleRequest) SetMinReviewers(v int32)`

SetMinReviewers sets MinReviewers field to given value.

### HasMinReviewers

`func (o *PatchedRequestRuleRequest) HasMinReviewers() bool`

HasMinReviewers returns a boolean if a field has been set.

### GetMinReviewersIsPerGroup

`func (o *PatchedRequestRuleRequest) GetMinReviewersIsPerGroup() bool`

GetMinReviewersIsPerGroup returns the MinReviewersIsPerGroup field if non-nil, zero value otherwise.

### GetMinReviewersIsPerGroupOk

`func (o *PatchedRequestRuleRequest) GetMinReviewersIsPerGroupOk() (*bool, bool)`

GetMinReviewersIsPerGroupOk returns a tuple with the MinReviewersIsPerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReviewersIsPerGroup

`func (o *PatchedRequestRuleRequest) SetMinReviewersIsPerGroup(v bool)`

SetMinReviewersIsPerGroup sets MinReviewersIsPerGroup field to given value.

### HasMinReviewersIsPerGroup

`func (o *PatchedRequestRuleRequest) HasMinReviewersIsPerGroup() bool`

HasMinReviewersIsPerGroup returns a boolean if a field has been set.

### GetRequestFlow

`func (o *PatchedRequestRuleRequest) GetRequestFlow() string`

GetRequestFlow returns the RequestFlow field if non-nil, zero value otherwise.

### GetRequestFlowOk

`func (o *PatchedRequestRuleRequest) GetRequestFlowOk() (*string, bool)`

GetRequestFlowOk returns a tuple with the RequestFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestFlow

`func (o *PatchedRequestRuleRequest) SetRequestFlow(v string)`

SetRequestFlow sets RequestFlow field to given value.

### HasRequestFlow

`func (o *PatchedRequestRuleRequest) HasRequestFlow() bool`

HasRequestFlow returns a boolean if a field has been set.

### SetRequestFlowNil

`func (o *PatchedRequestRuleRequest) SetRequestFlowNil(b bool)`

 SetRequestFlowNil sets the value for RequestFlow to be an explicit nil

### UnsetRequestFlow
`func (o *PatchedRequestRuleRequest) UnsetRequestFlow()`

UnsetRequestFlow ensures that no value is present for RequestFlow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


