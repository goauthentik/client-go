# PaginatedRadiusProviderList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]RadiusProvider**](RadiusProvider.md) |  | 

## Methods

### NewPaginatedRadiusProviderList

`func NewPaginatedRadiusProviderList(pagination Pagination, results []RadiusProvider, ) *PaginatedRadiusProviderList`

NewPaginatedRadiusProviderList instantiates a new PaginatedRadiusProviderList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedRadiusProviderListWithDefaults

`func NewPaginatedRadiusProviderListWithDefaults() *PaginatedRadiusProviderList`

NewPaginatedRadiusProviderListWithDefaults instantiates a new PaginatedRadiusProviderList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedRadiusProviderList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedRadiusProviderList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedRadiusProviderList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedRadiusProviderList) GetResults() []RadiusProvider`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedRadiusProviderList) GetResultsOk() (*[]RadiusProvider, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedRadiusProviderList) SetResults(v []RadiusProvider)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


