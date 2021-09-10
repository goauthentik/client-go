# \AuthenticatorsApi

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticatorsAdminDuoCreate**](AuthenticatorsApi.md#AuthenticatorsAdminDuoCreate) | **Post** /authenticators/admin/duo/ | 
[**AuthenticatorsAdminDuoDestroy**](AuthenticatorsApi.md#AuthenticatorsAdminDuoDestroy) | **Delete** /authenticators/admin/duo/{id}/ | 
[**AuthenticatorsAdminDuoList**](AuthenticatorsApi.md#AuthenticatorsAdminDuoList) | **Get** /authenticators/admin/duo/ | 
[**AuthenticatorsAdminDuoPartialUpdate**](AuthenticatorsApi.md#AuthenticatorsAdminDuoPartialUpdate) | **Patch** /authenticators/admin/duo/{id}/ | 
[**AuthenticatorsAdminDuoRetrieve**](AuthenticatorsApi.md#AuthenticatorsAdminDuoRetrieve) | **Get** /authenticators/admin/duo/{id}/ | 
[**AuthenticatorsAdminDuoUpdate**](AuthenticatorsApi.md#AuthenticatorsAdminDuoUpdate) | **Put** /authenticators/admin/duo/{id}/ | 
[**AuthenticatorsAdminStaticList**](AuthenticatorsApi.md#AuthenticatorsAdminStaticList) | **Get** /authenticators/admin/static/ | 
[**AuthenticatorsAdminStaticRetrieve**](AuthenticatorsApi.md#AuthenticatorsAdminStaticRetrieve) | **Get** /authenticators/admin/static/{id}/ | 
[**AuthenticatorsAdminTotpList**](AuthenticatorsApi.md#AuthenticatorsAdminTotpList) | **Get** /authenticators/admin/totp/ | 
[**AuthenticatorsAdminTotpRetrieve**](AuthenticatorsApi.md#AuthenticatorsAdminTotpRetrieve) | **Get** /authenticators/admin/totp/{id}/ | 
[**AuthenticatorsAdminWebauthnList**](AuthenticatorsApi.md#AuthenticatorsAdminWebauthnList) | **Get** /authenticators/admin/webauthn/ | 
[**AuthenticatorsAdminWebauthnRetrieve**](AuthenticatorsApi.md#AuthenticatorsAdminWebauthnRetrieve) | **Get** /authenticators/admin/webauthn/{id}/ | 
[**AuthenticatorsDuoDestroy**](AuthenticatorsApi.md#AuthenticatorsDuoDestroy) | **Delete** /authenticators/duo/{id}/ | 
[**AuthenticatorsDuoList**](AuthenticatorsApi.md#AuthenticatorsDuoList) | **Get** /authenticators/duo/ | 
[**AuthenticatorsDuoPartialUpdate**](AuthenticatorsApi.md#AuthenticatorsDuoPartialUpdate) | **Patch** /authenticators/duo/{id}/ | 
[**AuthenticatorsDuoRetrieve**](AuthenticatorsApi.md#AuthenticatorsDuoRetrieve) | **Get** /authenticators/duo/{id}/ | 
[**AuthenticatorsDuoUpdate**](AuthenticatorsApi.md#AuthenticatorsDuoUpdate) | **Put** /authenticators/duo/{id}/ | 
[**AuthenticatorsDuoUsedByList**](AuthenticatorsApi.md#AuthenticatorsDuoUsedByList) | **Get** /authenticators/duo/{id}/used_by/ | 
[**AuthenticatorsStaticDestroy**](AuthenticatorsApi.md#AuthenticatorsStaticDestroy) | **Delete** /authenticators/static/{id}/ | 
[**AuthenticatorsStaticList**](AuthenticatorsApi.md#AuthenticatorsStaticList) | **Get** /authenticators/static/ | 
[**AuthenticatorsStaticPartialUpdate**](AuthenticatorsApi.md#AuthenticatorsStaticPartialUpdate) | **Patch** /authenticators/static/{id}/ | 
[**AuthenticatorsStaticRetrieve**](AuthenticatorsApi.md#AuthenticatorsStaticRetrieve) | **Get** /authenticators/static/{id}/ | 
[**AuthenticatorsStaticUpdate**](AuthenticatorsApi.md#AuthenticatorsStaticUpdate) | **Put** /authenticators/static/{id}/ | 
[**AuthenticatorsStaticUsedByList**](AuthenticatorsApi.md#AuthenticatorsStaticUsedByList) | **Get** /authenticators/static/{id}/used_by/ | 
[**AuthenticatorsTotpDestroy**](AuthenticatorsApi.md#AuthenticatorsTotpDestroy) | **Delete** /authenticators/totp/{id}/ | 
[**AuthenticatorsTotpList**](AuthenticatorsApi.md#AuthenticatorsTotpList) | **Get** /authenticators/totp/ | 
[**AuthenticatorsTotpPartialUpdate**](AuthenticatorsApi.md#AuthenticatorsTotpPartialUpdate) | **Patch** /authenticators/totp/{id}/ | 
[**AuthenticatorsTotpRetrieve**](AuthenticatorsApi.md#AuthenticatorsTotpRetrieve) | **Get** /authenticators/totp/{id}/ | 
[**AuthenticatorsTotpUpdate**](AuthenticatorsApi.md#AuthenticatorsTotpUpdate) | **Put** /authenticators/totp/{id}/ | 
[**AuthenticatorsTotpUsedByList**](AuthenticatorsApi.md#AuthenticatorsTotpUsedByList) | **Get** /authenticators/totp/{id}/used_by/ | 
[**AuthenticatorsWebauthnDestroy**](AuthenticatorsApi.md#AuthenticatorsWebauthnDestroy) | **Delete** /authenticators/webauthn/{id}/ | 
[**AuthenticatorsWebauthnList**](AuthenticatorsApi.md#AuthenticatorsWebauthnList) | **Get** /authenticators/webauthn/ | 
[**AuthenticatorsWebauthnPartialUpdate**](AuthenticatorsApi.md#AuthenticatorsWebauthnPartialUpdate) | **Patch** /authenticators/webauthn/{id}/ | 
[**AuthenticatorsWebauthnRetrieve**](AuthenticatorsApi.md#AuthenticatorsWebauthnRetrieve) | **Get** /authenticators/webauthn/{id}/ | 
[**AuthenticatorsWebauthnUpdate**](AuthenticatorsApi.md#AuthenticatorsWebauthnUpdate) | **Put** /authenticators/webauthn/{id}/ | 
[**AuthenticatorsWebauthnUsedByList**](AuthenticatorsApi.md#AuthenticatorsWebauthnUsedByList) | **Get** /authenticators/webauthn/{id}/used_by/ | 



## AuthenticatorsAdminDuoCreate

> DuoDevice AuthenticatorsAdminDuoCreate(ctx).DuoDeviceRequest(duoDeviceRequest).Execute()





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
    duoDeviceRequest := *openapiclient.NewDuoDeviceRequest("Name_example") // DuoDeviceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsAdminDuoCreate(context.Background()).DuoDeviceRequest(duoDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsAdminDuoCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsAdminDuoCreate`: DuoDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsAdminDuoCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsAdminDuoCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **duoDeviceRequest** | [**DuoDeviceRequest**](DuoDeviceRequest.md) |  | 

### Return type

[**DuoDevice**](DuoDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsAdminDuoDestroy

> AuthenticatorsAdminDuoDestroy(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Duo Device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsAdminDuoDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsAdminDuoDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Duo Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsAdminDuoDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsAdminDuoList

> PaginatedDuoDeviceList AuthenticatorsAdminDuoList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsAdminDuoList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsAdminDuoList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsAdminDuoList`: PaginatedDuoDeviceList
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsAdminDuoList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsAdminDuoListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedDuoDeviceList**](PaginatedDuoDeviceList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsAdminDuoPartialUpdate

> DuoDevice AuthenticatorsAdminDuoPartialUpdate(ctx, id).PatchedDuoDeviceRequest(patchedDuoDeviceRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Duo Device.
    patchedDuoDeviceRequest := *openapiclient.NewPatchedDuoDeviceRequest() // PatchedDuoDeviceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsAdminDuoPartialUpdate(context.Background(), id).PatchedDuoDeviceRequest(patchedDuoDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsAdminDuoPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsAdminDuoPartialUpdate`: DuoDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsAdminDuoPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Duo Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsAdminDuoPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDuoDeviceRequest** | [**PatchedDuoDeviceRequest**](PatchedDuoDeviceRequest.md) |  | 

### Return type

[**DuoDevice**](DuoDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsAdminDuoRetrieve

> DuoDevice AuthenticatorsAdminDuoRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Duo Device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsAdminDuoRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsAdminDuoRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsAdminDuoRetrieve`: DuoDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsAdminDuoRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Duo Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsAdminDuoRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DuoDevice**](DuoDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsAdminDuoUpdate

> DuoDevice AuthenticatorsAdminDuoUpdate(ctx, id).DuoDeviceRequest(duoDeviceRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Duo Device.
    duoDeviceRequest := *openapiclient.NewDuoDeviceRequest("Name_example") // DuoDeviceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsAdminDuoUpdate(context.Background(), id).DuoDeviceRequest(duoDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsAdminDuoUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsAdminDuoUpdate`: DuoDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsAdminDuoUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Duo Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsAdminDuoUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duoDeviceRequest** | [**DuoDeviceRequest**](DuoDeviceRequest.md) |  | 

### Return type

[**DuoDevice**](DuoDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsAdminStaticList

> PaginatedStaticDeviceList AuthenticatorsAdminStaticList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsAdminStaticList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsAdminStaticList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsAdminStaticList`: PaginatedStaticDeviceList
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsAdminStaticList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsAdminStaticListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedStaticDeviceList**](PaginatedStaticDeviceList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsAdminStaticRetrieve

> StaticDevice AuthenticatorsAdminStaticRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this static device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsAdminStaticRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsAdminStaticRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsAdminStaticRetrieve`: StaticDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsAdminStaticRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this static device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsAdminStaticRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StaticDevice**](StaticDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsAdminTotpList

> PaginatedTOTPDeviceList AuthenticatorsAdminTotpList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsAdminTotpList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsAdminTotpList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsAdminTotpList`: PaginatedTOTPDeviceList
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsAdminTotpList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsAdminTotpListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedTOTPDeviceList**](PaginatedTOTPDeviceList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsAdminTotpRetrieve

> TOTPDevice AuthenticatorsAdminTotpRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this TOTP device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsAdminTotpRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsAdminTotpRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsAdminTotpRetrieve`: TOTPDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsAdminTotpRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this TOTP device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsAdminTotpRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TOTPDevice**](TOTPDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsAdminWebauthnList

> PaginatedWebAuthnDeviceList AuthenticatorsAdminWebauthnList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsAdminWebauthnList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsAdminWebauthnList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsAdminWebauthnList`: PaginatedWebAuthnDeviceList
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsAdminWebauthnList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsAdminWebauthnListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedWebAuthnDeviceList**](PaginatedWebAuthnDeviceList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsAdminWebauthnRetrieve

> WebAuthnDevice AuthenticatorsAdminWebauthnRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this WebAuthn Device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsAdminWebauthnRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsAdminWebauthnRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsAdminWebauthnRetrieve`: WebAuthnDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsAdminWebauthnRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this WebAuthn Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsAdminWebauthnRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebAuthnDevice**](WebAuthnDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsDuoDestroy

> AuthenticatorsDuoDestroy(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Duo Device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsDuoDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsDuoDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Duo Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsDuoDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsDuoList

> PaginatedDuoDeviceList AuthenticatorsDuoList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsDuoList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsDuoList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsDuoList`: PaginatedDuoDeviceList
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsDuoList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsDuoListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedDuoDeviceList**](PaginatedDuoDeviceList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsDuoPartialUpdate

> DuoDevice AuthenticatorsDuoPartialUpdate(ctx, id).PatchedDuoDeviceRequest(patchedDuoDeviceRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Duo Device.
    patchedDuoDeviceRequest := *openapiclient.NewPatchedDuoDeviceRequest() // PatchedDuoDeviceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsDuoPartialUpdate(context.Background(), id).PatchedDuoDeviceRequest(patchedDuoDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsDuoPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsDuoPartialUpdate`: DuoDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsDuoPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Duo Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsDuoPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDuoDeviceRequest** | [**PatchedDuoDeviceRequest**](PatchedDuoDeviceRequest.md) |  | 

### Return type

[**DuoDevice**](DuoDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsDuoRetrieve

> DuoDevice AuthenticatorsDuoRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Duo Device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsDuoRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsDuoRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsDuoRetrieve`: DuoDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsDuoRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Duo Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsDuoRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DuoDevice**](DuoDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsDuoUpdate

> DuoDevice AuthenticatorsDuoUpdate(ctx, id).DuoDeviceRequest(duoDeviceRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Duo Device.
    duoDeviceRequest := *openapiclient.NewDuoDeviceRequest("Name_example") // DuoDeviceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsDuoUpdate(context.Background(), id).DuoDeviceRequest(duoDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsDuoUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsDuoUpdate`: DuoDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsDuoUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Duo Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsDuoUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duoDeviceRequest** | [**DuoDeviceRequest**](DuoDeviceRequest.md) |  | 

### Return type

[**DuoDevice**](DuoDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsDuoUsedByList

> []UsedBy AuthenticatorsDuoUsedByList(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Duo Device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsDuoUsedByList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsDuoUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsDuoUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsDuoUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Duo Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsDuoUsedByListRequest struct via the builder pattern


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


## AuthenticatorsStaticDestroy

> AuthenticatorsStaticDestroy(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this static device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsStaticDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsStaticDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this static device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsStaticDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsStaticList

> PaginatedStaticDeviceList AuthenticatorsStaticList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsStaticList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsStaticList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsStaticList`: PaginatedStaticDeviceList
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsStaticList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsStaticListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedStaticDeviceList**](PaginatedStaticDeviceList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsStaticPartialUpdate

> StaticDevice AuthenticatorsStaticPartialUpdate(ctx, id).PatchedStaticDeviceRequest(patchedStaticDeviceRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this static device.
    patchedStaticDeviceRequest := *openapiclient.NewPatchedStaticDeviceRequest() // PatchedStaticDeviceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsStaticPartialUpdate(context.Background(), id).PatchedStaticDeviceRequest(patchedStaticDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsStaticPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsStaticPartialUpdate`: StaticDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsStaticPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this static device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsStaticPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedStaticDeviceRequest** | [**PatchedStaticDeviceRequest**](PatchedStaticDeviceRequest.md) |  | 

### Return type

[**StaticDevice**](StaticDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsStaticRetrieve

> StaticDevice AuthenticatorsStaticRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this static device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsStaticRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsStaticRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsStaticRetrieve`: StaticDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsStaticRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this static device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsStaticRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StaticDevice**](StaticDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsStaticUpdate

> StaticDevice AuthenticatorsStaticUpdate(ctx, id).StaticDeviceRequest(staticDeviceRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this static device.
    staticDeviceRequest := *openapiclient.NewStaticDeviceRequest("Name_example") // StaticDeviceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsStaticUpdate(context.Background(), id).StaticDeviceRequest(staticDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsStaticUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsStaticUpdate`: StaticDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsStaticUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this static device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsStaticUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **staticDeviceRequest** | [**StaticDeviceRequest**](StaticDeviceRequest.md) |  | 

### Return type

[**StaticDevice**](StaticDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsStaticUsedByList

> []UsedBy AuthenticatorsStaticUsedByList(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this static device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsStaticUsedByList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsStaticUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsStaticUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsStaticUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this static device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsStaticUsedByListRequest struct via the builder pattern


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


## AuthenticatorsTotpDestroy

> AuthenticatorsTotpDestroy(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this TOTP device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsTotpDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsTotpDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this TOTP device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsTotpDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsTotpList

> PaginatedTOTPDeviceList AuthenticatorsTotpList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsTotpList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsTotpList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsTotpList`: PaginatedTOTPDeviceList
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsTotpList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsTotpListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedTOTPDeviceList**](PaginatedTOTPDeviceList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsTotpPartialUpdate

> TOTPDevice AuthenticatorsTotpPartialUpdate(ctx, id).PatchedTOTPDeviceRequest(patchedTOTPDeviceRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this TOTP device.
    patchedTOTPDeviceRequest := *openapiclient.NewPatchedTOTPDeviceRequest() // PatchedTOTPDeviceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsTotpPartialUpdate(context.Background(), id).PatchedTOTPDeviceRequest(patchedTOTPDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsTotpPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsTotpPartialUpdate`: TOTPDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsTotpPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this TOTP device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsTotpPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedTOTPDeviceRequest** | [**PatchedTOTPDeviceRequest**](PatchedTOTPDeviceRequest.md) |  | 

### Return type

[**TOTPDevice**](TOTPDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsTotpRetrieve

> TOTPDevice AuthenticatorsTotpRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this TOTP device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsTotpRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsTotpRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsTotpRetrieve`: TOTPDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsTotpRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this TOTP device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsTotpRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TOTPDevice**](TOTPDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsTotpUpdate

> TOTPDevice AuthenticatorsTotpUpdate(ctx, id).TOTPDeviceRequest(tOTPDeviceRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this TOTP device.
    tOTPDeviceRequest := *openapiclient.NewTOTPDeviceRequest("Name_example") // TOTPDeviceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsTotpUpdate(context.Background(), id).TOTPDeviceRequest(tOTPDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsTotpUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsTotpUpdate`: TOTPDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsTotpUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this TOTP device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsTotpUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tOTPDeviceRequest** | [**TOTPDeviceRequest**](TOTPDeviceRequest.md) |  | 

### Return type

[**TOTPDevice**](TOTPDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsTotpUsedByList

> []UsedBy AuthenticatorsTotpUsedByList(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this TOTP device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsTotpUsedByList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsTotpUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsTotpUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsTotpUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this TOTP device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsTotpUsedByListRequest struct via the builder pattern


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


## AuthenticatorsWebauthnDestroy

> AuthenticatorsWebauthnDestroy(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this WebAuthn Device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsWebauthnDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsWebauthnDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this WebAuthn Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsWebauthnDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsWebauthnList

> PaginatedWebAuthnDeviceList AuthenticatorsWebauthnList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsWebauthnList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsWebauthnList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsWebauthnList`: PaginatedWebAuthnDeviceList
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsWebauthnList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsWebauthnListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedWebAuthnDeviceList**](PaginatedWebAuthnDeviceList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsWebauthnPartialUpdate

> WebAuthnDevice AuthenticatorsWebauthnPartialUpdate(ctx, id).PatchedWebAuthnDeviceRequest(patchedWebAuthnDeviceRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this WebAuthn Device.
    patchedWebAuthnDeviceRequest := *openapiclient.NewPatchedWebAuthnDeviceRequest() // PatchedWebAuthnDeviceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsWebauthnPartialUpdate(context.Background(), id).PatchedWebAuthnDeviceRequest(patchedWebAuthnDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsWebauthnPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsWebauthnPartialUpdate`: WebAuthnDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsWebauthnPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this WebAuthn Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsWebauthnPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWebAuthnDeviceRequest** | [**PatchedWebAuthnDeviceRequest**](PatchedWebAuthnDeviceRequest.md) |  | 

### Return type

[**WebAuthnDevice**](WebAuthnDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsWebauthnRetrieve

> WebAuthnDevice AuthenticatorsWebauthnRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this WebAuthn Device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsWebauthnRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsWebauthnRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsWebauthnRetrieve`: WebAuthnDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsWebauthnRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this WebAuthn Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsWebauthnRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebAuthnDevice**](WebAuthnDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsWebauthnUpdate

> WebAuthnDevice AuthenticatorsWebauthnUpdate(ctx, id).WebAuthnDeviceRequest(webAuthnDeviceRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this WebAuthn Device.
    webAuthnDeviceRequest := *openapiclient.NewWebAuthnDeviceRequest("Name_example") // WebAuthnDeviceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsWebauthnUpdate(context.Background(), id).WebAuthnDeviceRequest(webAuthnDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsWebauthnUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsWebauthnUpdate`: WebAuthnDevice
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsWebauthnUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this WebAuthn Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsWebauthnUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webAuthnDeviceRequest** | [**WebAuthnDeviceRequest**](WebAuthnDeviceRequest.md) |  | 

### Return type

[**WebAuthnDevice**](WebAuthnDevice.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticatorsWebauthnUsedByList

> []UsedBy AuthenticatorsWebauthnUsedByList(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this WebAuthn Device.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticatorsApi.AuthenticatorsWebauthnUsedByList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticatorsApi.AuthenticatorsWebauthnUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticatorsWebauthnUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `AuthenticatorsApi.AuthenticatorsWebauthnUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this WebAuthn Device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticatorsWebauthnUsedByListRequest struct via the builder pattern


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

