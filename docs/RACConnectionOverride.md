# RACConnectionOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Hostname/IP to connect to. Optionally specify the port. | 
**Protocol** | [**ProtocolEnum**](ProtocolEnum.md) |  | 

## Methods

### NewRACConnectionOverride

`func NewRACConnectionOverride(host string, protocol ProtocolEnum, ) *RACConnectionOverride`

NewRACConnectionOverride instantiates a new RACConnectionOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRACConnectionOverrideWithDefaults

`func NewRACConnectionOverrideWithDefaults() *RACConnectionOverride`

NewRACConnectionOverrideWithDefaults instantiates a new RACConnectionOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *RACConnectionOverride) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RACConnectionOverride) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RACConnectionOverride) SetHost(v string)`

SetHost sets Host field to given value.


### GetProtocol

`func (o *RACConnectionOverride) GetProtocol() ProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *RACConnectionOverride) GetProtocolOk() (*ProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *RACConnectionOverride) SetProtocol(v ProtocolEnum)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


