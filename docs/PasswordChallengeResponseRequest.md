# PasswordChallengeResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-password"]
**Password** | **string** |  | 

## Methods

### NewPasswordChallengeResponseRequest

`func NewPasswordChallengeResponseRequest(password string, ) *PasswordChallengeResponseRequest`

NewPasswordChallengeResponseRequest instantiates a new PasswordChallengeResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordChallengeResponseRequestWithDefaults

`func NewPasswordChallengeResponseRequestWithDefaults() *PasswordChallengeResponseRequest`

NewPasswordChallengeResponseRequestWithDefaults instantiates a new PasswordChallengeResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *PasswordChallengeResponseRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *PasswordChallengeResponseRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *PasswordChallengeResponseRequest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *PasswordChallengeResponseRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetPassword

`func (o *PasswordChallengeResponseRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PasswordChallengeResponseRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PasswordChallengeResponseRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


