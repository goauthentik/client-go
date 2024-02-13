# SystemTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | [readonly] 
**Name** | **string** |  | 
**FullName** | **string** | Get full name with UID | [readonly] 
**Uid** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**StartTimestamp** | **time.Time** |  | [readonly] 
**FinishTimestamp** | **time.Time** |  | [readonly] 
**Duration** | **float64** |  | [readonly] 
**Status** | [**SystemTaskStatusEnum**](SystemTaskStatusEnum.md) |  | 
**Messages** | **[]string** |  | 

## Methods

### NewSystemTask

`func NewSystemTask(uuid string, name string, fullName string, description string, startTimestamp time.Time, finishTimestamp time.Time, duration float64, status SystemTaskStatusEnum, messages []string, ) *SystemTask`

NewSystemTask instantiates a new SystemTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemTaskWithDefaults

`func NewSystemTaskWithDefaults() *SystemTask`

NewSystemTaskWithDefaults instantiates a new SystemTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *SystemTask) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SystemTask) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SystemTask) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *SystemTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemTask) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *SystemTask) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *SystemTask) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *SystemTask) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetUid

`func (o *SystemTask) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SystemTask) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SystemTask) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SystemTask) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDescription

`func (o *SystemTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemTask) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStartTimestamp

`func (o *SystemTask) GetStartTimestamp() time.Time`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *SystemTask) GetStartTimestampOk() (*time.Time, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *SystemTask) SetStartTimestamp(v time.Time)`

SetStartTimestamp sets StartTimestamp field to given value.


### GetFinishTimestamp

`func (o *SystemTask) GetFinishTimestamp() time.Time`

GetFinishTimestamp returns the FinishTimestamp field if non-nil, zero value otherwise.

### GetFinishTimestampOk

`func (o *SystemTask) GetFinishTimestampOk() (*time.Time, bool)`

GetFinishTimestampOk returns a tuple with the FinishTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTimestamp

`func (o *SystemTask) SetFinishTimestamp(v time.Time)`

SetFinishTimestamp sets FinishTimestamp field to given value.


### GetDuration

`func (o *SystemTask) GetDuration() float64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SystemTask) GetDurationOk() (*float64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SystemTask) SetDuration(v float64)`

SetDuration sets Duration field to given value.


### GetStatus

`func (o *SystemTask) GetStatus() SystemTaskStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SystemTask) GetStatusOk() (*SystemTaskStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SystemTask) SetStatus(v SystemTaskStatusEnum)`

SetStatus sets Status field to given value.


### GetMessages

`func (o *SystemTask) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *SystemTask) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *SystemTask) SetMessages(v []string)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


