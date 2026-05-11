# AuthenticatorValidationChallengeResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-authenticator-validate"]
**SelectedChallenge** | Pointer to [**DeviceChallengeRequest**](DeviceChallengeRequest.md) |  | [optional] 
**SelectedStage** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Webauthn** | Pointer to **map[string]interface{}** |  | [optional] 
**Duo** | Pointer to **int32** |  | [optional] 

## Methods

### NewAuthenticatorValidationChallengeResponseRequest

`func NewAuthenticatorValidationChallengeResponseRequest() *AuthenticatorValidationChallengeResponseRequest`

NewAuthenticatorValidationChallengeResponseRequest instantiates a new AuthenticatorValidationChallengeResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorValidationChallengeResponseRequestWithDefaults

`func NewAuthenticatorValidationChallengeResponseRequestWithDefaults() *AuthenticatorValidationChallengeResponseRequest`

NewAuthenticatorValidationChallengeResponseRequestWithDefaults instantiates a new AuthenticatorValidationChallengeResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *AuthenticatorValidationChallengeResponseRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AuthenticatorValidationChallengeResponseRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AuthenticatorValidationChallengeResponseRequest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AuthenticatorValidationChallengeResponseRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetSelectedChallenge

`func (o *AuthenticatorValidationChallengeResponseRequest) GetSelectedChallenge() DeviceChallengeRequest`

GetSelectedChallenge returns the SelectedChallenge field if non-nil, zero value otherwise.

### GetSelectedChallengeOk

`func (o *AuthenticatorValidationChallengeResponseRequest) GetSelectedChallengeOk() (*DeviceChallengeRequest, bool)`

GetSelectedChallengeOk returns a tuple with the SelectedChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedChallenge

`func (o *AuthenticatorValidationChallengeResponseRequest) SetSelectedChallenge(v DeviceChallengeRequest)`

SetSelectedChallenge sets SelectedChallenge field to given value.

### HasSelectedChallenge

`func (o *AuthenticatorValidationChallengeResponseRequest) HasSelectedChallenge() bool`

HasSelectedChallenge returns a boolean if a field has been set.

### GetSelectedStage

`func (o *AuthenticatorValidationChallengeResponseRequest) GetSelectedStage() string`

GetSelectedStage returns the SelectedStage field if non-nil, zero value otherwise.

### GetSelectedStageOk

`func (o *AuthenticatorValidationChallengeResponseRequest) GetSelectedStageOk() (*string, bool)`

GetSelectedStageOk returns a tuple with the SelectedStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedStage

`func (o *AuthenticatorValidationChallengeResponseRequest) SetSelectedStage(v string)`

SetSelectedStage sets SelectedStage field to given value.

### HasSelectedStage

`func (o *AuthenticatorValidationChallengeResponseRequest) HasSelectedStage() bool`

HasSelectedStage returns a boolean if a field has been set.

### GetCode

`func (o *AuthenticatorValidationChallengeResponseRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuthenticatorValidationChallengeResponseRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuthenticatorValidationChallengeResponseRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AuthenticatorValidationChallengeResponseRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetWebauthn

`func (o *AuthenticatorValidationChallengeResponseRequest) GetWebauthn() map[string]interface{}`

GetWebauthn returns the Webauthn field if non-nil, zero value otherwise.

### GetWebauthnOk

`func (o *AuthenticatorValidationChallengeResponseRequest) GetWebauthnOk() (*map[string]interface{}, bool)`

GetWebauthnOk returns a tuple with the Webauthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebauthn

`func (o *AuthenticatorValidationChallengeResponseRequest) SetWebauthn(v map[string]interface{})`

SetWebauthn sets Webauthn field to given value.

### HasWebauthn

`func (o *AuthenticatorValidationChallengeResponseRequest) HasWebauthn() bool`

HasWebauthn returns a boolean if a field has been set.

### GetDuo

`func (o *AuthenticatorValidationChallengeResponseRequest) GetDuo() int32`

GetDuo returns the Duo field if non-nil, zero value otherwise.

### GetDuoOk

`func (o *AuthenticatorValidationChallengeResponseRequest) GetDuoOk() (*int32, bool)`

GetDuoOk returns a tuple with the Duo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuo

`func (o *AuthenticatorValidationChallengeResponseRequest) SetDuo(v int32)`

SetDuo sets Duo field to given value.

### HasDuo

`func (o *AuthenticatorValidationChallengeResponseRequest) HasDuo() bool`

HasDuo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


