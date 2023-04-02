# AuthenticateWebAuthnStage

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
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 
**UserVerification** | Pointer to [**UserVerificationEnum**](UserVerificationEnum.md) |  | [optional] 
**AuthenticatorAttachment** | Pointer to [**NullableAuthenticatorAttachmentEnum**](AuthenticatorAttachmentEnum.md) |  | [optional] 
**ResidentKeyRequirement** | Pointer to [**ResidentKeyRequirementEnum**](ResidentKeyRequirementEnum.md) |  | [optional] 

## Methods

### NewAuthenticateWebAuthnStage

`func NewAuthenticateWebAuthnStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, ) *AuthenticateWebAuthnStage`

NewAuthenticateWebAuthnStage instantiates a new AuthenticateWebAuthnStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticateWebAuthnStageWithDefaults

`func NewAuthenticateWebAuthnStageWithDefaults() *AuthenticateWebAuthnStage`

NewAuthenticateWebAuthnStageWithDefaults instantiates a new AuthenticateWebAuthnStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *AuthenticateWebAuthnStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *AuthenticateWebAuthnStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *AuthenticateWebAuthnStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *AuthenticateWebAuthnStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticateWebAuthnStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticateWebAuthnStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *AuthenticateWebAuthnStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AuthenticateWebAuthnStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AuthenticateWebAuthnStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *AuthenticateWebAuthnStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *AuthenticateWebAuthnStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *AuthenticateWebAuthnStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *AuthenticateWebAuthnStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *AuthenticateWebAuthnStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *AuthenticateWebAuthnStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *AuthenticateWebAuthnStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *AuthenticateWebAuthnStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *AuthenticateWebAuthnStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *AuthenticateWebAuthnStage) GetFlowSet() []FlowSet`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *AuthenticateWebAuthnStage) GetFlowSetOk() (*[]FlowSet, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *AuthenticateWebAuthnStage) SetFlowSet(v []FlowSet)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *AuthenticateWebAuthnStage) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetConfigureFlow

`func (o *AuthenticateWebAuthnStage) GetConfigureFlow() string`

GetConfigureFlow returns the ConfigureFlow field if non-nil, zero value otherwise.

### GetConfigureFlowOk

`func (o *AuthenticateWebAuthnStage) GetConfigureFlowOk() (*string, bool)`

GetConfigureFlowOk returns a tuple with the ConfigureFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureFlow

`func (o *AuthenticateWebAuthnStage) SetConfigureFlow(v string)`

SetConfigureFlow sets ConfigureFlow field to given value.

### HasConfigureFlow

`func (o *AuthenticateWebAuthnStage) HasConfigureFlow() bool`

HasConfigureFlow returns a boolean if a field has been set.

### SetConfigureFlowNil

`func (o *AuthenticateWebAuthnStage) SetConfigureFlowNil(b bool)`

 SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil

### UnsetConfigureFlow
`func (o *AuthenticateWebAuthnStage) UnsetConfigureFlow()`

UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
### GetFriendlyName

`func (o *AuthenticateWebAuthnStage) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *AuthenticateWebAuthnStage) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *AuthenticateWebAuthnStage) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *AuthenticateWebAuthnStage) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *AuthenticateWebAuthnStage) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *AuthenticateWebAuthnStage) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetUserVerification

`func (o *AuthenticateWebAuthnStage) GetUserVerification() UserVerificationEnum`

GetUserVerification returns the UserVerification field if non-nil, zero value otherwise.

### GetUserVerificationOk

`func (o *AuthenticateWebAuthnStage) GetUserVerificationOk() (*UserVerificationEnum, bool)`

GetUserVerificationOk returns a tuple with the UserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerification

`func (o *AuthenticateWebAuthnStage) SetUserVerification(v UserVerificationEnum)`

SetUserVerification sets UserVerification field to given value.

### HasUserVerification

`func (o *AuthenticateWebAuthnStage) HasUserVerification() bool`

HasUserVerification returns a boolean if a field has been set.

### GetAuthenticatorAttachment

`func (o *AuthenticateWebAuthnStage) GetAuthenticatorAttachment() AuthenticatorAttachmentEnum`

GetAuthenticatorAttachment returns the AuthenticatorAttachment field if non-nil, zero value otherwise.

### GetAuthenticatorAttachmentOk

`func (o *AuthenticateWebAuthnStage) GetAuthenticatorAttachmentOk() (*AuthenticatorAttachmentEnum, bool)`

GetAuthenticatorAttachmentOk returns a tuple with the AuthenticatorAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorAttachment

`func (o *AuthenticateWebAuthnStage) SetAuthenticatorAttachment(v AuthenticatorAttachmentEnum)`

SetAuthenticatorAttachment sets AuthenticatorAttachment field to given value.

### HasAuthenticatorAttachment

`func (o *AuthenticateWebAuthnStage) HasAuthenticatorAttachment() bool`

HasAuthenticatorAttachment returns a boolean if a field has been set.

### SetAuthenticatorAttachmentNil

`func (o *AuthenticateWebAuthnStage) SetAuthenticatorAttachmentNil(b bool)`

 SetAuthenticatorAttachmentNil sets the value for AuthenticatorAttachment to be an explicit nil

### UnsetAuthenticatorAttachment
`func (o *AuthenticateWebAuthnStage) UnsetAuthenticatorAttachment()`

UnsetAuthenticatorAttachment ensures that no value is present for AuthenticatorAttachment, not even an explicit nil
### GetResidentKeyRequirement

`func (o *AuthenticateWebAuthnStage) GetResidentKeyRequirement() ResidentKeyRequirementEnum`

GetResidentKeyRequirement returns the ResidentKeyRequirement field if non-nil, zero value otherwise.

### GetResidentKeyRequirementOk

`func (o *AuthenticateWebAuthnStage) GetResidentKeyRequirementOk() (*ResidentKeyRequirementEnum, bool)`

GetResidentKeyRequirementOk returns a tuple with the ResidentKeyRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentKeyRequirement

`func (o *AuthenticateWebAuthnStage) SetResidentKeyRequirement(v ResidentKeyRequirementEnum)`

SetResidentKeyRequirement sets ResidentKeyRequirement field to given value.

### HasResidentKeyRequirement

`func (o *AuthenticateWebAuthnStage) HasResidentKeyRequirement() bool`

HasResidentKeyRequirement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


