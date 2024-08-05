/*
authentik

Making authentication simple.

API version: 2024.6.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ModelEnum the model 'ModelEnum'
type ModelEnum string

// List of ModelEnum
const (
	MODELENUM_TENANTS_DOMAIN                                            ModelEnum = "authentik_tenants.domain"
	MODELENUM_CRYPTO_CERTIFICATEKEYPAIR                                 ModelEnum = "authentik_crypto.certificatekeypair"
	MODELENUM_FLOWS_FLOW                                                ModelEnum = "authentik_flows.flow"
	MODELENUM_FLOWS_FLOWSTAGEBINDING                                    ModelEnum = "authentik_flows.flowstagebinding"
	MODELENUM_OUTPOSTS_DOCKERSERVICECONNECTION                          ModelEnum = "authentik_outposts.dockerserviceconnection"
	MODELENUM_OUTPOSTS_KUBERNETESSERVICECONNECTION                      ModelEnum = "authentik_outposts.kubernetesserviceconnection"
	MODELENUM_OUTPOSTS_OUTPOST                                          ModelEnum = "authentik_outposts.outpost"
	MODELENUM_POLICIES_DUMMY_DUMMYPOLICY                                ModelEnum = "authentik_policies_dummy.dummypolicy"
	MODELENUM_POLICIES_EVENT_MATCHER_EVENTMATCHERPOLICY                 ModelEnum = "authentik_policies_event_matcher.eventmatcherpolicy"
	MODELENUM_POLICIES_EXPIRY_PASSWORDEXPIRYPOLICY                      ModelEnum = "authentik_policies_expiry.passwordexpirypolicy"
	MODELENUM_POLICIES_EXPRESSION_EXPRESSIONPOLICY                      ModelEnum = "authentik_policies_expression.expressionpolicy"
	MODELENUM_POLICIES_PASSWORD_PASSWORDPOLICY                          ModelEnum = "authentik_policies_password.passwordpolicy"
	MODELENUM_POLICIES_REPUTATION_REPUTATIONPOLICY                      ModelEnum = "authentik_policies_reputation.reputationpolicy"
	MODELENUM_POLICIES_POLICYBINDING                                    ModelEnum = "authentik_policies.policybinding"
	MODELENUM_PROVIDERS_LDAP_LDAPPROVIDER                               ModelEnum = "authentik_providers_ldap.ldapprovider"
	MODELENUM_PROVIDERS_OAUTH2_SCOPEMAPPING                             ModelEnum = "authentik_providers_oauth2.scopemapping"
	MODELENUM_PROVIDERS_OAUTH2_OAUTH2PROVIDER                           ModelEnum = "authentik_providers_oauth2.oauth2provider"
	MODELENUM_PROVIDERS_PROXY_PROXYPROVIDER                             ModelEnum = "authentik_providers_proxy.proxyprovider"
	MODELENUM_PROVIDERS_RADIUS_RADIUSPROVIDER                           ModelEnum = "authentik_providers_radius.radiusprovider"
	MODELENUM_PROVIDERS_RADIUS_RADIUSPROVIDERPROPERTYMAPPING            ModelEnum = "authentik_providers_radius.radiusproviderpropertymapping"
	MODELENUM_PROVIDERS_SAML_SAMLPROVIDER                               ModelEnum = "authentik_providers_saml.samlprovider"
	MODELENUM_PROVIDERS_SAML_SAMLPROPERTYMAPPING                        ModelEnum = "authentik_providers_saml.samlpropertymapping"
	MODELENUM_PROVIDERS_SCIM_SCIMPROVIDER                               ModelEnum = "authentik_providers_scim.scimprovider"
	MODELENUM_PROVIDERS_SCIM_SCIMMAPPING                                ModelEnum = "authentik_providers_scim.scimmapping"
	MODELENUM_RBAC_ROLE                                                 ModelEnum = "authentik_rbac.role"
	MODELENUM_SOURCES_LDAP_LDAPSOURCE                                   ModelEnum = "authentik_sources_ldap.ldapsource"
	MODELENUM_SOURCES_LDAP_LDAPSOURCEPROPERTYMAPPING                    ModelEnum = "authentik_sources_ldap.ldapsourcepropertymapping"
	MODELENUM_SOURCES_OAUTH_OAUTHSOURCE                                 ModelEnum = "authentik_sources_oauth.oauthsource"
	MODELENUM_SOURCES_OAUTH_USEROAUTHSOURCECONNECTION                   ModelEnum = "authentik_sources_oauth.useroauthsourceconnection"
	MODELENUM_SOURCES_PLEX_PLEXSOURCE                                   ModelEnum = "authentik_sources_plex.plexsource"
	MODELENUM_SOURCES_PLEX_PLEXSOURCECONNECTION                         ModelEnum = "authentik_sources_plex.plexsourceconnection"
	MODELENUM_SOURCES_SAML_SAMLSOURCE                                   ModelEnum = "authentik_sources_saml.samlsource"
	MODELENUM_SOURCES_SAML_USERSAMLSOURCECONNECTION                     ModelEnum = "authentik_sources_saml.usersamlsourceconnection"
	MODELENUM_SOURCES_SCIM_SCIMSOURCE                                   ModelEnum = "authentik_sources_scim.scimsource"
	MODELENUM_SOURCES_SCIM_SCIMSOURCEPROPERTYMAPPING                    ModelEnum = "authentik_sources_scim.scimsourcepropertymapping"
	MODELENUM_STAGES_AUTHENTICATOR_DUO_AUTHENTICATORDUOSTAGE            ModelEnum = "authentik_stages_authenticator_duo.authenticatorduostage"
	MODELENUM_STAGES_AUTHENTICATOR_DUO_DUODEVICE                        ModelEnum = "authentik_stages_authenticator_duo.duodevice"
	MODELENUM_STAGES_AUTHENTICATOR_SMS_AUTHENTICATORSMSSTAGE            ModelEnum = "authentik_stages_authenticator_sms.authenticatorsmsstage"
	MODELENUM_STAGES_AUTHENTICATOR_SMS_SMSDEVICE                        ModelEnum = "authentik_stages_authenticator_sms.smsdevice"
	MODELENUM_STAGES_AUTHENTICATOR_STATIC_AUTHENTICATORSTATICSTAGE      ModelEnum = "authentik_stages_authenticator_static.authenticatorstaticstage"
	MODELENUM_STAGES_AUTHENTICATOR_STATIC_STATICDEVICE                  ModelEnum = "authentik_stages_authenticator_static.staticdevice"
	MODELENUM_STAGES_AUTHENTICATOR_TOTP_AUTHENTICATORTOTPSTAGE          ModelEnum = "authentik_stages_authenticator_totp.authenticatortotpstage"
	MODELENUM_STAGES_AUTHENTICATOR_TOTP_TOTPDEVICE                      ModelEnum = "authentik_stages_authenticator_totp.totpdevice"
	MODELENUM_STAGES_AUTHENTICATOR_VALIDATE_AUTHENTICATORVALIDATESTAGE  ModelEnum = "authentik_stages_authenticator_validate.authenticatorvalidatestage"
	MODELENUM_STAGES_AUTHENTICATOR_WEBAUTHN_AUTHENTICATORWEBAUTHNSTAGE  ModelEnum = "authentik_stages_authenticator_webauthn.authenticatorwebauthnstage"
	MODELENUM_STAGES_AUTHENTICATOR_WEBAUTHN_WEBAUTHNDEVICE              ModelEnum = "authentik_stages_authenticator_webauthn.webauthndevice"
	MODELENUM_STAGES_CAPTCHA_CAPTCHASTAGE                               ModelEnum = "authentik_stages_captcha.captchastage"
	MODELENUM_STAGES_CONSENT_CONSENTSTAGE                               ModelEnum = "authentik_stages_consent.consentstage"
	MODELENUM_STAGES_CONSENT_USERCONSENT                                ModelEnum = "authentik_stages_consent.userconsent"
	MODELENUM_STAGES_DENY_DENYSTAGE                                     ModelEnum = "authentik_stages_deny.denystage"
	MODELENUM_STAGES_DUMMY_DUMMYSTAGE                                   ModelEnum = "authentik_stages_dummy.dummystage"
	MODELENUM_STAGES_EMAIL_EMAILSTAGE                                   ModelEnum = "authentik_stages_email.emailstage"
	MODELENUM_STAGES_IDENTIFICATION_IDENTIFICATIONSTAGE                 ModelEnum = "authentik_stages_identification.identificationstage"
	MODELENUM_STAGES_INVITATION_INVITATIONSTAGE                         ModelEnum = "authentik_stages_invitation.invitationstage"
	MODELENUM_STAGES_INVITATION_INVITATION                              ModelEnum = "authentik_stages_invitation.invitation"
	MODELENUM_STAGES_PASSWORD_PASSWORDSTAGE                             ModelEnum = "authentik_stages_password.passwordstage"
	MODELENUM_STAGES_PROMPT_PROMPT                                      ModelEnum = "authentik_stages_prompt.prompt"
	MODELENUM_STAGES_PROMPT_PROMPTSTAGE                                 ModelEnum = "authentik_stages_prompt.promptstage"
	MODELENUM_STAGES_USER_DELETE_USERDELETESTAGE                        ModelEnum = "authentik_stages_user_delete.userdeletestage"
	MODELENUM_STAGES_USER_LOGIN_USERLOGINSTAGE                          ModelEnum = "authentik_stages_user_login.userloginstage"
	MODELENUM_STAGES_USER_LOGOUT_USERLOGOUTSTAGE                        ModelEnum = "authentik_stages_user_logout.userlogoutstage"
	MODELENUM_STAGES_USER_WRITE_USERWRITESTAGE                          ModelEnum = "authentik_stages_user_write.userwritestage"
	MODELENUM_BRANDS_BRAND                                              ModelEnum = "authentik_brands.brand"
	MODELENUM_BLUEPRINTS_BLUEPRINTINSTANCE                              ModelEnum = "authentik_blueprints.blueprintinstance"
	MODELENUM_CORE_GROUP                                                ModelEnum = "authentik_core.group"
	MODELENUM_CORE_USER                                                 ModelEnum = "authentik_core.user"
	MODELENUM_CORE_APPLICATION                                          ModelEnum = "authentik_core.application"
	MODELENUM_CORE_TOKEN                                                ModelEnum = "authentik_core.token"
	MODELENUM_ENTERPRISE_LICENSE                                        ModelEnum = "authentik_enterprise.license"
	MODELENUM_PROVIDERS_GOOGLE_WORKSPACE_GOOGLEWORKSPACEPROVIDER        ModelEnum = "authentik_providers_google_workspace.googleworkspaceprovider"
	MODELENUM_PROVIDERS_GOOGLE_WORKSPACE_GOOGLEWORKSPACEPROVIDERMAPPING ModelEnum = "authentik_providers_google_workspace.googleworkspaceprovidermapping"
	MODELENUM_PROVIDERS_MICROSOFT_ENTRA_MICROSOFTENTRAPROVIDER          ModelEnum = "authentik_providers_microsoft_entra.microsoftentraprovider"
	MODELENUM_PROVIDERS_MICROSOFT_ENTRA_MICROSOFTENTRAPROVIDERMAPPING   ModelEnum = "authentik_providers_microsoft_entra.microsoftentraprovidermapping"
	MODELENUM_PROVIDERS_RAC_RACPROVIDER                                 ModelEnum = "authentik_providers_rac.racprovider"
	MODELENUM_PROVIDERS_RAC_ENDPOINT                                    ModelEnum = "authentik_providers_rac.endpoint"
	MODELENUM_PROVIDERS_RAC_RACPROPERTYMAPPING                          ModelEnum = "authentik_providers_rac.racpropertymapping"
	MODELENUM_STAGES_SOURCE_SOURCESTAGE                                 ModelEnum = "authentik_stages_source.sourcestage"
	MODELENUM_EVENTS_EVENT                                              ModelEnum = "authentik_events.event"
	MODELENUM_EVENTS_NOTIFICATIONTRANSPORT                              ModelEnum = "authentik_events.notificationtransport"
	MODELENUM_EVENTS_NOTIFICATION                                       ModelEnum = "authentik_events.notification"
	MODELENUM_EVENTS_NOTIFICATIONRULE                                   ModelEnum = "authentik_events.notificationrule"
	MODELENUM_EVENTS_NOTIFICATIONWEBHOOKMAPPING                         ModelEnum = "authentik_events.notificationwebhookmapping"
)

// All allowed values of ModelEnum enum
var AllowedModelEnumEnumValues = []ModelEnum{
	"authentik_tenants.domain",
	"authentik_crypto.certificatekeypair",
	"authentik_flows.flow",
	"authentik_flows.flowstagebinding",
	"authentik_outposts.dockerserviceconnection",
	"authentik_outposts.kubernetesserviceconnection",
	"authentik_outposts.outpost",
	"authentik_policies_dummy.dummypolicy",
	"authentik_policies_event_matcher.eventmatcherpolicy",
	"authentik_policies_expiry.passwordexpirypolicy",
	"authentik_policies_expression.expressionpolicy",
	"authentik_policies_password.passwordpolicy",
	"authentik_policies_reputation.reputationpolicy",
	"authentik_policies.policybinding",
	"authentik_providers_ldap.ldapprovider",
	"authentik_providers_oauth2.scopemapping",
	"authentik_providers_oauth2.oauth2provider",
	"authentik_providers_proxy.proxyprovider",
	"authentik_providers_radius.radiusprovider",
	"authentik_providers_radius.radiusproviderpropertymapping",
	"authentik_providers_saml.samlprovider",
	"authentik_providers_saml.samlpropertymapping",
	"authentik_providers_scim.scimprovider",
	"authentik_providers_scim.scimmapping",
	"authentik_rbac.role",
	"authentik_sources_ldap.ldapsource",
	"authentik_sources_ldap.ldapsourcepropertymapping",
	"authentik_sources_oauth.oauthsource",
	"authentik_sources_oauth.useroauthsourceconnection",
	"authentik_sources_plex.plexsource",
	"authentik_sources_plex.plexsourceconnection",
	"authentik_sources_saml.samlsource",
	"authentik_sources_saml.usersamlsourceconnection",
	"authentik_sources_scim.scimsource",
	"authentik_sources_scim.scimsourcepropertymapping",
	"authentik_stages_authenticator_duo.authenticatorduostage",
	"authentik_stages_authenticator_duo.duodevice",
	"authentik_stages_authenticator_sms.authenticatorsmsstage",
	"authentik_stages_authenticator_sms.smsdevice",
	"authentik_stages_authenticator_static.authenticatorstaticstage",
	"authentik_stages_authenticator_static.staticdevice",
	"authentik_stages_authenticator_totp.authenticatortotpstage",
	"authentik_stages_authenticator_totp.totpdevice",
	"authentik_stages_authenticator_validate.authenticatorvalidatestage",
	"authentik_stages_authenticator_webauthn.authenticatorwebauthnstage",
	"authentik_stages_authenticator_webauthn.webauthndevice",
	"authentik_stages_captcha.captchastage",
	"authentik_stages_consent.consentstage",
	"authentik_stages_consent.userconsent",
	"authentik_stages_deny.denystage",
	"authentik_stages_dummy.dummystage",
	"authentik_stages_email.emailstage",
	"authentik_stages_identification.identificationstage",
	"authentik_stages_invitation.invitationstage",
	"authentik_stages_invitation.invitation",
	"authentik_stages_password.passwordstage",
	"authentik_stages_prompt.prompt",
	"authentik_stages_prompt.promptstage",
	"authentik_stages_user_delete.userdeletestage",
	"authentik_stages_user_login.userloginstage",
	"authentik_stages_user_logout.userlogoutstage",
	"authentik_stages_user_write.userwritestage",
	"authentik_brands.brand",
	"authentik_blueprints.blueprintinstance",
	"authentik_core.group",
	"authentik_core.user",
	"authentik_core.application",
	"authentik_core.token",
	"authentik_enterprise.license",
	"authentik_providers_google_workspace.googleworkspaceprovider",
	"authentik_providers_google_workspace.googleworkspaceprovidermapping",
	"authentik_providers_microsoft_entra.microsoftentraprovider",
	"authentik_providers_microsoft_entra.microsoftentraprovidermapping",
	"authentik_providers_rac.racprovider",
	"authentik_providers_rac.endpoint",
	"authentik_providers_rac.racpropertymapping",
	"authentik_stages_source.sourcestage",
	"authentik_events.event",
	"authentik_events.notificationtransport",
	"authentik_events.notification",
	"authentik_events.notificationrule",
	"authentik_events.notificationwebhookmapping",
}

func (v *ModelEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelEnum(value)
	for _, existing := range AllowedModelEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelEnum", value)
}

// NewModelEnumFromValue returns a pointer to a valid ModelEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelEnumFromValue(v string) (*ModelEnum, error) {
	ev := ModelEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelEnum: valid values are %v", v, AllowedModelEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelEnum) IsValid() bool {
	for _, existing := range AllowedModelEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModelEnum value
func (v ModelEnum) Ptr() *ModelEnum {
	return &v
}

type NullableModelEnum struct {
	value *ModelEnum
	isSet bool
}

func (v NullableModelEnum) Get() *ModelEnum {
	return v.value
}

func (v *NullableModelEnum) Set(val *ModelEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableModelEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableModelEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelEnum(val *ModelEnum) *NullableModelEnum {
	return &NullableModelEnum{value: val, isSet: true}
}

func (v NullableModelEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
