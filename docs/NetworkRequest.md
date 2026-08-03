# NetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** |  | 
**FirewallEnabled** | Pointer to **bool** |  | [optional] 
**Interfaces** | [**[]NetworkInterfaceRequest**](NetworkInterfaceRequest.md) |  | 
**Gateway** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkRequest

`func NewNetworkRequest(hostname string, interfaces []NetworkInterfaceRequest, ) *NetworkRequest`

NewNetworkRequest instantiates a new NetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRequestWithDefaults

`func NewNetworkRequestWithDefaults() *NetworkRequest`

NewNetworkRequestWithDefaults instantiates a new NetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *NetworkRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NetworkRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NetworkRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetFirewallEnabled

`func (o *NetworkRequest) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *NetworkRequest) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *NetworkRequest) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *NetworkRequest) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### GetInterfaces

`func (o *NetworkRequest) GetInterfaces() []NetworkInterfaceRequest`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *NetworkRequest) GetInterfacesOk() (*[]NetworkInterfaceRequest, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *NetworkRequest) SetInterfaces(v []NetworkInterfaceRequest)`

SetInterfaces sets Interfaces field to given value.


### GetGateway

`func (o *NetworkRequest) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *NetworkRequest) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *NetworkRequest) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *NetworkRequest) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


