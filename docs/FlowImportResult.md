# FlowImportResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | **[]map[string]interface{}** |  | [readonly] 
**Success** | **bool** |  | [readonly] 

## Methods

### NewFlowImportResult

`func NewFlowImportResult(logs []map[string]interface{}, success bool, ) *FlowImportResult`

NewFlowImportResult instantiates a new FlowImportResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowImportResultWithDefaults

`func NewFlowImportResultWithDefaults() *FlowImportResult`

NewFlowImportResultWithDefaults instantiates a new FlowImportResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *FlowImportResult) GetLogs() []map[string]interface{}`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *FlowImportResult) GetLogsOk() (*[]map[string]interface{}, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *FlowImportResult) SetLogs(v []map[string]interface{})`

SetLogs sets Logs field to given value.


### GetSuccess

`func (o *FlowImportResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *FlowImportResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *FlowImportResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


