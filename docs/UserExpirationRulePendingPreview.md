# UserExpirationRulePendingPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**User** | [**UserExpirationRuleUser**](UserExpirationRuleUser.md) |  | [readonly] 
**PreviousAction** | [**OffboardingActionEnum**](OffboardingActionEnum.md) |  | [readonly] 
**Action** | [**NullableOffboardingActionEnum**](OffboardingActionEnum.md) |  | [readonly] 
**PreviousScheduledAt** | **time.Time** |  | [readonly] 
**ScheduledAt** | **NullableTime** |  | [readonly] 
**PreviousRevokeSessions** | **bool** |  | [readonly] 
**RevokeSessions** | **NullableBool** |  | [readonly] 
**PreviousRevokeTokens** | **bool** |  | [readonly] 
**RevokeTokens** | **NullableBool** |  | [readonly] 

## Methods

### NewUserExpirationRulePendingPreview

`func NewUserExpirationRulePendingPreview(id string, user UserExpirationRuleUser, previousAction OffboardingActionEnum, action NullableOffboardingActionEnum, previousScheduledAt time.Time, scheduledAt NullableTime, previousRevokeSessions bool, revokeSessions NullableBool, previousRevokeTokens bool, revokeTokens NullableBool, ) *UserExpirationRulePendingPreview`

NewUserExpirationRulePendingPreview instantiates a new UserExpirationRulePendingPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserExpirationRulePendingPreviewWithDefaults

`func NewUserExpirationRulePendingPreviewWithDefaults() *UserExpirationRulePendingPreview`

NewUserExpirationRulePendingPreviewWithDefaults instantiates a new UserExpirationRulePendingPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserExpirationRulePendingPreview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserExpirationRulePendingPreview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserExpirationRulePendingPreview) SetId(v string)`

SetId sets Id field to given value.


### GetUser

`func (o *UserExpirationRulePendingPreview) GetUser() UserExpirationRuleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserExpirationRulePendingPreview) GetUserOk() (*UserExpirationRuleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserExpirationRulePendingPreview) SetUser(v UserExpirationRuleUser)`

SetUser sets User field to given value.


### GetPreviousAction

`func (o *UserExpirationRulePendingPreview) GetPreviousAction() OffboardingActionEnum`

GetPreviousAction returns the PreviousAction field if non-nil, zero value otherwise.

### GetPreviousActionOk

`func (o *UserExpirationRulePendingPreview) GetPreviousActionOk() (*OffboardingActionEnum, bool)`

GetPreviousActionOk returns a tuple with the PreviousAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAction

`func (o *UserExpirationRulePendingPreview) SetPreviousAction(v OffboardingActionEnum)`

SetPreviousAction sets PreviousAction field to given value.


### GetAction

`func (o *UserExpirationRulePendingPreview) GetAction() OffboardingActionEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UserExpirationRulePendingPreview) GetActionOk() (*OffboardingActionEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UserExpirationRulePendingPreview) SetAction(v OffboardingActionEnum)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *UserExpirationRulePendingPreview) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *UserExpirationRulePendingPreview) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetPreviousScheduledAt

`func (o *UserExpirationRulePendingPreview) GetPreviousScheduledAt() time.Time`

GetPreviousScheduledAt returns the PreviousScheduledAt field if non-nil, zero value otherwise.

### GetPreviousScheduledAtOk

`func (o *UserExpirationRulePendingPreview) GetPreviousScheduledAtOk() (*time.Time, bool)`

GetPreviousScheduledAtOk returns a tuple with the PreviousScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousScheduledAt

`func (o *UserExpirationRulePendingPreview) SetPreviousScheduledAt(v time.Time)`

SetPreviousScheduledAt sets PreviousScheduledAt field to given value.


### GetScheduledAt

`func (o *UserExpirationRulePendingPreview) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *UserExpirationRulePendingPreview) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *UserExpirationRulePendingPreview) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.


### SetScheduledAtNil

`func (o *UserExpirationRulePendingPreview) SetScheduledAtNil(b bool)`

 SetScheduledAtNil sets the value for ScheduledAt to be an explicit nil

### UnsetScheduledAt
`func (o *UserExpirationRulePendingPreview) UnsetScheduledAt()`

UnsetScheduledAt ensures that no value is present for ScheduledAt, not even an explicit nil
### GetPreviousRevokeSessions

`func (o *UserExpirationRulePendingPreview) GetPreviousRevokeSessions() bool`

GetPreviousRevokeSessions returns the PreviousRevokeSessions field if non-nil, zero value otherwise.

### GetPreviousRevokeSessionsOk

`func (o *UserExpirationRulePendingPreview) GetPreviousRevokeSessionsOk() (*bool, bool)`

GetPreviousRevokeSessionsOk returns a tuple with the PreviousRevokeSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRevokeSessions

`func (o *UserExpirationRulePendingPreview) SetPreviousRevokeSessions(v bool)`

SetPreviousRevokeSessions sets PreviousRevokeSessions field to given value.


### GetRevokeSessions

`func (o *UserExpirationRulePendingPreview) GetRevokeSessions() bool`

GetRevokeSessions returns the RevokeSessions field if non-nil, zero value otherwise.

### GetRevokeSessionsOk

`func (o *UserExpirationRulePendingPreview) GetRevokeSessionsOk() (*bool, bool)`

GetRevokeSessionsOk returns a tuple with the RevokeSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeSessions

`func (o *UserExpirationRulePendingPreview) SetRevokeSessions(v bool)`

SetRevokeSessions sets RevokeSessions field to given value.


### SetRevokeSessionsNil

`func (o *UserExpirationRulePendingPreview) SetRevokeSessionsNil(b bool)`

 SetRevokeSessionsNil sets the value for RevokeSessions to be an explicit nil

### UnsetRevokeSessions
`func (o *UserExpirationRulePendingPreview) UnsetRevokeSessions()`

UnsetRevokeSessions ensures that no value is present for RevokeSessions, not even an explicit nil
### GetPreviousRevokeTokens

`func (o *UserExpirationRulePendingPreview) GetPreviousRevokeTokens() bool`

GetPreviousRevokeTokens returns the PreviousRevokeTokens field if non-nil, zero value otherwise.

### GetPreviousRevokeTokensOk

`func (o *UserExpirationRulePendingPreview) GetPreviousRevokeTokensOk() (*bool, bool)`

GetPreviousRevokeTokensOk returns a tuple with the PreviousRevokeTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRevokeTokens

`func (o *UserExpirationRulePendingPreview) SetPreviousRevokeTokens(v bool)`

SetPreviousRevokeTokens sets PreviousRevokeTokens field to given value.


### GetRevokeTokens

`func (o *UserExpirationRulePendingPreview) GetRevokeTokens() bool`

GetRevokeTokens returns the RevokeTokens field if non-nil, zero value otherwise.

### GetRevokeTokensOk

`func (o *UserExpirationRulePendingPreview) GetRevokeTokensOk() (*bool, bool)`

GetRevokeTokensOk returns a tuple with the RevokeTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeTokens

`func (o *UserExpirationRulePendingPreview) SetRevokeTokens(v bool)`

SetRevokeTokens sets RevokeTokens field to given value.


### SetRevokeTokensNil

`func (o *UserExpirationRulePendingPreview) SetRevokeTokensNil(b bool)`

 SetRevokeTokensNil sets the value for RevokeTokens to be an explicit nil

### UnsetRevokeTokens
`func (o *UserExpirationRulePendingPreview) UnsetRevokeTokens()`

UnsetRevokeTokens ensures that no value is present for RevokeTokens, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


