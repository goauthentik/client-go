# WebAuthnDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**CreatedOn** | **time.Time** |  | [readonly] 
**DeviceType** | [**NullableWebAuthnDeviceDeviceType**](WebAuthnDeviceDeviceType.md) |  | 
**Aaguid** | **string** |  | [readonly] 
**User** | [**GroupMember**](GroupMember.md) |  | [readonly] 

## Methods

### NewWebAuthnDevice

`func NewWebAuthnDevice(pk int32, name string, createdOn time.Time, deviceType NullableWebAuthnDeviceDeviceType, aaguid string, user GroupMember, ) *WebAuthnDevice`

NewWebAuthnDevice instantiates a new WebAuthnDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebAuthnDeviceWithDefaults

`func NewWebAuthnDeviceWithDefaults() *WebAuthnDevice`

NewWebAuthnDeviceWithDefaults instantiates a new WebAuthnDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *WebAuthnDevice) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *WebAuthnDevice) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *WebAuthnDevice) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetName

`func (o *WebAuthnDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebAuthnDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebAuthnDevice) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedOn

`func (o *WebAuthnDevice) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *WebAuthnDevice) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *WebAuthnDevice) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.


### GetDeviceType

`func (o *WebAuthnDevice) GetDeviceType() WebAuthnDeviceDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *WebAuthnDevice) GetDeviceTypeOk() (*WebAuthnDeviceDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *WebAuthnDevice) SetDeviceType(v WebAuthnDeviceDeviceType)`

SetDeviceType sets DeviceType field to given value.


### SetDeviceTypeNil

`func (o *WebAuthnDevice) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *WebAuthnDevice) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetAaguid

`func (o *WebAuthnDevice) GetAaguid() string`

GetAaguid returns the Aaguid field if non-nil, zero value otherwise.

### GetAaguidOk

`func (o *WebAuthnDevice) GetAaguidOk() (*string, bool)`

GetAaguidOk returns a tuple with the Aaguid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaguid

`func (o *WebAuthnDevice) SetAaguid(v string)`

SetAaguid sets Aaguid field to given value.


### GetUser

`func (o *WebAuthnDevice) GetUser() GroupMember`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *WebAuthnDevice) GetUserOk() (*GroupMember, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *WebAuthnDevice) SetUser(v GroupMember)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


