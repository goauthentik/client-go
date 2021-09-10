# \OutpostsApi

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OutpostsInstancesCreate**](OutpostsApi.md#OutpostsInstancesCreate) | **Post** /outposts/instances/ | 
[**OutpostsInstancesDefaultSettingsRetrieve**](OutpostsApi.md#OutpostsInstancesDefaultSettingsRetrieve) | **Get** /outposts/instances/default_settings/ | 
[**OutpostsInstancesDestroy**](OutpostsApi.md#OutpostsInstancesDestroy) | **Delete** /outposts/instances/{uuid}/ | 
[**OutpostsInstancesHealthList**](OutpostsApi.md#OutpostsInstancesHealthList) | **Get** /outposts/instances/{uuid}/health/ | 
[**OutpostsInstancesList**](OutpostsApi.md#OutpostsInstancesList) | **Get** /outposts/instances/ | 
[**OutpostsInstancesPartialUpdate**](OutpostsApi.md#OutpostsInstancesPartialUpdate) | **Patch** /outposts/instances/{uuid}/ | 
[**OutpostsInstancesRetrieve**](OutpostsApi.md#OutpostsInstancesRetrieve) | **Get** /outposts/instances/{uuid}/ | 
[**OutpostsInstancesUpdate**](OutpostsApi.md#OutpostsInstancesUpdate) | **Put** /outposts/instances/{uuid}/ | 
[**OutpostsInstancesUsedByList**](OutpostsApi.md#OutpostsInstancesUsedByList) | **Get** /outposts/instances/{uuid}/used_by/ | 
[**OutpostsLdapList**](OutpostsApi.md#OutpostsLdapList) | **Get** /outposts/ldap/ | 
[**OutpostsLdapRetrieve**](OutpostsApi.md#OutpostsLdapRetrieve) | **Get** /outposts/ldap/{id}/ | 
[**OutpostsProxyList**](OutpostsApi.md#OutpostsProxyList) | **Get** /outposts/proxy/ | 
[**OutpostsProxyRetrieve**](OutpostsApi.md#OutpostsProxyRetrieve) | **Get** /outposts/proxy/{id}/ | 
[**OutpostsServiceConnectionsAllDestroy**](OutpostsApi.md#OutpostsServiceConnectionsAllDestroy) | **Delete** /outposts/service_connections/all/{uuid}/ | 
[**OutpostsServiceConnectionsAllList**](OutpostsApi.md#OutpostsServiceConnectionsAllList) | **Get** /outposts/service_connections/all/ | 
[**OutpostsServiceConnectionsAllRetrieve**](OutpostsApi.md#OutpostsServiceConnectionsAllRetrieve) | **Get** /outposts/service_connections/all/{uuid}/ | 
[**OutpostsServiceConnectionsAllStateRetrieve**](OutpostsApi.md#OutpostsServiceConnectionsAllStateRetrieve) | **Get** /outposts/service_connections/all/{uuid}/state/ | 
[**OutpostsServiceConnectionsAllTypesList**](OutpostsApi.md#OutpostsServiceConnectionsAllTypesList) | **Get** /outposts/service_connections/all/types/ | 
[**OutpostsServiceConnectionsAllUsedByList**](OutpostsApi.md#OutpostsServiceConnectionsAllUsedByList) | **Get** /outposts/service_connections/all/{uuid}/used_by/ | 
[**OutpostsServiceConnectionsDockerCreate**](OutpostsApi.md#OutpostsServiceConnectionsDockerCreate) | **Post** /outposts/service_connections/docker/ | 
[**OutpostsServiceConnectionsDockerDestroy**](OutpostsApi.md#OutpostsServiceConnectionsDockerDestroy) | **Delete** /outposts/service_connections/docker/{uuid}/ | 
[**OutpostsServiceConnectionsDockerList**](OutpostsApi.md#OutpostsServiceConnectionsDockerList) | **Get** /outposts/service_connections/docker/ | 
[**OutpostsServiceConnectionsDockerPartialUpdate**](OutpostsApi.md#OutpostsServiceConnectionsDockerPartialUpdate) | **Patch** /outposts/service_connections/docker/{uuid}/ | 
[**OutpostsServiceConnectionsDockerRetrieve**](OutpostsApi.md#OutpostsServiceConnectionsDockerRetrieve) | **Get** /outposts/service_connections/docker/{uuid}/ | 
[**OutpostsServiceConnectionsDockerUpdate**](OutpostsApi.md#OutpostsServiceConnectionsDockerUpdate) | **Put** /outposts/service_connections/docker/{uuid}/ | 
[**OutpostsServiceConnectionsDockerUsedByList**](OutpostsApi.md#OutpostsServiceConnectionsDockerUsedByList) | **Get** /outposts/service_connections/docker/{uuid}/used_by/ | 
[**OutpostsServiceConnectionsKubernetesCreate**](OutpostsApi.md#OutpostsServiceConnectionsKubernetesCreate) | **Post** /outposts/service_connections/kubernetes/ | 
[**OutpostsServiceConnectionsKubernetesDestroy**](OutpostsApi.md#OutpostsServiceConnectionsKubernetesDestroy) | **Delete** /outposts/service_connections/kubernetes/{uuid}/ | 
[**OutpostsServiceConnectionsKubernetesList**](OutpostsApi.md#OutpostsServiceConnectionsKubernetesList) | **Get** /outposts/service_connections/kubernetes/ | 
[**OutpostsServiceConnectionsKubernetesPartialUpdate**](OutpostsApi.md#OutpostsServiceConnectionsKubernetesPartialUpdate) | **Patch** /outposts/service_connections/kubernetes/{uuid}/ | 
[**OutpostsServiceConnectionsKubernetesRetrieve**](OutpostsApi.md#OutpostsServiceConnectionsKubernetesRetrieve) | **Get** /outposts/service_connections/kubernetes/{uuid}/ | 
[**OutpostsServiceConnectionsKubernetesUpdate**](OutpostsApi.md#OutpostsServiceConnectionsKubernetesUpdate) | **Put** /outposts/service_connections/kubernetes/{uuid}/ | 
[**OutpostsServiceConnectionsKubernetesUsedByList**](OutpostsApi.md#OutpostsServiceConnectionsKubernetesUsedByList) | **Get** /outposts/service_connections/kubernetes/{uuid}/used_by/ | 



## OutpostsInstancesCreate

> Outpost OutpostsInstancesCreate(ctx).OutpostRequest(outpostRequest).Execute()





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
    outpostRequest := *openapiclient.NewOutpostRequest("Name_example", openapiclient.OutpostTypeEnum("proxy"), []int32{int32(123)}, map[string]interface{}{"key": interface{}(123)}) // OutpostRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsInstancesCreate(context.Background()).OutpostRequest(outpostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsInstancesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsInstancesCreate`: Outpost
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsInstancesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsInstancesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outpostRequest** | [**OutpostRequest**](OutpostRequest.md) |  | 

### Return type

[**Outpost**](Outpost.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsInstancesDefaultSettingsRetrieve

> OutpostDefaultConfig OutpostsInstancesDefaultSettingsRetrieve(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsInstancesDefaultSettingsRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsInstancesDefaultSettingsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsInstancesDefaultSettingsRetrieve`: OutpostDefaultConfig
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsInstancesDefaultSettingsRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsInstancesDefaultSettingsRetrieveRequest struct via the builder pattern


### Return type

[**OutpostDefaultConfig**](OutpostDefaultConfig.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsInstancesDestroy

> OutpostsInstancesDestroy(ctx, uuid).Execute()





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
    uuid := TODO // string | A UUID string identifying this outpost.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsInstancesDestroy(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsInstancesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this outpost. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsInstancesDestroyRequest struct via the builder pattern


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


## OutpostsInstancesHealthList

> []OutpostHealth OutpostsInstancesHealthList(ctx, uuid).NameIcontains(nameIcontains).NameIexact(nameIexact).Ordering(ordering).ProvidersIsnull(providersIsnull).Search(search).ServiceConnectionNameIcontains(serviceConnectionNameIcontains).ServiceConnectionNameIexact(serviceConnectionNameIexact).Execute()





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
    uuid := TODO // string | A UUID string identifying this outpost.
    nameIcontains := "nameIcontains_example" // string |  (optional)
    nameIexact := "nameIexact_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    providersIsnull := true // bool |  (optional)
    search := "search_example" // string | A search term. (optional)
    serviceConnectionNameIcontains := "serviceConnectionNameIcontains_example" // string |  (optional)
    serviceConnectionNameIexact := "serviceConnectionNameIexact_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsInstancesHealthList(context.Background(), uuid).NameIcontains(nameIcontains).NameIexact(nameIexact).Ordering(ordering).ProvidersIsnull(providersIsnull).Search(search).ServiceConnectionNameIcontains(serviceConnectionNameIcontains).ServiceConnectionNameIexact(serviceConnectionNameIexact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsInstancesHealthList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsInstancesHealthList`: []OutpostHealth
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsInstancesHealthList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this outpost. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsInstancesHealthListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nameIcontains** | **string** |  | 
 **nameIexact** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **providersIsnull** | **bool** |  | 
 **search** | **string** | A search term. | 
 **serviceConnectionNameIcontains** | **string** |  | 
 **serviceConnectionNameIexact** | **string** |  | 

### Return type

[**[]OutpostHealth**](OutpostHealth.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsInstancesList

> PaginatedOutpostList OutpostsInstancesList(ctx).NameIcontains(nameIcontains).NameIexact(nameIexact).Ordering(ordering).Page(page).PageSize(pageSize).ProvidersIsnull(providersIsnull).Search(search).ServiceConnectionNameIcontains(serviceConnectionNameIcontains).ServiceConnectionNameIexact(serviceConnectionNameIexact).Execute()





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
    nameIcontains := "nameIcontains_example" // string |  (optional)
    nameIexact := "nameIexact_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    providersIsnull := true // bool |  (optional)
    search := "search_example" // string | A search term. (optional)
    serviceConnectionNameIcontains := "serviceConnectionNameIcontains_example" // string |  (optional)
    serviceConnectionNameIexact := "serviceConnectionNameIexact_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsInstancesList(context.Background()).NameIcontains(nameIcontains).NameIexact(nameIexact).Ordering(ordering).Page(page).PageSize(pageSize).ProvidersIsnull(providersIsnull).Search(search).ServiceConnectionNameIcontains(serviceConnectionNameIcontains).ServiceConnectionNameIexact(serviceConnectionNameIexact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsInstancesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsInstancesList`: PaginatedOutpostList
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsInstancesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsInstancesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameIcontains** | **string** |  | 
 **nameIexact** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **providersIsnull** | **bool** |  | 
 **search** | **string** | A search term. | 
 **serviceConnectionNameIcontains** | **string** |  | 
 **serviceConnectionNameIexact** | **string** |  | 

### Return type

[**PaginatedOutpostList**](PaginatedOutpostList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsInstancesPartialUpdate

> Outpost OutpostsInstancesPartialUpdate(ctx, uuid).PatchedOutpostRequest(patchedOutpostRequest).Execute()





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
    uuid := TODO // string | A UUID string identifying this outpost.
    patchedOutpostRequest := *openapiclient.NewPatchedOutpostRequest() // PatchedOutpostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsInstancesPartialUpdate(context.Background(), uuid).PatchedOutpostRequest(patchedOutpostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsInstancesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsInstancesPartialUpdate`: Outpost
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsInstancesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this outpost. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsInstancesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedOutpostRequest** | [**PatchedOutpostRequest**](PatchedOutpostRequest.md) |  | 

### Return type

[**Outpost**](Outpost.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsInstancesRetrieve

> Outpost OutpostsInstancesRetrieve(ctx, uuid).Execute()





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
    uuid := TODO // string | A UUID string identifying this outpost.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsInstancesRetrieve(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsInstancesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsInstancesRetrieve`: Outpost
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsInstancesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this outpost. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsInstancesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Outpost**](Outpost.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsInstancesUpdate

> Outpost OutpostsInstancesUpdate(ctx, uuid).OutpostRequest(outpostRequest).Execute()





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
    uuid := TODO // string | A UUID string identifying this outpost.
    outpostRequest := *openapiclient.NewOutpostRequest("Name_example", openapiclient.OutpostTypeEnum("proxy"), []int32{int32(123)}, map[string]interface{}{"key": interface{}(123)}) // OutpostRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsInstancesUpdate(context.Background(), uuid).OutpostRequest(outpostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsInstancesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsInstancesUpdate`: Outpost
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsInstancesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this outpost. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsInstancesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **outpostRequest** | [**OutpostRequest**](OutpostRequest.md) |  | 

### Return type

[**Outpost**](Outpost.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsInstancesUsedByList

> []UsedBy OutpostsInstancesUsedByList(ctx, uuid).Execute()





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
    uuid := TODO // string | A UUID string identifying this outpost.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsInstancesUsedByList(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsInstancesUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsInstancesUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsInstancesUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this outpost. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsInstancesUsedByListRequest struct via the builder pattern


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


## OutpostsLdapList

> PaginatedLDAPOutpostConfigList OutpostsLdapList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsLdapList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsLdapList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsLdapList`: PaginatedLDAPOutpostConfigList
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsLdapList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsLdapListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedLDAPOutpostConfigList**](PaginatedLDAPOutpostConfigList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsLdapRetrieve

> LDAPOutpostConfig OutpostsLdapRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this LDAP Provider.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsLdapRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsLdapRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsLdapRetrieve`: LDAPOutpostConfig
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsLdapRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this LDAP Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsLdapRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LDAPOutpostConfig**](LDAPOutpostConfig.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsProxyList

> PaginatedProxyOutpostConfigList OutpostsProxyList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsProxyList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsProxyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsProxyList`: PaginatedProxyOutpostConfigList
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsProxyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsProxyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedProxyOutpostConfigList**](PaginatedProxyOutpostConfigList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsProxyRetrieve

> ProxyOutpostConfig OutpostsProxyRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Proxy Provider.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsProxyRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsProxyRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsProxyRetrieve`: ProxyOutpostConfig
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsProxyRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Proxy Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsProxyRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProxyOutpostConfig**](ProxyOutpostConfig.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsServiceConnectionsAllDestroy

> OutpostsServiceConnectionsAllDestroy(ctx, uuid).Execute()





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
    uuid := TODO // string | A UUID string identifying this Outpost Service-Connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsAllDestroy(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsAllDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this Outpost Service-Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsAllDestroyRequest struct via the builder pattern


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


## OutpostsServiceConnectionsAllList

> PaginatedServiceConnectionList OutpostsServiceConnectionsAllList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsAllList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsAllList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsServiceConnectionsAllList`: PaginatedServiceConnectionList
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsServiceConnectionsAllList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsAllListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedServiceConnectionList**](PaginatedServiceConnectionList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsServiceConnectionsAllRetrieve

> ServiceConnection OutpostsServiceConnectionsAllRetrieve(ctx, uuid).Execute()





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
    uuid := TODO // string | A UUID string identifying this Outpost Service-Connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsAllRetrieve(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsAllRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsServiceConnectionsAllRetrieve`: ServiceConnection
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsServiceConnectionsAllRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this Outpost Service-Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsAllRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceConnection**](ServiceConnection.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsServiceConnectionsAllStateRetrieve

> ServiceConnectionState OutpostsServiceConnectionsAllStateRetrieve(ctx, uuid).Execute()





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
    uuid := TODO // string | A UUID string identifying this Outpost Service-Connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsAllStateRetrieve(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsAllStateRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsServiceConnectionsAllStateRetrieve`: ServiceConnectionState
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsServiceConnectionsAllStateRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this Outpost Service-Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsAllStateRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceConnectionState**](ServiceConnectionState.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsServiceConnectionsAllTypesList

> []TypeCreate OutpostsServiceConnectionsAllTypesList(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsAllTypesList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsAllTypesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsServiceConnectionsAllTypesList`: []TypeCreate
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsServiceConnectionsAllTypesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsAllTypesListRequest struct via the builder pattern


### Return type

[**[]TypeCreate**](TypeCreate.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsServiceConnectionsAllUsedByList

> []UsedBy OutpostsServiceConnectionsAllUsedByList(ctx, uuid).Execute()





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
    uuid := TODO // string | A UUID string identifying this Outpost Service-Connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsAllUsedByList(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsAllUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsServiceConnectionsAllUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsServiceConnectionsAllUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this Outpost Service-Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsAllUsedByListRequest struct via the builder pattern


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


## OutpostsServiceConnectionsDockerCreate

> DockerServiceConnection OutpostsServiceConnectionsDockerCreate(ctx).DockerServiceConnectionRequest(dockerServiceConnectionRequest).Execute()





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
    dockerServiceConnectionRequest := *openapiclient.NewDockerServiceConnectionRequest("Name_example", "Url_example") // DockerServiceConnectionRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsDockerCreate(context.Background()).DockerServiceConnectionRequest(dockerServiceConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsDockerCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsServiceConnectionsDockerCreate`: DockerServiceConnection
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsServiceConnectionsDockerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsDockerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dockerServiceConnectionRequest** | [**DockerServiceConnectionRequest**](DockerServiceConnectionRequest.md) |  | 

### Return type

[**DockerServiceConnection**](DockerServiceConnection.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsServiceConnectionsDockerDestroy

> OutpostsServiceConnectionsDockerDestroy(ctx, uuid).Execute()





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
    uuid := TODO // string | A UUID string identifying this Docker Service-Connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsDockerDestroy(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsDockerDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this Docker Service-Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsDockerDestroyRequest struct via the builder pattern


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


## OutpostsServiceConnectionsDockerList

> PaginatedDockerServiceConnectionList OutpostsServiceConnectionsDockerList(ctx).Local(local).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).TlsAuthentication(tlsAuthentication).TlsVerification(tlsVerification).Url(url).Execute()





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
    local := true // bool |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    tlsAuthentication := TODO // string |  (optional)
    tlsVerification := TODO // string |  (optional)
    url := "url_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsDockerList(context.Background()).Local(local).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).TlsAuthentication(tlsAuthentication).TlsVerification(tlsVerification).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsDockerList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsServiceConnectionsDockerList`: PaginatedDockerServiceConnectionList
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsServiceConnectionsDockerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsDockerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **local** | **bool** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **tlsAuthentication** | [**string**](string.md) |  | 
 **tlsVerification** | [**string**](string.md) |  | 
 **url** | **string** |  | 

### Return type

[**PaginatedDockerServiceConnectionList**](PaginatedDockerServiceConnectionList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsServiceConnectionsDockerPartialUpdate

> DockerServiceConnection OutpostsServiceConnectionsDockerPartialUpdate(ctx, uuid).PatchedDockerServiceConnectionRequest(patchedDockerServiceConnectionRequest).Execute()





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
    uuid := TODO // string | A UUID string identifying this Docker Service-Connection.
    patchedDockerServiceConnectionRequest := *openapiclient.NewPatchedDockerServiceConnectionRequest() // PatchedDockerServiceConnectionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsDockerPartialUpdate(context.Background(), uuid).PatchedDockerServiceConnectionRequest(patchedDockerServiceConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsDockerPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsServiceConnectionsDockerPartialUpdate`: DockerServiceConnection
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsServiceConnectionsDockerPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this Docker Service-Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsDockerPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDockerServiceConnectionRequest** | [**PatchedDockerServiceConnectionRequest**](PatchedDockerServiceConnectionRequest.md) |  | 

### Return type

[**DockerServiceConnection**](DockerServiceConnection.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsServiceConnectionsDockerRetrieve

> DockerServiceConnection OutpostsServiceConnectionsDockerRetrieve(ctx, uuid).Execute()





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
    uuid := TODO // string | A UUID string identifying this Docker Service-Connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsDockerRetrieve(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsDockerRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsServiceConnectionsDockerRetrieve`: DockerServiceConnection
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsServiceConnectionsDockerRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this Docker Service-Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsDockerRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DockerServiceConnection**](DockerServiceConnection.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsServiceConnectionsDockerUpdate

> DockerServiceConnection OutpostsServiceConnectionsDockerUpdate(ctx, uuid).DockerServiceConnectionRequest(dockerServiceConnectionRequest).Execute()





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
    uuid := TODO // string | A UUID string identifying this Docker Service-Connection.
    dockerServiceConnectionRequest := *openapiclient.NewDockerServiceConnectionRequest("Name_example", "Url_example") // DockerServiceConnectionRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsDockerUpdate(context.Background(), uuid).DockerServiceConnectionRequest(dockerServiceConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsDockerUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsServiceConnectionsDockerUpdate`: DockerServiceConnection
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsServiceConnectionsDockerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this Docker Service-Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsDockerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dockerServiceConnectionRequest** | [**DockerServiceConnectionRequest**](DockerServiceConnectionRequest.md) |  | 

### Return type

[**DockerServiceConnection**](DockerServiceConnection.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsServiceConnectionsDockerUsedByList

> []UsedBy OutpostsServiceConnectionsDockerUsedByList(ctx, uuid).Execute()





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
    uuid := TODO // string | A UUID string identifying this Docker Service-Connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsDockerUsedByList(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsDockerUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsServiceConnectionsDockerUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsServiceConnectionsDockerUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this Docker Service-Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsDockerUsedByListRequest struct via the builder pattern


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


## OutpostsServiceConnectionsKubernetesCreate

> KubernetesServiceConnection OutpostsServiceConnectionsKubernetesCreate(ctx).KubernetesServiceConnectionRequest(kubernetesServiceConnectionRequest).Execute()





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
    kubernetesServiceConnectionRequest := *openapiclient.NewKubernetesServiceConnectionRequest("Name_example") // KubernetesServiceConnectionRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsKubernetesCreate(context.Background()).KubernetesServiceConnectionRequest(kubernetesServiceConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsKubernetesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsServiceConnectionsKubernetesCreate`: KubernetesServiceConnection
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsServiceConnectionsKubernetesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsKubernetesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesServiceConnectionRequest** | [**KubernetesServiceConnectionRequest**](KubernetesServiceConnectionRequest.md) |  | 

### Return type

[**KubernetesServiceConnection**](KubernetesServiceConnection.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsServiceConnectionsKubernetesDestroy

> OutpostsServiceConnectionsKubernetesDestroy(ctx, uuid).Execute()





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
    uuid := TODO // string | A UUID string identifying this Kubernetes Service-Connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsKubernetesDestroy(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsKubernetesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this Kubernetes Service-Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsKubernetesDestroyRequest struct via the builder pattern


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


## OutpostsServiceConnectionsKubernetesList

> PaginatedKubernetesServiceConnectionList OutpostsServiceConnectionsKubernetesList(ctx).Local(local).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    local := true // bool |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsKubernetesList(context.Background()).Local(local).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsKubernetesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsServiceConnectionsKubernetesList`: PaginatedKubernetesServiceConnectionList
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsServiceConnectionsKubernetesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsKubernetesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **local** | **bool** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedKubernetesServiceConnectionList**](PaginatedKubernetesServiceConnectionList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsServiceConnectionsKubernetesPartialUpdate

> KubernetesServiceConnection OutpostsServiceConnectionsKubernetesPartialUpdate(ctx, uuid).PatchedKubernetesServiceConnectionRequest(patchedKubernetesServiceConnectionRequest).Execute()





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
    uuid := TODO // string | A UUID string identifying this Kubernetes Service-Connection.
    patchedKubernetesServiceConnectionRequest := *openapiclient.NewPatchedKubernetesServiceConnectionRequest() // PatchedKubernetesServiceConnectionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsKubernetesPartialUpdate(context.Background(), uuid).PatchedKubernetesServiceConnectionRequest(patchedKubernetesServiceConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsKubernetesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsServiceConnectionsKubernetesPartialUpdate`: KubernetesServiceConnection
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsServiceConnectionsKubernetesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this Kubernetes Service-Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsKubernetesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedKubernetesServiceConnectionRequest** | [**PatchedKubernetesServiceConnectionRequest**](PatchedKubernetesServiceConnectionRequest.md) |  | 

### Return type

[**KubernetesServiceConnection**](KubernetesServiceConnection.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsServiceConnectionsKubernetesRetrieve

> KubernetesServiceConnection OutpostsServiceConnectionsKubernetesRetrieve(ctx, uuid).Execute()





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
    uuid := TODO // string | A UUID string identifying this Kubernetes Service-Connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsKubernetesRetrieve(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsKubernetesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsServiceConnectionsKubernetesRetrieve`: KubernetesServiceConnection
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsServiceConnectionsKubernetesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this Kubernetes Service-Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsKubernetesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KubernetesServiceConnection**](KubernetesServiceConnection.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsServiceConnectionsKubernetesUpdate

> KubernetesServiceConnection OutpostsServiceConnectionsKubernetesUpdate(ctx, uuid).KubernetesServiceConnectionRequest(kubernetesServiceConnectionRequest).Execute()





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
    uuid := TODO // string | A UUID string identifying this Kubernetes Service-Connection.
    kubernetesServiceConnectionRequest := *openapiclient.NewKubernetesServiceConnectionRequest("Name_example") // KubernetesServiceConnectionRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsKubernetesUpdate(context.Background(), uuid).KubernetesServiceConnectionRequest(kubernetesServiceConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsKubernetesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsServiceConnectionsKubernetesUpdate`: KubernetesServiceConnection
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsServiceConnectionsKubernetesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this Kubernetes Service-Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsKubernetesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesServiceConnectionRequest** | [**KubernetesServiceConnectionRequest**](KubernetesServiceConnectionRequest.md) |  | 

### Return type

[**KubernetesServiceConnection**](KubernetesServiceConnection.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsServiceConnectionsKubernetesUsedByList

> []UsedBy OutpostsServiceConnectionsKubernetesUsedByList(ctx, uuid).Execute()





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
    uuid := TODO // string | A UUID string identifying this Kubernetes Service-Connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpostsApi.OutpostsServiceConnectionsKubernetesUsedByList(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpostsApi.OutpostsServiceConnectionsKubernetesUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpostsServiceConnectionsKubernetesUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `OutpostsApi.OutpostsServiceConnectionsKubernetesUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | A UUID string identifying this Kubernetes Service-Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsServiceConnectionsKubernetesUsedByListRequest struct via the builder pattern


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

