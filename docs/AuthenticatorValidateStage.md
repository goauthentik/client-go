# AuthenticatorValidateStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Component** | **string** | Get object type so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**FlowSet** | Pointer to [**[]FlowSet**](FlowSet.md) |  | [optional] 
**NotConfiguredAction** | Pointer to [**NotConfiguredActionEnum**](NotConfiguredActionEnum.md) |  | [optional] 
**DeviceClasses** | Pointer to [**[]DeviceClassesEnum**](DeviceClassesEnum.md) | Device classes which can be used to authenticate | [optional] 
**ConfigurationStages** | Pointer to **[]string** | Stages used to configure Authenticator when user doesn&#39;t have any compatible devices. After this configuration Stage passes, the user is not prompted again. | [optional] 
**LastAuthThreshold** | Pointer to **string** | If any of the user&#39;s device has been used within this threshold, this stage will be skipped | [optional] 
**WebauthnUserVerification** | Pointer to [**UserVerificationEnum**](UserVerificationEnum.md) |  | [optional] 

## Methods

### NewAuthenticatorValidateStage

`func NewAuthenticatorValidateStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, ) *AuthenticatorValidateStage`

NewAuthenticatorValidateStage instantiates a new AuthenticatorValidateStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorValidateStageWithDefaults

`func NewAuthenticatorValidateStageWithDefaults() *AuthenticatorValidateStage`

NewAuthenticatorValidateStageWithDefaults instantiates a new AuthenticatorValidateStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *AuthenticatorValidateStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *AuthenticatorValidateStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *AuthenticatorValidateStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *AuthenticatorValidateStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticatorValidateStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticatorValidateStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *AuthenticatorValidateStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AuthenticatorValidateStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AuthenticatorValidateStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *AuthenticatorValidateStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *AuthenticatorValidateStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *AuthenticatorValidateStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *AuthenticatorValidateStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *AuthenticatorValidateStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *AuthenticatorValidateStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *AuthenticatorValidateStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *AuthenticatorValidateStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *AuthenticatorValidateStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *AuthenticatorValidateStage) GetFlowSet() []FlowSet`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *AuthenticatorValidateStage) GetFlowSetOk() (*[]FlowSet, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *AuthenticatorValidateStage) SetFlowSet(v []FlowSet)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *AuthenticatorValidateStage) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetNotConfiguredAction

`func (o *AuthenticatorValidateStage) GetNotConfiguredAction() NotConfiguredActionEnum`

GetNotConfiguredAction returns the NotConfiguredAction field if non-nil, zero value otherwise.

### GetNotConfiguredActionOk

`func (o *AuthenticatorValidateStage) GetNotConfiguredActionOk() (*NotConfiguredActionEnum, bool)`

GetNotConfiguredActionOk returns a tuple with the NotConfiguredAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotConfiguredAction

`func (o *AuthenticatorValidateStage) SetNotConfiguredAction(v NotConfiguredActionEnum)`

SetNotConfiguredAction sets NotConfiguredAction field to given value.

### HasNotConfiguredAction

`func (o *AuthenticatorValidateStage) HasNotConfiguredAction() bool`

HasNotConfiguredAction returns a boolean if a field has been set.

### GetDeviceClasses

`func (o *AuthenticatorValidateStage) GetDeviceClasses() []DeviceClassesEnum`

GetDeviceClasses returns the DeviceClasses field if non-nil, zero value otherwise.

### GetDeviceClassesOk

`func (o *AuthenticatorValidateStage) GetDeviceClassesOk() (*[]DeviceClassesEnum, bool)`

GetDeviceClassesOk returns a tuple with the DeviceClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClasses

`func (o *AuthenticatorValidateStage) SetDeviceClasses(v []DeviceClassesEnum)`

SetDeviceClasses sets DeviceClasses field to given value.

### HasDeviceClasses

`func (o *AuthenticatorValidateStage) HasDeviceClasses() bool`

HasDeviceClasses returns a boolean if a field has been set.

### GetConfigurationStages

`func (o *AuthenticatorValidateStage) GetConfigurationStages() []string`

GetConfigurationStages returns the ConfigurationStages field if non-nil, zero value otherwise.

### GetConfigurationStagesOk

`func (o *AuthenticatorValidateStage) GetConfigurationStagesOk() (*[]string, bool)`

GetConfigurationStagesOk returns a tuple with the ConfigurationStages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationStages

`func (o *AuthenticatorValidateStage) SetConfigurationStages(v []string)`

SetConfigurationStages sets ConfigurationStages field to given value.

### HasConfigurationStages

`func (o *AuthenticatorValidateStage) HasConfigurationStages() bool`

HasConfigurationStages returns a boolean if a field has been set.

### GetLastAuthThreshold

`func (o *AuthenticatorValidateStage) GetLastAuthThreshold() string`

GetLastAuthThreshold returns the LastAuthThreshold field if non-nil, zero value otherwise.

### GetLastAuthThresholdOk

`func (o *AuthenticatorValidateStage) GetLastAuthThresholdOk() (*string, bool)`

GetLastAuthThresholdOk returns a tuple with the LastAuthThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuthThreshold

`func (o *AuthenticatorValidateStage) SetLastAuthThreshold(v string)`

SetLastAuthThreshold sets LastAuthThreshold field to given value.

### HasLastAuthThreshold

`func (o *AuthenticatorValidateStage) HasLastAuthThreshold() bool`

HasLastAuthThreshold returns a boolean if a field has been set.

### GetWebauthnUserVerification

`func (o *AuthenticatorValidateStage) GetWebauthnUserVerification() UserVerificationEnum`

GetWebauthnUserVerification returns the WebauthnUserVerification field if non-nil, zero value otherwise.

### GetWebauthnUserVerificationOk

`func (o *AuthenticatorValidateStage) GetWebauthnUserVerificationOk() (*UserVerificationEnum, bool)`

GetWebauthnUserVerificationOk returns a tuple with the WebauthnUserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebauthnUserVerification

`func (o *AuthenticatorValidateStage) SetWebauthnUserVerification(v UserVerificationEnum)`

SetWebauthnUserVerification sets WebauthnUserVerification field to given value.

### HasWebauthnUserVerification

`func (o *AuthenticatorValidateStage) HasWebauthnUserVerification() bool`

HasWebauthnUserVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


