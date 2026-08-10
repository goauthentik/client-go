# AgentCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | [**Agent**](Agent.md) |  | [readonly] 
**Token** | **string** |  | [readonly] 

## Methods

### NewAgentCreated

`func NewAgentCreated(agent Agent, token string, ) *AgentCreated`

NewAgentCreated instantiates a new AgentCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCreatedWithDefaults

`func NewAgentCreatedWithDefaults() *AgentCreated`

NewAgentCreatedWithDefaults instantiates a new AgentCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *AgentCreated) GetAgent() Agent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentCreated) GetAgentOk() (*Agent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentCreated) SetAgent(v Agent)`

SetAgent sets Agent field to given value.


### GetToken

`func (o *AgentCreated) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AgentCreated) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AgentCreated) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


