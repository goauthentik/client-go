# PatchedDockerServiceConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Local** | Pointer to **bool** | If enabled, use the local connection. Required Docker socket/Kubernetes Integration | [optional] 
**Url** | Pointer to **string** | Can be in the format of &#39;unix://&lt;path&gt;&#39; when connecting to a local docker daemon, or &#39;https://&lt;hostname&gt;:2376&#39; when connecting to a remote system. | [optional] 
**TlsVerification** | Pointer to **NullableString** | CA which the endpoint&#39;s Certificate is verified against. Can be left empty for no validation. | [optional] 
**TlsAuthentication** | Pointer to **NullableString** | Certificate/Key used for authentication. Can be left empty for no authentication. | [optional] 

## Methods

### NewPatchedDockerServiceConnectionRequest

`func NewPatchedDockerServiceConnectionRequest() *PatchedDockerServiceConnectionRequest`

NewPatchedDockerServiceConnectionRequest instantiates a new PatchedDockerServiceConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDockerServiceConnectionRequestWithDefaults

`func NewPatchedDockerServiceConnectionRequestWithDefaults() *PatchedDockerServiceConnectionRequest`

NewPatchedDockerServiceConnectionRequestWithDefaults instantiates a new PatchedDockerServiceConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedDockerServiceConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedDockerServiceConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedDockerServiceConnectionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedDockerServiceConnectionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocal

`func (o *PatchedDockerServiceConnectionRequest) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *PatchedDockerServiceConnectionRequest) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *PatchedDockerServiceConnectionRequest) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *PatchedDockerServiceConnectionRequest) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedDockerServiceConnectionRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedDockerServiceConnectionRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedDockerServiceConnectionRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedDockerServiceConnectionRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTlsVerification

`func (o *PatchedDockerServiceConnectionRequest) GetTlsVerification() string`

GetTlsVerification returns the TlsVerification field if non-nil, zero value otherwise.

### GetTlsVerificationOk

`func (o *PatchedDockerServiceConnectionRequest) GetTlsVerificationOk() (*string, bool)`

GetTlsVerificationOk returns a tuple with the TlsVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVerification

`func (o *PatchedDockerServiceConnectionRequest) SetTlsVerification(v string)`

SetTlsVerification sets TlsVerification field to given value.

### HasTlsVerification

`func (o *PatchedDockerServiceConnectionRequest) HasTlsVerification() bool`

HasTlsVerification returns a boolean if a field has been set.

### SetTlsVerificationNil

`func (o *PatchedDockerServiceConnectionRequest) SetTlsVerificationNil(b bool)`

 SetTlsVerificationNil sets the value for TlsVerification to be an explicit nil

### UnsetTlsVerification
`func (o *PatchedDockerServiceConnectionRequest) UnsetTlsVerification()`

UnsetTlsVerification ensures that no value is present for TlsVerification, not even an explicit nil
### GetTlsAuthentication

`func (o *PatchedDockerServiceConnectionRequest) GetTlsAuthentication() string`

GetTlsAuthentication returns the TlsAuthentication field if non-nil, zero value otherwise.

### GetTlsAuthenticationOk

`func (o *PatchedDockerServiceConnectionRequest) GetTlsAuthenticationOk() (*string, bool)`

GetTlsAuthenticationOk returns a tuple with the TlsAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAuthentication

`func (o *PatchedDockerServiceConnectionRequest) SetTlsAuthentication(v string)`

SetTlsAuthentication sets TlsAuthentication field to given value.

### HasTlsAuthentication

`func (o *PatchedDockerServiceConnectionRequest) HasTlsAuthentication() bool`

HasTlsAuthentication returns a boolean if a field has been set.

### SetTlsAuthenticationNil

`func (o *PatchedDockerServiceConnectionRequest) SetTlsAuthenticationNil(b bool)`

 SetTlsAuthenticationNil sets the value for TlsAuthentication to be an explicit nil

### UnsetTlsAuthentication
`func (o *PatchedDockerServiceConnectionRequest) UnsetTlsAuthentication()`

UnsetTlsAuthentication ensures that no value is present for TlsAuthentication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


