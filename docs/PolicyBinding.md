# PolicyBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Policy** | Pointer to **NullableString** |  | [optional] 
**Group** | Pointer to **NullableString** |  | [optional] 
**User** | Pointer to **NullableInt32** |  | [optional] 
**PolicyObj** | [**Policy**](Policy.md) |  | [readonly] 
**GroupObj** | [**Group**](Group.md) |  | [readonly] 
**UserObj** | [**User**](User.md) |  | [readonly] 
**Target** | **string** |  | 
**Negate** | Pointer to **bool** | Negates the outcome of the policy. Messages are unaffected. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Order** | **int32** |  | 
**Timeout** | Pointer to **int32** | Timeout after which Policy execution is terminated. | [optional] 
**FailureResult** | Pointer to **bool** | Result if the Policy execution fails. | [optional] 

## Methods

### NewPolicyBinding

`func NewPolicyBinding(pk string, policyObj Policy, groupObj Group, userObj User, target string, order int32, ) *PolicyBinding`

NewPolicyBinding instantiates a new PolicyBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyBindingWithDefaults

`func NewPolicyBindingWithDefaults() *PolicyBinding`

NewPolicyBindingWithDefaults instantiates a new PolicyBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *PolicyBinding) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *PolicyBinding) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *PolicyBinding) SetPk(v string)`

SetPk sets Pk field to given value.


### GetPolicy

`func (o *PolicyBinding) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PolicyBinding) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PolicyBinding) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PolicyBinding) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *PolicyBinding) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *PolicyBinding) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetGroup

`func (o *PolicyBinding) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PolicyBinding) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PolicyBinding) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PolicyBinding) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PolicyBinding) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PolicyBinding) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetUser

`func (o *PolicyBinding) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PolicyBinding) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PolicyBinding) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *PolicyBinding) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PolicyBinding) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PolicyBinding) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPolicyObj

`func (o *PolicyBinding) GetPolicyObj() Policy`

GetPolicyObj returns the PolicyObj field if non-nil, zero value otherwise.

### GetPolicyObjOk

`func (o *PolicyBinding) GetPolicyObjOk() (*Policy, bool)`

GetPolicyObjOk returns a tuple with the PolicyObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObj

`func (o *PolicyBinding) SetPolicyObj(v Policy)`

SetPolicyObj sets PolicyObj field to given value.


### GetGroupObj

`func (o *PolicyBinding) GetGroupObj() Group`

GetGroupObj returns the GroupObj field if non-nil, zero value otherwise.

### GetGroupObjOk

`func (o *PolicyBinding) GetGroupObjOk() (*Group, bool)`

GetGroupObjOk returns a tuple with the GroupObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObj

`func (o *PolicyBinding) SetGroupObj(v Group)`

SetGroupObj sets GroupObj field to given value.


### GetUserObj

`func (o *PolicyBinding) GetUserObj() User`

GetUserObj returns the UserObj field if non-nil, zero value otherwise.

### GetUserObjOk

`func (o *PolicyBinding) GetUserObjOk() (*User, bool)`

GetUserObjOk returns a tuple with the UserObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObj

`func (o *PolicyBinding) SetUserObj(v User)`

SetUserObj sets UserObj field to given value.


### GetTarget

`func (o *PolicyBinding) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PolicyBinding) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PolicyBinding) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetNegate

`func (o *PolicyBinding) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *PolicyBinding) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *PolicyBinding) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *PolicyBinding) HasNegate() bool`

HasNegate returns a boolean if a field has been set.

### GetEnabled

`func (o *PolicyBinding) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PolicyBinding) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PolicyBinding) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PolicyBinding) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOrder

`func (o *PolicyBinding) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PolicyBinding) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PolicyBinding) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetTimeout

`func (o *PolicyBinding) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PolicyBinding) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PolicyBinding) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PolicyBinding) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetFailureResult

`func (o *PolicyBinding) GetFailureResult() bool`

GetFailureResult returns the FailureResult field if non-nil, zero value otherwise.

### GetFailureResultOk

`func (o *PolicyBinding) GetFailureResultOk() (*bool, bool)`

GetFailureResultOk returns a tuple with the FailureResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureResult

`func (o *PolicyBinding) SetFailureResult(v bool)`

SetFailureResult sets FailureResult field to given value.

### HasFailureResult

`func (o *PolicyBinding) HasFailureResult() bool`

HasFailureResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


