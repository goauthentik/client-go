# EndpointDeviceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceUuid** | Pointer to **string** |  | [optional] 
**PbmUuid** | **string** |  | [readonly] 
**Name** | **string** |  | 
**AccessGroup** | Pointer to **NullableString** |  | [optional] 
**AccessGroupObj** | Pointer to [**DeviceAccessGroup**](DeviceAccessGroup.md) |  | [optional] 
**Expiring** | Pointer to **bool** |  | [optional] 
**Expires** | Pointer to **NullableTime** |  | [optional] 
**Facts** | [**DeviceFactSnapshot**](DeviceFactSnapshot.md) |  | [readonly] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**ConnectionsObj** | [**[]DeviceConnection**](DeviceConnection.md) |  | 
**Policies** | **[]string** |  | [readonly] 
**Connections** | **[]string** |  | [readonly] 

## Methods

### NewEndpointDeviceDetails

`func NewEndpointDeviceDetails(pbmUuid string, name string, facts DeviceFactSnapshot, connectionsObj []DeviceConnection, policies []string, connections []string, ) *EndpointDeviceDetails`

NewEndpointDeviceDetails instantiates a new EndpointDeviceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointDeviceDetailsWithDefaults

`func NewEndpointDeviceDetailsWithDefaults() *EndpointDeviceDetails`

NewEndpointDeviceDetailsWithDefaults instantiates a new EndpointDeviceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceUuid

`func (o *EndpointDeviceDetails) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *EndpointDeviceDetails) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *EndpointDeviceDetails) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.

### HasDeviceUuid

`func (o *EndpointDeviceDetails) HasDeviceUuid() bool`

HasDeviceUuid returns a boolean if a field has been set.

### GetPbmUuid

`func (o *EndpointDeviceDetails) GetPbmUuid() string`

GetPbmUuid returns the PbmUuid field if non-nil, zero value otherwise.

### GetPbmUuidOk

`func (o *EndpointDeviceDetails) GetPbmUuidOk() (*string, bool)`

GetPbmUuidOk returns a tuple with the PbmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbmUuid

`func (o *EndpointDeviceDetails) SetPbmUuid(v string)`

SetPbmUuid sets PbmUuid field to given value.


### GetName

`func (o *EndpointDeviceDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointDeviceDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointDeviceDetails) SetName(v string)`

SetName sets Name field to given value.


### GetAccessGroup

`func (o *EndpointDeviceDetails) GetAccessGroup() string`

GetAccessGroup returns the AccessGroup field if non-nil, zero value otherwise.

### GetAccessGroupOk

`func (o *EndpointDeviceDetails) GetAccessGroupOk() (*string, bool)`

GetAccessGroupOk returns a tuple with the AccessGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroup

`func (o *EndpointDeviceDetails) SetAccessGroup(v string)`

SetAccessGroup sets AccessGroup field to given value.

### HasAccessGroup

`func (o *EndpointDeviceDetails) HasAccessGroup() bool`

HasAccessGroup returns a boolean if a field has been set.

### SetAccessGroupNil

`func (o *EndpointDeviceDetails) SetAccessGroupNil(b bool)`

 SetAccessGroupNil sets the value for AccessGroup to be an explicit nil

### UnsetAccessGroup
`func (o *EndpointDeviceDetails) UnsetAccessGroup()`

UnsetAccessGroup ensures that no value is present for AccessGroup, not even an explicit nil
### GetAccessGroupObj

`func (o *EndpointDeviceDetails) GetAccessGroupObj() DeviceAccessGroup`

GetAccessGroupObj returns the AccessGroupObj field if non-nil, zero value otherwise.

### GetAccessGroupObjOk

`func (o *EndpointDeviceDetails) GetAccessGroupObjOk() (*DeviceAccessGroup, bool)`

GetAccessGroupObjOk returns a tuple with the AccessGroupObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroupObj

`func (o *EndpointDeviceDetails) SetAccessGroupObj(v DeviceAccessGroup)`

SetAccessGroupObj sets AccessGroupObj field to given value.

### HasAccessGroupObj

`func (o *EndpointDeviceDetails) HasAccessGroupObj() bool`

HasAccessGroupObj returns a boolean if a field has been set.

### GetExpiring

`func (o *EndpointDeviceDetails) GetExpiring() bool`

GetExpiring returns the Expiring field if non-nil, zero value otherwise.

### GetExpiringOk

`func (o *EndpointDeviceDetails) GetExpiringOk() (*bool, bool)`

GetExpiringOk returns a tuple with the Expiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiring

`func (o *EndpointDeviceDetails) SetExpiring(v bool)`

SetExpiring sets Expiring field to given value.

### HasExpiring

`func (o *EndpointDeviceDetails) HasExpiring() bool`

HasExpiring returns a boolean if a field has been set.

### GetExpires

`func (o *EndpointDeviceDetails) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *EndpointDeviceDetails) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *EndpointDeviceDetails) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *EndpointDeviceDetails) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *EndpointDeviceDetails) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *EndpointDeviceDetails) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetFacts

`func (o *EndpointDeviceDetails) GetFacts() DeviceFactSnapshot`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *EndpointDeviceDetails) GetFactsOk() (*DeviceFactSnapshot, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *EndpointDeviceDetails) SetFacts(v DeviceFactSnapshot)`

SetFacts sets Facts field to given value.


### GetAttributes

`func (o *EndpointDeviceDetails) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EndpointDeviceDetails) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EndpointDeviceDetails) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *EndpointDeviceDetails) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetConnectionsObj

`func (o *EndpointDeviceDetails) GetConnectionsObj() []DeviceConnection`

GetConnectionsObj returns the ConnectionsObj field if non-nil, zero value otherwise.

### GetConnectionsObjOk

`func (o *EndpointDeviceDetails) GetConnectionsObjOk() (*[]DeviceConnection, bool)`

GetConnectionsObjOk returns a tuple with the ConnectionsObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionsObj

`func (o *EndpointDeviceDetails) SetConnectionsObj(v []DeviceConnection)`

SetConnectionsObj sets ConnectionsObj field to given value.


### GetPolicies

`func (o *EndpointDeviceDetails) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *EndpointDeviceDetails) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *EndpointDeviceDetails) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### GetConnections

`func (o *EndpointDeviceDetails) GetConnections() []string`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *EndpointDeviceDetails) GetConnectionsOk() (*[]string, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *EndpointDeviceDetails) SetConnections(v []string)`

SetConnections sets Connections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


