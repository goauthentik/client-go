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

// PasswordExpiryPolicy Password Expiry Policy Serializer
type PasswordExpiryPolicy struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging  *bool  `json:"execution_logging,omitempty"`
	Component         string `json:"component"`
	VerboseName       string `json:"verbose_name"`
	VerboseNamePlural string `json:"verbose_name_plural"`
	MetaModelName     string `json:"meta_model_name"`
	BoundTo           int32  `json:"bound_to"`
	Days              int32  `json:"days"`
	DenyOnly          *bool  `json:"deny_only,omitempty"`
}

// NewPasswordExpiryPolicy instantiates a new PasswordExpiryPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordExpiryPolicy(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, boundTo int32, days int32) *PasswordExpiryPolicy {
	this := PasswordExpiryPolicy{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.BoundTo = boundTo
	this.Days = days
	return &this
}

// NewPasswordExpiryPolicyWithDefaults instantiates a new PasswordExpiryPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordExpiryPolicyWithDefaults() *PasswordExpiryPolicy {
	this := PasswordExpiryPolicy{}
	return &this
}

// GetPk returns the Pk field value
func (o *PasswordExpiryPolicy) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *PasswordExpiryPolicy) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *PasswordExpiryPolicy) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *PasswordExpiryPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PasswordExpiryPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PasswordExpiryPolicy) SetName(v string) {
	o.Name = v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *PasswordExpiryPolicy) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordExpiryPolicy) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *PasswordExpiryPolicy) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *PasswordExpiryPolicy) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetComponent returns the Component field value
func (o *PasswordExpiryPolicy) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *PasswordExpiryPolicy) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *PasswordExpiryPolicy) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *PasswordExpiryPolicy) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *PasswordExpiryPolicy) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *PasswordExpiryPolicy) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *PasswordExpiryPolicy) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *PasswordExpiryPolicy) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *PasswordExpiryPolicy) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *PasswordExpiryPolicy) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *PasswordExpiryPolicy) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *PasswordExpiryPolicy) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetBoundTo returns the BoundTo field value
func (o *PasswordExpiryPolicy) GetBoundTo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BoundTo
}

// GetBoundToOk returns a tuple with the BoundTo field value
// and a boolean to check if the value has been set.
func (o *PasswordExpiryPolicy) GetBoundToOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoundTo, true
}

// SetBoundTo sets field value
func (o *PasswordExpiryPolicy) SetBoundTo(v int32) {
	o.BoundTo = v
}

// GetDays returns the Days field value
func (o *PasswordExpiryPolicy) GetDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Days
}

// GetDaysOk returns a tuple with the Days field value
// and a boolean to check if the value has been set.
func (o *PasswordExpiryPolicy) GetDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Days, true
}

// SetDays sets field value
func (o *PasswordExpiryPolicy) SetDays(v int32) {
	o.Days = v
}

// GetDenyOnly returns the DenyOnly field value if set, zero value otherwise.
func (o *PasswordExpiryPolicy) GetDenyOnly() bool {
	if o == nil || o.DenyOnly == nil {
		var ret bool
		return ret
	}
	return *o.DenyOnly
}

// GetDenyOnlyOk returns a tuple with the DenyOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordExpiryPolicy) GetDenyOnlyOk() (*bool, bool) {
	if o == nil || o.DenyOnly == nil {
		return nil, false
	}
	return o.DenyOnly, true
}

// HasDenyOnly returns a boolean if a field has been set.
func (o *PasswordExpiryPolicy) HasDenyOnly() bool {
	if o != nil && o.DenyOnly != nil {
		return true
	}

	return false
}

// SetDenyOnly gets a reference to the given bool and assigns it to the DenyOnly field.
func (o *PasswordExpiryPolicy) SetDenyOnly(v bool) {
	o.DenyOnly = &v
}

func (o PasswordExpiryPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
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
		toSerialize["meta_model_name"] = o.MetaModelName
	}
	if true {
		toSerialize["bound_to"] = o.BoundTo
	}
	if true {
		toSerialize["days"] = o.Days
	}
	if o.DenyOnly != nil {
		toSerialize["deny_only"] = o.DenyOnly
	}
	return json.Marshal(toSerialize)
}

type NullablePasswordExpiryPolicy struct {
	value *PasswordExpiryPolicy
	isSet bool
}

func (v NullablePasswordExpiryPolicy) Get() *PasswordExpiryPolicy {
	return v.value
}

func (v *NullablePasswordExpiryPolicy) Set(val *PasswordExpiryPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordExpiryPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordExpiryPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordExpiryPolicy(val *PasswordExpiryPolicy) *NullablePasswordExpiryPolicy {
	return &NullablePasswordExpiryPolicy{value: val, isSet: true}
}

func (v NullablePasswordExpiryPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordExpiryPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
