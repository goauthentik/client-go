# PatchedRACPropertyMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Managed** | Pointer to **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Expression** | Pointer to **string** |  | [optional] 
**StaticSettings** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedRACPropertyMappingRequest

`func NewPatchedRACPropertyMappingRequest() *PatchedRACPropertyMappingRequest`

NewPatchedRACPropertyMappingRequest instantiates a new PatchedRACPropertyMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedRACPropertyMappingRequestWithDefaults

`func NewPatchedRACPropertyMappingRequestWithDefaults() *PatchedRACPropertyMappingRequest`

NewPatchedRACPropertyMappingRequestWithDefaults instantiates a new PatchedRACPropertyMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManaged

`func (o *PatchedRACPropertyMappingRequest) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *PatchedRACPropertyMappingRequest) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *PatchedRACPropertyMappingRequest) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *PatchedRACPropertyMappingRequest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *PatchedRACPropertyMappingRequest) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *PatchedRACPropertyMappingRequest) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetName

`func (o *PatchedRACPropertyMappingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedRACPropertyMappingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedRACPropertyMappingRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedRACPropertyMappingRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpression

`func (o *PatchedRACPropertyMappingRequest) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *PatchedRACPropertyMappingRequest) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *PatchedRACPropertyMappingRequest) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *PatchedRACPropertyMappingRequest) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetStaticSettings

`func (o *PatchedRACPropertyMappingRequest) GetStaticSettings() map[string]interface{}`

GetStaticSettings returns the StaticSettings field if non-nil, zero value otherwise.

### GetStaticSettingsOk

`func (o *PatchedRACPropertyMappingRequest) GetStaticSettingsOk() (*map[string]interface{}, bool)`

GetStaticSettingsOk returns a tuple with the StaticSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticSettings

`func (o *PatchedRACPropertyMappingRequest) SetStaticSettings(v map[string]interface{})`

SetStaticSettings sets StaticSettings field to given value.

### HasStaticSettings

`func (o *PatchedRACPropertyMappingRequest) HasStaticSettings() bool`

HasStaticSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


