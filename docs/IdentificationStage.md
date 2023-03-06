# IdentificationStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Component** | **string** | Get object type so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**FlowSet** | Pointer to [**[]FlowSet**](FlowSet.md) |  | [optional] 
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

### NewIdentificationStage

`func NewIdentificationStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, ) *IdentificationStage`

NewIdentificationStage instantiates a new IdentificationStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentificationStageWithDefaults

`func NewIdentificationStageWithDefaults() *IdentificationStage`

NewIdentificationStageWithDefaults instantiates a new IdentificationStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *IdentificationStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *IdentificationStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *IdentificationStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *IdentificationStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentificationStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentificationStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *IdentificationStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *IdentificationStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *IdentificationStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *IdentificationStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *IdentificationStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *IdentificationStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *IdentificationStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *IdentificationStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *IdentificationStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *IdentificationStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *IdentificationStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *IdentificationStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *IdentificationStage) GetFlowSet() []FlowSet`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *IdentificationStage) GetFlowSetOk() (*[]FlowSet, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *IdentificationStage) SetFlowSet(v []FlowSet)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *IdentificationStage) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetUserFields

`func (o *IdentificationStage) GetUserFields() []UserFieldsEnum`

GetUserFields returns the UserFields field if non-nil, zero value otherwise.

### GetUserFieldsOk

`func (o *IdentificationStage) GetUserFieldsOk() (*[]UserFieldsEnum, bool)`

GetUserFieldsOk returns a tuple with the UserFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFields

`func (o *IdentificationStage) SetUserFields(v []UserFieldsEnum)`

SetUserFields sets UserFields field to given value.

### HasUserFields

`func (o *IdentificationStage) HasUserFields() bool`

HasUserFields returns a boolean if a field has been set.

### GetPasswordStage

`func (o *IdentificationStage) GetPasswordStage() string`

GetPasswordStage returns the PasswordStage field if non-nil, zero value otherwise.

### GetPasswordStageOk

`func (o *IdentificationStage) GetPasswordStageOk() (*string, bool)`

GetPasswordStageOk returns a tuple with the PasswordStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordStage

`func (o *IdentificationStage) SetPasswordStage(v string)`

SetPasswordStage sets PasswordStage field to given value.

### HasPasswordStage

`func (o *IdentificationStage) HasPasswordStage() bool`

HasPasswordStage returns a boolean if a field has been set.

### SetPasswordStageNil

`func (o *IdentificationStage) SetPasswordStageNil(b bool)`

 SetPasswordStageNil sets the value for PasswordStage to be an explicit nil

### UnsetPasswordStage
`func (o *IdentificationStage) UnsetPasswordStage()`

UnsetPasswordStage ensures that no value is present for PasswordStage, not even an explicit nil
### GetCaseInsensitiveMatching

`func (o *IdentificationStage) GetCaseInsensitiveMatching() bool`

GetCaseInsensitiveMatching returns the CaseInsensitiveMatching field if non-nil, zero value otherwise.

### GetCaseInsensitiveMatchingOk

`func (o *IdentificationStage) GetCaseInsensitiveMatchingOk() (*bool, bool)`

GetCaseInsensitiveMatchingOk returns a tuple with the CaseInsensitiveMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitiveMatching

`func (o *IdentificationStage) SetCaseInsensitiveMatching(v bool)`

SetCaseInsensitiveMatching sets CaseInsensitiveMatching field to given value.

### HasCaseInsensitiveMatching

`func (o *IdentificationStage) HasCaseInsensitiveMatching() bool`

HasCaseInsensitiveMatching returns a boolean if a field has been set.

### GetShowMatchedUser

`func (o *IdentificationStage) GetShowMatchedUser() bool`

GetShowMatchedUser returns the ShowMatchedUser field if non-nil, zero value otherwise.

### GetShowMatchedUserOk

`func (o *IdentificationStage) GetShowMatchedUserOk() (*bool, bool)`

GetShowMatchedUserOk returns a tuple with the ShowMatchedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMatchedUser

`func (o *IdentificationStage) SetShowMatchedUser(v bool)`

SetShowMatchedUser sets ShowMatchedUser field to given value.

### HasShowMatchedUser

`func (o *IdentificationStage) HasShowMatchedUser() bool`

HasShowMatchedUser returns a boolean if a field has been set.

### GetEnrollmentFlow

`func (o *IdentificationStage) GetEnrollmentFlow() string`

GetEnrollmentFlow returns the EnrollmentFlow field if non-nil, zero value otherwise.

### GetEnrollmentFlowOk

`func (o *IdentificationStage) GetEnrollmentFlowOk() (*string, bool)`

GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlow

`func (o *IdentificationStage) SetEnrollmentFlow(v string)`

SetEnrollmentFlow sets EnrollmentFlow field to given value.

### HasEnrollmentFlow

`func (o *IdentificationStage) HasEnrollmentFlow() bool`

HasEnrollmentFlow returns a boolean if a field has been set.

### SetEnrollmentFlowNil

`func (o *IdentificationStage) SetEnrollmentFlowNil(b bool)`

 SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil

### UnsetEnrollmentFlow
`func (o *IdentificationStage) UnsetEnrollmentFlow()`

UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
### GetRecoveryFlow

`func (o *IdentificationStage) GetRecoveryFlow() string`

GetRecoveryFlow returns the RecoveryFlow field if non-nil, zero value otherwise.

### GetRecoveryFlowOk

`func (o *IdentificationStage) GetRecoveryFlowOk() (*string, bool)`

GetRecoveryFlowOk returns a tuple with the RecoveryFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryFlow

`func (o *IdentificationStage) SetRecoveryFlow(v string)`

SetRecoveryFlow sets RecoveryFlow field to given value.

### HasRecoveryFlow

`func (o *IdentificationStage) HasRecoveryFlow() bool`

HasRecoveryFlow returns a boolean if a field has been set.

### SetRecoveryFlowNil

`func (o *IdentificationStage) SetRecoveryFlowNil(b bool)`

 SetRecoveryFlowNil sets the value for RecoveryFlow to be an explicit nil

### UnsetRecoveryFlow
`func (o *IdentificationStage) UnsetRecoveryFlow()`

UnsetRecoveryFlow ensures that no value is present for RecoveryFlow, not even an explicit nil
### GetPasswordlessFlow

`func (o *IdentificationStage) GetPasswordlessFlow() string`

GetPasswordlessFlow returns the PasswordlessFlow field if non-nil, zero value otherwise.

### GetPasswordlessFlowOk

`func (o *IdentificationStage) GetPasswordlessFlowOk() (*string, bool)`

GetPasswordlessFlowOk returns a tuple with the PasswordlessFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordlessFlow

`func (o *IdentificationStage) SetPasswordlessFlow(v string)`

SetPasswordlessFlow sets PasswordlessFlow field to given value.

### HasPasswordlessFlow

`func (o *IdentificationStage) HasPasswordlessFlow() bool`

HasPasswordlessFlow returns a boolean if a field has been set.

### SetPasswordlessFlowNil

`func (o *IdentificationStage) SetPasswordlessFlowNil(b bool)`

 SetPasswordlessFlowNil sets the value for PasswordlessFlow to be an explicit nil

### UnsetPasswordlessFlow
`func (o *IdentificationStage) UnsetPasswordlessFlow()`

UnsetPasswordlessFlow ensures that no value is present for PasswordlessFlow, not even an explicit nil
### GetSources

`func (o *IdentificationStage) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *IdentificationStage) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *IdentificationStage) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *IdentificationStage) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetShowSourceLabels

`func (o *IdentificationStage) GetShowSourceLabels() bool`

GetShowSourceLabels returns the ShowSourceLabels field if non-nil, zero value otherwise.

### GetShowSourceLabelsOk

`func (o *IdentificationStage) GetShowSourceLabelsOk() (*bool, bool)`

GetShowSourceLabelsOk returns a tuple with the ShowSourceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSourceLabels

`func (o *IdentificationStage) SetShowSourceLabels(v bool)`

SetShowSourceLabels sets ShowSourceLabels field to given value.

### HasShowSourceLabels

`func (o *IdentificationStage) HasShowSourceLabels() bool`

HasShowSourceLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


