# PatchedUserExpirationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
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

### NewPatchedUserExpirationRuleRequest

`func NewPatchedUserExpirationRuleRequest() *PatchedUserExpirationRuleRequest`

NewPatchedUserExpirationRuleRequest instantiates a new PatchedUserExpirationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedUserExpirationRuleRequestWithDefaults

`func NewPatchedUserExpirationRuleRequestWithDefaults() *PatchedUserExpirationRuleRequest`

NewPatchedUserExpirationRuleRequestWithDefaults instantiates a new PatchedUserExpirationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *PatchedUserExpirationRuleRequest) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *PatchedUserExpirationRuleRequest) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *PatchedUserExpirationRuleRequest) SetPk(v string)`

SetPk sets Pk field to given value.

### HasPk

`func (o *PatchedUserExpirationRuleRequest) HasPk() bool`

HasPk returns a boolean if a field has been set.

### GetName

`func (o *PatchedUserExpirationRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedUserExpirationRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedUserExpirationRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedUserExpirationRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedUserExpirationRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedUserExpirationRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedUserExpirationRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedUserExpirationRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetActivityBasis

`func (o *PatchedUserExpirationRuleRequest) GetActivityBasis() ActivityBasisEnum`

GetActivityBasis returns the ActivityBasis field if non-nil, zero value otherwise.

### GetActivityBasisOk

`func (o *PatchedUserExpirationRuleRequest) GetActivityBasisOk() (*ActivityBasisEnum, bool)`

GetActivityBasisOk returns a tuple with the ActivityBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityBasis

`func (o *PatchedUserExpirationRuleRequest) SetActivityBasis(v ActivityBasisEnum)`

SetActivityBasis sets ActivityBasis field to given value.

### HasActivityBasis

`func (o *PatchedUserExpirationRuleRequest) HasActivityBasis() bool`

HasActivityBasis returns a boolean if a field has been set.

### GetGroup

`func (o *PatchedUserExpirationRuleRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedUserExpirationRuleRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedUserExpirationRuleRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedUserExpirationRuleRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedUserExpirationRuleRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedUserExpirationRuleRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetUserTypes

`func (o *PatchedUserExpirationRuleRequest) GetUserTypes() []UserTypeEnum`

GetUserTypes returns the UserTypes field if non-nil, zero value otherwise.

### GetUserTypesOk

`func (o *PatchedUserExpirationRuleRequest) GetUserTypesOk() (*[]UserTypeEnum, bool)`

GetUserTypesOk returns a tuple with the UserTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTypes

`func (o *PatchedUserExpirationRuleRequest) SetUserTypes(v []UserTypeEnum)`

SetUserTypes sets UserTypes field to given value.

### HasUserTypes

`func (o *PatchedUserExpirationRuleRequest) HasUserTypes() bool`

HasUserTypes returns a boolean if a field has been set.

### GetInactivityDuration

`func (o *PatchedUserExpirationRuleRequest) GetInactivityDuration() string`

GetInactivityDuration returns the InactivityDuration field if non-nil, zero value otherwise.

### GetInactivityDurationOk

`func (o *PatchedUserExpirationRuleRequest) GetInactivityDurationOk() (*string, bool)`

GetInactivityDurationOk returns a tuple with the InactivityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivityDuration

`func (o *PatchedUserExpirationRuleRequest) SetInactivityDuration(v string)`

SetInactivityDuration sets InactivityDuration field to given value.

### HasInactivityDuration

`func (o *PatchedUserExpirationRuleRequest) HasInactivityDuration() bool`

HasInactivityDuration returns a boolean if a field has been set.

### GetAction

`func (o *PatchedUserExpirationRuleRequest) GetAction() OffboardingActionEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PatchedUserExpirationRuleRequest) GetActionOk() (*OffboardingActionEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PatchedUserExpirationRuleRequest) SetAction(v OffboardingActionEnum)`

SetAction sets Action field to given value.

### HasAction

`func (o *PatchedUserExpirationRuleRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRevokeSessions

`func (o *PatchedUserExpirationRuleRequest) GetRevokeSessions() bool`

GetRevokeSessions returns the RevokeSessions field if non-nil, zero value otherwise.

### GetRevokeSessionsOk

`func (o *PatchedUserExpirationRuleRequest) GetRevokeSessionsOk() (*bool, bool)`

GetRevokeSessionsOk returns a tuple with the RevokeSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeSessions

`func (o *PatchedUserExpirationRuleRequest) SetRevokeSessions(v bool)`

SetRevokeSessions sets RevokeSessions field to given value.

### HasRevokeSessions

`func (o *PatchedUserExpirationRuleRequest) HasRevokeSessions() bool`

HasRevokeSessions returns a boolean if a field has been set.

### GetRevokeTokens

`func (o *PatchedUserExpirationRuleRequest) GetRevokeTokens() bool`

GetRevokeTokens returns the RevokeTokens field if non-nil, zero value otherwise.

### GetRevokeTokensOk

`func (o *PatchedUserExpirationRuleRequest) GetRevokeTokensOk() (*bool, bool)`

GetRevokeTokensOk returns a tuple with the RevokeTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeTokens

`func (o *PatchedUserExpirationRuleRequest) SetRevokeTokens(v bool)`

SetRevokeTokens sets RevokeTokens field to given value.

### HasRevokeTokens

`func (o *PatchedUserExpirationRuleRequest) HasRevokeTokens() bool`

HasRevokeTokens returns a boolean if a field has been set.

### GetWarnBefore

`func (o *PatchedUserExpirationRuleRequest) GetWarnBefore() string`

GetWarnBefore returns the WarnBefore field if non-nil, zero value otherwise.

### GetWarnBeforeOk

`func (o *PatchedUserExpirationRuleRequest) GetWarnBeforeOk() (*string, bool)`

GetWarnBeforeOk returns a tuple with the WarnBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnBefore

`func (o *PatchedUserExpirationRuleRequest) SetWarnBefore(v string)`

SetWarnBefore sets WarnBefore field to given value.

### HasWarnBefore

`func (o *PatchedUserExpirationRuleRequest) HasWarnBefore() bool`

HasWarnBefore returns a boolean if a field has been set.

### SetWarnBeforeNil

`func (o *PatchedUserExpirationRuleRequest) SetWarnBeforeNil(b bool)`

 SetWarnBeforeNil sets the value for WarnBefore to be an explicit nil

### UnsetWarnBefore
`func (o *PatchedUserExpirationRuleRequest) UnsetWarnBefore()`

UnsetWarnBefore ensures that no value is present for WarnBefore, not even an explicit nil
### GetNotificationTransports

`func (o *PatchedUserExpirationRuleRequest) GetNotificationTransports() []string`

GetNotificationTransports returns the NotificationTransports field if non-nil, zero value otherwise.

### GetNotificationTransportsOk

`func (o *PatchedUserExpirationRuleRequest) GetNotificationTransportsOk() (*[]string, bool)`

GetNotificationTransportsOk returns a tuple with the NotificationTransports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTransports

`func (o *PatchedUserExpirationRuleRequest) SetNotificationTransports(v []string)`

SetNotificationTransports sets NotificationTransports field to given value.

### HasNotificationTransports

`func (o *PatchedUserExpirationRuleRequest) HasNotificationTransports() bool`

HasNotificationTransports returns a boolean if a field has been set.

### GetExcludeSuperusers

`func (o *PatchedUserExpirationRuleRequest) GetExcludeSuperusers() bool`

GetExcludeSuperusers returns the ExcludeSuperusers field if non-nil, zero value otherwise.

### GetExcludeSuperusersOk

`func (o *PatchedUserExpirationRuleRequest) GetExcludeSuperusersOk() (*bool, bool)`

GetExcludeSuperusersOk returns a tuple with the ExcludeSuperusers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSuperusers

`func (o *PatchedUserExpirationRuleRequest) SetExcludeSuperusers(v bool)`

SetExcludeSuperusers sets ExcludeSuperusers field to given value.

### HasExcludeSuperusers

`func (o *PatchedUserExpirationRuleRequest) HasExcludeSuperusers() bool`

HasExcludeSuperusers returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *PatchedUserExpirationRuleRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *PatchedUserExpirationRuleRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *PatchedUserExpirationRuleRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *PatchedUserExpirationRuleRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


