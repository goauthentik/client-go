# PatchedAuthenticatorDuoStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**ConfigureFlow** | Pointer to **NullableString** | Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage. | [optional] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**ApiHostname** | Pointer to **string** |  | [optional] 
**AdminIntegrationKey** | Pointer to **string** |  | [optional] 
**AdminSecretKey** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedAuthenticatorDuoStageRequest

`func NewPatchedAuthenticatorDuoStageRequest() *PatchedAuthenticatorDuoStageRequest`

NewPatchedAuthenticatorDuoStageRequest instantiates a new PatchedAuthenticatorDuoStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAuthenticatorDuoStageRequestWithDefaults

`func NewPatchedAuthenticatorDuoStageRequestWithDefaults() *PatchedAuthenticatorDuoStageRequest`

NewPatchedAuthenticatorDuoStageRequestWithDefaults instantiates a new PatchedAuthenticatorDuoStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedAuthenticatorDuoStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedAuthenticatorDuoStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedAuthenticatorDuoStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedAuthenticatorDuoStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlowSet

`func (o *PatchedAuthenticatorDuoStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PatchedAuthenticatorDuoStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PatchedAuthenticatorDuoStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *PatchedAuthenticatorDuoStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetConfigureFlow

`func (o *PatchedAuthenticatorDuoStageRequest) GetConfigureFlow() string`

GetConfigureFlow returns the ConfigureFlow field if non-nil, zero value otherwise.

### GetConfigureFlowOk

`func (o *PatchedAuthenticatorDuoStageRequest) GetConfigureFlowOk() (*string, bool)`

GetConfigureFlowOk returns a tuple with the ConfigureFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureFlow

`func (o *PatchedAuthenticatorDuoStageRequest) SetConfigureFlow(v string)`

SetConfigureFlow sets ConfigureFlow field to given value.

### HasConfigureFlow

`func (o *PatchedAuthenticatorDuoStageRequest) HasConfigureFlow() bool`

HasConfigureFlow returns a boolean if a field has been set.

### SetConfigureFlowNil

`func (o *PatchedAuthenticatorDuoStageRequest) SetConfigureFlowNil(b bool)`

 SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil

### UnsetConfigureFlow
`func (o *PatchedAuthenticatorDuoStageRequest) UnsetConfigureFlow()`

UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
### GetFriendlyName

`func (o *PatchedAuthenticatorDuoStageRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *PatchedAuthenticatorDuoStageRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *PatchedAuthenticatorDuoStageRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *PatchedAuthenticatorDuoStageRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *PatchedAuthenticatorDuoStageRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *PatchedAuthenticatorDuoStageRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetClientId

`func (o *PatchedAuthenticatorDuoStageRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PatchedAuthenticatorDuoStageRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PatchedAuthenticatorDuoStageRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PatchedAuthenticatorDuoStageRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *PatchedAuthenticatorDuoStageRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PatchedAuthenticatorDuoStageRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PatchedAuthenticatorDuoStageRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PatchedAuthenticatorDuoStageRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetApiHostname

`func (o *PatchedAuthenticatorDuoStageRequest) GetApiHostname() string`

GetApiHostname returns the ApiHostname field if non-nil, zero value otherwise.

### GetApiHostnameOk

`func (o *PatchedAuthenticatorDuoStageRequest) GetApiHostnameOk() (*string, bool)`

GetApiHostnameOk returns a tuple with the ApiHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiHostname

`func (o *PatchedAuthenticatorDuoStageRequest) SetApiHostname(v string)`

SetApiHostname sets ApiHostname field to given value.

### HasApiHostname

`func (o *PatchedAuthenticatorDuoStageRequest) HasApiHostname() bool`

HasApiHostname returns a boolean if a field has been set.

### GetAdminIntegrationKey

`func (o *PatchedAuthenticatorDuoStageRequest) GetAdminIntegrationKey() string`

GetAdminIntegrationKey returns the AdminIntegrationKey field if non-nil, zero value otherwise.

### GetAdminIntegrationKeyOk

`func (o *PatchedAuthenticatorDuoStageRequest) GetAdminIntegrationKeyOk() (*string, bool)`

GetAdminIntegrationKeyOk returns a tuple with the AdminIntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminIntegrationKey

`func (o *PatchedAuthenticatorDuoStageRequest) SetAdminIntegrationKey(v string)`

SetAdminIntegrationKey sets AdminIntegrationKey field to given value.

### HasAdminIntegrationKey

`func (o *PatchedAuthenticatorDuoStageRequest) HasAdminIntegrationKey() bool`

HasAdminIntegrationKey returns a boolean if a field has been set.

### GetAdminSecretKey

`func (o *PatchedAuthenticatorDuoStageRequest) GetAdminSecretKey() string`

GetAdminSecretKey returns the AdminSecretKey field if non-nil, zero value otherwise.

### GetAdminSecretKeyOk

`func (o *PatchedAuthenticatorDuoStageRequest) GetAdminSecretKeyOk() (*string, bool)`

GetAdminSecretKeyOk returns a tuple with the AdminSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSecretKey

`func (o *PatchedAuthenticatorDuoStageRequest) SetAdminSecretKey(v string)`

SetAdminSecretKey sets AdminSecretKey field to given value.

### HasAdminSecretKey

`func (o *PatchedAuthenticatorDuoStageRequest) HasAdminSecretKey() bool`

HasAdminSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


