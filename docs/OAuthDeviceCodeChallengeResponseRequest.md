# OAuthDeviceCodeChallengeResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] [default to "ak-provider-oauth2-device-code"]
**Code** | **int32** |  | 

## Methods

### NewOAuthDeviceCodeChallengeResponseRequest

`func NewOAuthDeviceCodeChallengeResponseRequest(code int32, ) *OAuthDeviceCodeChallengeResponseRequest`

NewOAuthDeviceCodeChallengeResponseRequest instantiates a new OAuthDeviceCodeChallengeResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthDeviceCodeChallengeResponseRequestWithDefaults

`func NewOAuthDeviceCodeChallengeResponseRequestWithDefaults() *OAuthDeviceCodeChallengeResponseRequest`

NewOAuthDeviceCodeChallengeResponseRequestWithDefaults instantiates a new OAuthDeviceCodeChallengeResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *OAuthDeviceCodeChallengeResponseRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *OAuthDeviceCodeChallengeResponseRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *OAuthDeviceCodeChallengeResponseRequest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *OAuthDeviceCodeChallengeResponseRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetCode

`func (o *OAuthDeviceCodeChallengeResponseRequest) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OAuthDeviceCodeChallengeResponseRequest) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OAuthDeviceCodeChallengeResponseRequest) SetCode(v int32)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


