# PatchedEndpointDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceUuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **NullableString** |  | [optional] 
**GroupObj** | Pointer to [**DeviceGroupRequest**](DeviceGroupRequest.md) |  | [optional] 
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

### GetGroup

`func (o *PatchedEndpointDeviceRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedEndpointDeviceRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedEndpointDeviceRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedEndpointDeviceRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedEndpointDeviceRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedEndpointDeviceRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetGroupObj

`func (o *PatchedEndpointDeviceRequest) GetGroupObj() DeviceGroupRequest`

GetGroupObj returns the GroupObj field if non-nil, zero value otherwise.

### GetGroupObjOk

`func (o *PatchedEndpointDeviceRequest) GetGroupObjOk() (*DeviceGroupRequest, bool)`

GetGroupObjOk returns a tuple with the GroupObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObj

`func (o *PatchedEndpointDeviceRequest) SetGroupObj(v DeviceGroupRequest)`

SetGroupObj sets GroupObj field to given value.

### HasGroupObj

`func (o *PatchedEndpointDeviceRequest) HasGroupObj() bool`

HasGroupObj returns a boolean if a field has been set.

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


