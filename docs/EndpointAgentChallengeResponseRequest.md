# EndpointAgentChallengeResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-endpoint-agent"]
**Response** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEndpointAgentChallengeResponseRequest

`func NewEndpointAgentChallengeResponseRequest() *EndpointAgentChallengeResponseRequest`

NewEndpointAgentChallengeResponseRequest instantiates a new EndpointAgentChallengeResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointAgentChallengeResponseRequestWithDefaults

`func NewEndpointAgentChallengeResponseRequestWithDefaults() *EndpointAgentChallengeResponseRequest`

NewEndpointAgentChallengeResponseRequestWithDefaults instantiates a new EndpointAgentChallengeResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *EndpointAgentChallengeResponseRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *EndpointAgentChallengeResponseRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *EndpointAgentChallengeResponseRequest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *EndpointAgentChallengeResponseRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponse

`func (o *EndpointAgentChallengeResponseRequest) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *EndpointAgentChallengeResponseRequest) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *EndpointAgentChallengeResponseRequest) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *EndpointAgentChallengeResponseRequest) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *EndpointAgentChallengeResponseRequest) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *EndpointAgentChallengeResponseRequest) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


