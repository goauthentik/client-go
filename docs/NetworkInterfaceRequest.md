# NetworkInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**HardwareAddress** | **string** |  | 
**IpAddresses** | Pointer to **[]string** |  | [optional] 
**DnsServers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNetworkInterfaceRequest

`func NewNetworkInterfaceRequest(name string, hardwareAddress string, ) *NetworkInterfaceRequest`

NewNetworkInterfaceRequest instantiates a new NetworkInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceRequestWithDefaults

`func NewNetworkInterfaceRequestWithDefaults() *NetworkInterfaceRequest`

NewNetworkInterfaceRequestWithDefaults instantiates a new NetworkInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkInterfaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkInterfaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkInterfaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetHardwareAddress

`func (o *NetworkInterfaceRequest) GetHardwareAddress() string`

GetHardwareAddress returns the HardwareAddress field if non-nil, zero value otherwise.

### GetHardwareAddressOk

`func (o *NetworkInterfaceRequest) GetHardwareAddressOk() (*string, bool)`

GetHardwareAddressOk returns a tuple with the HardwareAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareAddress

`func (o *NetworkInterfaceRequest) SetHardwareAddress(v string)`

SetHardwareAddress sets HardwareAddress field to given value.


### GetIpAddresses

`func (o *NetworkInterfaceRequest) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *NetworkInterfaceRequest) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *NetworkInterfaceRequest) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *NetworkInterfaceRequest) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetDnsServers

`func (o *NetworkInterfaceRequest) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *NetworkInterfaceRequest) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *NetworkInterfaceRequest) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *NetworkInterfaceRequest) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


