# PaginatedRoleAssignedObjectPermissionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]RoleAssignedObjectPermission**](RoleAssignedObjectPermission.md) |  | 

## Methods

### NewPaginatedRoleAssignedObjectPermissionList

`func NewPaginatedRoleAssignedObjectPermissionList(pagination Pagination, results []RoleAssignedObjectPermission, ) *PaginatedRoleAssignedObjectPermissionList`

NewPaginatedRoleAssignedObjectPermissionList instantiates a new PaginatedRoleAssignedObjectPermissionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedRoleAssignedObjectPermissionListWithDefaults

`func NewPaginatedRoleAssignedObjectPermissionListWithDefaults() *PaginatedRoleAssignedObjectPermissionList`

NewPaginatedRoleAssignedObjectPermissionListWithDefaults instantiates a new PaginatedRoleAssignedObjectPermissionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedRoleAssignedObjectPermissionList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedRoleAssignedObjectPermissionList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedRoleAssignedObjectPermissionList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedRoleAssignedObjectPermissionList) GetResults() []RoleAssignedObjectPermission`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedRoleAssignedObjectPermissionList) GetResultsOk() (*[]RoleAssignedObjectPermission, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedRoleAssignedObjectPermissionList) SetResults(v []RoleAssignedObjectPermission)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


