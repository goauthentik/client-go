# PasswordStage

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
**Backends** | [**[]BackendsEnum**](BackendsEnum.md) | Selection of backends to test the password against. | 
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**FailedAttemptsBeforeCancel** | Pointer to **int32** | How many attempts a user has before the flow is canceled. To lock the user out, use a reputation policy and a user_write stage. | [optional] 

## Methods

### NewPasswordStage

`func NewPasswordStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, backends []BackendsEnum, ) *PasswordStage`

NewPasswordStage instantiates a new PasswordStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordStageWithDefaults

`func NewPasswordStageWithDefaults() *PasswordStage`

NewPasswordStageWithDefaults instantiates a new PasswordStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *PasswordStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *PasswordStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *PasswordStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *PasswordStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PasswordStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PasswordStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *PasswordStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *PasswordStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *PasswordStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *PasswordStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *PasswordStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *PasswordStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *PasswordStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *PasswordStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *PasswordStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *PasswordStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *PasswordStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *PasswordStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *PasswordStage) GetFlowSet() []FlowSet`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PasswordStage) GetFlowSetOk() (*[]FlowSet, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PasswordStage) SetFlowSet(v []FlowSet)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *PasswordStage) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetBackends

`func (o *PasswordStage) GetBackends() []BackendsEnum`

GetBackends returns the Backends field if non-nil, zero value otherwise.

### GetBackendsOk

`func (o *PasswordStage) GetBackendsOk() (*[]BackendsEnum, bool)`

GetBackendsOk returns a tuple with the Backends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackends

`func (o *PasswordStage) SetBackends(v []BackendsEnum)`

SetBackends sets Backends field to given value.


### GetConfigureFlow

`func (o *PasswordStage) GetConfigureFlow() string`

GetConfigureFlow returns the ConfigureFlow field if non-nil, zero value otherwise.

### GetConfigureFlowOk

`func (o *PasswordStage) GetConfigureFlowOk() (*string, bool)`

GetConfigureFlowOk returns a tuple with the ConfigureFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureFlow

`func (o *PasswordStage) SetConfigureFlow(v string)`

SetConfigureFlow sets ConfigureFlow field to given value.

### HasConfigureFlow

`func (o *PasswordStage) HasConfigureFlow() bool`

HasConfigureFlow returns a boolean if a field has been set.

### SetConfigureFlowNil

`func (o *PasswordStage) SetConfigureFlowNil(b bool)`

 SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil

### UnsetConfigureFlow
`func (o *PasswordStage) UnsetConfigureFlow()`

UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
### GetFailedAttemptsBeforeCancel

`func (o *PasswordStage) GetFailedAttemptsBeforeCancel() int32`

GetFailedAttemptsBeforeCancel returns the FailedAttemptsBeforeCancel field if non-nil, zero value otherwise.

### GetFailedAttemptsBeforeCancelOk

`func (o *PasswordStage) GetFailedAttemptsBeforeCancelOk() (*int32, bool)`

GetFailedAttemptsBeforeCancelOk returns a tuple with the FailedAttemptsBeforeCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAttemptsBeforeCancel

`func (o *PasswordStage) SetFailedAttemptsBeforeCancel(v int32)`

SetFailedAttemptsBeforeCancel sets FailedAttemptsBeforeCancel field to given value.

### HasFailedAttemptsBeforeCancel

`func (o *PasswordStage) HasFailedAttemptsBeforeCancel() bool`

HasFailedAttemptsBeforeCancel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


