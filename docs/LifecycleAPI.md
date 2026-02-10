# \LifecycleAPI

All URIs are relative to */api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LifecycleIterationsCreate**](LifecycleAPI.md#LifecycleIterationsCreate) | **Post** /lifecycle/iterations/ | 
[**LifecycleIterationsLatestRetrieve**](LifecycleAPI.md#LifecycleIterationsLatestRetrieve) | **Get** /lifecycle/iterations/latest/{content_type}/{object_id}/ | 
[**LifecycleIterationsListOpen**](LifecycleAPI.md#LifecycleIterationsListOpen) | **Get** /lifecycle/iterations/open/ | 
[**LifecycleReviewsCreate**](LifecycleAPI.md#LifecycleReviewsCreate) | **Post** /lifecycle/reviews/ | 
[**LifecycleRulesCreate**](LifecycleAPI.md#LifecycleRulesCreate) | **Post** /lifecycle/rules/ | 
[**LifecycleRulesDestroy**](LifecycleAPI.md#LifecycleRulesDestroy) | **Delete** /lifecycle/rules/{id}/ | 
[**LifecycleRulesList**](LifecycleAPI.md#LifecycleRulesList) | **Get** /lifecycle/rules/ | 
[**LifecycleRulesPartialUpdate**](LifecycleAPI.md#LifecycleRulesPartialUpdate) | **Patch** /lifecycle/rules/{id}/ | 
[**LifecycleRulesRetrieve**](LifecycleAPI.md#LifecycleRulesRetrieve) | **Get** /lifecycle/rules/{id}/ | 
[**LifecycleRulesUpdate**](LifecycleAPI.md#LifecycleRulesUpdate) | **Put** /lifecycle/rules/{id}/ | 



## LifecycleIterationsCreate

> LifecycleIteration LifecycleIterationsCreate(ctx).LifecycleIterationRequest(lifecycleIterationRequest).Execute()





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
	lifecycleIterationRequest := *openapiclient.NewLifecycleIterationRequest(openapiclient.ContentTypeEnum("authentik_core.application")) // LifecycleIterationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LifecycleAPI.LifecycleIterationsCreate(context.Background()).LifecycleIterationRequest(lifecycleIterationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LifecycleAPI.LifecycleIterationsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LifecycleIterationsCreate`: LifecycleIteration
	fmt.Fprintf(os.Stdout, "Response from `LifecycleAPI.LifecycleIterationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLifecycleIterationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lifecycleIterationRequest** | [**LifecycleIterationRequest**](LifecycleIterationRequest.md) |  | 

### Return type

[**LifecycleIteration**](LifecycleIteration.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LifecycleIterationsLatestRetrieve

> LifecycleIteration LifecycleIterationsLatestRetrieve(ctx, contentType, objectId).Execute()





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
	contentType := "contentType_example" // string | 
	objectId := "objectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LifecycleAPI.LifecycleIterationsLatestRetrieve(context.Background(), contentType, objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LifecycleAPI.LifecycleIterationsLatestRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LifecycleIterationsLatestRetrieve`: LifecycleIteration
	fmt.Fprintf(os.Stdout, "Response from `LifecycleAPI.LifecycleIterationsLatestRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentType** | **string** |  | 
**objectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLifecycleIterationsLatestRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LifecycleIteration**](LifecycleIteration.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LifecycleIterationsListOpen

> PaginatedLifecycleIterationList LifecycleIterationsListOpen(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).UserIsReviewer(userIsReviewer).Execute()





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
	userIsReviewer := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LifecycleAPI.LifecycleIterationsListOpen(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).UserIsReviewer(userIsReviewer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LifecycleAPI.LifecycleIterationsListOpen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LifecycleIterationsListOpen`: PaginatedLifecycleIterationList
	fmt.Fprintf(os.Stdout, "Response from `LifecycleAPI.LifecycleIterationsListOpen`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLifecycleIterationsListOpenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **userIsReviewer** | **bool** |  | 

### Return type

[**PaginatedLifecycleIterationList**](PaginatedLifecycleIterationList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LifecycleReviewsCreate

> Review LifecycleReviewsCreate(ctx).ReviewRequest(reviewRequest).Execute()





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
	reviewRequest := *openapiclient.NewReviewRequest("Iteration_example") // ReviewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LifecycleAPI.LifecycleReviewsCreate(context.Background()).ReviewRequest(reviewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LifecycleAPI.LifecycleReviewsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LifecycleReviewsCreate`: Review
	fmt.Fprintf(os.Stdout, "Response from `LifecycleAPI.LifecycleReviewsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLifecycleReviewsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reviewRequest** | [**ReviewRequest**](ReviewRequest.md) |  | 

### Return type

[**Review**](Review.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LifecycleRulesCreate

> LifecycleRule LifecycleRulesCreate(ctx).LifecycleRuleRequest(lifecycleRuleRequest).Execute()



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
	lifecycleRuleRequest := *openapiclient.NewLifecycleRuleRequest("Name_example", openapiclient.ContentTypeEnum("authentik_core.application"), []string{"Reviewers_example"}) // LifecycleRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LifecycleAPI.LifecycleRulesCreate(context.Background()).LifecycleRuleRequest(lifecycleRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LifecycleAPI.LifecycleRulesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LifecycleRulesCreate`: LifecycleRule
	fmt.Fprintf(os.Stdout, "Response from `LifecycleAPI.LifecycleRulesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLifecycleRulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lifecycleRuleRequest** | [**LifecycleRuleRequest**](LifecycleRuleRequest.md) |  | 

### Return type

[**LifecycleRule**](LifecycleRule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LifecycleRulesDestroy

> LifecycleRulesDestroy(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this lifecycle rule.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LifecycleAPI.LifecycleRulesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LifecycleAPI.LifecycleRulesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this lifecycle rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLifecycleRulesDestroyRequest struct via the builder pattern


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


## LifecycleRulesList

> PaginatedLifecycleRuleList LifecycleRulesList(ctx).ContentTypeModel(contentTypeModel).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()



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
	contentTypeModel := "contentTypeModel_example" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LifecycleAPI.LifecycleRulesList(context.Background()).ContentTypeModel(contentTypeModel).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LifecycleAPI.LifecycleRulesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LifecycleRulesList`: PaginatedLifecycleRuleList
	fmt.Fprintf(os.Stdout, "Response from `LifecycleAPI.LifecycleRulesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLifecycleRulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentTypeModel** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedLifecycleRuleList**](PaginatedLifecycleRuleList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LifecycleRulesPartialUpdate

> LifecycleRule LifecycleRulesPartialUpdate(ctx, id).PatchedLifecycleRuleRequest(patchedLifecycleRuleRequest).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this lifecycle rule.
	patchedLifecycleRuleRequest := *openapiclient.NewPatchedLifecycleRuleRequest() // PatchedLifecycleRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LifecycleAPI.LifecycleRulesPartialUpdate(context.Background(), id).PatchedLifecycleRuleRequest(patchedLifecycleRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LifecycleAPI.LifecycleRulesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LifecycleRulesPartialUpdate`: LifecycleRule
	fmt.Fprintf(os.Stdout, "Response from `LifecycleAPI.LifecycleRulesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this lifecycle rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLifecycleRulesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedLifecycleRuleRequest** | [**PatchedLifecycleRuleRequest**](PatchedLifecycleRuleRequest.md) |  | 

### Return type

[**LifecycleRule**](LifecycleRule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LifecycleRulesRetrieve

> LifecycleRule LifecycleRulesRetrieve(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this lifecycle rule.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LifecycleAPI.LifecycleRulesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LifecycleAPI.LifecycleRulesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LifecycleRulesRetrieve`: LifecycleRule
	fmt.Fprintf(os.Stdout, "Response from `LifecycleAPI.LifecycleRulesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this lifecycle rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLifecycleRulesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LifecycleRule**](LifecycleRule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LifecycleRulesUpdate

> LifecycleRule LifecycleRulesUpdate(ctx, id).LifecycleRuleRequest(lifecycleRuleRequest).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this lifecycle rule.
	lifecycleRuleRequest := *openapiclient.NewLifecycleRuleRequest("Name_example", openapiclient.ContentTypeEnum("authentik_core.application"), []string{"Reviewers_example"}) // LifecycleRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LifecycleAPI.LifecycleRulesUpdate(context.Background(), id).LifecycleRuleRequest(lifecycleRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LifecycleAPI.LifecycleRulesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LifecycleRulesUpdate`: LifecycleRule
	fmt.Fprintf(os.Stdout, "Response from `LifecycleAPI.LifecycleRulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this lifecycle rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLifecycleRulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lifecycleRuleRequest** | [**LifecycleRuleRequest**](LifecycleRuleRequest.md) |  | 

### Return type

[**LifecycleRule**](LifecycleRule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

