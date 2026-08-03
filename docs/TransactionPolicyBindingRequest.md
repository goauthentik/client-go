# TransactionPolicyBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to **NullableString** |  | [optional] 
**Group** | Pointer to **NullableString** |  | [optional] 
**User** | Pointer to **NullableInt32** |  | [optional] 
**Negate** | Pointer to **bool** | Negates the outcome of the policy. Messages are unaffected. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Order** | **int32** |  | 
**Timeout** | Pointer to **int32** | Timeout after which Policy execution is terminated. | [optional] 
**FailureResult** | Pointer to **bool** | Result if the Policy execution fails. | [optional] 

## Methods

### NewTransactionPolicyBindingRequest

`func NewTransactionPolicyBindingRequest(order int32, ) *TransactionPolicyBindingRequest`

NewTransactionPolicyBindingRequest instantiates a new TransactionPolicyBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionPolicyBindingRequestWithDefaults

`func NewTransactionPolicyBindingRequestWithDefaults() *TransactionPolicyBindingRequest`

NewTransactionPolicyBindingRequestWithDefaults instantiates a new TransactionPolicyBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *TransactionPolicyBindingRequest) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *TransactionPolicyBindingRequest) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *TransactionPolicyBindingRequest) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *TransactionPolicyBindingRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *TransactionPolicyBindingRequest) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *TransactionPolicyBindingRequest) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetGroup

`func (o *TransactionPolicyBindingRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *TransactionPolicyBindingRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *TransactionPolicyBindingRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *TransactionPolicyBindingRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *TransactionPolicyBindingRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *TransactionPolicyBindingRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetUser

`func (o *TransactionPolicyBindingRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TransactionPolicyBindingRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TransactionPolicyBindingRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *TransactionPolicyBindingRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *TransactionPolicyBindingRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *TransactionPolicyBindingRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetNegate

`func (o *TransactionPolicyBindingRequest) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *TransactionPolicyBindingRequest) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *TransactionPolicyBindingRequest) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *TransactionPolicyBindingRequest) HasNegate() bool`

HasNegate returns a boolean if a field has been set.

### GetEnabled

`func (o *TransactionPolicyBindingRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TransactionPolicyBindingRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TransactionPolicyBindingRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TransactionPolicyBindingRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOrder

`func (o *TransactionPolicyBindingRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TransactionPolicyBindingRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TransactionPolicyBindingRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetTimeout

`func (o *TransactionPolicyBindingRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TransactionPolicyBindingRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TransactionPolicyBindingRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TransactionPolicyBindingRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetFailureResult

`func (o *TransactionPolicyBindingRequest) GetFailureResult() bool`

GetFailureResult returns the FailureResult field if non-nil, zero value otherwise.

### GetFailureResultOk

`func (o *TransactionPolicyBindingRequest) GetFailureResultOk() (*bool, bool)`

GetFailureResultOk returns a tuple with the FailureResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureResult

`func (o *TransactionPolicyBindingRequest) SetFailureResult(v bool)`

SetFailureResult sets FailureResult field to given value.

### HasFailureResult

`func (o *TransactionPolicyBindingRequest) HasFailureResult() bool`

HasFailureResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


