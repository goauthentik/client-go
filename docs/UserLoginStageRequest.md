# UserLoginStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**SessionDuration** | Pointer to **string** | Determines how long a session lasts. Default of 0 means that the sessions lasts until the browser is closed. (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3) | [optional] 
**TerminateOtherSessions** | Pointer to **bool** | Terminate all other sessions of the user logging in. | [optional] 
**RememberMeOffset** | Pointer to **string** | Offset the session will be extended by when the user picks the remember me option. Default of 0 means that the remember me option will not be shown. (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3) | [optional] 
**NetworkBinding** | Pointer to [**NetworkBindingEnum**](NetworkBindingEnum.md) | Bind sessions created by this stage to the configured network | [optional] 
**GeoipBinding** | Pointer to [**GeoipBindingEnum**](GeoipBindingEnum.md) | Bind sessions created by this stage to the configured GeoIP location | [optional] 
**RememberDevice** | Pointer to **string** | When set to a non-zero value, authentik will save a cookie with a longer expiry,to remember the device the user is logging in from. (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3) | [optional] 

## Methods

### NewUserLoginStageRequest

`func NewUserLoginStageRequest(name string, ) *UserLoginStageRequest`

NewUserLoginStageRequest instantiates a new UserLoginStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLoginStageRequestWithDefaults

`func NewUserLoginStageRequestWithDefaults() *UserLoginStageRequest`

NewUserLoginStageRequestWithDefaults instantiates a new UserLoginStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserLoginStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserLoginStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserLoginStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowSet

`func (o *UserLoginStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *UserLoginStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *UserLoginStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *UserLoginStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetSessionDuration

`func (o *UserLoginStageRequest) GetSessionDuration() string`

GetSessionDuration returns the SessionDuration field if non-nil, zero value otherwise.

### GetSessionDurationOk

`func (o *UserLoginStageRequest) GetSessionDurationOk() (*string, bool)`

GetSessionDurationOk returns a tuple with the SessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDuration

`func (o *UserLoginStageRequest) SetSessionDuration(v string)`

SetSessionDuration sets SessionDuration field to given value.

### HasSessionDuration

`func (o *UserLoginStageRequest) HasSessionDuration() bool`

HasSessionDuration returns a boolean if a field has been set.

### GetTerminateOtherSessions

`func (o *UserLoginStageRequest) GetTerminateOtherSessions() bool`

GetTerminateOtherSessions returns the TerminateOtherSessions field if non-nil, zero value otherwise.

### GetTerminateOtherSessionsOk

`func (o *UserLoginStageRequest) GetTerminateOtherSessionsOk() (*bool, bool)`

GetTerminateOtherSessionsOk returns a tuple with the TerminateOtherSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminateOtherSessions

`func (o *UserLoginStageRequest) SetTerminateOtherSessions(v bool)`

SetTerminateOtherSessions sets TerminateOtherSessions field to given value.

### HasTerminateOtherSessions

`func (o *UserLoginStageRequest) HasTerminateOtherSessions() bool`

HasTerminateOtherSessions returns a boolean if a field has been set.

### GetRememberMeOffset

`func (o *UserLoginStageRequest) GetRememberMeOffset() string`

GetRememberMeOffset returns the RememberMeOffset field if non-nil, zero value otherwise.

### GetRememberMeOffsetOk

`func (o *UserLoginStageRequest) GetRememberMeOffsetOk() (*string, bool)`

GetRememberMeOffsetOk returns a tuple with the RememberMeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMeOffset

`func (o *UserLoginStageRequest) SetRememberMeOffset(v string)`

SetRememberMeOffset sets RememberMeOffset field to given value.

### HasRememberMeOffset

`func (o *UserLoginStageRequest) HasRememberMeOffset() bool`

HasRememberMeOffset returns a boolean if a field has been set.

### GetNetworkBinding

`func (o *UserLoginStageRequest) GetNetworkBinding() NetworkBindingEnum`

GetNetworkBinding returns the NetworkBinding field if non-nil, zero value otherwise.

### GetNetworkBindingOk

`func (o *UserLoginStageRequest) GetNetworkBindingOk() (*NetworkBindingEnum, bool)`

GetNetworkBindingOk returns a tuple with the NetworkBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBinding

`func (o *UserLoginStageRequest) SetNetworkBinding(v NetworkBindingEnum)`

SetNetworkBinding sets NetworkBinding field to given value.

### HasNetworkBinding

`func (o *UserLoginStageRequest) HasNetworkBinding() bool`

HasNetworkBinding returns a boolean if a field has been set.

### GetGeoipBinding

`func (o *UserLoginStageRequest) GetGeoipBinding() GeoipBindingEnum`

GetGeoipBinding returns the GeoipBinding field if non-nil, zero value otherwise.

### GetGeoipBindingOk

`func (o *UserLoginStageRequest) GetGeoipBindingOk() (*GeoipBindingEnum, bool)`

GetGeoipBindingOk returns a tuple with the GeoipBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoipBinding

`func (o *UserLoginStageRequest) SetGeoipBinding(v GeoipBindingEnum)`

SetGeoipBinding sets GeoipBinding field to given value.

### HasGeoipBinding

`func (o *UserLoginStageRequest) HasGeoipBinding() bool`

HasGeoipBinding returns a boolean if a field has been set.

### GetRememberDevice

`func (o *UserLoginStageRequest) GetRememberDevice() string`

GetRememberDevice returns the RememberDevice field if non-nil, zero value otherwise.

### GetRememberDeviceOk

`func (o *UserLoginStageRequest) GetRememberDeviceOk() (*string, bool)`

GetRememberDeviceOk returns a tuple with the RememberDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberDevice

`func (o *UserLoginStageRequest) SetRememberDevice(v string)`

SetRememberDevice sets RememberDevice field to given value.

### HasRememberDevice

`func (o *UserLoginStageRequest) HasRememberDevice() bool`

HasRememberDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


