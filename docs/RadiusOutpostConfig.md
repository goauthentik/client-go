# RadiusOutpostConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**ApplicationSlug** | **string** |  | 
**AuthFlowSlug** | **string** |  | 
**ClientNetworks** | Pointer to **string** | List of CIDRs (comma-separated) that clients can connect from. A more specific CIDR will match before a looser one. Clients connecting from a non-specified CIDR will be dropped. | [optional] 
**SharedSecret** | Pointer to **string** | Shared secret between clients and server to hash packets. | [optional] 
**MfaSupport** | Pointer to **bool** | When enabled, code-based multi-factor authentication can be used by appending a semicolon and the TOTP code to the password. This should only be enabled if all users that will bind to this provider have a TOTP device configured, as otherwise a password may incorrectly be rejected if it contains a semicolon. | [optional] 

## Methods

### NewRadiusOutpostConfig

`func NewRadiusOutpostConfig(pk int32, name string, applicationSlug string, authFlowSlug string, ) *RadiusOutpostConfig`

NewRadiusOutpostConfig instantiates a new RadiusOutpostConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusOutpostConfigWithDefaults

`func NewRadiusOutpostConfigWithDefaults() *RadiusOutpostConfig`

NewRadiusOutpostConfigWithDefaults instantiates a new RadiusOutpostConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *RadiusOutpostConfig) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *RadiusOutpostConfig) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *RadiusOutpostConfig) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetName

`func (o *RadiusOutpostConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RadiusOutpostConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RadiusOutpostConfig) SetName(v string)`

SetName sets Name field to given value.


### GetApplicationSlug

`func (o *RadiusOutpostConfig) GetApplicationSlug() string`

GetApplicationSlug returns the ApplicationSlug field if non-nil, zero value otherwise.

### GetApplicationSlugOk

`func (o *RadiusOutpostConfig) GetApplicationSlugOk() (*string, bool)`

GetApplicationSlugOk returns a tuple with the ApplicationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSlug

`func (o *RadiusOutpostConfig) SetApplicationSlug(v string)`

SetApplicationSlug sets ApplicationSlug field to given value.


### GetAuthFlowSlug

`func (o *RadiusOutpostConfig) GetAuthFlowSlug() string`

GetAuthFlowSlug returns the AuthFlowSlug field if non-nil, zero value otherwise.

### GetAuthFlowSlugOk

`func (o *RadiusOutpostConfig) GetAuthFlowSlugOk() (*string, bool)`

GetAuthFlowSlugOk returns a tuple with the AuthFlowSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFlowSlug

`func (o *RadiusOutpostConfig) SetAuthFlowSlug(v string)`

SetAuthFlowSlug sets AuthFlowSlug field to given value.


### GetClientNetworks

`func (o *RadiusOutpostConfig) GetClientNetworks() string`

GetClientNetworks returns the ClientNetworks field if non-nil, zero value otherwise.

### GetClientNetworksOk

`func (o *RadiusOutpostConfig) GetClientNetworksOk() (*string, bool)`

GetClientNetworksOk returns a tuple with the ClientNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNetworks

`func (o *RadiusOutpostConfig) SetClientNetworks(v string)`

SetClientNetworks sets ClientNetworks field to given value.

### HasClientNetworks

`func (o *RadiusOutpostConfig) HasClientNetworks() bool`

HasClientNetworks returns a boolean if a field has been set.

### GetSharedSecret

`func (o *RadiusOutpostConfig) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *RadiusOutpostConfig) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *RadiusOutpostConfig) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *RadiusOutpostConfig) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetMfaSupport

`func (o *RadiusOutpostConfig) GetMfaSupport() bool`

GetMfaSupport returns the MfaSupport field if non-nil, zero value otherwise.

### GetMfaSupportOk

`func (o *RadiusOutpostConfig) GetMfaSupportOk() (*bool, bool)`

GetMfaSupportOk returns a tuple with the MfaSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaSupport

`func (o *RadiusOutpostConfig) SetMfaSupport(v bool)`

SetMfaSupport sets MfaSupport field to given value.

### HasMfaSupport

`func (o *RadiusOutpostConfig) HasMfaSupport() bool`

HasMfaSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


