# PatchedAccountLockdownStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DeactivateUser** | Pointer to **bool** | Deactivate the user account (set is_active to False) | [optional] 
**SetUnusablePassword** | Pointer to **bool** | Set an unusable password for the user | [optional] 
**DeleteSessions** | Pointer to **bool** | Delete all active sessions for the user | [optional] 
**RevokeTokens** | Pointer to **bool** | Revoke all tokens for the user (API, app password, recovery, verification, OAuth) | [optional] 
**SelfServiceCompletionFlow** | Pointer to **NullableString** | Flow to redirect users to after self-service lockdown. This flow should not require authentication since the user&#39;s session is deleted. | [optional] 

## Methods

### NewPatchedAccountLockdownStageRequest

`func NewPatchedAccountLockdownStageRequest() *PatchedAccountLockdownStageRequest`

NewPatchedAccountLockdownStageRequest instantiates a new PatchedAccountLockdownStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAccountLockdownStageRequestWithDefaults

`func NewPatchedAccountLockdownStageRequestWithDefaults() *PatchedAccountLockdownStageRequest`

NewPatchedAccountLockdownStageRequestWithDefaults instantiates a new PatchedAccountLockdownStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedAccountLockdownStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedAccountLockdownStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedAccountLockdownStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedAccountLockdownStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeactivateUser

`func (o *PatchedAccountLockdownStageRequest) GetDeactivateUser() bool`

GetDeactivateUser returns the DeactivateUser field if non-nil, zero value otherwise.

### GetDeactivateUserOk

`func (o *PatchedAccountLockdownStageRequest) GetDeactivateUserOk() (*bool, bool)`

GetDeactivateUserOk returns a tuple with the DeactivateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivateUser

`func (o *PatchedAccountLockdownStageRequest) SetDeactivateUser(v bool)`

SetDeactivateUser sets DeactivateUser field to given value.

### HasDeactivateUser

`func (o *PatchedAccountLockdownStageRequest) HasDeactivateUser() bool`

HasDeactivateUser returns a boolean if a field has been set.

### GetSetUnusablePassword

`func (o *PatchedAccountLockdownStageRequest) GetSetUnusablePassword() bool`

GetSetUnusablePassword returns the SetUnusablePassword field if non-nil, zero value otherwise.

### GetSetUnusablePasswordOk

`func (o *PatchedAccountLockdownStageRequest) GetSetUnusablePasswordOk() (*bool, bool)`

GetSetUnusablePasswordOk returns a tuple with the SetUnusablePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetUnusablePassword

`func (o *PatchedAccountLockdownStageRequest) SetSetUnusablePassword(v bool)`

SetSetUnusablePassword sets SetUnusablePassword field to given value.

### HasSetUnusablePassword

`func (o *PatchedAccountLockdownStageRequest) HasSetUnusablePassword() bool`

HasSetUnusablePassword returns a boolean if a field has been set.

### GetDeleteSessions

`func (o *PatchedAccountLockdownStageRequest) GetDeleteSessions() bool`

GetDeleteSessions returns the DeleteSessions field if non-nil, zero value otherwise.

### GetDeleteSessionsOk

`func (o *PatchedAccountLockdownStageRequest) GetDeleteSessionsOk() (*bool, bool)`

GetDeleteSessionsOk returns a tuple with the DeleteSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteSessions

`func (o *PatchedAccountLockdownStageRequest) SetDeleteSessions(v bool)`

SetDeleteSessions sets DeleteSessions field to given value.

### HasDeleteSessions

`func (o *PatchedAccountLockdownStageRequest) HasDeleteSessions() bool`

HasDeleteSessions returns a boolean if a field has been set.

### GetRevokeTokens

`func (o *PatchedAccountLockdownStageRequest) GetRevokeTokens() bool`

GetRevokeTokens returns the RevokeTokens field if non-nil, zero value otherwise.

### GetRevokeTokensOk

`func (o *PatchedAccountLockdownStageRequest) GetRevokeTokensOk() (*bool, bool)`

GetRevokeTokensOk returns a tuple with the RevokeTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeTokens

`func (o *PatchedAccountLockdownStageRequest) SetRevokeTokens(v bool)`

SetRevokeTokens sets RevokeTokens field to given value.

### HasRevokeTokens

`func (o *PatchedAccountLockdownStageRequest) HasRevokeTokens() bool`

HasRevokeTokens returns a boolean if a field has been set.

### GetSelfServiceCompletionFlow

`func (o *PatchedAccountLockdownStageRequest) GetSelfServiceCompletionFlow() string`

GetSelfServiceCompletionFlow returns the SelfServiceCompletionFlow field if non-nil, zero value otherwise.

### GetSelfServiceCompletionFlowOk

`func (o *PatchedAccountLockdownStageRequest) GetSelfServiceCompletionFlowOk() (*string, bool)`

GetSelfServiceCompletionFlowOk returns a tuple with the SelfServiceCompletionFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceCompletionFlow

`func (o *PatchedAccountLockdownStageRequest) SetSelfServiceCompletionFlow(v string)`

SetSelfServiceCompletionFlow sets SelfServiceCompletionFlow field to given value.

### HasSelfServiceCompletionFlow

`func (o *PatchedAccountLockdownStageRequest) HasSelfServiceCompletionFlow() bool`

HasSelfServiceCompletionFlow returns a boolean if a field has been set.

### SetSelfServiceCompletionFlowNil

`func (o *PatchedAccountLockdownStageRequest) SetSelfServiceCompletionFlowNil(b bool)`

 SetSelfServiceCompletionFlowNil sets the value for SelfServiceCompletionFlow to be an explicit nil

### UnsetSelfServiceCompletionFlow
`func (o *PatchedAccountLockdownStageRequest) UnsetSelfServiceCompletionFlow()`

UnsetSelfServiceCompletionFlow ensures that no value is present for SelfServiceCompletionFlow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


