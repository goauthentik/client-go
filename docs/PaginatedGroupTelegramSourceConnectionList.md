# PaginatedGroupTelegramSourceConnectionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]GroupTelegramSourceConnection**](GroupTelegramSourceConnection.md) |  | 
**Autocomplete** | **map[string]interface{}** |  | 

## Methods

### NewPaginatedGroupTelegramSourceConnectionList

`func NewPaginatedGroupTelegramSourceConnectionList(pagination Pagination, results []GroupTelegramSourceConnection, autocomplete map[string]interface{}, ) *PaginatedGroupTelegramSourceConnectionList`

NewPaginatedGroupTelegramSourceConnectionList instantiates a new PaginatedGroupTelegramSourceConnectionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedGroupTelegramSourceConnectionListWithDefaults

`func NewPaginatedGroupTelegramSourceConnectionListWithDefaults() *PaginatedGroupTelegramSourceConnectionList`

NewPaginatedGroupTelegramSourceConnectionListWithDefaults instantiates a new PaginatedGroupTelegramSourceConnectionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedGroupTelegramSourceConnectionList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedGroupTelegramSourceConnectionList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedGroupTelegramSourceConnectionList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedGroupTelegramSourceConnectionList) GetResults() []GroupTelegramSourceConnection`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedGroupTelegramSourceConnectionList) GetResultsOk() (*[]GroupTelegramSourceConnection, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedGroupTelegramSourceConnectionList) SetResults(v []GroupTelegramSourceConnection)`

SetResults sets Results field to given value.


### GetAutocomplete

`func (o *PaginatedGroupTelegramSourceConnectionList) GetAutocomplete() map[string]interface{}`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *PaginatedGroupTelegramSourceConnectionList) GetAutocompleteOk() (*map[string]interface{}, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *PaginatedGroupTelegramSourceConnectionList) SetAutocomplete(v map[string]interface{})`

SetAutocomplete sets Autocomplete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


