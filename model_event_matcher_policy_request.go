/*
authentik

Making authentication simple.

API version: 2023.10.7
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// EventMatcherPolicyRequest Event Matcher Policy Serializer
type EventMatcherPolicyRequest struct {
	Name string `json:"name"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	// Match created events with this action type. When left empty, all action types will be matched.  * `login` - Login * `login_failed` - Login Failed * `logout` - Logout * `user_write` - User Write * `suspicious_request` - Suspicious Request * `password_set` - Password Set * `secret_view` - Secret View * `secret_rotate` - Secret Rotate * `invitation_used` - Invite Used * `authorize_application` - Authorize Application * `source_linked` - Source Linked * `impersonation_started` - Impersonation Started * `impersonation_ended` - Impersonation Ended * `flow_execution` - Flow Execution * `policy_execution` - Policy Execution * `policy_exception` - Policy Exception * `property_mapping_exception` - Property Mapping Exception * `system_task_execution` - System Task Execution * `system_task_exception` - System Task Exception * `system_exception` - System Exception * `configuration_error` - Configuration Error * `model_created` - Model Created * `model_updated` - Model Updated * `model_deleted` - Model Deleted * `email_sent` - Email Sent * `update_available` - Update Available * `custom_` - Custom Prefix
	Action NullableEventActions `json:"action,omitempty"`
	// Matches Event's Client IP (strict matching, for network matching use an Expression Policy)
	ClientIp NullableString `json:"client_ip,omitempty"`
	// Match events created by selected application. When left empty, all applications are matched.  * `authentik.tenants` - authentik Tenants * `authentik.admin` - authentik Admin * `authentik.api` - authentik API * `authentik.crypto` - authentik Crypto * `authentik.flows` - authentik Flows * `authentik.outposts` - authentik Outpost * `authentik.policies.dummy` - authentik Policies.Dummy * `authentik.policies.event_matcher` - authentik Policies.Event Matcher * `authentik.policies.expiry` - authentik Policies.Expiry * `authentik.policies.expression` - authentik Policies.Expression * `authentik.policies.password` - authentik Policies.Password * `authentik.policies.reputation` - authentik Policies.Reputation * `authentik.policies` - authentik Policies * `authentik.providers.ldap` - authentik Providers.LDAP * `authentik.providers.oauth2` - authentik Providers.OAuth2 * `authentik.providers.proxy` - authentik Providers.Proxy * `authentik.providers.radius` - authentik Providers.Radius * `authentik.providers.saml` - authentik Providers.SAML * `authentik.providers.scim` - authentik Providers.SCIM * `authentik.rbac` - authentik RBAC * `authentik.recovery` - authentik Recovery * `authentik.sources.ldap` - authentik Sources.LDAP * `authentik.sources.oauth` - authentik Sources.OAuth * `authentik.sources.plex` - authentik Sources.Plex * `authentik.sources.saml` - authentik Sources.SAML * `authentik.stages.authenticator` - authentik Stages.Authenticator * `authentik.stages.authenticator_duo` - authentik Stages.Authenticator.Duo * `authentik.stages.authenticator_sms` - authentik Stages.Authenticator.SMS * `authentik.stages.authenticator_static` - authentik Stages.Authenticator.Static * `authentik.stages.authenticator_totp` - authentik Stages.Authenticator.TOTP * `authentik.stages.authenticator_validate` - authentik Stages.Authenticator.Validate * `authentik.stages.authenticator_webauthn` - authentik Stages.Authenticator.WebAuthn * `authentik.stages.captcha` - authentik Stages.Captcha * `authentik.stages.consent` - authentik Stages.Consent * `authentik.stages.deny` - authentik Stages.Deny * `authentik.stages.dummy` - authentik Stages.Dummy * `authentik.stages.email` - authentik Stages.Email * `authentik.stages.identification` - authentik Stages.Identification * `authentik.stages.invitation` - authentik Stages.User Invitation * `authentik.stages.password` - authentik Stages.Password * `authentik.stages.prompt` - authentik Stages.Prompt * `authentik.stages.user_delete` - authentik Stages.User Delete * `authentik.stages.user_login` - authentik Stages.User Login * `authentik.stages.user_logout` - authentik Stages.User Logout * `authentik.stages.user_write` - authentik Stages.User Write * `authentik.brands` - authentik Brands * `authentik.blueprints` - authentik Blueprints * `authentik.core` - authentik Core * `authentik.enterprise` - authentik Enterprise * `authentik.enterprise.audit` - authentik Enterprise.Audit * `authentik.enterprise.providers.rac` - authentik Enterprise.Providers.RAC * `authentik.events` - authentik Events
	App NullableAppEnum `json:"app,omitempty"`
	// Match events created by selected model. When left empty, all models are matched. When an app is selected, all the application's models are matched.  * `authentik_tenants.domain` - Domain * `authentik_crypto.certificatekeypair` - Certificate-Key Pair * `authentik_flows.flow` - Flow * `authentik_flows.flowstagebinding` - Flow Stage Binding * `authentik_outposts.dockerserviceconnection` - Docker Service-Connection * `authentik_outposts.kubernetesserviceconnection` - Kubernetes Service-Connection * `authentik_outposts.outpost` - Outpost * `authentik_policies_dummy.dummypolicy` - Dummy Policy * `authentik_policies_event_matcher.eventmatcherpolicy` - Event Matcher Policy * `authentik_policies_expiry.passwordexpirypolicy` - Password Expiry Policy * `authentik_policies_expression.expressionpolicy` - Expression Policy * `authentik_policies_password.passwordpolicy` - Password Policy * `authentik_policies_reputation.reputationpolicy` - Reputation Policy * `authentik_policies.policybinding` - Policy Binding * `authentik_providers_ldap.ldapprovider` - LDAP Provider * `authentik_providers_oauth2.scopemapping` - Scope Mapping * `authentik_providers_oauth2.oauth2provider` - OAuth2/OpenID Provider * `authentik_providers_proxy.proxyprovider` - Proxy Provider * `authentik_providers_radius.radiusprovider` - Radius Provider * `authentik_providers_saml.samlprovider` - SAML Provider * `authentik_providers_saml.samlpropertymapping` - SAML Property Mapping * `authentik_providers_scim.scimprovider` - SCIM Provider * `authentik_providers_scim.scimmapping` - SCIM Mapping * `authentik_rbac.role` - Role * `authentik_sources_ldap.ldapsource` - LDAP Source * `authentik_sources_ldap.ldappropertymapping` - LDAP Property Mapping * `authentik_sources_oauth.oauthsource` - OAuth Source * `authentik_sources_oauth.useroauthsourceconnection` - User OAuth Source Connection * `authentik_sources_plex.plexsource` - Plex Source * `authentik_sources_plex.plexsourceconnection` - User Plex Source Connection * `authentik_sources_saml.samlsource` - SAML Source * `authentik_sources_saml.usersamlsourceconnection` - User SAML Source Connection * `authentik_stages_authenticator_duo.authenticatorduostage` - Duo Authenticator Setup Stage * `authentik_stages_authenticator_duo.duodevice` - Duo Device * `authentik_stages_authenticator_sms.authenticatorsmsstage` - SMS Authenticator Setup Stage * `authentik_stages_authenticator_sms.smsdevice` - SMS Device * `authentik_stages_authenticator_static.authenticatorstaticstage` - Static Authenticator Setup Stage * `authentik_stages_authenticator_static.staticdevice` - Static Device * `authentik_stages_authenticator_totp.authenticatortotpstage` - TOTP Authenticator Setup Stage * `authentik_stages_authenticator_totp.totpdevice` - TOTP Device * `authentik_stages_authenticator_validate.authenticatorvalidatestage` - Authenticator Validation Stage * `authentik_stages_authenticator_webauthn.authenticatewebauthnstage` - WebAuthn Authenticator Setup Stage * `authentik_stages_authenticator_webauthn.webauthndevice` - WebAuthn Device * `authentik_stages_captcha.captchastage` - Captcha Stage * `authentik_stages_consent.consentstage` - Consent Stage * `authentik_stages_consent.userconsent` - User Consent * `authentik_stages_deny.denystage` - Deny Stage * `authentik_stages_dummy.dummystage` - Dummy Stage * `authentik_stages_email.emailstage` - Email Stage * `authentik_stages_identification.identificationstage` - Identification Stage * `authentik_stages_invitation.invitationstage` - Invitation Stage * `authentik_stages_invitation.invitation` - Invitation * `authentik_stages_password.passwordstage` - Password Stage * `authentik_stages_prompt.prompt` - Prompt * `authentik_stages_prompt.promptstage` - Prompt Stage * `authentik_stages_user_delete.userdeletestage` - User Delete Stage * `authentik_stages_user_login.userloginstage` - User Login Stage * `authentik_stages_user_logout.userlogoutstage` - User Logout Stage * `authentik_stages_user_write.userwritestage` - User Write Stage * `authentik_brands.brand` - Brand * `authentik_blueprints.blueprintinstance` - Blueprint Instance * `authentik_core.group` - Group * `authentik_core.user` - User * `authentik_core.application` - Application * `authentik_core.token` - Token * `authentik_enterprise.license` - License * `authentik_providers_rac.racprovider` - RAC Provider * `authentik_providers_rac.endpoint` - RAC Endpoint * `authentik_providers_rac.racpropertymapping` - RAC Property Mapping * `authentik_events.event` - Event * `authentik_events.notificationtransport` - Notification Transport * `authentik_events.notification` - Notification * `authentik_events.notificationrule` - Notification Rule * `authentik_events.notificationwebhookmapping` - Webhook Mapping
	Model NullableModelEnum `json:"model,omitempty"`
}

// NewEventMatcherPolicyRequest instantiates a new EventMatcherPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventMatcherPolicyRequest(name string) *EventMatcherPolicyRequest {
	this := EventMatcherPolicyRequest{}
	this.Name = name
	return &this
}

// NewEventMatcherPolicyRequestWithDefaults instantiates a new EventMatcherPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventMatcherPolicyRequestWithDefaults() *EventMatcherPolicyRequest {
	this := EventMatcherPolicyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *EventMatcherPolicyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventMatcherPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventMatcherPolicyRequest) SetName(v string) {
	o.Name = v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *EventMatcherPolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMatcherPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *EventMatcherPolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *EventMatcherPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventMatcherPolicyRequest) GetAction() EventActions {
	if o == nil || o.Action.Get() == nil {
		var ret EventActions
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventMatcherPolicyRequest) GetActionOk() (*EventActions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *EventMatcherPolicyRequest) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableEventActions and assigns it to the Action field.
func (o *EventMatcherPolicyRequest) SetAction(v EventActions) {
	o.Action.Set(&v)
}

// SetActionNil sets the value for Action to be an explicit nil
func (o *EventMatcherPolicyRequest) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *EventMatcherPolicyRequest) UnsetAction() {
	o.Action.Unset()
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventMatcherPolicyRequest) GetClientIp() string {
	if o == nil || o.ClientIp.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientIp.Get()
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventMatcherPolicyRequest) GetClientIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientIp.Get(), o.ClientIp.IsSet()
}

// HasClientIp returns a boolean if a field has been set.
func (o *EventMatcherPolicyRequest) HasClientIp() bool {
	if o != nil && o.ClientIp.IsSet() {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given NullableString and assigns it to the ClientIp field.
func (o *EventMatcherPolicyRequest) SetClientIp(v string) {
	o.ClientIp.Set(&v)
}

// SetClientIpNil sets the value for ClientIp to be an explicit nil
func (o *EventMatcherPolicyRequest) SetClientIpNil() {
	o.ClientIp.Set(nil)
}

// UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
func (o *EventMatcherPolicyRequest) UnsetClientIp() {
	o.ClientIp.Unset()
}

// GetApp returns the App field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventMatcherPolicyRequest) GetApp() AppEnum {
	if o == nil || o.App.Get() == nil {
		var ret AppEnum
		return ret
	}
	return *o.App.Get()
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventMatcherPolicyRequest) GetAppOk() (*AppEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.App.Get(), o.App.IsSet()
}

// HasApp returns a boolean if a field has been set.
func (o *EventMatcherPolicyRequest) HasApp() bool {
	if o != nil && o.App.IsSet() {
		return true
	}

	return false
}

// SetApp gets a reference to the given NullableAppEnum and assigns it to the App field.
func (o *EventMatcherPolicyRequest) SetApp(v AppEnum) {
	o.App.Set(&v)
}

// SetAppNil sets the value for App to be an explicit nil
func (o *EventMatcherPolicyRequest) SetAppNil() {
	o.App.Set(nil)
}

// UnsetApp ensures that no value is present for App, not even an explicit nil
func (o *EventMatcherPolicyRequest) UnsetApp() {
	o.App.Unset()
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventMatcherPolicyRequest) GetModel() ModelEnum {
	if o == nil || o.Model.Get() == nil {
		var ret ModelEnum
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventMatcherPolicyRequest) GetModelOk() (*ModelEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *EventMatcherPolicyRequest) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableModelEnum and assigns it to the Model field.
func (o *EventMatcherPolicyRequest) SetModel(v ModelEnum) {
	o.Model.Set(&v)
}

// SetModelNil sets the value for Model to be an explicit nil
func (o *EventMatcherPolicyRequest) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *EventMatcherPolicyRequest) UnsetModel() {
	o.Model.Unset()
}

func (o EventMatcherPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.ClientIp.IsSet() {
		toSerialize["client_ip"] = o.ClientIp.Get()
	}
	if o.App.IsSet() {
		toSerialize["app"] = o.App.Get()
	}
	if o.Model.IsSet() {
		toSerialize["model"] = o.Model.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEventMatcherPolicyRequest struct {
	value *EventMatcherPolicyRequest
	isSet bool
}

func (v NullableEventMatcherPolicyRequest) Get() *EventMatcherPolicyRequest {
	return v.value
}

func (v *NullableEventMatcherPolicyRequest) Set(val *EventMatcherPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEventMatcherPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEventMatcherPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventMatcherPolicyRequest(val *EventMatcherPolicyRequest) *NullableEventMatcherPolicyRequest {
	return &NullableEventMatcherPolicyRequest{value: val, isSet: true}
}

func (v NullableEventMatcherPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventMatcherPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
