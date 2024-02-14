# PaginatedConnectionTokenList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]ConnectionToken**](ConnectionToken.md) |  | 

## Methods

### NewPaginatedConnectionTokenList

`func NewPaginatedConnectionTokenList(pagination Pagination, results []ConnectionToken, ) *PaginatedConnectionTokenList`

NewPaginatedConnectionTokenList instantiates a new PaginatedConnectionTokenList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedConnectionTokenListWithDefaults

`func NewPaginatedConnectionTokenListWithDefaults() *PaginatedConnectionTokenList`

NewPaginatedConnectionTokenListWithDefaults instantiates a new PaginatedConnectionTokenList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedConnectionTokenList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedConnectionTokenList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedConnectionTokenList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedConnectionTokenList) GetResults() []ConnectionToken`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedConnectionTokenList) GetResultsOk() (*[]ConnectionToken, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedConnectionTokenList) SetResults(v []ConnectionToken)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


