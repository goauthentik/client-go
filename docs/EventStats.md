# EventStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueUsers** | **int32** |  | 
**CountStep** | **map[string]interface{}** |  | 

## Methods

### NewEventStats

`func NewEventStats(uniqueUsers int32, countStep map[string]interface{}, ) *EventStats`

NewEventStats instantiates a new EventStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventStatsWithDefaults

`func NewEventStatsWithDefaults() *EventStats`

NewEventStatsWithDefaults instantiates a new EventStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueUsers

`func (o *EventStats) GetUniqueUsers() int32`

GetUniqueUsers returns the UniqueUsers field if non-nil, zero value otherwise.

### GetUniqueUsersOk

`func (o *EventStats) GetUniqueUsersOk() (*int32, bool)`

GetUniqueUsersOk returns a tuple with the UniqueUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueUsers

`func (o *EventStats) SetUniqueUsers(v int32)`

SetUniqueUsers sets UniqueUsers field to given value.


### GetCountStep

`func (o *EventStats) GetCountStep() map[string]interface{}`

GetCountStep returns the CountStep field if non-nil, zero value otherwise.

### GetCountStepOk

`func (o *EventStats) GetCountStepOk() (*map[string]interface{}, bool)`

GetCountStepOk returns a tuple with the CountStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountStep

`func (o *EventStats) SetCountStep(v map[string]interface{})`

SetCountStep sets CountStep field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


