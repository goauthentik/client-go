# PatchedFlowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** | Visible in the URL. | [optional] 
**Title** | Pointer to **string** | Shown as the Title in Flow pages. | [optional] 
**Designation** | Pointer to [**FlowDesignationEnum**](FlowDesignationEnum.md) |  | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**CompatibilityMode** | Pointer to **bool** | Enable compatibility mode, increases compatibility with password managers on mobile devices. | [optional] 
**Layout** | Pointer to [**LayoutEnum**](LayoutEnum.md) |  | [optional] 
**DeniedAction** | Pointer to [**DeniedActionEnum**](DeniedActionEnum.md) |  | [optional] 
**Authentication** | Pointer to [**AuthenticationEnum**](AuthenticationEnum.md) |  | [optional] 

## Methods

### NewPatchedFlowRequest

`func NewPatchedFlowRequest() *PatchedFlowRequest`

NewPatchedFlowRequest instantiates a new PatchedFlowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedFlowRequestWithDefaults

`func NewPatchedFlowRequestWithDefaults() *PatchedFlowRequest`

NewPatchedFlowRequestWithDefaults instantiates a new PatchedFlowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedFlowRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedFlowRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedFlowRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedFlowRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedFlowRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedFlowRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedFlowRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedFlowRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTitle

`func (o *PatchedFlowRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PatchedFlowRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PatchedFlowRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PatchedFlowRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDesignation

`func (o *PatchedFlowRequest) GetDesignation() FlowDesignationEnum`

GetDesignation returns the Designation field if non-nil, zero value otherwise.

### GetDesignationOk

`func (o *PatchedFlowRequest) GetDesignationOk() (*FlowDesignationEnum, bool)`

GetDesignationOk returns a tuple with the Designation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignation

`func (o *PatchedFlowRequest) SetDesignation(v FlowDesignationEnum)`

SetDesignation sets Designation field to given value.

### HasDesignation

`func (o *PatchedFlowRequest) HasDesignation() bool`

HasDesignation returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *PatchedFlowRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *PatchedFlowRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *PatchedFlowRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *PatchedFlowRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetCompatibilityMode

`func (o *PatchedFlowRequest) GetCompatibilityMode() bool`

GetCompatibilityMode returns the CompatibilityMode field if non-nil, zero value otherwise.

### GetCompatibilityModeOk

`func (o *PatchedFlowRequest) GetCompatibilityModeOk() (*bool, bool)`

GetCompatibilityModeOk returns a tuple with the CompatibilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityMode

`func (o *PatchedFlowRequest) SetCompatibilityMode(v bool)`

SetCompatibilityMode sets CompatibilityMode field to given value.

### HasCompatibilityMode

`func (o *PatchedFlowRequest) HasCompatibilityMode() bool`

HasCompatibilityMode returns a boolean if a field has been set.

### GetLayout

`func (o *PatchedFlowRequest) GetLayout() LayoutEnum`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *PatchedFlowRequest) GetLayoutOk() (*LayoutEnum, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *PatchedFlowRequest) SetLayout(v LayoutEnum)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *PatchedFlowRequest) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetDeniedAction

`func (o *PatchedFlowRequest) GetDeniedAction() DeniedActionEnum`

GetDeniedAction returns the DeniedAction field if non-nil, zero value otherwise.

### GetDeniedActionOk

`func (o *PatchedFlowRequest) GetDeniedActionOk() (*DeniedActionEnum, bool)`

GetDeniedActionOk returns a tuple with the DeniedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedAction

`func (o *PatchedFlowRequest) SetDeniedAction(v DeniedActionEnum)`

SetDeniedAction sets DeniedAction field to given value.

### HasDeniedAction

`func (o *PatchedFlowRequest) HasDeniedAction() bool`

HasDeniedAction returns a boolean if a field has been set.

### GetAuthentication

`func (o *PatchedFlowRequest) GetAuthentication() AuthenticationEnum`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *PatchedFlowRequest) GetAuthenticationOk() (*AuthenticationEnum, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *PatchedFlowRequest) SetAuthentication(v AuthenticationEnum)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *PatchedFlowRequest) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


