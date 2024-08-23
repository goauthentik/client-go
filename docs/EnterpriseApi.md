# \EnterpriseApi

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnterpriseLicenseCreate**](EnterpriseApi.md#EnterpriseLicenseCreate) | **Post** /enterprise/license/ | 
[**EnterpriseLicenseDestroy**](EnterpriseApi.md#EnterpriseLicenseDestroy) | **Delete** /enterprise/license/{license_uuid}/ | 
[**EnterpriseLicenseForecastRetrieve**](EnterpriseApi.md#EnterpriseLicenseForecastRetrieve) | **Get** /enterprise/license/forecast/ | 
[**EnterpriseLicenseInstallIdRetrieve**](EnterpriseApi.md#EnterpriseLicenseInstallIdRetrieve) | **Get** /enterprise/license/install_id/ | 
[**EnterpriseLicenseList**](EnterpriseApi.md#EnterpriseLicenseList) | **Get** /enterprise/license/ | 
[**EnterpriseLicensePartialUpdate**](EnterpriseApi.md#EnterpriseLicensePartialUpdate) | **Patch** /enterprise/license/{license_uuid}/ | 
[**EnterpriseLicenseRetrieve**](EnterpriseApi.md#EnterpriseLicenseRetrieve) | **Get** /enterprise/license/{license_uuid}/ | 
[**EnterpriseLicenseSummaryRetrieve**](EnterpriseApi.md#EnterpriseLicenseSummaryRetrieve) | **Get** /enterprise/license/summary/ | 
[**EnterpriseLicenseUpdate**](EnterpriseApi.md#EnterpriseLicenseUpdate) | **Put** /enterprise/license/{license_uuid}/ | 
[**EnterpriseLicenseUsedByList**](EnterpriseApi.md#EnterpriseLicenseUsedByList) | **Get** /enterprise/license/{license_uuid}/used_by/ | 



## EnterpriseLicenseCreate

> License EnterpriseLicenseCreate(ctx).LicenseRequest(licenseRequest).Execute()





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
    licenseRequest := *openapiclient.NewLicenseRequest("Key_example") // LicenseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseApi.EnterpriseLicenseCreate(context.Background()).LicenseRequest(licenseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseApi.EnterpriseLicenseCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseLicenseCreate`: License
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseApi.EnterpriseLicenseCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseLicenseCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **licenseRequest** | [**LicenseRequest**](LicenseRequest.md) |  | 

### Return type

[**License**](License.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseLicenseDestroy

> EnterpriseLicenseDestroy(ctx, licenseUuid).Execute()





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
    licenseUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this License.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseApi.EnterpriseLicenseDestroy(context.Background(), licenseUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseApi.EnterpriseLicenseDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licenseUuid** | **string** | A UUID string identifying this License. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseLicenseDestroyRequest struct via the builder pattern


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


## EnterpriseLicenseForecastRetrieve

> LicenseForecast EnterpriseLicenseForecastRetrieve(ctx).Execute()





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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseApi.EnterpriseLicenseForecastRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseApi.EnterpriseLicenseForecastRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseLicenseForecastRetrieve`: LicenseForecast
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseApi.EnterpriseLicenseForecastRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseLicenseForecastRetrieveRequest struct via the builder pattern


### Return type

[**LicenseForecast**](LicenseForecast.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseLicenseInstallIdRetrieve

> InstallID EnterpriseLicenseInstallIdRetrieve(ctx).Execute()





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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseApi.EnterpriseLicenseInstallIdRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseApi.EnterpriseLicenseInstallIdRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseLicenseInstallIdRetrieve`: InstallID
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseApi.EnterpriseLicenseInstallIdRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseLicenseInstallIdRetrieveRequest struct via the builder pattern


### Return type

[**InstallID**](InstallID.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseLicenseList

> PaginatedLicenseList EnterpriseLicenseList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseApi.EnterpriseLicenseList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseApi.EnterpriseLicenseList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseLicenseList`: PaginatedLicenseList
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseApi.EnterpriseLicenseList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseLicenseListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedLicenseList**](PaginatedLicenseList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseLicensePartialUpdate

> License EnterpriseLicensePartialUpdate(ctx, licenseUuid).PatchedLicenseRequest(patchedLicenseRequest).Execute()





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
    licenseUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this License.
    patchedLicenseRequest := *openapiclient.NewPatchedLicenseRequest() // PatchedLicenseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseApi.EnterpriseLicensePartialUpdate(context.Background(), licenseUuid).PatchedLicenseRequest(patchedLicenseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseApi.EnterpriseLicensePartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseLicensePartialUpdate`: License
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseApi.EnterpriseLicensePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licenseUuid** | **string** | A UUID string identifying this License. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseLicensePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedLicenseRequest** | [**PatchedLicenseRequest**](PatchedLicenseRequest.md) |  | 

### Return type

[**License**](License.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseLicenseRetrieve

> License EnterpriseLicenseRetrieve(ctx, licenseUuid).Execute()





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
    licenseUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this License.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseApi.EnterpriseLicenseRetrieve(context.Background(), licenseUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseApi.EnterpriseLicenseRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseLicenseRetrieve`: License
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseApi.EnterpriseLicenseRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licenseUuid** | **string** | A UUID string identifying this License. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseLicenseRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**License**](License.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseLicenseSummaryRetrieve

> LicenseSummary EnterpriseLicenseSummaryRetrieve(ctx).Cached(cached).Execute()





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
    cached := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseApi.EnterpriseLicenseSummaryRetrieve(context.Background()).Cached(cached).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseApi.EnterpriseLicenseSummaryRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseLicenseSummaryRetrieve`: LicenseSummary
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseApi.EnterpriseLicenseSummaryRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseLicenseSummaryRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cached** | **bool** |  | [default to true]

### Return type

[**LicenseSummary**](LicenseSummary.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseLicenseUpdate

> License EnterpriseLicenseUpdate(ctx, licenseUuid).LicenseRequest(licenseRequest).Execute()





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
    licenseUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this License.
    licenseRequest := *openapiclient.NewLicenseRequest("Key_example") // LicenseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseApi.EnterpriseLicenseUpdate(context.Background(), licenseUuid).LicenseRequest(licenseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseApi.EnterpriseLicenseUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseLicenseUpdate`: License
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseApi.EnterpriseLicenseUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licenseUuid** | **string** | A UUID string identifying this License. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseLicenseUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **licenseRequest** | [**LicenseRequest**](LicenseRequest.md) |  | 

### Return type

[**License**](License.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterpriseLicenseUsedByList

> []UsedBy EnterpriseLicenseUsedByList(ctx, licenseUuid).Execute()





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
    licenseUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this License.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnterpriseApi.EnterpriseLicenseUsedByList(context.Background(), licenseUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnterpriseApi.EnterpriseLicenseUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterpriseLicenseUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `EnterpriseApi.EnterpriseLicenseUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licenseUuid** | **string** | A UUID string identifying this License. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterpriseLicenseUsedByListRequest struct via the builder pattern


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

