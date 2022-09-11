# AuthenticatorSMSChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ChallengeChoices**](ChallengeChoices.md) |  | 
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-authenticator-sms"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**PendingUser** | **string** |  | 
**PendingUserAvatar** | **string** |  | 
**PhoneNumberRequired** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewAuthenticatorSMSChallenge

`func NewAuthenticatorSMSChallenge(type_ ChallengeChoices, pendingUser string, pendingUserAvatar string, ) *AuthenticatorSMSChallenge`

NewAuthenticatorSMSChallenge instantiates a new AuthenticatorSMSChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorSMSChallengeWithDefaults

`func NewAuthenticatorSMSChallengeWithDefaults() *AuthenticatorSMSChallenge`

NewAuthenticatorSMSChallengeWithDefaults instantiates a new AuthenticatorSMSChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthenticatorSMSChallenge) GetType() ChallengeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthenticatorSMSChallenge) GetTypeOk() (*ChallengeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthenticatorSMSChallenge) SetType(v ChallengeChoices)`

SetType sets Type field to given value.


### GetFlowInfo

`func (o *AuthenticatorSMSChallenge) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *AuthenticatorSMSChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *AuthenticatorSMSChallenge) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *AuthenticatorSMSChallenge) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *AuthenticatorSMSChallenge) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AuthenticatorSMSChallenge) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AuthenticatorSMSChallenge) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AuthenticatorSMSChallenge) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *AuthenticatorSMSChallenge) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *AuthenticatorSMSChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *AuthenticatorSMSChallenge) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *AuthenticatorSMSChallenge) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetPendingUser

`func (o *AuthenticatorSMSChallenge) GetPendingUser() string`

GetPendingUser returns the PendingUser field if non-nil, zero value otherwise.

### GetPendingUserOk

`func (o *AuthenticatorSMSChallenge) GetPendingUserOk() (*string, bool)`

GetPendingUserOk returns a tuple with the PendingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUser

`func (o *AuthenticatorSMSChallenge) SetPendingUser(v string)`

SetPendingUser sets PendingUser field to given value.


### GetPendingUserAvatar

`func (o *AuthenticatorSMSChallenge) GetPendingUserAvatar() string`

GetPendingUserAvatar returns the PendingUserAvatar field if non-nil, zero value otherwise.

### GetPendingUserAvatarOk

`func (o *AuthenticatorSMSChallenge) GetPendingUserAvatarOk() (*string, bool)`

GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUserAvatar

`func (o *AuthenticatorSMSChallenge) SetPendingUserAvatar(v string)`

SetPendingUserAvatar sets PendingUserAvatar field to given value.


### GetPhoneNumberRequired

`func (o *AuthenticatorSMSChallenge) GetPhoneNumberRequired() bool`

GetPhoneNumberRequired returns the PhoneNumberRequired field if non-nil, zero value otherwise.

### GetPhoneNumberRequiredOk

`func (o *AuthenticatorSMSChallenge) GetPhoneNumberRequiredOk() (*bool, bool)`

GetPhoneNumberRequiredOk returns a tuple with the PhoneNumberRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberRequired

`func (o *AuthenticatorSMSChallenge) SetPhoneNumberRequired(v bool)`

SetPhoneNumberRequired sets PhoneNumberRequired field to given value.

### HasPhoneNumberRequired

`func (o *AuthenticatorSMSChallenge) HasPhoneNumberRequired() bool`

HasPhoneNumberRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


