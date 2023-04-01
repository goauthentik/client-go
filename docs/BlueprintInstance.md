# BlueprintInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Path** | Pointer to **string** |  | [optional] [default to ""]
**Context** | Pointer to **map[string]interface{}** |  | [optional] 
**LastApplied** | **time.Time** |  | [readonly] 
**LastAppliedHash** | **string** |  | [readonly] 
**Status** | [**BlueprintInstanceStatusEnum**](BlueprintInstanceStatusEnum.md) |  | [readonly] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ManagedModels** | **[]string** |  | [readonly] 
**Metadata** | **map[string]interface{}** |  | [readonly] 
**Content** | Pointer to **string** |  | [optional] 

## Methods

### NewBlueprintInstance

`func NewBlueprintInstance(pk string, name string, lastApplied time.Time, lastAppliedHash string, status BlueprintInstanceStatusEnum, managedModels []string, metadata map[string]interface{}, ) *BlueprintInstance`

NewBlueprintInstance instantiates a new BlueprintInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintInstanceWithDefaults

`func NewBlueprintInstanceWithDefaults() *BlueprintInstance`

NewBlueprintInstanceWithDefaults instantiates a new BlueprintInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *BlueprintInstance) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *BlueprintInstance) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *BlueprintInstance) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *BlueprintInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintInstance) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *BlueprintInstance) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BlueprintInstance) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BlueprintInstance) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BlueprintInstance) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetContext

`func (o *BlueprintInstance) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *BlueprintInstance) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *BlueprintInstance) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *BlueprintInstance) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLastApplied

`func (o *BlueprintInstance) GetLastApplied() time.Time`

GetLastApplied returns the LastApplied field if non-nil, zero value otherwise.

### GetLastAppliedOk

`func (o *BlueprintInstance) GetLastAppliedOk() (*time.Time, bool)`

GetLastAppliedOk returns a tuple with the LastApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastApplied

`func (o *BlueprintInstance) SetLastApplied(v time.Time)`

SetLastApplied sets LastApplied field to given value.


### GetLastAppliedHash

`func (o *BlueprintInstance) GetLastAppliedHash() string`

GetLastAppliedHash returns the LastAppliedHash field if non-nil, zero value otherwise.

### GetLastAppliedHashOk

`func (o *BlueprintInstance) GetLastAppliedHashOk() (*string, bool)`

GetLastAppliedHashOk returns a tuple with the LastAppliedHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAppliedHash

`func (o *BlueprintInstance) SetLastAppliedHash(v string)`

SetLastAppliedHash sets LastAppliedHash field to given value.


### GetStatus

`func (o *BlueprintInstance) GetStatus() BlueprintInstanceStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlueprintInstance) GetStatusOk() (*BlueprintInstanceStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlueprintInstance) SetStatus(v BlueprintInstanceStatusEnum)`

SetStatus sets Status field to given value.


### GetEnabled

`func (o *BlueprintInstance) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BlueprintInstance) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BlueprintInstance) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BlueprintInstance) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetManagedModels

`func (o *BlueprintInstance) GetManagedModels() []string`

GetManagedModels returns the ManagedModels field if non-nil, zero value otherwise.

### GetManagedModelsOk

`func (o *BlueprintInstance) GetManagedModelsOk() (*[]string, bool)`

GetManagedModelsOk returns a tuple with the ManagedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedModels

`func (o *BlueprintInstance) SetManagedModels(v []string)`

SetManagedModels sets ManagedModels field to given value.


### GetMetadata

`func (o *BlueprintInstance) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BlueprintInstance) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BlueprintInstance) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetContent

`func (o *BlueprintInstance) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BlueprintInstance) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BlueprintInstance) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *BlueprintInstance) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


