# UserExpirationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**ActivityBasis** | Pointer to [**ActivityBasisEnum**](ActivityBasisEnum.md) | Measure inactivity from the last login or successful user activity. | [optional] 
**Group** | Pointer to **NullableString** | Only expire members of this group (including descendant groups). Leave empty to apply to every user. | [optional] 
**UserTypes** | Pointer to [**[]UserTypeEnum**](UserTypeEnum.md) | Only expire users of these types. Service accounts are excluded by default. | [optional] 
**InactivityDuration** | Pointer to **string** | How long a user may be inactive before they are expired. | [optional] 
**Action** | Pointer to [**OffboardingActionEnum**](OffboardingActionEnum.md) |  | [optional] 
**RevokeSessions** | Pointer to **bool** |  | [optional] 
**RevokeTokens** | Pointer to **bool** |  | [optional] 
**WarnBefore** | Pointer to **NullableString** | Schedule the offboarding and notify the user this long before they expire. Leave empty to expire users without warning. | [optional] 
**NotificationTransports** | Pointer to **[]string** | Select which transports should be used to warn the user. If none are selected, the notification will only be shown in the authentik UI. | [optional] 
**ExcludeSuperusers** | Pointer to **bool** | Never expire members of a superuser group. | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 

## Methods

### NewUserExpirationRuleRequest

`func NewUserExpirationRuleRequest(name string, ) *UserExpirationRuleRequest`

NewUserExpirationRuleRequest instantiates a new UserExpirationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserExpirationRuleRequestWithDefaults

`func NewUserExpirationRuleRequestWithDefaults() *UserExpirationRuleRequest`

NewUserExpirationRuleRequestWithDefaults instantiates a new UserExpirationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *UserExpirationRuleRequest) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *UserExpirationRuleRequest) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *UserExpirationRuleRequest) SetPk(v string)`

SetPk sets Pk field to given value.

### HasPk

`func (o *UserExpirationRuleRequest) HasPk() bool`

HasPk returns a boolean if a field has been set.

### GetName

`func (o *UserExpirationRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserExpirationRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserExpirationRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *UserExpirationRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserExpirationRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserExpirationRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserExpirationRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetActivityBasis

`func (o *UserExpirationRuleRequest) GetActivityBasis() ActivityBasisEnum`

GetActivityBasis returns the ActivityBasis field if non-nil, zero value otherwise.

### GetActivityBasisOk

`func (o *UserExpirationRuleRequest) GetActivityBasisOk() (*ActivityBasisEnum, bool)`

GetActivityBasisOk returns a tuple with the ActivityBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityBasis

`func (o *UserExpirationRuleRequest) SetActivityBasis(v ActivityBasisEnum)`

SetActivityBasis sets ActivityBasis field to given value.

### HasActivityBasis

`func (o *UserExpirationRuleRequest) HasActivityBasis() bool`

HasActivityBasis returns a boolean if a field has been set.

### GetGroup

`func (o *UserExpirationRuleRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UserExpirationRuleRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UserExpirationRuleRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UserExpirationRuleRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *UserExpirationRuleRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *UserExpirationRuleRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetUserTypes

`func (o *UserExpirationRuleRequest) GetUserTypes() []UserTypeEnum`

GetUserTypes returns the UserTypes field if non-nil, zero value otherwise.

### GetUserTypesOk

`func (o *UserExpirationRuleRequest) GetUserTypesOk() (*[]UserTypeEnum, bool)`

GetUserTypesOk returns a tuple with the UserTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTypes

`func (o *UserExpirationRuleRequest) SetUserTypes(v []UserTypeEnum)`

SetUserTypes sets UserTypes field to given value.

### HasUserTypes

`func (o *UserExpirationRuleRequest) HasUserTypes() bool`

HasUserTypes returns a boolean if a field has been set.

### GetInactivityDuration

`func (o *UserExpirationRuleRequest) GetInactivityDuration() string`

GetInactivityDuration returns the InactivityDuration field if non-nil, zero value otherwise.

### GetInactivityDurationOk

`func (o *UserExpirationRuleRequest) GetInactivityDurationOk() (*string, bool)`

GetInactivityDurationOk returns a tuple with the InactivityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivityDuration

`func (o *UserExpirationRuleRequest) SetInactivityDuration(v string)`

SetInactivityDuration sets InactivityDuration field to given value.

### HasInactivityDuration

`func (o *UserExpirationRuleRequest) HasInactivityDuration() bool`

HasInactivityDuration returns a boolean if a field has been set.

### GetAction

`func (o *UserExpirationRuleRequest) GetAction() OffboardingActionEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UserExpirationRuleRequest) GetActionOk() (*OffboardingActionEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UserExpirationRuleRequest) SetAction(v OffboardingActionEnum)`

SetAction sets Action field to given value.

### HasAction

`func (o *UserExpirationRuleRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRevokeSessions

`func (o *UserExpirationRuleRequest) GetRevokeSessions() bool`

GetRevokeSessions returns the RevokeSessions field if non-nil, zero value otherwise.

### GetRevokeSessionsOk

`func (o *UserExpirationRuleRequest) GetRevokeSessionsOk() (*bool, bool)`

GetRevokeSessionsOk returns a tuple with the RevokeSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeSessions

`func (o *UserExpirationRuleRequest) SetRevokeSessions(v bool)`

SetRevokeSessions sets RevokeSessions field to given value.

### HasRevokeSessions

`func (o *UserExpirationRuleRequest) HasRevokeSessions() bool`

HasRevokeSessions returns a boolean if a field has been set.

### GetRevokeTokens

`func (o *UserExpirationRuleRequest) GetRevokeTokens() bool`

GetRevokeTokens returns the RevokeTokens field if non-nil, zero value otherwise.

### GetRevokeTokensOk

`func (o *UserExpirationRuleRequest) GetRevokeTokensOk() (*bool, bool)`

GetRevokeTokensOk returns a tuple with the RevokeTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeTokens

`func (o *UserExpirationRuleRequest) SetRevokeTokens(v bool)`

SetRevokeTokens sets RevokeTokens field to given value.

### HasRevokeTokens

`func (o *UserExpirationRuleRequest) HasRevokeTokens() bool`

HasRevokeTokens returns a boolean if a field has been set.

### GetWarnBefore

`func (o *UserExpirationRuleRequest) GetWarnBefore() string`

GetWarnBefore returns the WarnBefore field if non-nil, zero value otherwise.

### GetWarnBeforeOk

`func (o *UserExpirationRuleRequest) GetWarnBeforeOk() (*string, bool)`

GetWarnBeforeOk returns a tuple with the WarnBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnBefore

`func (o *UserExpirationRuleRequest) SetWarnBefore(v string)`

SetWarnBefore sets WarnBefore field to given value.

### HasWarnBefore

`func (o *UserExpirationRuleRequest) HasWarnBefore() bool`

HasWarnBefore returns a boolean if a field has been set.

### SetWarnBeforeNil

`func (o *UserExpirationRuleRequest) SetWarnBeforeNil(b bool)`

 SetWarnBeforeNil sets the value for WarnBefore to be an explicit nil

### UnsetWarnBefore
`func (o *UserExpirationRuleRequest) UnsetWarnBefore()`

UnsetWarnBefore ensures that no value is present for WarnBefore, not even an explicit nil
### GetNotificationTransports

`func (o *UserExpirationRuleRequest) GetNotificationTransports() []string`

GetNotificationTransports returns the NotificationTransports field if non-nil, zero value otherwise.

### GetNotificationTransportsOk

`func (o *UserExpirationRuleRequest) GetNotificationTransportsOk() (*[]string, bool)`

GetNotificationTransportsOk returns a tuple with the NotificationTransports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTransports

`func (o *UserExpirationRuleRequest) SetNotificationTransports(v []string)`

SetNotificationTransports sets NotificationTransports field to given value.

### HasNotificationTransports

`func (o *UserExpirationRuleRequest) HasNotificationTransports() bool`

HasNotificationTransports returns a boolean if a field has been set.

### GetExcludeSuperusers

`func (o *UserExpirationRuleRequest) GetExcludeSuperusers() bool`

GetExcludeSuperusers returns the ExcludeSuperusers field if non-nil, zero value otherwise.

### GetExcludeSuperusersOk

`func (o *UserExpirationRuleRequest) GetExcludeSuperusersOk() (*bool, bool)`

GetExcludeSuperusersOk returns a tuple with the ExcludeSuperusers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSuperusers

`func (o *UserExpirationRuleRequest) SetExcludeSuperusers(v bool)`

SetExcludeSuperusers sets ExcludeSuperusers field to given value.

### HasExcludeSuperusers

`func (o *UserExpirationRuleRequest) HasExcludeSuperusers() bool`

HasExcludeSuperusers returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *UserExpirationRuleRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *UserExpirationRuleRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *UserExpirationRuleRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *UserExpirationRuleRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


