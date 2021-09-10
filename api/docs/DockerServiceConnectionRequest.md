# DockerServiceConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Local** | Pointer to **bool** | If enabled, use the local connection. Required Docker socket/Kubernetes Integration | [optional] 
**Url** | **string** | Can be in the format of &#39;unix://&lt;path&gt;&#39; when connecting to a local docker daemon, or &#39;https://&lt;hostname&gt;:2376&#39; when connecting to a remote system. | 
**TlsVerification** | Pointer to **NullableString** | CA which the endpoint&#39;s Certificate is verified against. Can be left empty for no validation. | [optional] 
**TlsAuthentication** | Pointer to **NullableString** | Certificate/Key used for authentication. Can be left empty for no authentication. | [optional] 

## Methods

### NewDockerServiceConnectionRequest

`func NewDockerServiceConnectionRequest(name string, url string, ) *DockerServiceConnectionRequest`

NewDockerServiceConnectionRequest instantiates a new DockerServiceConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerServiceConnectionRequestWithDefaults

`func NewDockerServiceConnectionRequestWithDefaults() *DockerServiceConnectionRequest`

NewDockerServiceConnectionRequestWithDefaults instantiates a new DockerServiceConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DockerServiceConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DockerServiceConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DockerServiceConnectionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLocal

`func (o *DockerServiceConnectionRequest) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *DockerServiceConnectionRequest) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *DockerServiceConnectionRequest) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *DockerServiceConnectionRequest) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetUrl

`func (o *DockerServiceConnectionRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DockerServiceConnectionRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DockerServiceConnectionRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTlsVerification

`func (o *DockerServiceConnectionRequest) GetTlsVerification() string`

GetTlsVerification returns the TlsVerification field if non-nil, zero value otherwise.

### GetTlsVerificationOk

`func (o *DockerServiceConnectionRequest) GetTlsVerificationOk() (*string, bool)`

GetTlsVerificationOk returns a tuple with the TlsVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVerification

`func (o *DockerServiceConnectionRequest) SetTlsVerification(v string)`

SetTlsVerification sets TlsVerification field to given value.

### HasTlsVerification

`func (o *DockerServiceConnectionRequest) HasTlsVerification() bool`

HasTlsVerification returns a boolean if a field has been set.

### SetTlsVerificationNil

`func (o *DockerServiceConnectionRequest) SetTlsVerificationNil(b bool)`

 SetTlsVerificationNil sets the value for TlsVerification to be an explicit nil

### UnsetTlsVerification
`func (o *DockerServiceConnectionRequest) UnsetTlsVerification()`

UnsetTlsVerification ensures that no value is present for TlsVerification, not even an explicit nil
### GetTlsAuthentication

`func (o *DockerServiceConnectionRequest) GetTlsAuthentication() string`

GetTlsAuthentication returns the TlsAuthentication field if non-nil, zero value otherwise.

### GetTlsAuthenticationOk

`func (o *DockerServiceConnectionRequest) GetTlsAuthenticationOk() (*string, bool)`

GetTlsAuthenticationOk returns a tuple with the TlsAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAuthentication

`func (o *DockerServiceConnectionRequest) SetTlsAuthentication(v string)`

SetTlsAuthentication sets TlsAuthentication field to given value.

### HasTlsAuthentication

`func (o *DockerServiceConnectionRequest) HasTlsAuthentication() bool`

HasTlsAuthentication returns a boolean if a field has been set.

### SetTlsAuthenticationNil

`func (o *DockerServiceConnectionRequest) SetTlsAuthenticationNil(b bool)`

 SetTlsAuthenticationNil sets the value for TlsAuthentication to be an explicit nil

### UnsetTlsAuthentication
`func (o *DockerServiceConnectionRequest) UnsetTlsAuthentication()`

UnsetTlsAuthentication ensures that no value is present for TlsAuthentication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


