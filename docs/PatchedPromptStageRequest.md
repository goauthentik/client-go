# PatchedPromptStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**Fields** | Pointer to **[]string** |  | [optional] 
**ValidationPolicies** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPatchedPromptStageRequest

`func NewPatchedPromptStageRequest() *PatchedPromptStageRequest`

NewPatchedPromptStageRequest instantiates a new PatchedPromptStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedPromptStageRequestWithDefaults

`func NewPatchedPromptStageRequestWithDefaults() *PatchedPromptStageRequest`

NewPatchedPromptStageRequestWithDefaults instantiates a new PatchedPromptStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedPromptStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedPromptStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedPromptStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedPromptStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlowSet

`func (o *PatchedPromptStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PatchedPromptStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PatchedPromptStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *PatchedPromptStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetFields

`func (o *PatchedPromptStageRequest) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PatchedPromptStageRequest) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PatchedPromptStageRequest) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *PatchedPromptStageRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetValidationPolicies

`func (o *PatchedPromptStageRequest) GetValidationPolicies() []string`

GetValidationPolicies returns the ValidationPolicies field if non-nil, zero value otherwise.

### GetValidationPoliciesOk

`func (o *PatchedPromptStageRequest) GetValidationPoliciesOk() (*[]string, bool)`

GetValidationPoliciesOk returns a tuple with the ValidationPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationPolicies

`func (o *PatchedPromptStageRequest) SetValidationPolicies(v []string)`

SetValidationPolicies sets ValidationPolicies field to given value.

### HasValidationPolicies

`func (o *PatchedPromptStageRequest) HasValidationPolicies() bool`

HasValidationPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


