# AuthenticatorValidationChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ChallengeChoices**](ChallengeChoices.md) |  | 
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-authenticator-validate"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**PendingUser** | **string** |  | 
**PendingUserAvatar** | **string** |  | 
**DeviceChallenges** | [**[]DeviceChallenge**](DeviceChallenge.md) |  | 
**ConfigurationStages** | [**[]SelectableStage**](SelectableStage.md) |  | 

## Methods

### NewAuthenticatorValidationChallenge

`func NewAuthenticatorValidationChallenge(type_ ChallengeChoices, pendingUser string, pendingUserAvatar string, deviceChallenges []DeviceChallenge, configurationStages []SelectableStage, ) *AuthenticatorValidationChallenge`

NewAuthenticatorValidationChallenge instantiates a new AuthenticatorValidationChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorValidationChallengeWithDefaults

`func NewAuthenticatorValidationChallengeWithDefaults() *AuthenticatorValidationChallenge`

NewAuthenticatorValidationChallengeWithDefaults instantiates a new AuthenticatorValidationChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthenticatorValidationChallenge) GetType() ChallengeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthenticatorValidationChallenge) GetTypeOk() (*ChallengeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthenticatorValidationChallenge) SetType(v ChallengeChoices)`

SetType sets Type field to given value.


### GetFlowInfo

`func (o *AuthenticatorValidationChallenge) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *AuthenticatorValidationChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *AuthenticatorValidationChallenge) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *AuthenticatorValidationChallenge) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *AuthenticatorValidationChallenge) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AuthenticatorValidationChallenge) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AuthenticatorValidationChallenge) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AuthenticatorValidationChallenge) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *AuthenticatorValidationChallenge) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *AuthenticatorValidationChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *AuthenticatorValidationChallenge) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *AuthenticatorValidationChallenge) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetPendingUser

`func (o *AuthenticatorValidationChallenge) GetPendingUser() string`

GetPendingUser returns the PendingUser field if non-nil, zero value otherwise.

### GetPendingUserOk

`func (o *AuthenticatorValidationChallenge) GetPendingUserOk() (*string, bool)`

GetPendingUserOk returns a tuple with the PendingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUser

`func (o *AuthenticatorValidationChallenge) SetPendingUser(v string)`

SetPendingUser sets PendingUser field to given value.


### GetPendingUserAvatar

`func (o *AuthenticatorValidationChallenge) GetPendingUserAvatar() string`

GetPendingUserAvatar returns the PendingUserAvatar field if non-nil, zero value otherwise.

### GetPendingUserAvatarOk

`func (o *AuthenticatorValidationChallenge) GetPendingUserAvatarOk() (*string, bool)`

GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUserAvatar

`func (o *AuthenticatorValidationChallenge) SetPendingUserAvatar(v string)`

SetPendingUserAvatar sets PendingUserAvatar field to given value.


### GetDeviceChallenges

`func (o *AuthenticatorValidationChallenge) GetDeviceChallenges() []DeviceChallenge`

GetDeviceChallenges returns the DeviceChallenges field if non-nil, zero value otherwise.

### GetDeviceChallengesOk

`func (o *AuthenticatorValidationChallenge) GetDeviceChallengesOk() (*[]DeviceChallenge, bool)`

GetDeviceChallengesOk returns a tuple with the DeviceChallenges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceChallenges

`func (o *AuthenticatorValidationChallenge) SetDeviceChallenges(v []DeviceChallenge)`

SetDeviceChallenges sets DeviceChallenges field to given value.


### GetConfigurationStages

`func (o *AuthenticatorValidationChallenge) GetConfigurationStages() []SelectableStage`

GetConfigurationStages returns the ConfigurationStages field if non-nil, zero value otherwise.

### GetConfigurationStagesOk

`func (o *AuthenticatorValidationChallenge) GetConfigurationStagesOk() (*[]SelectableStage, bool)`

GetConfigurationStagesOk returns a tuple with the ConfigurationStages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationStages

`func (o *AuthenticatorValidationChallenge) SetConfigurationStages(v []SelectableStage)`

SetConfigurationStages sets ConfigurationStages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


