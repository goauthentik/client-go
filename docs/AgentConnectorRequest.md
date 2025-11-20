# AgentConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorUuid** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**SnapshotExpiry** | Pointer to **string** |  | [optional] 
**NssUidOffset** | Pointer to **int32** |  | [optional] 
**NssGidOffset** | Pointer to **int32** |  | [optional] 
**AuthTerminateSessionOnExpiry** | Pointer to **bool** |  | [optional] 
**RefreshInterval** | Pointer to **string** |  | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** |  | [optional] 
**ChallengeKey** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAgentConnectorRequest

`func NewAgentConnectorRequest(name string, ) *AgentConnectorRequest`

NewAgentConnectorRequest instantiates a new AgentConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentConnectorRequestWithDefaults

`func NewAgentConnectorRequestWithDefaults() *AgentConnectorRequest`

NewAgentConnectorRequestWithDefaults instantiates a new AgentConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorUuid

`func (o *AgentConnectorRequest) GetConnectorUuid() string`

GetConnectorUuid returns the ConnectorUuid field if non-nil, zero value otherwise.

### GetConnectorUuidOk

`func (o *AgentConnectorRequest) GetConnectorUuidOk() (*string, bool)`

GetConnectorUuidOk returns a tuple with the ConnectorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorUuid

`func (o *AgentConnectorRequest) SetConnectorUuid(v string)`

SetConnectorUuid sets ConnectorUuid field to given value.

### HasConnectorUuid

`func (o *AgentConnectorRequest) HasConnectorUuid() bool`

HasConnectorUuid returns a boolean if a field has been set.

### GetName

`func (o *AgentConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentConnectorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *AgentConnectorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgentConnectorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgentConnectorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AgentConnectorRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSnapshotExpiry

`func (o *AgentConnectorRequest) GetSnapshotExpiry() string`

GetSnapshotExpiry returns the SnapshotExpiry field if non-nil, zero value otherwise.

### GetSnapshotExpiryOk

`func (o *AgentConnectorRequest) GetSnapshotExpiryOk() (*string, bool)`

GetSnapshotExpiryOk returns a tuple with the SnapshotExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotExpiry

`func (o *AgentConnectorRequest) SetSnapshotExpiry(v string)`

SetSnapshotExpiry sets SnapshotExpiry field to given value.

### HasSnapshotExpiry

`func (o *AgentConnectorRequest) HasSnapshotExpiry() bool`

HasSnapshotExpiry returns a boolean if a field has been set.

### GetNssUidOffset

`func (o *AgentConnectorRequest) GetNssUidOffset() int32`

GetNssUidOffset returns the NssUidOffset field if non-nil, zero value otherwise.

### GetNssUidOffsetOk

`func (o *AgentConnectorRequest) GetNssUidOffsetOk() (*int32, bool)`

GetNssUidOffsetOk returns a tuple with the NssUidOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssUidOffset

`func (o *AgentConnectorRequest) SetNssUidOffset(v int32)`

SetNssUidOffset sets NssUidOffset field to given value.

### HasNssUidOffset

`func (o *AgentConnectorRequest) HasNssUidOffset() bool`

HasNssUidOffset returns a boolean if a field has been set.

### GetNssGidOffset

`func (o *AgentConnectorRequest) GetNssGidOffset() int32`

GetNssGidOffset returns the NssGidOffset field if non-nil, zero value otherwise.

### GetNssGidOffsetOk

`func (o *AgentConnectorRequest) GetNssGidOffsetOk() (*int32, bool)`

GetNssGidOffsetOk returns a tuple with the NssGidOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssGidOffset

`func (o *AgentConnectorRequest) SetNssGidOffset(v int32)`

SetNssGidOffset sets NssGidOffset field to given value.

### HasNssGidOffset

`func (o *AgentConnectorRequest) HasNssGidOffset() bool`

HasNssGidOffset returns a boolean if a field has been set.

### GetAuthTerminateSessionOnExpiry

`func (o *AgentConnectorRequest) GetAuthTerminateSessionOnExpiry() bool`

GetAuthTerminateSessionOnExpiry returns the AuthTerminateSessionOnExpiry field if non-nil, zero value otherwise.

### GetAuthTerminateSessionOnExpiryOk

`func (o *AgentConnectorRequest) GetAuthTerminateSessionOnExpiryOk() (*bool, bool)`

GetAuthTerminateSessionOnExpiryOk returns a tuple with the AuthTerminateSessionOnExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTerminateSessionOnExpiry

`func (o *AgentConnectorRequest) SetAuthTerminateSessionOnExpiry(v bool)`

SetAuthTerminateSessionOnExpiry sets AuthTerminateSessionOnExpiry field to given value.

### HasAuthTerminateSessionOnExpiry

`func (o *AgentConnectorRequest) HasAuthTerminateSessionOnExpiry() bool`

HasAuthTerminateSessionOnExpiry returns a boolean if a field has been set.

### GetRefreshInterval

`func (o *AgentConnectorRequest) GetRefreshInterval() string`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *AgentConnectorRequest) GetRefreshIntervalOk() (*string, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *AgentConnectorRequest) SetRefreshInterval(v string)`

SetRefreshInterval sets RefreshInterval field to given value.

### HasRefreshInterval

`func (o *AgentConnectorRequest) HasRefreshInterval() bool`

HasRefreshInterval returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *AgentConnectorRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *AgentConnectorRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *AgentConnectorRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *AgentConnectorRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *AgentConnectorRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *AgentConnectorRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetChallengeKey

`func (o *AgentConnectorRequest) GetChallengeKey() string`

GetChallengeKey returns the ChallengeKey field if non-nil, zero value otherwise.

### GetChallengeKeyOk

`func (o *AgentConnectorRequest) GetChallengeKeyOk() (*string, bool)`

GetChallengeKeyOk returns a tuple with the ChallengeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeKey

`func (o *AgentConnectorRequest) SetChallengeKey(v string)`

SetChallengeKey sets ChallengeKey field to given value.

### HasChallengeKey

`func (o *AgentConnectorRequest) HasChallengeKey() bool`

HasChallengeKey returns a boolean if a field has been set.

### SetChallengeKeyNil

`func (o *AgentConnectorRequest) SetChallengeKeyNil(b bool)`

 SetChallengeKeyNil sets the value for ChallengeKey to be an explicit nil

### UnsetChallengeKey
`func (o *AgentConnectorRequest) UnsetChallengeKey()`

UnsetChallengeKey ensures that no value is present for ChallengeKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


