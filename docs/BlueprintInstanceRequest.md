# BlueprintInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Path** | Pointer to **string** |  | [optional] [default to ""]
**Context** | Pointer to **map[string]interface{}** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 

## Methods

### NewBlueprintInstanceRequest

`func NewBlueprintInstanceRequest(name string, ) *BlueprintInstanceRequest`

NewBlueprintInstanceRequest instantiates a new BlueprintInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintInstanceRequestWithDefaults

`func NewBlueprintInstanceRequestWithDefaults() *BlueprintInstanceRequest`

NewBlueprintInstanceRequestWithDefaults instantiates a new BlueprintInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintInstanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintInstanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintInstanceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *BlueprintInstanceRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BlueprintInstanceRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BlueprintInstanceRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BlueprintInstanceRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetContext

`func (o *BlueprintInstanceRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *BlueprintInstanceRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *BlueprintInstanceRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *BlueprintInstanceRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetEnabled

`func (o *BlueprintInstanceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BlueprintInstanceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BlueprintInstanceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BlueprintInstanceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetContent

`func (o *BlueprintInstanceRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BlueprintInstanceRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BlueprintInstanceRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *BlueprintInstanceRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


