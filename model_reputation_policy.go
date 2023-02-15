/*
authentik

Making authentication simple.

API version: 2023.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReputationPolicy Reputation Policy Serializer
type ReputationPolicy struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging  *bool  `json:"execution_logging,omitempty"`
	Component         string `json:"component"`
	VerboseName       string `json:"verbose_name"`
	VerboseNamePlural string `json:"verbose_name_plural"`
	MetaModelName     string `json:"meta_model_name"`
	BoundTo           int32  `json:"bound_to"`
	CheckIp           *bool  `json:"check_ip,omitempty"`
	CheckUsername     *bool  `json:"check_username,omitempty"`
	Threshold         *int32 `json:"threshold,omitempty"`
}

// NewReputationPolicy instantiates a new ReputationPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReputationPolicy(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, boundTo int32) *ReputationPolicy {
	this := ReputationPolicy{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.BoundTo = boundTo
	return &this
}

// NewReputationPolicyWithDefaults instantiates a new ReputationPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReputationPolicyWithDefaults() *ReputationPolicy {
	this := ReputationPolicy{}
	return &this
}

// GetPk returns the Pk field value
func (o *ReputationPolicy) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *ReputationPolicy) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *ReputationPolicy) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *ReputationPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReputationPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReputationPolicy) SetName(v string) {
	o.Name = v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *ReputationPolicy) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReputationPolicy) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *ReputationPolicy) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *ReputationPolicy) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetComponent returns the Component field value
func (o *ReputationPolicy) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ReputationPolicy) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *ReputationPolicy) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *ReputationPolicy) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *ReputationPolicy) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *ReputationPolicy) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *ReputationPolicy) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *ReputationPolicy) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *ReputationPolicy) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *ReputationPolicy) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *ReputationPolicy) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *ReputationPolicy) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetBoundTo returns the BoundTo field value
func (o *ReputationPolicy) GetBoundTo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BoundTo
}

// GetBoundToOk returns a tuple with the BoundTo field value
// and a boolean to check if the value has been set.
func (o *ReputationPolicy) GetBoundToOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoundTo, true
}

// SetBoundTo sets field value
func (o *ReputationPolicy) SetBoundTo(v int32) {
	o.BoundTo = v
}

// GetCheckIp returns the CheckIp field value if set, zero value otherwise.
func (o *ReputationPolicy) GetCheckIp() bool {
	if o == nil || o.CheckIp == nil {
		var ret bool
		return ret
	}
	return *o.CheckIp
}

// GetCheckIpOk returns a tuple with the CheckIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReputationPolicy) GetCheckIpOk() (*bool, bool) {
	if o == nil || o.CheckIp == nil {
		return nil, false
	}
	return o.CheckIp, true
}

// HasCheckIp returns a boolean if a field has been set.
func (o *ReputationPolicy) HasCheckIp() bool {
	if o != nil && o.CheckIp != nil {
		return true
	}

	return false
}

// SetCheckIp gets a reference to the given bool and assigns it to the CheckIp field.
func (o *ReputationPolicy) SetCheckIp(v bool) {
	o.CheckIp = &v
}

// GetCheckUsername returns the CheckUsername field value if set, zero value otherwise.
func (o *ReputationPolicy) GetCheckUsername() bool {
	if o == nil || o.CheckUsername == nil {
		var ret bool
		return ret
	}
	return *o.CheckUsername
}

// GetCheckUsernameOk returns a tuple with the CheckUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReputationPolicy) GetCheckUsernameOk() (*bool, bool) {
	if o == nil || o.CheckUsername == nil {
		return nil, false
	}
	return o.CheckUsername, true
}

// HasCheckUsername returns a boolean if a field has been set.
func (o *ReputationPolicy) HasCheckUsername() bool {
	if o != nil && o.CheckUsername != nil {
		return true
	}

	return false
}

// SetCheckUsername gets a reference to the given bool and assigns it to the CheckUsername field.
func (o *ReputationPolicy) SetCheckUsername(v bool) {
	o.CheckUsername = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *ReputationPolicy) GetThreshold() int32 {
	if o == nil || o.Threshold == nil {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReputationPolicy) GetThresholdOk() (*int32, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *ReputationPolicy) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *ReputationPolicy) SetThreshold(v int32) {
	o.Threshold = &v
}

func (o ReputationPolicy) MarshalJSON() ([]byte, error) {
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
	if o.CheckIp != nil {
		toSerialize["check_ip"] = o.CheckIp
	}
	if o.CheckUsername != nil {
		toSerialize["check_username"] = o.CheckUsername
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	return json.Marshal(toSerialize)
}

type NullableReputationPolicy struct {
	value *ReputationPolicy
	isSet bool
}

func (v NullableReputationPolicy) Get() *ReputationPolicy {
	return v.value
}

func (v *NullableReputationPolicy) Set(val *ReputationPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableReputationPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableReputationPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReputationPolicy(val *ReputationPolicy) *NullableReputationPolicy {
	return &NullableReputationPolicy{value: val, isSet: true}
}

func (v NullableReputationPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReputationPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
