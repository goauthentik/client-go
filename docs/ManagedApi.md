# \ManagedApi

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ManagedBlueprintsApplyCreate**](ManagedApi.md#ManagedBlueprintsApplyCreate) | **Post** /managed/blueprints/{instance_uuid}/apply/ | 
[**ManagedBlueprintsAvailableList**](ManagedApi.md#ManagedBlueprintsAvailableList) | **Get** /managed/blueprints/available/ | 
[**ManagedBlueprintsCreate**](ManagedApi.md#ManagedBlueprintsCreate) | **Post** /managed/blueprints/ | 
[**ManagedBlueprintsDestroy**](ManagedApi.md#ManagedBlueprintsDestroy) | **Delete** /managed/blueprints/{instance_uuid}/ | 
[**ManagedBlueprintsList**](ManagedApi.md#ManagedBlueprintsList) | **Get** /managed/blueprints/ | 
[**ManagedBlueprintsPartialUpdate**](ManagedApi.md#ManagedBlueprintsPartialUpdate) | **Patch** /managed/blueprints/{instance_uuid}/ | 
[**ManagedBlueprintsRetrieve**](ManagedApi.md#ManagedBlueprintsRetrieve) | **Get** /managed/blueprints/{instance_uuid}/ | 
[**ManagedBlueprintsUpdate**](ManagedApi.md#ManagedBlueprintsUpdate) | **Put** /managed/blueprints/{instance_uuid}/ | 
[**ManagedBlueprintsUsedByList**](ManagedApi.md#ManagedBlueprintsUsedByList) | **Get** /managed/blueprints/{instance_uuid}/used_by/ | 



## ManagedBlueprintsApplyCreate

> BlueprintInstance ManagedBlueprintsApplyCreate(ctx, instanceUuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    instanceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Blueprint Instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedApi.ManagedBlueprintsApplyCreate(context.Background(), instanceUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.ManagedBlueprintsApplyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManagedBlueprintsApplyCreate`: BlueprintInstance
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.ManagedBlueprintsApplyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceUuid** | **string** | A UUID string identifying this Blueprint Instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedBlueprintsApplyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlueprintInstance**](BlueprintInstance.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedBlueprintsAvailableList

> []BlueprintFile ManagedBlueprintsAvailableList(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedApi.ManagedBlueprintsAvailableList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.ManagedBlueprintsAvailableList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManagedBlueprintsAvailableList`: []BlueprintFile
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.ManagedBlueprintsAvailableList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiManagedBlueprintsAvailableListRequest struct via the builder pattern


### Return type

[**[]BlueprintFile**](BlueprintFile.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedBlueprintsCreate

> BlueprintInstance ManagedBlueprintsCreate(ctx).BlueprintInstanceRequest(blueprintInstanceRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    blueprintInstanceRequest := *openapiclient.NewBlueprintInstanceRequest("Name_example") // BlueprintInstanceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedApi.ManagedBlueprintsCreate(context.Background()).BlueprintInstanceRequest(blueprintInstanceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.ManagedBlueprintsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManagedBlueprintsCreate`: BlueprintInstance
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.ManagedBlueprintsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManagedBlueprintsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintInstanceRequest** | [**BlueprintInstanceRequest**](BlueprintInstanceRequest.md) |  | 

### Return type

[**BlueprintInstance**](BlueprintInstance.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedBlueprintsDestroy

> ManagedBlueprintsDestroy(ctx, instanceUuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    instanceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Blueprint Instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedApi.ManagedBlueprintsDestroy(context.Background(), instanceUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.ManagedBlueprintsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceUuid** | **string** | A UUID string identifying this Blueprint Instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedBlueprintsDestroyRequest struct via the builder pattern


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


## ManagedBlueprintsList

> PaginatedBlueprintInstanceList ManagedBlueprintsList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Path(path).Search(search).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    path := "path_example" // string |  (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedApi.ManagedBlueprintsList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Path(path).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.ManagedBlueprintsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManagedBlueprintsList`: PaginatedBlueprintInstanceList
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.ManagedBlueprintsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManagedBlueprintsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **path** | **string** |  | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedBlueprintInstanceList**](PaginatedBlueprintInstanceList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedBlueprintsPartialUpdate

> BlueprintInstance ManagedBlueprintsPartialUpdate(ctx, instanceUuid).PatchedBlueprintInstanceRequest(patchedBlueprintInstanceRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    instanceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Blueprint Instance.
    patchedBlueprintInstanceRequest := *openapiclient.NewPatchedBlueprintInstanceRequest() // PatchedBlueprintInstanceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedApi.ManagedBlueprintsPartialUpdate(context.Background(), instanceUuid).PatchedBlueprintInstanceRequest(patchedBlueprintInstanceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.ManagedBlueprintsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManagedBlueprintsPartialUpdate`: BlueprintInstance
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.ManagedBlueprintsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceUuid** | **string** | A UUID string identifying this Blueprint Instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedBlueprintsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedBlueprintInstanceRequest** | [**PatchedBlueprintInstanceRequest**](PatchedBlueprintInstanceRequest.md) |  | 

### Return type

[**BlueprintInstance**](BlueprintInstance.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedBlueprintsRetrieve

> BlueprintInstance ManagedBlueprintsRetrieve(ctx, instanceUuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    instanceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Blueprint Instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedApi.ManagedBlueprintsRetrieve(context.Background(), instanceUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.ManagedBlueprintsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManagedBlueprintsRetrieve`: BlueprintInstance
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.ManagedBlueprintsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceUuid** | **string** | A UUID string identifying this Blueprint Instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedBlueprintsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlueprintInstance**](BlueprintInstance.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedBlueprintsUpdate

> BlueprintInstance ManagedBlueprintsUpdate(ctx, instanceUuid).BlueprintInstanceRequest(blueprintInstanceRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    instanceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Blueprint Instance.
    blueprintInstanceRequest := *openapiclient.NewBlueprintInstanceRequest("Name_example") // BlueprintInstanceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedApi.ManagedBlueprintsUpdate(context.Background(), instanceUuid).BlueprintInstanceRequest(blueprintInstanceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.ManagedBlueprintsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManagedBlueprintsUpdate`: BlueprintInstance
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.ManagedBlueprintsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceUuid** | **string** | A UUID string identifying this Blueprint Instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedBlueprintsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blueprintInstanceRequest** | [**BlueprintInstanceRequest**](BlueprintInstanceRequest.md) |  | 

### Return type

[**BlueprintInstance**](BlueprintInstance.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedBlueprintsUsedByList

> []UsedBy ManagedBlueprintsUsedByList(ctx, instanceUuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    instanceUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Blueprint Instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedApi.ManagedBlueprintsUsedByList(context.Background(), instanceUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedApi.ManagedBlueprintsUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManagedBlueprintsUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `ManagedApi.ManagedBlueprintsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceUuid** | **string** | A UUID string identifying this Blueprint Instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedBlueprintsUsedByListRequest struct via the builder pattern


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

