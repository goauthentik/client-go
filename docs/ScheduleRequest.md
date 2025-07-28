# ScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelObjId** | Pointer to **NullableString** |  | [optional] 
**Crontab** | **string** | When to schedule tasks | 
**Paused** | Pointer to **bool** | Pause this schedule | [optional] 

## Methods

### NewScheduleRequest

`func NewScheduleRequest(crontab string, ) *ScheduleRequest`

NewScheduleRequest instantiates a new ScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleRequestWithDefaults

`func NewScheduleRequestWithDefaults() *ScheduleRequest`

NewScheduleRequestWithDefaults instantiates a new ScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelObjId

`func (o *ScheduleRequest) GetRelObjId() string`

GetRelObjId returns the RelObjId field if non-nil, zero value otherwise.

### GetRelObjIdOk

`func (o *ScheduleRequest) GetRelObjIdOk() (*string, bool)`

GetRelObjIdOk returns a tuple with the RelObjId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelObjId

`func (o *ScheduleRequest) SetRelObjId(v string)`

SetRelObjId sets RelObjId field to given value.

### HasRelObjId

`func (o *ScheduleRequest) HasRelObjId() bool`

HasRelObjId returns a boolean if a field has been set.

### SetRelObjIdNil

`func (o *ScheduleRequest) SetRelObjIdNil(b bool)`

 SetRelObjIdNil sets the value for RelObjId to be an explicit nil

### UnsetRelObjId
`func (o *ScheduleRequest) UnsetRelObjId()`

UnsetRelObjId ensures that no value is present for RelObjId, not even an explicit nil
### GetCrontab

`func (o *ScheduleRequest) GetCrontab() string`

GetCrontab returns the Crontab field if non-nil, zero value otherwise.

### GetCrontabOk

`func (o *ScheduleRequest) GetCrontabOk() (*string, bool)`

GetCrontabOk returns a tuple with the Crontab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrontab

`func (o *ScheduleRequest) SetCrontab(v string)`

SetCrontab sets Crontab field to given value.


### GetPaused

`func (o *ScheduleRequest) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *ScheduleRequest) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *ScheduleRequest) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *ScheduleRequest) HasPaused() bool`

HasPaused returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


