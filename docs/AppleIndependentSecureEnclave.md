# AppleIndependentSecureEnclave

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**User** | **int32** | The user that this device belongs to. | 
**AppleSecureEnclaveKey** | **string** |  | 
**AppleEnclaveKeyId** | **string** |  | 
**DeviceType** | **string** |  | 

## Methods

### NewAppleIndependentSecureEnclave

`func NewAppleIndependentSecureEnclave(user int32, appleSecureEnclaveKey string, appleEnclaveKeyId string, deviceType string, ) *AppleIndependentSecureEnclave`

NewAppleIndependentSecureEnclave instantiates a new AppleIndependentSecureEnclave object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppleIndependentSecureEnclaveWithDefaults

`func NewAppleIndependentSecureEnclaveWithDefaults() *AppleIndependentSecureEnclave`

NewAppleIndependentSecureEnclaveWithDefaults instantiates a new AppleIndependentSecureEnclave object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *AppleIndependentSecureEnclave) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppleIndependentSecureEnclave) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppleIndependentSecureEnclave) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AppleIndependentSecureEnclave) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetUser

`func (o *AppleIndependentSecureEnclave) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AppleIndependentSecureEnclave) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AppleIndependentSecureEnclave) SetUser(v int32)`

SetUser sets User field to given value.


### GetAppleSecureEnclaveKey

`func (o *AppleIndependentSecureEnclave) GetAppleSecureEnclaveKey() string`

GetAppleSecureEnclaveKey returns the AppleSecureEnclaveKey field if non-nil, zero value otherwise.

### GetAppleSecureEnclaveKeyOk

`func (o *AppleIndependentSecureEnclave) GetAppleSecureEnclaveKeyOk() (*string, bool)`

GetAppleSecureEnclaveKeyOk returns a tuple with the AppleSecureEnclaveKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleSecureEnclaveKey

`func (o *AppleIndependentSecureEnclave) SetAppleSecureEnclaveKey(v string)`

SetAppleSecureEnclaveKey sets AppleSecureEnclaveKey field to given value.


### GetAppleEnclaveKeyId

`func (o *AppleIndependentSecureEnclave) GetAppleEnclaveKeyId() string`

GetAppleEnclaveKeyId returns the AppleEnclaveKeyId field if non-nil, zero value otherwise.

### GetAppleEnclaveKeyIdOk

`func (o *AppleIndependentSecureEnclave) GetAppleEnclaveKeyIdOk() (*string, bool)`

GetAppleEnclaveKeyIdOk returns a tuple with the AppleEnclaveKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleEnclaveKeyId

`func (o *AppleIndependentSecureEnclave) SetAppleEnclaveKeyId(v string)`

SetAppleEnclaveKeyId sets AppleEnclaveKeyId field to given value.


### GetDeviceType

`func (o *AppleIndependentSecureEnclave) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *AppleIndependentSecureEnclave) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *AppleIndependentSecureEnclave) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


