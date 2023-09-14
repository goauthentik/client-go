# PatchedPolicyBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to **NullableString** |  | [optional] 
**Group** | Pointer to **NullableString** |  | [optional] 
**User** | Pointer to **NullableInt32** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**Negate** | Pointer to **bool** | Negates the outcome of the policy. Messages are unaffected. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Timeout** | Pointer to **int32** | Timeout after which Policy execution is terminated. | [optional] 
**FailureResult** | Pointer to **bool** | Result if the Policy execution fails. | [optional] 

## Methods

### NewPatchedPolicyBindingRequest

`func NewPatchedPolicyBindingRequest() *PatchedPolicyBindingRequest`

NewPatchedPolicyBindingRequest instantiates a new PatchedPolicyBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedPolicyBindingRequestWithDefaults

`func NewPatchedPolicyBindingRequestWithDefaults() *PatchedPolicyBindingRequest`

NewPatchedPolicyBindingRequestWithDefaults instantiates a new PatchedPolicyBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *PatchedPolicyBindingRequest) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PatchedPolicyBindingRequest) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PatchedPolicyBindingRequest) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PatchedPolicyBindingRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *PatchedPolicyBindingRequest) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *PatchedPolicyBindingRequest) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetGroup

`func (o *PatchedPolicyBindingRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedPolicyBindingRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedPolicyBindingRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedPolicyBindingRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedPolicyBindingRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedPolicyBindingRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetUser

`func (o *PatchedPolicyBindingRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedPolicyBindingRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedPolicyBindingRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedPolicyBindingRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PatchedPolicyBindingRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PatchedPolicyBindingRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetTarget

`func (o *PatchedPolicyBindingRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PatchedPolicyBindingRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PatchedPolicyBindingRequest) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *PatchedPolicyBindingRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetNegate

`func (o *PatchedPolicyBindingRequest) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *PatchedPolicyBindingRequest) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *PatchedPolicyBindingRequest) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *PatchedPolicyBindingRequest) HasNegate() bool`

HasNegate returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedPolicyBindingRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedPolicyBindingRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedPolicyBindingRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedPolicyBindingRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOrder

`func (o *PatchedPolicyBindingRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PatchedPolicyBindingRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PatchedPolicyBindingRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PatchedPolicyBindingRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetTimeout

`func (o *PatchedPolicyBindingRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PatchedPolicyBindingRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PatchedPolicyBindingRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PatchedPolicyBindingRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetFailureResult

`func (o *PatchedPolicyBindingRequest) GetFailureResult() bool`

GetFailureResult returns the FailureResult field if non-nil, zero value otherwise.

### GetFailureResultOk

`func (o *PatchedPolicyBindingRequest) GetFailureResultOk() (*bool, bool)`

GetFailureResultOk returns a tuple with the FailureResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureResult

`func (o *PatchedPolicyBindingRequest) SetFailureResult(v bool)`

SetFailureResult sets FailureResult field to given value.

### HasFailureResult

`func (o *PatchedPolicyBindingRequest) HasFailureResult() bool`

HasFailureResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


