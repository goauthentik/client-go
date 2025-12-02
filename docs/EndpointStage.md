# EndpointStage

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
**Connector** | **string** |  | 
**ConnectorObj** | [**Connector**](Connector.md) |  | [readonly] 
**Mode** | Pointer to [**StageModeEnum**](StageModeEnum.md) |  | [optional] 

## Methods

### NewEndpointStage

`func NewEndpointStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, connector string, connectorObj Connector, ) *EndpointStage`

NewEndpointStage instantiates a new EndpointStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointStageWithDefaults

`func NewEndpointStageWithDefaults() *EndpointStage`

NewEndpointStageWithDefaults instantiates a new EndpointStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *EndpointStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *EndpointStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *EndpointStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *EndpointStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *EndpointStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *EndpointStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *EndpointStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *EndpointStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *EndpointStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *EndpointStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *EndpointStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *EndpointStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *EndpointStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *EndpointStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *EndpointStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *EndpointStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *EndpointStage) GetFlowSet() []FlowSet`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *EndpointStage) GetFlowSetOk() (*[]FlowSet, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *EndpointStage) SetFlowSet(v []FlowSet)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *EndpointStage) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetConnector

`func (o *EndpointStage) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *EndpointStage) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *EndpointStage) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetConnectorObj

`func (o *EndpointStage) GetConnectorObj() Connector`

GetConnectorObj returns the ConnectorObj field if non-nil, zero value otherwise.

### GetConnectorObjOk

`func (o *EndpointStage) GetConnectorObjOk() (*Connector, bool)`

GetConnectorObjOk returns a tuple with the ConnectorObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorObj

`func (o *EndpointStage) SetConnectorObj(v Connector)`

SetConnectorObj sets ConnectorObj field to given value.


### GetMode

`func (o *EndpointStage) GetMode() StageModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EndpointStage) GetModeOk() (*StageModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EndpointStage) SetMode(v StageModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *EndpointStage) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


