# NotificationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Transports** | Pointer to **[]string** | Select which transports should be used to notify the user. If none are selected, the notification will only be shown in the authentik UI. | [optional] 
**Severity** | Pointer to [**SeverityEnum**](SeverityEnum.md) |  | [optional] 
**Group** | Pointer to **NullableString** | Define which group of users this notification should be sent and shown to. If left empty, Notification won&#39;t ben sent. | [optional] 

## Methods

### NewNotificationRuleRequest

`func NewNotificationRuleRequest(name string, ) *NotificationRuleRequest`

NewNotificationRuleRequest instantiates a new NotificationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleRequestWithDefaults

`func NewNotificationRuleRequestWithDefaults() *NotificationRuleRequest`

NewNotificationRuleRequestWithDefaults instantiates a new NotificationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NotificationRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTransports

`func (o *NotificationRuleRequest) GetTransports() []string`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *NotificationRuleRequest) GetTransportsOk() (*[]string, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *NotificationRuleRequest) SetTransports(v []string)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *NotificationRuleRequest) HasTransports() bool`

HasTransports returns a boolean if a field has been set.

### GetSeverity

`func (o *NotificationRuleRequest) GetSeverity() SeverityEnum`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *NotificationRuleRequest) GetSeverityOk() (*SeverityEnum, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *NotificationRuleRequest) SetSeverity(v SeverityEnum)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *NotificationRuleRequest) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetGroup

`func (o *NotificationRuleRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *NotificationRuleRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *NotificationRuleRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *NotificationRuleRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *NotificationRuleRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *NotificationRuleRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


