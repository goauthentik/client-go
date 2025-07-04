# PaginatedGoogleWorkspaceProviderGroupList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]GoogleWorkspaceProviderGroup**](GoogleWorkspaceProviderGroup.md) |  | 
**Autocomplete** | **map[string]interface{}** |  | 

## Methods

### NewPaginatedGoogleWorkspaceProviderGroupList

`func NewPaginatedGoogleWorkspaceProviderGroupList(pagination Pagination, results []GoogleWorkspaceProviderGroup, autocomplete map[string]interface{}, ) *PaginatedGoogleWorkspaceProviderGroupList`

NewPaginatedGoogleWorkspaceProviderGroupList instantiates a new PaginatedGoogleWorkspaceProviderGroupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedGoogleWorkspaceProviderGroupListWithDefaults

`func NewPaginatedGoogleWorkspaceProviderGroupListWithDefaults() *PaginatedGoogleWorkspaceProviderGroupList`

NewPaginatedGoogleWorkspaceProviderGroupListWithDefaults instantiates a new PaginatedGoogleWorkspaceProviderGroupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedGoogleWorkspaceProviderGroupList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedGoogleWorkspaceProviderGroupList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedGoogleWorkspaceProviderGroupList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedGoogleWorkspaceProviderGroupList) GetResults() []GoogleWorkspaceProviderGroup`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedGoogleWorkspaceProviderGroupList) GetResultsOk() (*[]GoogleWorkspaceProviderGroup, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedGoogleWorkspaceProviderGroupList) SetResults(v []GoogleWorkspaceProviderGroup)`

SetResults sets Results field to given value.


### GetAutocomplete

`func (o *PaginatedGoogleWorkspaceProviderGroupList) GetAutocomplete() map[string]interface{}`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *PaginatedGoogleWorkspaceProviderGroupList) GetAutocompleteOk() (*map[string]interface{}, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *PaginatedGoogleWorkspaceProviderGroupList) SetAutocomplete(v map[string]interface{})`

SetAutocomplete sets Autocomplete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


