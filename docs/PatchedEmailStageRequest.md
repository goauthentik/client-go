# PatchedEmailStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**UseGlobalSettings** | Pointer to **bool** | When enabled, global Email connection settings will be used and connection settings below will be ignored. | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**UseTls** | Pointer to **bool** |  | [optional] 
**UseSsl** | Pointer to **bool** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 
**FromAddress** | Pointer to **string** |  | [optional] 
**TokenExpiry** | Pointer to **int32** | Time in minutes the token sent is valid. | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**ActivateUserOnSuccess** | Pointer to **bool** | Activate users upon completion of stage. | [optional] 

## Methods

### NewPatchedEmailStageRequest

`func NewPatchedEmailStageRequest() *PatchedEmailStageRequest`

NewPatchedEmailStageRequest instantiates a new PatchedEmailStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEmailStageRequestWithDefaults

`func NewPatchedEmailStageRequestWithDefaults() *PatchedEmailStageRequest`

NewPatchedEmailStageRequestWithDefaults instantiates a new PatchedEmailStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedEmailStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEmailStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEmailStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEmailStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlowSet

`func (o *PatchedEmailStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PatchedEmailStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PatchedEmailStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *PatchedEmailStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetUseGlobalSettings

`func (o *PatchedEmailStageRequest) GetUseGlobalSettings() bool`

GetUseGlobalSettings returns the UseGlobalSettings field if non-nil, zero value otherwise.

### GetUseGlobalSettingsOk

`func (o *PatchedEmailStageRequest) GetUseGlobalSettingsOk() (*bool, bool)`

GetUseGlobalSettingsOk returns a tuple with the UseGlobalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGlobalSettings

`func (o *PatchedEmailStageRequest) SetUseGlobalSettings(v bool)`

SetUseGlobalSettings sets UseGlobalSettings field to given value.

### HasUseGlobalSettings

`func (o *PatchedEmailStageRequest) HasUseGlobalSettings() bool`

HasUseGlobalSettings returns a boolean if a field has been set.

### GetHost

`func (o *PatchedEmailStageRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PatchedEmailStageRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PatchedEmailStageRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PatchedEmailStageRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *PatchedEmailStageRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PatchedEmailStageRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PatchedEmailStageRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PatchedEmailStageRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *PatchedEmailStageRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchedEmailStageRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchedEmailStageRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchedEmailStageRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *PatchedEmailStageRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchedEmailStageRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchedEmailStageRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchedEmailStageRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUseTls

`func (o *PatchedEmailStageRequest) GetUseTls() bool`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *PatchedEmailStageRequest) GetUseTlsOk() (*bool, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *PatchedEmailStageRequest) SetUseTls(v bool)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *PatchedEmailStageRequest) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetUseSsl

`func (o *PatchedEmailStageRequest) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *PatchedEmailStageRequest) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *PatchedEmailStageRequest) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *PatchedEmailStageRequest) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### GetTimeout

`func (o *PatchedEmailStageRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PatchedEmailStageRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PatchedEmailStageRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PatchedEmailStageRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetFromAddress

`func (o *PatchedEmailStageRequest) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *PatchedEmailStageRequest) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *PatchedEmailStageRequest) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *PatchedEmailStageRequest) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetTokenExpiry

`func (o *PatchedEmailStageRequest) GetTokenExpiry() int32`

GetTokenExpiry returns the TokenExpiry field if non-nil, zero value otherwise.

### GetTokenExpiryOk

`func (o *PatchedEmailStageRequest) GetTokenExpiryOk() (*int32, bool)`

GetTokenExpiryOk returns a tuple with the TokenExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiry

`func (o *PatchedEmailStageRequest) SetTokenExpiry(v int32)`

SetTokenExpiry sets TokenExpiry field to given value.

### HasTokenExpiry

`func (o *PatchedEmailStageRequest) HasTokenExpiry() bool`

HasTokenExpiry returns a boolean if a field has been set.

### GetSubject

`func (o *PatchedEmailStageRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PatchedEmailStageRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PatchedEmailStageRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PatchedEmailStageRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTemplate

`func (o *PatchedEmailStageRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PatchedEmailStageRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PatchedEmailStageRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PatchedEmailStageRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetActivateUserOnSuccess

`func (o *PatchedEmailStageRequest) GetActivateUserOnSuccess() bool`

GetActivateUserOnSuccess returns the ActivateUserOnSuccess field if non-nil, zero value otherwise.

### GetActivateUserOnSuccessOk

`func (o *PatchedEmailStageRequest) GetActivateUserOnSuccessOk() (*bool, bool)`

GetActivateUserOnSuccessOk returns a tuple with the ActivateUserOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateUserOnSuccess

`func (o *PatchedEmailStageRequest) SetActivateUserOnSuccess(v bool)`

SetActivateUserOnSuccess sets ActivateUserOnSuccess field to given value.

### HasActivateUserOnSuccess

`func (o *PatchedEmailStageRequest) HasActivateUserOnSuccess() bool`

HasActivateUserOnSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


