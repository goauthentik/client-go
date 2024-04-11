# AuthenticatorWebAuthnStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 
**UserVerification** | Pointer to [**UserVerificationEnum**](UserVerificationEnum.md) |  | [optional] 
**AuthenticatorAttachment** | Pointer to [**NullableAuthenticatorAttachmentEnum**](AuthenticatorAttachmentEnum.md) |  | [optional] 
**ResidentKeyRequirement** | Pointer to [**ResidentKeyRequirementEnum**](ResidentKeyRequirementEnum.md) |  | [optional] 
**DeviceTypeRestrictions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAuthenticatorWebAuthnStageRequest

`func NewAuthenticatorWebAuthnStageRequest(name string, ) *AuthenticatorWebAuthnStageRequest`

NewAuthenticatorWebAuthnStageRequest instantiates a new AuthenticatorWebAuthnStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorWebAuthnStageRequestWithDefaults

`func NewAuthenticatorWebAuthnStageRequestWithDefaults() *AuthenticatorWebAuthnStageRequest`

NewAuthenticatorWebAuthnStageRequestWithDefaults instantiates a new AuthenticatorWebAuthnStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthenticatorWebAuthnStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticatorWebAuthnStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticatorWebAuthnStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowSet

`func (o *AuthenticatorWebAuthnStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *AuthenticatorWebAuthnStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *AuthenticatorWebAuthnStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *AuthenticatorWebAuthnStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetConfigureFlow

`func (o *AuthenticatorWebAuthnStageRequest) GetConfigureFlow() string`

GetConfigureFlow returns the ConfigureFlow field if non-nil, zero value otherwise.

### GetConfigureFlowOk

`func (o *AuthenticatorWebAuthnStageRequest) GetConfigureFlowOk() (*string, bool)`

GetConfigureFlowOk returns a tuple with the ConfigureFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureFlow

`func (o *AuthenticatorWebAuthnStageRequest) SetConfigureFlow(v string)`

SetConfigureFlow sets ConfigureFlow field to given value.

### HasConfigureFlow

`func (o *AuthenticatorWebAuthnStageRequest) HasConfigureFlow() bool`

HasConfigureFlow returns a boolean if a field has been set.

### SetConfigureFlowNil

`func (o *AuthenticatorWebAuthnStageRequest) SetConfigureFlowNil(b bool)`

 SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil

### UnsetConfigureFlow
`func (o *AuthenticatorWebAuthnStageRequest) UnsetConfigureFlow()`

UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
### GetFriendlyName

`func (o *AuthenticatorWebAuthnStageRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *AuthenticatorWebAuthnStageRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *AuthenticatorWebAuthnStageRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *AuthenticatorWebAuthnStageRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *AuthenticatorWebAuthnStageRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *AuthenticatorWebAuthnStageRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetUserVerification

`func (o *AuthenticatorWebAuthnStageRequest) GetUserVerification() UserVerificationEnum`

GetUserVerification returns the UserVerification field if non-nil, zero value otherwise.

### GetUserVerificationOk

`func (o *AuthenticatorWebAuthnStageRequest) GetUserVerificationOk() (*UserVerificationEnum, bool)`

GetUserVerificationOk returns a tuple with the UserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerification

`func (o *AuthenticatorWebAuthnStageRequest) SetUserVerification(v UserVerificationEnum)`

SetUserVerification sets UserVerification field to given value.

### HasUserVerification

`func (o *AuthenticatorWebAuthnStageRequest) HasUserVerification() bool`

HasUserVerification returns a boolean if a field has been set.

### GetAuthenticatorAttachment

`func (o *AuthenticatorWebAuthnStageRequest) GetAuthenticatorAttachment() AuthenticatorAttachmentEnum`

GetAuthenticatorAttachment returns the AuthenticatorAttachment field if non-nil, zero value otherwise.

### GetAuthenticatorAttachmentOk

`func (o *AuthenticatorWebAuthnStageRequest) GetAuthenticatorAttachmentOk() (*AuthenticatorAttachmentEnum, bool)`

GetAuthenticatorAttachmentOk returns a tuple with the AuthenticatorAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorAttachment

`func (o *AuthenticatorWebAuthnStageRequest) SetAuthenticatorAttachment(v AuthenticatorAttachmentEnum)`

SetAuthenticatorAttachment sets AuthenticatorAttachment field to given value.

### HasAuthenticatorAttachment

`func (o *AuthenticatorWebAuthnStageRequest) HasAuthenticatorAttachment() bool`

HasAuthenticatorAttachment returns a boolean if a field has been set.

### SetAuthenticatorAttachmentNil

`func (o *AuthenticatorWebAuthnStageRequest) SetAuthenticatorAttachmentNil(b bool)`

 SetAuthenticatorAttachmentNil sets the value for AuthenticatorAttachment to be an explicit nil

### UnsetAuthenticatorAttachment
`func (o *AuthenticatorWebAuthnStageRequest) UnsetAuthenticatorAttachment()`

UnsetAuthenticatorAttachment ensures that no value is present for AuthenticatorAttachment, not even an explicit nil
### GetResidentKeyRequirement

`func (o *AuthenticatorWebAuthnStageRequest) GetResidentKeyRequirement() ResidentKeyRequirementEnum`

GetResidentKeyRequirement returns the ResidentKeyRequirement field if non-nil, zero value otherwise.

### GetResidentKeyRequirementOk

`func (o *AuthenticatorWebAuthnStageRequest) GetResidentKeyRequirementOk() (*ResidentKeyRequirementEnum, bool)`

GetResidentKeyRequirementOk returns a tuple with the ResidentKeyRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentKeyRequirement

`func (o *AuthenticatorWebAuthnStageRequest) SetResidentKeyRequirement(v ResidentKeyRequirementEnum)`

SetResidentKeyRequirement sets ResidentKeyRequirement field to given value.

### HasResidentKeyRequirement

`func (o *AuthenticatorWebAuthnStageRequest) HasResidentKeyRequirement() bool`

HasResidentKeyRequirement returns a boolean if a field has been set.

### GetDeviceTypeRestrictions

`func (o *AuthenticatorWebAuthnStageRequest) GetDeviceTypeRestrictions() []string`

GetDeviceTypeRestrictions returns the DeviceTypeRestrictions field if non-nil, zero value otherwise.

### GetDeviceTypeRestrictionsOk

`func (o *AuthenticatorWebAuthnStageRequest) GetDeviceTypeRestrictionsOk() (*[]string, bool)`

GetDeviceTypeRestrictionsOk returns a tuple with the DeviceTypeRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypeRestrictions

`func (o *AuthenticatorWebAuthnStageRequest) SetDeviceTypeRestrictions(v []string)`

SetDeviceTypeRestrictions sets DeviceTypeRestrictions field to given value.

### HasDeviceTypeRestrictions

`func (o *AuthenticatorWebAuthnStageRequest) HasDeviceTypeRestrictions() bool`

HasDeviceTypeRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


