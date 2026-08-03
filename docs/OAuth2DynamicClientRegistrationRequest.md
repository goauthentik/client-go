# OAuth2DynamicClientRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **int32** |  | 
**DefaultApplicationGroup** | Pointer to **string** | Group to assign to automatically created applications. | [optional] 
**OverrideAuthorizationFlow** | Pointer to **NullableString** | Authorization flow applied to dynamically registered clients. | [optional] 
**OverrideInvalidationFlow** | Pointer to **NullableString** |  | [optional] 
**OverridePropertyMappings** | Pointer to **[]string** | Scope mappings applied to dynamically registered clients. | [optional] 
**AccessTokenValidity** | Pointer to **string** | Maximum access token validity for registered clients (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**RefreshTokenValidity** | Pointer to **string** | Maximum refresh token validity for registered clients (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**AllowedGrantTypes** | Pointer to [**[]GrantTypeEnum**](GrantTypeEnum.md) | If empty, all grant types are allowed. | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 

## Methods

### NewOAuth2DynamicClientRegistrationRequest

`func NewOAuth2DynamicClientRegistrationRequest(provider int32, ) *OAuth2DynamicClientRegistrationRequest`

NewOAuth2DynamicClientRegistrationRequest instantiates a new OAuth2DynamicClientRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2DynamicClientRegistrationRequestWithDefaults

`func NewOAuth2DynamicClientRegistrationRequestWithDefaults() *OAuth2DynamicClientRegistrationRequest`

NewOAuth2DynamicClientRegistrationRequestWithDefaults instantiates a new OAuth2DynamicClientRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *OAuth2DynamicClientRegistrationRequest) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OAuth2DynamicClientRegistrationRequest) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OAuth2DynamicClientRegistrationRequest) SetProvider(v int32)`

SetProvider sets Provider field to given value.


### GetDefaultApplicationGroup

`func (o *OAuth2DynamicClientRegistrationRequest) GetDefaultApplicationGroup() string`

GetDefaultApplicationGroup returns the DefaultApplicationGroup field if non-nil, zero value otherwise.

### GetDefaultApplicationGroupOk

`func (o *OAuth2DynamicClientRegistrationRequest) GetDefaultApplicationGroupOk() (*string, bool)`

GetDefaultApplicationGroupOk returns a tuple with the DefaultApplicationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApplicationGroup

`func (o *OAuth2DynamicClientRegistrationRequest) SetDefaultApplicationGroup(v string)`

SetDefaultApplicationGroup sets DefaultApplicationGroup field to given value.

### HasDefaultApplicationGroup

`func (o *OAuth2DynamicClientRegistrationRequest) HasDefaultApplicationGroup() bool`

HasDefaultApplicationGroup returns a boolean if a field has been set.

### GetOverrideAuthorizationFlow

`func (o *OAuth2DynamicClientRegistrationRequest) GetOverrideAuthorizationFlow() string`

GetOverrideAuthorizationFlow returns the OverrideAuthorizationFlow field if non-nil, zero value otherwise.

### GetOverrideAuthorizationFlowOk

`func (o *OAuth2DynamicClientRegistrationRequest) GetOverrideAuthorizationFlowOk() (*string, bool)`

GetOverrideAuthorizationFlowOk returns a tuple with the OverrideAuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideAuthorizationFlow

`func (o *OAuth2DynamicClientRegistrationRequest) SetOverrideAuthorizationFlow(v string)`

SetOverrideAuthorizationFlow sets OverrideAuthorizationFlow field to given value.

### HasOverrideAuthorizationFlow

`func (o *OAuth2DynamicClientRegistrationRequest) HasOverrideAuthorizationFlow() bool`

HasOverrideAuthorizationFlow returns a boolean if a field has been set.

### SetOverrideAuthorizationFlowNil

`func (o *OAuth2DynamicClientRegistrationRequest) SetOverrideAuthorizationFlowNil(b bool)`

 SetOverrideAuthorizationFlowNil sets the value for OverrideAuthorizationFlow to be an explicit nil

### UnsetOverrideAuthorizationFlow
`func (o *OAuth2DynamicClientRegistrationRequest) UnsetOverrideAuthorizationFlow()`

UnsetOverrideAuthorizationFlow ensures that no value is present for OverrideAuthorizationFlow, not even an explicit nil
### GetOverrideInvalidationFlow

`func (o *OAuth2DynamicClientRegistrationRequest) GetOverrideInvalidationFlow() string`

GetOverrideInvalidationFlow returns the OverrideInvalidationFlow field if non-nil, zero value otherwise.

### GetOverrideInvalidationFlowOk

`func (o *OAuth2DynamicClientRegistrationRequest) GetOverrideInvalidationFlowOk() (*string, bool)`

GetOverrideInvalidationFlowOk returns a tuple with the OverrideInvalidationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideInvalidationFlow

`func (o *OAuth2DynamicClientRegistrationRequest) SetOverrideInvalidationFlow(v string)`

SetOverrideInvalidationFlow sets OverrideInvalidationFlow field to given value.

### HasOverrideInvalidationFlow

`func (o *OAuth2DynamicClientRegistrationRequest) HasOverrideInvalidationFlow() bool`

HasOverrideInvalidationFlow returns a boolean if a field has been set.

### SetOverrideInvalidationFlowNil

`func (o *OAuth2DynamicClientRegistrationRequest) SetOverrideInvalidationFlowNil(b bool)`

 SetOverrideInvalidationFlowNil sets the value for OverrideInvalidationFlow to be an explicit nil

### UnsetOverrideInvalidationFlow
`func (o *OAuth2DynamicClientRegistrationRequest) UnsetOverrideInvalidationFlow()`

UnsetOverrideInvalidationFlow ensures that no value is present for OverrideInvalidationFlow, not even an explicit nil
### GetOverridePropertyMappings

`func (o *OAuth2DynamicClientRegistrationRequest) GetOverridePropertyMappings() []string`

GetOverridePropertyMappings returns the OverridePropertyMappings field if non-nil, zero value otherwise.

### GetOverridePropertyMappingsOk

`func (o *OAuth2DynamicClientRegistrationRequest) GetOverridePropertyMappingsOk() (*[]string, bool)`

GetOverridePropertyMappingsOk returns a tuple with the OverridePropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePropertyMappings

`func (o *OAuth2DynamicClientRegistrationRequest) SetOverridePropertyMappings(v []string)`

SetOverridePropertyMappings sets OverridePropertyMappings field to given value.

### HasOverridePropertyMappings

`func (o *OAuth2DynamicClientRegistrationRequest) HasOverridePropertyMappings() bool`

HasOverridePropertyMappings returns a boolean if a field has been set.

### GetAccessTokenValidity

`func (o *OAuth2DynamicClientRegistrationRequest) GetAccessTokenValidity() string`

GetAccessTokenValidity returns the AccessTokenValidity field if non-nil, zero value otherwise.

### GetAccessTokenValidityOk

`func (o *OAuth2DynamicClientRegistrationRequest) GetAccessTokenValidityOk() (*string, bool)`

GetAccessTokenValidityOk returns a tuple with the AccessTokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidity

`func (o *OAuth2DynamicClientRegistrationRequest) SetAccessTokenValidity(v string)`

SetAccessTokenValidity sets AccessTokenValidity field to given value.

### HasAccessTokenValidity

`func (o *OAuth2DynamicClientRegistrationRequest) HasAccessTokenValidity() bool`

HasAccessTokenValidity returns a boolean if a field has been set.

### GetRefreshTokenValidity

`func (o *OAuth2DynamicClientRegistrationRequest) GetRefreshTokenValidity() string`

GetRefreshTokenValidity returns the RefreshTokenValidity field if non-nil, zero value otherwise.

### GetRefreshTokenValidityOk

`func (o *OAuth2DynamicClientRegistrationRequest) GetRefreshTokenValidityOk() (*string, bool)`

GetRefreshTokenValidityOk returns a tuple with the RefreshTokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenValidity

`func (o *OAuth2DynamicClientRegistrationRequest) SetRefreshTokenValidity(v string)`

SetRefreshTokenValidity sets RefreshTokenValidity field to given value.

### HasRefreshTokenValidity

`func (o *OAuth2DynamicClientRegistrationRequest) HasRefreshTokenValidity() bool`

HasRefreshTokenValidity returns a boolean if a field has been set.

### GetAllowedGrantTypes

`func (o *OAuth2DynamicClientRegistrationRequest) GetAllowedGrantTypes() []GrantTypeEnum`

GetAllowedGrantTypes returns the AllowedGrantTypes field if non-nil, zero value otherwise.

### GetAllowedGrantTypesOk

`func (o *OAuth2DynamicClientRegistrationRequest) GetAllowedGrantTypesOk() (*[]GrantTypeEnum, bool)`

GetAllowedGrantTypesOk returns a tuple with the AllowedGrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedGrantTypes

`func (o *OAuth2DynamicClientRegistrationRequest) SetAllowedGrantTypes(v []GrantTypeEnum)`

SetAllowedGrantTypes sets AllowedGrantTypes field to given value.

### HasAllowedGrantTypes

`func (o *OAuth2DynamicClientRegistrationRequest) HasAllowedGrantTypes() bool`

HasAllowedGrantTypes returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *OAuth2DynamicClientRegistrationRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *OAuth2DynamicClientRegistrationRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *OAuth2DynamicClientRegistrationRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *OAuth2DynamicClientRegistrationRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


