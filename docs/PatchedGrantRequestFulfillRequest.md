# PatchedGrantRequestFulfillRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to [**RequestStatus**](RequestStatus.md) |  | [optional] 

## Methods

### NewPatchedGrantRequestFulfillRequest

`func NewPatchedGrantRequestFulfillRequest() *PatchedGrantRequestFulfillRequest`

NewPatchedGrantRequestFulfillRequest instantiates a new PatchedGrantRequestFulfillRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedGrantRequestFulfillRequestWithDefaults

`func NewPatchedGrantRequestFulfillRequestWithDefaults() *PatchedGrantRequestFulfillRequest`

NewPatchedGrantRequestFulfillRequestWithDefaults instantiates a new PatchedGrantRequestFulfillRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PatchedGrantRequestFulfillRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PatchedGrantRequestFulfillRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PatchedGrantRequestFulfillRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *PatchedGrantRequestFulfillRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedGrantRequestFulfillRequest) GetStatus() RequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedGrantRequestFulfillRequest) GetStatusOk() (*RequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedGrantRequestFulfillRequest) SetStatus(v RequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedGrantRequestFulfillRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


