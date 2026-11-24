# DockerServiceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Local** | Pointer to **bool** | If enabled, use the local connection. Required Docker socket/Kubernetes Integration | [optional] 
**Component** | **string** |  | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**Url** | **string** | Can be in the format of &#39;unix://&lt;path&gt;&#39; when connecting to a local docker daemon, or &#39;https://&lt;hostname&gt;:2376&#39; when connecting to a remote system. | 
**TlsVerification** | Pointer to **NullableString** | CA which the endpoint&#39;s Certificate is verified against. Can be left empty for no validation. | [optional] 
**TlsAuthentication** | Pointer to **NullableString** | Certificate/Key used for authentication. Can be left empty for no authentication. | [optional] 

## Methods

### NewDockerServiceConnection

`func NewDockerServiceConnection(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, url string, ) *DockerServiceConnection`

NewDockerServiceConnection instantiates a new DockerServiceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerServiceConnectionWithDefaults

`func NewDockerServiceConnectionWithDefaults() *DockerServiceConnection`

NewDockerServiceConnectionWithDefaults instantiates a new DockerServiceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *DockerServiceConnection) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *DockerServiceConnection) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *DockerServiceConnection) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *DockerServiceConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DockerServiceConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DockerServiceConnection) SetName(v string)`

SetName sets Name field to given value.


### GetLocal

`func (o *DockerServiceConnection) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *DockerServiceConnection) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *DockerServiceConnection) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *DockerServiceConnection) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetComponent

`func (o *DockerServiceConnection) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *DockerServiceConnection) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *DockerServiceConnection) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *DockerServiceConnection) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *DockerServiceConnection) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *DockerServiceConnection) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *DockerServiceConnection) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *DockerServiceConnection) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *DockerServiceConnection) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *DockerServiceConnection) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *DockerServiceConnection) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *DockerServiceConnection) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetUrl

`func (o *DockerServiceConnection) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DockerServiceConnection) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DockerServiceConnection) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTlsVerification

`func (o *DockerServiceConnection) GetTlsVerification() string`

GetTlsVerification returns the TlsVerification field if non-nil, zero value otherwise.

### GetTlsVerificationOk

`func (o *DockerServiceConnection) GetTlsVerificationOk() (*string, bool)`

GetTlsVerificationOk returns a tuple with the TlsVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVerification

`func (o *DockerServiceConnection) SetTlsVerification(v string)`

SetTlsVerification sets TlsVerification field to given value.

### HasTlsVerification

`func (o *DockerServiceConnection) HasTlsVerification() bool`

HasTlsVerification returns a boolean if a field has been set.

### SetTlsVerificationNil

`func (o *DockerServiceConnection) SetTlsVerificationNil(b bool)`

 SetTlsVerificationNil sets the value for TlsVerification to be an explicit nil

### UnsetTlsVerification
`func (o *DockerServiceConnection) UnsetTlsVerification()`

UnsetTlsVerification ensures that no value is present for TlsVerification, not even an explicit nil
### GetTlsAuthentication

`func (o *DockerServiceConnection) GetTlsAuthentication() string`

GetTlsAuthentication returns the TlsAuthentication field if non-nil, zero value otherwise.

### GetTlsAuthenticationOk

`func (o *DockerServiceConnection) GetTlsAuthenticationOk() (*string, bool)`

GetTlsAuthenticationOk returns a tuple with the TlsAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAuthentication

`func (o *DockerServiceConnection) SetTlsAuthentication(v string)`

SetTlsAuthentication sets TlsAuthentication field to given value.

### HasTlsAuthentication

`func (o *DockerServiceConnection) HasTlsAuthentication() bool`

HasTlsAuthentication returns a boolean if a field has been set.

### SetTlsAuthenticationNil

`func (o *DockerServiceConnection) SetTlsAuthenticationNil(b bool)`

 SetTlsAuthenticationNil sets the value for TlsAuthentication to be an explicit nil

### UnsetTlsAuthentication
`func (o *DockerServiceConnection) UnsetTlsAuthentication()`

UnsetTlsAuthentication ensures that no value is present for TlsAuthentication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


