# SystemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpHeaders** | **map[string]string** | Get HTTP Request headers | [readonly] 
**HttpHost** | **string** | Get HTTP host | [readonly] 
**HttpIsSecure** | **bool** | Get HTTP Secure flag | [readonly] 
**Runtime** | [**SystemInfoRuntime**](SystemInfoRuntime.md) |  | 
**Brand** | **string** | Currently active brand | [readonly] 
**ServerTime** | **time.Time** | Current server time | [readonly] 
**EmbeddedOutpostDisabled** | **bool** | Whether the embedded outpost is disabled | [readonly] 
**EmbeddedOutpostHost** | **string** | Get the FQDN configured on the embedded outpost | [readonly] 

## Methods

### NewSystemInfo

`func NewSystemInfo(httpHeaders map[string]string, httpHost string, httpIsSecure bool, runtime SystemInfoRuntime, brand string, serverTime time.Time, embeddedOutpostDisabled bool, embeddedOutpostHost string, ) *SystemInfo`

NewSystemInfo instantiates a new SystemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInfoWithDefaults

`func NewSystemInfoWithDefaults() *SystemInfo`

NewSystemInfoWithDefaults instantiates a new SystemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpHeaders

`func (o *SystemInfo) GetHttpHeaders() map[string]string`

GetHttpHeaders returns the HttpHeaders field if non-nil, zero value otherwise.

### GetHttpHeadersOk

`func (o *SystemInfo) GetHttpHeadersOk() (*map[string]string, bool)`

GetHttpHeadersOk returns a tuple with the HttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHeaders

`func (o *SystemInfo) SetHttpHeaders(v map[string]string)`

SetHttpHeaders sets HttpHeaders field to given value.


### GetHttpHost

`func (o *SystemInfo) GetHttpHost() string`

GetHttpHost returns the HttpHost field if non-nil, zero value otherwise.

### GetHttpHostOk

`func (o *SystemInfo) GetHttpHostOk() (*string, bool)`

GetHttpHostOk returns a tuple with the HttpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHost

`func (o *SystemInfo) SetHttpHost(v string)`

SetHttpHost sets HttpHost field to given value.


### GetHttpIsSecure

`func (o *SystemInfo) GetHttpIsSecure() bool`

GetHttpIsSecure returns the HttpIsSecure field if non-nil, zero value otherwise.

### GetHttpIsSecureOk

`func (o *SystemInfo) GetHttpIsSecureOk() (*bool, bool)`

GetHttpIsSecureOk returns a tuple with the HttpIsSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpIsSecure

`func (o *SystemInfo) SetHttpIsSecure(v bool)`

SetHttpIsSecure sets HttpIsSecure field to given value.


### GetRuntime

`func (o *SystemInfo) GetRuntime() SystemInfoRuntime`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *SystemInfo) GetRuntimeOk() (*SystemInfoRuntime, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *SystemInfo) SetRuntime(v SystemInfoRuntime)`

SetRuntime sets Runtime field to given value.


### GetBrand

`func (o *SystemInfo) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *SystemInfo) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *SystemInfo) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetServerTime

`func (o *SystemInfo) GetServerTime() time.Time`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *SystemInfo) GetServerTimeOk() (*time.Time, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *SystemInfo) SetServerTime(v time.Time)`

SetServerTime sets ServerTime field to given value.


### GetEmbeddedOutpostDisabled

`func (o *SystemInfo) GetEmbeddedOutpostDisabled() bool`

GetEmbeddedOutpostDisabled returns the EmbeddedOutpostDisabled field if non-nil, zero value otherwise.

### GetEmbeddedOutpostDisabledOk

`func (o *SystemInfo) GetEmbeddedOutpostDisabledOk() (*bool, bool)`

GetEmbeddedOutpostDisabledOk returns a tuple with the EmbeddedOutpostDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddedOutpostDisabled

`func (o *SystemInfo) SetEmbeddedOutpostDisabled(v bool)`

SetEmbeddedOutpostDisabled sets EmbeddedOutpostDisabled field to given value.


### GetEmbeddedOutpostHost

`func (o *SystemInfo) GetEmbeddedOutpostHost() string`

GetEmbeddedOutpostHost returns the EmbeddedOutpostHost field if non-nil, zero value otherwise.

### GetEmbeddedOutpostHostOk

`func (o *SystemInfo) GetEmbeddedOutpostHostOk() (*string, bool)`

GetEmbeddedOutpostHostOk returns a tuple with the EmbeddedOutpostHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddedOutpostHost

`func (o *SystemInfo) SetEmbeddedOutpostHost(v string)`

SetEmbeddedOutpostHost sets EmbeddedOutpostHost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


