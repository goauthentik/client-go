/*
authentik

Making authentication simple.

API version: 2024.8.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// PolicyEngineMode the model 'PolicyEngineMode'
type PolicyEngineMode string

// List of PolicyEngineMode
const (
	POLICYENGINEMODE_ALL PolicyEngineMode = "all"
	POLICYENGINEMODE_ANY PolicyEngineMode = "any"
)

// All allowed values of PolicyEngineMode enum
var AllowedPolicyEngineModeEnumValues = []PolicyEngineMode{
	"all",
	"any",
}

func (v *PolicyEngineMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PolicyEngineMode(value)
	for _, existing := range AllowedPolicyEngineModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PolicyEngineMode", value)
}

// NewPolicyEngineModeFromValue returns a pointer to a valid PolicyEngineMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPolicyEngineModeFromValue(v string) (*PolicyEngineMode, error) {
	ev := PolicyEngineMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PolicyEngineMode: valid values are %v", v, AllowedPolicyEngineModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PolicyEngineMode) IsValid() bool {
	for _, existing := range AllowedPolicyEngineModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PolicyEngineMode value
func (v PolicyEngineMode) Ptr() *PolicyEngineMode {
	return &v
}

type NullablePolicyEngineMode struct {
	value *PolicyEngineMode
	isSet bool
}

func (v NullablePolicyEngineMode) Get() *PolicyEngineMode {
	return v.value
}

func (v *NullablePolicyEngineMode) Set(val *PolicyEngineMode) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEngineMode) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEngineMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEngineMode(val *PolicyEngineMode) *NullablePolicyEngineMode {
	return &NullablePolicyEngineMode{value: val, isSet: true}
}

func (v NullablePolicyEngineMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEngineMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
