# NotificationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Transports** | Pointer to **[]string** | Select which transports should be used to notify the user. If none are selected, the notification will only be shown in the authentik UI. | [optional] 
**Severity** | Pointer to [**SeverityEnum**](SeverityEnum.md) |  | [optional] 
**Group** | Pointer to **NullableString** | Define which group of users this notification should be sent and shown to. If left empty, Notification won&#39;t ben sent. | [optional] 
**GroupObj** | [**Group**](Group.md) |  | [readonly] 

## Methods

### NewNotificationRule

`func NewNotificationRule(pk string, name string, groupObj Group, ) *NotificationRule`

NewNotificationRule instantiates a new NotificationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleWithDefaults

`func NewNotificationRuleWithDefaults() *NotificationRule`

NewNotificationRuleWithDefaults instantiates a new NotificationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *NotificationRule) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *NotificationRule) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *NotificationRule) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *NotificationRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRule) SetName(v string)`

SetName sets Name field to given value.


### GetTransports

`func (o *NotificationRule) GetTransports() []string`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *NotificationRule) GetTransportsOk() (*[]string, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *NotificationRule) SetTransports(v []string)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *NotificationRule) HasTransports() bool`

HasTransports returns a boolean if a field has been set.

### GetSeverity

`func (o *NotificationRule) GetSeverity() SeverityEnum`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *NotificationRule) GetSeverityOk() (*SeverityEnum, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *NotificationRule) SetSeverity(v SeverityEnum)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *NotificationRule) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetGroup

`func (o *NotificationRule) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *NotificationRule) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *NotificationRule) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *NotificationRule) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *NotificationRule) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *NotificationRule) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetGroupObj

`func (o *NotificationRule) GetGroupObj() Group`

GetGroupObj returns the GroupObj field if non-nil, zero value otherwise.

### GetGroupObjOk

`func (o *NotificationRule) GetGroupObjOk() (*Group, bool)`

GetGroupObjOk returns a tuple with the GroupObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObj

`func (o *NotificationRule) SetGroupObj(v Group)`

SetGroupObj sets GroupObj field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


