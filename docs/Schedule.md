# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Identifier** | **NullableString** | Unique schedule identifier | [readonly] 
**Uid** | **string** |  | [readonly] 
**ActorName** | **string** | Dramatiq actor to call | [readonly] 
**RelObjAppLabel** | **string** |  | [readonly] 
**RelObjModel** | **string** |  | [readonly] 
**RelObjId** | Pointer to **NullableString** |  | [optional] 
**Crontab** | **string** | When to schedule tasks | 
**Paused** | Pointer to **bool** | Pause this schedule | [optional] 
**NextRun** | **time.Time** |  | [readonly] 
**Description** | **NullableString** |  | [readonly] 
**LastTaskStatus** | [**NullableLastTaskStatusEnum**](LastTaskStatusEnum.md) |  | [readonly] 

## Methods

### NewSchedule

`func NewSchedule(id string, identifier NullableString, uid string, actorName string, relObjAppLabel string, relObjModel string, crontab string, nextRun time.Time, description NullableString, lastTaskStatus NullableLastTaskStatusEnum, ) *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Schedule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Schedule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Schedule) SetId(v string)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *Schedule) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Schedule) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Schedule) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### SetIdentifierNil

`func (o *Schedule) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *Schedule) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetUid

`func (o *Schedule) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Schedule) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Schedule) SetUid(v string)`

SetUid sets Uid field to given value.


### GetActorName

`func (o *Schedule) GetActorName() string`

GetActorName returns the ActorName field if non-nil, zero value otherwise.

### GetActorNameOk

`func (o *Schedule) GetActorNameOk() (*string, bool)`

GetActorNameOk returns a tuple with the ActorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorName

`func (o *Schedule) SetActorName(v string)`

SetActorName sets ActorName field to given value.


### GetRelObjAppLabel

`func (o *Schedule) GetRelObjAppLabel() string`

GetRelObjAppLabel returns the RelObjAppLabel field if non-nil, zero value otherwise.

### GetRelObjAppLabelOk

`func (o *Schedule) GetRelObjAppLabelOk() (*string, bool)`

GetRelObjAppLabelOk returns a tuple with the RelObjAppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelObjAppLabel

`func (o *Schedule) SetRelObjAppLabel(v string)`

SetRelObjAppLabel sets RelObjAppLabel field to given value.


### GetRelObjModel

`func (o *Schedule) GetRelObjModel() string`

GetRelObjModel returns the RelObjModel field if non-nil, zero value otherwise.

### GetRelObjModelOk

`func (o *Schedule) GetRelObjModelOk() (*string, bool)`

GetRelObjModelOk returns a tuple with the RelObjModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelObjModel

`func (o *Schedule) SetRelObjModel(v string)`

SetRelObjModel sets RelObjModel field to given value.


### GetRelObjId

`func (o *Schedule) GetRelObjId() string`

GetRelObjId returns the RelObjId field if non-nil, zero value otherwise.

### GetRelObjIdOk

`func (o *Schedule) GetRelObjIdOk() (*string, bool)`

GetRelObjIdOk returns a tuple with the RelObjId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelObjId

`func (o *Schedule) SetRelObjId(v string)`

SetRelObjId sets RelObjId field to given value.

### HasRelObjId

`func (o *Schedule) HasRelObjId() bool`

HasRelObjId returns a boolean if a field has been set.

### SetRelObjIdNil

`func (o *Schedule) SetRelObjIdNil(b bool)`

 SetRelObjIdNil sets the value for RelObjId to be an explicit nil

### UnsetRelObjId
`func (o *Schedule) UnsetRelObjId()`

UnsetRelObjId ensures that no value is present for RelObjId, not even an explicit nil
### GetCrontab

`func (o *Schedule) GetCrontab() string`

GetCrontab returns the Crontab field if non-nil, zero value otherwise.

### GetCrontabOk

`func (o *Schedule) GetCrontabOk() (*string, bool)`

GetCrontabOk returns a tuple with the Crontab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrontab

`func (o *Schedule) SetCrontab(v string)`

SetCrontab sets Crontab field to given value.


### GetPaused

`func (o *Schedule) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *Schedule) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *Schedule) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *Schedule) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetNextRun

`func (o *Schedule) GetNextRun() time.Time`

GetNextRun returns the NextRun field if non-nil, zero value otherwise.

### GetNextRunOk

`func (o *Schedule) GetNextRunOk() (*time.Time, bool)`

GetNextRunOk returns a tuple with the NextRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRun

`func (o *Schedule) SetNextRun(v time.Time)`

SetNextRun sets NextRun field to given value.


### GetDescription

`func (o *Schedule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Schedule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Schedule) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Schedule) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Schedule) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLastTaskStatus

`func (o *Schedule) GetLastTaskStatus() LastTaskStatusEnum`

GetLastTaskStatus returns the LastTaskStatus field if non-nil, zero value otherwise.

### GetLastTaskStatusOk

`func (o *Schedule) GetLastTaskStatusOk() (*LastTaskStatusEnum, bool)`

GetLastTaskStatusOk returns a tuple with the LastTaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTaskStatus

`func (o *Schedule) SetLastTaskStatus(v LastTaskStatusEnum)`

SetLastTaskStatus sets LastTaskStatus field to given value.


### SetLastTaskStatusNil

`func (o *Schedule) SetLastTaskStatusNil(b bool)`

 SetLastTaskStatusNil sets the value for LastTaskStatus to be an explicit nil

### UnsetLastTaskStatus
`func (o *Schedule) UnsetLastTaskStatus()`

UnsetLastTaskStatus ensures that no value is present for LastTaskStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


