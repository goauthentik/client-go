# MutualTLSStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Component** | **string** | Get object type so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**FlowSet** | Pointer to [**[]FlowSet**](FlowSet.md) |  | [optional] 
**Mode** | [**StageModeEnum**](StageModeEnum.md) |  | 
**CertificateAuthorities** | Pointer to **[]string** | Configure certificate authorities to validate the certificate against. This option has a higher priority than the &#x60;client_certificate&#x60; option on &#x60;Brand&#x60;. | [optional] 
**CertAttribute** | [**CertAttributeEnum**](CertAttributeEnum.md) |  | 
**UserAttribute** | [**UserAttributeEnum**](UserAttributeEnum.md) |  | 

## Methods

### NewMutualTLSStage

`func NewMutualTLSStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, mode StageModeEnum, certAttribute CertAttributeEnum, userAttribute UserAttributeEnum, ) *MutualTLSStage`

NewMutualTLSStage instantiates a new MutualTLSStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualTLSStageWithDefaults

`func NewMutualTLSStageWithDefaults() *MutualTLSStage`

NewMutualTLSStageWithDefaults instantiates a new MutualTLSStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *MutualTLSStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *MutualTLSStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *MutualTLSStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *MutualTLSStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MutualTLSStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MutualTLSStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *MutualTLSStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *MutualTLSStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *MutualTLSStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *MutualTLSStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *MutualTLSStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *MutualTLSStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *MutualTLSStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *MutualTLSStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *MutualTLSStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *MutualTLSStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *MutualTLSStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *MutualTLSStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *MutualTLSStage) GetFlowSet() []FlowSet`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *MutualTLSStage) GetFlowSetOk() (*[]FlowSet, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *MutualTLSStage) SetFlowSet(v []FlowSet)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *MutualTLSStage) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetMode

`func (o *MutualTLSStage) GetMode() StageModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MutualTLSStage) GetModeOk() (*StageModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MutualTLSStage) SetMode(v StageModeEnum)`

SetMode sets Mode field to given value.


### GetCertificateAuthorities

`func (o *MutualTLSStage) GetCertificateAuthorities() []string`

GetCertificateAuthorities returns the CertificateAuthorities field if non-nil, zero value otherwise.

### GetCertificateAuthoritiesOk

`func (o *MutualTLSStage) GetCertificateAuthoritiesOk() (*[]string, bool)`

GetCertificateAuthoritiesOk returns a tuple with the CertificateAuthorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorities

`func (o *MutualTLSStage) SetCertificateAuthorities(v []string)`

SetCertificateAuthorities sets CertificateAuthorities field to given value.

### HasCertificateAuthorities

`func (o *MutualTLSStage) HasCertificateAuthorities() bool`

HasCertificateAuthorities returns a boolean if a field has been set.

### GetCertAttribute

`func (o *MutualTLSStage) GetCertAttribute() CertAttributeEnum`

GetCertAttribute returns the CertAttribute field if non-nil, zero value otherwise.

### GetCertAttributeOk

`func (o *MutualTLSStage) GetCertAttributeOk() (*CertAttributeEnum, bool)`

GetCertAttributeOk returns a tuple with the CertAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertAttribute

`func (o *MutualTLSStage) SetCertAttribute(v CertAttributeEnum)`

SetCertAttribute sets CertAttribute field to given value.


### GetUserAttribute

`func (o *MutualTLSStage) GetUserAttribute() UserAttributeEnum`

GetUserAttribute returns the UserAttribute field if non-nil, zero value otherwise.

### GetUserAttributeOk

`func (o *MutualTLSStage) GetUserAttributeOk() (*UserAttributeEnum, bool)`

GetUserAttributeOk returns a tuple with the UserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttribute

`func (o *MutualTLSStage) SetUserAttribute(v UserAttributeEnum)`

SetUserAttribute sets UserAttribute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


