# PaginatedRACProviderList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]RACProvider**](RACProvider.md) |  | 

## Methods

### NewPaginatedRACProviderList

`func NewPaginatedRACProviderList(pagination Pagination, results []RACProvider, ) *PaginatedRACProviderList`

NewPaginatedRACProviderList instantiates a new PaginatedRACProviderList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedRACProviderListWithDefaults

`func NewPaginatedRACProviderListWithDefaults() *PaginatedRACProviderList`

NewPaginatedRACProviderListWithDefaults instantiates a new PaginatedRACProviderList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedRACProviderList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedRACProviderList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedRACProviderList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedRACProviderList) GetResults() []RACProvider`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedRACProviderList) GetResultsOk() (*[]RACProvider, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedRACProviderList) SetResults(v []RACProvider)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


