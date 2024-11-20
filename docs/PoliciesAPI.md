# \PoliciesAPI

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesAllCacheClearCreate**](PoliciesAPI.md#PoliciesAllCacheClearCreate) | **Post** /policies/all/cache_clear/ | 
[**PoliciesAllCacheInfoRetrieve**](PoliciesAPI.md#PoliciesAllCacheInfoRetrieve) | **Get** /policies/all/cache_info/ | 
[**PoliciesAllDestroy**](PoliciesAPI.md#PoliciesAllDestroy) | **Delete** /policies/all/{policy_uuid}/ | 
[**PoliciesAllList**](PoliciesAPI.md#PoliciesAllList) | **Get** /policies/all/ | 
[**PoliciesAllRetrieve**](PoliciesAPI.md#PoliciesAllRetrieve) | **Get** /policies/all/{policy_uuid}/ | 
[**PoliciesAllTestCreate**](PoliciesAPI.md#PoliciesAllTestCreate) | **Post** /policies/all/{policy_uuid}/test/ | 
[**PoliciesAllTypesList**](PoliciesAPI.md#PoliciesAllTypesList) | **Get** /policies/all/types/ | 
[**PoliciesAllUsedByList**](PoliciesAPI.md#PoliciesAllUsedByList) | **Get** /policies/all/{policy_uuid}/used_by/ | 
[**PoliciesBindingsCreate**](PoliciesAPI.md#PoliciesBindingsCreate) | **Post** /policies/bindings/ | 
[**PoliciesBindingsDestroy**](PoliciesAPI.md#PoliciesBindingsDestroy) | **Delete** /policies/bindings/{policy_binding_uuid}/ | 
[**PoliciesBindingsList**](PoliciesAPI.md#PoliciesBindingsList) | **Get** /policies/bindings/ | 
[**PoliciesBindingsPartialUpdate**](PoliciesAPI.md#PoliciesBindingsPartialUpdate) | **Patch** /policies/bindings/{policy_binding_uuid}/ | 
[**PoliciesBindingsRetrieve**](PoliciesAPI.md#PoliciesBindingsRetrieve) | **Get** /policies/bindings/{policy_binding_uuid}/ | 
[**PoliciesBindingsUpdate**](PoliciesAPI.md#PoliciesBindingsUpdate) | **Put** /policies/bindings/{policy_binding_uuid}/ | 
[**PoliciesBindingsUsedByList**](PoliciesAPI.md#PoliciesBindingsUsedByList) | **Get** /policies/bindings/{policy_binding_uuid}/used_by/ | 
[**PoliciesDummyCreate**](PoliciesAPI.md#PoliciesDummyCreate) | **Post** /policies/dummy/ | 
[**PoliciesDummyDestroy**](PoliciesAPI.md#PoliciesDummyDestroy) | **Delete** /policies/dummy/{policy_uuid}/ | 
[**PoliciesDummyList**](PoliciesAPI.md#PoliciesDummyList) | **Get** /policies/dummy/ | 
[**PoliciesDummyPartialUpdate**](PoliciesAPI.md#PoliciesDummyPartialUpdate) | **Patch** /policies/dummy/{policy_uuid}/ | 
[**PoliciesDummyRetrieve**](PoliciesAPI.md#PoliciesDummyRetrieve) | **Get** /policies/dummy/{policy_uuid}/ | 
[**PoliciesDummyUpdate**](PoliciesAPI.md#PoliciesDummyUpdate) | **Put** /policies/dummy/{policy_uuid}/ | 
[**PoliciesDummyUsedByList**](PoliciesAPI.md#PoliciesDummyUsedByList) | **Get** /policies/dummy/{policy_uuid}/used_by/ | 
[**PoliciesEventMatcherCreate**](PoliciesAPI.md#PoliciesEventMatcherCreate) | **Post** /policies/event_matcher/ | 
[**PoliciesEventMatcherDestroy**](PoliciesAPI.md#PoliciesEventMatcherDestroy) | **Delete** /policies/event_matcher/{policy_uuid}/ | 
[**PoliciesEventMatcherList**](PoliciesAPI.md#PoliciesEventMatcherList) | **Get** /policies/event_matcher/ | 
[**PoliciesEventMatcherPartialUpdate**](PoliciesAPI.md#PoliciesEventMatcherPartialUpdate) | **Patch** /policies/event_matcher/{policy_uuid}/ | 
[**PoliciesEventMatcherRetrieve**](PoliciesAPI.md#PoliciesEventMatcherRetrieve) | **Get** /policies/event_matcher/{policy_uuid}/ | 
[**PoliciesEventMatcherUpdate**](PoliciesAPI.md#PoliciesEventMatcherUpdate) | **Put** /policies/event_matcher/{policy_uuid}/ | 
[**PoliciesEventMatcherUsedByList**](PoliciesAPI.md#PoliciesEventMatcherUsedByList) | **Get** /policies/event_matcher/{policy_uuid}/used_by/ | 
[**PoliciesExpressionCreate**](PoliciesAPI.md#PoliciesExpressionCreate) | **Post** /policies/expression/ | 
[**PoliciesExpressionDestroy**](PoliciesAPI.md#PoliciesExpressionDestroy) | **Delete** /policies/expression/{policy_uuid}/ | 
[**PoliciesExpressionList**](PoliciesAPI.md#PoliciesExpressionList) | **Get** /policies/expression/ | 
[**PoliciesExpressionPartialUpdate**](PoliciesAPI.md#PoliciesExpressionPartialUpdate) | **Patch** /policies/expression/{policy_uuid}/ | 
[**PoliciesExpressionRetrieve**](PoliciesAPI.md#PoliciesExpressionRetrieve) | **Get** /policies/expression/{policy_uuid}/ | 
[**PoliciesExpressionUpdate**](PoliciesAPI.md#PoliciesExpressionUpdate) | **Put** /policies/expression/{policy_uuid}/ | 
[**PoliciesExpressionUsedByList**](PoliciesAPI.md#PoliciesExpressionUsedByList) | **Get** /policies/expression/{policy_uuid}/used_by/ | 
[**PoliciesGeoipCreate**](PoliciesAPI.md#PoliciesGeoipCreate) | **Post** /policies/geoip/ | 
[**PoliciesGeoipDestroy**](PoliciesAPI.md#PoliciesGeoipDestroy) | **Delete** /policies/geoip/{policy_uuid}/ | 
[**PoliciesGeoipIso3166List**](PoliciesAPI.md#PoliciesGeoipIso3166List) | **Get** /policies/geoip_iso3166/ | 
[**PoliciesGeoipList**](PoliciesAPI.md#PoliciesGeoipList) | **Get** /policies/geoip/ | 
[**PoliciesGeoipPartialUpdate**](PoliciesAPI.md#PoliciesGeoipPartialUpdate) | **Patch** /policies/geoip/{policy_uuid}/ | 
[**PoliciesGeoipRetrieve**](PoliciesAPI.md#PoliciesGeoipRetrieve) | **Get** /policies/geoip/{policy_uuid}/ | 
[**PoliciesGeoipUpdate**](PoliciesAPI.md#PoliciesGeoipUpdate) | **Put** /policies/geoip/{policy_uuid}/ | 
[**PoliciesGeoipUsedByList**](PoliciesAPI.md#PoliciesGeoipUsedByList) | **Get** /policies/geoip/{policy_uuid}/used_by/ | 
[**PoliciesPasswordCreate**](PoliciesAPI.md#PoliciesPasswordCreate) | **Post** /policies/password/ | 
[**PoliciesPasswordDestroy**](PoliciesAPI.md#PoliciesPasswordDestroy) | **Delete** /policies/password/{policy_uuid}/ | 
[**PoliciesPasswordExpiryCreate**](PoliciesAPI.md#PoliciesPasswordExpiryCreate) | **Post** /policies/password_expiry/ | 
[**PoliciesPasswordExpiryDestroy**](PoliciesAPI.md#PoliciesPasswordExpiryDestroy) | **Delete** /policies/password_expiry/{policy_uuid}/ | 
[**PoliciesPasswordExpiryList**](PoliciesAPI.md#PoliciesPasswordExpiryList) | **Get** /policies/password_expiry/ | 
[**PoliciesPasswordExpiryPartialUpdate**](PoliciesAPI.md#PoliciesPasswordExpiryPartialUpdate) | **Patch** /policies/password_expiry/{policy_uuid}/ | 
[**PoliciesPasswordExpiryRetrieve**](PoliciesAPI.md#PoliciesPasswordExpiryRetrieve) | **Get** /policies/password_expiry/{policy_uuid}/ | 
[**PoliciesPasswordExpiryUpdate**](PoliciesAPI.md#PoliciesPasswordExpiryUpdate) | **Put** /policies/password_expiry/{policy_uuid}/ | 
[**PoliciesPasswordExpiryUsedByList**](PoliciesAPI.md#PoliciesPasswordExpiryUsedByList) | **Get** /policies/password_expiry/{policy_uuid}/used_by/ | 
[**PoliciesPasswordList**](PoliciesAPI.md#PoliciesPasswordList) | **Get** /policies/password/ | 
[**PoliciesPasswordPartialUpdate**](PoliciesAPI.md#PoliciesPasswordPartialUpdate) | **Patch** /policies/password/{policy_uuid}/ | 
[**PoliciesPasswordRetrieve**](PoliciesAPI.md#PoliciesPasswordRetrieve) | **Get** /policies/password/{policy_uuid}/ | 
[**PoliciesPasswordUpdate**](PoliciesAPI.md#PoliciesPasswordUpdate) | **Put** /policies/password/{policy_uuid}/ | 
[**PoliciesPasswordUsedByList**](PoliciesAPI.md#PoliciesPasswordUsedByList) | **Get** /policies/password/{policy_uuid}/used_by/ | 
[**PoliciesReputationCreate**](PoliciesAPI.md#PoliciesReputationCreate) | **Post** /policies/reputation/ | 
[**PoliciesReputationDestroy**](PoliciesAPI.md#PoliciesReputationDestroy) | **Delete** /policies/reputation/{policy_uuid}/ | 
[**PoliciesReputationList**](PoliciesAPI.md#PoliciesReputationList) | **Get** /policies/reputation/ | 
[**PoliciesReputationPartialUpdate**](PoliciesAPI.md#PoliciesReputationPartialUpdate) | **Patch** /policies/reputation/{policy_uuid}/ | 
[**PoliciesReputationRetrieve**](PoliciesAPI.md#PoliciesReputationRetrieve) | **Get** /policies/reputation/{policy_uuid}/ | 
[**PoliciesReputationScoresDestroy**](PoliciesAPI.md#PoliciesReputationScoresDestroy) | **Delete** /policies/reputation/scores/{reputation_uuid}/ | 
[**PoliciesReputationScoresList**](PoliciesAPI.md#PoliciesReputationScoresList) | **Get** /policies/reputation/scores/ | 
[**PoliciesReputationScoresRetrieve**](PoliciesAPI.md#PoliciesReputationScoresRetrieve) | **Get** /policies/reputation/scores/{reputation_uuid}/ | 
[**PoliciesReputationScoresUsedByList**](PoliciesAPI.md#PoliciesReputationScoresUsedByList) | **Get** /policies/reputation/scores/{reputation_uuid}/used_by/ | 
[**PoliciesReputationUpdate**](PoliciesAPI.md#PoliciesReputationUpdate) | **Put** /policies/reputation/{policy_uuid}/ | 
[**PoliciesReputationUsedByList**](PoliciesAPI.md#PoliciesReputationUsedByList) | **Get** /policies/reputation/{policy_uuid}/used_by/ | 



## PoliciesAllCacheClearCreate

> PoliciesAllCacheClearCreate(ctx).Execute()





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
	r, err := apiClient.PoliciesAPI.PoliciesAllCacheClearCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesAllCacheClearCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesAllCacheClearCreateRequest struct via the builder pattern


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


## PoliciesAllCacheInfoRetrieve

> Cache PoliciesAllCacheInfoRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.PoliciesAPI.PoliciesAllCacheInfoRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesAllCacheInfoRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesAllCacheInfoRetrieve`: Cache
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesAllCacheInfoRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesAllCacheInfoRetrieveRequest struct via the builder pattern


### Return type

[**Cache**](Cache.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesAllDestroy

> PoliciesAllDestroy(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoliciesAPI.PoliciesAllDestroy(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesAllDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesAllDestroyRequest struct via the builder pattern


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


## PoliciesAllList

> PaginatedPolicyList PoliciesAllList(ctx).BindingsIsnull(bindingsIsnull).Ordering(ordering).Page(page).PageSize(pageSize).PromptstageIsnull(promptstageIsnull).Search(search).Execute()





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
	bindingsIsnull := true // bool |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	promptstageIsnull := true // bool |  (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesAllList(context.Background()).BindingsIsnull(bindingsIsnull).Ordering(ordering).Page(page).PageSize(pageSize).PromptstageIsnull(promptstageIsnull).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesAllList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesAllList`: PaginatedPolicyList
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesAllList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesAllListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindingsIsnull** | **bool** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **promptstageIsnull** | **bool** |  | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedPolicyList**](PaginatedPolicyList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesAllRetrieve

> Policy PoliciesAllRetrieve(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesAllRetrieve(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesAllRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesAllRetrieve`: Policy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesAllRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesAllRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Policy**](Policy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesAllTestCreate

> PolicyTestResult PoliciesAllTestCreate(ctx, policyUuid).PolicyTestRequest(policyTestRequest).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Policy.
	policyTestRequest := *openapiclient.NewPolicyTestRequest(int32(123)) // PolicyTestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesAllTestCreate(context.Background(), policyUuid).PolicyTestRequest(policyTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesAllTestCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesAllTestCreate`: PolicyTestResult
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesAllTestCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesAllTestCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyTestRequest** | [**PolicyTestRequest**](PolicyTestRequest.md) |  | 

### Return type

[**PolicyTestResult**](PolicyTestResult.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesAllTypesList

> []TypeCreate PoliciesAllTypesList(ctx).Execute()





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
	resp, r, err := apiClient.PoliciesAPI.PoliciesAllTypesList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesAllTypesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesAllTypesList`: []TypeCreate
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesAllTypesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesAllTypesListRequest struct via the builder pattern


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


## PoliciesAllUsedByList

> []UsedBy PoliciesAllUsedByList(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesAllUsedByList(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesAllUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesAllUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesAllUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesAllUsedByListRequest struct via the builder pattern


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


## PoliciesBindingsCreate

> PolicyBinding PoliciesBindingsCreate(ctx).PolicyBindingRequest(policyBindingRequest).Execute()





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
	policyBindingRequest := *openapiclient.NewPolicyBindingRequest("Target_example", int32(123)) // PolicyBindingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesBindingsCreate(context.Background()).PolicyBindingRequest(policyBindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesBindingsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesBindingsCreate`: PolicyBinding
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesBindingsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesBindingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyBindingRequest** | [**PolicyBindingRequest**](PolicyBindingRequest.md) |  | 

### Return type

[**PolicyBinding**](PolicyBinding.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesBindingsDestroy

> PoliciesBindingsDestroy(ctx, policyBindingUuid).Execute()





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
	policyBindingUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Policy Binding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoliciesAPI.PoliciesBindingsDestroy(context.Background(), policyBindingUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesBindingsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyBindingUuid** | **string** | A UUID string identifying this Policy Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesBindingsDestroyRequest struct via the builder pattern


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


## PoliciesBindingsList

> PaginatedPolicyBindingList PoliciesBindingsList(ctx).Enabled(enabled).Order(order).Ordering(ordering).Page(page).PageSize(pageSize).Policy(policy).PolicyIsnull(policyIsnull).Search(search).Target(target).TargetIn(targetIn).Timeout(timeout).Execute()





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
	enabled := true // bool |  (optional)
	order := int32(56) // int32 |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	policy := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	policyIsnull := true // bool |  (optional)
	search := "search_example" // string | A search term. (optional)
	target := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	targetIn := []string{"Inner_example"} // []string |  (optional)
	timeout := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesBindingsList(context.Background()).Enabled(enabled).Order(order).Ordering(ordering).Page(page).PageSize(pageSize).Policy(policy).PolicyIsnull(policyIsnull).Search(search).Target(target).TargetIn(targetIn).Timeout(timeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesBindingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesBindingsList`: PaginatedPolicyBindingList
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesBindingsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesBindingsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabled** | **bool** |  | 
 **order** | **int32** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **policy** | **string** |  | 
 **policyIsnull** | **bool** |  | 
 **search** | **string** | A search term. | 
 **target** | **string** |  | 
 **targetIn** | **[]string** |  | 
 **timeout** | **int32** |  | 

### Return type

[**PaginatedPolicyBindingList**](PaginatedPolicyBindingList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesBindingsPartialUpdate

> PolicyBinding PoliciesBindingsPartialUpdate(ctx, policyBindingUuid).PatchedPolicyBindingRequest(patchedPolicyBindingRequest).Execute()





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
	policyBindingUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Policy Binding.
	patchedPolicyBindingRequest := *openapiclient.NewPatchedPolicyBindingRequest() // PatchedPolicyBindingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesBindingsPartialUpdate(context.Background(), policyBindingUuid).PatchedPolicyBindingRequest(patchedPolicyBindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesBindingsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesBindingsPartialUpdate`: PolicyBinding
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesBindingsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyBindingUuid** | **string** | A UUID string identifying this Policy Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesBindingsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPolicyBindingRequest** | [**PatchedPolicyBindingRequest**](PatchedPolicyBindingRequest.md) |  | 

### Return type

[**PolicyBinding**](PolicyBinding.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesBindingsRetrieve

> PolicyBinding PoliciesBindingsRetrieve(ctx, policyBindingUuid).Execute()





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
	policyBindingUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Policy Binding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesBindingsRetrieve(context.Background(), policyBindingUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesBindingsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesBindingsRetrieve`: PolicyBinding
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesBindingsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyBindingUuid** | **string** | A UUID string identifying this Policy Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesBindingsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyBinding**](PolicyBinding.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesBindingsUpdate

> PolicyBinding PoliciesBindingsUpdate(ctx, policyBindingUuid).PolicyBindingRequest(policyBindingRequest).Execute()





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
	policyBindingUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Policy Binding.
	policyBindingRequest := *openapiclient.NewPolicyBindingRequest("Target_example", int32(123)) // PolicyBindingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesBindingsUpdate(context.Background(), policyBindingUuid).PolicyBindingRequest(policyBindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesBindingsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesBindingsUpdate`: PolicyBinding
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesBindingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyBindingUuid** | **string** | A UUID string identifying this Policy Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesBindingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyBindingRequest** | [**PolicyBindingRequest**](PolicyBindingRequest.md) |  | 

### Return type

[**PolicyBinding**](PolicyBinding.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesBindingsUsedByList

> []UsedBy PoliciesBindingsUsedByList(ctx, policyBindingUuid).Execute()





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
	policyBindingUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Policy Binding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesBindingsUsedByList(context.Background(), policyBindingUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesBindingsUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesBindingsUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesBindingsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyBindingUuid** | **string** | A UUID string identifying this Policy Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesBindingsUsedByListRequest struct via the builder pattern


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


## PoliciesDummyCreate

> DummyPolicy PoliciesDummyCreate(ctx).DummyPolicyRequest(dummyPolicyRequest).Execute()





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
	dummyPolicyRequest := *openapiclient.NewDummyPolicyRequest("Name_example") // DummyPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesDummyCreate(context.Background()).DummyPolicyRequest(dummyPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesDummyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesDummyCreate`: DummyPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesDummyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDummyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dummyPolicyRequest** | [**DummyPolicyRequest**](DummyPolicyRequest.md) |  | 

### Return type

[**DummyPolicy**](DummyPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesDummyDestroy

> PoliciesDummyDestroy(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Dummy Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoliciesAPI.PoliciesDummyDestroy(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesDummyDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Dummy Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDummyDestroyRequest struct via the builder pattern


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


## PoliciesDummyList

> PaginatedDummyPolicyList PoliciesDummyList(ctx).Created(created).ExecutionLogging(executionLogging).LastUpdated(lastUpdated).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PolicyUuid(policyUuid).Result(result).Search(search).WaitMax(waitMax).WaitMin(waitMin).Execute()





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
	created := time.Now() // time.Time |  (optional)
	executionLogging := true // bool |  (optional)
	lastUpdated := time.Now() // time.Time |  (optional)
	name := "name_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	result := true // bool |  (optional)
	search := "search_example" // string | A search term. (optional)
	waitMax := int32(56) // int32 |  (optional)
	waitMin := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesDummyList(context.Background()).Created(created).ExecutionLogging(executionLogging).LastUpdated(lastUpdated).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PolicyUuid(policyUuid).Result(result).Search(search).WaitMax(waitMax).WaitMin(waitMin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesDummyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesDummyList`: PaginatedDummyPolicyList
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesDummyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDummyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created** | **time.Time** |  | 
 **executionLogging** | **bool** |  | 
 **lastUpdated** | **time.Time** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **policyUuid** | **string** |  | 
 **result** | **bool** |  | 
 **search** | **string** | A search term. | 
 **waitMax** | **int32** |  | 
 **waitMin** | **int32** |  | 

### Return type

[**PaginatedDummyPolicyList**](PaginatedDummyPolicyList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesDummyPartialUpdate

> DummyPolicy PoliciesDummyPartialUpdate(ctx, policyUuid).PatchedDummyPolicyRequest(patchedDummyPolicyRequest).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Dummy Policy.
	patchedDummyPolicyRequest := *openapiclient.NewPatchedDummyPolicyRequest() // PatchedDummyPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesDummyPartialUpdate(context.Background(), policyUuid).PatchedDummyPolicyRequest(patchedDummyPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesDummyPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesDummyPartialUpdate`: DummyPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesDummyPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Dummy Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDummyPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDummyPolicyRequest** | [**PatchedDummyPolicyRequest**](PatchedDummyPolicyRequest.md) |  | 

### Return type

[**DummyPolicy**](DummyPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesDummyRetrieve

> DummyPolicy PoliciesDummyRetrieve(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Dummy Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesDummyRetrieve(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesDummyRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesDummyRetrieve`: DummyPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesDummyRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Dummy Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDummyRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DummyPolicy**](DummyPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesDummyUpdate

> DummyPolicy PoliciesDummyUpdate(ctx, policyUuid).DummyPolicyRequest(dummyPolicyRequest).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Dummy Policy.
	dummyPolicyRequest := *openapiclient.NewDummyPolicyRequest("Name_example") // DummyPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesDummyUpdate(context.Background(), policyUuid).DummyPolicyRequest(dummyPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesDummyUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesDummyUpdate`: DummyPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesDummyUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Dummy Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDummyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dummyPolicyRequest** | [**DummyPolicyRequest**](DummyPolicyRequest.md) |  | 

### Return type

[**DummyPolicy**](DummyPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesDummyUsedByList

> []UsedBy PoliciesDummyUsedByList(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Dummy Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesDummyUsedByList(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesDummyUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesDummyUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesDummyUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Dummy Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDummyUsedByListRequest struct via the builder pattern


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


## PoliciesEventMatcherCreate

> EventMatcherPolicy PoliciesEventMatcherCreate(ctx).EventMatcherPolicyRequest(eventMatcherPolicyRequest).Execute()





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
	eventMatcherPolicyRequest := *openapiclient.NewEventMatcherPolicyRequest("Name_example") // EventMatcherPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesEventMatcherCreate(context.Background()).EventMatcherPolicyRequest(eventMatcherPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesEventMatcherCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesEventMatcherCreate`: EventMatcherPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesEventMatcherCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesEventMatcherCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventMatcherPolicyRequest** | [**EventMatcherPolicyRequest**](EventMatcherPolicyRequest.md) |  | 

### Return type

[**EventMatcherPolicy**](EventMatcherPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesEventMatcherDestroy

> PoliciesEventMatcherDestroy(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Event Matcher Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoliciesAPI.PoliciesEventMatcherDestroy(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesEventMatcherDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Event Matcher Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesEventMatcherDestroyRequest struct via the builder pattern


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


## PoliciesEventMatcherList

> PaginatedEventMatcherPolicyList PoliciesEventMatcherList(ctx).Action(action).App(app).ClientIp(clientIp).Created(created).ExecutionLogging(executionLogging).LastUpdated(lastUpdated).Model(model).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PolicyUuid(policyUuid).Search(search).Execute()





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
	action := "action_example" // string | Match created events with this action type. When left empty, all action types will be matched.   (optional)
	app := "app_example" // string |  (optional)
	clientIp := "clientIp_example" // string |  (optional)
	created := time.Now() // time.Time |  (optional)
	executionLogging := true // bool |  (optional)
	lastUpdated := time.Now() // time.Time |  (optional)
	model := "model_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesEventMatcherList(context.Background()).Action(action).App(app).ClientIp(clientIp).Created(created).ExecutionLogging(executionLogging).LastUpdated(lastUpdated).Model(model).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PolicyUuid(policyUuid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesEventMatcherList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesEventMatcherList`: PaginatedEventMatcherPolicyList
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesEventMatcherList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesEventMatcherListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Match created events with this action type. When left empty, all action types will be matched.   | 
 **app** | **string** |  | 
 **clientIp** | **string** |  | 
 **created** | **time.Time** |  | 
 **executionLogging** | **bool** |  | 
 **lastUpdated** | **time.Time** |  | 
 **model** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **policyUuid** | **string** |  | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedEventMatcherPolicyList**](PaginatedEventMatcherPolicyList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesEventMatcherPartialUpdate

> EventMatcherPolicy PoliciesEventMatcherPartialUpdate(ctx, policyUuid).PatchedEventMatcherPolicyRequest(patchedEventMatcherPolicyRequest).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Event Matcher Policy.
	patchedEventMatcherPolicyRequest := *openapiclient.NewPatchedEventMatcherPolicyRequest() // PatchedEventMatcherPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesEventMatcherPartialUpdate(context.Background(), policyUuid).PatchedEventMatcherPolicyRequest(patchedEventMatcherPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesEventMatcherPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesEventMatcherPartialUpdate`: EventMatcherPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesEventMatcherPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Event Matcher Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesEventMatcherPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedEventMatcherPolicyRequest** | [**PatchedEventMatcherPolicyRequest**](PatchedEventMatcherPolicyRequest.md) |  | 

### Return type

[**EventMatcherPolicy**](EventMatcherPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesEventMatcherRetrieve

> EventMatcherPolicy PoliciesEventMatcherRetrieve(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Event Matcher Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesEventMatcherRetrieve(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesEventMatcherRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesEventMatcherRetrieve`: EventMatcherPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesEventMatcherRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Event Matcher Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesEventMatcherRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventMatcherPolicy**](EventMatcherPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesEventMatcherUpdate

> EventMatcherPolicy PoliciesEventMatcherUpdate(ctx, policyUuid).EventMatcherPolicyRequest(eventMatcherPolicyRequest).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Event Matcher Policy.
	eventMatcherPolicyRequest := *openapiclient.NewEventMatcherPolicyRequest("Name_example") // EventMatcherPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesEventMatcherUpdate(context.Background(), policyUuid).EventMatcherPolicyRequest(eventMatcherPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesEventMatcherUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesEventMatcherUpdate`: EventMatcherPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesEventMatcherUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Event Matcher Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesEventMatcherUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventMatcherPolicyRequest** | [**EventMatcherPolicyRequest**](EventMatcherPolicyRequest.md) |  | 

### Return type

[**EventMatcherPolicy**](EventMatcherPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesEventMatcherUsedByList

> []UsedBy PoliciesEventMatcherUsedByList(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Event Matcher Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesEventMatcherUsedByList(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesEventMatcherUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesEventMatcherUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesEventMatcherUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Event Matcher Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesEventMatcherUsedByListRequest struct via the builder pattern


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


## PoliciesExpressionCreate

> ExpressionPolicy PoliciesExpressionCreate(ctx).ExpressionPolicyRequest(expressionPolicyRequest).Execute()





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
	expressionPolicyRequest := *openapiclient.NewExpressionPolicyRequest("Name_example", "Expression_example") // ExpressionPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesExpressionCreate(context.Background()).ExpressionPolicyRequest(expressionPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesExpressionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesExpressionCreate`: ExpressionPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesExpressionCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesExpressionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expressionPolicyRequest** | [**ExpressionPolicyRequest**](ExpressionPolicyRequest.md) |  | 

### Return type

[**ExpressionPolicy**](ExpressionPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesExpressionDestroy

> PoliciesExpressionDestroy(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Expression Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoliciesAPI.PoliciesExpressionDestroy(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesExpressionDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Expression Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesExpressionDestroyRequest struct via the builder pattern


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


## PoliciesExpressionList

> PaginatedExpressionPolicyList PoliciesExpressionList(ctx).Created(created).ExecutionLogging(executionLogging).Expression(expression).LastUpdated(lastUpdated).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PolicyUuid(policyUuid).Search(search).Execute()





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
	created := time.Now() // time.Time |  (optional)
	executionLogging := true // bool |  (optional)
	expression := "expression_example" // string |  (optional)
	lastUpdated := time.Now() // time.Time |  (optional)
	name := "name_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesExpressionList(context.Background()).Created(created).ExecutionLogging(executionLogging).Expression(expression).LastUpdated(lastUpdated).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PolicyUuid(policyUuid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesExpressionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesExpressionList`: PaginatedExpressionPolicyList
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesExpressionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesExpressionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created** | **time.Time** |  | 
 **executionLogging** | **bool** |  | 
 **expression** | **string** |  | 
 **lastUpdated** | **time.Time** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **policyUuid** | **string** |  | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedExpressionPolicyList**](PaginatedExpressionPolicyList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesExpressionPartialUpdate

> ExpressionPolicy PoliciesExpressionPartialUpdate(ctx, policyUuid).PatchedExpressionPolicyRequest(patchedExpressionPolicyRequest).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Expression Policy.
	patchedExpressionPolicyRequest := *openapiclient.NewPatchedExpressionPolicyRequest() // PatchedExpressionPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesExpressionPartialUpdate(context.Background(), policyUuid).PatchedExpressionPolicyRequest(patchedExpressionPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesExpressionPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesExpressionPartialUpdate`: ExpressionPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesExpressionPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Expression Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesExpressionPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedExpressionPolicyRequest** | [**PatchedExpressionPolicyRequest**](PatchedExpressionPolicyRequest.md) |  | 

### Return type

[**ExpressionPolicy**](ExpressionPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesExpressionRetrieve

> ExpressionPolicy PoliciesExpressionRetrieve(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Expression Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesExpressionRetrieve(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesExpressionRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesExpressionRetrieve`: ExpressionPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesExpressionRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Expression Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesExpressionRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExpressionPolicy**](ExpressionPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesExpressionUpdate

> ExpressionPolicy PoliciesExpressionUpdate(ctx, policyUuid).ExpressionPolicyRequest(expressionPolicyRequest).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Expression Policy.
	expressionPolicyRequest := *openapiclient.NewExpressionPolicyRequest("Name_example", "Expression_example") // ExpressionPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesExpressionUpdate(context.Background(), policyUuid).ExpressionPolicyRequest(expressionPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesExpressionUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesExpressionUpdate`: ExpressionPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesExpressionUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Expression Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesExpressionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expressionPolicyRequest** | [**ExpressionPolicyRequest**](ExpressionPolicyRequest.md) |  | 

### Return type

[**ExpressionPolicy**](ExpressionPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesExpressionUsedByList

> []UsedBy PoliciesExpressionUsedByList(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Expression Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesExpressionUsedByList(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesExpressionUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesExpressionUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesExpressionUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Expression Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesExpressionUsedByListRequest struct via the builder pattern


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


## PoliciesGeoipCreate

> GeoIPPolicy PoliciesGeoipCreate(ctx).GeoIPPolicyRequest(geoIPPolicyRequest).Execute()





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
	geoIPPolicyRequest := *openapiclient.NewGeoIPPolicyRequest("Name_example", []openapiclient.CountryCodeEnum{openapiclient.CountryCodeEnum("AF")}) // GeoIPPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesGeoipCreate(context.Background()).GeoIPPolicyRequest(geoIPPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesGeoipCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesGeoipCreate`: GeoIPPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesGeoipCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGeoipCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **geoIPPolicyRequest** | [**GeoIPPolicyRequest**](GeoIPPolicyRequest.md) |  | 

### Return type

[**GeoIPPolicy**](GeoIPPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGeoipDestroy

> PoliciesGeoipDestroy(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this GeoIP Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoliciesAPI.PoliciesGeoipDestroy(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesGeoipDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this GeoIP Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGeoipDestroyRequest struct via the builder pattern


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


## PoliciesGeoipIso3166List

> []DetailedCountry PoliciesGeoipIso3166List(ctx).Execute()





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
	resp, r, err := apiClient.PoliciesAPI.PoliciesGeoipIso3166List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesGeoipIso3166List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesGeoipIso3166List`: []DetailedCountry
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesGeoipIso3166List`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGeoipIso3166ListRequest struct via the builder pattern


### Return type

[**[]DetailedCountry**](DetailedCountry.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGeoipList

> PaginatedGeoIPPolicyList PoliciesGeoipList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
	resp, r, err := apiClient.PoliciesAPI.PoliciesGeoipList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesGeoipList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesGeoipList`: PaginatedGeoIPPolicyList
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesGeoipList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGeoipListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedGeoIPPolicyList**](PaginatedGeoIPPolicyList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGeoipPartialUpdate

> GeoIPPolicy PoliciesGeoipPartialUpdate(ctx, policyUuid).PatchedGeoIPPolicyRequest(patchedGeoIPPolicyRequest).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this GeoIP Policy.
	patchedGeoIPPolicyRequest := *openapiclient.NewPatchedGeoIPPolicyRequest() // PatchedGeoIPPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesGeoipPartialUpdate(context.Background(), policyUuid).PatchedGeoIPPolicyRequest(patchedGeoIPPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesGeoipPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesGeoipPartialUpdate`: GeoIPPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesGeoipPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this GeoIP Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGeoipPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedGeoIPPolicyRequest** | [**PatchedGeoIPPolicyRequest**](PatchedGeoIPPolicyRequest.md) |  | 

### Return type

[**GeoIPPolicy**](GeoIPPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGeoipRetrieve

> GeoIPPolicy PoliciesGeoipRetrieve(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this GeoIP Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesGeoipRetrieve(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesGeoipRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesGeoipRetrieve`: GeoIPPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesGeoipRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this GeoIP Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGeoipRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GeoIPPolicy**](GeoIPPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGeoipUpdate

> GeoIPPolicy PoliciesGeoipUpdate(ctx, policyUuid).GeoIPPolicyRequest(geoIPPolicyRequest).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this GeoIP Policy.
	geoIPPolicyRequest := *openapiclient.NewGeoIPPolicyRequest("Name_example", []openapiclient.CountryCodeEnum{openapiclient.CountryCodeEnum("AF")}) // GeoIPPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesGeoipUpdate(context.Background(), policyUuid).GeoIPPolicyRequest(geoIPPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesGeoipUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesGeoipUpdate`: GeoIPPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesGeoipUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this GeoIP Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGeoipUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **geoIPPolicyRequest** | [**GeoIPPolicyRequest**](GeoIPPolicyRequest.md) |  | 

### Return type

[**GeoIPPolicy**](GeoIPPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesGeoipUsedByList

> []UsedBy PoliciesGeoipUsedByList(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this GeoIP Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesGeoipUsedByList(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesGeoipUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesGeoipUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesGeoipUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this GeoIP Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGeoipUsedByListRequest struct via the builder pattern


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


## PoliciesPasswordCreate

> PasswordPolicy PoliciesPasswordCreate(ctx).PasswordPolicyRequest(passwordPolicyRequest).Execute()





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
	passwordPolicyRequest := *openapiclient.NewPasswordPolicyRequest("Name_example") // PasswordPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesPasswordCreate(context.Background()).PasswordPolicyRequest(passwordPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesPasswordCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesPasswordCreate`: PasswordPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesPasswordCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPasswordCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordPolicyRequest** | [**PasswordPolicyRequest**](PasswordPolicyRequest.md) |  | 

### Return type

[**PasswordPolicy**](PasswordPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPasswordDestroy

> PoliciesPasswordDestroy(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Password Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoliciesAPI.PoliciesPasswordDestroy(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesPasswordDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Password Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPasswordDestroyRequest struct via the builder pattern


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


## PoliciesPasswordExpiryCreate

> PasswordExpiryPolicy PoliciesPasswordExpiryCreate(ctx).PasswordExpiryPolicyRequest(passwordExpiryPolicyRequest).Execute()





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
	passwordExpiryPolicyRequest := *openapiclient.NewPasswordExpiryPolicyRequest("Name_example", int32(123)) // PasswordExpiryPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesPasswordExpiryCreate(context.Background()).PasswordExpiryPolicyRequest(passwordExpiryPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesPasswordExpiryCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesPasswordExpiryCreate`: PasswordExpiryPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesPasswordExpiryCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPasswordExpiryCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordExpiryPolicyRequest** | [**PasswordExpiryPolicyRequest**](PasswordExpiryPolicyRequest.md) |  | 

### Return type

[**PasswordExpiryPolicy**](PasswordExpiryPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPasswordExpiryDestroy

> PoliciesPasswordExpiryDestroy(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Password Expiry Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoliciesAPI.PoliciesPasswordExpiryDestroy(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesPasswordExpiryDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Password Expiry Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPasswordExpiryDestroyRequest struct via the builder pattern


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


## PoliciesPasswordExpiryList

> PaginatedPasswordExpiryPolicyList PoliciesPasswordExpiryList(ctx).Created(created).Days(days).DenyOnly(denyOnly).ExecutionLogging(executionLogging).LastUpdated(lastUpdated).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PolicyUuid(policyUuid).Search(search).Execute()





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
	created := time.Now() // time.Time |  (optional)
	days := int32(56) // int32 |  (optional)
	denyOnly := true // bool |  (optional)
	executionLogging := true // bool |  (optional)
	lastUpdated := time.Now() // time.Time |  (optional)
	name := "name_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesPasswordExpiryList(context.Background()).Created(created).Days(days).DenyOnly(denyOnly).ExecutionLogging(executionLogging).LastUpdated(lastUpdated).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PolicyUuid(policyUuid).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesPasswordExpiryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesPasswordExpiryList`: PaginatedPasswordExpiryPolicyList
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesPasswordExpiryList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPasswordExpiryListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created** | **time.Time** |  | 
 **days** | **int32** |  | 
 **denyOnly** | **bool** |  | 
 **executionLogging** | **bool** |  | 
 **lastUpdated** | **time.Time** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **policyUuid** | **string** |  | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedPasswordExpiryPolicyList**](PaginatedPasswordExpiryPolicyList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPasswordExpiryPartialUpdate

> PasswordExpiryPolicy PoliciesPasswordExpiryPartialUpdate(ctx, policyUuid).PatchedPasswordExpiryPolicyRequest(patchedPasswordExpiryPolicyRequest).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Password Expiry Policy.
	patchedPasswordExpiryPolicyRequest := *openapiclient.NewPatchedPasswordExpiryPolicyRequest() // PatchedPasswordExpiryPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesPasswordExpiryPartialUpdate(context.Background(), policyUuid).PatchedPasswordExpiryPolicyRequest(patchedPasswordExpiryPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesPasswordExpiryPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesPasswordExpiryPartialUpdate`: PasswordExpiryPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesPasswordExpiryPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Password Expiry Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPasswordExpiryPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPasswordExpiryPolicyRequest** | [**PatchedPasswordExpiryPolicyRequest**](PatchedPasswordExpiryPolicyRequest.md) |  | 

### Return type

[**PasswordExpiryPolicy**](PasswordExpiryPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPasswordExpiryRetrieve

> PasswordExpiryPolicy PoliciesPasswordExpiryRetrieve(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Password Expiry Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesPasswordExpiryRetrieve(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesPasswordExpiryRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesPasswordExpiryRetrieve`: PasswordExpiryPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesPasswordExpiryRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Password Expiry Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPasswordExpiryRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PasswordExpiryPolicy**](PasswordExpiryPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPasswordExpiryUpdate

> PasswordExpiryPolicy PoliciesPasswordExpiryUpdate(ctx, policyUuid).PasswordExpiryPolicyRequest(passwordExpiryPolicyRequest).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Password Expiry Policy.
	passwordExpiryPolicyRequest := *openapiclient.NewPasswordExpiryPolicyRequest("Name_example", int32(123)) // PasswordExpiryPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesPasswordExpiryUpdate(context.Background(), policyUuid).PasswordExpiryPolicyRequest(passwordExpiryPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesPasswordExpiryUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesPasswordExpiryUpdate`: PasswordExpiryPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesPasswordExpiryUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Password Expiry Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPasswordExpiryUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **passwordExpiryPolicyRequest** | [**PasswordExpiryPolicyRequest**](PasswordExpiryPolicyRequest.md) |  | 

### Return type

[**PasswordExpiryPolicy**](PasswordExpiryPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPasswordExpiryUsedByList

> []UsedBy PoliciesPasswordExpiryUsedByList(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Password Expiry Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesPasswordExpiryUsedByList(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesPasswordExpiryUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesPasswordExpiryUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesPasswordExpiryUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Password Expiry Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPasswordExpiryUsedByListRequest struct via the builder pattern


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


## PoliciesPasswordList

> PaginatedPasswordPolicyList PoliciesPasswordList(ctx).AmountDigits(amountDigits).AmountLowercase(amountLowercase).AmountSymbols(amountSymbols).AmountUppercase(amountUppercase).CheckHaveIBeenPwned(checkHaveIBeenPwned).CheckStaticRules(checkStaticRules).CheckZxcvbn(checkZxcvbn).Created(created).ErrorMessage(errorMessage).ExecutionLogging(executionLogging).HibpAllowedCount(hibpAllowedCount).LastUpdated(lastUpdated).LengthMin(lengthMin).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PasswordField(passwordField).PolicyUuid(policyUuid).Search(search).SymbolCharset(symbolCharset).ZxcvbnScoreThreshold(zxcvbnScoreThreshold).Execute()





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
	amountDigits := int32(56) // int32 |  (optional)
	amountLowercase := int32(56) // int32 |  (optional)
	amountSymbols := int32(56) // int32 |  (optional)
	amountUppercase := int32(56) // int32 |  (optional)
	checkHaveIBeenPwned := true // bool |  (optional)
	checkStaticRules := true // bool |  (optional)
	checkZxcvbn := true // bool |  (optional)
	created := time.Now() // time.Time |  (optional)
	errorMessage := "errorMessage_example" // string |  (optional)
	executionLogging := true // bool |  (optional)
	hibpAllowedCount := int32(56) // int32 |  (optional)
	lastUpdated := time.Now() // time.Time |  (optional)
	lengthMin := int32(56) // int32 |  (optional)
	name := "name_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	passwordField := "passwordField_example" // string |  (optional)
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	search := "search_example" // string | A search term. (optional)
	symbolCharset := "symbolCharset_example" // string |  (optional)
	zxcvbnScoreThreshold := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesPasswordList(context.Background()).AmountDigits(amountDigits).AmountLowercase(amountLowercase).AmountSymbols(amountSymbols).AmountUppercase(amountUppercase).CheckHaveIBeenPwned(checkHaveIBeenPwned).CheckStaticRules(checkStaticRules).CheckZxcvbn(checkZxcvbn).Created(created).ErrorMessage(errorMessage).ExecutionLogging(executionLogging).HibpAllowedCount(hibpAllowedCount).LastUpdated(lastUpdated).LengthMin(lengthMin).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PasswordField(passwordField).PolicyUuid(policyUuid).Search(search).SymbolCharset(symbolCharset).ZxcvbnScoreThreshold(zxcvbnScoreThreshold).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesPasswordList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesPasswordList`: PaginatedPasswordPolicyList
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesPasswordList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPasswordListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amountDigits** | **int32** |  | 
 **amountLowercase** | **int32** |  | 
 **amountSymbols** | **int32** |  | 
 **amountUppercase** | **int32** |  | 
 **checkHaveIBeenPwned** | **bool** |  | 
 **checkStaticRules** | **bool** |  | 
 **checkZxcvbn** | **bool** |  | 
 **created** | **time.Time** |  | 
 **errorMessage** | **string** |  | 
 **executionLogging** | **bool** |  | 
 **hibpAllowedCount** | **int32** |  | 
 **lastUpdated** | **time.Time** |  | 
 **lengthMin** | **int32** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **passwordField** | **string** |  | 
 **policyUuid** | **string** |  | 
 **search** | **string** | A search term. | 
 **symbolCharset** | **string** |  | 
 **zxcvbnScoreThreshold** | **int32** |  | 

### Return type

[**PaginatedPasswordPolicyList**](PaginatedPasswordPolicyList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPasswordPartialUpdate

> PasswordPolicy PoliciesPasswordPartialUpdate(ctx, policyUuid).PatchedPasswordPolicyRequest(patchedPasswordPolicyRequest).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Password Policy.
	patchedPasswordPolicyRequest := *openapiclient.NewPatchedPasswordPolicyRequest() // PatchedPasswordPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesPasswordPartialUpdate(context.Background(), policyUuid).PatchedPasswordPolicyRequest(patchedPasswordPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesPasswordPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesPasswordPartialUpdate`: PasswordPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesPasswordPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Password Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPasswordPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPasswordPolicyRequest** | [**PatchedPasswordPolicyRequest**](PatchedPasswordPolicyRequest.md) |  | 

### Return type

[**PasswordPolicy**](PasswordPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPasswordRetrieve

> PasswordPolicy PoliciesPasswordRetrieve(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Password Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesPasswordRetrieve(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesPasswordRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesPasswordRetrieve`: PasswordPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesPasswordRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Password Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPasswordRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PasswordPolicy**](PasswordPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPasswordUpdate

> PasswordPolicy PoliciesPasswordUpdate(ctx, policyUuid).PasswordPolicyRequest(passwordPolicyRequest).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Password Policy.
	passwordPolicyRequest := *openapiclient.NewPasswordPolicyRequest("Name_example") // PasswordPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesPasswordUpdate(context.Background(), policyUuid).PasswordPolicyRequest(passwordPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesPasswordUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesPasswordUpdate`: PasswordPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesPasswordUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Password Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPasswordUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **passwordPolicyRequest** | [**PasswordPolicyRequest**](PasswordPolicyRequest.md) |  | 

### Return type

[**PasswordPolicy**](PasswordPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPasswordUsedByList

> []UsedBy PoliciesPasswordUsedByList(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Password Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesPasswordUsedByList(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesPasswordUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesPasswordUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesPasswordUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Password Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPasswordUsedByListRequest struct via the builder pattern


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


## PoliciesReputationCreate

> ReputationPolicy PoliciesReputationCreate(ctx).ReputationPolicyRequest(reputationPolicyRequest).Execute()





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
	reputationPolicyRequest := *openapiclient.NewReputationPolicyRequest("Name_example") // ReputationPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesReputationCreate(context.Background()).ReputationPolicyRequest(reputationPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesReputationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesReputationCreate`: ReputationPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesReputationCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reputationPolicyRequest** | [**ReputationPolicyRequest**](ReputationPolicyRequest.md) |  | 

### Return type

[**ReputationPolicy**](ReputationPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesReputationDestroy

> PoliciesReputationDestroy(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Reputation Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoliciesAPI.PoliciesReputationDestroy(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesReputationDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Reputation Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationDestroyRequest struct via the builder pattern


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


## PoliciesReputationList

> PaginatedReputationPolicyList PoliciesReputationList(ctx).CheckIp(checkIp).CheckUsername(checkUsername).Created(created).ExecutionLogging(executionLogging).LastUpdated(lastUpdated).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PolicyUuid(policyUuid).Search(search).Threshold(threshold).Execute()





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
	checkIp := true // bool |  (optional)
	checkUsername := true // bool |  (optional)
	created := time.Now() // time.Time |  (optional)
	executionLogging := true // bool |  (optional)
	lastUpdated := time.Now() // time.Time |  (optional)
	name := "name_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	search := "search_example" // string | A search term. (optional)
	threshold := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesReputationList(context.Background()).CheckIp(checkIp).CheckUsername(checkUsername).Created(created).ExecutionLogging(executionLogging).LastUpdated(lastUpdated).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PolicyUuid(policyUuid).Search(search).Threshold(threshold).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesReputationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesReputationList`: PaginatedReputationPolicyList
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesReputationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkIp** | **bool** |  | 
 **checkUsername** | **bool** |  | 
 **created** | **time.Time** |  | 
 **executionLogging** | **bool** |  | 
 **lastUpdated** | **time.Time** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **policyUuid** | **string** |  | 
 **search** | **string** | A search term. | 
 **threshold** | **int32** |  | 

### Return type

[**PaginatedReputationPolicyList**](PaginatedReputationPolicyList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesReputationPartialUpdate

> ReputationPolicy PoliciesReputationPartialUpdate(ctx, policyUuid).PatchedReputationPolicyRequest(patchedReputationPolicyRequest).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Reputation Policy.
	patchedReputationPolicyRequest := *openapiclient.NewPatchedReputationPolicyRequest() // PatchedReputationPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesReputationPartialUpdate(context.Background(), policyUuid).PatchedReputationPolicyRequest(patchedReputationPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesReputationPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesReputationPartialUpdate`: ReputationPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesReputationPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Reputation Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedReputationPolicyRequest** | [**PatchedReputationPolicyRequest**](PatchedReputationPolicyRequest.md) |  | 

### Return type

[**ReputationPolicy**](ReputationPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesReputationRetrieve

> ReputationPolicy PoliciesReputationRetrieve(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Reputation Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesReputationRetrieve(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesReputationRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesReputationRetrieve`: ReputationPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesReputationRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Reputation Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReputationPolicy**](ReputationPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesReputationScoresDestroy

> PoliciesReputationScoresDestroy(ctx, reputationUuid).Execute()





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
	reputationUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Reputation Score.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoliciesAPI.PoliciesReputationScoresDestroy(context.Background(), reputationUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesReputationScoresDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reputationUuid** | **string** | A UUID string identifying this Reputation Score. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationScoresDestroyRequest struct via the builder pattern


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


## PoliciesReputationScoresList

> PaginatedReputationList PoliciesReputationScoresList(ctx).Identifier(identifier).IdentifierIn(identifierIn).Ip(ip).Ordering(ordering).Page(page).PageSize(pageSize).Score(score).Search(search).Execute()





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
	identifier := "identifier_example" // string |  (optional)
	identifierIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
	ip := "ip_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	score := int32(56) // int32 |  (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesReputationScoresList(context.Background()).Identifier(identifier).IdentifierIn(identifierIn).Ip(ip).Ordering(ordering).Page(page).PageSize(pageSize).Score(score).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesReputationScoresList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesReputationScoresList`: PaginatedReputationList
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesReputationScoresList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationScoresListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identifier** | **string** |  | 
 **identifierIn** | **[]string** | Multiple values may be separated by commas. | 
 **ip** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **score** | **int32** |  | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedReputationList**](PaginatedReputationList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesReputationScoresRetrieve

> Reputation PoliciesReputationScoresRetrieve(ctx, reputationUuid).Execute()





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
	reputationUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Reputation Score.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesReputationScoresRetrieve(context.Background(), reputationUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesReputationScoresRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesReputationScoresRetrieve`: Reputation
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesReputationScoresRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reputationUuid** | **string** | A UUID string identifying this Reputation Score. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationScoresRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Reputation**](Reputation.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesReputationScoresUsedByList

> []UsedBy PoliciesReputationScoresUsedByList(ctx, reputationUuid).Execute()





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
	reputationUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Reputation Score.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesReputationScoresUsedByList(context.Background(), reputationUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesReputationScoresUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesReputationScoresUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesReputationScoresUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reputationUuid** | **string** | A UUID string identifying this Reputation Score. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationScoresUsedByListRequest struct via the builder pattern


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


## PoliciesReputationUpdate

> ReputationPolicy PoliciesReputationUpdate(ctx, policyUuid).ReputationPolicyRequest(reputationPolicyRequest).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Reputation Policy.
	reputationPolicyRequest := *openapiclient.NewReputationPolicyRequest("Name_example") // ReputationPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesReputationUpdate(context.Background(), policyUuid).ReputationPolicyRequest(reputationPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesReputationUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesReputationUpdate`: ReputationPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesReputationUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Reputation Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reputationPolicyRequest** | [**ReputationPolicyRequest**](ReputationPolicyRequest.md) |  | 

### Return type

[**ReputationPolicy**](ReputationPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesReputationUsedByList

> []UsedBy PoliciesReputationUsedByList(ctx, policyUuid).Execute()





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
	policyUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Reputation Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.PoliciesReputationUsedByList(context.Background(), policyUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.PoliciesReputationUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PoliciesReputationUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.PoliciesReputationUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | **string** | A UUID string identifying this Reputation Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationUsedByListRequest struct via the builder pattern


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

