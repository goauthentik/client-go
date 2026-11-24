# PatchedRACProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** | Flow used for authentication when the associated application is accessed by an un-authenticated user. | [optional] 
**AuthorizationFlow** | Pointer to **string** | Flow used when authorizing this provider. | [optional] 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**Settings** | Pointer to **map[string]interface{}** |  | [optional] 
**AccessGroup** | Pointer to **NullableString** | Only devices in this access group can be accessed through this provider. When left empty, every device the user has access to can be accessed. | [optional] 
**MaximumConnections** | Pointer to **int32** | Maximum concurrent connections to a single device. Can be set to -1 to disable the limit. | [optional] 
**AuthMode** | Pointer to [**RACProviderAuthModeEnum**](RACProviderAuthModeEnum.md) |  | [optional] 
**ConnectionExpiry** | Pointer to **string** | Determines how long a session lasts. Default of 0 means that the sessions lasts until the browser is closed. (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3) | [optional] 
**DeleteTokenOnDisconnect** | Pointer to **bool** | When set to true, connection tokens will be deleted upon disconnect. | [optional] 

## Methods

### NewPatchedRACProviderRequest

`func NewPatchedRACProviderRequest() *PatchedRACProviderRequest`

NewPatchedRACProviderRequest instantiates a new PatchedRACProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedRACProviderRequestWithDefaults

`func NewPatchedRACProviderRequestWithDefaults() *PatchedRACProviderRequest`

NewPatchedRACProviderRequestWithDefaults instantiates a new PatchedRACProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedRACProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedRACProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedRACProviderRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedRACProviderRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *PatchedRACProviderRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *PatchedRACProviderRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *PatchedRACProviderRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *PatchedRACProviderRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *PatchedRACProviderRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *PatchedRACProviderRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetAuthorizationFlow

`func (o *PatchedRACProviderRequest) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *PatchedRACProviderRequest) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *PatchedRACProviderRequest) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.

### HasAuthorizationFlow

`func (o *PatchedRACProviderRequest) HasAuthorizationFlow() bool`

HasAuthorizationFlow returns a boolean if a field has been set.

### GetPropertyMappings

`func (o *PatchedRACProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *PatchedRACProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *PatchedRACProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *PatchedRACProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetSettings

`func (o *PatchedRACProviderRequest) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PatchedRACProviderRequest) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PatchedRACProviderRequest) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PatchedRACProviderRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetAccessGroup

`func (o *PatchedRACProviderRequest) GetAccessGroup() string`

GetAccessGroup returns the AccessGroup field if non-nil, zero value otherwise.

### GetAccessGroupOk

`func (o *PatchedRACProviderRequest) GetAccessGroupOk() (*string, bool)`

GetAccessGroupOk returns a tuple with the AccessGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroup

`func (o *PatchedRACProviderRequest) SetAccessGroup(v string)`

SetAccessGroup sets AccessGroup field to given value.

### HasAccessGroup

`func (o *PatchedRACProviderRequest) HasAccessGroup() bool`

HasAccessGroup returns a boolean if a field has been set.

### SetAccessGroupNil

`func (o *PatchedRACProviderRequest) SetAccessGroupNil(b bool)`

 SetAccessGroupNil sets the value for AccessGroup to be an explicit nil

### UnsetAccessGroup
`func (o *PatchedRACProviderRequest) UnsetAccessGroup()`

UnsetAccessGroup ensures that no value is present for AccessGroup, not even an explicit nil
### GetMaximumConnections

`func (o *PatchedRACProviderRequest) GetMaximumConnections() int32`

GetMaximumConnections returns the MaximumConnections field if non-nil, zero value otherwise.

### GetMaximumConnectionsOk

`func (o *PatchedRACProviderRequest) GetMaximumConnectionsOk() (*int32, bool)`

GetMaximumConnectionsOk returns a tuple with the MaximumConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConnections

`func (o *PatchedRACProviderRequest) SetMaximumConnections(v int32)`

SetMaximumConnections sets MaximumConnections field to given value.

### HasMaximumConnections

`func (o *PatchedRACProviderRequest) HasMaximumConnections() bool`

HasMaximumConnections returns a boolean if a field has been set.

### GetAuthMode

`func (o *PatchedRACProviderRequest) GetAuthMode() RACProviderAuthModeEnum`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *PatchedRACProviderRequest) GetAuthModeOk() (*RACProviderAuthModeEnum, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *PatchedRACProviderRequest) SetAuthMode(v RACProviderAuthModeEnum)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *PatchedRACProviderRequest) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetConnectionExpiry

`func (o *PatchedRACProviderRequest) GetConnectionExpiry() string`

GetConnectionExpiry returns the ConnectionExpiry field if non-nil, zero value otherwise.

### GetConnectionExpiryOk

`func (o *PatchedRACProviderRequest) GetConnectionExpiryOk() (*string, bool)`

GetConnectionExpiryOk returns a tuple with the ConnectionExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionExpiry

`func (o *PatchedRACProviderRequest) SetConnectionExpiry(v string)`

SetConnectionExpiry sets ConnectionExpiry field to given value.

### HasConnectionExpiry

`func (o *PatchedRACProviderRequest) HasConnectionExpiry() bool`

HasConnectionExpiry returns a boolean if a field has been set.

### GetDeleteTokenOnDisconnect

`func (o *PatchedRACProviderRequest) GetDeleteTokenOnDisconnect() bool`

GetDeleteTokenOnDisconnect returns the DeleteTokenOnDisconnect field if non-nil, zero value otherwise.

### GetDeleteTokenOnDisconnectOk

`func (o *PatchedRACProviderRequest) GetDeleteTokenOnDisconnectOk() (*bool, bool)`

GetDeleteTokenOnDisconnectOk returns a tuple with the DeleteTokenOnDisconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteTokenOnDisconnect

`func (o *PatchedRACProviderRequest) SetDeleteTokenOnDisconnect(v bool)`

SetDeleteTokenOnDisconnect sets DeleteTokenOnDisconnect field to given value.

### HasDeleteTokenOnDisconnect

`func (o *PatchedRACProviderRequest) HasDeleteTokenOnDisconnect() bool`

HasDeleteTokenOnDisconnect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


