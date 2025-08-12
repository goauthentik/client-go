# \TenantsApi

All URIs are relative to */api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenantsDomainsCreate**](TenantsApi.md#TenantsDomainsCreate) | **Post** /tenants/domains/ | 
[**TenantsDomainsDestroy**](TenantsApi.md#TenantsDomainsDestroy) | **Delete** /tenants/domains/{id}/ | 
[**TenantsDomainsList**](TenantsApi.md#TenantsDomainsList) | **Get** /tenants/domains/ | 
[**TenantsDomainsPartialUpdate**](TenantsApi.md#TenantsDomainsPartialUpdate) | **Patch** /tenants/domains/{id}/ | 
[**TenantsDomainsRetrieve**](TenantsApi.md#TenantsDomainsRetrieve) | **Get** /tenants/domains/{id}/ | 
[**TenantsDomainsUpdate**](TenantsApi.md#TenantsDomainsUpdate) | **Put** /tenants/domains/{id}/ | 
[**TenantsTenantsCreate**](TenantsApi.md#TenantsTenantsCreate) | **Post** /tenants/tenants/ | 
[**TenantsTenantsCreateAdminGroupCreate**](TenantsApi.md#TenantsTenantsCreateAdminGroupCreate) | **Post** /tenants/tenants/{tenant_uuid}/create_admin_group/ | 
[**TenantsTenantsCreateRecoveryKeyCreate**](TenantsApi.md#TenantsTenantsCreateRecoveryKeyCreate) | **Post** /tenants/tenants/{tenant_uuid}/create_recovery_key/ | 
[**TenantsTenantsDestroy**](TenantsApi.md#TenantsTenantsDestroy) | **Delete** /tenants/tenants/{tenant_uuid}/ | 
[**TenantsTenantsList**](TenantsApi.md#TenantsTenantsList) | **Get** /tenants/tenants/ | 
[**TenantsTenantsPartialUpdate**](TenantsApi.md#TenantsTenantsPartialUpdate) | **Patch** /tenants/tenants/{tenant_uuid}/ | 
[**TenantsTenantsRetrieve**](TenantsApi.md#TenantsTenantsRetrieve) | **Get** /tenants/tenants/{tenant_uuid}/ | 
[**TenantsTenantsUpdate**](TenantsApi.md#TenantsTenantsUpdate) | **Put** /tenants/tenants/{tenant_uuid}/ | 



## TenantsDomainsCreate

> Domain TenantsDomainsCreate(ctx).DomainRequest(domainRequest).Execute()





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
    domainRequest := *openapiclient.NewDomainRequest("Domain_example", "Tenant_example") // DomainRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantsDomainsCreate(context.Background()).DomainRequest(domainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsDomainsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsDomainsCreate`: Domain
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantsDomainsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantsDomainsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainRequest** | [**DomainRequest**](DomainRequest.md) |  | 

### Return type

[**Domain**](Domain.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsDomainsDestroy

> TenantsDomainsDestroy(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Domain.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantsDomainsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsDomainsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantsDomainsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsDomainsList

> PaginatedDomainList TenantsDomainsList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantsDomainsList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsDomainsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsDomainsList`: PaginatedDomainList
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantsDomainsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantsDomainsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedDomainList**](PaginatedDomainList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsDomainsPartialUpdate

> Domain TenantsDomainsPartialUpdate(ctx, id).PatchedDomainRequest(patchedDomainRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Domain.
    patchedDomainRequest := *openapiclient.NewPatchedDomainRequest() // PatchedDomainRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantsDomainsPartialUpdate(context.Background(), id).PatchedDomainRequest(patchedDomainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsDomainsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsDomainsPartialUpdate`: Domain
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantsDomainsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantsDomainsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDomainRequest** | [**PatchedDomainRequest**](PatchedDomainRequest.md) |  | 

### Return type

[**Domain**](Domain.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsDomainsRetrieve

> Domain TenantsDomainsRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Domain.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantsDomainsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsDomainsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsDomainsRetrieve`: Domain
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantsDomainsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantsDomainsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Domain**](Domain.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsDomainsUpdate

> Domain TenantsDomainsUpdate(ctx, id).DomainRequest(domainRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Domain.
    domainRequest := *openapiclient.NewDomainRequest("Domain_example", "Tenant_example") // DomainRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantsDomainsUpdate(context.Background(), id).DomainRequest(domainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsDomainsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsDomainsUpdate`: Domain
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantsDomainsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantsDomainsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainRequest** | [**DomainRequest**](DomainRequest.md) |  | 

### Return type

[**Domain**](Domain.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsTenantsCreate

> Tenant TenantsTenantsCreate(ctx).TenantRequest(tenantRequest).Execute()





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
    tenantRequest := *openapiclient.NewTenantRequest("SchemaName_example", "Name_example") // TenantRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantsTenantsCreate(context.Background()).TenantRequest(tenantRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsTenantsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsTenantsCreate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantsTenantsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantsTenantsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantRequest** | [**TenantRequest**](TenantRequest.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsTenantsCreateAdminGroupCreate

> TenantsTenantsCreateAdminGroupCreate(ctx, tenantUuid).TenantAdminGroupRequestRequest(tenantAdminGroupRequestRequest).Execute()





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
    tenantUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Tenant.
    tenantAdminGroupRequestRequest := *openapiclient.NewTenantAdminGroupRequestRequest("User_example") // TenantAdminGroupRequestRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantsTenantsCreateAdminGroupCreate(context.Background(), tenantUuid).TenantAdminGroupRequestRequest(tenantAdminGroupRequestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsTenantsCreateAdminGroupCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantUuid** | **string** | A UUID string identifying this Tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantsTenantsCreateAdminGroupCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantAdminGroupRequestRequest** | [**TenantAdminGroupRequestRequest**](TenantAdminGroupRequestRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsTenantsCreateRecoveryKeyCreate

> TenantRecoveryKeyResponse TenantsTenantsCreateRecoveryKeyCreate(ctx, tenantUuid).TenantRecoveryKeyRequestRequest(tenantRecoveryKeyRequestRequest).Execute()





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
    tenantUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Tenant.
    tenantRecoveryKeyRequestRequest := *openapiclient.NewTenantRecoveryKeyRequestRequest("User_example", int32(123)) // TenantRecoveryKeyRequestRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantsTenantsCreateRecoveryKeyCreate(context.Background(), tenantUuid).TenantRecoveryKeyRequestRequest(tenantRecoveryKeyRequestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsTenantsCreateRecoveryKeyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsTenantsCreateRecoveryKeyCreate`: TenantRecoveryKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantsTenantsCreateRecoveryKeyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantUuid** | **string** | A UUID string identifying this Tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantsTenantsCreateRecoveryKeyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantRecoveryKeyRequestRequest** | [**TenantRecoveryKeyRequestRequest**](TenantRecoveryKeyRequestRequest.md) |  | 

### Return type

[**TenantRecoveryKeyResponse**](TenantRecoveryKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsTenantsDestroy

> TenantsTenantsDestroy(ctx, tenantUuid).Execute()





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
    tenantUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Tenant.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantsTenantsDestroy(context.Background(), tenantUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsTenantsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantUuid** | **string** | A UUID string identifying this Tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantsTenantsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsTenantsList

> PaginatedTenantList TenantsTenantsList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantsTenantsList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsTenantsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsTenantsList`: PaginatedTenantList
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantsTenantsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantsTenantsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedTenantList**](PaginatedTenantList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsTenantsPartialUpdate

> Tenant TenantsTenantsPartialUpdate(ctx, tenantUuid).PatchedTenantRequest(patchedTenantRequest).Execute()





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
    tenantUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Tenant.
    patchedTenantRequest := *openapiclient.NewPatchedTenantRequest() // PatchedTenantRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantsTenantsPartialUpdate(context.Background(), tenantUuid).PatchedTenantRequest(patchedTenantRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsTenantsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsTenantsPartialUpdate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantsTenantsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantUuid** | **string** | A UUID string identifying this Tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantsTenantsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedTenantRequest** | [**PatchedTenantRequest**](PatchedTenantRequest.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsTenantsRetrieve

> Tenant TenantsTenantsRetrieve(ctx, tenantUuid).Execute()





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
    tenantUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Tenant.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantsTenantsRetrieve(context.Background(), tenantUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsTenantsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsTenantsRetrieve`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantsTenantsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantUuid** | **string** | A UUID string identifying this Tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantsTenantsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsTenantsUpdate

> Tenant TenantsTenantsUpdate(ctx, tenantUuid).TenantRequest(tenantRequest).Execute()





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
    tenantUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Tenant.
    tenantRequest := *openapiclient.NewTenantRequest("SchemaName_example", "Name_example") // TenantRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantsApi.TenantsTenantsUpdate(context.Background(), tenantUuid).TenantRequest(tenantRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.TenantsTenantsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantsTenantsUpdate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.TenantsTenantsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantUuid** | **string** | A UUID string identifying this Tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantsTenantsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantRequest** | [**TenantRequest**](TenantRequest.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

