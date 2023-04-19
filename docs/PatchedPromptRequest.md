# PatchedPromptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FieldKey** | Pointer to **string** | Name of the form field, also used to store the value | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**PromptTypeEnum**](PromptTypeEnum.md) |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Placeholder** | Pointer to **string** | Optionally provide a short hint that describes the expected input value. When creating a fixed choice field, enable interpreting as expression and return a list to return multiple choices. | [optional] 
**InitialValue** | Pointer to **string** | Optionally pre-fill the input with an initial value. When creating a fixed choice field, enable interpreting as expression and return a list to return multiple default choices. | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**PromptstageSet** | Pointer to [**[]StageRequest**](StageRequest.md) |  | [optional] 
**SubText** | Pointer to **string** |  | [optional] 
**PlaceholderExpression** | Pointer to **bool** |  | [optional] 
**InitialValueExpression** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedPromptRequest

`func NewPatchedPromptRequest() *PatchedPromptRequest`

NewPatchedPromptRequest instantiates a new PatchedPromptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedPromptRequestWithDefaults

`func NewPatchedPromptRequestWithDefaults() *PatchedPromptRequest`

NewPatchedPromptRequestWithDefaults instantiates a new PatchedPromptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedPromptRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedPromptRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedPromptRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedPromptRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFieldKey

`func (o *PatchedPromptRequest) GetFieldKey() string`

GetFieldKey returns the FieldKey field if non-nil, zero value otherwise.

### GetFieldKeyOk

`func (o *PatchedPromptRequest) GetFieldKeyOk() (*string, bool)`

GetFieldKeyOk returns a tuple with the FieldKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldKey

`func (o *PatchedPromptRequest) SetFieldKey(v string)`

SetFieldKey sets FieldKey field to given value.

### HasFieldKey

`func (o *PatchedPromptRequest) HasFieldKey() bool`

HasFieldKey returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedPromptRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedPromptRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedPromptRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedPromptRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PatchedPromptRequest) GetType() PromptTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedPromptRequest) GetTypeOk() (*PromptTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedPromptRequest) SetType(v PromptTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedPromptRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRequired

`func (o *PatchedPromptRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *PatchedPromptRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *PatchedPromptRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *PatchedPromptRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetPlaceholder

`func (o *PatchedPromptRequest) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *PatchedPromptRequest) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *PatchedPromptRequest) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *PatchedPromptRequest) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetInitialValue

`func (o *PatchedPromptRequest) GetInitialValue() string`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *PatchedPromptRequest) GetInitialValueOk() (*string, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *PatchedPromptRequest) SetInitialValue(v string)`

SetInitialValue sets InitialValue field to given value.

### HasInitialValue

`func (o *PatchedPromptRequest) HasInitialValue() bool`

HasInitialValue returns a boolean if a field has been set.

### GetOrder

`func (o *PatchedPromptRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PatchedPromptRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PatchedPromptRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PatchedPromptRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPromptstageSet

`func (o *PatchedPromptRequest) GetPromptstageSet() []StageRequest`

GetPromptstageSet returns the PromptstageSet field if non-nil, zero value otherwise.

### GetPromptstageSetOk

`func (o *PatchedPromptRequest) GetPromptstageSetOk() (*[]StageRequest, bool)`

GetPromptstageSetOk returns a tuple with the PromptstageSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptstageSet

`func (o *PatchedPromptRequest) SetPromptstageSet(v []StageRequest)`

SetPromptstageSet sets PromptstageSet field to given value.

### HasPromptstageSet

`func (o *PatchedPromptRequest) HasPromptstageSet() bool`

HasPromptstageSet returns a boolean if a field has been set.

### GetSubText

`func (o *PatchedPromptRequest) GetSubText() string`

GetSubText returns the SubText field if non-nil, zero value otherwise.

### GetSubTextOk

`func (o *PatchedPromptRequest) GetSubTextOk() (*string, bool)`

GetSubTextOk returns a tuple with the SubText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubText

`func (o *PatchedPromptRequest) SetSubText(v string)`

SetSubText sets SubText field to given value.

### HasSubText

`func (o *PatchedPromptRequest) HasSubText() bool`

HasSubText returns a boolean if a field has been set.

### GetPlaceholderExpression

`func (o *PatchedPromptRequest) GetPlaceholderExpression() bool`

GetPlaceholderExpression returns the PlaceholderExpression field if non-nil, zero value otherwise.

### GetPlaceholderExpressionOk

`func (o *PatchedPromptRequest) GetPlaceholderExpressionOk() (*bool, bool)`

GetPlaceholderExpressionOk returns a tuple with the PlaceholderExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholderExpression

`func (o *PatchedPromptRequest) SetPlaceholderExpression(v bool)`

SetPlaceholderExpression sets PlaceholderExpression field to given value.

### HasPlaceholderExpression

`func (o *PatchedPromptRequest) HasPlaceholderExpression() bool`

HasPlaceholderExpression returns a boolean if a field has been set.

### GetInitialValueExpression

`func (o *PatchedPromptRequest) GetInitialValueExpression() bool`

GetInitialValueExpression returns the InitialValueExpression field if non-nil, zero value otherwise.

### GetInitialValueExpressionOk

`func (o *PatchedPromptRequest) GetInitialValueExpressionOk() (*bool, bool)`

GetInitialValueExpressionOk returns a tuple with the InitialValueExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValueExpression

`func (o *PatchedPromptRequest) SetInitialValueExpression(v bool)`

SetInitialValueExpression sets InitialValueExpression field to given value.

### HasInitialValueExpression

`func (o *PatchedPromptRequest) HasInitialValueExpression() bool`

HasInitialValueExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


