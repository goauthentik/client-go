# DiagramNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | [readonly] 
**Type** | [**DiagramNodeTypeEnum**](DiagramNodeTypeEnum.md) |  | [readonly] 
**Name** | **string** |  | [readonly] 
**VerboseName** | **string** |  | [readonly] 
**Model** | **string** |  | [readonly] 
**Pk** | **string** |  | [readonly] 
**Component** | **string** |  | [readonly] 
**BindingModel** | **string** |  | [readonly] 
**BindingPk** | **string** |  | [readonly] 
**BindingOrder** | **NullableInt32** |  | [readonly] 

## Methods

### NewDiagramNode

`func NewDiagramNode(identifier string, type_ DiagramNodeTypeEnum, name string, verboseName string, model string, pk string, component string, bindingModel string, bindingPk string, bindingOrder NullableInt32, ) *DiagramNode`

NewDiagramNode instantiates a new DiagramNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagramNodeWithDefaults

`func NewDiagramNodeWithDefaults() *DiagramNode`

NewDiagramNodeWithDefaults instantiates a new DiagramNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *DiagramNode) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *DiagramNode) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *DiagramNode) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetType

`func (o *DiagramNode) GetType() DiagramNodeTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiagramNode) GetTypeOk() (*DiagramNodeTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiagramNode) SetType(v DiagramNodeTypeEnum)`

SetType sets Type field to given value.


### GetName

`func (o *DiagramNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiagramNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiagramNode) SetName(v string)`

SetName sets Name field to given value.


### GetVerboseName

`func (o *DiagramNode) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *DiagramNode) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *DiagramNode) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetModel

`func (o *DiagramNode) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DiagramNode) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DiagramNode) SetModel(v string)`

SetModel sets Model field to given value.


### GetPk

`func (o *DiagramNode) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *DiagramNode) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *DiagramNode) SetPk(v string)`

SetPk sets Pk field to given value.


### GetComponent

`func (o *DiagramNode) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *DiagramNode) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *DiagramNode) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetBindingModel

`func (o *DiagramNode) GetBindingModel() string`

GetBindingModel returns the BindingModel field if non-nil, zero value otherwise.

### GetBindingModelOk

`func (o *DiagramNode) GetBindingModelOk() (*string, bool)`

GetBindingModelOk returns a tuple with the BindingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingModel

`func (o *DiagramNode) SetBindingModel(v string)`

SetBindingModel sets BindingModel field to given value.


### GetBindingPk

`func (o *DiagramNode) GetBindingPk() string`

GetBindingPk returns the BindingPk field if non-nil, zero value otherwise.

### GetBindingPkOk

`func (o *DiagramNode) GetBindingPkOk() (*string, bool)`

GetBindingPkOk returns a tuple with the BindingPk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingPk

`func (o *DiagramNode) SetBindingPk(v string)`

SetBindingPk sets BindingPk field to given value.


### GetBindingOrder

`func (o *DiagramNode) GetBindingOrder() int32`

GetBindingOrder returns the BindingOrder field if non-nil, zero value otherwise.

### GetBindingOrderOk

`func (o *DiagramNode) GetBindingOrderOk() (*int32, bool)`

GetBindingOrderOk returns a tuple with the BindingOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingOrder

`func (o *DiagramNode) SetBindingOrder(v int32)`

SetBindingOrder sets BindingOrder field to given value.


### SetBindingOrderNil

`func (o *DiagramNode) SetBindingOrderNil(b bool)`

 SetBindingOrderNil sets the value for BindingOrder to be an explicit nil

### UnsetBindingOrder
`func (o *DiagramNode) UnsetBindingOrder()`

UnsetBindingOrder ensures that no value is present for BindingOrder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


