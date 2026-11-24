# UserExpirationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | Pointer to **string** |  | [optional] 
**PbmUuid** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**ActivityBasis** | Pointer to [**ActivityBasisEnum**](ActivityBasisEnum.md) | Measure inactivity from the last login or successful user activity. | [optional] 
**Group** | Pointer to **NullableString** | Only expire members of this group (including descendant groups). Leave empty to apply to every user. | [optional] 
**GroupObj** | [**UserExpirationRuleGroup**](UserExpirationRuleGroup.md) |  | [readonly] 
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

### NewUserExpirationRule

`func NewUserExpirationRule(pbmUuid string, name string, groupObj UserExpirationRuleGroup, ) *UserExpirationRule`

NewUserExpirationRule instantiates a new UserExpirationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserExpirationRuleWithDefaults

`func NewUserExpirationRuleWithDefaults() *UserExpirationRule`

NewUserExpirationRuleWithDefaults instantiates a new UserExpirationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *UserExpirationRule) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *UserExpirationRule) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *UserExpirationRule) SetPk(v string)`

SetPk sets Pk field to given value.

### HasPk

`func (o *UserExpirationRule) HasPk() bool`

HasPk returns a boolean if a field has been set.

### GetPbmUuid

`func (o *UserExpirationRule) GetPbmUuid() string`

GetPbmUuid returns the PbmUuid field if non-nil, zero value otherwise.

### GetPbmUuidOk

`func (o *UserExpirationRule) GetPbmUuidOk() (*string, bool)`

GetPbmUuidOk returns a tuple with the PbmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbmUuid

`func (o *UserExpirationRule) SetPbmUuid(v string)`

SetPbmUuid sets PbmUuid field to given value.


### GetName

`func (o *UserExpirationRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserExpirationRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserExpirationRule) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *UserExpirationRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserExpirationRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserExpirationRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserExpirationRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetActivityBasis

`func (o *UserExpirationRule) GetActivityBasis() ActivityBasisEnum`

GetActivityBasis returns the ActivityBasis field if non-nil, zero value otherwise.

### GetActivityBasisOk

`func (o *UserExpirationRule) GetActivityBasisOk() (*ActivityBasisEnum, bool)`

GetActivityBasisOk returns a tuple with the ActivityBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityBasis

`func (o *UserExpirationRule) SetActivityBasis(v ActivityBasisEnum)`

SetActivityBasis sets ActivityBasis field to given value.

### HasActivityBasis

`func (o *UserExpirationRule) HasActivityBasis() bool`

HasActivityBasis returns a boolean if a field has been set.

### GetGroup

`func (o *UserExpirationRule) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UserExpirationRule) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UserExpirationRule) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UserExpirationRule) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *UserExpirationRule) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *UserExpirationRule) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetGroupObj

`func (o *UserExpirationRule) GetGroupObj() UserExpirationRuleGroup`

GetGroupObj returns the GroupObj field if non-nil, zero value otherwise.

### GetGroupObjOk

`func (o *UserExpirationRule) GetGroupObjOk() (*UserExpirationRuleGroup, bool)`

GetGroupObjOk returns a tuple with the GroupObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObj

`func (o *UserExpirationRule) SetGroupObj(v UserExpirationRuleGroup)`

SetGroupObj sets GroupObj field to given value.


### GetUserTypes

`func (o *UserExpirationRule) GetUserTypes() []UserTypeEnum`

GetUserTypes returns the UserTypes field if non-nil, zero value otherwise.

### GetUserTypesOk

`func (o *UserExpirationRule) GetUserTypesOk() (*[]UserTypeEnum, bool)`

GetUserTypesOk returns a tuple with the UserTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTypes

`func (o *UserExpirationRule) SetUserTypes(v []UserTypeEnum)`

SetUserTypes sets UserTypes field to given value.

### HasUserTypes

`func (o *UserExpirationRule) HasUserTypes() bool`

HasUserTypes returns a boolean if a field has been set.

### GetInactivityDuration

`func (o *UserExpirationRule) GetInactivityDuration() string`

GetInactivityDuration returns the InactivityDuration field if non-nil, zero value otherwise.

### GetInactivityDurationOk

`func (o *UserExpirationRule) GetInactivityDurationOk() (*string, bool)`

GetInactivityDurationOk returns a tuple with the InactivityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivityDuration

`func (o *UserExpirationRule) SetInactivityDuration(v string)`

SetInactivityDuration sets InactivityDuration field to given value.

### HasInactivityDuration

`func (o *UserExpirationRule) HasInactivityDuration() bool`

HasInactivityDuration returns a boolean if a field has been set.

### GetAction

`func (o *UserExpirationRule) GetAction() OffboardingActionEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UserExpirationRule) GetActionOk() (*OffboardingActionEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UserExpirationRule) SetAction(v OffboardingActionEnum)`

SetAction sets Action field to given value.

### HasAction

`func (o *UserExpirationRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRevokeSessions

`func (o *UserExpirationRule) GetRevokeSessions() bool`

GetRevokeSessions returns the RevokeSessions field if non-nil, zero value otherwise.

### GetRevokeSessionsOk

`func (o *UserExpirationRule) GetRevokeSessionsOk() (*bool, bool)`

GetRevokeSessionsOk returns a tuple with the RevokeSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeSessions

`func (o *UserExpirationRule) SetRevokeSessions(v bool)`

SetRevokeSessions sets RevokeSessions field to given value.

### HasRevokeSessions

`func (o *UserExpirationRule) HasRevokeSessions() bool`

HasRevokeSessions returns a boolean if a field has been set.

### GetRevokeTokens

`func (o *UserExpirationRule) GetRevokeTokens() bool`

GetRevokeTokens returns the RevokeTokens field if non-nil, zero value otherwise.

### GetRevokeTokensOk

`func (o *UserExpirationRule) GetRevokeTokensOk() (*bool, bool)`

GetRevokeTokensOk returns a tuple with the RevokeTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeTokens

`func (o *UserExpirationRule) SetRevokeTokens(v bool)`

SetRevokeTokens sets RevokeTokens field to given value.

### HasRevokeTokens

`func (o *UserExpirationRule) HasRevokeTokens() bool`

HasRevokeTokens returns a boolean if a field has been set.

### GetWarnBefore

`func (o *UserExpirationRule) GetWarnBefore() string`

GetWarnBefore returns the WarnBefore field if non-nil, zero value otherwise.

### GetWarnBeforeOk

`func (o *UserExpirationRule) GetWarnBeforeOk() (*string, bool)`

GetWarnBeforeOk returns a tuple with the WarnBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnBefore

`func (o *UserExpirationRule) SetWarnBefore(v string)`

SetWarnBefore sets WarnBefore field to given value.

### HasWarnBefore

`func (o *UserExpirationRule) HasWarnBefore() bool`

HasWarnBefore returns a boolean if a field has been set.

### SetWarnBeforeNil

`func (o *UserExpirationRule) SetWarnBeforeNil(b bool)`

 SetWarnBeforeNil sets the value for WarnBefore to be an explicit nil

### UnsetWarnBefore
`func (o *UserExpirationRule) UnsetWarnBefore()`

UnsetWarnBefore ensures that no value is present for WarnBefore, not even an explicit nil
### GetNotificationTransports

`func (o *UserExpirationRule) GetNotificationTransports() []string`

GetNotificationTransports returns the NotificationTransports field if non-nil, zero value otherwise.

### GetNotificationTransportsOk

`func (o *UserExpirationRule) GetNotificationTransportsOk() (*[]string, bool)`

GetNotificationTransportsOk returns a tuple with the NotificationTransports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTransports

`func (o *UserExpirationRule) SetNotificationTransports(v []string)`

SetNotificationTransports sets NotificationTransports field to given value.

### HasNotificationTransports

`func (o *UserExpirationRule) HasNotificationTransports() bool`

HasNotificationTransports returns a boolean if a field has been set.

### GetExcludeSuperusers

`func (o *UserExpirationRule) GetExcludeSuperusers() bool`

GetExcludeSuperusers returns the ExcludeSuperusers field if non-nil, zero value otherwise.

### GetExcludeSuperusersOk

`func (o *UserExpirationRule) GetExcludeSuperusersOk() (*bool, bool)`

GetExcludeSuperusersOk returns a tuple with the ExcludeSuperusers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSuperusers

`func (o *UserExpirationRule) SetExcludeSuperusers(v bool)`

SetExcludeSuperusers sets ExcludeSuperusers field to given value.

### HasExcludeSuperusers

`func (o *UserExpirationRule) HasExcludeSuperusers() bool`

HasExcludeSuperusers returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *UserExpirationRule) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *UserExpirationRule) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *UserExpirationRule) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *UserExpirationRule) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


