# FlowRequest

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
**Authentication** | Pointer to [**AuthenticationEnum**](AuthenticationEnum.md) |  | [optional] 

## Methods

### NewFlowRequest

`func NewFlowRequest(name string, slug string, title string, designation FlowDesignationEnum, ) *FlowRequest`

NewFlowRequest instantiates a new FlowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowRequestWithDefaults

`func NewFlowRequestWithDefaults() *FlowRequest`

NewFlowRequestWithDefaults instantiates a new FlowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FlowRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *FlowRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *FlowRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *FlowRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTitle

`func (o *FlowRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FlowRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FlowRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDesignation

`func (o *FlowRequest) GetDesignation() FlowDesignationEnum`

GetDesignation returns the Designation field if non-nil, zero value otherwise.

### GetDesignationOk

`func (o *FlowRequest) GetDesignationOk() (*FlowDesignationEnum, bool)`

GetDesignationOk returns a tuple with the Designation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignation

`func (o *FlowRequest) SetDesignation(v FlowDesignationEnum)`

SetDesignation sets Designation field to given value.


### GetPolicyEngineMode

`func (o *FlowRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *FlowRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *FlowRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *FlowRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetCompatibilityMode

`func (o *FlowRequest) GetCompatibilityMode() bool`

GetCompatibilityMode returns the CompatibilityMode field if non-nil, zero value otherwise.

### GetCompatibilityModeOk

`func (o *FlowRequest) GetCompatibilityModeOk() (*bool, bool)`

GetCompatibilityModeOk returns a tuple with the CompatibilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityMode

`func (o *FlowRequest) SetCompatibilityMode(v bool)`

SetCompatibilityMode sets CompatibilityMode field to given value.

### HasCompatibilityMode

`func (o *FlowRequest) HasCompatibilityMode() bool`

HasCompatibilityMode returns a boolean if a field has been set.

### GetLayout

`func (o *FlowRequest) GetLayout() LayoutEnum`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *FlowRequest) GetLayoutOk() (*LayoutEnum, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *FlowRequest) SetLayout(v LayoutEnum)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *FlowRequest) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetDeniedAction

`func (o *FlowRequest) GetDeniedAction() DeniedActionEnum`

GetDeniedAction returns the DeniedAction field if non-nil, zero value otherwise.

### GetDeniedActionOk

`func (o *FlowRequest) GetDeniedActionOk() (*DeniedActionEnum, bool)`

GetDeniedActionOk returns a tuple with the DeniedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedAction

`func (o *FlowRequest) SetDeniedAction(v DeniedActionEnum)`

SetDeniedAction sets DeniedAction field to given value.

### HasDeniedAction

`func (o *FlowRequest) HasDeniedAction() bool`

HasDeniedAction returns a boolean if a field has been set.

### GetAuthentication

`func (o *FlowRequest) GetAuthentication() AuthenticationEnum`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *FlowRequest) GetAuthenticationOk() (*AuthenticationEnum, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *FlowRequest) SetAuthentication(v AuthenticationEnum)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *FlowRequest) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


