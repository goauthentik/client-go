# RequestRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**Name** | **string** |  | 
**NotificationTransports** | Pointer to **[]string** |  | [optional] 
**NotificationMode** | Pointer to [**NotificationModeEnum**](NotificationModeEnum.md) |  | [optional] 
**MinReviewers** | Pointer to **int32** |  | [optional] 
**MinReviewersIsPerGroup** | Pointer to **bool** |  | [optional] 
**RequestFlow** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRequestRuleRequest

`func NewRequestRuleRequest(name string, ) *RequestRuleRequest`

NewRequestRuleRequest instantiates a new RequestRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestRuleRequestWithDefaults

`func NewRequestRuleRequestWithDefaults() *RequestRuleRequest`

NewRequestRuleRequestWithDefaults instantiates a new RequestRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RequestRuleRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RequestRuleRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RequestRuleRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RequestRuleRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *RequestRuleRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *RequestRuleRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *RequestRuleRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *RequestRuleRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetName

`func (o *RequestRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNotificationTransports

`func (o *RequestRuleRequest) GetNotificationTransports() []string`

GetNotificationTransports returns the NotificationTransports field if non-nil, zero value otherwise.

### GetNotificationTransportsOk

`func (o *RequestRuleRequest) GetNotificationTransportsOk() (*[]string, bool)`

GetNotificationTransportsOk returns a tuple with the NotificationTransports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTransports

`func (o *RequestRuleRequest) SetNotificationTransports(v []string)`

SetNotificationTransports sets NotificationTransports field to given value.

### HasNotificationTransports

`func (o *RequestRuleRequest) HasNotificationTransports() bool`

HasNotificationTransports returns a boolean if a field has been set.

### GetNotificationMode

`func (o *RequestRuleRequest) GetNotificationMode() NotificationModeEnum`

GetNotificationMode returns the NotificationMode field if non-nil, zero value otherwise.

### GetNotificationModeOk

`func (o *RequestRuleRequest) GetNotificationModeOk() (*NotificationModeEnum, bool)`

GetNotificationModeOk returns a tuple with the NotificationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMode

`func (o *RequestRuleRequest) SetNotificationMode(v NotificationModeEnum)`

SetNotificationMode sets NotificationMode field to given value.

### HasNotificationMode

`func (o *RequestRuleRequest) HasNotificationMode() bool`

HasNotificationMode returns a boolean if a field has been set.

### GetMinReviewers

`func (o *RequestRuleRequest) GetMinReviewers() int32`

GetMinReviewers returns the MinReviewers field if non-nil, zero value otherwise.

### GetMinReviewersOk

`func (o *RequestRuleRequest) GetMinReviewersOk() (*int32, bool)`

GetMinReviewersOk returns a tuple with the MinReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReviewers

`func (o *RequestRuleRequest) SetMinReviewers(v int32)`

SetMinReviewers sets MinReviewers field to given value.

### HasMinReviewers

`func (o *RequestRuleRequest) HasMinReviewers() bool`

HasMinReviewers returns a boolean if a field has been set.

### GetMinReviewersIsPerGroup

`func (o *RequestRuleRequest) GetMinReviewersIsPerGroup() bool`

GetMinReviewersIsPerGroup returns the MinReviewersIsPerGroup field if non-nil, zero value otherwise.

### GetMinReviewersIsPerGroupOk

`func (o *RequestRuleRequest) GetMinReviewersIsPerGroupOk() (*bool, bool)`

GetMinReviewersIsPerGroupOk returns a tuple with the MinReviewersIsPerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReviewersIsPerGroup

`func (o *RequestRuleRequest) SetMinReviewersIsPerGroup(v bool)`

SetMinReviewersIsPerGroup sets MinReviewersIsPerGroup field to given value.

### HasMinReviewersIsPerGroup

`func (o *RequestRuleRequest) HasMinReviewersIsPerGroup() bool`

HasMinReviewersIsPerGroup returns a boolean if a field has been set.

### GetRequestFlow

`func (o *RequestRuleRequest) GetRequestFlow() string`

GetRequestFlow returns the RequestFlow field if non-nil, zero value otherwise.

### GetRequestFlowOk

`func (o *RequestRuleRequest) GetRequestFlowOk() (*string, bool)`

GetRequestFlowOk returns a tuple with the RequestFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestFlow

`func (o *RequestRuleRequest) SetRequestFlow(v string)`

SetRequestFlow sets RequestFlow field to given value.

### HasRequestFlow

`func (o *RequestRuleRequest) HasRequestFlow() bool`

HasRequestFlow returns a boolean if a field has been set.

### SetRequestFlowNil

`func (o *RequestRuleRequest) SetRequestFlowNil(b bool)`

 SetRequestFlowNil sets the value for RequestFlow to be an explicit nil

### UnsetRequestFlow
`func (o *RequestRuleRequest) UnsetRequestFlow()`

UnsetRequestFlow ensures that no value is present for RequestFlow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


