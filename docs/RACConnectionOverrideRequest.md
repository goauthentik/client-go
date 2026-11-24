# RACConnectionOverrideRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Hostname/IP to connect to. Optionally specify the port. | 
**Protocol** | [**ProtocolEnum**](ProtocolEnum.md) |  | 

## Methods

### NewRACConnectionOverrideRequest

`func NewRACConnectionOverrideRequest(host string, protocol ProtocolEnum, ) *RACConnectionOverrideRequest`

NewRACConnectionOverrideRequest instantiates a new RACConnectionOverrideRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRACConnectionOverrideRequestWithDefaults

`func NewRACConnectionOverrideRequestWithDefaults() *RACConnectionOverrideRequest`

NewRACConnectionOverrideRequestWithDefaults instantiates a new RACConnectionOverrideRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *RACConnectionOverrideRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RACConnectionOverrideRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RACConnectionOverrideRequest) SetHost(v string)`

SetHost sets Host field to given value.


### GetProtocol

`func (o *RACConnectionOverrideRequest) GetProtocol() ProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *RACConnectionOverrideRequest) GetProtocolOk() (*ProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *RACConnectionOverrideRequest) SetProtocol(v ProtocolEnum)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


