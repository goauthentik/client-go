# PermissionCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Expression** | **string** |  | 
**Endpoints** | **[]string** |  | 

## Methods

### NewPermissionCheckRequest

`func NewPermissionCheckRequest(name string, expression string, endpoints []string, ) *PermissionCheckRequest`

NewPermissionCheckRequest instantiates a new PermissionCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionCheckRequestWithDefaults

`func NewPermissionCheckRequestWithDefaults() *PermissionCheckRequest`

NewPermissionCheckRequestWithDefaults instantiates a new PermissionCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PermissionCheckRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionCheckRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionCheckRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExpression

`func (o *PermissionCheckRequest) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *PermissionCheckRequest) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *PermissionCheckRequest) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetEndpoints

`func (o *PermissionCheckRequest) GetEndpoints() []string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *PermissionCheckRequest) GetEndpointsOk() (*[]string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *PermissionCheckRequest) SetEndpoints(v []string)`

SetEndpoints sets Endpoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


