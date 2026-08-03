# LDAPCheckAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasSearchPermission** | Pointer to **bool** |  | [optional] 
**Access** | [**PolicyTestResult**](PolicyTestResult.md) |  | 

## Methods

### NewLDAPCheckAccess

`func NewLDAPCheckAccess(access PolicyTestResult, ) *LDAPCheckAccess`

NewLDAPCheckAccess instantiates a new LDAPCheckAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPCheckAccessWithDefaults

`func NewLDAPCheckAccessWithDefaults() *LDAPCheckAccess`

NewLDAPCheckAccessWithDefaults instantiates a new LDAPCheckAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasSearchPermission

`func (o *LDAPCheckAccess) GetHasSearchPermission() bool`

GetHasSearchPermission returns the HasSearchPermission field if non-nil, zero value otherwise.

### GetHasSearchPermissionOk

`func (o *LDAPCheckAccess) GetHasSearchPermissionOk() (*bool, bool)`

GetHasSearchPermissionOk returns a tuple with the HasSearchPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSearchPermission

`func (o *LDAPCheckAccess) SetHasSearchPermission(v bool)`

SetHasSearchPermission sets HasSearchPermission field to given value.

### HasHasSearchPermission

`func (o *LDAPCheckAccess) HasHasSearchPermission() bool`

HasHasSearchPermission returns a boolean if a field has been set.

### GetAccess

`func (o *LDAPCheckAccess) GetAccess() PolicyTestResult`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *LDAPCheckAccess) GetAccessOk() (*PolicyTestResult, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *LDAPCheckAccess) SetAccess(v PolicyTestResult)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


