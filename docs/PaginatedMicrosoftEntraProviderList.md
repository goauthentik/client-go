# PaginatedMicrosoftEntraProviderList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]MicrosoftEntraProvider**](MicrosoftEntraProvider.md) |  | 
**Autocomplete** | **map[string]interface{}** |  | 

## Methods

### NewPaginatedMicrosoftEntraProviderList

`func NewPaginatedMicrosoftEntraProviderList(pagination Pagination, results []MicrosoftEntraProvider, autocomplete map[string]interface{}, ) *PaginatedMicrosoftEntraProviderList`

NewPaginatedMicrosoftEntraProviderList instantiates a new PaginatedMicrosoftEntraProviderList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedMicrosoftEntraProviderListWithDefaults

`func NewPaginatedMicrosoftEntraProviderListWithDefaults() *PaginatedMicrosoftEntraProviderList`

NewPaginatedMicrosoftEntraProviderListWithDefaults instantiates a new PaginatedMicrosoftEntraProviderList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedMicrosoftEntraProviderList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedMicrosoftEntraProviderList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedMicrosoftEntraProviderList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedMicrosoftEntraProviderList) GetResults() []MicrosoftEntraProvider`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedMicrosoftEntraProviderList) GetResultsOk() (*[]MicrosoftEntraProvider, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedMicrosoftEntraProviderList) SetResults(v []MicrosoftEntraProvider)`

SetResults sets Results field to given value.


### GetAutocomplete

`func (o *PaginatedMicrosoftEntraProviderList) GetAutocomplete() map[string]interface{}`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *PaginatedMicrosoftEntraProviderList) GetAutocompleteOk() (*map[string]interface{}, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *PaginatedMicrosoftEntraProviderList) SetAutocomplete(v map[string]interface{})`

SetAutocomplete sets Autocomplete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


