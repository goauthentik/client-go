# PasswordStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Backends** | [**[]BackendsEnum**](BackendsEnum.md) | Selection of backends to test the password against. | 
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**FailedAttemptsBeforeCancel** | Pointer to **int32** | How many attempts a user has before the flow is canceled. This only cancels the flow, it does not lock the user&#39;s password. | [optional] 
**FailedAttemptsBeforeLockout** | Pointer to **int32** | How many consecutive failed attempts lock the user&#39;s password until an administrator unlocks it. Set to 0 to never lock. | [optional] 
**LastAttemptWarningMessage** | Pointer to **string** | Warning shown when the user has one password attempt remaining. Leave blank to show no warning. | [optional] 
**LockoutMessage** | Pointer to **string** | Message shown when the user&#39;s password has been locked. Leave blank to show a generic authentication error. | [optional] 
**AllowShowPassword** | Pointer to **bool** | When enabled, provides a &#39;show password&#39; button with the password input field. | [optional] 

## Methods

### NewPasswordStageRequest

`func NewPasswordStageRequest(name string, backends []BackendsEnum, ) *PasswordStageRequest`

NewPasswordStageRequest instantiates a new PasswordStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordStageRequestWithDefaults

`func NewPasswordStageRequestWithDefaults() *PasswordStageRequest`

NewPasswordStageRequestWithDefaults instantiates a new PasswordStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PasswordStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PasswordStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PasswordStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetBackends

`func (o *PasswordStageRequest) GetBackends() []BackendsEnum`

GetBackends returns the Backends field if non-nil, zero value otherwise.

### GetBackendsOk

`func (o *PasswordStageRequest) GetBackendsOk() (*[]BackendsEnum, bool)`

GetBackendsOk returns a tuple with the Backends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackends

`func (o *PasswordStageRequest) SetBackends(v []BackendsEnum)`

SetBackends sets Backends field to given value.


### GetConfigureFlow

`func (o *PasswordStageRequest) GetConfigureFlow() string`

GetConfigureFlow returns the ConfigureFlow field if non-nil, zero value otherwise.

### GetConfigureFlowOk

`func (o *PasswordStageRequest) GetConfigureFlowOk() (*string, bool)`

GetConfigureFlowOk returns a tuple with the ConfigureFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureFlow

`func (o *PasswordStageRequest) SetConfigureFlow(v string)`

SetConfigureFlow sets ConfigureFlow field to given value.

### HasConfigureFlow

`func (o *PasswordStageRequest) HasConfigureFlow() bool`

HasConfigureFlow returns a boolean if a field has been set.

### SetConfigureFlowNil

`func (o *PasswordStageRequest) SetConfigureFlowNil(b bool)`

 SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil

### UnsetConfigureFlow
`func (o *PasswordStageRequest) UnsetConfigureFlow()`

UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
### GetFailedAttemptsBeforeCancel

`func (o *PasswordStageRequest) GetFailedAttemptsBeforeCancel() int32`

GetFailedAttemptsBeforeCancel returns the FailedAttemptsBeforeCancel field if non-nil, zero value otherwise.

### GetFailedAttemptsBeforeCancelOk

`func (o *PasswordStageRequest) GetFailedAttemptsBeforeCancelOk() (*int32, bool)`

GetFailedAttemptsBeforeCancelOk returns a tuple with the FailedAttemptsBeforeCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAttemptsBeforeCancel

`func (o *PasswordStageRequest) SetFailedAttemptsBeforeCancel(v int32)`

SetFailedAttemptsBeforeCancel sets FailedAttemptsBeforeCancel field to given value.

### HasFailedAttemptsBeforeCancel

`func (o *PasswordStageRequest) HasFailedAttemptsBeforeCancel() bool`

HasFailedAttemptsBeforeCancel returns a boolean if a field has been set.

### GetFailedAttemptsBeforeLockout

`func (o *PasswordStageRequest) GetFailedAttemptsBeforeLockout() int32`

GetFailedAttemptsBeforeLockout returns the FailedAttemptsBeforeLockout field if non-nil, zero value otherwise.

### GetFailedAttemptsBeforeLockoutOk

`func (o *PasswordStageRequest) GetFailedAttemptsBeforeLockoutOk() (*int32, bool)`

GetFailedAttemptsBeforeLockoutOk returns a tuple with the FailedAttemptsBeforeLockout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAttemptsBeforeLockout

`func (o *PasswordStageRequest) SetFailedAttemptsBeforeLockout(v int32)`

SetFailedAttemptsBeforeLockout sets FailedAttemptsBeforeLockout field to given value.

### HasFailedAttemptsBeforeLockout

`func (o *PasswordStageRequest) HasFailedAttemptsBeforeLockout() bool`

HasFailedAttemptsBeforeLockout returns a boolean if a field has been set.

### GetLastAttemptWarningMessage

`func (o *PasswordStageRequest) GetLastAttemptWarningMessage() string`

GetLastAttemptWarningMessage returns the LastAttemptWarningMessage field if non-nil, zero value otherwise.

### GetLastAttemptWarningMessageOk

`func (o *PasswordStageRequest) GetLastAttemptWarningMessageOk() (*string, bool)`

GetLastAttemptWarningMessageOk returns a tuple with the LastAttemptWarningMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAttemptWarningMessage

`func (o *PasswordStageRequest) SetLastAttemptWarningMessage(v string)`

SetLastAttemptWarningMessage sets LastAttemptWarningMessage field to given value.

### HasLastAttemptWarningMessage

`func (o *PasswordStageRequest) HasLastAttemptWarningMessage() bool`

HasLastAttemptWarningMessage returns a boolean if a field has been set.

### GetLockoutMessage

`func (o *PasswordStageRequest) GetLockoutMessage() string`

GetLockoutMessage returns the LockoutMessage field if non-nil, zero value otherwise.

### GetLockoutMessageOk

`func (o *PasswordStageRequest) GetLockoutMessageOk() (*string, bool)`

GetLockoutMessageOk returns a tuple with the LockoutMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutMessage

`func (o *PasswordStageRequest) SetLockoutMessage(v string)`

SetLockoutMessage sets LockoutMessage field to given value.

### HasLockoutMessage

`func (o *PasswordStageRequest) HasLockoutMessage() bool`

HasLockoutMessage returns a boolean if a field has been set.

### GetAllowShowPassword

`func (o *PasswordStageRequest) GetAllowShowPassword() bool`

GetAllowShowPassword returns the AllowShowPassword field if non-nil, zero value otherwise.

### GetAllowShowPasswordOk

`func (o *PasswordStageRequest) GetAllowShowPasswordOk() (*bool, bool)`

GetAllowShowPasswordOk returns a tuple with the AllowShowPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowShowPassword

`func (o *PasswordStageRequest) SetAllowShowPassword(v bool)`

SetAllowShowPassword sets AllowShowPassword field to given value.

### HasAllowShowPassword

`func (o *PasswordStageRequest) HasAllowShowPassword() bool`

HasAllowShowPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


