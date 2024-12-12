# PaginatedRedirectStageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]RedirectStage**](RedirectStage.md) |  | 

## Methods

### NewPaginatedRedirectStageList

`func NewPaginatedRedirectStageList(pagination Pagination, results []RedirectStage, ) *PaginatedRedirectStageList`

NewPaginatedRedirectStageList instantiates a new PaginatedRedirectStageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedRedirectStageListWithDefaults

`func NewPaginatedRedirectStageListWithDefaults() *PaginatedRedirectStageList`

NewPaginatedRedirectStageListWithDefaults instantiates a new PaginatedRedirectStageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedRedirectStageList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedRedirectStageList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedRedirectStageList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedRedirectStageList) GetResults() []RedirectStage`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedRedirectStageList) GetResultsOk() (*[]RedirectStage, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedRedirectStageList) SetResults(v []RedirectStage)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


