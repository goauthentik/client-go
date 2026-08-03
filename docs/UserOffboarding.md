# UserOffboarding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**User** | **int32** |  | 
**UserObj** | [**PartialUser**](PartialUser.md) |  | [readonly] 
**ScheduledAt** | **time.Time** | Absolute time at which the offboarding action is executed. | 
**Action** | Pointer to [**OffboardingActionEnum**](OffboardingActionEnum.md) |  | [optional] 
**RevokeSessions** | Pointer to **bool** | Revoke all of the user&#39;s sessions when offboarding. | [optional] 
**RevokeTokens** | Pointer to **bool** | Revoke all of the user&#39;s tokens when offboarding. | [optional] 
**Status** | [**OffboardingStatusEnum**](OffboardingStatusEnum.md) |  | [readonly] 
**CreatedByObj** | [**PartialUser**](PartialUser.md) |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ExecutedAt** | **NullableTime** |  | [readonly] 

## Methods

### NewUserOffboarding

`func NewUserOffboarding(id string, user int32, userObj PartialUser, scheduledAt time.Time, status OffboardingStatusEnum, createdByObj PartialUser, createdAt time.Time, executedAt NullableTime, ) *UserOffboarding`

NewUserOffboarding instantiates a new UserOffboarding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOffboardingWithDefaults

`func NewUserOffboardingWithDefaults() *UserOffboarding`

NewUserOffboardingWithDefaults instantiates a new UserOffboarding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserOffboarding) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserOffboarding) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserOffboarding) SetId(v string)`

SetId sets Id field to given value.


### GetUser

`func (o *UserOffboarding) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserOffboarding) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserOffboarding) SetUser(v int32)`

SetUser sets User field to given value.


### GetUserObj

`func (o *UserOffboarding) GetUserObj() PartialUser`

GetUserObj returns the UserObj field if non-nil, zero value otherwise.

### GetUserObjOk

`func (o *UserOffboarding) GetUserObjOk() (*PartialUser, bool)`

GetUserObjOk returns a tuple with the UserObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObj

`func (o *UserOffboarding) SetUserObj(v PartialUser)`

SetUserObj sets UserObj field to given value.


### GetScheduledAt

`func (o *UserOffboarding) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *UserOffboarding) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *UserOffboarding) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.


### GetAction

`func (o *UserOffboarding) GetAction() OffboardingActionEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UserOffboarding) GetActionOk() (*OffboardingActionEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UserOffboarding) SetAction(v OffboardingActionEnum)`

SetAction sets Action field to given value.

### HasAction

`func (o *UserOffboarding) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRevokeSessions

`func (o *UserOffboarding) GetRevokeSessions() bool`

GetRevokeSessions returns the RevokeSessions field if non-nil, zero value otherwise.

### GetRevokeSessionsOk

`func (o *UserOffboarding) GetRevokeSessionsOk() (*bool, bool)`

GetRevokeSessionsOk returns a tuple with the RevokeSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeSessions

`func (o *UserOffboarding) SetRevokeSessions(v bool)`

SetRevokeSessions sets RevokeSessions field to given value.

### HasRevokeSessions

`func (o *UserOffboarding) HasRevokeSessions() bool`

HasRevokeSessions returns a boolean if a field has been set.

### GetRevokeTokens

`func (o *UserOffboarding) GetRevokeTokens() bool`

GetRevokeTokens returns the RevokeTokens field if non-nil, zero value otherwise.

### GetRevokeTokensOk

`func (o *UserOffboarding) GetRevokeTokensOk() (*bool, bool)`

GetRevokeTokensOk returns a tuple with the RevokeTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeTokens

`func (o *UserOffboarding) SetRevokeTokens(v bool)`

SetRevokeTokens sets RevokeTokens field to given value.

### HasRevokeTokens

`func (o *UserOffboarding) HasRevokeTokens() bool`

HasRevokeTokens returns a boolean if a field has been set.

### GetStatus

`func (o *UserOffboarding) GetStatus() OffboardingStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserOffboarding) GetStatusOk() (*OffboardingStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserOffboarding) SetStatus(v OffboardingStatusEnum)`

SetStatus sets Status field to given value.


### GetCreatedByObj

`func (o *UserOffboarding) GetCreatedByObj() PartialUser`

GetCreatedByObj returns the CreatedByObj field if non-nil, zero value otherwise.

### GetCreatedByObjOk

`func (o *UserOffboarding) GetCreatedByObjOk() (*PartialUser, bool)`

GetCreatedByObjOk returns a tuple with the CreatedByObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByObj

`func (o *UserOffboarding) SetCreatedByObj(v PartialUser)`

SetCreatedByObj sets CreatedByObj field to given value.


### GetCreatedAt

`func (o *UserOffboarding) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserOffboarding) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserOffboarding) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExecutedAt

`func (o *UserOffboarding) GetExecutedAt() time.Time`

GetExecutedAt returns the ExecutedAt field if non-nil, zero value otherwise.

### GetExecutedAtOk

`func (o *UserOffboarding) GetExecutedAtOk() (*time.Time, bool)`

GetExecutedAtOk returns a tuple with the ExecutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAt

`func (o *UserOffboarding) SetExecutedAt(v time.Time)`

SetExecutedAt sets ExecutedAt field to given value.


### SetExecutedAtNil

`func (o *UserOffboarding) SetExecutedAtNil(b bool)`

 SetExecutedAtNil sets the value for ExecutedAt to be an explicit nil

### UnsetExecutedAt
`func (o *UserOffboarding) UnsetExecutedAt()`

UnsetExecutedAt ensures that no value is present for ExecutedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


