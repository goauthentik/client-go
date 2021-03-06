# \ProvidersApi

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvidersAllDestroy**](ProvidersApi.md#ProvidersAllDestroy) | **Delete** /providers/all/{id}/ | 
[**ProvidersAllList**](ProvidersApi.md#ProvidersAllList) | **Get** /providers/all/ | 
[**ProvidersAllRetrieve**](ProvidersApi.md#ProvidersAllRetrieve) | **Get** /providers/all/{id}/ | 
[**ProvidersAllTypesList**](ProvidersApi.md#ProvidersAllTypesList) | **Get** /providers/all/types/ | 
[**ProvidersAllUsedByList**](ProvidersApi.md#ProvidersAllUsedByList) | **Get** /providers/all/{id}/used_by/ | 
[**ProvidersLdapCreate**](ProvidersApi.md#ProvidersLdapCreate) | **Post** /providers/ldap/ | 
[**ProvidersLdapDestroy**](ProvidersApi.md#ProvidersLdapDestroy) | **Delete** /providers/ldap/{id}/ | 
[**ProvidersLdapList**](ProvidersApi.md#ProvidersLdapList) | **Get** /providers/ldap/ | 
[**ProvidersLdapPartialUpdate**](ProvidersApi.md#ProvidersLdapPartialUpdate) | **Patch** /providers/ldap/{id}/ | 
[**ProvidersLdapRetrieve**](ProvidersApi.md#ProvidersLdapRetrieve) | **Get** /providers/ldap/{id}/ | 
[**ProvidersLdapUpdate**](ProvidersApi.md#ProvidersLdapUpdate) | **Put** /providers/ldap/{id}/ | 
[**ProvidersLdapUsedByList**](ProvidersApi.md#ProvidersLdapUsedByList) | **Get** /providers/ldap/{id}/used_by/ | 
[**ProvidersOauth2Create**](ProvidersApi.md#ProvidersOauth2Create) | **Post** /providers/oauth2/ | 
[**ProvidersOauth2Destroy**](ProvidersApi.md#ProvidersOauth2Destroy) | **Delete** /providers/oauth2/{id}/ | 
[**ProvidersOauth2List**](ProvidersApi.md#ProvidersOauth2List) | **Get** /providers/oauth2/ | 
[**ProvidersOauth2PartialUpdate**](ProvidersApi.md#ProvidersOauth2PartialUpdate) | **Patch** /providers/oauth2/{id}/ | 
[**ProvidersOauth2Retrieve**](ProvidersApi.md#ProvidersOauth2Retrieve) | **Get** /providers/oauth2/{id}/ | 
[**ProvidersOauth2SetupUrlsRetrieve**](ProvidersApi.md#ProvidersOauth2SetupUrlsRetrieve) | **Get** /providers/oauth2/{id}/setup_urls/ | 
[**ProvidersOauth2Update**](ProvidersApi.md#ProvidersOauth2Update) | **Put** /providers/oauth2/{id}/ | 
[**ProvidersOauth2UsedByList**](ProvidersApi.md#ProvidersOauth2UsedByList) | **Get** /providers/oauth2/{id}/used_by/ | 
[**ProvidersProxyCreate**](ProvidersApi.md#ProvidersProxyCreate) | **Post** /providers/proxy/ | 
[**ProvidersProxyDestroy**](ProvidersApi.md#ProvidersProxyDestroy) | **Delete** /providers/proxy/{id}/ | 
[**ProvidersProxyList**](ProvidersApi.md#ProvidersProxyList) | **Get** /providers/proxy/ | 
[**ProvidersProxyPartialUpdate**](ProvidersApi.md#ProvidersProxyPartialUpdate) | **Patch** /providers/proxy/{id}/ | 
[**ProvidersProxyRetrieve**](ProvidersApi.md#ProvidersProxyRetrieve) | **Get** /providers/proxy/{id}/ | 
[**ProvidersProxyUpdate**](ProvidersApi.md#ProvidersProxyUpdate) | **Put** /providers/proxy/{id}/ | 
[**ProvidersProxyUsedByList**](ProvidersApi.md#ProvidersProxyUsedByList) | **Get** /providers/proxy/{id}/used_by/ | 
[**ProvidersSamlCreate**](ProvidersApi.md#ProvidersSamlCreate) | **Post** /providers/saml/ | 
[**ProvidersSamlDestroy**](ProvidersApi.md#ProvidersSamlDestroy) | **Delete** /providers/saml/{id}/ | 
[**ProvidersSamlImportMetadataCreate**](ProvidersApi.md#ProvidersSamlImportMetadataCreate) | **Post** /providers/saml/import_metadata/ | 
[**ProvidersSamlList**](ProvidersApi.md#ProvidersSamlList) | **Get** /providers/saml/ | 
[**ProvidersSamlMetadataRetrieve**](ProvidersApi.md#ProvidersSamlMetadataRetrieve) | **Get** /providers/saml/{id}/metadata/ | 
[**ProvidersSamlPartialUpdate**](ProvidersApi.md#ProvidersSamlPartialUpdate) | **Patch** /providers/saml/{id}/ | 
[**ProvidersSamlRetrieve**](ProvidersApi.md#ProvidersSamlRetrieve) | **Get** /providers/saml/{id}/ | 
[**ProvidersSamlUpdate**](ProvidersApi.md#ProvidersSamlUpdate) | **Put** /providers/saml/{id}/ | 
[**ProvidersSamlUsedByList**](ProvidersApi.md#ProvidersSamlUsedByList) | **Get** /providers/saml/{id}/used_by/ | 



## ProvidersAllDestroy

> ProvidersAllDestroy(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this provider.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersAllDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersAllDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersAllDestroyRequest struct via the builder pattern


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


## ProvidersAllList

> PaginatedProviderList ProvidersAllList(ctx).ApplicationIsnull(applicationIsnull).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    applicationIsnull := true // bool |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersAllList(context.Background()).ApplicationIsnull(applicationIsnull).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersAllList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersAllList`: PaginatedProviderList
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersAllList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvidersAllListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationIsnull** | **bool** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedProviderList**](PaginatedProviderList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersAllRetrieve

> Provider ProvidersAllRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this provider.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersAllRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersAllRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersAllRetrieve`: Provider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersAllRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersAllRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Provider**](Provider.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersAllTypesList

> []TypeCreate ProvidersAllTypesList(ctx).Execute()





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
    resp, r, err := apiClient.ProvidersApi.ProvidersAllTypesList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersAllTypesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersAllTypesList`: []TypeCreate
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersAllTypesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersAllTypesListRequest struct via the builder pattern


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


## ProvidersAllUsedByList

> []UsedBy ProvidersAllUsedByList(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this provider.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersAllUsedByList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersAllUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersAllUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersAllUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersAllUsedByListRequest struct via the builder pattern


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


## ProvidersLdapCreate

> LDAPProvider ProvidersLdapCreate(ctx).LDAPProviderRequest(lDAPProviderRequest).Execute()





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
    lDAPProviderRequest := *openapiclient.NewLDAPProviderRequest("Name_example", "AuthorizationFlow_example") // LDAPProviderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersLdapCreate(context.Background()).LDAPProviderRequest(lDAPProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersLdapCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersLdapCreate`: LDAPProvider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersLdapCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvidersLdapCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPProviderRequest** | [**LDAPProviderRequest**](LDAPProviderRequest.md) |  | 

### Return type

[**LDAPProvider**](LDAPProvider.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersLdapDestroy

> ProvidersLdapDestroy(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this LDAP Provider.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersLdapDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersLdapDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this LDAP Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersLdapDestroyRequest struct via the builder pattern


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


## ProvidersLdapList

> PaginatedLDAPProviderList ProvidersLdapList(ctx).ApplicationIsnull(applicationIsnull).AuthorizationFlowSlugIexact(authorizationFlowSlugIexact).BaseDnIexact(baseDnIexact).CertificateKpUuidIexact(certificateKpUuidIexact).CertificateNameIexact(certificateNameIexact).GidStartNumberIexact(gidStartNumberIexact).NameIexact(nameIexact).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).SearchGroupGroupUuidIexact(searchGroupGroupUuidIexact).SearchGroupNameIexact(searchGroupNameIexact).TlsServerNameIexact(tlsServerNameIexact).UidStartNumberIexact(uidStartNumberIexact).Execute()





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
    applicationIsnull := true // bool |  (optional)
    authorizationFlowSlugIexact := "authorizationFlowSlugIexact_example" // string |  (optional)
    baseDnIexact := "baseDnIexact_example" // string |  (optional)
    certificateKpUuidIexact := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    certificateNameIexact := "certificateNameIexact_example" // string |  (optional)
    gidStartNumberIexact := int32(56) // int32 |  (optional)
    nameIexact := "nameIexact_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    searchGroupGroupUuidIexact := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    searchGroupNameIexact := "searchGroupNameIexact_example" // string |  (optional)
    tlsServerNameIexact := "tlsServerNameIexact_example" // string |  (optional)
    uidStartNumberIexact := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersLdapList(context.Background()).ApplicationIsnull(applicationIsnull).AuthorizationFlowSlugIexact(authorizationFlowSlugIexact).BaseDnIexact(baseDnIexact).CertificateKpUuidIexact(certificateKpUuidIexact).CertificateNameIexact(certificateNameIexact).GidStartNumberIexact(gidStartNumberIexact).NameIexact(nameIexact).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).SearchGroupGroupUuidIexact(searchGroupGroupUuidIexact).SearchGroupNameIexact(searchGroupNameIexact).TlsServerNameIexact(tlsServerNameIexact).UidStartNumberIexact(uidStartNumberIexact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersLdapList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersLdapList`: PaginatedLDAPProviderList
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersLdapList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvidersLdapListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationIsnull** | **bool** |  | 
 **authorizationFlowSlugIexact** | **string** |  | 
 **baseDnIexact** | **string** |  | 
 **certificateKpUuidIexact** | **string** |  | 
 **certificateNameIexact** | **string** |  | 
 **gidStartNumberIexact** | **int32** |  | 
 **nameIexact** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **searchGroupGroupUuidIexact** | **string** |  | 
 **searchGroupNameIexact** | **string** |  | 
 **tlsServerNameIexact** | **string** |  | 
 **uidStartNumberIexact** | **int32** |  | 

### Return type

[**PaginatedLDAPProviderList**](PaginatedLDAPProviderList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersLdapPartialUpdate

> LDAPProvider ProvidersLdapPartialUpdate(ctx, id).PatchedLDAPProviderRequest(patchedLDAPProviderRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this LDAP Provider.
    patchedLDAPProviderRequest := *openapiclient.NewPatchedLDAPProviderRequest() // PatchedLDAPProviderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersLdapPartialUpdate(context.Background(), id).PatchedLDAPProviderRequest(patchedLDAPProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersLdapPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersLdapPartialUpdate`: LDAPProvider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersLdapPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this LDAP Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersLdapPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedLDAPProviderRequest** | [**PatchedLDAPProviderRequest**](PatchedLDAPProviderRequest.md) |  | 

### Return type

[**LDAPProvider**](LDAPProvider.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersLdapRetrieve

> LDAPProvider ProvidersLdapRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this LDAP Provider.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersLdapRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersLdapRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersLdapRetrieve`: LDAPProvider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersLdapRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this LDAP Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersLdapRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LDAPProvider**](LDAPProvider.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersLdapUpdate

> LDAPProvider ProvidersLdapUpdate(ctx, id).LDAPProviderRequest(lDAPProviderRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this LDAP Provider.
    lDAPProviderRequest := *openapiclient.NewLDAPProviderRequest("Name_example", "AuthorizationFlow_example") // LDAPProviderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersLdapUpdate(context.Background(), id).LDAPProviderRequest(lDAPProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersLdapUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersLdapUpdate`: LDAPProvider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersLdapUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this LDAP Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersLdapUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lDAPProviderRequest** | [**LDAPProviderRequest**](LDAPProviderRequest.md) |  | 

### Return type

[**LDAPProvider**](LDAPProvider.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersLdapUsedByList

> []UsedBy ProvidersLdapUsedByList(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this LDAP Provider.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersLdapUsedByList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersLdapUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersLdapUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersLdapUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this LDAP Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersLdapUsedByListRequest struct via the builder pattern


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


## ProvidersOauth2Create

> OAuth2Provider ProvidersOauth2Create(ctx).OAuth2ProviderRequest(oAuth2ProviderRequest).Execute()





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
    oAuth2ProviderRequest := *openapiclient.NewOAuth2ProviderRequest("Name_example", "AuthorizationFlow_example") // OAuth2ProviderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersOauth2Create(context.Background()).OAuth2ProviderRequest(oAuth2ProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersOauth2Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersOauth2Create`: OAuth2Provider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersOauth2Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvidersOauth2CreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oAuth2ProviderRequest** | [**OAuth2ProviderRequest**](OAuth2ProviderRequest.md) |  | 

### Return type

[**OAuth2Provider**](OAuth2Provider.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersOauth2Destroy

> ProvidersOauth2Destroy(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this OAuth2/OpenID Provider.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersOauth2Destroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersOauth2Destroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this OAuth2/OpenID Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersOauth2DestroyRequest struct via the builder pattern


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


## ProvidersOauth2List

> PaginatedOAuth2ProviderList ProvidersOauth2List(ctx).AccessCodeValidity(accessCodeValidity).Application(application).AuthorizationFlow(authorizationFlow).ClientId(clientId).ClientType(clientType).IncludeClaimsInIdToken(includeClaimsInIdToken).IssuerMode(issuerMode).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PropertyMappings(propertyMappings).RedirectUris(redirectUris).Search(search).SigningKey(signingKey).SubMode(subMode).TokenValidity(tokenValidity).Execute()





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
    accessCodeValidity := "accessCodeValidity_example" // string |  (optional)
    application := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    authorizationFlow := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    clientId := "clientId_example" // string |  (optional)
    clientType := "clientType_example" // string | Confidential clients are capable of maintaining the confidentiality of their credentials. Public clients are incapable (optional)
    includeClaimsInIdToken := true // bool |  (optional)
    issuerMode := "issuerMode_example" // string | Configure how the issuer field of the ID Token should be filled. (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    propertyMappings := []string{"Inner_example"} // []string |  (optional)
    redirectUris := "redirectUris_example" // string |  (optional)
    search := "search_example" // string | A search term. (optional)
    signingKey := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    subMode := "subMode_example" // string | Configure what data should be used as unique User Identifier. For most cases, the default should be fine. (optional)
    tokenValidity := "tokenValidity_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersOauth2List(context.Background()).AccessCodeValidity(accessCodeValidity).Application(application).AuthorizationFlow(authorizationFlow).ClientId(clientId).ClientType(clientType).IncludeClaimsInIdToken(includeClaimsInIdToken).IssuerMode(issuerMode).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PropertyMappings(propertyMappings).RedirectUris(redirectUris).Search(search).SigningKey(signingKey).SubMode(subMode).TokenValidity(tokenValidity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersOauth2List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersOauth2List`: PaginatedOAuth2ProviderList
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersOauth2List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvidersOauth2ListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessCodeValidity** | **string** |  | 
 **application** | **string** |  | 
 **authorizationFlow** | **string** |  | 
 **clientId** | **string** |  | 
 **clientType** | **string** | Confidential clients are capable of maintaining the confidentiality of their credentials. Public clients are incapable | 
 **includeClaimsInIdToken** | **bool** |  | 
 **issuerMode** | **string** | Configure how the issuer field of the ID Token should be filled. | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **propertyMappings** | **[]string** |  | 
 **redirectUris** | **string** |  | 
 **search** | **string** | A search term. | 
 **signingKey** | **string** |  | 
 **subMode** | **string** | Configure what data should be used as unique User Identifier. For most cases, the default should be fine. | 
 **tokenValidity** | **string** |  | 

### Return type

[**PaginatedOAuth2ProviderList**](PaginatedOAuth2ProviderList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersOauth2PartialUpdate

> OAuth2Provider ProvidersOauth2PartialUpdate(ctx, id).PatchedOAuth2ProviderRequest(patchedOAuth2ProviderRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this OAuth2/OpenID Provider.
    patchedOAuth2ProviderRequest := *openapiclient.NewPatchedOAuth2ProviderRequest() // PatchedOAuth2ProviderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersOauth2PartialUpdate(context.Background(), id).PatchedOAuth2ProviderRequest(patchedOAuth2ProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersOauth2PartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersOauth2PartialUpdate`: OAuth2Provider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersOauth2PartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this OAuth2/OpenID Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersOauth2PartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedOAuth2ProviderRequest** | [**PatchedOAuth2ProviderRequest**](PatchedOAuth2ProviderRequest.md) |  | 

### Return type

[**OAuth2Provider**](OAuth2Provider.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersOauth2Retrieve

> OAuth2Provider ProvidersOauth2Retrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this OAuth2/OpenID Provider.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersOauth2Retrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersOauth2Retrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersOauth2Retrieve`: OAuth2Provider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersOauth2Retrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this OAuth2/OpenID Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersOauth2RetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OAuth2Provider**](OAuth2Provider.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersOauth2SetupUrlsRetrieve

> OAuth2ProviderSetupURLs ProvidersOauth2SetupUrlsRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this OAuth2/OpenID Provider.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersOauth2SetupUrlsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersOauth2SetupUrlsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersOauth2SetupUrlsRetrieve`: OAuth2ProviderSetupURLs
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersOauth2SetupUrlsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this OAuth2/OpenID Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersOauth2SetupUrlsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OAuth2ProviderSetupURLs**](OAuth2ProviderSetupURLs.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersOauth2Update

> OAuth2Provider ProvidersOauth2Update(ctx, id).OAuth2ProviderRequest(oAuth2ProviderRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this OAuth2/OpenID Provider.
    oAuth2ProviderRequest := *openapiclient.NewOAuth2ProviderRequest("Name_example", "AuthorizationFlow_example") // OAuth2ProviderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersOauth2Update(context.Background(), id).OAuth2ProviderRequest(oAuth2ProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersOauth2Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersOauth2Update`: OAuth2Provider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersOauth2Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this OAuth2/OpenID Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersOauth2UpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oAuth2ProviderRequest** | [**OAuth2ProviderRequest**](OAuth2ProviderRequest.md) |  | 

### Return type

[**OAuth2Provider**](OAuth2Provider.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersOauth2UsedByList

> []UsedBy ProvidersOauth2UsedByList(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this OAuth2/OpenID Provider.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersOauth2UsedByList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersOauth2UsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersOauth2UsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersOauth2UsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this OAuth2/OpenID Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersOauth2UsedByListRequest struct via the builder pattern


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


## ProvidersProxyCreate

> ProxyProvider ProvidersProxyCreate(ctx).ProxyProviderRequest(proxyProviderRequest).Execute()





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
    proxyProviderRequest := *openapiclient.NewProxyProviderRequest("Name_example", "AuthorizationFlow_example", "ExternalHost_example") // ProxyProviderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersProxyCreate(context.Background()).ProxyProviderRequest(proxyProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersProxyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersProxyCreate`: ProxyProvider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersProxyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvidersProxyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proxyProviderRequest** | [**ProxyProviderRequest**](ProxyProviderRequest.md) |  | 

### Return type

[**ProxyProvider**](ProxyProvider.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersProxyDestroy

> ProvidersProxyDestroy(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Proxy Provider.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersProxyDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersProxyDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Proxy Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersProxyDestroyRequest struct via the builder pattern


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


## ProvidersProxyList

> PaginatedProxyProviderList ProvidersProxyList(ctx).ApplicationIsnull(applicationIsnull).AuthorizationFlowSlugIexact(authorizationFlowSlugIexact).BasicAuthEnabledIexact(basicAuthEnabledIexact).BasicAuthPasswordAttributeIexact(basicAuthPasswordAttributeIexact).BasicAuthUserAttributeIexact(basicAuthUserAttributeIexact).CertificateKpUuidIexact(certificateKpUuidIexact).CertificateNameIexact(certificateNameIexact).CookieDomainIexact(cookieDomainIexact).ExternalHostIexact(externalHostIexact).InternalHostIexact(internalHostIexact).InternalHostSslValidationIexact(internalHostSslValidationIexact).ModeIexact(modeIexact).NameIexact(nameIexact).Ordering(ordering).Page(page).PageSize(pageSize).PropertyMappingsIexact(propertyMappingsIexact).RedirectUrisIexact(redirectUrisIexact).Search(search).SkipPathRegexIexact(skipPathRegexIexact).Execute()





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
    applicationIsnull := true // bool |  (optional)
    authorizationFlowSlugIexact := "authorizationFlowSlugIexact_example" // string |  (optional)
    basicAuthEnabledIexact := true // bool |  (optional)
    basicAuthPasswordAttributeIexact := "basicAuthPasswordAttributeIexact_example" // string |  (optional)
    basicAuthUserAttributeIexact := "basicAuthUserAttributeIexact_example" // string |  (optional)
    certificateKpUuidIexact := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    certificateNameIexact := "certificateNameIexact_example" // string |  (optional)
    cookieDomainIexact := "cookieDomainIexact_example" // string |  (optional)
    externalHostIexact := "externalHostIexact_example" // string |  (optional)
    internalHostIexact := "internalHostIexact_example" // string |  (optional)
    internalHostSslValidationIexact := true // bool |  (optional)
    modeIexact := "modeIexact_example" // string |  (optional)
    nameIexact := "nameIexact_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    propertyMappingsIexact := []string{"Inner_example"} // []string |  (optional)
    redirectUrisIexact := "redirectUrisIexact_example" // string |  (optional)
    search := "search_example" // string | A search term. (optional)
    skipPathRegexIexact := "skipPathRegexIexact_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersProxyList(context.Background()).ApplicationIsnull(applicationIsnull).AuthorizationFlowSlugIexact(authorizationFlowSlugIexact).BasicAuthEnabledIexact(basicAuthEnabledIexact).BasicAuthPasswordAttributeIexact(basicAuthPasswordAttributeIexact).BasicAuthUserAttributeIexact(basicAuthUserAttributeIexact).CertificateKpUuidIexact(certificateKpUuidIexact).CertificateNameIexact(certificateNameIexact).CookieDomainIexact(cookieDomainIexact).ExternalHostIexact(externalHostIexact).InternalHostIexact(internalHostIexact).InternalHostSslValidationIexact(internalHostSslValidationIexact).ModeIexact(modeIexact).NameIexact(nameIexact).Ordering(ordering).Page(page).PageSize(pageSize).PropertyMappingsIexact(propertyMappingsIexact).RedirectUrisIexact(redirectUrisIexact).Search(search).SkipPathRegexIexact(skipPathRegexIexact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersProxyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersProxyList`: PaginatedProxyProviderList
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersProxyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvidersProxyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationIsnull** | **bool** |  | 
 **authorizationFlowSlugIexact** | **string** |  | 
 **basicAuthEnabledIexact** | **bool** |  | 
 **basicAuthPasswordAttributeIexact** | **string** |  | 
 **basicAuthUserAttributeIexact** | **string** |  | 
 **certificateKpUuidIexact** | **string** |  | 
 **certificateNameIexact** | **string** |  | 
 **cookieDomainIexact** | **string** |  | 
 **externalHostIexact** | **string** |  | 
 **internalHostIexact** | **string** |  | 
 **internalHostSslValidationIexact** | **bool** |  | 
 **modeIexact** | **string** |  | 
 **nameIexact** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **propertyMappingsIexact** | **[]string** |  | 
 **redirectUrisIexact** | **string** |  | 
 **search** | **string** | A search term. | 
 **skipPathRegexIexact** | **string** |  | 

### Return type

[**PaginatedProxyProviderList**](PaginatedProxyProviderList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersProxyPartialUpdate

> ProxyProvider ProvidersProxyPartialUpdate(ctx, id).PatchedProxyProviderRequest(patchedProxyProviderRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Proxy Provider.
    patchedProxyProviderRequest := *openapiclient.NewPatchedProxyProviderRequest() // PatchedProxyProviderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersProxyPartialUpdate(context.Background(), id).PatchedProxyProviderRequest(patchedProxyProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersProxyPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersProxyPartialUpdate`: ProxyProvider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersProxyPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Proxy Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersProxyPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedProxyProviderRequest** | [**PatchedProxyProviderRequest**](PatchedProxyProviderRequest.md) |  | 

### Return type

[**ProxyProvider**](ProxyProvider.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersProxyRetrieve

> ProxyProvider ProvidersProxyRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Proxy Provider.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersProxyRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersProxyRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersProxyRetrieve`: ProxyProvider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersProxyRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Proxy Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersProxyRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProxyProvider**](ProxyProvider.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersProxyUpdate

> ProxyProvider ProvidersProxyUpdate(ctx, id).ProxyProviderRequest(proxyProviderRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Proxy Provider.
    proxyProviderRequest := *openapiclient.NewProxyProviderRequest("Name_example", "AuthorizationFlow_example", "ExternalHost_example") // ProxyProviderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersProxyUpdate(context.Background(), id).ProxyProviderRequest(proxyProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersProxyUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersProxyUpdate`: ProxyProvider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersProxyUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Proxy Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersProxyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **proxyProviderRequest** | [**ProxyProviderRequest**](ProxyProviderRequest.md) |  | 

### Return type

[**ProxyProvider**](ProxyProvider.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersProxyUsedByList

> []UsedBy ProvidersProxyUsedByList(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this Proxy Provider.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersProxyUsedByList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersProxyUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersProxyUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersProxyUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Proxy Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersProxyUsedByListRequest struct via the builder pattern


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


## ProvidersSamlCreate

> SAMLProvider ProvidersSamlCreate(ctx).SAMLProviderRequest(sAMLProviderRequest).Execute()





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
    sAMLProviderRequest := *openapiclient.NewSAMLProviderRequest("Name_example", "AuthorizationFlow_example", "AcsUrl_example") // SAMLProviderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersSamlCreate(context.Background()).SAMLProviderRequest(sAMLProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersSamlCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersSamlCreate`: SAMLProvider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersSamlCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvidersSamlCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sAMLProviderRequest** | [**SAMLProviderRequest**](SAMLProviderRequest.md) |  | 

### Return type

[**SAMLProvider**](SAMLProvider.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersSamlDestroy

> ProvidersSamlDestroy(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this SAML Provider.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersSamlDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersSamlDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this SAML Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersSamlDestroyRequest struct via the builder pattern


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


## ProvidersSamlImportMetadataCreate

> ProvidersSamlImportMetadataCreate(ctx).Name(name).AuthorizationFlow(authorizationFlow).File(file).Execute()





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
    name := "name_example" // string | 
    authorizationFlow := "authorizationFlow_example" // string | 
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersSamlImportMetadataCreate(context.Background()).Name(name).AuthorizationFlow(authorizationFlow).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersSamlImportMetadataCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvidersSamlImportMetadataCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **authorizationFlow** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersSamlList

> PaginatedSAMLProviderList ProvidersSamlList(ctx).AcsUrl(acsUrl).AssertionValidNotBefore(assertionValidNotBefore).AssertionValidNotOnOrAfter(assertionValidNotOnOrAfter).Audience(audience).AuthorizationFlow(authorizationFlow).DigestAlgorithm(digestAlgorithm).Issuer(issuer).Name(name).NameIdMapping(nameIdMapping).Ordering(ordering).Page(page).PageSize(pageSize).PropertyMappings(propertyMappings).Search(search).SessionValidNotOnOrAfter(sessionValidNotOnOrAfter).SignatureAlgorithm(signatureAlgorithm).SigningKp(signingKp).SpBinding(spBinding).VerificationKp(verificationKp).Execute()





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
    acsUrl := "acsUrl_example" // string |  (optional)
    assertionValidNotBefore := "assertionValidNotBefore_example" // string |  (optional)
    assertionValidNotOnOrAfter := "assertionValidNotOnOrAfter_example" // string |  (optional)
    audience := "audience_example" // string |  (optional)
    authorizationFlow := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    digestAlgorithm := "digestAlgorithm_example" // string |  (optional)
    issuer := "issuer_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    nameIdMapping := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    propertyMappings := []string{"Inner_example"} // []string |  (optional)
    search := "search_example" // string | A search term. (optional)
    sessionValidNotOnOrAfter := "sessionValidNotOnOrAfter_example" // string |  (optional)
    signatureAlgorithm := "signatureAlgorithm_example" // string |  (optional)
    signingKp := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    spBinding := "spBinding_example" // string | This determines how authentik sends the response back to the Service Provider. (optional)
    verificationKp := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersSamlList(context.Background()).AcsUrl(acsUrl).AssertionValidNotBefore(assertionValidNotBefore).AssertionValidNotOnOrAfter(assertionValidNotOnOrAfter).Audience(audience).AuthorizationFlow(authorizationFlow).DigestAlgorithm(digestAlgorithm).Issuer(issuer).Name(name).NameIdMapping(nameIdMapping).Ordering(ordering).Page(page).PageSize(pageSize).PropertyMappings(propertyMappings).Search(search).SessionValidNotOnOrAfter(sessionValidNotOnOrAfter).SignatureAlgorithm(signatureAlgorithm).SigningKp(signingKp).SpBinding(spBinding).VerificationKp(verificationKp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersSamlList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersSamlList`: PaginatedSAMLProviderList
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersSamlList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvidersSamlListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acsUrl** | **string** |  | 
 **assertionValidNotBefore** | **string** |  | 
 **assertionValidNotOnOrAfter** | **string** |  | 
 **audience** | **string** |  | 
 **authorizationFlow** | **string** |  | 
 **digestAlgorithm** | **string** |  | 
 **issuer** | **string** |  | 
 **name** | **string** |  | 
 **nameIdMapping** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **propertyMappings** | **[]string** |  | 
 **search** | **string** | A search term. | 
 **sessionValidNotOnOrAfter** | **string** |  | 
 **signatureAlgorithm** | **string** |  | 
 **signingKp** | **string** |  | 
 **spBinding** | **string** | This determines how authentik sends the response back to the Service Provider. | 
 **verificationKp** | **string** |  | 

### Return type

[**PaginatedSAMLProviderList**](PaginatedSAMLProviderList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersSamlMetadataRetrieve

> SAMLMetadata ProvidersSamlMetadataRetrieve(ctx, id).Download(download).ForceBinding(forceBinding).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this SAML Provider.
    download := true // bool |  (optional)
    forceBinding := "forceBinding_example" // string | Optionally force the metadata to only include one binding. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersSamlMetadataRetrieve(context.Background(), id).Download(download).ForceBinding(forceBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersSamlMetadataRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersSamlMetadataRetrieve`: SAMLMetadata
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersSamlMetadataRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this SAML Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersSamlMetadataRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **download** | **bool** |  | 
 **forceBinding** | **string** | Optionally force the metadata to only include one binding. | 

### Return type

[**SAMLMetadata**](SAMLMetadata.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersSamlPartialUpdate

> SAMLProvider ProvidersSamlPartialUpdate(ctx, id).PatchedSAMLProviderRequest(patchedSAMLProviderRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this SAML Provider.
    patchedSAMLProviderRequest := *openapiclient.NewPatchedSAMLProviderRequest() // PatchedSAMLProviderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersSamlPartialUpdate(context.Background(), id).PatchedSAMLProviderRequest(patchedSAMLProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersSamlPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersSamlPartialUpdate`: SAMLProvider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersSamlPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this SAML Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersSamlPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedSAMLProviderRequest** | [**PatchedSAMLProviderRequest**](PatchedSAMLProviderRequest.md) |  | 

### Return type

[**SAMLProvider**](SAMLProvider.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersSamlRetrieve

> SAMLProvider ProvidersSamlRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this SAML Provider.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersSamlRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersSamlRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersSamlRetrieve`: SAMLProvider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersSamlRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this SAML Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersSamlRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SAMLProvider**](SAMLProvider.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersSamlUpdate

> SAMLProvider ProvidersSamlUpdate(ctx, id).SAMLProviderRequest(sAMLProviderRequest).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this SAML Provider.
    sAMLProviderRequest := *openapiclient.NewSAMLProviderRequest("Name_example", "AuthorizationFlow_example", "AcsUrl_example") // SAMLProviderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersSamlUpdate(context.Background(), id).SAMLProviderRequest(sAMLProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersSamlUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersSamlUpdate`: SAMLProvider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersSamlUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this SAML Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersSamlUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sAMLProviderRequest** | [**SAMLProviderRequest**](SAMLProviderRequest.md) |  | 

### Return type

[**SAMLProvider**](SAMLProvider.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersSamlUsedByList

> []UsedBy ProvidersSamlUsedByList(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this SAML Provider.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersSamlUsedByList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersSamlUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersSamlUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersSamlUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this SAML Provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersSamlUsedByListRequest struct via the builder pattern


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

