# RedirectChallengeResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] [default to "xak-flow-redirect"]
**To** | **string** |  | 

## Methods

### NewRedirectChallengeResponseRequest

`func NewRedirectChallengeResponseRequest(to string, ) *RedirectChallengeResponseRequest`

NewRedirectChallengeResponseRequest instantiates a new RedirectChallengeResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedirectChallengeResponseRequestWithDefaults

`func NewRedirectChallengeResponseRequestWithDefaults() *RedirectChallengeResponseRequest`

NewRedirectChallengeResponseRequestWithDefaults instantiates a new RedirectChallengeResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *RedirectChallengeResponseRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *RedirectChallengeResponseRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *RedirectChallengeResponseRequest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *RedirectChallengeResponseRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetTo

`func (o *RedirectChallengeResponseRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RedirectChallengeResponseRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RedirectChallengeResponseRequest) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


