# ReputationPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**CheckIp** | Pointer to **bool** |  | [optional] 
**CheckUsername** | Pointer to **bool** |  | [optional] 
**Threshold** | Pointer to **int32** |  | [optional] 

## Methods

### NewReputationPolicyRequest

`func NewReputationPolicyRequest(name string, ) *ReputationPolicyRequest`

NewReputationPolicyRequest instantiates a new ReputationPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReputationPolicyRequestWithDefaults

`func NewReputationPolicyRequestWithDefaults() *ReputationPolicyRequest`

NewReputationPolicyRequestWithDefaults instantiates a new ReputationPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReputationPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReputationPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReputationPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExecutionLogging

`func (o *ReputationPolicyRequest) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *ReputationPolicyRequest) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *ReputationPolicyRequest) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *ReputationPolicyRequest) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetCheckIp

`func (o *ReputationPolicyRequest) GetCheckIp() bool`

GetCheckIp returns the CheckIp field if non-nil, zero value otherwise.

### GetCheckIpOk

`func (o *ReputationPolicyRequest) GetCheckIpOk() (*bool, bool)`

GetCheckIpOk returns a tuple with the CheckIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIp

`func (o *ReputationPolicyRequest) SetCheckIp(v bool)`

SetCheckIp sets CheckIp field to given value.

### HasCheckIp

`func (o *ReputationPolicyRequest) HasCheckIp() bool`

HasCheckIp returns a boolean if a field has been set.

### GetCheckUsername

`func (o *ReputationPolicyRequest) GetCheckUsername() bool`

GetCheckUsername returns the CheckUsername field if non-nil, zero value otherwise.

### GetCheckUsernameOk

`func (o *ReputationPolicyRequest) GetCheckUsernameOk() (*bool, bool)`

GetCheckUsernameOk returns a tuple with the CheckUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUsername

`func (o *ReputationPolicyRequest) SetCheckUsername(v bool)`

SetCheckUsername sets CheckUsername field to given value.

### HasCheckUsername

`func (o *ReputationPolicyRequest) HasCheckUsername() bool`

HasCheckUsername returns a boolean if a field has been set.

### GetThreshold

`func (o *ReputationPolicyRequest) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ReputationPolicyRequest) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ReputationPolicyRequest) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ReputationPolicyRequest) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


