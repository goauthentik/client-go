# AuthenticatorDuoChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ChallengeChoices**](ChallengeChoices.md) |  | 
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-authenticator-duo"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**PendingUser** | **string** |  | 
**PendingUserAvatar** | **string** |  | 
**ActivationBarcode** | **string** |  | 
**ActivationCode** | **string** |  | 
**StageUuid** | **string** |  | 

## Methods

### NewAuthenticatorDuoChallenge

`func NewAuthenticatorDuoChallenge(type_ ChallengeChoices, pendingUser string, pendingUserAvatar string, activationBarcode string, activationCode string, stageUuid string, ) *AuthenticatorDuoChallenge`

NewAuthenticatorDuoChallenge instantiates a new AuthenticatorDuoChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorDuoChallengeWithDefaults

`func NewAuthenticatorDuoChallengeWithDefaults() *AuthenticatorDuoChallenge`

NewAuthenticatorDuoChallengeWithDefaults instantiates a new AuthenticatorDuoChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthenticatorDuoChallenge) GetType() ChallengeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthenticatorDuoChallenge) GetTypeOk() (*ChallengeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthenticatorDuoChallenge) SetType(v ChallengeChoices)`

SetType sets Type field to given value.


### GetFlowInfo

`func (o *AuthenticatorDuoChallenge) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *AuthenticatorDuoChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *AuthenticatorDuoChallenge) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *AuthenticatorDuoChallenge) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *AuthenticatorDuoChallenge) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AuthenticatorDuoChallenge) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AuthenticatorDuoChallenge) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AuthenticatorDuoChallenge) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *AuthenticatorDuoChallenge) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *AuthenticatorDuoChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *AuthenticatorDuoChallenge) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *AuthenticatorDuoChallenge) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetPendingUser

`func (o *AuthenticatorDuoChallenge) GetPendingUser() string`

GetPendingUser returns the PendingUser field if non-nil, zero value otherwise.

### GetPendingUserOk

`func (o *AuthenticatorDuoChallenge) GetPendingUserOk() (*string, bool)`

GetPendingUserOk returns a tuple with the PendingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUser

`func (o *AuthenticatorDuoChallenge) SetPendingUser(v string)`

SetPendingUser sets PendingUser field to given value.


### GetPendingUserAvatar

`func (o *AuthenticatorDuoChallenge) GetPendingUserAvatar() string`

GetPendingUserAvatar returns the PendingUserAvatar field if non-nil, zero value otherwise.

### GetPendingUserAvatarOk

`func (o *AuthenticatorDuoChallenge) GetPendingUserAvatarOk() (*string, bool)`

GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUserAvatar

`func (o *AuthenticatorDuoChallenge) SetPendingUserAvatar(v string)`

SetPendingUserAvatar sets PendingUserAvatar field to given value.


### GetActivationBarcode

`func (o *AuthenticatorDuoChallenge) GetActivationBarcode() string`

GetActivationBarcode returns the ActivationBarcode field if non-nil, zero value otherwise.

### GetActivationBarcodeOk

`func (o *AuthenticatorDuoChallenge) GetActivationBarcodeOk() (*string, bool)`

GetActivationBarcodeOk returns a tuple with the ActivationBarcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationBarcode

`func (o *AuthenticatorDuoChallenge) SetActivationBarcode(v string)`

SetActivationBarcode sets ActivationBarcode field to given value.


### GetActivationCode

`func (o *AuthenticatorDuoChallenge) GetActivationCode() string`

GetActivationCode returns the ActivationCode field if non-nil, zero value otherwise.

### GetActivationCodeOk

`func (o *AuthenticatorDuoChallenge) GetActivationCodeOk() (*string, bool)`

GetActivationCodeOk returns a tuple with the ActivationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationCode

`func (o *AuthenticatorDuoChallenge) SetActivationCode(v string)`

SetActivationCode sets ActivationCode field to given value.


### GetStageUuid

`func (o *AuthenticatorDuoChallenge) GetStageUuid() string`

GetStageUuid returns the StageUuid field if non-nil, zero value otherwise.

### GetStageUuidOk

`func (o *AuthenticatorDuoChallenge) GetStageUuidOk() (*string, bool)`

GetStageUuidOk returns a tuple with the StageUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageUuid

`func (o *AuthenticatorDuoChallenge) SetStageUuid(v string)`

SetStageUuid sets StageUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


