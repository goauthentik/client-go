# UserLoginChallengeResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-user-login"]
**RememberMe** | **bool** |  | 

## Methods

### NewUserLoginChallengeResponseRequest

`func NewUserLoginChallengeResponseRequest(rememberMe bool, ) *UserLoginChallengeResponseRequest`

NewUserLoginChallengeResponseRequest instantiates a new UserLoginChallengeResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLoginChallengeResponseRequestWithDefaults

`func NewUserLoginChallengeResponseRequestWithDefaults() *UserLoginChallengeResponseRequest`

NewUserLoginChallengeResponseRequestWithDefaults instantiates a new UserLoginChallengeResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *UserLoginChallengeResponseRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *UserLoginChallengeResponseRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *UserLoginChallengeResponseRequest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *UserLoginChallengeResponseRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetRememberMe

`func (o *UserLoginChallengeResponseRequest) GetRememberMe() bool`

GetRememberMe returns the RememberMe field if non-nil, zero value otherwise.

### GetRememberMeOk

`func (o *UserLoginChallengeResponseRequest) GetRememberMeOk() (*bool, bool)`

GetRememberMeOk returns a tuple with the RememberMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMe

`func (o *UserLoginChallengeResponseRequest) SetRememberMe(v bool)`

SetRememberMe sets RememberMe field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


