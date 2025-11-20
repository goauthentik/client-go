# DeviceFactsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Os** | Pointer to [**NullableDeviceFactsRequestOs**](DeviceFactsRequestOs.md) |  | [optional] 
**Disks** | Pointer to [**[]DiskRequest**](DiskRequest.md) |  | [optional] 
**Network** | Pointer to [**NullableDeviceFactsRequestNetwork**](DeviceFactsRequestNetwork.md) |  | [optional] 
**Hardware** | Pointer to [**NullableDeviceFactsRequestHardware**](DeviceFactsRequestHardware.md) |  | [optional] 
**Software** | Pointer to [**[]SoftwareRequest**](SoftwareRequest.md) |  | [optional] 
**Processes** | Pointer to [**[]ProcessRequest**](ProcessRequest.md) |  | [optional] 
**Vendor** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDeviceFactsRequest

`func NewDeviceFactsRequest() *DeviceFactsRequest`

NewDeviceFactsRequest instantiates a new DeviceFactsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceFactsRequestWithDefaults

`func NewDeviceFactsRequestWithDefaults() *DeviceFactsRequest`

NewDeviceFactsRequestWithDefaults instantiates a new DeviceFactsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOs

`func (o *DeviceFactsRequest) GetOs() DeviceFactsRequestOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *DeviceFactsRequest) GetOsOk() (*DeviceFactsRequestOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *DeviceFactsRequest) SetOs(v DeviceFactsRequestOs)`

SetOs sets Os field to given value.

### HasOs

`func (o *DeviceFactsRequest) HasOs() bool`

HasOs returns a boolean if a field has been set.

### SetOsNil

`func (o *DeviceFactsRequest) SetOsNil(b bool)`

 SetOsNil sets the value for Os to be an explicit nil

### UnsetOs
`func (o *DeviceFactsRequest) UnsetOs()`

UnsetOs ensures that no value is present for Os, not even an explicit nil
### GetDisks

`func (o *DeviceFactsRequest) GetDisks() []DiskRequest`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *DeviceFactsRequest) GetDisksOk() (*[]DiskRequest, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *DeviceFactsRequest) SetDisks(v []DiskRequest)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *DeviceFactsRequest) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### SetDisksNil

`func (o *DeviceFactsRequest) SetDisksNil(b bool)`

 SetDisksNil sets the value for Disks to be an explicit nil

### UnsetDisks
`func (o *DeviceFactsRequest) UnsetDisks()`

UnsetDisks ensures that no value is present for Disks, not even an explicit nil
### GetNetwork

`func (o *DeviceFactsRequest) GetNetwork() DeviceFactsRequestNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *DeviceFactsRequest) GetNetworkOk() (*DeviceFactsRequestNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *DeviceFactsRequest) SetNetwork(v DeviceFactsRequestNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *DeviceFactsRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *DeviceFactsRequest) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *DeviceFactsRequest) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetHardware

`func (o *DeviceFactsRequest) GetHardware() DeviceFactsRequestHardware`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *DeviceFactsRequest) GetHardwareOk() (*DeviceFactsRequestHardware, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *DeviceFactsRequest) SetHardware(v DeviceFactsRequestHardware)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *DeviceFactsRequest) HasHardware() bool`

HasHardware returns a boolean if a field has been set.

### SetHardwareNil

`func (o *DeviceFactsRequest) SetHardwareNil(b bool)`

 SetHardwareNil sets the value for Hardware to be an explicit nil

### UnsetHardware
`func (o *DeviceFactsRequest) UnsetHardware()`

UnsetHardware ensures that no value is present for Hardware, not even an explicit nil
### GetSoftware

`func (o *DeviceFactsRequest) GetSoftware() []SoftwareRequest`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *DeviceFactsRequest) GetSoftwareOk() (*[]SoftwareRequest, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *DeviceFactsRequest) SetSoftware(v []SoftwareRequest)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *DeviceFactsRequest) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### SetSoftwareNil

`func (o *DeviceFactsRequest) SetSoftwareNil(b bool)`

 SetSoftwareNil sets the value for Software to be an explicit nil

### UnsetSoftware
`func (o *DeviceFactsRequest) UnsetSoftware()`

UnsetSoftware ensures that no value is present for Software, not even an explicit nil
### GetProcesses

`func (o *DeviceFactsRequest) GetProcesses() []ProcessRequest`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *DeviceFactsRequest) GetProcessesOk() (*[]ProcessRequest, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *DeviceFactsRequest) SetProcesses(v []ProcessRequest)`

SetProcesses sets Processes field to given value.

### HasProcesses

`func (o *DeviceFactsRequest) HasProcesses() bool`

HasProcesses returns a boolean if a field has been set.

### SetProcessesNil

`func (o *DeviceFactsRequest) SetProcessesNil(b bool)`

 SetProcessesNil sets the value for Processes to be an explicit nil

### UnsetProcesses
`func (o *DeviceFactsRequest) UnsetProcesses()`

UnsetProcesses ensures that no value is present for Processes, not even an explicit nil
### GetVendor

`func (o *DeviceFactsRequest) GetVendor() map[string]interface{}`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *DeviceFactsRequest) GetVendorOk() (*map[string]interface{}, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *DeviceFactsRequest) SetVendor(v map[string]interface{})`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *DeviceFactsRequest) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


