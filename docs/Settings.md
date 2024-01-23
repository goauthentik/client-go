# Settings

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

### NewSettings

`func NewSettings() *Settings`

NewSettings instantiates a new Settings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsWithDefaults

`func NewSettingsWithDefaults() *Settings`

NewSettingsWithDefaults instantiates a new Settings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatars

`func (o *Settings) GetAvatars() string`

GetAvatars returns the Avatars field if non-nil, zero value otherwise.

### GetAvatarsOk

`func (o *Settings) GetAvatarsOk() (*string, bool)`

GetAvatarsOk returns a tuple with the Avatars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatars

`func (o *Settings) SetAvatars(v string)`

SetAvatars sets Avatars field to given value.

### HasAvatars

`func (o *Settings) HasAvatars() bool`

HasAvatars returns a boolean if a field has been set.

### GetDefaultUserChangeName

`func (o *Settings) GetDefaultUserChangeName() bool`

GetDefaultUserChangeName returns the DefaultUserChangeName field if non-nil, zero value otherwise.

### GetDefaultUserChangeNameOk

`func (o *Settings) GetDefaultUserChangeNameOk() (*bool, bool)`

GetDefaultUserChangeNameOk returns a tuple with the DefaultUserChangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserChangeName

`func (o *Settings) SetDefaultUserChangeName(v bool)`

SetDefaultUserChangeName sets DefaultUserChangeName field to given value.

### HasDefaultUserChangeName

`func (o *Settings) HasDefaultUserChangeName() bool`

HasDefaultUserChangeName returns a boolean if a field has been set.

### GetDefaultUserChangeEmail

`func (o *Settings) GetDefaultUserChangeEmail() bool`

GetDefaultUserChangeEmail returns the DefaultUserChangeEmail field if non-nil, zero value otherwise.

### GetDefaultUserChangeEmailOk

`func (o *Settings) GetDefaultUserChangeEmailOk() (*bool, bool)`

GetDefaultUserChangeEmailOk returns a tuple with the DefaultUserChangeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserChangeEmail

`func (o *Settings) SetDefaultUserChangeEmail(v bool)`

SetDefaultUserChangeEmail sets DefaultUserChangeEmail field to given value.

### HasDefaultUserChangeEmail

`func (o *Settings) HasDefaultUserChangeEmail() bool`

HasDefaultUserChangeEmail returns a boolean if a field has been set.

### GetDefaultUserChangeUsername

`func (o *Settings) GetDefaultUserChangeUsername() bool`

GetDefaultUserChangeUsername returns the DefaultUserChangeUsername field if non-nil, zero value otherwise.

### GetDefaultUserChangeUsernameOk

`func (o *Settings) GetDefaultUserChangeUsernameOk() (*bool, bool)`

GetDefaultUserChangeUsernameOk returns a tuple with the DefaultUserChangeUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserChangeUsername

`func (o *Settings) SetDefaultUserChangeUsername(v bool)`

SetDefaultUserChangeUsername sets DefaultUserChangeUsername field to given value.

### HasDefaultUserChangeUsername

`func (o *Settings) HasDefaultUserChangeUsername() bool`

HasDefaultUserChangeUsername returns a boolean if a field has been set.

### GetEventRetention

`func (o *Settings) GetEventRetention() string`

GetEventRetention returns the EventRetention field if non-nil, zero value otherwise.

### GetEventRetentionOk

`func (o *Settings) GetEventRetentionOk() (*string, bool)`

GetEventRetentionOk returns a tuple with the EventRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRetention

`func (o *Settings) SetEventRetention(v string)`

SetEventRetention sets EventRetention field to given value.

### HasEventRetention

`func (o *Settings) HasEventRetention() bool`

HasEventRetention returns a boolean if a field has been set.

### GetFooterLinks

`func (o *Settings) GetFooterLinks() interface{}`

GetFooterLinks returns the FooterLinks field if non-nil, zero value otherwise.

### GetFooterLinksOk

`func (o *Settings) GetFooterLinksOk() (*interface{}, bool)`

GetFooterLinksOk returns a tuple with the FooterLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterLinks

`func (o *Settings) SetFooterLinks(v interface{})`

SetFooterLinks sets FooterLinks field to given value.

### HasFooterLinks

`func (o *Settings) HasFooterLinks() bool`

HasFooterLinks returns a boolean if a field has been set.

### SetFooterLinksNil

`func (o *Settings) SetFooterLinksNil(b bool)`

 SetFooterLinksNil sets the value for FooterLinks to be an explicit nil

### UnsetFooterLinks
`func (o *Settings) UnsetFooterLinks()`

UnsetFooterLinks ensures that no value is present for FooterLinks, not even an explicit nil
### GetGdprCompliance

`func (o *Settings) GetGdprCompliance() bool`

GetGdprCompliance returns the GdprCompliance field if non-nil, zero value otherwise.

### GetGdprComplianceOk

`func (o *Settings) GetGdprComplianceOk() (*bool, bool)`

GetGdprComplianceOk returns a tuple with the GdprCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGdprCompliance

`func (o *Settings) SetGdprCompliance(v bool)`

SetGdprCompliance sets GdprCompliance field to given value.

### HasGdprCompliance

`func (o *Settings) HasGdprCompliance() bool`

HasGdprCompliance returns a boolean if a field has been set.

### GetImpersonation

`func (o *Settings) GetImpersonation() bool`

GetImpersonation returns the Impersonation field if non-nil, zero value otherwise.

### GetImpersonationOk

`func (o *Settings) GetImpersonationOk() (*bool, bool)`

GetImpersonationOk returns a tuple with the Impersonation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonation

`func (o *Settings) SetImpersonation(v bool)`

SetImpersonation sets Impersonation field to given value.

### HasImpersonation

`func (o *Settings) HasImpersonation() bool`

HasImpersonation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


