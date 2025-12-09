# SourceStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Component** | **string** | Get object type so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**FlowSet** | [**[]FlowSet**](FlowSet.md) |  | [readonly] 
**Source** | **string** |  | 
**ResumeTimeout** | Pointer to **string** | Amount of time a user can take to return from the source to continue the flow (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3) | [optional] 

## Methods

### NewSourceStage

`func NewSourceStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, flowSet []FlowSet, source string, ) *SourceStage`

NewSourceStage instantiates a new SourceStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceStageWithDefaults

`func NewSourceStageWithDefaults() *SourceStage`

NewSourceStageWithDefaults instantiates a new SourceStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *SourceStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *SourceStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *SourceStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *SourceStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *SourceStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *SourceStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *SourceStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *SourceStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *SourceStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *SourceStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *SourceStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *SourceStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *SourceStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *SourceStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *SourceStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *SourceStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *SourceStage) GetFlowSet() []FlowSet`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *SourceStage) GetFlowSetOk() (*[]FlowSet, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *SourceStage) SetFlowSet(v []FlowSet)`

SetFlowSet sets FlowSet field to given value.


### GetSource

`func (o *SourceStage) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SourceStage) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SourceStage) SetSource(v string)`

SetSource sets Source field to given value.


### GetResumeTimeout

`func (o *SourceStage) GetResumeTimeout() string`

GetResumeTimeout returns the ResumeTimeout field if non-nil, zero value otherwise.

### GetResumeTimeoutOk

`func (o *SourceStage) GetResumeTimeoutOk() (*string, bool)`

GetResumeTimeoutOk returns a tuple with the ResumeTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeTimeout

`func (o *SourceStage) SetResumeTimeout(v string)`

SetResumeTimeout sets ResumeTimeout field to given value.

### HasResumeTimeout

`func (o *SourceStage) HasResumeTimeout() bool`

HasResumeTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


