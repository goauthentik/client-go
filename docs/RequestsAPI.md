# \RequestsAPI

All URIs are relative to */api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestsGrantRequestsAgentCreate**](RequestsAPI.md#RequestsGrantRequestsAgentCreate) | **Post** /requests/grant-requests/agent/ | 
[**RequestsGrantRequestsCreate**](RequestsAPI.md#RequestsGrantRequestsCreate) | **Post** /requests/grant-requests/ | 
[**RequestsGrantRequestsDestroy**](RequestsAPI.md#RequestsGrantRequestsDestroy) | **Delete** /requests/grant-requests/{uuid}/ | 
[**RequestsGrantRequestsFulfillPartialUpdate**](RequestsAPI.md#RequestsGrantRequestsFulfillPartialUpdate) | **Patch** /requests/grant-requests/{uuid}/fulfill/ | 
[**RequestsGrantRequestsList**](RequestsAPI.md#RequestsGrantRequestsList) | **Get** /requests/grant-requests/ | 
[**RequestsGrantRequestsPendingReviewList**](RequestsAPI.md#RequestsGrantRequestsPendingReviewList) | **Get** /requests/grant-requests/pending_review/ | 
[**RequestsGrantRequestsRetrieve**](RequestsAPI.md#RequestsGrantRequestsRetrieve) | **Get** /requests/grant-requests/{uuid}/ | 
[**RequestsGrantRequestsRevokeDestroy**](RequestsAPI.md#RequestsGrantRequestsRevokeDestroy) | **Delete** /requests/grant-requests/{uuid}/revoke/ | 
[**RequestsRuleBindingsCreate**](RequestsAPI.md#RequestsRuleBindingsCreate) | **Post** /requests/rule-bindings/ | 
[**RequestsRuleBindingsDestroy**](RequestsAPI.md#RequestsRuleBindingsDestroy) | **Delete** /requests/rule-bindings/{uuid}/ | 
[**RequestsRuleBindingsList**](RequestsAPI.md#RequestsRuleBindingsList) | **Get** /requests/rule-bindings/ | 
[**RequestsRuleBindingsPartialUpdate**](RequestsAPI.md#RequestsRuleBindingsPartialUpdate) | **Patch** /requests/rule-bindings/{uuid}/ | 
[**RequestsRuleBindingsRetrieve**](RequestsAPI.md#RequestsRuleBindingsRetrieve) | **Get** /requests/rule-bindings/{uuid}/ | 
[**RequestsRuleBindingsUpdate**](RequestsAPI.md#RequestsRuleBindingsUpdate) | **Put** /requests/rule-bindings/{uuid}/ | 
[**RequestsRuleBindingsUsedByList**](RequestsAPI.md#RequestsRuleBindingsUsedByList) | **Get** /requests/rule-bindings/{uuid}/used_by/ | 
[**RequestsRuleChildBindingsCreate**](RequestsAPI.md#RequestsRuleChildBindingsCreate) | **Post** /requests/rule-child-bindings/ | 
[**RequestsRuleChildBindingsDestroy**](RequestsAPI.md#RequestsRuleChildBindingsDestroy) | **Delete** /requests/rule-child-bindings/{uuid}/ | 
[**RequestsRuleChildBindingsList**](RequestsAPI.md#RequestsRuleChildBindingsList) | **Get** /requests/rule-child-bindings/ | 
[**RequestsRuleChildBindingsPartialUpdate**](RequestsAPI.md#RequestsRuleChildBindingsPartialUpdate) | **Patch** /requests/rule-child-bindings/{uuid}/ | 
[**RequestsRuleChildBindingsRetrieve**](RequestsAPI.md#RequestsRuleChildBindingsRetrieve) | **Get** /requests/rule-child-bindings/{uuid}/ | 
[**RequestsRuleChildBindingsUpdate**](RequestsAPI.md#RequestsRuleChildBindingsUpdate) | **Put** /requests/rule-child-bindings/{uuid}/ | 
[**RequestsRuleChildBindingsUsedByList**](RequestsAPI.md#RequestsRuleChildBindingsUsedByList) | **Get** /requests/rule-child-bindings/{uuid}/used_by/ | 
[**RequestsRulesCreate**](RequestsAPI.md#RequestsRulesCreate) | **Post** /requests/rules/ | 
[**RequestsRulesDestroy**](RequestsAPI.md#RequestsRulesDestroy) | **Delete** /requests/rules/{uuid}/ | 
[**RequestsRulesList**](RequestsAPI.md#RequestsRulesList) | **Get** /requests/rules/ | 
[**RequestsRulesPartialUpdate**](RequestsAPI.md#RequestsRulesPartialUpdate) | **Patch** /requests/rules/{uuid}/ | 
[**RequestsRulesRetrieve**](RequestsAPI.md#RequestsRulesRetrieve) | **Get** /requests/rules/{uuid}/ | 
[**RequestsRulesUpdate**](RequestsAPI.md#RequestsRulesUpdate) | **Put** /requests/rules/{uuid}/ | 
[**RequestsRulesUsedByList**](RequestsAPI.md#RequestsRulesUsedByList) | **Get** /requests/rules/{uuid}/used_by/ | 



## RequestsGrantRequestsAgentCreate

> AgentGrantRequestCreated RequestsGrantRequestsAgentCreate(ctx).AgentGrantRequestCreateRequest(agentGrantRequestCreateRequest).Execute()





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
	agentGrantRequestCreateRequest := *openapiclient.NewAgentGrantRequestCreateRequest([]string{"Pbms_example"}) // AgentGrantRequestCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsGrantRequestsAgentCreate(context.Background()).AgentGrantRequestCreateRequest(agentGrantRequestCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsGrantRequestsAgentCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsGrantRequestsAgentCreate`: AgentGrantRequestCreated
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsGrantRequestsAgentCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestsGrantRequestsAgentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentGrantRequestCreateRequest** | [**AgentGrantRequestCreateRequest**](AgentGrantRequestCreateRequest.md) |  | 

### Return type

[**AgentGrantRequestCreated**](AgentGrantRequestCreated.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsGrantRequestsCreate

> Link RequestsGrantRequestsCreate(ctx).GrantRequestCreateRequest(grantRequestCreateRequest).Execute()



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
	grantRequestCreateRequest := *openapiclient.NewGrantRequestCreateRequest([]string{"Pbms_example"}) // GrantRequestCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsGrantRequestsCreate(context.Background()).GrantRequestCreateRequest(grantRequestCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsGrantRequestsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsGrantRequestsCreate`: Link
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsGrantRequestsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestsGrantRequestsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantRequestCreateRequest** | [**GrantRequestCreateRequest**](GrantRequestCreateRequest.md) |  | 

### Return type

[**Link**](Link.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsGrantRequestsDestroy

> RequestsGrantRequestsDestroy(ctx, uuid).Execute()



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Grant Request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RequestsAPI.RequestsGrantRequestsDestroy(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsGrantRequestsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Grant Request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsGrantRequestsDestroyRequest struct via the builder pattern


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


## RequestsGrantRequestsFulfillPartialUpdate

> RequestsGrantRequestsFulfillPartialUpdate(ctx, uuid).PatchedGrantRequestFulfillRequest(patchedGrantRequestFulfillRequest).Execute()



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Grant Request.
	patchedGrantRequestFulfillRequest := *openapiclient.NewPatchedGrantRequestFulfillRequest() // PatchedGrantRequestFulfillRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RequestsAPI.RequestsGrantRequestsFulfillPartialUpdate(context.Background(), uuid).PatchedGrantRequestFulfillRequest(patchedGrantRequestFulfillRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsGrantRequestsFulfillPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Grant Request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsGrantRequestsFulfillPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedGrantRequestFulfillRequest** | [**PatchedGrantRequestFulfillRequest**](PatchedGrantRequestFulfillRequest.md) |  | 

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


## RequestsGrantRequestsList

> PaginatedGrantRequestList RequestsGrantRequestsList(ctx).AgentOwner(agentOwner).CreatedBy(createdBy).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Status(status).Execute()



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
	agentOwner := int32(56) // int32 |  (optional)
	createdBy := int32(56) // int32 |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)
	status := openapiclient.RequestStatus("created") // RequestStatus |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsGrantRequestsList(context.Background()).AgentOwner(agentOwner).CreatedBy(createdBy).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsGrantRequestsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsGrantRequestsList`: PaginatedGrantRequestList
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsGrantRequestsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestsGrantRequestsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentOwner** | **int32** |  | 
 **createdBy** | **int32** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **status** | [**RequestStatus**](RequestStatus.md) |  | 

### Return type

[**PaginatedGrantRequestList**](PaginatedGrantRequestList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsGrantRequestsPendingReviewList

> PaginatedGrantRequestList RequestsGrantRequestsPendingReviewList(ctx).AgentOwner(agentOwner).CreatedBy(createdBy).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Status(status).Execute()





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
	agentOwner := int32(56) // int32 |  (optional)
	createdBy := int32(56) // int32 |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)
	status := openapiclient.RequestStatus("created") // RequestStatus |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsGrantRequestsPendingReviewList(context.Background()).AgentOwner(agentOwner).CreatedBy(createdBy).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsGrantRequestsPendingReviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsGrantRequestsPendingReviewList`: PaginatedGrantRequestList
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsGrantRequestsPendingReviewList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestsGrantRequestsPendingReviewListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentOwner** | **int32** |  | 
 **createdBy** | **int32** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **status** | [**RequestStatus**](RequestStatus.md) |  | 

### Return type

[**PaginatedGrantRequestList**](PaginatedGrantRequestList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsGrantRequestsRetrieve

> GrantRequest RequestsGrantRequestsRetrieve(ctx, uuid).Execute()



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Grant Request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsGrantRequestsRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsGrantRequestsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsGrantRequestsRetrieve`: GrantRequest
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsGrantRequestsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Grant Request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsGrantRequestsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GrantRequest**](GrantRequest.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsGrantRequestsRevokeDestroy

> RequestsGrantRequestsRevokeDestroy(ctx, uuid).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Grant Request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RequestsAPI.RequestsGrantRequestsRevokeDestroy(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsGrantRequestsRevokeDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Grant Request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsGrantRequestsRevokeDestroyRequest struct via the builder pattern


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


## RequestsRuleBindingsCreate

> RequestRuleBinding RequestsRuleBindingsCreate(ctx).RequestRuleBindingRequest(requestRuleBindingRequest).Execute()





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
	requestRuleBindingRequest := *openapiclient.NewRequestRuleBindingRequest("Rule_example", "Target_example") // RequestRuleBindingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRuleBindingsCreate(context.Background()).RequestRuleBindingRequest(requestRuleBindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRuleBindingsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRuleBindingsCreate`: RequestRuleBinding
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRuleBindingsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRuleBindingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestRuleBindingRequest** | [**RequestRuleBindingRequest**](RequestRuleBindingRequest.md) |  | 

### Return type

[**RequestRuleBinding**](RequestRuleBinding.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsRuleBindingsDestroy

> RequestsRuleBindingsDestroy(ctx, uuid).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Request Rule Binding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RequestsAPI.RequestsRuleBindingsDestroy(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRuleBindingsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Request Rule Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRuleBindingsDestroyRequest struct via the builder pattern


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


## RequestsRuleBindingsList

> PaginatedRequestRuleBindingList RequestsRuleBindingsList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Rule(rule).Search(search).Target(target).Execute()





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
	rule := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	search := "search_example" // string | A search term. (optional)
	target := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRuleBindingsList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Rule(rule).Search(search).Target(target).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRuleBindingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRuleBindingsList`: PaginatedRequestRuleBindingList
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRuleBindingsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRuleBindingsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **rule** | **string** |  | 
 **search** | **string** | A search term. | 
 **target** | **string** |  | 

### Return type

[**PaginatedRequestRuleBindingList**](PaginatedRequestRuleBindingList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsRuleBindingsPartialUpdate

> RequestRuleBinding RequestsRuleBindingsPartialUpdate(ctx, uuid).PatchedRequestRuleBindingRequest(patchedRequestRuleBindingRequest).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Request Rule Binding.
	patchedRequestRuleBindingRequest := *openapiclient.NewPatchedRequestRuleBindingRequest() // PatchedRequestRuleBindingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRuleBindingsPartialUpdate(context.Background(), uuid).PatchedRequestRuleBindingRequest(patchedRequestRuleBindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRuleBindingsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRuleBindingsPartialUpdate`: RequestRuleBinding
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRuleBindingsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Request Rule Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRuleBindingsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedRequestRuleBindingRequest** | [**PatchedRequestRuleBindingRequest**](PatchedRequestRuleBindingRequest.md) |  | 

### Return type

[**RequestRuleBinding**](RequestRuleBinding.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsRuleBindingsRetrieve

> RequestRuleBinding RequestsRuleBindingsRetrieve(ctx, uuid).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Request Rule Binding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRuleBindingsRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRuleBindingsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRuleBindingsRetrieve`: RequestRuleBinding
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRuleBindingsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Request Rule Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRuleBindingsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestRuleBinding**](RequestRuleBinding.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsRuleBindingsUpdate

> RequestRuleBinding RequestsRuleBindingsUpdate(ctx, uuid).RequestRuleBindingRequest(requestRuleBindingRequest).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Request Rule Binding.
	requestRuleBindingRequest := *openapiclient.NewRequestRuleBindingRequest("Rule_example", "Target_example") // RequestRuleBindingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRuleBindingsUpdate(context.Background(), uuid).RequestRuleBindingRequest(requestRuleBindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRuleBindingsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRuleBindingsUpdate`: RequestRuleBinding
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRuleBindingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Request Rule Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRuleBindingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestRuleBindingRequest** | [**RequestRuleBindingRequest**](RequestRuleBindingRequest.md) |  | 

### Return type

[**RequestRuleBinding**](RequestRuleBinding.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsRuleBindingsUsedByList

> []UsedBy RequestsRuleBindingsUsedByList(ctx, uuid).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Request Rule Binding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRuleBindingsUsedByList(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRuleBindingsUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRuleBindingsUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRuleBindingsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Request Rule Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRuleBindingsUsedByListRequest struct via the builder pattern


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


## RequestsRuleChildBindingsCreate

> RequestRuleChildBinding RequestsRuleChildBindingsCreate(ctx).RequestRuleChildBindingRequest(requestRuleChildBindingRequest).Execute()





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
	requestRuleChildBindingRequest := *openapiclient.NewRequestRuleChildBindingRequest("Binding_example", "Target_example") // RequestRuleChildBindingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRuleChildBindingsCreate(context.Background()).RequestRuleChildBindingRequest(requestRuleChildBindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRuleChildBindingsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRuleChildBindingsCreate`: RequestRuleChildBinding
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRuleChildBindingsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRuleChildBindingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestRuleChildBindingRequest** | [**RequestRuleChildBindingRequest**](RequestRuleChildBindingRequest.md) |  | 

### Return type

[**RequestRuleChildBinding**](RequestRuleChildBinding.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsRuleChildBindingsDestroy

> RequestsRuleChildBindingsDestroy(ctx, uuid).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Request Rule Child Binding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RequestsAPI.RequestsRuleChildBindingsDestroy(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRuleChildBindingsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Request Rule Child Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRuleChildBindingsDestroyRequest struct via the builder pattern


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


## RequestsRuleChildBindingsList

> PaginatedRequestRuleChildBindingList RequestsRuleChildBindingsList(ctx).Binding(binding).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Target(target).Execute()





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
	binding := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)
	target := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRuleChildBindingsList(context.Background()).Binding(binding).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Target(target).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRuleChildBindingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRuleChildBindingsList`: PaginatedRequestRuleChildBindingList
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRuleChildBindingsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRuleChildBindingsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **binding** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **target** | **string** |  | 

### Return type

[**PaginatedRequestRuleChildBindingList**](PaginatedRequestRuleChildBindingList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsRuleChildBindingsPartialUpdate

> RequestRuleChildBinding RequestsRuleChildBindingsPartialUpdate(ctx, uuid).PatchedRequestRuleChildBindingRequest(patchedRequestRuleChildBindingRequest).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Request Rule Child Binding.
	patchedRequestRuleChildBindingRequest := *openapiclient.NewPatchedRequestRuleChildBindingRequest() // PatchedRequestRuleChildBindingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRuleChildBindingsPartialUpdate(context.Background(), uuid).PatchedRequestRuleChildBindingRequest(patchedRequestRuleChildBindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRuleChildBindingsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRuleChildBindingsPartialUpdate`: RequestRuleChildBinding
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRuleChildBindingsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Request Rule Child Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRuleChildBindingsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedRequestRuleChildBindingRequest** | [**PatchedRequestRuleChildBindingRequest**](PatchedRequestRuleChildBindingRequest.md) |  | 

### Return type

[**RequestRuleChildBinding**](RequestRuleChildBinding.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsRuleChildBindingsRetrieve

> RequestRuleChildBinding RequestsRuleChildBindingsRetrieve(ctx, uuid).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Request Rule Child Binding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRuleChildBindingsRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRuleChildBindingsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRuleChildBindingsRetrieve`: RequestRuleChildBinding
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRuleChildBindingsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Request Rule Child Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRuleChildBindingsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestRuleChildBinding**](RequestRuleChildBinding.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsRuleChildBindingsUpdate

> RequestRuleChildBinding RequestsRuleChildBindingsUpdate(ctx, uuid).RequestRuleChildBindingRequest(requestRuleChildBindingRequest).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Request Rule Child Binding.
	requestRuleChildBindingRequest := *openapiclient.NewRequestRuleChildBindingRequest("Binding_example", "Target_example") // RequestRuleChildBindingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRuleChildBindingsUpdate(context.Background(), uuid).RequestRuleChildBindingRequest(requestRuleChildBindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRuleChildBindingsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRuleChildBindingsUpdate`: RequestRuleChildBinding
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRuleChildBindingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Request Rule Child Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRuleChildBindingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestRuleChildBindingRequest** | [**RequestRuleChildBindingRequest**](RequestRuleChildBindingRequest.md) |  | 

### Return type

[**RequestRuleChildBinding**](RequestRuleChildBinding.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsRuleChildBindingsUsedByList

> []UsedBy RequestsRuleChildBindingsUsedByList(ctx, uuid).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Request Rule Child Binding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRuleChildBindingsUsedByList(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRuleChildBindingsUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRuleChildBindingsUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRuleChildBindingsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Request Rule Child Binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRuleChildBindingsUsedByListRequest struct via the builder pattern


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


## RequestsRulesCreate

> RequestRule RequestsRulesCreate(ctx).RequestRuleRequest(requestRuleRequest).Execute()





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
	requestRuleRequest := *openapiclient.NewRequestRuleRequest("Name_example") // RequestRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRulesCreate(context.Background()).RequestRuleRequest(requestRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRulesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRulesCreate`: RequestRule
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRulesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestRuleRequest** | [**RequestRuleRequest**](RequestRuleRequest.md) |  | 

### Return type

[**RequestRule**](RequestRule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsRulesDestroy

> RequestsRulesDestroy(ctx, uuid).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Request Rule.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RequestsAPI.RequestsRulesDestroy(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRulesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Request Rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRulesDestroyRequest struct via the builder pattern


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


## RequestsRulesList

> PaginatedRequestRuleList RequestsRulesList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PbmUuid(pbmUuid).RequestFlowSlug(requestFlowSlug).Search(search).Execute()





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
	pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	requestFlowSlug := "requestFlowSlug_example" // string |  (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRulesList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PbmUuid(pbmUuid).RequestFlowSlug(requestFlowSlug).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRulesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRulesList`: PaginatedRequestRuleList
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRulesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **pbmUuid** | **string** |  | 
 **requestFlowSlug** | **string** |  | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedRequestRuleList**](PaginatedRequestRuleList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsRulesPartialUpdate

> RequestRule RequestsRulesPartialUpdate(ctx, uuid).PatchedRequestRuleRequest(patchedRequestRuleRequest).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Request Rule.
	patchedRequestRuleRequest := *openapiclient.NewPatchedRequestRuleRequest() // PatchedRequestRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRulesPartialUpdate(context.Background(), uuid).PatchedRequestRuleRequest(patchedRequestRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRulesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRulesPartialUpdate`: RequestRule
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRulesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Request Rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRulesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedRequestRuleRequest** | [**PatchedRequestRuleRequest**](PatchedRequestRuleRequest.md) |  | 

### Return type

[**RequestRule**](RequestRule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsRulesRetrieve

> RequestRule RequestsRulesRetrieve(ctx, uuid).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Request Rule.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRulesRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRulesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRulesRetrieve`: RequestRule
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRulesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Request Rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRulesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestRule**](RequestRule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsRulesUpdate

> RequestRule RequestsRulesUpdate(ctx, uuid).RequestRuleRequest(requestRuleRequest).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Request Rule.
	requestRuleRequest := *openapiclient.NewRequestRuleRequest("Name_example") // RequestRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRulesUpdate(context.Background(), uuid).RequestRuleRequest(requestRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRulesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRulesUpdate`: RequestRule
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Request Rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestRuleRequest** | [**RequestRuleRequest**](RequestRuleRequest.md) |  | 

### Return type

[**RequestRule**](RequestRule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsRulesUsedByList

> []UsedBy RequestsRulesUsedByList(ctx, uuid).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Request Rule.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.RequestsRulesUsedByList(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.RequestsRulesUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestsRulesUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.RequestsRulesUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Request Rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsRulesUsedByListRequest struct via the builder pattern


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

