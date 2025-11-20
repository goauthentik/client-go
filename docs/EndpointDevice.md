# EndpointDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceUuid** | Pointer to **string** |  | [optional] 
**PbmUuid** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Group** | Pointer to **NullableString** |  | [optional] 
**GroupObj** | [**DeviceGroup**](DeviceGroup.md) |  | 
**Expiring** | Pointer to **bool** |  | [optional] 
**Expires** | Pointer to **NullableTime** |  | [optional] 
**Facts** | [**DeviceFactSnapshot**](DeviceFactSnapshot.md) |  | [readonly] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEndpointDevice

`func NewEndpointDevice(pbmUuid string, name string, groupObj DeviceGroup, facts DeviceFactSnapshot, ) *EndpointDevice`

NewEndpointDevice instantiates a new EndpointDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointDeviceWithDefaults

`func NewEndpointDeviceWithDefaults() *EndpointDevice`

NewEndpointDeviceWithDefaults instantiates a new EndpointDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceUuid

`func (o *EndpointDevice) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *EndpointDevice) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *EndpointDevice) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.

### HasDeviceUuid

`func (o *EndpointDevice) HasDeviceUuid() bool`

HasDeviceUuid returns a boolean if a field has been set.

### GetPbmUuid

`func (o *EndpointDevice) GetPbmUuid() string`

GetPbmUuid returns the PbmUuid field if non-nil, zero value otherwise.

### GetPbmUuidOk

`func (o *EndpointDevice) GetPbmUuidOk() (*string, bool)`

GetPbmUuidOk returns a tuple with the PbmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbmUuid

`func (o *EndpointDevice) SetPbmUuid(v string)`

SetPbmUuid sets PbmUuid field to given value.


### GetName

`func (o *EndpointDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointDevice) SetName(v string)`

SetName sets Name field to given value.


### GetGroup

`func (o *EndpointDevice) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *EndpointDevice) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *EndpointDevice) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *EndpointDevice) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *EndpointDevice) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *EndpointDevice) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetGroupObj

`func (o *EndpointDevice) GetGroupObj() DeviceGroup`

GetGroupObj returns the GroupObj field if non-nil, zero value otherwise.

### GetGroupObjOk

`func (o *EndpointDevice) GetGroupObjOk() (*DeviceGroup, bool)`

GetGroupObjOk returns a tuple with the GroupObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObj

`func (o *EndpointDevice) SetGroupObj(v DeviceGroup)`

SetGroupObj sets GroupObj field to given value.


### GetExpiring

`func (o *EndpointDevice) GetExpiring() bool`

GetExpiring returns the Expiring field if non-nil, zero value otherwise.

### GetExpiringOk

`func (o *EndpointDevice) GetExpiringOk() (*bool, bool)`

GetExpiringOk returns a tuple with the Expiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiring

`func (o *EndpointDevice) SetExpiring(v bool)`

SetExpiring sets Expiring field to given value.

### HasExpiring

`func (o *EndpointDevice) HasExpiring() bool`

HasExpiring returns a boolean if a field has been set.

### GetExpires

`func (o *EndpointDevice) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *EndpointDevice) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *EndpointDevice) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *EndpointDevice) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *EndpointDevice) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *EndpointDevice) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetFacts

`func (o *EndpointDevice) GetFacts() DeviceFactSnapshot`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *EndpointDevice) GetFactsOk() (*DeviceFactSnapshot, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *EndpointDevice) SetFacts(v DeviceFactSnapshot)`

SetFacts sets Facts field to given value.


### GetAttributes

`func (o *EndpointDevice) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EndpointDevice) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EndpointDevice) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *EndpointDevice) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


