# PatchedKubernetesServiceConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Local** | Pointer to **bool** | If enabled, use the local connection. Required Docker socket/Kubernetes Integration | [optional] 
**Kubeconfig** | Pointer to **map[string]interface{}** | Paste your kubeconfig here. authentik will automatically use the currently selected context. | [optional] 
**VerifySsl** | Pointer to **bool** | Verify SSL Certificates of the Kubernetes API endpoint | [optional] 

## Methods

### NewPatchedKubernetesServiceConnectionRequest

`func NewPatchedKubernetesServiceConnectionRequest() *PatchedKubernetesServiceConnectionRequest`

NewPatchedKubernetesServiceConnectionRequest instantiates a new PatchedKubernetesServiceConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedKubernetesServiceConnectionRequestWithDefaults

`func NewPatchedKubernetesServiceConnectionRequestWithDefaults() *PatchedKubernetesServiceConnectionRequest`

NewPatchedKubernetesServiceConnectionRequestWithDefaults instantiates a new PatchedKubernetesServiceConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedKubernetesServiceConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedKubernetesServiceConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedKubernetesServiceConnectionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedKubernetesServiceConnectionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocal

`func (o *PatchedKubernetesServiceConnectionRequest) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *PatchedKubernetesServiceConnectionRequest) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *PatchedKubernetesServiceConnectionRequest) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *PatchedKubernetesServiceConnectionRequest) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetKubeconfig

`func (o *PatchedKubernetesServiceConnectionRequest) GetKubeconfig() map[string]interface{}`

GetKubeconfig returns the Kubeconfig field if non-nil, zero value otherwise.

### GetKubeconfigOk

`func (o *PatchedKubernetesServiceConnectionRequest) GetKubeconfigOk() (*map[string]interface{}, bool)`

GetKubeconfigOk returns a tuple with the Kubeconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeconfig

`func (o *PatchedKubernetesServiceConnectionRequest) SetKubeconfig(v map[string]interface{})`

SetKubeconfig sets Kubeconfig field to given value.

### HasKubeconfig

`func (o *PatchedKubernetesServiceConnectionRequest) HasKubeconfig() bool`

HasKubeconfig returns a boolean if a field has been set.

### GetVerifySsl

`func (o *PatchedKubernetesServiceConnectionRequest) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *PatchedKubernetesServiceConnectionRequest) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *PatchedKubernetesServiceConnectionRequest) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *PatchedKubernetesServiceConnectionRequest) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


