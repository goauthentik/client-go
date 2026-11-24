# RequestRuleBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**Rule** | **string** |  | 
**Target** | **string** |  | 
**ExpiryPending** | Pointer to **string** | How long a request against this binding stays pending before it automatically lapses if not approved or denied. | [optional] 
**ExpiryGrantedMax** | Pointer to **string** | The maximum duration a grant approved against this binding can last. | [optional] 

## Methods

### NewRequestRuleBindingRequest

`func NewRequestRuleBindingRequest(rule string, target string, ) *RequestRuleBindingRequest`

NewRequestRuleBindingRequest instantiates a new RequestRuleBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestRuleBindingRequestWithDefaults

`func NewRequestRuleBindingRequestWithDefaults() *RequestRuleBindingRequest`

NewRequestRuleBindingRequestWithDefaults instantiates a new RequestRuleBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RequestRuleBindingRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RequestRuleBindingRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RequestRuleBindingRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RequestRuleBindingRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *RequestRuleBindingRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *RequestRuleBindingRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *RequestRuleBindingRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *RequestRuleBindingRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetRule

`func (o *RequestRuleBindingRequest) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *RequestRuleBindingRequest) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *RequestRuleBindingRequest) SetRule(v string)`

SetRule sets Rule field to given value.


### GetTarget

`func (o *RequestRuleBindingRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RequestRuleBindingRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RequestRuleBindingRequest) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetExpiryPending

`func (o *RequestRuleBindingRequest) GetExpiryPending() string`

GetExpiryPending returns the ExpiryPending field if non-nil, zero value otherwise.

### GetExpiryPendingOk

`func (o *RequestRuleBindingRequest) GetExpiryPendingOk() (*string, bool)`

GetExpiryPendingOk returns a tuple with the ExpiryPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryPending

`func (o *RequestRuleBindingRequest) SetExpiryPending(v string)`

SetExpiryPending sets ExpiryPending field to given value.

### HasExpiryPending

`func (o *RequestRuleBindingRequest) HasExpiryPending() bool`

HasExpiryPending returns a boolean if a field has been set.

### GetExpiryGrantedMax

`func (o *RequestRuleBindingRequest) GetExpiryGrantedMax() string`

GetExpiryGrantedMax returns the ExpiryGrantedMax field if non-nil, zero value otherwise.

### GetExpiryGrantedMaxOk

`func (o *RequestRuleBindingRequest) GetExpiryGrantedMaxOk() (*string, bool)`

GetExpiryGrantedMaxOk returns a tuple with the ExpiryGrantedMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryGrantedMax

`func (o *RequestRuleBindingRequest) SetExpiryGrantedMax(v string)`

SetExpiryGrantedMax sets ExpiryGrantedMax field to given value.

### HasExpiryGrantedMax

`func (o *RequestRuleBindingRequest) HasExpiryGrantedMax() bool`

HasExpiryGrantedMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


