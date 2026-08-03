# OAuth2DynamicClientRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PbmUuid** | **string** |  | [readonly] 
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

### NewOAuth2DynamicClientRegistration

`func NewOAuth2DynamicClientRegistration(pbmUuid string, provider int32, ) *OAuth2DynamicClientRegistration`

NewOAuth2DynamicClientRegistration instantiates a new OAuth2DynamicClientRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2DynamicClientRegistrationWithDefaults

`func NewOAuth2DynamicClientRegistrationWithDefaults() *OAuth2DynamicClientRegistration`

NewOAuth2DynamicClientRegistrationWithDefaults instantiates a new OAuth2DynamicClientRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPbmUuid

`func (o *OAuth2DynamicClientRegistration) GetPbmUuid() string`

GetPbmUuid returns the PbmUuid field if non-nil, zero value otherwise.

### GetPbmUuidOk

`func (o *OAuth2DynamicClientRegistration) GetPbmUuidOk() (*string, bool)`

GetPbmUuidOk returns a tuple with the PbmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbmUuid

`func (o *OAuth2DynamicClientRegistration) SetPbmUuid(v string)`

SetPbmUuid sets PbmUuid field to given value.


### GetProvider

`func (o *OAuth2DynamicClientRegistration) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OAuth2DynamicClientRegistration) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OAuth2DynamicClientRegistration) SetProvider(v int32)`

SetProvider sets Provider field to given value.


### GetDefaultApplicationGroup

`func (o *OAuth2DynamicClientRegistration) GetDefaultApplicationGroup() string`

GetDefaultApplicationGroup returns the DefaultApplicationGroup field if non-nil, zero value otherwise.

### GetDefaultApplicationGroupOk

`func (o *OAuth2DynamicClientRegistration) GetDefaultApplicationGroupOk() (*string, bool)`

GetDefaultApplicationGroupOk returns a tuple with the DefaultApplicationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApplicationGroup

`func (o *OAuth2DynamicClientRegistration) SetDefaultApplicationGroup(v string)`

SetDefaultApplicationGroup sets DefaultApplicationGroup field to given value.

### HasDefaultApplicationGroup

`func (o *OAuth2DynamicClientRegistration) HasDefaultApplicationGroup() bool`

HasDefaultApplicationGroup returns a boolean if a field has been set.

### GetOverrideAuthorizationFlow

`func (o *OAuth2DynamicClientRegistration) GetOverrideAuthorizationFlow() string`

GetOverrideAuthorizationFlow returns the OverrideAuthorizationFlow field if non-nil, zero value otherwise.

### GetOverrideAuthorizationFlowOk

`func (o *OAuth2DynamicClientRegistration) GetOverrideAuthorizationFlowOk() (*string, bool)`

GetOverrideAuthorizationFlowOk returns a tuple with the OverrideAuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideAuthorizationFlow

`func (o *OAuth2DynamicClientRegistration) SetOverrideAuthorizationFlow(v string)`

SetOverrideAuthorizationFlow sets OverrideAuthorizationFlow field to given value.

### HasOverrideAuthorizationFlow

`func (o *OAuth2DynamicClientRegistration) HasOverrideAuthorizationFlow() bool`

HasOverrideAuthorizationFlow returns a boolean if a field has been set.

### SetOverrideAuthorizationFlowNil

`func (o *OAuth2DynamicClientRegistration) SetOverrideAuthorizationFlowNil(b bool)`

 SetOverrideAuthorizationFlowNil sets the value for OverrideAuthorizationFlow to be an explicit nil

### UnsetOverrideAuthorizationFlow
`func (o *OAuth2DynamicClientRegistration) UnsetOverrideAuthorizationFlow()`

UnsetOverrideAuthorizationFlow ensures that no value is present for OverrideAuthorizationFlow, not even an explicit nil
### GetOverrideInvalidationFlow

`func (o *OAuth2DynamicClientRegistration) GetOverrideInvalidationFlow() string`

GetOverrideInvalidationFlow returns the OverrideInvalidationFlow field if non-nil, zero value otherwise.

### GetOverrideInvalidationFlowOk

`func (o *OAuth2DynamicClientRegistration) GetOverrideInvalidationFlowOk() (*string, bool)`

GetOverrideInvalidationFlowOk returns a tuple with the OverrideInvalidationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideInvalidationFlow

`func (o *OAuth2DynamicClientRegistration) SetOverrideInvalidationFlow(v string)`

SetOverrideInvalidationFlow sets OverrideInvalidationFlow field to given value.

### HasOverrideInvalidationFlow

`func (o *OAuth2DynamicClientRegistration) HasOverrideInvalidationFlow() bool`

HasOverrideInvalidationFlow returns a boolean if a field has been set.

### SetOverrideInvalidationFlowNil

`func (o *OAuth2DynamicClientRegistration) SetOverrideInvalidationFlowNil(b bool)`

 SetOverrideInvalidationFlowNil sets the value for OverrideInvalidationFlow to be an explicit nil

### UnsetOverrideInvalidationFlow
`func (o *OAuth2DynamicClientRegistration) UnsetOverrideInvalidationFlow()`

UnsetOverrideInvalidationFlow ensures that no value is present for OverrideInvalidationFlow, not even an explicit nil
### GetOverridePropertyMappings

`func (o *OAuth2DynamicClientRegistration) GetOverridePropertyMappings() []string`

GetOverridePropertyMappings returns the OverridePropertyMappings field if non-nil, zero value otherwise.

### GetOverridePropertyMappingsOk

`func (o *OAuth2DynamicClientRegistration) GetOverridePropertyMappingsOk() (*[]string, bool)`

GetOverridePropertyMappingsOk returns a tuple with the OverridePropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePropertyMappings

`func (o *OAuth2DynamicClientRegistration) SetOverridePropertyMappings(v []string)`

SetOverridePropertyMappings sets OverridePropertyMappings field to given value.

### HasOverridePropertyMappings

`func (o *OAuth2DynamicClientRegistration) HasOverridePropertyMappings() bool`

HasOverridePropertyMappings returns a boolean if a field has been set.

### GetAccessTokenValidity

`func (o *OAuth2DynamicClientRegistration) GetAccessTokenValidity() string`

GetAccessTokenValidity returns the AccessTokenValidity field if non-nil, zero value otherwise.

### GetAccessTokenValidityOk

`func (o *OAuth2DynamicClientRegistration) GetAccessTokenValidityOk() (*string, bool)`

GetAccessTokenValidityOk returns a tuple with the AccessTokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidity

`func (o *OAuth2DynamicClientRegistration) SetAccessTokenValidity(v string)`

SetAccessTokenValidity sets AccessTokenValidity field to given value.

### HasAccessTokenValidity

`func (o *OAuth2DynamicClientRegistration) HasAccessTokenValidity() bool`

HasAccessTokenValidity returns a boolean if a field has been set.

### GetRefreshTokenValidity

`func (o *OAuth2DynamicClientRegistration) GetRefreshTokenValidity() string`

GetRefreshTokenValidity returns the RefreshTokenValidity field if non-nil, zero value otherwise.

### GetRefreshTokenValidityOk

`func (o *OAuth2DynamicClientRegistration) GetRefreshTokenValidityOk() (*string, bool)`

GetRefreshTokenValidityOk returns a tuple with the RefreshTokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenValidity

`func (o *OAuth2DynamicClientRegistration) SetRefreshTokenValidity(v string)`

SetRefreshTokenValidity sets RefreshTokenValidity field to given value.

### HasRefreshTokenValidity

`func (o *OAuth2DynamicClientRegistration) HasRefreshTokenValidity() bool`

HasRefreshTokenValidity returns a boolean if a field has been set.

### GetAllowedGrantTypes

`func (o *OAuth2DynamicClientRegistration) GetAllowedGrantTypes() []GrantTypeEnum`

GetAllowedGrantTypes returns the AllowedGrantTypes field if non-nil, zero value otherwise.

### GetAllowedGrantTypesOk

`func (o *OAuth2DynamicClientRegistration) GetAllowedGrantTypesOk() (*[]GrantTypeEnum, bool)`

GetAllowedGrantTypesOk returns a tuple with the AllowedGrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedGrantTypes

`func (o *OAuth2DynamicClientRegistration) SetAllowedGrantTypes(v []GrantTypeEnum)`

SetAllowedGrantTypes sets AllowedGrantTypes field to given value.

### HasAllowedGrantTypes

`func (o *OAuth2DynamicClientRegistration) HasAllowedGrantTypes() bool`

HasAllowedGrantTypes returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *OAuth2DynamicClientRegistration) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *OAuth2DynamicClientRegistration) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *OAuth2DynamicClientRegistration) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *OAuth2DynamicClientRegistration) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


