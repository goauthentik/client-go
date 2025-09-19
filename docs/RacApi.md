# \RacApi

All URIs are relative to */api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RacConnectionTokensDestroy**](RacApi.md#RacConnectionTokensDestroy) | **Delete** /rac/connection_tokens/{connection_token_uuid}/ | 
[**RacConnectionTokensList**](RacApi.md#RacConnectionTokensList) | **Get** /rac/connection_tokens/ | 
[**RacConnectionTokensPartialUpdate**](RacApi.md#RacConnectionTokensPartialUpdate) | **Patch** /rac/connection_tokens/{connection_token_uuid}/ | 
[**RacConnectionTokensRetrieve**](RacApi.md#RacConnectionTokensRetrieve) | **Get** /rac/connection_tokens/{connection_token_uuid}/ | 
[**RacConnectionTokensUpdate**](RacApi.md#RacConnectionTokensUpdate) | **Put** /rac/connection_tokens/{connection_token_uuid}/ | 
[**RacConnectionTokensUsedByList**](RacApi.md#RacConnectionTokensUsedByList) | **Get** /rac/connection_tokens/{connection_token_uuid}/used_by/ | 
[**RacEndpointsCreate**](RacApi.md#RacEndpointsCreate) | **Post** /rac/endpoints/ | 
[**RacEndpointsDestroy**](RacApi.md#RacEndpointsDestroy) | **Delete** /rac/endpoints/{pbm_uuid}/ | 
[**RacEndpointsList**](RacApi.md#RacEndpointsList) | **Get** /rac/endpoints/ | 
[**RacEndpointsPartialUpdate**](RacApi.md#RacEndpointsPartialUpdate) | **Patch** /rac/endpoints/{pbm_uuid}/ | 
[**RacEndpointsRetrieve**](RacApi.md#RacEndpointsRetrieve) | **Get** /rac/endpoints/{pbm_uuid}/ | 
[**RacEndpointsUpdate**](RacApi.md#RacEndpointsUpdate) | **Put** /rac/endpoints/{pbm_uuid}/ | 
[**RacEndpointsUsedByList**](RacApi.md#RacEndpointsUsedByList) | **Get** /rac/endpoints/{pbm_uuid}/used_by/ | 



## RacConnectionTokensDestroy

> RacConnectionTokensDestroy(ctx, connectionTokenUuid).Execute()





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
    connectionTokenUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this RAC Connection token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RacApi.RacConnectionTokensDestroy(context.Background(), connectionTokenUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RacApi.RacConnectionTokensDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionTokenUuid** | **string** | A UUID string identifying this RAC Connection token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRacConnectionTokensDestroyRequest struct via the builder pattern


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


## RacConnectionTokensList

> PaginatedConnectionTokenList RacConnectionTokensList(ctx).Endpoint(endpoint).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Search(search).SessionUser(sessionUser).Execute()





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
    endpoint := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    provider := int32(56) // int32 |  (optional)
    search := "search_example" // string | A search term. (optional)
    sessionUser := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RacApi.RacConnectionTokensList(context.Background()).Endpoint(endpoint).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Search(search).SessionUser(sessionUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RacApi.RacConnectionTokensList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RacConnectionTokensList`: PaginatedConnectionTokenList
    fmt.Fprintf(os.Stdout, "Response from `RacApi.RacConnectionTokensList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRacConnectionTokensListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpoint** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **provider** | **int32** |  | 
 **search** | **string** | A search term. | 
 **sessionUser** | **int32** |  | 

### Return type

[**PaginatedConnectionTokenList**](PaginatedConnectionTokenList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RacConnectionTokensPartialUpdate

> ConnectionToken RacConnectionTokensPartialUpdate(ctx, connectionTokenUuid).PatchedConnectionTokenRequest(patchedConnectionTokenRequest).Execute()





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
    connectionTokenUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this RAC Connection token.
    patchedConnectionTokenRequest := *openapiclient.NewPatchedConnectionTokenRequest() // PatchedConnectionTokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RacApi.RacConnectionTokensPartialUpdate(context.Background(), connectionTokenUuid).PatchedConnectionTokenRequest(patchedConnectionTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RacApi.RacConnectionTokensPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RacConnectionTokensPartialUpdate`: ConnectionToken
    fmt.Fprintf(os.Stdout, "Response from `RacApi.RacConnectionTokensPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionTokenUuid** | **string** | A UUID string identifying this RAC Connection token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRacConnectionTokensPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedConnectionTokenRequest** | [**PatchedConnectionTokenRequest**](PatchedConnectionTokenRequest.md) |  | 

### Return type

[**ConnectionToken**](ConnectionToken.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RacConnectionTokensRetrieve

> ConnectionToken RacConnectionTokensRetrieve(ctx, connectionTokenUuid).Execute()





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
    connectionTokenUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this RAC Connection token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RacApi.RacConnectionTokensRetrieve(context.Background(), connectionTokenUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RacApi.RacConnectionTokensRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RacConnectionTokensRetrieve`: ConnectionToken
    fmt.Fprintf(os.Stdout, "Response from `RacApi.RacConnectionTokensRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionTokenUuid** | **string** | A UUID string identifying this RAC Connection token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRacConnectionTokensRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectionToken**](ConnectionToken.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RacConnectionTokensUpdate

> ConnectionToken RacConnectionTokensUpdate(ctx, connectionTokenUuid).ConnectionTokenRequest(connectionTokenRequest).Execute()





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
    connectionTokenUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this RAC Connection token.
    connectionTokenRequest := *openapiclient.NewConnectionTokenRequest(int32(123), "Endpoint_example") // ConnectionTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RacApi.RacConnectionTokensUpdate(context.Background(), connectionTokenUuid).ConnectionTokenRequest(connectionTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RacApi.RacConnectionTokensUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RacConnectionTokensUpdate`: ConnectionToken
    fmt.Fprintf(os.Stdout, "Response from `RacApi.RacConnectionTokensUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionTokenUuid** | **string** | A UUID string identifying this RAC Connection token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRacConnectionTokensUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionTokenRequest** | [**ConnectionTokenRequest**](ConnectionTokenRequest.md) |  | 

### Return type

[**ConnectionToken**](ConnectionToken.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RacConnectionTokensUsedByList

> []UsedBy RacConnectionTokensUsedByList(ctx, connectionTokenUuid).Execute()





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
    connectionTokenUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this RAC Connection token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RacApi.RacConnectionTokensUsedByList(context.Background(), connectionTokenUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RacApi.RacConnectionTokensUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RacConnectionTokensUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `RacApi.RacConnectionTokensUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionTokenUuid** | **string** | A UUID string identifying this RAC Connection token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRacConnectionTokensUsedByListRequest struct via the builder pattern


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
    search := "search_example" // string | A search term. (optional)
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
 **search** | **string** | A search term. | 
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

