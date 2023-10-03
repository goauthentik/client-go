# AuthenticatorSMSStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Component** | **string** | Get object type so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**FlowSet** | Pointer to [**[]FlowSet**](FlowSet.md) |  | [optional] 
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 
**Provider** | [**ProviderEnum**](ProviderEnum.md) |  | 
**FromNumber** | **string** |  | 
**AccountSid** | **string** |  | 
**Auth** | **string** |  | 
**AuthPassword** | Pointer to **string** |  | [optional] 
**AuthType** | Pointer to [**AuthTypeEnum**](AuthTypeEnum.md) |  | [optional] 
**VerifyOnly** | Pointer to **bool** | When enabled, the Phone number is only used during enrollment to verify the users authenticity. Only a hash of the phone number is saved to ensure it is not reused in the future. | [optional] 
**Mapping** | Pointer to **NullableString** | Optionally modify the payload being sent to custom providers. | [optional] 

## Methods

### NewAuthenticatorSMSStage

`func NewAuthenticatorSMSStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, provider ProviderEnum, fromNumber string, accountSid string, auth string, ) *AuthenticatorSMSStage`

NewAuthenticatorSMSStage instantiates a new AuthenticatorSMSStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorSMSStageWithDefaults

`func NewAuthenticatorSMSStageWithDefaults() *AuthenticatorSMSStage`

NewAuthenticatorSMSStageWithDefaults instantiates a new AuthenticatorSMSStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *AuthenticatorSMSStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *AuthenticatorSMSStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *AuthenticatorSMSStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *AuthenticatorSMSStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticatorSMSStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticatorSMSStage) SetName(v string)`

SetName sets Name field to given value.


### GetComponent

`func (o *AuthenticatorSMSStage) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AuthenticatorSMSStage) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AuthenticatorSMSStage) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *AuthenticatorSMSStage) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *AuthenticatorSMSStage) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *AuthenticatorSMSStage) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *AuthenticatorSMSStage) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *AuthenticatorSMSStage) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *AuthenticatorSMSStage) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *AuthenticatorSMSStage) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *AuthenticatorSMSStage) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *AuthenticatorSMSStage) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetFlowSet

`func (o *AuthenticatorSMSStage) GetFlowSet() []FlowSet`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *AuthenticatorSMSStage) GetFlowSetOk() (*[]FlowSet, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *AuthenticatorSMSStage) SetFlowSet(v []FlowSet)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *AuthenticatorSMSStage) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetConfigureFlow

`func (o *AuthenticatorSMSStage) GetConfigureFlow() string`

GetConfigureFlow returns the ConfigureFlow field if non-nil, zero value otherwise.

### GetConfigureFlowOk

`func (o *AuthenticatorSMSStage) GetConfigureFlowOk() (*string, bool)`

GetConfigureFlowOk returns a tuple with the ConfigureFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureFlow

`func (o *AuthenticatorSMSStage) SetConfigureFlow(v string)`

SetConfigureFlow sets ConfigureFlow field to given value.

### HasConfigureFlow

`func (o *AuthenticatorSMSStage) HasConfigureFlow() bool`

HasConfigureFlow returns a boolean if a field has been set.

### SetConfigureFlowNil

`func (o *AuthenticatorSMSStage) SetConfigureFlowNil(b bool)`

 SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil

### UnsetConfigureFlow
`func (o *AuthenticatorSMSStage) UnsetConfigureFlow()`

UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
### GetFriendlyName

`func (o *AuthenticatorSMSStage) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *AuthenticatorSMSStage) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *AuthenticatorSMSStage) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *AuthenticatorSMSStage) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *AuthenticatorSMSStage) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *AuthenticatorSMSStage) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetProvider

`func (o *AuthenticatorSMSStage) GetProvider() ProviderEnum`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AuthenticatorSMSStage) GetProviderOk() (*ProviderEnum, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AuthenticatorSMSStage) SetProvider(v ProviderEnum)`

SetProvider sets Provider field to given value.


### GetFromNumber

`func (o *AuthenticatorSMSStage) GetFromNumber() string`

GetFromNumber returns the FromNumber field if non-nil, zero value otherwise.

### GetFromNumberOk

`func (o *AuthenticatorSMSStage) GetFromNumberOk() (*string, bool)`

GetFromNumberOk returns a tuple with the FromNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromNumber

`func (o *AuthenticatorSMSStage) SetFromNumber(v string)`

SetFromNumber sets FromNumber field to given value.


### GetAccountSid

`func (o *AuthenticatorSMSStage) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *AuthenticatorSMSStage) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *AuthenticatorSMSStage) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.


### GetAuth

`func (o *AuthenticatorSMSStage) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *AuthenticatorSMSStage) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *AuthenticatorSMSStage) SetAuth(v string)`

SetAuth sets Auth field to given value.


### GetAuthPassword

`func (o *AuthenticatorSMSStage) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *AuthenticatorSMSStage) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *AuthenticatorSMSStage) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *AuthenticatorSMSStage) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### GetAuthType

`func (o *AuthenticatorSMSStage) GetAuthType() AuthTypeEnum`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AuthenticatorSMSStage) GetAuthTypeOk() (*AuthTypeEnum, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AuthenticatorSMSStage) SetAuthType(v AuthTypeEnum)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *AuthenticatorSMSStage) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetVerifyOnly

`func (o *AuthenticatorSMSStage) GetVerifyOnly() bool`

GetVerifyOnly returns the VerifyOnly field if non-nil, zero value otherwise.

### GetVerifyOnlyOk

`func (o *AuthenticatorSMSStage) GetVerifyOnlyOk() (*bool, bool)`

GetVerifyOnlyOk returns a tuple with the VerifyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyOnly

`func (o *AuthenticatorSMSStage) SetVerifyOnly(v bool)`

SetVerifyOnly sets VerifyOnly field to given value.

### HasVerifyOnly

`func (o *AuthenticatorSMSStage) HasVerifyOnly() bool`

HasVerifyOnly returns a boolean if a field has been set.

### GetMapping

`func (o *AuthenticatorSMSStage) GetMapping() string`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *AuthenticatorSMSStage) GetMappingOk() (*string, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *AuthenticatorSMSStage) SetMapping(v string)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *AuthenticatorSMSStage) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### SetMappingNil

`func (o *AuthenticatorSMSStage) SetMappingNil(b bool)`

 SetMappingNil sets the value for Mapping to be an explicit nil

### UnsetMapping
`func (o *AuthenticatorSMSStage) UnsetMapping()`

UnsetMapping ensures that no value is present for Mapping, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


