/*
authentik

Making authentication simple.

API version: 2022.1.3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ExpressionPolicyRequest Group Membership Policy Serializer
type ExpressionPolicyRequest struct {
	Name NullableString `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool  `json:"execution_logging,omitempty"`
	Expression       string `json:"expression"`
}

// NewExpressionPolicyRequest instantiates a new ExpressionPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpressionPolicyRequest(expression string) *ExpressionPolicyRequest {
	this := ExpressionPolicyRequest{}
	this.Expression = expression
	return &this
}

// NewExpressionPolicyRequestWithDefaults instantiates a new ExpressionPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpressionPolicyRequestWithDefaults() *ExpressionPolicyRequest {
	this := ExpressionPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpressionPolicyRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpressionPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ExpressionPolicyRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ExpressionPolicyRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ExpressionPolicyRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ExpressionPolicyRequest) UnsetName() {
	o.Name.Unset()
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *ExpressionPolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpressionPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *ExpressionPolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *ExpressionPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetExpression returns the Expression field value
func (o *ExpressionPolicyRequest) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *ExpressionPolicyRequest) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *ExpressionPolicyRequest) SetExpression(v string) {
	o.Expression = v
}

func (o ExpressionPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if true {
		toSerialize["expression"] = o.Expression
	}
	return json.Marshal(toSerialize)
}

type NullableExpressionPolicyRequest struct {
	value *ExpressionPolicyRequest
	isSet bool
}

func (v NullableExpressionPolicyRequest) Get() *ExpressionPolicyRequest {
	return v.value
}

func (v *NullableExpressionPolicyRequest) Set(val *ExpressionPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExpressionPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExpressionPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpressionPolicyRequest(val *ExpressionPolicyRequest) *NullableExpressionPolicyRequest {
	return &NullableExpressionPolicyRequest{value: val, isSet: true}
}

func (v NullableExpressionPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpressionPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
