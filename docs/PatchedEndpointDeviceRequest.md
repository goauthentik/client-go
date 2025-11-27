# PatchedEndpointDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceUuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AccessGroup** | Pointer to **NullableString** |  | [optional] 
**AccessGroupObj** | Pointer to [**DeviceAccessGroupRequest**](DeviceAccessGroupRequest.md) |  | [optional] 
**Expiring** | Pointer to **bool** |  | [optional] 
**Expires** | Pointer to **NullableTime** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedEndpointDeviceRequest

`func NewPatchedEndpointDeviceRequest() *PatchedEndpointDeviceRequest`

NewPatchedEndpointDeviceRequest instantiates a new PatchedEndpointDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEndpointDeviceRequestWithDefaults

`func NewPatchedEndpointDeviceRequestWithDefaults() *PatchedEndpointDeviceRequest`

NewPatchedEndpointDeviceRequestWithDefaults instantiates a new PatchedEndpointDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceUuid

`func (o *PatchedEndpointDeviceRequest) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *PatchedEndpointDeviceRequest) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *PatchedEndpointDeviceRequest) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.

### HasDeviceUuid

`func (o *PatchedEndpointDeviceRequest) HasDeviceUuid() bool`

HasDeviceUuid returns a boolean if a field has been set.

### GetName

`func (o *PatchedEndpointDeviceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEndpointDeviceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEndpointDeviceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEndpointDeviceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccessGroup

`func (o *PatchedEndpointDeviceRequest) GetAccessGroup() string`

GetAccessGroup returns the AccessGroup field if non-nil, zero value otherwise.

### GetAccessGroupOk

`func (o *PatchedEndpointDeviceRequest) GetAccessGroupOk() (*string, bool)`

GetAccessGroupOk returns a tuple with the AccessGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroup

`func (o *PatchedEndpointDeviceRequest) SetAccessGroup(v string)`

SetAccessGroup sets AccessGroup field to given value.

### HasAccessGroup

`func (o *PatchedEndpointDeviceRequest) HasAccessGroup() bool`

HasAccessGroup returns a boolean if a field has been set.

### SetAccessGroupNil

`func (o *PatchedEndpointDeviceRequest) SetAccessGroupNil(b bool)`

 SetAccessGroupNil sets the value for AccessGroup to be an explicit nil

### UnsetAccessGroup
`func (o *PatchedEndpointDeviceRequest) UnsetAccessGroup()`

UnsetAccessGroup ensures that no value is present for AccessGroup, not even an explicit nil
### GetAccessGroupObj

`func (o *PatchedEndpointDeviceRequest) GetAccessGroupObj() DeviceAccessGroupRequest`

GetAccessGroupObj returns the AccessGroupObj field if non-nil, zero value otherwise.

### GetAccessGroupObjOk

`func (o *PatchedEndpointDeviceRequest) GetAccessGroupObjOk() (*DeviceAccessGroupRequest, bool)`

GetAccessGroupObjOk returns a tuple with the AccessGroupObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroupObj

`func (o *PatchedEndpointDeviceRequest) SetAccessGroupObj(v DeviceAccessGroupRequest)`

SetAccessGroupObj sets AccessGroupObj field to given value.

### HasAccessGroupObj

`func (o *PatchedEndpointDeviceRequest) HasAccessGroupObj() bool`

HasAccessGroupObj returns a boolean if a field has been set.

### GetExpiring

`func (o *PatchedEndpointDeviceRequest) GetExpiring() bool`

GetExpiring returns the Expiring field if non-nil, zero value otherwise.

### GetExpiringOk

`func (o *PatchedEndpointDeviceRequest) GetExpiringOk() (*bool, bool)`

GetExpiringOk returns a tuple with the Expiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiring

`func (o *PatchedEndpointDeviceRequest) SetExpiring(v bool)`

SetExpiring sets Expiring field to given value.

### HasExpiring

`func (o *PatchedEndpointDeviceRequest) HasExpiring() bool`

HasExpiring returns a boolean if a field has been set.

### GetExpires

`func (o *PatchedEndpointDeviceRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *PatchedEndpointDeviceRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *PatchedEndpointDeviceRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *PatchedEndpointDeviceRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *PatchedEndpointDeviceRequest) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *PatchedEndpointDeviceRequest) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetAttributes

`func (o *PatchedEndpointDeviceRequest) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PatchedEndpointDeviceRequest) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PatchedEndpointDeviceRequest) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PatchedEndpointDeviceRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


