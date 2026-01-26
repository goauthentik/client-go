# InvitationStage

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
**ContinueFlowWithoutInvitation** | Pointer to **bool** | If this flag is set, this Stage will jump to the next Stage when no Invitation is given. By default this Stage will cancel the Flow when no invitation is given. | [optional] 

## Methods

### NewInvitationStage

`func NewInvitationStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, ) *InvitationStage`

NewInvitationStage instantiates a new InvitationStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationStageWithDefaults

`func NewInvitationStageWithDefaults() *InvitationStage`

NewInvitationStageWithDefaults instantiates a new InvitationStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *InvitationStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *InvitationStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *InvitationStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *InvitationStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvitationStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvitationStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *InvitationStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *InvitationStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *InvitationStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *InvitationStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *InvitationStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *InvitationStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *InvitationStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *InvitationStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *InvitationStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *InvitationStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *InvitationStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *InvitationStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *InvitationStage) GetFlowSet() []FlowSet`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *InvitationStage) GetFlowSetOk() (*[]FlowSet, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *InvitationStage) SetFlowSet(v []FlowSet)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *InvitationStage) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetContinueFlowWithoutInvitation

`func (o *InvitationStage) GetContinueFlowWithoutInvitation() bool`

GetContinueFlowWithoutInvitation returns the ContinueFlowWithoutInvitation field if non-nil, zero value otherwise.

### GetContinueFlowWithoutInvitationOk

`func (o *InvitationStage) GetContinueFlowWithoutInvitationOk() (*bool, bool)`

GetContinueFlowWithoutInvitationOk returns a tuple with the ContinueFlowWithoutInvitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueFlowWithoutInvitation

`func (o *InvitationStage) SetContinueFlowWithoutInvitation(v bool)`

SetContinueFlowWithoutInvitation sets ContinueFlowWithoutInvitation field to given value.

### HasContinueFlowWithoutInvitation

`func (o *InvitationStage) HasContinueFlowWithoutInvitation() bool`

HasContinueFlowWithoutInvitation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


