# EnrollmentToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenUuid** | **string** |  | [readonly] 
**DeviceGroup** | Pointer to **NullableString** |  | [optional] 
**DeviceGroupObj** | [**DeviceGroup**](DeviceGroup.md) |  | [readonly] 
**Connector** | **string** |  | 
**Name** | **string** |  | 
**Expiring** | Pointer to **bool** |  | [optional] 
**Expires** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewEnrollmentToken

`func NewEnrollmentToken(tokenUuid string, deviceGroupObj DeviceGroup, connector string, name string, ) *EnrollmentToken`

NewEnrollmentToken instantiates a new EnrollmentToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentTokenWithDefaults

`func NewEnrollmentTokenWithDefaults() *EnrollmentToken`

NewEnrollmentTokenWithDefaults instantiates a new EnrollmentToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenUuid

`func (o *EnrollmentToken) GetTokenUuid() string`

GetTokenUuid returns the TokenUuid field if non-nil, zero value otherwise.

### GetTokenUuidOk

`func (o *EnrollmentToken) GetTokenUuidOk() (*string, bool)`

GetTokenUuidOk returns a tuple with the TokenUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUuid

`func (o *EnrollmentToken) SetTokenUuid(v string)`

SetTokenUuid sets TokenUuid field to given value.


### GetDeviceGroup

`func (o *EnrollmentToken) GetDeviceGroup() string`

GetDeviceGroup returns the DeviceGroup field if non-nil, zero value otherwise.

### GetDeviceGroupOk

`func (o *EnrollmentToken) GetDeviceGroupOk() (*string, bool)`

GetDeviceGroupOk returns a tuple with the DeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroup

`func (o *EnrollmentToken) SetDeviceGroup(v string)`

SetDeviceGroup sets DeviceGroup field to given value.

### HasDeviceGroup

`func (o *EnrollmentToken) HasDeviceGroup() bool`

HasDeviceGroup returns a boolean if a field has been set.

### SetDeviceGroupNil

`func (o *EnrollmentToken) SetDeviceGroupNil(b bool)`

 SetDeviceGroupNil sets the value for DeviceGroup to be an explicit nil

### UnsetDeviceGroup
`func (o *EnrollmentToken) UnsetDeviceGroup()`

UnsetDeviceGroup ensures that no value is present for DeviceGroup, not even an explicit nil
### GetDeviceGroupObj

`func (o *EnrollmentToken) GetDeviceGroupObj() DeviceGroup`

GetDeviceGroupObj returns the DeviceGroupObj field if non-nil, zero value otherwise.

### GetDeviceGroupObjOk

`func (o *EnrollmentToken) GetDeviceGroupObjOk() (*DeviceGroup, bool)`

GetDeviceGroupObjOk returns a tuple with the DeviceGroupObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroupObj

`func (o *EnrollmentToken) SetDeviceGroupObj(v DeviceGroup)`

SetDeviceGroupObj sets DeviceGroupObj field to given value.


### GetConnector

`func (o *EnrollmentToken) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *EnrollmentToken) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *EnrollmentToken) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetName

`func (o *EnrollmentToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnrollmentToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnrollmentToken) SetName(v string)`

SetName sets Name field to given value.


### GetExpiring

`func (o *EnrollmentToken) GetExpiring() bool`

GetExpiring returns the Expiring field if non-nil, zero value otherwise.

### GetExpiringOk

`func (o *EnrollmentToken) GetExpiringOk() (*bool, bool)`

GetExpiringOk returns a tuple with the Expiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiring

`func (o *EnrollmentToken) SetExpiring(v bool)`

SetExpiring sets Expiring field to given value.

### HasExpiring

`func (o *EnrollmentToken) HasExpiring() bool`

HasExpiring returns a boolean if a field has been set.

### GetExpires

`func (o *EnrollmentToken) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *EnrollmentToken) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *EnrollmentToken) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *EnrollmentToken) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *EnrollmentToken) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *EnrollmentToken) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


