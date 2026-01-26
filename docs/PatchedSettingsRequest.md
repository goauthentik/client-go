# PatchedSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatars** | Pointer to **string** | Configure how authentik should show avatars for users. | [optional] 
**DefaultUserChangeName** | Pointer to **bool** | Enable the ability for users to change their name. | [optional] 
**DefaultUserChangeEmail** | Pointer to **bool** | Enable the ability for users to change their email address. | [optional] 
**DefaultUserChangeUsername** | Pointer to **bool** | Enable the ability for users to change their username. | [optional] 
**EventRetention** | Pointer to **string** | Events will be deleted after this duration.(Format: weeks&#x3D;3;days&#x3D;2;hours&#x3D;3,seconds&#x3D;2). | [optional] 
**ReputationLowerLimit** | Pointer to **int32** | Reputation cannot decrease lower than this value. Zero or negative. | [optional] 
**ReputationUpperLimit** | Pointer to **int32** | Reputation cannot increase higher than this value. Zero or positive. | [optional] 
**FooterLinks** | Pointer to **interface{}** |  | [optional] 
**GdprCompliance** | Pointer to **bool** | When enabled, all the events caused by a user will be deleted upon the user&#39;s deletion. | [optional] 
**Impersonation** | Pointer to **bool** | Globally enable/disable impersonation. | [optional] 
**ImpersonationRequireReason** | Pointer to **bool** | Require administrators to provide a reason for impersonating a user. | [optional] 
**DefaultTokenDuration** | Pointer to **string** | Default token duration | [optional] 
**DefaultTokenLength** | Pointer to **int32** | Default token length | [optional] 
**Flags** | Pointer to [**PatchedSettingsRequestFlags**](PatchedSettingsRequestFlags.md) |  | [optional] 

## Methods

### NewPatchedSettingsRequest

`func NewPatchedSettingsRequest() *PatchedSettingsRequest`

NewPatchedSettingsRequest instantiates a new PatchedSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSettingsRequestWithDefaults

`func NewPatchedSettingsRequestWithDefaults() *PatchedSettingsRequest`

NewPatchedSettingsRequestWithDefaults instantiates a new PatchedSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatars

`func (o *PatchedSettingsRequest) GetAvatars() string`

GetAvatars returns the Avatars field if non-nil, zero value otherwise.

### GetAvatarsOk

`func (o *PatchedSettingsRequest) GetAvatarsOk() (*string, bool)`

GetAvatarsOk returns a tuple with the Avatars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatars

`func (o *PatchedSettingsRequest) SetAvatars(v string)`

SetAvatars sets Avatars field to given value.

### HasAvatars

`func (o *PatchedSettingsRequest) HasAvatars() bool`

HasAvatars returns a boolean if a field has been set.

### GetDefaultUserChangeName

`func (o *PatchedSettingsRequest) GetDefaultUserChangeName() bool`

GetDefaultUserChangeName returns the DefaultUserChangeName field if non-nil, zero value otherwise.

### GetDefaultUserChangeNameOk

`func (o *PatchedSettingsRequest) GetDefaultUserChangeNameOk() (*bool, bool)`

GetDefaultUserChangeNameOk returns a tuple with the DefaultUserChangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserChangeName

`func (o *PatchedSettingsRequest) SetDefaultUserChangeName(v bool)`

SetDefaultUserChangeName sets DefaultUserChangeName field to given value.

### HasDefaultUserChangeName

`func (o *PatchedSettingsRequest) HasDefaultUserChangeName() bool`

HasDefaultUserChangeName returns a boolean if a field has been set.

### GetDefaultUserChangeEmail

`func (o *PatchedSettingsRequest) GetDefaultUserChangeEmail() bool`

GetDefaultUserChangeEmail returns the DefaultUserChangeEmail field if non-nil, zero value otherwise.

### GetDefaultUserChangeEmailOk

`func (o *PatchedSettingsRequest) GetDefaultUserChangeEmailOk() (*bool, bool)`

GetDefaultUserChangeEmailOk returns a tuple with the DefaultUserChangeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserChangeEmail

`func (o *PatchedSettingsRequest) SetDefaultUserChangeEmail(v bool)`

SetDefaultUserChangeEmail sets DefaultUserChangeEmail field to given value.

### HasDefaultUserChangeEmail

`func (o *PatchedSettingsRequest) HasDefaultUserChangeEmail() bool`

HasDefaultUserChangeEmail returns a boolean if a field has been set.

### GetDefaultUserChangeUsername

`func (o *PatchedSettingsRequest) GetDefaultUserChangeUsername() bool`

GetDefaultUserChangeUsername returns the DefaultUserChangeUsername field if non-nil, zero value otherwise.

### GetDefaultUserChangeUsernameOk

`func (o *PatchedSettingsRequest) GetDefaultUserChangeUsernameOk() (*bool, bool)`

GetDefaultUserChangeUsernameOk returns a tuple with the DefaultUserChangeUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserChangeUsername

`func (o *PatchedSettingsRequest) SetDefaultUserChangeUsername(v bool)`

SetDefaultUserChangeUsername sets DefaultUserChangeUsername field to given value.

### HasDefaultUserChangeUsername

`func (o *PatchedSettingsRequest) HasDefaultUserChangeUsername() bool`

HasDefaultUserChangeUsername returns a boolean if a field has been set.

### GetEventRetention

`func (o *PatchedSettingsRequest) GetEventRetention() string`

GetEventRetention returns the EventRetention field if non-nil, zero value otherwise.

### GetEventRetentionOk

`func (o *PatchedSettingsRequest) GetEventRetentionOk() (*string, bool)`

GetEventRetentionOk returns a tuple with the EventRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRetention

`func (o *PatchedSettingsRequest) SetEventRetention(v string)`

SetEventRetention sets EventRetention field to given value.

### HasEventRetention

`func (o *PatchedSettingsRequest) HasEventRetention() bool`

HasEventRetention returns a boolean if a field has been set.

### GetReputationLowerLimit

`func (o *PatchedSettingsRequest) GetReputationLowerLimit() int32`

GetReputationLowerLimit returns the ReputationLowerLimit field if non-nil, zero value otherwise.

### GetReputationLowerLimitOk

`func (o *PatchedSettingsRequest) GetReputationLowerLimitOk() (*int32, bool)`

GetReputationLowerLimitOk returns a tuple with the ReputationLowerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputationLowerLimit

`func (o *PatchedSettingsRequest) SetReputationLowerLimit(v int32)`

SetReputationLowerLimit sets ReputationLowerLimit field to given value.

### HasReputationLowerLimit

`func (o *PatchedSettingsRequest) HasReputationLowerLimit() bool`

HasReputationLowerLimit returns a boolean if a field has been set.

### GetReputationUpperLimit

`func (o *PatchedSettingsRequest) GetReputationUpperLimit() int32`

GetReputationUpperLimit returns the ReputationUpperLimit field if non-nil, zero value otherwise.

### GetReputationUpperLimitOk

`func (o *PatchedSettingsRequest) GetReputationUpperLimitOk() (*int32, bool)`

GetReputationUpperLimitOk returns a tuple with the ReputationUpperLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputationUpperLimit

`func (o *PatchedSettingsRequest) SetReputationUpperLimit(v int32)`

SetReputationUpperLimit sets ReputationUpperLimit field to given value.

### HasReputationUpperLimit

`func (o *PatchedSettingsRequest) HasReputationUpperLimit() bool`

HasReputationUpperLimit returns a boolean if a field has been set.

### GetFooterLinks

`func (o *PatchedSettingsRequest) GetFooterLinks() interface{}`

GetFooterLinks returns the FooterLinks field if non-nil, zero value otherwise.

### GetFooterLinksOk

`func (o *PatchedSettingsRequest) GetFooterLinksOk() (*interface{}, bool)`

GetFooterLinksOk returns a tuple with the FooterLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterLinks

`func (o *PatchedSettingsRequest) SetFooterLinks(v interface{})`

SetFooterLinks sets FooterLinks field to given value.

### HasFooterLinks

`func (o *PatchedSettingsRequest) HasFooterLinks() bool`

HasFooterLinks returns a boolean if a field has been set.

### SetFooterLinksNil

`func (o *PatchedSettingsRequest) SetFooterLinksNil(b bool)`

 SetFooterLinksNil sets the value for FooterLinks to be an explicit nil

### UnsetFooterLinks
`func (o *PatchedSettingsRequest) UnsetFooterLinks()`

UnsetFooterLinks ensures that no value is present for FooterLinks, not even an explicit nil
### GetGdprCompliance

`func (o *PatchedSettingsRequest) GetGdprCompliance() bool`

GetGdprCompliance returns the GdprCompliance field if non-nil, zero value otherwise.

### GetGdprComplianceOk

`func (o *PatchedSettingsRequest) GetGdprComplianceOk() (*bool, bool)`

GetGdprComplianceOk returns a tuple with the GdprCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGdprCompliance

`func (o *PatchedSettingsRequest) SetGdprCompliance(v bool)`

SetGdprCompliance sets GdprCompliance field to given value.

### HasGdprCompliance

`func (o *PatchedSettingsRequest) HasGdprCompliance() bool`

HasGdprCompliance returns a boolean if a field has been set.

### GetImpersonation

`func (o *PatchedSettingsRequest) GetImpersonation() bool`

GetImpersonation returns the Impersonation field if non-nil, zero value otherwise.

### GetImpersonationOk

`func (o *PatchedSettingsRequest) GetImpersonationOk() (*bool, bool)`

GetImpersonationOk returns a tuple with the Impersonation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonation

`func (o *PatchedSettingsRequest) SetImpersonation(v bool)`

SetImpersonation sets Impersonation field to given value.

### HasImpersonation

`func (o *PatchedSettingsRequest) HasImpersonation() bool`

HasImpersonation returns a boolean if a field has been set.

### GetImpersonationRequireReason

`func (o *PatchedSettingsRequest) GetImpersonationRequireReason() bool`

GetImpersonationRequireReason returns the ImpersonationRequireReason field if non-nil, zero value otherwise.

### GetImpersonationRequireReasonOk

`func (o *PatchedSettingsRequest) GetImpersonationRequireReasonOk() (*bool, bool)`

GetImpersonationRequireReasonOk returns a tuple with the ImpersonationRequireReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonationRequireReason

`func (o *PatchedSettingsRequest) SetImpersonationRequireReason(v bool)`

SetImpersonationRequireReason sets ImpersonationRequireReason field to given value.

### HasImpersonationRequireReason

`func (o *PatchedSettingsRequest) HasImpersonationRequireReason() bool`

HasImpersonationRequireReason returns a boolean if a field has been set.

### GetDefaultTokenDuration

`func (o *PatchedSettingsRequest) GetDefaultTokenDuration() string`

GetDefaultTokenDuration returns the DefaultTokenDuration field if non-nil, zero value otherwise.

### GetDefaultTokenDurationOk

`func (o *PatchedSettingsRequest) GetDefaultTokenDurationOk() (*string, bool)`

GetDefaultTokenDurationOk returns a tuple with the DefaultTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTokenDuration

`func (o *PatchedSettingsRequest) SetDefaultTokenDuration(v string)`

SetDefaultTokenDuration sets DefaultTokenDuration field to given value.

### HasDefaultTokenDuration

`func (o *PatchedSettingsRequest) HasDefaultTokenDuration() bool`

HasDefaultTokenDuration returns a boolean if a field has been set.

### GetDefaultTokenLength

`func (o *PatchedSettingsRequest) GetDefaultTokenLength() int32`

GetDefaultTokenLength returns the DefaultTokenLength field if non-nil, zero value otherwise.

### GetDefaultTokenLengthOk

`func (o *PatchedSettingsRequest) GetDefaultTokenLengthOk() (*int32, bool)`

GetDefaultTokenLengthOk returns a tuple with the DefaultTokenLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTokenLength

`func (o *PatchedSettingsRequest) SetDefaultTokenLength(v int32)`

SetDefaultTokenLength sets DefaultTokenLength field to given value.

### HasDefaultTokenLength

`func (o *PatchedSettingsRequest) HasDefaultTokenLength() bool`

HasDefaultTokenLength returns a boolean if a field has been set.

### GetFlags

`func (o *PatchedSettingsRequest) GetFlags() PatchedSettingsRequestFlags`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *PatchedSettingsRequest) GetFlagsOk() (*PatchedSettingsRequestFlags, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *PatchedSettingsRequest) SetFlags(v PatchedSettingsRequestFlags)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *PatchedSettingsRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


