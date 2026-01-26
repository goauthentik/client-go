# PaginatedGeoIPPolicyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]GeoIPPolicy**](GeoIPPolicy.md) |  | 

## Methods

### NewPaginatedGeoIPPolicyList

`func NewPaginatedGeoIPPolicyList(pagination Pagination, results []GeoIPPolicy, ) *PaginatedGeoIPPolicyList`

NewPaginatedGeoIPPolicyList instantiates a new PaginatedGeoIPPolicyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedGeoIPPolicyListWithDefaults

`func NewPaginatedGeoIPPolicyListWithDefaults() *PaginatedGeoIPPolicyList`

NewPaginatedGeoIPPolicyListWithDefaults instantiates a new PaginatedGeoIPPolicyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedGeoIPPolicyList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedGeoIPPolicyList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedGeoIPPolicyList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedGeoIPPolicyList) GetResults() []GeoIPPolicy`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedGeoIPPolicyList) GetResultsOk() (*[]GeoIPPolicy, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedGeoIPPolicyList) SetResults(v []GeoIPPolicy)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


