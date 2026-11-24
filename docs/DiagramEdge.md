# DiagramEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origin** | **string** |  | [readonly] 
**Target** | **string** |  | [readonly] 
**Type** | [**DiagramEdgeTypeEnum**](DiagramEdgeTypeEnum.md) |  | [readonly] 

## Methods

### NewDiagramEdge

`func NewDiagramEdge(origin string, target string, type_ DiagramEdgeTypeEnum, ) *DiagramEdge`

NewDiagramEdge instantiates a new DiagramEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagramEdgeWithDefaults

`func NewDiagramEdgeWithDefaults() *DiagramEdge`

NewDiagramEdgeWithDefaults instantiates a new DiagramEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigin

`func (o *DiagramEdge) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *DiagramEdge) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *DiagramEdge) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetTarget

`func (o *DiagramEdge) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DiagramEdge) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DiagramEdge) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetType

`func (o *DiagramEdge) GetType() DiagramEdgeTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiagramEdge) GetTypeOk() (*DiagramEdgeTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiagramEdge) SetType(v DiagramEdgeTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


