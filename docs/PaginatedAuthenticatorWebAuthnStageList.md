# PaginatedAuthenticatorWebAuthnStageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]AuthenticatorWebAuthnStage**](AuthenticatorWebAuthnStage.md) |  | 
**Autocomplete** | **map[string]interface{}** |  | 

## Methods

### NewPaginatedAuthenticatorWebAuthnStageList

`func NewPaginatedAuthenticatorWebAuthnStageList(pagination Pagination, results []AuthenticatorWebAuthnStage, autocomplete map[string]interface{}, ) *PaginatedAuthenticatorWebAuthnStageList`

NewPaginatedAuthenticatorWebAuthnStageList instantiates a new PaginatedAuthenticatorWebAuthnStageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedAuthenticatorWebAuthnStageListWithDefaults

`func NewPaginatedAuthenticatorWebAuthnStageListWithDefaults() *PaginatedAuthenticatorWebAuthnStageList`

NewPaginatedAuthenticatorWebAuthnStageListWithDefaults instantiates a new PaginatedAuthenticatorWebAuthnStageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedAuthenticatorWebAuthnStageList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedAuthenticatorWebAuthnStageList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedAuthenticatorWebAuthnStageList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedAuthenticatorWebAuthnStageList) GetResults() []AuthenticatorWebAuthnStage`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedAuthenticatorWebAuthnStageList) GetResultsOk() (*[]AuthenticatorWebAuthnStage, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedAuthenticatorWebAuthnStageList) SetResults(v []AuthenticatorWebAuthnStage)`

SetResults sets Results field to given value.


### GetAutocomplete

`func (o *PaginatedAuthenticatorWebAuthnStageList) GetAutocomplete() map[string]interface{}`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *PaginatedAuthenticatorWebAuthnStageList) GetAutocompleteOk() (*map[string]interface{}, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *PaginatedAuthenticatorWebAuthnStageList) SetAutocomplete(v map[string]interface{})`

SetAutocomplete sets Autocomplete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


