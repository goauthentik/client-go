# \RbacAPI

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RbacInitialPermissionsCreate**](RbacAPI.md#RbacInitialPermissionsCreate) | **Post** /rbac/initial_permissions/ | 
[**RbacInitialPermissionsDestroy**](RbacAPI.md#RbacInitialPermissionsDestroy) | **Delete** /rbac/initial_permissions/{id}/ | 
[**RbacInitialPermissionsList**](RbacAPI.md#RbacInitialPermissionsList) | **Get** /rbac/initial_permissions/ | 
[**RbacInitialPermissionsPartialUpdate**](RbacAPI.md#RbacInitialPermissionsPartialUpdate) | **Patch** /rbac/initial_permissions/{id}/ | 
[**RbacInitialPermissionsRetrieve**](RbacAPI.md#RbacInitialPermissionsRetrieve) | **Get** /rbac/initial_permissions/{id}/ | 
[**RbacInitialPermissionsUpdate**](RbacAPI.md#RbacInitialPermissionsUpdate) | **Put** /rbac/initial_permissions/{id}/ | 
[**RbacInitialPermissionsUsedByList**](RbacAPI.md#RbacInitialPermissionsUsedByList) | **Get** /rbac/initial_permissions/{id}/used_by/ | 
[**RbacPermissionsAssignedByRolesAssign**](RbacAPI.md#RbacPermissionsAssignedByRolesAssign) | **Post** /rbac/permissions/assigned_by_roles/{uuid}/assign/ | 
[**RbacPermissionsAssignedByRolesList**](RbacAPI.md#RbacPermissionsAssignedByRolesList) | **Get** /rbac/permissions/assigned_by_roles/ | 
[**RbacPermissionsAssignedByRolesUnassignPartialUpdate**](RbacAPI.md#RbacPermissionsAssignedByRolesUnassignPartialUpdate) | **Patch** /rbac/permissions/assigned_by_roles/{uuid}/unassign/ | 
[**RbacPermissionsAssignedByUsersAssign**](RbacAPI.md#RbacPermissionsAssignedByUsersAssign) | **Post** /rbac/permissions/assigned_by_users/{id}/assign/ | 
[**RbacPermissionsAssignedByUsersList**](RbacAPI.md#RbacPermissionsAssignedByUsersList) | **Get** /rbac/permissions/assigned_by_users/ | 
[**RbacPermissionsAssignedByUsersUnassignPartialUpdate**](RbacAPI.md#RbacPermissionsAssignedByUsersUnassignPartialUpdate) | **Patch** /rbac/permissions/assigned_by_users/{id}/unassign/ | 
[**RbacPermissionsList**](RbacAPI.md#RbacPermissionsList) | **Get** /rbac/permissions/ | 
[**RbacPermissionsRetrieve**](RbacAPI.md#RbacPermissionsRetrieve) | **Get** /rbac/permissions/{id}/ | 
[**RbacPermissionsRolesDestroy**](RbacAPI.md#RbacPermissionsRolesDestroy) | **Delete** /rbac/permissions/roles/{id}/ | 
[**RbacPermissionsRolesList**](RbacAPI.md#RbacPermissionsRolesList) | **Get** /rbac/permissions/roles/ | 
[**RbacPermissionsRolesPartialUpdate**](RbacAPI.md#RbacPermissionsRolesPartialUpdate) | **Patch** /rbac/permissions/roles/{id}/ | 
[**RbacPermissionsRolesRetrieve**](RbacAPI.md#RbacPermissionsRolesRetrieve) | **Get** /rbac/permissions/roles/{id}/ | 
[**RbacPermissionsRolesUpdate**](RbacAPI.md#RbacPermissionsRolesUpdate) | **Put** /rbac/permissions/roles/{id}/ | 
[**RbacPermissionsUsersDestroy**](RbacAPI.md#RbacPermissionsUsersDestroy) | **Delete** /rbac/permissions/users/{id}/ | 
[**RbacPermissionsUsersList**](RbacAPI.md#RbacPermissionsUsersList) | **Get** /rbac/permissions/users/ | 
[**RbacPermissionsUsersPartialUpdate**](RbacAPI.md#RbacPermissionsUsersPartialUpdate) | **Patch** /rbac/permissions/users/{id}/ | 
[**RbacPermissionsUsersRetrieve**](RbacAPI.md#RbacPermissionsUsersRetrieve) | **Get** /rbac/permissions/users/{id}/ | 
[**RbacPermissionsUsersUpdate**](RbacAPI.md#RbacPermissionsUsersUpdate) | **Put** /rbac/permissions/users/{id}/ | 
[**RbacRolesCreate**](RbacAPI.md#RbacRolesCreate) | **Post** /rbac/roles/ | 
[**RbacRolesDestroy**](RbacAPI.md#RbacRolesDestroy) | **Delete** /rbac/roles/{uuid}/ | 
[**RbacRolesList**](RbacAPI.md#RbacRolesList) | **Get** /rbac/roles/ | 
[**RbacRolesPartialUpdate**](RbacAPI.md#RbacRolesPartialUpdate) | **Patch** /rbac/roles/{uuid}/ | 
[**RbacRolesRetrieve**](RbacAPI.md#RbacRolesRetrieve) | **Get** /rbac/roles/{uuid}/ | 
[**RbacRolesUpdate**](RbacAPI.md#RbacRolesUpdate) | **Put** /rbac/roles/{uuid}/ | 
[**RbacRolesUsedByList**](RbacAPI.md#RbacRolesUsedByList) | **Get** /rbac/roles/{uuid}/used_by/ | 



## RbacInitialPermissionsCreate

> InitialPermissions RbacInitialPermissionsCreate(ctx).InitialPermissionsRequest(initialPermissionsRequest).Execute()





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
	initialPermissionsRequest := *openapiclient.NewInitialPermissionsRequest("Name_example", openapiclient.InitialPermissionsModeEnum("user"), "Role_example") // InitialPermissionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacInitialPermissionsCreate(context.Background()).InitialPermissionsRequest(initialPermissionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacInitialPermissionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacInitialPermissionsCreate`: InitialPermissions
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacInitialPermissionsCreate`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Initial Permissions.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbacAPI.RbacInitialPermissionsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacInitialPermissionsDestroy``: %v\n", err)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacInitialPermissionsList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacInitialPermissionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacInitialPermissionsList`: PaginatedInitialPermissionsList
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacInitialPermissionsList`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Initial Permissions.
	patchedInitialPermissionsRequest := *openapiclient.NewPatchedInitialPermissionsRequest() // PatchedInitialPermissionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacInitialPermissionsPartialUpdate(context.Background(), id).PatchedInitialPermissionsRequest(patchedInitialPermissionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacInitialPermissionsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacInitialPermissionsPartialUpdate`: InitialPermissions
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacInitialPermissionsPartialUpdate`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Initial Permissions.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacInitialPermissionsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacInitialPermissionsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacInitialPermissionsRetrieve`: InitialPermissions
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacInitialPermissionsRetrieve`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Initial Permissions.
	initialPermissionsRequest := *openapiclient.NewInitialPermissionsRequest("Name_example", openapiclient.InitialPermissionsModeEnum("user"), "Role_example") // InitialPermissionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacInitialPermissionsUpdate(context.Background(), id).InitialPermissionsRequest(initialPermissionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacInitialPermissionsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacInitialPermissionsUpdate`: InitialPermissions
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacInitialPermissionsUpdate`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Initial Permissions.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacInitialPermissionsUsedByList(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacInitialPermissionsUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacInitialPermissionsUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacInitialPermissionsUsedByList`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.
	permissionAssignRequest := *openapiclient.NewPermissionAssignRequest([]string{"Permissions_example"}) // PermissionAssignRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacPermissionsAssignedByRolesAssign(context.Background(), uuid).PermissionAssignRequest(permissionAssignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsAssignedByRolesAssign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacPermissionsAssignedByRolesAssign`: []PermissionAssignResult
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacPermissionsAssignedByRolesAssign`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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
	resp, r, err := apiClient.RbacAPI.RbacPermissionsAssignedByRolesList(context.Background()).Model(model).ObjectPk(objectPk).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsAssignedByRolesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacPermissionsAssignedByRolesList`: PaginatedRoleAssignedObjectPermissionList
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacPermissionsAssignedByRolesList`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.
	patchedPermissionAssignRequest := *openapiclient.NewPatchedPermissionAssignRequest() // PatchedPermissionAssignRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbacAPI.RbacPermissionsAssignedByRolesUnassignPartialUpdate(context.Background(), uuid).PatchedPermissionAssignRequest(patchedPermissionAssignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsAssignedByRolesUnassignPartialUpdate``: %v\n", err)
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


## RbacPermissionsAssignedByUsersAssign

> []PermissionAssignResult RbacPermissionsAssignedByUsersAssign(ctx, id).PermissionAssignRequest(permissionAssignRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this User.
	permissionAssignRequest := *openapiclient.NewPermissionAssignRequest([]string{"Permissions_example"}) // PermissionAssignRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacPermissionsAssignedByUsersAssign(context.Background(), id).PermissionAssignRequest(permissionAssignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsAssignedByUsersAssign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacPermissionsAssignedByUsersAssign`: []PermissionAssignResult
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacPermissionsAssignedByUsersAssign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsAssignedByUsersAssignRequest struct via the builder pattern


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


## RbacPermissionsAssignedByUsersList

> PaginatedUserAssignedObjectPermissionList RbacPermissionsAssignedByUsersList(ctx).Model(model).ObjectPk(objectPk).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
	model := "model_example" // string | 
	objectPk := "objectPk_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacPermissionsAssignedByUsersList(context.Background()).Model(model).ObjectPk(objectPk).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsAssignedByUsersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacPermissionsAssignedByUsersList`: PaginatedUserAssignedObjectPermissionList
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacPermissionsAssignedByUsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsAssignedByUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | **string** |  | 
 **objectPk** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedUserAssignedObjectPermissionList**](PaginatedUserAssignedObjectPermissionList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsAssignedByUsersUnassignPartialUpdate

> RbacPermissionsAssignedByUsersUnassignPartialUpdate(ctx, id).PatchedPermissionAssignRequest(patchedPermissionAssignRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this User.
	patchedPermissionAssignRequest := *openapiclient.NewPatchedPermissionAssignRequest() // PatchedPermissionAssignRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbacAPI.RbacPermissionsAssignedByUsersUnassignPartialUpdate(context.Background(), id).PatchedPermissionAssignRequest(patchedPermissionAssignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsAssignedByUsersUnassignPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsAssignedByUsersUnassignPartialUpdateRequest struct via the builder pattern


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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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
	resp, r, err := apiClient.RbacAPI.RbacPermissionsList(context.Background()).Codename(codename).ContentTypeAppLabel(contentTypeAppLabel).ContentTypeModel(contentTypeModel).Ordering(ordering).Page(page).PageSize(pageSize).Role(role).Search(search).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacPermissionsList`: PaginatedPermissionList
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacPermissionsList`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this permission.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacPermissionsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacPermissionsRetrieve`: Permission
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacPermissionsRetrieve`: %v\n", resp)
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


## RbacPermissionsRolesDestroy

> RbacPermissionsRolesDestroy(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this group object permission.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbacAPI.RbacPermissionsRolesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsRolesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group object permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsRolesDestroyRequest struct via the builder pattern


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


## RbacPermissionsRolesList

> PaginatedExtraRoleObjectPermissionList RbacPermissionsRolesList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Uuid(uuid).Execute()





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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacPermissionsRolesList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Uuid(uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsRolesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacPermissionsRolesList`: PaginatedExtraRoleObjectPermissionList
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacPermissionsRolesList`: %v\n", resp)
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


## RbacPermissionsRolesPartialUpdate

> ExtraRoleObjectPermission RbacPermissionsRolesPartialUpdate(ctx, id).PatchedExtraRoleObjectPermissionRequest(patchedExtraRoleObjectPermissionRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this group object permission.
	patchedExtraRoleObjectPermissionRequest := *openapiclient.NewPatchedExtraRoleObjectPermissionRequest() // PatchedExtraRoleObjectPermissionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacPermissionsRolesPartialUpdate(context.Background(), id).PatchedExtraRoleObjectPermissionRequest(patchedExtraRoleObjectPermissionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsRolesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacPermissionsRolesPartialUpdate`: ExtraRoleObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacPermissionsRolesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group object permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsRolesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedExtraRoleObjectPermissionRequest** | [**PatchedExtraRoleObjectPermissionRequest**](PatchedExtraRoleObjectPermissionRequest.md) |  | 

### Return type

[**ExtraRoleObjectPermission**](ExtraRoleObjectPermission.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsRolesRetrieve

> ExtraRoleObjectPermission RbacPermissionsRolesRetrieve(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this group object permission.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacPermissionsRolesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsRolesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacPermissionsRolesRetrieve`: ExtraRoleObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacPermissionsRolesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group object permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsRolesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExtraRoleObjectPermission**](ExtraRoleObjectPermission.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsRolesUpdate

> ExtraRoleObjectPermission RbacPermissionsRolesUpdate(ctx, id).ExtraRoleObjectPermissionRequest(extraRoleObjectPermissionRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this group object permission.
	extraRoleObjectPermissionRequest := *openapiclient.NewExtraRoleObjectPermissionRequest("ObjectPk_example") // ExtraRoleObjectPermissionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacPermissionsRolesUpdate(context.Background(), id).ExtraRoleObjectPermissionRequest(extraRoleObjectPermissionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsRolesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacPermissionsRolesUpdate`: ExtraRoleObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacPermissionsRolesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group object permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsRolesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extraRoleObjectPermissionRequest** | [**ExtraRoleObjectPermissionRequest**](ExtraRoleObjectPermissionRequest.md) |  | 

### Return type

[**ExtraRoleObjectPermission**](ExtraRoleObjectPermission.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsUsersDestroy

> RbacPermissionsUsersDestroy(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this user object permission.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbacAPI.RbacPermissionsUsersDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsUsersDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user object permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsUsersDestroyRequest struct via the builder pattern


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


## RbacPermissionsUsersList

> PaginatedExtraUserObjectPermissionList RbacPermissionsUsersList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).UserId(userId).Execute()





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
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)
	userId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacPermissionsUsersList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsUsersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacPermissionsUsersList`: PaginatedExtraUserObjectPermissionList
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacPermissionsUsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **userId** | **int32** |  | 

### Return type

[**PaginatedExtraUserObjectPermissionList**](PaginatedExtraUserObjectPermissionList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsUsersPartialUpdate

> ExtraUserObjectPermission RbacPermissionsUsersPartialUpdate(ctx, id).PatchedExtraUserObjectPermissionRequest(patchedExtraUserObjectPermissionRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this user object permission.
	patchedExtraUserObjectPermissionRequest := *openapiclient.NewPatchedExtraUserObjectPermissionRequest() // PatchedExtraUserObjectPermissionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacPermissionsUsersPartialUpdate(context.Background(), id).PatchedExtraUserObjectPermissionRequest(patchedExtraUserObjectPermissionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsUsersPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacPermissionsUsersPartialUpdate`: ExtraUserObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacPermissionsUsersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user object permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsUsersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedExtraUserObjectPermissionRequest** | [**PatchedExtraUserObjectPermissionRequest**](PatchedExtraUserObjectPermissionRequest.md) |  | 

### Return type

[**ExtraUserObjectPermission**](ExtraUserObjectPermission.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsUsersRetrieve

> ExtraUserObjectPermission RbacPermissionsUsersRetrieve(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this user object permission.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacPermissionsUsersRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsUsersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacPermissionsUsersRetrieve`: ExtraUserObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacPermissionsUsersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user object permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsUsersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExtraUserObjectPermission**](ExtraUserObjectPermission.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsUsersUpdate

> ExtraUserObjectPermission RbacPermissionsUsersUpdate(ctx, id).ExtraUserObjectPermissionRequest(extraUserObjectPermissionRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this user object permission.
	extraUserObjectPermissionRequest := *openapiclient.NewExtraUserObjectPermissionRequest("ObjectPk_example") // ExtraUserObjectPermissionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacPermissionsUsersUpdate(context.Background(), id).ExtraUserObjectPermissionRequest(extraUserObjectPermissionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacPermissionsUsersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacPermissionsUsersUpdate`: ExtraUserObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacPermissionsUsersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user object permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsUsersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extraUserObjectPermissionRequest** | [**ExtraUserObjectPermissionRequest**](ExtraUserObjectPermissionRequest.md) |  | 

### Return type

[**ExtraUserObjectPermission**](ExtraUserObjectPermission.md)

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	roleRequest := *openapiclient.NewRoleRequest("Name_example") // RoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacRolesCreate(context.Background()).RoleRequest(roleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacRolesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacRolesCreate`: Role
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacRolesCreate`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbacAPI.RbacRolesDestroy(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacRolesDestroy``: %v\n", err)
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

> PaginatedRoleList RbacRolesList(ctx).GroupName(groupName).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
	groupName := "groupName_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacRolesList(context.Background()).GroupName(groupName).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacRolesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacRolesList`: PaginatedRoleList
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacRolesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRbacRolesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.
	patchedRoleRequest := *openapiclient.NewPatchedRoleRequest() // PatchedRoleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacRolesPartialUpdate(context.Background(), uuid).PatchedRoleRequest(patchedRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacRolesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacRolesPartialUpdate`: Role
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacRolesPartialUpdate`: %v\n", resp)
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


## RbacRolesRetrieve

> Role RbacRolesRetrieve(ctx, uuid).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacRolesRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacRolesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacRolesRetrieve`: Role
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacRolesRetrieve`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.
	roleRequest := *openapiclient.NewRoleRequest("Name_example") // RoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacRolesUpdate(context.Background(), uuid).RoleRequest(roleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacRolesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacRolesUpdate`: Role
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacRolesUpdate`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbacAPI.RbacRolesUsedByList(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbacAPI.RbacRolesUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RbacRolesUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `RbacAPI.RbacRolesUsedByList`: %v\n", resp)
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

