# AuthenticatorTOTPChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ChallengeChoices**](ChallengeChoices.md) |  | 
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-authenticator-totp"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**PendingUser** | **string** |  | 
**PendingUserAvatar** | **string** |  | 
**ConfigUrl** | **string** |  | 

## Methods

### NewAuthenticatorTOTPChallenge

`func NewAuthenticatorTOTPChallenge(type_ ChallengeChoices, pendingUser string, pendingUserAvatar string, configUrl string, ) *AuthenticatorTOTPChallenge`

NewAuthenticatorTOTPChallenge instantiates a new AuthenticatorTOTPChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorTOTPChallengeWithDefaults

`func NewAuthenticatorTOTPChallengeWithDefaults() *AuthenticatorTOTPChallenge`

NewAuthenticatorTOTPChallengeWithDefaults instantiates a new AuthenticatorTOTPChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthenticatorTOTPChallenge) GetType() ChallengeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthenticatorTOTPChallenge) GetTypeOk() (*ChallengeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthenticatorTOTPChallenge) SetType(v ChallengeChoices)`

SetType sets Type field to given value.


### GetFlowInfo

`func (o *AuthenticatorTOTPChallenge) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *AuthenticatorTOTPChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *AuthenticatorTOTPChallenge) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *AuthenticatorTOTPChallenge) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *AuthenticatorTOTPChallenge) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AuthenticatorTOTPChallenge) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AuthenticatorTOTPChallenge) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AuthenticatorTOTPChallenge) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *AuthenticatorTOTPChallenge) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *AuthenticatorTOTPChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *AuthenticatorTOTPChallenge) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *AuthenticatorTOTPChallenge) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetPendingUser

`func (o *AuthenticatorTOTPChallenge) GetPendingUser() string`

GetPendingUser returns the PendingUser field if non-nil, zero value otherwise.

### GetPendingUserOk

`func (o *AuthenticatorTOTPChallenge) GetPendingUserOk() (*string, bool)`

GetPendingUserOk returns a tuple with the PendingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUser

`func (o *AuthenticatorTOTPChallenge) SetPendingUser(v string)`

SetPendingUser sets PendingUser field to given value.


### GetPendingUserAvatar

`func (o *AuthenticatorTOTPChallenge) GetPendingUserAvatar() string`

GetPendingUserAvatar returns the PendingUserAvatar field if non-nil, zero value otherwise.

### GetPendingUserAvatarOk

`func (o *AuthenticatorTOTPChallenge) GetPendingUserAvatarOk() (*string, bool)`

GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUserAvatar

`func (o *AuthenticatorTOTPChallenge) SetPendingUserAvatar(v string)`

SetPendingUserAvatar sets PendingUserAvatar field to given value.


### GetConfigUrl

`func (o *AuthenticatorTOTPChallenge) GetConfigUrl() string`

GetConfigUrl returns the ConfigUrl field if non-nil, zero value otherwise.

### GetConfigUrlOk

`func (o *AuthenticatorTOTPChallenge) GetConfigUrlOk() (*string, bool)`

GetConfigUrlOk returns a tuple with the ConfigUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigUrl

`func (o *AuthenticatorTOTPChallenge) SetConfigUrl(v string)`

SetConfigUrl sets ConfigUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


