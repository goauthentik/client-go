# UserLoginStage

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
**SessionDuration** | Pointer to **string** | Determines how long a session lasts. Default of 0 means that the sessions lasts until the browser is closed. (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3) | [optional] 
**TerminateOtherSessions** | Pointer to **bool** | Terminate all other sessions of the user logging in. | [optional] 
**RememberMeOffset** | Pointer to **string** | Offset the session will be extended by when the user picks the remember me option. Default of 0 means that the remember me option will not be shown. (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3) | [optional] 

## Methods

### NewUserLoginStage

`func NewUserLoginStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, ) *UserLoginStage`

NewUserLoginStage instantiates a new UserLoginStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLoginStageWithDefaults

`func NewUserLoginStageWithDefaults() *UserLoginStage`

NewUserLoginStageWithDefaults instantiates a new UserLoginStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *UserLoginStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *UserLoginStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *UserLoginStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *UserLoginStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserLoginStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserLoginStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *UserLoginStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *UserLoginStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *UserLoginStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *UserLoginStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *UserLoginStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *UserLoginStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *UserLoginStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *UserLoginStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *UserLoginStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *UserLoginStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *UserLoginStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *UserLoginStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *UserLoginStage) GetFlowSet() []FlowSet`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *UserLoginStage) GetFlowSetOk() (*[]FlowSet, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *UserLoginStage) SetFlowSet(v []FlowSet)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *UserLoginStage) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetSessionDuration

`func (o *UserLoginStage) GetSessionDuration() string`

GetSessionDuration returns the SessionDuration field if non-nil, zero value otherwise.

### GetSessionDurationOk

`func (o *UserLoginStage) GetSessionDurationOk() (*string, bool)`

GetSessionDurationOk returns a tuple with the SessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDuration

`func (o *UserLoginStage) SetSessionDuration(v string)`

SetSessionDuration sets SessionDuration field to given value.

### HasSessionDuration

`func (o *UserLoginStage) HasSessionDuration() bool`

HasSessionDuration returns a boolean if a field has been set.

### GetTerminateOtherSessions

`func (o *UserLoginStage) GetTerminateOtherSessions() bool`

GetTerminateOtherSessions returns the TerminateOtherSessions field if non-nil, zero value otherwise.

### GetTerminateOtherSessionsOk

`func (o *UserLoginStage) GetTerminateOtherSessionsOk() (*bool, bool)`

GetTerminateOtherSessionsOk returns a tuple with the TerminateOtherSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminateOtherSessions

`func (o *UserLoginStage) SetTerminateOtherSessions(v bool)`

SetTerminateOtherSessions sets TerminateOtherSessions field to given value.

### HasTerminateOtherSessions

`func (o *UserLoginStage) HasTerminateOtherSessions() bool`

HasTerminateOtherSessions returns a boolean if a field has been set.

### GetRememberMeOffset

`func (o *UserLoginStage) GetRememberMeOffset() string`

GetRememberMeOffset returns the RememberMeOffset field if non-nil, zero value otherwise.

### GetRememberMeOffsetOk

`func (o *UserLoginStage) GetRememberMeOffsetOk() (*string, bool)`

GetRememberMeOffsetOk returns a tuple with the RememberMeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMeOffset

`func (o *UserLoginStage) SetRememberMeOffset(v string)`

SetRememberMeOffset sets RememberMeOffset field to given value.

### HasRememberMeOffset

`func (o *UserLoginStage) HasRememberMeOffset() bool`

HasRememberMeOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


