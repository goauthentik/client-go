# PaginatedOAuth2ProviderList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]OAuth2Provider**](OAuth2Provider.md) |  | 

## Methods

### NewPaginatedOAuth2ProviderList

`func NewPaginatedOAuth2ProviderList(pagination Pagination, results []OAuth2Provider, ) *PaginatedOAuth2ProviderList`

NewPaginatedOAuth2ProviderList instantiates a new PaginatedOAuth2ProviderList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedOAuth2ProviderListWithDefaults

`func NewPaginatedOAuth2ProviderListWithDefaults() *PaginatedOAuth2ProviderList`

NewPaginatedOAuth2ProviderListWithDefaults instantiates a new PaginatedOAuth2ProviderList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedOAuth2ProviderList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedOAuth2ProviderList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedOAuth2ProviderList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedOAuth2ProviderList) GetResults() []OAuth2Provider`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedOAuth2ProviderList) GetResultsOk() (*[]OAuth2Provider, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedOAuth2ProviderList) SetResults(v []OAuth2Provider)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


