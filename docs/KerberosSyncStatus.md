# KerberosSyncStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRunning** | **bool** |  | [readonly] 
**Tasks** | [**[]SystemTask**](SystemTask.md) |  | [readonly] 

## Methods

### NewKerberosSyncStatus

`func NewKerberosSyncStatus(isRunning bool, tasks []SystemTask, ) *KerberosSyncStatus`

NewKerberosSyncStatus instantiates a new KerberosSyncStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKerberosSyncStatusWithDefaults

`func NewKerberosSyncStatusWithDefaults() *KerberosSyncStatus`

NewKerberosSyncStatusWithDefaults instantiates a new KerberosSyncStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRunning

`func (o *KerberosSyncStatus) GetIsRunning() bool`

GetIsRunning returns the IsRunning field if non-nil, zero value otherwise.

### GetIsRunningOk

`func (o *KerberosSyncStatus) GetIsRunningOk() (*bool, bool)`

GetIsRunningOk returns a tuple with the IsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRunning

`func (o *KerberosSyncStatus) SetIsRunning(v bool)`

SetIsRunning sets IsRunning field to given value.


### GetTasks

`func (o *KerberosSyncStatus) GetTasks() []SystemTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *KerberosSyncStatus) GetTasksOk() (*[]SystemTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *KerberosSyncStatus) SetTasks(v []SystemTask)`

SetTasks sets Tasks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


