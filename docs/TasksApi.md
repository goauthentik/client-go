# \TasksApi

All URIs are relative to */api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TasksSchedulesList**](TasksApi.md#TasksSchedulesList) | **Get** /tasks/schedules/ | 
[**TasksSchedulesPartialUpdate**](TasksApi.md#TasksSchedulesPartialUpdate) | **Patch** /tasks/schedules/{id}/ | 
[**TasksSchedulesRetrieve**](TasksApi.md#TasksSchedulesRetrieve) | **Get** /tasks/schedules/{id}/ | 
[**TasksSchedulesSendCreate**](TasksApi.md#TasksSchedulesSendCreate) | **Post** /tasks/schedules/{id}/send/ | 
[**TasksSchedulesUpdate**](TasksApi.md#TasksSchedulesUpdate) | **Put** /tasks/schedules/{id}/ | 
[**TasksTasksList**](TasksApi.md#TasksTasksList) | **Get** /tasks/tasks/ | 
[**TasksTasksRetrieve**](TasksApi.md#TasksTasksRetrieve) | **Get** /tasks/tasks/{message_id}/ | 
[**TasksTasksRetryCreate**](TasksApi.md#TasksTasksRetryCreate) | **Post** /tasks/tasks/{message_id}/retry/ | 
[**TasksTasksStatusRetrieve**](TasksApi.md#TasksTasksStatusRetrieve) | **Get** /tasks/tasks/status/ | 
[**TasksWorkersList**](TasksApi.md#TasksWorkersList) | **Get** /tasks/workers | 



## TasksSchedulesList

> PaginatedScheduleList TasksSchedulesList(ctx).ActorName(actorName).Ordering(ordering).Page(page).PageSize(pageSize).Paused(paused).RelObjContentTypeAppLabel(relObjContentTypeAppLabel).RelObjContentTypeModel(relObjContentTypeModel).RelObjId(relObjId).RelObjIdIsnull(relObjIdIsnull).Search(search).Execute()



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
    actorName := "actorName_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    paused := true // bool |  (optional)
    relObjContentTypeAppLabel := "relObjContentTypeAppLabel_example" // string |  (optional)
    relObjContentTypeModel := "relObjContentTypeModel_example" // string |  (optional)
    relObjId := "relObjId_example" // string |  (optional)
    relObjIdIsnull := true // bool |  (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksApi.TasksSchedulesList(context.Background()).ActorName(actorName).Ordering(ordering).Page(page).PageSize(pageSize).Paused(paused).RelObjContentTypeAppLabel(relObjContentTypeAppLabel).RelObjContentTypeModel(relObjContentTypeModel).RelObjId(relObjId).RelObjIdIsnull(relObjIdIsnull).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksSchedulesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksSchedulesList`: PaginatedScheduleList
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksSchedulesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksSchedulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actorName** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **paused** | **bool** |  | 
 **relObjContentTypeAppLabel** | **string** |  | 
 **relObjContentTypeModel** | **string** |  | 
 **relObjId** | **string** |  | 
 **relObjIdIsnull** | **bool** |  | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedScheduleList**](PaginatedScheduleList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksSchedulesPartialUpdate

> Schedule TasksSchedulesPartialUpdate(ctx, id).PatchedScheduleRequest(patchedScheduleRequest).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Schedule.
    patchedScheduleRequest := *openapiclient.NewPatchedScheduleRequest() // PatchedScheduleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksApi.TasksSchedulesPartialUpdate(context.Background(), id).PatchedScheduleRequest(patchedScheduleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksSchedulesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksSchedulesPartialUpdate`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksSchedulesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this Schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksSchedulesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedScheduleRequest** | [**PatchedScheduleRequest**](PatchedScheduleRequest.md) |  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksSchedulesRetrieve

> Schedule TasksSchedulesRetrieve(ctx, id).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Schedule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksApi.TasksSchedulesRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksSchedulesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksSchedulesRetrieve`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksSchedulesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this Schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksSchedulesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Schedule**](Schedule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksSchedulesSendCreate

> TasksSchedulesSendCreate(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Schedule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksApi.TasksSchedulesSendCreate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksSchedulesSendCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this Schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksSchedulesSendCreateRequest struct via the builder pattern


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


## TasksSchedulesUpdate

> Schedule TasksSchedulesUpdate(ctx, id).ScheduleRequest(scheduleRequest).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Schedule.
    scheduleRequest := *openapiclient.NewScheduleRequest("Crontab_example") // ScheduleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksApi.TasksSchedulesUpdate(context.Background(), id).ScheduleRequest(scheduleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksSchedulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksSchedulesUpdate`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksSchedulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this Schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksSchedulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduleRequest** | [**ScheduleRequest**](ScheduleRequest.md) |  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTasksList

> PaginatedTaskList TasksTasksList(ctx).ActorName(actorName).AggregatedStatus(aggregatedStatus).Ordering(ordering).Page(page).PageSize(pageSize).QueueName(queueName).RelObjContentTypeAppLabel(relObjContentTypeAppLabel).RelObjContentTypeModel(relObjContentTypeModel).RelObjId(relObjId).RelObjIdIsnull(relObjIdIsnull).Search(search).State(state).Execute()



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
    actorName := "actorName_example" // string |  (optional)
    aggregatedStatus := []string{"AggregatedStatus_example"} // []string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    queueName := "queueName_example" // string |  (optional)
    relObjContentTypeAppLabel := "relObjContentTypeAppLabel_example" // string |  (optional)
    relObjContentTypeModel := "relObjContentTypeModel_example" // string |  (optional)
    relObjId := "relObjId_example" // string |  (optional)
    relObjIdIsnull := true // bool |  (optional)
    search := "search_example" // string | A search term. (optional)
    state := "state_example" // string | Task status   (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksApi.TasksTasksList(context.Background()).ActorName(actorName).AggregatedStatus(aggregatedStatus).Ordering(ordering).Page(page).PageSize(pageSize).QueueName(queueName).RelObjContentTypeAppLabel(relObjContentTypeAppLabel).RelObjContentTypeModel(relObjContentTypeModel).RelObjId(relObjId).RelObjIdIsnull(relObjIdIsnull).Search(search).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksTasksList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksTasksList`: PaginatedTaskList
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksTasksList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksTasksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actorName** | **string** |  | 
 **aggregatedStatus** | **[]string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **queueName** | **string** |  | 
 **relObjContentTypeAppLabel** | **string** |  | 
 **relObjContentTypeModel** | **string** |  | 
 **relObjId** | **string** |  | 
 **relObjIdIsnull** | **bool** |  | 
 **search** | **string** | A search term. | 
 **state** | **string** | Task status   | 

### Return type

[**PaginatedTaskList**](PaginatedTaskList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTasksRetrieve

> Task TasksTasksRetrieve(ctx, messageId).Execute()



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
    messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Task.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksApi.TasksTasksRetrieve(context.Background(), messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksTasksRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksTasksRetrieve`: Task
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksTasksRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | A UUID string identifying this Task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTasksRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Task**](Task.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTasksRetryCreate

> TasksTasksRetryCreate(ctx, messageId).Execute()





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
    messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Task.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksApi.TasksTasksRetryCreate(context.Background(), messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksTasksRetryCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | A UUID string identifying this Task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTasksRetryCreateRequest struct via the builder pattern


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


## TasksTasksStatusRetrieve

> GlobalTaskStatus TasksTasksStatusRetrieve(ctx).Execute()





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
    resp, r, err := apiClient.TasksApi.TasksTasksStatusRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksTasksStatusRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksTasksStatusRetrieve`: GlobalTaskStatus
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksTasksStatusRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTasksStatusRetrieveRequest struct via the builder pattern


### Return type

[**GlobalTaskStatus**](GlobalTaskStatus.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksWorkersList

> []Worker TasksWorkersList(ctx).Execute()





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
    resp, r, err := apiClient.TasksApi.TasksWorkersList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksWorkersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksWorkersList`: []Worker
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksWorkersList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTasksWorkersListRequest struct via the builder pattern


### Return type

[**[]Worker**](Worker.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

