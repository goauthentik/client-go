/*
authentik

Making authentication simple.

API version: 2023.10.6
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PolicyBinding PolicyBinding Serializer
type PolicyBinding struct {
	Pk        string         `json:"pk"`
	Policy    NullableString `json:"policy,omitempty"`
	Group     NullableString `json:"group,omitempty"`
	User      NullableInt32  `json:"user,omitempty"`
	PolicyObj Policy         `json:"policy_obj"`
	GroupObj  Group          `json:"group_obj"`
	UserObj   User           `json:"user_obj"`
	Target    string         `json:"target"`
	// Negates the outcome of the policy. Messages are unaffected.
	Negate  *bool `json:"negate,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Order   int32 `json:"order"`
	// Timeout after which Policy execution is terminated.
	Timeout *int32 `json:"timeout,omitempty"`
	// Result if the Policy execution fails.
	FailureResult *bool `json:"failure_result,omitempty"`
}

// NewPolicyBinding instantiates a new PolicyBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyBinding(pk string, policyObj Policy, groupObj Group, userObj User, target string, order int32) *PolicyBinding {
	this := PolicyBinding{}
	this.Pk = pk
	this.PolicyObj = policyObj
	this.GroupObj = groupObj
	this.UserObj = userObj
	this.Target = target
	this.Order = order
	return &this
}

// NewPolicyBindingWithDefaults instantiates a new PolicyBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyBindingWithDefaults() *PolicyBinding {
	this := PolicyBinding{}
	return &this
}

// GetPk returns the Pk field value
func (o *PolicyBinding) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *PolicyBinding) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *PolicyBinding) SetPk(v string) {
	o.Pk = v
}

// GetPolicy returns the Policy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyBinding) GetPolicy() string {
	if o == nil || o.Policy.Get() == nil {
		var ret string
		return ret
	}
	return *o.Policy.Get()
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyBinding) GetPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Policy.Get(), o.Policy.IsSet()
}

// HasPolicy returns a boolean if a field has been set.
func (o *PolicyBinding) HasPolicy() bool {
	if o != nil && o.Policy.IsSet() {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given NullableString and assigns it to the Policy field.
func (o *PolicyBinding) SetPolicy(v string) {
	o.Policy.Set(&v)
}

// SetPolicyNil sets the value for Policy to be an explicit nil
func (o *PolicyBinding) SetPolicyNil() {
	o.Policy.Set(nil)
}

// UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
func (o *PolicyBinding) UnsetPolicy() {
	o.Policy.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyBinding) GetGroup() string {
	if o == nil || o.Group.Get() == nil {
		var ret string
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyBinding) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *PolicyBinding) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableString and assigns it to the Group field.
func (o *PolicyBinding) SetGroup(v string) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil
func (o *PolicyBinding) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *PolicyBinding) UnsetGroup() {
	o.Group.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyBinding) GetUser() int32 {
	if o == nil || o.User.Get() == nil {
		var ret int32
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyBinding) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *PolicyBinding) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableInt32 and assigns it to the User field.
func (o *PolicyBinding) SetUser(v int32) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *PolicyBinding) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *PolicyBinding) UnsetUser() {
	o.User.Unset()
}

// GetPolicyObj returns the PolicyObj field value
func (o *PolicyBinding) GetPolicyObj() Policy {
	if o == nil {
		var ret Policy
		return ret
	}

	return o.PolicyObj
}

// GetPolicyObjOk returns a tuple with the PolicyObj field value
// and a boolean to check if the value has been set.
func (o *PolicyBinding) GetPolicyObjOk() (*Policy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyObj, true
}

// SetPolicyObj sets field value
func (o *PolicyBinding) SetPolicyObj(v Policy) {
	o.PolicyObj = v
}

// GetGroupObj returns the GroupObj field value
func (o *PolicyBinding) GetGroupObj() Group {
	if o == nil {
		var ret Group
		return ret
	}

	return o.GroupObj
}

// GetGroupObjOk returns a tuple with the GroupObj field value
// and a boolean to check if the value has been set.
func (o *PolicyBinding) GetGroupObjOk() (*Group, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupObj, true
}

// SetGroupObj sets field value
func (o *PolicyBinding) SetGroupObj(v Group) {
	o.GroupObj = v
}

// GetUserObj returns the UserObj field value
func (o *PolicyBinding) GetUserObj() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.UserObj
}

// GetUserObjOk returns a tuple with the UserObj field value
// and a boolean to check if the value has been set.
func (o *PolicyBinding) GetUserObjOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserObj, true
}

// SetUserObj sets field value
func (o *PolicyBinding) SetUserObj(v User) {
	o.UserObj = v
}

// GetTarget returns the Target field value
func (o *PolicyBinding) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *PolicyBinding) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *PolicyBinding) SetTarget(v string) {
	o.Target = v
}

// GetNegate returns the Negate field value if set, zero value otherwise.
func (o *PolicyBinding) GetNegate() bool {
	if o == nil || o.Negate == nil {
		var ret bool
		return ret
	}
	return *o.Negate
}

// GetNegateOk returns a tuple with the Negate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyBinding) GetNegateOk() (*bool, bool) {
	if o == nil || o.Negate == nil {
		return nil, false
	}
	return o.Negate, true
}

// HasNegate returns a boolean if a field has been set.
func (o *PolicyBinding) HasNegate() bool {
	if o != nil && o.Negate != nil {
		return true
	}

	return false
}

// SetNegate gets a reference to the given bool and assigns it to the Negate field.
func (o *PolicyBinding) SetNegate(v bool) {
	o.Negate = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PolicyBinding) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyBinding) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PolicyBinding) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PolicyBinding) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOrder returns the Order field value
func (o *PolicyBinding) GetOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *PolicyBinding) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *PolicyBinding) SetOrder(v int32) {
	o.Order = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *PolicyBinding) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyBinding) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *PolicyBinding) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *PolicyBinding) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetFailureResult returns the FailureResult field value if set, zero value otherwise.
func (o *PolicyBinding) GetFailureResult() bool {
	if o == nil || o.FailureResult == nil {
		var ret bool
		return ret
	}
	return *o.FailureResult
}

// GetFailureResultOk returns a tuple with the FailureResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyBinding) GetFailureResultOk() (*bool, bool) {
	if o == nil || o.FailureResult == nil {
		return nil, false
	}
	return o.FailureResult, true
}

// HasFailureResult returns a boolean if a field has been set.
func (o *PolicyBinding) HasFailureResult() bool {
	if o != nil && o.FailureResult != nil {
		return true
	}

	return false
}

// SetFailureResult gets a reference to the given bool and assigns it to the FailureResult field.
func (o *PolicyBinding) SetFailureResult(v bool) {
	o.FailureResult = &v
}

func (o PolicyBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if o.Policy.IsSet() {
		toSerialize["policy"] = o.Policy.Get()
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if true {
		toSerialize["policy_obj"] = o.PolicyObj
	}
	if true {
		toSerialize["group_obj"] = o.GroupObj
	}
	if true {
		toSerialize["user_obj"] = o.UserObj
	}
	if true {
		toSerialize["target"] = o.Target
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

type NullablePolicyBinding struct {
	value *PolicyBinding
	isSet bool
}

func (v NullablePolicyBinding) Get() *PolicyBinding {
	return v.value
}

func (v *NullablePolicyBinding) Set(val *PolicyBinding) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyBinding) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyBinding(val *PolicyBinding) *NullablePolicyBinding {
	return &NullablePolicyBinding{value: val, isSet: true}
}

func (v NullablePolicyBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
