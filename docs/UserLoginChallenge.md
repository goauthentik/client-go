# UserLoginChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ChallengeChoices**](ChallengeChoices.md) |  | 
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-user-login"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**PendingUser** | **string** |  | 
**PendingUserAvatar** | **string** |  | 

## Methods

### NewUserLoginChallenge

`func NewUserLoginChallenge(type_ ChallengeChoices, pendingUser string, pendingUserAvatar string, ) *UserLoginChallenge`

NewUserLoginChallenge instantiates a new UserLoginChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLoginChallengeWithDefaults

`func NewUserLoginChallengeWithDefaults() *UserLoginChallenge`

NewUserLoginChallengeWithDefaults instantiates a new UserLoginChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserLoginChallenge) GetType() ChallengeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserLoginChallenge) GetTypeOk() (*ChallengeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserLoginChallenge) SetType(v ChallengeChoices)`

SetType sets Type field to given value.


### GetFlowInfo

`func (o *UserLoginChallenge) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *UserLoginChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *UserLoginChallenge) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *UserLoginChallenge) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *UserLoginChallenge) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *UserLoginChallenge) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *UserLoginChallenge) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *UserLoginChallenge) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *UserLoginChallenge) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *UserLoginChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *UserLoginChallenge) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *UserLoginChallenge) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetPendingUser

`func (o *UserLoginChallenge) GetPendingUser() string`

GetPendingUser returns the PendingUser field if non-nil, zero value otherwise.

### GetPendingUserOk

`func (o *UserLoginChallenge) GetPendingUserOk() (*string, bool)`

GetPendingUserOk returns a tuple with the PendingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUser

`func (o *UserLoginChallenge) SetPendingUser(v string)`

SetPendingUser sets PendingUser field to given value.


### GetPendingUserAvatar

`func (o *UserLoginChallenge) GetPendingUserAvatar() string`

GetPendingUserAvatar returns the PendingUserAvatar field if non-nil, zero value otherwise.

### GetPendingUserAvatarOk

`func (o *UserLoginChallenge) GetPendingUserAvatarOk() (*string, bool)`

GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUserAvatar

`func (o *UserLoginChallenge) SetPendingUserAvatar(v string)`

SetPendingUserAvatar sets PendingUserAvatar field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


