# DeviceFactSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**DeviceFacts**](DeviceFacts.md) |  | 
**Connection** | **string** |  | 
**Created** | **time.Time** |  | [readonly] 
**Expires** | **NullableTime** |  | [readonly] 
**Vendor** | [**VendorEnum**](VendorEnum.md) |  | [readonly] 

## Methods

### NewDeviceFactSnapshot

`func NewDeviceFactSnapshot(data DeviceFacts, connection string, created time.Time, expires NullableTime, vendor VendorEnum, ) *DeviceFactSnapshot`

NewDeviceFactSnapshot instantiates a new DeviceFactSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceFactSnapshotWithDefaults

`func NewDeviceFactSnapshotWithDefaults() *DeviceFactSnapshot`

NewDeviceFactSnapshotWithDefaults instantiates a new DeviceFactSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DeviceFactSnapshot) GetData() DeviceFacts`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeviceFactSnapshot) GetDataOk() (*DeviceFacts, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeviceFactSnapshot) SetData(v DeviceFacts)`

SetData sets Data field to given value.


### GetConnection

`func (o *DeviceFactSnapshot) GetConnection() string`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *DeviceFactSnapshot) GetConnectionOk() (*string, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *DeviceFactSnapshot) SetConnection(v string)`

SetConnection sets Connection field to given value.


### GetCreated

`func (o *DeviceFactSnapshot) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeviceFactSnapshot) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeviceFactSnapshot) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetExpires

`func (o *DeviceFactSnapshot) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *DeviceFactSnapshot) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *DeviceFactSnapshot) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.


### SetExpiresNil

`func (o *DeviceFactSnapshot) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *DeviceFactSnapshot) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetVendor

`func (o *DeviceFactSnapshot) GetVendor() VendorEnum`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *DeviceFactSnapshot) GetVendorOk() (*VendorEnum, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *DeviceFactSnapshot) SetVendor(v VendorEnum)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


