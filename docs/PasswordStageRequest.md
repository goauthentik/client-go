# PasswordStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**Backends** | [**[]BackendsEnum**](BackendsEnum.md) | Selection of backends to test the password against. | 
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**FailedAttemptsBeforeCancel** | Pointer to **int32** | How many attempts a user has before the flow is canceled. To lock the user out, use a reputation policy and a user_write stage. | [optional] 

## Methods

### NewPasswordStageRequest

`func NewPasswordStageRequest(name string, backends []BackendsEnum, ) *PasswordStageRequest`

NewPasswordStageRequest instantiates a new PasswordStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordStageRequestWithDefaults

`func NewPasswordStageRequestWithDefaults() *PasswordStageRequest`

NewPasswordStageRequestWithDefaults instantiates a new PasswordStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PasswordStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PasswordStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PasswordStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowSet

`func (o *PasswordStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PasswordStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PasswordStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *PasswordStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetBackends

`func (o *PasswordStageRequest) GetBackends() []BackendsEnum`

GetBackends returns the Backends field if non-nil, zero value otherwise.

### GetBackendsOk

`func (o *PasswordStageRequest) GetBackendsOk() (*[]BackendsEnum, bool)`

GetBackendsOk returns a tuple with the Backends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackends

`func (o *PasswordStageRequest) SetBackends(v []BackendsEnum)`

SetBackends sets Backends field to given value.


### GetConfigureFlow

`func (o *PasswordStageRequest) GetConfigureFlow() string`

GetConfigureFlow returns the ConfigureFlow field if non-nil, zero value otherwise.

### GetConfigureFlowOk

`func (o *PasswordStageRequest) GetConfigureFlowOk() (*string, bool)`

GetConfigureFlowOk returns a tuple with the ConfigureFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureFlow

`func (o *PasswordStageRequest) SetConfigureFlow(v string)`

SetConfigureFlow sets ConfigureFlow field to given value.

### HasConfigureFlow

`func (o *PasswordStageRequest) HasConfigureFlow() bool`

HasConfigureFlow returns a boolean if a field has been set.

### SetConfigureFlowNil

`func (o *PasswordStageRequest) SetConfigureFlowNil(b bool)`

 SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil

### UnsetConfigureFlow
`func (o *PasswordStageRequest) UnsetConfigureFlow()`

UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
### GetFailedAttemptsBeforeCancel

`func (o *PasswordStageRequest) GetFailedAttemptsBeforeCancel() int32`

GetFailedAttemptsBeforeCancel returns the FailedAttemptsBeforeCancel field if non-nil, zero value otherwise.

### GetFailedAttemptsBeforeCancelOk

`func (o *PasswordStageRequest) GetFailedAttemptsBeforeCancelOk() (*int32, bool)`

GetFailedAttemptsBeforeCancelOk returns a tuple with the FailedAttemptsBeforeCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAttemptsBeforeCancel

`func (o *PasswordStageRequest) SetFailedAttemptsBeforeCancel(v int32)`

SetFailedAttemptsBeforeCancel sets FailedAttemptsBeforeCancel field to given value.

### HasFailedAttemptsBeforeCancel

`func (o *PasswordStageRequest) HasFailedAttemptsBeforeCancel() bool`

HasFailedAttemptsBeforeCancel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


