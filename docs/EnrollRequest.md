# EnrollRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceSerial** | **string** |  | 
**DeviceName** | **string** |  | 

## Methods

### NewEnrollRequest

`func NewEnrollRequest(deviceSerial string, deviceName string, ) *EnrollRequest`

NewEnrollRequest instantiates a new EnrollRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollRequestWithDefaults

`func NewEnrollRequestWithDefaults() *EnrollRequest`

NewEnrollRequestWithDefaults instantiates a new EnrollRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceSerial

`func (o *EnrollRequest) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *EnrollRequest) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *EnrollRequest) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.


### GetDeviceName

`func (o *EnrollRequest) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *EnrollRequest) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *EnrollRequest) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


