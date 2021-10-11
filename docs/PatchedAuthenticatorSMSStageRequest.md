# PatchedAuthenticatorSMSStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FlowSet** | Pointer to [**[]FlowRequest**](FlowRequest.md) |  | [optional] 
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**Provider** | Pointer to [**ProviderEnum**](ProviderEnum.md) |  | [optional] 
**FromNumber** | Pointer to **string** |  | [optional] 
**TwilioAccountSid** | Pointer to **string** |  | [optional] 
**TwilioAuth** | Pointer to **string** |  | [optional] 

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

`func (o *PatchedAuthenticatorSMSStageRequest) GetFlowSet() []FlowRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PatchedAuthenticatorSMSStageRequest) GetFlowSetOk() (*[]FlowRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PatchedAuthenticatorSMSStageRequest) SetFlowSet(v []FlowRequest)`

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

### GetTwilioAccountSid

`func (o *PatchedAuthenticatorSMSStageRequest) GetTwilioAccountSid() string`

GetTwilioAccountSid returns the TwilioAccountSid field if non-nil, zero value otherwise.

### GetTwilioAccountSidOk

`func (o *PatchedAuthenticatorSMSStageRequest) GetTwilioAccountSidOk() (*string, bool)`

GetTwilioAccountSidOk returns a tuple with the TwilioAccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAccountSid

`func (o *PatchedAuthenticatorSMSStageRequest) SetTwilioAccountSid(v string)`

SetTwilioAccountSid sets TwilioAccountSid field to given value.

### HasTwilioAccountSid

`func (o *PatchedAuthenticatorSMSStageRequest) HasTwilioAccountSid() bool`

HasTwilioAccountSid returns a boolean if a field has been set.

### GetTwilioAuth

`func (o *PatchedAuthenticatorSMSStageRequest) GetTwilioAuth() string`

GetTwilioAuth returns the TwilioAuth field if non-nil, zero value otherwise.

### GetTwilioAuthOk

`func (o *PatchedAuthenticatorSMSStageRequest) GetTwilioAuthOk() (*string, bool)`

GetTwilioAuthOk returns a tuple with the TwilioAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuth

`func (o *PatchedAuthenticatorSMSStageRequest) SetTwilioAuth(v string)`

SetTwilioAuth sets TwilioAuth field to given value.

### HasTwilioAuth

`func (o *PatchedAuthenticatorSMSStageRequest) HasTwilioAuth() bool`

HasTwilioAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


