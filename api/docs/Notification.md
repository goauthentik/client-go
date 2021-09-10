# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Severity** | [**SeverityEnum**](SeverityEnum.md) |  | [readonly] 
**Body** | **string** |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Event** | Pointer to [**Event**](Event.md) |  | [optional] 
**Seen** | Pointer to **bool** |  | [optional] 

## Methods

### NewNotification

`func NewNotification(pk string, severity SeverityEnum, body string, created time.Time, ) *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *Notification) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *Notification) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *Notification) SetPk(v string)`

SetPk sets Pk field to given value.


### GetSeverity

`func (o *Notification) GetSeverity() SeverityEnum`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Notification) GetSeverityOk() (*SeverityEnum, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Notification) SetSeverity(v SeverityEnum)`

SetSeverity sets Severity field to given value.


### GetBody

`func (o *Notification) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Notification) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Notification) SetBody(v string)`

SetBody sets Body field to given value.


### GetCreated

`func (o *Notification) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Notification) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Notification) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetEvent

`func (o *Notification) GetEvent() Event`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *Notification) GetEventOk() (*Event, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *Notification) SetEvent(v Event)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *Notification) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetSeen

`func (o *Notification) GetSeen() bool`

GetSeen returns the Seen field if non-nil, zero value otherwise.

### GetSeenOk

`func (o *Notification) GetSeenOk() (*bool, bool)`

GetSeenOk returns a tuple with the Seen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeen

`func (o *Notification) SetSeen(v bool)`

SetSeen sets Seen field to given value.

### HasSeen

`func (o *Notification) HasSeen() bool`

HasSeen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


