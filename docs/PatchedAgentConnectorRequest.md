# PatchedAgentConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorUuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**SnapshotExpiry** | Pointer to **string** |  | [optional] 
**NssUidOffset** | Pointer to **int32** |  | [optional] 
**NssGidOffset** | Pointer to **int32** |  | [optional] 
**AuthTerminateSessionOnExpiry** | Pointer to **bool** |  | [optional] 
**RefreshInterval** | Pointer to **string** |  | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** |  | [optional] 
**ChallengeKey** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPatchedAgentConnectorRequest

`func NewPatchedAgentConnectorRequest() *PatchedAgentConnectorRequest`

NewPatchedAgentConnectorRequest instantiates a new PatchedAgentConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAgentConnectorRequestWithDefaults

`func NewPatchedAgentConnectorRequestWithDefaults() *PatchedAgentConnectorRequest`

NewPatchedAgentConnectorRequestWithDefaults instantiates a new PatchedAgentConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorUuid

`func (o *PatchedAgentConnectorRequest) GetConnectorUuid() string`

GetConnectorUuid returns the ConnectorUuid field if non-nil, zero value otherwise.

### GetConnectorUuidOk

`func (o *PatchedAgentConnectorRequest) GetConnectorUuidOk() (*string, bool)`

GetConnectorUuidOk returns a tuple with the ConnectorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorUuid

`func (o *PatchedAgentConnectorRequest) SetConnectorUuid(v string)`

SetConnectorUuid sets ConnectorUuid field to given value.

### HasConnectorUuid

`func (o *PatchedAgentConnectorRequest) HasConnectorUuid() bool`

HasConnectorUuid returns a boolean if a field has been set.

### GetName

`func (o *PatchedAgentConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedAgentConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedAgentConnectorRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedAgentConnectorRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedAgentConnectorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedAgentConnectorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedAgentConnectorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedAgentConnectorRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSnapshotExpiry

`func (o *PatchedAgentConnectorRequest) GetSnapshotExpiry() string`

GetSnapshotExpiry returns the SnapshotExpiry field if non-nil, zero value otherwise.

### GetSnapshotExpiryOk

`func (o *PatchedAgentConnectorRequest) GetSnapshotExpiryOk() (*string, bool)`

GetSnapshotExpiryOk returns a tuple with the SnapshotExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotExpiry

`func (o *PatchedAgentConnectorRequest) SetSnapshotExpiry(v string)`

SetSnapshotExpiry sets SnapshotExpiry field to given value.

### HasSnapshotExpiry

`func (o *PatchedAgentConnectorRequest) HasSnapshotExpiry() bool`

HasSnapshotExpiry returns a boolean if a field has been set.

### GetNssUidOffset

`func (o *PatchedAgentConnectorRequest) GetNssUidOffset() int32`

GetNssUidOffset returns the NssUidOffset field if non-nil, zero value otherwise.

### GetNssUidOffsetOk

`func (o *PatchedAgentConnectorRequest) GetNssUidOffsetOk() (*int32, bool)`

GetNssUidOffsetOk returns a tuple with the NssUidOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssUidOffset

`func (o *PatchedAgentConnectorRequest) SetNssUidOffset(v int32)`

SetNssUidOffset sets NssUidOffset field to given value.

### HasNssUidOffset

`func (o *PatchedAgentConnectorRequest) HasNssUidOffset() bool`

HasNssUidOffset returns a boolean if a field has been set.

### GetNssGidOffset

`func (o *PatchedAgentConnectorRequest) GetNssGidOffset() int32`

GetNssGidOffset returns the NssGidOffset field if non-nil, zero value otherwise.

### GetNssGidOffsetOk

`func (o *PatchedAgentConnectorRequest) GetNssGidOffsetOk() (*int32, bool)`

GetNssGidOffsetOk returns a tuple with the NssGidOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssGidOffset

`func (o *PatchedAgentConnectorRequest) SetNssGidOffset(v int32)`

SetNssGidOffset sets NssGidOffset field to given value.

### HasNssGidOffset

`func (o *PatchedAgentConnectorRequest) HasNssGidOffset() bool`

HasNssGidOffset returns a boolean if a field has been set.

### GetAuthTerminateSessionOnExpiry

`func (o *PatchedAgentConnectorRequest) GetAuthTerminateSessionOnExpiry() bool`

GetAuthTerminateSessionOnExpiry returns the AuthTerminateSessionOnExpiry field if non-nil, zero value otherwise.

### GetAuthTerminateSessionOnExpiryOk

`func (o *PatchedAgentConnectorRequest) GetAuthTerminateSessionOnExpiryOk() (*bool, bool)`

GetAuthTerminateSessionOnExpiryOk returns a tuple with the AuthTerminateSessionOnExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTerminateSessionOnExpiry

`func (o *PatchedAgentConnectorRequest) SetAuthTerminateSessionOnExpiry(v bool)`

SetAuthTerminateSessionOnExpiry sets AuthTerminateSessionOnExpiry field to given value.

### HasAuthTerminateSessionOnExpiry

`func (o *PatchedAgentConnectorRequest) HasAuthTerminateSessionOnExpiry() bool`

HasAuthTerminateSessionOnExpiry returns a boolean if a field has been set.

### GetRefreshInterval

`func (o *PatchedAgentConnectorRequest) GetRefreshInterval() string`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *PatchedAgentConnectorRequest) GetRefreshIntervalOk() (*string, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *PatchedAgentConnectorRequest) SetRefreshInterval(v string)`

SetRefreshInterval sets RefreshInterval field to given value.

### HasRefreshInterval

`func (o *PatchedAgentConnectorRequest) HasRefreshInterval() bool`

HasRefreshInterval returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *PatchedAgentConnectorRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *PatchedAgentConnectorRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *PatchedAgentConnectorRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *PatchedAgentConnectorRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *PatchedAgentConnectorRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *PatchedAgentConnectorRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetChallengeKey

`func (o *PatchedAgentConnectorRequest) GetChallengeKey() string`

GetChallengeKey returns the ChallengeKey field if non-nil, zero value otherwise.

### GetChallengeKeyOk

`func (o *PatchedAgentConnectorRequest) GetChallengeKeyOk() (*string, bool)`

GetChallengeKeyOk returns a tuple with the ChallengeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeKey

`func (o *PatchedAgentConnectorRequest) SetChallengeKey(v string)`

SetChallengeKey sets ChallengeKey field to given value.

### HasChallengeKey

`func (o *PatchedAgentConnectorRequest) HasChallengeKey() bool`

HasChallengeKey returns a boolean if a field has been set.

### SetChallengeKeyNil

`func (o *PatchedAgentConnectorRequest) SetChallengeKeyNil(b bool)`

 SetChallengeKeyNil sets the value for ChallengeKey to be an explicit nil

### UnsetChallengeKey
`func (o *PatchedAgentConnectorRequest) UnsetChallengeKey()`

UnsetChallengeKey ensures that no value is present for ChallengeKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


