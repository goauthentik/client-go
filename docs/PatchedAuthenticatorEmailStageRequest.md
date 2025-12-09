# PatchedAuthenticatorEmailStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**UseGlobalSettings** | Pointer to **bool** | When enabled, global Email connection settings will be used and connection settings below will be ignored. | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**UseTls** | Pointer to **bool** |  | [optional] 
**UseSsl** | Pointer to **bool** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 
**FromAddress** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**TokenExpiry** | Pointer to **string** | Time the token sent is valid (Format: hours&#x3D;3,minutes&#x3D;17,seconds&#x3D;300). | [optional] 
**Template** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedAuthenticatorEmailStageRequest

`func NewPatchedAuthenticatorEmailStageRequest() *PatchedAuthenticatorEmailStageRequest`

NewPatchedAuthenticatorEmailStageRequest instantiates a new PatchedAuthenticatorEmailStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAuthenticatorEmailStageRequestWithDefaults

`func NewPatchedAuthenticatorEmailStageRequestWithDefaults() *PatchedAuthenticatorEmailStageRequest`

NewPatchedAuthenticatorEmailStageRequestWithDefaults instantiates a new PatchedAuthenticatorEmailStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedAuthenticatorEmailStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedAuthenticatorEmailStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedAuthenticatorEmailStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedAuthenticatorEmailStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfigureFlow

`func (o *PatchedAuthenticatorEmailStageRequest) GetConfigureFlow() string`

GetConfigureFlow returns the ConfigureFlow field if non-nil, zero value otherwise.

### GetConfigureFlowOk

`func (o *PatchedAuthenticatorEmailStageRequest) GetConfigureFlowOk() (*string, bool)`

GetConfigureFlowOk returns a tuple with the ConfigureFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureFlow

`func (o *PatchedAuthenticatorEmailStageRequest) SetConfigureFlow(v string)`

SetConfigureFlow sets ConfigureFlow field to given value.

### HasConfigureFlow

`func (o *PatchedAuthenticatorEmailStageRequest) HasConfigureFlow() bool`

HasConfigureFlow returns a boolean if a field has been set.

### SetConfigureFlowNil

`func (o *PatchedAuthenticatorEmailStageRequest) SetConfigureFlowNil(b bool)`

 SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil

### UnsetConfigureFlow
`func (o *PatchedAuthenticatorEmailStageRequest) UnsetConfigureFlow()`

UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
### GetFriendlyName

`func (o *PatchedAuthenticatorEmailStageRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *PatchedAuthenticatorEmailStageRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *PatchedAuthenticatorEmailStageRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *PatchedAuthenticatorEmailStageRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetUseGlobalSettings

`func (o *PatchedAuthenticatorEmailStageRequest) GetUseGlobalSettings() bool`

GetUseGlobalSettings returns the UseGlobalSettings field if non-nil, zero value otherwise.

### GetUseGlobalSettingsOk

`func (o *PatchedAuthenticatorEmailStageRequest) GetUseGlobalSettingsOk() (*bool, bool)`

GetUseGlobalSettingsOk returns a tuple with the UseGlobalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGlobalSettings

`func (o *PatchedAuthenticatorEmailStageRequest) SetUseGlobalSettings(v bool)`

SetUseGlobalSettings sets UseGlobalSettings field to given value.

### HasUseGlobalSettings

`func (o *PatchedAuthenticatorEmailStageRequest) HasUseGlobalSettings() bool`

HasUseGlobalSettings returns a boolean if a field has been set.

### GetHost

`func (o *PatchedAuthenticatorEmailStageRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PatchedAuthenticatorEmailStageRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PatchedAuthenticatorEmailStageRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PatchedAuthenticatorEmailStageRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *PatchedAuthenticatorEmailStageRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PatchedAuthenticatorEmailStageRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PatchedAuthenticatorEmailStageRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PatchedAuthenticatorEmailStageRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *PatchedAuthenticatorEmailStageRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchedAuthenticatorEmailStageRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchedAuthenticatorEmailStageRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchedAuthenticatorEmailStageRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *PatchedAuthenticatorEmailStageRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchedAuthenticatorEmailStageRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchedAuthenticatorEmailStageRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchedAuthenticatorEmailStageRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUseTls

`func (o *PatchedAuthenticatorEmailStageRequest) GetUseTls() bool`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *PatchedAuthenticatorEmailStageRequest) GetUseTlsOk() (*bool, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *PatchedAuthenticatorEmailStageRequest) SetUseTls(v bool)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *PatchedAuthenticatorEmailStageRequest) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetUseSsl

`func (o *PatchedAuthenticatorEmailStageRequest) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *PatchedAuthenticatorEmailStageRequest) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *PatchedAuthenticatorEmailStageRequest) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *PatchedAuthenticatorEmailStageRequest) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### GetTimeout

`func (o *PatchedAuthenticatorEmailStageRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PatchedAuthenticatorEmailStageRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PatchedAuthenticatorEmailStageRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PatchedAuthenticatorEmailStageRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetFromAddress

`func (o *PatchedAuthenticatorEmailStageRequest) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *PatchedAuthenticatorEmailStageRequest) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *PatchedAuthenticatorEmailStageRequest) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *PatchedAuthenticatorEmailStageRequest) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetSubject

`func (o *PatchedAuthenticatorEmailStageRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PatchedAuthenticatorEmailStageRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PatchedAuthenticatorEmailStageRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PatchedAuthenticatorEmailStageRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTokenExpiry

`func (o *PatchedAuthenticatorEmailStageRequest) GetTokenExpiry() string`

GetTokenExpiry returns the TokenExpiry field if non-nil, zero value otherwise.

### GetTokenExpiryOk

`func (o *PatchedAuthenticatorEmailStageRequest) GetTokenExpiryOk() (*string, bool)`

GetTokenExpiryOk returns a tuple with the TokenExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiry

`func (o *PatchedAuthenticatorEmailStageRequest) SetTokenExpiry(v string)`

SetTokenExpiry sets TokenExpiry field to given value.

### HasTokenExpiry

`func (o *PatchedAuthenticatorEmailStageRequest) HasTokenExpiry() bool`

HasTokenExpiry returns a boolean if a field has been set.

### GetTemplate

`func (o *PatchedAuthenticatorEmailStageRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PatchedAuthenticatorEmailStageRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PatchedAuthenticatorEmailStageRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PatchedAuthenticatorEmailStageRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


