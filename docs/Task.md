# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** |  | [optional] 
**QueueName** | Pointer to **string** | Queue name | [optional] 
**ActorName** | **string** | Dramatiq actor name | 
**State** | Pointer to [**StateEnum**](StateEnum.md) | Task status | [optional] 
**Mtime** | Pointer to **time.Time** | Task last modified time | [optional] 
**RelObjAppLabel** | **string** |  | [readonly] 
**RelObjModel** | **string** |  | [readonly] 
**RelObjId** | Pointer to **NullableString** |  | [optional] 
**Uid** | **string** |  | [readonly] 
**Messages** | [**[]LogEvent**](LogEvent.md) |  | 
**PreviousMessages** | [**[]LogEvent**](LogEvent.md) |  | 
**AggregatedStatus** | [**TaskAggregatedStatusEnum**](TaskAggregatedStatusEnum.md) |  | 
**Description** | **NullableString** |  | [readonly] 

## Methods

### NewTask

`func NewTask(actorName string, relObjAppLabel string, relObjModel string, uid string, messages []LogEvent, previousMessages []LogEvent, aggregatedStatus TaskAggregatedStatusEnum, description NullableString, ) *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *Task) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *Task) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *Task) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *Task) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetQueueName

`func (o *Task) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *Task) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *Task) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.

### HasQueueName

`func (o *Task) HasQueueName() bool`

HasQueueName returns a boolean if a field has been set.

### GetActorName

`func (o *Task) GetActorName() string`

GetActorName returns the ActorName field if non-nil, zero value otherwise.

### GetActorNameOk

`func (o *Task) GetActorNameOk() (*string, bool)`

GetActorNameOk returns a tuple with the ActorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorName

`func (o *Task) SetActorName(v string)`

SetActorName sets ActorName field to given value.


### GetState

`func (o *Task) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Task) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Task) SetState(v StateEnum)`

SetState sets State field to given value.

### HasState

`func (o *Task) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMtime

`func (o *Task) GetMtime() time.Time`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *Task) GetMtimeOk() (*time.Time, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *Task) SetMtime(v time.Time)`

SetMtime sets Mtime field to given value.

### HasMtime

`func (o *Task) HasMtime() bool`

HasMtime returns a boolean if a field has been set.

### GetRelObjAppLabel

`func (o *Task) GetRelObjAppLabel() string`

GetRelObjAppLabel returns the RelObjAppLabel field if non-nil, zero value otherwise.

### GetRelObjAppLabelOk

`func (o *Task) GetRelObjAppLabelOk() (*string, bool)`

GetRelObjAppLabelOk returns a tuple with the RelObjAppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelObjAppLabel

`func (o *Task) SetRelObjAppLabel(v string)`

SetRelObjAppLabel sets RelObjAppLabel field to given value.


### GetRelObjModel

`func (o *Task) GetRelObjModel() string`

GetRelObjModel returns the RelObjModel field if non-nil, zero value otherwise.

### GetRelObjModelOk

`func (o *Task) GetRelObjModelOk() (*string, bool)`

GetRelObjModelOk returns a tuple with the RelObjModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelObjModel

`func (o *Task) SetRelObjModel(v string)`

SetRelObjModel sets RelObjModel field to given value.


### GetRelObjId

`func (o *Task) GetRelObjId() string`

GetRelObjId returns the RelObjId field if non-nil, zero value otherwise.

### GetRelObjIdOk

`func (o *Task) GetRelObjIdOk() (*string, bool)`

GetRelObjIdOk returns a tuple with the RelObjId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelObjId

`func (o *Task) SetRelObjId(v string)`

SetRelObjId sets RelObjId field to given value.

### HasRelObjId

`func (o *Task) HasRelObjId() bool`

HasRelObjId returns a boolean if a field has been set.

### SetRelObjIdNil

`func (o *Task) SetRelObjIdNil(b bool)`

 SetRelObjIdNil sets the value for RelObjId to be an explicit nil

### UnsetRelObjId
`func (o *Task) UnsetRelObjId()`

UnsetRelObjId ensures that no value is present for RelObjId, not even an explicit nil
### GetUid

`func (o *Task) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Task) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Task) SetUid(v string)`

SetUid sets Uid field to given value.


### GetMessages

`func (o *Task) GetMessages() []LogEvent`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *Task) GetMessagesOk() (*[]LogEvent, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *Task) SetMessages(v []LogEvent)`

SetMessages sets Messages field to given value.


### GetPreviousMessages

`func (o *Task) GetPreviousMessages() []LogEvent`

GetPreviousMessages returns the PreviousMessages field if non-nil, zero value otherwise.

### GetPreviousMessagesOk

`func (o *Task) GetPreviousMessagesOk() (*[]LogEvent, bool)`

GetPreviousMessagesOk returns a tuple with the PreviousMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousMessages

`func (o *Task) SetPreviousMessages(v []LogEvent)`

SetPreviousMessages sets PreviousMessages field to given value.


### GetAggregatedStatus

`func (o *Task) GetAggregatedStatus() TaskAggregatedStatusEnum`

GetAggregatedStatus returns the AggregatedStatus field if non-nil, zero value otherwise.

### GetAggregatedStatusOk

`func (o *Task) GetAggregatedStatusOk() (*TaskAggregatedStatusEnum, bool)`

GetAggregatedStatusOk returns a tuple with the AggregatedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedStatus

`func (o *Task) SetAggregatedStatus(v TaskAggregatedStatusEnum)`

SetAggregatedStatus sets AggregatedStatus field to given value.


### GetDescription

`func (o *Task) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Task) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Task) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Task) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Task) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


