# \CoreAPI

All URIs are relative to */api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreApplicationEntitlementsCreate**](CoreAPI.md#CoreApplicationEntitlementsCreate) | **Post** /core/application_entitlements/ | 
[**CoreApplicationEntitlementsDestroy**](CoreAPI.md#CoreApplicationEntitlementsDestroy) | **Delete** /core/application_entitlements/{pbm_uuid}/ | 
[**CoreApplicationEntitlementsList**](CoreAPI.md#CoreApplicationEntitlementsList) | **Get** /core/application_entitlements/ | 
[**CoreApplicationEntitlementsPartialUpdate**](CoreAPI.md#CoreApplicationEntitlementsPartialUpdate) | **Patch** /core/application_entitlements/{pbm_uuid}/ | 
[**CoreApplicationEntitlementsRetrieve**](CoreAPI.md#CoreApplicationEntitlementsRetrieve) | **Get** /core/application_entitlements/{pbm_uuid}/ | 
[**CoreApplicationEntitlementsUpdate**](CoreAPI.md#CoreApplicationEntitlementsUpdate) | **Put** /core/application_entitlements/{pbm_uuid}/ | 
[**CoreApplicationEntitlementsUsedByList**](CoreAPI.md#CoreApplicationEntitlementsUsedByList) | **Get** /core/application_entitlements/{pbm_uuid}/used_by/ | 
[**CoreApplicationsCheckAccessRetrieve**](CoreAPI.md#CoreApplicationsCheckAccessRetrieve) | **Get** /core/applications/{slug}/check_access/ | 
[**CoreApplicationsCreate**](CoreAPI.md#CoreApplicationsCreate) | **Post** /core/applications/ | 
[**CoreApplicationsDestroy**](CoreAPI.md#CoreApplicationsDestroy) | **Delete** /core/applications/{slug}/ | 
[**CoreApplicationsList**](CoreAPI.md#CoreApplicationsList) | **Get** /core/applications/ | 
[**CoreApplicationsPartialUpdate**](CoreAPI.md#CoreApplicationsPartialUpdate) | **Patch** /core/applications/{slug}/ | 
[**CoreApplicationsRetrieve**](CoreAPI.md#CoreApplicationsRetrieve) | **Get** /core/applications/{slug}/ | 
[**CoreApplicationsSetIconCreate**](CoreAPI.md#CoreApplicationsSetIconCreate) | **Post** /core/applications/{slug}/set_icon/ | 
[**CoreApplicationsSetIconUrlCreate**](CoreAPI.md#CoreApplicationsSetIconUrlCreate) | **Post** /core/applications/{slug}/set_icon_url/ | 
[**CoreApplicationsUpdate**](CoreAPI.md#CoreApplicationsUpdate) | **Put** /core/applications/{slug}/ | 
[**CoreApplicationsUsedByList**](CoreAPI.md#CoreApplicationsUsedByList) | **Get** /core/applications/{slug}/used_by/ | 
[**CoreAuthenticatedSessionsDestroy**](CoreAPI.md#CoreAuthenticatedSessionsDestroy) | **Delete** /core/authenticated_sessions/{uuid}/ | 
[**CoreAuthenticatedSessionsList**](CoreAPI.md#CoreAuthenticatedSessionsList) | **Get** /core/authenticated_sessions/ | 
[**CoreAuthenticatedSessionsRetrieve**](CoreAPI.md#CoreAuthenticatedSessionsRetrieve) | **Get** /core/authenticated_sessions/{uuid}/ | 
[**CoreAuthenticatedSessionsUsedByList**](CoreAPI.md#CoreAuthenticatedSessionsUsedByList) | **Get** /core/authenticated_sessions/{uuid}/used_by/ | 
[**CoreBrandsCreate**](CoreAPI.md#CoreBrandsCreate) | **Post** /core/brands/ | 
[**CoreBrandsCurrentRetrieve**](CoreAPI.md#CoreBrandsCurrentRetrieve) | **Get** /core/brands/current/ | 
[**CoreBrandsDestroy**](CoreAPI.md#CoreBrandsDestroy) | **Delete** /core/brands/{brand_uuid}/ | 
[**CoreBrandsList**](CoreAPI.md#CoreBrandsList) | **Get** /core/brands/ | 
[**CoreBrandsPartialUpdate**](CoreAPI.md#CoreBrandsPartialUpdate) | **Patch** /core/brands/{brand_uuid}/ | 
[**CoreBrandsRetrieve**](CoreAPI.md#CoreBrandsRetrieve) | **Get** /core/brands/{brand_uuid}/ | 
[**CoreBrandsUpdate**](CoreAPI.md#CoreBrandsUpdate) | **Put** /core/brands/{brand_uuid}/ | 
[**CoreBrandsUsedByList**](CoreAPI.md#CoreBrandsUsedByList) | **Get** /core/brands/{brand_uuid}/used_by/ | 
[**CoreGroupsAddUserCreate**](CoreAPI.md#CoreGroupsAddUserCreate) | **Post** /core/groups/{group_uuid}/add_user/ | 
[**CoreGroupsCreate**](CoreAPI.md#CoreGroupsCreate) | **Post** /core/groups/ | 
[**CoreGroupsDestroy**](CoreAPI.md#CoreGroupsDestroy) | **Delete** /core/groups/{group_uuid}/ | 
[**CoreGroupsList**](CoreAPI.md#CoreGroupsList) | **Get** /core/groups/ | 
[**CoreGroupsPartialUpdate**](CoreAPI.md#CoreGroupsPartialUpdate) | **Patch** /core/groups/{group_uuid}/ | 
[**CoreGroupsRemoveUserCreate**](CoreAPI.md#CoreGroupsRemoveUserCreate) | **Post** /core/groups/{group_uuid}/remove_user/ | 
[**CoreGroupsRetrieve**](CoreAPI.md#CoreGroupsRetrieve) | **Get** /core/groups/{group_uuid}/ | 
[**CoreGroupsUpdate**](CoreAPI.md#CoreGroupsUpdate) | **Put** /core/groups/{group_uuid}/ | 
[**CoreGroupsUsedByList**](CoreAPI.md#CoreGroupsUsedByList) | **Get** /core/groups/{group_uuid}/used_by/ | 
[**CoreTokensCreate**](CoreAPI.md#CoreTokensCreate) | **Post** /core/tokens/ | 
[**CoreTokensDestroy**](CoreAPI.md#CoreTokensDestroy) | **Delete** /core/tokens/{identifier}/ | 
[**CoreTokensList**](CoreAPI.md#CoreTokensList) | **Get** /core/tokens/ | 
[**CoreTokensPartialUpdate**](CoreAPI.md#CoreTokensPartialUpdate) | **Patch** /core/tokens/{identifier}/ | 
[**CoreTokensRetrieve**](CoreAPI.md#CoreTokensRetrieve) | **Get** /core/tokens/{identifier}/ | 
[**CoreTokensSetKeyCreate**](CoreAPI.md#CoreTokensSetKeyCreate) | **Post** /core/tokens/{identifier}/set_key/ | 
[**CoreTokensUpdate**](CoreAPI.md#CoreTokensUpdate) | **Put** /core/tokens/{identifier}/ | 
[**CoreTokensUsedByList**](CoreAPI.md#CoreTokensUsedByList) | **Get** /core/tokens/{identifier}/used_by/ | 
[**CoreTokensViewKeyRetrieve**](CoreAPI.md#CoreTokensViewKeyRetrieve) | **Get** /core/tokens/{identifier}/view_key/ | 
[**CoreTransactionalApplicationsUpdate**](CoreAPI.md#CoreTransactionalApplicationsUpdate) | **Put** /core/transactional/applications/ | 
[**CoreUserConsentDestroy**](CoreAPI.md#CoreUserConsentDestroy) | **Delete** /core/user_consent/{id}/ | 
[**CoreUserConsentList**](CoreAPI.md#CoreUserConsentList) | **Get** /core/user_consent/ | 
[**CoreUserConsentRetrieve**](CoreAPI.md#CoreUserConsentRetrieve) | **Get** /core/user_consent/{id}/ | 
[**CoreUserConsentUsedByList**](CoreAPI.md#CoreUserConsentUsedByList) | **Get** /core/user_consent/{id}/used_by/ | 
[**CoreUsersCreate**](CoreAPI.md#CoreUsersCreate) | **Post** /core/users/ | 
[**CoreUsersDestroy**](CoreAPI.md#CoreUsersDestroy) | **Delete** /core/users/{id}/ | 
[**CoreUsersImpersonateCreate**](CoreAPI.md#CoreUsersImpersonateCreate) | **Post** /core/users/{id}/impersonate/ | 
[**CoreUsersImpersonateEndRetrieve**](CoreAPI.md#CoreUsersImpersonateEndRetrieve) | **Get** /core/users/impersonate_end/ | 
[**CoreUsersList**](CoreAPI.md#CoreUsersList) | **Get** /core/users/ | 
[**CoreUsersMeRetrieve**](CoreAPI.md#CoreUsersMeRetrieve) | **Get** /core/users/me/ | 
[**CoreUsersPartialUpdate**](CoreAPI.md#CoreUsersPartialUpdate) | **Patch** /core/users/{id}/ | 
[**CoreUsersPathsRetrieve**](CoreAPI.md#CoreUsersPathsRetrieve) | **Get** /core/users/paths/ | 
[**CoreUsersRecoveryCreate**](CoreAPI.md#CoreUsersRecoveryCreate) | **Post** /core/users/{id}/recovery/ | 
[**CoreUsersRecoveryEmailCreate**](CoreAPI.md#CoreUsersRecoveryEmailCreate) | **Post** /core/users/{id}/recovery_email/ | 
[**CoreUsersRetrieve**](CoreAPI.md#CoreUsersRetrieve) | **Get** /core/users/{id}/ | 
[**CoreUsersServiceAccountCreate**](CoreAPI.md#CoreUsersServiceAccountCreate) | **Post** /core/users/service_account/ | 
[**CoreUsersSetPasswordCreate**](CoreAPI.md#CoreUsersSetPasswordCreate) | **Post** /core/users/{id}/set_password/ | 
[**CoreUsersUpdate**](CoreAPI.md#CoreUsersUpdate) | **Put** /core/users/{id}/ | 
[**CoreUsersUsedByList**](CoreAPI.md#CoreUsersUsedByList) | **Get** /core/users/{id}/used_by/ | 



## CoreApplicationEntitlementsCreate

> ApplicationEntitlement CoreApplicationEntitlementsCreate(ctx).ApplicationEntitlementRequest(applicationEntitlementRequest).Execute()





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
	applicationEntitlementRequest := *openapiclient.NewApplicationEntitlementRequest("Name_example", "App_example") // ApplicationEntitlementRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreApplicationEntitlementsCreate(context.Background()).ApplicationEntitlementRequest(applicationEntitlementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreApplicationEntitlementsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreApplicationEntitlementsCreate`: ApplicationEntitlement
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreApplicationEntitlementsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreApplicationEntitlementsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationEntitlementRequest** | [**ApplicationEntitlementRequest**](ApplicationEntitlementRequest.md) |  | 

### Return type

[**ApplicationEntitlement**](ApplicationEntitlement.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreApplicationEntitlementsDestroy

> CoreApplicationEntitlementsDestroy(ctx, pbmUuid).Execute()





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
	pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Application Entitlement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreApplicationEntitlementsDestroy(context.Background(), pbmUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreApplicationEntitlementsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this Application Entitlement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreApplicationEntitlementsDestroyRequest struct via the builder pattern


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


## CoreApplicationEntitlementsList

> PaginatedApplicationEntitlementList CoreApplicationEntitlementsList(ctx).App(app).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PbmUuid(pbmUuid).Search(search).Execute()





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
	app := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	name := "name_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreApplicationEntitlementsList(context.Background()).App(app).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PbmUuid(pbmUuid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreApplicationEntitlementsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreApplicationEntitlementsList`: PaginatedApplicationEntitlementList
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreApplicationEntitlementsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreApplicationEntitlementsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **pbmUuid** | **string** |  | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedApplicationEntitlementList**](PaginatedApplicationEntitlementList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreApplicationEntitlementsPartialUpdate

> ApplicationEntitlement CoreApplicationEntitlementsPartialUpdate(ctx, pbmUuid).PatchedApplicationEntitlementRequest(patchedApplicationEntitlementRequest).Execute()





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
	pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Application Entitlement.
	patchedApplicationEntitlementRequest := *openapiclient.NewPatchedApplicationEntitlementRequest() // PatchedApplicationEntitlementRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreApplicationEntitlementsPartialUpdate(context.Background(), pbmUuid).PatchedApplicationEntitlementRequest(patchedApplicationEntitlementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreApplicationEntitlementsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreApplicationEntitlementsPartialUpdate`: ApplicationEntitlement
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreApplicationEntitlementsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this Application Entitlement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreApplicationEntitlementsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedApplicationEntitlementRequest** | [**PatchedApplicationEntitlementRequest**](PatchedApplicationEntitlementRequest.md) |  | 

### Return type

[**ApplicationEntitlement**](ApplicationEntitlement.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreApplicationEntitlementsRetrieve

> ApplicationEntitlement CoreApplicationEntitlementsRetrieve(ctx, pbmUuid).Execute()





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
	pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Application Entitlement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreApplicationEntitlementsRetrieve(context.Background(), pbmUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreApplicationEntitlementsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreApplicationEntitlementsRetrieve`: ApplicationEntitlement
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreApplicationEntitlementsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this Application Entitlement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreApplicationEntitlementsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationEntitlement**](ApplicationEntitlement.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreApplicationEntitlementsUpdate

> ApplicationEntitlement CoreApplicationEntitlementsUpdate(ctx, pbmUuid).ApplicationEntitlementRequest(applicationEntitlementRequest).Execute()





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
	pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Application Entitlement.
	applicationEntitlementRequest := *openapiclient.NewApplicationEntitlementRequest("Name_example", "App_example") // ApplicationEntitlementRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreApplicationEntitlementsUpdate(context.Background(), pbmUuid).ApplicationEntitlementRequest(applicationEntitlementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreApplicationEntitlementsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreApplicationEntitlementsUpdate`: ApplicationEntitlement
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreApplicationEntitlementsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this Application Entitlement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreApplicationEntitlementsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationEntitlementRequest** | [**ApplicationEntitlementRequest**](ApplicationEntitlementRequest.md) |  | 

### Return type

[**ApplicationEntitlement**](ApplicationEntitlement.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreApplicationEntitlementsUsedByList

> []UsedBy CoreApplicationEntitlementsUsedByList(ctx, pbmUuid).Execute()





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
	pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Application Entitlement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreApplicationEntitlementsUsedByList(context.Background(), pbmUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreApplicationEntitlementsUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreApplicationEntitlementsUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreApplicationEntitlementsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this Application Entitlement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreApplicationEntitlementsUsedByListRequest struct via the builder pattern


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


## CoreApplicationsCheckAccessRetrieve

> PolicyTestResult CoreApplicationsCheckAccessRetrieve(ctx, slug).ForUser(forUser).Execute()





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
	slug := "slug_example" // string | 
	forUser := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreApplicationsCheckAccessRetrieve(context.Background(), slug).ForUser(forUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreApplicationsCheckAccessRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreApplicationsCheckAccessRetrieve`: PolicyTestResult
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreApplicationsCheckAccessRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreApplicationsCheckAccessRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forUser** | **int32** |  | 

### Return type

[**PolicyTestResult**](PolicyTestResult.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreApplicationsCreate

> Application CoreApplicationsCreate(ctx).ApplicationRequest(applicationRequest).Execute()





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
	applicationRequest := *openapiclient.NewApplicationRequest("Name_example", "Slug_example") // ApplicationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreApplicationsCreate(context.Background()).ApplicationRequest(applicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreApplicationsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreApplicationsCreate`: Application
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreApplicationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreApplicationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationRequest** | [**ApplicationRequest**](ApplicationRequest.md) |  | 

### Return type

[**Application**](Application.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreApplicationsDestroy

> CoreApplicationsDestroy(ctx, slug).Execute()





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
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreApplicationsDestroy(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreApplicationsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreApplicationsDestroyRequest struct via the builder pattern


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


## CoreApplicationsList

> PaginatedApplicationList CoreApplicationsList(ctx).ForUser(forUser).Group(group).MetaDescription(metaDescription).MetaLaunchUrl(metaLaunchUrl).MetaPublisher(metaPublisher).Name(name).OnlyWithLaunchUrl(onlyWithLaunchUrl).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Slug(slug).SuperuserFullList(superuserFullList).Execute()





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
	forUser := int32(56) // int32 |  (optional)
	group := "group_example" // string |  (optional)
	metaDescription := "metaDescription_example" // string |  (optional)
	metaLaunchUrl := "metaLaunchUrl_example" // string |  (optional)
	metaPublisher := "metaPublisher_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	onlyWithLaunchUrl := true // bool |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)
	slug := "slug_example" // string |  (optional)
	superuserFullList := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreApplicationsList(context.Background()).ForUser(forUser).Group(group).MetaDescription(metaDescription).MetaLaunchUrl(metaLaunchUrl).MetaPublisher(metaPublisher).Name(name).OnlyWithLaunchUrl(onlyWithLaunchUrl).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Slug(slug).SuperuserFullList(superuserFullList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreApplicationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreApplicationsList`: PaginatedApplicationList
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreApplicationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreApplicationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forUser** | **int32** |  | 
 **group** | **string** |  | 
 **metaDescription** | **string** |  | 
 **metaLaunchUrl** | **string** |  | 
 **metaPublisher** | **string** |  | 
 **name** | **string** |  | 
 **onlyWithLaunchUrl** | **bool** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **slug** | **string** |  | 
 **superuserFullList** | **bool** |  | 

### Return type

[**PaginatedApplicationList**](PaginatedApplicationList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreApplicationsPartialUpdate

> Application CoreApplicationsPartialUpdate(ctx, slug).PatchedApplicationRequest(patchedApplicationRequest).Execute()





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
	slug := "slug_example" // string | 
	patchedApplicationRequest := *openapiclient.NewPatchedApplicationRequest() // PatchedApplicationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreApplicationsPartialUpdate(context.Background(), slug).PatchedApplicationRequest(patchedApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreApplicationsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreApplicationsPartialUpdate`: Application
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreApplicationsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreApplicationsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedApplicationRequest** | [**PatchedApplicationRequest**](PatchedApplicationRequest.md) |  | 

### Return type

[**Application**](Application.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreApplicationsRetrieve

> Application CoreApplicationsRetrieve(ctx, slug).Execute()





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
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreApplicationsRetrieve(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreApplicationsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreApplicationsRetrieve`: Application
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreApplicationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreApplicationsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Application**](Application.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreApplicationsSetIconCreate

> CoreApplicationsSetIconCreate(ctx, slug).File(file).Clear(clear).Execute()





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
	slug := "slug_example" // string | 
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)
	clear := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreApplicationsSetIconCreate(context.Background(), slug).File(file).Clear(clear).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreApplicationsSetIconCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreApplicationsSetIconCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 
 **clear** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreApplicationsSetIconUrlCreate

> CoreApplicationsSetIconUrlCreate(ctx, slug).FilePathRequest(filePathRequest).Execute()





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
	slug := "slug_example" // string | 
	filePathRequest := *openapiclient.NewFilePathRequest("Url_example") // FilePathRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreApplicationsSetIconUrlCreate(context.Background(), slug).FilePathRequest(filePathRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreApplicationsSetIconUrlCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreApplicationsSetIconUrlCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filePathRequest** | [**FilePathRequest**](FilePathRequest.md) |  | 

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


## CoreApplicationsUpdate

> Application CoreApplicationsUpdate(ctx, slug).ApplicationRequest(applicationRequest).Execute()





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
	slug := "slug_example" // string | 
	applicationRequest := *openapiclient.NewApplicationRequest("Name_example", "Slug_example") // ApplicationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreApplicationsUpdate(context.Background(), slug).ApplicationRequest(applicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreApplicationsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreApplicationsUpdate`: Application
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreApplicationsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreApplicationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationRequest** | [**ApplicationRequest**](ApplicationRequest.md) |  | 

### Return type

[**Application**](Application.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreApplicationsUsedByList

> []UsedBy CoreApplicationsUsedByList(ctx, slug).Execute()





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
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreApplicationsUsedByList(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreApplicationsUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreApplicationsUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreApplicationsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreApplicationsUsedByListRequest struct via the builder pattern


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


## CoreAuthenticatedSessionsDestroy

> CoreAuthenticatedSessionsDestroy(ctx, uuid).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreAuthenticatedSessionsDestroy(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreAuthenticatedSessionsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreAuthenticatedSessionsDestroyRequest struct via the builder pattern


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


## CoreAuthenticatedSessionsList

> PaginatedAuthenticatedSessionList CoreAuthenticatedSessionsList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).SessionLastIp(sessionLastIp).SessionLastUserAgent(sessionLastUserAgent).UserUsername(userUsername).Execute()





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
	sessionLastIp := "sessionLastIp_example" // string |  (optional)
	sessionLastUserAgent := "sessionLastUserAgent_example" // string |  (optional)
	userUsername := "userUsername_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreAuthenticatedSessionsList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).SessionLastIp(sessionLastIp).SessionLastUserAgent(sessionLastUserAgent).UserUsername(userUsername).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreAuthenticatedSessionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreAuthenticatedSessionsList`: PaginatedAuthenticatedSessionList
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreAuthenticatedSessionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreAuthenticatedSessionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **sessionLastIp** | **string** |  | 
 **sessionLastUserAgent** | **string** |  | 
 **userUsername** | **string** |  | 

### Return type

[**PaginatedAuthenticatedSessionList**](PaginatedAuthenticatedSessionList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreAuthenticatedSessionsRetrieve

> AuthenticatedSession CoreAuthenticatedSessionsRetrieve(ctx, uuid).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreAuthenticatedSessionsRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreAuthenticatedSessionsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreAuthenticatedSessionsRetrieve`: AuthenticatedSession
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreAuthenticatedSessionsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreAuthenticatedSessionsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticatedSession**](AuthenticatedSession.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreAuthenticatedSessionsUsedByList

> []UsedBy CoreAuthenticatedSessionsUsedByList(ctx, uuid).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreAuthenticatedSessionsUsedByList(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreAuthenticatedSessionsUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreAuthenticatedSessionsUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreAuthenticatedSessionsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreAuthenticatedSessionsUsedByListRequest struct via the builder pattern


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


## CoreBrandsCreate

> Brand CoreBrandsCreate(ctx).BrandRequest(brandRequest).Execute()





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
	brandRequest := *openapiclient.NewBrandRequest("Domain_example") // BrandRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreBrandsCreate(context.Background()).BrandRequest(brandRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreBrandsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreBrandsCreate`: Brand
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreBrandsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreBrandsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **brandRequest** | [**BrandRequest**](BrandRequest.md) |  | 

### Return type

[**Brand**](Brand.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreBrandsCurrentRetrieve

> CurrentBrand CoreBrandsCurrentRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.CoreAPI.CoreBrandsCurrentRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreBrandsCurrentRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreBrandsCurrentRetrieve`: CurrentBrand
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreBrandsCurrentRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCoreBrandsCurrentRetrieveRequest struct via the builder pattern


### Return type

[**CurrentBrand**](CurrentBrand.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreBrandsDestroy

> CoreBrandsDestroy(ctx, brandUuid).Execute()





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
	brandUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Brand.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreBrandsDestroy(context.Background(), brandUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreBrandsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandUuid** | **string** | A UUID string identifying this Brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreBrandsDestroyRequest struct via the builder pattern


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


## CoreBrandsList

> PaginatedBrandList CoreBrandsList(ctx).BrandUuid(brandUuid).BrandingDefaultFlowBackground(brandingDefaultFlowBackground).BrandingFavicon(brandingFavicon).BrandingLogo(brandingLogo).BrandingTitle(brandingTitle).ClientCertificates(clientCertificates).Default_(default_).Domain(domain).FlowAuthentication(flowAuthentication).FlowDeviceCode(flowDeviceCode).FlowInvalidation(flowInvalidation).FlowRecovery(flowRecovery).FlowUnenrollment(flowUnenrollment).FlowUserSettings(flowUserSettings).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).WebCertificate(webCertificate).Execute()





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
	brandUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	brandingDefaultFlowBackground := "brandingDefaultFlowBackground_example" // string |  (optional)
	brandingFavicon := "brandingFavicon_example" // string |  (optional)
	brandingLogo := "brandingLogo_example" // string |  (optional)
	brandingTitle := "brandingTitle_example" // string |  (optional)
	clientCertificates := []string{"Inner_example"} // []string |  (optional)
	default_ := true // bool |  (optional)
	domain := "domain_example" // string |  (optional)
	flowAuthentication := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	flowDeviceCode := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	flowInvalidation := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	flowRecovery := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	flowUnenrollment := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	flowUserSettings := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)
	webCertificate := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreBrandsList(context.Background()).BrandUuid(brandUuid).BrandingDefaultFlowBackground(brandingDefaultFlowBackground).BrandingFavicon(brandingFavicon).BrandingLogo(brandingLogo).BrandingTitle(brandingTitle).ClientCertificates(clientCertificates).Default_(default_).Domain(domain).FlowAuthentication(flowAuthentication).FlowDeviceCode(flowDeviceCode).FlowInvalidation(flowInvalidation).FlowRecovery(flowRecovery).FlowUnenrollment(flowUnenrollment).FlowUserSettings(flowUserSettings).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).WebCertificate(webCertificate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreBrandsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreBrandsList`: PaginatedBrandList
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreBrandsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreBrandsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **brandUuid** | **string** |  | 
 **brandingDefaultFlowBackground** | **string** |  | 
 **brandingFavicon** | **string** |  | 
 **brandingLogo** | **string** |  | 
 **brandingTitle** | **string** |  | 
 **clientCertificates** | **[]string** |  | 
 **default_** | **bool** |  | 
 **domain** | **string** |  | 
 **flowAuthentication** | **string** |  | 
 **flowDeviceCode** | **string** |  | 
 **flowInvalidation** | **string** |  | 
 **flowRecovery** | **string** |  | 
 **flowUnenrollment** | **string** |  | 
 **flowUserSettings** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **webCertificate** | **string** |  | 

### Return type

[**PaginatedBrandList**](PaginatedBrandList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreBrandsPartialUpdate

> Brand CoreBrandsPartialUpdate(ctx, brandUuid).PatchedBrandRequest(patchedBrandRequest).Execute()





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
	brandUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Brand.
	patchedBrandRequest := *openapiclient.NewPatchedBrandRequest() // PatchedBrandRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreBrandsPartialUpdate(context.Background(), brandUuid).PatchedBrandRequest(patchedBrandRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreBrandsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreBrandsPartialUpdate`: Brand
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreBrandsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandUuid** | **string** | A UUID string identifying this Brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreBrandsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedBrandRequest** | [**PatchedBrandRequest**](PatchedBrandRequest.md) |  | 

### Return type

[**Brand**](Brand.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreBrandsRetrieve

> Brand CoreBrandsRetrieve(ctx, brandUuid).Execute()





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
	brandUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Brand.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreBrandsRetrieve(context.Background(), brandUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreBrandsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreBrandsRetrieve`: Brand
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreBrandsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandUuid** | **string** | A UUID string identifying this Brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreBrandsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Brand**](Brand.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreBrandsUpdate

> Brand CoreBrandsUpdate(ctx, brandUuid).BrandRequest(brandRequest).Execute()





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
	brandUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Brand.
	brandRequest := *openapiclient.NewBrandRequest("Domain_example") // BrandRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreBrandsUpdate(context.Background(), brandUuid).BrandRequest(brandRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreBrandsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreBrandsUpdate`: Brand
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreBrandsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandUuid** | **string** | A UUID string identifying this Brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreBrandsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **brandRequest** | [**BrandRequest**](BrandRequest.md) |  | 

### Return type

[**Brand**](Brand.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreBrandsUsedByList

> []UsedBy CoreBrandsUsedByList(ctx, brandUuid).Execute()





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
	brandUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Brand.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreBrandsUsedByList(context.Background(), brandUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreBrandsUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreBrandsUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreBrandsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandUuid** | **string** | A UUID string identifying this Brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreBrandsUsedByListRequest struct via the builder pattern


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


## CoreGroupsAddUserCreate

> CoreGroupsAddUserCreate(ctx, groupUuid).UserAccountRequest(userAccountRequest).Execute()





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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Group.
	userAccountRequest := *openapiclient.NewUserAccountRequest(int32(123)) // UserAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreGroupsAddUserCreate(context.Background(), groupUuid).UserAccountRequest(userAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreGroupsAddUserCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | A UUID string identifying this Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreGroupsAddUserCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userAccountRequest** | [**UserAccountRequest**](UserAccountRequest.md) |  | 

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


## CoreGroupsCreate

> Group CoreGroupsCreate(ctx).GroupRequest(groupRequest).Execute()





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
	groupRequest := *openapiclient.NewGroupRequest("Name_example") // GroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreGroupsCreate(context.Background()).GroupRequest(groupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreGroupsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreGroupsCreate`: Group
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreGroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupRequest** | [**GroupRequest**](GroupRequest.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreGroupsDestroy

> CoreGroupsDestroy(ctx, groupUuid).Execute()





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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreGroupsDestroy(context.Background(), groupUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreGroupsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | A UUID string identifying this Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreGroupsDestroyRequest struct via the builder pattern


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


## CoreGroupsList

> PaginatedGroupList CoreGroupsList(ctx).Attributes(attributes).IncludeChildren(includeChildren).IncludeUsers(includeUsers).IsSuperuser(isSuperuser).MembersByPk(membersByPk).MembersByUsername(membersByUsername).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
	attributes := "attributes_example" // string | Attributes (optional)
	includeChildren := true // bool |  (optional) (default to false)
	includeUsers := true // bool |  (optional) (default to true)
	isSuperuser := true // bool |  (optional)
	membersByPk := []int32{int32(123)} // []int32 |  (optional)
	membersByUsername := []string{"Inner_example"} // []string | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (optional)
	name := "name_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreGroupsList(context.Background()).Attributes(attributes).IncludeChildren(includeChildren).IncludeUsers(includeUsers).IsSuperuser(isSuperuser).MembersByPk(membersByPk).MembersByUsername(membersByUsername).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreGroupsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreGroupsList`: PaginatedGroupList
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreGroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributes** | **string** | Attributes | 
 **includeChildren** | **bool** |  | [default to false]
 **includeUsers** | **bool** |  | [default to true]
 **isSuperuser** | **bool** |  | 
 **membersByPk** | **[]int32** |  | 
 **membersByUsername** | **[]string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedGroupList**](PaginatedGroupList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreGroupsPartialUpdate

> Group CoreGroupsPartialUpdate(ctx, groupUuid).PatchedGroupRequest(patchedGroupRequest).Execute()





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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Group.
	patchedGroupRequest := *openapiclient.NewPatchedGroupRequest() // PatchedGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreGroupsPartialUpdate(context.Background(), groupUuid).PatchedGroupRequest(patchedGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreGroupsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreGroupsPartialUpdate`: Group
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreGroupsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | A UUID string identifying this Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreGroupsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedGroupRequest** | [**PatchedGroupRequest**](PatchedGroupRequest.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreGroupsRemoveUserCreate

> CoreGroupsRemoveUserCreate(ctx, groupUuid).UserAccountRequest(userAccountRequest).Execute()





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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Group.
	userAccountRequest := *openapiclient.NewUserAccountRequest(int32(123)) // UserAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreGroupsRemoveUserCreate(context.Background(), groupUuid).UserAccountRequest(userAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreGroupsRemoveUserCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | A UUID string identifying this Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreGroupsRemoveUserCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userAccountRequest** | [**UserAccountRequest**](UserAccountRequest.md) |  | 

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


## CoreGroupsRetrieve

> Group CoreGroupsRetrieve(ctx, groupUuid).IncludeChildren(includeChildren).IncludeUsers(includeUsers).Execute()





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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Group.
	includeChildren := true // bool |  (optional) (default to false)
	includeUsers := true // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreGroupsRetrieve(context.Background(), groupUuid).IncludeChildren(includeChildren).IncludeUsers(includeUsers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreGroupsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreGroupsRetrieve`: Group
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreGroupsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | A UUID string identifying this Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreGroupsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeChildren** | **bool** |  | [default to false]
 **includeUsers** | **bool** |  | [default to true]

### Return type

[**Group**](Group.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreGroupsUpdate

> Group CoreGroupsUpdate(ctx, groupUuid).GroupRequest(groupRequest).Execute()





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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Group.
	groupRequest := *openapiclient.NewGroupRequest("Name_example") // GroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreGroupsUpdate(context.Background(), groupUuid).GroupRequest(groupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreGroupsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreGroupsUpdate`: Group
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreGroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | A UUID string identifying this Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreGroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupRequest** | [**GroupRequest**](GroupRequest.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreGroupsUsedByList

> []UsedBy CoreGroupsUsedByList(ctx, groupUuid).Execute()





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
	groupUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreGroupsUsedByList(context.Background(), groupUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreGroupsUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreGroupsUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreGroupsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | A UUID string identifying this Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreGroupsUsedByListRequest struct via the builder pattern


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


## CoreTokensCreate

> Token CoreTokensCreate(ctx).TokenRequest(tokenRequest).Execute()





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
	tokenRequest := *openapiclient.NewTokenRequest("Identifier_example") // TokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreTokensCreate(context.Background()).TokenRequest(tokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreTokensCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreTokensCreate`: Token
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreTokensCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreTokensCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreTokensDestroy

> CoreTokensDestroy(ctx, identifier).Execute()





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
	identifier := "identifier_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreTokensDestroy(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreTokensDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreTokensDestroyRequest struct via the builder pattern


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


## CoreTokensList

> PaginatedTokenList CoreTokensList(ctx).Description(description).Expires(expires).Expiring(expiring).Identifier(identifier).Intent(intent).Managed(managed).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).UserUsername(userUsername).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string |  (optional)
	expires := time.Now() // time.Time |  (optional)
	expiring := true // bool |  (optional)
	identifier := "identifier_example" // string |  (optional)
	intent := "intent_example" // string |  (optional)
	managed := "managed_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)
	userUsername := "userUsername_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreTokensList(context.Background()).Description(description).Expires(expires).Expiring(expiring).Identifier(identifier).Intent(intent).Managed(managed).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).UserUsername(userUsername).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreTokensList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreTokensList`: PaginatedTokenList
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreTokensList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreTokensListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** |  | 
 **expires** | **time.Time** |  | 
 **expiring** | **bool** |  | 
 **identifier** | **string** |  | 
 **intent** | **string** |  | 
 **managed** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **userUsername** | **string** |  | 

### Return type

[**PaginatedTokenList**](PaginatedTokenList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreTokensPartialUpdate

> Token CoreTokensPartialUpdate(ctx, identifier).PatchedTokenRequest(patchedTokenRequest).Execute()





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
	identifier := "identifier_example" // string | 
	patchedTokenRequest := *openapiclient.NewPatchedTokenRequest() // PatchedTokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreTokensPartialUpdate(context.Background(), identifier).PatchedTokenRequest(patchedTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreTokensPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreTokensPartialUpdate`: Token
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreTokensPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreTokensPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedTokenRequest** | [**PatchedTokenRequest**](PatchedTokenRequest.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreTokensRetrieve

> Token CoreTokensRetrieve(ctx, identifier).Execute()





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
	identifier := "identifier_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreTokensRetrieve(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreTokensRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreTokensRetrieve`: Token
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreTokensRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreTokensRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Token**](Token.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreTokensSetKeyCreate

> CoreTokensSetKeyCreate(ctx, identifier).TokenSetKeyRequest(tokenSetKeyRequest).Execute()





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
	identifier := "identifier_example" // string | 
	tokenSetKeyRequest := *openapiclient.NewTokenSetKeyRequest("Key_example") // TokenSetKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreTokensSetKeyCreate(context.Background(), identifier).TokenSetKeyRequest(tokenSetKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreTokensSetKeyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreTokensSetKeyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenSetKeyRequest** | [**TokenSetKeyRequest**](TokenSetKeyRequest.md) |  | 

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


## CoreTokensUpdate

> Token CoreTokensUpdate(ctx, identifier).TokenRequest(tokenRequest).Execute()





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
	identifier := "identifier_example" // string | 
	tokenRequest := *openapiclient.NewTokenRequest("Identifier_example") // TokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreTokensUpdate(context.Background(), identifier).TokenRequest(tokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreTokensUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreTokensUpdate`: Token
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreTokensUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreTokensUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreTokensUsedByList

> []UsedBy CoreTokensUsedByList(ctx, identifier).Execute()





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
	identifier := "identifier_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreTokensUsedByList(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreTokensUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreTokensUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreTokensUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreTokensUsedByListRequest struct via the builder pattern


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


## CoreTokensViewKeyRetrieve

> TokenView CoreTokensViewKeyRetrieve(ctx, identifier).Execute()





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
	identifier := "identifier_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreTokensViewKeyRetrieve(context.Background(), identifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreTokensViewKeyRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreTokensViewKeyRetrieve`: TokenView
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreTokensViewKeyRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreTokensViewKeyRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenView**](TokenView.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreTransactionalApplicationsUpdate

> TransactionApplicationResponse CoreTransactionalApplicationsUpdate(ctx).TransactionApplicationRequest(transactionApplicationRequest).Execute()





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
	transactionApplicationRequest := *openapiclient.NewTransactionApplicationRequest(*openapiclient.NewApplicationRequest("Name_example", "Slug_example"), openapiclient.ProviderModelEnum("authentik_providers_google_workspace.googleworkspaceprovider"), openapiclient.modelRequest{GoogleWorkspaceProviderRequest: openapiclient.NewGoogleWorkspaceProviderRequest("Name_example", "DelegatedSubject_example", map[string]interface{}{"key": interface{}(123)}, "DefaultGroupEmailDomain_example")}) // TransactionApplicationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreTransactionalApplicationsUpdate(context.Background()).TransactionApplicationRequest(transactionApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreTransactionalApplicationsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreTransactionalApplicationsUpdate`: TransactionApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreTransactionalApplicationsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreTransactionalApplicationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionApplicationRequest** | [**TransactionApplicationRequest**](TransactionApplicationRequest.md) |  | 

### Return type

[**TransactionApplicationResponse**](TransactionApplicationResponse.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreUserConsentDestroy

> CoreUserConsentDestroy(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this User Consent.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreUserConsentDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUserConsentDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this User Consent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreUserConsentDestroyRequest struct via the builder pattern


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


## CoreUserConsentList

> PaginatedUserConsentList CoreUserConsentList(ctx).Application(application).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).User(user).Execute()





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
	application := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)
	user := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreUserConsentList(context.Background()).Application(application).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUserConsentList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreUserConsentList`: PaginatedUserConsentList
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreUserConsentList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreUserConsentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **user** | **int32** |  | 

### Return type

[**PaginatedUserConsentList**](PaginatedUserConsentList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreUserConsentRetrieve

> UserConsent CoreUserConsentRetrieve(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this User Consent.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreUserConsentRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUserConsentRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreUserConsentRetrieve`: UserConsent
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreUserConsentRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this User Consent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreUserConsentRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserConsent**](UserConsent.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreUserConsentUsedByList

> []UsedBy CoreUserConsentUsedByList(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this User Consent.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreUserConsentUsedByList(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUserConsentUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreUserConsentUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreUserConsentUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this User Consent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreUserConsentUsedByListRequest struct via the builder pattern


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


## CoreUsersCreate

> User CoreUsersCreate(ctx).UserRequest(userRequest).Execute()





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
	userRequest := *openapiclient.NewUserRequest("Username_example", "Name_example") // UserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreUsersCreate(context.Background()).UserRequest(userRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUsersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreUsersCreate`: User
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreUsersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreUsersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRequest** | [**UserRequest**](UserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreUsersDestroy

> CoreUsersDestroy(ctx, id).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreUsersDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUsersDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCoreUsersDestroyRequest struct via the builder pattern


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


## CoreUsersImpersonateCreate

> CoreUsersImpersonateCreate(ctx, id).ImpersonationRequest(impersonationRequest).Execute()





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
	impersonationRequest := *openapiclient.NewImpersonationRequest("Reason_example") // ImpersonationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreUsersImpersonateCreate(context.Background(), id).ImpersonationRequest(impersonationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUsersImpersonateCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCoreUsersImpersonateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **impersonationRequest** | [**ImpersonationRequest**](ImpersonationRequest.md) |  | 

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


## CoreUsersImpersonateEndRetrieve

> CoreUsersImpersonateEndRetrieve(ctx).Execute()





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
	r, err := apiClient.CoreAPI.CoreUsersImpersonateEndRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUsersImpersonateEndRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCoreUsersImpersonateEndRetrieveRequest struct via the builder pattern


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


## CoreUsersList

> PaginatedUserList CoreUsersList(ctx).Attributes(attributes).DateJoined(dateJoined).DateJoinedGt(dateJoinedGt).DateJoinedLt(dateJoinedLt).Email(email).GroupsByName(groupsByName).GroupsByPk(groupsByPk).IncludeGroups(includeGroups).IsActive(isActive).IsSuperuser(isSuperuser).LastUpdated(lastUpdated).LastUpdatedGt(lastUpdatedGt).LastUpdatedLt(lastUpdatedLt).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Path(path).PathStartswith(pathStartswith).Search(search).Type_(type_).Username(username).Uuid(uuid).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	attributes := "attributes_example" // string | Attributes (optional)
	dateJoined := time.Now() // time.Time |  (optional)
	dateJoinedGt := time.Now() // time.Time |  (optional)
	dateJoinedLt := time.Now() // time.Time |  (optional)
	email := "email_example" // string |  (optional)
	groupsByName := []string{"Inner_example"} // []string |  (optional)
	groupsByPk := []string{"Inner_example"} // []string |  (optional)
	includeGroups := true // bool |  (optional) (default to true)
	isActive := true // bool |  (optional)
	isSuperuser := true // bool |  (optional)
	lastUpdated := time.Now() // time.Time |  (optional)
	lastUpdatedGt := time.Now() // time.Time |  (optional)
	lastUpdatedLt := time.Now() // time.Time |  (optional)
	name := "name_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	path := "path_example" // string |  (optional)
	pathStartswith := "pathStartswith_example" // string |  (optional)
	search := "search_example" // string | A search term. (optional)
	type_ := []string{"Type_example"} // []string |  (optional)
	username := "username_example" // string |  (optional)
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreUsersList(context.Background()).Attributes(attributes).DateJoined(dateJoined).DateJoinedGt(dateJoinedGt).DateJoinedLt(dateJoinedLt).Email(email).GroupsByName(groupsByName).GroupsByPk(groupsByPk).IncludeGroups(includeGroups).IsActive(isActive).IsSuperuser(isSuperuser).LastUpdated(lastUpdated).LastUpdatedGt(lastUpdatedGt).LastUpdatedLt(lastUpdatedLt).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Path(path).PathStartswith(pathStartswith).Search(search).Type_(type_).Username(username).Uuid(uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUsersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreUsersList`: PaginatedUserList
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreUsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributes** | **string** | Attributes | 
 **dateJoined** | **time.Time** |  | 
 **dateJoinedGt** | **time.Time** |  | 
 **dateJoinedLt** | **time.Time** |  | 
 **email** | **string** |  | 
 **groupsByName** | **[]string** |  | 
 **groupsByPk** | **[]string** |  | 
 **includeGroups** | **bool** |  | [default to true]
 **isActive** | **bool** |  | 
 **isSuperuser** | **bool** |  | 
 **lastUpdated** | **time.Time** |  | 
 **lastUpdatedGt** | **time.Time** |  | 
 **lastUpdatedLt** | **time.Time** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **path** | **string** |  | 
 **pathStartswith** | **string** |  | 
 **search** | **string** | A search term. | 
 **type_** | **[]string** |  | 
 **username** | **string** |  | 
 **uuid** | **string** |  | 

### Return type

[**PaginatedUserList**](PaginatedUserList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreUsersMeRetrieve

> SessionUser CoreUsersMeRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.CoreAPI.CoreUsersMeRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUsersMeRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreUsersMeRetrieve`: SessionUser
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreUsersMeRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCoreUsersMeRetrieveRequest struct via the builder pattern


### Return type

[**SessionUser**](SessionUser.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreUsersPartialUpdate

> User CoreUsersPartialUpdate(ctx, id).PatchedUserRequest(patchedUserRequest).Execute()





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
	patchedUserRequest := *openapiclient.NewPatchedUserRequest() // PatchedUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreUsersPartialUpdate(context.Background(), id).PatchedUserRequest(patchedUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUsersPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreUsersPartialUpdate`: User
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreUsersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreUsersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedUserRequest** | [**PatchedUserRequest**](PatchedUserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreUsersPathsRetrieve

> UserPath CoreUsersPathsRetrieve(ctx).Search(search).Execute()





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
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreUsersPathsRetrieve(context.Background()).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUsersPathsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreUsersPathsRetrieve`: UserPath
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreUsersPathsRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreUsersPathsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | A search term. | 

### Return type

[**UserPath**](UserPath.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreUsersRecoveryCreate

> Link CoreUsersRecoveryCreate(ctx, id).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreUsersRecoveryCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUsersRecoveryCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreUsersRecoveryCreate`: Link
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreUsersRecoveryCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreUsersRecoveryCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Link**](Link.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreUsersRecoveryEmailCreate

> CoreUsersRecoveryEmailCreate(ctx, id).EmailStage(emailStage).Execute()





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
	emailStage := "emailStage_example" // string | 
	id := int32(56) // int32 | A unique integer value identifying this User.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreUsersRecoveryEmailCreate(context.Background(), id).EmailStage(emailStage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUsersRecoveryEmailCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCoreUsersRecoveryEmailCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailStage** | **string** |  | 


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


## CoreUsersRetrieve

> User CoreUsersRetrieve(ctx, id).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreUsersRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUsersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreUsersRetrieve`: User
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreUsersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreUsersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreUsersServiceAccountCreate

> UserServiceAccountResponse CoreUsersServiceAccountCreate(ctx).UserServiceAccountRequest(userServiceAccountRequest).Execute()





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
	userServiceAccountRequest := *openapiclient.NewUserServiceAccountRequest("Name_example") // UserServiceAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreUsersServiceAccountCreate(context.Background()).UserServiceAccountRequest(userServiceAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUsersServiceAccountCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreUsersServiceAccountCreate`: UserServiceAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreUsersServiceAccountCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreUsersServiceAccountCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userServiceAccountRequest** | [**UserServiceAccountRequest**](UserServiceAccountRequest.md) |  | 

### Return type

[**UserServiceAccountResponse**](UserServiceAccountResponse.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreUsersSetPasswordCreate

> CoreUsersSetPasswordCreate(ctx, id).UserPasswordSetRequest(userPasswordSetRequest).Execute()





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
	userPasswordSetRequest := *openapiclient.NewUserPasswordSetRequest("Password_example") // UserPasswordSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreUsersSetPasswordCreate(context.Background(), id).UserPasswordSetRequest(userPasswordSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUsersSetPasswordCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCoreUsersSetPasswordCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userPasswordSetRequest** | [**UserPasswordSetRequest**](UserPasswordSetRequest.md) |  | 

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


## CoreUsersUpdate

> User CoreUsersUpdate(ctx, id).UserRequest(userRequest).Execute()





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
	userRequest := *openapiclient.NewUserRequest("Username_example", "Name_example") // UserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreUsersUpdate(context.Background(), id).UserRequest(userRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUsersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreUsersUpdate`: User
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreUsersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreUsersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userRequest** | [**UserRequest**](UserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreUsersUsedByList

> []UsedBy CoreUsersUsedByList(ctx, id).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreUsersUsedByList(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreUsersUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreUsersUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreUsersUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreUsersUsedByListRequest struct via the builder pattern


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

