# \Oauth2Api

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Oauth2AccessTokensDestroy**](Oauth2Api.md#Oauth2AccessTokensDestroy) | **Delete** /oauth2/access_tokens/{id}/ | 
[**Oauth2AccessTokensList**](Oauth2Api.md#Oauth2AccessTokensList) | **Get** /oauth2/access_tokens/ | 
[**Oauth2AccessTokensRetrieve**](Oauth2Api.md#Oauth2AccessTokensRetrieve) | **Get** /oauth2/access_tokens/{id}/ | 
[**Oauth2AccessTokensUsedByList**](Oauth2Api.md#Oauth2AccessTokensUsedByList) | **Get** /oauth2/access_tokens/{id}/used_by/ | 
[**Oauth2AuthorizationCodesDestroy**](Oauth2Api.md#Oauth2AuthorizationCodesDestroy) | **Delete** /oauth2/authorization_codes/{id}/ | 
[**Oauth2AuthorizationCodesList**](Oauth2Api.md#Oauth2AuthorizationCodesList) | **Get** /oauth2/authorization_codes/ | 
[**Oauth2AuthorizationCodesRetrieve**](Oauth2Api.md#Oauth2AuthorizationCodesRetrieve) | **Get** /oauth2/authorization_codes/{id}/ | 
[**Oauth2AuthorizationCodesUsedByList**](Oauth2Api.md#Oauth2AuthorizationCodesUsedByList) | **Get** /oauth2/authorization_codes/{id}/used_by/ | 
[**Oauth2RefreshTokensDestroy**](Oauth2Api.md#Oauth2RefreshTokensDestroy) | **Delete** /oauth2/refresh_tokens/{id}/ | 
[**Oauth2RefreshTokensList**](Oauth2Api.md#Oauth2RefreshTokensList) | **Get** /oauth2/refresh_tokens/ | 
[**Oauth2RefreshTokensRetrieve**](Oauth2Api.md#Oauth2RefreshTokensRetrieve) | **Get** /oauth2/refresh_tokens/{id}/ | 
[**Oauth2RefreshTokensUsedByList**](Oauth2Api.md#Oauth2RefreshTokensUsedByList) | **Get** /oauth2/refresh_tokens/{id}/used_by/ | 



## Oauth2AccessTokensDestroy

> Oauth2AccessTokensDestroy(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this OAuth2 Access Token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.Oauth2Api.Oauth2AccessTokensDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.Oauth2AccessTokensDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this OAuth2 Access Token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2AccessTokensDestroyRequest struct via the builder pattern


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


## Oauth2AccessTokensList

> PaginatedTokenModelList Oauth2AccessTokensList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Search(search).User(user).Execute()





### Example

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
    provider := int32(56) // int32 |  (optional)
    search := "search_example" // string | A search term. (optional)
    user := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Oauth2Api.Oauth2AccessTokensList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Search(search).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.Oauth2AccessTokensList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Oauth2AccessTokensList`: PaginatedTokenModelList
    fmt.Fprintf(os.Stdout, "Response from `Oauth2Api.Oauth2AccessTokensList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauth2AccessTokensListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **provider** | **int32** |  | 
 **search** | **string** | A search term. | 
 **user** | **int32** |  | 

### Return type

[**PaginatedTokenModelList**](PaginatedTokenModelList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Oauth2AccessTokensRetrieve

> TokenModel Oauth2AccessTokensRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this OAuth2 Access Token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Oauth2Api.Oauth2AccessTokensRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.Oauth2AccessTokensRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Oauth2AccessTokensRetrieve`: TokenModel
    fmt.Fprintf(os.Stdout, "Response from `Oauth2Api.Oauth2AccessTokensRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this OAuth2 Access Token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2AccessTokensRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenModel**](TokenModel.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Oauth2AccessTokensUsedByList

> []UsedBy Oauth2AccessTokensUsedByList(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this OAuth2 Access Token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Oauth2Api.Oauth2AccessTokensUsedByList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.Oauth2AccessTokensUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Oauth2AccessTokensUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `Oauth2Api.Oauth2AccessTokensUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this OAuth2 Access Token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2AccessTokensUsedByListRequest struct via the builder pattern


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


## Oauth2AuthorizationCodesDestroy

> Oauth2AuthorizationCodesDestroy(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this Authorization Code.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.Oauth2Api.Oauth2AuthorizationCodesDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.Oauth2AuthorizationCodesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Authorization Code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2AuthorizationCodesDestroyRequest struct via the builder pattern


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


## Oauth2AuthorizationCodesList

> PaginatedExpiringBaseGrantModelList Oauth2AuthorizationCodesList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Search(search).User(user).Execute()





### Example

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
    provider := int32(56) // int32 |  (optional)
    search := "search_example" // string | A search term. (optional)
    user := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Oauth2Api.Oauth2AuthorizationCodesList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Search(search).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.Oauth2AuthorizationCodesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Oauth2AuthorizationCodesList`: PaginatedExpiringBaseGrantModelList
    fmt.Fprintf(os.Stdout, "Response from `Oauth2Api.Oauth2AuthorizationCodesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauth2AuthorizationCodesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **provider** | **int32** |  | 
 **search** | **string** | A search term. | 
 **user** | **int32** |  | 

### Return type

[**PaginatedExpiringBaseGrantModelList**](PaginatedExpiringBaseGrantModelList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Oauth2AuthorizationCodesRetrieve

> ExpiringBaseGrantModel Oauth2AuthorizationCodesRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this Authorization Code.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Oauth2Api.Oauth2AuthorizationCodesRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.Oauth2AuthorizationCodesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Oauth2AuthorizationCodesRetrieve`: ExpiringBaseGrantModel
    fmt.Fprintf(os.Stdout, "Response from `Oauth2Api.Oauth2AuthorizationCodesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Authorization Code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2AuthorizationCodesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExpiringBaseGrantModel**](ExpiringBaseGrantModel.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Oauth2AuthorizationCodesUsedByList

> []UsedBy Oauth2AuthorizationCodesUsedByList(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this Authorization Code.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Oauth2Api.Oauth2AuthorizationCodesUsedByList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.Oauth2AuthorizationCodesUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Oauth2AuthorizationCodesUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `Oauth2Api.Oauth2AuthorizationCodesUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Authorization Code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2AuthorizationCodesUsedByListRequest struct via the builder pattern


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


## Oauth2RefreshTokensDestroy

> Oauth2RefreshTokensDestroy(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this OAuth2 Refresh Token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.Oauth2Api.Oauth2RefreshTokensDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.Oauth2RefreshTokensDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this OAuth2 Refresh Token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2RefreshTokensDestroyRequest struct via the builder pattern


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


## Oauth2RefreshTokensList

> PaginatedTokenModelList Oauth2RefreshTokensList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Search(search).User(user).Execute()





### Example

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
    provider := int32(56) // int32 |  (optional)
    search := "search_example" // string | A search term. (optional)
    user := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Oauth2Api.Oauth2RefreshTokensList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Search(search).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.Oauth2RefreshTokensList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Oauth2RefreshTokensList`: PaginatedTokenModelList
    fmt.Fprintf(os.Stdout, "Response from `Oauth2Api.Oauth2RefreshTokensList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauth2RefreshTokensListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **provider** | **int32** |  | 
 **search** | **string** | A search term. | 
 **user** | **int32** |  | 

### Return type

[**PaginatedTokenModelList**](PaginatedTokenModelList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Oauth2RefreshTokensRetrieve

> TokenModel Oauth2RefreshTokensRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this OAuth2 Refresh Token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Oauth2Api.Oauth2RefreshTokensRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.Oauth2RefreshTokensRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Oauth2RefreshTokensRetrieve`: TokenModel
    fmt.Fprintf(os.Stdout, "Response from `Oauth2Api.Oauth2RefreshTokensRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this OAuth2 Refresh Token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2RefreshTokensRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenModel**](TokenModel.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Oauth2RefreshTokensUsedByList

> []UsedBy Oauth2RefreshTokensUsedByList(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this OAuth2 Refresh Token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Oauth2Api.Oauth2RefreshTokensUsedByList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2Api.Oauth2RefreshTokensUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Oauth2RefreshTokensUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `Oauth2Api.Oauth2RefreshTokensUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this OAuth2 Refresh Token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2RefreshTokensUsedByListRequest struct via the builder pattern


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

