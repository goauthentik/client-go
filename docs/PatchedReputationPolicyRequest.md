# PatchedReputationPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**CheckIp** | Pointer to **bool** |  | [optional] 
**CheckUsername** | Pointer to **bool** |  | [optional] 
**Threshold** | Pointer to **int32** |  | [optional] 

## Methods

### NewPatchedReputationPolicyRequest

`func NewPatchedReputationPolicyRequest() *PatchedReputationPolicyRequest`

NewPatchedReputationPolicyRequest instantiates a new PatchedReputationPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedReputationPolicyRequestWithDefaults

`func NewPatchedReputationPolicyRequestWithDefaults() *PatchedReputationPolicyRequest`

NewPatchedReputationPolicyRequestWithDefaults instantiates a new PatchedReputationPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedReputationPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedReputationPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedReputationPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedReputationPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExecutionLogging

`func (o *PatchedReputationPolicyRequest) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *PatchedReputationPolicyRequest) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *PatchedReputationPolicyRequest) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *PatchedReputationPolicyRequest) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetCheckIp

`func (o *PatchedReputationPolicyRequest) GetCheckIp() bool`

GetCheckIp returns the CheckIp field if non-nil, zero value otherwise.

### GetCheckIpOk

`func (o *PatchedReputationPolicyRequest) GetCheckIpOk() (*bool, bool)`

GetCheckIpOk returns a tuple with the CheckIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIp

`func (o *PatchedReputationPolicyRequest) SetCheckIp(v bool)`

SetCheckIp sets CheckIp field to given value.

### HasCheckIp

`func (o *PatchedReputationPolicyRequest) HasCheckIp() bool`

HasCheckIp returns a boolean if a field has been set.

### GetCheckUsername

`func (o *PatchedReputationPolicyRequest) GetCheckUsername() bool`

GetCheckUsername returns the CheckUsername field if non-nil, zero value otherwise.

### GetCheckUsernameOk

`func (o *PatchedReputationPolicyRequest) GetCheckUsernameOk() (*bool, bool)`

GetCheckUsernameOk returns a tuple with the CheckUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUsername

`func (o *PatchedReputationPolicyRequest) SetCheckUsername(v bool)`

SetCheckUsername sets CheckUsername field to given value.

### HasCheckUsername

`func (o *PatchedReputationPolicyRequest) HasCheckUsername() bool`

HasCheckUsername returns a boolean if a field has been set.

### GetThreshold

`func (o *PatchedReputationPolicyRequest) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *PatchedReputationPolicyRequest) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *PatchedReputationPolicyRequest) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *PatchedReputationPolicyRequest) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


