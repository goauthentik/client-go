/*
authentik

Making authentication simple.

API version: 2023.8.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// AppEnum * `authentik.admin` - authentik Admin * `authentik.api` - authentik API * `authentik.crypto` - authentik Crypto * `authentik.events` - authentik Events * `authentik.flows` - authentik Flows * `authentik.outposts` - authentik Outpost * `authentik.policies.dummy` - authentik Policies.Dummy * `authentik.policies.event_matcher` - authentik Policies.Event Matcher * `authentik.policies.expiry` - authentik Policies.Expiry * `authentik.policies.expression` - authentik Policies.Expression * `authentik.policies.password` - authentik Policies.Password * `authentik.policies.reputation` - authentik Policies.Reputation * `authentik.policies` - authentik Policies * `authentik.providers.ldap` - authentik Providers.LDAP * `authentik.providers.oauth2` - authentik Providers.OAuth2 * `authentik.providers.proxy` - authentik Providers.Proxy * `authentik.providers.radius` - authentik Providers.Radius * `authentik.providers.saml` - authentik Providers.SAML * `authentik.providers.scim` - authentik Providers.SCIM * `authentik.rbac` - authentik RBAC * `authentik.recovery` - authentik Recovery * `authentik.sources.ldap` - authentik Sources.LDAP * `authentik.sources.oauth` - authentik Sources.OAuth * `authentik.sources.plex` - authentik Sources.Plex * `authentik.sources.saml` - authentik Sources.SAML * `authentik.stages.authenticator` - authentik Stages.Authenticator * `authentik.stages.authenticator_duo` - authentik Stages.Authenticator.Duo * `authentik.stages.authenticator_sms` - authentik Stages.Authenticator.SMS * `authentik.stages.authenticator_static` - authentik Stages.Authenticator.Static * `authentik.stages.authenticator_totp` - authentik Stages.Authenticator.TOTP * `authentik.stages.authenticator_validate` - authentik Stages.Authenticator.Validate * `authentik.stages.authenticator_webauthn` - authentik Stages.Authenticator.WebAuthn * `authentik.stages.captcha` - authentik Stages.Captcha * `authentik.stages.consent` - authentik Stages.Consent * `authentik.stages.deny` - authentik Stages.Deny * `authentik.stages.dummy` - authentik Stages.Dummy * `authentik.stages.email` - authentik Stages.Email * `authentik.stages.identification` - authentik Stages.Identification * `authentik.stages.invitation` - authentik Stages.User Invitation * `authentik.stages.password` - authentik Stages.Password * `authentik.stages.prompt` - authentik Stages.Prompt * `authentik.stages.user_delete` - authentik Stages.User Delete * `authentik.stages.user_login` - authentik Stages.User Login * `authentik.stages.user_logout` - authentik Stages.User Logout * `authentik.stages.user_write` - authentik Stages.User Write * `authentik.tenants` - authentik Tenants * `authentik.blueprints` - authentik Blueprints * `authentik.core` - authentik Core * `authentik.enterprise` - authentik Enterprise
type AppEnum string

// List of AppEnum
const (
	APPENUM_ADMIN                         AppEnum = "authentik.admin"
	APPENUM_API                           AppEnum = "authentik.api"
	APPENUM_CRYPTO                        AppEnum = "authentik.crypto"
	APPENUM_EVENTS                        AppEnum = "authentik.events"
	APPENUM_FLOWS                         AppEnum = "authentik.flows"
	APPENUM_OUTPOSTS                      AppEnum = "authentik.outposts"
	APPENUM_POLICIES_DUMMY                AppEnum = "authentik.policies.dummy"
	APPENUM_POLICIES_EVENT_MATCHER        AppEnum = "authentik.policies.event_matcher"
	APPENUM_POLICIES_EXPIRY               AppEnum = "authentik.policies.expiry"
	APPENUM_POLICIES_EXPRESSION           AppEnum = "authentik.policies.expression"
	APPENUM_POLICIES_PASSWORD             AppEnum = "authentik.policies.password"
	APPENUM_POLICIES_REPUTATION           AppEnum = "authentik.policies.reputation"
	APPENUM_POLICIES                      AppEnum = "authentik.policies"
	APPENUM_PROVIDERS_LDAP                AppEnum = "authentik.providers.ldap"
	APPENUM_PROVIDERS_OAUTH2              AppEnum = "authentik.providers.oauth2"
	APPENUM_PROVIDERS_PROXY               AppEnum = "authentik.providers.proxy"
	APPENUM_PROVIDERS_RADIUS              AppEnum = "authentik.providers.radius"
	APPENUM_PROVIDERS_SAML                AppEnum = "authentik.providers.saml"
	APPENUM_PROVIDERS_SCIM                AppEnum = "authentik.providers.scim"
	APPENUM_RBAC                          AppEnum = "authentik.rbac"
	APPENUM_RECOVERY                      AppEnum = "authentik.recovery"
	APPENUM_SOURCES_LDAP                  AppEnum = "authentik.sources.ldap"
	APPENUM_SOURCES_OAUTH                 AppEnum = "authentik.sources.oauth"
	APPENUM_SOURCES_PLEX                  AppEnum = "authentik.sources.plex"
	APPENUM_SOURCES_SAML                  AppEnum = "authentik.sources.saml"
	APPENUM_STAGES_AUTHENTICATOR          AppEnum = "authentik.stages.authenticator"
	APPENUM_STAGES_AUTHENTICATOR_DUO      AppEnum = "authentik.stages.authenticator_duo"
	APPENUM_STAGES_AUTHENTICATOR_SMS      AppEnum = "authentik.stages.authenticator_sms"
	APPENUM_STAGES_AUTHENTICATOR_STATIC   AppEnum = "authentik.stages.authenticator_static"
	APPENUM_STAGES_AUTHENTICATOR_TOTP     AppEnum = "authentik.stages.authenticator_totp"
	APPENUM_STAGES_AUTHENTICATOR_VALIDATE AppEnum = "authentik.stages.authenticator_validate"
	APPENUM_STAGES_AUTHENTICATOR_WEBAUTHN AppEnum = "authentik.stages.authenticator_webauthn"
	APPENUM_STAGES_CAPTCHA                AppEnum = "authentik.stages.captcha"
	APPENUM_STAGES_CONSENT                AppEnum = "authentik.stages.consent"
	APPENUM_STAGES_DENY                   AppEnum = "authentik.stages.deny"
	APPENUM_STAGES_DUMMY                  AppEnum = "authentik.stages.dummy"
	APPENUM_STAGES_EMAIL                  AppEnum = "authentik.stages.email"
	APPENUM_STAGES_IDENTIFICATION         AppEnum = "authentik.stages.identification"
	APPENUM_STAGES_INVITATION             AppEnum = "authentik.stages.invitation"
	APPENUM_STAGES_PASSWORD               AppEnum = "authentik.stages.password"
	APPENUM_STAGES_PROMPT                 AppEnum = "authentik.stages.prompt"
	APPENUM_STAGES_USER_DELETE            AppEnum = "authentik.stages.user_delete"
	APPENUM_STAGES_USER_LOGIN             AppEnum = "authentik.stages.user_login"
	APPENUM_STAGES_USER_LOGOUT            AppEnum = "authentik.stages.user_logout"
	APPENUM_STAGES_USER_WRITE             AppEnum = "authentik.stages.user_write"
	APPENUM_TENANTS                       AppEnum = "authentik.tenants"
	APPENUM_BLUEPRINTS                    AppEnum = "authentik.blueprints"
	APPENUM_CORE                          AppEnum = "authentik.core"
	APPENUM_ENTERPRISE                    AppEnum = "authentik.enterprise"
)

// All allowed values of AppEnum enum
var AllowedAppEnumEnumValues = []AppEnum{
	"authentik.admin",
	"authentik.api",
	"authentik.crypto",
	"authentik.events",
	"authentik.flows",
	"authentik.outposts",
	"authentik.policies.dummy",
	"authentik.policies.event_matcher",
	"authentik.policies.expiry",
	"authentik.policies.expression",
	"authentik.policies.password",
	"authentik.policies.reputation",
	"authentik.policies",
	"authentik.providers.ldap",
	"authentik.providers.oauth2",
	"authentik.providers.proxy",
	"authentik.providers.radius",
	"authentik.providers.saml",
	"authentik.providers.scim",
	"authentik.rbac",
	"authentik.recovery",
	"authentik.sources.ldap",
	"authentik.sources.oauth",
	"authentik.sources.plex",
	"authentik.sources.saml",
	"authentik.stages.authenticator",
	"authentik.stages.authenticator_duo",
	"authentik.stages.authenticator_sms",
	"authentik.stages.authenticator_static",
	"authentik.stages.authenticator_totp",
	"authentik.stages.authenticator_validate",
	"authentik.stages.authenticator_webauthn",
	"authentik.stages.captcha",
	"authentik.stages.consent",
	"authentik.stages.deny",
	"authentik.stages.dummy",
	"authentik.stages.email",
	"authentik.stages.identification",
	"authentik.stages.invitation",
	"authentik.stages.password",
	"authentik.stages.prompt",
	"authentik.stages.user_delete",
	"authentik.stages.user_login",
	"authentik.stages.user_logout",
	"authentik.stages.user_write",
	"authentik.tenants",
	"authentik.blueprints",
	"authentik.core",
	"authentik.enterprise",
}

func (v *AppEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppEnum(value)
	for _, existing := range AllowedAppEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppEnum", value)
}

// NewAppEnumFromValue returns a pointer to a valid AppEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppEnumFromValue(v string) (*AppEnum, error) {
	ev := AppEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppEnum: valid values are %v", v, AllowedAppEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppEnum) IsValid() bool {
	for _, existing := range AllowedAppEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppEnum value
func (v AppEnum) Ptr() *AppEnum {
	return &v
}

type NullableAppEnum struct {
	value *AppEnum
	isSet bool
}

func (v NullableAppEnum) Get() *AppEnum {
	return v.value
}

func (v *NullableAppEnum) Set(val *AppEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEnum(val *AppEnum) *NullableAppEnum {
	return &NullableAppEnum{value: val, isSet: true}
}

func (v NullableAppEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
