# AuthenticatorSMSStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Component** | **string** |  | [readonly] 
**VerboseName** | **string** |  | [readonly] 
**VerboseNamePlural** | **string** |  | [readonly] 
**FlowSet** | Pointer to [**[]Flow**](Flow.md) |  | [optional] 
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**Provider** | [**ProviderEnum**](ProviderEnum.md) |  | 
**FromNumber** | **string** |  | 
**TwilioAccountSid** | **string** |  | 
**TwilioAuth** | **string** |  | 

## Methods

### NewAuthenticatorSMSStage

`func NewAuthenticatorSMSStage(pk string, name string, component string, verboseName string, verboseNamePlural string, provider ProviderEnum, fromNumber string, twilioAccountSid string, twilioAuth string, ) *AuthenticatorSMSStage`

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


### GetFlowSet

`func (o *AuthenticatorSMSStage) GetFlowSet() []Flow`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *AuthenticatorSMSStage) GetFlowSetOk() (*[]Flow, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *AuthenticatorSMSStage) SetFlowSet(v []Flow)`

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


### GetTwilioAccountSid

`func (o *AuthenticatorSMSStage) GetTwilioAccountSid() string`

GetTwilioAccountSid returns the TwilioAccountSid field if non-nil, zero value otherwise.

### GetTwilioAccountSidOk

`func (o *AuthenticatorSMSStage) GetTwilioAccountSidOk() (*string, bool)`

GetTwilioAccountSidOk returns a tuple with the TwilioAccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAccountSid

`func (o *AuthenticatorSMSStage) SetTwilioAccountSid(v string)`

SetTwilioAccountSid sets TwilioAccountSid field to given value.


### GetTwilioAuth

`func (o *AuthenticatorSMSStage) GetTwilioAuth() string`

GetTwilioAuth returns the TwilioAuth field if non-nil, zero value otherwise.

### GetTwilioAuthOk

`func (o *AuthenticatorSMSStage) GetTwilioAuthOk() (*string, bool)`

GetTwilioAuthOk returns a tuple with the TwilioAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuth

`func (o *AuthenticatorSMSStage) SetTwilioAuth(v string)`

SetTwilioAuth sets TwilioAuth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


