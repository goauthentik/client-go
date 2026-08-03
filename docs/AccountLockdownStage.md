# AccountLockdownStage

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
**DeactivateUser** | Pointer to **bool** | Deactivate the user account (set is_active to False) | [optional] 
**SetUnusablePassword** | Pointer to **bool** | Set an unusable password for the user | [optional] 
**DeleteSessions** | Pointer to **bool** | Delete all active sessions for the user | [optional] 
**RevokeTokens** | Pointer to **bool** | Revoke all tokens for the user (API, app password, recovery, verification, OAuth) | [optional] 
**SelfServiceCompletionFlow** | Pointer to **NullableString** | Flow to redirect users to after self-service lockdown. This flow should not require authentication since the user&#39;s session is deleted. | [optional] 

## Methods

### NewAccountLockdownStage

`func NewAccountLockdownStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, flowSet []FlowSet, ) *AccountLockdownStage`

NewAccountLockdownStage instantiates a new AccountLockdownStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountLockdownStageWithDefaults

`func NewAccountLockdownStageWithDefaults() *AccountLockdownStage`

NewAccountLockdownStageWithDefaults instantiates a new AccountLockdownStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *AccountLockdownStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *AccountLockdownStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *AccountLockdownStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *AccountLockdownStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountLockdownStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountLockdownStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *AccountLockdownStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AccountLockdownStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AccountLockdownStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *AccountLockdownStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *AccountLockdownStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *AccountLockdownStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *AccountLockdownStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *AccountLockdownStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *AccountLockdownStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *AccountLockdownStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *AccountLockdownStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *AccountLockdownStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *AccountLockdownStage) GetFlowSet() []FlowSet`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *AccountLockdownStage) GetFlowSetOk() (*[]FlowSet, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *AccountLockdownStage) SetFlowSet(v []FlowSet)`

SetFlowSet sets FlowSet field to given value.


### GetDeactivateUser

`func (o *AccountLockdownStage) GetDeactivateUser() bool`

GetDeactivateUser returns the DeactivateUser field if non-nil, zero value otherwise.

### GetDeactivateUserOk

`func (o *AccountLockdownStage) GetDeactivateUserOk() (*bool, bool)`

GetDeactivateUserOk returns a tuple with the DeactivateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivateUser

`func (o *AccountLockdownStage) SetDeactivateUser(v bool)`

SetDeactivateUser sets DeactivateUser field to given value.

### HasDeactivateUser

`func (o *AccountLockdownStage) HasDeactivateUser() bool`

HasDeactivateUser returns a boolean if a field has been set.

### GetSetUnusablePassword

`func (o *AccountLockdownStage) GetSetUnusablePassword() bool`

GetSetUnusablePassword returns the SetUnusablePassword field if non-nil, zero value otherwise.

### GetSetUnusablePasswordOk

`func (o *AccountLockdownStage) GetSetUnusablePasswordOk() (*bool, bool)`

GetSetUnusablePasswordOk returns a tuple with the SetUnusablePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetUnusablePassword

`func (o *AccountLockdownStage) SetSetUnusablePassword(v bool)`

SetSetUnusablePassword sets SetUnusablePassword field to given value.

### HasSetUnusablePassword

`func (o *AccountLockdownStage) HasSetUnusablePassword() bool`

HasSetUnusablePassword returns a boolean if a field has been set.

### GetDeleteSessions

`func (o *AccountLockdownStage) GetDeleteSessions() bool`

GetDeleteSessions returns the DeleteSessions field if non-nil, zero value otherwise.

### GetDeleteSessionsOk

`func (o *AccountLockdownStage) GetDeleteSessionsOk() (*bool, bool)`

GetDeleteSessionsOk returns a tuple with the DeleteSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteSessions

`func (o *AccountLockdownStage) SetDeleteSessions(v bool)`

SetDeleteSessions sets DeleteSessions field to given value.

### HasDeleteSessions

`func (o *AccountLockdownStage) HasDeleteSessions() bool`

HasDeleteSessions returns a boolean if a field has been set.

### GetRevokeTokens

`func (o *AccountLockdownStage) GetRevokeTokens() bool`

GetRevokeTokens returns the RevokeTokens field if non-nil, zero value otherwise.

### GetRevokeTokensOk

`func (o *AccountLockdownStage) GetRevokeTokensOk() (*bool, bool)`

GetRevokeTokensOk returns a tuple with the RevokeTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeTokens

`func (o *AccountLockdownStage) SetRevokeTokens(v bool)`

SetRevokeTokens sets RevokeTokens field to given value.

### HasRevokeTokens

`func (o *AccountLockdownStage) HasRevokeTokens() bool`

HasRevokeTokens returns a boolean if a field has been set.

### GetSelfServiceCompletionFlow

`func (o *AccountLockdownStage) GetSelfServiceCompletionFlow() string`

GetSelfServiceCompletionFlow returns the SelfServiceCompletionFlow field if non-nil, zero value otherwise.

### GetSelfServiceCompletionFlowOk

`func (o *AccountLockdownStage) GetSelfServiceCompletionFlowOk() (*string, bool)`

GetSelfServiceCompletionFlowOk returns a tuple with the SelfServiceCompletionFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceCompletionFlow

`func (o *AccountLockdownStage) SetSelfServiceCompletionFlow(v string)`

SetSelfServiceCompletionFlow sets SelfServiceCompletionFlow field to given value.

### HasSelfServiceCompletionFlow

`func (o *AccountLockdownStage) HasSelfServiceCompletionFlow() bool`

HasSelfServiceCompletionFlow returns a boolean if a field has been set.

### SetSelfServiceCompletionFlowNil

`func (o *AccountLockdownStage) SetSelfServiceCompletionFlowNil(b bool)`

 SetSelfServiceCompletionFlowNil sets the value for SelfServiceCompletionFlow to be an explicit nil

### UnsetSelfServiceCompletionFlow
`func (o *AccountLockdownStage) UnsetSelfServiceCompletionFlow()`

UnsetSelfServiceCompletionFlow ensures that no value is present for SelfServiceCompletionFlow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


