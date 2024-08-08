# PaginatedGroupPlexSourceConnectionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]GroupPlexSourceConnection**](GroupPlexSourceConnection.md) |  | 

## Methods

### NewPaginatedGroupPlexSourceConnectionList

`func NewPaginatedGroupPlexSourceConnectionList(pagination Pagination, results []GroupPlexSourceConnection, ) *PaginatedGroupPlexSourceConnectionList`

NewPaginatedGroupPlexSourceConnectionList instantiates a new PaginatedGroupPlexSourceConnectionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedGroupPlexSourceConnectionListWithDefaults

`func NewPaginatedGroupPlexSourceConnectionListWithDefaults() *PaginatedGroupPlexSourceConnectionList`

NewPaginatedGroupPlexSourceConnectionListWithDefaults instantiates a new PaginatedGroupPlexSourceConnectionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedGroupPlexSourceConnectionList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedGroupPlexSourceConnectionList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedGroupPlexSourceConnectionList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedGroupPlexSourceConnectionList) GetResults() []GroupPlexSourceConnection`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedGroupPlexSourceConnectionList) GetResultsOk() (*[]GroupPlexSourceConnection, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedGroupPlexSourceConnectionList) SetResults(v []GroupPlexSourceConnection)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


