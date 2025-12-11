# ChallengeTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Email** | Pointer to **NullableString** |  | [optional] 
**EmailRequired** | Pointer to **bool** |  | [optional] [default to true]
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
**Interactive** | **bool** |  | 
**HeaderText** | Pointer to **string** |  | [optional] 
**Permissions** | [**[]ConsentPermission**](ConsentPermission.md) |  | 
**AdditionalPermissions** | [**[]ConsentPermission**](ConsentPermission.md) |  | 
**Token** | **string** |  | 
**Name** | **string** |  | 
**Challenge** | **string** |  | 
**ChallengeIdleTimeout** | **int32** |  | 
**RequestId** | **string** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Traceback** | Pointer to **string** |  | [optional] 
**LoadingOverlay** | Pointer to **bool** |  | [optional] [default to false]
**LoadingText** | **string** |  | 
**UserFields** | **[]string** |  | 
**PasswordFields** | **bool** |  | 
**AllowShowPassword** | Pointer to **bool** |  | [optional] [default to false]
**ApplicationPre** | Pointer to **string** |  | [optional] 
**FlowDesignation** | [**FlowDesignationEnum**](FlowDesignationEnum.md) |  | 
**CaptchaStage** | Pointer to [**NullableIdentificationChallengeCaptchaStage**](IdentificationChallengeCaptchaStage.md) |  | [optional] 
**EnrollUrl** | Pointer to **string** |  | [optional] 
**RecoveryUrl** | Pointer to **string** |  | [optional] 
**PasswordlessUrl** | Pointer to **string** |  | [optional] 
**PrimaryAction** | **string** |  | 
**Sources** | Pointer to [**[]LoginSource**](LoginSource.md) |  | [optional] 
**ShowSourceLabels** | **bool** |  | 
**EnableRememberMe** | Pointer to **bool** |  | [optional] [default to true]
**PasskeyChallenge** | Pointer to **map[string]interface{}** |  | [optional] 
**LogoutUrls** | Pointer to **[]map[string]interface{}** |  | [optional] 
**PostUrl** | Pointer to **string** |  | [optional] 
**SamlRequest** | Pointer to **string** |  | [optional] 
**RelayState** | Pointer to **string** |  | [optional] 
**ProviderName** | Pointer to **string** |  | [optional] 
**Binding** | Pointer to **string** |  | [optional] 
**RedirectUrl** | Pointer to **string** |  | [optional] 
**IsComplete** | Pointer to **bool** |  | [optional] [default to false]
**Slug** | **string** |  | 
**Fields** | [**[]StagePrompt**](StagePrompt.md) |  | 
**To** | **string** |  | 
**ApplicationName** | Pointer to **string** |  | [optional] 
**ApplicationLaunchUrl** | Pointer to **string** |  | [optional] 
**InvalidationFlowUrl** | Pointer to **string** |  | [optional] 
**BrandName** | **string** |  | 
**Body** | **string** |  | 
**BotUsername** | **string** | Telegram bot username | 
**RequestMessageAccess** | **bool** |  | 

## Methods

### NewChallengeTypes

`func NewChallengeTypes(pendingUser string, pendingUserAvatar string, clientId string, scope string, redirectUri string, state string, activationBarcode string, activationCode string, stageUuid string, codes []string, configUrl string, deviceChallenges []DeviceChallenge, configurationStages []SelectableStage, registration map[string]interface{}, url string, attrs map[string]string, siteKey string, jsUrl string, interactive bool, permissions []ConsentPermission, additionalPermissions []ConsentPermission, token string, name string, challenge string, challengeIdleTimeout int32, requestId string, loadingText string, userFields []string, passwordFields bool, flowDesignation FlowDesignationEnum, primaryAction string, showSourceLabels bool, slug string, fields []StagePrompt, to string, brandName string, body string, botUsername string, requestMessageAccess bool, ) *ChallengeTypes`

NewChallengeTypes instantiates a new ChallengeTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeTypesWithDefaults

`func NewChallengeTypesWithDefaults() *ChallengeTypes`

NewChallengeTypesWithDefaults instantiates a new ChallengeTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetEmail

`func (o *ChallengeTypes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ChallengeTypes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ChallengeTypes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ChallengeTypes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ChallengeTypes) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ChallengeTypes) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailRequired

`func (o *ChallengeTypes) GetEmailRequired() bool`

GetEmailRequired returns the EmailRequired field if non-nil, zero value otherwise.

### GetEmailRequiredOk

`func (o *ChallengeTypes) GetEmailRequiredOk() (*bool, bool)`

GetEmailRequiredOk returns a tuple with the EmailRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRequired

`func (o *ChallengeTypes) SetEmailRequired(v bool)`

SetEmailRequired sets EmailRequired field to given value.

### HasEmailRequired

`func (o *ChallengeTypes) HasEmailRequired() bool`

HasEmailRequired returns a boolean if a field has been set.

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


### GetInteractive

`func (o *ChallengeTypes) GetInteractive() bool`

GetInteractive returns the Interactive field if non-nil, zero value otherwise.

### GetInteractiveOk

`func (o *ChallengeTypes) GetInteractiveOk() (*bool, bool)`

GetInteractiveOk returns a tuple with the Interactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractive

`func (o *ChallengeTypes) SetInteractive(v bool)`

SetInteractive sets Interactive field to given value.


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


### GetName

`func (o *ChallengeTypes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChallengeTypes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChallengeTypes) SetName(v string)`

SetName sets Name field to given value.


### GetChallenge

`func (o *ChallengeTypes) GetChallenge() string`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *ChallengeTypes) GetChallengeOk() (*string, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *ChallengeTypes) SetChallenge(v string)`

SetChallenge sets Challenge field to given value.


### GetChallengeIdleTimeout

`func (o *ChallengeTypes) GetChallengeIdleTimeout() int32`

GetChallengeIdleTimeout returns the ChallengeIdleTimeout field if non-nil, zero value otherwise.

### GetChallengeIdleTimeoutOk

`func (o *ChallengeTypes) GetChallengeIdleTimeoutOk() (*int32, bool)`

GetChallengeIdleTimeoutOk returns a tuple with the ChallengeIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeIdleTimeout

`func (o *ChallengeTypes) SetChallengeIdleTimeout(v int32)`

SetChallengeIdleTimeout sets ChallengeIdleTimeout field to given value.


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

### GetLoadingOverlay

`func (o *ChallengeTypes) GetLoadingOverlay() bool`

GetLoadingOverlay returns the LoadingOverlay field if non-nil, zero value otherwise.

### GetLoadingOverlayOk

`func (o *ChallengeTypes) GetLoadingOverlayOk() (*bool, bool)`

GetLoadingOverlayOk returns a tuple with the LoadingOverlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadingOverlay

`func (o *ChallengeTypes) SetLoadingOverlay(v bool)`

SetLoadingOverlay sets LoadingOverlay field to given value.

### HasLoadingOverlay

`func (o *ChallengeTypes) HasLoadingOverlay() bool`

HasLoadingOverlay returns a boolean if a field has been set.

### GetLoadingText

`func (o *ChallengeTypes) GetLoadingText() string`

GetLoadingText returns the LoadingText field if non-nil, zero value otherwise.

### GetLoadingTextOk

`func (o *ChallengeTypes) GetLoadingTextOk() (*string, bool)`

GetLoadingTextOk returns a tuple with the LoadingText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadingText

`func (o *ChallengeTypes) SetLoadingText(v string)`

SetLoadingText sets LoadingText field to given value.


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


### GetAllowShowPassword

`func (o *ChallengeTypes) GetAllowShowPassword() bool`

GetAllowShowPassword returns the AllowShowPassword field if non-nil, zero value otherwise.

### GetAllowShowPasswordOk

`func (o *ChallengeTypes) GetAllowShowPasswordOk() (*bool, bool)`

GetAllowShowPasswordOk returns a tuple with the AllowShowPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowShowPassword

`func (o *ChallengeTypes) SetAllowShowPassword(v bool)`

SetAllowShowPassword sets AllowShowPassword field to given value.

### HasAllowShowPassword

`func (o *ChallengeTypes) HasAllowShowPassword() bool`

HasAllowShowPassword returns a boolean if a field has been set.

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

### GetFlowDesignation

`func (o *ChallengeTypes) GetFlowDesignation() FlowDesignationEnum`

GetFlowDesignation returns the FlowDesignation field if non-nil, zero value otherwise.

### GetFlowDesignationOk

`func (o *ChallengeTypes) GetFlowDesignationOk() (*FlowDesignationEnum, bool)`

GetFlowDesignationOk returns a tuple with the FlowDesignation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDesignation

`func (o *ChallengeTypes) SetFlowDesignation(v FlowDesignationEnum)`

SetFlowDesignation sets FlowDesignation field to given value.


### GetCaptchaStage

`func (o *ChallengeTypes) GetCaptchaStage() IdentificationChallengeCaptchaStage`

GetCaptchaStage returns the CaptchaStage field if non-nil, zero value otherwise.

### GetCaptchaStageOk

`func (o *ChallengeTypes) GetCaptchaStageOk() (*IdentificationChallengeCaptchaStage, bool)`

GetCaptchaStageOk returns a tuple with the CaptchaStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaStage

`func (o *ChallengeTypes) SetCaptchaStage(v IdentificationChallengeCaptchaStage)`

SetCaptchaStage sets CaptchaStage field to given value.

### HasCaptchaStage

`func (o *ChallengeTypes) HasCaptchaStage() bool`

HasCaptchaStage returns a boolean if a field has been set.

### SetCaptchaStageNil

`func (o *ChallengeTypes) SetCaptchaStageNil(b bool)`

 SetCaptchaStageNil sets the value for CaptchaStage to be an explicit nil

### UnsetCaptchaStage
`func (o *ChallengeTypes) UnsetCaptchaStage()`

UnsetCaptchaStage ensures that no value is present for CaptchaStage, not even an explicit nil
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


### GetEnableRememberMe

`func (o *ChallengeTypes) GetEnableRememberMe() bool`

GetEnableRememberMe returns the EnableRememberMe field if non-nil, zero value otherwise.

### GetEnableRememberMeOk

`func (o *ChallengeTypes) GetEnableRememberMeOk() (*bool, bool)`

GetEnableRememberMeOk returns a tuple with the EnableRememberMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRememberMe

`func (o *ChallengeTypes) SetEnableRememberMe(v bool)`

SetEnableRememberMe sets EnableRememberMe field to given value.

### HasEnableRememberMe

`func (o *ChallengeTypes) HasEnableRememberMe() bool`

HasEnableRememberMe returns a boolean if a field has been set.

### GetPasskeyChallenge

`func (o *ChallengeTypes) GetPasskeyChallenge() map[string]interface{}`

GetPasskeyChallenge returns the PasskeyChallenge field if non-nil, zero value otherwise.

### GetPasskeyChallengeOk

`func (o *ChallengeTypes) GetPasskeyChallengeOk() (*map[string]interface{}, bool)`

GetPasskeyChallengeOk returns a tuple with the PasskeyChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasskeyChallenge

`func (o *ChallengeTypes) SetPasskeyChallenge(v map[string]interface{})`

SetPasskeyChallenge sets PasskeyChallenge field to given value.

### HasPasskeyChallenge

`func (o *ChallengeTypes) HasPasskeyChallenge() bool`

HasPasskeyChallenge returns a boolean if a field has been set.

### SetPasskeyChallengeNil

`func (o *ChallengeTypes) SetPasskeyChallengeNil(b bool)`

 SetPasskeyChallengeNil sets the value for PasskeyChallenge to be an explicit nil

### UnsetPasskeyChallenge
`func (o *ChallengeTypes) UnsetPasskeyChallenge()`

UnsetPasskeyChallenge ensures that no value is present for PasskeyChallenge, not even an explicit nil
### GetLogoutUrls

`func (o *ChallengeTypes) GetLogoutUrls() []map[string]interface{}`

GetLogoutUrls returns the LogoutUrls field if non-nil, zero value otherwise.

### GetLogoutUrlsOk

`func (o *ChallengeTypes) GetLogoutUrlsOk() (*[]map[string]interface{}, bool)`

GetLogoutUrlsOk returns a tuple with the LogoutUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrls

`func (o *ChallengeTypes) SetLogoutUrls(v []map[string]interface{})`

SetLogoutUrls sets LogoutUrls field to given value.

### HasLogoutUrls

`func (o *ChallengeTypes) HasLogoutUrls() bool`

HasLogoutUrls returns a boolean if a field has been set.

### GetPostUrl

`func (o *ChallengeTypes) GetPostUrl() string`

GetPostUrl returns the PostUrl field if non-nil, zero value otherwise.

### GetPostUrlOk

`func (o *ChallengeTypes) GetPostUrlOk() (*string, bool)`

GetPostUrlOk returns a tuple with the PostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUrl

`func (o *ChallengeTypes) SetPostUrl(v string)`

SetPostUrl sets PostUrl field to given value.

### HasPostUrl

`func (o *ChallengeTypes) HasPostUrl() bool`

HasPostUrl returns a boolean if a field has been set.

### GetSamlRequest

`func (o *ChallengeTypes) GetSamlRequest() string`

GetSamlRequest returns the SamlRequest field if non-nil, zero value otherwise.

### GetSamlRequestOk

`func (o *ChallengeTypes) GetSamlRequestOk() (*string, bool)`

GetSamlRequestOk returns a tuple with the SamlRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlRequest

`func (o *ChallengeTypes) SetSamlRequest(v string)`

SetSamlRequest sets SamlRequest field to given value.

### HasSamlRequest

`func (o *ChallengeTypes) HasSamlRequest() bool`

HasSamlRequest returns a boolean if a field has been set.

### GetRelayState

`func (o *ChallengeTypes) GetRelayState() string`

GetRelayState returns the RelayState field if non-nil, zero value otherwise.

### GetRelayStateOk

`func (o *ChallengeTypes) GetRelayStateOk() (*string, bool)`

GetRelayStateOk returns a tuple with the RelayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayState

`func (o *ChallengeTypes) SetRelayState(v string)`

SetRelayState sets RelayState field to given value.

### HasRelayState

`func (o *ChallengeTypes) HasRelayState() bool`

HasRelayState returns a boolean if a field has been set.

### GetProviderName

`func (o *ChallengeTypes) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ChallengeTypes) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ChallengeTypes) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ChallengeTypes) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetBinding

`func (o *ChallengeTypes) GetBinding() string`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *ChallengeTypes) GetBindingOk() (*string, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *ChallengeTypes) SetBinding(v string)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *ChallengeTypes) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *ChallengeTypes) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *ChallengeTypes) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *ChallengeTypes) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *ChallengeTypes) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetIsComplete

`func (o *ChallengeTypes) GetIsComplete() bool`

GetIsComplete returns the IsComplete field if non-nil, zero value otherwise.

### GetIsCompleteOk

`func (o *ChallengeTypes) GetIsCompleteOk() (*bool, bool)`

GetIsCompleteOk returns a tuple with the IsComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComplete

`func (o *ChallengeTypes) SetIsComplete(v bool)`

SetIsComplete sets IsComplete field to given value.

### HasIsComplete

`func (o *ChallengeTypes) HasIsComplete() bool`

HasIsComplete returns a boolean if a field has been set.

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


### GetApplicationName

`func (o *ChallengeTypes) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *ChallengeTypes) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *ChallengeTypes) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *ChallengeTypes) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetApplicationLaunchUrl

`func (o *ChallengeTypes) GetApplicationLaunchUrl() string`

GetApplicationLaunchUrl returns the ApplicationLaunchUrl field if non-nil, zero value otherwise.

### GetApplicationLaunchUrlOk

`func (o *ChallengeTypes) GetApplicationLaunchUrlOk() (*string, bool)`

GetApplicationLaunchUrlOk returns a tuple with the ApplicationLaunchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationLaunchUrl

`func (o *ChallengeTypes) SetApplicationLaunchUrl(v string)`

SetApplicationLaunchUrl sets ApplicationLaunchUrl field to given value.

### HasApplicationLaunchUrl

`func (o *ChallengeTypes) HasApplicationLaunchUrl() bool`

HasApplicationLaunchUrl returns a boolean if a field has been set.

### GetInvalidationFlowUrl

`func (o *ChallengeTypes) GetInvalidationFlowUrl() string`

GetInvalidationFlowUrl returns the InvalidationFlowUrl field if non-nil, zero value otherwise.

### GetInvalidationFlowUrlOk

`func (o *ChallengeTypes) GetInvalidationFlowUrlOk() (*string, bool)`

GetInvalidationFlowUrlOk returns a tuple with the InvalidationFlowUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidationFlowUrl

`func (o *ChallengeTypes) SetInvalidationFlowUrl(v string)`

SetInvalidationFlowUrl sets InvalidationFlowUrl field to given value.

### HasInvalidationFlowUrl

`func (o *ChallengeTypes) HasInvalidationFlowUrl() bool`

HasInvalidationFlowUrl returns a boolean if a field has been set.

### GetBrandName

`func (o *ChallengeTypes) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *ChallengeTypes) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *ChallengeTypes) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.


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


### GetBotUsername

`func (o *ChallengeTypes) GetBotUsername() string`

GetBotUsername returns the BotUsername field if non-nil, zero value otherwise.

### GetBotUsernameOk

`func (o *ChallengeTypes) GetBotUsernameOk() (*string, bool)`

GetBotUsernameOk returns a tuple with the BotUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotUsername

`func (o *ChallengeTypes) SetBotUsername(v string)`

SetBotUsername sets BotUsername field to given value.


### GetRequestMessageAccess

`func (o *ChallengeTypes) GetRequestMessageAccess() bool`

GetRequestMessageAccess returns the RequestMessageAccess field if non-nil, zero value otherwise.

### GetRequestMessageAccessOk

`func (o *ChallengeTypes) GetRequestMessageAccessOk() (*bool, bool)`

GetRequestMessageAccessOk returns a tuple with the RequestMessageAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMessageAccess

`func (o *ChallengeTypes) SetRequestMessageAccess(v bool)`

SetRequestMessageAccess sets RequestMessageAccess field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


