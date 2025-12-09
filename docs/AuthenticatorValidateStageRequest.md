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
**WebauthnAllowedDeviceTypes** | Pointer to **[]string** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


