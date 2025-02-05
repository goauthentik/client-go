# PatchedSSFProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**SigningKey** | Pointer to **string** | Key used to sign the SSF Events. | [optional] 
**OidcAuthProviders** | Pointer to **[]int32** |  | [optional] 
**EventRetention** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedSSFProviderRequest

`func NewPatchedSSFProviderRequest() *PatchedSSFProviderRequest`

NewPatchedSSFProviderRequest instantiates a new PatchedSSFProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSSFProviderRequestWithDefaults

`func NewPatchedSSFProviderRequestWithDefaults() *PatchedSSFProviderRequest`

NewPatchedSSFProviderRequestWithDefaults instantiates a new PatchedSSFProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedSSFProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedSSFProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedSSFProviderRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedSSFProviderRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSigningKey

`func (o *PatchedSSFProviderRequest) GetSigningKey() string`

GetSigningKey returns the SigningKey field if non-nil, zero value otherwise.

### GetSigningKeyOk

`func (o *PatchedSSFProviderRequest) GetSigningKeyOk() (*string, bool)`

GetSigningKeyOk returns a tuple with the SigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey

`func (o *PatchedSSFProviderRequest) SetSigningKey(v string)`

SetSigningKey sets SigningKey field to given value.

### HasSigningKey

`func (o *PatchedSSFProviderRequest) HasSigningKey() bool`

HasSigningKey returns a boolean if a field has been set.

### GetOidcAuthProviders

`func (o *PatchedSSFProviderRequest) GetOidcAuthProviders() []int32`

GetOidcAuthProviders returns the OidcAuthProviders field if non-nil, zero value otherwise.

### GetOidcAuthProvidersOk

`func (o *PatchedSSFProviderRequest) GetOidcAuthProvidersOk() (*[]int32, bool)`

GetOidcAuthProvidersOk returns a tuple with the OidcAuthProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAuthProviders

`func (o *PatchedSSFProviderRequest) SetOidcAuthProviders(v []int32)`

SetOidcAuthProviders sets OidcAuthProviders field to given value.

### HasOidcAuthProviders

`func (o *PatchedSSFProviderRequest) HasOidcAuthProviders() bool`

HasOidcAuthProviders returns a boolean if a field has been set.

### GetEventRetention

`func (o *PatchedSSFProviderRequest) GetEventRetention() string`

GetEventRetention returns the EventRetention field if non-nil, zero value otherwise.

### GetEventRetentionOk

`func (o *PatchedSSFProviderRequest) GetEventRetentionOk() (*string, bool)`

GetEventRetentionOk returns a tuple with the EventRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRetention

`func (o *PatchedSSFProviderRequest) SetEventRetention(v string)`

SetEventRetention sets EventRetention field to given value.

### HasEventRetention

`func (o *PatchedSSFProviderRequest) HasEventRetention() bool`

HasEventRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


