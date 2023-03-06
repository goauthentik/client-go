# KubernetesServiceConnection

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
**Kubeconfig** | Pointer to **map[string]interface{}** | Paste your kubeconfig here. authentik will automatically use the currently selected context. | [optional] 
**VerifySsl** | Pointer to **bool** | Verify SSL Certificates of the Kubernetes API endpoint | [optional] 

## Methods

### NewKubernetesServiceConnection

`func NewKubernetesServiceConnection(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, ) *KubernetesServiceConnection`

NewKubernetesServiceConnection instantiates a new KubernetesServiceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesServiceConnectionWithDefaults

`func NewKubernetesServiceConnectionWithDefaults() *KubernetesServiceConnection`

NewKubernetesServiceConnectionWithDefaults instantiates a new KubernetesServiceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *KubernetesServiceConnection) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *KubernetesServiceConnection) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *KubernetesServiceConnection) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *KubernetesServiceConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesServiceConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesServiceConnection) SetName(v string)`

SetName sets Name field to given value.


### GetLocal

`func (o *KubernetesServiceConnection) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *KubernetesServiceConnection) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *KubernetesServiceConnection) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *KubernetesServiceConnection) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetComponent

`func (o *KubernetesServiceConnection) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *KubernetesServiceConnection) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *KubernetesServiceConnection) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *KubernetesServiceConnection) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *KubernetesServiceConnection) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *KubernetesServiceConnection) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *KubernetesServiceConnection) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *KubernetesServiceConnection) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *KubernetesServiceConnection) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *KubernetesServiceConnection) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *KubernetesServiceConnection) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *KubernetesServiceConnection) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetKubeconfig

`func (o *KubernetesServiceConnection) GetKubeconfig() map[string]interface{}`

GetKubeconfig returns the Kubeconfig field if non-nil, zero value otherwise.

### GetKubeconfigOk

`func (o *KubernetesServiceConnection) GetKubeconfigOk() (*map[string]interface{}, bool)`

GetKubeconfigOk returns a tuple with the Kubeconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeconfig

`func (o *KubernetesServiceConnection) SetKubeconfig(v map[string]interface{})`

SetKubeconfig sets Kubeconfig field to given value.

### HasKubeconfig

`func (o *KubernetesServiceConnection) HasKubeconfig() bool`

HasKubeconfig returns a boolean if a field has been set.

### GetVerifySsl

`func (o *KubernetesServiceConnection) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *KubernetesServiceConnection) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *KubernetesServiceConnection) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *KubernetesServiceConnection) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


