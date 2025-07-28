# \EventsApi

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsEventsActionsList**](EventsApi.md#EventsEventsActionsList) | **Get** /events/events/actions/ | 
[**EventsEventsCreate**](EventsApi.md#EventsEventsCreate) | **Post** /events/events/ | 
[**EventsEventsDestroy**](EventsApi.md#EventsEventsDestroy) | **Delete** /events/events/{event_uuid}/ | 
[**EventsEventsList**](EventsApi.md#EventsEventsList) | **Get** /events/events/ | 
[**EventsEventsPartialUpdate**](EventsApi.md#EventsEventsPartialUpdate) | **Patch** /events/events/{event_uuid}/ | 
[**EventsEventsRetrieve**](EventsApi.md#EventsEventsRetrieve) | **Get** /events/events/{event_uuid}/ | 
[**EventsEventsTopPerUserList**](EventsApi.md#EventsEventsTopPerUserList) | **Get** /events/events/top_per_user/ | 
[**EventsEventsUpdate**](EventsApi.md#EventsEventsUpdate) | **Put** /events/events/{event_uuid}/ | 
[**EventsEventsVolumeList**](EventsApi.md#EventsEventsVolumeList) | **Get** /events/events/volume/ | 
[**EventsNotificationsDestroy**](EventsApi.md#EventsNotificationsDestroy) | **Delete** /events/notifications/{uuid}/ | 
[**EventsNotificationsList**](EventsApi.md#EventsNotificationsList) | **Get** /events/notifications/ | 
[**EventsNotificationsMarkAllSeenCreate**](EventsApi.md#EventsNotificationsMarkAllSeenCreate) | **Post** /events/notifications/mark_all_seen/ | 
[**EventsNotificationsPartialUpdate**](EventsApi.md#EventsNotificationsPartialUpdate) | **Patch** /events/notifications/{uuid}/ | 
[**EventsNotificationsRetrieve**](EventsApi.md#EventsNotificationsRetrieve) | **Get** /events/notifications/{uuid}/ | 
[**EventsNotificationsUpdate**](EventsApi.md#EventsNotificationsUpdate) | **Put** /events/notifications/{uuid}/ | 
[**EventsNotificationsUsedByList**](EventsApi.md#EventsNotificationsUsedByList) | **Get** /events/notifications/{uuid}/used_by/ | 
[**EventsRulesCreate**](EventsApi.md#EventsRulesCreate) | **Post** /events/rules/ | 
[**EventsRulesDestroy**](EventsApi.md#EventsRulesDestroy) | **Delete** /events/rules/{pbm_uuid}/ | 
[**EventsRulesList**](EventsApi.md#EventsRulesList) | **Get** /events/rules/ | 
[**EventsRulesPartialUpdate**](EventsApi.md#EventsRulesPartialUpdate) | **Patch** /events/rules/{pbm_uuid}/ | 
[**EventsRulesRetrieve**](EventsApi.md#EventsRulesRetrieve) | **Get** /events/rules/{pbm_uuid}/ | 
[**EventsRulesUpdate**](EventsApi.md#EventsRulesUpdate) | **Put** /events/rules/{pbm_uuid}/ | 
[**EventsRulesUsedByList**](EventsApi.md#EventsRulesUsedByList) | **Get** /events/rules/{pbm_uuid}/used_by/ | 
[**EventsTransportsCreate**](EventsApi.md#EventsTransportsCreate) | **Post** /events/transports/ | 
[**EventsTransportsDestroy**](EventsApi.md#EventsTransportsDestroy) | **Delete** /events/transports/{uuid}/ | 
[**EventsTransportsList**](EventsApi.md#EventsTransportsList) | **Get** /events/transports/ | 
[**EventsTransportsPartialUpdate**](EventsApi.md#EventsTransportsPartialUpdate) | **Patch** /events/transports/{uuid}/ | 
[**EventsTransportsRetrieve**](EventsApi.md#EventsTransportsRetrieve) | **Get** /events/transports/{uuid}/ | 
[**EventsTransportsTestCreate**](EventsApi.md#EventsTransportsTestCreate) | **Post** /events/transports/{uuid}/test/ | 
[**EventsTransportsUpdate**](EventsApi.md#EventsTransportsUpdate) | **Put** /events/transports/{uuid}/ | 
[**EventsTransportsUsedByList**](EventsApi.md#EventsTransportsUsedByList) | **Get** /events/transports/{uuid}/used_by/ | 



## EventsEventsActionsList

> []TypeCreate EventsEventsActionsList(ctx).Execute()





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
    resp, r, err := apiClient.EventsApi.EventsEventsActionsList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsEventsActionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsEventsActionsList`: []TypeCreate
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsEventsActionsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsActionsListRequest struct via the builder pattern


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


## EventsEventsCreate

> Event EventsEventsCreate(ctx).EventRequest(eventRequest).Execute()





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
    eventRequest := *openapiclient.NewEventRequest(openapiclient.EventActions("login"), "App_example") // EventRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsEventsCreate(context.Background()).EventRequest(eventRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsEventsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsEventsCreate`: Event
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsEventsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventRequest** | [**EventRequest**](EventRequest.md) |  | 

### Return type

[**Event**](Event.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsEventsDestroy

> EventsEventsDestroy(ctx, eventUuid).Execute()





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
    eventUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Event.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsEventsDestroy(context.Background(), eventUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsEventsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventUuid** | **string** | A UUID string identifying this Event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsDestroyRequest struct via the builder pattern


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


## EventsEventsList

> PaginatedEventList EventsEventsList(ctx).Action(action).Actions(actions).BrandName(brandName).ClientIp(clientIp).ContextAuthorizedApp(contextAuthorizedApp).ContextModelApp(contextModelApp).ContextModelName(contextModelName).ContextModelPk(contextModelPk).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Username(username).Execute()





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
    action := "action_example" // string |  (optional)
    actions := []string{"Actions_example"} // []string |  (optional)
    brandName := "brandName_example" // string | Brand name (optional)
    clientIp := "clientIp_example" // string |  (optional)
    contextAuthorizedApp := "contextAuthorizedApp_example" // string | Context Authorized application (optional)
    contextModelApp := "contextModelApp_example" // string | Context Model App (optional)
    contextModelName := "contextModelName_example" // string | Context Model Name (optional)
    contextModelPk := "contextModelPk_example" // string | Context Model Primary Key (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    username := "username_example" // string | Username (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsEventsList(context.Background()).Action(action).Actions(actions).BrandName(brandName).ClientIp(clientIp).ContextAuthorizedApp(contextAuthorizedApp).ContextModelApp(contextModelApp).ContextModelName(contextModelName).ContextModelPk(contextModelPk).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsEventsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsEventsList`: PaginatedEventList
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsEventsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** |  | 
 **actions** | **[]string** |  | 
 **brandName** | **string** | Brand name | 
 **clientIp** | **string** |  | 
 **contextAuthorizedApp** | **string** | Context Authorized application | 
 **contextModelApp** | **string** | Context Model App | 
 **contextModelName** | **string** | Context Model Name | 
 **contextModelPk** | **string** | Context Model Primary Key | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **username** | **string** | Username | 

### Return type

[**PaginatedEventList**](PaginatedEventList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsEventsPartialUpdate

> Event EventsEventsPartialUpdate(ctx, eventUuid).PatchedEventRequest(patchedEventRequest).Execute()





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
    eventUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Event.
    patchedEventRequest := *openapiclient.NewPatchedEventRequest() // PatchedEventRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsEventsPartialUpdate(context.Background(), eventUuid).PatchedEventRequest(patchedEventRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsEventsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsEventsPartialUpdate`: Event
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsEventsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventUuid** | **string** | A UUID string identifying this Event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedEventRequest** | [**PatchedEventRequest**](PatchedEventRequest.md) |  | 

### Return type

[**Event**](Event.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsEventsRetrieve

> Event EventsEventsRetrieve(ctx, eventUuid).Execute()





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
    eventUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Event.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsEventsRetrieve(context.Background(), eventUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsEventsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsEventsRetrieve`: Event
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsEventsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventUuid** | **string** | A UUID string identifying this Event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Event**](Event.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsEventsTopPerUserList

> []EventTopPerUser EventsEventsTopPerUserList(ctx).Action(action).TopN(topN).Execute()





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
    action := "action_example" // string |  (optional)
    topN := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsEventsTopPerUserList(context.Background()).Action(action).TopN(topN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsEventsTopPerUserList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsEventsTopPerUserList`: []EventTopPerUser
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsEventsTopPerUserList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsTopPerUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** |  | 
 **topN** | **int32** |  | 

### Return type

[**[]EventTopPerUser**](EventTopPerUser.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsEventsUpdate

> Event EventsEventsUpdate(ctx, eventUuid).EventRequest(eventRequest).Execute()





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
    eventUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Event.
    eventRequest := *openapiclient.NewEventRequest(openapiclient.EventActions("login"), "App_example") // EventRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsEventsUpdate(context.Background(), eventUuid).EventRequest(eventRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsEventsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsEventsUpdate`: Event
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsEventsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventUuid** | **string** | A UUID string identifying this Event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventRequest** | [**EventRequest**](EventRequest.md) |  | 

### Return type

[**Event**](Event.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsEventsVolumeList

> []EventVolume EventsEventsVolumeList(ctx).Action(action).Actions(actions).BrandName(brandName).ClientIp(clientIp).ContextAuthorizedApp(contextAuthorizedApp).ContextModelApp(contextModelApp).ContextModelName(contextModelName).ContextModelPk(contextModelPk).HistoryDays(historyDays).Ordering(ordering).Search(search).Username(username).Execute()





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
    action := "action_example" // string |  (optional)
    actions := []string{"Actions_example"} // []string |  (optional)
    brandName := "brandName_example" // string | Brand name (optional)
    clientIp := "clientIp_example" // string |  (optional)
    contextAuthorizedApp := "contextAuthorizedApp_example" // string | Context Authorized application (optional)
    contextModelApp := "contextModelApp_example" // string | Context Model App (optional)
    contextModelName := "contextModelName_example" // string | Context Model Name (optional)
    contextModelPk := "contextModelPk_example" // string | Context Model Primary Key (optional)
    historyDays := float32(8.14) // float32 |  (optional) (default to 7)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    search := "search_example" // string | A search term. (optional)
    username := "username_example" // string | Username (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsEventsVolumeList(context.Background()).Action(action).Actions(actions).BrandName(brandName).ClientIp(clientIp).ContextAuthorizedApp(contextAuthorizedApp).ContextModelApp(contextModelApp).ContextModelName(contextModelName).ContextModelPk(contextModelPk).HistoryDays(historyDays).Ordering(ordering).Search(search).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsEventsVolumeList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsEventsVolumeList`: []EventVolume
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsEventsVolumeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsVolumeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** |  | 
 **actions** | **[]string** |  | 
 **brandName** | **string** | Brand name | 
 **clientIp** | **string** |  | 
 **contextAuthorizedApp** | **string** | Context Authorized application | 
 **contextModelApp** | **string** | Context Model App | 
 **contextModelName** | **string** | Context Model Name | 
 **contextModelPk** | **string** | Context Model Primary Key | 
 **historyDays** | **float32** |  | [default to 7]
 **ordering** | **string** | Which field to use when ordering the results. | 
 **search** | **string** | A search term. | 
 **username** | **string** | Username | 

### Return type

[**[]EventVolume**](EventVolume.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsNotificationsDestroy

> EventsNotificationsDestroy(ctx, uuid).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsNotificationsDestroy(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsNotificationsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsNotificationsDestroyRequest struct via the builder pattern


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


## EventsNotificationsList

> PaginatedNotificationList EventsNotificationsList(ctx).Body(body).Created(created).Event(event).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Seen(seen).Severity(severity).User(user).Execute()





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
    body := "body_example" // string |  (optional)
    created := time.Now() // time.Time |  (optional)
    event := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    seen := true // bool |  (optional)
    severity := "severity_example" // string |  (optional)
    user := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsNotificationsList(context.Background()).Body(body).Created(created).Event(event).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Seen(seen).Severity(severity).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsNotificationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsNotificationsList`: PaginatedNotificationList
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsNotificationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsNotificationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 
 **created** | **time.Time** |  | 
 **event** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **seen** | **bool** |  | 
 **severity** | **string** |  | 
 **user** | **int32** |  | 

### Return type

[**PaginatedNotificationList**](PaginatedNotificationList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsNotificationsMarkAllSeenCreate

> EventsNotificationsMarkAllSeenCreate(ctx).Execute()





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
    resp, r, err := apiClient.EventsApi.EventsNotificationsMarkAllSeenCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsNotificationsMarkAllSeenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEventsNotificationsMarkAllSeenCreateRequest struct via the builder pattern


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


## EventsNotificationsPartialUpdate

> Notification EventsNotificationsPartialUpdate(ctx, uuid).PatchedNotificationRequest(patchedNotificationRequest).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification.
    patchedNotificationRequest := *openapiclient.NewPatchedNotificationRequest() // PatchedNotificationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsNotificationsPartialUpdate(context.Background(), uuid).PatchedNotificationRequest(patchedNotificationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsNotificationsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsNotificationsPartialUpdate`: Notification
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsNotificationsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsNotificationsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedNotificationRequest** | [**PatchedNotificationRequest**](PatchedNotificationRequest.md) |  | 

### Return type

[**Notification**](Notification.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsNotificationsRetrieve

> Notification EventsNotificationsRetrieve(ctx, uuid).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsNotificationsRetrieve(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsNotificationsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsNotificationsRetrieve`: Notification
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsNotificationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsNotificationsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Notification**](Notification.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsNotificationsUpdate

> Notification EventsNotificationsUpdate(ctx, uuid).NotificationRequest(notificationRequest).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification.
    notificationRequest := *openapiclient.NewNotificationRequest() // NotificationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsNotificationsUpdate(context.Background(), uuid).NotificationRequest(notificationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsNotificationsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsNotificationsUpdate`: Notification
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsNotificationsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsNotificationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationRequest** | [**NotificationRequest**](NotificationRequest.md) |  | 

### Return type

[**Notification**](Notification.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsNotificationsUsedByList

> []UsedBy EventsNotificationsUsedByList(ctx, uuid).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsNotificationsUsedByList(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsNotificationsUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsNotificationsUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsNotificationsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsNotificationsUsedByListRequest struct via the builder pattern


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


## EventsRulesCreate

> NotificationRule EventsRulesCreate(ctx).NotificationRuleRequest(notificationRuleRequest).Execute()





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
    notificationRuleRequest := *openapiclient.NewNotificationRuleRequest("Name_example") // NotificationRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsRulesCreate(context.Background()).NotificationRuleRequest(notificationRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsRulesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsRulesCreate`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsRulesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsRulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationRuleRequest** | [**NotificationRuleRequest**](NotificationRuleRequest.md) |  | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsRulesDestroy

> EventsRulesDestroy(ctx, pbmUuid).Execute()





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
    pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsRulesDestroy(context.Background(), pbmUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsRulesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this Notification Rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsRulesDestroyRequest struct via the builder pattern


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


## EventsRulesList

> PaginatedNotificationRuleList EventsRulesList(ctx).DestinationGroupName(destinationGroupName).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Severity(severity).Execute()





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
    destinationGroupName := "destinationGroupName_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    severity := "severity_example" // string | Controls which severity level the created notifications will have.   (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsRulesList(context.Background()).DestinationGroupName(destinationGroupName).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Severity(severity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsRulesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsRulesList`: PaginatedNotificationRuleList
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsRulesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsRulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationGroupName** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **severity** | **string** | Controls which severity level the created notifications will have.   | 

### Return type

[**PaginatedNotificationRuleList**](PaginatedNotificationRuleList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsRulesPartialUpdate

> NotificationRule EventsRulesPartialUpdate(ctx, pbmUuid).PatchedNotificationRuleRequest(patchedNotificationRuleRequest).Execute()





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
    pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Rule.
    patchedNotificationRuleRequest := *openapiclient.NewPatchedNotificationRuleRequest() // PatchedNotificationRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsRulesPartialUpdate(context.Background(), pbmUuid).PatchedNotificationRuleRequest(patchedNotificationRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsRulesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsRulesPartialUpdate`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsRulesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this Notification Rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsRulesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedNotificationRuleRequest** | [**PatchedNotificationRuleRequest**](PatchedNotificationRuleRequest.md) |  | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsRulesRetrieve

> NotificationRule EventsRulesRetrieve(ctx, pbmUuid).Execute()





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
    pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsRulesRetrieve(context.Background(), pbmUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsRulesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsRulesRetrieve`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsRulesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this Notification Rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsRulesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsRulesUpdate

> NotificationRule EventsRulesUpdate(ctx, pbmUuid).NotificationRuleRequest(notificationRuleRequest).Execute()





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
    pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Rule.
    notificationRuleRequest := *openapiclient.NewNotificationRuleRequest("Name_example") // NotificationRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsRulesUpdate(context.Background(), pbmUuid).NotificationRuleRequest(notificationRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsRulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsRulesUpdate`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsRulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this Notification Rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsRulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationRuleRequest** | [**NotificationRuleRequest**](NotificationRuleRequest.md) |  | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsRulesUsedByList

> []UsedBy EventsRulesUsedByList(ctx, pbmUuid).Execute()





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
    pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsRulesUsedByList(context.Background(), pbmUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsRulesUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsRulesUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsRulesUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this Notification Rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsRulesUsedByListRequest struct via the builder pattern


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


## EventsTransportsCreate

> NotificationTransport EventsTransportsCreate(ctx).NotificationTransportRequest(notificationTransportRequest).Execute()





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
    notificationTransportRequest := *openapiclient.NewNotificationTransportRequest("Name_example") // NotificationTransportRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsTransportsCreate(context.Background()).NotificationTransportRequest(notificationTransportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsTransportsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsTransportsCreate`: NotificationTransport
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsTransportsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsTransportsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationTransportRequest** | [**NotificationTransportRequest**](NotificationTransportRequest.md) |  | 

### Return type

[**NotificationTransport**](NotificationTransport.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsTransportsDestroy

> EventsTransportsDestroy(ctx, uuid).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Transport.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsTransportsDestroy(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsTransportsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification Transport. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsTransportsDestroyRequest struct via the builder pattern


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


## EventsTransportsList

> PaginatedNotificationTransportList EventsTransportsList(ctx).Mode(mode).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).SendOnce(sendOnce).WebhookUrl(webhookUrl).Execute()





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
    mode := "mode_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    sendOnce := true // bool |  (optional)
    webhookUrl := "webhookUrl_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsTransportsList(context.Background()).Mode(mode).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).SendOnce(sendOnce).WebhookUrl(webhookUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsTransportsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsTransportsList`: PaginatedNotificationTransportList
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsTransportsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsTransportsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mode** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **sendOnce** | **bool** |  | 
 **webhookUrl** | **string** |  | 

### Return type

[**PaginatedNotificationTransportList**](PaginatedNotificationTransportList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsTransportsPartialUpdate

> NotificationTransport EventsTransportsPartialUpdate(ctx, uuid).PatchedNotificationTransportRequest(patchedNotificationTransportRequest).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Transport.
    patchedNotificationTransportRequest := *openapiclient.NewPatchedNotificationTransportRequest() // PatchedNotificationTransportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsTransportsPartialUpdate(context.Background(), uuid).PatchedNotificationTransportRequest(patchedNotificationTransportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsTransportsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsTransportsPartialUpdate`: NotificationTransport
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsTransportsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification Transport. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsTransportsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedNotificationTransportRequest** | [**PatchedNotificationTransportRequest**](PatchedNotificationTransportRequest.md) |  | 

### Return type

[**NotificationTransport**](NotificationTransport.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsTransportsRetrieve

> NotificationTransport EventsTransportsRetrieve(ctx, uuid).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Transport.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsTransportsRetrieve(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsTransportsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsTransportsRetrieve`: NotificationTransport
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsTransportsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification Transport. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsTransportsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationTransport**](NotificationTransport.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsTransportsTestCreate

> NotificationTransportTest EventsTransportsTestCreate(ctx, uuid).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Transport.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsTransportsTestCreate(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsTransportsTestCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsTransportsTestCreate`: NotificationTransportTest
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsTransportsTestCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification Transport. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsTransportsTestCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationTransportTest**](NotificationTransportTest.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsTransportsUpdate

> NotificationTransport EventsTransportsUpdate(ctx, uuid).NotificationTransportRequest(notificationTransportRequest).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Transport.
    notificationTransportRequest := *openapiclient.NewNotificationTransportRequest("Name_example") // NotificationTransportRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsTransportsUpdate(context.Background(), uuid).NotificationTransportRequest(notificationTransportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsTransportsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsTransportsUpdate`: NotificationTransport
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsTransportsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification Transport. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsTransportsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationTransportRequest** | [**NotificationTransportRequest**](NotificationTransportRequest.md) |  | 

### Return type

[**NotificationTransport**](NotificationTransport.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsTransportsUsedByList

> []UsedBy EventsTransportsUsedByList(ctx, uuid).Execute()





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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Transport.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.EventsTransportsUsedByList(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.EventsTransportsUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsTransportsUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.EventsTransportsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification Transport. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsTransportsUsedByListRequest struct via the builder pattern


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

