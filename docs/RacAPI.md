# \RacAPI

All URIs are relative to */api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RacConnectionTokensDestroy**](RacAPI.md#RacConnectionTokensDestroy) | **Delete** /rac/connection_tokens/{connection_token_uuid}/ | 
[**RacConnectionTokensList**](RacAPI.md#RacConnectionTokensList) | **Get** /rac/connection_tokens/ | 
[**RacConnectionTokensPartialUpdate**](RacAPI.md#RacConnectionTokensPartialUpdate) | **Patch** /rac/connection_tokens/{connection_token_uuid}/ | 
[**RacConnectionTokensRetrieve**](RacAPI.md#RacConnectionTokensRetrieve) | **Get** /rac/connection_tokens/{connection_token_uuid}/ | 
[**RacConnectionTokensUpdate**](RacAPI.md#RacConnectionTokensUpdate) | **Put** /rac/connection_tokens/{connection_token_uuid}/ | 
[**RacConnectionTokensUsedByList**](RacAPI.md#RacConnectionTokensUsedByList) | **Get** /rac/connection_tokens/{connection_token_uuid}/used_by/ | 
[**RacDevicesList**](RacAPI.md#RacDevicesList) | **Get** /rac/devices/ | 



## RacConnectionTokensDestroy

> RacConnectionTokensDestroy(ctx, connectionTokenUuid).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	connectionTokenUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this RAC Connection token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RacAPI.RacConnectionTokensDestroy(context.Background(), connectionTokenUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RacAPI.RacConnectionTokensDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionTokenUuid** | **string** | A UUID string identifying this RAC Connection token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRacConnectionTokensDestroyRequest struct via the builder pattern


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


## RacConnectionTokensList

> PaginatedConnectionTokenList RacConnectionTokensList(ctx).Device(device).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Search(search).SessionUser(sessionUser).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	device := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	provider := int32(56) // int32 |  (optional)
	search := "search_example" // string | A search term. (optional)
	sessionUser := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RacAPI.RacConnectionTokensList(context.Background()).Device(device).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Search(search).SessionUser(sessionUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RacAPI.RacConnectionTokensList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RacConnectionTokensList`: PaginatedConnectionTokenList
	fmt.Fprintf(os.Stdout, "Response from `RacAPI.RacConnectionTokensList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRacConnectionTokensListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **provider** | **int32** |  | 
 **search** | **string** | A search term. | 
 **sessionUser** | **int32** |  | 

### Return type

[**PaginatedConnectionTokenList**](PaginatedConnectionTokenList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RacConnectionTokensPartialUpdate

> ConnectionToken RacConnectionTokensPartialUpdate(ctx, connectionTokenUuid).PatchedConnectionTokenRequest(patchedConnectionTokenRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	connectionTokenUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this RAC Connection token.
	patchedConnectionTokenRequest := *openapiclient.NewPatchedConnectionTokenRequest() // PatchedConnectionTokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RacAPI.RacConnectionTokensPartialUpdate(context.Background(), connectionTokenUuid).PatchedConnectionTokenRequest(patchedConnectionTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RacAPI.RacConnectionTokensPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RacConnectionTokensPartialUpdate`: ConnectionToken
	fmt.Fprintf(os.Stdout, "Response from `RacAPI.RacConnectionTokensPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionTokenUuid** | **string** | A UUID string identifying this RAC Connection token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRacConnectionTokensPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedConnectionTokenRequest** | [**PatchedConnectionTokenRequest**](PatchedConnectionTokenRequest.md) |  | 

### Return type

[**ConnectionToken**](ConnectionToken.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RacConnectionTokensRetrieve

> ConnectionToken RacConnectionTokensRetrieve(ctx, connectionTokenUuid).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	connectionTokenUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this RAC Connection token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RacAPI.RacConnectionTokensRetrieve(context.Background(), connectionTokenUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RacAPI.RacConnectionTokensRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RacConnectionTokensRetrieve`: ConnectionToken
	fmt.Fprintf(os.Stdout, "Response from `RacAPI.RacConnectionTokensRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionTokenUuid** | **string** | A UUID string identifying this RAC Connection token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRacConnectionTokensRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectionToken**](ConnectionToken.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RacConnectionTokensUpdate

> ConnectionToken RacConnectionTokensUpdate(ctx, connectionTokenUuid).ConnectionTokenRequest(connectionTokenRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	connectionTokenUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this RAC Connection token.
	connectionTokenRequest := *openapiclient.NewConnectionTokenRequest(int32(123), "Device_example") // ConnectionTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RacAPI.RacConnectionTokensUpdate(context.Background(), connectionTokenUuid).ConnectionTokenRequest(connectionTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RacAPI.RacConnectionTokensUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RacConnectionTokensUpdate`: ConnectionToken
	fmt.Fprintf(os.Stdout, "Response from `RacAPI.RacConnectionTokensUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionTokenUuid** | **string** | A UUID string identifying this RAC Connection token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRacConnectionTokensUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionTokenRequest** | [**ConnectionTokenRequest**](ConnectionTokenRequest.md) |  | 

### Return type

[**ConnectionToken**](ConnectionToken.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RacConnectionTokensUsedByList

> []UsedBy RacConnectionTokensUsedByList(ctx, connectionTokenUuid).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	connectionTokenUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this RAC Connection token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RacAPI.RacConnectionTokensUsedByList(context.Background(), connectionTokenUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RacAPI.RacConnectionTokensUsedByList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RacConnectionTokensUsedByList`: []UsedBy
	fmt.Fprintf(os.Stdout, "Response from `RacAPI.RacConnectionTokensUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionTokenUuid** | **string** | A UUID string identifying this RAC Connection token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRacConnectionTokensUsedByListRequest struct via the builder pattern


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


## RacDevicesList

> PaginatedRACDeviceList RacDevicesList(ctx).Provider(provider).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	provider := int32(56) // int32 | 
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RacAPI.RacDevicesList(context.Background()).Provider(provider).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RacAPI.RacDevicesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RacDevicesList`: PaginatedRACDeviceList
	fmt.Fprintf(os.Stdout, "Response from `RacAPI.RacDevicesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRacDevicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **int32** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedRACDeviceList**](PaginatedRACDeviceList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

