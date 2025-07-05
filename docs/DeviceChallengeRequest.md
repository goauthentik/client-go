# DeviceChallengeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceClass** | [**DeviceClassesEnum**](DeviceClassesEnum.md) |  | 
**DeviceUid** | **string** |  | 
**Challenge** | **map[string]interface{}** |  | 
**LastUsed** | **NullableTime** |  | 

## Methods

### NewDeviceChallengeRequest

`func NewDeviceChallengeRequest(deviceClass DeviceClassesEnum, deviceUid string, challenge map[string]interface{}, lastUsed NullableTime, ) *DeviceChallengeRequest`

NewDeviceChallengeRequest instantiates a new DeviceChallengeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceChallengeRequestWithDefaults

`func NewDeviceChallengeRequestWithDefaults() *DeviceChallengeRequest`

NewDeviceChallengeRequestWithDefaults instantiates a new DeviceChallengeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceClass

`func (o *DeviceChallengeRequest) GetDeviceClass() DeviceClassesEnum`

GetDeviceClass returns the DeviceClass field if non-nil, zero value otherwise.

### GetDeviceClassOk

`func (o *DeviceChallengeRequest) GetDeviceClassOk() (*DeviceClassesEnum, bool)`

GetDeviceClassOk returns a tuple with the DeviceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClass

`func (o *DeviceChallengeRequest) SetDeviceClass(v DeviceClassesEnum)`

SetDeviceClass sets DeviceClass field to given value.


### GetDeviceUid

`func (o *DeviceChallengeRequest) GetDeviceUid() string`

GetDeviceUid returns the DeviceUid field if non-nil, zero value otherwise.

### GetDeviceUidOk

`func (o *DeviceChallengeRequest) GetDeviceUidOk() (*string, bool)`

GetDeviceUidOk returns a tuple with the DeviceUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUid

`func (o *DeviceChallengeRequest) SetDeviceUid(v string)`

SetDeviceUid sets DeviceUid field to given value.


### GetChallenge

`func (o *DeviceChallengeRequest) GetChallenge() map[string]interface{}`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *DeviceChallengeRequest) GetChallengeOk() (*map[string]interface{}, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *DeviceChallengeRequest) SetChallenge(v map[string]interface{})`

SetChallenge sets Challenge field to given value.


### GetLastUsed

`func (o *DeviceChallengeRequest) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *DeviceChallengeRequest) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *DeviceChallengeRequest) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.


### SetLastUsedNil

`func (o *DeviceChallengeRequest) SetLastUsedNil(b bool)`

 SetLastUsedNil sets the value for LastUsed to be an explicit nil

### UnsetLastUsed
`func (o *DeviceChallengeRequest) UnsetLastUsed()`

UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


