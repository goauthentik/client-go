# AuthenticatorValidateStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**NotConfiguredAction** | Pointer to [**NotConfiguredActionEnum**](NotConfiguredActionEnum.md) |  | [optional] 
**DeviceClasses** | Pointer to [**[]DeviceClassesEnum**](DeviceClassesEnum.md) | Device classes which can be used to authenticate | [optional] 
**ConfigurationStages** | Pointer to **[]string** | Stages used to configure Authenticator when user doesn&#39;t have any compatible devices. After this configuration Stage passes, the user is not prompted again. | [optional] 
**LastAuthThreshold** | Pointer to **string** | If any of the user&#39;s device has been used within this threshold, this stage will be skipped | [optional] 
**WebauthnUserVerification** | Pointer to [**UserVerificationEnum**](UserVerificationEnum.md) | Enforce user verification for WebAuthn devices. | [optional] 
**WebauthnHints** | Pointer to [**[]WebAuthnHintEnum**](WebAuthnHintEnum.md) |  | [optional] 
**WebauthnAllowedDeviceTypes** | Pointer to **[]string** |  | [optional] 
**EmailOtpThrottlingFactor** | Pointer to **float64** |  | [optional] 
**SmsOtpThrottlingFactor** | Pointer to **float64** |  | [optional] 
**TotpOtpThrottlingFactor** | Pointer to **float64** |  | [optional] 
**StaticOtpThrottlingFactor** | Pointer to **float64** |  | [optional] 

## Methods

### NewAuthenticatorValidateStageRequest

`func NewAuthenticatorValidateStageRequest(name string, ) *AuthenticatorValidateStageRequest`

NewAuthenticatorValidateStageRequest instantiates a new AuthenticatorValidateStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorValidateStageRequestWithDefaults

`func NewAuthenticatorValidateStageRequestWithDefaults() *AuthenticatorValidateStageRequest`

NewAuthenticatorValidateStageRequestWithDefaults instantiates a new AuthenticatorValidateStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthenticatorValidateStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticatorValidateStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticatorValidateStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNotConfiguredAction

`func (o *AuthenticatorValidateStageRequest) GetNotConfiguredAction() NotConfiguredActionEnum`

GetNotConfiguredAction returns the NotConfiguredAction field if non-nil, zero value otherwise.

### GetNotConfiguredActionOk

`func (o *AuthenticatorValidateStageRequest) GetNotConfiguredActionOk() (*NotConfiguredActionEnum, bool)`

GetNotConfiguredActionOk returns a tuple with the NotConfiguredAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotConfiguredAction

`func (o *AuthenticatorValidateStageRequest) SetNotConfiguredAction(v NotConfiguredActionEnum)`

SetNotConfiguredAction sets NotConfiguredAction field to given value.

### HasNotConfiguredAction

`func (o *AuthenticatorValidateStageRequest) HasNotConfiguredAction() bool`

HasNotConfiguredAction returns a boolean if a field has been set.

### GetDeviceClasses

`func (o *AuthenticatorValidateStageRequest) GetDeviceClasses() []DeviceClassesEnum`

GetDeviceClasses returns the DeviceClasses field if non-nil, zero value otherwise.

### GetDeviceClassesOk

`func (o *AuthenticatorValidateStageRequest) GetDeviceClassesOk() (*[]DeviceClassesEnum, bool)`

GetDeviceClassesOk returns a tuple with the DeviceClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClasses

`func (o *AuthenticatorValidateStageRequest) SetDeviceClasses(v []DeviceClassesEnum)`

SetDeviceClasses sets DeviceClasses field to given value.

### HasDeviceClasses

`func (o *AuthenticatorValidateStageRequest) HasDeviceClasses() bool`

HasDeviceClasses returns a boolean if a field has been set.

### GetConfigurationStages

`func (o *AuthenticatorValidateStageRequest) GetConfigurationStages() []string`

GetConfigurationStages returns the ConfigurationStages field if non-nil, zero value otherwise.

### GetConfigurationStagesOk

`func (o *AuthenticatorValidateStageRequest) GetConfigurationStagesOk() (*[]string, bool)`

GetConfigurationStagesOk returns a tuple with the ConfigurationStages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationStages

`func (o *AuthenticatorValidateStageRequest) SetConfigurationStages(v []string)`

SetConfigurationStages sets ConfigurationStages field to given value.

### HasConfigurationStages

`func (o *AuthenticatorValidateStageRequest) HasConfigurationStages() bool`

HasConfigurationStages returns a boolean if a field has been set.

### GetLastAuthThreshold

`func (o *AuthenticatorValidateStageRequest) GetLastAuthThreshold() string`

GetLastAuthThreshold returns the LastAuthThreshold field if non-nil, zero value otherwise.

### GetLastAuthThresholdOk

`func (o *AuthenticatorValidateStageRequest) GetLastAuthThresholdOk() (*string, bool)`

GetLastAuthThresholdOk returns a tuple with the LastAuthThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuthThreshold

`func (o *AuthenticatorValidateStageRequest) SetLastAuthThreshold(v string)`

SetLastAuthThreshold sets LastAuthThreshold field to given value.

### HasLastAuthThreshold

`func (o *AuthenticatorValidateStageRequest) HasLastAuthThreshold() bool`

HasLastAuthThreshold returns a boolean if a field has been set.

### GetWebauthnUserVerification

`func (o *AuthenticatorValidateStageRequest) GetWebauthnUserVerification() UserVerificationEnum`

GetWebauthnUserVerification returns the WebauthnUserVerification field if non-nil, zero value otherwise.

### GetWebauthnUserVerificationOk

`func (o *AuthenticatorValidateStageRequest) GetWebauthnUserVerificationOk() (*UserVerificationEnum, bool)`

GetWebauthnUserVerificationOk returns a tuple with the WebauthnUserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebauthnUserVerification

`func (o *AuthenticatorValidateStageRequest) SetWebauthnUserVerification(v UserVerificationEnum)`

SetWebauthnUserVerification sets WebauthnUserVerification field to given value.

### HasWebauthnUserVerification

`func (o *AuthenticatorValidateStageRequest) HasWebauthnUserVerification() bool`

HasWebauthnUserVerification returns a boolean if a field has been set.

### GetWebauthnHints

`func (o *AuthenticatorValidateStageRequest) GetWebauthnHints() []WebAuthnHintEnum`

GetWebauthnHints returns the WebauthnHints field if non-nil, zero value otherwise.

### GetWebauthnHintsOk

`func (o *AuthenticatorValidateStageRequest) GetWebauthnHintsOk() (*[]WebAuthnHintEnum, bool)`

GetWebauthnHintsOk returns a tuple with the WebauthnHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebauthnHints

`func (o *AuthenticatorValidateStageRequest) SetWebauthnHints(v []WebAuthnHintEnum)`

SetWebauthnHints sets WebauthnHints field to given value.

### HasWebauthnHints

`func (o *AuthenticatorValidateStageRequest) HasWebauthnHints() bool`

HasWebauthnHints returns a boolean if a field has been set.

### GetWebauthnAllowedDeviceTypes

`func (o *AuthenticatorValidateStageRequest) GetWebauthnAllowedDeviceTypes() []string`

GetWebauthnAllowedDeviceTypes returns the WebauthnAllowedDeviceTypes field if non-nil, zero value otherwise.

### GetWebauthnAllowedDeviceTypesOk

`func (o *AuthenticatorValidateStageRequest) GetWebauthnAllowedDeviceTypesOk() (*[]string, bool)`

GetWebauthnAllowedDeviceTypesOk returns a tuple with the WebauthnAllowedDeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebauthnAllowedDeviceTypes

`func (o *AuthenticatorValidateStageRequest) SetWebauthnAllowedDeviceTypes(v []string)`

SetWebauthnAllowedDeviceTypes sets WebauthnAllowedDeviceTypes field to given value.

### HasWebauthnAllowedDeviceTypes

`func (o *AuthenticatorValidateStageRequest) HasWebauthnAllowedDeviceTypes() bool`

HasWebauthnAllowedDeviceTypes returns a boolean if a field has been set.

### GetEmailOtpThrottlingFactor

`func (o *AuthenticatorValidateStageRequest) GetEmailOtpThrottlingFactor() float64`

GetEmailOtpThrottlingFactor returns the EmailOtpThrottlingFactor field if non-nil, zero value otherwise.

### GetEmailOtpThrottlingFactorOk

`func (o *AuthenticatorValidateStageRequest) GetEmailOtpThrottlingFactorOk() (*float64, bool)`

GetEmailOtpThrottlingFactorOk returns a tuple with the EmailOtpThrottlingFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOtpThrottlingFactor

`func (o *AuthenticatorValidateStageRequest) SetEmailOtpThrottlingFactor(v float64)`

SetEmailOtpThrottlingFactor sets EmailOtpThrottlingFactor field to given value.

### HasEmailOtpThrottlingFactor

`func (o *AuthenticatorValidateStageRequest) HasEmailOtpThrottlingFactor() bool`

HasEmailOtpThrottlingFactor returns a boolean if a field has been set.

### GetSmsOtpThrottlingFactor

`func (o *AuthenticatorValidateStageRequest) GetSmsOtpThrottlingFactor() float64`

GetSmsOtpThrottlingFactor returns the SmsOtpThrottlingFactor field if non-nil, zero value otherwise.

### GetSmsOtpThrottlingFactorOk

`func (o *AuthenticatorValidateStageRequest) GetSmsOtpThrottlingFactorOk() (*float64, bool)`

GetSmsOtpThrottlingFactorOk returns a tuple with the SmsOtpThrottlingFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsOtpThrottlingFactor

`func (o *AuthenticatorValidateStageRequest) SetSmsOtpThrottlingFactor(v float64)`

SetSmsOtpThrottlingFactor sets SmsOtpThrottlingFactor field to given value.

### HasSmsOtpThrottlingFactor

`func (o *AuthenticatorValidateStageRequest) HasSmsOtpThrottlingFactor() bool`

HasSmsOtpThrottlingFactor returns a boolean if a field has been set.

### GetTotpOtpThrottlingFactor

`func (o *AuthenticatorValidateStageRequest) GetTotpOtpThrottlingFactor() float64`

GetTotpOtpThrottlingFactor returns the TotpOtpThrottlingFactor field if non-nil, zero value otherwise.

### GetTotpOtpThrottlingFactorOk

`func (o *AuthenticatorValidateStageRequest) GetTotpOtpThrottlingFactorOk() (*float64, bool)`

GetTotpOtpThrottlingFactorOk returns a tuple with the TotpOtpThrottlingFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpOtpThrottlingFactor

`func (o *AuthenticatorValidateStageRequest) SetTotpOtpThrottlingFactor(v float64)`

SetTotpOtpThrottlingFactor sets TotpOtpThrottlingFactor field to given value.

### HasTotpOtpThrottlingFactor

`func (o *AuthenticatorValidateStageRequest) HasTotpOtpThrottlingFactor() bool`

HasTotpOtpThrottlingFactor returns a boolean if a field has been set.

### GetStaticOtpThrottlingFactor

`func (o *AuthenticatorValidateStageRequest) GetStaticOtpThrottlingFactor() float64`

GetStaticOtpThrottlingFactor returns the StaticOtpThrottlingFactor field if non-nil, zero value otherwise.

### GetStaticOtpThrottlingFactorOk

`func (o *AuthenticatorValidateStageRequest) GetStaticOtpThrottlingFactorOk() (*float64, bool)`

GetStaticOtpThrottlingFactorOk returns a tuple with the StaticOtpThrottlingFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticOtpThrottlingFactor

`func (o *AuthenticatorValidateStageRequest) SetStaticOtpThrottlingFactor(v float64)`

SetStaticOtpThrottlingFactor sets StaticOtpThrottlingFactor field to given value.

### HasStaticOtpThrottlingFactor

`func (o *AuthenticatorValidateStageRequest) HasStaticOtpThrottlingFactor() bool`

HasStaticOtpThrottlingFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


