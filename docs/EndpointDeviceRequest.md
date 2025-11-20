# EndpointDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceUuid** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Group** | Pointer to **NullableString** |  | [optional] 
**GroupObj** | [**DeviceGroupRequest**](DeviceGroupRequest.md) |  | 
**Expiring** | Pointer to **bool** |  | [optional] 
**Expires** | Pointer to **NullableTime** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEndpointDeviceRequest

`func NewEndpointDeviceRequest(name string, groupObj DeviceGroupRequest, ) *EndpointDeviceRequest`

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


### GetGroup

`func (o *EndpointDeviceRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *EndpointDeviceRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *EndpointDeviceRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *EndpointDeviceRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *EndpointDeviceRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *EndpointDeviceRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetGroupObj

`func (o *EndpointDeviceRequest) GetGroupObj() DeviceGroupRequest`

GetGroupObj returns the GroupObj field if non-nil, zero value otherwise.

### GetGroupObjOk

`func (o *EndpointDeviceRequest) GetGroupObjOk() (*DeviceGroupRequest, bool)`

GetGroupObjOk returns a tuple with the GroupObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObj

`func (o *EndpointDeviceRequest) SetGroupObj(v DeviceGroupRequest)`

SetGroupObj sets GroupObj field to given value.


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


