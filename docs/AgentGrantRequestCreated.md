# AgentGrantRequestCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantRequest** | [**GrantRequest**](GrantRequest.md) |  | [readonly] 
**FulfillUrl** | **string** |  | [readonly] 

## Methods

### NewAgentGrantRequestCreated

`func NewAgentGrantRequestCreated(grantRequest GrantRequest, fulfillUrl string, ) *AgentGrantRequestCreated`

NewAgentGrantRequestCreated instantiates a new AgentGrantRequestCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentGrantRequestCreatedWithDefaults

`func NewAgentGrantRequestCreatedWithDefaults() *AgentGrantRequestCreated`

NewAgentGrantRequestCreatedWithDefaults instantiates a new AgentGrantRequestCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantRequest

`func (o *AgentGrantRequestCreated) GetGrantRequest() GrantRequest`

GetGrantRequest returns the GrantRequest field if non-nil, zero value otherwise.

### GetGrantRequestOk

`func (o *AgentGrantRequestCreated) GetGrantRequestOk() (*GrantRequest, bool)`

GetGrantRequestOk returns a tuple with the GrantRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantRequest

`func (o *AgentGrantRequestCreated) SetGrantRequest(v GrantRequest)`

SetGrantRequest sets GrantRequest field to given value.


### GetFulfillUrl

`func (o *AgentGrantRequestCreated) GetFulfillUrl() string`

GetFulfillUrl returns the FulfillUrl field if non-nil, zero value otherwise.

### GetFulfillUrlOk

`func (o *AgentGrantRequestCreated) GetFulfillUrlOk() (*string, bool)`

GetFulfillUrlOk returns a tuple with the FulfillUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillUrl

`func (o *AgentGrantRequestCreated) SetFulfillUrl(v string)`

SetFulfillUrl sets FulfillUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


