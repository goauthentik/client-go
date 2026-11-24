# FlowDiagram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | [**[]DiagramNode**](DiagramNode.md) |  | [readonly] 
**Edges** | [**[]DiagramEdge**](DiagramEdge.md) |  | [readonly] 

## Methods

### NewFlowDiagram

`func NewFlowDiagram(nodes []DiagramNode, edges []DiagramEdge, ) *FlowDiagram`

NewFlowDiagram instantiates a new FlowDiagram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowDiagramWithDefaults

`func NewFlowDiagramWithDefaults() *FlowDiagram`

NewFlowDiagramWithDefaults instantiates a new FlowDiagram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *FlowDiagram) GetNodes() []DiagramNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *FlowDiagram) GetNodesOk() (*[]DiagramNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *FlowDiagram) SetNodes(v []DiagramNode)`

SetNodes sets Nodes field to given value.


### GetEdges

`func (o *FlowDiagram) GetEdges() []DiagramEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *FlowDiagram) GetEdgesOk() (*[]DiagramEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *FlowDiagram) SetEdges(v []DiagramEdge)`

SetEdges sets Edges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


