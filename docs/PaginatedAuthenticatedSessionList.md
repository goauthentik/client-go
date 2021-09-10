# PaginatedAuthenticatedSessionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**PaginatedApplicationListPagination**](PaginatedApplicationListPagination.md) |  | 
**Results** | [**[]AuthenticatedSession**](AuthenticatedSession.md) |  | 

## Methods

### NewPaginatedAuthenticatedSessionList

`func NewPaginatedAuthenticatedSessionList(pagination PaginatedApplicationListPagination, results []AuthenticatedSession, ) *PaginatedAuthenticatedSessionList`

NewPaginatedAuthenticatedSessionList instantiates a new PaginatedAuthenticatedSessionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedAuthenticatedSessionListWithDefaults

`func NewPaginatedAuthenticatedSessionListWithDefaults() *PaginatedAuthenticatedSessionList`

NewPaginatedAuthenticatedSessionListWithDefaults instantiates a new PaginatedAuthenticatedSessionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedAuthenticatedSessionList) GetPagination() PaginatedApplicationListPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedAuthenticatedSessionList) GetPaginationOk() (*PaginatedApplicationListPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedAuthenticatedSessionList) SetPagination(v PaginatedApplicationListPagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedAuthenticatedSessionList) GetResults() []AuthenticatedSession`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedAuthenticatedSessionList) GetResultsOk() (*[]AuthenticatedSession, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedAuthenticatedSessionList) SetResults(v []AuthenticatedSession)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


