# DeviceConnectionLatestSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**DeviceFacts**](DeviceFacts.md) |  | 
**Connection** | **string** |  | 
**Created** | **time.Time** |  | [readonly] 
**Expires** | **NullableTime** |  | [readonly] 
**Vendor** | [**VendorEnum**](VendorEnum.md) |  | [readonly] 

## Methods

### NewDeviceConnectionLatestSnapshot

`func NewDeviceConnectionLatestSnapshot(data DeviceFacts, connection string, created time.Time, expires NullableTime, vendor VendorEnum, ) *DeviceConnectionLatestSnapshot`

NewDeviceConnectionLatestSnapshot instantiates a new DeviceConnectionLatestSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceConnectionLatestSnapshotWithDefaults

`func NewDeviceConnectionLatestSnapshotWithDefaults() *DeviceConnectionLatestSnapshot`

NewDeviceConnectionLatestSnapshotWithDefaults instantiates a new DeviceConnectionLatestSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DeviceConnectionLatestSnapshot) GetData() DeviceFacts`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeviceConnectionLatestSnapshot) GetDataOk() (*DeviceFacts, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeviceConnectionLatestSnapshot) SetData(v DeviceFacts)`

SetData sets Data field to given value.


### GetConnection

`func (o *DeviceConnectionLatestSnapshot) GetConnection() string`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *DeviceConnectionLatestSnapshot) GetConnectionOk() (*string, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *DeviceConnectionLatestSnapshot) SetConnection(v string)`

SetConnection sets Connection field to given value.


### GetCreated

`func (o *DeviceConnectionLatestSnapshot) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeviceConnectionLatestSnapshot) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeviceConnectionLatestSnapshot) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetExpires

`func (o *DeviceConnectionLatestSnapshot) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *DeviceConnectionLatestSnapshot) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *DeviceConnectionLatestSnapshot) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.


### SetExpiresNil

`func (o *DeviceConnectionLatestSnapshot) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *DeviceConnectionLatestSnapshot) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetVendor

`func (o *DeviceConnectionLatestSnapshot) GetVendor() VendorEnum`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *DeviceConnectionLatestSnapshot) GetVendorOk() (*VendorEnum, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *DeviceConnectionLatestSnapshot) SetVendor(v VendorEnum)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


