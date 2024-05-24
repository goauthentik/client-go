# PropertyMappingTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **int32** |  | [optional] 
**Context** | Pointer to **map[string]interface{}** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 

## Methods

### NewPropertyMappingTestRequest

`func NewPropertyMappingTestRequest() *PropertyMappingTestRequest`

NewPropertyMappingTestRequest instantiates a new PropertyMappingTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyMappingTestRequestWithDefaults

`func NewPropertyMappingTestRequestWithDefaults() *PropertyMappingTestRequest`

NewPropertyMappingTestRequestWithDefaults instantiates a new PropertyMappingTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *PropertyMappingTestRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PropertyMappingTestRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PropertyMappingTestRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *PropertyMappingTestRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetContext

`func (o *PropertyMappingTestRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PropertyMappingTestRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PropertyMappingTestRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *PropertyMappingTestRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetGroup

`func (o *PropertyMappingTestRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PropertyMappingTestRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PropertyMappingTestRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PropertyMappingTestRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


