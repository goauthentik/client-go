# FlowSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**PolicybindingmodelPtrId** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** | Visible in the URL. | 
**Title** | **string** | Shown as the Title in Flow pages. | 
**Designation** | [**FlowDesignationEnum**](FlowDesignationEnum.md) |  | 
**Background** | **string** | Get the URL to the background image. If the name is /static or starts with http it is returned as-is | [readonly] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**CompatibilityMode** | Pointer to **bool** | Enable compatibility mode, increases compatibility with password managers on mobile devices. | [optional] 
**ExportUrl** | **string** | Get export URL for flow | [readonly] 
**Layout** | Pointer to [**LayoutEnum**](LayoutEnum.md) |  | [optional] 
**DeniedAction** | Pointer to [**DeniedActionEnum**](DeniedActionEnum.md) |  | [optional] 

## Methods

### NewFlowSet

`func NewFlowSet(pk string, policybindingmodelPtrId string, name string, slug string, title string, designation FlowDesignationEnum, background string, exportUrl string, ) *FlowSet`

NewFlowSet instantiates a new FlowSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowSetWithDefaults

`func NewFlowSetWithDefaults() *FlowSet`

NewFlowSetWithDefaults instantiates a new FlowSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *FlowSet) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *FlowSet) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *FlowSet) SetPk(v string)`

SetPk sets Pk field to given value.


### GetPolicybindingmodelPtrId

`func (o *FlowSet) GetPolicybindingmodelPtrId() string`

GetPolicybindingmodelPtrId returns the PolicybindingmodelPtrId field if non-nil, zero value otherwise.

### GetPolicybindingmodelPtrIdOk

`func (o *FlowSet) GetPolicybindingmodelPtrIdOk() (*string, bool)`

GetPolicybindingmodelPtrIdOk returns a tuple with the PolicybindingmodelPtrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicybindingmodelPtrId

`func (o *FlowSet) SetPolicybindingmodelPtrId(v string)`

SetPolicybindingmodelPtrId sets PolicybindingmodelPtrId field to given value.


### GetName

`func (o *FlowSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowSet) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *FlowSet) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *FlowSet) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *FlowSet) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTitle

`func (o *FlowSet) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FlowSet) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FlowSet) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDesignation

`func (o *FlowSet) GetDesignation() FlowDesignationEnum`

GetDesignation returns the Designation field if non-nil, zero value otherwise.

### GetDesignationOk

`func (o *FlowSet) GetDesignationOk() (*FlowDesignationEnum, bool)`

GetDesignationOk returns a tuple with the Designation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignation

`func (o *FlowSet) SetDesignation(v FlowDesignationEnum)`

SetDesignation sets Designation field to given value.


### GetBackground

`func (o *FlowSet) GetBackground() string`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *FlowSet) GetBackgroundOk() (*string, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *FlowSet) SetBackground(v string)`

SetBackground sets Background field to given value.


### GetPolicyEngineMode

`func (o *FlowSet) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *FlowSet) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *FlowSet) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *FlowSet) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetCompatibilityMode

`func (o *FlowSet) GetCompatibilityMode() bool`

GetCompatibilityMode returns the CompatibilityMode field if non-nil, zero value otherwise.

### GetCompatibilityModeOk

`func (o *FlowSet) GetCompatibilityModeOk() (*bool, bool)`

GetCompatibilityModeOk returns a tuple with the CompatibilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityMode

`func (o *FlowSet) SetCompatibilityMode(v bool)`

SetCompatibilityMode sets CompatibilityMode field to given value.

### HasCompatibilityMode

`func (o *FlowSet) HasCompatibilityMode() bool`

HasCompatibilityMode returns a boolean if a field has been set.

### GetExportUrl

`func (o *FlowSet) GetExportUrl() string`

GetExportUrl returns the ExportUrl field if non-nil, zero value otherwise.

### GetExportUrlOk

`func (o *FlowSet) GetExportUrlOk() (*string, bool)`

GetExportUrlOk returns a tuple with the ExportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportUrl

`func (o *FlowSet) SetExportUrl(v string)`

SetExportUrl sets ExportUrl field to given value.


### GetLayout

`func (o *FlowSet) GetLayout() LayoutEnum`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *FlowSet) GetLayoutOk() (*LayoutEnum, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *FlowSet) SetLayout(v LayoutEnum)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *FlowSet) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetDeniedAction

`func (o *FlowSet) GetDeniedAction() DeniedActionEnum`

GetDeniedAction returns the DeniedAction field if non-nil, zero value otherwise.

### GetDeniedActionOk

`func (o *FlowSet) GetDeniedActionOk() (*DeniedActionEnum, bool)`

GetDeniedActionOk returns a tuple with the DeniedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedAction

`func (o *FlowSet) SetDeniedAction(v DeniedActionEnum)`

SetDeniedAction sets DeniedAction field to given value.

### HasDeniedAction

`func (o *FlowSet) HasDeniedAction() bool`

HasDeniedAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


