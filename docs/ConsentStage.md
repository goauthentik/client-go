# ConsentStage

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
**Mode** | Pointer to [**ConsentStageModeEnum**](ConsentStageModeEnum.md) |  | [optional] 
**ConsentExpireIn** | Pointer to **string** | Offset after which consent expires. (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 

## Methods

### NewConsentStage

`func NewConsentStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, ) *ConsentStage`

NewConsentStage instantiates a new ConsentStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentStageWithDefaults

`func NewConsentStageWithDefaults() *ConsentStage`

NewConsentStageWithDefaults instantiates a new ConsentStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *ConsentStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *ConsentStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *ConsentStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *ConsentStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsentStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsentStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *ConsentStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ConsentStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ConsentStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *ConsentStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *ConsentStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *ConsentStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *ConsentStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *ConsentStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *ConsentStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *ConsentStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *ConsentStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *ConsentStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *ConsentStage) GetFlowSet() []FlowSet`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *ConsentStage) GetFlowSetOk() (*[]FlowSet, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *ConsentStage) SetFlowSet(v []FlowSet)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *ConsentStage) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetMode

`func (o *ConsentStage) GetMode() ConsentStageModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ConsentStage) GetModeOk() (*ConsentStageModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ConsentStage) SetMode(v ConsentStageModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ConsentStage) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetConsentExpireIn

`func (o *ConsentStage) GetConsentExpireIn() string`

GetConsentExpireIn returns the ConsentExpireIn field if non-nil, zero value otherwise.

### GetConsentExpireInOk

`func (o *ConsentStage) GetConsentExpireInOk() (*string, bool)`

GetConsentExpireInOk returns a tuple with the ConsentExpireIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentExpireIn

`func (o *ConsentStage) SetConsentExpireIn(v string)`

SetConsentExpireIn sets ConsentExpireIn field to given value.

### HasConsentExpireIn

`func (o *ConsentStage) HasConsentExpireIn() bool`

HasConsentExpireIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


