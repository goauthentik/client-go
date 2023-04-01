# PatchedNotificationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Transports** | Pointer to **[]string** | Select which transports should be used to notify the user. If none are selected, the notification will only be shown in the authentik UI. | [optional] 
**Severity** | Pointer to [**SeverityEnum**](SeverityEnum.md) |  | [optional] 
**Group** | Pointer to **NullableString** | Define which group of users this notification should be sent and shown to. If left empty, Notification won&#39;t ben sent. | [optional] 

## Methods

### NewPatchedNotificationRuleRequest

`func NewPatchedNotificationRuleRequest() *PatchedNotificationRuleRequest`

NewPatchedNotificationRuleRequest instantiates a new PatchedNotificationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedNotificationRuleRequestWithDefaults

`func NewPatchedNotificationRuleRequestWithDefaults() *PatchedNotificationRuleRequest`

NewPatchedNotificationRuleRequestWithDefaults instantiates a new PatchedNotificationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedNotificationRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedNotificationRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedNotificationRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedNotificationRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTransports

`func (o *PatchedNotificationRuleRequest) GetTransports() []string`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *PatchedNotificationRuleRequest) GetTransportsOk() (*[]string, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *PatchedNotificationRuleRequest) SetTransports(v []string)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *PatchedNotificationRuleRequest) HasTransports() bool`

HasTransports returns a boolean if a field has been set.

### GetSeverity

`func (o *PatchedNotificationRuleRequest) GetSeverity() SeverityEnum`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *PatchedNotificationRuleRequest) GetSeverityOk() (*SeverityEnum, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *PatchedNotificationRuleRequest) SetSeverity(v SeverityEnum)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *PatchedNotificationRuleRequest) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetGroup

`func (o *PatchedNotificationRuleRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedNotificationRuleRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedNotificationRuleRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedNotificationRuleRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedNotificationRuleRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedNotificationRuleRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


