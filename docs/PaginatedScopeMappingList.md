# PaginatedScopeMappingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**PaginatedApplicationListPagination**](PaginatedApplicationListPagination.md) |  | 
**Results** | [**[]ScopeMapping**](ScopeMapping.md) |  | 

## Methods

### NewPaginatedScopeMappingList

`func NewPaginatedScopeMappingList(pagination PaginatedApplicationListPagination, results []ScopeMapping, ) *PaginatedScopeMappingList`

NewPaginatedScopeMappingList instantiates a new PaginatedScopeMappingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedScopeMappingListWithDefaults

`func NewPaginatedScopeMappingListWithDefaults() *PaginatedScopeMappingList`

NewPaginatedScopeMappingListWithDefaults instantiates a new PaginatedScopeMappingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedScopeMappingList) GetPagination() PaginatedApplicationListPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedScopeMappingList) GetPaginationOk() (*PaginatedApplicationListPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedScopeMappingList) SetPagination(v PaginatedApplicationListPagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedScopeMappingList) GetResults() []ScopeMapping`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedScopeMappingList) GetResultsOk() (*[]ScopeMapping, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedScopeMappingList) SetResults(v []ScopeMapping)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


