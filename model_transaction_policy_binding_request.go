/*
authentik

Making authentication simple.

API version: 2024.12.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TransactionPolicyBindingRequest PolicyBindingSerializer which does not require target as target is set implicitly
type TransactionPolicyBindingRequest struct {
	Policy NullableString `json:"policy,omitempty"`
	Group  NullableString `json:"group,omitempty"`
	User   NullableInt32  `json:"user,omitempty"`
	// Negates the outcome of the policy. Messages are unaffected.
	Negate  *bool `json:"negate,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Order   int32 `json:"order"`
	// Timeout after which Policy execution is terminated.
	Timeout *int32 `json:"timeout,omitempty"`
	// Result if the Policy execution fails.
	FailureResult *bool `json:"failure_result,omitempty"`
}

// NewTransactionPolicyBindingRequest instantiates a new TransactionPolicyBindingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionPolicyBindingRequest(order int32) *TransactionPolicyBindingRequest {
	this := TransactionPolicyBindingRequest{}
	this.Order = order
	return &this
}

// NewTransactionPolicyBindingRequestWithDefaults instantiates a new TransactionPolicyBindingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionPolicyBindingRequestWithDefaults() *TransactionPolicyBindingRequest {
	this := TransactionPolicyBindingRequest{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionPolicyBindingRequest) GetPolicy() string {
	if o == nil || o.Policy.Get() == nil {
		var ret string
		return ret
	}
	return *o.Policy.Get()
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionPolicyBindingRequest) GetPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Policy.Get(), o.Policy.IsSet()
}

// HasPolicy returns a boolean if a field has been set.
func (o *TransactionPolicyBindingRequest) HasPolicy() bool {
	if o != nil && o.Policy.IsSet() {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given NullableString and assigns it to the Policy field.
func (o *TransactionPolicyBindingRequest) SetPolicy(v string) {
	o.Policy.Set(&v)
}

// SetPolicyNil sets the value for Policy to be an explicit nil
func (o *TransactionPolicyBindingRequest) SetPolicyNil() {
	o.Policy.Set(nil)
}

// UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
func (o *TransactionPolicyBindingRequest) UnsetPolicy() {
	o.Policy.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionPolicyBindingRequest) GetGroup() string {
	if o == nil || o.Group.Get() == nil {
		var ret string
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionPolicyBindingRequest) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *TransactionPolicyBindingRequest) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableString and assigns it to the Group field.
func (o *TransactionPolicyBindingRequest) SetGroup(v string) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil
func (o *TransactionPolicyBindingRequest) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *TransactionPolicyBindingRequest) UnsetGroup() {
	o.Group.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionPolicyBindingRequest) GetUser() int32 {
	if o == nil || o.User.Get() == nil {
		var ret int32
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionPolicyBindingRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *TransactionPolicyBindingRequest) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableInt32 and assigns it to the User field.
func (o *TransactionPolicyBindingRequest) SetUser(v int32) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *TransactionPolicyBindingRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *TransactionPolicyBindingRequest) UnsetUser() {
	o.User.Unset()
}

// GetNegate returns the Negate field value if set, zero value otherwise.
func (o *TransactionPolicyBindingRequest) GetNegate() bool {
	if o == nil || o.Negate == nil {
		var ret bool
		return ret
	}
	return *o.Negate
}

// GetNegateOk returns a tuple with the Negate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPolicyBindingRequest) GetNegateOk() (*bool, bool) {
	if o == nil || o.Negate == nil {
		return nil, false
	}
	return o.Negate, true
}

// HasNegate returns a boolean if a field has been set.
func (o *TransactionPolicyBindingRequest) HasNegate() bool {
	if o != nil && o.Negate != nil {
		return true
	}

	return false
}

// SetNegate gets a reference to the given bool and assigns it to the Negate field.
func (o *TransactionPolicyBindingRequest) SetNegate(v bool) {
	o.Negate = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TransactionPolicyBindingRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPolicyBindingRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TransactionPolicyBindingRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TransactionPolicyBindingRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOrder returns the Order field value
func (o *TransactionPolicyBindingRequest) GetOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *TransactionPolicyBindingRequest) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *TransactionPolicyBindingRequest) SetOrder(v int32) {
	o.Order = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *TransactionPolicyBindingRequest) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPolicyBindingRequest) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *TransactionPolicyBindingRequest) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *TransactionPolicyBindingRequest) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetFailureResult returns the FailureResult field value if set, zero value otherwise.
func (o *TransactionPolicyBindingRequest) GetFailureResult() bool {
	if o == nil || o.FailureResult == nil {
		var ret bool
		return ret
	}
	return *o.FailureResult
}

// GetFailureResultOk returns a tuple with the FailureResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPolicyBindingRequest) GetFailureResultOk() (*bool, bool) {
	if o == nil || o.FailureResult == nil {
		return nil, false
	}
	return o.FailureResult, true
}

// HasFailureResult returns a boolean if a field has been set.
func (o *TransactionPolicyBindingRequest) HasFailureResult() bool {
	if o != nil && o.FailureResult != nil {
		return true
	}

	return false
}

// SetFailureResult gets a reference to the given bool and assigns it to the FailureResult field.
func (o *TransactionPolicyBindingRequest) SetFailureResult(v bool) {
	o.FailureResult = &v
}

func (o TransactionPolicyBindingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Policy.IsSet() {
		toSerialize["policy"] = o.Policy.Get()
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.Negate != nil {
		toSerialize["negate"] = o.Negate
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["order"] = o.Order
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.FailureResult != nil {
		toSerialize["failure_result"] = o.FailureResult
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionPolicyBindingRequest struct {
	value *TransactionPolicyBindingRequest
	isSet bool
}

func (v NullableTransactionPolicyBindingRequest) Get() *TransactionPolicyBindingRequest {
	return v.value
}

func (v *NullableTransactionPolicyBindingRequest) Set(val *TransactionPolicyBindingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionPolicyBindingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionPolicyBindingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionPolicyBindingRequest(val *TransactionPolicyBindingRequest) *NullableTransactionPolicyBindingRequest {
	return &NullableTransactionPolicyBindingRequest{value: val, isSet: true}
}

func (v NullableTransactionPolicyBindingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionPolicyBindingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
