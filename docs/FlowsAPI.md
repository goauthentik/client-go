# \FlowsAPI

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowsBindingsCreate**](FlowsAPI.md#FlowsBindingsCreate) | **Post** /flows/bindings/ | 
[**FlowsBindingsDestroy**](FlowsAPI.md#FlowsBindingsDestroy) | **Delete** /flows/bindings/{fsb_uuid}/ | 
[**FlowsBindingsList**](FlowsAPI.md#FlowsBindingsList) | **Get** /flows/bindings/ | 
[**FlowsBindingsPartialUpdate**](FlowsAPI.md#FlowsBindingsPartialUpdate) | **Patch** /flows/bindings/{fsb_uuid}/ | 
[**FlowsBindingsRetrieve**](FlowsAPI.md#FlowsBindingsRetrieve) | **Get** /flows/bindings/{fsb_uuid}/ | 
[**FlowsBindingsUpdate**](FlowsAPI.md#FlowsBindingsUpdate) | **Put** /flows/bindings/{fsb_uuid}/ | 
[**FlowsBindingsUsedByList**](FlowsAPI.md#FlowsBindingsUsedByList) | **Get** /flows/bindings/{fsb_uuid}/used_by/ | 
[**FlowsExecutorGet**](FlowsAPI.md#FlowsExecutorGet) | **Get** /flows/executor/{flow_slug}/ | 
[**FlowsExecutorSolve**](FlowsAPI.md#FlowsExecutorSolve) | **Post** /flows/executor/{flow_slug}/ | 
[**FlowsInspectorGet**](FlowsAPI.md#FlowsInspectorGet) | **Get** /flows/inspector/{flow_slug}/ | 
[**FlowsInstancesCacheClearCreate**](FlowsAPI.md#FlowsInstancesCacheClearCreate) | **Post** /flows/instances/cache_clear/ | 
[**FlowsInstancesCacheInfoRetrieve**](FlowsAPI.md#FlowsInstancesCacheInfoRetrieve) | **Get** /flows/instances/cache_info/ | 
[**FlowsInstancesCreate**](FlowsAPI.md#FlowsInstancesCreate) | **Post** /flows/instances/ | 
[**FlowsInstancesDestroy**](FlowsAPI.md#FlowsInstancesDestroy) | **Delete** /flows/instances/{slug}/ | 
[**FlowsInstancesDiagramRetrieve**](FlowsAPI.md#FlowsInstancesDiagramRetrieve) | **Get** /flows/instances/{slug}/diagram/ | 
[**FlowsInstancesExecuteRetrieve**](FlowsAPI.md#FlowsInstancesExecuteRetrieve) | **Get** /flows/instances/{slug}/execute/ | 
[**FlowsInstancesExportRetrieve**](FlowsAPI.md#FlowsInstancesExportRetrieve) | **Get** /flows/instances/{slug}/export/ | 
[**FlowsInstancesImportCreate**](FlowsAPI.md#FlowsInstancesImportCreate) | **Post** /flows/instances/import/ | 
[**FlowsInstancesList**](FlowsAPI.md#FlowsInstancesList) | **Get** /flows/instances/ | 
[**FlowsInstancesPartialUpdate**](FlowsAPI.md#FlowsInstancesPartialUpdate) | **Patch** /flows/instances/{slug}/ | 
[**FlowsInstancesRetrieve**](FlowsAPI.md#FlowsInstancesRetrieve) | **Get** /flows/instances/{slug}/ | 
[**FlowsInstancesSetBackgroundCreate**](FlowsAPI.md#FlowsInstancesSetBackgroundCreate) | **Post** /flows/instances/{slug}/set_background/ | 
[**FlowsInstancesSetBackgroundUrlCreate**](FlowsAPI.md#FlowsInstancesSetBackgroundUrlCreate) | **Post** /flows/instances/{slug}/set_background_url/ | 
[**FlowsInstancesUpdate**](FlowsAPI.md#FlowsInstancesUpdate) | **Put** /flows/instances/{slug}/ | 
[**FlowsInstancesUsedByList**](FlowsAPI.md#FlowsInstancesUsedByList) | **Get** /flows/instances/{slug}/used_by/ | 



## FlowsBindingsCreate

> FlowStageBinding FlowsBindingsCreate(ctx).FlowStageBindingRequest(flowStageBindingRequest).Execute()





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
	flowStageBindingRequest := *openapiclient.NewFlowStageBindingRequest("Target_example", "Stage_example", int32(123)) // FlowStageBindingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowsBindingsCreate(context.Background()).FlowStageBindingRequest(flowStageBindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsBindingsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsBindingsCreate`: FlowStageBinding
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsBindingsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowsBindingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowStageBindingRequest** | [**FlowStageBindingRequest**](FlowStageBindingRequest.md) |  | 

### Return type

[**FlowStageBinding**](FlowStageBinding.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsBindingsDestroy

> FlowsBindingsDestroy(ctx, fsbUuid).Execute()





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
	fsbUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Flow Stage Binding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowsAPI.FlowsBindingsDestroy(context.Background(), fsbUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsBindingsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsbUuid** | **string** | A UUID string identifying this Flow Stage Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowsBindingsDestroyRequest struct via the builder pattern


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


## FlowsBindingsList

> PaginatedFlowStageBindingList FlowsBindingsList(ctx).EvaluateOnPlan(evaluateOnPlan).FsbUuid(fsbUuid).InvalidResponseAction(invalidResponseAction).Order(order).Ordering(ordering).Page(page).PageSize(pageSize).PbmUuid(pbmUuid).Policies(policies).PolicyEngineMode(policyEngineMode).ReEvaluatePolicies(reEvaluatePolicies).Search(search).Stage(stage).Target(target).Execute()





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
	evaluateOnPlan := true // bool |  (optional)
	fsbUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	invalidResponseAction := "invalidResponseAction_example" // string | Configure how the flow executor should handle an invalid response to a challenge. RETRY returns the error message and a similar challenge to the executor. RESTART restarts the flow from the beginning, and RESTART_WITH_CONTEXT restarts the flow while keeping the current context.   (optional)
	order := int32(56) // int32 |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	policies := []string{"Inner_example"} // []string |  (optional)
	policyEngineMode := "policyEngineMode_example" // string |  (optional)
	reEvaluatePolicies := true // bool |  (optional)
	search := "search_example" // string | A search term. (optional)
	stage := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	target := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowsBindingsList(context.Background()).EvaluateOnPlan(evaluateOnPlan).FsbUuid(fsbUuid).InvalidResponseAction(invalidResponseAction).Order(order).Ordering(ordering).Page(page).PageSize(pageSize).PbmUuid(pbmUuid).Policies(policies).PolicyEngineMode(policyEngineMode).ReEvaluatePolicies(reEvaluatePolicies).Search(search).Stage(stage).Target(target).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsBindingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsBindingsList`: PaginatedFlowStageBindingList
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsBindingsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowsBindingsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evaluateOnPlan** | **bool** |  | 
 **fsbUuid** | **string** |  | 
 **invalidResponseAction** | **string** | Configure how the flow executor should handle an invalid response to a challenge. RETRY returns the error message and a similar challenge to the executor. RESTART restarts the flow from the beginning, and RESTART_WITH_CONTEXT restarts the flow while keeping the current context.   | 
 **order** | **int32** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **pbmUuid** | **string** |  | 
 **policies** | **[]string** |  | 
 **policyEngineMode** | **string** |  | 
 **reEvaluatePolicies** | **bool** |  | 
 **search** | **string** | A search term. | 
 **stage** | **string** |  | 
 **target** | **string** |  | 

### Return type

[**PaginatedFlowStageBindingList**](PaginatedFlowStageBindingList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsBindingsPartialUpdate

> FlowStageBinding FlowsBindingsPartialUpdate(ctx, fsbUuid).PatchedFlowStageBindingRequest(patchedFlowStageBindingRequest).Execute()





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
	fsbUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Flow Stage Binding.
	patchedFlowStageBindingRequest := *openapiclient.NewPatchedFlowStageBindingRequest() // PatchedFlowStageBindingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowsBindingsPartialUpdate(context.Background(), fsbUuid).PatchedFlowStageBindingRequest(patchedFlowStageBindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsBindingsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsBindingsPartialUpdate`: FlowStageBinding
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsBindingsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsbUuid** | **string** | A UUID string identifying this Flow Stage Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowsBindingsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedFlowStageBindingRequest** | [**PatchedFlowStageBindingRequest**](PatchedFlowStageBindingRequest.md) |  | 

### Return type

[**FlowStageBinding**](FlowStageBinding.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsBindingsRetrieve

> FlowStageBinding FlowsBindingsRetrieve(ctx, fsbUuid).Execute()





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
	fsbUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Flow Stage Binding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowsBindingsRetrieve(context.Background(), fsbUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsBindingsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsBindingsRetrieve`: FlowStageBinding
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsBindingsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsbUuid** | **string** | A UUID string identifying this Flow Stage Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowsBindingsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlowStageBinding**](FlowStageBinding.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsBindingsUpdate

> FlowStageBinding FlowsBindingsUpdate(ctx, fsbUuid).FlowStageBindingRequest(flowStageBindingRequest).Execute()





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
	fsbUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Flow Stage Binding.
	flowStageBindingRequest := *openapiclient.NewFlowStageBindingRequest("Target_example", "Stage_example", int32(123)) // FlowStageBindingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowsBindingsUpdate(context.Background(), fsbUuid).FlowStageBindingRequest(flowStageBindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsBindingsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsBindingsUpdate`: FlowStageBinding
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsBindingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsbUuid** | **string** | A UUID string identifying this Flow Stage Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowsBindingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flowStageBindingRequest** | [**FlowStageBindingRequest**](FlowStageBindingRequest.md) |  | 

### Return type

[**FlowStageBinding**](FlowStageBinding.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsBindingsUsedByList

> []UsedBy FlowsBindingsUsedByList(ctx, fsbUuid).Execute()





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
	fsbUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Flow Stage Binding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowsBindingsUsedByList(context.Background(), fsbUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsBindingsUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsBindingsUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsBindingsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsbUuid** | **string** | A UUID string identifying this Flow Stage Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowsBindingsUsedByListRequest struct via the builder pattern


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


## FlowsExecutorGet

> ChallengeTypes FlowsExecutorGet(ctx, flowSlug).Query(query).Execute()





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
	flowSlug := "flowSlug_example" // string | 
	query := "query_example" // string | Querystring as received

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowsExecutorGet(context.Background(), flowSlug).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsExecutorGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsExecutorGet`: ChallengeTypes
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsExecutorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowsExecutorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | Querystring as received | 

### Return type

[**ChallengeTypes**](ChallengeTypes.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsExecutorSolve

> ChallengeTypes FlowsExecutorSolve(ctx, flowSlug).Query(query).FlowChallengeResponseRequest(flowChallengeResponseRequest).Execute()





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
	flowSlug := "flowSlug_example" // string | 
	query := "query_example" // string | Querystring as received
	flowChallengeResponseRequest := openapiclient.FlowChallengeResponseRequest{AppleChallengeResponseRequest: openapiclient.NewAppleChallengeResponseRequest()} // FlowChallengeResponseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowsExecutorSolve(context.Background(), flowSlug).Query(query).FlowChallengeResponseRequest(flowChallengeResponseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsExecutorSolve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsExecutorSolve`: ChallengeTypes
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsExecutorSolve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowsExecutorSolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | Querystring as received | 
 **flowChallengeResponseRequest** | [**FlowChallengeResponseRequest**](FlowChallengeResponseRequest.md) |  | 

### Return type

[**ChallengeTypes**](ChallengeTypes.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsInspectorGet

> FlowInspection FlowsInspectorGet(ctx, flowSlug).Execute()





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
	flowSlug := "flowSlug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowsInspectorGet(context.Background(), flowSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsInspectorGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsInspectorGet`: FlowInspection
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsInspectorGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowsInspectorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlowInspection**](FlowInspection.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsInstancesCacheClearCreate

> FlowsInstancesCacheClearCreate(ctx).Execute()





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
	r, err := apiClient.FlowsAPI.FlowsInstancesCacheClearCreate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsInstancesCacheClearCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowsInstancesCacheClearCreateRequest struct via the builder pattern


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


## FlowsInstancesCacheInfoRetrieve

> Cache FlowsInstancesCacheInfoRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.FlowsAPI.FlowsInstancesCacheInfoRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsInstancesCacheInfoRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsInstancesCacheInfoRetrieve`: Cache
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsInstancesCacheInfoRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowsInstancesCacheInfoRetrieveRequest struct via the builder pattern


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


## FlowsInstancesCreate

> Flow FlowsInstancesCreate(ctx).FlowRequest(flowRequest).Execute()





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
	flowRequest := *openapiclient.NewFlowRequest("Name_example", "Slug_example", "Title_example", openapiclient.FlowDesignationEnum("authentication")) // FlowRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowsInstancesCreate(context.Background()).FlowRequest(flowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsInstancesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsInstancesCreate`: Flow
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsInstancesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowsInstancesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowRequest** | [**FlowRequest**](FlowRequest.md) |  | 

### Return type

[**Flow**](Flow.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsInstancesDestroy

> FlowsInstancesDestroy(ctx, slug).Execute()





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
	r, err := apiClient.FlowsAPI.FlowsInstancesDestroy(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsInstancesDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiFlowsInstancesDestroyRequest struct via the builder pattern


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


## FlowsInstancesDiagramRetrieve

> FlowDiagram FlowsInstancesDiagramRetrieve(ctx, slug).Execute()





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
	resp, r, err := apiClient.FlowsAPI.FlowsInstancesDiagramRetrieve(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsInstancesDiagramRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsInstancesDiagramRetrieve`: FlowDiagram
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsInstancesDiagramRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowsInstancesDiagramRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlowDiagram**](FlowDiagram.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsInstancesExecuteRetrieve

> Link FlowsInstancesExecuteRetrieve(ctx, slug).Execute()





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
	resp, r, err := apiClient.FlowsAPI.FlowsInstancesExecuteRetrieve(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsInstancesExecuteRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsInstancesExecuteRetrieve`: Link
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsInstancesExecuteRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowsInstancesExecuteRetrieveRequest struct via the builder pattern


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


## FlowsInstancesExportRetrieve

> *os.File FlowsInstancesExportRetrieve(ctx, slug).Execute()





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
	resp, r, err := apiClient.FlowsAPI.FlowsInstancesExportRetrieve(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsInstancesExportRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsInstancesExportRetrieve`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsInstancesExportRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowsInstancesExportRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsInstancesImportCreate

> FlowImportResult FlowsInstancesImportCreate(ctx).File(file).Clear(clear).Execute()





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
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)
	clear := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowsInstancesImportCreate(context.Background()).File(file).Clear(clear).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsInstancesImportCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsInstancesImportCreate`: FlowImportResult
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsInstancesImportCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowsInstancesImportCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **clear** | **bool** |  | [default to false]

### Return type

[**FlowImportResult**](FlowImportResult.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsInstancesList

> PaginatedFlowList FlowsInstancesList(ctx).DeniedAction(deniedAction).Designation(designation).FlowUuid(flowUuid).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Slug(slug).Execute()





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
	deniedAction := "deniedAction_example" // string | Configure what should happen when a flow denies access to a user.   (optional)
	designation := "designation_example" // string | Decides what this Flow is used for. For example, the Authentication flow is redirect to when an un-authenticated user visits authentik.   (optional)
	flowUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	name := "name_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)
	slug := "slug_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowsInstancesList(context.Background()).DeniedAction(deniedAction).Designation(designation).FlowUuid(flowUuid).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Slug(slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsInstancesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsInstancesList`: PaginatedFlowList
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsInstancesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowsInstancesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deniedAction** | **string** | Configure what should happen when a flow denies access to a user.   | 
 **designation** | **string** | Decides what this Flow is used for. For example, the Authentication flow is redirect to when an un-authenticated user visits authentik.   | 
 **flowUuid** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **slug** | **string** |  | 

### Return type

[**PaginatedFlowList**](PaginatedFlowList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsInstancesPartialUpdate

> Flow FlowsInstancesPartialUpdate(ctx, slug).PatchedFlowRequest(patchedFlowRequest).Execute()





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
	patchedFlowRequest := *openapiclient.NewPatchedFlowRequest() // PatchedFlowRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowsInstancesPartialUpdate(context.Background(), slug).PatchedFlowRequest(patchedFlowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsInstancesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsInstancesPartialUpdate`: Flow
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsInstancesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowsInstancesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedFlowRequest** | [**PatchedFlowRequest**](PatchedFlowRequest.md) |  | 

### Return type

[**Flow**](Flow.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsInstancesRetrieve

> Flow FlowsInstancesRetrieve(ctx, slug).Execute()





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
	resp, r, err := apiClient.FlowsAPI.FlowsInstancesRetrieve(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsInstancesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsInstancesRetrieve`: Flow
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsInstancesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowsInstancesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Flow**](Flow.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsInstancesSetBackgroundCreate

> FlowsInstancesSetBackgroundCreate(ctx, slug).File(file).Clear(clear).Execute()





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
	r, err := apiClient.FlowsAPI.FlowsInstancesSetBackgroundCreate(context.Background(), slug).File(file).Clear(clear).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsInstancesSetBackgroundCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiFlowsInstancesSetBackgroundCreateRequest struct via the builder pattern


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


## FlowsInstancesSetBackgroundUrlCreate

> FlowsInstancesSetBackgroundUrlCreate(ctx, slug).FilePathRequest(filePathRequest).Execute()





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
	r, err := apiClient.FlowsAPI.FlowsInstancesSetBackgroundUrlCreate(context.Background(), slug).FilePathRequest(filePathRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsInstancesSetBackgroundUrlCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiFlowsInstancesSetBackgroundUrlCreateRequest struct via the builder pattern


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


## FlowsInstancesUpdate

> Flow FlowsInstancesUpdate(ctx, slug).FlowRequest(flowRequest).Execute()





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
	flowRequest := *openapiclient.NewFlowRequest("Name_example", "Slug_example", "Title_example", openapiclient.FlowDesignationEnum("authentication")) // FlowRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowsAPI.FlowsInstancesUpdate(context.Background(), slug).FlowRequest(flowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsInstancesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsInstancesUpdate`: Flow
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsInstancesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowsInstancesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flowRequest** | [**FlowRequest**](FlowRequest.md) |  | 

### Return type

[**Flow**](Flow.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsInstancesUsedByList

> []UsedBy FlowsInstancesUsedByList(ctx, slug).Execute()





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
	resp, r, err := apiClient.FlowsAPI.FlowsInstancesUsedByList(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowsAPI.FlowsInstancesUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowsInstancesUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `FlowsAPI.FlowsInstancesUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowsInstancesUsedByListRequest struct via the builder pattern


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

