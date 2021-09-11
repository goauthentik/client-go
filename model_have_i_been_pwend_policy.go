/*
authentik

Making authentication simple.

API version: 2021.8.5
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// HaveIBeenPwendPolicy Have I Been Pwned Policy Serializer
type HaveIBeenPwendPolicy struct {
	Pk string `json:"pk"`
	Name NullableString `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	Component string `json:"component"`
	VerboseName string `json:"verbose_name"`
	VerboseNamePlural string `json:"verbose_name_plural"`
	BoundTo int32 `json:"bound_to"`
	// Field key to check, field keys defined in Prompt stages are available.
	PasswordField *string `json:"password_field,omitempty"`
	AllowedCount *int32 `json:"allowed_count,omitempty"`
}

// NewHaveIBeenPwendPolicy instantiates a new HaveIBeenPwendPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHaveIBeenPwendPolicy(pk string, component string, verboseName string, verboseNamePlural string, boundTo int32) *HaveIBeenPwendPolicy {
	this := HaveIBeenPwendPolicy{}
	this.Pk = pk
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.BoundTo = boundTo
	return &this
}

// NewHaveIBeenPwendPolicyWithDefaults instantiates a new HaveIBeenPwendPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHaveIBeenPwendPolicyWithDefaults() *HaveIBeenPwendPolicy {
	this := HaveIBeenPwendPolicy{}
	return &this
}

// GetPk returns the Pk field value
func (o *HaveIBeenPwendPolicy) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *HaveIBeenPwendPolicy) GetPkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *HaveIBeenPwendPolicy) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HaveIBeenPwendPolicy) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HaveIBeenPwendPolicy) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *HaveIBeenPwendPolicy) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *HaveIBeenPwendPolicy) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *HaveIBeenPwendPolicy) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *HaveIBeenPwendPolicy) UnsetName() {
	o.Name.Unset()
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *HaveIBeenPwendPolicy) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaveIBeenPwendPolicy) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *HaveIBeenPwendPolicy) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *HaveIBeenPwendPolicy) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetComponent returns the Component field value
func (o *HaveIBeenPwendPolicy) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *HaveIBeenPwendPolicy) GetComponentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *HaveIBeenPwendPolicy) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *HaveIBeenPwendPolicy) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *HaveIBeenPwendPolicy) GetVerboseNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *HaveIBeenPwendPolicy) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *HaveIBeenPwendPolicy) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *HaveIBeenPwendPolicy) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *HaveIBeenPwendPolicy) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetBoundTo returns the BoundTo field value
func (o *HaveIBeenPwendPolicy) GetBoundTo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BoundTo
}

// GetBoundToOk returns a tuple with the BoundTo field value
// and a boolean to check if the value has been set.
func (o *HaveIBeenPwendPolicy) GetBoundToOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BoundTo, true
}

// SetBoundTo sets field value
func (o *HaveIBeenPwendPolicy) SetBoundTo(v int32) {
	o.BoundTo = v
}

// GetPasswordField returns the PasswordField field value if set, zero value otherwise.
func (o *HaveIBeenPwendPolicy) GetPasswordField() string {
	if o == nil || o.PasswordField == nil {
		var ret string
		return ret
	}
	return *o.PasswordField
}

// GetPasswordFieldOk returns a tuple with the PasswordField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaveIBeenPwendPolicy) GetPasswordFieldOk() (*string, bool) {
	if o == nil || o.PasswordField == nil {
		return nil, false
	}
	return o.PasswordField, true
}

// HasPasswordField returns a boolean if a field has been set.
func (o *HaveIBeenPwendPolicy) HasPasswordField() bool {
	if o != nil && o.PasswordField != nil {
		return true
	}

	return false
}

// SetPasswordField gets a reference to the given string and assigns it to the PasswordField field.
func (o *HaveIBeenPwendPolicy) SetPasswordField(v string) {
	o.PasswordField = &v
}

// GetAllowedCount returns the AllowedCount field value if set, zero value otherwise.
func (o *HaveIBeenPwendPolicy) GetAllowedCount() int32 {
	if o == nil || o.AllowedCount == nil {
		var ret int32
		return ret
	}
	return *o.AllowedCount
}

// GetAllowedCountOk returns a tuple with the AllowedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaveIBeenPwendPolicy) GetAllowedCountOk() (*int32, bool) {
	if o == nil || o.AllowedCount == nil {
		return nil, false
	}
	return o.AllowedCount, true
}

// HasAllowedCount returns a boolean if a field has been set.
func (o *HaveIBeenPwendPolicy) HasAllowedCount() bool {
	if o != nil && o.AllowedCount != nil {
		return true
	}

	return false
}

// SetAllowedCount gets a reference to the given int32 and assigns it to the AllowedCount field.
func (o *HaveIBeenPwendPolicy) SetAllowedCount(v int32) {
	o.AllowedCount = &v
}

func (o HaveIBeenPwendPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["verbose_name"] = o.VerboseName
	}
	if true {
		toSerialize["verbose_name_plural"] = o.VerboseNamePlural
	}
	if true {
		toSerialize["bound_to"] = o.BoundTo
	}
	if o.PasswordField != nil {
		toSerialize["password_field"] = o.PasswordField
	}
	if o.AllowedCount != nil {
		toSerialize["allowed_count"] = o.AllowedCount
	}
	return json.Marshal(toSerialize)
}

type NullableHaveIBeenPwendPolicy struct {
	value *HaveIBeenPwendPolicy
	isSet bool
}

func (v NullableHaveIBeenPwendPolicy) Get() *HaveIBeenPwendPolicy {
	return v.value
}

func (v *NullableHaveIBeenPwendPolicy) Set(val *HaveIBeenPwendPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHaveIBeenPwendPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHaveIBeenPwendPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHaveIBeenPwendPolicy(val *HaveIBeenPwendPolicy) *NullableHaveIBeenPwendPolicy {
	return &NullableHaveIBeenPwendPolicy{value: val, isSet: true}
}

func (v NullableHaveIBeenPwendPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHaveIBeenPwendPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


