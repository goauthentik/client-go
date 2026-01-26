# AuthenticatorDuoStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**ApiHostname** | **string** |  | 
**AdminIntegrationKey** | Pointer to **string** |  | [optional] 
**AdminSecretKey** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthenticatorDuoStageRequest

`func NewAuthenticatorDuoStageRequest(name string, clientId string, clientSecret string, apiHostname string, ) *AuthenticatorDuoStageRequest`

NewAuthenticatorDuoStageRequest instantiates a new AuthenticatorDuoStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorDuoStageRequestWithDefaults

`func NewAuthenticatorDuoStageRequestWithDefaults() *AuthenticatorDuoStageRequest`

NewAuthenticatorDuoStageRequestWithDefaults instantiates a new AuthenticatorDuoStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthenticatorDuoStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticatorDuoStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticatorDuoStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowSet

`func (o *AuthenticatorDuoStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *AuthenticatorDuoStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *AuthenticatorDuoStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *AuthenticatorDuoStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetConfigureFlow

`func (o *AuthenticatorDuoStageRequest) GetConfigureFlow() string`

GetConfigureFlow returns the ConfigureFlow field if non-nil, zero value otherwise.

### GetConfigureFlowOk

`func (o *AuthenticatorDuoStageRequest) GetConfigureFlowOk() (*string, bool)`

GetConfigureFlowOk returns a tuple with the ConfigureFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureFlow

`func (o *AuthenticatorDuoStageRequest) SetConfigureFlow(v string)`

SetConfigureFlow sets ConfigureFlow field to given value.

### HasConfigureFlow

`func (o *AuthenticatorDuoStageRequest) HasConfigureFlow() bool`

HasConfigureFlow returns a boolean if a field has been set.

### SetConfigureFlowNil

`func (o *AuthenticatorDuoStageRequest) SetConfigureFlowNil(b bool)`

 SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil

### UnsetConfigureFlow
`func (o *AuthenticatorDuoStageRequest) UnsetConfigureFlow()`

UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
### GetFriendlyName

`func (o *AuthenticatorDuoStageRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *AuthenticatorDuoStageRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *AuthenticatorDuoStageRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *AuthenticatorDuoStageRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *AuthenticatorDuoStageRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *AuthenticatorDuoStageRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetClientId

`func (o *AuthenticatorDuoStageRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AuthenticatorDuoStageRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AuthenticatorDuoStageRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *AuthenticatorDuoStageRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AuthenticatorDuoStageRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AuthenticatorDuoStageRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetApiHostname

`func (o *AuthenticatorDuoStageRequest) GetApiHostname() string`

GetApiHostname returns the ApiHostname field if non-nil, zero value otherwise.

### GetApiHostnameOk

`func (o *AuthenticatorDuoStageRequest) GetApiHostnameOk() (*string, bool)`

GetApiHostnameOk returns a tuple with the ApiHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiHostname

`func (o *AuthenticatorDuoStageRequest) SetApiHostname(v string)`

SetApiHostname sets ApiHostname field to given value.


### GetAdminIntegrationKey

`func (o *AuthenticatorDuoStageRequest) GetAdminIntegrationKey() string`

GetAdminIntegrationKey returns the AdminIntegrationKey field if non-nil, zero value otherwise.

### GetAdminIntegrationKeyOk

`func (o *AuthenticatorDuoStageRequest) GetAdminIntegrationKeyOk() (*string, bool)`

GetAdminIntegrationKeyOk returns a tuple with the AdminIntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminIntegrationKey

`func (o *AuthenticatorDuoStageRequest) SetAdminIntegrationKey(v string)`

SetAdminIntegrationKey sets AdminIntegrationKey field to given value.

### HasAdminIntegrationKey

`func (o *AuthenticatorDuoStageRequest) HasAdminIntegrationKey() bool`

HasAdminIntegrationKey returns a boolean if a field has been set.

### GetAdminSecretKey

`func (o *AuthenticatorDuoStageRequest) GetAdminSecretKey() string`

GetAdminSecretKey returns the AdminSecretKey field if non-nil, zero value otherwise.

### GetAdminSecretKeyOk

`func (o *AuthenticatorDuoStageRequest) GetAdminSecretKeyOk() (*string, bool)`

GetAdminSecretKeyOk returns a tuple with the AdminSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSecretKey

`func (o *AuthenticatorDuoStageRequest) SetAdminSecretKey(v string)`

SetAdminSecretKey sets AdminSecretKey field to given value.

### HasAdminSecretKey

`func (o *AuthenticatorDuoStageRequest) HasAdminSecretKey() bool`

HasAdminSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


