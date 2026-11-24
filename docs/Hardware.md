# Hardware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to **string** |  | [optional] 
**Serial** | **string** |  | 
**CpuName** | Pointer to **string** |  | [optional] 
**CpuCount** | Pointer to **int32** |  | [optional] 
**MemoryBytes** | Pointer to **int64** |  | [optional] 

## Methods

### NewHardware

`func NewHardware(serial string, ) *Hardware`

NewHardware instantiates a new Hardware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHardwareWithDefaults

`func NewHardwareWithDefaults() *Hardware`

NewHardwareWithDefaults instantiates a new Hardware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *Hardware) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Hardware) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Hardware) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Hardware) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetManufacturer

`func (o *Hardware) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *Hardware) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *Hardware) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *Hardware) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetSerial

`func (o *Hardware) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Hardware) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Hardware) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetCpuName

`func (o *Hardware) GetCpuName() string`

GetCpuName returns the CpuName field if non-nil, zero value otherwise.

### GetCpuNameOk

`func (o *Hardware) GetCpuNameOk() (*string, bool)`

GetCpuNameOk returns a tuple with the CpuName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuName

`func (o *Hardware) SetCpuName(v string)`

SetCpuName sets CpuName field to given value.

### HasCpuName

`func (o *Hardware) HasCpuName() bool`

HasCpuName returns a boolean if a field has been set.

### GetCpuCount

`func (o *Hardware) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *Hardware) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *Hardware) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *Hardware) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetMemoryBytes

`func (o *Hardware) GetMemoryBytes() int64`

GetMemoryBytes returns the MemoryBytes field if non-nil, zero value otherwise.

### GetMemoryBytesOk

`func (o *Hardware) GetMemoryBytesOk() (*int64, bool)`

GetMemoryBytesOk returns a tuple with the MemoryBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryBytes

`func (o *Hardware) SetMemoryBytes(v int64)`

SetMemoryBytes sets MemoryBytes field to given value.

### HasMemoryBytes

`func (o *Hardware) HasMemoryBytes() bool`

HasMemoryBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


