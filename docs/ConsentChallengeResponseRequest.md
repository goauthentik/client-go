# ConsentChallengeResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-consent"]
**Token** | **string** |  | 

## Methods

### NewConsentChallengeResponseRequest

`func NewConsentChallengeResponseRequest(token string, ) *ConsentChallengeResponseRequest`

NewConsentChallengeResponseRequest instantiates a new ConsentChallengeResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentChallengeResponseRequestWithDefaults

`func NewConsentChallengeResponseRequestWithDefaults() *ConsentChallengeResponseRequest`

NewConsentChallengeResponseRequestWithDefaults instantiates a new ConsentChallengeResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *ConsentChallengeResponseRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ConsentChallengeResponseRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ConsentChallengeResponseRequest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ConsentChallengeResponseRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetToken

`func (o *ConsentChallengeResponseRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ConsentChallengeResponseRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ConsentChallengeResponseRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


