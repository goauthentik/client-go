# RequestRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**PbmUuid** | **string** |  | [readonly] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**Name** | **string** |  | 
**Targets** | **[]string** |  | [readonly] 
**NotificationTransports** | Pointer to **[]string** |  | [optional] 
**NotificationMode** | Pointer to [**NotificationModeEnum**](NotificationModeEnum.md) |  | [optional] 
**MinReviewers** | Pointer to **int32** |  | [optional] 
**MinReviewersIsPerGroup** | Pointer to **bool** |  | [optional] 
**RequestFlow** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRequestRule

`func NewRequestRule(pbmUuid string, name string, targets []string, ) *RequestRule`

NewRequestRule instantiates a new RequestRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestRuleWithDefaults

`func NewRequestRuleWithDefaults() *RequestRule`

NewRequestRuleWithDefaults instantiates a new RequestRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RequestRule) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RequestRule) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RequestRule) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RequestRule) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetPbmUuid

`func (o *RequestRule) GetPbmUuid() string`

GetPbmUuid returns the PbmUuid field if non-nil, zero value otherwise.

### GetPbmUuidOk

`func (o *RequestRule) GetPbmUuidOk() (*string, bool)`

GetPbmUuidOk returns a tuple with the PbmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbmUuid

`func (o *RequestRule) SetPbmUuid(v string)`

SetPbmUuid sets PbmUuid field to given value.


### GetPolicyEngineMode

`func (o *RequestRule) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *RequestRule) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *RequestRule) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *RequestRule) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetName

`func (o *RequestRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestRule) SetName(v string)`

SetName sets Name field to given value.


### GetTargets

`func (o *RequestRule) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *RequestRule) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *RequestRule) SetTargets(v []string)`

SetTargets sets Targets field to given value.


### GetNotificationTransports

`func (o *RequestRule) GetNotificationTransports() []string`

GetNotificationTransports returns the NotificationTransports field if non-nil, zero value otherwise.

### GetNotificationTransportsOk

`func (o *RequestRule) GetNotificationTransportsOk() (*[]string, bool)`

GetNotificationTransportsOk returns a tuple with the NotificationTransports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTransports

`func (o *RequestRule) SetNotificationTransports(v []string)`

SetNotificationTransports sets NotificationTransports field to given value.

### HasNotificationTransports

`func (o *RequestRule) HasNotificationTransports() bool`

HasNotificationTransports returns a boolean if a field has been set.

### GetNotificationMode

`func (o *RequestRule) GetNotificationMode() NotificationModeEnum`

GetNotificationMode returns the NotificationMode field if non-nil, zero value otherwise.

### GetNotificationModeOk

`func (o *RequestRule) GetNotificationModeOk() (*NotificationModeEnum, bool)`

GetNotificationModeOk returns a tuple with the NotificationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMode

`func (o *RequestRule) SetNotificationMode(v NotificationModeEnum)`

SetNotificationMode sets NotificationMode field to given value.

### HasNotificationMode

`func (o *RequestRule) HasNotificationMode() bool`

HasNotificationMode returns a boolean if a field has been set.

### GetMinReviewers

`func (o *RequestRule) GetMinReviewers() int32`

GetMinReviewers returns the MinReviewers field if non-nil, zero value otherwise.

### GetMinReviewersOk

`func (o *RequestRule) GetMinReviewersOk() (*int32, bool)`

GetMinReviewersOk returns a tuple with the MinReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReviewers

`func (o *RequestRule) SetMinReviewers(v int32)`

SetMinReviewers sets MinReviewers field to given value.

### HasMinReviewers

`func (o *RequestRule) HasMinReviewers() bool`

HasMinReviewers returns a boolean if a field has been set.

### GetMinReviewersIsPerGroup

`func (o *RequestRule) GetMinReviewersIsPerGroup() bool`

GetMinReviewersIsPerGroup returns the MinReviewersIsPerGroup field if non-nil, zero value otherwise.

### GetMinReviewersIsPerGroupOk

`func (o *RequestRule) GetMinReviewersIsPerGroupOk() (*bool, bool)`

GetMinReviewersIsPerGroupOk returns a tuple with the MinReviewersIsPerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReviewersIsPerGroup

`func (o *RequestRule) SetMinReviewersIsPerGroup(v bool)`

SetMinReviewersIsPerGroup sets MinReviewersIsPerGroup field to given value.

### HasMinReviewersIsPerGroup

`func (o *RequestRule) HasMinReviewersIsPerGroup() bool`

HasMinReviewersIsPerGroup returns a boolean if a field has been set.

### GetRequestFlow

`func (o *RequestRule) GetRequestFlow() string`

GetRequestFlow returns the RequestFlow field if non-nil, zero value otherwise.

### GetRequestFlowOk

`func (o *RequestRule) GetRequestFlowOk() (*string, bool)`

GetRequestFlowOk returns a tuple with the RequestFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestFlow

`func (o *RequestRule) SetRequestFlow(v string)`

SetRequestFlow sets RequestFlow field to given value.

### HasRequestFlow

`func (o *RequestRule) HasRequestFlow() bool`

HasRequestFlow returns a boolean if a field has been set.

### SetRequestFlowNil

`func (o *RequestRule) SetRequestFlowNil(b bool)`

 SetRequestFlowNil sets the value for RequestFlow to be an explicit nil

### UnsetRequestFlow
`func (o *RequestRule) UnsetRequestFlow()`

UnsetRequestFlow ensures that no value is present for RequestFlow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


