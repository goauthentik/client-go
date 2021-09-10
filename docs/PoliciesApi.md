# \PoliciesApi

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesAllCacheClearCreate**](PoliciesApi.md#PoliciesAllCacheClearCreate) | **Post** /policies/all/cache_clear/ | 
[**PoliciesAllCacheInfoRetrieve**](PoliciesApi.md#PoliciesAllCacheInfoRetrieve) | **Get** /policies/all/cache_info/ | 
[**PoliciesAllDestroy**](PoliciesApi.md#PoliciesAllDestroy) | **Delete** /policies/all/{policy_uuid}/ | 
[**PoliciesAllList**](PoliciesApi.md#PoliciesAllList) | **Get** /policies/all/ | 
[**PoliciesAllRetrieve**](PoliciesApi.md#PoliciesAllRetrieve) | **Get** /policies/all/{policy_uuid}/ | 
[**PoliciesAllTestCreate**](PoliciesApi.md#PoliciesAllTestCreate) | **Post** /policies/all/{policy_uuid}/test/ | 
[**PoliciesAllTypesList**](PoliciesApi.md#PoliciesAllTypesList) | **Get** /policies/all/types/ | 
[**PoliciesAllUsedByList**](PoliciesApi.md#PoliciesAllUsedByList) | **Get** /policies/all/{policy_uuid}/used_by/ | 
[**PoliciesBindingsCreate**](PoliciesApi.md#PoliciesBindingsCreate) | **Post** /policies/bindings/ | 
[**PoliciesBindingsDestroy**](PoliciesApi.md#PoliciesBindingsDestroy) | **Delete** /policies/bindings/{policy_binding_uuid}/ | 
[**PoliciesBindingsList**](PoliciesApi.md#PoliciesBindingsList) | **Get** /policies/bindings/ | 
[**PoliciesBindingsPartialUpdate**](PoliciesApi.md#PoliciesBindingsPartialUpdate) | **Patch** /policies/bindings/{policy_binding_uuid}/ | 
[**PoliciesBindingsRetrieve**](PoliciesApi.md#PoliciesBindingsRetrieve) | **Get** /policies/bindings/{policy_binding_uuid}/ | 
[**PoliciesBindingsUpdate**](PoliciesApi.md#PoliciesBindingsUpdate) | **Put** /policies/bindings/{policy_binding_uuid}/ | 
[**PoliciesBindingsUsedByList**](PoliciesApi.md#PoliciesBindingsUsedByList) | **Get** /policies/bindings/{policy_binding_uuid}/used_by/ | 
[**PoliciesDummyCreate**](PoliciesApi.md#PoliciesDummyCreate) | **Post** /policies/dummy/ | 
[**PoliciesDummyDestroy**](PoliciesApi.md#PoliciesDummyDestroy) | **Delete** /policies/dummy/{policy_uuid}/ | 
[**PoliciesDummyList**](PoliciesApi.md#PoliciesDummyList) | **Get** /policies/dummy/ | 
[**PoliciesDummyPartialUpdate**](PoliciesApi.md#PoliciesDummyPartialUpdate) | **Patch** /policies/dummy/{policy_uuid}/ | 
[**PoliciesDummyRetrieve**](PoliciesApi.md#PoliciesDummyRetrieve) | **Get** /policies/dummy/{policy_uuid}/ | 
[**PoliciesDummyUpdate**](PoliciesApi.md#PoliciesDummyUpdate) | **Put** /policies/dummy/{policy_uuid}/ | 
[**PoliciesDummyUsedByList**](PoliciesApi.md#PoliciesDummyUsedByList) | **Get** /policies/dummy/{policy_uuid}/used_by/ | 
[**PoliciesEventMatcherCreate**](PoliciesApi.md#PoliciesEventMatcherCreate) | **Post** /policies/event_matcher/ | 
[**PoliciesEventMatcherDestroy**](PoliciesApi.md#PoliciesEventMatcherDestroy) | **Delete** /policies/event_matcher/{policy_uuid}/ | 
[**PoliciesEventMatcherList**](PoliciesApi.md#PoliciesEventMatcherList) | **Get** /policies/event_matcher/ | 
[**PoliciesEventMatcherPartialUpdate**](PoliciesApi.md#PoliciesEventMatcherPartialUpdate) | **Patch** /policies/event_matcher/{policy_uuid}/ | 
[**PoliciesEventMatcherRetrieve**](PoliciesApi.md#PoliciesEventMatcherRetrieve) | **Get** /policies/event_matcher/{policy_uuid}/ | 
[**PoliciesEventMatcherUpdate**](PoliciesApi.md#PoliciesEventMatcherUpdate) | **Put** /policies/event_matcher/{policy_uuid}/ | 
[**PoliciesEventMatcherUsedByList**](PoliciesApi.md#PoliciesEventMatcherUsedByList) | **Get** /policies/event_matcher/{policy_uuid}/used_by/ | 
[**PoliciesExpressionCreate**](PoliciesApi.md#PoliciesExpressionCreate) | **Post** /policies/expression/ | 
[**PoliciesExpressionDestroy**](PoliciesApi.md#PoliciesExpressionDestroy) | **Delete** /policies/expression/{policy_uuid}/ | 
[**PoliciesExpressionList**](PoliciesApi.md#PoliciesExpressionList) | **Get** /policies/expression/ | 
[**PoliciesExpressionPartialUpdate**](PoliciesApi.md#PoliciesExpressionPartialUpdate) | **Patch** /policies/expression/{policy_uuid}/ | 
[**PoliciesExpressionRetrieve**](PoliciesApi.md#PoliciesExpressionRetrieve) | **Get** /policies/expression/{policy_uuid}/ | 
[**PoliciesExpressionUpdate**](PoliciesApi.md#PoliciesExpressionUpdate) | **Put** /policies/expression/{policy_uuid}/ | 
[**PoliciesExpressionUsedByList**](PoliciesApi.md#PoliciesExpressionUsedByList) | **Get** /policies/expression/{policy_uuid}/used_by/ | 
[**PoliciesHaveibeenpwnedCreate**](PoliciesApi.md#PoliciesHaveibeenpwnedCreate) | **Post** /policies/haveibeenpwned/ | 
[**PoliciesHaveibeenpwnedDestroy**](PoliciesApi.md#PoliciesHaveibeenpwnedDestroy) | **Delete** /policies/haveibeenpwned/{policy_uuid}/ | 
[**PoliciesHaveibeenpwnedList**](PoliciesApi.md#PoliciesHaveibeenpwnedList) | **Get** /policies/haveibeenpwned/ | 
[**PoliciesHaveibeenpwnedPartialUpdate**](PoliciesApi.md#PoliciesHaveibeenpwnedPartialUpdate) | **Patch** /policies/haveibeenpwned/{policy_uuid}/ | 
[**PoliciesHaveibeenpwnedRetrieve**](PoliciesApi.md#PoliciesHaveibeenpwnedRetrieve) | **Get** /policies/haveibeenpwned/{policy_uuid}/ | 
[**PoliciesHaveibeenpwnedUpdate**](PoliciesApi.md#PoliciesHaveibeenpwnedUpdate) | **Put** /policies/haveibeenpwned/{policy_uuid}/ | 
[**PoliciesHaveibeenpwnedUsedByList**](PoliciesApi.md#PoliciesHaveibeenpwnedUsedByList) | **Get** /policies/haveibeenpwned/{policy_uuid}/used_by/ | 
[**PoliciesPasswordCreate**](PoliciesApi.md#PoliciesPasswordCreate) | **Post** /policies/password/ | 
[**PoliciesPasswordDestroy**](PoliciesApi.md#PoliciesPasswordDestroy) | **Delete** /policies/password/{policy_uuid}/ | 
[**PoliciesPasswordExpiryCreate**](PoliciesApi.md#PoliciesPasswordExpiryCreate) | **Post** /policies/password_expiry/ | 
[**PoliciesPasswordExpiryDestroy**](PoliciesApi.md#PoliciesPasswordExpiryDestroy) | **Delete** /policies/password_expiry/{policy_uuid}/ | 
[**PoliciesPasswordExpiryList**](PoliciesApi.md#PoliciesPasswordExpiryList) | **Get** /policies/password_expiry/ | 
[**PoliciesPasswordExpiryPartialUpdate**](PoliciesApi.md#PoliciesPasswordExpiryPartialUpdate) | **Patch** /policies/password_expiry/{policy_uuid}/ | 
[**PoliciesPasswordExpiryRetrieve**](PoliciesApi.md#PoliciesPasswordExpiryRetrieve) | **Get** /policies/password_expiry/{policy_uuid}/ | 
[**PoliciesPasswordExpiryUpdate**](PoliciesApi.md#PoliciesPasswordExpiryUpdate) | **Put** /policies/password_expiry/{policy_uuid}/ | 
[**PoliciesPasswordExpiryUsedByList**](PoliciesApi.md#PoliciesPasswordExpiryUsedByList) | **Get** /policies/password_expiry/{policy_uuid}/used_by/ | 
[**PoliciesPasswordList**](PoliciesApi.md#PoliciesPasswordList) | **Get** /policies/password/ | 
[**PoliciesPasswordPartialUpdate**](PoliciesApi.md#PoliciesPasswordPartialUpdate) | **Patch** /policies/password/{policy_uuid}/ | 
[**PoliciesPasswordRetrieve**](PoliciesApi.md#PoliciesPasswordRetrieve) | **Get** /policies/password/{policy_uuid}/ | 
[**PoliciesPasswordUpdate**](PoliciesApi.md#PoliciesPasswordUpdate) | **Put** /policies/password/{policy_uuid}/ | 
[**PoliciesPasswordUsedByList**](PoliciesApi.md#PoliciesPasswordUsedByList) | **Get** /policies/password/{policy_uuid}/used_by/ | 
[**PoliciesReputationCreate**](PoliciesApi.md#PoliciesReputationCreate) | **Post** /policies/reputation/ | 
[**PoliciesReputationDestroy**](PoliciesApi.md#PoliciesReputationDestroy) | **Delete** /policies/reputation/{policy_uuid}/ | 
[**PoliciesReputationIpsDestroy**](PoliciesApi.md#PoliciesReputationIpsDestroy) | **Delete** /policies/reputation/ips/{id}/ | 
[**PoliciesReputationIpsList**](PoliciesApi.md#PoliciesReputationIpsList) | **Get** /policies/reputation/ips/ | 
[**PoliciesReputationIpsRetrieve**](PoliciesApi.md#PoliciesReputationIpsRetrieve) | **Get** /policies/reputation/ips/{id}/ | 
[**PoliciesReputationIpsUsedByList**](PoliciesApi.md#PoliciesReputationIpsUsedByList) | **Get** /policies/reputation/ips/{id}/used_by/ | 
[**PoliciesReputationList**](PoliciesApi.md#PoliciesReputationList) | **Get** /policies/reputation/ | 
[**PoliciesReputationPartialUpdate**](PoliciesApi.md#PoliciesReputationPartialUpdate) | **Patch** /policies/reputation/{policy_uuid}/ | 
[**PoliciesReputationRetrieve**](PoliciesApi.md#PoliciesReputationRetrieve) | **Get** /policies/reputation/{policy_uuid}/ | 
[**PoliciesReputationUpdate**](PoliciesApi.md#PoliciesReputationUpdate) | **Put** /policies/reputation/{policy_uuid}/ | 
[**PoliciesReputationUsedByList**](PoliciesApi.md#PoliciesReputationUsedByList) | **Get** /policies/reputation/{policy_uuid}/used_by/ | 
[**PoliciesReputationUsersDestroy**](PoliciesApi.md#PoliciesReputationUsersDestroy) | **Delete** /policies/reputation/users/{id}/ | 
[**PoliciesReputationUsersList**](PoliciesApi.md#PoliciesReputationUsersList) | **Get** /policies/reputation/users/ | 
[**PoliciesReputationUsersRetrieve**](PoliciesApi.md#PoliciesReputationUsersRetrieve) | **Get** /policies/reputation/users/{id}/ | 
[**PoliciesReputationUsersUsedByList**](PoliciesApi.md#PoliciesReputationUsersUsedByList) | **Get** /policies/reputation/users/{id}/used_by/ | 



## PoliciesAllCacheClearCreate

> PoliciesAllCacheClearCreate(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesAllCacheClearCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesAllCacheClearCreate``: %v\n", err)
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
- **Accept**: Not defined

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesAllCacheInfoRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesAllCacheInfoRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesAllCacheInfoRetrieve`: Cache
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesAllCacheInfoRetrieve`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesAllDestroy(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesAllDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Policy. | 

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
- **Accept**: Not defined

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
    openapiclient "./openapi"
)

func main() {
    bindingsIsnull := true // bool |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    promptstageIsnull := true // bool |  (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesAllList(context.Background()).BindingsIsnull(bindingsIsnull).Ordering(ordering).Page(page).PageSize(pageSize).PromptstageIsnull(promptstageIsnull).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesAllList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesAllList`: PaginatedPolicyList
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesAllList`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesAllRetrieve(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesAllRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesAllRetrieve`: Policy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesAllRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Policy.
    policyTestRequest := *openapiclient.NewPolicyTestRequest(int32(123)) // PolicyTestRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesAllTestCreate(context.Background(), policyUuid).PolicyTestRequest(policyTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesAllTestCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesAllTestCreate`: PolicyTestResult
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesAllTestCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Policy. | 

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesAllTypesList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesAllTypesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesAllTypesList`: []TypeCreate
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesAllTypesList`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesAllUsedByList(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesAllUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesAllUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesAllUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyBindingRequest := *openapiclient.NewPolicyBindingRequest("Target_example", int32(123)) // PolicyBindingRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesBindingsCreate(context.Background()).PolicyBindingRequest(policyBindingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesBindingsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesBindingsCreate`: PolicyBinding
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesBindingsCreate`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    policyBindingUuid := TODO // string | A UUID string identifying this Policy Binding.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesBindingsDestroy(context.Background(), policyBindingUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesBindingsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyBindingUuid** | [**string**](.md) | A UUID string identifying this Policy Binding. | 

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesBindingsList

> PaginatedPolicyBindingList PoliciesBindingsList(ctx).Enabled(enabled).Order(order).Ordering(ordering).Page(page).PageSize(pageSize).Policy(policy).Search(search).Target(target).Timeout(timeout).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    enabled := true // bool |  (optional)
    order := int32(56) // int32 |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    policy := TODO // string |  (optional)
    search := "search_example" // string | A search term. (optional)
    target := TODO // string |  (optional)
    timeout := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesBindingsList(context.Background()).Enabled(enabled).Order(order).Ordering(ordering).Page(page).PageSize(pageSize).Policy(policy).Search(search).Target(target).Timeout(timeout).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesBindingsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesBindingsList`: PaginatedPolicyBindingList
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesBindingsList`: %v\n", resp)
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
 **policy** | [**string**](string.md) |  | 
 **search** | **string** | A search term. | 
 **target** | [**string**](string.md) |  | 
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
    openapiclient "./openapi"
)

func main() {
    policyBindingUuid := TODO // string | A UUID string identifying this Policy Binding.
    patchedPolicyBindingRequest := *openapiclient.NewPatchedPolicyBindingRequest() // PatchedPolicyBindingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesBindingsPartialUpdate(context.Background(), policyBindingUuid).PatchedPolicyBindingRequest(patchedPolicyBindingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesBindingsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesBindingsPartialUpdate`: PolicyBinding
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesBindingsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyBindingUuid** | [**string**](.md) | A UUID string identifying this Policy Binding. | 

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
    openapiclient "./openapi"
)

func main() {
    policyBindingUuid := TODO // string | A UUID string identifying this Policy Binding.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesBindingsRetrieve(context.Background(), policyBindingUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesBindingsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesBindingsRetrieve`: PolicyBinding
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesBindingsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyBindingUuid** | [**string**](.md) | A UUID string identifying this Policy Binding. | 

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
    openapiclient "./openapi"
)

func main() {
    policyBindingUuid := TODO // string | A UUID string identifying this Policy Binding.
    policyBindingRequest := *openapiclient.NewPolicyBindingRequest("Target_example", int32(123)) // PolicyBindingRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesBindingsUpdate(context.Background(), policyBindingUuid).PolicyBindingRequest(policyBindingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesBindingsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesBindingsUpdate`: PolicyBinding
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesBindingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyBindingUuid** | [**string**](.md) | A UUID string identifying this Policy Binding. | 

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
    openapiclient "./openapi"
)

func main() {
    policyBindingUuid := TODO // string | A UUID string identifying this Policy Binding.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesBindingsUsedByList(context.Background(), policyBindingUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesBindingsUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesBindingsUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesBindingsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyBindingUuid** | [**string**](.md) | A UUID string identifying this Policy Binding. | 

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
    openapiclient "./openapi"
)

func main() {
    dummyPolicyRequest := *openapiclient.NewDummyPolicyRequest() // DummyPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesDummyCreate(context.Background()).DummyPolicyRequest(dummyPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesDummyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesDummyCreate`: DummyPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesDummyCreate`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Dummy Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesDummyDestroy(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesDummyDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Dummy Policy. | 

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
- **Accept**: Not defined

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
    openapiclient "./openapi"
)

func main() {
    created := time.Now() // time.Time |  (optional)
    executionLogging := true // bool |  (optional)
    lastUpdated := time.Now() // time.Time |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    policyUuid := TODO // string |  (optional)
    result := true // bool |  (optional)
    search := "search_example" // string | A search term. (optional)
    waitMax := int32(56) // int32 |  (optional)
    waitMin := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesDummyList(context.Background()).Created(created).ExecutionLogging(executionLogging).LastUpdated(lastUpdated).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PolicyUuid(policyUuid).Result(result).Search(search).WaitMax(waitMax).WaitMin(waitMin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesDummyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesDummyList`: PaginatedDummyPolicyList
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesDummyList`: %v\n", resp)
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
 **policyUuid** | [**string**](string.md) |  | 
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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Dummy Policy.
    patchedDummyPolicyRequest := *openapiclient.NewPatchedDummyPolicyRequest() // PatchedDummyPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesDummyPartialUpdate(context.Background(), policyUuid).PatchedDummyPolicyRequest(patchedDummyPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesDummyPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesDummyPartialUpdate`: DummyPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesDummyPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Dummy Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Dummy Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesDummyRetrieve(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesDummyRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesDummyRetrieve`: DummyPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesDummyRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Dummy Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Dummy Policy.
    dummyPolicyRequest := *openapiclient.NewDummyPolicyRequest() // DummyPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesDummyUpdate(context.Background(), policyUuid).DummyPolicyRequest(dummyPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesDummyUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesDummyUpdate`: DummyPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesDummyUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Dummy Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Dummy Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesDummyUsedByList(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesDummyUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesDummyUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesDummyUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Dummy Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    eventMatcherPolicyRequest := *openapiclient.NewEventMatcherPolicyRequest() // EventMatcherPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesEventMatcherCreate(context.Background()).EventMatcherPolicyRequest(eventMatcherPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesEventMatcherCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesEventMatcherCreate`: EventMatcherPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesEventMatcherCreate`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Event Matcher Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesEventMatcherDestroy(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesEventMatcherDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Event Matcher Policy. | 

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesEventMatcherList

> PaginatedEventMatcherPolicyList PoliciesEventMatcherList(ctx).Action(action).App(app).ClientIp(clientIp).Created(created).ExecutionLogging(executionLogging).LastUpdated(lastUpdated).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PolicyUuid(policyUuid).Search(search).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    action := "action_example" // string | Match created events with this action type. When left empty, all action types will be matched. (optional)
    app := "app_example" // string | Match events created by selected application. When left empty, all applications are matched. (optional)
    clientIp := "clientIp_example" // string |  (optional)
    created := time.Now() // time.Time |  (optional)
    executionLogging := true // bool |  (optional)
    lastUpdated := time.Now() // time.Time |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    policyUuid := TODO // string |  (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesEventMatcherList(context.Background()).Action(action).App(app).ClientIp(clientIp).Created(created).ExecutionLogging(executionLogging).LastUpdated(lastUpdated).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PolicyUuid(policyUuid).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesEventMatcherList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesEventMatcherList`: PaginatedEventMatcherPolicyList
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesEventMatcherList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesEventMatcherListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Match created events with this action type. When left empty, all action types will be matched. | 
 **app** | **string** | Match events created by selected application. When left empty, all applications are matched. | 
 **clientIp** | **string** |  | 
 **created** | **time.Time** |  | 
 **executionLogging** | **bool** |  | 
 **lastUpdated** | **time.Time** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **policyUuid** | [**string**](string.md) |  | 
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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Event Matcher Policy.
    patchedEventMatcherPolicyRequest := *openapiclient.NewPatchedEventMatcherPolicyRequest() // PatchedEventMatcherPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesEventMatcherPartialUpdate(context.Background(), policyUuid).PatchedEventMatcherPolicyRequest(patchedEventMatcherPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesEventMatcherPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesEventMatcherPartialUpdate`: EventMatcherPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesEventMatcherPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Event Matcher Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Event Matcher Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesEventMatcherRetrieve(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesEventMatcherRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesEventMatcherRetrieve`: EventMatcherPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesEventMatcherRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Event Matcher Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Event Matcher Policy.
    eventMatcherPolicyRequest := *openapiclient.NewEventMatcherPolicyRequest() // EventMatcherPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesEventMatcherUpdate(context.Background(), policyUuid).EventMatcherPolicyRequest(eventMatcherPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesEventMatcherUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesEventMatcherUpdate`: EventMatcherPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesEventMatcherUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Event Matcher Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Event Matcher Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesEventMatcherUsedByList(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesEventMatcherUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesEventMatcherUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesEventMatcherUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Event Matcher Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    expressionPolicyRequest := *openapiclient.NewExpressionPolicyRequest("Expression_example") // ExpressionPolicyRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesExpressionCreate(context.Background()).ExpressionPolicyRequest(expressionPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesExpressionCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesExpressionCreate`: ExpressionPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesExpressionCreate`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Expression Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesExpressionDestroy(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesExpressionDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Expression Policy. | 

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
- **Accept**: Not defined

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
    openapiclient "./openapi"
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
    policyUuid := TODO // string |  (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesExpressionList(context.Background()).Created(created).ExecutionLogging(executionLogging).Expression(expression).LastUpdated(lastUpdated).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PolicyUuid(policyUuid).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesExpressionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesExpressionList`: PaginatedExpressionPolicyList
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesExpressionList`: %v\n", resp)
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
 **policyUuid** | [**string**](string.md) |  | 
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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Expression Policy.
    patchedExpressionPolicyRequest := *openapiclient.NewPatchedExpressionPolicyRequest() // PatchedExpressionPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesExpressionPartialUpdate(context.Background(), policyUuid).PatchedExpressionPolicyRequest(patchedExpressionPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesExpressionPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesExpressionPartialUpdate`: ExpressionPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesExpressionPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Expression Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Expression Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesExpressionRetrieve(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesExpressionRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesExpressionRetrieve`: ExpressionPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesExpressionRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Expression Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Expression Policy.
    expressionPolicyRequest := *openapiclient.NewExpressionPolicyRequest("Expression_example") // ExpressionPolicyRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesExpressionUpdate(context.Background(), policyUuid).ExpressionPolicyRequest(expressionPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesExpressionUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesExpressionUpdate`: ExpressionPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesExpressionUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Expression Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Expression Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesExpressionUsedByList(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesExpressionUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesExpressionUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesExpressionUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Expression Policy. | 

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


## PoliciesHaveibeenpwnedCreate

> HaveIBeenPwendPolicy PoliciesHaveibeenpwnedCreate(ctx).HaveIBeenPwendPolicyRequest(haveIBeenPwendPolicyRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    haveIBeenPwendPolicyRequest := *openapiclient.NewHaveIBeenPwendPolicyRequest() // HaveIBeenPwendPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesHaveibeenpwnedCreate(context.Background()).HaveIBeenPwendPolicyRequest(haveIBeenPwendPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesHaveibeenpwnedCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesHaveibeenpwnedCreate`: HaveIBeenPwendPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesHaveibeenpwnedCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesHaveibeenpwnedCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **haveIBeenPwendPolicyRequest** | [**HaveIBeenPwendPolicyRequest**](HaveIBeenPwendPolicyRequest.md) |  | 

### Return type

[**HaveIBeenPwendPolicy**](HaveIBeenPwendPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesHaveibeenpwnedDestroy

> PoliciesHaveibeenpwnedDestroy(ctx, policyUuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Have I Been Pwned Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesHaveibeenpwnedDestroy(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesHaveibeenpwnedDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Have I Been Pwned Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesHaveibeenpwnedDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesHaveibeenpwnedList

> PaginatedHaveIBeenPwendPolicyList PoliciesHaveibeenpwnedList(ctx).AllowedCount(allowedCount).Created(created).ExecutionLogging(executionLogging).LastUpdated(lastUpdated).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PasswordField(passwordField).PolicyUuid(policyUuid).Search(search).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    allowedCount := int32(56) // int32 |  (optional)
    created := time.Now() // time.Time |  (optional)
    executionLogging := true // bool |  (optional)
    lastUpdated := time.Now() // time.Time |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    passwordField := "passwordField_example" // string |  (optional)
    policyUuid := TODO // string |  (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesHaveibeenpwnedList(context.Background()).AllowedCount(allowedCount).Created(created).ExecutionLogging(executionLogging).LastUpdated(lastUpdated).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PasswordField(passwordField).PolicyUuid(policyUuid).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesHaveibeenpwnedList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesHaveibeenpwnedList`: PaginatedHaveIBeenPwendPolicyList
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesHaveibeenpwnedList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesHaveibeenpwnedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowedCount** | **int32** |  | 
 **created** | **time.Time** |  | 
 **executionLogging** | **bool** |  | 
 **lastUpdated** | **time.Time** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **passwordField** | **string** |  | 
 **policyUuid** | [**string**](string.md) |  | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedHaveIBeenPwendPolicyList**](PaginatedHaveIBeenPwendPolicyList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesHaveibeenpwnedPartialUpdate

> HaveIBeenPwendPolicy PoliciesHaveibeenpwnedPartialUpdate(ctx, policyUuid).PatchedHaveIBeenPwendPolicyRequest(patchedHaveIBeenPwendPolicyRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Have I Been Pwned Policy.
    patchedHaveIBeenPwendPolicyRequest := *openapiclient.NewPatchedHaveIBeenPwendPolicyRequest() // PatchedHaveIBeenPwendPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesHaveibeenpwnedPartialUpdate(context.Background(), policyUuid).PatchedHaveIBeenPwendPolicyRequest(patchedHaveIBeenPwendPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesHaveibeenpwnedPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesHaveibeenpwnedPartialUpdate`: HaveIBeenPwendPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesHaveibeenpwnedPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Have I Been Pwned Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesHaveibeenpwnedPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedHaveIBeenPwendPolicyRequest** | [**PatchedHaveIBeenPwendPolicyRequest**](PatchedHaveIBeenPwendPolicyRequest.md) |  | 

### Return type

[**HaveIBeenPwendPolicy**](HaveIBeenPwendPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesHaveibeenpwnedRetrieve

> HaveIBeenPwendPolicy PoliciesHaveibeenpwnedRetrieve(ctx, policyUuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Have I Been Pwned Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesHaveibeenpwnedRetrieve(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesHaveibeenpwnedRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesHaveibeenpwnedRetrieve`: HaveIBeenPwendPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesHaveibeenpwnedRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Have I Been Pwned Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesHaveibeenpwnedRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HaveIBeenPwendPolicy**](HaveIBeenPwendPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesHaveibeenpwnedUpdate

> HaveIBeenPwendPolicy PoliciesHaveibeenpwnedUpdate(ctx, policyUuid).HaveIBeenPwendPolicyRequest(haveIBeenPwendPolicyRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Have I Been Pwned Policy.
    haveIBeenPwendPolicyRequest := *openapiclient.NewHaveIBeenPwendPolicyRequest() // HaveIBeenPwendPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesHaveibeenpwnedUpdate(context.Background(), policyUuid).HaveIBeenPwendPolicyRequest(haveIBeenPwendPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesHaveibeenpwnedUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesHaveibeenpwnedUpdate`: HaveIBeenPwendPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesHaveibeenpwnedUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Have I Been Pwned Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesHaveibeenpwnedUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **haveIBeenPwendPolicyRequest** | [**HaveIBeenPwendPolicyRequest**](HaveIBeenPwendPolicyRequest.md) |  | 

### Return type

[**HaveIBeenPwendPolicy**](HaveIBeenPwendPolicy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesHaveibeenpwnedUsedByList

> []UsedBy PoliciesHaveibeenpwnedUsedByList(ctx, policyUuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Have I Been Pwned Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesHaveibeenpwnedUsedByList(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesHaveibeenpwnedUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesHaveibeenpwnedUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesHaveibeenpwnedUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Have I Been Pwned Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesHaveibeenpwnedUsedByListRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    passwordPolicyRequest := *openapiclient.NewPasswordPolicyRequest("ErrorMessage_example") // PasswordPolicyRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesPasswordCreate(context.Background()).PasswordPolicyRequest(passwordPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesPasswordCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPasswordCreate`: PasswordPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesPasswordCreate`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Password Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesPasswordDestroy(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesPasswordDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Password Policy. | 

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
- **Accept**: Not defined

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
    openapiclient "./openapi"
)

func main() {
    passwordExpiryPolicyRequest := *openapiclient.NewPasswordExpiryPolicyRequest(int32(123)) // PasswordExpiryPolicyRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesPasswordExpiryCreate(context.Background()).PasswordExpiryPolicyRequest(passwordExpiryPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesPasswordExpiryCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPasswordExpiryCreate`: PasswordExpiryPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesPasswordExpiryCreate`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Password Expiry Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesPasswordExpiryDestroy(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesPasswordExpiryDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Password Expiry Policy. | 

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
- **Accept**: Not defined

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
    openapiclient "./openapi"
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
    policyUuid := TODO // string |  (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesPasswordExpiryList(context.Background()).Created(created).Days(days).DenyOnly(denyOnly).ExecutionLogging(executionLogging).LastUpdated(lastUpdated).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PolicyUuid(policyUuid).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesPasswordExpiryList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPasswordExpiryList`: PaginatedPasswordExpiryPolicyList
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesPasswordExpiryList`: %v\n", resp)
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
 **policyUuid** | [**string**](string.md) |  | 
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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Password Expiry Policy.
    patchedPasswordExpiryPolicyRequest := *openapiclient.NewPatchedPasswordExpiryPolicyRequest() // PatchedPasswordExpiryPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesPasswordExpiryPartialUpdate(context.Background(), policyUuid).PatchedPasswordExpiryPolicyRequest(patchedPasswordExpiryPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesPasswordExpiryPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPasswordExpiryPartialUpdate`: PasswordExpiryPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesPasswordExpiryPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Password Expiry Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Password Expiry Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesPasswordExpiryRetrieve(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesPasswordExpiryRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPasswordExpiryRetrieve`: PasswordExpiryPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesPasswordExpiryRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Password Expiry Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Password Expiry Policy.
    passwordExpiryPolicyRequest := *openapiclient.NewPasswordExpiryPolicyRequest(int32(123)) // PasswordExpiryPolicyRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesPasswordExpiryUpdate(context.Background(), policyUuid).PasswordExpiryPolicyRequest(passwordExpiryPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesPasswordExpiryUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPasswordExpiryUpdate`: PasswordExpiryPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesPasswordExpiryUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Password Expiry Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Password Expiry Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesPasswordExpiryUsedByList(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesPasswordExpiryUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPasswordExpiryUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesPasswordExpiryUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Password Expiry Policy. | 

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

> PaginatedPasswordPolicyList PoliciesPasswordList(ctx).AmountLowercase(amountLowercase).AmountSymbols(amountSymbols).AmountUppercase(amountUppercase).Created(created).ErrorMessage(errorMessage).ExecutionLogging(executionLogging).LastUpdated(lastUpdated).LengthMin(lengthMin).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PasswordField(passwordField).PolicyUuid(policyUuid).Search(search).SymbolCharset(symbolCharset).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    amountLowercase := int32(56) // int32 |  (optional)
    amountSymbols := int32(56) // int32 |  (optional)
    amountUppercase := int32(56) // int32 |  (optional)
    created := time.Now() // time.Time |  (optional)
    errorMessage := "errorMessage_example" // string |  (optional)
    executionLogging := true // bool |  (optional)
    lastUpdated := time.Now() // time.Time |  (optional)
    lengthMin := int32(56) // int32 |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    passwordField := "passwordField_example" // string |  (optional)
    policyUuid := TODO // string |  (optional)
    search := "search_example" // string | A search term. (optional)
    symbolCharset := "symbolCharset_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesPasswordList(context.Background()).AmountLowercase(amountLowercase).AmountSymbols(amountSymbols).AmountUppercase(amountUppercase).Created(created).ErrorMessage(errorMessage).ExecutionLogging(executionLogging).LastUpdated(lastUpdated).LengthMin(lengthMin).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PasswordField(passwordField).PolicyUuid(policyUuid).Search(search).SymbolCharset(symbolCharset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesPasswordList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPasswordList`: PaginatedPasswordPolicyList
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesPasswordList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPasswordListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amountLowercase** | **int32** |  | 
 **amountSymbols** | **int32** |  | 
 **amountUppercase** | **int32** |  | 
 **created** | **time.Time** |  | 
 **errorMessage** | **string** |  | 
 **executionLogging** | **bool** |  | 
 **lastUpdated** | **time.Time** |  | 
 **lengthMin** | **int32** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **passwordField** | **string** |  | 
 **policyUuid** | [**string**](string.md) |  | 
 **search** | **string** | A search term. | 
 **symbolCharset** | **string** |  | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Password Policy.
    patchedPasswordPolicyRequest := *openapiclient.NewPatchedPasswordPolicyRequest() // PatchedPasswordPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesPasswordPartialUpdate(context.Background(), policyUuid).PatchedPasswordPolicyRequest(patchedPasswordPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesPasswordPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPasswordPartialUpdate`: PasswordPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesPasswordPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Password Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Password Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesPasswordRetrieve(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesPasswordRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPasswordRetrieve`: PasswordPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesPasswordRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Password Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Password Policy.
    passwordPolicyRequest := *openapiclient.NewPasswordPolicyRequest("ErrorMessage_example") // PasswordPolicyRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesPasswordUpdate(context.Background(), policyUuid).PasswordPolicyRequest(passwordPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesPasswordUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPasswordUpdate`: PasswordPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesPasswordUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Password Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Password Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesPasswordUsedByList(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesPasswordUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPasswordUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesPasswordUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Password Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    reputationPolicyRequest := *openapiclient.NewReputationPolicyRequest() // ReputationPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesReputationCreate(context.Background()).ReputationPolicyRequest(reputationPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesReputationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesReputationCreate`: ReputationPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesReputationCreate`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Reputation Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesReputationDestroy(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesReputationDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Reputation Policy. | 

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesReputationIpsDestroy

> PoliciesReputationIpsDestroy(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this ip reputation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesReputationIpsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesReputationIpsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this ip reputation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationIpsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesReputationIpsList

> PaginatedIPReputationList PoliciesReputationIpsList(ctx).Ip(ip).Ordering(ordering).Page(page).PageSize(pageSize).Score(score).Search(search).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ip := "ip_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    score := int32(56) // int32 |  (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesReputationIpsList(context.Background()).Ip(ip).Ordering(ordering).Page(page).PageSize(pageSize).Score(score).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesReputationIpsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesReputationIpsList`: PaginatedIPReputationList
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesReputationIpsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationIpsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **score** | **int32** |  | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedIPReputationList**](PaginatedIPReputationList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesReputationIpsRetrieve

> IPReputation PoliciesReputationIpsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this ip reputation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesReputationIpsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesReputationIpsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesReputationIpsRetrieve`: IPReputation
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesReputationIpsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this ip reputation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationIpsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IPReputation**](IPReputation.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesReputationIpsUsedByList

> []UsedBy PoliciesReputationIpsUsedByList(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this ip reputation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesReputationIpsUsedByList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesReputationIpsUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesReputationIpsUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesReputationIpsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this ip reputation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationIpsUsedByListRequest struct via the builder pattern


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
    openapiclient "./openapi"
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
    policyUuid := TODO // string |  (optional)
    search := "search_example" // string | A search term. (optional)
    threshold := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesReputationList(context.Background()).CheckIp(checkIp).CheckUsername(checkUsername).Created(created).ExecutionLogging(executionLogging).LastUpdated(lastUpdated).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PolicyUuid(policyUuid).Search(search).Threshold(threshold).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesReputationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesReputationList`: PaginatedReputationPolicyList
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesReputationList`: %v\n", resp)
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
 **policyUuid** | [**string**](string.md) |  | 
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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Reputation Policy.
    patchedReputationPolicyRequest := *openapiclient.NewPatchedReputationPolicyRequest() // PatchedReputationPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesReputationPartialUpdate(context.Background(), policyUuid).PatchedReputationPolicyRequest(patchedReputationPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesReputationPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesReputationPartialUpdate`: ReputationPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesReputationPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Reputation Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Reputation Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesReputationRetrieve(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesReputationRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesReputationRetrieve`: ReputationPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesReputationRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Reputation Policy. | 

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


## PoliciesReputationUpdate

> ReputationPolicy PoliciesReputationUpdate(ctx, policyUuid).ReputationPolicyRequest(reputationPolicyRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Reputation Policy.
    reputationPolicyRequest := *openapiclient.NewReputationPolicyRequest() // ReputationPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesReputationUpdate(context.Background(), policyUuid).ReputationPolicyRequest(reputationPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesReputationUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesReputationUpdate`: ReputationPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesReputationUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Reputation Policy. | 

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
    openapiclient "./openapi"
)

func main() {
    policyUuid := TODO // string | A UUID string identifying this Reputation Policy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesReputationUsedByList(context.Background(), policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesReputationUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesReputationUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesReputationUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyUuid** | [**string**](.md) | A UUID string identifying this Reputation Policy. | 

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


## PoliciesReputationUsersDestroy

> PoliciesReputationUsersDestroy(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this user reputation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesReputationUsersDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesReputationUsersDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user reputation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationUsersDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesReputationUsersList

> PaginatedUserReputationList PoliciesReputationUsersList(ctx).Ordering(ordering).Page(page).PageSize(pageSize).Score(score).Search(search).Username(username).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    score := int32(56) // int32 |  (optional)
    search := "search_example" // string | A search term. (optional)
    username := "username_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesReputationUsersList(context.Background()).Ordering(ordering).Page(page).PageSize(pageSize).Score(score).Search(search).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesReputationUsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesReputationUsersList`: PaginatedUserReputationList
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesReputationUsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **score** | **int32** |  | 
 **search** | **string** | A search term. | 
 **username** | **string** |  | 

### Return type

[**PaginatedUserReputationList**](PaginatedUserReputationList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesReputationUsersRetrieve

> UserReputation PoliciesReputationUsersRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this user reputation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesReputationUsersRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesReputationUsersRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesReputationUsersRetrieve`: UserReputation
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesReputationUsersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user reputation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationUsersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserReputation**](UserReputation.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesReputationUsersUsedByList

> []UsedBy PoliciesReputationUsersUsedByList(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this user reputation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesApi.PoliciesReputationUsersUsedByList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesApi.PoliciesReputationUsersUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesReputationUsersUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesApi.PoliciesReputationUsersUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user reputation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesReputationUsersUsedByListRequest struct via the builder pattern


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

