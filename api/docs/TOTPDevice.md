# TOTPDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The human-readable name of this device. | 
**Pk** | **int32** |  | [readonly] 

## Methods

### NewTOTPDevice

`func NewTOTPDevice(name string, pk int32, ) *TOTPDevice`

NewTOTPDevice instantiates a new TOTPDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTOTPDeviceWithDefaults

`func NewTOTPDeviceWithDefaults() *TOTPDevice`

NewTOTPDeviceWithDefaults instantiates a new TOTPDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TOTPDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TOTPDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TOTPDevice) SetName(v string)`

SetName sets Name field to given value.


### GetPk

`func (o *TOTPDevice) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *TOTPDevice) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *TOTPDevice) SetPk(v int32)`

SetPk sets Pk field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


