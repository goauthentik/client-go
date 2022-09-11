# AuthenticatorSMSChallengeResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-authenticator-sms"]
**Code** | Pointer to **int32** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthenticatorSMSChallengeResponseRequest

`func NewAuthenticatorSMSChallengeResponseRequest() *AuthenticatorSMSChallengeResponseRequest`

NewAuthenticatorSMSChallengeResponseRequest instantiates a new AuthenticatorSMSChallengeResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorSMSChallengeResponseRequestWithDefaults

`func NewAuthenticatorSMSChallengeResponseRequestWithDefaults() *AuthenticatorSMSChallengeResponseRequest`

NewAuthenticatorSMSChallengeResponseRequestWithDefaults instantiates a new AuthenticatorSMSChallengeResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *AuthenticatorSMSChallengeResponseRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AuthenticatorSMSChallengeResponseRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AuthenticatorSMSChallengeResponseRequest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AuthenticatorSMSChallengeResponseRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetCode

`func (o *AuthenticatorSMSChallengeResponseRequest) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuthenticatorSMSChallengeResponseRequest) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuthenticatorSMSChallengeResponseRequest) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *AuthenticatorSMSChallengeResponseRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *AuthenticatorSMSChallengeResponseRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *AuthenticatorSMSChallengeResponseRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *AuthenticatorSMSChallengeResponseRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *AuthenticatorSMSChallengeResponseRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


