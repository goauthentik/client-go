# KubernetesServiceConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Local** | Pointer to **bool** | If enabled, use the local connection. Required Docker socket/Kubernetes Integration | [optional] 
**Kubeconfig** | Pointer to **map[string]interface{}** | Paste your kubeconfig here. authentik will automatically use the currently selected context. | [optional] 

## Methods

### NewKubernetesServiceConnectionRequest

`func NewKubernetesServiceConnectionRequest(name string, ) *KubernetesServiceConnectionRequest`

NewKubernetesServiceConnectionRequest instantiates a new KubernetesServiceConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesServiceConnectionRequestWithDefaults

`func NewKubernetesServiceConnectionRequestWithDefaults() *KubernetesServiceConnectionRequest`

NewKubernetesServiceConnectionRequestWithDefaults instantiates a new KubernetesServiceConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesServiceConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesServiceConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesServiceConnectionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLocal

`func (o *KubernetesServiceConnectionRequest) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *KubernetesServiceConnectionRequest) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *KubernetesServiceConnectionRequest) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *KubernetesServiceConnectionRequest) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetKubeconfig

`func (o *KubernetesServiceConnectionRequest) GetKubeconfig() map[string]interface{}`

GetKubeconfig returns the Kubeconfig field if non-nil, zero value otherwise.

### GetKubeconfigOk

`func (o *KubernetesServiceConnectionRequest) GetKubeconfigOk() (*map[string]interface{}, bool)`

GetKubeconfigOk returns a tuple with the Kubeconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeconfig

`func (o *KubernetesServiceConnectionRequest) SetKubeconfig(v map[string]interface{})`

SetKubeconfig sets Kubeconfig field to given value.

### HasKubeconfig

`func (o *KubernetesServiceConnectionRequest) HasKubeconfig() bool`

HasKubeconfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


