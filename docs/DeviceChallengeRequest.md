# DeviceChallengeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceClass** | **string** |  | 
**DeviceUid** | **string** |  | 
**Challenge** | **map[string]interface{}** |  | 

## Methods

### NewDeviceChallengeRequest

`func NewDeviceChallengeRequest(deviceClass string, deviceUid string, challenge map[string]interface{}, ) *DeviceChallengeRequest`

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

`func (o *DeviceChallengeRequest) GetDeviceClass() string`

GetDeviceClass returns the DeviceClass field if non-nil, zero value otherwise.

### GetDeviceClassOk

`func (o *DeviceChallengeRequest) GetDeviceClassOk() (*string, bool)`

GetDeviceClassOk returns a tuple with the DeviceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClass

`func (o *DeviceChallengeRequest) SetDeviceClass(v string)`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


