# BlueprintImportResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | [**[]LogEvent**](LogEvent.md) |  | [readonly] 
**Success** | **bool** |  | [readonly] 

## Methods

### NewBlueprintImportResult

`func NewBlueprintImportResult(logs []LogEvent, success bool, ) *BlueprintImportResult`

NewBlueprintImportResult instantiates a new BlueprintImportResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintImportResultWithDefaults

`func NewBlueprintImportResultWithDefaults() *BlueprintImportResult`

NewBlueprintImportResultWithDefaults instantiates a new BlueprintImportResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *BlueprintImportResult) GetLogs() []LogEvent`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *BlueprintImportResult) GetLogsOk() (*[]LogEvent, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *BlueprintImportResult) SetLogs(v []LogEvent)`

SetLogs sets Logs field to given value.


### GetSuccess

`func (o *BlueprintImportResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *BlueprintImportResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *BlueprintImportResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


