# PaginatedProxyOutpostConfigList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]ProxyOutpostConfig**](ProxyOutpostConfig.md) |  | 

## Methods

### NewPaginatedProxyOutpostConfigList

`func NewPaginatedProxyOutpostConfigList(pagination Pagination, results []ProxyOutpostConfig, ) *PaginatedProxyOutpostConfigList`

NewPaginatedProxyOutpostConfigList instantiates a new PaginatedProxyOutpostConfigList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedProxyOutpostConfigListWithDefaults

`func NewPaginatedProxyOutpostConfigListWithDefaults() *PaginatedProxyOutpostConfigList`

NewPaginatedProxyOutpostConfigListWithDefaults instantiates a new PaginatedProxyOutpostConfigList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedProxyOutpostConfigList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedProxyOutpostConfigList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedProxyOutpostConfigList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedProxyOutpostConfigList) GetResults() []ProxyOutpostConfig`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedProxyOutpostConfigList) GetResultsOk() (*[]ProxyOutpostConfig, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedProxyOutpostConfigList) SetResults(v []ProxyOutpostConfig)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


