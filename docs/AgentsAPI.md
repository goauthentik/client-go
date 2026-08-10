# \AgentsAPI

All URIs are relative to */api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentsAgentsCreate**](AgentsAPI.md#AgentsAgentsCreate) | **Post** /agents/agents/ | 
[**AgentsAgentsDestroy**](AgentsAPI.md#AgentsAgentsDestroy) | **Delete** /agents/agents/{id}/ | 
[**AgentsAgentsList**](AgentsAPI.md#AgentsAgentsList) | **Get** /agents/agents/ | 
[**AgentsAgentsPartialUpdate**](AgentsAPI.md#AgentsAgentsPartialUpdate) | **Patch** /agents/agents/{id}/ | 
[**AgentsAgentsRetrieve**](AgentsAPI.md#AgentsAgentsRetrieve) | **Get** /agents/agents/{id}/ | 
[**AgentsAgentsUpdate**](AgentsAPI.md#AgentsAgentsUpdate) | **Put** /agents/agents/{id}/ | 



## AgentsAgentsCreate

> AgentCreated AgentsAgentsCreate(ctx).AgentCreateRequest(agentCreateRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	agentCreateRequest := *openapiclient.NewAgentCreateRequest() // AgentCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.AgentsAgentsCreate(context.Background()).AgentCreateRequest(agentCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.AgentsAgentsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsAgentsCreate`: AgentCreated
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.AgentsAgentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentsAgentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentCreateRequest** | [**AgentCreateRequest**](AgentCreateRequest.md) |  | 

### Return type

[**AgentCreated**](AgentCreated.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsAgentsDestroy

> AgentsAgentsDestroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Agent.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentsAPI.AgentsAgentsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.AgentsAgentsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Agent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsAgentsDestroyRequest struct via the builder pattern


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


## AgentsAgentsList

> PaginatedAgentList AgentsAgentsList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Parent(parent).Search(search).Execute()





### Example

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
	parent := int32(56) // int32 |  (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.AgentsAgentsList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Parent(parent).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.AgentsAgentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsAgentsList`: PaginatedAgentList
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.AgentsAgentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentsAgentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **parent** | **int32** |  | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedAgentList**](PaginatedAgentList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsAgentsPartialUpdate

> Agent AgentsAgentsPartialUpdate(ctx, id).PatchedAgentRequest(patchedAgentRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Agent.
	patchedAgentRequest := *openapiclient.NewPatchedAgentRequest() // PatchedAgentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.AgentsAgentsPartialUpdate(context.Background(), id).PatchedAgentRequest(patchedAgentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.AgentsAgentsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsAgentsPartialUpdate`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.AgentsAgentsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Agent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsAgentsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAgentRequest** | [**PatchedAgentRequest**](PatchedAgentRequest.md) |  | 

### Return type

[**Agent**](Agent.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsAgentsRetrieve

> Agent AgentsAgentsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Agent.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.AgentsAgentsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.AgentsAgentsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsAgentsRetrieve`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.AgentsAgentsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Agent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsAgentsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Agent**](Agent.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentsAgentsUpdate

> Agent AgentsAgentsUpdate(ctx, id).AgentRequest(agentRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this Agent.
	agentRequest := *openapiclient.NewAgentRequest("Username_example", "Name_example") // AgentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.AgentsAgentsUpdate(context.Background(), id).AgentRequest(agentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.AgentsAgentsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentsAgentsUpdate`: Agent
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.AgentsAgentsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Agent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentsAgentsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentRequest** | [**AgentRequest**](AgentRequest.md) |  | 

### Return type

[**Agent**](Agent.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

