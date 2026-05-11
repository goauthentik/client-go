# PatchedSourceStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**ResumeTimeout** | Pointer to **string** | Amount of time a user can take to return from the source to continue the flow (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3) | [optional] 

## Methods

### NewPatchedSourceStageRequest

`func NewPatchedSourceStageRequest() *PatchedSourceStageRequest`

NewPatchedSourceStageRequest instantiates a new PatchedSourceStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSourceStageRequestWithDefaults

`func NewPatchedSourceStageRequestWithDefaults() *PatchedSourceStageRequest`

NewPatchedSourceStageRequestWithDefaults instantiates a new PatchedSourceStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedSourceStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedSourceStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedSourceStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedSourceStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSource

`func (o *PatchedSourceStageRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PatchedSourceStageRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PatchedSourceStageRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PatchedSourceStageRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetResumeTimeout

`func (o *PatchedSourceStageRequest) GetResumeTimeout() string`

GetResumeTimeout returns the ResumeTimeout field if non-nil, zero value otherwise.

### GetResumeTimeoutOk

`func (o *PatchedSourceStageRequest) GetResumeTimeoutOk() (*string, bool)`

GetResumeTimeoutOk returns a tuple with the ResumeTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeTimeout

`func (o *PatchedSourceStageRequest) SetResumeTimeout(v string)`

SetResumeTimeout sets ResumeTimeout field to given value.

### HasResumeTimeout

`func (o *PatchedSourceStageRequest) HasResumeTimeout() bool`

HasResumeTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


