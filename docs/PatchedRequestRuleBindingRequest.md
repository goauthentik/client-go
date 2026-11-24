# PatchedRequestRuleBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**Rule** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**ExpiryPending** | Pointer to **string** | How long a request against this binding stays pending before it automatically lapses if not approved or denied. | [optional] 
**ExpiryGrantedMax** | Pointer to **string** | The maximum duration a grant approved against this binding can last. | [optional] 

## Methods

### NewPatchedRequestRuleBindingRequest

`func NewPatchedRequestRuleBindingRequest() *PatchedRequestRuleBindingRequest`

NewPatchedRequestRuleBindingRequest instantiates a new PatchedRequestRuleBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedRequestRuleBindingRequestWithDefaults

`func NewPatchedRequestRuleBindingRequestWithDefaults() *PatchedRequestRuleBindingRequest`

NewPatchedRequestRuleBindingRequestWithDefaults instantiates a new PatchedRequestRuleBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PatchedRequestRuleBindingRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PatchedRequestRuleBindingRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PatchedRequestRuleBindingRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PatchedRequestRuleBindingRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *PatchedRequestRuleBindingRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *PatchedRequestRuleBindingRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *PatchedRequestRuleBindingRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *PatchedRequestRuleBindingRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetRule

`func (o *PatchedRequestRuleBindingRequest) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *PatchedRequestRuleBindingRequest) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *PatchedRequestRuleBindingRequest) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *PatchedRequestRuleBindingRequest) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetTarget

`func (o *PatchedRequestRuleBindingRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PatchedRequestRuleBindingRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PatchedRequestRuleBindingRequest) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *PatchedRequestRuleBindingRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetExpiryPending

`func (o *PatchedRequestRuleBindingRequest) GetExpiryPending() string`

GetExpiryPending returns the ExpiryPending field if non-nil, zero value otherwise.

### GetExpiryPendingOk

`func (o *PatchedRequestRuleBindingRequest) GetExpiryPendingOk() (*string, bool)`

GetExpiryPendingOk returns a tuple with the ExpiryPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryPending

`func (o *PatchedRequestRuleBindingRequest) SetExpiryPending(v string)`

SetExpiryPending sets ExpiryPending field to given value.

### HasExpiryPending

`func (o *PatchedRequestRuleBindingRequest) HasExpiryPending() bool`

HasExpiryPending returns a boolean if a field has been set.

### GetExpiryGrantedMax

`func (o *PatchedRequestRuleBindingRequest) GetExpiryGrantedMax() string`

GetExpiryGrantedMax returns the ExpiryGrantedMax field if non-nil, zero value otherwise.

### GetExpiryGrantedMaxOk

`func (o *PatchedRequestRuleBindingRequest) GetExpiryGrantedMaxOk() (*string, bool)`

GetExpiryGrantedMaxOk returns a tuple with the ExpiryGrantedMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryGrantedMax

`func (o *PatchedRequestRuleBindingRequest) SetExpiryGrantedMax(v string)`

SetExpiryGrantedMax sets ExpiryGrantedMax field to given value.

### HasExpiryGrantedMax

`func (o *PatchedRequestRuleBindingRequest) HasExpiryGrantedMax() bool`

HasExpiryGrantedMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


