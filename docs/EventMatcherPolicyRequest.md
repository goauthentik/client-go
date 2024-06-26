# EventMatcherPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**Action** | Pointer to [**NullableEventActions**](EventActions.md) | Match created events with this action type. When left empty, all action types will be matched. | [optional] 
**ClientIp** | Pointer to **NullableString** | Matches Event&#39;s Client IP (strict matching, for network matching use an Expression Policy) | [optional] 
**App** | Pointer to [**NullableAppEnum**](AppEnum.md) | Match events created by selected application. When left empty, all applications are matched. | [optional] 
**Model** | Pointer to [**NullableModelEnum**](ModelEnum.md) | Match events created by selected model. When left empty, all models are matched. When an app is selected, all the application&#39;s models are matched. | [optional] 

## Methods

### NewEventMatcherPolicyRequest

`func NewEventMatcherPolicyRequest(name string, ) *EventMatcherPolicyRequest`

NewEventMatcherPolicyRequest instantiates a new EventMatcherPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventMatcherPolicyRequestWithDefaults

`func NewEventMatcherPolicyRequestWithDefaults() *EventMatcherPolicyRequest`

NewEventMatcherPolicyRequestWithDefaults instantiates a new EventMatcherPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventMatcherPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventMatcherPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventMatcherPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExecutionLogging

`func (o *EventMatcherPolicyRequest) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *EventMatcherPolicyRequest) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *EventMatcherPolicyRequest) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *EventMatcherPolicyRequest) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetAction

`func (o *EventMatcherPolicyRequest) GetAction() EventActions`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *EventMatcherPolicyRequest) GetActionOk() (*EventActions, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *EventMatcherPolicyRequest) SetAction(v EventActions)`

SetAction sets Action field to given value.

### HasAction

`func (o *EventMatcherPolicyRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *EventMatcherPolicyRequest) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *EventMatcherPolicyRequest) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetClientIp

`func (o *EventMatcherPolicyRequest) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *EventMatcherPolicyRequest) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *EventMatcherPolicyRequest) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *EventMatcherPolicyRequest) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### SetClientIpNil

`func (o *EventMatcherPolicyRequest) SetClientIpNil(b bool)`

 SetClientIpNil sets the value for ClientIp to be an explicit nil

### UnsetClientIp
`func (o *EventMatcherPolicyRequest) UnsetClientIp()`

UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
### GetApp

`func (o *EventMatcherPolicyRequest) GetApp() AppEnum`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *EventMatcherPolicyRequest) GetAppOk() (*AppEnum, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *EventMatcherPolicyRequest) SetApp(v AppEnum)`

SetApp sets App field to given value.

### HasApp

`func (o *EventMatcherPolicyRequest) HasApp() bool`

HasApp returns a boolean if a field has been set.

### SetAppNil

`func (o *EventMatcherPolicyRequest) SetAppNil(b bool)`

 SetAppNil sets the value for App to be an explicit nil

### UnsetApp
`func (o *EventMatcherPolicyRequest) UnsetApp()`

UnsetApp ensures that no value is present for App, not even an explicit nil
### GetModel

`func (o *EventMatcherPolicyRequest) GetModel() ModelEnum`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EventMatcherPolicyRequest) GetModelOk() (*ModelEnum, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EventMatcherPolicyRequest) SetModel(v ModelEnum)`

SetModel sets Model field to given value.

### HasModel

`func (o *EventMatcherPolicyRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *EventMatcherPolicyRequest) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *EventMatcherPolicyRequest) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


