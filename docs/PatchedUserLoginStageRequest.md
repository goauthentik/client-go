# PatchedUserLoginStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**SessionDuration** | Pointer to **string** | Determines how long a session lasts. Default of 0 means that the sessions lasts until the browser is closed. (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3) | [optional] 
**TerminateOtherSessions** | Pointer to **bool** | Terminate all other sessions of the user logging in. | [optional] 
**RememberMeOffset** | Pointer to **string** | Offset the session will be extended by when the user picks the remember me option. Default of 0 means that the remember me option will not be shown. (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3) | [optional] 
**NetworkBinding** | Pointer to [**NetworkBindingEnum**](NetworkBindingEnum.md) | Bind sessions created by this stage to the configured network | [optional] 
**GeoipBinding** | Pointer to [**GeoipBindingEnum**](GeoipBindingEnum.md) | Bind sessions created by this stage to the configured GeoIP location | [optional] 
**RememberDevice** | Pointer to **string** | When set to a non-zero value, authentik will save a cookie with a longer expiry,to remember the device the user is logging in from. (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3) | [optional] 

## Methods

### NewPatchedUserLoginStageRequest

`func NewPatchedUserLoginStageRequest() *PatchedUserLoginStageRequest`

NewPatchedUserLoginStageRequest instantiates a new PatchedUserLoginStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedUserLoginStageRequestWithDefaults

`func NewPatchedUserLoginStageRequestWithDefaults() *PatchedUserLoginStageRequest`

NewPatchedUserLoginStageRequestWithDefaults instantiates a new PatchedUserLoginStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedUserLoginStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedUserLoginStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedUserLoginStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedUserLoginStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSessionDuration

`func (o *PatchedUserLoginStageRequest) GetSessionDuration() string`

GetSessionDuration returns the SessionDuration field if non-nil, zero value otherwise.

### GetSessionDurationOk

`func (o *PatchedUserLoginStageRequest) GetSessionDurationOk() (*string, bool)`

GetSessionDurationOk returns a tuple with the SessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDuration

`func (o *PatchedUserLoginStageRequest) SetSessionDuration(v string)`

SetSessionDuration sets SessionDuration field to given value.

### HasSessionDuration

`func (o *PatchedUserLoginStageRequest) HasSessionDuration() bool`

HasSessionDuration returns a boolean if a field has been set.

### GetTerminateOtherSessions

`func (o *PatchedUserLoginStageRequest) GetTerminateOtherSessions() bool`

GetTerminateOtherSessions returns the TerminateOtherSessions field if non-nil, zero value otherwise.

### GetTerminateOtherSessionsOk

`func (o *PatchedUserLoginStageRequest) GetTerminateOtherSessionsOk() (*bool, bool)`

GetTerminateOtherSessionsOk returns a tuple with the TerminateOtherSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminateOtherSessions

`func (o *PatchedUserLoginStageRequest) SetTerminateOtherSessions(v bool)`

SetTerminateOtherSessions sets TerminateOtherSessions field to given value.

### HasTerminateOtherSessions

`func (o *PatchedUserLoginStageRequest) HasTerminateOtherSessions() bool`

HasTerminateOtherSessions returns a boolean if a field has been set.

### GetRememberMeOffset

`func (o *PatchedUserLoginStageRequest) GetRememberMeOffset() string`

GetRememberMeOffset returns the RememberMeOffset field if non-nil, zero value otherwise.

### GetRememberMeOffsetOk

`func (o *PatchedUserLoginStageRequest) GetRememberMeOffsetOk() (*string, bool)`

GetRememberMeOffsetOk returns a tuple with the RememberMeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMeOffset

`func (o *PatchedUserLoginStageRequest) SetRememberMeOffset(v string)`

SetRememberMeOffset sets RememberMeOffset field to given value.

### HasRememberMeOffset

`func (o *PatchedUserLoginStageRequest) HasRememberMeOffset() bool`

HasRememberMeOffset returns a boolean if a field has been set.

### GetNetworkBinding

`func (o *PatchedUserLoginStageRequest) GetNetworkBinding() NetworkBindingEnum`

GetNetworkBinding returns the NetworkBinding field if non-nil, zero value otherwise.

### GetNetworkBindingOk

`func (o *PatchedUserLoginStageRequest) GetNetworkBindingOk() (*NetworkBindingEnum, bool)`

GetNetworkBindingOk returns a tuple with the NetworkBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBinding

`func (o *PatchedUserLoginStageRequest) SetNetworkBinding(v NetworkBindingEnum)`

SetNetworkBinding sets NetworkBinding field to given value.

### HasNetworkBinding

`func (o *PatchedUserLoginStageRequest) HasNetworkBinding() bool`

HasNetworkBinding returns a boolean if a field has been set.

### GetGeoipBinding

`func (o *PatchedUserLoginStageRequest) GetGeoipBinding() GeoipBindingEnum`

GetGeoipBinding returns the GeoipBinding field if non-nil, zero value otherwise.

### GetGeoipBindingOk

`func (o *PatchedUserLoginStageRequest) GetGeoipBindingOk() (*GeoipBindingEnum, bool)`

GetGeoipBindingOk returns a tuple with the GeoipBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoipBinding

`func (o *PatchedUserLoginStageRequest) SetGeoipBinding(v GeoipBindingEnum)`

SetGeoipBinding sets GeoipBinding field to given value.

### HasGeoipBinding

`func (o *PatchedUserLoginStageRequest) HasGeoipBinding() bool`

HasGeoipBinding returns a boolean if a field has been set.

### GetRememberDevice

`func (o *PatchedUserLoginStageRequest) GetRememberDevice() string`

GetRememberDevice returns the RememberDevice field if non-nil, zero value otherwise.

### GetRememberDeviceOk

`func (o *PatchedUserLoginStageRequest) GetRememberDeviceOk() (*string, bool)`

GetRememberDeviceOk returns a tuple with the RememberDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberDevice

`func (o *PatchedUserLoginStageRequest) SetRememberDevice(v string)`

SetRememberDevice sets RememberDevice field to given value.

### HasRememberDevice

`func (o *PatchedUserLoginStageRequest) HasRememberDevice() bool`

HasRememberDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


