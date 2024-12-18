# IdentificationChallengeCaptchaStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-captcha"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**PendingUser** | **string** |  | 
**PendingUserAvatar** | **string** |  | 
**SiteKey** | **string** |  | 
**JsUrl** | **string** |  | 
**Interactive** | **bool** |  | 

## Methods

### NewIdentificationChallengeCaptchaStage

`func NewIdentificationChallengeCaptchaStage(pendingUser string, pendingUserAvatar string, siteKey string, jsUrl string, interactive bool, ) *IdentificationChallengeCaptchaStage`

NewIdentificationChallengeCaptchaStage instantiates a new IdentificationChallengeCaptchaStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentificationChallengeCaptchaStageWithDefaults

`func NewIdentificationChallengeCaptchaStageWithDefaults() *IdentificationChallengeCaptchaStage`

NewIdentificationChallengeCaptchaStageWithDefaults instantiates a new IdentificationChallengeCaptchaStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowInfo

`func (o *IdentificationChallengeCaptchaStage) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *IdentificationChallengeCaptchaStage) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *IdentificationChallengeCaptchaStage) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *IdentificationChallengeCaptchaStage) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *IdentificationChallengeCaptchaStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *IdentificationChallengeCaptchaStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *IdentificationChallengeCaptchaStage) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *IdentificationChallengeCaptchaStage) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *IdentificationChallengeCaptchaStage) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *IdentificationChallengeCaptchaStage) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *IdentificationChallengeCaptchaStage) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *IdentificationChallengeCaptchaStage) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetPendingUser

`func (o *IdentificationChallengeCaptchaStage) GetPendingUser() string`

GetPendingUser returns the PendingUser field if non-nil, zero value otherwise.

### GetPendingUserOk

`func (o *IdentificationChallengeCaptchaStage) GetPendingUserOk() (*string, bool)`

GetPendingUserOk returns a tuple with the PendingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUser

`func (o *IdentificationChallengeCaptchaStage) SetPendingUser(v string)`

SetPendingUser sets PendingUser field to given value.


### GetPendingUserAvatar

`func (o *IdentificationChallengeCaptchaStage) GetPendingUserAvatar() string`

GetPendingUserAvatar returns the PendingUserAvatar field if non-nil, zero value otherwise.

### GetPendingUserAvatarOk

`func (o *IdentificationChallengeCaptchaStage) GetPendingUserAvatarOk() (*string, bool)`

GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUserAvatar

`func (o *IdentificationChallengeCaptchaStage) SetPendingUserAvatar(v string)`

SetPendingUserAvatar sets PendingUserAvatar field to given value.


### GetSiteKey

`func (o *IdentificationChallengeCaptchaStage) GetSiteKey() string`

GetSiteKey returns the SiteKey field if non-nil, zero value otherwise.

### GetSiteKeyOk

`func (o *IdentificationChallengeCaptchaStage) GetSiteKeyOk() (*string, bool)`

GetSiteKeyOk returns a tuple with the SiteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteKey

`func (o *IdentificationChallengeCaptchaStage) SetSiteKey(v string)`

SetSiteKey sets SiteKey field to given value.


### GetJsUrl

`func (o *IdentificationChallengeCaptchaStage) GetJsUrl() string`

GetJsUrl returns the JsUrl field if non-nil, zero value otherwise.

### GetJsUrlOk

`func (o *IdentificationChallengeCaptchaStage) GetJsUrlOk() (*string, bool)`

GetJsUrlOk returns a tuple with the JsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsUrl

`func (o *IdentificationChallengeCaptchaStage) SetJsUrl(v string)`

SetJsUrl sets JsUrl field to given value.


### GetInteractive

`func (o *IdentificationChallengeCaptchaStage) GetInteractive() bool`

GetInteractive returns the Interactive field if non-nil, zero value otherwise.

### GetInteractiveOk

`func (o *IdentificationChallengeCaptchaStage) GetInteractiveOk() (*bool, bool)`

GetInteractiveOk returns a tuple with the Interactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractive

`func (o *IdentificationChallengeCaptchaStage) SetInteractive(v bool)`

SetInteractive sets Interactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


