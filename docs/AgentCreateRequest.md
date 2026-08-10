# AgentCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parent** | Pointer to **int32** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Expiring** | Pointer to **bool** |  | [optional] [default to false]
**Expires** | Pointer to **NullableTime** |  | [optional] 
**PolicyBehavior** | Pointer to [**PolicyBehaviorEnum**](PolicyBehaviorEnum.md) |  | [optional] [default to POLICYBEHAVIORENUM_MIRROR]

## Methods

### NewAgentCreateRequest

`func NewAgentCreateRequest() *AgentCreateRequest`

NewAgentCreateRequest instantiates a new AgentCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCreateRequestWithDefaults

`func NewAgentCreateRequestWithDefaults() *AgentCreateRequest`

NewAgentCreateRequestWithDefaults instantiates a new AgentCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParent

`func (o *AgentCreateRequest) GetParent() int32`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AgentCreateRequest) GetParentOk() (*int32, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AgentCreateRequest) SetParent(v int32)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AgentCreateRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetLabel

`func (o *AgentCreateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AgentCreateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AgentCreateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AgentCreateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetExpiring

`func (o *AgentCreateRequest) GetExpiring() bool`

GetExpiring returns the Expiring field if non-nil, zero value otherwise.

### GetExpiringOk

`func (o *AgentCreateRequest) GetExpiringOk() (*bool, bool)`

GetExpiringOk returns a tuple with the Expiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiring

`func (o *AgentCreateRequest) SetExpiring(v bool)`

SetExpiring sets Expiring field to given value.

### HasExpiring

`func (o *AgentCreateRequest) HasExpiring() bool`

HasExpiring returns a boolean if a field has been set.

### GetExpires

`func (o *AgentCreateRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *AgentCreateRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *AgentCreateRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *AgentCreateRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *AgentCreateRequest) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *AgentCreateRequest) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetPolicyBehavior

`func (o *AgentCreateRequest) GetPolicyBehavior() PolicyBehaviorEnum`

GetPolicyBehavior returns the PolicyBehavior field if non-nil, zero value otherwise.

### GetPolicyBehaviorOk

`func (o *AgentCreateRequest) GetPolicyBehaviorOk() (*PolicyBehaviorEnum, bool)`

GetPolicyBehaviorOk returns a tuple with the PolicyBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyBehavior

`func (o *AgentCreateRequest) SetPolicyBehavior(v PolicyBehaviorEnum)`

SetPolicyBehavior sets PolicyBehavior field to given value.

### HasPolicyBehavior

`func (o *AgentCreateRequest) HasPolicyBehavior() bool`

HasPolicyBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


