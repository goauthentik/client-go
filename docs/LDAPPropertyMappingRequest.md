# LDAPPropertyMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Managed** | Pointer to **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 
**Name** | **string** |  | 
**Expression** | **string** |  | 
**ObjectField** | **string** |  | 

## Methods

### NewLDAPPropertyMappingRequest

`func NewLDAPPropertyMappingRequest(name string, expression string, objectField string, ) *LDAPPropertyMappingRequest`

NewLDAPPropertyMappingRequest instantiates a new LDAPPropertyMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPPropertyMappingRequestWithDefaults

`func NewLDAPPropertyMappingRequestWithDefaults() *LDAPPropertyMappingRequest`

NewLDAPPropertyMappingRequestWithDefaults instantiates a new LDAPPropertyMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManaged

`func (o *LDAPPropertyMappingRequest) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *LDAPPropertyMappingRequest) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *LDAPPropertyMappingRequest) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *LDAPPropertyMappingRequest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *LDAPPropertyMappingRequest) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *LDAPPropertyMappingRequest) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetName

`func (o *LDAPPropertyMappingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LDAPPropertyMappingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LDAPPropertyMappingRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExpression

`func (o *LDAPPropertyMappingRequest) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *LDAPPropertyMappingRequest) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *LDAPPropertyMappingRequest) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetObjectField

`func (o *LDAPPropertyMappingRequest) GetObjectField() string`

GetObjectField returns the ObjectField field if non-nil, zero value otherwise.

### GetObjectFieldOk

`func (o *LDAPPropertyMappingRequest) GetObjectFieldOk() (*string, bool)`

GetObjectFieldOk returns a tuple with the ObjectField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectField

`func (o *LDAPPropertyMappingRequest) SetObjectField(v string)`

SetObjectField sets ObjectField field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


