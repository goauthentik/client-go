# RACPropertyMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Managed** | Pointer to **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 
**Name** | **string** |  | 
**Expression** | Pointer to **string** |  | [optional] 
**StaticSettings** | **map[string]interface{}** |  | 

## Methods

### NewRACPropertyMappingRequest

`func NewRACPropertyMappingRequest(name string, staticSettings map[string]interface{}, ) *RACPropertyMappingRequest`

NewRACPropertyMappingRequest instantiates a new RACPropertyMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRACPropertyMappingRequestWithDefaults

`func NewRACPropertyMappingRequestWithDefaults() *RACPropertyMappingRequest`

NewRACPropertyMappingRequestWithDefaults instantiates a new RACPropertyMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManaged

`func (o *RACPropertyMappingRequest) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *RACPropertyMappingRequest) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *RACPropertyMappingRequest) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *RACPropertyMappingRequest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *RACPropertyMappingRequest) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *RACPropertyMappingRequest) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetName

`func (o *RACPropertyMappingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RACPropertyMappingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RACPropertyMappingRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExpression

`func (o *RACPropertyMappingRequest) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *RACPropertyMappingRequest) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *RACPropertyMappingRequest) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *RACPropertyMappingRequest) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetStaticSettings

`func (o *RACPropertyMappingRequest) GetStaticSettings() map[string]interface{}`

GetStaticSettings returns the StaticSettings field if non-nil, zero value otherwise.

### GetStaticSettingsOk

`func (o *RACPropertyMappingRequest) GetStaticSettingsOk() (*map[string]interface{}, bool)`

GetStaticSettingsOk returns a tuple with the StaticSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticSettings

`func (o *RACPropertyMappingRequest) SetStaticSettings(v map[string]interface{})`

SetStaticSettings sets StaticSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


