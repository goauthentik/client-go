# AuthenticatorEmailChallengeResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-authenticator-email"]
**Code** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthenticatorEmailChallengeResponseRequest

`func NewAuthenticatorEmailChallengeResponseRequest() *AuthenticatorEmailChallengeResponseRequest`

NewAuthenticatorEmailChallengeResponseRequest instantiates a new AuthenticatorEmailChallengeResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorEmailChallengeResponseRequestWithDefaults

`func NewAuthenticatorEmailChallengeResponseRequestWithDefaults() *AuthenticatorEmailChallengeResponseRequest`

NewAuthenticatorEmailChallengeResponseRequestWithDefaults instantiates a new AuthenticatorEmailChallengeResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *AuthenticatorEmailChallengeResponseRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AuthenticatorEmailChallengeResponseRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AuthenticatorEmailChallengeResponseRequest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AuthenticatorEmailChallengeResponseRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetCode

`func (o *AuthenticatorEmailChallengeResponseRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuthenticatorEmailChallengeResponseRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuthenticatorEmailChallengeResponseRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AuthenticatorEmailChallengeResponseRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetEmail

`func (o *AuthenticatorEmailChallengeResponseRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthenticatorEmailChallengeResponseRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthenticatorEmailChallengeResponseRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthenticatorEmailChallengeResponseRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


