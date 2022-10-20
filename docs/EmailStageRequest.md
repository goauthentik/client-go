# EmailStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
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

### NewEmailStageRequest

`func NewEmailStageRequest(name string, ) *EmailStageRequest`

NewEmailStageRequest instantiates a new EmailStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailStageRequestWithDefaults

`func NewEmailStageRequestWithDefaults() *EmailStageRequest`

NewEmailStageRequestWithDefaults instantiates a new EmailStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EmailStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowSet

`func (o *EmailStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *EmailStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *EmailStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *EmailStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetUseGlobalSettings

`func (o *EmailStageRequest) GetUseGlobalSettings() bool`

GetUseGlobalSettings returns the UseGlobalSettings field if non-nil, zero value otherwise.

### GetUseGlobalSettingsOk

`func (o *EmailStageRequest) GetUseGlobalSettingsOk() (*bool, bool)`

GetUseGlobalSettingsOk returns a tuple with the UseGlobalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGlobalSettings

`func (o *EmailStageRequest) SetUseGlobalSettings(v bool)`

SetUseGlobalSettings sets UseGlobalSettings field to given value.

### HasUseGlobalSettings

`func (o *EmailStageRequest) HasUseGlobalSettings() bool`

HasUseGlobalSettings returns a boolean if a field has been set.

### GetHost

`func (o *EmailStageRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EmailStageRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EmailStageRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EmailStageRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *EmailStageRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EmailStageRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EmailStageRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *EmailStageRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *EmailStageRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EmailStageRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EmailStageRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EmailStageRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *EmailStageRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EmailStageRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EmailStageRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EmailStageRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUseTls

`func (o *EmailStageRequest) GetUseTls() bool`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *EmailStageRequest) GetUseTlsOk() (*bool, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *EmailStageRequest) SetUseTls(v bool)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *EmailStageRequest) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetUseSsl

`func (o *EmailStageRequest) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *EmailStageRequest) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *EmailStageRequest) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *EmailStageRequest) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### GetTimeout

`func (o *EmailStageRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *EmailStageRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *EmailStageRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *EmailStageRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetFromAddress

`func (o *EmailStageRequest) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *EmailStageRequest) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *EmailStageRequest) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *EmailStageRequest) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetTokenExpiry

`func (o *EmailStageRequest) GetTokenExpiry() int32`

GetTokenExpiry returns the TokenExpiry field if non-nil, zero value otherwise.

### GetTokenExpiryOk

`func (o *EmailStageRequest) GetTokenExpiryOk() (*int32, bool)`

GetTokenExpiryOk returns a tuple with the TokenExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiry

`func (o *EmailStageRequest) SetTokenExpiry(v int32)`

SetTokenExpiry sets TokenExpiry field to given value.

### HasTokenExpiry

`func (o *EmailStageRequest) HasTokenExpiry() bool`

HasTokenExpiry returns a boolean if a field has been set.

### GetSubject

`func (o *EmailStageRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailStageRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailStageRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailStageRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTemplate

`func (o *EmailStageRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EmailStageRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EmailStageRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *EmailStageRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetActivateUserOnSuccess

`func (o *EmailStageRequest) GetActivateUserOnSuccess() bool`

GetActivateUserOnSuccess returns the ActivateUserOnSuccess field if non-nil, zero value otherwise.

### GetActivateUserOnSuccessOk

`func (o *EmailStageRequest) GetActivateUserOnSuccessOk() (*bool, bool)`

GetActivateUserOnSuccessOk returns a tuple with the ActivateUserOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateUserOnSuccess

`func (o *EmailStageRequest) SetActivateUserOnSuccess(v bool)`

SetActivateUserOnSuccess sets ActivateUserOnSuccess field to given value.

### HasActivateUserOnSuccess

`func (o *EmailStageRequest) HasActivateUserOnSuccess() bool`

HasActivateUserOnSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


