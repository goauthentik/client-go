# IdentificationChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ChallengeChoices**](ChallengeChoices.md) |  | 
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-identification"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**UserFields** | **[]string** |  | 
**PasswordFields** | **bool** |  | 
**ApplicationPre** | Pointer to **string** |  | [optional] 
**EnrollUrl** | Pointer to **string** |  | [optional] 
**RecoveryUrl** | Pointer to **string** |  | [optional] 
**PasswordlessUrl** | Pointer to **string** |  | [optional] 
**PrimaryAction** | **string** |  | 
**Sources** | Pointer to [**[]LoginSource**](LoginSource.md) |  | [optional] 
**ShowSourceLabels** | **bool** |  | 

## Methods

### NewIdentificationChallenge

`func NewIdentificationChallenge(type_ ChallengeChoices, userFields []string, passwordFields bool, primaryAction string, showSourceLabels bool, ) *IdentificationChallenge`

NewIdentificationChallenge instantiates a new IdentificationChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentificationChallengeWithDefaults

`func NewIdentificationChallengeWithDefaults() *IdentificationChallenge`

NewIdentificationChallengeWithDefaults instantiates a new IdentificationChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IdentificationChallenge) GetType() ChallengeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentificationChallenge) GetTypeOk() (*ChallengeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentificationChallenge) SetType(v ChallengeChoices)`

SetType sets Type field to given value.


### GetFlowInfo

`func (o *IdentificationChallenge) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *IdentificationChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *IdentificationChallenge) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *IdentificationChallenge) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *IdentificationChallenge) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *IdentificationChallenge) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *IdentificationChallenge) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *IdentificationChallenge) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *IdentificationChallenge) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *IdentificationChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *IdentificationChallenge) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *IdentificationChallenge) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetUserFields

`func (o *IdentificationChallenge) GetUserFields() []string`

GetUserFields returns the UserFields field if non-nil, zero value otherwise.

### GetUserFieldsOk

`func (o *IdentificationChallenge) GetUserFieldsOk() (*[]string, bool)`

GetUserFieldsOk returns a tuple with the UserFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFields

`func (o *IdentificationChallenge) SetUserFields(v []string)`

SetUserFields sets UserFields field to given value.


### SetUserFieldsNil

`func (o *IdentificationChallenge) SetUserFieldsNil(b bool)`

 SetUserFieldsNil sets the value for UserFields to be an explicit nil

### UnsetUserFields
`func (o *IdentificationChallenge) UnsetUserFields()`

UnsetUserFields ensures that no value is present for UserFields, not even an explicit nil
### GetPasswordFields

`func (o *IdentificationChallenge) GetPasswordFields() bool`

GetPasswordFields returns the PasswordFields field if non-nil, zero value otherwise.

### GetPasswordFieldsOk

`func (o *IdentificationChallenge) GetPasswordFieldsOk() (*bool, bool)`

GetPasswordFieldsOk returns a tuple with the PasswordFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFields

`func (o *IdentificationChallenge) SetPasswordFields(v bool)`

SetPasswordFields sets PasswordFields field to given value.


### GetApplicationPre

`func (o *IdentificationChallenge) GetApplicationPre() string`

GetApplicationPre returns the ApplicationPre field if non-nil, zero value otherwise.

### GetApplicationPreOk

`func (o *IdentificationChallenge) GetApplicationPreOk() (*string, bool)`

GetApplicationPreOk returns a tuple with the ApplicationPre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPre

`func (o *IdentificationChallenge) SetApplicationPre(v string)`

SetApplicationPre sets ApplicationPre field to given value.

### HasApplicationPre

`func (o *IdentificationChallenge) HasApplicationPre() bool`

HasApplicationPre returns a boolean if a field has been set.

### GetEnrollUrl

`func (o *IdentificationChallenge) GetEnrollUrl() string`

GetEnrollUrl returns the EnrollUrl field if non-nil, zero value otherwise.

### GetEnrollUrlOk

`func (o *IdentificationChallenge) GetEnrollUrlOk() (*string, bool)`

GetEnrollUrlOk returns a tuple with the EnrollUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollUrl

`func (o *IdentificationChallenge) SetEnrollUrl(v string)`

SetEnrollUrl sets EnrollUrl field to given value.

### HasEnrollUrl

`func (o *IdentificationChallenge) HasEnrollUrl() bool`

HasEnrollUrl returns a boolean if a field has been set.

### GetRecoveryUrl

`func (o *IdentificationChallenge) GetRecoveryUrl() string`

GetRecoveryUrl returns the RecoveryUrl field if non-nil, zero value otherwise.

### GetRecoveryUrlOk

`func (o *IdentificationChallenge) GetRecoveryUrlOk() (*string, bool)`

GetRecoveryUrlOk returns a tuple with the RecoveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryUrl

`func (o *IdentificationChallenge) SetRecoveryUrl(v string)`

SetRecoveryUrl sets RecoveryUrl field to given value.

### HasRecoveryUrl

`func (o *IdentificationChallenge) HasRecoveryUrl() bool`

HasRecoveryUrl returns a boolean if a field has been set.

### GetPasswordlessUrl

`func (o *IdentificationChallenge) GetPasswordlessUrl() string`

GetPasswordlessUrl returns the PasswordlessUrl field if non-nil, zero value otherwise.

### GetPasswordlessUrlOk

`func (o *IdentificationChallenge) GetPasswordlessUrlOk() (*string, bool)`

GetPasswordlessUrlOk returns a tuple with the PasswordlessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordlessUrl

`func (o *IdentificationChallenge) SetPasswordlessUrl(v string)`

SetPasswordlessUrl sets PasswordlessUrl field to given value.

### HasPasswordlessUrl

`func (o *IdentificationChallenge) HasPasswordlessUrl() bool`

HasPasswordlessUrl returns a boolean if a field has been set.

### GetPrimaryAction

`func (o *IdentificationChallenge) GetPrimaryAction() string`

GetPrimaryAction returns the PrimaryAction field if non-nil, zero value otherwise.

### GetPrimaryActionOk

`func (o *IdentificationChallenge) GetPrimaryActionOk() (*string, bool)`

GetPrimaryActionOk returns a tuple with the PrimaryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAction

`func (o *IdentificationChallenge) SetPrimaryAction(v string)`

SetPrimaryAction sets PrimaryAction field to given value.


### GetSources

`func (o *IdentificationChallenge) GetSources() []LoginSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *IdentificationChallenge) GetSourcesOk() (*[]LoginSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *IdentificationChallenge) SetSources(v []LoginSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *IdentificationChallenge) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetShowSourceLabels

`func (o *IdentificationChallenge) GetShowSourceLabels() bool`

GetShowSourceLabels returns the ShowSourceLabels field if non-nil, zero value otherwise.

### GetShowSourceLabelsOk

`func (o *IdentificationChallenge) GetShowSourceLabelsOk() (*bool, bool)`

GetShowSourceLabelsOk returns a tuple with the ShowSourceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSourceLabels

`func (o *IdentificationChallenge) SetShowSourceLabels(v bool)`

SetShowSourceLabels sets ShowSourceLabels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


