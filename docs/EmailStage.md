# EmailStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Component** | **string** | Get object type so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**FlowSet** | [**[]FlowSet**](FlowSet.md) |  | [readonly] 
**UseGlobalSettings** | Pointer to **bool** | When enabled, global Email connection settings will be used and connection settings below will be ignored. | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**UseTls** | Pointer to **bool** |  | [optional] 
**UseSsl** | Pointer to **bool** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 
**FromAddress** | Pointer to **string** |  | [optional] 
**TokenExpiry** | Pointer to **string** | Time the token sent is valid (Format: hours&#x3D;3,minutes&#x3D;17,seconds&#x3D;300). | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**ActivateUserOnSuccess** | Pointer to **bool** | Activate users upon completion of stage. | [optional] 
**RecoveryMaxAttempts** | Pointer to **int32** |  | [optional] 
**RecoveryCacheTimeout** | Pointer to **string** | The time window used to count recent account recovery attempts. If the number of attempts exceed recovery_max_attempts within this period, further attempts will be rate-limited. (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 

## Methods

### NewEmailStage

`func NewEmailStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, flowSet []FlowSet, ) *EmailStage`

NewEmailStage instantiates a new EmailStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailStageWithDefaults

`func NewEmailStageWithDefaults() *EmailStage`

NewEmailStageWithDefaults instantiates a new EmailStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *EmailStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *EmailStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *EmailStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *EmailStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *EmailStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *EmailStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *EmailStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *EmailStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *EmailStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *EmailStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *EmailStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *EmailStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *EmailStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *EmailStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *EmailStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *EmailStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *EmailStage) GetFlowSet() []FlowSet`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *EmailStage) GetFlowSetOk() (*[]FlowSet, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *EmailStage) SetFlowSet(v []FlowSet)`

SetFlowSet sets FlowSet field to given value.


### GetUseGlobalSettings

`func (o *EmailStage) GetUseGlobalSettings() bool`

GetUseGlobalSettings returns the UseGlobalSettings field if non-nil, zero value otherwise.

### GetUseGlobalSettingsOk

`func (o *EmailStage) GetUseGlobalSettingsOk() (*bool, bool)`

GetUseGlobalSettingsOk returns a tuple with the UseGlobalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGlobalSettings

`func (o *EmailStage) SetUseGlobalSettings(v bool)`

SetUseGlobalSettings sets UseGlobalSettings field to given value.

### HasUseGlobalSettings

`func (o *EmailStage) HasUseGlobalSettings() bool`

HasUseGlobalSettings returns a boolean if a field has been set.

### GetHost

`func (o *EmailStage) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EmailStage) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EmailStage) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EmailStage) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *EmailStage) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EmailStage) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EmailStage) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *EmailStage) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *EmailStage) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EmailStage) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EmailStage) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EmailStage) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUseTls

`func (o *EmailStage) GetUseTls() bool`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *EmailStage) GetUseTlsOk() (*bool, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *EmailStage) SetUseTls(v bool)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *EmailStage) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetUseSsl

`func (o *EmailStage) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *EmailStage) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *EmailStage) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *EmailStage) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### GetTimeout

`func (o *EmailStage) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *EmailStage) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *EmailStage) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *EmailStage) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetFromAddress

`func (o *EmailStage) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *EmailStage) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *EmailStage) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *EmailStage) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetTokenExpiry

`func (o *EmailStage) GetTokenExpiry() string`

GetTokenExpiry returns the TokenExpiry field if non-nil, zero value otherwise.

### GetTokenExpiryOk

`func (o *EmailStage) GetTokenExpiryOk() (*string, bool)`

GetTokenExpiryOk returns a tuple with the TokenExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiry

`func (o *EmailStage) SetTokenExpiry(v string)`

SetTokenExpiry sets TokenExpiry field to given value.

### HasTokenExpiry

`func (o *EmailStage) HasTokenExpiry() bool`

HasTokenExpiry returns a boolean if a field has been set.

### GetSubject

`func (o *EmailStage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailStage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailStage) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailStage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTemplate

`func (o *EmailStage) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EmailStage) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EmailStage) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *EmailStage) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetActivateUserOnSuccess

`func (o *EmailStage) GetActivateUserOnSuccess() bool`

GetActivateUserOnSuccess returns the ActivateUserOnSuccess field if non-nil, zero value otherwise.

### GetActivateUserOnSuccessOk

`func (o *EmailStage) GetActivateUserOnSuccessOk() (*bool, bool)`

GetActivateUserOnSuccessOk returns a tuple with the ActivateUserOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateUserOnSuccess

`func (o *EmailStage) SetActivateUserOnSuccess(v bool)`

SetActivateUserOnSuccess sets ActivateUserOnSuccess field to given value.

### HasActivateUserOnSuccess

`func (o *EmailStage) HasActivateUserOnSuccess() bool`

HasActivateUserOnSuccess returns a boolean if a field has been set.

### GetRecoveryMaxAttempts

`func (o *EmailStage) GetRecoveryMaxAttempts() int32`

GetRecoveryMaxAttempts returns the RecoveryMaxAttempts field if non-nil, zero value otherwise.

### GetRecoveryMaxAttemptsOk

`func (o *EmailStage) GetRecoveryMaxAttemptsOk() (*int32, bool)`

GetRecoveryMaxAttemptsOk returns a tuple with the RecoveryMaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryMaxAttempts

`func (o *EmailStage) SetRecoveryMaxAttempts(v int32)`

SetRecoveryMaxAttempts sets RecoveryMaxAttempts field to given value.

### HasRecoveryMaxAttempts

`func (o *EmailStage) HasRecoveryMaxAttempts() bool`

HasRecoveryMaxAttempts returns a boolean if a field has been set.

### GetRecoveryCacheTimeout

`func (o *EmailStage) GetRecoveryCacheTimeout() string`

GetRecoveryCacheTimeout returns the RecoveryCacheTimeout field if non-nil, zero value otherwise.

### GetRecoveryCacheTimeoutOk

`func (o *EmailStage) GetRecoveryCacheTimeoutOk() (*string, bool)`

GetRecoveryCacheTimeoutOk returns a tuple with the RecoveryCacheTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryCacheTimeout

`func (o *EmailStage) SetRecoveryCacheTimeout(v string)`

SetRecoveryCacheTimeout sets RecoveryCacheTimeout field to given value.

### HasRecoveryCacheTimeout

`func (o *EmailStage) HasRecoveryCacheTimeout() bool`

HasRecoveryCacheTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


