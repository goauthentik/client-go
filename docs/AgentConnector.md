# AgentConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorUuid** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Component** | **string** | Get object component so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**SnapshotExpiry** | Pointer to **string** |  | [optional] 
**AuthSessionDuration** | Pointer to **string** |  | [optional] 
**AuthTerminateSessionOnExpiry** | Pointer to **bool** |  | [optional] 
**RefreshInterval** | Pointer to **string** |  | [optional] 
**AuthorizationFlow** | Pointer to **NullableString** |  | [optional] 
**NssUidOffset** | Pointer to **int32** |  | [optional] 
**NssGidOffset** | Pointer to **int32** |  | [optional] 
**ChallengeKey** | Pointer to **NullableString** |  | [optional] 
**ChallengeIdleTimeout** | Pointer to **string** |  | [optional] 
**ChallengeTriggerCheckIn** | Pointer to **bool** |  | [optional] 
**JwtFederationProviders** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewAgentConnector

`func NewAgentConnector(name string, component string, verboseName string, verboseNamePlural string, metaModelName string, ) *AgentConnector`

NewAgentConnector instantiates a new AgentConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentConnectorWithDefaults

`func NewAgentConnectorWithDefaults() *AgentConnector`

NewAgentConnectorWithDefaults instantiates a new AgentConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorUuid

`func (o *AgentConnector) GetConnectorUuid() string`

GetConnectorUuid returns the ConnectorUuid field if non-nil, zero value otherwise.

### GetConnectorUuidOk

`func (o *AgentConnector) GetConnectorUuidOk() (*string, bool)`

GetConnectorUuidOk returns a tuple with the ConnectorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorUuid

`func (o *AgentConnector) SetConnectorUuid(v string)`

SetConnectorUuid sets ConnectorUuid field to given value.

### HasConnectorUuid

`func (o *AgentConnector) HasConnectorUuid() bool`

HasConnectorUuid returns a boolean if a field has been set.

### GetName

`func (o *AgentConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentConnector) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *AgentConnector) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgentConnector) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgentConnector) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AgentConnector) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetComponent

`func (o *AgentConnector) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AgentConnector) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AgentConnector) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *AgentConnector) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *AgentConnector) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *AgentConnector) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *AgentConnector) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *AgentConnector) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *AgentConnector) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *AgentConnector) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *AgentConnector) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *AgentConnector) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetSnapshotExpiry

`func (o *AgentConnector) GetSnapshotExpiry() string`

GetSnapshotExpiry returns the SnapshotExpiry field if non-nil, zero value otherwise.

### GetSnapshotExpiryOk

`func (o *AgentConnector) GetSnapshotExpiryOk() (*string, bool)`

GetSnapshotExpiryOk returns a tuple with the SnapshotExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotExpiry

`func (o *AgentConnector) SetSnapshotExpiry(v string)`

SetSnapshotExpiry sets SnapshotExpiry field to given value.

### HasSnapshotExpiry

`func (o *AgentConnector) HasSnapshotExpiry() bool`

HasSnapshotExpiry returns a boolean if a field has been set.

### GetAuthSessionDuration

`func (o *AgentConnector) GetAuthSessionDuration() string`

GetAuthSessionDuration returns the AuthSessionDuration field if non-nil, zero value otherwise.

### GetAuthSessionDurationOk

`func (o *AgentConnector) GetAuthSessionDurationOk() (*string, bool)`

GetAuthSessionDurationOk returns a tuple with the AuthSessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSessionDuration

`func (o *AgentConnector) SetAuthSessionDuration(v string)`

SetAuthSessionDuration sets AuthSessionDuration field to given value.

### HasAuthSessionDuration

`func (o *AgentConnector) HasAuthSessionDuration() bool`

HasAuthSessionDuration returns a boolean if a field has been set.

### GetAuthTerminateSessionOnExpiry

`func (o *AgentConnector) GetAuthTerminateSessionOnExpiry() bool`

GetAuthTerminateSessionOnExpiry returns the AuthTerminateSessionOnExpiry field if non-nil, zero value otherwise.

### GetAuthTerminateSessionOnExpiryOk

`func (o *AgentConnector) GetAuthTerminateSessionOnExpiryOk() (*bool, bool)`

GetAuthTerminateSessionOnExpiryOk returns a tuple with the AuthTerminateSessionOnExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTerminateSessionOnExpiry

`func (o *AgentConnector) SetAuthTerminateSessionOnExpiry(v bool)`

SetAuthTerminateSessionOnExpiry sets AuthTerminateSessionOnExpiry field to given value.

### HasAuthTerminateSessionOnExpiry

`func (o *AgentConnector) HasAuthTerminateSessionOnExpiry() bool`

HasAuthTerminateSessionOnExpiry returns a boolean if a field has been set.

### GetRefreshInterval

`func (o *AgentConnector) GetRefreshInterval() string`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *AgentConnector) GetRefreshIntervalOk() (*string, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *AgentConnector) SetRefreshInterval(v string)`

SetRefreshInterval sets RefreshInterval field to given value.

### HasRefreshInterval

`func (o *AgentConnector) HasRefreshInterval() bool`

HasRefreshInterval returns a boolean if a field has been set.

### GetAuthorizationFlow

`func (o *AgentConnector) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *AgentConnector) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *AgentConnector) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.

### HasAuthorizationFlow

`func (o *AgentConnector) HasAuthorizationFlow() bool`

HasAuthorizationFlow returns a boolean if a field has been set.

### SetAuthorizationFlowNil

`func (o *AgentConnector) SetAuthorizationFlowNil(b bool)`

 SetAuthorizationFlowNil sets the value for AuthorizationFlow to be an explicit nil

### UnsetAuthorizationFlow
`func (o *AgentConnector) UnsetAuthorizationFlow()`

UnsetAuthorizationFlow ensures that no value is present for AuthorizationFlow, not even an explicit nil
### GetNssUidOffset

`func (o *AgentConnector) GetNssUidOffset() int32`

GetNssUidOffset returns the NssUidOffset field if non-nil, zero value otherwise.

### GetNssUidOffsetOk

`func (o *AgentConnector) GetNssUidOffsetOk() (*int32, bool)`

GetNssUidOffsetOk returns a tuple with the NssUidOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssUidOffset

`func (o *AgentConnector) SetNssUidOffset(v int32)`

SetNssUidOffset sets NssUidOffset field to given value.

### HasNssUidOffset

`func (o *AgentConnector) HasNssUidOffset() bool`

HasNssUidOffset returns a boolean if a field has been set.

### GetNssGidOffset

`func (o *AgentConnector) GetNssGidOffset() int32`

GetNssGidOffset returns the NssGidOffset field if non-nil, zero value otherwise.

### GetNssGidOffsetOk

`func (o *AgentConnector) GetNssGidOffsetOk() (*int32, bool)`

GetNssGidOffsetOk returns a tuple with the NssGidOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssGidOffset

`func (o *AgentConnector) SetNssGidOffset(v int32)`

SetNssGidOffset sets NssGidOffset field to given value.

### HasNssGidOffset

`func (o *AgentConnector) HasNssGidOffset() bool`

HasNssGidOffset returns a boolean if a field has been set.

### GetChallengeKey

`func (o *AgentConnector) GetChallengeKey() string`

GetChallengeKey returns the ChallengeKey field if non-nil, zero value otherwise.

### GetChallengeKeyOk

`func (o *AgentConnector) GetChallengeKeyOk() (*string, bool)`

GetChallengeKeyOk returns a tuple with the ChallengeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeKey

`func (o *AgentConnector) SetChallengeKey(v string)`

SetChallengeKey sets ChallengeKey field to given value.

### HasChallengeKey

`func (o *AgentConnector) HasChallengeKey() bool`

HasChallengeKey returns a boolean if a field has been set.

### SetChallengeKeyNil

`func (o *AgentConnector) SetChallengeKeyNil(b bool)`

 SetChallengeKeyNil sets the value for ChallengeKey to be an explicit nil

### UnsetChallengeKey
`func (o *AgentConnector) UnsetChallengeKey()`

UnsetChallengeKey ensures that no value is present for ChallengeKey, not even an explicit nil
### GetChallengeIdleTimeout

`func (o *AgentConnector) GetChallengeIdleTimeout() string`

GetChallengeIdleTimeout returns the ChallengeIdleTimeout field if non-nil, zero value otherwise.

### GetChallengeIdleTimeoutOk

`func (o *AgentConnector) GetChallengeIdleTimeoutOk() (*string, bool)`

GetChallengeIdleTimeoutOk returns a tuple with the ChallengeIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeIdleTimeout

`func (o *AgentConnector) SetChallengeIdleTimeout(v string)`

SetChallengeIdleTimeout sets ChallengeIdleTimeout field to given value.

### HasChallengeIdleTimeout

`func (o *AgentConnector) HasChallengeIdleTimeout() bool`

HasChallengeIdleTimeout returns a boolean if a field has been set.

### GetChallengeTriggerCheckIn

`func (o *AgentConnector) GetChallengeTriggerCheckIn() bool`

GetChallengeTriggerCheckIn returns the ChallengeTriggerCheckIn field if non-nil, zero value otherwise.

### GetChallengeTriggerCheckInOk

`func (o *AgentConnector) GetChallengeTriggerCheckInOk() (*bool, bool)`

GetChallengeTriggerCheckInOk returns a tuple with the ChallengeTriggerCheckIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeTriggerCheckIn

`func (o *AgentConnector) SetChallengeTriggerCheckIn(v bool)`

SetChallengeTriggerCheckIn sets ChallengeTriggerCheckIn field to given value.

### HasChallengeTriggerCheckIn

`func (o *AgentConnector) HasChallengeTriggerCheckIn() bool`

HasChallengeTriggerCheckIn returns a boolean if a field has been set.

### GetJwtFederationProviders

`func (o *AgentConnector) GetJwtFederationProviders() []int32`

GetJwtFederationProviders returns the JwtFederationProviders field if non-nil, zero value otherwise.

### GetJwtFederationProvidersOk

`func (o *AgentConnector) GetJwtFederationProvidersOk() (*[]int32, bool)`

GetJwtFederationProvidersOk returns a tuple with the JwtFederationProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtFederationProviders

`func (o *AgentConnector) SetJwtFederationProviders(v []int32)`

SetJwtFederationProviders sets JwtFederationProviders field to given value.

### HasJwtFederationProviders

`func (o *AgentConnector) HasJwtFederationProviders() bool`

HasJwtFederationProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


