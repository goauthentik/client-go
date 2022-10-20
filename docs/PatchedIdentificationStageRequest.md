# PatchedIdentificationStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**UserFields** | Pointer to [**[]UserFieldsEnum**](UserFieldsEnum.md) | Fields of the user object to match against. (Hold shift to select multiple options) | [optional] 
**PasswordStage** | Pointer to **NullableString** | When set, shows a password field, instead of showing the password field as seaprate step. | [optional] 
**CaseInsensitiveMatching** | Pointer to **bool** | When enabled, user fields are matched regardless of their casing. | [optional] 
**ShowMatchedUser** | Pointer to **bool** | When a valid username/email has been entered, and this option is enabled, the user&#39;s username and avatar will be shown. Otherwise, the text that the user entered will be shown | [optional] 
**EnrollmentFlow** | Pointer to **NullableString** | Optional enrollment flow, which is linked at the bottom of the page. | [optional] 
**RecoveryFlow** | Pointer to **NullableString** | Optional recovery flow, which is linked at the bottom of the page. | [optional] 
**PasswordlessFlow** | Pointer to **NullableString** | Optional passwordless flow, which is linked at the bottom of the page. | [optional] 
**Sources** | Pointer to **[]string** | Specify which sources should be shown. | [optional] 
**ShowSourceLabels** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedIdentificationStageRequest

`func NewPatchedIdentificationStageRequest() *PatchedIdentificationStageRequest`

NewPatchedIdentificationStageRequest instantiates a new PatchedIdentificationStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedIdentificationStageRequestWithDefaults

`func NewPatchedIdentificationStageRequestWithDefaults() *PatchedIdentificationStageRequest`

NewPatchedIdentificationStageRequestWithDefaults instantiates a new PatchedIdentificationStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedIdentificationStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedIdentificationStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedIdentificationStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedIdentificationStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlowSet

`func (o *PatchedIdentificationStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PatchedIdentificationStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PatchedIdentificationStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *PatchedIdentificationStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetUserFields

`func (o *PatchedIdentificationStageRequest) GetUserFields() []UserFieldsEnum`

GetUserFields returns the UserFields field if non-nil, zero value otherwise.

### GetUserFieldsOk

`func (o *PatchedIdentificationStageRequest) GetUserFieldsOk() (*[]UserFieldsEnum, bool)`

GetUserFieldsOk returns a tuple with the UserFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFields

`func (o *PatchedIdentificationStageRequest) SetUserFields(v []UserFieldsEnum)`

SetUserFields sets UserFields field to given value.

### HasUserFields

`func (o *PatchedIdentificationStageRequest) HasUserFields() bool`

HasUserFields returns a boolean if a field has been set.

### GetPasswordStage

`func (o *PatchedIdentificationStageRequest) GetPasswordStage() string`

GetPasswordStage returns the PasswordStage field if non-nil, zero value otherwise.

### GetPasswordStageOk

`func (o *PatchedIdentificationStageRequest) GetPasswordStageOk() (*string, bool)`

GetPasswordStageOk returns a tuple with the PasswordStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordStage

`func (o *PatchedIdentificationStageRequest) SetPasswordStage(v string)`

SetPasswordStage sets PasswordStage field to given value.

### HasPasswordStage

`func (o *PatchedIdentificationStageRequest) HasPasswordStage() bool`

HasPasswordStage returns a boolean if a field has been set.

### SetPasswordStageNil

`func (o *PatchedIdentificationStageRequest) SetPasswordStageNil(b bool)`

 SetPasswordStageNil sets the value for PasswordStage to be an explicit nil

### UnsetPasswordStage
`func (o *PatchedIdentificationStageRequest) UnsetPasswordStage()`

UnsetPasswordStage ensures that no value is present for PasswordStage, not even an explicit nil
### GetCaseInsensitiveMatching

`func (o *PatchedIdentificationStageRequest) GetCaseInsensitiveMatching() bool`

GetCaseInsensitiveMatching returns the CaseInsensitiveMatching field if non-nil, zero value otherwise.

### GetCaseInsensitiveMatchingOk

`func (o *PatchedIdentificationStageRequest) GetCaseInsensitiveMatchingOk() (*bool, bool)`

GetCaseInsensitiveMatchingOk returns a tuple with the CaseInsensitiveMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitiveMatching

`func (o *PatchedIdentificationStageRequest) SetCaseInsensitiveMatching(v bool)`

SetCaseInsensitiveMatching sets CaseInsensitiveMatching field to given value.

### HasCaseInsensitiveMatching

`func (o *PatchedIdentificationStageRequest) HasCaseInsensitiveMatching() bool`

HasCaseInsensitiveMatching returns a boolean if a field has been set.

### GetShowMatchedUser

`func (o *PatchedIdentificationStageRequest) GetShowMatchedUser() bool`

GetShowMatchedUser returns the ShowMatchedUser field if non-nil, zero value otherwise.

### GetShowMatchedUserOk

`func (o *PatchedIdentificationStageRequest) GetShowMatchedUserOk() (*bool, bool)`

GetShowMatchedUserOk returns a tuple with the ShowMatchedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMatchedUser

`func (o *PatchedIdentificationStageRequest) SetShowMatchedUser(v bool)`

SetShowMatchedUser sets ShowMatchedUser field to given value.

### HasShowMatchedUser

`func (o *PatchedIdentificationStageRequest) HasShowMatchedUser() bool`

HasShowMatchedUser returns a boolean if a field has been set.

### GetEnrollmentFlow

`func (o *PatchedIdentificationStageRequest) GetEnrollmentFlow() string`

GetEnrollmentFlow returns the EnrollmentFlow field if non-nil, zero value otherwise.

### GetEnrollmentFlowOk

`func (o *PatchedIdentificationStageRequest) GetEnrollmentFlowOk() (*string, bool)`

GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlow

`func (o *PatchedIdentificationStageRequest) SetEnrollmentFlow(v string)`

SetEnrollmentFlow sets EnrollmentFlow field to given value.

### HasEnrollmentFlow

`func (o *PatchedIdentificationStageRequest) HasEnrollmentFlow() bool`

HasEnrollmentFlow returns a boolean if a field has been set.

### SetEnrollmentFlowNil

`func (o *PatchedIdentificationStageRequest) SetEnrollmentFlowNil(b bool)`

 SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil

### UnsetEnrollmentFlow
`func (o *PatchedIdentificationStageRequest) UnsetEnrollmentFlow()`

UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
### GetRecoveryFlow

`func (o *PatchedIdentificationStageRequest) GetRecoveryFlow() string`

GetRecoveryFlow returns the RecoveryFlow field if non-nil, zero value otherwise.

### GetRecoveryFlowOk

`func (o *PatchedIdentificationStageRequest) GetRecoveryFlowOk() (*string, bool)`

GetRecoveryFlowOk returns a tuple with the RecoveryFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryFlow

`func (o *PatchedIdentificationStageRequest) SetRecoveryFlow(v string)`

SetRecoveryFlow sets RecoveryFlow field to given value.

### HasRecoveryFlow

`func (o *PatchedIdentificationStageRequest) HasRecoveryFlow() bool`

HasRecoveryFlow returns a boolean if a field has been set.

### SetRecoveryFlowNil

`func (o *PatchedIdentificationStageRequest) SetRecoveryFlowNil(b bool)`

 SetRecoveryFlowNil sets the value for RecoveryFlow to be an explicit nil

### UnsetRecoveryFlow
`func (o *PatchedIdentificationStageRequest) UnsetRecoveryFlow()`

UnsetRecoveryFlow ensures that no value is present for RecoveryFlow, not even an explicit nil
### GetPasswordlessFlow

`func (o *PatchedIdentificationStageRequest) GetPasswordlessFlow() string`

GetPasswordlessFlow returns the PasswordlessFlow field if non-nil, zero value otherwise.

### GetPasswordlessFlowOk

`func (o *PatchedIdentificationStageRequest) GetPasswordlessFlowOk() (*string, bool)`

GetPasswordlessFlowOk returns a tuple with the PasswordlessFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordlessFlow

`func (o *PatchedIdentificationStageRequest) SetPasswordlessFlow(v string)`

SetPasswordlessFlow sets PasswordlessFlow field to given value.

### HasPasswordlessFlow

`func (o *PatchedIdentificationStageRequest) HasPasswordlessFlow() bool`

HasPasswordlessFlow returns a boolean if a field has been set.

### SetPasswordlessFlowNil

`func (o *PatchedIdentificationStageRequest) SetPasswordlessFlowNil(b bool)`

 SetPasswordlessFlowNil sets the value for PasswordlessFlow to be an explicit nil

### UnsetPasswordlessFlow
`func (o *PatchedIdentificationStageRequest) UnsetPasswordlessFlow()`

UnsetPasswordlessFlow ensures that no value is present for PasswordlessFlow, not even an explicit nil
### GetSources

`func (o *PatchedIdentificationStageRequest) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *PatchedIdentificationStageRequest) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *PatchedIdentificationStageRequest) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *PatchedIdentificationStageRequest) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetShowSourceLabels

`func (o *PatchedIdentificationStageRequest) GetShowSourceLabels() bool`

GetShowSourceLabels returns the ShowSourceLabels field if non-nil, zero value otherwise.

### GetShowSourceLabelsOk

`func (o *PatchedIdentificationStageRequest) GetShowSourceLabelsOk() (*bool, bool)`

GetShowSourceLabelsOk returns a tuple with the ShowSourceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSourceLabels

`func (o *PatchedIdentificationStageRequest) SetShowSourceLabels(v bool)`

SetShowSourceLabels sets ShowSourceLabels field to given value.

### HasShowSourceLabels

`func (o *PatchedIdentificationStageRequest) HasShowSourceLabels() bool`

HasShowSourceLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


