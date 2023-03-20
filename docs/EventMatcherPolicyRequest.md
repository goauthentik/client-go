# EventMatcherPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**Action** | Pointer to [**NullableEventActions**](EventActions.md) | Match created events with this action type. When left empty, all action types will be matched.  * &#x60;login&#x60; - Login * &#x60;login_failed&#x60; - Login Failed * &#x60;logout&#x60; - Logout * &#x60;user_write&#x60; - User Write * &#x60;suspicious_request&#x60; - Suspicious Request * &#x60;password_set&#x60; - Password Set * &#x60;secret_view&#x60; - Secret View * &#x60;secret_rotate&#x60; - Secret Rotate * &#x60;invitation_used&#x60; - Invite Used * &#x60;authorize_application&#x60; - Authorize Application * &#x60;source_linked&#x60; - Source Linked * &#x60;impersonation_started&#x60; - Impersonation Started * &#x60;impersonation_ended&#x60; - Impersonation Ended * &#x60;flow_execution&#x60; - Flow Execution * &#x60;policy_execution&#x60; - Policy Execution * &#x60;policy_exception&#x60; - Policy Exception * &#x60;property_mapping_exception&#x60; - Property Mapping Exception * &#x60;system_task_execution&#x60; - System Task Execution * &#x60;system_task_exception&#x60; - System Task Exception * &#x60;system_exception&#x60; - System Exception * &#x60;configuration_error&#x60; - Configuration Error * &#x60;model_created&#x60; - Model Created * &#x60;model_updated&#x60; - Model Updated * &#x60;model_deleted&#x60; - Model Deleted * &#x60;email_sent&#x60; - Email Sent * &#x60;update_available&#x60; - Update Available * &#x60;custom_&#x60; - Custom Prefix | [optional] 
**ClientIp** | Pointer to **string** | Matches Event&#39;s Client IP (strict matching, for network matching use an Expression Policy) | [optional] 
**App** | Pointer to [**NullableAppEnum**](AppEnum.md) | Match events created by selected application. When left empty, all applications are matched.  * &#x60;authentik.admin&#x60; - authentik Admin * &#x60;authentik.api&#x60; - authentik API * &#x60;authentik.crypto&#x60; - authentik Crypto * &#x60;authentik.events&#x60; - authentik Events * &#x60;authentik.flows&#x60; - authentik Flows * &#x60;authentik.lib&#x60; - authentik lib * &#x60;authentik.outposts&#x60; - authentik Outpost * &#x60;authentik.policies.dummy&#x60; - authentik Policies.Dummy * &#x60;authentik.policies.event_matcher&#x60; - authentik Policies.Event Matcher * &#x60;authentik.policies.expiry&#x60; - authentik Policies.Expiry * &#x60;authentik.policies.expression&#x60; - authentik Policies.Expression * &#x60;authentik.policies.password&#x60; - authentik Policies.Password * &#x60;authentik.policies.reputation&#x60; - authentik Policies.Reputation * &#x60;authentik.policies&#x60; - authentik Policies * &#x60;authentik.providers.ldap&#x60; - authentik Providers.LDAP * &#x60;authentik.providers.oauth2&#x60; - authentik Providers.OAuth2 * &#x60;authentik.providers.proxy&#x60; - authentik Providers.Proxy * &#x60;authentik.providers.radius&#x60; - authentik Providers.Radius * &#x60;authentik.providers.saml&#x60; - authentik Providers.SAML * &#x60;authentik.providers.scim&#x60; - authentik Providers.SCIM * &#x60;authentik.recovery&#x60; - authentik Recovery * &#x60;authentik.sources.ldap&#x60; - authentik Sources.LDAP * &#x60;authentik.sources.oauth&#x60; - authentik Sources.OAuth * &#x60;authentik.sources.plex&#x60; - authentik Sources.Plex * &#x60;authentik.sources.saml&#x60; - authentik Sources.SAML * &#x60;authentik.stages.authenticator_duo&#x60; - authentik Stages.Authenticator.Duo * &#x60;authentik.stages.authenticator_sms&#x60; - authentik Stages.Authenticator.SMS * &#x60;authentik.stages.authenticator_static&#x60; - authentik Stages.Authenticator.Static * &#x60;authentik.stages.authenticator_totp&#x60; - authentik Stages.Authenticator.TOTP * &#x60;authentik.stages.authenticator_validate&#x60; - authentik Stages.Authenticator.Validate * &#x60;authentik.stages.authenticator_webauthn&#x60; - authentik Stages.Authenticator.WebAuthn * &#x60;authentik.stages.captcha&#x60; - authentik Stages.Captcha * &#x60;authentik.stages.consent&#x60; - authentik Stages.Consent * &#x60;authentik.stages.deny&#x60; - authentik Stages.Deny * &#x60;authentik.stages.dummy&#x60; - authentik Stages.Dummy * &#x60;authentik.stages.email&#x60; - authentik Stages.Email * &#x60;authentik.stages.identification&#x60; - authentik Stages.Identification * &#x60;authentik.stages.invitation&#x60; - authentik Stages.User Invitation * &#x60;authentik.stages.password&#x60; - authentik Stages.Password * &#x60;authentik.stages.prompt&#x60; - authentik Stages.Prompt * &#x60;authentik.stages.user_delete&#x60; - authentik Stages.User Delete * &#x60;authentik.stages.user_login&#x60; - authentik Stages.User Login * &#x60;authentik.stages.user_logout&#x60; - authentik Stages.User Logout * &#x60;authentik.stages.user_write&#x60; - authentik Stages.User Write * &#x60;authentik.tenants&#x60; - authentik Tenants * &#x60;authentik.blueprints&#x60; - authentik Blueprints * &#x60;authentik.core&#x60; - authentik Core | [optional] 

## Methods

### NewEventMatcherPolicyRequest

`func NewEventMatcherPolicyRequest(name string, ) *EventMatcherPolicyRequest`

NewEventMatcherPolicyRequest instantiates a new EventMatcherPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventMatcherPolicyRequestWithDefaults

`func NewEventMatcherPolicyRequestWithDefaults() *EventMatcherPolicyRequest`

NewEventMatcherPolicyRequestWithDefaults instantiates a new EventMatcherPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventMatcherPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventMatcherPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventMatcherPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExecutionLogging

`func (o *EventMatcherPolicyRequest) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *EventMatcherPolicyRequest) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *EventMatcherPolicyRequest) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *EventMatcherPolicyRequest) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetAction

`func (o *EventMatcherPolicyRequest) GetAction() EventActions`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *EventMatcherPolicyRequest) GetActionOk() (*EventActions, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *EventMatcherPolicyRequest) SetAction(v EventActions)`

SetAction sets Action field to given value.

### HasAction

`func (o *EventMatcherPolicyRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *EventMatcherPolicyRequest) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *EventMatcherPolicyRequest) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetClientIp

`func (o *EventMatcherPolicyRequest) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *EventMatcherPolicyRequest) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *EventMatcherPolicyRequest) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *EventMatcherPolicyRequest) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetApp

`func (o *EventMatcherPolicyRequest) GetApp() AppEnum`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *EventMatcherPolicyRequest) GetAppOk() (*AppEnum, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *EventMatcherPolicyRequest) SetApp(v AppEnum)`

SetApp sets App field to given value.

### HasApp

`func (o *EventMatcherPolicyRequest) HasApp() bool`

HasApp returns a boolean if a field has been set.

### SetAppNil

`func (o *EventMatcherPolicyRequest) SetAppNil(b bool)`

 SetAppNil sets the value for App to be an explicit nil

### UnsetApp
`func (o *EventMatcherPolicyRequest) UnsetApp()`

UnsetApp ensures that no value is present for App, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


