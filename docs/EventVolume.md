# EventVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**EventActions**](EventActions.md) |  | 
**Time** | **time.Time** |  | 
**Count** | **int32** |  | 

## Methods

### NewEventVolume

`func NewEventVolume(action EventActions, time time.Time, count int32, ) *EventVolume`

NewEventVolume instantiates a new EventVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventVolumeWithDefaults

`func NewEventVolumeWithDefaults() *EventVolume`

NewEventVolumeWithDefaults instantiates a new EventVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *EventVolume) GetAction() EventActions`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *EventVolume) GetActionOk() (*EventActions, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *EventVolume) SetAction(v EventActions)`

SetAction sets Action field to given value.


### GetTime

`func (o *EventVolume) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *EventVolume) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *EventVolume) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetCount

`func (o *EventVolume) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EventVolume) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EventVolume) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


