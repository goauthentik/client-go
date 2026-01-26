# PromptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FieldKey** | **string** | Name of the form field, also used to store the value | 
**Label** | **string** |  | 
**Type** | [**PromptTypeEnum**](PromptTypeEnum.md) |  | 
**Required** | Pointer to **bool** |  | [optional] 
**Placeholder** | Pointer to **string** | Optionally provide a short hint that describes the expected input value. When creating a fixed choice field, enable interpreting as expression and return a list to return multiple choices. | [optional] 
**InitialValue** | Pointer to **string** | Optionally pre-fill the input with an initial value. When creating a fixed choice field, enable interpreting as expression and return a list to return multiple default choices. | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**PromptstageSet** | Pointer to [**[]StageRequest**](StageRequest.md) |  | [optional] 
**SubText** | Pointer to **string** |  | [optional] 
**PlaceholderExpression** | Pointer to **bool** |  | [optional] 
**InitialValueExpression** | Pointer to **bool** |  | [optional] 

## Methods

### NewPromptRequest

`func NewPromptRequest(name string, fieldKey string, label string, type_ PromptTypeEnum, ) *PromptRequest`

NewPromptRequest instantiates a new PromptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptRequestWithDefaults

`func NewPromptRequestWithDefaults() *PromptRequest`

NewPromptRequestWithDefaults instantiates a new PromptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PromptRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PromptRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PromptRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFieldKey

`func (o *PromptRequest) GetFieldKey() string`

GetFieldKey returns the FieldKey field if non-nil, zero value otherwise.

### GetFieldKeyOk

`func (o *PromptRequest) GetFieldKeyOk() (*string, bool)`

GetFieldKeyOk returns a tuple with the FieldKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldKey

`func (o *PromptRequest) SetFieldKey(v string)`

SetFieldKey sets FieldKey field to given value.


### GetLabel

`func (o *PromptRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PromptRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PromptRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *PromptRequest) GetType() PromptTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromptRequest) GetTypeOk() (*PromptTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromptRequest) SetType(v PromptTypeEnum)`

SetType sets Type field to given value.


### GetRequired

`func (o *PromptRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *PromptRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *PromptRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *PromptRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetPlaceholder

`func (o *PromptRequest) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *PromptRequest) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *PromptRequest) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *PromptRequest) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetInitialValue

`func (o *PromptRequest) GetInitialValue() string`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *PromptRequest) GetInitialValueOk() (*string, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *PromptRequest) SetInitialValue(v string)`

SetInitialValue sets InitialValue field to given value.

### HasInitialValue

`func (o *PromptRequest) HasInitialValue() bool`

HasInitialValue returns a boolean if a field has been set.

### GetOrder

`func (o *PromptRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PromptRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PromptRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PromptRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPromptstageSet

`func (o *PromptRequest) GetPromptstageSet() []StageRequest`

GetPromptstageSet returns the PromptstageSet field if non-nil, zero value otherwise.

### GetPromptstageSetOk

`func (o *PromptRequest) GetPromptstageSetOk() (*[]StageRequest, bool)`

GetPromptstageSetOk returns a tuple with the PromptstageSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptstageSet

`func (o *PromptRequest) SetPromptstageSet(v []StageRequest)`

SetPromptstageSet sets PromptstageSet field to given value.

### HasPromptstageSet

`func (o *PromptRequest) HasPromptstageSet() bool`

HasPromptstageSet returns a boolean if a field has been set.

### GetSubText

`func (o *PromptRequest) GetSubText() string`

GetSubText returns the SubText field if non-nil, zero value otherwise.

### GetSubTextOk

`func (o *PromptRequest) GetSubTextOk() (*string, bool)`

GetSubTextOk returns a tuple with the SubText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubText

`func (o *PromptRequest) SetSubText(v string)`

SetSubText sets SubText field to given value.

### HasSubText

`func (o *PromptRequest) HasSubText() bool`

HasSubText returns a boolean if a field has been set.

### GetPlaceholderExpression

`func (o *PromptRequest) GetPlaceholderExpression() bool`

GetPlaceholderExpression returns the PlaceholderExpression field if non-nil, zero value otherwise.

### GetPlaceholderExpressionOk

`func (o *PromptRequest) GetPlaceholderExpressionOk() (*bool, bool)`

GetPlaceholderExpressionOk returns a tuple with the PlaceholderExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholderExpression

`func (o *PromptRequest) SetPlaceholderExpression(v bool)`

SetPlaceholderExpression sets PlaceholderExpression field to given value.

### HasPlaceholderExpression

`func (o *PromptRequest) HasPlaceholderExpression() bool`

HasPlaceholderExpression returns a boolean if a field has been set.

### GetInitialValueExpression

`func (o *PromptRequest) GetInitialValueExpression() bool`

GetInitialValueExpression returns the InitialValueExpression field if non-nil, zero value otherwise.

### GetInitialValueExpressionOk

`func (o *PromptRequest) GetInitialValueExpressionOk() (*bool, bool)`

GetInitialValueExpressionOk returns a tuple with the InitialValueExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValueExpression

`func (o *PromptRequest) SetInitialValueExpression(v bool)`

SetInitialValueExpression sets InitialValueExpression field to given value.

### HasInitialValueExpression

`func (o *PromptRequest) HasInitialValueExpression() bool`

HasInitialValueExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


