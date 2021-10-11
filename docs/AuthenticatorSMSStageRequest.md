# AuthenticatorSMSStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FlowSet** | Pointer to [**[]FlowRequest**](FlowRequest.md) |  | [optional] 
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**Provider** | [**ProviderEnum**](ProviderEnum.md) |  | 
**FromNumber** | **string** |  | 
**TwilioAccountSid** | **string** |  | 
**TwilioAuth** | **string** |  | 

## Methods

### NewAuthenticatorSMSStageRequest

`func NewAuthenticatorSMSStageRequest(name string, provider ProviderEnum, fromNumber string, twilioAccountSid string, twilioAuth string, ) *AuthenticatorSMSStageRequest`

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

`func (o *AuthenticatorSMSStageRequest) GetFlowSet() []FlowRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *AuthenticatorSMSStageRequest) GetFlowSetOk() (*[]FlowRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *AuthenticatorSMSStageRequest) SetFlowSet(v []FlowRequest)`

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


### GetTwilioAccountSid

`func (o *AuthenticatorSMSStageRequest) GetTwilioAccountSid() string`

GetTwilioAccountSid returns the TwilioAccountSid field if non-nil, zero value otherwise.

### GetTwilioAccountSidOk

`func (o *AuthenticatorSMSStageRequest) GetTwilioAccountSidOk() (*string, bool)`

GetTwilioAccountSidOk returns a tuple with the TwilioAccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAccountSid

`func (o *AuthenticatorSMSStageRequest) SetTwilioAccountSid(v string)`

SetTwilioAccountSid sets TwilioAccountSid field to given value.


### GetTwilioAuth

`func (o *AuthenticatorSMSStageRequest) GetTwilioAuth() string`

GetTwilioAuth returns the TwilioAuth field if non-nil, zero value otherwise.

### GetTwilioAuthOk

`func (o *AuthenticatorSMSStageRequest) GetTwilioAuthOk() (*string, bool)`

GetTwilioAuthOk returns a tuple with the TwilioAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuth

`func (o *AuthenticatorSMSStageRequest) SetTwilioAuth(v string)`

SetTwilioAuth sets TwilioAuth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


