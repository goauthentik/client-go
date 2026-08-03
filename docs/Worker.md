# Worker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkerId** | **string** |  | 
**Version** | **string** |  | 
**VersionMatching** | **bool** |  | 

## Methods

### NewWorker

`func NewWorker(workerId string, version string, versionMatching bool, ) *Worker`

NewWorker instantiates a new Worker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkerWithDefaults

`func NewWorkerWithDefaults() *Worker`

NewWorkerWithDefaults instantiates a new Worker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkerId

`func (o *Worker) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *Worker) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *Worker) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.


### GetVersion

`func (o *Worker) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Worker) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Worker) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVersionMatching

`func (o *Worker) GetVersionMatching() bool`

GetVersionMatching returns the VersionMatching field if non-nil, zero value otherwise.

### GetVersionMatchingOk

`func (o *Worker) GetVersionMatchingOk() (*bool, bool)`

GetVersionMatchingOk returns a tuple with the VersionMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionMatching

`func (o *Worker) SetVersionMatching(v bool)`

SetVersionMatching sets VersionMatching field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


