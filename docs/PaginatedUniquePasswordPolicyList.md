# PaginatedUniquePasswordPolicyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]UniquePasswordPolicy**](UniquePasswordPolicy.md) |  | 
**Autocomplete** | **map[string]interface{}** |  | 

## Methods

### NewPaginatedUniquePasswordPolicyList

`func NewPaginatedUniquePasswordPolicyList(pagination Pagination, results []UniquePasswordPolicy, autocomplete map[string]interface{}, ) *PaginatedUniquePasswordPolicyList`

NewPaginatedUniquePasswordPolicyList instantiates a new PaginatedUniquePasswordPolicyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedUniquePasswordPolicyListWithDefaults

`func NewPaginatedUniquePasswordPolicyListWithDefaults() *PaginatedUniquePasswordPolicyList`

NewPaginatedUniquePasswordPolicyListWithDefaults instantiates a new PaginatedUniquePasswordPolicyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedUniquePasswordPolicyList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedUniquePasswordPolicyList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedUniquePasswordPolicyList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedUniquePasswordPolicyList) GetResults() []UniquePasswordPolicy`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedUniquePasswordPolicyList) GetResultsOk() (*[]UniquePasswordPolicy, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedUniquePasswordPolicyList) SetResults(v []UniquePasswordPolicy)`

SetResults sets Results field to given value.


### GetAutocomplete

`func (o *PaginatedUniquePasswordPolicyList) GetAutocomplete() map[string]interface{}`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *PaginatedUniquePasswordPolicyList) GetAutocompleteOk() (*map[string]interface{}, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *PaginatedUniquePasswordPolicyList) SetAutocomplete(v map[string]interface{})`

SetAutocomplete sets Autocomplete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


