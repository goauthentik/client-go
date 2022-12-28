/*
authentik

Making authentication simple.

API version: 2022.12.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedPolicyBindingRequest PolicyBinding Serializer
type PatchedPolicyBindingRequest struct {
	Policy NullableString `json:"policy,omitempty"`
	Group  NullableString `json:"group,omitempty"`
	User   NullableInt32  `json:"user,omitempty"`
	Target *string        `json:"target,omitempty"`
	// Negates the outcome of the policy. Messages are unaffected.
	Negate  *bool  `json:"negate,omitempty"`
	Enabled *bool  `json:"enabled,omitempty"`
	Order   *int32 `json:"order,omitempty"`
	// Timeout after which Policy execution is terminated.
	Timeout *int32 `json:"timeout,omitempty"`
}

// NewPatchedPolicyBindingRequest instantiates a new PatchedPolicyBindingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedPolicyBindingRequest() *PatchedPolicyBindingRequest {
	this := PatchedPolicyBindingRequest{}
	return &this
}

// NewPatchedPolicyBindingRequestWithDefaults instantiates a new PatchedPolicyBindingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedPolicyBindingRequestWithDefaults() *PatchedPolicyBindingRequest {
	this := PatchedPolicyBindingRequest{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedPolicyBindingRequest) GetPolicy() string {
	if o == nil || o.Policy.Get() == nil {
		var ret string
		return ret
	}
	return *o.Policy.Get()
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedPolicyBindingRequest) GetPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Policy.Get(), o.Policy.IsSet()
}

// HasPolicy returns a boolean if a field has been set.
func (o *PatchedPolicyBindingRequest) HasPolicy() bool {
	if o != nil && o.Policy.IsSet() {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given NullableString and assigns it to the Policy field.
func (o *PatchedPolicyBindingRequest) SetPolicy(v string) {
	o.Policy.Set(&v)
}

// SetPolicyNil sets the value for Policy to be an explicit nil
func (o *PatchedPolicyBindingRequest) SetPolicyNil() {
	o.Policy.Set(nil)
}

// UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
func (o *PatchedPolicyBindingRequest) UnsetPolicy() {
	o.Policy.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedPolicyBindingRequest) GetGroup() string {
	if o == nil || o.Group.Get() == nil {
		var ret string
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedPolicyBindingRequest) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *PatchedPolicyBindingRequest) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableString and assigns it to the Group field.
func (o *PatchedPolicyBindingRequest) SetGroup(v string) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil
func (o *PatchedPolicyBindingRequest) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *PatchedPolicyBindingRequest) UnsetGroup() {
	o.Group.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedPolicyBindingRequest) GetUser() int32 {
	if o == nil || o.User.Get() == nil {
		var ret int32
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedPolicyBindingRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedPolicyBindingRequest) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableInt32 and assigns it to the User field.
func (o *PatchedPolicyBindingRequest) SetUser(v int32) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *PatchedPolicyBindingRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *PatchedPolicyBindingRequest) UnsetUser() {
	o.User.Unset()
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *PatchedPolicyBindingRequest) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPolicyBindingRequest) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *PatchedPolicyBindingRequest) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *PatchedPolicyBindingRequest) SetTarget(v string) {
	o.Target = &v
}

// GetNegate returns the Negate field value if set, zero value otherwise.
func (o *PatchedPolicyBindingRequest) GetNegate() bool {
	if o == nil || o.Negate == nil {
		var ret bool
		return ret
	}
	return *o.Negate
}

// GetNegateOk returns a tuple with the Negate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPolicyBindingRequest) GetNegateOk() (*bool, bool) {
	if o == nil || o.Negate == nil {
		return nil, false
	}
	return o.Negate, true
}

// HasNegate returns a boolean if a field has been set.
func (o *PatchedPolicyBindingRequest) HasNegate() bool {
	if o != nil && o.Negate != nil {
		return true
	}

	return false
}

// SetNegate gets a reference to the given bool and assigns it to the Negate field.
func (o *PatchedPolicyBindingRequest) SetNegate(v bool) {
	o.Negate = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchedPolicyBindingRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPolicyBindingRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchedPolicyBindingRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchedPolicyBindingRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *PatchedPolicyBindingRequest) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPolicyBindingRequest) GetOrderOk() (*int32, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *PatchedPolicyBindingRequest) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *PatchedPolicyBindingRequest) SetOrder(v int32) {
	o.Order = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *PatchedPolicyBindingRequest) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPolicyBindingRequest) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *PatchedPolicyBindingRequest) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *PatchedPolicyBindingRequest) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o PatchedPolicyBindingRequest) MarshalJSON() ([]byte, error) {
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
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Negate != nil {
		toSerialize["negate"] = o.Negate
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedPolicyBindingRequest struct {
	value *PatchedPolicyBindingRequest
	isSet bool
}

func (v NullablePatchedPolicyBindingRequest) Get() *PatchedPolicyBindingRequest {
	return v.value
}

func (v *NullablePatchedPolicyBindingRequest) Set(val *PatchedPolicyBindingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedPolicyBindingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedPolicyBindingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedPolicyBindingRequest(val *PatchedPolicyBindingRequest) *NullablePatchedPolicyBindingRequest {
	return &NullablePatchedPolicyBindingRequest{value: val, isSet: true}
}

func (v NullablePatchedPolicyBindingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedPolicyBindingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
