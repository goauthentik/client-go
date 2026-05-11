# DiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Mountpoint** | **string** |  | 
**Label** | Pointer to **string** |  | [optional] 
**CapacityTotalBytes** | Pointer to **int64** |  | [optional] 
**CapacityUsedBytes** | Pointer to **int64** |  | [optional] 
**EncryptionEnabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewDiskRequest

`func NewDiskRequest(name string, mountpoint string, ) *DiskRequest`

NewDiskRequest instantiates a new DiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskRequestWithDefaults

`func NewDiskRequestWithDefaults() *DiskRequest`

NewDiskRequestWithDefaults instantiates a new DiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DiskRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiskRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiskRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMountpoint

`func (o *DiskRequest) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *DiskRequest) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *DiskRequest) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.


### GetLabel

`func (o *DiskRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DiskRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DiskRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DiskRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCapacityTotalBytes

`func (o *DiskRequest) GetCapacityTotalBytes() int64`

GetCapacityTotalBytes returns the CapacityTotalBytes field if non-nil, zero value otherwise.

### GetCapacityTotalBytesOk

`func (o *DiskRequest) GetCapacityTotalBytesOk() (*int64, bool)`

GetCapacityTotalBytesOk returns a tuple with the CapacityTotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityTotalBytes

`func (o *DiskRequest) SetCapacityTotalBytes(v int64)`

SetCapacityTotalBytes sets CapacityTotalBytes field to given value.

### HasCapacityTotalBytes

`func (o *DiskRequest) HasCapacityTotalBytes() bool`

HasCapacityTotalBytes returns a boolean if a field has been set.

### GetCapacityUsedBytes

`func (o *DiskRequest) GetCapacityUsedBytes() int64`

GetCapacityUsedBytes returns the CapacityUsedBytes field if non-nil, zero value otherwise.

### GetCapacityUsedBytesOk

`func (o *DiskRequest) GetCapacityUsedBytesOk() (*int64, bool)`

GetCapacityUsedBytesOk returns a tuple with the CapacityUsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityUsedBytes

`func (o *DiskRequest) SetCapacityUsedBytes(v int64)`

SetCapacityUsedBytes sets CapacityUsedBytes field to given value.

### HasCapacityUsedBytes

`func (o *DiskRequest) HasCapacityUsedBytes() bool`

HasCapacityUsedBytes returns a boolean if a field has been set.

### GetEncryptionEnabled

`func (o *DiskRequest) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *DiskRequest) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *DiskRequest) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *DiskRequest) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


