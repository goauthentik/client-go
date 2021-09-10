# AuthenticatorWebAuthnChallengeResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-authenticator-webauthn"]
**Response** | **map[string]interface{}** |  | 

## Methods

### NewAuthenticatorWebAuthnChallengeResponseRequest

`func NewAuthenticatorWebAuthnChallengeResponseRequest(response map[string]interface{}, ) *AuthenticatorWebAuthnChallengeResponseRequest`

NewAuthenticatorWebAuthnChallengeResponseRequest instantiates a new AuthenticatorWebAuthnChallengeResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorWebAuthnChallengeResponseRequestWithDefaults

`func NewAuthenticatorWebAuthnChallengeResponseRequestWithDefaults() *AuthenticatorWebAuthnChallengeResponseRequest`

NewAuthenticatorWebAuthnChallengeResponseRequestWithDefaults instantiates a new AuthenticatorWebAuthnChallengeResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *AuthenticatorWebAuthnChallengeResponseRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AuthenticatorWebAuthnChallengeResponseRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AuthenticatorWebAuthnChallengeResponseRequest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AuthenticatorWebAuthnChallengeResponseRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponse

`func (o *AuthenticatorWebAuthnChallengeResponseRequest) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *AuthenticatorWebAuthnChallengeResponseRequest) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *AuthenticatorWebAuthnChallengeResponseRequest) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


