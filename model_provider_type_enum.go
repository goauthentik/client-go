/*
authentik

Making authentication simple.

API version: 2024.12.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ProviderTypeEnum the model 'ProviderTypeEnum'
type ProviderTypeEnum string

// List of ProviderTypeEnum
const (
	PROVIDERTYPEENUM_APPLE         ProviderTypeEnum = "apple"
	PROVIDERTYPEENUM_OPENIDCONNECT ProviderTypeEnum = "openidconnect"
	PROVIDERTYPEENUM_AZUREAD       ProviderTypeEnum = "azuread"
	PROVIDERTYPEENUM_DISCORD       ProviderTypeEnum = "discord"
	PROVIDERTYPEENUM_FACEBOOK      ProviderTypeEnum = "facebook"
	PROVIDERTYPEENUM_GITHUB        ProviderTypeEnum = "github"
	PROVIDERTYPEENUM_GITLAB        ProviderTypeEnum = "gitlab"
	PROVIDERTYPEENUM_GOOGLE        ProviderTypeEnum = "google"
	PROVIDERTYPEENUM_MAILCOW       ProviderTypeEnum = "mailcow"
	PROVIDERTYPEENUM_OKTA          ProviderTypeEnum = "okta"
	PROVIDERTYPEENUM_PATREON       ProviderTypeEnum = "patreon"
	PROVIDERTYPEENUM_REDDIT        ProviderTypeEnum = "reddit"
	PROVIDERTYPEENUM_TWITCH        ProviderTypeEnum = "twitch"
	PROVIDERTYPEENUM_TWITTER       ProviderTypeEnum = "twitter"
)

// All allowed values of ProviderTypeEnum enum
var AllowedProviderTypeEnumEnumValues = []ProviderTypeEnum{
	"apple",
	"openidconnect",
	"azuread",
	"discord",
	"facebook",
	"github",
	"gitlab",
	"google",
	"mailcow",
	"okta",
	"patreon",
	"reddit",
	"twitch",
	"twitter",
}

func (v *ProviderTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProviderTypeEnum(value)
	for _, existing := range AllowedProviderTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProviderTypeEnum", value)
}

// NewProviderTypeEnumFromValue returns a pointer to a valid ProviderTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProviderTypeEnumFromValue(v string) (*ProviderTypeEnum, error) {
	ev := ProviderTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProviderTypeEnum: valid values are %v", v, AllowedProviderTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProviderTypeEnum) IsValid() bool {
	for _, existing := range AllowedProviderTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProviderTypeEnum value
func (v ProviderTypeEnum) Ptr() *ProviderTypeEnum {
	return &v
}

type NullableProviderTypeEnum struct {
	value *ProviderTypeEnum
	isSet bool
}

func (v NullableProviderTypeEnum) Get() *ProviderTypeEnum {
	return v.value
}

func (v *NullableProviderTypeEnum) Set(val *ProviderTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderTypeEnum(val *ProviderTypeEnum) *NullableProviderTypeEnum {
	return &NullableProviderTypeEnum{value: val, isSet: true}
}

func (v NullableProviderTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
