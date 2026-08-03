# SourceStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Source** | **string** |  | 
**ResumeTimeout** | Pointer to **string** | Amount of time a user can take to return from the source to continue the flow (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3) | [optional] 

## Methods

### NewSourceStageRequest

`func NewSourceStageRequest(name string, source string, ) *SourceStageRequest`

NewSourceStageRequest instantiates a new SourceStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceStageRequestWithDefaults

`func NewSourceStageRequestWithDefaults() *SourceStageRequest`

NewSourceStageRequestWithDefaults instantiates a new SourceStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourceStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *SourceStageRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SourceStageRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SourceStageRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetResumeTimeout

`func (o *SourceStageRequest) GetResumeTimeout() string`

GetResumeTimeout returns the ResumeTimeout field if non-nil, zero value otherwise.

### GetResumeTimeoutOk

`func (o *SourceStageRequest) GetResumeTimeoutOk() (*string, bool)`

GetResumeTimeoutOk returns a tuple with the ResumeTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeTimeout

`func (o *SourceStageRequest) SetResumeTimeout(v string)`

SetResumeTimeout sets ResumeTimeout field to given value.

### HasResumeTimeout

`func (o *SourceStageRequest) HasResumeTimeout() bool`

HasResumeTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


