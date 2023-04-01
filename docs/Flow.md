# Flow

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
**Stages** | **[]string** |  | [readonly] 
**Policies** | **[]string** |  | [readonly] 
**CacheCount** | **int32** | Get count of cached flows | [readonly] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**CompatibilityMode** | Pointer to **bool** | Enable compatibility mode, increases compatibility with password managers on mobile devices. | [optional] 
**ExportUrl** | **string** | Get export URL for flow | [readonly] 
**Layout** | Pointer to [**LayoutEnum**](LayoutEnum.md) |  | [optional] 
**DeniedAction** | Pointer to [**DeniedActionEnum**](DeniedActionEnum.md) |  | [optional] 
**Authentication** | Pointer to [**AuthenticationEnum**](AuthenticationEnum.md) |  | [optional] 

## Methods

### NewFlow

`func NewFlow(pk string, policybindingmodelPtrId string, name string, slug string, title string, designation FlowDesignationEnum, background string, stages []string, policies []string, cacheCount int32, exportUrl string, ) *Flow`

NewFlow instantiates a new Flow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowWithDefaults

`func NewFlowWithDefaults() *Flow`

NewFlowWithDefaults instantiates a new Flow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *Flow) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *Flow) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *Flow) SetPk(v string)`

SetPk sets Pk field to given value.


### GetPolicybindingmodelPtrId

`func (o *Flow) GetPolicybindingmodelPtrId() string`

GetPolicybindingmodelPtrId returns the PolicybindingmodelPtrId field if non-nil, zero value otherwise.

### GetPolicybindingmodelPtrIdOk

`func (o *Flow) GetPolicybindingmodelPtrIdOk() (*string, bool)`

GetPolicybindingmodelPtrIdOk returns a tuple with the PolicybindingmodelPtrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicybindingmodelPtrId

`func (o *Flow) SetPolicybindingmodelPtrId(v string)`

SetPolicybindingmodelPtrId sets PolicybindingmodelPtrId field to given value.


### GetName

`func (o *Flow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Flow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Flow) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *Flow) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Flow) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Flow) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTitle

`func (o *Flow) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Flow) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Flow) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDesignation

`func (o *Flow) GetDesignation() FlowDesignationEnum`

GetDesignation returns the Designation field if non-nil, zero value otherwise.

### GetDesignationOk

`func (o *Flow) GetDesignationOk() (*FlowDesignationEnum, bool)`

GetDesignationOk returns a tuple with the Designation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignation

`func (o *Flow) SetDesignation(v FlowDesignationEnum)`

SetDesignation sets Designation field to given value.


### GetBackground

`func (o *Flow) GetBackground() string`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *Flow) GetBackgroundOk() (*string, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *Flow) SetBackground(v string)`

SetBackground sets Background field to given value.


### GetStages

`func (o *Flow) GetStages() []string`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *Flow) GetStagesOk() (*[]string, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *Flow) SetStages(v []string)`

SetStages sets Stages field to given value.


### GetPolicies

`func (o *Flow) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *Flow) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *Flow) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### GetCacheCount

`func (o *Flow) GetCacheCount() int32`

GetCacheCount returns the CacheCount field if non-nil, zero value otherwise.

### GetCacheCountOk

`func (o *Flow) GetCacheCountOk() (*int32, bool)`

GetCacheCountOk returns a tuple with the CacheCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCount

`func (o *Flow) SetCacheCount(v int32)`

SetCacheCount sets CacheCount field to given value.


### GetPolicyEngineMode

`func (o *Flow) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *Flow) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *Flow) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *Flow) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetCompatibilityMode

`func (o *Flow) GetCompatibilityMode() bool`

GetCompatibilityMode returns the CompatibilityMode field if non-nil, zero value otherwise.

### GetCompatibilityModeOk

`func (o *Flow) GetCompatibilityModeOk() (*bool, bool)`

GetCompatibilityModeOk returns a tuple with the CompatibilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityMode

`func (o *Flow) SetCompatibilityMode(v bool)`

SetCompatibilityMode sets CompatibilityMode field to given value.

### HasCompatibilityMode

`func (o *Flow) HasCompatibilityMode() bool`

HasCompatibilityMode returns a boolean if a field has been set.

### GetExportUrl

`func (o *Flow) GetExportUrl() string`

GetExportUrl returns the ExportUrl field if non-nil, zero value otherwise.

### GetExportUrlOk

`func (o *Flow) GetExportUrlOk() (*string, bool)`

GetExportUrlOk returns a tuple with the ExportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportUrl

`func (o *Flow) SetExportUrl(v string)`

SetExportUrl sets ExportUrl field to given value.


### GetLayout

`func (o *Flow) GetLayout() LayoutEnum`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *Flow) GetLayoutOk() (*LayoutEnum, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *Flow) SetLayout(v LayoutEnum)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *Flow) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetDeniedAction

`func (o *Flow) GetDeniedAction() DeniedActionEnum`

GetDeniedAction returns the DeniedAction field if non-nil, zero value otherwise.

### GetDeniedActionOk

`func (o *Flow) GetDeniedActionOk() (*DeniedActionEnum, bool)`

GetDeniedActionOk returns a tuple with the DeniedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedAction

`func (o *Flow) SetDeniedAction(v DeniedActionEnum)`

SetDeniedAction sets DeniedAction field to given value.

### HasDeniedAction

`func (o *Flow) HasDeniedAction() bool`

HasDeniedAction returns a boolean if a field has been set.

### GetAuthentication

`func (o *Flow) GetAuthentication() AuthenticationEnum`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *Flow) GetAuthenticationOk() (*AuthenticationEnum, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *Flow) SetAuthentication(v AuthenticationEnum)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *Flow) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


