# AgentPSSODeviceState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceRegistered** | **bool** |  | 
**SignKeyId** | **string** |  | 
**EncKeyId** | **string** |  | 
**Users** | [**[]AgentPSSODeviceStateUser**](AgentPSSODeviceStateUser.md) |  | 

## Methods

### NewAgentPSSODeviceState

`func NewAgentPSSODeviceState(deviceRegistered bool, signKeyId string, encKeyId string, users []AgentPSSODeviceStateUser, ) *AgentPSSODeviceState`

NewAgentPSSODeviceState instantiates a new AgentPSSODeviceState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentPSSODeviceStateWithDefaults

`func NewAgentPSSODeviceStateWithDefaults() *AgentPSSODeviceState`

NewAgentPSSODeviceStateWithDefaults instantiates a new AgentPSSODeviceState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceRegistered

`func (o *AgentPSSODeviceState) GetDeviceRegistered() bool`

GetDeviceRegistered returns the DeviceRegistered field if non-nil, zero value otherwise.

### GetDeviceRegisteredOk

`func (o *AgentPSSODeviceState) GetDeviceRegisteredOk() (*bool, bool)`

GetDeviceRegisteredOk returns a tuple with the DeviceRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistered

`func (o *AgentPSSODeviceState) SetDeviceRegistered(v bool)`

SetDeviceRegistered sets DeviceRegistered field to given value.


### GetSignKeyId

`func (o *AgentPSSODeviceState) GetSignKeyId() string`

GetSignKeyId returns the SignKeyId field if non-nil, zero value otherwise.

### GetSignKeyIdOk

`func (o *AgentPSSODeviceState) GetSignKeyIdOk() (*string, bool)`

GetSignKeyIdOk returns a tuple with the SignKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignKeyId

`func (o *AgentPSSODeviceState) SetSignKeyId(v string)`

SetSignKeyId sets SignKeyId field to given value.


### GetEncKeyId

`func (o *AgentPSSODeviceState) GetEncKeyId() string`

GetEncKeyId returns the EncKeyId field if non-nil, zero value otherwise.

### GetEncKeyIdOk

`func (o *AgentPSSODeviceState) GetEncKeyIdOk() (*string, bool)`

GetEncKeyIdOk returns a tuple with the EncKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncKeyId

`func (o *AgentPSSODeviceState) SetEncKeyId(v string)`

SetEncKeyId sets EncKeyId field to given value.


### GetUsers

`func (o *AgentPSSODeviceState) GetUsers() []AgentPSSODeviceStateUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AgentPSSODeviceState) GetUsersOk() (*[]AgentPSSODeviceStateUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AgentPSSODeviceState) SetUsers(v []AgentPSSODeviceStateUser)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


