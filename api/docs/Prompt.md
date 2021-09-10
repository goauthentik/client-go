# Prompt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**FieldKey** | **string** | Name of the form field, also used to store the value | 
**Label** | **string** |  | 
**Type** | [**PromptTypeEnum**](PromptTypeEnum.md) |  | 
**Required** | Pointer to **bool** |  | [optional] 
**Placeholder** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**PromptstageSet** | Pointer to [**[]Stage**](Stage.md) |  | [optional] 

## Methods

### NewPrompt

`func NewPrompt(pk string, fieldKey string, label string, type_ PromptTypeEnum, ) *Prompt`

NewPrompt instantiates a new Prompt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptWithDefaults

`func NewPromptWithDefaults() *Prompt`

NewPromptWithDefaults instantiates a new Prompt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *Prompt) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *Prompt) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *Prompt) SetPk(v string)`

SetPk sets Pk field to given value.


### GetFieldKey

`func (o *Prompt) GetFieldKey() string`

GetFieldKey returns the FieldKey field if non-nil, zero value otherwise.

### GetFieldKeyOk

`func (o *Prompt) GetFieldKeyOk() (*string, bool)`

GetFieldKeyOk returns a tuple with the FieldKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldKey

`func (o *Prompt) SetFieldKey(v string)`

SetFieldKey sets FieldKey field to given value.


### GetLabel

`func (o *Prompt) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Prompt) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Prompt) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *Prompt) GetType() PromptTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Prompt) GetTypeOk() (*PromptTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Prompt) SetType(v PromptTypeEnum)`

SetType sets Type field to given value.


### GetRequired

`func (o *Prompt) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *Prompt) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *Prompt) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *Prompt) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetPlaceholder

`func (o *Prompt) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *Prompt) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *Prompt) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *Prompt) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetOrder

`func (o *Prompt) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Prompt) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Prompt) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Prompt) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPromptstageSet

`func (o *Prompt) GetPromptstageSet() []Stage`

GetPromptstageSet returns the PromptstageSet field if non-nil, zero value otherwise.

### GetPromptstageSetOk

`func (o *Prompt) GetPromptstageSetOk() (*[]Stage, bool)`

GetPromptstageSetOk returns a tuple with the PromptstageSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptstageSet

`func (o *Prompt) SetPromptstageSet(v []Stage)`

SetPromptstageSet sets PromptstageSet field to given value.

### HasPromptstageSet

`func (o *Prompt) HasPromptstageSet() bool`

HasPromptstageSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


