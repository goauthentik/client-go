# LDAPPropertyMapping

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
**ObjectField** | **string** |  | 

## Methods

### NewLDAPPropertyMapping

`func NewLDAPPropertyMapping(pk string, name string, expression string, component string, verboseName string, verboseNamePlural string, metaModelName string, objectField string, ) *LDAPPropertyMapping`

NewLDAPPropertyMapping instantiates a new LDAPPropertyMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPPropertyMappingWithDefaults

`func NewLDAPPropertyMappingWithDefaults() *LDAPPropertyMapping`

NewLDAPPropertyMappingWithDefaults instantiates a new LDAPPropertyMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *LDAPPropertyMapping) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *LDAPPropertyMapping) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *LDAPPropertyMapping) SetPk(v string)`

SetPk sets Pk field to given value.


### GetManaged

`func (o *LDAPPropertyMapping) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *LDAPPropertyMapping) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *LDAPPropertyMapping) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *LDAPPropertyMapping) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *LDAPPropertyMapping) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *LDAPPropertyMapping) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetName

`func (o *LDAPPropertyMapping) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LDAPPropertyMapping) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LDAPPropertyMapping) SetName(v string)`

SetName sets Name field to given value.


### GetExpression

`func (o *LDAPPropertyMapping) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *LDAPPropertyMapping) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *LDAPPropertyMapping) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetComponent

`func (o *LDAPPropertyMapping) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *LDAPPropertyMapping) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *LDAPPropertyMapping) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *LDAPPropertyMapping) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *LDAPPropertyMapping) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *LDAPPropertyMapping) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *LDAPPropertyMapping) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *LDAPPropertyMapping) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *LDAPPropertyMapping) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *LDAPPropertyMapping) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *LDAPPropertyMapping) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *LDAPPropertyMapping) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetObjectField

`func (o *LDAPPropertyMapping) GetObjectField() string`

GetObjectField returns the ObjectField field if non-nil, zero value otherwise.

### GetObjectFieldOk

`func (o *LDAPPropertyMapping) GetObjectFieldOk() (*string, bool)`

GetObjectFieldOk returns a tuple with the ObjectField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectField

`func (o *LDAPPropertyMapping) SetObjectField(v string)`

SetObjectField sets ObjectField field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


