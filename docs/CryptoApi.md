# \CryptoApi

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CryptoCertificatekeypairsCreate**](CryptoApi.md#CryptoCertificatekeypairsCreate) | **Post** /crypto/certificatekeypairs/ | 
[**CryptoCertificatekeypairsDestroy**](CryptoApi.md#CryptoCertificatekeypairsDestroy) | **Delete** /crypto/certificatekeypairs/{kp_uuid}/ | 
[**CryptoCertificatekeypairsGenerateCreate**](CryptoApi.md#CryptoCertificatekeypairsGenerateCreate) | **Post** /crypto/certificatekeypairs/generate/ | 
[**CryptoCertificatekeypairsList**](CryptoApi.md#CryptoCertificatekeypairsList) | **Get** /crypto/certificatekeypairs/ | 
[**CryptoCertificatekeypairsPartialUpdate**](CryptoApi.md#CryptoCertificatekeypairsPartialUpdate) | **Patch** /crypto/certificatekeypairs/{kp_uuid}/ | 
[**CryptoCertificatekeypairsRetrieve**](CryptoApi.md#CryptoCertificatekeypairsRetrieve) | **Get** /crypto/certificatekeypairs/{kp_uuid}/ | 
[**CryptoCertificatekeypairsUpdate**](CryptoApi.md#CryptoCertificatekeypairsUpdate) | **Put** /crypto/certificatekeypairs/{kp_uuid}/ | 
[**CryptoCertificatekeypairsUsedByList**](CryptoApi.md#CryptoCertificatekeypairsUsedByList) | **Get** /crypto/certificatekeypairs/{kp_uuid}/used_by/ | 
[**CryptoCertificatekeypairsViewCertificateRetrieve**](CryptoApi.md#CryptoCertificatekeypairsViewCertificateRetrieve) | **Get** /crypto/certificatekeypairs/{kp_uuid}/view_certificate/ | 
[**CryptoCertificatekeypairsViewPrivateKeyRetrieve**](CryptoApi.md#CryptoCertificatekeypairsViewPrivateKeyRetrieve) | **Get** /crypto/certificatekeypairs/{kp_uuid}/view_private_key/ | 



## CryptoCertificatekeypairsCreate

> CertificateKeyPair CryptoCertificatekeypairsCreate(ctx).CertificateKeyPairRequest(certificateKeyPairRequest).Execute()





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
    certificateKeyPairRequest := *openapiclient.NewCertificateKeyPairRequest("Name_example", "CertificateData_example") // CertificateKeyPairRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CryptoApi.CryptoCertificatekeypairsCreate(context.Background()).CertificateKeyPairRequest(certificateKeyPairRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CryptoApi.CryptoCertificatekeypairsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CryptoCertificatekeypairsCreate`: CertificateKeyPair
    fmt.Fprintf(os.Stdout, "Response from `CryptoApi.CryptoCertificatekeypairsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCryptoCertificatekeypairsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateKeyPairRequest** | [**CertificateKeyPairRequest**](CertificateKeyPairRequest.md) |  | 

### Return type

[**CertificateKeyPair**](CertificateKeyPair.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CryptoCertificatekeypairsDestroy

> CryptoCertificatekeypairsDestroy(ctx, kpUuid).Execute()





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
    kpUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Certificate-Key Pair.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CryptoApi.CryptoCertificatekeypairsDestroy(context.Background(), kpUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CryptoApi.CryptoCertificatekeypairsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kpUuid** | **string** | A UUID string identifying this Certificate-Key Pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCryptoCertificatekeypairsDestroyRequest struct via the builder pattern


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


## CryptoCertificatekeypairsGenerateCreate

> CertificateKeyPair CryptoCertificatekeypairsGenerateCreate(ctx).CertificateGenerationRequest(certificateGenerationRequest).Execute()





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
    certificateGenerationRequest := *openapiclient.NewCertificateGenerationRequest("CommonName_example", int32(123)) // CertificateGenerationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CryptoApi.CryptoCertificatekeypairsGenerateCreate(context.Background()).CertificateGenerationRequest(certificateGenerationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CryptoApi.CryptoCertificatekeypairsGenerateCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CryptoCertificatekeypairsGenerateCreate`: CertificateKeyPair
    fmt.Fprintf(os.Stdout, "Response from `CryptoApi.CryptoCertificatekeypairsGenerateCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCryptoCertificatekeypairsGenerateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateGenerationRequest** | [**CertificateGenerationRequest**](CertificateGenerationRequest.md) |  | 

### Return type

[**CertificateKeyPair**](CertificateKeyPair.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CryptoCertificatekeypairsList

> PaginatedCertificateKeyPairList CryptoCertificatekeypairsList(ctx).HasKey(hasKey).IncludeDetails(includeDetails).Managed(managed).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    hasKey := true // bool | Only return certificate-key pairs with keys (optional)
    includeDetails := true // bool |  (optional) (default to true)
    managed := "managed_example" // string | managed (optional)
    name := "name_example" // string | name (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CryptoApi.CryptoCertificatekeypairsList(context.Background()).HasKey(hasKey).IncludeDetails(includeDetails).Managed(managed).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CryptoApi.CryptoCertificatekeypairsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CryptoCertificatekeypairsList`: PaginatedCertificateKeyPairList
    fmt.Fprintf(os.Stdout, "Response from `CryptoApi.CryptoCertificatekeypairsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCryptoCertificatekeypairsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hasKey** | **bool** | Only return certificate-key pairs with keys | 
 **includeDetails** | **bool** |  | [default to true]
 **managed** | **string** | managed | 
 **name** | **string** | name | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedCertificateKeyPairList**](PaginatedCertificateKeyPairList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CryptoCertificatekeypairsPartialUpdate

> CertificateKeyPair CryptoCertificatekeypairsPartialUpdate(ctx, kpUuid).PatchedCertificateKeyPairRequest(patchedCertificateKeyPairRequest).Execute()





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
    kpUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Certificate-Key Pair.
    patchedCertificateKeyPairRequest := *openapiclient.NewPatchedCertificateKeyPairRequest() // PatchedCertificateKeyPairRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CryptoApi.CryptoCertificatekeypairsPartialUpdate(context.Background(), kpUuid).PatchedCertificateKeyPairRequest(patchedCertificateKeyPairRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CryptoApi.CryptoCertificatekeypairsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CryptoCertificatekeypairsPartialUpdate`: CertificateKeyPair
    fmt.Fprintf(os.Stdout, "Response from `CryptoApi.CryptoCertificatekeypairsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kpUuid** | **string** | A UUID string identifying this Certificate-Key Pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCryptoCertificatekeypairsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedCertificateKeyPairRequest** | [**PatchedCertificateKeyPairRequest**](PatchedCertificateKeyPairRequest.md) |  | 

### Return type

[**CertificateKeyPair**](CertificateKeyPair.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CryptoCertificatekeypairsRetrieve

> CertificateKeyPair CryptoCertificatekeypairsRetrieve(ctx, kpUuid).Execute()





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
    kpUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Certificate-Key Pair.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CryptoApi.CryptoCertificatekeypairsRetrieve(context.Background(), kpUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CryptoApi.CryptoCertificatekeypairsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CryptoCertificatekeypairsRetrieve`: CertificateKeyPair
    fmt.Fprintf(os.Stdout, "Response from `CryptoApi.CryptoCertificatekeypairsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kpUuid** | **string** | A UUID string identifying this Certificate-Key Pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCryptoCertificatekeypairsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CertificateKeyPair**](CertificateKeyPair.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CryptoCertificatekeypairsUpdate

> CertificateKeyPair CryptoCertificatekeypairsUpdate(ctx, kpUuid).CertificateKeyPairRequest(certificateKeyPairRequest).Execute()





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
    kpUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Certificate-Key Pair.
    certificateKeyPairRequest := *openapiclient.NewCertificateKeyPairRequest("Name_example", "CertificateData_example") // CertificateKeyPairRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CryptoApi.CryptoCertificatekeypairsUpdate(context.Background(), kpUuid).CertificateKeyPairRequest(certificateKeyPairRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CryptoApi.CryptoCertificatekeypairsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CryptoCertificatekeypairsUpdate`: CertificateKeyPair
    fmt.Fprintf(os.Stdout, "Response from `CryptoApi.CryptoCertificatekeypairsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kpUuid** | **string** | A UUID string identifying this Certificate-Key Pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCryptoCertificatekeypairsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificateKeyPairRequest** | [**CertificateKeyPairRequest**](CertificateKeyPairRequest.md) |  | 

### Return type

[**CertificateKeyPair**](CertificateKeyPair.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CryptoCertificatekeypairsUsedByList

> []UsedBy CryptoCertificatekeypairsUsedByList(ctx, kpUuid).Execute()





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
    kpUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Certificate-Key Pair.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CryptoApi.CryptoCertificatekeypairsUsedByList(context.Background(), kpUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CryptoApi.CryptoCertificatekeypairsUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CryptoCertificatekeypairsUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `CryptoApi.CryptoCertificatekeypairsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kpUuid** | **string** | A UUID string identifying this Certificate-Key Pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCryptoCertificatekeypairsUsedByListRequest struct via the builder pattern


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


## CryptoCertificatekeypairsViewCertificateRetrieve

> CertificateData CryptoCertificatekeypairsViewCertificateRetrieve(ctx, kpUuid).Download(download).Execute()





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
    kpUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Certificate-Key Pair.
    download := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CryptoApi.CryptoCertificatekeypairsViewCertificateRetrieve(context.Background(), kpUuid).Download(download).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CryptoApi.CryptoCertificatekeypairsViewCertificateRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CryptoCertificatekeypairsViewCertificateRetrieve`: CertificateData
    fmt.Fprintf(os.Stdout, "Response from `CryptoApi.CryptoCertificatekeypairsViewCertificateRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kpUuid** | **string** | A UUID string identifying this Certificate-Key Pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCryptoCertificatekeypairsViewCertificateRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **download** | **bool** |  | 

### Return type

[**CertificateData**](CertificateData.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CryptoCertificatekeypairsViewPrivateKeyRetrieve

> CertificateData CryptoCertificatekeypairsViewPrivateKeyRetrieve(ctx, kpUuid).Download(download).Execute()





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
    kpUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Certificate-Key Pair.
    download := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CryptoApi.CryptoCertificatekeypairsViewPrivateKeyRetrieve(context.Background(), kpUuid).Download(download).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CryptoApi.CryptoCertificatekeypairsViewPrivateKeyRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CryptoCertificatekeypairsViewPrivateKeyRetrieve`: CertificateData
    fmt.Fprintf(os.Stdout, "Response from `CryptoApi.CryptoCertificatekeypairsViewPrivateKeyRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kpUuid** | **string** | A UUID string identifying this Certificate-Key Pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCryptoCertificatekeypairsViewPrivateKeyRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **download** | **bool** |  | 

### Return type

[**CertificateData**](CertificateData.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

