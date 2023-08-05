# PaginatedLDAPSourceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]LDAPSource**](LDAPSource.md) |  | 

## Methods

### NewPaginatedLDAPSourceList

`func NewPaginatedLDAPSourceList(pagination Pagination, results []LDAPSource, ) *PaginatedLDAPSourceList`

NewPaginatedLDAPSourceList instantiates a new PaginatedLDAPSourceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedLDAPSourceListWithDefaults

`func NewPaginatedLDAPSourceListWithDefaults() *PaginatedLDAPSourceList`

NewPaginatedLDAPSourceListWithDefaults instantiates a new PaginatedLDAPSourceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedLDAPSourceList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedLDAPSourceList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedLDAPSourceList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedLDAPSourceList) GetResults() []LDAPSource`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedLDAPSourceList) GetResultsOk() (*[]LDAPSource, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedLDAPSourceList) SetResults(v []LDAPSource)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


