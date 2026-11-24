# PaginatedUserSAMLSourceConnectionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]UserSAMLSourceConnection**](UserSAMLSourceConnection.md) |  | 
**Autocomplete** | **map[string]interface{}** |  | 

## Methods

### NewPaginatedUserSAMLSourceConnectionList

`func NewPaginatedUserSAMLSourceConnectionList(pagination Pagination, results []UserSAMLSourceConnection, autocomplete map[string]interface{}, ) *PaginatedUserSAMLSourceConnectionList`

NewPaginatedUserSAMLSourceConnectionList instantiates a new PaginatedUserSAMLSourceConnectionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedUserSAMLSourceConnectionListWithDefaults

`func NewPaginatedUserSAMLSourceConnectionListWithDefaults() *PaginatedUserSAMLSourceConnectionList`

NewPaginatedUserSAMLSourceConnectionListWithDefaults instantiates a new PaginatedUserSAMLSourceConnectionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedUserSAMLSourceConnectionList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedUserSAMLSourceConnectionList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedUserSAMLSourceConnectionList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedUserSAMLSourceConnectionList) GetResults() []UserSAMLSourceConnection`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedUserSAMLSourceConnectionList) GetResultsOk() (*[]UserSAMLSourceConnection, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedUserSAMLSourceConnectionList) SetResults(v []UserSAMLSourceConnection)`

SetResults sets Results field to given value.


### GetAutocomplete

`func (o *PaginatedUserSAMLSourceConnectionList) GetAutocomplete() map[string]interface{}`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *PaginatedUserSAMLSourceConnectionList) GetAutocompleteOk() (*map[string]interface{}, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *PaginatedUserSAMLSourceConnectionList) SetAutocomplete(v map[string]interface{})`

SetAutocomplete sets Autocomplete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


