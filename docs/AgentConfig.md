# AgentConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NssUidOffset** | **int32** |  | 
**NssGidOffset** | **int32** |  | 
**AuthenticationFlow** | **string** |  | 
**AuthTerminateSessionOnExpiry** | **bool** |  | 
**RefreshInterval** | **int32** |  | [readonly] 

## Methods

### NewAgentConfig

`func NewAgentConfig(nssUidOffset int32, nssGidOffset int32, authenticationFlow string, authTerminateSessionOnExpiry bool, refreshInterval int32, ) *AgentConfig`

NewAgentConfig instantiates a new AgentConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentConfigWithDefaults

`func NewAgentConfigWithDefaults() *AgentConfig`

NewAgentConfigWithDefaults instantiates a new AgentConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNssUidOffset

`func (o *AgentConfig) GetNssUidOffset() int32`

GetNssUidOffset returns the NssUidOffset field if non-nil, zero value otherwise.

### GetNssUidOffsetOk

`func (o *AgentConfig) GetNssUidOffsetOk() (*int32, bool)`

GetNssUidOffsetOk returns a tuple with the NssUidOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssUidOffset

`func (o *AgentConfig) SetNssUidOffset(v int32)`

SetNssUidOffset sets NssUidOffset field to given value.


### GetNssGidOffset

`func (o *AgentConfig) GetNssGidOffset() int32`

GetNssGidOffset returns the NssGidOffset field if non-nil, zero value otherwise.

### GetNssGidOffsetOk

`func (o *AgentConfig) GetNssGidOffsetOk() (*int32, bool)`

GetNssGidOffsetOk returns a tuple with the NssGidOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssGidOffset

`func (o *AgentConfig) SetNssGidOffset(v int32)`

SetNssGidOffset sets NssGidOffset field to given value.


### GetAuthenticationFlow

`func (o *AgentConfig) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *AgentConfig) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *AgentConfig) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.


### GetAuthTerminateSessionOnExpiry

`func (o *AgentConfig) GetAuthTerminateSessionOnExpiry() bool`

GetAuthTerminateSessionOnExpiry returns the AuthTerminateSessionOnExpiry field if non-nil, zero value otherwise.

### GetAuthTerminateSessionOnExpiryOk

`func (o *AgentConfig) GetAuthTerminateSessionOnExpiryOk() (*bool, bool)`

GetAuthTerminateSessionOnExpiryOk returns a tuple with the AuthTerminateSessionOnExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTerminateSessionOnExpiry

`func (o *AgentConfig) SetAuthTerminateSessionOnExpiry(v bool)`

SetAuthTerminateSessionOnExpiry sets AuthTerminateSessionOnExpiry field to given value.


### GetRefreshInterval

`func (o *AgentConfig) GetRefreshInterval() int32`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *AgentConfig) GetRefreshIntervalOk() (*int32, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *AgentConfig) SetRefreshInterval(v int32)`

SetRefreshInterval sets RefreshInterval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


