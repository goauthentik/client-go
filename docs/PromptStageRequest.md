# PromptStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**Fields** | **[]string** |  | 
**ValidationPolicies** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPromptStageRequest

`func NewPromptStageRequest(name string, fields []string, ) *PromptStageRequest`

NewPromptStageRequest instantiates a new PromptStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptStageRequestWithDefaults

`func NewPromptStageRequestWithDefaults() *PromptStageRequest`

NewPromptStageRequestWithDefaults instantiates a new PromptStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PromptStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PromptStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PromptStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowSet

`func (o *PromptStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PromptStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PromptStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *PromptStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetFields

`func (o *PromptStageRequest) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PromptStageRequest) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PromptStageRequest) SetFields(v []string)`

SetFields sets Fields field to given value.


### GetValidationPolicies

`func (o *PromptStageRequest) GetValidationPolicies() []string`

GetValidationPolicies returns the ValidationPolicies field if non-nil, zero value otherwise.

### GetValidationPoliciesOk

`func (o *PromptStageRequest) GetValidationPoliciesOk() (*[]string, bool)`

GetValidationPoliciesOk returns a tuple with the ValidationPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationPolicies

`func (o *PromptStageRequest) SetValidationPolicies(v []string)`

SetValidationPolicies sets ValidationPolicies field to given value.

### HasValidationPolicies

`func (o *PromptStageRequest) HasValidationPolicies() bool`

HasValidationPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


