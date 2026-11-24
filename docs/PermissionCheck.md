# PermissionCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Expression** | **string** |  | 
**Endpoints** | **[]string** |  | 
**UnknownEndpoints** | **[]string** |  | [readonly] 
**UnknownFields** | **map[string][]string** |  | [readonly] 

## Methods

### NewPermissionCheck

`func NewPermissionCheck(pk string, name string, expression string, endpoints []string, unknownEndpoints []string, unknownFields map[string][]string, ) *PermissionCheck`

NewPermissionCheck instantiates a new PermissionCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionCheckWithDefaults

`func NewPermissionCheckWithDefaults() *PermissionCheck`

NewPermissionCheckWithDefaults instantiates a new PermissionCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *PermissionCheck) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *PermissionCheck) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *PermissionCheck) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *PermissionCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionCheck) SetName(v string)`

SetName sets Name field to given value.


### GetExpression

`func (o *PermissionCheck) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *PermissionCheck) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *PermissionCheck) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetEndpoints

`func (o *PermissionCheck) GetEndpoints() []string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *PermissionCheck) GetEndpointsOk() (*[]string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *PermissionCheck) SetEndpoints(v []string)`

SetEndpoints sets Endpoints field to given value.


### GetUnknownEndpoints

`func (o *PermissionCheck) GetUnknownEndpoints() []string`

GetUnknownEndpoints returns the UnknownEndpoints field if non-nil, zero value otherwise.

### GetUnknownEndpointsOk

`func (o *PermissionCheck) GetUnknownEndpointsOk() (*[]string, bool)`

GetUnknownEndpointsOk returns a tuple with the UnknownEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownEndpoints

`func (o *PermissionCheck) SetUnknownEndpoints(v []string)`

SetUnknownEndpoints sets UnknownEndpoints field to given value.


### GetUnknownFields

`func (o *PermissionCheck) GetUnknownFields() map[string][]string`

GetUnknownFields returns the UnknownFields field if non-nil, zero value otherwise.

### GetUnknownFieldsOk

`func (o *PermissionCheck) GetUnknownFieldsOk() (*map[string][]string, bool)`

GetUnknownFieldsOk returns a tuple with the UnknownFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownFields

`func (o *PermissionCheck) SetUnknownFields(v map[string][]string)`

SetUnknownFields sets UnknownFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


