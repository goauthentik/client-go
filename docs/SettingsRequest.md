# SettingsRequest

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

## Methods

### NewSettingsRequest

`func NewSettingsRequest() *SettingsRequest`

NewSettingsRequest instantiates a new SettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsRequestWithDefaults

`func NewSettingsRequestWithDefaults() *SettingsRequest`

NewSettingsRequestWithDefaults instantiates a new SettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatars

`func (o *SettingsRequest) GetAvatars() string`

GetAvatars returns the Avatars field if non-nil, zero value otherwise.

### GetAvatarsOk

`func (o *SettingsRequest) GetAvatarsOk() (*string, bool)`

GetAvatarsOk returns a tuple with the Avatars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatars

`func (o *SettingsRequest) SetAvatars(v string)`

SetAvatars sets Avatars field to given value.

### HasAvatars

`func (o *SettingsRequest) HasAvatars() bool`

HasAvatars returns a boolean if a field has been set.

### GetDefaultUserChangeName

`func (o *SettingsRequest) GetDefaultUserChangeName() bool`

GetDefaultUserChangeName returns the DefaultUserChangeName field if non-nil, zero value otherwise.

### GetDefaultUserChangeNameOk

`func (o *SettingsRequest) GetDefaultUserChangeNameOk() (*bool, bool)`

GetDefaultUserChangeNameOk returns a tuple with the DefaultUserChangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserChangeName

`func (o *SettingsRequest) SetDefaultUserChangeName(v bool)`

SetDefaultUserChangeName sets DefaultUserChangeName field to given value.

### HasDefaultUserChangeName

`func (o *SettingsRequest) HasDefaultUserChangeName() bool`

HasDefaultUserChangeName returns a boolean if a field has been set.

### GetDefaultUserChangeEmail

`func (o *SettingsRequest) GetDefaultUserChangeEmail() bool`

GetDefaultUserChangeEmail returns the DefaultUserChangeEmail field if non-nil, zero value otherwise.

### GetDefaultUserChangeEmailOk

`func (o *SettingsRequest) GetDefaultUserChangeEmailOk() (*bool, bool)`

GetDefaultUserChangeEmailOk returns a tuple with the DefaultUserChangeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserChangeEmail

`func (o *SettingsRequest) SetDefaultUserChangeEmail(v bool)`

SetDefaultUserChangeEmail sets DefaultUserChangeEmail field to given value.

### HasDefaultUserChangeEmail

`func (o *SettingsRequest) HasDefaultUserChangeEmail() bool`

HasDefaultUserChangeEmail returns a boolean if a field has been set.

### GetDefaultUserChangeUsername

`func (o *SettingsRequest) GetDefaultUserChangeUsername() bool`

GetDefaultUserChangeUsername returns the DefaultUserChangeUsername field if non-nil, zero value otherwise.

### GetDefaultUserChangeUsernameOk

`func (o *SettingsRequest) GetDefaultUserChangeUsernameOk() (*bool, bool)`

GetDefaultUserChangeUsernameOk returns a tuple with the DefaultUserChangeUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserChangeUsername

`func (o *SettingsRequest) SetDefaultUserChangeUsername(v bool)`

SetDefaultUserChangeUsername sets DefaultUserChangeUsername field to given value.

### HasDefaultUserChangeUsername

`func (o *SettingsRequest) HasDefaultUserChangeUsername() bool`

HasDefaultUserChangeUsername returns a boolean if a field has been set.

### GetEventRetention

`func (o *SettingsRequest) GetEventRetention() string`

GetEventRetention returns the EventRetention field if non-nil, zero value otherwise.

### GetEventRetentionOk

`func (o *SettingsRequest) GetEventRetentionOk() (*string, bool)`

GetEventRetentionOk returns a tuple with the EventRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRetention

`func (o *SettingsRequest) SetEventRetention(v string)`

SetEventRetention sets EventRetention field to given value.

### HasEventRetention

`func (o *SettingsRequest) HasEventRetention() bool`

HasEventRetention returns a boolean if a field has been set.

### GetReputationLowerLimit

`func (o *SettingsRequest) GetReputationLowerLimit() int32`

GetReputationLowerLimit returns the ReputationLowerLimit field if non-nil, zero value otherwise.

### GetReputationLowerLimitOk

`func (o *SettingsRequest) GetReputationLowerLimitOk() (*int32, bool)`

GetReputationLowerLimitOk returns a tuple with the ReputationLowerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputationLowerLimit

`func (o *SettingsRequest) SetReputationLowerLimit(v int32)`

SetReputationLowerLimit sets ReputationLowerLimit field to given value.

### HasReputationLowerLimit

`func (o *SettingsRequest) HasReputationLowerLimit() bool`

HasReputationLowerLimit returns a boolean if a field has been set.

### GetReputationUpperLimit

`func (o *SettingsRequest) GetReputationUpperLimit() int32`

GetReputationUpperLimit returns the ReputationUpperLimit field if non-nil, zero value otherwise.

### GetReputationUpperLimitOk

`func (o *SettingsRequest) GetReputationUpperLimitOk() (*int32, bool)`

GetReputationUpperLimitOk returns a tuple with the ReputationUpperLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReputationUpperLimit

`func (o *SettingsRequest) SetReputationUpperLimit(v int32)`

SetReputationUpperLimit sets ReputationUpperLimit field to given value.

### HasReputationUpperLimit

`func (o *SettingsRequest) HasReputationUpperLimit() bool`

HasReputationUpperLimit returns a boolean if a field has been set.

### GetFooterLinks

`func (o *SettingsRequest) GetFooterLinks() interface{}`

GetFooterLinks returns the FooterLinks field if non-nil, zero value otherwise.

### GetFooterLinksOk

`func (o *SettingsRequest) GetFooterLinksOk() (*interface{}, bool)`

GetFooterLinksOk returns a tuple with the FooterLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterLinks

`func (o *SettingsRequest) SetFooterLinks(v interface{})`

SetFooterLinks sets FooterLinks field to given value.

### HasFooterLinks

`func (o *SettingsRequest) HasFooterLinks() bool`

HasFooterLinks returns a boolean if a field has been set.

### SetFooterLinksNil

`func (o *SettingsRequest) SetFooterLinksNil(b bool)`

 SetFooterLinksNil sets the value for FooterLinks to be an explicit nil

### UnsetFooterLinks
`func (o *SettingsRequest) UnsetFooterLinks()`

UnsetFooterLinks ensures that no value is present for FooterLinks, not even an explicit nil
### GetGdprCompliance

`func (o *SettingsRequest) GetGdprCompliance() bool`

GetGdprCompliance returns the GdprCompliance field if non-nil, zero value otherwise.

### GetGdprComplianceOk

`func (o *SettingsRequest) GetGdprComplianceOk() (*bool, bool)`

GetGdprComplianceOk returns a tuple with the GdprCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGdprCompliance

`func (o *SettingsRequest) SetGdprCompliance(v bool)`

SetGdprCompliance sets GdprCompliance field to given value.

### HasGdprCompliance

`func (o *SettingsRequest) HasGdprCompliance() bool`

HasGdprCompliance returns a boolean if a field has been set.

### GetImpersonation

`func (o *SettingsRequest) GetImpersonation() bool`

GetImpersonation returns the Impersonation field if non-nil, zero value otherwise.

### GetImpersonationOk

`func (o *SettingsRequest) GetImpersonationOk() (*bool, bool)`

GetImpersonationOk returns a tuple with the Impersonation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonation

`func (o *SettingsRequest) SetImpersonation(v bool)`

SetImpersonation sets Impersonation field to given value.

### HasImpersonation

`func (o *SettingsRequest) HasImpersonation() bool`

HasImpersonation returns a boolean if a field has been set.

### GetImpersonationRequireReason

`func (o *SettingsRequest) GetImpersonationRequireReason() bool`

GetImpersonationRequireReason returns the ImpersonationRequireReason field if non-nil, zero value otherwise.

### GetImpersonationRequireReasonOk

`func (o *SettingsRequest) GetImpersonationRequireReasonOk() (*bool, bool)`

GetImpersonationRequireReasonOk returns a tuple with the ImpersonationRequireReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonationRequireReason

`func (o *SettingsRequest) SetImpersonationRequireReason(v bool)`

SetImpersonationRequireReason sets ImpersonationRequireReason field to given value.

### HasImpersonationRequireReason

`func (o *SettingsRequest) HasImpersonationRequireReason() bool`

HasImpersonationRequireReason returns a boolean if a field has been set.

### GetDefaultTokenDuration

`func (o *SettingsRequest) GetDefaultTokenDuration() string`

GetDefaultTokenDuration returns the DefaultTokenDuration field if non-nil, zero value otherwise.

### GetDefaultTokenDurationOk

`func (o *SettingsRequest) GetDefaultTokenDurationOk() (*string, bool)`

GetDefaultTokenDurationOk returns a tuple with the DefaultTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTokenDuration

`func (o *SettingsRequest) SetDefaultTokenDuration(v string)`

SetDefaultTokenDuration sets DefaultTokenDuration field to given value.

### HasDefaultTokenDuration

`func (o *SettingsRequest) HasDefaultTokenDuration() bool`

HasDefaultTokenDuration returns a boolean if a field has been set.

### GetDefaultTokenLength

`func (o *SettingsRequest) GetDefaultTokenLength() int32`

GetDefaultTokenLength returns the DefaultTokenLength field if non-nil, zero value otherwise.

### GetDefaultTokenLengthOk

`func (o *SettingsRequest) GetDefaultTokenLengthOk() (*int32, bool)`

GetDefaultTokenLengthOk returns a tuple with the DefaultTokenLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTokenLength

`func (o *SettingsRequest) SetDefaultTokenLength(v int32)`

SetDefaultTokenLength sets DefaultTokenLength field to given value.

### HasDefaultTokenLength

`func (o *SettingsRequest) HasDefaultTokenLength() bool`

HasDefaultTokenLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


