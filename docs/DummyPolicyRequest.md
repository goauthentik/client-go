# DummyPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**Result** | Pointer to **bool** |  | [optional] 
**WaitMin** | Pointer to **int32** |  | [optional] 
**WaitMax** | Pointer to **int32** |  | [optional] 

## Methods

### NewDummyPolicyRequest

`func NewDummyPolicyRequest(name string, ) *DummyPolicyRequest`

NewDummyPolicyRequest instantiates a new DummyPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDummyPolicyRequestWithDefaults

`func NewDummyPolicyRequestWithDefaults() *DummyPolicyRequest`

NewDummyPolicyRequestWithDefaults instantiates a new DummyPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DummyPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DummyPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DummyPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExecutionLogging

`func (o *DummyPolicyRequest) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *DummyPolicyRequest) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *DummyPolicyRequest) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *DummyPolicyRequest) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetResult

`func (o *DummyPolicyRequest) GetResult() bool`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DummyPolicyRequest) GetResultOk() (*bool, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DummyPolicyRequest) SetResult(v bool)`

SetResult sets Result field to given value.

### HasResult

`func (o *DummyPolicyRequest) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetWaitMin

`func (o *DummyPolicyRequest) GetWaitMin() int32`

GetWaitMin returns the WaitMin field if non-nil, zero value otherwise.

### GetWaitMinOk

`func (o *DummyPolicyRequest) GetWaitMinOk() (*int32, bool)`

GetWaitMinOk returns a tuple with the WaitMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitMin

`func (o *DummyPolicyRequest) SetWaitMin(v int32)`

SetWaitMin sets WaitMin field to given value.

### HasWaitMin

`func (o *DummyPolicyRequest) HasWaitMin() bool`

HasWaitMin returns a boolean if a field has been set.

### GetWaitMax

`func (o *DummyPolicyRequest) GetWaitMax() int32`

GetWaitMax returns the WaitMax field if non-nil, zero value otherwise.

### GetWaitMaxOk

`func (o *DummyPolicyRequest) GetWaitMaxOk() (*int32, bool)`

GetWaitMaxOk returns a tuple with the WaitMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitMax

`func (o *DummyPolicyRequest) SetWaitMax(v int32)`

SetWaitMax sets WaitMax field to given value.

### HasWaitMax

`func (o *DummyPolicyRequest) HasWaitMax() bool`

HasWaitMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


