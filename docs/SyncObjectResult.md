# SyncObjectResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]LogEvent**](LogEvent.md) |  | [readonly] 

## Methods

### NewSyncObjectResult

`func NewSyncObjectResult(messages []LogEvent, ) *SyncObjectResult`

NewSyncObjectResult instantiates a new SyncObjectResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncObjectResultWithDefaults

`func NewSyncObjectResultWithDefaults() *SyncObjectResult`

NewSyncObjectResultWithDefaults instantiates a new SyncObjectResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *SyncObjectResult) GetMessages() []LogEvent`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *SyncObjectResult) GetMessagesOk() (*[]LogEvent, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *SyncObjectResult) SetMessages(v []LogEvent)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


