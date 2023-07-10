/*
authentik

Making authentication simple.

API version: 2023.6.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// DigestAlgorithmEnum * `http://www.w3.org/2000/09/xmldsig#sha1` - SHA1 * `http://www.w3.org/2001/04/xmlenc#sha256` - SHA256 * `http://www.w3.org/2001/04/xmldsig-more#sha384` - SHA384 * `http://www.w3.org/2001/04/xmlenc#sha512` - SHA512
type DigestAlgorithmEnum string

// List of DigestAlgorithmEnum
const (
	DIGESTALGORITHMENUM__2000_09_XMLDSIGSHA1        DigestAlgorithmEnum = "http://www.w3.org/2000/09/xmldsig#sha1"
	DIGESTALGORITHMENUM__2001_04_XMLENCSHA256       DigestAlgorithmEnum = "http://www.w3.org/2001/04/xmlenc#sha256"
	DIGESTALGORITHMENUM__2001_04_XMLDSIG_MORESHA384 DigestAlgorithmEnum = "http://www.w3.org/2001/04/xmldsig-more#sha384"
	DIGESTALGORITHMENUM__2001_04_XMLENCSHA512       DigestAlgorithmEnum = "http://www.w3.org/2001/04/xmlenc#sha512"
)

// All allowed values of DigestAlgorithmEnum enum
var AllowedDigestAlgorithmEnumEnumValues = []DigestAlgorithmEnum{
	"http://www.w3.org/2000/09/xmldsig#sha1",
	"http://www.w3.org/2001/04/xmlenc#sha256",
	"http://www.w3.org/2001/04/xmldsig-more#sha384",
	"http://www.w3.org/2001/04/xmlenc#sha512",
}

func (v *DigestAlgorithmEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DigestAlgorithmEnum(value)
	for _, existing := range AllowedDigestAlgorithmEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DigestAlgorithmEnum", value)
}

// NewDigestAlgorithmEnumFromValue returns a pointer to a valid DigestAlgorithmEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDigestAlgorithmEnumFromValue(v string) (*DigestAlgorithmEnum, error) {
	ev := DigestAlgorithmEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DigestAlgorithmEnum: valid values are %v", v, AllowedDigestAlgorithmEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DigestAlgorithmEnum) IsValid() bool {
	for _, existing := range AllowedDigestAlgorithmEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DigestAlgorithmEnum value
func (v DigestAlgorithmEnum) Ptr() *DigestAlgorithmEnum {
	return &v
}

type NullableDigestAlgorithmEnum struct {
	value *DigestAlgorithmEnum
	isSet bool
}

func (v NullableDigestAlgorithmEnum) Get() *DigestAlgorithmEnum {
	return v.value
}

func (v *NullableDigestAlgorithmEnum) Set(val *DigestAlgorithmEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableDigestAlgorithmEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableDigestAlgorithmEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigestAlgorithmEnum(val *DigestAlgorithmEnum) *NullableDigestAlgorithmEnum {
	return &NullableDigestAlgorithmEnum{value: val, isSet: true}
}

func (v NullableDigestAlgorithmEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigestAlgorithmEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
