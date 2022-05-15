# ContextualFlowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Background** | Pointer to **string** |  | [optional] 
**CancelUrl** | **string** |  | 
**Layout** | [**LayoutEnum**](LayoutEnum.md) |  | 

## Methods

### NewContextualFlowInfo

`func NewContextualFlowInfo(cancelUrl string, layout LayoutEnum, ) *ContextualFlowInfo`

NewContextualFlowInfo instantiates a new ContextualFlowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextualFlowInfoWithDefaults

`func NewContextualFlowInfoWithDefaults() *ContextualFlowInfo`

NewContextualFlowInfoWithDefaults instantiates a new ContextualFlowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ContextualFlowInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContextualFlowInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContextualFlowInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ContextualFlowInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBackground

`func (o *ContextualFlowInfo) GetBackground() string`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *ContextualFlowInfo) GetBackgroundOk() (*string, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *ContextualFlowInfo) SetBackground(v string)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *ContextualFlowInfo) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### GetCancelUrl

`func (o *ContextualFlowInfo) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *ContextualFlowInfo) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *ContextualFlowInfo) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.


### GetLayout

`func (o *ContextualFlowInfo) GetLayout() LayoutEnum`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *ContextualFlowInfo) GetLayoutOk() (*LayoutEnum, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *ContextualFlowInfo) SetLayout(v LayoutEnum)`

SetLayout sets Layout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


