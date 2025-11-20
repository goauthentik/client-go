# PatchedEnrollmentTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceGroup** | Pointer to **NullableString** |  | [optional] 
**Connector** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Expiring** | Pointer to **bool** |  | [optional] 
**Expires** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewPatchedEnrollmentTokenRequest

`func NewPatchedEnrollmentTokenRequest() *PatchedEnrollmentTokenRequest`

NewPatchedEnrollmentTokenRequest instantiates a new PatchedEnrollmentTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEnrollmentTokenRequestWithDefaults

`func NewPatchedEnrollmentTokenRequestWithDefaults() *PatchedEnrollmentTokenRequest`

NewPatchedEnrollmentTokenRequestWithDefaults instantiates a new PatchedEnrollmentTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceGroup

`func (o *PatchedEnrollmentTokenRequest) GetDeviceGroup() string`

GetDeviceGroup returns the DeviceGroup field if non-nil, zero value otherwise.

### GetDeviceGroupOk

`func (o *PatchedEnrollmentTokenRequest) GetDeviceGroupOk() (*string, bool)`

GetDeviceGroupOk returns a tuple with the DeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroup

`func (o *PatchedEnrollmentTokenRequest) SetDeviceGroup(v string)`

SetDeviceGroup sets DeviceGroup field to given value.

### HasDeviceGroup

`func (o *PatchedEnrollmentTokenRequest) HasDeviceGroup() bool`

HasDeviceGroup returns a boolean if a field has been set.

### SetDeviceGroupNil

`func (o *PatchedEnrollmentTokenRequest) SetDeviceGroupNil(b bool)`

 SetDeviceGroupNil sets the value for DeviceGroup to be an explicit nil

### UnsetDeviceGroup
`func (o *PatchedEnrollmentTokenRequest) UnsetDeviceGroup()`

UnsetDeviceGroup ensures that no value is present for DeviceGroup, not even an explicit nil
### GetConnector

`func (o *PatchedEnrollmentTokenRequest) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *PatchedEnrollmentTokenRequest) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *PatchedEnrollmentTokenRequest) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *PatchedEnrollmentTokenRequest) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetName

`func (o *PatchedEnrollmentTokenRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEnrollmentTokenRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEnrollmentTokenRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEnrollmentTokenRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpiring

`func (o *PatchedEnrollmentTokenRequest) GetExpiring() bool`

GetExpiring returns the Expiring field if non-nil, zero value otherwise.

### GetExpiringOk

`func (o *PatchedEnrollmentTokenRequest) GetExpiringOk() (*bool, bool)`

GetExpiringOk returns a tuple with the Expiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiring

`func (o *PatchedEnrollmentTokenRequest) SetExpiring(v bool)`

SetExpiring sets Expiring field to given value.

### HasExpiring

`func (o *PatchedEnrollmentTokenRequest) HasExpiring() bool`

HasExpiring returns a boolean if a field has been set.

### GetExpires

`func (o *PatchedEnrollmentTokenRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *PatchedEnrollmentTokenRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *PatchedEnrollmentTokenRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *PatchedEnrollmentTokenRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *PatchedEnrollmentTokenRequest) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *PatchedEnrollmentTokenRequest) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


