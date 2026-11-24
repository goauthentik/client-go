# SyncStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRunning** | **bool** |  | 
**LastSuccessfulSync** | Pointer to **time.Time** |  | [optional] 
**LastSyncStatus** | Pointer to [**TaskAggregatedStatusEnum**](TaskAggregatedStatusEnum.md) |  | [optional] 

## Methods

### NewSyncStatus

`func NewSyncStatus(isRunning bool, ) *SyncStatus`

NewSyncStatus instantiates a new SyncStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncStatusWithDefaults

`func NewSyncStatusWithDefaults() *SyncStatus`

NewSyncStatusWithDefaults instantiates a new SyncStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRunning

`func (o *SyncStatus) GetIsRunning() bool`

GetIsRunning returns the IsRunning field if non-nil, zero value otherwise.

### GetIsRunningOk

`func (o *SyncStatus) GetIsRunningOk() (*bool, bool)`

GetIsRunningOk returns a tuple with the IsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRunning

`func (o *SyncStatus) SetIsRunning(v bool)`

SetIsRunning sets IsRunning field to given value.


### GetLastSuccessfulSync

`func (o *SyncStatus) GetLastSuccessfulSync() time.Time`

GetLastSuccessfulSync returns the LastSuccessfulSync field if non-nil, zero value otherwise.

### GetLastSuccessfulSyncOk

`func (o *SyncStatus) GetLastSuccessfulSyncOk() (*time.Time, bool)`

GetLastSuccessfulSyncOk returns a tuple with the LastSuccessfulSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulSync

`func (o *SyncStatus) SetLastSuccessfulSync(v time.Time)`

SetLastSuccessfulSync sets LastSuccessfulSync field to given value.

### HasLastSuccessfulSync

`func (o *SyncStatus) HasLastSuccessfulSync() bool`

HasLastSuccessfulSync returns a boolean if a field has been set.

### GetLastSyncStatus

`func (o *SyncStatus) GetLastSyncStatus() TaskAggregatedStatusEnum`

GetLastSyncStatus returns the LastSyncStatus field if non-nil, zero value otherwise.

### GetLastSyncStatusOk

`func (o *SyncStatus) GetLastSyncStatusOk() (*TaskAggregatedStatusEnum, bool)`

GetLastSyncStatusOk returns a tuple with the LastSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncStatus

`func (o *SyncStatus) SetLastSyncStatus(v TaskAggregatedStatusEnum)`

SetLastSyncStatus sets LastSyncStatus field to given value.

### HasLastSyncStatus

`func (o *SyncStatus) HasLastSyncStatus() bool`

HasLastSyncStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


