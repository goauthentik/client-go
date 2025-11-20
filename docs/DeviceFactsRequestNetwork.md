# DeviceFactsRequestNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** |  | 
**FirewallEnabled** | Pointer to **bool** |  | [optional] 
**Interfaces** | [**[]NetworkInterfaceRequest**](NetworkInterfaceRequest.md) |  | 
**Gateway** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceFactsRequestNetwork

`func NewDeviceFactsRequestNetwork(hostname string, interfaces []NetworkInterfaceRequest, ) *DeviceFactsRequestNetwork`

NewDeviceFactsRequestNetwork instantiates a new DeviceFactsRequestNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceFactsRequestNetworkWithDefaults

`func NewDeviceFactsRequestNetworkWithDefaults() *DeviceFactsRequestNetwork`

NewDeviceFactsRequestNetworkWithDefaults instantiates a new DeviceFactsRequestNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *DeviceFactsRequestNetwork) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DeviceFactsRequestNetwork) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DeviceFactsRequestNetwork) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetFirewallEnabled

`func (o *DeviceFactsRequestNetwork) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *DeviceFactsRequestNetwork) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *DeviceFactsRequestNetwork) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *DeviceFactsRequestNetwork) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### GetInterfaces

`func (o *DeviceFactsRequestNetwork) GetInterfaces() []NetworkInterfaceRequest`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *DeviceFactsRequestNetwork) GetInterfacesOk() (*[]NetworkInterfaceRequest, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *DeviceFactsRequestNetwork) SetInterfaces(v []NetworkInterfaceRequest)`

SetInterfaces sets Interfaces field to given value.


### GetGateway

`func (o *DeviceFactsRequestNetwork) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *DeviceFactsRequestNetwork) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *DeviceFactsRequestNetwork) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *DeviceFactsRequestNetwork) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


