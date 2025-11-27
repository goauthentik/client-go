# EndpointDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceUuid** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**AccessGroup** | Pointer to **NullableString** |  | [optional] 
**AccessGroupObj** | [**DeviceAccessGroupRequest**](DeviceAccessGroupRequest.md) |  | 
**Expiring** | Pointer to **bool** |  | [optional] 
**Expires** | Pointer to **NullableTime** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEndpointDeviceRequest

`func NewEndpointDeviceRequest(name string, accessGroupObj DeviceAccessGroupRequest, ) *EndpointDeviceRequest`

NewEndpointDeviceRequest instantiates a new EndpointDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointDeviceRequestWithDefaults

`func NewEndpointDeviceRequestWithDefaults() *EndpointDeviceRequest`

NewEndpointDeviceRequestWithDefaults instantiates a new EndpointDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceUuid

`func (o *EndpointDeviceRequest) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *EndpointDeviceRequest) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *EndpointDeviceRequest) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.

### HasDeviceUuid

`func (o *EndpointDeviceRequest) HasDeviceUuid() bool`

HasDeviceUuid returns a boolean if a field has been set.

### GetName

`func (o *EndpointDeviceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointDeviceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointDeviceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAccessGroup

`func (o *EndpointDeviceRequest) GetAccessGroup() string`

GetAccessGroup returns the AccessGroup field if non-nil, zero value otherwise.

### GetAccessGroupOk

`func (o *EndpointDeviceRequest) GetAccessGroupOk() (*string, bool)`

GetAccessGroupOk returns a tuple with the AccessGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroup

`func (o *EndpointDeviceRequest) SetAccessGroup(v string)`

SetAccessGroup sets AccessGroup field to given value.

### HasAccessGroup

`func (o *EndpointDeviceRequest) HasAccessGroup() bool`

HasAccessGroup returns a boolean if a field has been set.

### SetAccessGroupNil

`func (o *EndpointDeviceRequest) SetAccessGroupNil(b bool)`

 SetAccessGroupNil sets the value for AccessGroup to be an explicit nil

### UnsetAccessGroup
`func (o *EndpointDeviceRequest) UnsetAccessGroup()`

UnsetAccessGroup ensures that no value is present for AccessGroup, not even an explicit nil
### GetAccessGroupObj

`func (o *EndpointDeviceRequest) GetAccessGroupObj() DeviceAccessGroupRequest`

GetAccessGroupObj returns the AccessGroupObj field if non-nil, zero value otherwise.

### GetAccessGroupObjOk

`func (o *EndpointDeviceRequest) GetAccessGroupObjOk() (*DeviceAccessGroupRequest, bool)`

GetAccessGroupObjOk returns a tuple with the AccessGroupObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroupObj

`func (o *EndpointDeviceRequest) SetAccessGroupObj(v DeviceAccessGroupRequest)`

SetAccessGroupObj sets AccessGroupObj field to given value.


### GetExpiring

`func (o *EndpointDeviceRequest) GetExpiring() bool`

GetExpiring returns the Expiring field if non-nil, zero value otherwise.

### GetExpiringOk

`func (o *EndpointDeviceRequest) GetExpiringOk() (*bool, bool)`

GetExpiringOk returns a tuple with the Expiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiring

`func (o *EndpointDeviceRequest) SetExpiring(v bool)`

SetExpiring sets Expiring field to given value.

### HasExpiring

`func (o *EndpointDeviceRequest) HasExpiring() bool`

HasExpiring returns a boolean if a field has been set.

### GetExpires

`func (o *EndpointDeviceRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *EndpointDeviceRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *EndpointDeviceRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *EndpointDeviceRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *EndpointDeviceRequest) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *EndpointDeviceRequest) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetAttributes

`func (o *EndpointDeviceRequest) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EndpointDeviceRequest) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EndpointDeviceRequest) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *EndpointDeviceRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


