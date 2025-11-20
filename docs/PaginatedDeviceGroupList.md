# PaginatedDeviceGroupList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Results** | [**[]DeviceGroup**](DeviceGroup.md) |  | 
**Autocomplete** | **map[string]interface{}** |  | 

## Methods

### NewPaginatedDeviceGroupList

`func NewPaginatedDeviceGroupList(pagination Pagination, results []DeviceGroup, autocomplete map[string]interface{}, ) *PaginatedDeviceGroupList`

NewPaginatedDeviceGroupList instantiates a new PaginatedDeviceGroupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedDeviceGroupListWithDefaults

`func NewPaginatedDeviceGroupListWithDefaults() *PaginatedDeviceGroupList`

NewPaginatedDeviceGroupListWithDefaults instantiates a new PaginatedDeviceGroupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PaginatedDeviceGroupList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PaginatedDeviceGroupList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PaginatedDeviceGroupList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetResults

`func (o *PaginatedDeviceGroupList) GetResults() []DeviceGroup`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedDeviceGroupList) GetResultsOk() (*[]DeviceGroup, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedDeviceGroupList) SetResults(v []DeviceGroup)`

SetResults sets Results field to given value.


### GetAutocomplete

`func (o *PaginatedDeviceGroupList) GetAutocomplete() map[string]interface{}`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *PaginatedDeviceGroupList) GetAutocompleteOk() (*map[string]interface{}, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *PaginatedDeviceGroupList) SetAutocomplete(v map[string]interface{})`

SetAutocomplete sets Autocomplete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


