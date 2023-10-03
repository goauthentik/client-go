# AuthenticatorSMSStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
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

### NewAuthenticatorSMSStageRequest

`func NewAuthenticatorSMSStageRequest(name string, provider ProviderEnum, fromNumber string, accountSid string, auth string, ) *AuthenticatorSMSStageRequest`

NewAuthenticatorSMSStageRequest instantiates a new AuthenticatorSMSStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorSMSStageRequestWithDefaults

`func NewAuthenticatorSMSStageRequestWithDefaults() *AuthenticatorSMSStageRequest`

NewAuthenticatorSMSStageRequestWithDefaults instantiates a new AuthenticatorSMSStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthenticatorSMSStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticatorSMSStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticatorSMSStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowSet

`func (o *AuthenticatorSMSStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *AuthenticatorSMSStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *AuthenticatorSMSStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *AuthenticatorSMSStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetConfigureFlow

`func (o *AuthenticatorSMSStageRequest) GetConfigureFlow() string`

GetConfigureFlow returns the ConfigureFlow field if non-nil, zero value otherwise.

### GetConfigureFlowOk

`func (o *AuthenticatorSMSStageRequest) GetConfigureFlowOk() (*string, bool)`

GetConfigureFlowOk returns a tuple with the ConfigureFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureFlow

`func (o *AuthenticatorSMSStageRequest) SetConfigureFlow(v string)`

SetConfigureFlow sets ConfigureFlow field to given value.

### HasConfigureFlow

`func (o *AuthenticatorSMSStageRequest) HasConfigureFlow() bool`

HasConfigureFlow returns a boolean if a field has been set.

### SetConfigureFlowNil

`func (o *AuthenticatorSMSStageRequest) SetConfigureFlowNil(b bool)`

 SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil

### UnsetConfigureFlow
`func (o *AuthenticatorSMSStageRequest) UnsetConfigureFlow()`

UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
### GetFriendlyName

`func (o *AuthenticatorSMSStageRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *AuthenticatorSMSStageRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *AuthenticatorSMSStageRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *AuthenticatorSMSStageRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *AuthenticatorSMSStageRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *AuthenticatorSMSStageRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetProvider

`func (o *AuthenticatorSMSStageRequest) GetProvider() ProviderEnum`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AuthenticatorSMSStageRequest) GetProviderOk() (*ProviderEnum, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AuthenticatorSMSStageRequest) SetProvider(v ProviderEnum)`

SetProvider sets Provider field to given value.


### GetFromNumber

`func (o *AuthenticatorSMSStageRequest) GetFromNumber() string`

GetFromNumber returns the FromNumber field if non-nil, zero value otherwise.

### GetFromNumberOk

`func (o *AuthenticatorSMSStageRequest) GetFromNumberOk() (*string, bool)`

GetFromNumberOk returns a tuple with the FromNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromNumber

`func (o *AuthenticatorSMSStageRequest) SetFromNumber(v string)`

SetFromNumber sets FromNumber field to given value.


### GetAccountSid

`func (o *AuthenticatorSMSStageRequest) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *AuthenticatorSMSStageRequest) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *AuthenticatorSMSStageRequest) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.


### GetAuth

`func (o *AuthenticatorSMSStageRequest) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *AuthenticatorSMSStageRequest) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *AuthenticatorSMSStageRequest) SetAuth(v string)`

SetAuth sets Auth field to given value.


### GetAuthPassword

`func (o *AuthenticatorSMSStageRequest) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *AuthenticatorSMSStageRequest) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *AuthenticatorSMSStageRequest) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *AuthenticatorSMSStageRequest) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### GetAuthType

`func (o *AuthenticatorSMSStageRequest) GetAuthType() AuthTypeEnum`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AuthenticatorSMSStageRequest) GetAuthTypeOk() (*AuthTypeEnum, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AuthenticatorSMSStageRequest) SetAuthType(v AuthTypeEnum)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *AuthenticatorSMSStageRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetVerifyOnly

`func (o *AuthenticatorSMSStageRequest) GetVerifyOnly() bool`

GetVerifyOnly returns the VerifyOnly field if non-nil, zero value otherwise.

### GetVerifyOnlyOk

`func (o *AuthenticatorSMSStageRequest) GetVerifyOnlyOk() (*bool, bool)`

GetVerifyOnlyOk returns a tuple with the VerifyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyOnly

`func (o *AuthenticatorSMSStageRequest) SetVerifyOnly(v bool)`

SetVerifyOnly sets VerifyOnly field to given value.

### HasVerifyOnly

`func (o *AuthenticatorSMSStageRequest) HasVerifyOnly() bool`

HasVerifyOnly returns a boolean if a field has been set.

### GetMapping

`func (o *AuthenticatorSMSStageRequest) GetMapping() string`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *AuthenticatorSMSStageRequest) GetMappingOk() (*string, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *AuthenticatorSMSStageRequest) SetMapping(v string)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *AuthenticatorSMSStageRequest) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### SetMappingNil

`func (o *AuthenticatorSMSStageRequest) SetMappingNil(b bool)`

 SetMappingNil sets the value for Mapping to be an explicit nil

### UnsetMapping
`func (o *AuthenticatorSMSStageRequest) UnsetMapping()`

UnsetMapping ensures that no value is present for Mapping, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


