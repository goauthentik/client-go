# PatchedPasswordStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**Backends** | Pointer to [**[]BackendsEnum**](BackendsEnum.md) | Selection of backends to test the password against. | [optional] 
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**FailedAttemptsBeforeCancel** | Pointer to **int32** | How many attempts a user has before the flow is canceled. To lock the user out, use a reputation policy and a user_write stage. | [optional] 

## Methods

### NewPatchedPasswordStageRequest

`func NewPatchedPasswordStageRequest() *PatchedPasswordStageRequest`

NewPatchedPasswordStageRequest instantiates a new PatchedPasswordStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedPasswordStageRequestWithDefaults

`func NewPatchedPasswordStageRequestWithDefaults() *PatchedPasswordStageRequest`

NewPatchedPasswordStageRequestWithDefaults instantiates a new PatchedPasswordStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedPasswordStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedPasswordStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedPasswordStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedPasswordStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlowSet

`func (o *PatchedPasswordStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PatchedPasswordStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PatchedPasswordStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *PatchedPasswordStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetBackends

`func (o *PatchedPasswordStageRequest) GetBackends() []BackendsEnum`

GetBackends returns the Backends field if non-nil, zero value otherwise.

### GetBackendsOk

`func (o *PatchedPasswordStageRequest) GetBackendsOk() (*[]BackendsEnum, bool)`

GetBackendsOk returns a tuple with the Backends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackends

`func (o *PatchedPasswordStageRequest) SetBackends(v []BackendsEnum)`

SetBackends sets Backends field to given value.

### HasBackends

`func (o *PatchedPasswordStageRequest) HasBackends() bool`

HasBackends returns a boolean if a field has been set.

### GetConfigureFlow

`func (o *PatchedPasswordStageRequest) GetConfigureFlow() string`

GetConfigureFlow returns the ConfigureFlow field if non-nil, zero value otherwise.

### GetConfigureFlowOk

`func (o *PatchedPasswordStageRequest) GetConfigureFlowOk() (*string, bool)`

GetConfigureFlowOk returns a tuple with the ConfigureFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureFlow

`func (o *PatchedPasswordStageRequest) SetConfigureFlow(v string)`

SetConfigureFlow sets ConfigureFlow field to given value.

### HasConfigureFlow

`func (o *PatchedPasswordStageRequest) HasConfigureFlow() bool`

HasConfigureFlow returns a boolean if a field has been set.

### SetConfigureFlowNil

`func (o *PatchedPasswordStageRequest) SetConfigureFlowNil(b bool)`

 SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil

### UnsetConfigureFlow
`func (o *PatchedPasswordStageRequest) UnsetConfigureFlow()`

UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
### GetFailedAttemptsBeforeCancel

`func (o *PatchedPasswordStageRequest) GetFailedAttemptsBeforeCancel() int32`

GetFailedAttemptsBeforeCancel returns the FailedAttemptsBeforeCancel field if non-nil, zero value otherwise.

### GetFailedAttemptsBeforeCancelOk

`func (o *PatchedPasswordStageRequest) GetFailedAttemptsBeforeCancelOk() (*int32, bool)`

GetFailedAttemptsBeforeCancelOk returns a tuple with the FailedAttemptsBeforeCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAttemptsBeforeCancel

`func (o *PatchedPasswordStageRequest) SetFailedAttemptsBeforeCancel(v int32)`

SetFailedAttemptsBeforeCancel sets FailedAttemptsBeforeCancel field to given value.

### HasFailedAttemptsBeforeCancel

`func (o *PatchedPasswordStageRequest) HasFailedAttemptsBeforeCancel() bool`

HasFailedAttemptsBeforeCancel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


