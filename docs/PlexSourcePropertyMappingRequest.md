# PlexSourcePropertyMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Managed** | Pointer to **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 
**Name** | **string** |  | 
**Expression** | **string** |  | 

## Methods

### NewPlexSourcePropertyMappingRequest

`func NewPlexSourcePropertyMappingRequest(name string, expression string, ) *PlexSourcePropertyMappingRequest`

NewPlexSourcePropertyMappingRequest instantiates a new PlexSourcePropertyMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlexSourcePropertyMappingRequestWithDefaults

`func NewPlexSourcePropertyMappingRequestWithDefaults() *PlexSourcePropertyMappingRequest`

NewPlexSourcePropertyMappingRequestWithDefaults instantiates a new PlexSourcePropertyMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManaged

`func (o *PlexSourcePropertyMappingRequest) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *PlexSourcePropertyMappingRequest) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *PlexSourcePropertyMappingRequest) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *PlexSourcePropertyMappingRequest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *PlexSourcePropertyMappingRequest) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *PlexSourcePropertyMappingRequest) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetName

`func (o *PlexSourcePropertyMappingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlexSourcePropertyMappingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlexSourcePropertyMappingRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExpression

`func (o *PlexSourcePropertyMappingRequest) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *PlexSourcePropertyMappingRequest) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *PlexSourcePropertyMappingRequest) SetExpression(v string)`

SetExpression sets Expression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


