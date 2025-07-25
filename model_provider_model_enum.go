/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ProviderModelEnum the model 'ProviderModelEnum'
type ProviderModelEnum string

// List of ProviderModelEnum
const (
	PROVIDERMODELENUM_GOOGLE_WORKSPACE_GOOGLEWORKSPACEPROVIDER ProviderModelEnum = "authentik_providers_google_workspace.googleworkspaceprovider"
	PROVIDERMODELENUM_LDAP_LDAPPROVIDER                        ProviderModelEnum = "authentik_providers_ldap.ldapprovider"
	PROVIDERMODELENUM_MICROSOFT_ENTRA_MICROSOFTENTRAPROVIDER   ProviderModelEnum = "authentik_providers_microsoft_entra.microsoftentraprovider"
	PROVIDERMODELENUM_OAUTH2_OAUTH2PROVIDER                    ProviderModelEnum = "authentik_providers_oauth2.oauth2provider"
	PROVIDERMODELENUM_PROXY_PROXYPROVIDER                      ProviderModelEnum = "authentik_providers_proxy.proxyprovider"
	PROVIDERMODELENUM_RAC_RACPROVIDER                          ProviderModelEnum = "authentik_providers_rac.racprovider"
	PROVIDERMODELENUM_RADIUS_RADIUSPROVIDER                    ProviderModelEnum = "authentik_providers_radius.radiusprovider"
	PROVIDERMODELENUM_SAML_SAMLPROVIDER                        ProviderModelEnum = "authentik_providers_saml.samlprovider"
	PROVIDERMODELENUM_SCIM_SCIMPROVIDER                        ProviderModelEnum = "authentik_providers_scim.scimprovider"
	PROVIDERMODELENUM_SSF_SSFPROVIDER                          ProviderModelEnum = "authentik_providers_ssf.ssfprovider"
)

// All allowed values of ProviderModelEnum enum
var AllowedProviderModelEnumEnumValues = []ProviderModelEnum{
	"authentik_providers_google_workspace.googleworkspaceprovider",
	"authentik_providers_ldap.ldapprovider",
	"authentik_providers_microsoft_entra.microsoftentraprovider",
	"authentik_providers_oauth2.oauth2provider",
	"authentik_providers_proxy.proxyprovider",
	"authentik_providers_rac.racprovider",
	"authentik_providers_radius.radiusprovider",
	"authentik_providers_saml.samlprovider",
	"authentik_providers_scim.scimprovider",
	"authentik_providers_ssf.ssfprovider",
}

func (v *ProviderModelEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProviderModelEnum(value)
	for _, existing := range AllowedProviderModelEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProviderModelEnum", value)
}

// NewProviderModelEnumFromValue returns a pointer to a valid ProviderModelEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProviderModelEnumFromValue(v string) (*ProviderModelEnum, error) {
	ev := ProviderModelEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProviderModelEnum: valid values are %v", v, AllowedProviderModelEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProviderModelEnum) IsValid() bool {
	for _, existing := range AllowedProviderModelEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProviderModelEnum value
func (v ProviderModelEnum) Ptr() *ProviderModelEnum {
	return &v
}

type NullableProviderModelEnum struct {
	value *ProviderModelEnum
	isSet bool
}

func (v NullableProviderModelEnum) Get() *ProviderModelEnum {
	return v.value
}

func (v *NullableProviderModelEnum) Set(val *ProviderModelEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderModelEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderModelEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderModelEnum(val *ProviderModelEnum) *NullableProviderModelEnum {
	return &NullableProviderModelEnum{value: val, isSet: true}
}

func (v NullableProviderModelEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderModelEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
