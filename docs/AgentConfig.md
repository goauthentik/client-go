# AgentConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** |  | [readonly] 
**RefreshInterval** | **int32** |  | [readonly] 
**AuthorizationFlow** | **NullableString** |  | [readonly] 
**JwksAuth** | **map[string]interface{}** |  | [readonly] 
**JwksChallenge** | **map[string]interface{}** |  | [readonly] 
**NssUidOffset** | **int32** |  | 
**NssGidOffset** | **int32** |  | 
**AuthTerminateSessionOnExpiry** | **bool** |  | 
**SystemConfig** | [**Config**](Config.md) |  | [readonly] 

## Methods

### NewAgentConfig

`func NewAgentConfig(deviceId string, refreshInterval int32, authorizationFlow NullableString, jwksAuth map[string]interface{}, jwksChallenge map[string]interface{}, nssUidOffset int32, nssGidOffset int32, authTerminateSessionOnExpiry bool, systemConfig Config, ) *AgentConfig`

NewAgentConfig instantiates a new AgentConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentConfigWithDefaults

`func NewAgentConfigWithDefaults() *AgentConfig`

NewAgentConfigWithDefaults instantiates a new AgentConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *AgentConfig) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *AgentConfig) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *AgentConfig) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


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


### GetAuthorizationFlow

`func (o *AgentConfig) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *AgentConfig) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *AgentConfig) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.


### SetAuthorizationFlowNil

`func (o *AgentConfig) SetAuthorizationFlowNil(b bool)`

 SetAuthorizationFlowNil sets the value for AuthorizationFlow to be an explicit nil

### UnsetAuthorizationFlow
`func (o *AgentConfig) UnsetAuthorizationFlow()`

UnsetAuthorizationFlow ensures that no value is present for AuthorizationFlow, not even an explicit nil
### GetJwksAuth

`func (o *AgentConfig) GetJwksAuth() map[string]interface{}`

GetJwksAuth returns the JwksAuth field if non-nil, zero value otherwise.

### GetJwksAuthOk

`func (o *AgentConfig) GetJwksAuthOk() (*map[string]interface{}, bool)`

GetJwksAuthOk returns a tuple with the JwksAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksAuth

`func (o *AgentConfig) SetJwksAuth(v map[string]interface{})`

SetJwksAuth sets JwksAuth field to given value.


### GetJwksChallenge

`func (o *AgentConfig) GetJwksChallenge() map[string]interface{}`

GetJwksChallenge returns the JwksChallenge field if non-nil, zero value otherwise.

### GetJwksChallengeOk

`func (o *AgentConfig) GetJwksChallengeOk() (*map[string]interface{}, bool)`

GetJwksChallengeOk returns a tuple with the JwksChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksChallenge

`func (o *AgentConfig) SetJwksChallenge(v map[string]interface{})`

SetJwksChallenge sets JwksChallenge field to given value.


### SetJwksChallengeNil

`func (o *AgentConfig) SetJwksChallengeNil(b bool)`

 SetJwksChallengeNil sets the value for JwksChallenge to be an explicit nil

### UnsetJwksChallenge
`func (o *AgentConfig) UnsetJwksChallenge()`

UnsetJwksChallenge ensures that no value is present for JwksChallenge, not even an explicit nil
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


### GetSystemConfig

`func (o *AgentConfig) GetSystemConfig() Config`

GetSystemConfig returns the SystemConfig field if non-nil, zero value otherwise.

### GetSystemConfigOk

`func (o *AgentConfig) GetSystemConfigOk() (*Config, bool)`

GetSystemConfigOk returns a tuple with the SystemConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemConfig

`func (o *AgentConfig) SetSystemConfig(v Config)`

SetSystemConfig sets SystemConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


