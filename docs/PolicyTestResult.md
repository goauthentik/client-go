# PolicyTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passing** | **bool** |  | 
**Messages** | **[]string** |  | [readonly] 
**LogMessages** | **[]map[string]interface{}** |  | [readonly] 

## Methods

### NewPolicyTestResult

`func NewPolicyTestResult(passing bool, messages []string, logMessages []map[string]interface{}, ) *PolicyTestResult`

NewPolicyTestResult instantiates a new PolicyTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyTestResultWithDefaults

`func NewPolicyTestResultWithDefaults() *PolicyTestResult`

NewPolicyTestResultWithDefaults instantiates a new PolicyTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassing

`func (o *PolicyTestResult) GetPassing() bool`

GetPassing returns the Passing field if non-nil, zero value otherwise.

### GetPassingOk

`func (o *PolicyTestResult) GetPassingOk() (*bool, bool)`

GetPassingOk returns a tuple with the Passing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassing

`func (o *PolicyTestResult) SetPassing(v bool)`

SetPassing sets Passing field to given value.


### GetMessages

`func (o *PolicyTestResult) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *PolicyTestResult) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *PolicyTestResult) SetMessages(v []string)`

SetMessages sets Messages field to given value.


### GetLogMessages

`func (o *PolicyTestResult) GetLogMessages() []map[string]interface{}`

GetLogMessages returns the LogMessages field if non-nil, zero value otherwise.

### GetLogMessagesOk

`func (o *PolicyTestResult) GetLogMessagesOk() (*[]map[string]interface{}, bool)`

GetLogMessagesOk returns a tuple with the LogMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMessages

`func (o *PolicyTestResult) SetLogMessages(v []map[string]interface{})`

SetLogMessages sets LogMessages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


