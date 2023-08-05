# PaginatedFlowStageBindingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]FlowStageBinding**](FlowStageBinding.md) |  | 

## Methods

### NewPaginatedFlowStageBindingList

`func NewPaginatedFlowStageBindingList(pagination Pagination, results []FlowStageBinding, ) *PaginatedFlowStageBindingList`

NewPaginatedFlowStageBindingList instantiates a new PaginatedFlowStageBindingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedFlowStageBindingListWithDefaults

`func NewPaginatedFlowStageBindingListWithDefaults() *PaginatedFlowStageBindingList`

NewPaginatedFlowStageBindingListWithDefaults instantiates a new PaginatedFlowStageBindingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedFlowStageBindingList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedFlowStageBindingList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedFlowStageBindingList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedFlowStageBindingList) GetResults() []FlowStageBinding`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedFlowStageBindingList) GetResultsOk() (*[]FlowStageBinding, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedFlowStageBindingList) SetResults(v []FlowStageBinding)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


