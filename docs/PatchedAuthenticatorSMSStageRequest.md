# PatchedAuthenticatorSMSStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 
**Provider** | Pointer to [**ProviderEnum**](ProviderEnum.md) |  | [optional] 
**FromNumber** | Pointer to **string** |  | [optional] 
**AccountSid** | Pointer to **string** |  | [optional] 
**Auth** | Pointer to **string** |  | [optional] 
**AuthPassword** | Pointer to **string** |  | [optional] 
**AuthType** | Pointer to [**AuthTypeEnum**](AuthTypeEnum.md) |  | [optional] 
**VerifyOnly** | Pointer to **bool** | When enabled, the Phone number is only used during enrollment to verify the users authenticity. Only a hash of the phone number is saved to ensure it is not reused in the future. | [optional] 
**Mapping** | Pointer to **NullableString** | Optionally modify the payload being sent to custom providers. | [optional] 

## Methods

### NewPatchedAuthenticatorSMSStageRequest

`func NewPatchedAuthenticatorSMSStageRequest() *PatchedAuthenticatorSMSStageRequest`

NewPatchedAuthenticatorSMSStageRequest instantiates a new PatchedAuthenticatorSMSStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAuthenticatorSMSStageRequestWithDefaults

`func NewPatchedAuthenticatorSMSStageRequestWithDefaults() *PatchedAuthenticatorSMSStageRequest`

NewPatchedAuthenticatorSMSStageRequestWithDefaults instantiates a new PatchedAuthenticatorSMSStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedAuthenticatorSMSStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedAuthenticatorSMSStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedAuthenticatorSMSStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedAuthenticatorSMSStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlowSet

`func (o *PatchedAuthenticatorSMSStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PatchedAuthenticatorSMSStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PatchedAuthenticatorSMSStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *PatchedAuthenticatorSMSStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetConfigureFlow

`func (o *PatchedAuthenticatorSMSStageRequest) GetConfigureFlow() string`

GetConfigureFlow returns the ConfigureFlow field if non-nil, zero value otherwise.

### GetConfigureFlowOk

`func (o *PatchedAuthenticatorSMSStageRequest) GetConfigureFlowOk() (*string, bool)`

GetConfigureFlowOk returns a tuple with the ConfigureFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureFlow

`func (o *PatchedAuthenticatorSMSStageRequest) SetConfigureFlow(v string)`

SetConfigureFlow sets ConfigureFlow field to given value.

### HasConfigureFlow

`func (o *PatchedAuthenticatorSMSStageRequest) HasConfigureFlow() bool`

HasConfigureFlow returns a boolean if a field has been set.

### SetConfigureFlowNil

`func (o *PatchedAuthenticatorSMSStageRequest) SetConfigureFlowNil(b bool)`

 SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil

### UnsetConfigureFlow
`func (o *PatchedAuthenticatorSMSStageRequest) UnsetConfigureFlow()`

UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
### GetFriendlyName

`func (o *PatchedAuthenticatorSMSStageRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *PatchedAuthenticatorSMSStageRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *PatchedAuthenticatorSMSStageRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *PatchedAuthenticatorSMSStageRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *PatchedAuthenticatorSMSStageRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *PatchedAuthenticatorSMSStageRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetProvider

`func (o *PatchedAuthenticatorSMSStageRequest) GetProvider() ProviderEnum`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PatchedAuthenticatorSMSStageRequest) GetProviderOk() (*ProviderEnum, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PatchedAuthenticatorSMSStageRequest) SetProvider(v ProviderEnum)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PatchedAuthenticatorSMSStageRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetFromNumber

`func (o *PatchedAuthenticatorSMSStageRequest) GetFromNumber() string`

GetFromNumber returns the FromNumber field if non-nil, zero value otherwise.

### GetFromNumberOk

`func (o *PatchedAuthenticatorSMSStageRequest) GetFromNumberOk() (*string, bool)`

GetFromNumberOk returns a tuple with the FromNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromNumber

`func (o *PatchedAuthenticatorSMSStageRequest) SetFromNumber(v string)`

SetFromNumber sets FromNumber field to given value.

### HasFromNumber

`func (o *PatchedAuthenticatorSMSStageRequest) HasFromNumber() bool`

HasFromNumber returns a boolean if a field has been set.

### GetAccountSid

`func (o *PatchedAuthenticatorSMSStageRequest) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *PatchedAuthenticatorSMSStageRequest) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *PatchedAuthenticatorSMSStageRequest) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *PatchedAuthenticatorSMSStageRequest) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### GetAuth

`func (o *PatchedAuthenticatorSMSStageRequest) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *PatchedAuthenticatorSMSStageRequest) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *PatchedAuthenticatorSMSStageRequest) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *PatchedAuthenticatorSMSStageRequest) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetAuthPassword

`func (o *PatchedAuthenticatorSMSStageRequest) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *PatchedAuthenticatorSMSStageRequest) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *PatchedAuthenticatorSMSStageRequest) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *PatchedAuthenticatorSMSStageRequest) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### GetAuthType

`func (o *PatchedAuthenticatorSMSStageRequest) GetAuthType() AuthTypeEnum`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *PatchedAuthenticatorSMSStageRequest) GetAuthTypeOk() (*AuthTypeEnum, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *PatchedAuthenticatorSMSStageRequest) SetAuthType(v AuthTypeEnum)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *PatchedAuthenticatorSMSStageRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetVerifyOnly

`func (o *PatchedAuthenticatorSMSStageRequest) GetVerifyOnly() bool`

GetVerifyOnly returns the VerifyOnly field if non-nil, zero value otherwise.

### GetVerifyOnlyOk

`func (o *PatchedAuthenticatorSMSStageRequest) GetVerifyOnlyOk() (*bool, bool)`

GetVerifyOnlyOk returns a tuple with the VerifyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyOnly

`func (o *PatchedAuthenticatorSMSStageRequest) SetVerifyOnly(v bool)`

SetVerifyOnly sets VerifyOnly field to given value.

### HasVerifyOnly

`func (o *PatchedAuthenticatorSMSStageRequest) HasVerifyOnly() bool`

HasVerifyOnly returns a boolean if a field has been set.

### GetMapping

`func (o *PatchedAuthenticatorSMSStageRequest) GetMapping() string`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *PatchedAuthenticatorSMSStageRequest) GetMappingOk() (*string, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *PatchedAuthenticatorSMSStageRequest) SetMapping(v string)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *PatchedAuthenticatorSMSStageRequest) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### SetMappingNil

`func (o *PatchedAuthenticatorSMSStageRequest) SetMappingNil(b bool)`

 SetMappingNil sets the value for Mapping to be an explicit nil

### UnsetMapping
`func (o *PatchedAuthenticatorSMSStageRequest) UnsetMapping()`

UnsetMapping ensures that no value is present for Mapping, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


