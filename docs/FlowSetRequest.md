# FlowSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Slug** | **string** | Visible in the URL. | 
**Title** | **string** | Shown as the Title in Flow pages. | 
**Designation** | [**FlowDesignationEnum**](FlowDesignationEnum.md) |  | 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**CompatibilityMode** | Pointer to **bool** | Enable compatibility mode, increases compatibility with password managers on mobile devices. | [optional] 
**Layout** | Pointer to [**LayoutEnum**](LayoutEnum.md) |  | [optional] 
**DeniedAction** | Pointer to [**DeniedActionEnum**](DeniedActionEnum.md) |  | [optional] 

## Methods

### NewFlowSetRequest

`func NewFlowSetRequest(name string, slug string, title string, designation FlowDesignationEnum, ) *FlowSetRequest`

NewFlowSetRequest instantiates a new FlowSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowSetRequestWithDefaults

`func NewFlowSetRequestWithDefaults() *FlowSetRequest`

NewFlowSetRequestWithDefaults instantiates a new FlowSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FlowSetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowSetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowSetRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *FlowSetRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *FlowSetRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *FlowSetRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTitle

`func (o *FlowSetRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FlowSetRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FlowSetRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDesignation

`func (o *FlowSetRequest) GetDesignation() FlowDesignationEnum`

GetDesignation returns the Designation field if non-nil, zero value otherwise.

### GetDesignationOk

`func (o *FlowSetRequest) GetDesignationOk() (*FlowDesignationEnum, bool)`

GetDesignationOk returns a tuple with the Designation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignation

`func (o *FlowSetRequest) SetDesignation(v FlowDesignationEnum)`

SetDesignation sets Designation field to given value.


### GetPolicyEngineMode

`func (o *FlowSetRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *FlowSetRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *FlowSetRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *FlowSetRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetCompatibilityMode

`func (o *FlowSetRequest) GetCompatibilityMode() bool`

GetCompatibilityMode returns the CompatibilityMode field if non-nil, zero value otherwise.

### GetCompatibilityModeOk

`func (o *FlowSetRequest) GetCompatibilityModeOk() (*bool, bool)`

GetCompatibilityModeOk returns a tuple with the CompatibilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityMode

`func (o *FlowSetRequest) SetCompatibilityMode(v bool)`

SetCompatibilityMode sets CompatibilityMode field to given value.

### HasCompatibilityMode

`func (o *FlowSetRequest) HasCompatibilityMode() bool`

HasCompatibilityMode returns a boolean if a field has been set.

### GetLayout

`func (o *FlowSetRequest) GetLayout() LayoutEnum`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *FlowSetRequest) GetLayoutOk() (*LayoutEnum, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *FlowSetRequest) SetLayout(v LayoutEnum)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *FlowSetRequest) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetDeniedAction

`func (o *FlowSetRequest) GetDeniedAction() DeniedActionEnum`

GetDeniedAction returns the DeniedAction field if non-nil, zero value otherwise.

### GetDeniedActionOk

`func (o *FlowSetRequest) GetDeniedActionOk() (*DeniedActionEnum, bool)`

GetDeniedActionOk returns a tuple with the DeniedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedAction

`func (o *FlowSetRequest) SetDeniedAction(v DeniedActionEnum)`

SetDeniedAction sets DeniedAction field to given value.

### HasDeniedAction

`func (o *FlowSetRequest) HasDeniedAction() bool`

HasDeniedAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


