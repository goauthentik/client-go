/*
authentik

Making authentication simple.

API version: 2024.8.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// NameIdPolicyEnum the model 'NameIdPolicyEnum'
type NameIdPolicyEnum string

// List of NameIdPolicyEnum
const (
	NAMEIDPOLICYENUM__1_1NAMEID_FORMATEMAIL_ADDRESS                 NameIdPolicyEnum = "urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress"
	NAMEIDPOLICYENUM__2_0NAMEID_FORMATPERSISTENT                    NameIdPolicyEnum = "urn:oasis:names:tc:SAML:2.0:nameid-format:persistent"
	NAMEIDPOLICYENUM__2_0NAMEID_FORMATX509_SUBJECT_NAME             NameIdPolicyEnum = "urn:oasis:names:tc:SAML:2.0:nameid-format:X509SubjectName"
	NAMEIDPOLICYENUM__2_0NAMEID_FORMATWINDOWS_DOMAIN_QUALIFIED_NAME NameIdPolicyEnum = "urn:oasis:names:tc:SAML:2.0:nameid-format:WindowsDomainQualifiedName"
	NAMEIDPOLICYENUM__2_0NAMEID_FORMATTRANSIENT                     NameIdPolicyEnum = "urn:oasis:names:tc:SAML:2.0:nameid-format:transient"
)

// All allowed values of NameIdPolicyEnum enum
var AllowedNameIdPolicyEnumEnumValues = []NameIdPolicyEnum{
	"urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress",
	"urn:oasis:names:tc:SAML:2.0:nameid-format:persistent",
	"urn:oasis:names:tc:SAML:2.0:nameid-format:X509SubjectName",
	"urn:oasis:names:tc:SAML:2.0:nameid-format:WindowsDomainQualifiedName",
	"urn:oasis:names:tc:SAML:2.0:nameid-format:transient",
}

func (v *NameIdPolicyEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NameIdPolicyEnum(value)
	for _, existing := range AllowedNameIdPolicyEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NameIdPolicyEnum", value)
}

// NewNameIdPolicyEnumFromValue returns a pointer to a valid NameIdPolicyEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNameIdPolicyEnumFromValue(v string) (*NameIdPolicyEnum, error) {
	ev := NameIdPolicyEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NameIdPolicyEnum: valid values are %v", v, AllowedNameIdPolicyEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NameIdPolicyEnum) IsValid() bool {
	for _, existing := range AllowedNameIdPolicyEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NameIdPolicyEnum value
func (v NameIdPolicyEnum) Ptr() *NameIdPolicyEnum {
	return &v
}

type NullableNameIdPolicyEnum struct {
	value *NameIdPolicyEnum
	isSet bool
}

func (v NullableNameIdPolicyEnum) Get() *NameIdPolicyEnum {
	return v.value
}

func (v *NullableNameIdPolicyEnum) Set(val *NameIdPolicyEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableNameIdPolicyEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableNameIdPolicyEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameIdPolicyEnum(val *NameIdPolicyEnum) *NullableNameIdPolicyEnum {
	return &NullableNameIdPolicyEnum{value: val, isSet: true}
}

func (v NullableNameIdPolicyEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameIdPolicyEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
