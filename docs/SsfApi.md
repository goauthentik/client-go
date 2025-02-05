# \SsfApi

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SsfStreamsList**](SsfApi.md#SsfStreamsList) | **Get** /ssf/streams/ | 
[**SsfStreamsRetrieve**](SsfApi.md#SsfStreamsRetrieve) | **Get** /ssf/streams/{uuid}/ | 



## SsfStreamsList

> PaginatedSSFStreamList SsfStreamsList(ctx).DeliveryMethod(deliveryMethod).EndpointUrl(endpointUrl).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Search(search).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deliveryMethod := "deliveryMethod_example" // string |  (optional)
    endpointUrl := "endpointUrl_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    provider := int32(56) // int32 |  (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsfApi.SsfStreamsList(context.Background()).DeliveryMethod(deliveryMethod).EndpointUrl(endpointUrl).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsfApi.SsfStreamsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SsfStreamsList`: PaginatedSSFStreamList
    fmt.Fprintf(os.Stdout, "Response from `SsfApi.SsfStreamsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSsfStreamsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deliveryMethod** | **string** |  | 
 **endpointUrl** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **provider** | **int32** |  | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedSSFStreamList**](PaginatedSSFStreamList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SsfStreamsRetrieve

> SSFStream SsfStreamsRetrieve(ctx, uuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this SSF Stream.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsfApi.SsfStreamsRetrieve(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsfApi.SsfStreamsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SsfStreamsRetrieve`: SSFStream
    fmt.Fprintf(os.Stdout, "Response from `SsfApi.SsfStreamsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this SSF Stream. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSsfStreamsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SSFStream**](SSFStream.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

