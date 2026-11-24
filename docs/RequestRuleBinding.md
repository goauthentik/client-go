# RequestRuleBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**PbmUuid** | **string** |  | [readonly] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**Rule** | **string** |  | 
**RuleObj** | [**RequestRule**](RequestRule.md) |  | [readonly] 
**Target** | **string** |  | 
**Related** | **[]string** |  | [readonly] 
**RelatedObj** | [**[]RelatedTarget**](RelatedTarget.md) |  | [readonly] 
**ExpiryPending** | Pointer to **string** | How long a request against this binding stays pending before it automatically lapses if not approved or denied. | [optional] 
**ExpiryGrantedMax** | Pointer to **string** | The maximum duration a grant approved against this binding can last. | [optional] 

## Methods

### NewRequestRuleBinding

`func NewRequestRuleBinding(pbmUuid string, rule string, ruleObj RequestRule, target string, related []string, relatedObj []RelatedTarget, ) *RequestRuleBinding`

NewRequestRuleBinding instantiates a new RequestRuleBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestRuleBindingWithDefaults

`func NewRequestRuleBindingWithDefaults() *RequestRuleBinding`

NewRequestRuleBindingWithDefaults instantiates a new RequestRuleBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RequestRuleBinding) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RequestRuleBinding) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RequestRuleBinding) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RequestRuleBinding) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetPbmUuid

`func (o *RequestRuleBinding) GetPbmUuid() string`

GetPbmUuid returns the PbmUuid field if non-nil, zero value otherwise.

### GetPbmUuidOk

`func (o *RequestRuleBinding) GetPbmUuidOk() (*string, bool)`

GetPbmUuidOk returns a tuple with the PbmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbmUuid

`func (o *RequestRuleBinding) SetPbmUuid(v string)`

SetPbmUuid sets PbmUuid field to given value.


### GetPolicyEngineMode

`func (o *RequestRuleBinding) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *RequestRuleBinding) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *RequestRuleBinding) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *RequestRuleBinding) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetRule

`func (o *RequestRuleBinding) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *RequestRuleBinding) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *RequestRuleBinding) SetRule(v string)`

SetRule sets Rule field to given value.


### GetRuleObj

`func (o *RequestRuleBinding) GetRuleObj() RequestRule`

GetRuleObj returns the RuleObj field if non-nil, zero value otherwise.

### GetRuleObjOk

`func (o *RequestRuleBinding) GetRuleObjOk() (*RequestRule, bool)`

GetRuleObjOk returns a tuple with the RuleObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleObj

`func (o *RequestRuleBinding) SetRuleObj(v RequestRule)`

SetRuleObj sets RuleObj field to given value.


### GetTarget

`func (o *RequestRuleBinding) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RequestRuleBinding) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RequestRuleBinding) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetRelated

`func (o *RequestRuleBinding) GetRelated() []string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *RequestRuleBinding) GetRelatedOk() (*[]string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *RequestRuleBinding) SetRelated(v []string)`

SetRelated sets Related field to given value.


### GetRelatedObj

`func (o *RequestRuleBinding) GetRelatedObj() []RelatedTarget`

GetRelatedObj returns the RelatedObj field if non-nil, zero value otherwise.

### GetRelatedObjOk

`func (o *RequestRuleBinding) GetRelatedObjOk() (*[]RelatedTarget, bool)`

GetRelatedObjOk returns a tuple with the RelatedObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObj

`func (o *RequestRuleBinding) SetRelatedObj(v []RelatedTarget)`

SetRelatedObj sets RelatedObj field to given value.


### GetExpiryPending

`func (o *RequestRuleBinding) GetExpiryPending() string`

GetExpiryPending returns the ExpiryPending field if non-nil, zero value otherwise.

### GetExpiryPendingOk

`func (o *RequestRuleBinding) GetExpiryPendingOk() (*string, bool)`

GetExpiryPendingOk returns a tuple with the ExpiryPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryPending

`func (o *RequestRuleBinding) SetExpiryPending(v string)`

SetExpiryPending sets ExpiryPending field to given value.

### HasExpiryPending

`func (o *RequestRuleBinding) HasExpiryPending() bool`

HasExpiryPending returns a boolean if a field has been set.

### GetExpiryGrantedMax

`func (o *RequestRuleBinding) GetExpiryGrantedMax() string`

GetExpiryGrantedMax returns the ExpiryGrantedMax field if non-nil, zero value otherwise.

### GetExpiryGrantedMaxOk

`func (o *RequestRuleBinding) GetExpiryGrantedMaxOk() (*string, bool)`

GetExpiryGrantedMaxOk returns a tuple with the ExpiryGrantedMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryGrantedMax

`func (o *RequestRuleBinding) SetExpiryGrantedMax(v string)`

SetExpiryGrantedMax sets ExpiryGrantedMax field to given value.

### HasExpiryGrantedMax

`func (o *RequestRuleBinding) HasExpiryGrantedMax() bool`

HasExpiryGrantedMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


