# PaginatedIdentificationStageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**PaginatedApplicationListPagination**](PaginatedApplicationListPagination.md) |  | 
**Results** | [**[]IdentificationStage**](IdentificationStage.md) |  | 

## Methods

### NewPaginatedIdentificationStageList

`func NewPaginatedIdentificationStageList(pagination PaginatedApplicationListPagination, results []IdentificationStage, ) *PaginatedIdentificationStageList`

NewPaginatedIdentificationStageList instantiates a new PaginatedIdentificationStageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedIdentificationStageListWithDefaults

`func NewPaginatedIdentificationStageListWithDefaults() *PaginatedIdentificationStageList`

NewPaginatedIdentificationStageListWithDefaults instantiates a new PaginatedIdentificationStageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedIdentificationStageList) GetPagination() PaginatedApplicationListPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedIdentificationStageList) GetPaginationOk() (*PaginatedApplicationListPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedIdentificationStageList) SetPagination(v PaginatedApplicationListPagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedIdentificationStageList) GetResults() []IdentificationStage`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedIdentificationStageList) GetResultsOk() (*[]IdentificationStage, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedIdentificationStageList) SetResults(v []IdentificationStage)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


