# PolicyBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to **NullableString** |  | [optional] 
**Group** | Pointer to **NullableString** |  | [optional] 
**User** | Pointer to **NullableInt32** |  | [optional] 
**Target** | **string** |  | 
**Negate** | Pointer to **bool** | Negates the outcome of the policy. Messages are unaffected. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Order** | **int32** |  | 
**Timeout** | Pointer to **int32** | Timeout after which Policy execution is terminated. | [optional] 
**FailureResult** | Pointer to **bool** | Result if the Policy execution fails. | [optional] 

## Methods

### NewPolicyBindingRequest

`func NewPolicyBindingRequest(target string, order int32, ) *PolicyBindingRequest`

NewPolicyBindingRequest instantiates a new PolicyBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyBindingRequestWithDefaults

`func NewPolicyBindingRequestWithDefaults() *PolicyBindingRequest`

NewPolicyBindingRequestWithDefaults instantiates a new PolicyBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *PolicyBindingRequest) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PolicyBindingRequest) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PolicyBindingRequest) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PolicyBindingRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *PolicyBindingRequest) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *PolicyBindingRequest) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetGroup

`func (o *PolicyBindingRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PolicyBindingRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PolicyBindingRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PolicyBindingRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PolicyBindingRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PolicyBindingRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetUser

`func (o *PolicyBindingRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PolicyBindingRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PolicyBindingRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *PolicyBindingRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PolicyBindingRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PolicyBindingRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetTarget

`func (o *PolicyBindingRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PolicyBindingRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PolicyBindingRequest) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetNegate

`func (o *PolicyBindingRequest) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *PolicyBindingRequest) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *PolicyBindingRequest) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *PolicyBindingRequest) HasNegate() bool`

HasNegate returns a boolean if a field has been set.

### GetEnabled

`func (o *PolicyBindingRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PolicyBindingRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PolicyBindingRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PolicyBindingRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOrder

`func (o *PolicyBindingRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PolicyBindingRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PolicyBindingRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetTimeout

`func (o *PolicyBindingRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PolicyBindingRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PolicyBindingRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PolicyBindingRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetFailureResult

`func (o *PolicyBindingRequest) GetFailureResult() bool`

GetFailureResult returns the FailureResult field if non-nil, zero value otherwise.

### GetFailureResultOk

`func (o *PolicyBindingRequest) GetFailureResultOk() (*bool, bool)`

GetFailureResultOk returns a tuple with the FailureResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureResult

`func (o *PolicyBindingRequest) SetFailureResult(v bool)`

SetFailureResult sets FailureResult field to given value.

### HasFailureResult

`func (o *PolicyBindingRequest) HasFailureResult() bool`

HasFailureResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


