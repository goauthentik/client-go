/*
authentik

Making authentication simple.

API version: 2023.3.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedEventMatcherPolicyRequest Event Matcher Policy Serializer
type PatchedEventMatcherPolicyRequest struct {
	Name *string `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	// Match created events with this action type. When left empty, all action types will be matched.  * `login` - Login * `login_failed` - Login Failed * `logout` - Logout * `user_write` - User Write * `suspicious_request` - Suspicious Request * `password_set` - Password Set * `secret_view` - Secret View * `secret_rotate` - Secret Rotate * `invitation_used` - Invite Used * `authorize_application` - Authorize Application * `source_linked` - Source Linked * `impersonation_started` - Impersonation Started * `impersonation_ended` - Impersonation Ended * `flow_execution` - Flow Execution * `policy_execution` - Policy Execution * `policy_exception` - Policy Exception * `property_mapping_exception` - Property Mapping Exception * `system_task_execution` - System Task Execution * `system_task_exception` - System Task Exception * `system_exception` - System Exception * `configuration_error` - Configuration Error * `model_created` - Model Created * `model_updated` - Model Updated * `model_deleted` - Model Deleted * `email_sent` - Email Sent * `update_available` - Update Available * `custom_` - Custom Prefix
	Action NullableEventActions `json:"action,omitempty"`
	// Matches Event's Client IP (strict matching, for network matching use an Expression Policy)
	ClientIp *string `json:"client_ip,omitempty"`
	// Match events created by selected application. When left empty, all applications are matched.  * `authentik.admin` - authentik Admin * `authentik.api` - authentik API * `authentik.crypto` - authentik Crypto * `authentik.events` - authentik Events * `authentik.flows` - authentik Flows * `authentik.lib` - authentik lib * `authentik.outposts` - authentik Outpost * `authentik.policies.dummy` - authentik Policies.Dummy * `authentik.policies.event_matcher` - authentik Policies.Event Matcher * `authentik.policies.expiry` - authentik Policies.Expiry * `authentik.policies.expression` - authentik Policies.Expression * `authentik.policies.password` - authentik Policies.Password * `authentik.policies.reputation` - authentik Policies.Reputation * `authentik.policies` - authentik Policies * `authentik.providers.ldap` - authentik Providers.LDAP * `authentik.providers.oauth2` - authentik Providers.OAuth2 * `authentik.providers.proxy` - authentik Providers.Proxy * `authentik.providers.saml` - authentik Providers.SAML * `authentik.providers.scim` - authentik Providers.SCIM * `authentik.recovery` - authentik Recovery * `authentik.sources.ldap` - authentik Sources.LDAP * `authentik.sources.oauth` - authentik Sources.OAuth * `authentik.sources.plex` - authentik Sources.Plex * `authentik.sources.saml` - authentik Sources.SAML * `authentik.stages.authenticator_duo` - authentik Stages.Authenticator.Duo * `authentik.stages.authenticator_sms` - authentik Stages.Authenticator.SMS * `authentik.stages.authenticator_static` - authentik Stages.Authenticator.Static * `authentik.stages.authenticator_totp` - authentik Stages.Authenticator.TOTP * `authentik.stages.authenticator_validate` - authentik Stages.Authenticator.Validate * `authentik.stages.authenticator_webauthn` - authentik Stages.Authenticator.WebAuthn * `authentik.stages.captcha` - authentik Stages.Captcha * `authentik.stages.consent` - authentik Stages.Consent * `authentik.stages.deny` - authentik Stages.Deny * `authentik.stages.dummy` - authentik Stages.Dummy * `authentik.stages.email` - authentik Stages.Email * `authentik.stages.identification` - authentik Stages.Identification * `authentik.stages.invitation` - authentik Stages.User Invitation * `authentik.stages.password` - authentik Stages.Password * `authentik.stages.prompt` - authentik Stages.Prompt * `authentik.stages.user_delete` - authentik Stages.User Delete * `authentik.stages.user_login` - authentik Stages.User Login * `authentik.stages.user_logout` - authentik Stages.User Logout * `authentik.stages.user_write` - authentik Stages.User Write * `authentik.tenants` - authentik Tenants * `authentik.blueprints` - authentik Blueprints * `authentik.core` - authentik Core
	App NullableAppEnum `json:"app,omitempty"`
}

// NewPatchedEventMatcherPolicyRequest instantiates a new PatchedEventMatcherPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedEventMatcherPolicyRequest() *PatchedEventMatcherPolicyRequest {
	this := PatchedEventMatcherPolicyRequest{}
	return &this
}

// NewPatchedEventMatcherPolicyRequestWithDefaults instantiates a new PatchedEventMatcherPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedEventMatcherPolicyRequestWithDefaults() *PatchedEventMatcherPolicyRequest {
	this := PatchedEventMatcherPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedEventMatcherPolicyRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEventMatcherPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedEventMatcherPolicyRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedEventMatcherPolicyRequest) SetName(v string) {
	o.Name = &v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *PatchedEventMatcherPolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEventMatcherPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *PatchedEventMatcherPolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *PatchedEventMatcherPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEventMatcherPolicyRequest) GetAction() EventActions {
	if o == nil || o.Action.Get() == nil {
		var ret EventActions
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEventMatcherPolicyRequest) GetActionOk() (*EventActions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *PatchedEventMatcherPolicyRequest) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableEventActions and assigns it to the Action field.
func (o *PatchedEventMatcherPolicyRequest) SetAction(v EventActions) {
	o.Action.Set(&v)
}

// SetActionNil sets the value for Action to be an explicit nil
func (o *PatchedEventMatcherPolicyRequest) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *PatchedEventMatcherPolicyRequest) UnsetAction() {
	o.Action.Unset()
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *PatchedEventMatcherPolicyRequest) GetClientIp() string {
	if o == nil || o.ClientIp == nil {
		var ret string
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEventMatcherPolicyRequest) GetClientIpOk() (*string, bool) {
	if o == nil || o.ClientIp == nil {
		return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *PatchedEventMatcherPolicyRequest) HasClientIp() bool {
	if o != nil && o.ClientIp != nil {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given string and assigns it to the ClientIp field.
func (o *PatchedEventMatcherPolicyRequest) SetClientIp(v string) {
	o.ClientIp = &v
}

// GetApp returns the App field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEventMatcherPolicyRequest) GetApp() AppEnum {
	if o == nil || o.App.Get() == nil {
		var ret AppEnum
		return ret
	}
	return *o.App.Get()
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEventMatcherPolicyRequest) GetAppOk() (*AppEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.App.Get(), o.App.IsSet()
}

// HasApp returns a boolean if a field has been set.
func (o *PatchedEventMatcherPolicyRequest) HasApp() bool {
	if o != nil && o.App.IsSet() {
		return true
	}

	return false
}

// SetApp gets a reference to the given NullableAppEnum and assigns it to the App field.
func (o *PatchedEventMatcherPolicyRequest) SetApp(v AppEnum) {
	o.App.Set(&v)
}

// SetAppNil sets the value for App to be an explicit nil
func (o *PatchedEventMatcherPolicyRequest) SetAppNil() {
	o.App.Set(nil)
}

// UnsetApp ensures that no value is present for App, not even an explicit nil
func (o *PatchedEventMatcherPolicyRequest) UnsetApp() {
	o.App.Unset()
}

func (o PatchedEventMatcherPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.ClientIp != nil {
		toSerialize["client_ip"] = o.ClientIp
	}
	if o.App.IsSet() {
		toSerialize["app"] = o.App.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedEventMatcherPolicyRequest struct {
	value *PatchedEventMatcherPolicyRequest
	isSet bool
}

func (v NullablePatchedEventMatcherPolicyRequest) Get() *PatchedEventMatcherPolicyRequest {
	return v.value
}

func (v *NullablePatchedEventMatcherPolicyRequest) Set(val *PatchedEventMatcherPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedEventMatcherPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedEventMatcherPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedEventMatcherPolicyRequest(val *PatchedEventMatcherPolicyRequest) *NullablePatchedEventMatcherPolicyRequest {
	return &NullablePatchedEventMatcherPolicyRequest{value: val, isSet: true}
}

func (v NullablePatchedEventMatcherPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedEventMatcherPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
