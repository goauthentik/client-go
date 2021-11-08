/*
authentik

Making authentication simple.

API version: 2021.10.3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ChallengeTypes - struct for ChallengeTypes
type ChallengeTypes struct {
	AccessDeniedChallenge            *AccessDeniedChallenge
	AuthenticatorDuoChallenge        *AuthenticatorDuoChallenge
	AuthenticatorSMSChallenge        *AuthenticatorSMSChallenge
	AuthenticatorStaticChallenge     *AuthenticatorStaticChallenge
	AuthenticatorTOTPChallenge       *AuthenticatorTOTPChallenge
	AuthenticatorValidationChallenge *AuthenticatorValidationChallenge
	AuthenticatorWebAuthnChallenge   *AuthenticatorWebAuthnChallenge
	AutosubmitChallenge              *AutosubmitChallenge
	CaptchaChallenge                 *CaptchaChallenge
	ConsentChallenge                 *ConsentChallenge
	DummyChallenge                   *DummyChallenge
	EmailChallenge                   *EmailChallenge
	IdentificationChallenge          *IdentificationChallenge
	PasswordChallenge                *PasswordChallenge
	PlexAuthenticationChallenge      *PlexAuthenticationChallenge
	PromptChallenge                  *PromptChallenge
	RedirectChallenge                *RedirectChallenge
	ShellChallenge                   *ShellChallenge
}

// AccessDeniedChallengeAsChallengeTypes is a convenience function that returns AccessDeniedChallenge wrapped in ChallengeTypes
func AccessDeniedChallengeAsChallengeTypes(v *AccessDeniedChallenge) ChallengeTypes {
	return ChallengeTypes{AccessDeniedChallenge: v}
}

// AuthenticatorDuoChallengeAsChallengeTypes is a convenience function that returns AuthenticatorDuoChallenge wrapped in ChallengeTypes
func AuthenticatorDuoChallengeAsChallengeTypes(v *AuthenticatorDuoChallenge) ChallengeTypes {
	return ChallengeTypes{AuthenticatorDuoChallenge: v}
}

// AuthenticatorSMSChallengeAsChallengeTypes is a convenience function that returns AuthenticatorSMSChallenge wrapped in ChallengeTypes
func AuthenticatorSMSChallengeAsChallengeTypes(v *AuthenticatorSMSChallenge) ChallengeTypes {
	return ChallengeTypes{AuthenticatorSMSChallenge: v}
}

// AuthenticatorStaticChallengeAsChallengeTypes is a convenience function that returns AuthenticatorStaticChallenge wrapped in ChallengeTypes
func AuthenticatorStaticChallengeAsChallengeTypes(v *AuthenticatorStaticChallenge) ChallengeTypes {
	return ChallengeTypes{AuthenticatorStaticChallenge: v}
}

// AuthenticatorTOTPChallengeAsChallengeTypes is a convenience function that returns AuthenticatorTOTPChallenge wrapped in ChallengeTypes
func AuthenticatorTOTPChallengeAsChallengeTypes(v *AuthenticatorTOTPChallenge) ChallengeTypes {
	return ChallengeTypes{AuthenticatorTOTPChallenge: v}
}

// AuthenticatorValidationChallengeAsChallengeTypes is a convenience function that returns AuthenticatorValidationChallenge wrapped in ChallengeTypes
func AuthenticatorValidationChallengeAsChallengeTypes(v *AuthenticatorValidationChallenge) ChallengeTypes {
	return ChallengeTypes{AuthenticatorValidationChallenge: v}
}

// AuthenticatorWebAuthnChallengeAsChallengeTypes is a convenience function that returns AuthenticatorWebAuthnChallenge wrapped in ChallengeTypes
func AuthenticatorWebAuthnChallengeAsChallengeTypes(v *AuthenticatorWebAuthnChallenge) ChallengeTypes {
	return ChallengeTypes{AuthenticatorWebAuthnChallenge: v}
}

// AutosubmitChallengeAsChallengeTypes is a convenience function that returns AutosubmitChallenge wrapped in ChallengeTypes
func AutosubmitChallengeAsChallengeTypes(v *AutosubmitChallenge) ChallengeTypes {
	return ChallengeTypes{AutosubmitChallenge: v}
}

// CaptchaChallengeAsChallengeTypes is a convenience function that returns CaptchaChallenge wrapped in ChallengeTypes
func CaptchaChallengeAsChallengeTypes(v *CaptchaChallenge) ChallengeTypes {
	return ChallengeTypes{CaptchaChallenge: v}
}

// ConsentChallengeAsChallengeTypes is a convenience function that returns ConsentChallenge wrapped in ChallengeTypes
func ConsentChallengeAsChallengeTypes(v *ConsentChallenge) ChallengeTypes {
	return ChallengeTypes{ConsentChallenge: v}
}

// DummyChallengeAsChallengeTypes is a convenience function that returns DummyChallenge wrapped in ChallengeTypes
func DummyChallengeAsChallengeTypes(v *DummyChallenge) ChallengeTypes {
	return ChallengeTypes{DummyChallenge: v}
}

// EmailChallengeAsChallengeTypes is a convenience function that returns EmailChallenge wrapped in ChallengeTypes
func EmailChallengeAsChallengeTypes(v *EmailChallenge) ChallengeTypes {
	return ChallengeTypes{EmailChallenge: v}
}

// IdentificationChallengeAsChallengeTypes is a convenience function that returns IdentificationChallenge wrapped in ChallengeTypes
func IdentificationChallengeAsChallengeTypes(v *IdentificationChallenge) ChallengeTypes {
	return ChallengeTypes{IdentificationChallenge: v}
}

// PasswordChallengeAsChallengeTypes is a convenience function that returns PasswordChallenge wrapped in ChallengeTypes
func PasswordChallengeAsChallengeTypes(v *PasswordChallenge) ChallengeTypes {
	return ChallengeTypes{PasswordChallenge: v}
}

// PlexAuthenticationChallengeAsChallengeTypes is a convenience function that returns PlexAuthenticationChallenge wrapped in ChallengeTypes
func PlexAuthenticationChallengeAsChallengeTypes(v *PlexAuthenticationChallenge) ChallengeTypes {
	return ChallengeTypes{PlexAuthenticationChallenge: v}
}

// PromptChallengeAsChallengeTypes is a convenience function that returns PromptChallenge wrapped in ChallengeTypes
func PromptChallengeAsChallengeTypes(v *PromptChallenge) ChallengeTypes {
	return ChallengeTypes{PromptChallenge: v}
}

// RedirectChallengeAsChallengeTypes is a convenience function that returns RedirectChallenge wrapped in ChallengeTypes
func RedirectChallengeAsChallengeTypes(v *RedirectChallenge) ChallengeTypes {
	return ChallengeTypes{RedirectChallenge: v}
}

// ShellChallengeAsChallengeTypes is a convenience function that returns ShellChallenge wrapped in ChallengeTypes
func ShellChallengeAsChallengeTypes(v *ShellChallenge) ChallengeTypes {
	return ChallengeTypes{ShellChallenge: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ChallengeTypes) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'AccessDeniedChallenge'
	if jsonDict["component"] == "AccessDeniedChallenge" {
		// try to unmarshal JSON data into AccessDeniedChallenge
		err = json.Unmarshal(data, &dst.AccessDeniedChallenge)
		if err == nil {
			return nil // data stored in dst.AccessDeniedChallenge, return on the first match
		} else {
			dst.AccessDeniedChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as AccessDeniedChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorDuoChallenge'
	if jsonDict["component"] == "AuthenticatorDuoChallenge" {
		// try to unmarshal JSON data into AuthenticatorDuoChallenge
		err = json.Unmarshal(data, &dst.AuthenticatorDuoChallenge)
		if err == nil {
			return nil // data stored in dst.AuthenticatorDuoChallenge, return on the first match
		} else {
			dst.AuthenticatorDuoChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as AuthenticatorDuoChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorSMSChallenge'
	if jsonDict["component"] == "AuthenticatorSMSChallenge" {
		// try to unmarshal JSON data into AuthenticatorSMSChallenge
		err = json.Unmarshal(data, &dst.AuthenticatorSMSChallenge)
		if err == nil {
			return nil // data stored in dst.AuthenticatorSMSChallenge, return on the first match
		} else {
			dst.AuthenticatorSMSChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as AuthenticatorSMSChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorStaticChallenge'
	if jsonDict["component"] == "AuthenticatorStaticChallenge" {
		// try to unmarshal JSON data into AuthenticatorStaticChallenge
		err = json.Unmarshal(data, &dst.AuthenticatorStaticChallenge)
		if err == nil {
			return nil // data stored in dst.AuthenticatorStaticChallenge, return on the first match
		} else {
			dst.AuthenticatorStaticChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as AuthenticatorStaticChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorTOTPChallenge'
	if jsonDict["component"] == "AuthenticatorTOTPChallenge" {
		// try to unmarshal JSON data into AuthenticatorTOTPChallenge
		err = json.Unmarshal(data, &dst.AuthenticatorTOTPChallenge)
		if err == nil {
			return nil // data stored in dst.AuthenticatorTOTPChallenge, return on the first match
		} else {
			dst.AuthenticatorTOTPChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as AuthenticatorTOTPChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorValidationChallenge'
	if jsonDict["component"] == "AuthenticatorValidationChallenge" {
		// try to unmarshal JSON data into AuthenticatorValidationChallenge
		err = json.Unmarshal(data, &dst.AuthenticatorValidationChallenge)
		if err == nil {
			return nil // data stored in dst.AuthenticatorValidationChallenge, return on the first match
		} else {
			dst.AuthenticatorValidationChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as AuthenticatorValidationChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthenticatorWebAuthnChallenge'
	if jsonDict["component"] == "AuthenticatorWebAuthnChallenge" {
		// try to unmarshal JSON data into AuthenticatorWebAuthnChallenge
		err = json.Unmarshal(data, &dst.AuthenticatorWebAuthnChallenge)
		if err == nil {
			return nil // data stored in dst.AuthenticatorWebAuthnChallenge, return on the first match
		} else {
			dst.AuthenticatorWebAuthnChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as AuthenticatorWebAuthnChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AutosubmitChallenge'
	if jsonDict["component"] == "AutosubmitChallenge" {
		// try to unmarshal JSON data into AutosubmitChallenge
		err = json.Unmarshal(data, &dst.AutosubmitChallenge)
		if err == nil {
			return nil // data stored in dst.AutosubmitChallenge, return on the first match
		} else {
			dst.AutosubmitChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as AutosubmitChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CaptchaChallenge'
	if jsonDict["component"] == "CaptchaChallenge" {
		// try to unmarshal JSON data into CaptchaChallenge
		err = json.Unmarshal(data, &dst.CaptchaChallenge)
		if err == nil {
			return nil // data stored in dst.CaptchaChallenge, return on the first match
		} else {
			dst.CaptchaChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as CaptchaChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ConsentChallenge'
	if jsonDict["component"] == "ConsentChallenge" {
		// try to unmarshal JSON data into ConsentChallenge
		err = json.Unmarshal(data, &dst.ConsentChallenge)
		if err == nil {
			return nil // data stored in dst.ConsentChallenge, return on the first match
		} else {
			dst.ConsentChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as ConsentChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DummyChallenge'
	if jsonDict["component"] == "DummyChallenge" {
		// try to unmarshal JSON data into DummyChallenge
		err = json.Unmarshal(data, &dst.DummyChallenge)
		if err == nil {
			return nil // data stored in dst.DummyChallenge, return on the first match
		} else {
			dst.DummyChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as DummyChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EmailChallenge'
	if jsonDict["component"] == "EmailChallenge" {
		// try to unmarshal JSON data into EmailChallenge
		err = json.Unmarshal(data, &dst.EmailChallenge)
		if err == nil {
			return nil // data stored in dst.EmailChallenge, return on the first match
		} else {
			dst.EmailChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as EmailChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'IdentificationChallenge'
	if jsonDict["component"] == "IdentificationChallenge" {
		// try to unmarshal JSON data into IdentificationChallenge
		err = json.Unmarshal(data, &dst.IdentificationChallenge)
		if err == nil {
			return nil // data stored in dst.IdentificationChallenge, return on the first match
		} else {
			dst.IdentificationChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as IdentificationChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PasswordChallenge'
	if jsonDict["component"] == "PasswordChallenge" {
		// try to unmarshal JSON data into PasswordChallenge
		err = json.Unmarshal(data, &dst.PasswordChallenge)
		if err == nil {
			return nil // data stored in dst.PasswordChallenge, return on the first match
		} else {
			dst.PasswordChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as PasswordChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PlexAuthenticationChallenge'
	if jsonDict["component"] == "PlexAuthenticationChallenge" {
		// try to unmarshal JSON data into PlexAuthenticationChallenge
		err = json.Unmarshal(data, &dst.PlexAuthenticationChallenge)
		if err == nil {
			return nil // data stored in dst.PlexAuthenticationChallenge, return on the first match
		} else {
			dst.PlexAuthenticationChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as PlexAuthenticationChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PromptChallenge'
	if jsonDict["component"] == "PromptChallenge" {
		// try to unmarshal JSON data into PromptChallenge
		err = json.Unmarshal(data, &dst.PromptChallenge)
		if err == nil {
			return nil // data stored in dst.PromptChallenge, return on the first match
		} else {
			dst.PromptChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as PromptChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RedirectChallenge'
	if jsonDict["component"] == "RedirectChallenge" {
		// try to unmarshal JSON data into RedirectChallenge
		err = json.Unmarshal(data, &dst.RedirectChallenge)
		if err == nil {
			return nil // data stored in dst.RedirectChallenge, return on the first match
		} else {
			dst.RedirectChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as RedirectChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ShellChallenge'
	if jsonDict["component"] == "ShellChallenge" {
		// try to unmarshal JSON data into ShellChallenge
		err = json.Unmarshal(data, &dst.ShellChallenge)
		if err == nil {
			return nil // data stored in dst.ShellChallenge, return on the first match
		} else {
			dst.ShellChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as ShellChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-flow-sources-plex'
	if jsonDict["component"] == "ak-flow-sources-plex" {
		// try to unmarshal JSON data into PlexAuthenticationChallenge
		err = json.Unmarshal(data, &dst.PlexAuthenticationChallenge)
		if err == nil {
			return nil // data stored in dst.PlexAuthenticationChallenge, return on the first match
		} else {
			dst.PlexAuthenticationChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as PlexAuthenticationChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-access-denied'
	if jsonDict["component"] == "ak-stage-access-denied" {
		// try to unmarshal JSON data into AccessDeniedChallenge
		err = json.Unmarshal(data, &dst.AccessDeniedChallenge)
		if err == nil {
			return nil // data stored in dst.AccessDeniedChallenge, return on the first match
		} else {
			dst.AccessDeniedChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as AccessDeniedChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-authenticator-duo'
	if jsonDict["component"] == "ak-stage-authenticator-duo" {
		// try to unmarshal JSON data into AuthenticatorDuoChallenge
		err = json.Unmarshal(data, &dst.AuthenticatorDuoChallenge)
		if err == nil {
			return nil // data stored in dst.AuthenticatorDuoChallenge, return on the first match
		} else {
			dst.AuthenticatorDuoChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as AuthenticatorDuoChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-authenticator-sms'
	if jsonDict["component"] == "ak-stage-authenticator-sms" {
		// try to unmarshal JSON data into AuthenticatorSMSChallenge
		err = json.Unmarshal(data, &dst.AuthenticatorSMSChallenge)
		if err == nil {
			return nil // data stored in dst.AuthenticatorSMSChallenge, return on the first match
		} else {
			dst.AuthenticatorSMSChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as AuthenticatorSMSChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-authenticator-static'
	if jsonDict["component"] == "ak-stage-authenticator-static" {
		// try to unmarshal JSON data into AuthenticatorStaticChallenge
		err = json.Unmarshal(data, &dst.AuthenticatorStaticChallenge)
		if err == nil {
			return nil // data stored in dst.AuthenticatorStaticChallenge, return on the first match
		} else {
			dst.AuthenticatorStaticChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as AuthenticatorStaticChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-authenticator-totp'
	if jsonDict["component"] == "ak-stage-authenticator-totp" {
		// try to unmarshal JSON data into AuthenticatorTOTPChallenge
		err = json.Unmarshal(data, &dst.AuthenticatorTOTPChallenge)
		if err == nil {
			return nil // data stored in dst.AuthenticatorTOTPChallenge, return on the first match
		} else {
			dst.AuthenticatorTOTPChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as AuthenticatorTOTPChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-authenticator-validate'
	if jsonDict["component"] == "ak-stage-authenticator-validate" {
		// try to unmarshal JSON data into AuthenticatorValidationChallenge
		err = json.Unmarshal(data, &dst.AuthenticatorValidationChallenge)
		if err == nil {
			return nil // data stored in dst.AuthenticatorValidationChallenge, return on the first match
		} else {
			dst.AuthenticatorValidationChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as AuthenticatorValidationChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-authenticator-webauthn'
	if jsonDict["component"] == "ak-stage-authenticator-webauthn" {
		// try to unmarshal JSON data into AuthenticatorWebAuthnChallenge
		err = json.Unmarshal(data, &dst.AuthenticatorWebAuthnChallenge)
		if err == nil {
			return nil // data stored in dst.AuthenticatorWebAuthnChallenge, return on the first match
		} else {
			dst.AuthenticatorWebAuthnChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as AuthenticatorWebAuthnChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-autosubmit'
	if jsonDict["component"] == "ak-stage-autosubmit" {
		// try to unmarshal JSON data into AutosubmitChallenge
		err = json.Unmarshal(data, &dst.AutosubmitChallenge)
		if err == nil {
			return nil // data stored in dst.AutosubmitChallenge, return on the first match
		} else {
			dst.AutosubmitChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as AutosubmitChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-captcha'
	if jsonDict["component"] == "ak-stage-captcha" {
		// try to unmarshal JSON data into CaptchaChallenge
		err = json.Unmarshal(data, &dst.CaptchaChallenge)
		if err == nil {
			return nil // data stored in dst.CaptchaChallenge, return on the first match
		} else {
			dst.CaptchaChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as CaptchaChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-consent'
	if jsonDict["component"] == "ak-stage-consent" {
		// try to unmarshal JSON data into ConsentChallenge
		err = json.Unmarshal(data, &dst.ConsentChallenge)
		if err == nil {
			return nil // data stored in dst.ConsentChallenge, return on the first match
		} else {
			dst.ConsentChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as ConsentChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-dummy'
	if jsonDict["component"] == "ak-stage-dummy" {
		// try to unmarshal JSON data into DummyChallenge
		err = json.Unmarshal(data, &dst.DummyChallenge)
		if err == nil {
			return nil // data stored in dst.DummyChallenge, return on the first match
		} else {
			dst.DummyChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as DummyChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-email'
	if jsonDict["component"] == "ak-stage-email" {
		// try to unmarshal JSON data into EmailChallenge
		err = json.Unmarshal(data, &dst.EmailChallenge)
		if err == nil {
			return nil // data stored in dst.EmailChallenge, return on the first match
		} else {
			dst.EmailChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as EmailChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-identification'
	if jsonDict["component"] == "ak-stage-identification" {
		// try to unmarshal JSON data into IdentificationChallenge
		err = json.Unmarshal(data, &dst.IdentificationChallenge)
		if err == nil {
			return nil // data stored in dst.IdentificationChallenge, return on the first match
		} else {
			dst.IdentificationChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as IdentificationChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-password'
	if jsonDict["component"] == "ak-stage-password" {
		// try to unmarshal JSON data into PasswordChallenge
		err = json.Unmarshal(data, &dst.PasswordChallenge)
		if err == nil {
			return nil // data stored in dst.PasswordChallenge, return on the first match
		} else {
			dst.PasswordChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as PasswordChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-stage-prompt'
	if jsonDict["component"] == "ak-stage-prompt" {
		// try to unmarshal JSON data into PromptChallenge
		err = json.Unmarshal(data, &dst.PromptChallenge)
		if err == nil {
			return nil // data stored in dst.PromptChallenge, return on the first match
		} else {
			dst.PromptChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as PromptChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'xak-flow-redirect'
	if jsonDict["component"] == "xak-flow-redirect" {
		// try to unmarshal JSON data into RedirectChallenge
		err = json.Unmarshal(data, &dst.RedirectChallenge)
		if err == nil {
			return nil // data stored in dst.RedirectChallenge, return on the first match
		} else {
			dst.RedirectChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as RedirectChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'xak-flow-shell'
	if jsonDict["component"] == "xak-flow-shell" {
		// try to unmarshal JSON data into ShellChallenge
		err = json.Unmarshal(data, &dst.ShellChallenge)
		if err == nil {
			return nil // data stored in dst.ShellChallenge, return on the first match
		} else {
			dst.ShellChallenge = nil
			return fmt.Errorf("Failed to unmarshal ChallengeTypes as ShellChallenge: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ChallengeTypes) MarshalJSON() ([]byte, error) {
	if src.AccessDeniedChallenge != nil {
		return json.Marshal(&src.AccessDeniedChallenge)
	}

	if src.AuthenticatorDuoChallenge != nil {
		return json.Marshal(&src.AuthenticatorDuoChallenge)
	}

	if src.AuthenticatorSMSChallenge != nil {
		return json.Marshal(&src.AuthenticatorSMSChallenge)
	}

	if src.AuthenticatorStaticChallenge != nil {
		return json.Marshal(&src.AuthenticatorStaticChallenge)
	}

	if src.AuthenticatorTOTPChallenge != nil {
		return json.Marshal(&src.AuthenticatorTOTPChallenge)
	}

	if src.AuthenticatorValidationChallenge != nil {
		return json.Marshal(&src.AuthenticatorValidationChallenge)
	}

	if src.AuthenticatorWebAuthnChallenge != nil {
		return json.Marshal(&src.AuthenticatorWebAuthnChallenge)
	}

	if src.AutosubmitChallenge != nil {
		return json.Marshal(&src.AutosubmitChallenge)
	}

	if src.CaptchaChallenge != nil {
		return json.Marshal(&src.CaptchaChallenge)
	}

	if src.ConsentChallenge != nil {
		return json.Marshal(&src.ConsentChallenge)
	}

	if src.DummyChallenge != nil {
		return json.Marshal(&src.DummyChallenge)
	}

	if src.EmailChallenge != nil {
		return json.Marshal(&src.EmailChallenge)
	}

	if src.IdentificationChallenge != nil {
		return json.Marshal(&src.IdentificationChallenge)
	}

	if src.PasswordChallenge != nil {
		return json.Marshal(&src.PasswordChallenge)
	}

	if src.PlexAuthenticationChallenge != nil {
		return json.Marshal(&src.PlexAuthenticationChallenge)
	}

	if src.PromptChallenge != nil {
		return json.Marshal(&src.PromptChallenge)
	}

	if src.RedirectChallenge != nil {
		return json.Marshal(&src.RedirectChallenge)
	}

	if src.ShellChallenge != nil {
		return json.Marshal(&src.ShellChallenge)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ChallengeTypes) GetActualInstance() interface{} {
	if obj.AccessDeniedChallenge != nil {
		return obj.AccessDeniedChallenge
	}

	if obj.AuthenticatorDuoChallenge != nil {
		return obj.AuthenticatorDuoChallenge
	}

	if obj.AuthenticatorSMSChallenge != nil {
		return obj.AuthenticatorSMSChallenge
	}

	if obj.AuthenticatorStaticChallenge != nil {
		return obj.AuthenticatorStaticChallenge
	}

	if obj.AuthenticatorTOTPChallenge != nil {
		return obj.AuthenticatorTOTPChallenge
	}

	if obj.AuthenticatorValidationChallenge != nil {
		return obj.AuthenticatorValidationChallenge
	}

	if obj.AuthenticatorWebAuthnChallenge != nil {
		return obj.AuthenticatorWebAuthnChallenge
	}

	if obj.AutosubmitChallenge != nil {
		return obj.AutosubmitChallenge
	}

	if obj.CaptchaChallenge != nil {
		return obj.CaptchaChallenge
	}

	if obj.ConsentChallenge != nil {
		return obj.ConsentChallenge
	}

	if obj.DummyChallenge != nil {
		return obj.DummyChallenge
	}

	if obj.EmailChallenge != nil {
		return obj.EmailChallenge
	}

	if obj.IdentificationChallenge != nil {
		return obj.IdentificationChallenge
	}

	if obj.PasswordChallenge != nil {
		return obj.PasswordChallenge
	}

	if obj.PlexAuthenticationChallenge != nil {
		return obj.PlexAuthenticationChallenge
	}

	if obj.PromptChallenge != nil {
		return obj.PromptChallenge
	}

	if obj.RedirectChallenge != nil {
		return obj.RedirectChallenge
	}

	if obj.ShellChallenge != nil {
		return obj.ShellChallenge
	}

	// all schemas are nil
	return nil
}

type NullableChallengeTypes struct {
	value *ChallengeTypes
	isSet bool
}

func (v NullableChallengeTypes) Get() *ChallengeTypes {
	return v.value
}

func (v *NullableChallengeTypes) Set(val *ChallengeTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableChallengeTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableChallengeTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChallengeTypes(val *ChallengeTypes) *NullableChallengeTypes {
	return &NullableChallengeTypes{value: val, isSet: true}
}

func (v NullableChallengeTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChallengeTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
