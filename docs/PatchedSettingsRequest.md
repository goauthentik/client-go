# PatchedSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatars** | Pointer to **string** | Configure how authentik should show avatars for users. | [optional] 
**DefaultUserChangeName** | Pointer to **bool** | Enable the ability for users to change their name. | [optional] 
**DefaultUserChangeEmail** | Pointer to **bool** | Enable the ability for users to change their email address. | [optional] 
**DefaultUserChangeUsername** | Pointer to **bool** | Enable the ability for users to change their username. | [optional] 
**EventRetention** | Pointer to **string** | Events will be deleted after this duration.(Format: weeks&#x3D;3;days&#x3D;2;hours&#x3D;3,seconds&#x3D;2). | [optional] 
**FooterLinks** | Pointer to **interface{}** | The option configures the footer links on the flow executor pages. | [optional] 
**GdprCompliance** | Pointer to **bool** | When enabled, all the events caused by a user will be deleted upon the user&#39;s deletion. | [optional] 
**Impersonation** | Pointer to **bool** | Globally enable/disable impersonation. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


