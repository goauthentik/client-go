# PatchedAuthenticateWebAuthnStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 
**UserVerification** | Pointer to [**UserVerificationEnum**](UserVerificationEnum.md) |  | [optional] 
**AuthenticatorAttachment** | Pointer to [**NullableAuthenticatorAttachmentEnum**](AuthenticatorAttachmentEnum.md) |  | [optional] 
**ResidentKeyRequirement** | Pointer to [**ResidentKeyRequirementEnum**](ResidentKeyRequirementEnum.md) |  | [optional] 

## Methods

### NewPatchedAuthenticateWebAuthnStageRequest

`func NewPatchedAuthenticateWebAuthnStageRequest() *PatchedAuthenticateWebAuthnStageRequest`

NewPatchedAuthenticateWebAuthnStageRequest instantiates a new PatchedAuthenticateWebAuthnStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAuthenticateWebAuthnStageRequestWithDefaults

`func NewPatchedAuthenticateWebAuthnStageRequestWithDefaults() *PatchedAuthenticateWebAuthnStageRequest`

NewPatchedAuthenticateWebAuthnStageRequestWithDefaults instantiates a new PatchedAuthenticateWebAuthnStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedAuthenticateWebAuthnStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedAuthenticateWebAuthnStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedAuthenticateWebAuthnStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedAuthenticateWebAuthnStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlowSet

`func (o *PatchedAuthenticateWebAuthnStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PatchedAuthenticateWebAuthnStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PatchedAuthenticateWebAuthnStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *PatchedAuthenticateWebAuthnStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetConfigureFlow

`func (o *PatchedAuthenticateWebAuthnStageRequest) GetConfigureFlow() string`

GetConfigureFlow returns the ConfigureFlow field if non-nil, zero value otherwise.

### GetConfigureFlowOk

`func (o *PatchedAuthenticateWebAuthnStageRequest) GetConfigureFlowOk() (*string, bool)`

GetConfigureFlowOk returns a tuple with the ConfigureFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureFlow

`func (o *PatchedAuthenticateWebAuthnStageRequest) SetConfigureFlow(v string)`

SetConfigureFlow sets ConfigureFlow field to given value.

### HasConfigureFlow

`func (o *PatchedAuthenticateWebAuthnStageRequest) HasConfigureFlow() bool`

HasConfigureFlow returns a boolean if a field has been set.

### SetConfigureFlowNil

`func (o *PatchedAuthenticateWebAuthnStageRequest) SetConfigureFlowNil(b bool)`

 SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil

### UnsetConfigureFlow
`func (o *PatchedAuthenticateWebAuthnStageRequest) UnsetConfigureFlow()`

UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
### GetFriendlyName

`func (o *PatchedAuthenticateWebAuthnStageRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *PatchedAuthenticateWebAuthnStageRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *PatchedAuthenticateWebAuthnStageRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *PatchedAuthenticateWebAuthnStageRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *PatchedAuthenticateWebAuthnStageRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *PatchedAuthenticateWebAuthnStageRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetUserVerification

`func (o *PatchedAuthenticateWebAuthnStageRequest) GetUserVerification() UserVerificationEnum`

GetUserVerification returns the UserVerification field if non-nil, zero value otherwise.

### GetUserVerificationOk

`func (o *PatchedAuthenticateWebAuthnStageRequest) GetUserVerificationOk() (*UserVerificationEnum, bool)`

GetUserVerificationOk returns a tuple with the UserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerification

`func (o *PatchedAuthenticateWebAuthnStageRequest) SetUserVerification(v UserVerificationEnum)`

SetUserVerification sets UserVerification field to given value.

### HasUserVerification

`func (o *PatchedAuthenticateWebAuthnStageRequest) HasUserVerification() bool`

HasUserVerification returns a boolean if a field has been set.

### GetAuthenticatorAttachment

`func (o *PatchedAuthenticateWebAuthnStageRequest) GetAuthenticatorAttachment() AuthenticatorAttachmentEnum`

GetAuthenticatorAttachment returns the AuthenticatorAttachment field if non-nil, zero value otherwise.

### GetAuthenticatorAttachmentOk

`func (o *PatchedAuthenticateWebAuthnStageRequest) GetAuthenticatorAttachmentOk() (*AuthenticatorAttachmentEnum, bool)`

GetAuthenticatorAttachmentOk returns a tuple with the AuthenticatorAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorAttachment

`func (o *PatchedAuthenticateWebAuthnStageRequest) SetAuthenticatorAttachment(v AuthenticatorAttachmentEnum)`

SetAuthenticatorAttachment sets AuthenticatorAttachment field to given value.

### HasAuthenticatorAttachment

`func (o *PatchedAuthenticateWebAuthnStageRequest) HasAuthenticatorAttachment() bool`

HasAuthenticatorAttachment returns a boolean if a field has been set.

### SetAuthenticatorAttachmentNil

`func (o *PatchedAuthenticateWebAuthnStageRequest) SetAuthenticatorAttachmentNil(b bool)`

 SetAuthenticatorAttachmentNil sets the value for AuthenticatorAttachment to be an explicit nil

### UnsetAuthenticatorAttachment
`func (o *PatchedAuthenticateWebAuthnStageRequest) UnsetAuthenticatorAttachment()`

UnsetAuthenticatorAttachment ensures that no value is present for AuthenticatorAttachment, not even an explicit nil
### GetResidentKeyRequirement

`func (o *PatchedAuthenticateWebAuthnStageRequest) GetResidentKeyRequirement() ResidentKeyRequirementEnum`

GetResidentKeyRequirement returns the ResidentKeyRequirement field if non-nil, zero value otherwise.

### GetResidentKeyRequirementOk

`func (o *PatchedAuthenticateWebAuthnStageRequest) GetResidentKeyRequirementOk() (*ResidentKeyRequirementEnum, bool)`

GetResidentKeyRequirementOk returns a tuple with the ResidentKeyRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentKeyRequirement

`func (o *PatchedAuthenticateWebAuthnStageRequest) SetResidentKeyRequirement(v ResidentKeyRequirementEnum)`

SetResidentKeyRequirement sets ResidentKeyRequirement field to given value.

### HasResidentKeyRequirement

`func (o *PatchedAuthenticateWebAuthnStageRequest) HasResidentKeyRequirement() bool`

HasResidentKeyRequirement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


