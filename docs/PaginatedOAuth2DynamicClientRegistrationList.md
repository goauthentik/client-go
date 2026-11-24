# PaginatedOAuth2DynamicClientRegistrationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]OAuth2DynamicClientRegistration**](OAuth2DynamicClientRegistration.md) |  | 
**Autocomplete** | **map[string]interface{}** |  | 

## Methods

### NewPaginatedOAuth2DynamicClientRegistrationList

`func NewPaginatedOAuth2DynamicClientRegistrationList(pagination Pagination, results []OAuth2DynamicClientRegistration, autocomplete map[string]interface{}, ) *PaginatedOAuth2DynamicClientRegistrationList`

NewPaginatedOAuth2DynamicClientRegistrationList instantiates a new PaginatedOAuth2DynamicClientRegistrationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedOAuth2DynamicClientRegistrationListWithDefaults

`func NewPaginatedOAuth2DynamicClientRegistrationListWithDefaults() *PaginatedOAuth2DynamicClientRegistrationList`

NewPaginatedOAuth2DynamicClientRegistrationListWithDefaults instantiates a new PaginatedOAuth2DynamicClientRegistrationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedOAuth2DynamicClientRegistrationList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedOAuth2DynamicClientRegistrationList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedOAuth2DynamicClientRegistrationList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedOAuth2DynamicClientRegistrationList) GetResults() []OAuth2DynamicClientRegistration`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedOAuth2DynamicClientRegistrationList) GetResultsOk() (*[]OAuth2DynamicClientRegistration, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedOAuth2DynamicClientRegistrationList) SetResults(v []OAuth2DynamicClientRegistration)`

SetResults sets Results field to given value.


### GetAutocomplete

`func (o *PaginatedOAuth2DynamicClientRegistrationList) GetAutocomplete() map[string]interface{}`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *PaginatedOAuth2DynamicClientRegistrationList) GetAutocompleteOk() (*map[string]interface{}, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *PaginatedOAuth2DynamicClientRegistrationList) SetAutocomplete(v map[string]interface{})`

SetAutocomplete sets Autocomplete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


