# RACDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceUuid** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Protocols** | [**[]RACDeviceProtocol**](RACDeviceProtocol.md) |  | [readonly] 
**OverridePk** | **NullableInt32** | Primary key of this device&#39;s connection override, if it has one | [readonly] 

## Methods

### NewRACDevice

`func NewRACDevice(name string, protocols []RACDeviceProtocol, overridePk NullableInt32, ) *RACDevice`

NewRACDevice instantiates a new RACDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRACDeviceWithDefaults

`func NewRACDeviceWithDefaults() *RACDevice`

NewRACDeviceWithDefaults instantiates a new RACDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceUuid

`func (o *RACDevice) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *RACDevice) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *RACDevice) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.

### HasDeviceUuid

`func (o *RACDevice) HasDeviceUuid() bool`

HasDeviceUuid returns a boolean if a field has been set.

### GetName

`func (o *RACDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RACDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RACDevice) SetName(v string)`

SetName sets Name field to given value.


### GetProtocols

`func (o *RACDevice) GetProtocols() []RACDeviceProtocol`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *RACDevice) GetProtocolsOk() (*[]RACDeviceProtocol, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *RACDevice) SetProtocols(v []RACDeviceProtocol)`

SetProtocols sets Protocols field to given value.


### GetOverridePk

`func (o *RACDevice) GetOverridePk() int32`

GetOverridePk returns the OverridePk field if non-nil, zero value otherwise.

### GetOverridePkOk

`func (o *RACDevice) GetOverridePkOk() (*int32, bool)`

GetOverridePkOk returns a tuple with the OverridePk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePk

`func (o *RACDevice) SetOverridePk(v int32)`

SetOverridePk sets OverridePk field to given value.


### SetOverridePkNil

`func (o *RACDevice) SetOverridePkNil(b bool)`

 SetOverridePkNil sets the value for OverridePk to be an explicit nil

### UnsetOverridePk
`func (o *RACDevice) UnsetOverridePk()`

UnsetOverridePk ensures that no value is present for OverridePk, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


