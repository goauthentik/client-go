# PaginatedUserPlexSourceConnectionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]UserPlexSourceConnection**](UserPlexSourceConnection.md) |  | 
**Autocomplete** | **map[string]interface{}** |  | 

## Methods

### NewPaginatedUserPlexSourceConnectionList

`func NewPaginatedUserPlexSourceConnectionList(pagination Pagination, results []UserPlexSourceConnection, autocomplete map[string]interface{}, ) *PaginatedUserPlexSourceConnectionList`

NewPaginatedUserPlexSourceConnectionList instantiates a new PaginatedUserPlexSourceConnectionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedUserPlexSourceConnectionListWithDefaults

`func NewPaginatedUserPlexSourceConnectionListWithDefaults() *PaginatedUserPlexSourceConnectionList`

NewPaginatedUserPlexSourceConnectionListWithDefaults instantiates a new PaginatedUserPlexSourceConnectionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedUserPlexSourceConnectionList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedUserPlexSourceConnectionList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedUserPlexSourceConnectionList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedUserPlexSourceConnectionList) GetResults() []UserPlexSourceConnection`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedUserPlexSourceConnectionList) GetResultsOk() (*[]UserPlexSourceConnection, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedUserPlexSourceConnectionList) SetResults(v []UserPlexSourceConnection)`

SetResults sets Results field to given value.


### GetAutocomplete

`func (o *PaginatedUserPlexSourceConnectionList) GetAutocomplete() map[string]interface{}`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *PaginatedUserPlexSourceConnectionList) GetAutocompleteOk() (*map[string]interface{}, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *PaginatedUserPlexSourceConnectionList) SetAutocomplete(v map[string]interface{})`

SetAutocomplete sets Autocomplete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


