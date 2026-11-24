# SystemInfoPostgresql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**Jit** | **string** |  | 
**JitSource** | **string** |  | 
**Ssl** | **bool** |  | 
**SslVersion** | **NullableString** |  | 
**SslCipher** | **NullableString** |  | 

## Methods

### NewSystemInfoPostgresql

`func NewSystemInfoPostgresql(version string, jit string, jitSource string, ssl bool, sslVersion NullableString, sslCipher NullableString, ) *SystemInfoPostgresql`

NewSystemInfoPostgresql instantiates a new SystemInfoPostgresql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInfoPostgresqlWithDefaults

`func NewSystemInfoPostgresqlWithDefaults() *SystemInfoPostgresql`

NewSystemInfoPostgresqlWithDefaults instantiates a new SystemInfoPostgresql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *SystemInfoPostgresql) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SystemInfoPostgresql) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SystemInfoPostgresql) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetJit

`func (o *SystemInfoPostgresql) GetJit() string`

GetJit returns the Jit field if non-nil, zero value otherwise.

### GetJitOk

`func (o *SystemInfoPostgresql) GetJitOk() (*string, bool)`

GetJitOk returns a tuple with the Jit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJit

`func (o *SystemInfoPostgresql) SetJit(v string)`

SetJit sets Jit field to given value.


### GetJitSource

`func (o *SystemInfoPostgresql) GetJitSource() string`

GetJitSource returns the JitSource field if non-nil, zero value otherwise.

### GetJitSourceOk

`func (o *SystemInfoPostgresql) GetJitSourceOk() (*string, bool)`

GetJitSourceOk returns a tuple with the JitSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitSource

`func (o *SystemInfoPostgresql) SetJitSource(v string)`

SetJitSource sets JitSource field to given value.


### GetSsl

`func (o *SystemInfoPostgresql) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *SystemInfoPostgresql) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *SystemInfoPostgresql) SetSsl(v bool)`

SetSsl sets Ssl field to given value.


### GetSslVersion

`func (o *SystemInfoPostgresql) GetSslVersion() string`

GetSslVersion returns the SslVersion field if non-nil, zero value otherwise.

### GetSslVersionOk

`func (o *SystemInfoPostgresql) GetSslVersionOk() (*string, bool)`

GetSslVersionOk returns a tuple with the SslVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVersion

`func (o *SystemInfoPostgresql) SetSslVersion(v string)`

SetSslVersion sets SslVersion field to given value.


### SetSslVersionNil

`func (o *SystemInfoPostgresql) SetSslVersionNil(b bool)`

 SetSslVersionNil sets the value for SslVersion to be an explicit nil

### UnsetSslVersion
`func (o *SystemInfoPostgresql) UnsetSslVersion()`

UnsetSslVersion ensures that no value is present for SslVersion, not even an explicit nil
### GetSslCipher

`func (o *SystemInfoPostgresql) GetSslCipher() string`

GetSslCipher returns the SslCipher field if non-nil, zero value otherwise.

### GetSslCipherOk

`func (o *SystemInfoPostgresql) GetSslCipherOk() (*string, bool)`

GetSslCipherOk returns a tuple with the SslCipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCipher

`func (o *SystemInfoPostgresql) SetSslCipher(v string)`

SetSslCipher sets SslCipher field to given value.


### SetSslCipherNil

`func (o *SystemInfoPostgresql) SetSslCipherNil(b bool)`

 SetSslCipherNil sets the value for SslCipher to be an explicit nil

### UnsetSslCipher
`func (o *SystemInfoPostgresql) UnsetSslCipher()`

UnsetSslCipher ensures that no value is present for SslCipher, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


