/*
authentik

Making authentication simple.

API version: 2023.3.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PatchedExpressionPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedExpressionPolicyRequest{}

// PatchedExpressionPolicyRequest Group Membership Policy Serializer
type PatchedExpressionPolicyRequest struct {
	Name *string `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool   `json:"execution_logging,omitempty"`
	Expression       *string `json:"expression,omitempty"`
}

// NewPatchedExpressionPolicyRequest instantiates a new PatchedExpressionPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedExpressionPolicyRequest() *PatchedExpressionPolicyRequest {
	this := PatchedExpressionPolicyRequest{}
	return &this
}

// NewPatchedExpressionPolicyRequestWithDefaults instantiates a new PatchedExpressionPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedExpressionPolicyRequestWithDefaults() *PatchedExpressionPolicyRequest {
	this := PatchedExpressionPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedExpressionPolicyRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExpressionPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedExpressionPolicyRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedExpressionPolicyRequest) SetName(v string) {
	o.Name = &v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *PatchedExpressionPolicyRequest) GetExecutionLogging() bool {
	if o == nil || IsNil(o.ExecutionLogging) {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExpressionPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || IsNil(o.ExecutionLogging) {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *PatchedExpressionPolicyRequest) HasExecutionLogging() bool {
	if o != nil && !IsNil(o.ExecutionLogging) {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *PatchedExpressionPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *PatchedExpressionPolicyRequest) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExpressionPolicyRequest) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *PatchedExpressionPolicyRequest) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *PatchedExpressionPolicyRequest) SetExpression(v string) {
	o.Expression = &v
}

func (o PatchedExpressionPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedExpressionPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ExecutionLogging) {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	return toSerialize, nil
}

type NullablePatchedExpressionPolicyRequest struct {
	value *PatchedExpressionPolicyRequest
	isSet bool
}

func (v NullablePatchedExpressionPolicyRequest) Get() *PatchedExpressionPolicyRequest {
	return v.value
}

func (v *NullablePatchedExpressionPolicyRequest) Set(val *PatchedExpressionPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedExpressionPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedExpressionPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedExpressionPolicyRequest(val *PatchedExpressionPolicyRequest) *NullablePatchedExpressionPolicyRequest {
	return &NullablePatchedExpressionPolicyRequest{value: val, isSet: true}
}

func (v NullablePatchedExpressionPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedExpressionPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
