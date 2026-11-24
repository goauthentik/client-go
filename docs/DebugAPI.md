# \DebugAPI

All URIs are relative to */api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DebugLogCreate**](DebugAPI.md#DebugLogCreate) | **Post** /debug/log/ | 



## DebugLogCreate

> DebugLogCreate(ctx).ServerLogRequest(serverLogRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serverLogRequest := *openapiclient.NewServerLogRequest("Message_example") // ServerLogRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DebugAPI.DebugLogCreate(context.Background()).ServerLogRequest(serverLogRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DebugAPI.DebugLogCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDebugLogCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverLogRequest** | [**ServerLogRequest**](ServerLogRequest.md) |  | 

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

