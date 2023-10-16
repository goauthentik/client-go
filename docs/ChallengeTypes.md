# ChallengeTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ChallengeChoices**](ChallengeChoices.md) |  | 
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-user-login"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**PendingUser** | **string** |  | 
**PendingUserAvatar** | **string** |  | 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**ClientId** | **string** |  | 
**Scope** | **string** |  | 
**RedirectUri** | **string** |  | 
**State** | **string** |  | 
**ActivationBarcode** | **string** |  | 
**ActivationCode** | **string** |  | 
**StageUuid** | **string** |  | 
**PhoneNumberRequired** | Pointer to **bool** |  | [optional] [default to true]
**Codes** | **[]string** |  | 
**ConfigUrl** | **string** |  | 
**DeviceChallenges** | [**[]DeviceChallenge**](DeviceChallenge.md) |  | 
**ConfigurationStages** | [**[]SelectableStage**](SelectableStage.md) |  | 
**Registration** | **map[string]interface{}** |  | 
**Url** | **string** |  | 
**Attrs** | **map[string]string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**SiteKey** | **string** |  | 
**JsUrl** | **string** |  | 
**HeaderText** | Pointer to **string** |  | [optional] 
**Permissions** | [**[]ConsentPermission**](ConsentPermission.md) |  | 
**AdditionalPermissions** | [**[]ConsentPermission**](ConsentPermission.md) |  | 
**Token** | **string** |  | 
**RequestId** | **string** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Traceback** | Pointer to **string** |  | [optional] 
**UserFields** | **[]string** |  | 
**PasswordFields** | **bool** |  | 
**ApplicationPre** | Pointer to **string** |  | [optional] 
**EnrollUrl** | Pointer to **string** |  | [optional] 
**RecoveryUrl** | Pointer to **string** |  | [optional] 
**PasswordlessUrl** | Pointer to **string** |  | [optional] 
**PrimaryAction** | **string** |  | 
**Sources** | Pointer to [**[]LoginSource**](LoginSource.md) |  | [optional] 
**ShowSourceLabels** | **bool** |  | 
**Slug** | **string** |  | 
**Fields** | [**[]StagePrompt**](StagePrompt.md) |  | 
**To** | **string** |  | 
**Body** | **string** |  | 

## Methods

### NewChallengeTypes

`func NewChallengeTypes(type_ ChallengeChoices, pendingUser string, pendingUserAvatar string, clientId string, scope string, redirectUri string, state string, activationBarcode string, activationCode string, stageUuid string, codes []string, configUrl string, deviceChallenges []DeviceChallenge, configurationStages []SelectableStage, registration map[string]interface{}, url string, attrs map[string]string, siteKey string, jsUrl string, permissions []ConsentPermission, additionalPermissions []ConsentPermission, token string, requestId string, userFields []string, passwordFields bool, primaryAction string, showSourceLabels bool, slug string, fields []StagePrompt, to string, body string, ) *ChallengeTypes`

NewChallengeTypes instantiates a new ChallengeTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeTypesWithDefaults

`func NewChallengeTypesWithDefaults() *ChallengeTypes`

NewChallengeTypesWithDefaults instantiates a new ChallengeTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChallengeTypes) GetType() ChallengeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChallengeTypes) GetTypeOk() (*ChallengeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChallengeTypes) SetType(v ChallengeChoices)`

SetType sets Type field to given value.


### GetFlowInfo

`func (o *ChallengeTypes) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *ChallengeTypes) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *ChallengeTypes) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *ChallengeTypes) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *ChallengeTypes) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ChallengeTypes) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ChallengeTypes) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ChallengeTypes) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *ChallengeTypes) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *ChallengeTypes) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *ChallengeTypes) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *ChallengeTypes) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetPendingUser

`func (o *ChallengeTypes) GetPendingUser() string`

GetPendingUser returns the PendingUser field if non-nil, zero value otherwise.

### GetPendingUserOk

`func (o *ChallengeTypes) GetPendingUserOk() (*string, bool)`

GetPendingUserOk returns a tuple with the PendingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUser

`func (o *ChallengeTypes) SetPendingUser(v string)`

SetPendingUser sets PendingUser field to given value.


### GetPendingUserAvatar

`func (o *ChallengeTypes) GetPendingUserAvatar() string`

GetPendingUserAvatar returns the PendingUserAvatar field if non-nil, zero value otherwise.

### GetPendingUserAvatarOk

`func (o *ChallengeTypes) GetPendingUserAvatarOk() (*string, bool)`

GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUserAvatar

`func (o *ChallengeTypes) SetPendingUserAvatar(v string)`

SetPendingUserAvatar sets PendingUserAvatar field to given value.


### GetErrorMessage

`func (o *ChallengeTypes) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ChallengeTypes) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ChallengeTypes) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ChallengeTypes) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetClientId

`func (o *ChallengeTypes) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ChallengeTypes) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ChallengeTypes) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetScope

`func (o *ChallengeTypes) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ChallengeTypes) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ChallengeTypes) SetScope(v string)`

SetScope sets Scope field to given value.


### GetRedirectUri

`func (o *ChallengeTypes) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *ChallengeTypes) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *ChallengeTypes) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.


### GetState

`func (o *ChallengeTypes) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ChallengeTypes) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ChallengeTypes) SetState(v string)`

SetState sets State field to given value.


### GetActivationBarcode

`func (o *ChallengeTypes) GetActivationBarcode() string`

GetActivationBarcode returns the ActivationBarcode field if non-nil, zero value otherwise.

### GetActivationBarcodeOk

`func (o *ChallengeTypes) GetActivationBarcodeOk() (*string, bool)`

GetActivationBarcodeOk returns a tuple with the ActivationBarcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationBarcode

`func (o *ChallengeTypes) SetActivationBarcode(v string)`

SetActivationBarcode sets ActivationBarcode field to given value.


### GetActivationCode

`func (o *ChallengeTypes) GetActivationCode() string`

GetActivationCode returns the ActivationCode field if non-nil, zero value otherwise.

### GetActivationCodeOk

`func (o *ChallengeTypes) GetActivationCodeOk() (*string, bool)`

GetActivationCodeOk returns a tuple with the ActivationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationCode

`func (o *ChallengeTypes) SetActivationCode(v string)`

SetActivationCode sets ActivationCode field to given value.


### GetStageUuid

`func (o *ChallengeTypes) GetStageUuid() string`

GetStageUuid returns the StageUuid field if non-nil, zero value otherwise.

### GetStageUuidOk

`func (o *ChallengeTypes) GetStageUuidOk() (*string, bool)`

GetStageUuidOk returns a tuple with the StageUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageUuid

`func (o *ChallengeTypes) SetStageUuid(v string)`

SetStageUuid sets StageUuid field to given value.


### GetPhoneNumberRequired

`func (o *ChallengeTypes) GetPhoneNumberRequired() bool`

GetPhoneNumberRequired returns the PhoneNumberRequired field if non-nil, zero value otherwise.

### GetPhoneNumberRequiredOk

`func (o *ChallengeTypes) GetPhoneNumberRequiredOk() (*bool, bool)`

GetPhoneNumberRequiredOk returns a tuple with the PhoneNumberRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberRequired

`func (o *ChallengeTypes) SetPhoneNumberRequired(v bool)`

SetPhoneNumberRequired sets PhoneNumberRequired field to given value.

### HasPhoneNumberRequired

`func (o *ChallengeTypes) HasPhoneNumberRequired() bool`

HasPhoneNumberRequired returns a boolean if a field has been set.

### GetCodes

`func (o *ChallengeTypes) GetCodes() []string`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *ChallengeTypes) GetCodesOk() (*[]string, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *ChallengeTypes) SetCodes(v []string)`

SetCodes sets Codes field to given value.


### GetConfigUrl

`func (o *ChallengeTypes) GetConfigUrl() string`

GetConfigUrl returns the ConfigUrl field if non-nil, zero value otherwise.

### GetConfigUrlOk

`func (o *ChallengeTypes) GetConfigUrlOk() (*string, bool)`

GetConfigUrlOk returns a tuple with the ConfigUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigUrl

`func (o *ChallengeTypes) SetConfigUrl(v string)`

SetConfigUrl sets ConfigUrl field to given value.


### GetDeviceChallenges

`func (o *ChallengeTypes) GetDeviceChallenges() []DeviceChallenge`

GetDeviceChallenges returns the DeviceChallenges field if non-nil, zero value otherwise.

### GetDeviceChallengesOk

`func (o *ChallengeTypes) GetDeviceChallengesOk() (*[]DeviceChallenge, bool)`

GetDeviceChallengesOk returns a tuple with the DeviceChallenges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceChallenges

`func (o *ChallengeTypes) SetDeviceChallenges(v []DeviceChallenge)`

SetDeviceChallenges sets DeviceChallenges field to given value.


### GetConfigurationStages

`func (o *ChallengeTypes) GetConfigurationStages() []SelectableStage`

GetConfigurationStages returns the ConfigurationStages field if non-nil, zero value otherwise.

### GetConfigurationStagesOk

`func (o *ChallengeTypes) GetConfigurationStagesOk() (*[]SelectableStage, bool)`

GetConfigurationStagesOk returns a tuple with the ConfigurationStages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationStages

`func (o *ChallengeTypes) SetConfigurationStages(v []SelectableStage)`

SetConfigurationStages sets ConfigurationStages field to given value.


### GetRegistration

`func (o *ChallengeTypes) GetRegistration() map[string]interface{}`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *ChallengeTypes) GetRegistrationOk() (*map[string]interface{}, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *ChallengeTypes) SetRegistration(v map[string]interface{})`

SetRegistration sets Registration field to given value.


### GetUrl

`func (o *ChallengeTypes) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChallengeTypes) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChallengeTypes) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAttrs

`func (o *ChallengeTypes) GetAttrs() map[string]string`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *ChallengeTypes) GetAttrsOk() (*map[string]string, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *ChallengeTypes) SetAttrs(v map[string]string)`

SetAttrs sets Attrs field to given value.


### GetTitle

`func (o *ChallengeTypes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChallengeTypes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChallengeTypes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ChallengeTypes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSiteKey

`func (o *ChallengeTypes) GetSiteKey() string`

GetSiteKey returns the SiteKey field if non-nil, zero value otherwise.

### GetSiteKeyOk

`func (o *ChallengeTypes) GetSiteKeyOk() (*string, bool)`

GetSiteKeyOk returns a tuple with the SiteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteKey

`func (o *ChallengeTypes) SetSiteKey(v string)`

SetSiteKey sets SiteKey field to given value.


### GetJsUrl

`func (o *ChallengeTypes) GetJsUrl() string`

GetJsUrl returns the JsUrl field if non-nil, zero value otherwise.

### GetJsUrlOk

`func (o *ChallengeTypes) GetJsUrlOk() (*string, bool)`

GetJsUrlOk returns a tuple with the JsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsUrl

`func (o *ChallengeTypes) SetJsUrl(v string)`

SetJsUrl sets JsUrl field to given value.


### GetHeaderText

`func (o *ChallengeTypes) GetHeaderText() string`

GetHeaderText returns the HeaderText field if non-nil, zero value otherwise.

### GetHeaderTextOk

`func (o *ChallengeTypes) GetHeaderTextOk() (*string, bool)`

GetHeaderTextOk returns a tuple with the HeaderText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderText

`func (o *ChallengeTypes) SetHeaderText(v string)`

SetHeaderText sets HeaderText field to given value.

### HasHeaderText

`func (o *ChallengeTypes) HasHeaderText() bool`

HasHeaderText returns a boolean if a field has been set.

### GetPermissions

`func (o *ChallengeTypes) GetPermissions() []ConsentPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ChallengeTypes) GetPermissionsOk() (*[]ConsentPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ChallengeTypes) SetPermissions(v []ConsentPermission)`

SetPermissions sets Permissions field to given value.


### GetAdditionalPermissions

`func (o *ChallengeTypes) GetAdditionalPermissions() []ConsentPermission`

GetAdditionalPermissions returns the AdditionalPermissions field if non-nil, zero value otherwise.

### GetAdditionalPermissionsOk

`func (o *ChallengeTypes) GetAdditionalPermissionsOk() (*[]ConsentPermission, bool)`

GetAdditionalPermissionsOk returns a tuple with the AdditionalPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPermissions

`func (o *ChallengeTypes) SetAdditionalPermissions(v []ConsentPermission)`

SetAdditionalPermissions sets AdditionalPermissions field to given value.


### GetToken

`func (o *ChallengeTypes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ChallengeTypes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ChallengeTypes) SetToken(v string)`

SetToken sets Token field to given value.


### GetRequestId

`func (o *ChallengeTypes) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ChallengeTypes) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ChallengeTypes) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetError

`func (o *ChallengeTypes) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ChallengeTypes) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ChallengeTypes) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ChallengeTypes) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTraceback

`func (o *ChallengeTypes) GetTraceback() string`

GetTraceback returns the Traceback field if non-nil, zero value otherwise.

### GetTracebackOk

`func (o *ChallengeTypes) GetTracebackOk() (*string, bool)`

GetTracebackOk returns a tuple with the Traceback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceback

`func (o *ChallengeTypes) SetTraceback(v string)`

SetTraceback sets Traceback field to given value.

### HasTraceback

`func (o *ChallengeTypes) HasTraceback() bool`

HasTraceback returns a boolean if a field has been set.

### GetUserFields

`func (o *ChallengeTypes) GetUserFields() []string`

GetUserFields returns the UserFields field if non-nil, zero value otherwise.

### GetUserFieldsOk

`func (o *ChallengeTypes) GetUserFieldsOk() (*[]string, bool)`

GetUserFieldsOk returns a tuple with the UserFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFields

`func (o *ChallengeTypes) SetUserFields(v []string)`

SetUserFields sets UserFields field to given value.


### SetUserFieldsNil

`func (o *ChallengeTypes) SetUserFieldsNil(b bool)`

 SetUserFieldsNil sets the value for UserFields to be an explicit nil

### UnsetUserFields
`func (o *ChallengeTypes) UnsetUserFields()`

UnsetUserFields ensures that no value is present for UserFields, not even an explicit nil
### GetPasswordFields

`func (o *ChallengeTypes) GetPasswordFields() bool`

GetPasswordFields returns the PasswordFields field if non-nil, zero value otherwise.

### GetPasswordFieldsOk

`func (o *ChallengeTypes) GetPasswordFieldsOk() (*bool, bool)`

GetPasswordFieldsOk returns a tuple with the PasswordFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFields

`func (o *ChallengeTypes) SetPasswordFields(v bool)`

SetPasswordFields sets PasswordFields field to given value.


### GetApplicationPre

`func (o *ChallengeTypes) GetApplicationPre() string`

GetApplicationPre returns the ApplicationPre field if non-nil, zero value otherwise.

### GetApplicationPreOk

`func (o *ChallengeTypes) GetApplicationPreOk() (*string, bool)`

GetApplicationPreOk returns a tuple with the ApplicationPre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPre

`func (o *ChallengeTypes) SetApplicationPre(v string)`

SetApplicationPre sets ApplicationPre field to given value.

### HasApplicationPre

`func (o *ChallengeTypes) HasApplicationPre() bool`

HasApplicationPre returns a boolean if a field has been set.

### GetEnrollUrl

`func (o *ChallengeTypes) GetEnrollUrl() string`

GetEnrollUrl returns the EnrollUrl field if non-nil, zero value otherwise.

### GetEnrollUrlOk

`func (o *ChallengeTypes) GetEnrollUrlOk() (*string, bool)`

GetEnrollUrlOk returns a tuple with the EnrollUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollUrl

`func (o *ChallengeTypes) SetEnrollUrl(v string)`

SetEnrollUrl sets EnrollUrl field to given value.

### HasEnrollUrl

`func (o *ChallengeTypes) HasEnrollUrl() bool`

HasEnrollUrl returns a boolean if a field has been set.

### GetRecoveryUrl

`func (o *ChallengeTypes) GetRecoveryUrl() string`

GetRecoveryUrl returns the RecoveryUrl field if non-nil, zero value otherwise.

### GetRecoveryUrlOk

`func (o *ChallengeTypes) GetRecoveryUrlOk() (*string, bool)`

GetRecoveryUrlOk returns a tuple with the RecoveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryUrl

`func (o *ChallengeTypes) SetRecoveryUrl(v string)`

SetRecoveryUrl sets RecoveryUrl field to given value.

### HasRecoveryUrl

`func (o *ChallengeTypes) HasRecoveryUrl() bool`

HasRecoveryUrl returns a boolean if a field has been set.

### GetPasswordlessUrl

`func (o *ChallengeTypes) GetPasswordlessUrl() string`

GetPasswordlessUrl returns the PasswordlessUrl field if non-nil, zero value otherwise.

### GetPasswordlessUrlOk

`func (o *ChallengeTypes) GetPasswordlessUrlOk() (*string, bool)`

GetPasswordlessUrlOk returns a tuple with the PasswordlessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordlessUrl

`func (o *ChallengeTypes) SetPasswordlessUrl(v string)`

SetPasswordlessUrl sets PasswordlessUrl field to given value.

### HasPasswordlessUrl

`func (o *ChallengeTypes) HasPasswordlessUrl() bool`

HasPasswordlessUrl returns a boolean if a field has been set.

### GetPrimaryAction

`func (o *ChallengeTypes) GetPrimaryAction() string`

GetPrimaryAction returns the PrimaryAction field if non-nil, zero value otherwise.

### GetPrimaryActionOk

`func (o *ChallengeTypes) GetPrimaryActionOk() (*string, bool)`

GetPrimaryActionOk returns a tuple with the PrimaryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAction

`func (o *ChallengeTypes) SetPrimaryAction(v string)`

SetPrimaryAction sets PrimaryAction field to given value.


### GetSources

`func (o *ChallengeTypes) GetSources() []LoginSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ChallengeTypes) GetSourcesOk() (*[]LoginSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ChallengeTypes) SetSources(v []LoginSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *ChallengeTypes) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetShowSourceLabels

`func (o *ChallengeTypes) GetShowSourceLabels() bool`

GetShowSourceLabels returns the ShowSourceLabels field if non-nil, zero value otherwise.

### GetShowSourceLabelsOk

`func (o *ChallengeTypes) GetShowSourceLabelsOk() (*bool, bool)`

GetShowSourceLabelsOk returns a tuple with the ShowSourceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSourceLabels

`func (o *ChallengeTypes) SetShowSourceLabels(v bool)`

SetShowSourceLabels sets ShowSourceLabels field to given value.


### GetSlug

`func (o *ChallengeTypes) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ChallengeTypes) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ChallengeTypes) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetFields

`func (o *ChallengeTypes) GetFields() []StagePrompt`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ChallengeTypes) GetFieldsOk() (*[]StagePrompt, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ChallengeTypes) SetFields(v []StagePrompt)`

SetFields sets Fields field to given value.


### GetTo

`func (o *ChallengeTypes) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ChallengeTypes) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ChallengeTypes) SetTo(v string)`

SetTo sets To field to given value.


### GetBody

`func (o *ChallengeTypes) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ChallengeTypes) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ChallengeTypes) SetBody(v string)`

SetBody sets Body field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


