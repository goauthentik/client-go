# WebAuthnDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**CreatedOn** | **time.Time** |  | [readonly] 

## Methods

### NewWebAuthnDevice

`func NewWebAuthnDevice(pk int32, name string, createdOn time.Time, ) *WebAuthnDevice`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


