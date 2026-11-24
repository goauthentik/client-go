# NetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**HardwareAddress** | **string** |  | 
**IpAddresses** | Pointer to **[]string** |  | [optional] 
**DnsServers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNetworkInterface

`func NewNetworkInterface(name string, hardwareAddress string, ) *NetworkInterface`

NewNetworkInterface instantiates a new NetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceWithDefaults

`func NewNetworkInterfaceWithDefaults() *NetworkInterface`

NewNetworkInterfaceWithDefaults instantiates a new NetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkInterface) SetName(v string)`

SetName sets Name field to given value.


### GetHardwareAddress

`func (o *NetworkInterface) GetHardwareAddress() string`

GetHardwareAddress returns the HardwareAddress field if non-nil, zero value otherwise.

### GetHardwareAddressOk

`func (o *NetworkInterface) GetHardwareAddressOk() (*string, bool)`

GetHardwareAddressOk returns a tuple with the HardwareAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareAddress

`func (o *NetworkInterface) SetHardwareAddress(v string)`

SetHardwareAddress sets HardwareAddress field to given value.


### GetIpAddresses

`func (o *NetworkInterface) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *NetworkInterface) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *NetworkInterface) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *NetworkInterface) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetDnsServers

`func (o *NetworkInterface) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *NetworkInterface) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *NetworkInterface) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *NetworkInterface) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


