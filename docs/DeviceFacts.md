# DeviceFacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Os** | Pointer to [**NullableDeviceFactsOs**](DeviceFactsOs.md) |  | [optional] 
**Disks** | Pointer to [**[]Disk**](Disk.md) |  | [optional] 
**Network** | Pointer to [**NullableDeviceFactsNetwork**](DeviceFactsNetwork.md) |  | [optional] 
**Hardware** | Pointer to [**NullableDeviceFactsHardware**](DeviceFactsHardware.md) |  | [optional] 
**Software** | Pointer to [**[]Software**](Software.md) |  | [optional] 
**Processes** | Pointer to [**[]Process**](Process.md) |  | [optional] 
**Users** | Pointer to [**[]DeviceUser**](DeviceUser.md) |  | [optional] 
**Groups** | Pointer to [**[]DeviceGroup**](DeviceGroup.md) |  | [optional] 
**Vendor** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDeviceFacts

`func NewDeviceFacts() *DeviceFacts`

NewDeviceFacts instantiates a new DeviceFacts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceFactsWithDefaults

`func NewDeviceFactsWithDefaults() *DeviceFacts`

NewDeviceFactsWithDefaults instantiates a new DeviceFacts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOs

`func (o *DeviceFacts) GetOs() DeviceFactsOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *DeviceFacts) GetOsOk() (*DeviceFactsOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *DeviceFacts) SetOs(v DeviceFactsOs)`

SetOs sets Os field to given value.

### HasOs

`func (o *DeviceFacts) HasOs() bool`

HasOs returns a boolean if a field has been set.

### SetOsNil

`func (o *DeviceFacts) SetOsNil(b bool)`

 SetOsNil sets the value for Os to be an explicit nil

### UnsetOs
`func (o *DeviceFacts) UnsetOs()`

UnsetOs ensures that no value is present for Os, not even an explicit nil
### GetDisks

`func (o *DeviceFacts) GetDisks() []Disk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *DeviceFacts) GetDisksOk() (*[]Disk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *DeviceFacts) SetDisks(v []Disk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *DeviceFacts) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### SetDisksNil

`func (o *DeviceFacts) SetDisksNil(b bool)`

 SetDisksNil sets the value for Disks to be an explicit nil

### UnsetDisks
`func (o *DeviceFacts) UnsetDisks()`

UnsetDisks ensures that no value is present for Disks, not even an explicit nil
### GetNetwork

`func (o *DeviceFacts) GetNetwork() DeviceFactsNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *DeviceFacts) GetNetworkOk() (*DeviceFactsNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *DeviceFacts) SetNetwork(v DeviceFactsNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *DeviceFacts) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *DeviceFacts) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *DeviceFacts) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetHardware

`func (o *DeviceFacts) GetHardware() DeviceFactsHardware`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *DeviceFacts) GetHardwareOk() (*DeviceFactsHardware, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *DeviceFacts) SetHardware(v DeviceFactsHardware)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *DeviceFacts) HasHardware() bool`

HasHardware returns a boolean if a field has been set.

### SetHardwareNil

`func (o *DeviceFacts) SetHardwareNil(b bool)`

 SetHardwareNil sets the value for Hardware to be an explicit nil

### UnsetHardware
`func (o *DeviceFacts) UnsetHardware()`

UnsetHardware ensures that no value is present for Hardware, not even an explicit nil
### GetSoftware

`func (o *DeviceFacts) GetSoftware() []Software`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *DeviceFacts) GetSoftwareOk() (*[]Software, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *DeviceFacts) SetSoftware(v []Software)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *DeviceFacts) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### SetSoftwareNil

`func (o *DeviceFacts) SetSoftwareNil(b bool)`

 SetSoftwareNil sets the value for Software to be an explicit nil

### UnsetSoftware
`func (o *DeviceFacts) UnsetSoftware()`

UnsetSoftware ensures that no value is present for Software, not even an explicit nil
### GetProcesses

`func (o *DeviceFacts) GetProcesses() []Process`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *DeviceFacts) GetProcessesOk() (*[]Process, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *DeviceFacts) SetProcesses(v []Process)`

SetProcesses sets Processes field to given value.

### HasProcesses

`func (o *DeviceFacts) HasProcesses() bool`

HasProcesses returns a boolean if a field has been set.

### SetProcessesNil

`func (o *DeviceFacts) SetProcessesNil(b bool)`

 SetProcessesNil sets the value for Processes to be an explicit nil

### UnsetProcesses
`func (o *DeviceFacts) UnsetProcesses()`

UnsetProcesses ensures that no value is present for Processes, not even an explicit nil
### GetUsers

`func (o *DeviceFacts) GetUsers() []DeviceUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *DeviceFacts) GetUsersOk() (*[]DeviceUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *DeviceFacts) SetUsers(v []DeviceUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *DeviceFacts) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *DeviceFacts) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *DeviceFacts) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetGroups

`func (o *DeviceFacts) GetGroups() []DeviceGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *DeviceFacts) GetGroupsOk() (*[]DeviceGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *DeviceFacts) SetGroups(v []DeviceGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *DeviceFacts) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *DeviceFacts) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *DeviceFacts) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetVendor

`func (o *DeviceFacts) GetVendor() map[string]interface{}`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *DeviceFacts) GetVendorOk() (*map[string]interface{}, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *DeviceFacts) SetVendor(v map[string]interface{})`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *DeviceFacts) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


