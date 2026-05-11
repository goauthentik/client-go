# PatchedTelegramSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Source&#39;s display Name. | [optional] 
**Slug** | Pointer to **string** | Internal source name, used in URLs. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Promoted** | Pointer to **bool** | When enabled, this source will be displayed as a prominent button on the login page, instead of a small icon. | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** | Flow to use when authenticating existing users. | [optional] 
**EnrollmentFlow** | Pointer to **NullableString** | Flow to use when enrolling new users. | [optional] 
**UserPropertyMappings** | Pointer to **[]string** |  | [optional] 
**GroupPropertyMappings** | Pointer to **[]string** |  | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**UserMatchingMode** | Pointer to [**UserMatchingModeEnum**](UserMatchingModeEnum.md) | How the source determines if an existing user should be authenticated or a new user enrolled. | [optional] 
**UserPathTemplate** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**BotUsername** | Pointer to **string** | Telegram bot username | [optional] 
**BotToken** | Pointer to **string** | Telegram bot token | [optional] 
**RequestMessageAccess** | Pointer to **bool** | Request access to send messages from your bot. | [optional] 
**PreAuthenticationFlow** | Pointer to **string** | Flow used before authentication. | [optional] 

## Methods

### NewPatchedTelegramSourceRequest

`func NewPatchedTelegramSourceRequest() *PatchedTelegramSourceRequest`

NewPatchedTelegramSourceRequest instantiates a new PatchedTelegramSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTelegramSourceRequestWithDefaults

`func NewPatchedTelegramSourceRequestWithDefaults() *PatchedTelegramSourceRequest`

NewPatchedTelegramSourceRequestWithDefaults instantiates a new PatchedTelegramSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedTelegramSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedTelegramSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedTelegramSourceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedTelegramSourceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedTelegramSourceRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedTelegramSourceRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedTelegramSourceRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedTelegramSourceRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedTelegramSourceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedTelegramSourceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedTelegramSourceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedTelegramSourceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPromoted

`func (o *PatchedTelegramSourceRequest) GetPromoted() bool`

GetPromoted returns the Promoted field if non-nil, zero value otherwise.

### GetPromotedOk

`func (o *PatchedTelegramSourceRequest) GetPromotedOk() (*bool, bool)`

GetPromotedOk returns a tuple with the Promoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoted

`func (o *PatchedTelegramSourceRequest) SetPromoted(v bool)`

SetPromoted sets Promoted field to given value.

### HasPromoted

`func (o *PatchedTelegramSourceRequest) HasPromoted() bool`

HasPromoted returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *PatchedTelegramSourceRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *PatchedTelegramSourceRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *PatchedTelegramSourceRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *PatchedTelegramSourceRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *PatchedTelegramSourceRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *PatchedTelegramSourceRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetEnrollmentFlow

`func (o *PatchedTelegramSourceRequest) GetEnrollmentFlow() string`

GetEnrollmentFlow returns the EnrollmentFlow field if non-nil, zero value otherwise.

### GetEnrollmentFlowOk

`func (o *PatchedTelegramSourceRequest) GetEnrollmentFlowOk() (*string, bool)`

GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlow

`func (o *PatchedTelegramSourceRequest) SetEnrollmentFlow(v string)`

SetEnrollmentFlow sets EnrollmentFlow field to given value.

### HasEnrollmentFlow

`func (o *PatchedTelegramSourceRequest) HasEnrollmentFlow() bool`

HasEnrollmentFlow returns a boolean if a field has been set.

### SetEnrollmentFlowNil

`func (o *PatchedTelegramSourceRequest) SetEnrollmentFlowNil(b bool)`

 SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil

### UnsetEnrollmentFlow
`func (o *PatchedTelegramSourceRequest) UnsetEnrollmentFlow()`

UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
### GetUserPropertyMappings

`func (o *PatchedTelegramSourceRequest) GetUserPropertyMappings() []string`

GetUserPropertyMappings returns the UserPropertyMappings field if non-nil, zero value otherwise.

### GetUserPropertyMappingsOk

`func (o *PatchedTelegramSourceRequest) GetUserPropertyMappingsOk() (*[]string, bool)`

GetUserPropertyMappingsOk returns a tuple with the UserPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPropertyMappings

`func (o *PatchedTelegramSourceRequest) SetUserPropertyMappings(v []string)`

SetUserPropertyMappings sets UserPropertyMappings field to given value.

### HasUserPropertyMappings

`func (o *PatchedTelegramSourceRequest) HasUserPropertyMappings() bool`

HasUserPropertyMappings returns a boolean if a field has been set.

### GetGroupPropertyMappings

`func (o *PatchedTelegramSourceRequest) GetGroupPropertyMappings() []string`

GetGroupPropertyMappings returns the GroupPropertyMappings field if non-nil, zero value otherwise.

### GetGroupPropertyMappingsOk

`func (o *PatchedTelegramSourceRequest) GetGroupPropertyMappingsOk() (*[]string, bool)`

GetGroupPropertyMappingsOk returns a tuple with the GroupPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPropertyMappings

`func (o *PatchedTelegramSourceRequest) SetGroupPropertyMappings(v []string)`

SetGroupPropertyMappings sets GroupPropertyMappings field to given value.

### HasGroupPropertyMappings

`func (o *PatchedTelegramSourceRequest) HasGroupPropertyMappings() bool`

HasGroupPropertyMappings returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *PatchedTelegramSourceRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *PatchedTelegramSourceRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *PatchedTelegramSourceRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *PatchedTelegramSourceRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetUserMatchingMode

`func (o *PatchedTelegramSourceRequest) GetUserMatchingMode() UserMatchingModeEnum`

GetUserMatchingMode returns the UserMatchingMode field if non-nil, zero value otherwise.

### GetUserMatchingModeOk

`func (o *PatchedTelegramSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool)`

GetUserMatchingModeOk returns a tuple with the UserMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMatchingMode

`func (o *PatchedTelegramSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum)`

SetUserMatchingMode sets UserMatchingMode field to given value.

### HasUserMatchingMode

`func (o *PatchedTelegramSourceRequest) HasUserMatchingMode() bool`

HasUserMatchingMode returns a boolean if a field has been set.

### GetUserPathTemplate

`func (o *PatchedTelegramSourceRequest) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *PatchedTelegramSourceRequest) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *PatchedTelegramSourceRequest) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *PatchedTelegramSourceRequest) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.

### GetIcon

`func (o *PatchedTelegramSourceRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *PatchedTelegramSourceRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *PatchedTelegramSourceRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *PatchedTelegramSourceRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetBotUsername

`func (o *PatchedTelegramSourceRequest) GetBotUsername() string`

GetBotUsername returns the BotUsername field if non-nil, zero value otherwise.

### GetBotUsernameOk

`func (o *PatchedTelegramSourceRequest) GetBotUsernameOk() (*string, bool)`

GetBotUsernameOk returns a tuple with the BotUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotUsername

`func (o *PatchedTelegramSourceRequest) SetBotUsername(v string)`

SetBotUsername sets BotUsername field to given value.

### HasBotUsername

`func (o *PatchedTelegramSourceRequest) HasBotUsername() bool`

HasBotUsername returns a boolean if a field has been set.

### GetBotToken

`func (o *PatchedTelegramSourceRequest) GetBotToken() string`

GetBotToken returns the BotToken field if non-nil, zero value otherwise.

### GetBotTokenOk

`func (o *PatchedTelegramSourceRequest) GetBotTokenOk() (*string, bool)`

GetBotTokenOk returns a tuple with the BotToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotToken

`func (o *PatchedTelegramSourceRequest) SetBotToken(v string)`

SetBotToken sets BotToken field to given value.

### HasBotToken

`func (o *PatchedTelegramSourceRequest) HasBotToken() bool`

HasBotToken returns a boolean if a field has been set.

### GetRequestMessageAccess

`func (o *PatchedTelegramSourceRequest) GetRequestMessageAccess() bool`

GetRequestMessageAccess returns the RequestMessageAccess field if non-nil, zero value otherwise.

### GetRequestMessageAccessOk

`func (o *PatchedTelegramSourceRequest) GetRequestMessageAccessOk() (*bool, bool)`

GetRequestMessageAccessOk returns a tuple with the RequestMessageAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMessageAccess

`func (o *PatchedTelegramSourceRequest) SetRequestMessageAccess(v bool)`

SetRequestMessageAccess sets RequestMessageAccess field to given value.

### HasRequestMessageAccess

`func (o *PatchedTelegramSourceRequest) HasRequestMessageAccess() bool`

HasRequestMessageAccess returns a boolean if a field has been set.

### GetPreAuthenticationFlow

`func (o *PatchedTelegramSourceRequest) GetPreAuthenticationFlow() string`

GetPreAuthenticationFlow returns the PreAuthenticationFlow field if non-nil, zero value otherwise.

### GetPreAuthenticationFlowOk

`func (o *PatchedTelegramSourceRequest) GetPreAuthenticationFlowOk() (*string, bool)`

GetPreAuthenticationFlowOk returns a tuple with the PreAuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthenticationFlow

`func (o *PatchedTelegramSourceRequest) SetPreAuthenticationFlow(v string)`

SetPreAuthenticationFlow sets PreAuthenticationFlow field to given value.

### HasPreAuthenticationFlow

`func (o *PatchedTelegramSourceRequest) HasPreAuthenticationFlow() bool`

HasPreAuthenticationFlow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


