# SAMLPropertyMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Managed** | Pointer to **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 
**Name** | **string** |  | 
**Expression** | **string** |  | 
**Component** | **string** | Get object&#39;s component so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**SamlName** | **string** |  | 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSAMLPropertyMapping

`func NewSAMLPropertyMapping(pk string, name string, expression string, component string, verboseName string, verboseNamePlural string, metaModelName string, samlName string, ) *SAMLPropertyMapping`

NewSAMLPropertyMapping instantiates a new SAMLPropertyMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLPropertyMappingWithDefaults

`func NewSAMLPropertyMappingWithDefaults() *SAMLPropertyMapping`

NewSAMLPropertyMappingWithDefaults instantiates a new SAMLPropertyMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *SAMLPropertyMapping) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *SAMLPropertyMapping) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *SAMLPropertyMapping) SetPk(v string)`

SetPk sets Pk field to given value.


### GetManaged

`func (o *SAMLPropertyMapping) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *SAMLPropertyMapping) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *SAMLPropertyMapping) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *SAMLPropertyMapping) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *SAMLPropertyMapping) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *SAMLPropertyMapping) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetName

`func (o *SAMLPropertyMapping) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SAMLPropertyMapping) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SAMLPropertyMapping) SetName(v string)`

SetName sets Name field to given value.


### GetExpression

`func (o *SAMLPropertyMapping) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SAMLPropertyMapping) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SAMLPropertyMapping) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetComponent

`func (o *SAMLPropertyMapping) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *SAMLPropertyMapping) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *SAMLPropertyMapping) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *SAMLPropertyMapping) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *SAMLPropertyMapping) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *SAMLPropertyMapping) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *SAMLPropertyMapping) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *SAMLPropertyMapping) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *SAMLPropertyMapping) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *SAMLPropertyMapping) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *SAMLPropertyMapping) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *SAMLPropertyMapping) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetSamlName

`func (o *SAMLPropertyMapping) GetSamlName() string`

GetSamlName returns the SamlName field if non-nil, zero value otherwise.

### GetSamlNameOk

`func (o *SAMLPropertyMapping) GetSamlNameOk() (*string, bool)`

GetSamlNameOk returns a tuple with the SamlName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlName

`func (o *SAMLPropertyMapping) SetSamlName(v string)`

SetSamlName sets SamlName field to given value.


### GetFriendlyName

`func (o *SAMLPropertyMapping) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *SAMLPropertyMapping) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *SAMLPropertyMapping) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *SAMLPropertyMapping) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *SAMLPropertyMapping) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *SAMLPropertyMapping) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


