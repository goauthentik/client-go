# FlowChallengeResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** |  | [optional] [default to "ak-stage-user-login"]
**Code** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**SelectedChallenge** | Pointer to [**DeviceChallengeRequest**](DeviceChallengeRequest.md) |  | [optional] 
**SelectedStage** | Pointer to **string** |  | [optional] 
**Webauthn** | Pointer to **map[string]interface{}** |  | [optional] 
**Duo** | Pointer to **int32** |  | [optional] 
**Response** | **map[string]interface{}** |  | 
**Token** | **string** |  | 
**UidField** | **string** |  | 
**Password** | **string** |  | 
**CaptchaToken** | Pointer to **NullableString** |  | [optional] 
**To** | **string** |  | 
**RememberMe** | **bool** |  | 

## Methods

### NewFlowChallengeResponseRequest

`func NewFlowChallengeResponseRequest(code string, response map[string]interface{}, token string, uidField string, password string, to string, rememberMe bool, ) *FlowChallengeResponseRequest`

NewFlowChallengeResponseRequest instantiates a new FlowChallengeResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowChallengeResponseRequestWithDefaults

`func NewFlowChallengeResponseRequestWithDefaults() *FlowChallengeResponseRequest`

NewFlowChallengeResponseRequestWithDefaults instantiates a new FlowChallengeResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *FlowChallengeResponseRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *FlowChallengeResponseRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *FlowChallengeResponseRequest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *FlowChallengeResponseRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetCode

`func (o *FlowChallengeResponseRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FlowChallengeResponseRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FlowChallengeResponseRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetEmail

`func (o *FlowChallengeResponseRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FlowChallengeResponseRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FlowChallengeResponseRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FlowChallengeResponseRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *FlowChallengeResponseRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *FlowChallengeResponseRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *FlowChallengeResponseRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *FlowChallengeResponseRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetSelectedChallenge

`func (o *FlowChallengeResponseRequest) GetSelectedChallenge() DeviceChallengeRequest`

GetSelectedChallenge returns the SelectedChallenge field if non-nil, zero value otherwise.

### GetSelectedChallengeOk

`func (o *FlowChallengeResponseRequest) GetSelectedChallengeOk() (*DeviceChallengeRequest, bool)`

GetSelectedChallengeOk returns a tuple with the SelectedChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedChallenge

`func (o *FlowChallengeResponseRequest) SetSelectedChallenge(v DeviceChallengeRequest)`

SetSelectedChallenge sets SelectedChallenge field to given value.

### HasSelectedChallenge

`func (o *FlowChallengeResponseRequest) HasSelectedChallenge() bool`

HasSelectedChallenge returns a boolean if a field has been set.

### GetSelectedStage

`func (o *FlowChallengeResponseRequest) GetSelectedStage() string`

GetSelectedStage returns the SelectedStage field if non-nil, zero value otherwise.

### GetSelectedStageOk

`func (o *FlowChallengeResponseRequest) GetSelectedStageOk() (*string, bool)`

GetSelectedStageOk returns a tuple with the SelectedStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedStage

`func (o *FlowChallengeResponseRequest) SetSelectedStage(v string)`

SetSelectedStage sets SelectedStage field to given value.

### HasSelectedStage

`func (o *FlowChallengeResponseRequest) HasSelectedStage() bool`

HasSelectedStage returns a boolean if a field has been set.

### GetWebauthn

`func (o *FlowChallengeResponseRequest) GetWebauthn() map[string]interface{}`

GetWebauthn returns the Webauthn field if non-nil, zero value otherwise.

### GetWebauthnOk

`func (o *FlowChallengeResponseRequest) GetWebauthnOk() (*map[string]interface{}, bool)`

GetWebauthnOk returns a tuple with the Webauthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebauthn

`func (o *FlowChallengeResponseRequest) SetWebauthn(v map[string]interface{})`

SetWebauthn sets Webauthn field to given value.

### HasWebauthn

`func (o *FlowChallengeResponseRequest) HasWebauthn() bool`

HasWebauthn returns a boolean if a field has been set.

### GetDuo

`func (o *FlowChallengeResponseRequest) GetDuo() int32`

GetDuo returns the Duo field if non-nil, zero value otherwise.

### GetDuoOk

`func (o *FlowChallengeResponseRequest) GetDuoOk() (*int32, bool)`

GetDuoOk returns a tuple with the Duo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuo

`func (o *FlowChallengeResponseRequest) SetDuo(v int32)`

SetDuo sets Duo field to given value.

### HasDuo

`func (o *FlowChallengeResponseRequest) HasDuo() bool`

HasDuo returns a boolean if a field has been set.

### GetResponse

`func (o *FlowChallengeResponseRequest) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *FlowChallengeResponseRequest) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *FlowChallengeResponseRequest) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.


### GetToken

`func (o *FlowChallengeResponseRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FlowChallengeResponseRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FlowChallengeResponseRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUidField

`func (o *FlowChallengeResponseRequest) GetUidField() string`

GetUidField returns the UidField field if non-nil, zero value otherwise.

### GetUidFieldOk

`func (o *FlowChallengeResponseRequest) GetUidFieldOk() (*string, bool)`

GetUidFieldOk returns a tuple with the UidField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidField

`func (o *FlowChallengeResponseRequest) SetUidField(v string)`

SetUidField sets UidField field to given value.


### GetPassword

`func (o *FlowChallengeResponseRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FlowChallengeResponseRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FlowChallengeResponseRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetCaptchaToken

`func (o *FlowChallengeResponseRequest) GetCaptchaToken() string`

GetCaptchaToken returns the CaptchaToken field if non-nil, zero value otherwise.

### GetCaptchaTokenOk

`func (o *FlowChallengeResponseRequest) GetCaptchaTokenOk() (*string, bool)`

GetCaptchaTokenOk returns a tuple with the CaptchaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaToken

`func (o *FlowChallengeResponseRequest) SetCaptchaToken(v string)`

SetCaptchaToken sets CaptchaToken field to given value.

### HasCaptchaToken

`func (o *FlowChallengeResponseRequest) HasCaptchaToken() bool`

HasCaptchaToken returns a boolean if a field has been set.

### SetCaptchaTokenNil

`func (o *FlowChallengeResponseRequest) SetCaptchaTokenNil(b bool)`

 SetCaptchaTokenNil sets the value for CaptchaToken to be an explicit nil

### UnsetCaptchaToken
`func (o *FlowChallengeResponseRequest) UnsetCaptchaToken()`

UnsetCaptchaToken ensures that no value is present for CaptchaToken, not even an explicit nil
### GetTo

`func (o *FlowChallengeResponseRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *FlowChallengeResponseRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *FlowChallengeResponseRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetRememberMe

`func (o *FlowChallengeResponseRequest) GetRememberMe() bool`

GetRememberMe returns the RememberMe field if non-nil, zero value otherwise.

### GetRememberMeOk

`func (o *FlowChallengeResponseRequest) GetRememberMeOk() (*bool, bool)`

GetRememberMeOk returns a tuple with the RememberMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMe

`func (o *FlowChallengeResponseRequest) SetRememberMe(v bool)`

SetRememberMe sets RememberMe field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


