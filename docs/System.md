# System

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpHeaders** | **map[string]string** | Get HTTP Request headers | [readonly] 
**HttpHost** | **string** | Get HTTP host | [readonly] 
**HttpIsSecure** | **bool** | Get HTTP Secure flag | [readonly] 
**Runtime** | [**SystemRuntime**](SystemRuntime.md) |  | 
**Tenant** | **string** | Currently active tenant | [readonly] 
**ServerTime** | **time.Time** | Current server time | [readonly] 
**EmbeddedOutpostHost** | **string** | Get the FQDN configured on the embedded outpost | [readonly] 

## Methods

### NewSystem

`func NewSystem(httpHeaders map[string]string, httpHost string, httpIsSecure bool, runtime SystemRuntime, tenant string, serverTime time.Time, embeddedOutpostHost string, ) *System`

NewSystem instantiates a new System object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemWithDefaults

`func NewSystemWithDefaults() *System`

NewSystemWithDefaults instantiates a new System object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpHeaders

`func (o *System) GetHttpHeaders() map[string]string`

GetHttpHeaders returns the HttpHeaders field if non-nil, zero value otherwise.

### GetHttpHeadersOk

`func (o *System) GetHttpHeadersOk() (*map[string]string, bool)`

GetHttpHeadersOk returns a tuple with the HttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHeaders

`func (o *System) SetHttpHeaders(v map[string]string)`

SetHttpHeaders sets HttpHeaders field to given value.


### GetHttpHost

`func (o *System) GetHttpHost() string`

GetHttpHost returns the HttpHost field if non-nil, zero value otherwise.

### GetHttpHostOk

`func (o *System) GetHttpHostOk() (*string, bool)`

GetHttpHostOk returns a tuple with the HttpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHost

`func (o *System) SetHttpHost(v string)`

SetHttpHost sets HttpHost field to given value.


### GetHttpIsSecure

`func (o *System) GetHttpIsSecure() bool`

GetHttpIsSecure returns the HttpIsSecure field if non-nil, zero value otherwise.

### GetHttpIsSecureOk

`func (o *System) GetHttpIsSecureOk() (*bool, bool)`

GetHttpIsSecureOk returns a tuple with the HttpIsSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpIsSecure

`func (o *System) SetHttpIsSecure(v bool)`

SetHttpIsSecure sets HttpIsSecure field to given value.


### GetRuntime

`func (o *System) GetRuntime() SystemRuntime`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *System) GetRuntimeOk() (*SystemRuntime, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *System) SetRuntime(v SystemRuntime)`

SetRuntime sets Runtime field to given value.


### GetTenant

`func (o *System) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *System) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *System) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetServerTime

`func (o *System) GetServerTime() time.Time`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *System) GetServerTimeOk() (*time.Time, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *System) SetServerTime(v time.Time)`

SetServerTime sets ServerTime field to given value.


### GetEmbeddedOutpostHost

`func (o *System) GetEmbeddedOutpostHost() string`

GetEmbeddedOutpostHost returns the EmbeddedOutpostHost field if non-nil, zero value otherwise.

### GetEmbeddedOutpostHostOk

`func (o *System) GetEmbeddedOutpostHostOk() (*string, bool)`

GetEmbeddedOutpostHostOk returns a tuple with the EmbeddedOutpostHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddedOutpostHost

`func (o *System) SetEmbeddedOutpostHost(v string)`

SetEmbeddedOutpostHost sets EmbeddedOutpostHost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


