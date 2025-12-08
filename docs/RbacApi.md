# \RbacApi

All URIs are relative to */api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RbacInitialPermissionsCreate**](RbacApi.md#RbacInitialPermissionsCreate) | **Post** /rbac/initial_permissions/ | 
[**RbacInitialPermissionsDestroy**](RbacApi.md#RbacInitialPermissionsDestroy) | **Delete** /rbac/initial_permissions/{id}/ | 
[**RbacInitialPermissionsList**](RbacApi.md#RbacInitialPermissionsList) | **Get** /rbac/initial_permissions/ | 
[**RbacInitialPermissionsPartialUpdate**](RbacApi.md#RbacInitialPermissionsPartialUpdate) | **Patch** /rbac/initial_permissions/{id}/ | 
[**RbacInitialPermissionsRetrieve**](RbacApi.md#RbacInitialPermissionsRetrieve) | **Get** /rbac/initial_permissions/{id}/ | 
[**RbacInitialPermissionsUpdate**](RbacApi.md#RbacInitialPermissionsUpdate) | **Put** /rbac/initial_permissions/{id}/ | 
[**RbacInitialPermissionsUsedByList**](RbacApi.md#RbacInitialPermissionsUsedByList) | **Get** /rbac/initial_permissions/{id}/used_by/ | 
[**RbacPermissionsAssignedByRolesAssign**](RbacApi.md#RbacPermissionsAssignedByRolesAssign) | **Post** /rbac/permissions/assigned_by_roles/{uuid}/assign/ | 
[**RbacPermissionsAssignedByRolesList**](RbacApi.md#RbacPermissionsAssignedByRolesList) | **Get** /rbac/permissions/assigned_by_roles/ | 
[**RbacPermissionsAssignedByRolesUnassignPartialUpdate**](RbacApi.md#RbacPermissionsAssignedByRolesUnassignPartialUpdate) | **Patch** /rbac/permissions/assigned_by_roles/{uuid}/unassign/ | 
[**RbacPermissionsList**](RbacApi.md#RbacPermissionsList) | **Get** /rbac/permissions/ | 
[**RbacPermissionsRetrieve**](RbacApi.md#RbacPermissionsRetrieve) | **Get** /rbac/permissions/{id}/ | 
[**RbacPermissionsRolesList**](RbacApi.md#RbacPermissionsRolesList) | **Get** /rbac/permissions/roles/ | 
[**RbacRolesAddUserCreate**](RbacApi.md#RbacRolesAddUserCreate) | **Post** /rbac/roles/{uuid}/add_user/ | 
[**RbacRolesCreate**](RbacApi.md#RbacRolesCreate) | **Post** /rbac/roles/ | 
[**RbacRolesDestroy**](RbacApi.md#RbacRolesDestroy) | **Delete** /rbac/roles/{uuid}/ | 
[**RbacRolesList**](RbacApi.md#RbacRolesList) | **Get** /rbac/roles/ | 
[**RbacRolesPartialUpdate**](RbacApi.md#RbacRolesPartialUpdate) | **Patch** /rbac/roles/{uuid}/ | 
[**RbacRolesRemoveUserCreate**](RbacApi.md#RbacRolesRemoveUserCreate) | **Post** /rbac/roles/{uuid}/remove_user/ | 
[**RbacRolesRetrieve**](RbacApi.md#RbacRolesRetrieve) | **Get** /rbac/roles/{uuid}/ | 
[**RbacRolesUpdate**](RbacApi.md#RbacRolesUpdate) | **Put** /rbac/roles/{uuid}/ | 
[**RbacRolesUsedByList**](RbacApi.md#RbacRolesUsedByList) | **Get** /rbac/roles/{uuid}/used_by/ | 



## RbacInitialPermissionsCreate

> InitialPermissions RbacInitialPermissionsCreate(ctx).InitialPermissionsRequest(initialPermissionsRequest).Execute()





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
    initialPermissionsRequest := *openapiclient.NewInitialPermissionsRequest("Name_example", "Role_example") // InitialPermissionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacInitialPermissionsCreate(context.Background()).InitialPermissionsRequest(initialPermissionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacInitialPermissionsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacInitialPermissionsCreate`: InitialPermissions
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacInitialPermissionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRbacInitialPermissionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **initialPermissionsRequest** | [**InitialPermissionsRequest**](InitialPermissionsRequest.md) |  | 

### Return type

[**InitialPermissions**](InitialPermissions.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacInitialPermissionsDestroy

> RbacInitialPermissionsDestroy(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Initial Permissions.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacInitialPermissionsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacInitialPermissionsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Initial Permissions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacInitialPermissionsDestroyRequest struct via the builder pattern


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


## RbacInitialPermissionsList

> PaginatedInitialPermissionsList RbacInitialPermissionsList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacInitialPermissionsList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacInitialPermissionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacInitialPermissionsList`: PaginatedInitialPermissionsList
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacInitialPermissionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRbacInitialPermissionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedInitialPermissionsList**](PaginatedInitialPermissionsList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacInitialPermissionsPartialUpdate

> InitialPermissions RbacInitialPermissionsPartialUpdate(ctx, id).PatchedInitialPermissionsRequest(patchedInitialPermissionsRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Initial Permissions.
    patchedInitialPermissionsRequest := *openapiclient.NewPatchedInitialPermissionsRequest() // PatchedInitialPermissionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacInitialPermissionsPartialUpdate(context.Background(), id).PatchedInitialPermissionsRequest(patchedInitialPermissionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacInitialPermissionsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacInitialPermissionsPartialUpdate`: InitialPermissions
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacInitialPermissionsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Initial Permissions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacInitialPermissionsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedInitialPermissionsRequest** | [**PatchedInitialPermissionsRequest**](PatchedInitialPermissionsRequest.md) |  | 

### Return type

[**InitialPermissions**](InitialPermissions.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacInitialPermissionsRetrieve

> InitialPermissions RbacInitialPermissionsRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Initial Permissions.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacInitialPermissionsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacInitialPermissionsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacInitialPermissionsRetrieve`: InitialPermissions
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacInitialPermissionsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Initial Permissions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacInitialPermissionsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InitialPermissions**](InitialPermissions.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacInitialPermissionsUpdate

> InitialPermissions RbacInitialPermissionsUpdate(ctx, id).InitialPermissionsRequest(initialPermissionsRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Initial Permissions.
    initialPermissionsRequest := *openapiclient.NewInitialPermissionsRequest("Name_example", "Role_example") // InitialPermissionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacInitialPermissionsUpdate(context.Background(), id).InitialPermissionsRequest(initialPermissionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacInitialPermissionsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacInitialPermissionsUpdate`: InitialPermissions
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacInitialPermissionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Initial Permissions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacInitialPermissionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **initialPermissionsRequest** | [**InitialPermissionsRequest**](InitialPermissionsRequest.md) |  | 

### Return type

[**InitialPermissions**](InitialPermissions.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacInitialPermissionsUsedByList

> []UsedBy RbacInitialPermissionsUsedByList(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Initial Permissions.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacInitialPermissionsUsedByList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacInitialPermissionsUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacInitialPermissionsUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacInitialPermissionsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Initial Permissions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacInitialPermissionsUsedByListRequest struct via the builder pattern


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


## RbacPermissionsAssignedByRolesAssign

> []PermissionAssignResult RbacPermissionsAssignedByRolesAssign(ctx, uuid).PermissionAssignRequest(permissionAssignRequest).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.
    permissionAssignRequest := *openapiclient.NewPermissionAssignRequest([]string{"Permissions_example"}) // PermissionAssignRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacPermissionsAssignedByRolesAssign(context.Background(), uuid).PermissionAssignRequest(permissionAssignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacPermissionsAssignedByRolesAssign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacPermissionsAssignedByRolesAssign`: []PermissionAssignResult
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacPermissionsAssignedByRolesAssign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsAssignedByRolesAssignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permissionAssignRequest** | [**PermissionAssignRequest**](PermissionAssignRequest.md) |  | 

### Return type

[**[]PermissionAssignResult**](PermissionAssignResult.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsAssignedByRolesList

> PaginatedRoleAssignedObjectPermissionList RbacPermissionsAssignedByRolesList(ctx).Model(model).ObjectPk(objectPk).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    model := "model_example" // string | 
    objectPk := "objectPk_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacPermissionsAssignedByRolesList(context.Background()).Model(model).ObjectPk(objectPk).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacPermissionsAssignedByRolesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacPermissionsAssignedByRolesList`: PaginatedRoleAssignedObjectPermissionList
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacPermissionsAssignedByRolesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsAssignedByRolesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | **string** |  | 
 **objectPk** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedRoleAssignedObjectPermissionList**](PaginatedRoleAssignedObjectPermissionList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsAssignedByRolesUnassignPartialUpdate

> RbacPermissionsAssignedByRolesUnassignPartialUpdate(ctx, uuid).PatchedPermissionAssignRequest(patchedPermissionAssignRequest).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.
    patchedPermissionAssignRequest := *openapiclient.NewPatchedPermissionAssignRequest() // PatchedPermissionAssignRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacPermissionsAssignedByRolesUnassignPartialUpdate(context.Background(), uuid).PatchedPermissionAssignRequest(patchedPermissionAssignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacPermissionsAssignedByRolesUnassignPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsAssignedByRolesUnassignPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPermissionAssignRequest** | [**PatchedPermissionAssignRequest**](PatchedPermissionAssignRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsList

> PaginatedPermissionList RbacPermissionsList(ctx).Codename(codename).ContentTypeAppLabel(contentTypeAppLabel).ContentTypeModel(contentTypeModel).Ordering(ordering).Page(page).PageSize(pageSize).Role(role).Search(search).User(user).Execute()





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
    codename := "codename_example" // string |  (optional)
    contentTypeAppLabel := "contentTypeAppLabel_example" // string |  (optional)
    contentTypeModel := "contentTypeModel_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    role := "role_example" // string |  (optional)
    search := "search_example" // string | A search term. (optional)
    user := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacPermissionsList(context.Background()).Codename(codename).ContentTypeAppLabel(contentTypeAppLabel).ContentTypeModel(contentTypeModel).Ordering(ordering).Page(page).PageSize(pageSize).Role(role).Search(search).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacPermissionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacPermissionsList`: PaginatedPermissionList
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacPermissionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codename** | **string** |  | 
 **contentTypeAppLabel** | **string** |  | 
 **contentTypeModel** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **role** | **string** |  | 
 **search** | **string** | A search term. | 
 **user** | **int32** |  | 

### Return type

[**PaginatedPermissionList**](PaginatedPermissionList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsRetrieve

> Permission RbacPermissionsRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this permission.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacPermissionsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacPermissionsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacPermissionsRetrieve`: Permission
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacPermissionsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Permission**](Permission.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsRolesList

> PaginatedExtraRoleObjectPermissionList RbacPermissionsRolesList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Uuid(uuid).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacPermissionsRolesList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Uuid(uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacPermissionsRolesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacPermissionsRolesList`: PaginatedExtraRoleObjectPermissionList
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacPermissionsRolesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsRolesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **uuid** | **string** |  | 

### Return type

[**PaginatedExtraRoleObjectPermissionList**](PaginatedExtraRoleObjectPermissionList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacRolesAddUserCreate

> RbacRolesAddUserCreate(ctx, uuid).UserAccountSerializerForRoleRequest(userAccountSerializerForRoleRequest).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.
    userAccountSerializerForRoleRequest := *openapiclient.NewUserAccountSerializerForRoleRequest(int32(123)) // UserAccountSerializerForRoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacRolesAddUserCreate(context.Background(), uuid).UserAccountSerializerForRoleRequest(userAccountSerializerForRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacRolesAddUserCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacRolesAddUserCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userAccountSerializerForRoleRequest** | [**UserAccountSerializerForRoleRequest**](UserAccountSerializerForRoleRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacRolesCreate

> Role RbacRolesCreate(ctx).RoleRequest(roleRequest).Execute()





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
    roleRequest := *openapiclient.NewRoleRequest("Name_example") // RoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacRolesCreate(context.Background()).RoleRequest(roleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacRolesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacRolesCreate`: Role
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacRolesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRbacRolesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleRequest** | [**RoleRequest**](RoleRequest.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacRolesDestroy

> RbacRolesDestroy(ctx, uuid).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacRolesDestroy(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacRolesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacRolesDestroyRequest struct via the builder pattern


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


## RbacRolesList

> PaginatedRoleList RbacRolesList(ctx).Managed(managed).ManagedIsnull(managedIsnull).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Users(users).Execute()





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
    managed := []string{"Inner_example"} // []string |  (optional)
    managedIsnull := true // bool |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    users := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacRolesList(context.Background()).Managed(managed).ManagedIsnull(managedIsnull).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Users(users).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacRolesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacRolesList`: PaginatedRoleList
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacRolesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRbacRolesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managed** | **[]string** |  | 
 **managedIsnull** | **bool** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **users** | **[]int32** |  | 

### Return type

[**PaginatedRoleList**](PaginatedRoleList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacRolesPartialUpdate

> Role RbacRolesPartialUpdate(ctx, uuid).PatchedRoleRequest(patchedRoleRequest).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.
    patchedRoleRequest := *openapiclient.NewPatchedRoleRequest() // PatchedRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacRolesPartialUpdate(context.Background(), uuid).PatchedRoleRequest(patchedRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacRolesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacRolesPartialUpdate`: Role
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacRolesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacRolesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedRoleRequest** | [**PatchedRoleRequest**](PatchedRoleRequest.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacRolesRemoveUserCreate

> RbacRolesRemoveUserCreate(ctx, uuid).UserAccountSerializerForRoleRequest(userAccountSerializerForRoleRequest).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.
    userAccountSerializerForRoleRequest := *openapiclient.NewUserAccountSerializerForRoleRequest(int32(123)) // UserAccountSerializerForRoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacRolesRemoveUserCreate(context.Background(), uuid).UserAccountSerializerForRoleRequest(userAccountSerializerForRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacRolesRemoveUserCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacRolesRemoveUserCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userAccountSerializerForRoleRequest** | [**UserAccountSerializerForRoleRequest**](UserAccountSerializerForRoleRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacRolesRetrieve

> Role RbacRolesRetrieve(ctx, uuid).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacRolesRetrieve(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacRolesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacRolesRetrieve`: Role
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacRolesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacRolesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Role**](Role.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacRolesUpdate

> Role RbacRolesUpdate(ctx, uuid).RoleRequest(roleRequest).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.
    roleRequest := *openapiclient.NewRoleRequest("Name_example") // RoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacRolesUpdate(context.Background(), uuid).RoleRequest(roleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacRolesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacRolesUpdate`: Role
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacRolesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacRolesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleRequest** | [**RoleRequest**](RoleRequest.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacRolesUsedByList

> []UsedBy RbacRolesUsedByList(ctx, uuid).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacRolesUsedByList(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacRolesUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacRolesUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacRolesUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacRolesUsedByListRequest struct via the builder pattern


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

