# NativeLogoutChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowInfo** | Pointer to [**ContextualFlowInfo**](ContextualFlowInfo.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] [default to "ak-provider-saml-native-logout"]
**ResponseErrors** | Pointer to [**map[string][]ErrorDetail**](array.md) |  | [optional] 
**PostUrl** | Pointer to **string** |  | [optional] 
**SamlRequest** | Pointer to **string** |  | [optional] 
**RelayState** | Pointer to **string** |  | [optional] 
**ProviderName** | Pointer to **string** |  | [optional] 
**Binding** | Pointer to **string** |  | [optional] 
**RedirectUrl** | Pointer to **string** |  | [optional] 
**IsComplete** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewNativeLogoutChallenge

`func NewNativeLogoutChallenge() *NativeLogoutChallenge`

NewNativeLogoutChallenge instantiates a new NativeLogoutChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNativeLogoutChallengeWithDefaults

`func NewNativeLogoutChallengeWithDefaults() *NativeLogoutChallenge`

NewNativeLogoutChallengeWithDefaults instantiates a new NativeLogoutChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowInfo

`func (o *NativeLogoutChallenge) GetFlowInfo() ContextualFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *NativeLogoutChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *NativeLogoutChallenge) SetFlowInfo(v ContextualFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *NativeLogoutChallenge) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetComponent

`func (o *NativeLogoutChallenge) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *NativeLogoutChallenge) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *NativeLogoutChallenge) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *NativeLogoutChallenge) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetResponseErrors

`func (o *NativeLogoutChallenge) GetResponseErrors() map[string][]ErrorDetail`

GetResponseErrors returns the ResponseErrors field if non-nil, zero value otherwise.

### GetResponseErrorsOk

`func (o *NativeLogoutChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool)`

GetResponseErrorsOk returns a tuple with the ResponseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseErrors

`func (o *NativeLogoutChallenge) SetResponseErrors(v map[string][]ErrorDetail)`

SetResponseErrors sets ResponseErrors field to given value.

### HasResponseErrors

`func (o *NativeLogoutChallenge) HasResponseErrors() bool`

HasResponseErrors returns a boolean if a field has been set.

### GetPostUrl

`func (o *NativeLogoutChallenge) GetPostUrl() string`

GetPostUrl returns the PostUrl field if non-nil, zero value otherwise.

### GetPostUrlOk

`func (o *NativeLogoutChallenge) GetPostUrlOk() (*string, bool)`

GetPostUrlOk returns a tuple with the PostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUrl

`func (o *NativeLogoutChallenge) SetPostUrl(v string)`

SetPostUrl sets PostUrl field to given value.

### HasPostUrl

`func (o *NativeLogoutChallenge) HasPostUrl() bool`

HasPostUrl returns a boolean if a field has been set.

### GetSamlRequest

`func (o *NativeLogoutChallenge) GetSamlRequest() string`

GetSamlRequest returns the SamlRequest field if non-nil, zero value otherwise.

### GetSamlRequestOk

`func (o *NativeLogoutChallenge) GetSamlRequestOk() (*string, bool)`

GetSamlRequestOk returns a tuple with the SamlRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlRequest

`func (o *NativeLogoutChallenge) SetSamlRequest(v string)`

SetSamlRequest sets SamlRequest field to given value.

### HasSamlRequest

`func (o *NativeLogoutChallenge) HasSamlRequest() bool`

HasSamlRequest returns a boolean if a field has been set.

### GetRelayState

`func (o *NativeLogoutChallenge) GetRelayState() string`

GetRelayState returns the RelayState field if non-nil, zero value otherwise.

### GetRelayStateOk

`func (o *NativeLogoutChallenge) GetRelayStateOk() (*string, bool)`

GetRelayStateOk returns a tuple with the RelayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayState

`func (o *NativeLogoutChallenge) SetRelayState(v string)`

SetRelayState sets RelayState field to given value.

### HasRelayState

`func (o *NativeLogoutChallenge) HasRelayState() bool`

HasRelayState returns a boolean if a field has been set.

### GetProviderName

`func (o *NativeLogoutChallenge) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *NativeLogoutChallenge) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *NativeLogoutChallenge) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *NativeLogoutChallenge) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetBinding

`func (o *NativeLogoutChallenge) GetBinding() string`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *NativeLogoutChallenge) GetBindingOk() (*string, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *NativeLogoutChallenge) SetBinding(v string)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *NativeLogoutChallenge) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *NativeLogoutChallenge) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *NativeLogoutChallenge) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *NativeLogoutChallenge) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *NativeLogoutChallenge) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetIsComplete

`func (o *NativeLogoutChallenge) GetIsComplete() bool`

GetIsComplete returns the IsComplete field if non-nil, zero value otherwise.

### GetIsCompleteOk

`func (o *NativeLogoutChallenge) GetIsCompleteOk() (*bool, bool)`

GetIsCompleteOk returns a tuple with the IsComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComplete

`func (o *NativeLogoutChallenge) SetIsComplete(v bool)`

SetIsComplete sets IsComplete field to given value.

### HasIsComplete

`func (o *NativeLogoutChallenge) HasIsComplete() bool`

HasIsComplete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


