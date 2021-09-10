# PolicyTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **int32** |  | 
**Context** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPolicyTestRequest

`func NewPolicyTestRequest(user int32, ) *PolicyTestRequest`

NewPolicyTestRequest instantiates a new PolicyTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyTestRequestWithDefaults

`func NewPolicyTestRequestWithDefaults() *PolicyTestRequest`

NewPolicyTestRequestWithDefaults instantiates a new PolicyTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *PolicyTestRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PolicyTestRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PolicyTestRequest) SetUser(v int32)`

SetUser sets User field to given value.


### GetContext

`func (o *PolicyTestRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PolicyTestRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PolicyTestRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *PolicyTestRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


