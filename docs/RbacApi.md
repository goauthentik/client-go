# \RbacApi

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RbacPermissionsAssignedByRolesAssignCreate**](RbacApi.md#RbacPermissionsAssignedByRolesAssignCreate) | **Post** /rbac/permissions/assigned_by_roles/{uuid}/assign/ | 
[**RbacPermissionsAssignedByRolesList**](RbacApi.md#RbacPermissionsAssignedByRolesList) | **Get** /rbac/permissions/assigned_by_roles/ | 
[**RbacPermissionsAssignedByRolesUnassignPartialUpdate**](RbacApi.md#RbacPermissionsAssignedByRolesUnassignPartialUpdate) | **Patch** /rbac/permissions/assigned_by_roles/{uuid}/unassign/ | 
[**RbacPermissionsAssignedByUsersAssignCreate**](RbacApi.md#RbacPermissionsAssignedByUsersAssignCreate) | **Post** /rbac/permissions/assigned_by_users/{id}/assign/ | 
[**RbacPermissionsAssignedByUsersList**](RbacApi.md#RbacPermissionsAssignedByUsersList) | **Get** /rbac/permissions/assigned_by_users/ | 
[**RbacPermissionsAssignedByUsersUnassignPartialUpdate**](RbacApi.md#RbacPermissionsAssignedByUsersUnassignPartialUpdate) | **Patch** /rbac/permissions/assigned_by_users/{id}/unassign/ | 
[**RbacPermissionsList**](RbacApi.md#RbacPermissionsList) | **Get** /rbac/permissions/ | 
[**RbacPermissionsRetrieve**](RbacApi.md#RbacPermissionsRetrieve) | **Get** /rbac/permissions/{id}/ | 
[**RbacPermissionsRolesList**](RbacApi.md#RbacPermissionsRolesList) | **Get** /rbac/permissions/roles/ | 
[**RbacPermissionsUsersList**](RbacApi.md#RbacPermissionsUsersList) | **Get** /rbac/permissions/users/ | 
[**RbacRolesCreate**](RbacApi.md#RbacRolesCreate) | **Post** /rbac/roles/ | 
[**RbacRolesDestroy**](RbacApi.md#RbacRolesDestroy) | **Delete** /rbac/roles/{uuid}/ | 
[**RbacRolesList**](RbacApi.md#RbacRolesList) | **Get** /rbac/roles/ | 
[**RbacRolesPartialUpdate**](RbacApi.md#RbacRolesPartialUpdate) | **Patch** /rbac/roles/{uuid}/ | 
[**RbacRolesRetrieve**](RbacApi.md#RbacRolesRetrieve) | **Get** /rbac/roles/{uuid}/ | 
[**RbacRolesUpdate**](RbacApi.md#RbacRolesUpdate) | **Put** /rbac/roles/{uuid}/ | 
[**RbacRolesUsedByList**](RbacApi.md#RbacRolesUsedByList) | **Get** /rbac/roles/{uuid}/used_by/ | 



## RbacPermissionsAssignedByRolesAssignCreate

> RbacPermissionsAssignedByRolesAssignCreate(ctx, uuid).PermissionAssignRequest(permissionAssignRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.
    permissionAssignRequest := *openapiclient.NewPermissionAssignRequest([]string{"Permissions_example"}) // PermissionAssignRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacPermissionsAssignedByRolesAssignCreate(context.Background(), uuid).PermissionAssignRequest(permissionAssignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacPermissionsAssignedByRolesAssignCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsAssignedByRolesAssignCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permissionAssignRequest** | [**PermissionAssignRequest**](PermissionAssignRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsAssignedByRolesList

> PaginatedRoleAssignedObjectPermissionList RbacPermissionsAssignedByRolesList(ctx).Model(model).ObjectPk(objectPk).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    model := "model_example" // string | * `authentik_tenants.domain` - Domain * `authentik_crypto.certificatekeypair` - Certificate-Key Pair * `authentik_events.event` - Event * `authentik_events.notificationtransport` - Notification Transport * `authentik_events.notification` - Notification * `authentik_events.notificationrule` - Notification Rule * `authentik_events.notificationwebhookmapping` - Webhook Mapping * `authentik_flows.flow` - Flow * `authentik_flows.flowstagebinding` - Flow Stage Binding * `authentik_outposts.dockerserviceconnection` - Docker Service-Connection * `authentik_outposts.kubernetesserviceconnection` - Kubernetes Service-Connection * `authentik_outposts.outpost` - Outpost * `authentik_policies_dummy.dummypolicy` - Dummy Policy * `authentik_policies_event_matcher.eventmatcherpolicy` - Event Matcher Policy * `authentik_policies_expiry.passwordexpirypolicy` - Password Expiry Policy * `authentik_policies_expression.expressionpolicy` - Expression Policy * `authentik_policies_password.passwordpolicy` - Password Policy * `authentik_policies_reputation.reputationpolicy` - Reputation Policy * `authentik_policies_reputation.reputation` - Reputation Score * `authentik_policies.policybinding` - Policy Binding * `authentik_providers_ldap.ldapprovider` - LDAP Provider * `authentik_providers_oauth2.scopemapping` - Scope Mapping * `authentik_providers_oauth2.oauth2provider` - OAuth2/OpenID Provider * `authentik_providers_oauth2.authorizationcode` - Authorization Code * `authentik_providers_oauth2.accesstoken` - OAuth2 Access Token * `authentik_providers_oauth2.refreshtoken` - OAuth2 Refresh Token * `authentik_providers_proxy.proxyprovider` - Proxy Provider * `authentik_providers_radius.radiusprovider` - Radius Provider * `authentik_providers_saml.samlprovider` - SAML Provider * `authentik_providers_saml.samlpropertymapping` - SAML Property Mapping * `authentik_providers_scim.scimprovider` - SCIM Provider * `authentik_providers_scim.scimmapping` - SCIM Mapping * `authentik_rbac.role` - Role * `authentik_sources_ldap.ldapsource` - LDAP Source * `authentik_sources_ldap.ldappropertymapping` - LDAP Property Mapping * `authentik_sources_oauth.oauthsource` - OAuth Source * `authentik_sources_oauth.useroauthsourceconnection` - User OAuth Source Connection * `authentik_sources_plex.plexsource` - Plex Source * `authentik_sources_plex.plexsourceconnection` - User Plex Source Connection * `authentik_sources_saml.samlsource` - SAML Source * `authentik_sources_saml.usersamlsourceconnection` - User SAML Source Connection * `authentik_stages_authenticator_duo.authenticatorduostage` - Duo Authenticator Setup Stage * `authentik_stages_authenticator_duo.duodevice` - Duo Device * `authentik_stages_authenticator_sms.authenticatorsmsstage` - SMS Authenticator Setup Stage * `authentik_stages_authenticator_sms.smsdevice` - SMS Device * `authentik_stages_authenticator_static.authenticatorstaticstage` - Static Authenticator Stage * `authentik_stages_authenticator_static.staticdevice` - Static Device * `authentik_stages_authenticator_totp.authenticatortotpstage` - TOTP Authenticator Setup Stage * `authentik_stages_authenticator_totp.totpdevice` - TOTP Device * `authentik_stages_authenticator_validate.authenticatorvalidatestage` - Authenticator Validation Stage * `authentik_stages_authenticator_webauthn.authenticatewebauthnstage` - WebAuthn Authenticator Setup Stage * `authentik_stages_authenticator_webauthn.webauthndevice` - WebAuthn Device * `authentik_stages_captcha.captchastage` - Captcha Stage * `authentik_stages_consent.consentstage` - Consent Stage * `authentik_stages_consent.userconsent` - User Consent * `authentik_stages_deny.denystage` - Deny Stage * `authentik_stages_dummy.dummystage` - Dummy Stage * `authentik_stages_email.emailstage` - Email Stage * `authentik_stages_identification.identificationstage` - Identification Stage * `authentik_stages_invitation.invitationstage` - Invitation Stage * `authentik_stages_invitation.invitation` - Invitation * `authentik_stages_password.passwordstage` - Password Stage * `authentik_stages_prompt.prompt` - Prompt * `authentik_stages_prompt.promptstage` - Prompt Stage * `authentik_stages_user_delete.userdeletestage` - User Delete Stage * `authentik_stages_user_login.userloginstage` - User Login Stage * `authentik_stages_user_logout.userlogoutstage` - User Logout Stage * `authentik_stages_user_write.userwritestage` - User Write Stage * `authentik_brands.brand` - Brand * `authentik_blueprints.blueprintinstance` - Blueprint Instance * `authentik_core.group` - Group * `authentik_core.user` - User * `authentik_core.application` - Application * `authentik_core.token` - Token * `authentik_enterprise.license` - License * `authentik_providers_rac.racprovider` - RAC Provider * `authentik_providers_rac.endpoint` - RAC Endpoint * `authentik_providers_rac.racpropertymapping` - RAC Property Mapping
    objectPk := "objectPk_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacPermissionsAssignedByRolesList(context.Background()).Model(model).ObjectPk(objectPk).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacPermissionsAssignedByRolesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacPermissionsAssignedByRolesList`: PaginatedRoleAssignedObjectPermissionList
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacPermissionsAssignedByRolesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsAssignedByRolesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | **string** | * &#x60;authentik_tenants.domain&#x60; - Domain * &#x60;authentik_crypto.certificatekeypair&#x60; - Certificate-Key Pair * &#x60;authentik_events.event&#x60; - Event * &#x60;authentik_events.notificationtransport&#x60; - Notification Transport * &#x60;authentik_events.notification&#x60; - Notification * &#x60;authentik_events.notificationrule&#x60; - Notification Rule * &#x60;authentik_events.notificationwebhookmapping&#x60; - Webhook Mapping * &#x60;authentik_flows.flow&#x60; - Flow * &#x60;authentik_flows.flowstagebinding&#x60; - Flow Stage Binding * &#x60;authentik_outposts.dockerserviceconnection&#x60; - Docker Service-Connection * &#x60;authentik_outposts.kubernetesserviceconnection&#x60; - Kubernetes Service-Connection * &#x60;authentik_outposts.outpost&#x60; - Outpost * &#x60;authentik_policies_dummy.dummypolicy&#x60; - Dummy Policy * &#x60;authentik_policies_event_matcher.eventmatcherpolicy&#x60; - Event Matcher Policy * &#x60;authentik_policies_expiry.passwordexpirypolicy&#x60; - Password Expiry Policy * &#x60;authentik_policies_expression.expressionpolicy&#x60; - Expression Policy * &#x60;authentik_policies_password.passwordpolicy&#x60; - Password Policy * &#x60;authentik_policies_reputation.reputationpolicy&#x60; - Reputation Policy * &#x60;authentik_policies_reputation.reputation&#x60; - Reputation Score * &#x60;authentik_policies.policybinding&#x60; - Policy Binding * &#x60;authentik_providers_ldap.ldapprovider&#x60; - LDAP Provider * &#x60;authentik_providers_oauth2.scopemapping&#x60; - Scope Mapping * &#x60;authentik_providers_oauth2.oauth2provider&#x60; - OAuth2/OpenID Provider * &#x60;authentik_providers_oauth2.authorizationcode&#x60; - Authorization Code * &#x60;authentik_providers_oauth2.accesstoken&#x60; - OAuth2 Access Token * &#x60;authentik_providers_oauth2.refreshtoken&#x60; - OAuth2 Refresh Token * &#x60;authentik_providers_proxy.proxyprovider&#x60; - Proxy Provider * &#x60;authentik_providers_radius.radiusprovider&#x60; - Radius Provider * &#x60;authentik_providers_saml.samlprovider&#x60; - SAML Provider * &#x60;authentik_providers_saml.samlpropertymapping&#x60; - SAML Property Mapping * &#x60;authentik_providers_scim.scimprovider&#x60; - SCIM Provider * &#x60;authentik_providers_scim.scimmapping&#x60; - SCIM Mapping * &#x60;authentik_rbac.role&#x60; - Role * &#x60;authentik_sources_ldap.ldapsource&#x60; - LDAP Source * &#x60;authentik_sources_ldap.ldappropertymapping&#x60; - LDAP Property Mapping * &#x60;authentik_sources_oauth.oauthsource&#x60; - OAuth Source * &#x60;authentik_sources_oauth.useroauthsourceconnection&#x60; - User OAuth Source Connection * &#x60;authentik_sources_plex.plexsource&#x60; - Plex Source * &#x60;authentik_sources_plex.plexsourceconnection&#x60; - User Plex Source Connection * &#x60;authentik_sources_saml.samlsource&#x60; - SAML Source * &#x60;authentik_sources_saml.usersamlsourceconnection&#x60; - User SAML Source Connection * &#x60;authentik_stages_authenticator_duo.authenticatorduostage&#x60; - Duo Authenticator Setup Stage * &#x60;authentik_stages_authenticator_duo.duodevice&#x60; - Duo Device * &#x60;authentik_stages_authenticator_sms.authenticatorsmsstage&#x60; - SMS Authenticator Setup Stage * &#x60;authentik_stages_authenticator_sms.smsdevice&#x60; - SMS Device * &#x60;authentik_stages_authenticator_static.authenticatorstaticstage&#x60; - Static Authenticator Stage * &#x60;authentik_stages_authenticator_static.staticdevice&#x60; - Static Device * &#x60;authentik_stages_authenticator_totp.authenticatortotpstage&#x60; - TOTP Authenticator Setup Stage * &#x60;authentik_stages_authenticator_totp.totpdevice&#x60; - TOTP Device * &#x60;authentik_stages_authenticator_validate.authenticatorvalidatestage&#x60; - Authenticator Validation Stage * &#x60;authentik_stages_authenticator_webauthn.authenticatewebauthnstage&#x60; - WebAuthn Authenticator Setup Stage * &#x60;authentik_stages_authenticator_webauthn.webauthndevice&#x60; - WebAuthn Device * &#x60;authentik_stages_captcha.captchastage&#x60; - Captcha Stage * &#x60;authentik_stages_consent.consentstage&#x60; - Consent Stage * &#x60;authentik_stages_consent.userconsent&#x60; - User Consent * &#x60;authentik_stages_deny.denystage&#x60; - Deny Stage * &#x60;authentik_stages_dummy.dummystage&#x60; - Dummy Stage * &#x60;authentik_stages_email.emailstage&#x60; - Email Stage * &#x60;authentik_stages_identification.identificationstage&#x60; - Identification Stage * &#x60;authentik_stages_invitation.invitationstage&#x60; - Invitation Stage * &#x60;authentik_stages_invitation.invitation&#x60; - Invitation * &#x60;authentik_stages_password.passwordstage&#x60; - Password Stage * &#x60;authentik_stages_prompt.prompt&#x60; - Prompt * &#x60;authentik_stages_prompt.promptstage&#x60; - Prompt Stage * &#x60;authentik_stages_user_delete.userdeletestage&#x60; - User Delete Stage * &#x60;authentik_stages_user_login.userloginstage&#x60; - User Login Stage * &#x60;authentik_stages_user_logout.userlogoutstage&#x60; - User Logout Stage * &#x60;authentik_stages_user_write.userwritestage&#x60; - User Write Stage * &#x60;authentik_brands.brand&#x60; - Brand * &#x60;authentik_blueprints.blueprintinstance&#x60; - Blueprint Instance * &#x60;authentik_core.group&#x60; - Group * &#x60;authentik_core.user&#x60; - User * &#x60;authentik_core.application&#x60; - Application * &#x60;authentik_core.token&#x60; - Token * &#x60;authentik_enterprise.license&#x60; - License * &#x60;authentik_providers_rac.racprovider&#x60; - RAC Provider * &#x60;authentik_providers_rac.endpoint&#x60; - RAC Endpoint * &#x60;authentik_providers_rac.racpropertymapping&#x60; - RAC Property Mapping | 
 **objectPk** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedRoleAssignedObjectPermissionList**](PaginatedRoleAssignedObjectPermissionList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsAssignedByRolesUnassignPartialUpdate

> RbacPermissionsAssignedByRolesUnassignPartialUpdate(ctx, uuid).PatchedPermissionAssignRequest(patchedPermissionAssignRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.
    patchedPermissionAssignRequest := *openapiclient.NewPatchedPermissionAssignRequest() // PatchedPermissionAssignRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacPermissionsAssignedByRolesUnassignPartialUpdate(context.Background(), uuid).PatchedPermissionAssignRequest(patchedPermissionAssignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacPermissionsAssignedByRolesUnassignPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsAssignedByRolesUnassignPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPermissionAssignRequest** | [**PatchedPermissionAssignRequest**](PatchedPermissionAssignRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsAssignedByUsersAssignCreate

> RbacPermissionsAssignedByUsersAssignCreate(ctx, id).PermissionAssignRequest(permissionAssignRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this User.
    permissionAssignRequest := *openapiclient.NewPermissionAssignRequest([]string{"Permissions_example"}) // PermissionAssignRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacPermissionsAssignedByUsersAssignCreate(context.Background(), id).PermissionAssignRequest(permissionAssignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacPermissionsAssignedByUsersAssignCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsAssignedByUsersAssignCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permissionAssignRequest** | [**PermissionAssignRequest**](PermissionAssignRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsAssignedByUsersList

> PaginatedUserAssignedObjectPermissionList RbacPermissionsAssignedByUsersList(ctx).Model(model).ObjectPk(objectPk).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    model := "model_example" // string | * `authentik_tenants.domain` - Domain * `authentik_crypto.certificatekeypair` - Certificate-Key Pair * `authentik_events.event` - Event * `authentik_events.notificationtransport` - Notification Transport * `authentik_events.notification` - Notification * `authentik_events.notificationrule` - Notification Rule * `authentik_events.notificationwebhookmapping` - Webhook Mapping * `authentik_flows.flow` - Flow * `authentik_flows.flowstagebinding` - Flow Stage Binding * `authentik_outposts.dockerserviceconnection` - Docker Service-Connection * `authentik_outposts.kubernetesserviceconnection` - Kubernetes Service-Connection * `authentik_outposts.outpost` - Outpost * `authentik_policies_dummy.dummypolicy` - Dummy Policy * `authentik_policies_event_matcher.eventmatcherpolicy` - Event Matcher Policy * `authentik_policies_expiry.passwordexpirypolicy` - Password Expiry Policy * `authentik_policies_expression.expressionpolicy` - Expression Policy * `authentik_policies_password.passwordpolicy` - Password Policy * `authentik_policies_reputation.reputationpolicy` - Reputation Policy * `authentik_policies_reputation.reputation` - Reputation Score * `authentik_policies.policybinding` - Policy Binding * `authentik_providers_ldap.ldapprovider` - LDAP Provider * `authentik_providers_oauth2.scopemapping` - Scope Mapping * `authentik_providers_oauth2.oauth2provider` - OAuth2/OpenID Provider * `authentik_providers_oauth2.authorizationcode` - Authorization Code * `authentik_providers_oauth2.accesstoken` - OAuth2 Access Token * `authentik_providers_oauth2.refreshtoken` - OAuth2 Refresh Token * `authentik_providers_proxy.proxyprovider` - Proxy Provider * `authentik_providers_radius.radiusprovider` - Radius Provider * `authentik_providers_saml.samlprovider` - SAML Provider * `authentik_providers_saml.samlpropertymapping` - SAML Property Mapping * `authentik_providers_scim.scimprovider` - SCIM Provider * `authentik_providers_scim.scimmapping` - SCIM Mapping * `authentik_rbac.role` - Role * `authentik_sources_ldap.ldapsource` - LDAP Source * `authentik_sources_ldap.ldappropertymapping` - LDAP Property Mapping * `authentik_sources_oauth.oauthsource` - OAuth Source * `authentik_sources_oauth.useroauthsourceconnection` - User OAuth Source Connection * `authentik_sources_plex.plexsource` - Plex Source * `authentik_sources_plex.plexsourceconnection` - User Plex Source Connection * `authentik_sources_saml.samlsource` - SAML Source * `authentik_sources_saml.usersamlsourceconnection` - User SAML Source Connection * `authentik_stages_authenticator_duo.authenticatorduostage` - Duo Authenticator Setup Stage * `authentik_stages_authenticator_duo.duodevice` - Duo Device * `authentik_stages_authenticator_sms.authenticatorsmsstage` - SMS Authenticator Setup Stage * `authentik_stages_authenticator_sms.smsdevice` - SMS Device * `authentik_stages_authenticator_static.authenticatorstaticstage` - Static Authenticator Stage * `authentik_stages_authenticator_static.staticdevice` - Static Device * `authentik_stages_authenticator_totp.authenticatortotpstage` - TOTP Authenticator Setup Stage * `authentik_stages_authenticator_totp.totpdevice` - TOTP Device * `authentik_stages_authenticator_validate.authenticatorvalidatestage` - Authenticator Validation Stage * `authentik_stages_authenticator_webauthn.authenticatewebauthnstage` - WebAuthn Authenticator Setup Stage * `authentik_stages_authenticator_webauthn.webauthndevice` - WebAuthn Device * `authentik_stages_captcha.captchastage` - Captcha Stage * `authentik_stages_consent.consentstage` - Consent Stage * `authentik_stages_consent.userconsent` - User Consent * `authentik_stages_deny.denystage` - Deny Stage * `authentik_stages_dummy.dummystage` - Dummy Stage * `authentik_stages_email.emailstage` - Email Stage * `authentik_stages_identification.identificationstage` - Identification Stage * `authentik_stages_invitation.invitationstage` - Invitation Stage * `authentik_stages_invitation.invitation` - Invitation * `authentik_stages_password.passwordstage` - Password Stage * `authentik_stages_prompt.prompt` - Prompt * `authentik_stages_prompt.promptstage` - Prompt Stage * `authentik_stages_user_delete.userdeletestage` - User Delete Stage * `authentik_stages_user_login.userloginstage` - User Login Stage * `authentik_stages_user_logout.userlogoutstage` - User Logout Stage * `authentik_stages_user_write.userwritestage` - User Write Stage * `authentik_brands.brand` - Brand * `authentik_blueprints.blueprintinstance` - Blueprint Instance * `authentik_core.group` - Group * `authentik_core.user` - User * `authentik_core.application` - Application * `authentik_core.token` - Token * `authentik_enterprise.license` - License * `authentik_providers_rac.racprovider` - RAC Provider * `authentik_providers_rac.endpoint` - RAC Endpoint * `authentik_providers_rac.racpropertymapping` - RAC Property Mapping
    objectPk := "objectPk_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacPermissionsAssignedByUsersList(context.Background()).Model(model).ObjectPk(objectPk).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacPermissionsAssignedByUsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacPermissionsAssignedByUsersList`: PaginatedUserAssignedObjectPermissionList
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacPermissionsAssignedByUsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsAssignedByUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | **string** | * &#x60;authentik_tenants.domain&#x60; - Domain * &#x60;authentik_crypto.certificatekeypair&#x60; - Certificate-Key Pair * &#x60;authentik_events.event&#x60; - Event * &#x60;authentik_events.notificationtransport&#x60; - Notification Transport * &#x60;authentik_events.notification&#x60; - Notification * &#x60;authentik_events.notificationrule&#x60; - Notification Rule * &#x60;authentik_events.notificationwebhookmapping&#x60; - Webhook Mapping * &#x60;authentik_flows.flow&#x60; - Flow * &#x60;authentik_flows.flowstagebinding&#x60; - Flow Stage Binding * &#x60;authentik_outposts.dockerserviceconnection&#x60; - Docker Service-Connection * &#x60;authentik_outposts.kubernetesserviceconnection&#x60; - Kubernetes Service-Connection * &#x60;authentik_outposts.outpost&#x60; - Outpost * &#x60;authentik_policies_dummy.dummypolicy&#x60; - Dummy Policy * &#x60;authentik_policies_event_matcher.eventmatcherpolicy&#x60; - Event Matcher Policy * &#x60;authentik_policies_expiry.passwordexpirypolicy&#x60; - Password Expiry Policy * &#x60;authentik_policies_expression.expressionpolicy&#x60; - Expression Policy * &#x60;authentik_policies_password.passwordpolicy&#x60; - Password Policy * &#x60;authentik_policies_reputation.reputationpolicy&#x60; - Reputation Policy * &#x60;authentik_policies_reputation.reputation&#x60; - Reputation Score * &#x60;authentik_policies.policybinding&#x60; - Policy Binding * &#x60;authentik_providers_ldap.ldapprovider&#x60; - LDAP Provider * &#x60;authentik_providers_oauth2.scopemapping&#x60; - Scope Mapping * &#x60;authentik_providers_oauth2.oauth2provider&#x60; - OAuth2/OpenID Provider * &#x60;authentik_providers_oauth2.authorizationcode&#x60; - Authorization Code * &#x60;authentik_providers_oauth2.accesstoken&#x60; - OAuth2 Access Token * &#x60;authentik_providers_oauth2.refreshtoken&#x60; - OAuth2 Refresh Token * &#x60;authentik_providers_proxy.proxyprovider&#x60; - Proxy Provider * &#x60;authentik_providers_radius.radiusprovider&#x60; - Radius Provider * &#x60;authentik_providers_saml.samlprovider&#x60; - SAML Provider * &#x60;authentik_providers_saml.samlpropertymapping&#x60; - SAML Property Mapping * &#x60;authentik_providers_scim.scimprovider&#x60; - SCIM Provider * &#x60;authentik_providers_scim.scimmapping&#x60; - SCIM Mapping * &#x60;authentik_rbac.role&#x60; - Role * &#x60;authentik_sources_ldap.ldapsource&#x60; - LDAP Source * &#x60;authentik_sources_ldap.ldappropertymapping&#x60; - LDAP Property Mapping * &#x60;authentik_sources_oauth.oauthsource&#x60; - OAuth Source * &#x60;authentik_sources_oauth.useroauthsourceconnection&#x60; - User OAuth Source Connection * &#x60;authentik_sources_plex.plexsource&#x60; - Plex Source * &#x60;authentik_sources_plex.plexsourceconnection&#x60; - User Plex Source Connection * &#x60;authentik_sources_saml.samlsource&#x60; - SAML Source * &#x60;authentik_sources_saml.usersamlsourceconnection&#x60; - User SAML Source Connection * &#x60;authentik_stages_authenticator_duo.authenticatorduostage&#x60; - Duo Authenticator Setup Stage * &#x60;authentik_stages_authenticator_duo.duodevice&#x60; - Duo Device * &#x60;authentik_stages_authenticator_sms.authenticatorsmsstage&#x60; - SMS Authenticator Setup Stage * &#x60;authentik_stages_authenticator_sms.smsdevice&#x60; - SMS Device * &#x60;authentik_stages_authenticator_static.authenticatorstaticstage&#x60; - Static Authenticator Stage * &#x60;authentik_stages_authenticator_static.staticdevice&#x60; - Static Device * &#x60;authentik_stages_authenticator_totp.authenticatortotpstage&#x60; - TOTP Authenticator Setup Stage * &#x60;authentik_stages_authenticator_totp.totpdevice&#x60; - TOTP Device * &#x60;authentik_stages_authenticator_validate.authenticatorvalidatestage&#x60; - Authenticator Validation Stage * &#x60;authentik_stages_authenticator_webauthn.authenticatewebauthnstage&#x60; - WebAuthn Authenticator Setup Stage * &#x60;authentik_stages_authenticator_webauthn.webauthndevice&#x60; - WebAuthn Device * &#x60;authentik_stages_captcha.captchastage&#x60; - Captcha Stage * &#x60;authentik_stages_consent.consentstage&#x60; - Consent Stage * &#x60;authentik_stages_consent.userconsent&#x60; - User Consent * &#x60;authentik_stages_deny.denystage&#x60; - Deny Stage * &#x60;authentik_stages_dummy.dummystage&#x60; - Dummy Stage * &#x60;authentik_stages_email.emailstage&#x60; - Email Stage * &#x60;authentik_stages_identification.identificationstage&#x60; - Identification Stage * &#x60;authentik_stages_invitation.invitationstage&#x60; - Invitation Stage * &#x60;authentik_stages_invitation.invitation&#x60; - Invitation * &#x60;authentik_stages_password.passwordstage&#x60; - Password Stage * &#x60;authentik_stages_prompt.prompt&#x60; - Prompt * &#x60;authentik_stages_prompt.promptstage&#x60; - Prompt Stage * &#x60;authentik_stages_user_delete.userdeletestage&#x60; - User Delete Stage * &#x60;authentik_stages_user_login.userloginstage&#x60; - User Login Stage * &#x60;authentik_stages_user_logout.userlogoutstage&#x60; - User Logout Stage * &#x60;authentik_stages_user_write.userwritestage&#x60; - User Write Stage * &#x60;authentik_brands.brand&#x60; - Brand * &#x60;authentik_blueprints.blueprintinstance&#x60; - Blueprint Instance * &#x60;authentik_core.group&#x60; - Group * &#x60;authentik_core.user&#x60; - User * &#x60;authentik_core.application&#x60; - Application * &#x60;authentik_core.token&#x60; - Token * &#x60;authentik_enterprise.license&#x60; - License * &#x60;authentik_providers_rac.racprovider&#x60; - RAC Provider * &#x60;authentik_providers_rac.endpoint&#x60; - RAC Endpoint * &#x60;authentik_providers_rac.racpropertymapping&#x60; - RAC Property Mapping | 
 **objectPk** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedUserAssignedObjectPermissionList**](PaginatedUserAssignedObjectPermissionList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsAssignedByUsersUnassignPartialUpdate

> RbacPermissionsAssignedByUsersUnassignPartialUpdate(ctx, id).PatchedPermissionAssignRequest(patchedPermissionAssignRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this User.
    patchedPermissionAssignRequest := *openapiclient.NewPatchedPermissionAssignRequest() // PatchedPermissionAssignRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacPermissionsAssignedByUsersUnassignPartialUpdate(context.Background(), id).PatchedPermissionAssignRequest(patchedPermissionAssignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacPermissionsAssignedByUsersUnassignPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsAssignedByUsersUnassignPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPermissionAssignRequest** | [**PatchedPermissionAssignRequest**](PatchedPermissionAssignRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsList

> PaginatedPermissionList RbacPermissionsList(ctx).Codename(codename).ContentTypeAppLabel(contentTypeAppLabel).ContentTypeModel(contentTypeModel).Ordering(ordering).Page(page).PageSize(pageSize).Role(role).Search(search).User(user).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    codename := "codename_example" // string |  (optional)
    contentTypeAppLabel := "contentTypeAppLabel_example" // string |  (optional)
    contentTypeModel := "contentTypeModel_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    role := "role_example" // string |  (optional)
    search := "search_example" // string | A search term. (optional)
    user := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacPermissionsList(context.Background()).Codename(codename).ContentTypeAppLabel(contentTypeAppLabel).ContentTypeModel(contentTypeModel).Ordering(ordering).Page(page).PageSize(pageSize).Role(role).Search(search).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacPermissionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacPermissionsList`: PaginatedPermissionList
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacPermissionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codename** | **string** |  | 
 **contentTypeAppLabel** | **string** |  | 
 **contentTypeModel** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **role** | **string** |  | 
 **search** | **string** | A search term. | 
 **user** | **int32** |  | 

### Return type

[**PaginatedPermissionList**](PaginatedPermissionList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsRetrieve

> Permission RbacPermissionsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this permission.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacPermissionsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacPermissionsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacPermissionsRetrieve`: Permission
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacPermissionsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Permission**](Permission.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsRolesList

> PaginatedExtraRoleObjectPermissionList RbacPermissionsRolesList(ctx).Uuid(uuid).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacPermissionsRolesList(context.Background()).Uuid(uuid).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacPermissionsRolesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacPermissionsRolesList`: PaginatedExtraRoleObjectPermissionList
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacPermissionsRolesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsRolesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedExtraRoleObjectPermissionList**](PaginatedExtraRoleObjectPermissionList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacPermissionsUsersList

> PaginatedExtraUserObjectPermissionList RbacPermissionsUsersList(ctx).UserId(userId).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userId := int32(56) // int32 | 
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacPermissionsUsersList(context.Background()).UserId(userId).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacPermissionsUsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacPermissionsUsersList`: PaginatedExtraUserObjectPermissionList
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacPermissionsUsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRbacPermissionsUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedExtraUserObjectPermissionList**](PaginatedExtraUserObjectPermissionList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacRolesCreate

> Role RbacRolesCreate(ctx).RoleRequest(roleRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleRequest := *openapiclient.NewRoleRequest("Name_example") // RoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacRolesCreate(context.Background()).RoleRequest(roleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacRolesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacRolesCreate`: Role
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacRolesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRbacRolesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleRequest** | [**RoleRequest**](RoleRequest.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacRolesDestroy

> RbacRolesDestroy(ctx, uuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacRolesDestroy(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacRolesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacRolesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacRolesList

> PaginatedRoleList RbacRolesList(ctx).GroupName(groupName).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupName := "groupName_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacRolesList(context.Background()).GroupName(groupName).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacRolesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacRolesList`: PaginatedRoleList
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacRolesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRbacRolesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedRoleList**](PaginatedRoleList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacRolesPartialUpdate

> Role RbacRolesPartialUpdate(ctx, uuid).PatchedRoleRequest(patchedRoleRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.
    patchedRoleRequest := *openapiclient.NewPatchedRoleRequest() // PatchedRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacRolesPartialUpdate(context.Background(), uuid).PatchedRoleRequest(patchedRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacRolesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacRolesPartialUpdate`: Role
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacRolesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacRolesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedRoleRequest** | [**PatchedRoleRequest**](PatchedRoleRequest.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacRolesRetrieve

> Role RbacRolesRetrieve(ctx, uuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacRolesRetrieve(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacRolesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacRolesRetrieve`: Role
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacRolesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacRolesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Role**](Role.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacRolesUpdate

> Role RbacRolesUpdate(ctx, uuid).RoleRequest(roleRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.
    roleRequest := *openapiclient.NewRoleRequest("Name_example") // RoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacRolesUpdate(context.Background(), uuid).RoleRequest(roleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacRolesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacRolesUpdate`: Role
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacRolesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacRolesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleRequest** | [**RoleRequest**](RoleRequest.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RbacRolesUsedByList

> []UsedBy RbacRolesUsedByList(ctx, uuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacApi.RbacRolesUsedByList(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacApi.RbacRolesUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RbacRolesUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `RbacApi.RbacRolesUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRbacRolesUsedByListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UsedBy**](UsedBy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

