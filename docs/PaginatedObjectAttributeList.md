# PaginatedObjectAttributeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]ObjectAttribute**](ObjectAttribute.md) |  | 
**Autocomplete** | **map[string]interface{}** |  | 

## Methods

### NewPaginatedObjectAttributeList

`func NewPaginatedObjectAttributeList(pagination Pagination, results []ObjectAttribute, autocomplete map[string]interface{}, ) *PaginatedObjectAttributeList`

NewPaginatedObjectAttributeList instantiates a new PaginatedObjectAttributeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedObjectAttributeListWithDefaults

`func NewPaginatedObjectAttributeListWithDefaults() *PaginatedObjectAttributeList`

NewPaginatedObjectAttributeListWithDefaults instantiates a new PaginatedObjectAttributeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedObjectAttributeList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedObjectAttributeList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedObjectAttributeList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedObjectAttributeList) GetResults() []ObjectAttribute`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedObjectAttributeList) GetResultsOk() (*[]ObjectAttribute, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedObjectAttributeList) SetResults(v []ObjectAttribute)`

SetResults sets Results field to given value.


### GetAutocomplete

`func (o *PaginatedObjectAttributeList) GetAutocomplete() map[string]interface{}`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *PaginatedObjectAttributeList) GetAutocompleteOk() (*map[string]interface{}, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *PaginatedObjectAttributeList) SetAutocomplete(v map[string]interface{})`

SetAutocomplete sets Autocomplete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


