# EventTopPerUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | **map[string]interface{}** |  | 
**CountedEvents** | **int32** |  | 
**UniqueUsers** | **int32** |  | 

## Methods

### NewEventTopPerUser

`func NewEventTopPerUser(application map[string]interface{}, countedEvents int32, uniqueUsers int32, ) *EventTopPerUser`

NewEventTopPerUser instantiates a new EventTopPerUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTopPerUserWithDefaults

`func NewEventTopPerUserWithDefaults() *EventTopPerUser`

NewEventTopPerUserWithDefaults instantiates a new EventTopPerUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *EventTopPerUser) GetApplication() map[string]interface{}`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *EventTopPerUser) GetApplicationOk() (*map[string]interface{}, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *EventTopPerUser) SetApplication(v map[string]interface{})`

SetApplication sets Application field to given value.


### GetCountedEvents

`func (o *EventTopPerUser) GetCountedEvents() int32`

GetCountedEvents returns the CountedEvents field if non-nil, zero value otherwise.

### GetCountedEventsOk

`func (o *EventTopPerUser) GetCountedEventsOk() (*int32, bool)`

GetCountedEventsOk returns a tuple with the CountedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountedEvents

`func (o *EventTopPerUser) SetCountedEvents(v int32)`

SetCountedEvents sets CountedEvents field to given value.


### GetUniqueUsers

`func (o *EventTopPerUser) GetUniqueUsers() int32`

GetUniqueUsers returns the UniqueUsers field if non-nil, zero value otherwise.

### GetUniqueUsersOk

`func (o *EventTopPerUser) GetUniqueUsersOk() (*int32, bool)`

GetUniqueUsersOk returns a tuple with the UniqueUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueUsers

`func (o *EventTopPerUser) SetUniqueUsers(v int32)`

SetUniqueUsers sets UniqueUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


