# DataExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**RequestedBy** | [**RequestedBy**](RequestedBy.md) |  | [readonly] 
**RequestedOn** | **time.Time** |  | [readonly] 
**ContentType** | [**ContentType**](ContentType.md) |  | [readonly] 
**QueryParams** | **map[string]interface{}** |  | 
**FileUrl** | **string** |  | [readonly] 
**Completed** | **bool** |  | [readonly] 

## Methods

### NewDataExport

`func NewDataExport(id string, requestedBy RequestedBy, requestedOn time.Time, contentType ContentType, queryParams map[string]interface{}, fileUrl string, completed bool, ) *DataExport`

NewDataExport instantiates a new DataExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataExportWithDefaults

`func NewDataExportWithDefaults() *DataExport`

NewDataExportWithDefaults instantiates a new DataExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataExport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataExport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataExport) SetId(v string)`

SetId sets Id field to given value.


### GetRequestedBy

`func (o *DataExport) GetRequestedBy() RequestedBy`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *DataExport) GetRequestedByOk() (*RequestedBy, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *DataExport) SetRequestedBy(v RequestedBy)`

SetRequestedBy sets RequestedBy field to given value.


### GetRequestedOn

`func (o *DataExport) GetRequestedOn() time.Time`

GetRequestedOn returns the RequestedOn field if non-nil, zero value otherwise.

### GetRequestedOnOk

`func (o *DataExport) GetRequestedOnOk() (*time.Time, bool)`

GetRequestedOnOk returns a tuple with the RequestedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedOn

`func (o *DataExport) SetRequestedOn(v time.Time)`

SetRequestedOn sets RequestedOn field to given value.


### GetContentType

`func (o *DataExport) GetContentType() ContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *DataExport) GetContentTypeOk() (*ContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *DataExport) SetContentType(v ContentType)`

SetContentType sets ContentType field to given value.


### GetQueryParams

`func (o *DataExport) GetQueryParams() map[string]interface{}`

GetQueryParams returns the QueryParams field if non-nil, zero value otherwise.

### GetQueryParamsOk

`func (o *DataExport) GetQueryParamsOk() (*map[string]interface{}, bool)`

GetQueryParamsOk returns a tuple with the QueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParams

`func (o *DataExport) SetQueryParams(v map[string]interface{})`

SetQueryParams sets QueryParams field to given value.


### GetFileUrl

`func (o *DataExport) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *DataExport) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *DataExport) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetCompleted

`func (o *DataExport) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *DataExport) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *DataExport) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


