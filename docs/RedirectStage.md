# RedirectStage

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
**KeepContext** | Pointer to **bool** |  | [optional] 
**Mode** | [**RedirectStageModeEnum**](RedirectStageModeEnum.md) |  | 
**TargetStatic** | Pointer to **string** |  | [optional] 
**TargetFlow** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRedirectStage

`func NewRedirectStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, mode RedirectStageModeEnum, ) *RedirectStage`

NewRedirectStage instantiates a new RedirectStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedirectStageWithDefaults

`func NewRedirectStageWithDefaults() *RedirectStage`

NewRedirectStageWithDefaults instantiates a new RedirectStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *RedirectStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *RedirectStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *RedirectStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *RedirectStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RedirectStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RedirectStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *RedirectStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *RedirectStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *RedirectStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *RedirectStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *RedirectStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *RedirectStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *RedirectStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *RedirectStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *RedirectStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *RedirectStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *RedirectStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *RedirectStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *RedirectStage) GetFlowSet() []FlowSet`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *RedirectStage) GetFlowSetOk() (*[]FlowSet, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *RedirectStage) SetFlowSet(v []FlowSet)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *RedirectStage) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetKeepContext

`func (o *RedirectStage) GetKeepContext() bool`

GetKeepContext returns the KeepContext field if non-nil, zero value otherwise.

### GetKeepContextOk

`func (o *RedirectStage) GetKeepContextOk() (*bool, bool)`

GetKeepContextOk returns a tuple with the KeepContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepContext

`func (o *RedirectStage) SetKeepContext(v bool)`

SetKeepContext sets KeepContext field to given value.

### HasKeepContext

`func (o *RedirectStage) HasKeepContext() bool`

HasKeepContext returns a boolean if a field has been set.

### GetMode

`func (o *RedirectStage) GetMode() RedirectStageModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RedirectStage) GetModeOk() (*RedirectStageModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RedirectStage) SetMode(v RedirectStageModeEnum)`

SetMode sets Mode field to given value.


### GetTargetStatic

`func (o *RedirectStage) GetTargetStatic() string`

GetTargetStatic returns the TargetStatic field if non-nil, zero value otherwise.

### GetTargetStaticOk

`func (o *RedirectStage) GetTargetStaticOk() (*string, bool)`

GetTargetStaticOk returns a tuple with the TargetStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStatic

`func (o *RedirectStage) SetTargetStatic(v string)`

SetTargetStatic sets TargetStatic field to given value.

### HasTargetStatic

`func (o *RedirectStage) HasTargetStatic() bool`

HasTargetStatic returns a boolean if a field has been set.

### GetTargetFlow

`func (o *RedirectStage) GetTargetFlow() string`

GetTargetFlow returns the TargetFlow field if non-nil, zero value otherwise.

### GetTargetFlowOk

`func (o *RedirectStage) GetTargetFlowOk() (*string, bool)`

GetTargetFlowOk returns a tuple with the TargetFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFlow

`func (o *RedirectStage) SetTargetFlow(v string)`

SetTargetFlow sets TargetFlow field to given value.

### HasTargetFlow

`func (o *RedirectStage) HasTargetFlow() bool`

HasTargetFlow returns a boolean if a field has been set.

### SetTargetFlowNil

`func (o *RedirectStage) SetTargetFlowNil(b bool)`

 SetTargetFlowNil sets the value for TargetFlow to be an explicit nil

### UnsetTargetFlow
`func (o *RedirectStage) UnsetTargetFlow()`

UnsetTargetFlow ensures that no value is present for TargetFlow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


