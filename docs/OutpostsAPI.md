# \OutpostsAPI

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OutpostsInstancesCreate**](OutpostsAPI.md#OutpostsInstancesCreate) | **Post** /outposts/instances/ | 
[**OutpostsInstancesDefaultSettingsRetrieve**](OutpostsAPI.md#OutpostsInstancesDefaultSettingsRetrieve) | **Get** /outposts/instances/default_settings/ | 
[**OutpostsInstancesDestroy**](OutpostsAPI.md#OutpostsInstancesDestroy) | **Delete** /outposts/instances/{uuid}/ | 
[**OutpostsInstancesHealthList**](OutpostsAPI.md#OutpostsInstancesHealthList) | **Get** /outposts/instances/{uuid}/health/ | 
[**OutpostsInstancesList**](OutpostsAPI.md#OutpostsInstancesList) | **Get** /outposts/instances/ | 
[**OutpostsInstancesPartialUpdate**](OutpostsAPI.md#OutpostsInstancesPartialUpdate) | **Patch** /outposts/instances/{uuid}/ | 
[**OutpostsInstancesRetrieve**](OutpostsAPI.md#OutpostsInstancesRetrieve) | **Get** /outposts/instances/{uuid}/ | 
[**OutpostsInstancesUpdate**](OutpostsAPI.md#OutpostsInstancesUpdate) | **Put** /outposts/instances/{uuid}/ | 
[**OutpostsInstancesUsedByList**](OutpostsAPI.md#OutpostsInstancesUsedByList) | **Get** /outposts/instances/{uuid}/used_by/ | 
[**OutpostsLdapAccessCheck**](OutpostsAPI.md#OutpostsLdapAccessCheck) | **Get** /outposts/ldap/{id}/check_access/ | 
[**OutpostsLdapList**](OutpostsAPI.md#OutpostsLdapList) | **Get** /outposts/ldap/ | 
[**OutpostsProxyList**](OutpostsAPI.md#OutpostsProxyList) | **Get** /outposts/proxy/ | 
[**OutpostsRadiusAccessCheck**](OutpostsAPI.md#OutpostsRadiusAccessCheck) | **Get** /outposts/radius/{id}/check_access/ | 
[**OutpostsRadiusList**](OutpostsAPI.md#OutpostsRadiusList) | **Get** /outposts/radius/ | 
[**OutpostsServiceConnectionsAllDestroy**](OutpostsAPI.md#OutpostsServiceConnectionsAllDestroy) | **Delete** /outposts/service_connections/all/{uuid}/ | 
[**OutpostsServiceConnectionsAllList**](OutpostsAPI.md#OutpostsServiceConnectionsAllList) | **Get** /outposts/service_connections/all/ | 
[**OutpostsServiceConnectionsAllRetrieve**](OutpostsAPI.md#OutpostsServiceConnectionsAllRetrieve) | **Get** /outposts/service_connections/all/{uuid}/ | 
[**OutpostsServiceConnectionsAllStateRetrieve**](OutpostsAPI.md#OutpostsServiceConnectionsAllStateRetrieve) | **Get** /outposts/service_connections/all/{uuid}/state/ | 
[**OutpostsServiceConnectionsAllTypesList**](OutpostsAPI.md#OutpostsServiceConnectionsAllTypesList) | **Get** /outposts/service_connections/all/types/ | 
[**OutpostsServiceConnectionsAllUsedByList**](OutpostsAPI.md#OutpostsServiceConnectionsAllUsedByList) | **Get** /outposts/service_connections/all/{uuid}/used_by/ | 
[**OutpostsServiceConnectionsDockerCreate**](OutpostsAPI.md#OutpostsServiceConnectionsDockerCreate) | **Post** /outposts/service_connections/docker/ | 
[**OutpostsServiceConnectionsDockerDestroy**](OutpostsAPI.md#OutpostsServiceConnectionsDockerDestroy) | **Delete** /outposts/service_connections/docker/{uuid}/ | 
[**OutpostsServiceConnectionsDockerList**](OutpostsAPI.md#OutpostsServiceConnectionsDockerList) | **Get** /outposts/service_connections/docker/ | 
[**OutpostsServiceConnectionsDockerPartialUpdate**](OutpostsAPI.md#OutpostsServiceConnectionsDockerPartialUpdate) | **Patch** /outposts/service_connections/docker/{uuid}/ | 
[**OutpostsServiceConnectionsDockerRetrieve**](OutpostsAPI.md#OutpostsServiceConnectionsDockerRetrieve) | **Get** /outposts/service_connections/docker/{uuid}/ | 
[**OutpostsServiceConnectionsDockerUpdate**](OutpostsAPI.md#OutpostsServiceConnectionsDockerUpdate) | **Put** /outposts/service_connections/docker/{uuid}/ | 
[**OutpostsServiceConnectionsDockerUsedByList**](OutpostsAPI.md#OutpostsServiceConnectionsDockerUsedByList) | **Get** /outposts/service_connections/docker/{uuid}/used_by/ | 
[**OutpostsServiceConnectionsKubernetesCreate**](OutpostsAPI.md#OutpostsServiceConnectionsKubernetesCreate) | **Post** /outposts/service_connections/kubernetes/ | 
[**OutpostsServiceConnectionsKubernetesDestroy**](OutpostsAPI.md#OutpostsServiceConnectionsKubernetesDestroy) | **Delete** /outposts/service_connections/kubernetes/{uuid}/ | 
[**OutpostsServiceConnectionsKubernetesList**](OutpostsAPI.md#OutpostsServiceConnectionsKubernetesList) | **Get** /outposts/service_connections/kubernetes/ | 
[**OutpostsServiceConnectionsKubernetesPartialUpdate**](OutpostsAPI.md#OutpostsServiceConnectionsKubernetesPartialUpdate) | **Patch** /outposts/service_connections/kubernetes/{uuid}/ | 
[**OutpostsServiceConnectionsKubernetesRetrieve**](OutpostsAPI.md#OutpostsServiceConnectionsKubernetesRetrieve) | **Get** /outposts/service_connections/kubernetes/{uuid}/ | 
[**OutpostsServiceConnectionsKubernetesUpdate**](OutpostsAPI.md#OutpostsServiceConnectionsKubernetesUpdate) | **Put** /outposts/service_connections/kubernetes/{uuid}/ | 
[**OutpostsServiceConnectionsKubernetesUsedByList**](OutpostsAPI.md#OutpostsServiceConnectionsKubernetesUsedByList) | **Get** /outposts/service_connections/kubernetes/{uuid}/used_by/ | 



## OutpostsInstancesCreate

> Outpost OutpostsInstancesCreate(ctx).OutpostRequest(outpostRequest).Execute()





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
	outpostRequest := *openapiclient.NewOutpostRequest("Name_example", openapiclient.OutpostTypeEnum("proxy"), []int32{int32(123)}, map[string]interface{}{"key": interface{}(123)}) // OutpostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsInstancesCreate(context.Background()).OutpostRequest(outpostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsInstancesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsInstancesCreate`: Outpost
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsInstancesCreate`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsInstancesDefaultSettingsRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsInstancesDefaultSettingsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsInstancesDefaultSettingsRetrieve`: OutpostDefaultConfig
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsInstancesDefaultSettingsRetrieve`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Outpost.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OutpostsAPI.OutpostsInstancesDestroy(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsInstancesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Outpost. | 

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsInstancesHealthList

> []OutpostHealth OutpostsInstancesHealthList(ctx, uuid).ManagedIcontains(managedIcontains).ManagedIexact(managedIexact).NameIcontains(nameIcontains).NameIexact(nameIexact).Ordering(ordering).ProvidersIsnull(providersIsnull).ProvidersByPk(providersByPk).Search(search).ServiceConnectionNameIcontains(serviceConnectionNameIcontains).ServiceConnectionNameIexact(serviceConnectionNameIexact).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Outpost.
	managedIcontains := "managedIcontains_example" // string |  (optional)
	managedIexact := "managedIexact_example" // string |  (optional)
	nameIcontains := "nameIcontains_example" // string |  (optional)
	nameIexact := "nameIexact_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	providersIsnull := true // bool |  (optional)
	providersByPk := []int32{int32(123)} // []int32 |  (optional)
	search := "search_example" // string | A search term. (optional)
	serviceConnectionNameIcontains := "serviceConnectionNameIcontains_example" // string |  (optional)
	serviceConnectionNameIexact := "serviceConnectionNameIexact_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsInstancesHealthList(context.Background(), uuid).ManagedIcontains(managedIcontains).ManagedIexact(managedIexact).NameIcontains(nameIcontains).NameIexact(nameIexact).Ordering(ordering).ProvidersIsnull(providersIsnull).ProvidersByPk(providersByPk).Search(search).ServiceConnectionNameIcontains(serviceConnectionNameIcontains).ServiceConnectionNameIexact(serviceConnectionNameIexact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsInstancesHealthList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsInstancesHealthList`: []OutpostHealth
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsInstancesHealthList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Outpost. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsInstancesHealthListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managedIcontains** | **string** |  | 
 **managedIexact** | **string** |  | 
 **nameIcontains** | **string** |  | 
 **nameIexact** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **providersIsnull** | **bool** |  | 
 **providersByPk** | **[]int32** |  | 
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

> PaginatedOutpostList OutpostsInstancesList(ctx).ManagedIcontains(managedIcontains).ManagedIexact(managedIexact).NameIcontains(nameIcontains).NameIexact(nameIexact).Ordering(ordering).Page(page).PageSize(pageSize).ProvidersIsnull(providersIsnull).ProvidersByPk(providersByPk).Search(search).ServiceConnectionNameIcontains(serviceConnectionNameIcontains).ServiceConnectionNameIexact(serviceConnectionNameIexact).Execute()





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
	managedIcontains := "managedIcontains_example" // string |  (optional)
	managedIexact := "managedIexact_example" // string |  (optional)
	nameIcontains := "nameIcontains_example" // string |  (optional)
	nameIexact := "nameIexact_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	providersIsnull := true // bool |  (optional)
	providersByPk := []int32{int32(123)} // []int32 |  (optional)
	search := "search_example" // string | A search term. (optional)
	serviceConnectionNameIcontains := "serviceConnectionNameIcontains_example" // string |  (optional)
	serviceConnectionNameIexact := "serviceConnectionNameIexact_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsInstancesList(context.Background()).ManagedIcontains(managedIcontains).ManagedIexact(managedIexact).NameIcontains(nameIcontains).NameIexact(nameIexact).Ordering(ordering).Page(page).PageSize(pageSize).ProvidersIsnull(providersIsnull).ProvidersByPk(providersByPk).Search(search).ServiceConnectionNameIcontains(serviceConnectionNameIcontains).ServiceConnectionNameIexact(serviceConnectionNameIexact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsInstancesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsInstancesList`: PaginatedOutpostList
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsInstancesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsInstancesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managedIcontains** | **string** |  | 
 **managedIexact** | **string** |  | 
 **nameIcontains** | **string** |  | 
 **nameIexact** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **providersIsnull** | **bool** |  | 
 **providersByPk** | **[]int32** |  | 
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Outpost.
	patchedOutpostRequest := *openapiclient.NewPatchedOutpostRequest() // PatchedOutpostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsInstancesPartialUpdate(context.Background(), uuid).PatchedOutpostRequest(patchedOutpostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsInstancesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsInstancesPartialUpdate`: Outpost
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsInstancesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Outpost. | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Outpost.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsInstancesRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsInstancesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsInstancesRetrieve`: Outpost
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsInstancesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Outpost. | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Outpost.
	outpostRequest := *openapiclient.NewOutpostRequest("Name_example", openapiclient.OutpostTypeEnum("proxy"), []int32{int32(123)}, map[string]interface{}{"key": interface{}(123)}) // OutpostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsInstancesUpdate(context.Background(), uuid).OutpostRequest(outpostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsInstancesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsInstancesUpdate`: Outpost
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsInstancesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Outpost. | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Outpost.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsInstancesUsedByList(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsInstancesUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsInstancesUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsInstancesUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Outpost. | 

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


## OutpostsLdapAccessCheck

> LDAPCheckAccess OutpostsLdapAccessCheck(ctx, id).AppSlug(appSlug).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this LDAP Provider.
	appSlug := "appSlug_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsLdapAccessCheck(context.Background(), id).AppSlug(appSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsLdapAccessCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsLdapAccessCheck`: LDAPCheckAccess
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsLdapAccessCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this LDAP Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsLdapAccessCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appSlug** | **string** |  | 

### Return type

[**LDAPCheckAccess**](LDAPCheckAccess.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsLdapList

> PaginatedLDAPOutpostConfigList OutpostsLdapList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
	resp, r, err := apiClient.OutpostsAPI.OutpostsLdapList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsLdapList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsLdapList`: PaginatedLDAPOutpostConfigList
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsLdapList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsLdapListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
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


## OutpostsProxyList

> PaginatedProxyOutpostConfigList OutpostsProxyList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
	resp, r, err := apiClient.OutpostsAPI.OutpostsProxyList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsProxyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsProxyList`: PaginatedProxyOutpostConfigList
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsProxyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsProxyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
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


## OutpostsRadiusAccessCheck

> RadiusCheckAccess OutpostsRadiusAccessCheck(ctx, id).AppSlug(appSlug).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this Radius Provider.
	appSlug := "appSlug_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsRadiusAccessCheck(context.Background(), id).AppSlug(appSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsRadiusAccessCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsRadiusAccessCheck`: RadiusCheckAccess
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsRadiusAccessCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Radius Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsRadiusAccessCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appSlug** | **string** |  | 

### Return type

[**RadiusCheckAccess**](RadiusCheckAccess.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutpostsRadiusList

> PaginatedRadiusOutpostConfigList OutpostsRadiusList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
	resp, r, err := apiClient.OutpostsAPI.OutpostsRadiusList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsRadiusList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsRadiusList`: PaginatedRadiusOutpostConfigList
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsRadiusList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOutpostsRadiusListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedRadiusOutpostConfigList**](PaginatedRadiusOutpostConfigList.md)

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Outpost Service-Connection.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsAllDestroy(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsAllDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Outpost Service-Connection. | 

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
- **Accept**: application/json

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
	resp, r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsAllList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsAllList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsServiceConnectionsAllList`: PaginatedServiceConnectionList
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsServiceConnectionsAllList`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Outpost Service-Connection.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsAllRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsAllRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsServiceConnectionsAllRetrieve`: ServiceConnection
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsServiceConnectionsAllRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Outpost Service-Connection. | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Outpost Service-Connection.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsAllStateRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsAllStateRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsServiceConnectionsAllStateRetrieve`: ServiceConnectionState
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsServiceConnectionsAllStateRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Outpost Service-Connection. | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsAllTypesList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsAllTypesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsServiceConnectionsAllTypesList`: []TypeCreate
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsServiceConnectionsAllTypesList`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Outpost Service-Connection.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsAllUsedByList(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsAllUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsServiceConnectionsAllUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsServiceConnectionsAllUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Outpost Service-Connection. | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dockerServiceConnectionRequest := *openapiclient.NewDockerServiceConnectionRequest("Name_example", "Url_example") // DockerServiceConnectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsDockerCreate(context.Background()).DockerServiceConnectionRequest(dockerServiceConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsDockerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsServiceConnectionsDockerCreate`: DockerServiceConnection
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsServiceConnectionsDockerCreate`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Docker Service-Connection.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsDockerDestroy(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsDockerDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Docker Service-Connection. | 

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
- **Accept**: application/json

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	local := true // bool |  (optional)
	name := "name_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)
	tlsAuthentication := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	tlsVerification := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	url := "url_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsDockerList(context.Background()).Local(local).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).TlsAuthentication(tlsAuthentication).TlsVerification(tlsVerification).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsDockerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsServiceConnectionsDockerList`: PaginatedDockerServiceConnectionList
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsServiceConnectionsDockerList`: %v\n", resp)
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
 **tlsAuthentication** | **string** |  | 
 **tlsVerification** | **string** |  | 
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Docker Service-Connection.
	patchedDockerServiceConnectionRequest := *openapiclient.NewPatchedDockerServiceConnectionRequest() // PatchedDockerServiceConnectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsDockerPartialUpdate(context.Background(), uuid).PatchedDockerServiceConnectionRequest(patchedDockerServiceConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsDockerPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsServiceConnectionsDockerPartialUpdate`: DockerServiceConnection
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsServiceConnectionsDockerPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Docker Service-Connection. | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Docker Service-Connection.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsDockerRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsDockerRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsServiceConnectionsDockerRetrieve`: DockerServiceConnection
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsServiceConnectionsDockerRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Docker Service-Connection. | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Docker Service-Connection.
	dockerServiceConnectionRequest := *openapiclient.NewDockerServiceConnectionRequest("Name_example", "Url_example") // DockerServiceConnectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsDockerUpdate(context.Background(), uuid).DockerServiceConnectionRequest(dockerServiceConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsDockerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsServiceConnectionsDockerUpdate`: DockerServiceConnection
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsServiceConnectionsDockerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Docker Service-Connection. | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Docker Service-Connection.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsDockerUsedByList(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsDockerUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsServiceConnectionsDockerUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsServiceConnectionsDockerUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Docker Service-Connection. | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	kubernetesServiceConnectionRequest := *openapiclient.NewKubernetesServiceConnectionRequest("Name_example") // KubernetesServiceConnectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsKubernetesCreate(context.Background()).KubernetesServiceConnectionRequest(kubernetesServiceConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsKubernetesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsServiceConnectionsKubernetesCreate`: KubernetesServiceConnection
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsServiceConnectionsKubernetesCreate`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Kubernetes Service-Connection.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsKubernetesDestroy(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsKubernetesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Kubernetes Service-Connection. | 

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
- **Accept**: application/json

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	local := true // bool |  (optional)
	name := "name_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsKubernetesList(context.Background()).Local(local).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsKubernetesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsServiceConnectionsKubernetesList`: PaginatedKubernetesServiceConnectionList
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsServiceConnectionsKubernetesList`: %v\n", resp)
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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Kubernetes Service-Connection.
	patchedKubernetesServiceConnectionRequest := *openapiclient.NewPatchedKubernetesServiceConnectionRequest() // PatchedKubernetesServiceConnectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsKubernetesPartialUpdate(context.Background(), uuid).PatchedKubernetesServiceConnectionRequest(patchedKubernetesServiceConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsKubernetesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsServiceConnectionsKubernetesPartialUpdate`: KubernetesServiceConnection
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsServiceConnectionsKubernetesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Kubernetes Service-Connection. | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Kubernetes Service-Connection.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsKubernetesRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsKubernetesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsServiceConnectionsKubernetesRetrieve`: KubernetesServiceConnection
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsServiceConnectionsKubernetesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Kubernetes Service-Connection. | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Kubernetes Service-Connection.
	kubernetesServiceConnectionRequest := *openapiclient.NewKubernetesServiceConnectionRequest("Name_example") // KubernetesServiceConnectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsKubernetesUpdate(context.Background(), uuid).KubernetesServiceConnectionRequest(kubernetesServiceConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsKubernetesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsServiceConnectionsKubernetesUpdate`: KubernetesServiceConnection
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsServiceConnectionsKubernetesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Kubernetes Service-Connection. | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Kubernetes Service-Connection.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutpostsAPI.OutpostsServiceConnectionsKubernetesUsedByList(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutpostsAPI.OutpostsServiceConnectionsKubernetesUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutpostsServiceConnectionsKubernetesUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `OutpostsAPI.OutpostsServiceConnectionsKubernetesUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Kubernetes Service-Connection. | 

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

