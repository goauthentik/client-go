# Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorReportingEnabled** | **bool** |  | [readonly] 
**ErrorReportingEnvironment** | **string** |  | [readonly] 
**ErrorReportingSendPii** | **bool** |  | [readonly] 
**Capabilities** | [**[]CapabilitiesEnum**](CapabilitiesEnum.md) |  | 
**CacheTimeout** | **int32** |  | 
**CacheTimeoutFlows** | **int32** |  | 
**CacheTimeoutPolicies** | **int32** |  | 
**CacheTimeoutReputation** | **int32** |  | 

## Methods

### NewConfig

`func NewConfig(errorReportingEnabled bool, errorReportingEnvironment string, errorReportingSendPii bool, capabilities []CapabilitiesEnum, cacheTimeout int32, cacheTimeoutFlows int32, cacheTimeoutPolicies int32, cacheTimeoutReputation int32, ) *Config`

NewConfig instantiates a new Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigWithDefaults

`func NewConfigWithDefaults() *Config`

NewConfigWithDefaults instantiates a new Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorReportingEnabled

`func (o *Config) GetErrorReportingEnabled() bool`

GetErrorReportingEnabled returns the ErrorReportingEnabled field if non-nil, zero value otherwise.

### GetErrorReportingEnabledOk

`func (o *Config) GetErrorReportingEnabledOk() (*bool, bool)`

GetErrorReportingEnabledOk returns a tuple with the ErrorReportingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReportingEnabled

`func (o *Config) SetErrorReportingEnabled(v bool)`

SetErrorReportingEnabled sets ErrorReportingEnabled field to given value.


### GetErrorReportingEnvironment

`func (o *Config) GetErrorReportingEnvironment() string`

GetErrorReportingEnvironment returns the ErrorReportingEnvironment field if non-nil, zero value otherwise.

### GetErrorReportingEnvironmentOk

`func (o *Config) GetErrorReportingEnvironmentOk() (*string, bool)`

GetErrorReportingEnvironmentOk returns a tuple with the ErrorReportingEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReportingEnvironment

`func (o *Config) SetErrorReportingEnvironment(v string)`

SetErrorReportingEnvironment sets ErrorReportingEnvironment field to given value.


### GetErrorReportingSendPii

`func (o *Config) GetErrorReportingSendPii() bool`

GetErrorReportingSendPii returns the ErrorReportingSendPii field if non-nil, zero value otherwise.

### GetErrorReportingSendPiiOk

`func (o *Config) GetErrorReportingSendPiiOk() (*bool, bool)`

GetErrorReportingSendPiiOk returns a tuple with the ErrorReportingSendPii field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReportingSendPii

`func (o *Config) SetErrorReportingSendPii(v bool)`

SetErrorReportingSendPii sets ErrorReportingSendPii field to given value.


### GetCapabilities

`func (o *Config) GetCapabilities() []CapabilitiesEnum`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Config) GetCapabilitiesOk() (*[]CapabilitiesEnum, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Config) SetCapabilities(v []CapabilitiesEnum)`

SetCapabilities sets Capabilities field to given value.


### GetCacheTimeout

`func (o *Config) GetCacheTimeout() int32`

GetCacheTimeout returns the CacheTimeout field if non-nil, zero value otherwise.

### GetCacheTimeoutOk

`func (o *Config) GetCacheTimeoutOk() (*int32, bool)`

GetCacheTimeoutOk returns a tuple with the CacheTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeout

`func (o *Config) SetCacheTimeout(v int32)`

SetCacheTimeout sets CacheTimeout field to given value.


### GetCacheTimeoutFlows

`func (o *Config) GetCacheTimeoutFlows() int32`

GetCacheTimeoutFlows returns the CacheTimeoutFlows field if non-nil, zero value otherwise.

### GetCacheTimeoutFlowsOk

`func (o *Config) GetCacheTimeoutFlowsOk() (*int32, bool)`

GetCacheTimeoutFlowsOk returns a tuple with the CacheTimeoutFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeoutFlows

`func (o *Config) SetCacheTimeoutFlows(v int32)`

SetCacheTimeoutFlows sets CacheTimeoutFlows field to given value.


### GetCacheTimeoutPolicies

`func (o *Config) GetCacheTimeoutPolicies() int32`

GetCacheTimeoutPolicies returns the CacheTimeoutPolicies field if non-nil, zero value otherwise.

### GetCacheTimeoutPoliciesOk

`func (o *Config) GetCacheTimeoutPoliciesOk() (*int32, bool)`

GetCacheTimeoutPoliciesOk returns a tuple with the CacheTimeoutPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeoutPolicies

`func (o *Config) SetCacheTimeoutPolicies(v int32)`

SetCacheTimeoutPolicies sets CacheTimeoutPolicies field to given value.


### GetCacheTimeoutReputation

`func (o *Config) GetCacheTimeoutReputation() int32`

GetCacheTimeoutReputation returns the CacheTimeoutReputation field if non-nil, zero value otherwise.

### GetCacheTimeoutReputationOk

`func (o *Config) GetCacheTimeoutReputationOk() (*int32, bool)`

GetCacheTimeoutReputationOk returns a tuple with the CacheTimeoutReputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeoutReputation

`func (o *Config) SetCacheTimeoutReputation(v int32)`

SetCacheTimeoutReputation sets CacheTimeoutReputation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


