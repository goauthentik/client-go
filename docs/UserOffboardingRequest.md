# UserOffboardingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **int32** |  | 
**ScheduledAt** | **time.Time** | Absolute time at which the offboarding action is executed. | 
**Action** | Pointer to [**OffboardingActionEnum**](OffboardingActionEnum.md) |  | [optional] 
**RevokeSessions** | Pointer to **bool** | Revoke all of the user&#39;s sessions when offboarding. | [optional] 
**RevokeTokens** | Pointer to **bool** | Revoke all of the user&#39;s tokens when offboarding. | [optional] 

## Methods

### NewUserOffboardingRequest

`func NewUserOffboardingRequest(user int32, scheduledAt time.Time, ) *UserOffboardingRequest`

NewUserOffboardingRequest instantiates a new UserOffboardingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOffboardingRequestWithDefaults

`func NewUserOffboardingRequestWithDefaults() *UserOffboardingRequest`

NewUserOffboardingRequestWithDefaults instantiates a new UserOffboardingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserOffboardingRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserOffboardingRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserOffboardingRequest) SetUser(v int32)`

SetUser sets User field to given value.


### GetScheduledAt

`func (o *UserOffboardingRequest) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *UserOffboardingRequest) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *UserOffboardingRequest) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.


### GetAction

`func (o *UserOffboardingRequest) GetAction() OffboardingActionEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UserOffboardingRequest) GetActionOk() (*OffboardingActionEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UserOffboardingRequest) SetAction(v OffboardingActionEnum)`

SetAction sets Action field to given value.

### HasAction

`func (o *UserOffboardingRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRevokeSessions

`func (o *UserOffboardingRequest) GetRevokeSessions() bool`

GetRevokeSessions returns the RevokeSessions field if non-nil, zero value otherwise.

### GetRevokeSessionsOk

`func (o *UserOffboardingRequest) GetRevokeSessionsOk() (*bool, bool)`

GetRevokeSessionsOk returns a tuple with the RevokeSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeSessions

`func (o *UserOffboardingRequest) SetRevokeSessions(v bool)`

SetRevokeSessions sets RevokeSessions field to given value.

### HasRevokeSessions

`func (o *UserOffboardingRequest) HasRevokeSessions() bool`

HasRevokeSessions returns a boolean if a field has been set.

### GetRevokeTokens

`func (o *UserOffboardingRequest) GetRevokeTokens() bool`

GetRevokeTokens returns the RevokeTokens field if non-nil, zero value otherwise.

### GetRevokeTokensOk

`func (o *UserOffboardingRequest) GetRevokeTokensOk() (*bool, bool)`

GetRevokeTokensOk returns a tuple with the RevokeTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeTokens

`func (o *UserOffboardingRequest) SetRevokeTokens(v bool)`

SetRevokeTokens sets RevokeTokens field to given value.

### HasRevokeTokens

`func (o *UserOffboardingRequest) HasRevokeTokens() bool`

HasRevokeTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


