# InvitationFlowObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**PolicybindingmodelPtrId** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** | Visible in the URL. | 
**Title** | **string** | Shown as the Title in Flow pages. | 
**Designation** | [**NullableFlowDesignationEnum**](FlowDesignationEnum.md) | Decides what this Flow is used for. For example, the Authentication flow is redirect to when an un-authenticated user visits authentik.  * &#x60;authentication&#x60; - Authentication * &#x60;authorization&#x60; - Authorization * &#x60;invalidation&#x60; - Invalidation * &#x60;enrollment&#x60; - Enrollment * &#x60;unenrollment&#x60; - Unrenollment * &#x60;recovery&#x60; - Recovery * &#x60;stage_configuration&#x60; - Stage Configuration | 
**Background** | **string** | Get the URL to the background image. If the name is /static or starts with http it is returned as-is | [readonly] 
**Stages** | **[]string** |  | [readonly] 
**Policies** | **[]string** |  | [readonly] 
**CacheCount** | **int32** | Get count of cached flows | [readonly] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**CompatibilityMode** | Pointer to **bool** | Enable compatibility mode, increases compatibility with password managers on mobile devices. | [optional] 
**ExportUrl** | **string** | Get export URL for flow | [readonly] 
**Layout** | Pointer to [**LayoutEnum**](LayoutEnum.md) |  | [optional] 
**DeniedAction** | Pointer to [**NullableDeniedActionEnum**](DeniedActionEnum.md) | Configure what should happen when a flow denies access to a user.  * &#x60;message_continue&#x60; - Message Continue * &#x60;message&#x60; - Message * &#x60;continue&#x60; - Continue | [optional] 
**Authentication** | Pointer to [**NullableAuthenticationEnum**](AuthenticationEnum.md) | Required level of authentication and authorization to access a flow.  * &#x60;none&#x60; - None * &#x60;require_authenticated&#x60; - Require Authenticated * &#x60;require_unauthenticated&#x60; - Require Unauthenticated * &#x60;require_superuser&#x60; - Require Superuser | [optional] 

## Methods

### NewInvitationFlowObj

`func NewInvitationFlowObj(pk string, policybindingmodelPtrId string, name string, slug string, title string, designation NullableFlowDesignationEnum, background string, stages []string, policies []string, cacheCount int32, exportUrl string, ) *InvitationFlowObj`

NewInvitationFlowObj instantiates a new InvitationFlowObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationFlowObjWithDefaults

`func NewInvitationFlowObjWithDefaults() *InvitationFlowObj`

NewInvitationFlowObjWithDefaults instantiates a new InvitationFlowObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *InvitationFlowObj) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *InvitationFlowObj) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *InvitationFlowObj) SetPk(v string)`

SetPk sets Pk field to given value.


### GetPolicybindingmodelPtrId

`func (o *InvitationFlowObj) GetPolicybindingmodelPtrId() string`

GetPolicybindingmodelPtrId returns the PolicybindingmodelPtrId field if non-nil, zero value otherwise.

### GetPolicybindingmodelPtrIdOk

`func (o *InvitationFlowObj) GetPolicybindingmodelPtrIdOk() (*string, bool)`

GetPolicybindingmodelPtrIdOk returns a tuple with the PolicybindingmodelPtrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicybindingmodelPtrId

`func (o *InvitationFlowObj) SetPolicybindingmodelPtrId(v string)`

SetPolicybindingmodelPtrId sets PolicybindingmodelPtrId field to given value.


### GetName

`func (o *InvitationFlowObj) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvitationFlowObj) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvitationFlowObj) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *InvitationFlowObj) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *InvitationFlowObj) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *InvitationFlowObj) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTitle

`func (o *InvitationFlowObj) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InvitationFlowObj) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InvitationFlowObj) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDesignation

`func (o *InvitationFlowObj) GetDesignation() FlowDesignationEnum`

GetDesignation returns the Designation field if non-nil, zero value otherwise.

### GetDesignationOk

`func (o *InvitationFlowObj) GetDesignationOk() (*FlowDesignationEnum, bool)`

GetDesignationOk returns a tuple with the Designation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignation

`func (o *InvitationFlowObj) SetDesignation(v FlowDesignationEnum)`

SetDesignation sets Designation field to given value.


### SetDesignationNil

`func (o *InvitationFlowObj) SetDesignationNil(b bool)`

 SetDesignationNil sets the value for Designation to be an explicit nil

### UnsetDesignation
`func (o *InvitationFlowObj) UnsetDesignation()`

UnsetDesignation ensures that no value is present for Designation, not even an explicit nil
### GetBackground

`func (o *InvitationFlowObj) GetBackground() string`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *InvitationFlowObj) GetBackgroundOk() (*string, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *InvitationFlowObj) SetBackground(v string)`

SetBackground sets Background field to given value.


### GetStages

`func (o *InvitationFlowObj) GetStages() []string`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InvitationFlowObj) GetStagesOk() (*[]string, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InvitationFlowObj) SetStages(v []string)`

SetStages sets Stages field to given value.


### GetPolicies

`func (o *InvitationFlowObj) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *InvitationFlowObj) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *InvitationFlowObj) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### GetCacheCount

`func (o *InvitationFlowObj) GetCacheCount() int32`

GetCacheCount returns the CacheCount field if non-nil, zero value otherwise.

### GetCacheCountOk

`func (o *InvitationFlowObj) GetCacheCountOk() (*int32, bool)`

GetCacheCountOk returns a tuple with the CacheCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCount

`func (o *InvitationFlowObj) SetCacheCount(v int32)`

SetCacheCount sets CacheCount field to given value.


### GetPolicyEngineMode

`func (o *InvitationFlowObj) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *InvitationFlowObj) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *InvitationFlowObj) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *InvitationFlowObj) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetCompatibilityMode

`func (o *InvitationFlowObj) GetCompatibilityMode() bool`

GetCompatibilityMode returns the CompatibilityMode field if non-nil, zero value otherwise.

### GetCompatibilityModeOk

`func (o *InvitationFlowObj) GetCompatibilityModeOk() (*bool, bool)`

GetCompatibilityModeOk returns a tuple with the CompatibilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityMode

`func (o *InvitationFlowObj) SetCompatibilityMode(v bool)`

SetCompatibilityMode sets CompatibilityMode field to given value.

### HasCompatibilityMode

`func (o *InvitationFlowObj) HasCompatibilityMode() bool`

HasCompatibilityMode returns a boolean if a field has been set.

### GetExportUrl

`func (o *InvitationFlowObj) GetExportUrl() string`

GetExportUrl returns the ExportUrl field if non-nil, zero value otherwise.

### GetExportUrlOk

`func (o *InvitationFlowObj) GetExportUrlOk() (*string, bool)`

GetExportUrlOk returns a tuple with the ExportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportUrl

`func (o *InvitationFlowObj) SetExportUrl(v string)`

SetExportUrl sets ExportUrl field to given value.


### GetLayout

`func (o *InvitationFlowObj) GetLayout() LayoutEnum`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *InvitationFlowObj) GetLayoutOk() (*LayoutEnum, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *InvitationFlowObj) SetLayout(v LayoutEnum)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *InvitationFlowObj) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetDeniedAction

`func (o *InvitationFlowObj) GetDeniedAction() DeniedActionEnum`

GetDeniedAction returns the DeniedAction field if non-nil, zero value otherwise.

### GetDeniedActionOk

`func (o *InvitationFlowObj) GetDeniedActionOk() (*DeniedActionEnum, bool)`

GetDeniedActionOk returns a tuple with the DeniedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedAction

`func (o *InvitationFlowObj) SetDeniedAction(v DeniedActionEnum)`

SetDeniedAction sets DeniedAction field to given value.

### HasDeniedAction

`func (o *InvitationFlowObj) HasDeniedAction() bool`

HasDeniedAction returns a boolean if a field has been set.

### SetDeniedActionNil

`func (o *InvitationFlowObj) SetDeniedActionNil(b bool)`

 SetDeniedActionNil sets the value for DeniedAction to be an explicit nil

### UnsetDeniedAction
`func (o *InvitationFlowObj) UnsetDeniedAction()`

UnsetDeniedAction ensures that no value is present for DeniedAction, not even an explicit nil
### GetAuthentication

`func (o *InvitationFlowObj) GetAuthentication() AuthenticationEnum`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InvitationFlowObj) GetAuthenticationOk() (*AuthenticationEnum, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InvitationFlowObj) SetAuthentication(v AuthenticationEnum)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InvitationFlowObj) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### SetAuthenticationNil

`func (o *InvitationFlowObj) SetAuthenticationNil(b bool)`

 SetAuthenticationNil sets the value for Authentication to be an explicit nil

### UnsetAuthentication
`func (o *InvitationFlowObj) UnsetAuthentication()`

UnsetAuthentication ensures that no value is present for Authentication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


