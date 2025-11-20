# DeviceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | **string** |  | 
**Connector** | **string** |  | 
**ConnectorObj** | [**Connector**](Connector.md) |  | [readonly] 
**LatestSnapshot** | [**NullableDeviceConnectionLatestSnapshot**](DeviceConnectionLatestSnapshot.md) |  | 

## Methods

### NewDeviceConnection

`func NewDeviceConnection(device string, connector string, connectorObj Connector, latestSnapshot NullableDeviceConnectionLatestSnapshot, ) *DeviceConnection`

NewDeviceConnection instantiates a new DeviceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceConnectionWithDefaults

`func NewDeviceConnectionWithDefaults() *DeviceConnection`

NewDeviceConnectionWithDefaults instantiates a new DeviceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *DeviceConnection) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DeviceConnection) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DeviceConnection) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetConnector

`func (o *DeviceConnection) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *DeviceConnection) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *DeviceConnection) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetConnectorObj

`func (o *DeviceConnection) GetConnectorObj() Connector`

GetConnectorObj returns the ConnectorObj field if non-nil, zero value otherwise.

### GetConnectorObjOk

`func (o *DeviceConnection) GetConnectorObjOk() (*Connector, bool)`

GetConnectorObjOk returns a tuple with the ConnectorObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorObj

`func (o *DeviceConnection) SetConnectorObj(v Connector)`

SetConnectorObj sets ConnectorObj field to given value.


### GetLatestSnapshot

`func (o *DeviceConnection) GetLatestSnapshot() DeviceConnectionLatestSnapshot`

GetLatestSnapshot returns the LatestSnapshot field if non-nil, zero value otherwise.

### GetLatestSnapshotOk

`func (o *DeviceConnection) GetLatestSnapshotOk() (*DeviceConnectionLatestSnapshot, bool)`

GetLatestSnapshotOk returns a tuple with the LatestSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSnapshot

`func (o *DeviceConnection) SetLatestSnapshot(v DeviceConnectionLatestSnapshot)`

SetLatestSnapshot sets LatestSnapshot field to given value.


### SetLatestSnapshotNil

`func (o *DeviceConnection) SetLatestSnapshotNil(b bool)`

 SetLatestSnapshotNil sets the value for LatestSnapshot to be an explicit nil

### UnsetLatestSnapshot
`func (o *DeviceConnection) UnsetLatestSnapshot()`

UnsetLatestSnapshot ensures that no value is present for LatestSnapshot, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


