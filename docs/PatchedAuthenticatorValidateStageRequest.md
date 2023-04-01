# PatchedAuthenticatorValidateStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**NotConfiguredAction** | Pointer to [**NotConfiguredActionEnum**](NotConfiguredActionEnum.md) |  | [optional] 
**DeviceClasses** | Pointer to [**[]DeviceClassesEnum**](DeviceClassesEnum.md) | Device classes which can be used to authenticate | [optional] 
**ConfigurationStages** | Pointer to **[]string** | Stages used to configure Authenticator when user doesn&#39;t have any compatible devices. After this configuration Stage passes, the user is not prompted again. | [optional] 
**LastAuthThreshold** | Pointer to **string** | If any of the user&#39;s device has been used within this threshold, this stage will be skipped | [optional] 
**WebauthnUserVerification** | Pointer to [**UserVerificationEnum**](UserVerificationEnum.md) |  | [optional] 

## Methods

### NewPatchedAuthenticatorValidateStageRequest

`func NewPatchedAuthenticatorValidateStageRequest() *PatchedAuthenticatorValidateStageRequest`

NewPatchedAuthenticatorValidateStageRequest instantiates a new PatchedAuthenticatorValidateStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAuthenticatorValidateStageRequestWithDefaults

`func NewPatchedAuthenticatorValidateStageRequestWithDefaults() *PatchedAuthenticatorValidateStageRequest`

NewPatchedAuthenticatorValidateStageRequestWithDefaults instantiates a new PatchedAuthenticatorValidateStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedAuthenticatorValidateStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedAuthenticatorValidateStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedAuthenticatorValidateStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedAuthenticatorValidateStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlowSet

`func (o *PatchedAuthenticatorValidateStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PatchedAuthenticatorValidateStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PatchedAuthenticatorValidateStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *PatchedAuthenticatorValidateStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetNotConfiguredAction

`func (o *PatchedAuthenticatorValidateStageRequest) GetNotConfiguredAction() NotConfiguredActionEnum`

GetNotConfiguredAction returns the NotConfiguredAction field if non-nil, zero value otherwise.

### GetNotConfiguredActionOk

`func (o *PatchedAuthenticatorValidateStageRequest) GetNotConfiguredActionOk() (*NotConfiguredActionEnum, bool)`

GetNotConfiguredActionOk returns a tuple with the NotConfiguredAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotConfiguredAction

`func (o *PatchedAuthenticatorValidateStageRequest) SetNotConfiguredAction(v NotConfiguredActionEnum)`

SetNotConfiguredAction sets NotConfiguredAction field to given value.

### HasNotConfiguredAction

`func (o *PatchedAuthenticatorValidateStageRequest) HasNotConfiguredAction() bool`

HasNotConfiguredAction returns a boolean if a field has been set.

### GetDeviceClasses

`func (o *PatchedAuthenticatorValidateStageRequest) GetDeviceClasses() []DeviceClassesEnum`

GetDeviceClasses returns the DeviceClasses field if non-nil, zero value otherwise.

### GetDeviceClassesOk

`func (o *PatchedAuthenticatorValidateStageRequest) GetDeviceClassesOk() (*[]DeviceClassesEnum, bool)`

GetDeviceClassesOk returns a tuple with the DeviceClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClasses

`func (o *PatchedAuthenticatorValidateStageRequest) SetDeviceClasses(v []DeviceClassesEnum)`

SetDeviceClasses sets DeviceClasses field to given value.

### HasDeviceClasses

`func (o *PatchedAuthenticatorValidateStageRequest) HasDeviceClasses() bool`

HasDeviceClasses returns a boolean if a field has been set.

### GetConfigurationStages

`func (o *PatchedAuthenticatorValidateStageRequest) GetConfigurationStages() []string`

GetConfigurationStages returns the ConfigurationStages field if non-nil, zero value otherwise.

### GetConfigurationStagesOk

`func (o *PatchedAuthenticatorValidateStageRequest) GetConfigurationStagesOk() (*[]string, bool)`

GetConfigurationStagesOk returns a tuple with the ConfigurationStages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationStages

`func (o *PatchedAuthenticatorValidateStageRequest) SetConfigurationStages(v []string)`

SetConfigurationStages sets ConfigurationStages field to given value.

### HasConfigurationStages

`func (o *PatchedAuthenticatorValidateStageRequest) HasConfigurationStages() bool`

HasConfigurationStages returns a boolean if a field has been set.

### GetLastAuthThreshold

`func (o *PatchedAuthenticatorValidateStageRequest) GetLastAuthThreshold() string`

GetLastAuthThreshold returns the LastAuthThreshold field if non-nil, zero value otherwise.

### GetLastAuthThresholdOk

`func (o *PatchedAuthenticatorValidateStageRequest) GetLastAuthThresholdOk() (*string, bool)`

GetLastAuthThresholdOk returns a tuple with the LastAuthThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuthThreshold

`func (o *PatchedAuthenticatorValidateStageRequest) SetLastAuthThreshold(v string)`

SetLastAuthThreshold sets LastAuthThreshold field to given value.

### HasLastAuthThreshold

`func (o *PatchedAuthenticatorValidateStageRequest) HasLastAuthThreshold() bool`

HasLastAuthThreshold returns a boolean if a field has been set.

### GetWebauthnUserVerification

`func (o *PatchedAuthenticatorValidateStageRequest) GetWebauthnUserVerification() UserVerificationEnum`

GetWebauthnUserVerification returns the WebauthnUserVerification field if non-nil, zero value otherwise.

### GetWebauthnUserVerificationOk

`func (o *PatchedAuthenticatorValidateStageRequest) GetWebauthnUserVerificationOk() (*UserVerificationEnum, bool)`

GetWebauthnUserVerificationOk returns a tuple with the WebauthnUserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebauthnUserVerification

`func (o *PatchedAuthenticatorValidateStageRequest) SetWebauthnUserVerification(v UserVerificationEnum)`

SetWebauthnUserVerification sets WebauthnUserVerification field to given value.

### HasWebauthnUserVerification

`func (o *PatchedAuthenticatorValidateStageRequest) HasWebauthnUserVerification() bool`

HasWebauthnUserVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


