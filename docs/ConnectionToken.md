# ConnectionToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Provider** | **int32** |  | 
**ProviderObj** | [**RACProvider**](RACProvider.md) |  | [readonly] 
**Endpoint** | **string** |  | [readonly] 
**EndpointObj** | [**Endpoint**](Endpoint.md) |  | [readonly] 
**User** | [**GroupMember**](GroupMember.md) |  | [readonly] 

## Methods

### NewConnectionToken

`func NewConnectionToken(pk string, provider int32, providerObj RACProvider, endpoint string, endpointObj Endpoint, user GroupMember, ) *ConnectionToken`

NewConnectionToken instantiates a new ConnectionToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionTokenWithDefaults

`func NewConnectionTokenWithDefaults() *ConnectionToken`

NewConnectionTokenWithDefaults instantiates a new ConnectionToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *ConnectionToken) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *ConnectionToken) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *ConnectionToken) SetPk(v string)`

SetPk sets Pk field to given value.


### GetProvider

`func (o *ConnectionToken) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ConnectionToken) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ConnectionToken) SetProvider(v int32)`

SetProvider sets Provider field to given value.


### GetProviderObj

`func (o *ConnectionToken) GetProviderObj() RACProvider`

GetProviderObj returns the ProviderObj field if non-nil, zero value otherwise.

### GetProviderObjOk

`func (o *ConnectionToken) GetProviderObjOk() (*RACProvider, bool)`

GetProviderObjOk returns a tuple with the ProviderObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderObj

`func (o *ConnectionToken) SetProviderObj(v RACProvider)`

SetProviderObj sets ProviderObj field to given value.


### GetEndpoint

`func (o *ConnectionToken) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ConnectionToken) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ConnectionToken) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetEndpointObj

`func (o *ConnectionToken) GetEndpointObj() Endpoint`

GetEndpointObj returns the EndpointObj field if non-nil, zero value otherwise.

### GetEndpointObjOk

`func (o *ConnectionToken) GetEndpointObjOk() (*Endpoint, bool)`

GetEndpointObjOk returns a tuple with the EndpointObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointObj

`func (o *ConnectionToken) SetEndpointObj(v Endpoint)`

SetEndpointObj sets EndpointObj field to given value.


### GetUser

`func (o *ConnectionToken) GetUser() GroupMember`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ConnectionToken) GetUserOk() (*GroupMember, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ConnectionToken) SetUser(v GroupMember)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


