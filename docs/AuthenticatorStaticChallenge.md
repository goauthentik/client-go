# AuthenticatorStaticChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ChallengeChoices**](ChallengeChoices.md) |  | 
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-authenticator-static"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**PendingUser** | **string** |  | 
**PendingUserAvatar** | **string** |  | 
**Codes** | **[]string** |  | 

## Methods

### NewAuthenticatorStaticChallenge

`func NewAuthenticatorStaticChallenge(type_ ChallengeChoices, pendingUser string, pendingUserAvatar string, codes []string, ) *AuthenticatorStaticChallenge`

NewAuthenticatorStaticChallenge instantiates a new AuthenticatorStaticChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorStaticChallengeWithDefaults

`func NewAuthenticatorStaticChallengeWithDefaults() *AuthenticatorStaticChallenge`

NewAuthenticatorStaticChallengeWithDefaults instantiates a new AuthenticatorStaticChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthenticatorStaticChallenge) GetType() ChallengeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthenticatorStaticChallenge) GetTypeOk() (*ChallengeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthenticatorStaticChallenge) SetType(v ChallengeChoices)`

SetType sets Type field to given value.


### GetFlowInfo

`func (o *AuthenticatorStaticChallenge) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *AuthenticatorStaticChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *AuthenticatorStaticChallenge) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *AuthenticatorStaticChallenge) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *AuthenticatorStaticChallenge) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AuthenticatorStaticChallenge) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AuthenticatorStaticChallenge) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AuthenticatorStaticChallenge) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *AuthenticatorStaticChallenge) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *AuthenticatorStaticChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *AuthenticatorStaticChallenge) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *AuthenticatorStaticChallenge) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetPendingUser

`func (o *AuthenticatorStaticChallenge) GetPendingUser() string`

GetPendingUser returns the PendingUser field if non-nil, zero value otherwise.

### GetPendingUserOk

`func (o *AuthenticatorStaticChallenge) GetPendingUserOk() (*string, bool)`

GetPendingUserOk returns a tuple with the PendingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUser

`func (o *AuthenticatorStaticChallenge) SetPendingUser(v string)`

SetPendingUser sets PendingUser field to given value.


### GetPendingUserAvatar

`func (o *AuthenticatorStaticChallenge) GetPendingUserAvatar() string`

GetPendingUserAvatar returns the PendingUserAvatar field if non-nil, zero value otherwise.

### GetPendingUserAvatarOk

`func (o *AuthenticatorStaticChallenge) GetPendingUserAvatarOk() (*string, bool)`

GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUserAvatar

`func (o *AuthenticatorStaticChallenge) SetPendingUserAvatar(v string)`

SetPendingUserAvatar sets PendingUserAvatar field to given value.


### GetCodes

`func (o *AuthenticatorStaticChallenge) GetCodes() []string`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *AuthenticatorStaticChallenge) GetCodesOk() (*[]string, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *AuthenticatorStaticChallenge) SetCodes(v []string)`

SetCodes sets Codes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


