# SyncObjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncObjectModel** | [**SyncObjectModelEnum**](SyncObjectModelEnum.md) |  | 
**SyncObjectId** | **string** |  | 
**OverrideDryRun** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSyncObjectRequest

`func NewSyncObjectRequest(syncObjectModel SyncObjectModelEnum, syncObjectId string, ) *SyncObjectRequest`

NewSyncObjectRequest instantiates a new SyncObjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncObjectRequestWithDefaults

`func NewSyncObjectRequestWithDefaults() *SyncObjectRequest`

NewSyncObjectRequestWithDefaults instantiates a new SyncObjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyncObjectModel

`func (o *SyncObjectRequest) GetSyncObjectModel() SyncObjectModelEnum`

GetSyncObjectModel returns the SyncObjectModel field if non-nil, zero value otherwise.

### GetSyncObjectModelOk

`func (o *SyncObjectRequest) GetSyncObjectModelOk() (*SyncObjectModelEnum, bool)`

GetSyncObjectModelOk returns a tuple with the SyncObjectModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncObjectModel

`func (o *SyncObjectRequest) SetSyncObjectModel(v SyncObjectModelEnum)`

SetSyncObjectModel sets SyncObjectModel field to given value.


### GetSyncObjectId

`func (o *SyncObjectRequest) GetSyncObjectId() string`

GetSyncObjectId returns the SyncObjectId field if non-nil, zero value otherwise.

### GetSyncObjectIdOk

`func (o *SyncObjectRequest) GetSyncObjectIdOk() (*string, bool)`

GetSyncObjectIdOk returns a tuple with the SyncObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncObjectId

`func (o *SyncObjectRequest) SetSyncObjectId(v string)`

SetSyncObjectId sets SyncObjectId field to given value.


### GetOverrideDryRun

`func (o *SyncObjectRequest) GetOverrideDryRun() bool`

GetOverrideDryRun returns the OverrideDryRun field if non-nil, zero value otherwise.

### GetOverrideDryRunOk

`func (o *SyncObjectRequest) GetOverrideDryRunOk() (*bool, bool)`

GetOverrideDryRunOk returns a tuple with the OverrideDryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideDryRun

`func (o *SyncObjectRequest) SetOverrideDryRun(v bool)`

SetOverrideDryRun sets OverrideDryRun field to given value.

### HasOverrideDryRun

`func (o *SyncObjectRequest) HasOverrideDryRun() bool`

HasOverrideDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


