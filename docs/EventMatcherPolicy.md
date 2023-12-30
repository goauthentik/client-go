# EventMatcherPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**Component** | **string** | Get object component so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**BoundTo** | **int32** | Return objects policy is bound to | [readonly] 
**Action** | Pointer to [**NullableEventActions**](EventActions.md) | Match created events with this action type. When left empty, all action types will be matched.  * &#x60;login&#x60; - Login * &#x60;login_failed&#x60; - Login Failed * &#x60;logout&#x60; - Logout * &#x60;user_write&#x60; - User Write * &#x60;suspicious_request&#x60; - Suspicious Request * &#x60;password_set&#x60; - Password Set * &#x60;secret_view&#x60; - Secret View * &#x60;secret_rotate&#x60; - Secret Rotate * &#x60;invitation_used&#x60; - Invite Used * &#x60;authorize_application&#x60; - Authorize Application * &#x60;source_linked&#x60; - Source Linked * &#x60;impersonation_started&#x60; - Impersonation Started * &#x60;impersonation_ended&#x60; - Impersonation Ended * &#x60;flow_execution&#x60; - Flow Execution * &#x60;policy_execution&#x60; - Policy Execution * &#x60;policy_exception&#x60; - Policy Exception * &#x60;property_mapping_exception&#x60; - Property Mapping Exception * &#x60;system_task_execution&#x60; - System Task Execution * &#x60;system_task_exception&#x60; - System Task Exception * &#x60;system_exception&#x60; - System Exception * &#x60;configuration_error&#x60; - Configuration Error * &#x60;model_created&#x60; - Model Created * &#x60;model_updated&#x60; - Model Updated * &#x60;model_deleted&#x60; - Model Deleted * &#x60;email_sent&#x60; - Email Sent * &#x60;update_available&#x60; - Update Available * &#x60;custom_&#x60; - Custom Prefix | [optional] 
**ClientIp** | Pointer to **NullableString** | Matches Event&#39;s Client IP (strict matching, for network matching use an Expression Policy) | [optional] 
**App** | Pointer to [**NullableAppEnum**](AppEnum.md) | Match events created by selected application. When left empty, all applications are matched.  * &#x60;authentik.admin&#x60; - authentik Admin * &#x60;authentik.api&#x60; - authentik API * &#x60;authentik.crypto&#x60; - authentik Crypto * &#x60;authentik.events&#x60; - authentik Events * &#x60;authentik.flows&#x60; - authentik Flows * &#x60;authentik.outposts&#x60; - authentik Outpost * &#x60;authentik.policies.dummy&#x60; - authentik Policies.Dummy * &#x60;authentik.policies.event_matcher&#x60; - authentik Policies.Event Matcher * &#x60;authentik.policies.expiry&#x60; - authentik Policies.Expiry * &#x60;authentik.policies.expression&#x60; - authentik Policies.Expression * &#x60;authentik.policies.password&#x60; - authentik Policies.Password * &#x60;authentik.policies.reputation&#x60; - authentik Policies.Reputation * &#x60;authentik.policies&#x60; - authentik Policies * &#x60;authentik.providers.ldap&#x60; - authentik Providers.LDAP * &#x60;authentik.providers.oauth2&#x60; - authentik Providers.OAuth2 * &#x60;authentik.providers.proxy&#x60; - authentik Providers.Proxy * &#x60;authentik.providers.radius&#x60; - authentik Providers.Radius * &#x60;authentik.providers.saml&#x60; - authentik Providers.SAML * &#x60;authentik.providers.scim&#x60; - authentik Providers.SCIM * &#x60;authentik.rbac&#x60; - authentik RBAC * &#x60;authentik.recovery&#x60; - authentik Recovery * &#x60;authentik.sources.ldap&#x60; - authentik Sources.LDAP * &#x60;authentik.sources.oauth&#x60; - authentik Sources.OAuth * &#x60;authentik.sources.plex&#x60; - authentik Sources.Plex * &#x60;authentik.sources.saml&#x60; - authentik Sources.SAML * &#x60;authentik.stages.authenticator&#x60; - authentik Stages.Authenticator * &#x60;authentik.stages.authenticator_duo&#x60; - authentik Stages.Authenticator.Duo * &#x60;authentik.stages.authenticator_sms&#x60; - authentik Stages.Authenticator.SMS * &#x60;authentik.stages.authenticator_static&#x60; - authentik Stages.Authenticator.Static * &#x60;authentik.stages.authenticator_totp&#x60; - authentik Stages.Authenticator.TOTP * &#x60;authentik.stages.authenticator_validate&#x60; - authentik Stages.Authenticator.Validate * &#x60;authentik.stages.authenticator_webauthn&#x60; - authentik Stages.Authenticator.WebAuthn * &#x60;authentik.stages.captcha&#x60; - authentik Stages.Captcha * &#x60;authentik.stages.consent&#x60; - authentik Stages.Consent * &#x60;authentik.stages.deny&#x60; - authentik Stages.Deny * &#x60;authentik.stages.dummy&#x60; - authentik Stages.Dummy * &#x60;authentik.stages.email&#x60; - authentik Stages.Email * &#x60;authentik.stages.identification&#x60; - authentik Stages.Identification * &#x60;authentik.stages.invitation&#x60; - authentik Stages.User Invitation * &#x60;authentik.stages.password&#x60; - authentik Stages.Password * &#x60;authentik.stages.prompt&#x60; - authentik Stages.Prompt * &#x60;authentik.stages.user_delete&#x60; - authentik Stages.User Delete * &#x60;authentik.stages.user_login&#x60; - authentik Stages.User Login * &#x60;authentik.stages.user_logout&#x60; - authentik Stages.User Logout * &#x60;authentik.stages.user_write&#x60; - authentik Stages.User Write * &#x60;authentik.tenants&#x60; - authentik Tenants * &#x60;authentik.blueprints&#x60; - authentik Blueprints * &#x60;authentik.core&#x60; - authentik Core * &#x60;authentik.enterprise&#x60; - authentik Enterprise * &#x60;authentik.enterprise.providers.rac&#x60; - authentik Enterprise.Providers.RAC | [optional] 
**Model** | Pointer to [**NullableModelEnum**](ModelEnum.md) | Match events created by selected model. When left empty, all models are matched. When an app is selected, all the application&#39;s models are matched.  * &#x60;authentik_crypto.certificatekeypair&#x60; - Certificate-Key Pair * &#x60;authentik_events.event&#x60; - Event * &#x60;authentik_events.notificationtransport&#x60; - Notification Transport * &#x60;authentik_events.notification&#x60; - Notification * &#x60;authentik_events.notificationrule&#x60; - Notification Rule * &#x60;authentik_events.notificationwebhookmapping&#x60; - Webhook Mapping * &#x60;authentik_flows.flow&#x60; - Flow * &#x60;authentik_flows.flowstagebinding&#x60; - Flow Stage Binding * &#x60;authentik_outposts.dockerserviceconnection&#x60; - Docker Service-Connection * &#x60;authentik_outposts.kubernetesserviceconnection&#x60; - Kubernetes Service-Connection * &#x60;authentik_outposts.outpost&#x60; - Outpost * &#x60;authentik_policies_dummy.dummypolicy&#x60; - Dummy Policy * &#x60;authentik_policies_event_matcher.eventmatcherpolicy&#x60; - Event Matcher Policy * &#x60;authentik_policies_expiry.passwordexpirypolicy&#x60; - Password Expiry Policy * &#x60;authentik_policies_expression.expressionpolicy&#x60; - Expression Policy * &#x60;authentik_policies_password.passwordpolicy&#x60; - Password Policy * &#x60;authentik_policies_reputation.reputationpolicy&#x60; - Reputation Policy * &#x60;authentik_policies_reputation.reputation&#x60; - Reputation Score * &#x60;authentik_policies.policybinding&#x60; - Policy Binding * &#x60;authentik_providers_ldap.ldapprovider&#x60; - LDAP Provider * &#x60;authentik_providers_oauth2.scopemapping&#x60; - Scope Mapping * &#x60;authentik_providers_oauth2.oauth2provider&#x60; - OAuth2/OpenID Provider * &#x60;authentik_providers_oauth2.authorizationcode&#x60; - Authorization Code * &#x60;authentik_providers_oauth2.accesstoken&#x60; - OAuth2 Access Token * &#x60;authentik_providers_oauth2.refreshtoken&#x60; - OAuth2 Refresh Token * &#x60;authentik_providers_proxy.proxyprovider&#x60; - Proxy Provider * &#x60;authentik_providers_radius.radiusprovider&#x60; - Radius Provider * &#x60;authentik_providers_saml.samlprovider&#x60; - SAML Provider * &#x60;authentik_providers_saml.samlpropertymapping&#x60; - SAML Property Mapping * &#x60;authentik_providers_scim.scimprovider&#x60; - SCIM Provider * &#x60;authentik_providers_scim.scimmapping&#x60; - SCIM Mapping * &#x60;authentik_rbac.role&#x60; - Role * &#x60;authentik_sources_ldap.ldapsource&#x60; - LDAP Source * &#x60;authentik_sources_ldap.ldappropertymapping&#x60; - LDAP Property Mapping * &#x60;authentik_sources_oauth.oauthsource&#x60; - OAuth Source * &#x60;authentik_sources_oauth.useroauthsourceconnection&#x60; - User OAuth Source Connection * &#x60;authentik_sources_plex.plexsource&#x60; - Plex Source * &#x60;authentik_sources_plex.plexsourceconnection&#x60; - User Plex Source Connection * &#x60;authentik_sources_saml.samlsource&#x60; - SAML Source * &#x60;authentik_sources_saml.usersamlsourceconnection&#x60; - User SAML Source Connection * &#x60;authentik_stages_authenticator_duo.authenticatorduostage&#x60; - Duo Authenticator Setup Stage * &#x60;authentik_stages_authenticator_duo.duodevice&#x60; - Duo Device * &#x60;authentik_stages_authenticator_sms.authenticatorsmsstage&#x60; - SMS Authenticator Setup Stage * &#x60;authentik_stages_authenticator_sms.smsdevice&#x60; - SMS Device * &#x60;authentik_stages_authenticator_static.authenticatorstaticstage&#x60; - Static Authenticator Stage * &#x60;authentik_stages_authenticator_static.staticdevice&#x60; - Static Device * &#x60;authentik_stages_authenticator_totp.authenticatortotpstage&#x60; - TOTP Authenticator Setup Stage * &#x60;authentik_stages_authenticator_totp.totpdevice&#x60; - TOTP Device * &#x60;authentik_stages_authenticator_validate.authenticatorvalidatestage&#x60; - Authenticator Validation Stage * &#x60;authentik_stages_authenticator_webauthn.authenticatewebauthnstage&#x60; - WebAuthn Authenticator Setup Stage * &#x60;authentik_stages_authenticator_webauthn.webauthndevice&#x60; - WebAuthn Device * &#x60;authentik_stages_captcha.captchastage&#x60; - Captcha Stage * &#x60;authentik_stages_consent.consentstage&#x60; - Consent Stage * &#x60;authentik_stages_consent.userconsent&#x60; - User Consent * &#x60;authentik_stages_deny.denystage&#x60; - Deny Stage * &#x60;authentik_stages_dummy.dummystage&#x60; - Dummy Stage * &#x60;authentik_stages_email.emailstage&#x60; - Email Stage * &#x60;authentik_stages_identification.identificationstage&#x60; - Identification Stage * &#x60;authentik_stages_invitation.invitationstage&#x60; - Invitation Stage * &#x60;authentik_stages_invitation.invitation&#x60; - Invitation * &#x60;authentik_stages_password.passwordstage&#x60; - Password Stage * &#x60;authentik_stages_prompt.prompt&#x60; - Prompt * &#x60;authentik_stages_prompt.promptstage&#x60; - Prompt Stage * &#x60;authentik_stages_user_delete.userdeletestage&#x60; - User Delete Stage * &#x60;authentik_stages_user_login.userloginstage&#x60; - User Login Stage * &#x60;authentik_stages_user_logout.userlogoutstage&#x60; - User Logout Stage * &#x60;authentik_stages_user_write.userwritestage&#x60; - User Write Stage * &#x60;authentik_tenants.tenant&#x60; - Tenant * &#x60;authentik_blueprints.blueprintinstance&#x60; - Blueprint Instance * &#x60;authentik_core.group&#x60; - Group * &#x60;authentik_core.user&#x60; - User * &#x60;authentik_core.application&#x60; - Application * &#x60;authentik_core.token&#x60; - Token * &#x60;authentik_enterprise.license&#x60; - License * &#x60;authentik_providers_rac.racprovider&#x60; - RAC Provider * &#x60;authentik_providers_rac.endpoint&#x60; - RAC Endpoint * &#x60;authentik_providers_rac.racpropertymapping&#x60; - RAC Property Mapping | [optional] 

## Methods

### NewEventMatcherPolicy

`func NewEventMatcherPolicy(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, boundTo int32, ) *EventMatcherPolicy`

NewEventMatcherPolicy instantiates a new EventMatcherPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventMatcherPolicyWithDefaults

`func NewEventMatcherPolicyWithDefaults() *EventMatcherPolicy`

NewEventMatcherPolicyWithDefaults instantiates a new EventMatcherPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *EventMatcherPolicy) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *EventMatcherPolicy) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *EventMatcherPolicy) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *EventMatcherPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventMatcherPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventMatcherPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetExecutionLogging

`func (o *EventMatcherPolicy) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *EventMatcherPolicy) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *EventMatcherPolicy) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *EventMatcherPolicy) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetComponent

`func (o *EventMatcherPolicy) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *EventMatcherPolicy) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *EventMatcherPolicy) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *EventMatcherPolicy) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *EventMatcherPolicy) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *EventMatcherPolicy) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *EventMatcherPolicy) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *EventMatcherPolicy) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *EventMatcherPolicy) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *EventMatcherPolicy) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *EventMatcherPolicy) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *EventMatcherPolicy) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetBoundTo

`func (o *EventMatcherPolicy) GetBoundTo() int32`

GetBoundTo returns the BoundTo field if non-nil, zero value otherwise.

### GetBoundToOk

`func (o *EventMatcherPolicy) GetBoundToOk() (*int32, bool)`

GetBoundToOk returns a tuple with the BoundTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundTo

`func (o *EventMatcherPolicy) SetBoundTo(v int32)`

SetBoundTo sets BoundTo field to given value.


### GetAction

`func (o *EventMatcherPolicy) GetAction() EventActions`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *EventMatcherPolicy) GetActionOk() (*EventActions, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *EventMatcherPolicy) SetAction(v EventActions)`

SetAction sets Action field to given value.

### HasAction

`func (o *EventMatcherPolicy) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *EventMatcherPolicy) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *EventMatcherPolicy) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetClientIp

`func (o *EventMatcherPolicy) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *EventMatcherPolicy) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *EventMatcherPolicy) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *EventMatcherPolicy) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### SetClientIpNil

`func (o *EventMatcherPolicy) SetClientIpNil(b bool)`

 SetClientIpNil sets the value for ClientIp to be an explicit nil

### UnsetClientIp
`func (o *EventMatcherPolicy) UnsetClientIp()`

UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
### GetApp

`func (o *EventMatcherPolicy) GetApp() AppEnum`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *EventMatcherPolicy) GetAppOk() (*AppEnum, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *EventMatcherPolicy) SetApp(v AppEnum)`

SetApp sets App field to given value.

### HasApp

`func (o *EventMatcherPolicy) HasApp() bool`

HasApp returns a boolean if a field has been set.

### SetAppNil

`func (o *EventMatcherPolicy) SetAppNil(b bool)`

 SetAppNil sets the value for App to be an explicit nil

### UnsetApp
`func (o *EventMatcherPolicy) UnsetApp()`

UnsetApp ensures that no value is present for App, not even an explicit nil
### GetModel

`func (o *EventMatcherPolicy) GetModel() ModelEnum`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EventMatcherPolicy) GetModelOk() (*ModelEnum, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EventMatcherPolicy) SetModel(v ModelEnum)`

SetModel sets Model field to given value.

### HasModel

`func (o *EventMatcherPolicy) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *EventMatcherPolicy) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *EventMatcherPolicy) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


