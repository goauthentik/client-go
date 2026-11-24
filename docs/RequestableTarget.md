# RequestableTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**PbmUuid** | **string** |  | [readonly] 
**Label** | **string** |  | [readonly] 
**Parent** | [**NullableApplication**](Application.md) |  | [readonly] 

## Methods

### NewRequestableTarget

`func NewRequestableTarget(verboseName string, verboseNamePlural string, metaModelName string, pbmUuid string, label string, parent NullableApplication, ) *RequestableTarget`

NewRequestableTarget instantiates a new RequestableTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestableTargetWithDefaults

`func NewRequestableTargetWithDefaults() *RequestableTarget`

NewRequestableTargetWithDefaults instantiates a new RequestableTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerboseName

`func (o *RequestableTarget) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *RequestableTarget) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *RequestableTarget) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *RequestableTarget) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *RequestableTarget) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *RequestableTarget) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *RequestableTarget) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *RequestableTarget) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *RequestableTarget) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetPbmUuid

`func (o *RequestableTarget) GetPbmUuid() string`

GetPbmUuid returns the PbmUuid field if non-nil, zero value otherwise.

### GetPbmUuidOk

`func (o *RequestableTarget) GetPbmUuidOk() (*string, bool)`

GetPbmUuidOk returns a tuple with the PbmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbmUuid

`func (o *RequestableTarget) SetPbmUuid(v string)`

SetPbmUuid sets PbmUuid field to given value.


### GetLabel

`func (o *RequestableTarget) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RequestableTarget) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RequestableTarget) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetParent

`func (o *RequestableTarget) GetParent() Application`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *RequestableTarget) GetParentOk() (*Application, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *RequestableTarget) SetParent(v Application)`

SetParent sets Parent field to given value.


### SetParentNil

`func (o *RequestableTarget) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *RequestableTarget) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


