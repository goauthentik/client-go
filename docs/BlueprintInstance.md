# BlueprintInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Path** | **string** |  | 
**Context** | **map[string]interface{}** |  | 
**LastApplied** | **time.Time** |  | [readonly] 
**Status** | [**BlueprintInstanceStatusEnum**](BlueprintInstanceStatusEnum.md) |  | 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewBlueprintInstance

`func NewBlueprintInstance(name string, path string, context map[string]interface{}, lastApplied time.Time, status BlueprintInstanceStatusEnum, ) *BlueprintInstance`

NewBlueprintInstance instantiates a new BlueprintInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintInstanceWithDefaults

`func NewBlueprintInstanceWithDefaults() *BlueprintInstance`

NewBlueprintInstanceWithDefaults instantiates a new BlueprintInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


