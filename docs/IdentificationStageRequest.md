# IdentificationStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
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

### NewIdentificationStageRequest

`func NewIdentificationStageRequest(name string, ) *IdentificationStageRequest`

NewIdentificationStageRequest instantiates a new IdentificationStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentificationStageRequestWithDefaults

`func NewIdentificationStageRequestWithDefaults() *IdentificationStageRequest`

NewIdentificationStageRequestWithDefaults instantiates a new IdentificationStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdentificationStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentificationStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentificationStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowSet

`func (o *IdentificationStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *IdentificationStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *IdentificationStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *IdentificationStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetUserFields

`func (o *IdentificationStageRequest) GetUserFields() []UserFieldsEnum`

GetUserFields returns the UserFields field if non-nil, zero value otherwise.

### GetUserFieldsOk

`func (o *IdentificationStageRequest) GetUserFieldsOk() (*[]UserFieldsEnum, bool)`

GetUserFieldsOk returns a tuple with the UserFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFields

`func (o *IdentificationStageRequest) SetUserFields(v []UserFieldsEnum)`

SetUserFields sets UserFields field to given value.

### HasUserFields

`func (o *IdentificationStageRequest) HasUserFields() bool`

HasUserFields returns a boolean if a field has been set.

### GetPasswordStage

`func (o *IdentificationStageRequest) GetPasswordStage() string`

GetPasswordStage returns the PasswordStage field if non-nil, zero value otherwise.

### GetPasswordStageOk

`func (o *IdentificationStageRequest) GetPasswordStageOk() (*string, bool)`

GetPasswordStageOk returns a tuple with the PasswordStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordStage

`func (o *IdentificationStageRequest) SetPasswordStage(v string)`

SetPasswordStage sets PasswordStage field to given value.

### HasPasswordStage

`func (o *IdentificationStageRequest) HasPasswordStage() bool`

HasPasswordStage returns a boolean if a field has been set.

### SetPasswordStageNil

`func (o *IdentificationStageRequest) SetPasswordStageNil(b bool)`

 SetPasswordStageNil sets the value for PasswordStage to be an explicit nil

### UnsetPasswordStage
`func (o *IdentificationStageRequest) UnsetPasswordStage()`

UnsetPasswordStage ensures that no value is present for PasswordStage, not even an explicit nil
### GetCaseInsensitiveMatching

`func (o *IdentificationStageRequest) GetCaseInsensitiveMatching() bool`

GetCaseInsensitiveMatching returns the CaseInsensitiveMatching field if non-nil, zero value otherwise.

### GetCaseInsensitiveMatchingOk

`func (o *IdentificationStageRequest) GetCaseInsensitiveMatchingOk() (*bool, bool)`

GetCaseInsensitiveMatchingOk returns a tuple with the CaseInsensitiveMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitiveMatching

`func (o *IdentificationStageRequest) SetCaseInsensitiveMatching(v bool)`

SetCaseInsensitiveMatching sets CaseInsensitiveMatching field to given value.

### HasCaseInsensitiveMatching

`func (o *IdentificationStageRequest) HasCaseInsensitiveMatching() bool`

HasCaseInsensitiveMatching returns a boolean if a field has been set.

### GetShowMatchedUser

`func (o *IdentificationStageRequest) GetShowMatchedUser() bool`

GetShowMatchedUser returns the ShowMatchedUser field if non-nil, zero value otherwise.

### GetShowMatchedUserOk

`func (o *IdentificationStageRequest) GetShowMatchedUserOk() (*bool, bool)`

GetShowMatchedUserOk returns a tuple with the ShowMatchedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMatchedUser

`func (o *IdentificationStageRequest) SetShowMatchedUser(v bool)`

SetShowMatchedUser sets ShowMatchedUser field to given value.

### HasShowMatchedUser

`func (o *IdentificationStageRequest) HasShowMatchedUser() bool`

HasShowMatchedUser returns a boolean if a field has been set.

### GetEnrollmentFlow

`func (o *IdentificationStageRequest) GetEnrollmentFlow() string`

GetEnrollmentFlow returns the EnrollmentFlow field if non-nil, zero value otherwise.

### GetEnrollmentFlowOk

`func (o *IdentificationStageRequest) GetEnrollmentFlowOk() (*string, bool)`

GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlow

`func (o *IdentificationStageRequest) SetEnrollmentFlow(v string)`

SetEnrollmentFlow sets EnrollmentFlow field to given value.

### HasEnrollmentFlow

`func (o *IdentificationStageRequest) HasEnrollmentFlow() bool`

HasEnrollmentFlow returns a boolean if a field has been set.

### SetEnrollmentFlowNil

`func (o *IdentificationStageRequest) SetEnrollmentFlowNil(b bool)`

 SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil

### UnsetEnrollmentFlow
`func (o *IdentificationStageRequest) UnsetEnrollmentFlow()`

UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
### GetRecoveryFlow

`func (o *IdentificationStageRequest) GetRecoveryFlow() string`

GetRecoveryFlow returns the RecoveryFlow field if non-nil, zero value otherwise.

### GetRecoveryFlowOk

`func (o *IdentificationStageRequest) GetRecoveryFlowOk() (*string, bool)`

GetRecoveryFlowOk returns a tuple with the RecoveryFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryFlow

`func (o *IdentificationStageRequest) SetRecoveryFlow(v string)`

SetRecoveryFlow sets RecoveryFlow field to given value.

### HasRecoveryFlow

`func (o *IdentificationStageRequest) HasRecoveryFlow() bool`

HasRecoveryFlow returns a boolean if a field has been set.

### SetRecoveryFlowNil

`func (o *IdentificationStageRequest) SetRecoveryFlowNil(b bool)`

 SetRecoveryFlowNil sets the value for RecoveryFlow to be an explicit nil

### UnsetRecoveryFlow
`func (o *IdentificationStageRequest) UnsetRecoveryFlow()`

UnsetRecoveryFlow ensures that no value is present for RecoveryFlow, not even an explicit nil
### GetPasswordlessFlow

`func (o *IdentificationStageRequest) GetPasswordlessFlow() string`

GetPasswordlessFlow returns the PasswordlessFlow field if non-nil, zero value otherwise.

### GetPasswordlessFlowOk

`func (o *IdentificationStageRequest) GetPasswordlessFlowOk() (*string, bool)`

GetPasswordlessFlowOk returns a tuple with the PasswordlessFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordlessFlow

`func (o *IdentificationStageRequest) SetPasswordlessFlow(v string)`

SetPasswordlessFlow sets PasswordlessFlow field to given value.

### HasPasswordlessFlow

`func (o *IdentificationStageRequest) HasPasswordlessFlow() bool`

HasPasswordlessFlow returns a boolean if a field has been set.

### SetPasswordlessFlowNil

`func (o *IdentificationStageRequest) SetPasswordlessFlowNil(b bool)`

 SetPasswordlessFlowNil sets the value for PasswordlessFlow to be an explicit nil

### UnsetPasswordlessFlow
`func (o *IdentificationStageRequest) UnsetPasswordlessFlow()`

UnsetPasswordlessFlow ensures that no value is present for PasswordlessFlow, not even an explicit nil
### GetSources

`func (o *IdentificationStageRequest) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *IdentificationStageRequest) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *IdentificationStageRequest) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *IdentificationStageRequest) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetShowSourceLabels

`func (o *IdentificationStageRequest) GetShowSourceLabels() bool`

GetShowSourceLabels returns the ShowSourceLabels field if non-nil, zero value otherwise.

### GetShowSourceLabelsOk

`func (o *IdentificationStageRequest) GetShowSourceLabelsOk() (*bool, bool)`

GetShowSourceLabelsOk returns a tuple with the ShowSourceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSourceLabels

`func (o *IdentificationStageRequest) SetShowSourceLabels(v bool)`

SetShowSourceLabels sets ShowSourceLabels field to given value.

### HasShowSourceLabels

`func (o *IdentificationStageRequest) HasShowSourceLabels() bool`

HasShowSourceLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


