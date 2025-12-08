# PatchedDeviceUserBindingRequest

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
**IsPrimary** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedDeviceUserBindingRequest

`func NewPatchedDeviceUserBindingRequest() *PatchedDeviceUserBindingRequest`

NewPatchedDeviceUserBindingRequest instantiates a new PatchedDeviceUserBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDeviceUserBindingRequestWithDefaults

`func NewPatchedDeviceUserBindingRequestWithDefaults() *PatchedDeviceUserBindingRequest`

NewPatchedDeviceUserBindingRequestWithDefaults instantiates a new PatchedDeviceUserBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *PatchedDeviceUserBindingRequest) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PatchedDeviceUserBindingRequest) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PatchedDeviceUserBindingRequest) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PatchedDeviceUserBindingRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *PatchedDeviceUserBindingRequest) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *PatchedDeviceUserBindingRequest) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetGroup

`func (o *PatchedDeviceUserBindingRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedDeviceUserBindingRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedDeviceUserBindingRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedDeviceUserBindingRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedDeviceUserBindingRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedDeviceUserBindingRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetUser

`func (o *PatchedDeviceUserBindingRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedDeviceUserBindingRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedDeviceUserBindingRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedDeviceUserBindingRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PatchedDeviceUserBindingRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PatchedDeviceUserBindingRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetTarget

`func (o *PatchedDeviceUserBindingRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PatchedDeviceUserBindingRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PatchedDeviceUserBindingRequest) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *PatchedDeviceUserBindingRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetNegate

`func (o *PatchedDeviceUserBindingRequest) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *PatchedDeviceUserBindingRequest) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *PatchedDeviceUserBindingRequest) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *PatchedDeviceUserBindingRequest) HasNegate() bool`

HasNegate returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedDeviceUserBindingRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedDeviceUserBindingRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedDeviceUserBindingRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedDeviceUserBindingRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOrder

`func (o *PatchedDeviceUserBindingRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PatchedDeviceUserBindingRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PatchedDeviceUserBindingRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PatchedDeviceUserBindingRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetTimeout

`func (o *PatchedDeviceUserBindingRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PatchedDeviceUserBindingRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PatchedDeviceUserBindingRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PatchedDeviceUserBindingRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetFailureResult

`func (o *PatchedDeviceUserBindingRequest) GetFailureResult() bool`

GetFailureResult returns the FailureResult field if non-nil, zero value otherwise.

### GetFailureResultOk

`func (o *PatchedDeviceUserBindingRequest) GetFailureResultOk() (*bool, bool)`

GetFailureResultOk returns a tuple with the FailureResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureResult

`func (o *PatchedDeviceUserBindingRequest) SetFailureResult(v bool)`

SetFailureResult sets FailureResult field to given value.

### HasFailureResult

`func (o *PatchedDeviceUserBindingRequest) HasFailureResult() bool`

HasFailureResult returns a boolean if a field has been set.

### GetIsPrimary

`func (o *PatchedDeviceUserBindingRequest) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *PatchedDeviceUserBindingRequest) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *PatchedDeviceUserBindingRequest) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *PatchedDeviceUserBindingRequest) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


