# PatchedEventMatcherPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**Action** | Pointer to [**NullableEventActions**](EventActions.md) | Match created events with this action type. When left empty, all action types will be matched. | [optional] 
**ClientIp** | Pointer to **string** | Matches Event&#39;s Client IP (strict matching, for network matching use an Expression Policy) | [optional] 
**App** | Pointer to [**NullableAppEnum**](AppEnum.md) | Match events created by selected application. When left empty, all applications are matched. | [optional] 

## Methods

### NewPatchedEventMatcherPolicyRequest

`func NewPatchedEventMatcherPolicyRequest() *PatchedEventMatcherPolicyRequest`

NewPatchedEventMatcherPolicyRequest instantiates a new PatchedEventMatcherPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEventMatcherPolicyRequestWithDefaults

`func NewPatchedEventMatcherPolicyRequestWithDefaults() *PatchedEventMatcherPolicyRequest`

NewPatchedEventMatcherPolicyRequestWithDefaults instantiates a new PatchedEventMatcherPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedEventMatcherPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEventMatcherPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEventMatcherPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEventMatcherPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExecutionLogging

`func (o *PatchedEventMatcherPolicyRequest) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *PatchedEventMatcherPolicyRequest) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *PatchedEventMatcherPolicyRequest) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *PatchedEventMatcherPolicyRequest) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetAction

`func (o *PatchedEventMatcherPolicyRequest) GetAction() EventActions`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PatchedEventMatcherPolicyRequest) GetActionOk() (*EventActions, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PatchedEventMatcherPolicyRequest) SetAction(v EventActions)`

SetAction sets Action field to given value.

### HasAction

`func (o *PatchedEventMatcherPolicyRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *PatchedEventMatcherPolicyRequest) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *PatchedEventMatcherPolicyRequest) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetClientIp

`func (o *PatchedEventMatcherPolicyRequest) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *PatchedEventMatcherPolicyRequest) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *PatchedEventMatcherPolicyRequest) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *PatchedEventMatcherPolicyRequest) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetApp

`func (o *PatchedEventMatcherPolicyRequest) GetApp() AppEnum`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *PatchedEventMatcherPolicyRequest) GetAppOk() (*AppEnum, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *PatchedEventMatcherPolicyRequest) SetApp(v AppEnum)`

SetApp sets App field to given value.

### HasApp

`func (o *PatchedEventMatcherPolicyRequest) HasApp() bool`

HasApp returns a boolean if a field has been set.

### SetAppNil

`func (o *PatchedEventMatcherPolicyRequest) SetAppNil(b bool)`

 SetAppNil sets the value for App to be an explicit nil

### UnsetApp
`func (o *PatchedEventMatcherPolicyRequest) UnsetApp()`

UnsetApp ensures that no value is present for App, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


