# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskName** | **string** |  | 
**TaskDescription** | **string** |  | 
**TaskFinishTimestamp** | **time.Time** |  | 
**TaskDuration** | **int32** | Get the duration a task took to run | [readonly] 
**Status** | [**TaskStatusEnum**](TaskStatusEnum.md) |  | 
**Messages** | **[]interface{}** |  | 

## Methods

### NewTask

`func NewTask(taskName string, taskDescription string, taskFinishTimestamp time.Time, taskDuration int32, status TaskStatusEnum, messages []interface{}, ) *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskName

`func (o *Task) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *Task) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *Task) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetTaskDescription

`func (o *Task) GetTaskDescription() string`

GetTaskDescription returns the TaskDescription field if non-nil, zero value otherwise.

### GetTaskDescriptionOk

`func (o *Task) GetTaskDescriptionOk() (*string, bool)`

GetTaskDescriptionOk returns a tuple with the TaskDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDescription

`func (o *Task) SetTaskDescription(v string)`

SetTaskDescription sets TaskDescription field to given value.


### GetTaskFinishTimestamp

`func (o *Task) GetTaskFinishTimestamp() time.Time`

GetTaskFinishTimestamp returns the TaskFinishTimestamp field if non-nil, zero value otherwise.

### GetTaskFinishTimestampOk

`func (o *Task) GetTaskFinishTimestampOk() (*time.Time, bool)`

GetTaskFinishTimestampOk returns a tuple with the TaskFinishTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskFinishTimestamp

`func (o *Task) SetTaskFinishTimestamp(v time.Time)`

SetTaskFinishTimestamp sets TaskFinishTimestamp field to given value.


### GetTaskDuration

`func (o *Task) GetTaskDuration() int32`

GetTaskDuration returns the TaskDuration field if non-nil, zero value otherwise.

### GetTaskDurationOk

`func (o *Task) GetTaskDurationOk() (*int32, bool)`

GetTaskDurationOk returns a tuple with the TaskDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDuration

`func (o *Task) SetTaskDuration(v int32)`

SetTaskDuration sets TaskDuration field to given value.


### GetStatus

`func (o *Task) GetStatus() TaskStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Task) GetStatusOk() (*TaskStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Task) SetStatus(v TaskStatusEnum)`

SetStatus sets Status field to given value.


### GetMessages

`func (o *Task) GetMessages() []interface{}`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *Task) GetMessagesOk() (*[]interface{}, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *Task) SetMessages(v []interface{})`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


