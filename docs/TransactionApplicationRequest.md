# TransactionApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | [**ApplicationRequest**](ApplicationRequest.md) |  | 
**ProviderModel** | [**ProviderModelEnum**](ProviderModelEnum.md) |  | 
**Provider** | [**ModelRequest**](ModelRequest.md) |  | 
**PolicyBindings** | Pointer to [**[]TransactionPolicyBindingRequest**](TransactionPolicyBindingRequest.md) |  | [optional] 

## Methods

### NewTransactionApplicationRequest

`func NewTransactionApplicationRequest(app ApplicationRequest, providerModel ProviderModelEnum, provider ModelRequest, ) *TransactionApplicationRequest`

NewTransactionApplicationRequest instantiates a new TransactionApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionApplicationRequestWithDefaults

`func NewTransactionApplicationRequestWithDefaults() *TransactionApplicationRequest`

NewTransactionApplicationRequestWithDefaults instantiates a new TransactionApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *TransactionApplicationRequest) GetApp() ApplicationRequest`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *TransactionApplicationRequest) GetAppOk() (*ApplicationRequest, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *TransactionApplicationRequest) SetApp(v ApplicationRequest)`

SetApp sets App field to given value.


### GetProviderModel

`func (o *TransactionApplicationRequest) GetProviderModel() ProviderModelEnum`

GetProviderModel returns the ProviderModel field if non-nil, zero value otherwise.

### GetProviderModelOk

`func (o *TransactionApplicationRequest) GetProviderModelOk() (*ProviderModelEnum, bool)`

GetProviderModelOk returns a tuple with the ProviderModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderModel

`func (o *TransactionApplicationRequest) SetProviderModel(v ProviderModelEnum)`

SetProviderModel sets ProviderModel field to given value.


### GetProvider

`func (o *TransactionApplicationRequest) GetProvider() ModelRequest`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *TransactionApplicationRequest) GetProviderOk() (*ModelRequest, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *TransactionApplicationRequest) SetProvider(v ModelRequest)`

SetProvider sets Provider field to given value.


### GetPolicyBindings

`func (o *TransactionApplicationRequest) GetPolicyBindings() []TransactionPolicyBindingRequest`

GetPolicyBindings returns the PolicyBindings field if non-nil, zero value otherwise.

### GetPolicyBindingsOk

`func (o *TransactionApplicationRequest) GetPolicyBindingsOk() (*[]TransactionPolicyBindingRequest, bool)`

GetPolicyBindingsOk returns a tuple with the PolicyBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyBindings

`func (o *TransactionApplicationRequest) SetPolicyBindings(v []TransactionPolicyBindingRequest)`

SetPolicyBindings sets PolicyBindings field to given value.

### HasPolicyBindings

`func (o *TransactionApplicationRequest) HasPolicyBindings() bool`

HasPolicyBindings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


