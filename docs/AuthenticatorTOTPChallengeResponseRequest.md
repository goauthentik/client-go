# AuthenticatorTOTPChallengeResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-authenticator-totp"]
**Code** | **int32** |  | 

## Methods

### NewAuthenticatorTOTPChallengeResponseRequest

`func NewAuthenticatorTOTPChallengeResponseRequest(code int32, ) *AuthenticatorTOTPChallengeResponseRequest`

NewAuthenticatorTOTPChallengeResponseRequest instantiates a new AuthenticatorTOTPChallengeResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorTOTPChallengeResponseRequestWithDefaults

`func NewAuthenticatorTOTPChallengeResponseRequestWithDefaults() *AuthenticatorTOTPChallengeResponseRequest`

NewAuthenticatorTOTPChallengeResponseRequestWithDefaults instantiates a new AuthenticatorTOTPChallengeResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *AuthenticatorTOTPChallengeResponseRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AuthenticatorTOTPChallengeResponseRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AuthenticatorTOTPChallengeResponseRequest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AuthenticatorTOTPChallengeResponseRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetCode

`func (o *AuthenticatorTOTPChallengeResponseRequest) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuthenticatorTOTPChallengeResponseRequest) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuthenticatorTOTPChallengeResponseRequest) SetCode(v int32)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


