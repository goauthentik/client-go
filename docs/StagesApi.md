# \StagesApi

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StagesAllDestroy**](StagesApi.md#StagesAllDestroy) | **Delete** /stages/all/{stage_uuid}/ | 
[**StagesAllList**](StagesApi.md#StagesAllList) | **Get** /stages/all/ | 
[**StagesAllRetrieve**](StagesApi.md#StagesAllRetrieve) | **Get** /stages/all/{stage_uuid}/ | 
[**StagesAllTypesList**](StagesApi.md#StagesAllTypesList) | **Get** /stages/all/types/ | 
[**StagesAllUsedByList**](StagesApi.md#StagesAllUsedByList) | **Get** /stages/all/{stage_uuid}/used_by/ | 
[**StagesAllUserSettingsList**](StagesApi.md#StagesAllUserSettingsList) | **Get** /stages/all/user_settings/ | 
[**StagesAuthenticatorDuoCreate**](StagesApi.md#StagesAuthenticatorDuoCreate) | **Post** /stages/authenticator/duo/ | 
[**StagesAuthenticatorDuoDestroy**](StagesApi.md#StagesAuthenticatorDuoDestroy) | **Delete** /stages/authenticator/duo/{stage_uuid}/ | 
[**StagesAuthenticatorDuoEnrollmentStatusCreate**](StagesApi.md#StagesAuthenticatorDuoEnrollmentStatusCreate) | **Post** /stages/authenticator/duo/{stage_uuid}/enrollment_status/ | 
[**StagesAuthenticatorDuoImportDeviceManualCreate**](StagesApi.md#StagesAuthenticatorDuoImportDeviceManualCreate) | **Post** /stages/authenticator/duo/{stage_uuid}/import_device_manual/ | 
[**StagesAuthenticatorDuoImportDevicesAutomaticCreate**](StagesApi.md#StagesAuthenticatorDuoImportDevicesAutomaticCreate) | **Post** /stages/authenticator/duo/{stage_uuid}/import_devices_automatic/ | 
[**StagesAuthenticatorDuoList**](StagesApi.md#StagesAuthenticatorDuoList) | **Get** /stages/authenticator/duo/ | 
[**StagesAuthenticatorDuoPartialUpdate**](StagesApi.md#StagesAuthenticatorDuoPartialUpdate) | **Patch** /stages/authenticator/duo/{stage_uuid}/ | 
[**StagesAuthenticatorDuoRetrieve**](StagesApi.md#StagesAuthenticatorDuoRetrieve) | **Get** /stages/authenticator/duo/{stage_uuid}/ | 
[**StagesAuthenticatorDuoUpdate**](StagesApi.md#StagesAuthenticatorDuoUpdate) | **Put** /stages/authenticator/duo/{stage_uuid}/ | 
[**StagesAuthenticatorDuoUsedByList**](StagesApi.md#StagesAuthenticatorDuoUsedByList) | **Get** /stages/authenticator/duo/{stage_uuid}/used_by/ | 
[**StagesAuthenticatorSmsCreate**](StagesApi.md#StagesAuthenticatorSmsCreate) | **Post** /stages/authenticator/sms/ | 
[**StagesAuthenticatorSmsDestroy**](StagesApi.md#StagesAuthenticatorSmsDestroy) | **Delete** /stages/authenticator/sms/{stage_uuid}/ | 
[**StagesAuthenticatorSmsList**](StagesApi.md#StagesAuthenticatorSmsList) | **Get** /stages/authenticator/sms/ | 
[**StagesAuthenticatorSmsPartialUpdate**](StagesApi.md#StagesAuthenticatorSmsPartialUpdate) | **Patch** /stages/authenticator/sms/{stage_uuid}/ | 
[**StagesAuthenticatorSmsRetrieve**](StagesApi.md#StagesAuthenticatorSmsRetrieve) | **Get** /stages/authenticator/sms/{stage_uuid}/ | 
[**StagesAuthenticatorSmsUpdate**](StagesApi.md#StagesAuthenticatorSmsUpdate) | **Put** /stages/authenticator/sms/{stage_uuid}/ | 
[**StagesAuthenticatorSmsUsedByList**](StagesApi.md#StagesAuthenticatorSmsUsedByList) | **Get** /stages/authenticator/sms/{stage_uuid}/used_by/ | 
[**StagesAuthenticatorStaticCreate**](StagesApi.md#StagesAuthenticatorStaticCreate) | **Post** /stages/authenticator/static/ | 
[**StagesAuthenticatorStaticDestroy**](StagesApi.md#StagesAuthenticatorStaticDestroy) | **Delete** /stages/authenticator/static/{stage_uuid}/ | 
[**StagesAuthenticatorStaticList**](StagesApi.md#StagesAuthenticatorStaticList) | **Get** /stages/authenticator/static/ | 
[**StagesAuthenticatorStaticPartialUpdate**](StagesApi.md#StagesAuthenticatorStaticPartialUpdate) | **Patch** /stages/authenticator/static/{stage_uuid}/ | 
[**StagesAuthenticatorStaticRetrieve**](StagesApi.md#StagesAuthenticatorStaticRetrieve) | **Get** /stages/authenticator/static/{stage_uuid}/ | 
[**StagesAuthenticatorStaticUpdate**](StagesApi.md#StagesAuthenticatorStaticUpdate) | **Put** /stages/authenticator/static/{stage_uuid}/ | 
[**StagesAuthenticatorStaticUsedByList**](StagesApi.md#StagesAuthenticatorStaticUsedByList) | **Get** /stages/authenticator/static/{stage_uuid}/used_by/ | 
[**StagesAuthenticatorTotpCreate**](StagesApi.md#StagesAuthenticatorTotpCreate) | **Post** /stages/authenticator/totp/ | 
[**StagesAuthenticatorTotpDestroy**](StagesApi.md#StagesAuthenticatorTotpDestroy) | **Delete** /stages/authenticator/totp/{stage_uuid}/ | 
[**StagesAuthenticatorTotpList**](StagesApi.md#StagesAuthenticatorTotpList) | **Get** /stages/authenticator/totp/ | 
[**StagesAuthenticatorTotpPartialUpdate**](StagesApi.md#StagesAuthenticatorTotpPartialUpdate) | **Patch** /stages/authenticator/totp/{stage_uuid}/ | 
[**StagesAuthenticatorTotpRetrieve**](StagesApi.md#StagesAuthenticatorTotpRetrieve) | **Get** /stages/authenticator/totp/{stage_uuid}/ | 
[**StagesAuthenticatorTotpUpdate**](StagesApi.md#StagesAuthenticatorTotpUpdate) | **Put** /stages/authenticator/totp/{stage_uuid}/ | 
[**StagesAuthenticatorTotpUsedByList**](StagesApi.md#StagesAuthenticatorTotpUsedByList) | **Get** /stages/authenticator/totp/{stage_uuid}/used_by/ | 
[**StagesAuthenticatorValidateCreate**](StagesApi.md#StagesAuthenticatorValidateCreate) | **Post** /stages/authenticator/validate/ | 
[**StagesAuthenticatorValidateDestroy**](StagesApi.md#StagesAuthenticatorValidateDestroy) | **Delete** /stages/authenticator/validate/{stage_uuid}/ | 
[**StagesAuthenticatorValidateList**](StagesApi.md#StagesAuthenticatorValidateList) | **Get** /stages/authenticator/validate/ | 
[**StagesAuthenticatorValidatePartialUpdate**](StagesApi.md#StagesAuthenticatorValidatePartialUpdate) | **Patch** /stages/authenticator/validate/{stage_uuid}/ | 
[**StagesAuthenticatorValidateRetrieve**](StagesApi.md#StagesAuthenticatorValidateRetrieve) | **Get** /stages/authenticator/validate/{stage_uuid}/ | 
[**StagesAuthenticatorValidateUpdate**](StagesApi.md#StagesAuthenticatorValidateUpdate) | **Put** /stages/authenticator/validate/{stage_uuid}/ | 
[**StagesAuthenticatorValidateUsedByList**](StagesApi.md#StagesAuthenticatorValidateUsedByList) | **Get** /stages/authenticator/validate/{stage_uuid}/used_by/ | 
[**StagesAuthenticatorWebauthnCreate**](StagesApi.md#StagesAuthenticatorWebauthnCreate) | **Post** /stages/authenticator/webauthn/ | 
[**StagesAuthenticatorWebauthnDestroy**](StagesApi.md#StagesAuthenticatorWebauthnDestroy) | **Delete** /stages/authenticator/webauthn/{stage_uuid}/ | 
[**StagesAuthenticatorWebauthnList**](StagesApi.md#StagesAuthenticatorWebauthnList) | **Get** /stages/authenticator/webauthn/ | 
[**StagesAuthenticatorWebauthnPartialUpdate**](StagesApi.md#StagesAuthenticatorWebauthnPartialUpdate) | **Patch** /stages/authenticator/webauthn/{stage_uuid}/ | 
[**StagesAuthenticatorWebauthnRetrieve**](StagesApi.md#StagesAuthenticatorWebauthnRetrieve) | **Get** /stages/authenticator/webauthn/{stage_uuid}/ | 
[**StagesAuthenticatorWebauthnUpdate**](StagesApi.md#StagesAuthenticatorWebauthnUpdate) | **Put** /stages/authenticator/webauthn/{stage_uuid}/ | 
[**StagesAuthenticatorWebauthnUsedByList**](StagesApi.md#StagesAuthenticatorWebauthnUsedByList) | **Get** /stages/authenticator/webauthn/{stage_uuid}/used_by/ | 
[**StagesCaptchaCreate**](StagesApi.md#StagesCaptchaCreate) | **Post** /stages/captcha/ | 
[**StagesCaptchaDestroy**](StagesApi.md#StagesCaptchaDestroy) | **Delete** /stages/captcha/{stage_uuid}/ | 
[**StagesCaptchaList**](StagesApi.md#StagesCaptchaList) | **Get** /stages/captcha/ | 
[**StagesCaptchaPartialUpdate**](StagesApi.md#StagesCaptchaPartialUpdate) | **Patch** /stages/captcha/{stage_uuid}/ | 
[**StagesCaptchaRetrieve**](StagesApi.md#StagesCaptchaRetrieve) | **Get** /stages/captcha/{stage_uuid}/ | 
[**StagesCaptchaUpdate**](StagesApi.md#StagesCaptchaUpdate) | **Put** /stages/captcha/{stage_uuid}/ | 
[**StagesCaptchaUsedByList**](StagesApi.md#StagesCaptchaUsedByList) | **Get** /stages/captcha/{stage_uuid}/used_by/ | 
[**StagesConsentCreate**](StagesApi.md#StagesConsentCreate) | **Post** /stages/consent/ | 
[**StagesConsentDestroy**](StagesApi.md#StagesConsentDestroy) | **Delete** /stages/consent/{stage_uuid}/ | 
[**StagesConsentList**](StagesApi.md#StagesConsentList) | **Get** /stages/consent/ | 
[**StagesConsentPartialUpdate**](StagesApi.md#StagesConsentPartialUpdate) | **Patch** /stages/consent/{stage_uuid}/ | 
[**StagesConsentRetrieve**](StagesApi.md#StagesConsentRetrieve) | **Get** /stages/consent/{stage_uuid}/ | 
[**StagesConsentUpdate**](StagesApi.md#StagesConsentUpdate) | **Put** /stages/consent/{stage_uuid}/ | 
[**StagesConsentUsedByList**](StagesApi.md#StagesConsentUsedByList) | **Get** /stages/consent/{stage_uuid}/used_by/ | 
[**StagesDenyCreate**](StagesApi.md#StagesDenyCreate) | **Post** /stages/deny/ | 
[**StagesDenyDestroy**](StagesApi.md#StagesDenyDestroy) | **Delete** /stages/deny/{stage_uuid}/ | 
[**StagesDenyList**](StagesApi.md#StagesDenyList) | **Get** /stages/deny/ | 
[**StagesDenyPartialUpdate**](StagesApi.md#StagesDenyPartialUpdate) | **Patch** /stages/deny/{stage_uuid}/ | 
[**StagesDenyRetrieve**](StagesApi.md#StagesDenyRetrieve) | **Get** /stages/deny/{stage_uuid}/ | 
[**StagesDenyUpdate**](StagesApi.md#StagesDenyUpdate) | **Put** /stages/deny/{stage_uuid}/ | 
[**StagesDenyUsedByList**](StagesApi.md#StagesDenyUsedByList) | **Get** /stages/deny/{stage_uuid}/used_by/ | 
[**StagesDummyCreate**](StagesApi.md#StagesDummyCreate) | **Post** /stages/dummy/ | 
[**StagesDummyDestroy**](StagesApi.md#StagesDummyDestroy) | **Delete** /stages/dummy/{stage_uuid}/ | 
[**StagesDummyList**](StagesApi.md#StagesDummyList) | **Get** /stages/dummy/ | 
[**StagesDummyPartialUpdate**](StagesApi.md#StagesDummyPartialUpdate) | **Patch** /stages/dummy/{stage_uuid}/ | 
[**StagesDummyRetrieve**](StagesApi.md#StagesDummyRetrieve) | **Get** /stages/dummy/{stage_uuid}/ | 
[**StagesDummyUpdate**](StagesApi.md#StagesDummyUpdate) | **Put** /stages/dummy/{stage_uuid}/ | 
[**StagesDummyUsedByList**](StagesApi.md#StagesDummyUsedByList) | **Get** /stages/dummy/{stage_uuid}/used_by/ | 
[**StagesEmailCreate**](StagesApi.md#StagesEmailCreate) | **Post** /stages/email/ | 
[**StagesEmailDestroy**](StagesApi.md#StagesEmailDestroy) | **Delete** /stages/email/{stage_uuid}/ | 
[**StagesEmailList**](StagesApi.md#StagesEmailList) | **Get** /stages/email/ | 
[**StagesEmailPartialUpdate**](StagesApi.md#StagesEmailPartialUpdate) | **Patch** /stages/email/{stage_uuid}/ | 
[**StagesEmailRetrieve**](StagesApi.md#StagesEmailRetrieve) | **Get** /stages/email/{stage_uuid}/ | 
[**StagesEmailTemplatesList**](StagesApi.md#StagesEmailTemplatesList) | **Get** /stages/email/templates/ | 
[**StagesEmailUpdate**](StagesApi.md#StagesEmailUpdate) | **Put** /stages/email/{stage_uuid}/ | 
[**StagesEmailUsedByList**](StagesApi.md#StagesEmailUsedByList) | **Get** /stages/email/{stage_uuid}/used_by/ | 
[**StagesIdentificationCreate**](StagesApi.md#StagesIdentificationCreate) | **Post** /stages/identification/ | 
[**StagesIdentificationDestroy**](StagesApi.md#StagesIdentificationDestroy) | **Delete** /stages/identification/{stage_uuid}/ | 
[**StagesIdentificationList**](StagesApi.md#StagesIdentificationList) | **Get** /stages/identification/ | 
[**StagesIdentificationPartialUpdate**](StagesApi.md#StagesIdentificationPartialUpdate) | **Patch** /stages/identification/{stage_uuid}/ | 
[**StagesIdentificationRetrieve**](StagesApi.md#StagesIdentificationRetrieve) | **Get** /stages/identification/{stage_uuid}/ | 
[**StagesIdentificationUpdate**](StagesApi.md#StagesIdentificationUpdate) | **Put** /stages/identification/{stage_uuid}/ | 
[**StagesIdentificationUsedByList**](StagesApi.md#StagesIdentificationUsedByList) | **Get** /stages/identification/{stage_uuid}/used_by/ | 
[**StagesInvitationInvitationsCreate**](StagesApi.md#StagesInvitationInvitationsCreate) | **Post** /stages/invitation/invitations/ | 
[**StagesInvitationInvitationsDestroy**](StagesApi.md#StagesInvitationInvitationsDestroy) | **Delete** /stages/invitation/invitations/{invite_uuid}/ | 
[**StagesInvitationInvitationsList**](StagesApi.md#StagesInvitationInvitationsList) | **Get** /stages/invitation/invitations/ | 
[**StagesInvitationInvitationsPartialUpdate**](StagesApi.md#StagesInvitationInvitationsPartialUpdate) | **Patch** /stages/invitation/invitations/{invite_uuid}/ | 
[**StagesInvitationInvitationsRetrieve**](StagesApi.md#StagesInvitationInvitationsRetrieve) | **Get** /stages/invitation/invitations/{invite_uuid}/ | 
[**StagesInvitationInvitationsUpdate**](StagesApi.md#StagesInvitationInvitationsUpdate) | **Put** /stages/invitation/invitations/{invite_uuid}/ | 
[**StagesInvitationInvitationsUsedByList**](StagesApi.md#StagesInvitationInvitationsUsedByList) | **Get** /stages/invitation/invitations/{invite_uuid}/used_by/ | 
[**StagesInvitationStagesCreate**](StagesApi.md#StagesInvitationStagesCreate) | **Post** /stages/invitation/stages/ | 
[**StagesInvitationStagesDestroy**](StagesApi.md#StagesInvitationStagesDestroy) | **Delete** /stages/invitation/stages/{stage_uuid}/ | 
[**StagesInvitationStagesList**](StagesApi.md#StagesInvitationStagesList) | **Get** /stages/invitation/stages/ | 
[**StagesInvitationStagesPartialUpdate**](StagesApi.md#StagesInvitationStagesPartialUpdate) | **Patch** /stages/invitation/stages/{stage_uuid}/ | 
[**StagesInvitationStagesRetrieve**](StagesApi.md#StagesInvitationStagesRetrieve) | **Get** /stages/invitation/stages/{stage_uuid}/ | 
[**StagesInvitationStagesUpdate**](StagesApi.md#StagesInvitationStagesUpdate) | **Put** /stages/invitation/stages/{stage_uuid}/ | 
[**StagesInvitationStagesUsedByList**](StagesApi.md#StagesInvitationStagesUsedByList) | **Get** /stages/invitation/stages/{stage_uuid}/used_by/ | 
[**StagesPasswordCreate**](StagesApi.md#StagesPasswordCreate) | **Post** /stages/password/ | 
[**StagesPasswordDestroy**](StagesApi.md#StagesPasswordDestroy) | **Delete** /stages/password/{stage_uuid}/ | 
[**StagesPasswordList**](StagesApi.md#StagesPasswordList) | **Get** /stages/password/ | 
[**StagesPasswordPartialUpdate**](StagesApi.md#StagesPasswordPartialUpdate) | **Patch** /stages/password/{stage_uuid}/ | 
[**StagesPasswordRetrieve**](StagesApi.md#StagesPasswordRetrieve) | **Get** /stages/password/{stage_uuid}/ | 
[**StagesPasswordUpdate**](StagesApi.md#StagesPasswordUpdate) | **Put** /stages/password/{stage_uuid}/ | 
[**StagesPasswordUsedByList**](StagesApi.md#StagesPasswordUsedByList) | **Get** /stages/password/{stage_uuid}/used_by/ | 
[**StagesPromptPromptsCreate**](StagesApi.md#StagesPromptPromptsCreate) | **Post** /stages/prompt/prompts/ | 
[**StagesPromptPromptsDestroy**](StagesApi.md#StagesPromptPromptsDestroy) | **Delete** /stages/prompt/prompts/{prompt_uuid}/ | 
[**StagesPromptPromptsList**](StagesApi.md#StagesPromptPromptsList) | **Get** /stages/prompt/prompts/ | 
[**StagesPromptPromptsPartialUpdate**](StagesApi.md#StagesPromptPromptsPartialUpdate) | **Patch** /stages/prompt/prompts/{prompt_uuid}/ | 
[**StagesPromptPromptsRetrieve**](StagesApi.md#StagesPromptPromptsRetrieve) | **Get** /stages/prompt/prompts/{prompt_uuid}/ | 
[**StagesPromptPromptsUpdate**](StagesApi.md#StagesPromptPromptsUpdate) | **Put** /stages/prompt/prompts/{prompt_uuid}/ | 
[**StagesPromptPromptsUsedByList**](StagesApi.md#StagesPromptPromptsUsedByList) | **Get** /stages/prompt/prompts/{prompt_uuid}/used_by/ | 
[**StagesPromptStagesCreate**](StagesApi.md#StagesPromptStagesCreate) | **Post** /stages/prompt/stages/ | 
[**StagesPromptStagesDestroy**](StagesApi.md#StagesPromptStagesDestroy) | **Delete** /stages/prompt/stages/{stage_uuid}/ | 
[**StagesPromptStagesList**](StagesApi.md#StagesPromptStagesList) | **Get** /stages/prompt/stages/ | 
[**StagesPromptStagesPartialUpdate**](StagesApi.md#StagesPromptStagesPartialUpdate) | **Patch** /stages/prompt/stages/{stage_uuid}/ | 
[**StagesPromptStagesRetrieve**](StagesApi.md#StagesPromptStagesRetrieve) | **Get** /stages/prompt/stages/{stage_uuid}/ | 
[**StagesPromptStagesUpdate**](StagesApi.md#StagesPromptStagesUpdate) | **Put** /stages/prompt/stages/{stage_uuid}/ | 
[**StagesPromptStagesUsedByList**](StagesApi.md#StagesPromptStagesUsedByList) | **Get** /stages/prompt/stages/{stage_uuid}/used_by/ | 
[**StagesUserDeleteCreate**](StagesApi.md#StagesUserDeleteCreate) | **Post** /stages/user_delete/ | 
[**StagesUserDeleteDestroy**](StagesApi.md#StagesUserDeleteDestroy) | **Delete** /stages/user_delete/{stage_uuid}/ | 
[**StagesUserDeleteList**](StagesApi.md#StagesUserDeleteList) | **Get** /stages/user_delete/ | 
[**StagesUserDeletePartialUpdate**](StagesApi.md#StagesUserDeletePartialUpdate) | **Patch** /stages/user_delete/{stage_uuid}/ | 
[**StagesUserDeleteRetrieve**](StagesApi.md#StagesUserDeleteRetrieve) | **Get** /stages/user_delete/{stage_uuid}/ | 
[**StagesUserDeleteUpdate**](StagesApi.md#StagesUserDeleteUpdate) | **Put** /stages/user_delete/{stage_uuid}/ | 
[**StagesUserDeleteUsedByList**](StagesApi.md#StagesUserDeleteUsedByList) | **Get** /stages/user_delete/{stage_uuid}/used_by/ | 
[**StagesUserLoginCreate**](StagesApi.md#StagesUserLoginCreate) | **Post** /stages/user_login/ | 
[**StagesUserLoginDestroy**](StagesApi.md#StagesUserLoginDestroy) | **Delete** /stages/user_login/{stage_uuid}/ | 
[**StagesUserLoginList**](StagesApi.md#StagesUserLoginList) | **Get** /stages/user_login/ | 
[**StagesUserLoginPartialUpdate**](StagesApi.md#StagesUserLoginPartialUpdate) | **Patch** /stages/user_login/{stage_uuid}/ | 
[**StagesUserLoginRetrieve**](StagesApi.md#StagesUserLoginRetrieve) | **Get** /stages/user_login/{stage_uuid}/ | 
[**StagesUserLoginUpdate**](StagesApi.md#StagesUserLoginUpdate) | **Put** /stages/user_login/{stage_uuid}/ | 
[**StagesUserLoginUsedByList**](StagesApi.md#StagesUserLoginUsedByList) | **Get** /stages/user_login/{stage_uuid}/used_by/ | 
[**StagesUserLogoutCreate**](StagesApi.md#StagesUserLogoutCreate) | **Post** /stages/user_logout/ | 
[**StagesUserLogoutDestroy**](StagesApi.md#StagesUserLogoutDestroy) | **Delete** /stages/user_logout/{stage_uuid}/ | 
[**StagesUserLogoutList**](StagesApi.md#StagesUserLogoutList) | **Get** /stages/user_logout/ | 
[**StagesUserLogoutPartialUpdate**](StagesApi.md#StagesUserLogoutPartialUpdate) | **Patch** /stages/user_logout/{stage_uuid}/ | 
[**StagesUserLogoutRetrieve**](StagesApi.md#StagesUserLogoutRetrieve) | **Get** /stages/user_logout/{stage_uuid}/ | 
[**StagesUserLogoutUpdate**](StagesApi.md#StagesUserLogoutUpdate) | **Put** /stages/user_logout/{stage_uuid}/ | 
[**StagesUserLogoutUsedByList**](StagesApi.md#StagesUserLogoutUsedByList) | **Get** /stages/user_logout/{stage_uuid}/used_by/ | 
[**StagesUserWriteCreate**](StagesApi.md#StagesUserWriteCreate) | **Post** /stages/user_write/ | 
[**StagesUserWriteDestroy**](StagesApi.md#StagesUserWriteDestroy) | **Delete** /stages/user_write/{stage_uuid}/ | 
[**StagesUserWriteList**](StagesApi.md#StagesUserWriteList) | **Get** /stages/user_write/ | 
[**StagesUserWritePartialUpdate**](StagesApi.md#StagesUserWritePartialUpdate) | **Patch** /stages/user_write/{stage_uuid}/ | 
[**StagesUserWriteRetrieve**](StagesApi.md#StagesUserWriteRetrieve) | **Get** /stages/user_write/{stage_uuid}/ | 
[**StagesUserWriteUpdate**](StagesApi.md#StagesUserWriteUpdate) | **Put** /stages/user_write/{stage_uuid}/ | 
[**StagesUserWriteUsedByList**](StagesApi.md#StagesUserWriteUsedByList) | **Get** /stages/user_write/{stage_uuid}/used_by/ | 



## StagesAllDestroy

> StagesAllDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAllDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAllDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAllDestroyRequest struct via the builder pattern


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


## StagesAllList

> PaginatedStageList StagesAllList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAllList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAllList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAllList`: PaginatedStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAllList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesAllListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedStageList**](PaginatedStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAllRetrieve

> Stage StagesAllRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAllRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAllRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAllRetrieve`: Stage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAllRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAllRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Stage**](Stage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAllTypesList

> []TypeCreate StagesAllTypesList(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAllTypesList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAllTypesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAllTypesList`: []TypeCreate
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAllTypesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAllTypesListRequest struct via the builder pattern


### Return type

[**[]TypeCreate**](TypeCreate.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAllUsedByList

> []UsedBy StagesAllUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAllUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAllUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAllUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAllUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAllUsedByListRequest struct via the builder pattern


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


## StagesAllUserSettingsList

> []UserSetting StagesAllUserSettingsList(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAllUserSettingsList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAllUserSettingsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAllUserSettingsList`: []UserSetting
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAllUserSettingsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAllUserSettingsListRequest struct via the builder pattern


### Return type

[**[]UserSetting**](UserSetting.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorDuoCreate

> AuthenticatorDuoStage StagesAuthenticatorDuoCreate(ctx).AuthenticatorDuoStageRequest(authenticatorDuoStageRequest).Execute()





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
    authenticatorDuoStageRequest := *openapiclient.NewAuthenticatorDuoStageRequest("Name_example", "ClientId_example", "ClientSecret_example", "ApiHostname_example") // AuthenticatorDuoStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorDuoCreate(context.Background()).AuthenticatorDuoStageRequest(authenticatorDuoStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorDuoCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorDuoCreate`: AuthenticatorDuoStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorDuoCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorDuoCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticatorDuoStageRequest** | [**AuthenticatorDuoStageRequest**](AuthenticatorDuoStageRequest.md) |  | 

### Return type

[**AuthenticatorDuoStage**](AuthenticatorDuoStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorDuoDestroy

> StagesAuthenticatorDuoDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Duo Authenticator Setup Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorDuoDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorDuoDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Duo Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorDuoDestroyRequest struct via the builder pattern


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


## StagesAuthenticatorDuoEnrollmentStatusCreate

> DuoDeviceEnrollmentStatus StagesAuthenticatorDuoEnrollmentStatusCreate(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Duo Authenticator Setup Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorDuoEnrollmentStatusCreate(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorDuoEnrollmentStatusCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorDuoEnrollmentStatusCreate`: DuoDeviceEnrollmentStatus
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorDuoEnrollmentStatusCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Duo Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorDuoEnrollmentStatusCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DuoDeviceEnrollmentStatus**](DuoDeviceEnrollmentStatus.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorDuoImportDeviceManualCreate

> StagesAuthenticatorDuoImportDeviceManualCreate(ctx, stageUuid).AuthenticatorDuoStageManualDeviceImportRequest(authenticatorDuoStageManualDeviceImportRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Duo Authenticator Setup Stage.
    authenticatorDuoStageManualDeviceImportRequest := *openapiclient.NewAuthenticatorDuoStageManualDeviceImportRequest("DuoUserId_example", "Username_example") // AuthenticatorDuoStageManualDeviceImportRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorDuoImportDeviceManualCreate(context.Background(), stageUuid).AuthenticatorDuoStageManualDeviceImportRequest(authenticatorDuoStageManualDeviceImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorDuoImportDeviceManualCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Duo Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorDuoImportDeviceManualCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticatorDuoStageManualDeviceImportRequest** | [**AuthenticatorDuoStageManualDeviceImportRequest**](AuthenticatorDuoStageManualDeviceImportRequest.md) |  | 

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


## StagesAuthenticatorDuoImportDevicesAutomaticCreate

> AuthenticatorDuoStageDeviceImportResponse StagesAuthenticatorDuoImportDevicesAutomaticCreate(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Duo Authenticator Setup Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorDuoImportDevicesAutomaticCreate(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorDuoImportDevicesAutomaticCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorDuoImportDevicesAutomaticCreate`: AuthenticatorDuoStageDeviceImportResponse
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorDuoImportDevicesAutomaticCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Duo Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorDuoImportDevicesAutomaticCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticatorDuoStageDeviceImportResponse**](AuthenticatorDuoStageDeviceImportResponse.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorDuoList

> PaginatedAuthenticatorDuoStageList StagesAuthenticatorDuoList(ctx).ApiHostname(apiHostname).ClientId(clientId).ConfigureFlow(configureFlow).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    apiHostname := "apiHostname_example" // string |  (optional)
    clientId := "clientId_example" // string |  (optional)
    configureFlow := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorDuoList(context.Background()).ApiHostname(apiHostname).ClientId(clientId).ConfigureFlow(configureFlow).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorDuoList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorDuoList`: PaginatedAuthenticatorDuoStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorDuoList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorDuoListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiHostname** | **string** |  | 
 **clientId** | **string** |  | 
 **configureFlow** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedAuthenticatorDuoStageList**](PaginatedAuthenticatorDuoStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorDuoPartialUpdate

> AuthenticatorDuoStage StagesAuthenticatorDuoPartialUpdate(ctx, stageUuid).PatchedAuthenticatorDuoStageRequest(patchedAuthenticatorDuoStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Duo Authenticator Setup Stage.
    patchedAuthenticatorDuoStageRequest := *openapiclient.NewPatchedAuthenticatorDuoStageRequest() // PatchedAuthenticatorDuoStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorDuoPartialUpdate(context.Background(), stageUuid).PatchedAuthenticatorDuoStageRequest(patchedAuthenticatorDuoStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorDuoPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorDuoPartialUpdate`: AuthenticatorDuoStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorDuoPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Duo Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorDuoPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAuthenticatorDuoStageRequest** | [**PatchedAuthenticatorDuoStageRequest**](PatchedAuthenticatorDuoStageRequest.md) |  | 

### Return type

[**AuthenticatorDuoStage**](AuthenticatorDuoStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorDuoRetrieve

> AuthenticatorDuoStage StagesAuthenticatorDuoRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Duo Authenticator Setup Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorDuoRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorDuoRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorDuoRetrieve`: AuthenticatorDuoStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorDuoRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Duo Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorDuoRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticatorDuoStage**](AuthenticatorDuoStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorDuoUpdate

> AuthenticatorDuoStage StagesAuthenticatorDuoUpdate(ctx, stageUuid).AuthenticatorDuoStageRequest(authenticatorDuoStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Duo Authenticator Setup Stage.
    authenticatorDuoStageRequest := *openapiclient.NewAuthenticatorDuoStageRequest("Name_example", "ClientId_example", "ClientSecret_example", "ApiHostname_example") // AuthenticatorDuoStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorDuoUpdate(context.Background(), stageUuid).AuthenticatorDuoStageRequest(authenticatorDuoStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorDuoUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorDuoUpdate`: AuthenticatorDuoStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorDuoUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Duo Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorDuoUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticatorDuoStageRequest** | [**AuthenticatorDuoStageRequest**](AuthenticatorDuoStageRequest.md) |  | 

### Return type

[**AuthenticatorDuoStage**](AuthenticatorDuoStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorDuoUsedByList

> []UsedBy StagesAuthenticatorDuoUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Duo Authenticator Setup Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorDuoUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorDuoUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorDuoUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorDuoUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Duo Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorDuoUsedByListRequest struct via the builder pattern


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


## StagesAuthenticatorSmsCreate

> AuthenticatorSMSStage StagesAuthenticatorSmsCreate(ctx).AuthenticatorSMSStageRequest(authenticatorSMSStageRequest).Execute()





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
    authenticatorSMSStageRequest := *openapiclient.NewAuthenticatorSMSStageRequest("Name_example", openapiclient.ProviderEnum("twilio"), "FromNumber_example", "AccountSid_example", "Auth_example") // AuthenticatorSMSStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorSmsCreate(context.Background()).AuthenticatorSMSStageRequest(authenticatorSMSStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorSmsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorSmsCreate`: AuthenticatorSMSStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorSmsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorSmsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticatorSMSStageRequest** | [**AuthenticatorSMSStageRequest**](AuthenticatorSMSStageRequest.md) |  | 

### Return type

[**AuthenticatorSMSStage**](AuthenticatorSMSStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorSmsDestroy

> StagesAuthenticatorSmsDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this SMS Authenticator Setup Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorSmsDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorSmsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this SMS Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorSmsDestroyRequest struct via the builder pattern


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


## StagesAuthenticatorSmsList

> PaginatedAuthenticatorSMSStageList StagesAuthenticatorSmsList(ctx).AccountSid(accountSid).Auth(auth).AuthPassword(authPassword).AuthType(authType).ConfigureFlow(configureFlow).FromNumber(fromNumber).Mapping(mapping).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Search(search).StageUuid(stageUuid).VerifyOnly(verifyOnly).Execute()





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
    accountSid := "accountSid_example" // string |  (optional)
    auth := "auth_example" // string |  (optional)
    authPassword := "authPassword_example" // string |  (optional)
    authType := "authType_example" // string | * `basic` - Basic * `bearer` - Bearer  * `basic` - Basic * `bearer` - Bearer (optional)
    configureFlow := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    fromNumber := "fromNumber_example" // string |  (optional)
    mapping := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    provider := "provider_example" // string | * `twilio` - Twilio * `generic` - Generic  * `twilio` - Twilio * `generic` - Generic (optional)
    search := "search_example" // string | A search term. (optional)
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    verifyOnly := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorSmsList(context.Background()).AccountSid(accountSid).Auth(auth).AuthPassword(authPassword).AuthType(authType).ConfigureFlow(configureFlow).FromNumber(fromNumber).Mapping(mapping).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Provider(provider).Search(search).StageUuid(stageUuid).VerifyOnly(verifyOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorSmsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorSmsList`: PaginatedAuthenticatorSMSStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorSmsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorSmsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountSid** | **string** |  | 
 **auth** | **string** |  | 
 **authPassword** | **string** |  | 
 **authType** | **string** | * &#x60;basic&#x60; - Basic * &#x60;bearer&#x60; - Bearer  * &#x60;basic&#x60; - Basic * &#x60;bearer&#x60; - Bearer | 
 **configureFlow** | **string** |  | 
 **fromNumber** | **string** |  | 
 **mapping** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **provider** | **string** | * &#x60;twilio&#x60; - Twilio * &#x60;generic&#x60; - Generic  * &#x60;twilio&#x60; - Twilio * &#x60;generic&#x60; - Generic | 
 **search** | **string** | A search term. | 
 **stageUuid** | **string** |  | 
 **verifyOnly** | **bool** |  | 

### Return type

[**PaginatedAuthenticatorSMSStageList**](PaginatedAuthenticatorSMSStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorSmsPartialUpdate

> AuthenticatorSMSStage StagesAuthenticatorSmsPartialUpdate(ctx, stageUuid).PatchedAuthenticatorSMSStageRequest(patchedAuthenticatorSMSStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this SMS Authenticator Setup Stage.
    patchedAuthenticatorSMSStageRequest := *openapiclient.NewPatchedAuthenticatorSMSStageRequest() // PatchedAuthenticatorSMSStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorSmsPartialUpdate(context.Background(), stageUuid).PatchedAuthenticatorSMSStageRequest(patchedAuthenticatorSMSStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorSmsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorSmsPartialUpdate`: AuthenticatorSMSStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorSmsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this SMS Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorSmsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAuthenticatorSMSStageRequest** | [**PatchedAuthenticatorSMSStageRequest**](PatchedAuthenticatorSMSStageRequest.md) |  | 

### Return type

[**AuthenticatorSMSStage**](AuthenticatorSMSStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorSmsRetrieve

> AuthenticatorSMSStage StagesAuthenticatorSmsRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this SMS Authenticator Setup Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorSmsRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorSmsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorSmsRetrieve`: AuthenticatorSMSStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorSmsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this SMS Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorSmsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticatorSMSStage**](AuthenticatorSMSStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorSmsUpdate

> AuthenticatorSMSStage StagesAuthenticatorSmsUpdate(ctx, stageUuid).AuthenticatorSMSStageRequest(authenticatorSMSStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this SMS Authenticator Setup Stage.
    authenticatorSMSStageRequest := *openapiclient.NewAuthenticatorSMSStageRequest("Name_example", openapiclient.ProviderEnum("twilio"), "FromNumber_example", "AccountSid_example", "Auth_example") // AuthenticatorSMSStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorSmsUpdate(context.Background(), stageUuid).AuthenticatorSMSStageRequest(authenticatorSMSStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorSmsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorSmsUpdate`: AuthenticatorSMSStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorSmsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this SMS Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorSmsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticatorSMSStageRequest** | [**AuthenticatorSMSStageRequest**](AuthenticatorSMSStageRequest.md) |  | 

### Return type

[**AuthenticatorSMSStage**](AuthenticatorSMSStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorSmsUsedByList

> []UsedBy StagesAuthenticatorSmsUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this SMS Authenticator Setup Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorSmsUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorSmsUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorSmsUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorSmsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this SMS Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorSmsUsedByListRequest struct via the builder pattern


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


## StagesAuthenticatorStaticCreate

> AuthenticatorStaticStage StagesAuthenticatorStaticCreate(ctx).AuthenticatorStaticStageRequest(authenticatorStaticStageRequest).Execute()





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
    authenticatorStaticStageRequest := *openapiclient.NewAuthenticatorStaticStageRequest("Name_example") // AuthenticatorStaticStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorStaticCreate(context.Background()).AuthenticatorStaticStageRequest(authenticatorStaticStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorStaticCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorStaticCreate`: AuthenticatorStaticStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorStaticCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorStaticCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticatorStaticStageRequest** | [**AuthenticatorStaticStageRequest**](AuthenticatorStaticStageRequest.md) |  | 

### Return type

[**AuthenticatorStaticStage**](AuthenticatorStaticStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorStaticDestroy

> StagesAuthenticatorStaticDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Static Authenticator Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorStaticDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorStaticDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Static Authenticator Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorStaticDestroyRequest struct via the builder pattern


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


## StagesAuthenticatorStaticList

> PaginatedAuthenticatorStaticStageList StagesAuthenticatorStaticList(ctx).ConfigureFlow(configureFlow).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).TokenCount(tokenCount).Execute()





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
    configureFlow := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    tokenCount := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorStaticList(context.Background()).ConfigureFlow(configureFlow).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).TokenCount(tokenCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorStaticList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorStaticList`: PaginatedAuthenticatorStaticStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorStaticList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorStaticListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configureFlow** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **stageUuid** | **string** |  | 
 **tokenCount** | **int32** |  | 

### Return type

[**PaginatedAuthenticatorStaticStageList**](PaginatedAuthenticatorStaticStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorStaticPartialUpdate

> AuthenticatorStaticStage StagesAuthenticatorStaticPartialUpdate(ctx, stageUuid).PatchedAuthenticatorStaticStageRequest(patchedAuthenticatorStaticStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Static Authenticator Stage.
    patchedAuthenticatorStaticStageRequest := *openapiclient.NewPatchedAuthenticatorStaticStageRequest() // PatchedAuthenticatorStaticStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorStaticPartialUpdate(context.Background(), stageUuid).PatchedAuthenticatorStaticStageRequest(patchedAuthenticatorStaticStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorStaticPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorStaticPartialUpdate`: AuthenticatorStaticStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorStaticPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Static Authenticator Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorStaticPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAuthenticatorStaticStageRequest** | [**PatchedAuthenticatorStaticStageRequest**](PatchedAuthenticatorStaticStageRequest.md) |  | 

### Return type

[**AuthenticatorStaticStage**](AuthenticatorStaticStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorStaticRetrieve

> AuthenticatorStaticStage StagesAuthenticatorStaticRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Static Authenticator Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorStaticRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorStaticRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorStaticRetrieve`: AuthenticatorStaticStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorStaticRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Static Authenticator Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorStaticRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticatorStaticStage**](AuthenticatorStaticStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorStaticUpdate

> AuthenticatorStaticStage StagesAuthenticatorStaticUpdate(ctx, stageUuid).AuthenticatorStaticStageRequest(authenticatorStaticStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Static Authenticator Stage.
    authenticatorStaticStageRequest := *openapiclient.NewAuthenticatorStaticStageRequest("Name_example") // AuthenticatorStaticStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorStaticUpdate(context.Background(), stageUuid).AuthenticatorStaticStageRequest(authenticatorStaticStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorStaticUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorStaticUpdate`: AuthenticatorStaticStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorStaticUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Static Authenticator Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorStaticUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticatorStaticStageRequest** | [**AuthenticatorStaticStageRequest**](AuthenticatorStaticStageRequest.md) |  | 

### Return type

[**AuthenticatorStaticStage**](AuthenticatorStaticStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorStaticUsedByList

> []UsedBy StagesAuthenticatorStaticUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Static Authenticator Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorStaticUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorStaticUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorStaticUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorStaticUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Static Authenticator Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorStaticUsedByListRequest struct via the builder pattern


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


## StagesAuthenticatorTotpCreate

> AuthenticatorTOTPStage StagesAuthenticatorTotpCreate(ctx).AuthenticatorTOTPStageRequest(authenticatorTOTPStageRequest).Execute()





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
    authenticatorTOTPStageRequest := *openapiclient.NewAuthenticatorTOTPStageRequest("Name_example", "TODO") // AuthenticatorTOTPStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorTotpCreate(context.Background()).AuthenticatorTOTPStageRequest(authenticatorTOTPStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorTotpCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorTotpCreate`: AuthenticatorTOTPStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorTotpCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorTotpCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticatorTOTPStageRequest** | [**AuthenticatorTOTPStageRequest**](AuthenticatorTOTPStageRequest.md) |  | 

### Return type

[**AuthenticatorTOTPStage**](AuthenticatorTOTPStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorTotpDestroy

> StagesAuthenticatorTotpDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this TOTP Authenticator Setup Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorTotpDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorTotpDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this TOTP Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorTotpDestroyRequest struct via the builder pattern


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


## StagesAuthenticatorTotpList

> PaginatedAuthenticatorTOTPStageList StagesAuthenticatorTotpList(ctx).ConfigureFlow(configureFlow).Digits(digits).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).Execute()





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
    configureFlow := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    digits := int32(56) // int32 | * `6` - 6 digits, widely compatible * `8` - 8 digits, not compatible with apps like Google Authenticator  * `6` - 6 digits, widely compatible * `8` - 8 digits, not compatible with apps like Google Authenticator (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorTotpList(context.Background()).ConfigureFlow(configureFlow).Digits(digits).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorTotpList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorTotpList`: PaginatedAuthenticatorTOTPStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorTotpList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorTotpListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configureFlow** | **string** |  | 
 **digits** | **int32** | * &#x60;6&#x60; - 6 digits, widely compatible * &#x60;8&#x60; - 8 digits, not compatible with apps like Google Authenticator  * &#x60;6&#x60; - 6 digits, widely compatible * &#x60;8&#x60; - 8 digits, not compatible with apps like Google Authenticator | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **stageUuid** | **string** |  | 

### Return type

[**PaginatedAuthenticatorTOTPStageList**](PaginatedAuthenticatorTOTPStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorTotpPartialUpdate

> AuthenticatorTOTPStage StagesAuthenticatorTotpPartialUpdate(ctx, stageUuid).PatchedAuthenticatorTOTPStageRequest(patchedAuthenticatorTOTPStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this TOTP Authenticator Setup Stage.
    patchedAuthenticatorTOTPStageRequest := *openapiclient.NewPatchedAuthenticatorTOTPStageRequest() // PatchedAuthenticatorTOTPStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorTotpPartialUpdate(context.Background(), stageUuid).PatchedAuthenticatorTOTPStageRequest(patchedAuthenticatorTOTPStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorTotpPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorTotpPartialUpdate`: AuthenticatorTOTPStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorTotpPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this TOTP Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorTotpPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAuthenticatorTOTPStageRequest** | [**PatchedAuthenticatorTOTPStageRequest**](PatchedAuthenticatorTOTPStageRequest.md) |  | 

### Return type

[**AuthenticatorTOTPStage**](AuthenticatorTOTPStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorTotpRetrieve

> AuthenticatorTOTPStage StagesAuthenticatorTotpRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this TOTP Authenticator Setup Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorTotpRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorTotpRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorTotpRetrieve`: AuthenticatorTOTPStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorTotpRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this TOTP Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorTotpRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticatorTOTPStage**](AuthenticatorTOTPStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorTotpUpdate

> AuthenticatorTOTPStage StagesAuthenticatorTotpUpdate(ctx, stageUuid).AuthenticatorTOTPStageRequest(authenticatorTOTPStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this TOTP Authenticator Setup Stage.
    authenticatorTOTPStageRequest := *openapiclient.NewAuthenticatorTOTPStageRequest("Name_example", "TODO") // AuthenticatorTOTPStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorTotpUpdate(context.Background(), stageUuid).AuthenticatorTOTPStageRequest(authenticatorTOTPStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorTotpUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorTotpUpdate`: AuthenticatorTOTPStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorTotpUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this TOTP Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorTotpUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticatorTOTPStageRequest** | [**AuthenticatorTOTPStageRequest**](AuthenticatorTOTPStageRequest.md) |  | 

### Return type

[**AuthenticatorTOTPStage**](AuthenticatorTOTPStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorTotpUsedByList

> []UsedBy StagesAuthenticatorTotpUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this TOTP Authenticator Setup Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorTotpUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorTotpUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorTotpUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorTotpUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this TOTP Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorTotpUsedByListRequest struct via the builder pattern


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


## StagesAuthenticatorValidateCreate

> AuthenticatorValidateStage StagesAuthenticatorValidateCreate(ctx).AuthenticatorValidateStageRequest(authenticatorValidateStageRequest).Execute()





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
    authenticatorValidateStageRequest := *openapiclient.NewAuthenticatorValidateStageRequest("Name_example") // AuthenticatorValidateStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorValidateCreate(context.Background()).AuthenticatorValidateStageRequest(authenticatorValidateStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorValidateCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorValidateCreate`: AuthenticatorValidateStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorValidateCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorValidateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticatorValidateStageRequest** | [**AuthenticatorValidateStageRequest**](AuthenticatorValidateStageRequest.md) |  | 

### Return type

[**AuthenticatorValidateStage**](AuthenticatorValidateStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorValidateDestroy

> StagesAuthenticatorValidateDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Authenticator Validation Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorValidateDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorValidateDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Authenticator Validation Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorValidateDestroyRequest struct via the builder pattern


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


## StagesAuthenticatorValidateList

> PaginatedAuthenticatorValidateStageList StagesAuthenticatorValidateList(ctx).ConfigurationStages(configurationStages).Name(name).NotConfiguredAction(notConfiguredAction).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    configurationStages := []string{"Inner_example"} // []string |  (optional)
    name := "name_example" // string |  (optional)
    notConfiguredAction := "notConfiguredAction_example" // string | * `skip` - Skip * `deny` - Deny * `configure` - Configure  * `skip` - Skip * `deny` - Deny * `configure` - Configure (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorValidateList(context.Background()).ConfigurationStages(configurationStages).Name(name).NotConfiguredAction(notConfiguredAction).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorValidateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorValidateList`: PaginatedAuthenticatorValidateStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorValidateList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorValidateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationStages** | **[]string** |  | 
 **name** | **string** |  | 
 **notConfiguredAction** | **string** | * &#x60;skip&#x60; - Skip * &#x60;deny&#x60; - Deny * &#x60;configure&#x60; - Configure  * &#x60;skip&#x60; - Skip * &#x60;deny&#x60; - Deny * &#x60;configure&#x60; - Configure | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedAuthenticatorValidateStageList**](PaginatedAuthenticatorValidateStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorValidatePartialUpdate

> AuthenticatorValidateStage StagesAuthenticatorValidatePartialUpdate(ctx, stageUuid).PatchedAuthenticatorValidateStageRequest(patchedAuthenticatorValidateStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Authenticator Validation Stage.
    patchedAuthenticatorValidateStageRequest := *openapiclient.NewPatchedAuthenticatorValidateStageRequest() // PatchedAuthenticatorValidateStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorValidatePartialUpdate(context.Background(), stageUuid).PatchedAuthenticatorValidateStageRequest(patchedAuthenticatorValidateStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorValidatePartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorValidatePartialUpdate`: AuthenticatorValidateStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorValidatePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Authenticator Validation Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorValidatePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAuthenticatorValidateStageRequest** | [**PatchedAuthenticatorValidateStageRequest**](PatchedAuthenticatorValidateStageRequest.md) |  | 

### Return type

[**AuthenticatorValidateStage**](AuthenticatorValidateStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorValidateRetrieve

> AuthenticatorValidateStage StagesAuthenticatorValidateRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Authenticator Validation Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorValidateRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorValidateRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorValidateRetrieve`: AuthenticatorValidateStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorValidateRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Authenticator Validation Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorValidateRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticatorValidateStage**](AuthenticatorValidateStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorValidateUpdate

> AuthenticatorValidateStage StagesAuthenticatorValidateUpdate(ctx, stageUuid).AuthenticatorValidateStageRequest(authenticatorValidateStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Authenticator Validation Stage.
    authenticatorValidateStageRequest := *openapiclient.NewAuthenticatorValidateStageRequest("Name_example") // AuthenticatorValidateStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorValidateUpdate(context.Background(), stageUuid).AuthenticatorValidateStageRequest(authenticatorValidateStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorValidateUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorValidateUpdate`: AuthenticatorValidateStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorValidateUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Authenticator Validation Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorValidateUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticatorValidateStageRequest** | [**AuthenticatorValidateStageRequest**](AuthenticatorValidateStageRequest.md) |  | 

### Return type

[**AuthenticatorValidateStage**](AuthenticatorValidateStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorValidateUsedByList

> []UsedBy StagesAuthenticatorValidateUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Authenticator Validation Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorValidateUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorValidateUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorValidateUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorValidateUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Authenticator Validation Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorValidateUsedByListRequest struct via the builder pattern


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


## StagesAuthenticatorWebauthnCreate

> AuthenticateWebAuthnStage StagesAuthenticatorWebauthnCreate(ctx).AuthenticateWebAuthnStageRequest(authenticateWebAuthnStageRequest).Execute()





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
    authenticateWebAuthnStageRequest := *openapiclient.NewAuthenticateWebAuthnStageRequest("Name_example") // AuthenticateWebAuthnStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorWebauthnCreate(context.Background()).AuthenticateWebAuthnStageRequest(authenticateWebAuthnStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorWebauthnCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorWebauthnCreate`: AuthenticateWebAuthnStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorWebauthnCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorWebauthnCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticateWebAuthnStageRequest** | [**AuthenticateWebAuthnStageRequest**](AuthenticateWebAuthnStageRequest.md) |  | 

### Return type

[**AuthenticateWebAuthnStage**](AuthenticateWebAuthnStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorWebauthnDestroy

> StagesAuthenticatorWebauthnDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this WebAuthn Authenticator Setup Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorWebauthnDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorWebauthnDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this WebAuthn Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorWebauthnDestroyRequest struct via the builder pattern


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


## StagesAuthenticatorWebauthnList

> PaginatedAuthenticateWebAuthnStageList StagesAuthenticatorWebauthnList(ctx).AuthenticatorAttachment(authenticatorAttachment).ConfigureFlow(configureFlow).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).ResidentKeyRequirement(residentKeyRequirement).Search(search).StageUuid(stageUuid).UserVerification(userVerification).Execute()





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
    authenticatorAttachment := "authenticatorAttachment_example" // string | * `platform` - Platform * `cross-platform` - Cross Platform  * `platform` - Platform * `cross-platform` - Cross Platform (optional)
    configureFlow := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    residentKeyRequirement := "residentKeyRequirement_example" // string | * `discouraged` - Discouraged * `preferred` - Preferred * `required` - Required  * `discouraged` - Discouraged * `preferred` - Preferred * `required` - Required (optional)
    search := "search_example" // string | A search term. (optional)
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    userVerification := "userVerification_example" // string | * `required` - Required * `preferred` - Preferred * `discouraged` - Discouraged  * `required` - Required * `preferred` - Preferred * `discouraged` - Discouraged (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorWebauthnList(context.Background()).AuthenticatorAttachment(authenticatorAttachment).ConfigureFlow(configureFlow).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).ResidentKeyRequirement(residentKeyRequirement).Search(search).StageUuid(stageUuid).UserVerification(userVerification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorWebauthnList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorWebauthnList`: PaginatedAuthenticateWebAuthnStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorWebauthnList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorWebauthnListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticatorAttachment** | **string** | * &#x60;platform&#x60; - Platform * &#x60;cross-platform&#x60; - Cross Platform  * &#x60;platform&#x60; - Platform * &#x60;cross-platform&#x60; - Cross Platform | 
 **configureFlow** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **residentKeyRequirement** | **string** | * &#x60;discouraged&#x60; - Discouraged * &#x60;preferred&#x60; - Preferred * &#x60;required&#x60; - Required  * &#x60;discouraged&#x60; - Discouraged * &#x60;preferred&#x60; - Preferred * &#x60;required&#x60; - Required | 
 **search** | **string** | A search term. | 
 **stageUuid** | **string** |  | 
 **userVerification** | **string** | * &#x60;required&#x60; - Required * &#x60;preferred&#x60; - Preferred * &#x60;discouraged&#x60; - Discouraged  * &#x60;required&#x60; - Required * &#x60;preferred&#x60; - Preferred * &#x60;discouraged&#x60; - Discouraged | 

### Return type

[**PaginatedAuthenticateWebAuthnStageList**](PaginatedAuthenticateWebAuthnStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorWebauthnPartialUpdate

> AuthenticateWebAuthnStage StagesAuthenticatorWebauthnPartialUpdate(ctx, stageUuid).PatchedAuthenticateWebAuthnStageRequest(patchedAuthenticateWebAuthnStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this WebAuthn Authenticator Setup Stage.
    patchedAuthenticateWebAuthnStageRequest := *openapiclient.NewPatchedAuthenticateWebAuthnStageRequest() // PatchedAuthenticateWebAuthnStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorWebauthnPartialUpdate(context.Background(), stageUuid).PatchedAuthenticateWebAuthnStageRequest(patchedAuthenticateWebAuthnStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorWebauthnPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorWebauthnPartialUpdate`: AuthenticateWebAuthnStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorWebauthnPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this WebAuthn Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorWebauthnPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAuthenticateWebAuthnStageRequest** | [**PatchedAuthenticateWebAuthnStageRequest**](PatchedAuthenticateWebAuthnStageRequest.md) |  | 

### Return type

[**AuthenticateWebAuthnStage**](AuthenticateWebAuthnStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorWebauthnRetrieve

> AuthenticateWebAuthnStage StagesAuthenticatorWebauthnRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this WebAuthn Authenticator Setup Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorWebauthnRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorWebauthnRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorWebauthnRetrieve`: AuthenticateWebAuthnStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorWebauthnRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this WebAuthn Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorWebauthnRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticateWebAuthnStage**](AuthenticateWebAuthnStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorWebauthnUpdate

> AuthenticateWebAuthnStage StagesAuthenticatorWebauthnUpdate(ctx, stageUuid).AuthenticateWebAuthnStageRequest(authenticateWebAuthnStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this WebAuthn Authenticator Setup Stage.
    authenticateWebAuthnStageRequest := *openapiclient.NewAuthenticateWebAuthnStageRequest("Name_example") // AuthenticateWebAuthnStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorWebauthnUpdate(context.Background(), stageUuid).AuthenticateWebAuthnStageRequest(authenticateWebAuthnStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorWebauthnUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorWebauthnUpdate`: AuthenticateWebAuthnStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorWebauthnUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this WebAuthn Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorWebauthnUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticateWebAuthnStageRequest** | [**AuthenticateWebAuthnStageRequest**](AuthenticateWebAuthnStageRequest.md) |  | 

### Return type

[**AuthenticateWebAuthnStage**](AuthenticateWebAuthnStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesAuthenticatorWebauthnUsedByList

> []UsedBy StagesAuthenticatorWebauthnUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this WebAuthn Authenticator Setup Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesAuthenticatorWebauthnUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesAuthenticatorWebauthnUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesAuthenticatorWebauthnUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesAuthenticatorWebauthnUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this WebAuthn Authenticator Setup Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesAuthenticatorWebauthnUsedByListRequest struct via the builder pattern


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


## StagesCaptchaCreate

> CaptchaStage StagesCaptchaCreate(ctx).CaptchaStageRequest(captchaStageRequest).Execute()





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
    captchaStageRequest := *openapiclient.NewCaptchaStageRequest("Name_example", "PublicKey_example", "PrivateKey_example") // CaptchaStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesCaptchaCreate(context.Background()).CaptchaStageRequest(captchaStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesCaptchaCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesCaptchaCreate`: CaptchaStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesCaptchaCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesCaptchaCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **captchaStageRequest** | [**CaptchaStageRequest**](CaptchaStageRequest.md) |  | 

### Return type

[**CaptchaStage**](CaptchaStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesCaptchaDestroy

> StagesCaptchaDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Captcha Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesCaptchaDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesCaptchaDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Captcha Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesCaptchaDestroyRequest struct via the builder pattern


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


## StagesCaptchaList

> PaginatedCaptchaStageList StagesCaptchaList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PublicKey(publicKey).Search(search).Execute()





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
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    publicKey := "publicKey_example" // string |  (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesCaptchaList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PublicKey(publicKey).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesCaptchaList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesCaptchaList`: PaginatedCaptchaStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesCaptchaList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesCaptchaListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **publicKey** | **string** |  | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedCaptchaStageList**](PaginatedCaptchaStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesCaptchaPartialUpdate

> CaptchaStage StagesCaptchaPartialUpdate(ctx, stageUuid).PatchedCaptchaStageRequest(patchedCaptchaStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Captcha Stage.
    patchedCaptchaStageRequest := *openapiclient.NewPatchedCaptchaStageRequest() // PatchedCaptchaStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesCaptchaPartialUpdate(context.Background(), stageUuid).PatchedCaptchaStageRequest(patchedCaptchaStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesCaptchaPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesCaptchaPartialUpdate`: CaptchaStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesCaptchaPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Captcha Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesCaptchaPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedCaptchaStageRequest** | [**PatchedCaptchaStageRequest**](PatchedCaptchaStageRequest.md) |  | 

### Return type

[**CaptchaStage**](CaptchaStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesCaptchaRetrieve

> CaptchaStage StagesCaptchaRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Captcha Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesCaptchaRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesCaptchaRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesCaptchaRetrieve`: CaptchaStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesCaptchaRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Captcha Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesCaptchaRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CaptchaStage**](CaptchaStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesCaptchaUpdate

> CaptchaStage StagesCaptchaUpdate(ctx, stageUuid).CaptchaStageRequest(captchaStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Captcha Stage.
    captchaStageRequest := *openapiclient.NewCaptchaStageRequest("Name_example", "PublicKey_example", "PrivateKey_example") // CaptchaStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesCaptchaUpdate(context.Background(), stageUuid).CaptchaStageRequest(captchaStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesCaptchaUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesCaptchaUpdate`: CaptchaStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesCaptchaUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Captcha Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesCaptchaUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **captchaStageRequest** | [**CaptchaStageRequest**](CaptchaStageRequest.md) |  | 

### Return type

[**CaptchaStage**](CaptchaStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesCaptchaUsedByList

> []UsedBy StagesCaptchaUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Captcha Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesCaptchaUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesCaptchaUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesCaptchaUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesCaptchaUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Captcha Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesCaptchaUsedByListRequest struct via the builder pattern


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


## StagesConsentCreate

> ConsentStage StagesConsentCreate(ctx).ConsentStageRequest(consentStageRequest).Execute()





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
    consentStageRequest := *openapiclient.NewConsentStageRequest("Name_example") // ConsentStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesConsentCreate(context.Background()).ConsentStageRequest(consentStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesConsentCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesConsentCreate`: ConsentStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesConsentCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesConsentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consentStageRequest** | [**ConsentStageRequest**](ConsentStageRequest.md) |  | 

### Return type

[**ConsentStage**](ConsentStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesConsentDestroy

> StagesConsentDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Consent Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesConsentDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesConsentDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Consent Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesConsentDestroyRequest struct via the builder pattern


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


## StagesConsentList

> PaginatedConsentStageList StagesConsentList(ctx).ConsentExpireIn(consentExpireIn).Mode(mode).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).Execute()





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
    consentExpireIn := "consentExpireIn_example" // string |  (optional)
    mode := "mode_example" // string | * `always_require` - Always Require * `permanent` - Permanent * `expiring` - Expiring  * `always_require` - Always Require * `permanent` - Permanent * `expiring` - Expiring (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesConsentList(context.Background()).ConsentExpireIn(consentExpireIn).Mode(mode).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesConsentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesConsentList`: PaginatedConsentStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesConsentList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesConsentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consentExpireIn** | **string** |  | 
 **mode** | **string** | * &#x60;always_require&#x60; - Always Require * &#x60;permanent&#x60; - Permanent * &#x60;expiring&#x60; - Expiring  * &#x60;always_require&#x60; - Always Require * &#x60;permanent&#x60; - Permanent * &#x60;expiring&#x60; - Expiring | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **stageUuid** | **string** |  | 

### Return type

[**PaginatedConsentStageList**](PaginatedConsentStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesConsentPartialUpdate

> ConsentStage StagesConsentPartialUpdate(ctx, stageUuid).PatchedConsentStageRequest(patchedConsentStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Consent Stage.
    patchedConsentStageRequest := *openapiclient.NewPatchedConsentStageRequest() // PatchedConsentStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesConsentPartialUpdate(context.Background(), stageUuid).PatchedConsentStageRequest(patchedConsentStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesConsentPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesConsentPartialUpdate`: ConsentStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesConsentPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Consent Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesConsentPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedConsentStageRequest** | [**PatchedConsentStageRequest**](PatchedConsentStageRequest.md) |  | 

### Return type

[**ConsentStage**](ConsentStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesConsentRetrieve

> ConsentStage StagesConsentRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Consent Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesConsentRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesConsentRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesConsentRetrieve`: ConsentStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesConsentRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Consent Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesConsentRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsentStage**](ConsentStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesConsentUpdate

> ConsentStage StagesConsentUpdate(ctx, stageUuid).ConsentStageRequest(consentStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Consent Stage.
    consentStageRequest := *openapiclient.NewConsentStageRequest("Name_example") // ConsentStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesConsentUpdate(context.Background(), stageUuid).ConsentStageRequest(consentStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesConsentUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesConsentUpdate`: ConsentStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesConsentUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Consent Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesConsentUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consentStageRequest** | [**ConsentStageRequest**](ConsentStageRequest.md) |  | 

### Return type

[**ConsentStage**](ConsentStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesConsentUsedByList

> []UsedBy StagesConsentUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Consent Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesConsentUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesConsentUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesConsentUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesConsentUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Consent Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesConsentUsedByListRequest struct via the builder pattern


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


## StagesDenyCreate

> DenyStage StagesDenyCreate(ctx).DenyStageRequest(denyStageRequest).Execute()





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
    denyStageRequest := *openapiclient.NewDenyStageRequest("Name_example") // DenyStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesDenyCreate(context.Background()).DenyStageRequest(denyStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesDenyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesDenyCreate`: DenyStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesDenyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesDenyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **denyStageRequest** | [**DenyStageRequest**](DenyStageRequest.md) |  | 

### Return type

[**DenyStage**](DenyStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesDenyDestroy

> StagesDenyDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Deny Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesDenyDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesDenyDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Deny Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesDenyDestroyRequest struct via the builder pattern


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


## StagesDenyList

> PaginatedDenyStageList StagesDenyList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).Execute()





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
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesDenyList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesDenyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesDenyList`: PaginatedDenyStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesDenyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesDenyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **stageUuid** | **string** |  | 

### Return type

[**PaginatedDenyStageList**](PaginatedDenyStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesDenyPartialUpdate

> DenyStage StagesDenyPartialUpdate(ctx, stageUuid).PatchedDenyStageRequest(patchedDenyStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Deny Stage.
    patchedDenyStageRequest := *openapiclient.NewPatchedDenyStageRequest() // PatchedDenyStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesDenyPartialUpdate(context.Background(), stageUuid).PatchedDenyStageRequest(patchedDenyStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesDenyPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesDenyPartialUpdate`: DenyStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesDenyPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Deny Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesDenyPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDenyStageRequest** | [**PatchedDenyStageRequest**](PatchedDenyStageRequest.md) |  | 

### Return type

[**DenyStage**](DenyStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesDenyRetrieve

> DenyStage StagesDenyRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Deny Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesDenyRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesDenyRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesDenyRetrieve`: DenyStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesDenyRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Deny Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesDenyRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DenyStage**](DenyStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesDenyUpdate

> DenyStage StagesDenyUpdate(ctx, stageUuid).DenyStageRequest(denyStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Deny Stage.
    denyStageRequest := *openapiclient.NewDenyStageRequest("Name_example") // DenyStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesDenyUpdate(context.Background(), stageUuid).DenyStageRequest(denyStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesDenyUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesDenyUpdate`: DenyStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesDenyUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Deny Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesDenyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **denyStageRequest** | [**DenyStageRequest**](DenyStageRequest.md) |  | 

### Return type

[**DenyStage**](DenyStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesDenyUsedByList

> []UsedBy StagesDenyUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Deny Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesDenyUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesDenyUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesDenyUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesDenyUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Deny Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesDenyUsedByListRequest struct via the builder pattern


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


## StagesDummyCreate

> DummyStage StagesDummyCreate(ctx).DummyStageRequest(dummyStageRequest).Execute()





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
    dummyStageRequest := *openapiclient.NewDummyStageRequest("Name_example") // DummyStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesDummyCreate(context.Background()).DummyStageRequest(dummyStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesDummyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesDummyCreate`: DummyStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesDummyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesDummyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dummyStageRequest** | [**DummyStageRequest**](DummyStageRequest.md) |  | 

### Return type

[**DummyStage**](DummyStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesDummyDestroy

> StagesDummyDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Dummy Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesDummyDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesDummyDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Dummy Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesDummyDestroyRequest struct via the builder pattern


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


## StagesDummyList

> PaginatedDummyStageList StagesDummyList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).ThrowError(throwError).Execute()





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
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    throwError := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesDummyList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).ThrowError(throwError).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesDummyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesDummyList`: PaginatedDummyStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesDummyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesDummyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **stageUuid** | **string** |  | 
 **throwError** | **bool** |  | 

### Return type

[**PaginatedDummyStageList**](PaginatedDummyStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesDummyPartialUpdate

> DummyStage StagesDummyPartialUpdate(ctx, stageUuid).PatchedDummyStageRequest(patchedDummyStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Dummy Stage.
    patchedDummyStageRequest := *openapiclient.NewPatchedDummyStageRequest() // PatchedDummyStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesDummyPartialUpdate(context.Background(), stageUuid).PatchedDummyStageRequest(patchedDummyStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesDummyPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesDummyPartialUpdate`: DummyStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesDummyPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Dummy Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesDummyPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDummyStageRequest** | [**PatchedDummyStageRequest**](PatchedDummyStageRequest.md) |  | 

### Return type

[**DummyStage**](DummyStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesDummyRetrieve

> DummyStage StagesDummyRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Dummy Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesDummyRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesDummyRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesDummyRetrieve`: DummyStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesDummyRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Dummy Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesDummyRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DummyStage**](DummyStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesDummyUpdate

> DummyStage StagesDummyUpdate(ctx, stageUuid).DummyStageRequest(dummyStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Dummy Stage.
    dummyStageRequest := *openapiclient.NewDummyStageRequest("Name_example") // DummyStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesDummyUpdate(context.Background(), stageUuid).DummyStageRequest(dummyStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesDummyUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesDummyUpdate`: DummyStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesDummyUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Dummy Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesDummyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dummyStageRequest** | [**DummyStageRequest**](DummyStageRequest.md) |  | 

### Return type

[**DummyStage**](DummyStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesDummyUsedByList

> []UsedBy StagesDummyUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Dummy Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesDummyUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesDummyUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesDummyUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesDummyUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Dummy Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesDummyUsedByListRequest struct via the builder pattern


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


## StagesEmailCreate

> EmailStage StagesEmailCreate(ctx).EmailStageRequest(emailStageRequest).Execute()





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
    emailStageRequest := *openapiclient.NewEmailStageRequest("Name_example") // EmailStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesEmailCreate(context.Background()).EmailStageRequest(emailStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesEmailCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesEmailCreate`: EmailStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesEmailCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesEmailCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailStageRequest** | [**EmailStageRequest**](EmailStageRequest.md) |  | 

### Return type

[**EmailStage**](EmailStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesEmailDestroy

> StagesEmailDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Email Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesEmailDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesEmailDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Email Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesEmailDestroyRequest struct via the builder pattern


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


## StagesEmailList

> PaginatedEmailStageList StagesEmailList(ctx).ActivateUserOnSuccess(activateUserOnSuccess).FromAddress(fromAddress).Host(host).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Port(port).Search(search).Subject(subject).Template(template).Timeout(timeout).TokenExpiry(tokenExpiry).UseGlobalSettings(useGlobalSettings).UseSsl(useSsl).UseTls(useTls).Username(username).Execute()





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
    activateUserOnSuccess := true // bool |  (optional)
    fromAddress := "fromAddress_example" // string |  (optional)
    host := "host_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    port := int32(56) // int32 |  (optional)
    search := "search_example" // string | A search term. (optional)
    subject := "subject_example" // string |  (optional)
    template := "template_example" // string |  (optional)
    timeout := int32(56) // int32 |  (optional)
    tokenExpiry := int32(56) // int32 |  (optional)
    useGlobalSettings := true // bool |  (optional)
    useSsl := true // bool |  (optional)
    useTls := true // bool |  (optional)
    username := "username_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesEmailList(context.Background()).ActivateUserOnSuccess(activateUserOnSuccess).FromAddress(fromAddress).Host(host).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Port(port).Search(search).Subject(subject).Template(template).Timeout(timeout).TokenExpiry(tokenExpiry).UseGlobalSettings(useGlobalSettings).UseSsl(useSsl).UseTls(useTls).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesEmailList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesEmailList`: PaginatedEmailStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesEmailList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesEmailListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activateUserOnSuccess** | **bool** |  | 
 **fromAddress** | **string** |  | 
 **host** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **port** | **int32** |  | 
 **search** | **string** | A search term. | 
 **subject** | **string** |  | 
 **template** | **string** |  | 
 **timeout** | **int32** |  | 
 **tokenExpiry** | **int32** |  | 
 **useGlobalSettings** | **bool** |  | 
 **useSsl** | **bool** |  | 
 **useTls** | **bool** |  | 
 **username** | **string** |  | 

### Return type

[**PaginatedEmailStageList**](PaginatedEmailStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesEmailPartialUpdate

> EmailStage StagesEmailPartialUpdate(ctx, stageUuid).PatchedEmailStageRequest(patchedEmailStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Email Stage.
    patchedEmailStageRequest := *openapiclient.NewPatchedEmailStageRequest() // PatchedEmailStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesEmailPartialUpdate(context.Background(), stageUuid).PatchedEmailStageRequest(patchedEmailStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesEmailPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesEmailPartialUpdate`: EmailStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesEmailPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Email Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesEmailPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedEmailStageRequest** | [**PatchedEmailStageRequest**](PatchedEmailStageRequest.md) |  | 

### Return type

[**EmailStage**](EmailStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesEmailRetrieve

> EmailStage StagesEmailRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Email Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesEmailRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesEmailRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesEmailRetrieve`: EmailStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesEmailRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Email Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesEmailRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailStage**](EmailStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesEmailTemplatesList

> []TypeCreate StagesEmailTemplatesList(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesEmailTemplatesList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesEmailTemplatesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesEmailTemplatesList`: []TypeCreate
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesEmailTemplatesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStagesEmailTemplatesListRequest struct via the builder pattern


### Return type

[**[]TypeCreate**](TypeCreate.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesEmailUpdate

> EmailStage StagesEmailUpdate(ctx, stageUuid).EmailStageRequest(emailStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Email Stage.
    emailStageRequest := *openapiclient.NewEmailStageRequest("Name_example") // EmailStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesEmailUpdate(context.Background(), stageUuid).EmailStageRequest(emailStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesEmailUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesEmailUpdate`: EmailStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesEmailUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Email Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesEmailUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailStageRequest** | [**EmailStageRequest**](EmailStageRequest.md) |  | 

### Return type

[**EmailStage**](EmailStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesEmailUsedByList

> []UsedBy StagesEmailUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Email Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesEmailUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesEmailUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesEmailUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesEmailUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Email Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesEmailUsedByListRequest struct via the builder pattern


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


## StagesIdentificationCreate

> IdentificationStage StagesIdentificationCreate(ctx).IdentificationStageRequest(identificationStageRequest).Execute()





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
    identificationStageRequest := *openapiclient.NewIdentificationStageRequest("Name_example") // IdentificationStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesIdentificationCreate(context.Background()).IdentificationStageRequest(identificationStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesIdentificationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesIdentificationCreate`: IdentificationStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesIdentificationCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesIdentificationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identificationStageRequest** | [**IdentificationStageRequest**](IdentificationStageRequest.md) |  | 

### Return type

[**IdentificationStage**](IdentificationStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesIdentificationDestroy

> StagesIdentificationDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Identification Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesIdentificationDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesIdentificationDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Identification Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesIdentificationDestroyRequest struct via the builder pattern


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


## StagesIdentificationList

> PaginatedIdentificationStageList StagesIdentificationList(ctx).CaseInsensitiveMatching(caseInsensitiveMatching).EnrollmentFlow(enrollmentFlow).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PasswordStage(passwordStage).PasswordlessFlow(passwordlessFlow).RecoveryFlow(recoveryFlow).Search(search).ShowMatchedUser(showMatchedUser).ShowSourceLabels(showSourceLabels).Execute()





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
    caseInsensitiveMatching := true // bool |  (optional)
    enrollmentFlow := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    passwordStage := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    passwordlessFlow := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    recoveryFlow := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    search := "search_example" // string | A search term. (optional)
    showMatchedUser := true // bool |  (optional)
    showSourceLabels := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesIdentificationList(context.Background()).CaseInsensitiveMatching(caseInsensitiveMatching).EnrollmentFlow(enrollmentFlow).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).PasswordStage(passwordStage).PasswordlessFlow(passwordlessFlow).RecoveryFlow(recoveryFlow).Search(search).ShowMatchedUser(showMatchedUser).ShowSourceLabels(showSourceLabels).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesIdentificationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesIdentificationList`: PaginatedIdentificationStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesIdentificationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesIdentificationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **caseInsensitiveMatching** | **bool** |  | 
 **enrollmentFlow** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **passwordStage** | **string** |  | 
 **passwordlessFlow** | **string** |  | 
 **recoveryFlow** | **string** |  | 
 **search** | **string** | A search term. | 
 **showMatchedUser** | **bool** |  | 
 **showSourceLabels** | **bool** |  | 

### Return type

[**PaginatedIdentificationStageList**](PaginatedIdentificationStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesIdentificationPartialUpdate

> IdentificationStage StagesIdentificationPartialUpdate(ctx, stageUuid).PatchedIdentificationStageRequest(patchedIdentificationStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Identification Stage.
    patchedIdentificationStageRequest := *openapiclient.NewPatchedIdentificationStageRequest() // PatchedIdentificationStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesIdentificationPartialUpdate(context.Background(), stageUuid).PatchedIdentificationStageRequest(patchedIdentificationStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesIdentificationPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesIdentificationPartialUpdate`: IdentificationStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesIdentificationPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Identification Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesIdentificationPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedIdentificationStageRequest** | [**PatchedIdentificationStageRequest**](PatchedIdentificationStageRequest.md) |  | 

### Return type

[**IdentificationStage**](IdentificationStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesIdentificationRetrieve

> IdentificationStage StagesIdentificationRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Identification Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesIdentificationRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesIdentificationRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesIdentificationRetrieve`: IdentificationStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesIdentificationRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Identification Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesIdentificationRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentificationStage**](IdentificationStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesIdentificationUpdate

> IdentificationStage StagesIdentificationUpdate(ctx, stageUuid).IdentificationStageRequest(identificationStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Identification Stage.
    identificationStageRequest := *openapiclient.NewIdentificationStageRequest("Name_example") // IdentificationStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesIdentificationUpdate(context.Background(), stageUuid).IdentificationStageRequest(identificationStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesIdentificationUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesIdentificationUpdate`: IdentificationStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesIdentificationUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Identification Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesIdentificationUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identificationStageRequest** | [**IdentificationStageRequest**](IdentificationStageRequest.md) |  | 

### Return type

[**IdentificationStage**](IdentificationStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesIdentificationUsedByList

> []UsedBy StagesIdentificationUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Identification Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesIdentificationUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesIdentificationUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesIdentificationUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesIdentificationUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Identification Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesIdentificationUsedByListRequest struct via the builder pattern


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


## StagesInvitationInvitationsCreate

> Invitation StagesInvitationInvitationsCreate(ctx).InvitationRequest(invitationRequest).Execute()





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
    invitationRequest := *openapiclient.NewInvitationRequest("Name_example") // InvitationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesInvitationInvitationsCreate(context.Background()).InvitationRequest(invitationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesInvitationInvitationsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesInvitationInvitationsCreate`: Invitation
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesInvitationInvitationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesInvitationInvitationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationRequest** | [**InvitationRequest**](InvitationRequest.md) |  | 

### Return type

[**Invitation**](Invitation.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesInvitationInvitationsDestroy

> StagesInvitationInvitationsDestroy(ctx, inviteUuid).Execute()





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
    inviteUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Invitation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesInvitationInvitationsDestroy(context.Background(), inviteUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesInvitationInvitationsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteUuid** | **string** | A UUID string identifying this Invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesInvitationInvitationsDestroyRequest struct via the builder pattern


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


## StagesInvitationInvitationsList

> PaginatedInvitationList StagesInvitationInvitationsList(ctx).CreatedByUsername(createdByUsername).Expires(expires).FlowSlug(flowSlug).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    createdByUsername := "createdByUsername_example" // string |  (optional)
    expires := time.Now() // time.Time |  (optional)
    flowSlug := "flowSlug_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesInvitationInvitationsList(context.Background()).CreatedByUsername(createdByUsername).Expires(expires).FlowSlug(flowSlug).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesInvitationInvitationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesInvitationInvitationsList`: PaginatedInvitationList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesInvitationInvitationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesInvitationInvitationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdByUsername** | **string** |  | 
 **expires** | **time.Time** |  | 
 **flowSlug** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedInvitationList**](PaginatedInvitationList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesInvitationInvitationsPartialUpdate

> Invitation StagesInvitationInvitationsPartialUpdate(ctx, inviteUuid).PatchedInvitationRequest(patchedInvitationRequest).Execute()





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
    inviteUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Invitation.
    patchedInvitationRequest := *openapiclient.NewPatchedInvitationRequest() // PatchedInvitationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesInvitationInvitationsPartialUpdate(context.Background(), inviteUuid).PatchedInvitationRequest(patchedInvitationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesInvitationInvitationsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesInvitationInvitationsPartialUpdate`: Invitation
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesInvitationInvitationsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteUuid** | **string** | A UUID string identifying this Invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesInvitationInvitationsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedInvitationRequest** | [**PatchedInvitationRequest**](PatchedInvitationRequest.md) |  | 

### Return type

[**Invitation**](Invitation.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesInvitationInvitationsRetrieve

> Invitation StagesInvitationInvitationsRetrieve(ctx, inviteUuid).Execute()





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
    inviteUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Invitation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesInvitationInvitationsRetrieve(context.Background(), inviteUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesInvitationInvitationsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesInvitationInvitationsRetrieve`: Invitation
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesInvitationInvitationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteUuid** | **string** | A UUID string identifying this Invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesInvitationInvitationsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invitation**](Invitation.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesInvitationInvitationsUpdate

> Invitation StagesInvitationInvitationsUpdate(ctx, inviteUuid).InvitationRequest(invitationRequest).Execute()





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
    inviteUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Invitation.
    invitationRequest := *openapiclient.NewInvitationRequest("Name_example") // InvitationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesInvitationInvitationsUpdate(context.Background(), inviteUuid).InvitationRequest(invitationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesInvitationInvitationsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesInvitationInvitationsUpdate`: Invitation
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesInvitationInvitationsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteUuid** | **string** | A UUID string identifying this Invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesInvitationInvitationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invitationRequest** | [**InvitationRequest**](InvitationRequest.md) |  | 

### Return type

[**Invitation**](Invitation.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesInvitationInvitationsUsedByList

> []UsedBy StagesInvitationInvitationsUsedByList(ctx, inviteUuid).Execute()





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
    inviteUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Invitation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesInvitationInvitationsUsedByList(context.Background(), inviteUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesInvitationInvitationsUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesInvitationInvitationsUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesInvitationInvitationsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteUuid** | **string** | A UUID string identifying this Invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesInvitationInvitationsUsedByListRequest struct via the builder pattern


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


## StagesInvitationStagesCreate

> InvitationStage StagesInvitationStagesCreate(ctx).InvitationStageRequest(invitationStageRequest).Execute()





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
    invitationStageRequest := *openapiclient.NewInvitationStageRequest("Name_example") // InvitationStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesInvitationStagesCreate(context.Background()).InvitationStageRequest(invitationStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesInvitationStagesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesInvitationStagesCreate`: InvitationStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesInvitationStagesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesInvitationStagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationStageRequest** | [**InvitationStageRequest**](InvitationStageRequest.md) |  | 

### Return type

[**InvitationStage**](InvitationStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesInvitationStagesDestroy

> StagesInvitationStagesDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Invitation Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesInvitationStagesDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesInvitationStagesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Invitation Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesInvitationStagesDestroyRequest struct via the builder pattern


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


## StagesInvitationStagesList

> PaginatedInvitationStageList StagesInvitationStagesList(ctx).ContinueFlowWithoutInvitation(continueFlowWithoutInvitation).Name(name).NoFlows(noFlows).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).Execute()





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
    continueFlowWithoutInvitation := true // bool |  (optional)
    name := "name_example" // string |  (optional)
    noFlows := true // bool |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesInvitationStagesList(context.Background()).ContinueFlowWithoutInvitation(continueFlowWithoutInvitation).Name(name).NoFlows(noFlows).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesInvitationStagesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesInvitationStagesList`: PaginatedInvitationStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesInvitationStagesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesInvitationStagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **continueFlowWithoutInvitation** | **bool** |  | 
 **name** | **string** |  | 
 **noFlows** | **bool** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **stageUuid** | **string** |  | 

### Return type

[**PaginatedInvitationStageList**](PaginatedInvitationStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesInvitationStagesPartialUpdate

> InvitationStage StagesInvitationStagesPartialUpdate(ctx, stageUuid).PatchedInvitationStageRequest(patchedInvitationStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Invitation Stage.
    patchedInvitationStageRequest := *openapiclient.NewPatchedInvitationStageRequest() // PatchedInvitationStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesInvitationStagesPartialUpdate(context.Background(), stageUuid).PatchedInvitationStageRequest(patchedInvitationStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesInvitationStagesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesInvitationStagesPartialUpdate`: InvitationStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesInvitationStagesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Invitation Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesInvitationStagesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedInvitationStageRequest** | [**PatchedInvitationStageRequest**](PatchedInvitationStageRequest.md) |  | 

### Return type

[**InvitationStage**](InvitationStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesInvitationStagesRetrieve

> InvitationStage StagesInvitationStagesRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Invitation Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesInvitationStagesRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesInvitationStagesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesInvitationStagesRetrieve`: InvitationStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesInvitationStagesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Invitation Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesInvitationStagesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvitationStage**](InvitationStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesInvitationStagesUpdate

> InvitationStage StagesInvitationStagesUpdate(ctx, stageUuid).InvitationStageRequest(invitationStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Invitation Stage.
    invitationStageRequest := *openapiclient.NewInvitationStageRequest("Name_example") // InvitationStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesInvitationStagesUpdate(context.Background(), stageUuid).InvitationStageRequest(invitationStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesInvitationStagesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesInvitationStagesUpdate`: InvitationStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesInvitationStagesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Invitation Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesInvitationStagesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invitationStageRequest** | [**InvitationStageRequest**](InvitationStageRequest.md) |  | 

### Return type

[**InvitationStage**](InvitationStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesInvitationStagesUsedByList

> []UsedBy StagesInvitationStagesUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Invitation Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesInvitationStagesUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesInvitationStagesUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesInvitationStagesUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesInvitationStagesUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Invitation Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesInvitationStagesUsedByListRequest struct via the builder pattern


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


## StagesPasswordCreate

> PasswordStage StagesPasswordCreate(ctx).PasswordStageRequest(passwordStageRequest).Execute()





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
    passwordStageRequest := *openapiclient.NewPasswordStageRequest("Name_example", []openapiclient.BackendsEnum{openapiclient.BackendsEnum("authentik.core.auth.InbuiltBackend")}) // PasswordStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPasswordCreate(context.Background()).PasswordStageRequest(passwordStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPasswordCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPasswordCreate`: PasswordStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPasswordCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesPasswordCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordStageRequest** | [**PasswordStageRequest**](PasswordStageRequest.md) |  | 

### Return type

[**PasswordStage**](PasswordStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesPasswordDestroy

> StagesPasswordDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Password Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPasswordDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPasswordDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Password Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesPasswordDestroyRequest struct via the builder pattern


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


## StagesPasswordList

> PaginatedPasswordStageList StagesPasswordList(ctx).ConfigureFlow(configureFlow).FailedAttemptsBeforeCancel(failedAttemptsBeforeCancel).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()





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
    configureFlow := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    failedAttemptsBeforeCancel := int32(56) // int32 |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPasswordList(context.Background()).ConfigureFlow(configureFlow).FailedAttemptsBeforeCancel(failedAttemptsBeforeCancel).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPasswordList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPasswordList`: PaginatedPasswordStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPasswordList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesPasswordListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configureFlow** | **string** |  | 
 **failedAttemptsBeforeCancel** | **int32** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedPasswordStageList**](PaginatedPasswordStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesPasswordPartialUpdate

> PasswordStage StagesPasswordPartialUpdate(ctx, stageUuid).PatchedPasswordStageRequest(patchedPasswordStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Password Stage.
    patchedPasswordStageRequest := *openapiclient.NewPatchedPasswordStageRequest() // PatchedPasswordStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPasswordPartialUpdate(context.Background(), stageUuid).PatchedPasswordStageRequest(patchedPasswordStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPasswordPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPasswordPartialUpdate`: PasswordStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPasswordPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Password Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesPasswordPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPasswordStageRequest** | [**PatchedPasswordStageRequest**](PatchedPasswordStageRequest.md) |  | 

### Return type

[**PasswordStage**](PasswordStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesPasswordRetrieve

> PasswordStage StagesPasswordRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Password Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPasswordRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPasswordRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPasswordRetrieve`: PasswordStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPasswordRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Password Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesPasswordRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PasswordStage**](PasswordStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesPasswordUpdate

> PasswordStage StagesPasswordUpdate(ctx, stageUuid).PasswordStageRequest(passwordStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Password Stage.
    passwordStageRequest := *openapiclient.NewPasswordStageRequest("Name_example", []openapiclient.BackendsEnum{openapiclient.BackendsEnum("authentik.core.auth.InbuiltBackend")}) // PasswordStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPasswordUpdate(context.Background(), stageUuid).PasswordStageRequest(passwordStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPasswordUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPasswordUpdate`: PasswordStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPasswordUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Password Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesPasswordUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **passwordStageRequest** | [**PasswordStageRequest**](PasswordStageRequest.md) |  | 

### Return type

[**PasswordStage**](PasswordStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesPasswordUsedByList

> []UsedBy StagesPasswordUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Password Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPasswordUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPasswordUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPasswordUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPasswordUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Password Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesPasswordUsedByListRequest struct via the builder pattern


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


## StagesPromptPromptsCreate

> Prompt StagesPromptPromptsCreate(ctx).PromptRequest(promptRequest).Execute()





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
    promptRequest := *openapiclient.NewPromptRequest("Name_example", "FieldKey_example", "Label_example", openapiclient.PromptTypeEnum("text")) // PromptRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPromptPromptsCreate(context.Background()).PromptRequest(promptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPromptPromptsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPromptPromptsCreate`: Prompt
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPromptPromptsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesPromptPromptsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promptRequest** | [**PromptRequest**](PromptRequest.md) |  | 

### Return type

[**Prompt**](Prompt.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesPromptPromptsDestroy

> StagesPromptPromptsDestroy(ctx, promptUuid).Execute()





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
    promptUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Prompt.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPromptPromptsDestroy(context.Background(), promptUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPromptPromptsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptUuid** | **string** | A UUID string identifying this Prompt. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesPromptPromptsDestroyRequest struct via the builder pattern


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


## StagesPromptPromptsList

> PaginatedPromptList StagesPromptPromptsList(ctx).FieldKey(fieldKey).Label(label).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Placeholder(placeholder).Search(search).Type_(type_).Execute()





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
    fieldKey := "fieldKey_example" // string |  (optional)
    label := "label_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    placeholder := "placeholder_example" // string |  (optional)
    search := "search_example" // string | A search term. (optional)
    type_ := "type__example" // string | * `text` - Text: Simple Text input * `text_area` - Text area: Multiline Text Input. * `text_read_only` - Text (read-only): Simple Text input, but cannot be edited. * `text_area_read_only` - Text area (read-only): Multiline Text input, but cannot be edited. * `username` - Username: Same as Text input, but checks for and prevents duplicate usernames. * `email` - Email: Text field with Email type. * `password` - Password: Masked input, password is validated against sources. Policies still have to be applied to this Stage. If two of these are used in the same stage, they are ensured to be identical. * `number` - Number * `checkbox` - Checkbox * `radio-button-group` - Fixed choice field rendered as a group of radio buttons. * `dropdown` - Fixed choice field rendered as a dropdown. * `date` - Date * `date-time` - Date Time * `file` - File: File upload for arbitrary files. File content will be available in flow context as data-URI * `separator` - Separator: Static Separator Line * `hidden` - Hidden: Hidden field, can be used to insert data into form. * `static` - Static: Static value, displayed as-is. * `ak-locale` - authentik: Selection of locales authentik supports  * `text` - Text: Simple Text input * `text_area` - Text area: Multiline Text Input. * `text_read_only` - Text (read-only): Simple Text input, but cannot be edited. * `text_area_read_only` - Text area (read-only): Multiline Text input, but cannot be edited. * `username` - Username: Same as Text input, but checks for and prevents duplicate usernames. * `email` - Email: Text field with Email type. * `password` - Password: Masked input, password is validated against sources. Policies still have to be applied to this Stage. If two of these are used in the same stage, they are ensured to be identical. * `number` - Number * `checkbox` - Checkbox * `radio-button-group` - Fixed choice field rendered as a group of radio buttons. * `dropdown` - Fixed choice field rendered as a dropdown. * `date` - Date * `date-time` - Date Time * `file` - File: File upload for arbitrary files. File content will be available in flow context as data-URI * `separator` - Separator: Static Separator Line * `hidden` - Hidden: Hidden field, can be used to insert data into form. * `static` - Static: Static value, displayed as-is. * `ak-locale` - authentik: Selection of locales authentik supports (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPromptPromptsList(context.Background()).FieldKey(fieldKey).Label(label).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Placeholder(placeholder).Search(search).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPromptPromptsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPromptPromptsList`: PaginatedPromptList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPromptPromptsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesPromptPromptsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldKey** | **string** |  | 
 **label** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **placeholder** | **string** |  | 
 **search** | **string** | A search term. | 
 **type_** | **string** | * &#x60;text&#x60; - Text: Simple Text input * &#x60;text_area&#x60; - Text area: Multiline Text Input. * &#x60;text_read_only&#x60; - Text (read-only): Simple Text input, but cannot be edited. * &#x60;text_area_read_only&#x60; - Text area (read-only): Multiline Text input, but cannot be edited. * &#x60;username&#x60; - Username: Same as Text input, but checks for and prevents duplicate usernames. * &#x60;email&#x60; - Email: Text field with Email type. * &#x60;password&#x60; - Password: Masked input, password is validated against sources. Policies still have to be applied to this Stage. If two of these are used in the same stage, they are ensured to be identical. * &#x60;number&#x60; - Number * &#x60;checkbox&#x60; - Checkbox * &#x60;radio-button-group&#x60; - Fixed choice field rendered as a group of radio buttons. * &#x60;dropdown&#x60; - Fixed choice field rendered as a dropdown. * &#x60;date&#x60; - Date * &#x60;date-time&#x60; - Date Time * &#x60;file&#x60; - File: File upload for arbitrary files. File content will be available in flow context as data-URI * &#x60;separator&#x60; - Separator: Static Separator Line * &#x60;hidden&#x60; - Hidden: Hidden field, can be used to insert data into form. * &#x60;static&#x60; - Static: Static value, displayed as-is. * &#x60;ak-locale&#x60; - authentik: Selection of locales authentik supports  * &#x60;text&#x60; - Text: Simple Text input * &#x60;text_area&#x60; - Text area: Multiline Text Input. * &#x60;text_read_only&#x60; - Text (read-only): Simple Text input, but cannot be edited. * &#x60;text_area_read_only&#x60; - Text area (read-only): Multiline Text input, but cannot be edited. * &#x60;username&#x60; - Username: Same as Text input, but checks for and prevents duplicate usernames. * &#x60;email&#x60; - Email: Text field with Email type. * &#x60;password&#x60; - Password: Masked input, password is validated against sources. Policies still have to be applied to this Stage. If two of these are used in the same stage, they are ensured to be identical. * &#x60;number&#x60; - Number * &#x60;checkbox&#x60; - Checkbox * &#x60;radio-button-group&#x60; - Fixed choice field rendered as a group of radio buttons. * &#x60;dropdown&#x60; - Fixed choice field rendered as a dropdown. * &#x60;date&#x60; - Date * &#x60;date-time&#x60; - Date Time * &#x60;file&#x60; - File: File upload for arbitrary files. File content will be available in flow context as data-URI * &#x60;separator&#x60; - Separator: Static Separator Line * &#x60;hidden&#x60; - Hidden: Hidden field, can be used to insert data into form. * &#x60;static&#x60; - Static: Static value, displayed as-is. * &#x60;ak-locale&#x60; - authentik: Selection of locales authentik supports | 

### Return type

[**PaginatedPromptList**](PaginatedPromptList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesPromptPromptsPartialUpdate

> Prompt StagesPromptPromptsPartialUpdate(ctx, promptUuid).PatchedPromptRequest(patchedPromptRequest).Execute()





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
    promptUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Prompt.
    patchedPromptRequest := *openapiclient.NewPatchedPromptRequest() // PatchedPromptRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPromptPromptsPartialUpdate(context.Background(), promptUuid).PatchedPromptRequest(patchedPromptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPromptPromptsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPromptPromptsPartialUpdate`: Prompt
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPromptPromptsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptUuid** | **string** | A UUID string identifying this Prompt. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesPromptPromptsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPromptRequest** | [**PatchedPromptRequest**](PatchedPromptRequest.md) |  | 

### Return type

[**Prompt**](Prompt.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesPromptPromptsRetrieve

> Prompt StagesPromptPromptsRetrieve(ctx, promptUuid).Execute()





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
    promptUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Prompt.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPromptPromptsRetrieve(context.Background(), promptUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPromptPromptsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPromptPromptsRetrieve`: Prompt
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPromptPromptsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptUuid** | **string** | A UUID string identifying this Prompt. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesPromptPromptsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Prompt**](Prompt.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesPromptPromptsUpdate

> Prompt StagesPromptPromptsUpdate(ctx, promptUuid).PromptRequest(promptRequest).Execute()





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
    promptUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Prompt.
    promptRequest := *openapiclient.NewPromptRequest("Name_example", "FieldKey_example", "Label_example", openapiclient.PromptTypeEnum("text")) // PromptRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPromptPromptsUpdate(context.Background(), promptUuid).PromptRequest(promptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPromptPromptsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPromptPromptsUpdate`: Prompt
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPromptPromptsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptUuid** | **string** | A UUID string identifying this Prompt. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesPromptPromptsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **promptRequest** | [**PromptRequest**](PromptRequest.md) |  | 

### Return type

[**Prompt**](Prompt.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesPromptPromptsUsedByList

> []UsedBy StagesPromptPromptsUsedByList(ctx, promptUuid).Execute()





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
    promptUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Prompt.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPromptPromptsUsedByList(context.Background(), promptUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPromptPromptsUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPromptPromptsUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPromptPromptsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promptUuid** | **string** | A UUID string identifying this Prompt. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesPromptPromptsUsedByListRequest struct via the builder pattern


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


## StagesPromptStagesCreate

> PromptStage StagesPromptStagesCreate(ctx).PromptStageRequest(promptStageRequest).Execute()





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
    promptStageRequest := *openapiclient.NewPromptStageRequest("Name_example", []string{"Fields_example"}) // PromptStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPromptStagesCreate(context.Background()).PromptStageRequest(promptStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPromptStagesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPromptStagesCreate`: PromptStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPromptStagesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesPromptStagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promptStageRequest** | [**PromptStageRequest**](PromptStageRequest.md) |  | 

### Return type

[**PromptStage**](PromptStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesPromptStagesDestroy

> StagesPromptStagesDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Prompt Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPromptStagesDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPromptStagesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Prompt Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesPromptStagesDestroyRequest struct via the builder pattern


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


## StagesPromptStagesList

> PaginatedPromptStageList StagesPromptStagesList(ctx).Fields(fields).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).ValidationPolicies(validationPolicies).Execute()





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
    fields := []string{"Inner_example"} // []string |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    validationPolicies := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPromptStagesList(context.Background()).Fields(fields).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).ValidationPolicies(validationPolicies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPromptStagesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPromptStagesList`: PaginatedPromptStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPromptStagesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesPromptStagesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **stageUuid** | **string** |  | 
 **validationPolicies** | **[]string** |  | 

### Return type

[**PaginatedPromptStageList**](PaginatedPromptStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesPromptStagesPartialUpdate

> PromptStage StagesPromptStagesPartialUpdate(ctx, stageUuid).PatchedPromptStageRequest(patchedPromptStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Prompt Stage.
    patchedPromptStageRequest := *openapiclient.NewPatchedPromptStageRequest() // PatchedPromptStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPromptStagesPartialUpdate(context.Background(), stageUuid).PatchedPromptStageRequest(patchedPromptStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPromptStagesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPromptStagesPartialUpdate`: PromptStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPromptStagesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Prompt Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesPromptStagesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPromptStageRequest** | [**PatchedPromptStageRequest**](PatchedPromptStageRequest.md) |  | 

### Return type

[**PromptStage**](PromptStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesPromptStagesRetrieve

> PromptStage StagesPromptStagesRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Prompt Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPromptStagesRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPromptStagesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPromptStagesRetrieve`: PromptStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPromptStagesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Prompt Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesPromptStagesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PromptStage**](PromptStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesPromptStagesUpdate

> PromptStage StagesPromptStagesUpdate(ctx, stageUuid).PromptStageRequest(promptStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Prompt Stage.
    promptStageRequest := *openapiclient.NewPromptStageRequest("Name_example", []string{"Fields_example"}) // PromptStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPromptStagesUpdate(context.Background(), stageUuid).PromptStageRequest(promptStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPromptStagesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPromptStagesUpdate`: PromptStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPromptStagesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Prompt Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesPromptStagesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **promptStageRequest** | [**PromptStageRequest**](PromptStageRequest.md) |  | 

### Return type

[**PromptStage**](PromptStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesPromptStagesUsedByList

> []UsedBy StagesPromptStagesUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Prompt Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesPromptStagesUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesPromptStagesUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesPromptStagesUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesPromptStagesUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this Prompt Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesPromptStagesUsedByListRequest struct via the builder pattern


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


## StagesUserDeleteCreate

> UserDeleteStage StagesUserDeleteCreate(ctx).UserDeleteStageRequest(userDeleteStageRequest).Execute()





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
    userDeleteStageRequest := *openapiclient.NewUserDeleteStageRequest("Name_example") // UserDeleteStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserDeleteCreate(context.Background()).UserDeleteStageRequest(userDeleteStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserDeleteCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserDeleteCreate`: UserDeleteStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserDeleteCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserDeleteCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userDeleteStageRequest** | [**UserDeleteStageRequest**](UserDeleteStageRequest.md) |  | 

### Return type

[**UserDeleteStage**](UserDeleteStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserDeleteDestroy

> StagesUserDeleteDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Delete Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserDeleteDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserDeleteDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Delete Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserDeleteDestroyRequest struct via the builder pattern


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


## StagesUserDeleteList

> PaginatedUserDeleteStageList StagesUserDeleteList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).Execute()





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
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserDeleteList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserDeleteList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserDeleteList`: PaginatedUserDeleteStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserDeleteList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserDeleteListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **stageUuid** | **string** |  | 

### Return type

[**PaginatedUserDeleteStageList**](PaginatedUserDeleteStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserDeletePartialUpdate

> UserDeleteStage StagesUserDeletePartialUpdate(ctx, stageUuid).PatchedUserDeleteStageRequest(patchedUserDeleteStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Delete Stage.
    patchedUserDeleteStageRequest := *openapiclient.NewPatchedUserDeleteStageRequest() // PatchedUserDeleteStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserDeletePartialUpdate(context.Background(), stageUuid).PatchedUserDeleteStageRequest(patchedUserDeleteStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserDeletePartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserDeletePartialUpdate`: UserDeleteStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserDeletePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Delete Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserDeletePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedUserDeleteStageRequest** | [**PatchedUserDeleteStageRequest**](PatchedUserDeleteStageRequest.md) |  | 

### Return type

[**UserDeleteStage**](UserDeleteStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserDeleteRetrieve

> UserDeleteStage StagesUserDeleteRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Delete Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserDeleteRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserDeleteRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserDeleteRetrieve`: UserDeleteStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserDeleteRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Delete Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserDeleteRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserDeleteStage**](UserDeleteStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserDeleteUpdate

> UserDeleteStage StagesUserDeleteUpdate(ctx, stageUuid).UserDeleteStageRequest(userDeleteStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Delete Stage.
    userDeleteStageRequest := *openapiclient.NewUserDeleteStageRequest("Name_example") // UserDeleteStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserDeleteUpdate(context.Background(), stageUuid).UserDeleteStageRequest(userDeleteStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserDeleteUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserDeleteUpdate`: UserDeleteStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserDeleteUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Delete Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserDeleteUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userDeleteStageRequest** | [**UserDeleteStageRequest**](UserDeleteStageRequest.md) |  | 

### Return type

[**UserDeleteStage**](UserDeleteStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserDeleteUsedByList

> []UsedBy StagesUserDeleteUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Delete Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserDeleteUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserDeleteUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserDeleteUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserDeleteUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Delete Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserDeleteUsedByListRequest struct via the builder pattern


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


## StagesUserLoginCreate

> UserLoginStage StagesUserLoginCreate(ctx).UserLoginStageRequest(userLoginStageRequest).Execute()





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
    userLoginStageRequest := *openapiclient.NewUserLoginStageRequest("Name_example") // UserLoginStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserLoginCreate(context.Background()).UserLoginStageRequest(userLoginStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserLoginCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserLoginCreate`: UserLoginStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserLoginCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserLoginCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userLoginStageRequest** | [**UserLoginStageRequest**](UserLoginStageRequest.md) |  | 

### Return type

[**UserLoginStage**](UserLoginStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserLoginDestroy

> StagesUserLoginDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Login Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserLoginDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserLoginDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Login Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserLoginDestroyRequest struct via the builder pattern


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


## StagesUserLoginList

> PaginatedUserLoginStageList StagesUserLoginList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).RememberMeOffset(rememberMeOffset).Search(search).SessionDuration(sessionDuration).StageUuid(stageUuid).TerminateOtherSessions(terminateOtherSessions).Execute()





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
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    rememberMeOffset := "rememberMeOffset_example" // string |  (optional)
    search := "search_example" // string | A search term. (optional)
    sessionDuration := "sessionDuration_example" // string |  (optional)
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    terminateOtherSessions := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserLoginList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).RememberMeOffset(rememberMeOffset).Search(search).SessionDuration(sessionDuration).StageUuid(stageUuid).TerminateOtherSessions(terminateOtherSessions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserLoginList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserLoginList`: PaginatedUserLoginStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserLoginList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserLoginListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **rememberMeOffset** | **string** |  | 
 **search** | **string** | A search term. | 
 **sessionDuration** | **string** |  | 
 **stageUuid** | **string** |  | 
 **terminateOtherSessions** | **bool** |  | 

### Return type

[**PaginatedUserLoginStageList**](PaginatedUserLoginStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserLoginPartialUpdate

> UserLoginStage StagesUserLoginPartialUpdate(ctx, stageUuid).PatchedUserLoginStageRequest(patchedUserLoginStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Login Stage.
    patchedUserLoginStageRequest := *openapiclient.NewPatchedUserLoginStageRequest() // PatchedUserLoginStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserLoginPartialUpdate(context.Background(), stageUuid).PatchedUserLoginStageRequest(patchedUserLoginStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserLoginPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserLoginPartialUpdate`: UserLoginStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserLoginPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Login Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserLoginPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedUserLoginStageRequest** | [**PatchedUserLoginStageRequest**](PatchedUserLoginStageRequest.md) |  | 

### Return type

[**UserLoginStage**](UserLoginStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserLoginRetrieve

> UserLoginStage StagesUserLoginRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Login Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserLoginRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserLoginRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserLoginRetrieve`: UserLoginStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserLoginRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Login Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserLoginRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserLoginStage**](UserLoginStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserLoginUpdate

> UserLoginStage StagesUserLoginUpdate(ctx, stageUuid).UserLoginStageRequest(userLoginStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Login Stage.
    userLoginStageRequest := *openapiclient.NewUserLoginStageRequest("Name_example") // UserLoginStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserLoginUpdate(context.Background(), stageUuid).UserLoginStageRequest(userLoginStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserLoginUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserLoginUpdate`: UserLoginStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserLoginUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Login Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserLoginUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userLoginStageRequest** | [**UserLoginStageRequest**](UserLoginStageRequest.md) |  | 

### Return type

[**UserLoginStage**](UserLoginStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserLoginUsedByList

> []UsedBy StagesUserLoginUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Login Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserLoginUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserLoginUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserLoginUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserLoginUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Login Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserLoginUsedByListRequest struct via the builder pattern


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


## StagesUserLogoutCreate

> UserLogoutStage StagesUserLogoutCreate(ctx).UserLogoutStageRequest(userLogoutStageRequest).Execute()





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
    userLogoutStageRequest := *openapiclient.NewUserLogoutStageRequest("Name_example") // UserLogoutStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserLogoutCreate(context.Background()).UserLogoutStageRequest(userLogoutStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserLogoutCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserLogoutCreate`: UserLogoutStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserLogoutCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserLogoutCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userLogoutStageRequest** | [**UserLogoutStageRequest**](UserLogoutStageRequest.md) |  | 

### Return type

[**UserLogoutStage**](UserLogoutStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserLogoutDestroy

> StagesUserLogoutDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Logout Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserLogoutDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserLogoutDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Logout Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserLogoutDestroyRequest struct via the builder pattern


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


## StagesUserLogoutList

> PaginatedUserLogoutStageList StagesUserLogoutList(ctx).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).Execute()





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
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserLogoutList(context.Background()).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserLogoutList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserLogoutList`: PaginatedUserLogoutStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserLogoutList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserLogoutListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **stageUuid** | **string** |  | 

### Return type

[**PaginatedUserLogoutStageList**](PaginatedUserLogoutStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserLogoutPartialUpdate

> UserLogoutStage StagesUserLogoutPartialUpdate(ctx, stageUuid).PatchedUserLogoutStageRequest(patchedUserLogoutStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Logout Stage.
    patchedUserLogoutStageRequest := *openapiclient.NewPatchedUserLogoutStageRequest() // PatchedUserLogoutStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserLogoutPartialUpdate(context.Background(), stageUuid).PatchedUserLogoutStageRequest(patchedUserLogoutStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserLogoutPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserLogoutPartialUpdate`: UserLogoutStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserLogoutPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Logout Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserLogoutPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedUserLogoutStageRequest** | [**PatchedUserLogoutStageRequest**](PatchedUserLogoutStageRequest.md) |  | 

### Return type

[**UserLogoutStage**](UserLogoutStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserLogoutRetrieve

> UserLogoutStage StagesUserLogoutRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Logout Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserLogoutRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserLogoutRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserLogoutRetrieve`: UserLogoutStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserLogoutRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Logout Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserLogoutRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserLogoutStage**](UserLogoutStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserLogoutUpdate

> UserLogoutStage StagesUserLogoutUpdate(ctx, stageUuid).UserLogoutStageRequest(userLogoutStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Logout Stage.
    userLogoutStageRequest := *openapiclient.NewUserLogoutStageRequest("Name_example") // UserLogoutStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserLogoutUpdate(context.Background(), stageUuid).UserLogoutStageRequest(userLogoutStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserLogoutUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserLogoutUpdate`: UserLogoutStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserLogoutUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Logout Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserLogoutUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userLogoutStageRequest** | [**UserLogoutStageRequest**](UserLogoutStageRequest.md) |  | 

### Return type

[**UserLogoutStage**](UserLogoutStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserLogoutUsedByList

> []UsedBy StagesUserLogoutUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Logout Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserLogoutUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserLogoutUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserLogoutUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserLogoutUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Logout Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserLogoutUsedByListRequest struct via the builder pattern


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


## StagesUserWriteCreate

> UserWriteStage StagesUserWriteCreate(ctx).UserWriteStageRequest(userWriteStageRequest).Execute()





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
    userWriteStageRequest := *openapiclient.NewUserWriteStageRequest("Name_example") // UserWriteStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserWriteCreate(context.Background()).UserWriteStageRequest(userWriteStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserWriteCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserWriteCreate`: UserWriteStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserWriteCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserWriteCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userWriteStageRequest** | [**UserWriteStageRequest**](UserWriteStageRequest.md) |  | 

### Return type

[**UserWriteStage**](UserWriteStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserWriteDestroy

> StagesUserWriteDestroy(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Write Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserWriteDestroy(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserWriteDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Write Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserWriteDestroyRequest struct via the builder pattern


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


## StagesUserWriteList

> PaginatedUserWriteStageList StagesUserWriteList(ctx).CreateUsersAsInactive(createUsersAsInactive).CreateUsersGroup(createUsersGroup).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).UserCreationMode(userCreationMode).UserPathTemplate(userPathTemplate).Execute()





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
    createUsersAsInactive := true // bool |  (optional)
    createUsersGroup := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    userCreationMode := "userCreationMode_example" // string | * `never_create` - Never Create * `create_when_required` - Create When Required * `always_create` - Always Create  * `never_create` - Never Create * `create_when_required` - Create When Required * `always_create` - Always Create (optional)
    userPathTemplate := "userPathTemplate_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserWriteList(context.Background()).CreateUsersAsInactive(createUsersAsInactive).CreateUsersGroup(createUsersGroup).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).StageUuid(stageUuid).UserCreationMode(userCreationMode).UserPathTemplate(userPathTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserWriteList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserWriteList`: PaginatedUserWriteStageList
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserWriteList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserWriteListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUsersAsInactive** | **bool** |  | 
 **createUsersGroup** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **stageUuid** | **string** |  | 
 **userCreationMode** | **string** | * &#x60;never_create&#x60; - Never Create * &#x60;create_when_required&#x60; - Create When Required * &#x60;always_create&#x60; - Always Create  * &#x60;never_create&#x60; - Never Create * &#x60;create_when_required&#x60; - Create When Required * &#x60;always_create&#x60; - Always Create | 
 **userPathTemplate** | **string** |  | 

### Return type

[**PaginatedUserWriteStageList**](PaginatedUserWriteStageList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserWritePartialUpdate

> UserWriteStage StagesUserWritePartialUpdate(ctx, stageUuid).PatchedUserWriteStageRequest(patchedUserWriteStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Write Stage.
    patchedUserWriteStageRequest := *openapiclient.NewPatchedUserWriteStageRequest() // PatchedUserWriteStageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserWritePartialUpdate(context.Background(), stageUuid).PatchedUserWriteStageRequest(patchedUserWriteStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserWritePartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserWritePartialUpdate`: UserWriteStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserWritePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Write Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserWritePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedUserWriteStageRequest** | [**PatchedUserWriteStageRequest**](PatchedUserWriteStageRequest.md) |  | 

### Return type

[**UserWriteStage**](UserWriteStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserWriteRetrieve

> UserWriteStage StagesUserWriteRetrieve(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Write Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserWriteRetrieve(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserWriteRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserWriteRetrieve`: UserWriteStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserWriteRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Write Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserWriteRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserWriteStage**](UserWriteStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserWriteUpdate

> UserWriteStage StagesUserWriteUpdate(ctx, stageUuid).UserWriteStageRequest(userWriteStageRequest).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Write Stage.
    userWriteStageRequest := *openapiclient.NewUserWriteStageRequest("Name_example") // UserWriteStageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserWriteUpdate(context.Background(), stageUuid).UserWriteStageRequest(userWriteStageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserWriteUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserWriteUpdate`: UserWriteStage
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserWriteUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Write Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserWriteUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userWriteStageRequest** | [**UserWriteStageRequest**](UserWriteStageRequest.md) |  | 

### Return type

[**UserWriteStage**](UserWriteStage.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StagesUserWriteUsedByList

> []UsedBy StagesUserWriteUsedByList(ctx, stageUuid).Execute()





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
    stageUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this User Write Stage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagesApi.StagesUserWriteUsedByList(context.Background(), stageUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagesApi.StagesUserWriteUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StagesUserWriteUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `StagesApi.StagesUserWriteUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageUuid** | **string** | A UUID string identifying this User Write Stage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStagesUserWriteUsedByListRequest struct via the builder pattern


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

