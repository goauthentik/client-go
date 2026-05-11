# AuthenticatorEndpointGDTCStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**Credentials** | **map[string]interface{}** |  | 

## Methods

### NewAuthenticatorEndpointGDTCStageRequest

`func NewAuthenticatorEndpointGDTCStageRequest(name string, credentials map[string]interface{}, ) *AuthenticatorEndpointGDTCStageRequest`

NewAuthenticatorEndpointGDTCStageRequest instantiates a new AuthenticatorEndpointGDTCStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorEndpointGDTCStageRequestWithDefaults

`func NewAuthenticatorEndpointGDTCStageRequestWithDefaults() *AuthenticatorEndpointGDTCStageRequest`

NewAuthenticatorEndpointGDTCStageRequestWithDefaults instantiates a new AuthenticatorEndpointGDTCStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthenticatorEndpointGDTCStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticatorEndpointGDTCStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticatorEndpointGDTCStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetConfigureFlow

`func (o *AuthenticatorEndpointGDTCStageRequest) GetConfigureFlow() string`

GetConfigureFlow returns the ConfigureFlow field if non-nil, zero value otherwise.

### GetConfigureFlowOk

`func (o *AuthenticatorEndpointGDTCStageRequest) GetConfigureFlowOk() (*string, bool)`

GetConfigureFlowOk returns a tuple with the ConfigureFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureFlow

`func (o *AuthenticatorEndpointGDTCStageRequest) SetConfigureFlow(v string)`

SetConfigureFlow sets ConfigureFlow field to given value.

### HasConfigureFlow

`func (o *AuthenticatorEndpointGDTCStageRequest) HasConfigureFlow() bool`

HasConfigureFlow returns a boolean if a field has been set.

### SetConfigureFlowNil

`func (o *AuthenticatorEndpointGDTCStageRequest) SetConfigureFlowNil(b bool)`

 SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil

### UnsetConfigureFlow
`func (o *AuthenticatorEndpointGDTCStageRequest) UnsetConfigureFlow()`

UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
### GetFriendlyName

`func (o *AuthenticatorEndpointGDTCStageRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *AuthenticatorEndpointGDTCStageRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *AuthenticatorEndpointGDTCStageRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *AuthenticatorEndpointGDTCStageRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetCredentials

`func (o *AuthenticatorEndpointGDTCStageRequest) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AuthenticatorEndpointGDTCStageRequest) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AuthenticatorEndpointGDTCStageRequest) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


