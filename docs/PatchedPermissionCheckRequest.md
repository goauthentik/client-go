# PatchedPermissionCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Expression** | Pointer to **string** |  | [optional] 
**Endpoints** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPatchedPermissionCheckRequest

`func NewPatchedPermissionCheckRequest() *PatchedPermissionCheckRequest`

NewPatchedPermissionCheckRequest instantiates a new PatchedPermissionCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedPermissionCheckRequestWithDefaults

`func NewPatchedPermissionCheckRequestWithDefaults() *PatchedPermissionCheckRequest`

NewPatchedPermissionCheckRequestWithDefaults instantiates a new PatchedPermissionCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedPermissionCheckRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedPermissionCheckRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedPermissionCheckRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedPermissionCheckRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpression

`func (o *PatchedPermissionCheckRequest) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *PatchedPermissionCheckRequest) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *PatchedPermissionCheckRequest) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *PatchedPermissionCheckRequest) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetEndpoints

`func (o *PatchedPermissionCheckRequest) GetEndpoints() []string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *PatchedPermissionCheckRequest) GetEndpointsOk() (*[]string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *PatchedPermissionCheckRequest) SetEndpoints(v []string)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *PatchedPermissionCheckRequest) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


