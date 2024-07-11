/*
authentik

Making authentication simple.

API version: 2024.6.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ModelRequest - struct for ModelRequest
type ModelRequest struct {
	GoogleWorkspaceProviderRequest *GoogleWorkspaceProviderRequest
	LDAPProviderRequest            *LDAPProviderRequest
	MicrosoftEntraProviderRequest  *MicrosoftEntraProviderRequest
	OAuth2ProviderRequest          *OAuth2ProviderRequest
	ProxyProviderRequest           *ProxyProviderRequest
	RACProviderRequest             *RACProviderRequest
	RadiusProviderRequest          *RadiusProviderRequest
	SAMLProviderRequest            *SAMLProviderRequest
	SCIMProviderRequest            *SCIMProviderRequest
}

// GoogleWorkspaceProviderRequestAsModelRequest is a convenience function that returns GoogleWorkspaceProviderRequest wrapped in ModelRequest
func GoogleWorkspaceProviderRequestAsModelRequest(v *GoogleWorkspaceProviderRequest) ModelRequest {
	return ModelRequest{
		GoogleWorkspaceProviderRequest: v,
	}
}

// LDAPProviderRequestAsModelRequest is a convenience function that returns LDAPProviderRequest wrapped in ModelRequest
func LDAPProviderRequestAsModelRequest(v *LDAPProviderRequest) ModelRequest {
	return ModelRequest{
		LDAPProviderRequest: v,
	}
}

// MicrosoftEntraProviderRequestAsModelRequest is a convenience function that returns MicrosoftEntraProviderRequest wrapped in ModelRequest
func MicrosoftEntraProviderRequestAsModelRequest(v *MicrosoftEntraProviderRequest) ModelRequest {
	return ModelRequest{
		MicrosoftEntraProviderRequest: v,
	}
}

// OAuth2ProviderRequestAsModelRequest is a convenience function that returns OAuth2ProviderRequest wrapped in ModelRequest
func OAuth2ProviderRequestAsModelRequest(v *OAuth2ProviderRequest) ModelRequest {
	return ModelRequest{
		OAuth2ProviderRequest: v,
	}
}

// ProxyProviderRequestAsModelRequest is a convenience function that returns ProxyProviderRequest wrapped in ModelRequest
func ProxyProviderRequestAsModelRequest(v *ProxyProviderRequest) ModelRequest {
	return ModelRequest{
		ProxyProviderRequest: v,
	}
}

// RACProviderRequestAsModelRequest is a convenience function that returns RACProviderRequest wrapped in ModelRequest
func RACProviderRequestAsModelRequest(v *RACProviderRequest) ModelRequest {
	return ModelRequest{
		RACProviderRequest: v,
	}
}

// RadiusProviderRequestAsModelRequest is a convenience function that returns RadiusProviderRequest wrapped in ModelRequest
func RadiusProviderRequestAsModelRequest(v *RadiusProviderRequest) ModelRequest {
	return ModelRequest{
		RadiusProviderRequest: v,
	}
}

// SAMLProviderRequestAsModelRequest is a convenience function that returns SAMLProviderRequest wrapped in ModelRequest
func SAMLProviderRequestAsModelRequest(v *SAMLProviderRequest) ModelRequest {
	return ModelRequest{
		SAMLProviderRequest: v,
	}
}

// SCIMProviderRequestAsModelRequest is a convenience function that returns SCIMProviderRequest wrapped in ModelRequest
func SCIMProviderRequestAsModelRequest(v *SCIMProviderRequest) ModelRequest {
	return ModelRequest{
		SCIMProviderRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ModelRequest) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'GoogleWorkspaceProviderRequest'
	if jsonDict["provider_model"] == "GoogleWorkspaceProviderRequest" {
		// try to unmarshal JSON data into GoogleWorkspaceProviderRequest
		err = json.Unmarshal(data, &dst.GoogleWorkspaceProviderRequest)
		if err == nil {
			return nil // data stored in dst.GoogleWorkspaceProviderRequest, return on the first match
		} else {
			dst.GoogleWorkspaceProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as GoogleWorkspaceProviderRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LDAPProviderRequest'
	if jsonDict["provider_model"] == "LDAPProviderRequest" {
		// try to unmarshal JSON data into LDAPProviderRequest
		err = json.Unmarshal(data, &dst.LDAPProviderRequest)
		if err == nil {
			return nil // data stored in dst.LDAPProviderRequest, return on the first match
		} else {
			dst.LDAPProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as LDAPProviderRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MicrosoftEntraProviderRequest'
	if jsonDict["provider_model"] == "MicrosoftEntraProviderRequest" {
		// try to unmarshal JSON data into MicrosoftEntraProviderRequest
		err = json.Unmarshal(data, &dst.MicrosoftEntraProviderRequest)
		if err == nil {
			return nil // data stored in dst.MicrosoftEntraProviderRequest, return on the first match
		} else {
			dst.MicrosoftEntraProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as MicrosoftEntraProviderRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OAuth2ProviderRequest'
	if jsonDict["provider_model"] == "OAuth2ProviderRequest" {
		// try to unmarshal JSON data into OAuth2ProviderRequest
		err = json.Unmarshal(data, &dst.OAuth2ProviderRequest)
		if err == nil {
			return nil // data stored in dst.OAuth2ProviderRequest, return on the first match
		} else {
			dst.OAuth2ProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as OAuth2ProviderRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ProxyProviderRequest'
	if jsonDict["provider_model"] == "ProxyProviderRequest" {
		// try to unmarshal JSON data into ProxyProviderRequest
		err = json.Unmarshal(data, &dst.ProxyProviderRequest)
		if err == nil {
			return nil // data stored in dst.ProxyProviderRequest, return on the first match
		} else {
			dst.ProxyProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as ProxyProviderRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RACProviderRequest'
	if jsonDict["provider_model"] == "RACProviderRequest" {
		// try to unmarshal JSON data into RACProviderRequest
		err = json.Unmarshal(data, &dst.RACProviderRequest)
		if err == nil {
			return nil // data stored in dst.RACProviderRequest, return on the first match
		} else {
			dst.RACProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as RACProviderRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RadiusProviderRequest'
	if jsonDict["provider_model"] == "RadiusProviderRequest" {
		// try to unmarshal JSON data into RadiusProviderRequest
		err = json.Unmarshal(data, &dst.RadiusProviderRequest)
		if err == nil {
			return nil // data stored in dst.RadiusProviderRequest, return on the first match
		} else {
			dst.RadiusProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as RadiusProviderRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SAMLProviderRequest'
	if jsonDict["provider_model"] == "SAMLProviderRequest" {
		// try to unmarshal JSON data into SAMLProviderRequest
		err = json.Unmarshal(data, &dst.SAMLProviderRequest)
		if err == nil {
			return nil // data stored in dst.SAMLProviderRequest, return on the first match
		} else {
			dst.SAMLProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as SAMLProviderRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SCIMProviderRequest'
	if jsonDict["provider_model"] == "SCIMProviderRequest" {
		// try to unmarshal JSON data into SCIMProviderRequest
		err = json.Unmarshal(data, &dst.SCIMProviderRequest)
		if err == nil {
			return nil // data stored in dst.SCIMProviderRequest, return on the first match
		} else {
			dst.SCIMProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as SCIMProviderRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'authentik_providers_google_workspace.googleworkspaceprovider'
	if jsonDict["provider_model"] == "authentik_providers_google_workspace.googleworkspaceprovider" {
		// try to unmarshal JSON data into GoogleWorkspaceProviderRequest
		err = json.Unmarshal(data, &dst.GoogleWorkspaceProviderRequest)
		if err == nil {
			return nil // data stored in dst.GoogleWorkspaceProviderRequest, return on the first match
		} else {
			dst.GoogleWorkspaceProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as GoogleWorkspaceProviderRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'authentik_providers_ldap.ldapprovider'
	if jsonDict["provider_model"] == "authentik_providers_ldap.ldapprovider" {
		// try to unmarshal JSON data into LDAPProviderRequest
		err = json.Unmarshal(data, &dst.LDAPProviderRequest)
		if err == nil {
			return nil // data stored in dst.LDAPProviderRequest, return on the first match
		} else {
			dst.LDAPProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as LDAPProviderRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'authentik_providers_microsoft_entra.microsoftentraprovider'
	if jsonDict["provider_model"] == "authentik_providers_microsoft_entra.microsoftentraprovider" {
		// try to unmarshal JSON data into MicrosoftEntraProviderRequest
		err = json.Unmarshal(data, &dst.MicrosoftEntraProviderRequest)
		if err == nil {
			return nil // data stored in dst.MicrosoftEntraProviderRequest, return on the first match
		} else {
			dst.MicrosoftEntraProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as MicrosoftEntraProviderRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'authentik_providers_oauth2.oauth2provider'
	if jsonDict["provider_model"] == "authentik_providers_oauth2.oauth2provider" {
		// try to unmarshal JSON data into OAuth2ProviderRequest
		err = json.Unmarshal(data, &dst.OAuth2ProviderRequest)
		if err == nil {
			return nil // data stored in dst.OAuth2ProviderRequest, return on the first match
		} else {
			dst.OAuth2ProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as OAuth2ProviderRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'authentik_providers_proxy.proxyprovider'
	if jsonDict["provider_model"] == "authentik_providers_proxy.proxyprovider" {
		// try to unmarshal JSON data into ProxyProviderRequest
		err = json.Unmarshal(data, &dst.ProxyProviderRequest)
		if err == nil {
			return nil // data stored in dst.ProxyProviderRequest, return on the first match
		} else {
			dst.ProxyProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as ProxyProviderRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'authentik_providers_rac.racprovider'
	if jsonDict["provider_model"] == "authentik_providers_rac.racprovider" {
		// try to unmarshal JSON data into RACProviderRequest
		err = json.Unmarshal(data, &dst.RACProviderRequest)
		if err == nil {
			return nil // data stored in dst.RACProviderRequest, return on the first match
		} else {
			dst.RACProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as RACProviderRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'authentik_providers_radius.radiusprovider'
	if jsonDict["provider_model"] == "authentik_providers_radius.radiusprovider" {
		// try to unmarshal JSON data into RadiusProviderRequest
		err = json.Unmarshal(data, &dst.RadiusProviderRequest)
		if err == nil {
			return nil // data stored in dst.RadiusProviderRequest, return on the first match
		} else {
			dst.RadiusProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as RadiusProviderRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'authentik_providers_saml.samlprovider'
	if jsonDict["provider_model"] == "authentik_providers_saml.samlprovider" {
		// try to unmarshal JSON data into SAMLProviderRequest
		err = json.Unmarshal(data, &dst.SAMLProviderRequest)
		if err == nil {
			return nil // data stored in dst.SAMLProviderRequest, return on the first match
		} else {
			dst.SAMLProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as SAMLProviderRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'authentik_providers_scim.scimprovider'
	if jsonDict["provider_model"] == "authentik_providers_scim.scimprovider" {
		// try to unmarshal JSON data into SCIMProviderRequest
		err = json.Unmarshal(data, &dst.SCIMProviderRequest)
		if err == nil {
			return nil // data stored in dst.SCIMProviderRequest, return on the first match
		} else {
			dst.SCIMProviderRequest = nil
			return fmt.Errorf("Failed to unmarshal ModelRequest as SCIMProviderRequest: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ModelRequest) MarshalJSON() ([]byte, error) {
	if src.GoogleWorkspaceProviderRequest != nil {
		return json.Marshal(&src.GoogleWorkspaceProviderRequest)
	}

	if src.LDAPProviderRequest != nil {
		return json.Marshal(&src.LDAPProviderRequest)
	}

	if src.MicrosoftEntraProviderRequest != nil {
		return json.Marshal(&src.MicrosoftEntraProviderRequest)
	}

	if src.OAuth2ProviderRequest != nil {
		return json.Marshal(&src.OAuth2ProviderRequest)
	}

	if src.ProxyProviderRequest != nil {
		return json.Marshal(&src.ProxyProviderRequest)
	}

	if src.RACProviderRequest != nil {
		return json.Marshal(&src.RACProviderRequest)
	}

	if src.RadiusProviderRequest != nil {
		return json.Marshal(&src.RadiusProviderRequest)
	}

	if src.SAMLProviderRequest != nil {
		return json.Marshal(&src.SAMLProviderRequest)
	}

	if src.SCIMProviderRequest != nil {
		return json.Marshal(&src.SCIMProviderRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ModelRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GoogleWorkspaceProviderRequest != nil {
		return obj.GoogleWorkspaceProviderRequest
	}

	if obj.LDAPProviderRequest != nil {
		return obj.LDAPProviderRequest
	}

	if obj.MicrosoftEntraProviderRequest != nil {
		return obj.MicrosoftEntraProviderRequest
	}

	if obj.OAuth2ProviderRequest != nil {
		return obj.OAuth2ProviderRequest
	}

	if obj.ProxyProviderRequest != nil {
		return obj.ProxyProviderRequest
	}

	if obj.RACProviderRequest != nil {
		return obj.RACProviderRequest
	}

	if obj.RadiusProviderRequest != nil {
		return obj.RadiusProviderRequest
	}

	if obj.SAMLProviderRequest != nil {
		return obj.SAMLProviderRequest
	}

	if obj.SCIMProviderRequest != nil {
		return obj.SCIMProviderRequest
	}

	// all schemas are nil
	return nil
}

type NullableModelRequest struct {
	value *ModelRequest
	isSet bool
}

func (v NullableModelRequest) Get() *ModelRequest {
	return v.value
}

func (v *NullableModelRequest) Set(val *ModelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRequest(val *ModelRequest) *NullableModelRequest {
	return &NullableModelRequest{value: val, isSet: true}
}

func (v NullableModelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
