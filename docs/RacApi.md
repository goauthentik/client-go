# \RacApi

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RacEndpointsCreate**](RacApi.md#RacEndpointsCreate) | **Post** /rac/endpoints/ | 
[**RacEndpointsDestroy**](RacApi.md#RacEndpointsDestroy) | **Delete** /rac/endpoints/{pbm_uuid}/ | 
[**RacEndpointsList**](RacApi.md#RacEndpointsList) | **Get** /rac/endpoints/ | 
[**RacEndpointsPartialUpdate**](RacApi.md#RacEndpointsPartialUpdate) | **Patch** /rac/endpoints/{pbm_uuid}/ | 
[**RacEndpointsRetrieve**](RacApi.md#RacEndpointsRetrieve) | **Get** /rac/endpoints/{pbm_uuid}/ | 
[**RacEndpointsUpdate**](RacApi.md#RacEndpointsUpdate) | **Put** /rac/endpoints/{pbm_uuid}/ | 
[**RacEndpointsUsedByList**](RacApi.md#RacEndpointsUsedByList) | **Get** /rac/endpoints/{pbm_uuid}/used_by/ | 



## RacEndpointsCreate

> Endpoint RacEndpointsCreate(ctx).EndpointRequest(endpointRequest).Execute()





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
    endpointRequest := *openapiclient.NewEndpointRequest("Name_example", int32(123), openapiclient.ProtocolEnum("rdp"), "Host_example", openapiclient.AuthModeEnum("static")) // EndpointRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RacApi.RacEndpointsCreate(context.Background()).EndpointRequest(endpointRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RacApi.RacEndpointsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RacEndpointsCreate`: Endpoint
    fmt.Fprintf(os.Stdout, "Response from `RacApi.RacEndpointsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRacEndpointsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointRequest** | [**EndpointRequest**](EndpointRequest.md) |  | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RacEndpointsDestroy

> RacEndpointsDestroy(ctx, pbmUuid).Execute()





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
    pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this RAC Endpoint.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RacApi.RacEndpointsDestroy(context.Background(), pbmUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RacApi.RacEndpointsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this RAC Endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRacEndpointsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RacEndpointsList

> PaginatedEndpointList RacEndpointsList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Search(search).SuperuserFullList(superuserFullList).Execute()





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
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    provider := int32(56) // int32 |  (optional)
    search := "search_example" // string |  (optional)
    superuserFullList := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RacApi.RacEndpointsList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Search(search).SuperuserFullList(superuserFullList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RacApi.RacEndpointsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RacEndpointsList`: PaginatedEndpointList
    fmt.Fprintf(os.Stdout, "Response from `RacApi.RacEndpointsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRacEndpointsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **provider** | **int32** |  | 
 **search** | **string** |  | 
 **superuserFullList** | **bool** |  | 

### Return type

[**PaginatedEndpointList**](PaginatedEndpointList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RacEndpointsPartialUpdate

> Endpoint RacEndpointsPartialUpdate(ctx, pbmUuid).PatchedEndpointRequest(patchedEndpointRequest).Execute()





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
    pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this RAC Endpoint.
    patchedEndpointRequest := *openapiclient.NewPatchedEndpointRequest() // PatchedEndpointRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RacApi.RacEndpointsPartialUpdate(context.Background(), pbmUuid).PatchedEndpointRequest(patchedEndpointRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RacApi.RacEndpointsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RacEndpointsPartialUpdate`: Endpoint
    fmt.Fprintf(os.Stdout, "Response from `RacApi.RacEndpointsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this RAC Endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRacEndpointsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedEndpointRequest** | [**PatchedEndpointRequest**](PatchedEndpointRequest.md) |  | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RacEndpointsRetrieve

> Endpoint RacEndpointsRetrieve(ctx, pbmUuid).Execute()





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
    pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this RAC Endpoint.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RacApi.RacEndpointsRetrieve(context.Background(), pbmUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RacApi.RacEndpointsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RacEndpointsRetrieve`: Endpoint
    fmt.Fprintf(os.Stdout, "Response from `RacApi.RacEndpointsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this RAC Endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRacEndpointsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RacEndpointsUpdate

> Endpoint RacEndpointsUpdate(ctx, pbmUuid).EndpointRequest(endpointRequest).Execute()





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
    pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this RAC Endpoint.
    endpointRequest := *openapiclient.NewEndpointRequest("Name_example", int32(123), openapiclient.ProtocolEnum("rdp"), "Host_example", openapiclient.AuthModeEnum("static")) // EndpointRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RacApi.RacEndpointsUpdate(context.Background(), pbmUuid).EndpointRequest(endpointRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RacApi.RacEndpointsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RacEndpointsUpdate`: Endpoint
    fmt.Fprintf(os.Stdout, "Response from `RacApi.RacEndpointsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this RAC Endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRacEndpointsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointRequest** | [**EndpointRequest**](EndpointRequest.md) |  | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RacEndpointsUsedByList

> []UsedBy RacEndpointsUsedByList(ctx, pbmUuid).Execute()





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
    pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this RAC Endpoint.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RacApi.RacEndpointsUsedByList(context.Background(), pbmUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RacApi.RacEndpointsUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RacEndpointsUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `RacApi.RacEndpointsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this RAC Endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRacEndpointsUsedByListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UsedBy**](UsedBy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

