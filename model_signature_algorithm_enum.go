/*
authentik

Making authentication simple.

API version: 2023.10.5
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// SignatureAlgorithmEnum * `http://www.w3.org/2000/09/xmldsig#rsa-sha1` - RSA-SHA1 * `http://www.w3.org/2001/04/xmldsig-more#rsa-sha256` - RSA-SHA256 * `http://www.w3.org/2001/04/xmldsig-more#rsa-sha384` - RSA-SHA384 * `http://www.w3.org/2001/04/xmldsig-more#rsa-sha512` - RSA-SHA512 * `http://www.w3.org/2000/09/xmldsig#dsa-sha1` - DSA-SHA1
type SignatureAlgorithmEnum string

// List of SignatureAlgorithmEnum
const (
	SIGNATUREALGORITHMENUM__2000_09_XMLDSIGRSA_SHA1        SignatureAlgorithmEnum = "http://www.w3.org/2000/09/xmldsig#rsa-sha1"
	SIGNATUREALGORITHMENUM__2001_04_XMLDSIG_MORERSA_SHA256 SignatureAlgorithmEnum = "http://www.w3.org/2001/04/xmldsig-more#rsa-sha256"
	SIGNATUREALGORITHMENUM__2001_04_XMLDSIG_MORERSA_SHA384 SignatureAlgorithmEnum = "http://www.w3.org/2001/04/xmldsig-more#rsa-sha384"
	SIGNATUREALGORITHMENUM__2001_04_XMLDSIG_MORERSA_SHA512 SignatureAlgorithmEnum = "http://www.w3.org/2001/04/xmldsig-more#rsa-sha512"
	SIGNATUREALGORITHMENUM__2000_09_XMLDSIGDSA_SHA1        SignatureAlgorithmEnum = "http://www.w3.org/2000/09/xmldsig#dsa-sha1"
)

// All allowed values of SignatureAlgorithmEnum enum
var AllowedSignatureAlgorithmEnumEnumValues = []SignatureAlgorithmEnum{
	"http://www.w3.org/2000/09/xmldsig#rsa-sha1",
	"http://www.w3.org/2001/04/xmldsig-more#rsa-sha256",
	"http://www.w3.org/2001/04/xmldsig-more#rsa-sha384",
	"http://www.w3.org/2001/04/xmldsig-more#rsa-sha512",
	"http://www.w3.org/2000/09/xmldsig#dsa-sha1",
}

func (v *SignatureAlgorithmEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SignatureAlgorithmEnum(value)
	for _, existing := range AllowedSignatureAlgorithmEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SignatureAlgorithmEnum", value)
}

// NewSignatureAlgorithmEnumFromValue returns a pointer to a valid SignatureAlgorithmEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSignatureAlgorithmEnumFromValue(v string) (*SignatureAlgorithmEnum, error) {
	ev := SignatureAlgorithmEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SignatureAlgorithmEnum: valid values are %v", v, AllowedSignatureAlgorithmEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SignatureAlgorithmEnum) IsValid() bool {
	for _, existing := range AllowedSignatureAlgorithmEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SignatureAlgorithmEnum value
func (v SignatureAlgorithmEnum) Ptr() *SignatureAlgorithmEnum {
	return &v
}

type NullableSignatureAlgorithmEnum struct {
	value *SignatureAlgorithmEnum
	isSet bool
}

func (v NullableSignatureAlgorithmEnum) Get() *SignatureAlgorithmEnum {
	return v.value
}

func (v *NullableSignatureAlgorithmEnum) Set(val *SignatureAlgorithmEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureAlgorithmEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureAlgorithmEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureAlgorithmEnum(val *SignatureAlgorithmEnum) *NullableSignatureAlgorithmEnum {
	return &NullableSignatureAlgorithmEnum{value: val, isSet: true}
}

func (v NullableSignatureAlgorithmEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureAlgorithmEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
