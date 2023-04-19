# StagePrompt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldKey** | **string** |  | 
**Label** | **string** |  | 
**Type** | [**PromptTypeEnum**](PromptTypeEnum.md) |  | 
**Required** | **bool** |  | 
**Placeholder** | **string** |  | 
**InitialValue** | **string** |  | 
**Order** | **int32** |  | 
**SubText** | **string** |  | 
**Choices** | **[]string** |  | 

## Methods

### NewStagePrompt

`func NewStagePrompt(fieldKey string, label string, type_ PromptTypeEnum, required bool, placeholder string, initialValue string, order int32, subText string, choices []string, ) *StagePrompt`

NewStagePrompt instantiates a new StagePrompt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStagePromptWithDefaults

`func NewStagePromptWithDefaults() *StagePrompt`

NewStagePromptWithDefaults instantiates a new StagePrompt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldKey

`func (o *StagePrompt) GetFieldKey() string`

GetFieldKey returns the FieldKey field if non-nil, zero value otherwise.

### GetFieldKeyOk

`func (o *StagePrompt) GetFieldKeyOk() (*string, bool)`

GetFieldKeyOk returns a tuple with the FieldKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldKey

`func (o *StagePrompt) SetFieldKey(v string)`

SetFieldKey sets FieldKey field to given value.


### GetLabel

`func (o *StagePrompt) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *StagePrompt) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *StagePrompt) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *StagePrompt) GetType() PromptTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StagePrompt) GetTypeOk() (*PromptTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StagePrompt) SetType(v PromptTypeEnum)`

SetType sets Type field to given value.


### GetRequired

`func (o *StagePrompt) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *StagePrompt) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *StagePrompt) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetPlaceholder

`func (o *StagePrompt) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *StagePrompt) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *StagePrompt) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.


### GetInitialValue

`func (o *StagePrompt) GetInitialValue() string`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *StagePrompt) GetInitialValueOk() (*string, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *StagePrompt) SetInitialValue(v string)`

SetInitialValue sets InitialValue field to given value.


### GetOrder

`func (o *StagePrompt) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *StagePrompt) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *StagePrompt) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetSubText

`func (o *StagePrompt) GetSubText() string`

GetSubText returns the SubText field if non-nil, zero value otherwise.

### GetSubTextOk

`func (o *StagePrompt) GetSubTextOk() (*string, bool)`

GetSubTextOk returns a tuple with the SubText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubText

`func (o *StagePrompt) SetSubText(v string)`

SetSubText sets SubText field to given value.


### GetChoices

`func (o *StagePrompt) GetChoices() []string`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *StagePrompt) GetChoicesOk() (*[]string, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *StagePrompt) SetChoices(v []string)`

SetChoices sets Choices field to given value.


### SetChoicesNil

`func (o *StagePrompt) SetChoicesNil(b bool)`

 SetChoicesNil sets the value for Choices to be an explicit nil

### UnsetChoices
`func (o *StagePrompt) UnsetChoices()`

UnsetChoices ensures that no value is present for Choices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


