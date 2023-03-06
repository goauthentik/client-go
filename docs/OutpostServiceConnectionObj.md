# OutpostServiceConnectionObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Local** | Pointer to **bool** | If enabled, use the local connection. Required Docker socket/Kubernetes Integration | [optional] 
**Component** | **string** | Return component used to edit this object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 

## Methods

### NewOutpostServiceConnectionObj

`func NewOutpostServiceConnectionObj(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, ) *OutpostServiceConnectionObj`

NewOutpostServiceConnectionObj instantiates a new OutpostServiceConnectionObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutpostServiceConnectionObjWithDefaults

`func NewOutpostServiceConnectionObjWithDefaults() *OutpostServiceConnectionObj`

NewOutpostServiceConnectionObjWithDefaults instantiates a new OutpostServiceConnectionObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *OutpostServiceConnectionObj) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *OutpostServiceConnectionObj) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *OutpostServiceConnectionObj) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *OutpostServiceConnectionObj) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutpostServiceConnectionObj) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutpostServiceConnectionObj) SetName(v string)`

SetName sets Name field to given value.


### GetLocal

`func (o *OutpostServiceConnectionObj) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *OutpostServiceConnectionObj) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *OutpostServiceConnectionObj) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *OutpostServiceConnectionObj) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetComponent

`func (o *OutpostServiceConnectionObj) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *OutpostServiceConnectionObj) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *OutpostServiceConnectionObj) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *OutpostServiceConnectionObj) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *OutpostServiceConnectionObj) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *OutpostServiceConnectionObj) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *OutpostServiceConnectionObj) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *OutpostServiceConnectionObj) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *OutpostServiceConnectionObj) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *OutpostServiceConnectionObj) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *OutpostServiceConnectionObj) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *OutpostServiceConnectionObj) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


