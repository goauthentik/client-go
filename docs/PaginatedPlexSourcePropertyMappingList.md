# PaginatedPlexSourcePropertyMappingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]PlexSourcePropertyMapping**](PlexSourcePropertyMapping.md) |  | 

## Methods

### NewPaginatedPlexSourcePropertyMappingList

`func NewPaginatedPlexSourcePropertyMappingList(pagination Pagination, results []PlexSourcePropertyMapping, ) *PaginatedPlexSourcePropertyMappingList`

NewPaginatedPlexSourcePropertyMappingList instantiates a new PaginatedPlexSourcePropertyMappingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedPlexSourcePropertyMappingListWithDefaults

`func NewPaginatedPlexSourcePropertyMappingListWithDefaults() *PaginatedPlexSourcePropertyMappingList`

NewPaginatedPlexSourcePropertyMappingListWithDefaults instantiates a new PaginatedPlexSourcePropertyMappingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedPlexSourcePropertyMappingList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedPlexSourcePropertyMappingList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedPlexSourcePropertyMappingList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedPlexSourcePropertyMappingList) GetResults() []PlexSourcePropertyMapping`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedPlexSourcePropertyMappingList) GetResultsOk() (*[]PlexSourcePropertyMapping, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedPlexSourcePropertyMappingList) SetResults(v []PlexSourcePropertyMapping)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


