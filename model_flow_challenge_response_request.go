/*
authentik

Making authentication simple.

API version: 2022.4.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// FlowChallengeResponseRequest - struct for FlowChallengeResponseRequest
type FlowChallengeResponseRequest struct {
	AppleChallengeResponseRequest                   *AppleChallengeResponseRequest
	AuthenticatorDuoChallengeResponseRequest        *AuthenticatorDuoChallengeResponseRequest
	AuthenticatorSMSChallengeResponseRequest        *AuthenticatorSMSChallengeResponseRequest
	AuthenticatorStaticChallengeResponseRequest     *AuthenticatorStaticChallengeResponseRequest
	AuthenticatorTOTPChallengeResponseRequest       *AuthenticatorTOTPChallengeResponseRequest
	AuthenticatorValidationChallengeResponseRequest *AuthenticatorValidationChallengeResponseRequest
	AuthenticatorWebAuthnChallengeResponseRequest   *AuthenticatorWebAuthnChallengeResponseRequest
	AutoSubmitChallengeResponseRequest              *AutoSubmitChallengeResponseRequest
	CaptchaChallengeResponseRequest                 *CaptchaChallengeResponseRequest
	ConsentChallengeResponseRequest                 *ConsentChallengeResponseRequest
	DummyChallengeResponseRequest                   *DummyChallengeResponseRequest
	EmailChallengeResponseRequest                   *EmailChallengeResponseRequest
	IdentificationChallengeResponseRequest          *IdentificationChallengeResponseRequest
	PasswordChallengeResponseRequest                *PasswordChallengeResponseRequest
	PlexAuthenticationChallengeResponseRequest      *PlexAuthenticationChallengeResponseRequest
	PromptChallengeResponseRequest                  *PromptChallengeResponseRequest
}

// AppleChallengeResponseRequestAsFlowChallengeResponseRequest is a convenience function that returns AppleChallengeResponseRequest wrapped in FlowChallengeResponseRequest
func AppleChallengeResponseRequestAsFlowChallengeResponseRequest(v *AppleChallengeResponseRequest) FlowChallengeResponseRequest {
	return FlowChallengeResponseRequest{AppleChallengeResponseRequest: v}
}

// AuthenticatorDuoChallengeResponseRequestAsFlowChallengeResponseRequest is a convenience function that returns AuthenticatorDuoChallengeResponseRequest wrapped in FlowChallengeResponseRequest
func AuthenticatorDuoChallengeResponseRequestAsFlowChallengeResponseRequest(v *AuthenticatorDuoChallengeResponseRequest) FlowChallengeResponseRequest {
	return FlowChallengeResponseRequest{AuthenticatorDuoChallengeResponseRequest: v}
}

// AuthenticatorSMSChallengeResponseRequestAsFlowChallengeResponseRequest is a convenience function that returns AuthenticatorSMSChallengeResponseRequest wrapped in FlowChallengeResponseRequest
func AuthenticatorSMSChallengeResponseRequestAsFlowChallengeResponseRequest(v *AuthenticatorSMSChallengeResponseRequest) FlowChallengeResponseRequest {
	return FlowChallengeResponseRequest{AuthenticatorSMSChallengeResponseRequest: v}
}

// AuthenticatorStaticChallengeResponseRequestAsFlowChallengeResponseRequest is a convenience function that returns AuthenticatorStaticChallengeResponseRequest wrapped in FlowChallengeResponseRequest
func AuthenticatorStaticChallengeResponseRequestAsFlowChallengeResponseRequest(v *AuthenticatorStaticChallengeResponseRequest) FlowChallengeResponseRequest {
	return FlowChallengeResponseRequest{AuthenticatorStaticChallengeResponseRequest: v}
}

// AuthenticatorTOTPChallengeResponseRequestAsFlowChallengeResponseRequest is a convenience function that returns AuthenticatorTOTPChallengeResponseRequest wrapped in FlowChallengeResponseRequest
func AuthenticatorTOTPChallengeResponseRequestAsFlowChallengeResponseRequest(v *AuthenticatorTOTPChallengeResponseRequest) FlowChallengeResponseRequest {
	return FlowChallengeResponseRequest{AuthenticatorTOTPChallengeResponseRequest: v}
}

// AuthenticatorValidationChallengeResponseRequestAsFlowChallengeResponseRequest is a convenience function that returns AuthenticatorValidationChallengeResponseRequest wrapped in FlowChallengeResponseRequest
func AuthenticatorValidationChallengeResponseRequestAsFlowChallengeResponseRequest(v *AuthenticatorValidationChallengeResponseRequest) FlowChallengeResponseRequest {
	return FlowChallengeResponseRequest{AuthenticatorValidationChallengeResponseRequest: v}
}

// AuthenticatorWebAuthnChallengeResponseRequestAsFlowChallengeResponseRequest is a convenience function that returns AuthenticatorWebAuthnChallengeResponseRequest wrapped in FlowChallengeResponseRequest
func AuthenticatorWebAuthnChallengeResponseRequestAsFlowChallengeResponseRequest(v *AuthenticatorWebAuthnChallengeResponseRequest) FlowChallengeResponseRequest {
	return FlowChallengeResponseRequest{AuthenticatorWebAuthnChallengeResponseRequest: v}
}

// AutoSubmitChallengeResponseRequestAsFlowChallengeResponseRequest is a convenience function that returns AutoSubmitChallengeResponseRequest wrapped in FlowChallengeResponseRequest
func AutoSubmitChallengeResponseRequestAsFlowChallengeResponseRequest(v *AutoSubmitChallengeResponseRequest) FlowChallengeResponseRequest {
	return FlowChallengeResponseRequest{AutoSubmitChallengeResponseRequest: v}
}

// CaptchaChallengeResponseRequestAsFlowChallengeResponseRequest is a convenience function that returns CaptchaChallengeResponseRequest wrapped in FlowChallengeResponseRequest
func CaptchaChallengeResponseRequestAsFlowChallengeResponseRequest(v *CaptchaChallengeResponseRequest) FlowChallengeResponseRequest {
	return FlowChallengeResponseRequest{CaptchaChallengeResponseRequest: v}
}

// ConsentChallengeResponseRequestAsFlowChallengeResponseRequest is a convenience function that returns ConsentChallengeResponseRequest wrapped in FlowChallengeResponseRequest
func ConsentChallengeResponseRequestAsFlowChallengeResponseRequest(v *ConsentChallengeResponseRequest) FlowChallengeResponseRequest {
	return FlowChallengeResponseRequest{ConsentChallengeResponseRequest: v}
}

// DummyChallengeResponseRequestAsFlowChallengeResponseRequest is a convenience function that returns DummyChallengeResponseRequest wrapped in FlowChallengeResponseRequest
func DummyChallengeResponseRequestAsFlowChallengeResponseRequest(v *DummyChallengeResponseRequest) FlowChallengeResponseRequest {
	return FlowChallengeResponseRequest{DummyChallengeResponseRequest: v}
}

// EmailChallengeResponseRequestAsFlowChallengeResponseRequest is a convenience function that returns EmailChallengeResponseRequest wrapped in FlowChallengeResponseRequest
func EmailChallengeResponseRequestAsFlowChallengeResponseRequest(v *EmailChallengeResponseRequest) FlowChallengeResponseRequest {
	return FlowChallengeResponseRequest{EmailChallengeResponseRequest: v}
}

// IdentificationChallengeResponseRequestAsFlowChallengeResponseRequest is a convenience function that returns IdentificationChallengeResponseRequest wrapped in FlowChallengeResponseRequest
func IdentificationChallengeResponseRequestAsFlowChallengeResponseRequest(v *IdentificationChallengeResponseRequest) FlowChallengeResponseRequest {
	return FlowChallengeResponseRequest{IdentificationChallengeResponseRequest: v}
}

// PasswordChallengeResponseRequestAsFlowChallengeResponseRequest is a convenience function that returns PasswordChallengeResponseRequest wrapped in FlowChallengeResponseRequest
func PasswordChallengeResponseRequestAsFlowChallengeResponseRequest(v *PasswordChallengeResponseRequest) FlowChallengeResponseRequest {
	return FlowChallengeResponseRequest{PasswordChallengeResponseRequest: v}
}

// PlexAuthenticationChallengeResponseRequestAsFlowChallengeResponseRequest is a convenience function that returns PlexAuthenticationChallengeResponseRequest wrapped in FlowChallengeResponseRequest
func PlexAuthenticationChallengeResponseRequestAsFlowChallengeResponseRequest(v *PlexAuthenticationChallengeResponseRequest) FlowChallengeResponseRequest {
	return FlowChallengeResponseRequest{PlexAuthenticationChallengeResponseRequest: v}
}

// PromptChallengeResponseRequestAsFlowChallengeResponseRequest is a convenience function that returns PromptChallengeResponseRequest wrapped in FlowChallengeResponseRequest
func PromptChallengeResponseRequestAsFlowChallengeResponseRequest(v *PromptChallengeResponseRequest) FlowChallengeResponseRequest {
	return FlowChallengeResponseRequest{PromptChallengeResponseRequest: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FlowChallengeResponseRequest) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'AppleChallengeResponseRequest'
	if jsonDict["component"] == "AppleChallengeResponseRequest" {
		// try to unmarshal JSON data into AppleChallengeResponseRequest
		err = json.Unmarshal(data, &dst.AppleChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.AppleChallengeResponseRequest, return on the first match
		} else {
			dst.AppleChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as AppleChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorDuoChallengeResponseRequest'
	if jsonDict["component"] == "AuthenticatorDuoChallengeResponseRequest" {
		// try to unmarshal JSON data into AuthenticatorDuoChallengeResponseRequest
		err = json.Unmarshal(data, &dst.AuthenticatorDuoChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.AuthenticatorDuoChallengeResponseRequest, return on the first match
		} else {
			dst.AuthenticatorDuoChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as AuthenticatorDuoChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorSMSChallengeResponseRequest'
	if jsonDict["component"] == "AuthenticatorSMSChallengeResponseRequest" {
		// try to unmarshal JSON data into AuthenticatorSMSChallengeResponseRequest
		err = json.Unmarshal(data, &dst.AuthenticatorSMSChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.AuthenticatorSMSChallengeResponseRequest, return on the first match
		} else {
			dst.AuthenticatorSMSChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as AuthenticatorSMSChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorStaticChallengeResponseRequest'
	if jsonDict["component"] == "AuthenticatorStaticChallengeResponseRequest" {
		// try to unmarshal JSON data into AuthenticatorStaticChallengeResponseRequest
		err = json.Unmarshal(data, &dst.AuthenticatorStaticChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.AuthenticatorStaticChallengeResponseRequest, return on the first match
		} else {
			dst.AuthenticatorStaticChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as AuthenticatorStaticChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorTOTPChallengeResponseRequest'
	if jsonDict["component"] == "AuthenticatorTOTPChallengeResponseRequest" {
		// try to unmarshal JSON data into AuthenticatorTOTPChallengeResponseRequest
		err = json.Unmarshal(data, &dst.AuthenticatorTOTPChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.AuthenticatorTOTPChallengeResponseRequest, return on the first match
		} else {
			dst.AuthenticatorTOTPChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as AuthenticatorTOTPChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorValidationChallengeResponseRequest'
	if jsonDict["component"] == "AuthenticatorValidationChallengeResponseRequest" {
		// try to unmarshal JSON data into AuthenticatorValidationChallengeResponseRequest
		err = json.Unmarshal(data, &dst.AuthenticatorValidationChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.AuthenticatorValidationChallengeResponseRequest, return on the first match
		} else {
			dst.AuthenticatorValidationChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as AuthenticatorValidationChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorWebAuthnChallengeResponseRequest'
	if jsonDict["component"] == "AuthenticatorWebAuthnChallengeResponseRequest" {
		// try to unmarshal JSON data into AuthenticatorWebAuthnChallengeResponseRequest
		err = json.Unmarshal(data, &dst.AuthenticatorWebAuthnChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.AuthenticatorWebAuthnChallengeResponseRequest, return on the first match
		} else {
			dst.AuthenticatorWebAuthnChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as AuthenticatorWebAuthnChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AutoSubmitChallengeResponseRequest'
	if jsonDict["component"] == "AutoSubmitChallengeResponseRequest" {
		// try to unmarshal JSON data into AutoSubmitChallengeResponseRequest
		err = json.Unmarshal(data, &dst.AutoSubmitChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.AutoSubmitChallengeResponseRequest, return on the first match
		} else {
			dst.AutoSubmitChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as AutoSubmitChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CaptchaChallengeResponseRequest'
	if jsonDict["component"] == "CaptchaChallengeResponseRequest" {
		// try to unmarshal JSON data into CaptchaChallengeResponseRequest
		err = json.Unmarshal(data, &dst.CaptchaChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.CaptchaChallengeResponseRequest, return on the first match
		} else {
			dst.CaptchaChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as CaptchaChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ConsentChallengeResponseRequest'
	if jsonDict["component"] == "ConsentChallengeResponseRequest" {
		// try to unmarshal JSON data into ConsentChallengeResponseRequest
		err = json.Unmarshal(data, &dst.ConsentChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.ConsentChallengeResponseRequest, return on the first match
		} else {
			dst.ConsentChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as ConsentChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DummyChallengeResponseRequest'
	if jsonDict["component"] == "DummyChallengeResponseRequest" {
		// try to unmarshal JSON data into DummyChallengeResponseRequest
		err = json.Unmarshal(data, &dst.DummyChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.DummyChallengeResponseRequest, return on the first match
		} else {
			dst.DummyChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as DummyChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EmailChallengeResponseRequest'
	if jsonDict["component"] == "EmailChallengeResponseRequest" {
		// try to unmarshal JSON data into EmailChallengeResponseRequest
		err = json.Unmarshal(data, &dst.EmailChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.EmailChallengeResponseRequest, return on the first match
		} else {
			dst.EmailChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as EmailChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'IdentificationChallengeResponseRequest'
	if jsonDict["component"] == "IdentificationChallengeResponseRequest" {
		// try to unmarshal JSON data into IdentificationChallengeResponseRequest
		err = json.Unmarshal(data, &dst.IdentificationChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.IdentificationChallengeResponseRequest, return on the first match
		} else {
			dst.IdentificationChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as IdentificationChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PasswordChallengeResponseRequest'
	if jsonDict["component"] == "PasswordChallengeResponseRequest" {
		// try to unmarshal JSON data into PasswordChallengeResponseRequest
		err = json.Unmarshal(data, &dst.PasswordChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.PasswordChallengeResponseRequest, return on the first match
		} else {
			dst.PasswordChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as PasswordChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PlexAuthenticationChallengeResponseRequest'
	if jsonDict["component"] == "PlexAuthenticationChallengeResponseRequest" {
		// try to unmarshal JSON data into PlexAuthenticationChallengeResponseRequest
		err = json.Unmarshal(data, &dst.PlexAuthenticationChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.PlexAuthenticationChallengeResponseRequest, return on the first match
		} else {
			dst.PlexAuthenticationChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as PlexAuthenticationChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PromptChallengeResponseRequest'
	if jsonDict["component"] == "PromptChallengeResponseRequest" {
		// try to unmarshal JSON data into PromptChallengeResponseRequest
		err = json.Unmarshal(data, &dst.PromptChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.PromptChallengeResponseRequest, return on the first match
		} else {
			dst.PromptChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as PromptChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-flow-sources-oauth-apple'
	if jsonDict["component"] == "ak-flow-sources-oauth-apple" {
		// try to unmarshal JSON data into AppleChallengeResponseRequest
		err = json.Unmarshal(data, &dst.AppleChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.AppleChallengeResponseRequest, return on the first match
		} else {
			dst.AppleChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as AppleChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-flow-sources-plex'
	if jsonDict["component"] == "ak-flow-sources-plex" {
		// try to unmarshal JSON data into PlexAuthenticationChallengeResponseRequest
		err = json.Unmarshal(data, &dst.PlexAuthenticationChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.PlexAuthenticationChallengeResponseRequest, return on the first match
		} else {
			dst.PlexAuthenticationChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as PlexAuthenticationChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-authenticator-duo'
	if jsonDict["component"] == "ak-stage-authenticator-duo" {
		// try to unmarshal JSON data into AuthenticatorDuoChallengeResponseRequest
		err = json.Unmarshal(data, &dst.AuthenticatorDuoChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.AuthenticatorDuoChallengeResponseRequest, return on the first match
		} else {
			dst.AuthenticatorDuoChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as AuthenticatorDuoChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-authenticator-sms'
	if jsonDict["component"] == "ak-stage-authenticator-sms" {
		// try to unmarshal JSON data into AuthenticatorSMSChallengeResponseRequest
		err = json.Unmarshal(data, &dst.AuthenticatorSMSChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.AuthenticatorSMSChallengeResponseRequest, return on the first match
		} else {
			dst.AuthenticatorSMSChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as AuthenticatorSMSChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-authenticator-static'
	if jsonDict["component"] == "ak-stage-authenticator-static" {
		// try to unmarshal JSON data into AuthenticatorStaticChallengeResponseRequest
		err = json.Unmarshal(data, &dst.AuthenticatorStaticChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.AuthenticatorStaticChallengeResponseRequest, return on the first match
		} else {
			dst.AuthenticatorStaticChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as AuthenticatorStaticChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-authenticator-totp'
	if jsonDict["component"] == "ak-stage-authenticator-totp" {
		// try to unmarshal JSON data into AuthenticatorTOTPChallengeResponseRequest
		err = json.Unmarshal(data, &dst.AuthenticatorTOTPChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.AuthenticatorTOTPChallengeResponseRequest, return on the first match
		} else {
			dst.AuthenticatorTOTPChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as AuthenticatorTOTPChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-authenticator-validate'
	if jsonDict["component"] == "ak-stage-authenticator-validate" {
		// try to unmarshal JSON data into AuthenticatorValidationChallengeResponseRequest
		err = json.Unmarshal(data, &dst.AuthenticatorValidationChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.AuthenticatorValidationChallengeResponseRequest, return on the first match
		} else {
			dst.AuthenticatorValidationChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as AuthenticatorValidationChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-authenticator-webauthn'
	if jsonDict["component"] == "ak-stage-authenticator-webauthn" {
		// try to unmarshal JSON data into AuthenticatorWebAuthnChallengeResponseRequest
		err = json.Unmarshal(data, &dst.AuthenticatorWebAuthnChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.AuthenticatorWebAuthnChallengeResponseRequest, return on the first match
		} else {
			dst.AuthenticatorWebAuthnChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as AuthenticatorWebAuthnChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-autosubmit'
	if jsonDict["component"] == "ak-stage-autosubmit" {
		// try to unmarshal JSON data into AutoSubmitChallengeResponseRequest
		err = json.Unmarshal(data, &dst.AutoSubmitChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.AutoSubmitChallengeResponseRequest, return on the first match
		} else {
			dst.AutoSubmitChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as AutoSubmitChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-captcha'
	if jsonDict["component"] == "ak-stage-captcha" {
		// try to unmarshal JSON data into CaptchaChallengeResponseRequest
		err = json.Unmarshal(data, &dst.CaptchaChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.CaptchaChallengeResponseRequest, return on the first match
		} else {
			dst.CaptchaChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as CaptchaChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-consent'
	if jsonDict["component"] == "ak-stage-consent" {
		// try to unmarshal JSON data into ConsentChallengeResponseRequest
		err = json.Unmarshal(data, &dst.ConsentChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.ConsentChallengeResponseRequest, return on the first match
		} else {
			dst.ConsentChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as ConsentChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-dummy'
	if jsonDict["component"] == "ak-stage-dummy" {
		// try to unmarshal JSON data into DummyChallengeResponseRequest
		err = json.Unmarshal(data, &dst.DummyChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.DummyChallengeResponseRequest, return on the first match
		} else {
			dst.DummyChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as DummyChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-email'
	if jsonDict["component"] == "ak-stage-email" {
		// try to unmarshal JSON data into EmailChallengeResponseRequest
		err = json.Unmarshal(data, &dst.EmailChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.EmailChallengeResponseRequest, return on the first match
		} else {
			dst.EmailChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as EmailChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-identification'
	if jsonDict["component"] == "ak-stage-identification" {
		// try to unmarshal JSON data into IdentificationChallengeResponseRequest
		err = json.Unmarshal(data, &dst.IdentificationChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.IdentificationChallengeResponseRequest, return on the first match
		} else {
			dst.IdentificationChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as IdentificationChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-password'
	if jsonDict["component"] == "ak-stage-password" {
		// try to unmarshal JSON data into PasswordChallengeResponseRequest
		err = json.Unmarshal(data, &dst.PasswordChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.PasswordChallengeResponseRequest, return on the first match
		} else {
			dst.PasswordChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as PasswordChallengeResponseRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-prompt'
	if jsonDict["component"] == "ak-stage-prompt" {
		// try to unmarshal JSON data into PromptChallengeResponseRequest
		err = json.Unmarshal(data, &dst.PromptChallengeResponseRequest)
		if err == nil {
			return nil // data stored in dst.PromptChallengeResponseRequest, return on the first match
		} else {
			dst.PromptChallengeResponseRequest = nil
			return fmt.Errorf("Failed to unmarshal FlowChallengeResponseRequest as PromptChallengeResponseRequest: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FlowChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	if src.AppleChallengeResponseRequest != nil {
		return json.Marshal(&src.AppleChallengeResponseRequest)
	}

	if src.AuthenticatorDuoChallengeResponseRequest != nil {
		return json.Marshal(&src.AuthenticatorDuoChallengeResponseRequest)
	}

	if src.AuthenticatorSMSChallengeResponseRequest != nil {
		return json.Marshal(&src.AuthenticatorSMSChallengeResponseRequest)
	}

	if src.AuthenticatorStaticChallengeResponseRequest != nil {
		return json.Marshal(&src.AuthenticatorStaticChallengeResponseRequest)
	}

	if src.AuthenticatorTOTPChallengeResponseRequest != nil {
		return json.Marshal(&src.AuthenticatorTOTPChallengeResponseRequest)
	}

	if src.AuthenticatorValidationChallengeResponseRequest != nil {
		return json.Marshal(&src.AuthenticatorValidationChallengeResponseRequest)
	}

	if src.AuthenticatorWebAuthnChallengeResponseRequest != nil {
		return json.Marshal(&src.AuthenticatorWebAuthnChallengeResponseRequest)
	}

	if src.AutoSubmitChallengeResponseRequest != nil {
		return json.Marshal(&src.AutoSubmitChallengeResponseRequest)
	}

	if src.CaptchaChallengeResponseRequest != nil {
		return json.Marshal(&src.CaptchaChallengeResponseRequest)
	}

	if src.ConsentChallengeResponseRequest != nil {
		return json.Marshal(&src.ConsentChallengeResponseRequest)
	}

	if src.DummyChallengeResponseRequest != nil {
		return json.Marshal(&src.DummyChallengeResponseRequest)
	}

	if src.EmailChallengeResponseRequest != nil {
		return json.Marshal(&src.EmailChallengeResponseRequest)
	}

	if src.IdentificationChallengeResponseRequest != nil {
		return json.Marshal(&src.IdentificationChallengeResponseRequest)
	}

	if src.PasswordChallengeResponseRequest != nil {
		return json.Marshal(&src.PasswordChallengeResponseRequest)
	}

	if src.PlexAuthenticationChallengeResponseRequest != nil {
		return json.Marshal(&src.PlexAuthenticationChallengeResponseRequest)
	}

	if src.PromptChallengeResponseRequest != nil {
		return json.Marshal(&src.PromptChallengeResponseRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FlowChallengeResponseRequest) GetActualInstance() interface{} {
	if obj.AppleChallengeResponseRequest != nil {
		return obj.AppleChallengeResponseRequest
	}

	if obj.AuthenticatorDuoChallengeResponseRequest != nil {
		return obj.AuthenticatorDuoChallengeResponseRequest
	}

	if obj.AuthenticatorSMSChallengeResponseRequest != nil {
		return obj.AuthenticatorSMSChallengeResponseRequest
	}

	if obj.AuthenticatorStaticChallengeResponseRequest != nil {
		return obj.AuthenticatorStaticChallengeResponseRequest
	}

	if obj.AuthenticatorTOTPChallengeResponseRequest != nil {
		return obj.AuthenticatorTOTPChallengeResponseRequest
	}

	if obj.AuthenticatorValidationChallengeResponseRequest != nil {
		return obj.AuthenticatorValidationChallengeResponseRequest
	}

	if obj.AuthenticatorWebAuthnChallengeResponseRequest != nil {
		return obj.AuthenticatorWebAuthnChallengeResponseRequest
	}

	if obj.AutoSubmitChallengeResponseRequest != nil {
		return obj.AutoSubmitChallengeResponseRequest
	}

	if obj.CaptchaChallengeResponseRequest != nil {
		return obj.CaptchaChallengeResponseRequest
	}

	if obj.ConsentChallengeResponseRequest != nil {
		return obj.ConsentChallengeResponseRequest
	}

	if obj.DummyChallengeResponseRequest != nil {
		return obj.DummyChallengeResponseRequest
	}

	if obj.EmailChallengeResponseRequest != nil {
		return obj.EmailChallengeResponseRequest
	}

	if obj.IdentificationChallengeResponseRequest != nil {
		return obj.IdentificationChallengeResponseRequest
	}

	if obj.PasswordChallengeResponseRequest != nil {
		return obj.PasswordChallengeResponseRequest
	}

	if obj.PlexAuthenticationChallengeResponseRequest != nil {
		return obj.PlexAuthenticationChallengeResponseRequest
	}

	if obj.PromptChallengeResponseRequest != nil {
		return obj.PromptChallengeResponseRequest
	}

	// all schemas are nil
	return nil
}

type NullableFlowChallengeResponseRequest struct {
	value *FlowChallengeResponseRequest
	isSet bool
}

func (v NullableFlowChallengeResponseRequest) Get() *FlowChallengeResponseRequest {
	return v.value
}

func (v *NullableFlowChallengeResponseRequest) Set(val *FlowChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowChallengeResponseRequest(val *FlowChallengeResponseRequest) *NullableFlowChallengeResponseRequest {
	return &NullableFlowChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableFlowChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
