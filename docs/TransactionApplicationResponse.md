# TransactionApplicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applied** | **bool** |  | 
**Logs** | **[]string** |  | 

## Methods

### NewTransactionApplicationResponse

`func NewTransactionApplicationResponse(applied bool, logs []string, ) *TransactionApplicationResponse`

NewTransactionApplicationResponse instantiates a new TransactionApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionApplicationResponseWithDefaults

`func NewTransactionApplicationResponseWithDefaults() *TransactionApplicationResponse`

NewTransactionApplicationResponseWithDefaults instantiates a new TransactionApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplied

`func (o *TransactionApplicationResponse) GetApplied() bool`

GetApplied returns the Applied field if non-nil, zero value otherwise.

### GetAppliedOk

`func (o *TransactionApplicationResponse) GetAppliedOk() (*bool, bool)`

GetAppliedOk returns a tuple with the Applied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplied

`func (o *TransactionApplicationResponse) SetApplied(v bool)`

SetApplied sets Applied field to given value.


### GetLogs

`func (o *TransactionApplicationResponse) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *TransactionApplicationResponse) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *TransactionApplicationResponse) SetLogs(v []string)`

SetLogs sets Logs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


