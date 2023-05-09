# PatchedLDAPPropertyMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Managed** | Pointer to **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Expression** | Pointer to **string** |  | [optional] 
**ObjectField** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedLDAPPropertyMappingRequest

`func NewPatchedLDAPPropertyMappingRequest() *PatchedLDAPPropertyMappingRequest`

NewPatchedLDAPPropertyMappingRequest instantiates a new PatchedLDAPPropertyMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedLDAPPropertyMappingRequestWithDefaults

`func NewPatchedLDAPPropertyMappingRequestWithDefaults() *PatchedLDAPPropertyMappingRequest`

NewPatchedLDAPPropertyMappingRequestWithDefaults instantiates a new PatchedLDAPPropertyMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManaged

`func (o *PatchedLDAPPropertyMappingRequest) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *PatchedLDAPPropertyMappingRequest) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *PatchedLDAPPropertyMappingRequest) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *PatchedLDAPPropertyMappingRequest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *PatchedLDAPPropertyMappingRequest) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *PatchedLDAPPropertyMappingRequest) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetName

`func (o *PatchedLDAPPropertyMappingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedLDAPPropertyMappingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedLDAPPropertyMappingRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedLDAPPropertyMappingRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpression

`func (o *PatchedLDAPPropertyMappingRequest) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *PatchedLDAPPropertyMappingRequest) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *PatchedLDAPPropertyMappingRequest) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *PatchedLDAPPropertyMappingRequest) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetObjectField

`func (o *PatchedLDAPPropertyMappingRequest) GetObjectField() string`

GetObjectField returns the ObjectField field if non-nil, zero value otherwise.

### GetObjectFieldOk

`func (o *PatchedLDAPPropertyMappingRequest) GetObjectFieldOk() (*string, bool)`

GetObjectFieldOk returns a tuple with the ObjectField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectField

`func (o *PatchedLDAPPropertyMappingRequest) SetObjectField(v string)`

SetObjectField sets ObjectField field to given value.

### HasObjectField

`func (o *PatchedLDAPPropertyMappingRequest) HasObjectField() bool`

HasObjectField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


