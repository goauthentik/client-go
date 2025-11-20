# DeviceFactsNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** |  | 
**FirewallEnabled** | Pointer to **bool** |  | [optional] 
**Interfaces** | [**[]NetworkInterface**](NetworkInterface.md) |  | 
**Gateway** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceFactsNetwork

`func NewDeviceFactsNetwork(hostname string, interfaces []NetworkInterface, ) *DeviceFactsNetwork`

NewDeviceFactsNetwork instantiates a new DeviceFactsNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceFactsNetworkWithDefaults

`func NewDeviceFactsNetworkWithDefaults() *DeviceFactsNetwork`

NewDeviceFactsNetworkWithDefaults instantiates a new DeviceFactsNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *DeviceFactsNetwork) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DeviceFactsNetwork) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DeviceFactsNetwork) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetFirewallEnabled

`func (o *DeviceFactsNetwork) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *DeviceFactsNetwork) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *DeviceFactsNetwork) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *DeviceFactsNetwork) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### GetInterfaces

`func (o *DeviceFactsNetwork) GetInterfaces() []NetworkInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *DeviceFactsNetwork) GetInterfacesOk() (*[]NetworkInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *DeviceFactsNetwork) SetInterfaces(v []NetworkInterface)`

SetInterfaces sets Interfaces field to given value.


### GetGateway

`func (o *DeviceFactsNetwork) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *DeviceFactsNetwork) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *DeviceFactsNetwork) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *DeviceFactsNetwork) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


