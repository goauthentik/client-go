# AuthenticatorWebAuthnStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Component** | **string** | Get object type so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**FlowSet** | [**[]FlowSet**](FlowSet.md) |  | [readonly] 
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**UserVerification** | Pointer to [**UserVerificationEnum**](UserVerificationEnum.md) |  | [optional] 
**AuthenticatorAttachment** | Pointer to [**NullableAuthenticatorAttachmentEnum**](AuthenticatorAttachmentEnum.md) |  | [optional] 
**ResidentKeyRequirement** | Pointer to [**ResidentKeyRequirementEnum**](ResidentKeyRequirementEnum.md) |  | [optional] 
**DeviceTypeRestrictions** | Pointer to **[]string** |  | [optional] 
**DeviceTypeRestrictionsObj** | [**[]WebAuthnDeviceType**](WebAuthnDeviceType.md) |  | [readonly] 
**MaxAttempts** | Pointer to **int32** |  | [optional] 

## Methods

### NewAuthenticatorWebAuthnStage

`func NewAuthenticatorWebAuthnStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, flowSet []FlowSet, deviceTypeRestrictionsObj []WebAuthnDeviceType, ) *AuthenticatorWebAuthnStage`

NewAuthenticatorWebAuthnStage instantiates a new AuthenticatorWebAuthnStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorWebAuthnStageWithDefaults

`func NewAuthenticatorWebAuthnStageWithDefaults() *AuthenticatorWebAuthnStage`

NewAuthenticatorWebAuthnStageWithDefaults instantiates a new AuthenticatorWebAuthnStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *AuthenticatorWebAuthnStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *AuthenticatorWebAuthnStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *AuthenticatorWebAuthnStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *AuthenticatorWebAuthnStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticatorWebAuthnStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticatorWebAuthnStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *AuthenticatorWebAuthnStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AuthenticatorWebAuthnStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AuthenticatorWebAuthnStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *AuthenticatorWebAuthnStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *AuthenticatorWebAuthnStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *AuthenticatorWebAuthnStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *AuthenticatorWebAuthnStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *AuthenticatorWebAuthnStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *AuthenticatorWebAuthnStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *AuthenticatorWebAuthnStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *AuthenticatorWebAuthnStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *AuthenticatorWebAuthnStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *AuthenticatorWebAuthnStage) GetFlowSet() []FlowSet`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *AuthenticatorWebAuthnStage) GetFlowSetOk() (*[]FlowSet, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *AuthenticatorWebAuthnStage) SetFlowSet(v []FlowSet)`

SetFlowSet sets FlowSet field to given value.


### GetConfigureFlow

`func (o *AuthenticatorWebAuthnStage) GetConfigureFlow() string`

GetConfigureFlow returns the ConfigureFlow field if non-nil, zero value otherwise.

### GetConfigureFlowOk

`func (o *AuthenticatorWebAuthnStage) GetConfigureFlowOk() (*string, bool)`

GetConfigureFlowOk returns a tuple with the ConfigureFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureFlow

`func (o *AuthenticatorWebAuthnStage) SetConfigureFlow(v string)`

SetConfigureFlow sets ConfigureFlow field to given value.

### HasConfigureFlow

`func (o *AuthenticatorWebAuthnStage) HasConfigureFlow() bool`

HasConfigureFlow returns a boolean if a field has been set.

### SetConfigureFlowNil

`func (o *AuthenticatorWebAuthnStage) SetConfigureFlowNil(b bool)`

 SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil

### UnsetConfigureFlow
`func (o *AuthenticatorWebAuthnStage) UnsetConfigureFlow()`

UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
### GetFriendlyName

`func (o *AuthenticatorWebAuthnStage) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *AuthenticatorWebAuthnStage) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *AuthenticatorWebAuthnStage) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *AuthenticatorWebAuthnStage) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetUserVerification

`func (o *AuthenticatorWebAuthnStage) GetUserVerification() UserVerificationEnum`

GetUserVerification returns the UserVerification field if non-nil, zero value otherwise.

### GetUserVerificationOk

`func (o *AuthenticatorWebAuthnStage) GetUserVerificationOk() (*UserVerificationEnum, bool)`

GetUserVerificationOk returns a tuple with the UserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerification

`func (o *AuthenticatorWebAuthnStage) SetUserVerification(v UserVerificationEnum)`

SetUserVerification sets UserVerification field to given value.

### HasUserVerification

`func (o *AuthenticatorWebAuthnStage) HasUserVerification() bool`

HasUserVerification returns a boolean if a field has been set.

### GetAuthenticatorAttachment

`func (o *AuthenticatorWebAuthnStage) GetAuthenticatorAttachment() AuthenticatorAttachmentEnum`

GetAuthenticatorAttachment returns the AuthenticatorAttachment field if non-nil, zero value otherwise.

### GetAuthenticatorAttachmentOk

`func (o *AuthenticatorWebAuthnStage) GetAuthenticatorAttachmentOk() (*AuthenticatorAttachmentEnum, bool)`

GetAuthenticatorAttachmentOk returns a tuple with the AuthenticatorAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorAttachment

`func (o *AuthenticatorWebAuthnStage) SetAuthenticatorAttachment(v AuthenticatorAttachmentEnum)`

SetAuthenticatorAttachment sets AuthenticatorAttachment field to given value.

### HasAuthenticatorAttachment

`func (o *AuthenticatorWebAuthnStage) HasAuthenticatorAttachment() bool`

HasAuthenticatorAttachment returns a boolean if a field has been set.

### SetAuthenticatorAttachmentNil

`func (o *AuthenticatorWebAuthnStage) SetAuthenticatorAttachmentNil(b bool)`

 SetAuthenticatorAttachmentNil sets the value for AuthenticatorAttachment to be an explicit nil

### UnsetAuthenticatorAttachment
`func (o *AuthenticatorWebAuthnStage) UnsetAuthenticatorAttachment()`

UnsetAuthenticatorAttachment ensures that no value is present for AuthenticatorAttachment, not even an explicit nil
### GetResidentKeyRequirement

`func (o *AuthenticatorWebAuthnStage) GetResidentKeyRequirement() ResidentKeyRequirementEnum`

GetResidentKeyRequirement returns the ResidentKeyRequirement field if non-nil, zero value otherwise.

### GetResidentKeyRequirementOk

`func (o *AuthenticatorWebAuthnStage) GetResidentKeyRequirementOk() (*ResidentKeyRequirementEnum, bool)`

GetResidentKeyRequirementOk returns a tuple with the ResidentKeyRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentKeyRequirement

`func (o *AuthenticatorWebAuthnStage) SetResidentKeyRequirement(v ResidentKeyRequirementEnum)`

SetResidentKeyRequirement sets ResidentKeyRequirement field to given value.

### HasResidentKeyRequirement

`func (o *AuthenticatorWebAuthnStage) HasResidentKeyRequirement() bool`

HasResidentKeyRequirement returns a boolean if a field has been set.

### GetDeviceTypeRestrictions

`func (o *AuthenticatorWebAuthnStage) GetDeviceTypeRestrictions() []string`

GetDeviceTypeRestrictions returns the DeviceTypeRestrictions field if non-nil, zero value otherwise.

### GetDeviceTypeRestrictionsOk

`func (o *AuthenticatorWebAuthnStage) GetDeviceTypeRestrictionsOk() (*[]string, bool)`

GetDeviceTypeRestrictionsOk returns a tuple with the DeviceTypeRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypeRestrictions

`func (o *AuthenticatorWebAuthnStage) SetDeviceTypeRestrictions(v []string)`

SetDeviceTypeRestrictions sets DeviceTypeRestrictions field to given value.

### HasDeviceTypeRestrictions

`func (o *AuthenticatorWebAuthnStage) HasDeviceTypeRestrictions() bool`

HasDeviceTypeRestrictions returns a boolean if a field has been set.

### GetDeviceTypeRestrictionsObj

`func (o *AuthenticatorWebAuthnStage) GetDeviceTypeRestrictionsObj() []WebAuthnDeviceType`

GetDeviceTypeRestrictionsObj returns the DeviceTypeRestrictionsObj field if non-nil, zero value otherwise.

### GetDeviceTypeRestrictionsObjOk

`func (o *AuthenticatorWebAuthnStage) GetDeviceTypeRestrictionsObjOk() (*[]WebAuthnDeviceType, bool)`

GetDeviceTypeRestrictionsObjOk returns a tuple with the DeviceTypeRestrictionsObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypeRestrictionsObj

`func (o *AuthenticatorWebAuthnStage) SetDeviceTypeRestrictionsObj(v []WebAuthnDeviceType)`

SetDeviceTypeRestrictionsObj sets DeviceTypeRestrictionsObj field to given value.


### GetMaxAttempts

`func (o *AuthenticatorWebAuthnStage) GetMaxAttempts() int32`

GetMaxAttempts returns the MaxAttempts field if non-nil, zero value otherwise.

### GetMaxAttemptsOk

`func (o *AuthenticatorWebAuthnStage) GetMaxAttemptsOk() (*int32, bool)`

GetMaxAttemptsOk returns a tuple with the MaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttempts

`func (o *AuthenticatorWebAuthnStage) SetMaxAttempts(v int32)`

SetMaxAttempts sets MaxAttempts field to given value.

### HasMaxAttempts

`func (o *AuthenticatorWebAuthnStage) HasMaxAttempts() bool`

HasMaxAttempts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


