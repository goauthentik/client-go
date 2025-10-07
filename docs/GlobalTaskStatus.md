# GlobalTaskStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queued** | **int32** |  | [readonly] 
**Consumed** | **int32** |  | [readonly] 
**Preprocess** | **int32** |  | [readonly] 
**Running** | **int32** |  | [readonly] 
**Postprocess** | **int32** |  | [readonly] 
**Rejected** | **int32** |  | [readonly] 
**Done** | **int32** |  | [readonly] 
**Info** | **int32** |  | [readonly] 
**Warning** | **int32** |  | [readonly] 
**Error** | **int32** |  | [readonly] 

## Methods

### NewGlobalTaskStatus

`func NewGlobalTaskStatus(queued int32, consumed int32, preprocess int32, running int32, postprocess int32, rejected int32, done int32, info int32, warning int32, error_ int32, ) *GlobalTaskStatus`

NewGlobalTaskStatus instantiates a new GlobalTaskStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalTaskStatusWithDefaults

`func NewGlobalTaskStatusWithDefaults() *GlobalTaskStatus`

NewGlobalTaskStatusWithDefaults instantiates a new GlobalTaskStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueued

`func (o *GlobalTaskStatus) GetQueued() int32`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *GlobalTaskStatus) GetQueuedOk() (*int32, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *GlobalTaskStatus) SetQueued(v int32)`

SetQueued sets Queued field to given value.


### GetConsumed

`func (o *GlobalTaskStatus) GetConsumed() int32`

GetConsumed returns the Consumed field if non-nil, zero value otherwise.

### GetConsumedOk

`func (o *GlobalTaskStatus) GetConsumedOk() (*int32, bool)`

GetConsumedOk returns a tuple with the Consumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumed

`func (o *GlobalTaskStatus) SetConsumed(v int32)`

SetConsumed sets Consumed field to given value.


### GetPreprocess

`func (o *GlobalTaskStatus) GetPreprocess() int32`

GetPreprocess returns the Preprocess field if non-nil, zero value otherwise.

### GetPreprocessOk

`func (o *GlobalTaskStatus) GetPreprocessOk() (*int32, bool)`

GetPreprocessOk returns a tuple with the Preprocess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprocess

`func (o *GlobalTaskStatus) SetPreprocess(v int32)`

SetPreprocess sets Preprocess field to given value.


### GetRunning

`func (o *GlobalTaskStatus) GetRunning() int32`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *GlobalTaskStatus) GetRunningOk() (*int32, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *GlobalTaskStatus) SetRunning(v int32)`

SetRunning sets Running field to given value.


### GetPostprocess

`func (o *GlobalTaskStatus) GetPostprocess() int32`

GetPostprocess returns the Postprocess field if non-nil, zero value otherwise.

### GetPostprocessOk

`func (o *GlobalTaskStatus) GetPostprocessOk() (*int32, bool)`

GetPostprocessOk returns a tuple with the Postprocess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostprocess

`func (o *GlobalTaskStatus) SetPostprocess(v int32)`

SetPostprocess sets Postprocess field to given value.


### GetRejected

`func (o *GlobalTaskStatus) GetRejected() int32`

GetRejected returns the Rejected field if non-nil, zero value otherwise.

### GetRejectedOk

`func (o *GlobalTaskStatus) GetRejectedOk() (*int32, bool)`

GetRejectedOk returns a tuple with the Rejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejected

`func (o *GlobalTaskStatus) SetRejected(v int32)`

SetRejected sets Rejected field to given value.


### GetDone

`func (o *GlobalTaskStatus) GetDone() int32`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *GlobalTaskStatus) GetDoneOk() (*int32, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *GlobalTaskStatus) SetDone(v int32)`

SetDone sets Done field to given value.


### GetInfo

`func (o *GlobalTaskStatus) GetInfo() int32`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *GlobalTaskStatus) GetInfoOk() (*int32, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *GlobalTaskStatus) SetInfo(v int32)`

SetInfo sets Info field to given value.


### GetWarning

`func (o *GlobalTaskStatus) GetWarning() int32`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *GlobalTaskStatus) GetWarningOk() (*int32, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *GlobalTaskStatus) SetWarning(v int32)`

SetWarning sets Warning field to given value.


### GetError

`func (o *GlobalTaskStatus) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GlobalTaskStatus) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GlobalTaskStatus) SetError(v int32)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


