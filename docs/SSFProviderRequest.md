# SSFProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**SigningKey** | **string** | Key used to sign the SSF Events. | 
**OidcAuthProviders** | Pointer to **[]int32** |  | [optional] 
**EventRetention** | Pointer to **string** |  | [optional] 

## Methods

### NewSSFProviderRequest

`func NewSSFProviderRequest(name string, signingKey string, ) *SSFProviderRequest`

NewSSFProviderRequest instantiates a new SSFProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSFProviderRequestWithDefaults

`func NewSSFProviderRequestWithDefaults() *SSFProviderRequest`

NewSSFProviderRequestWithDefaults instantiates a new SSFProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SSFProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SSFProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SSFProviderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSigningKey

`func (o *SSFProviderRequest) GetSigningKey() string`

GetSigningKey returns the SigningKey field if non-nil, zero value otherwise.

### GetSigningKeyOk

`func (o *SSFProviderRequest) GetSigningKeyOk() (*string, bool)`

GetSigningKeyOk returns a tuple with the SigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey

`func (o *SSFProviderRequest) SetSigningKey(v string)`

SetSigningKey sets SigningKey field to given value.


### GetOidcAuthProviders

`func (o *SSFProviderRequest) GetOidcAuthProviders() []int32`

GetOidcAuthProviders returns the OidcAuthProviders field if non-nil, zero value otherwise.

### GetOidcAuthProvidersOk

`func (o *SSFProviderRequest) GetOidcAuthProvidersOk() (*[]int32, bool)`

GetOidcAuthProvidersOk returns a tuple with the OidcAuthProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAuthProviders

`func (o *SSFProviderRequest) SetOidcAuthProviders(v []int32)`

SetOidcAuthProviders sets OidcAuthProviders field to given value.

### HasOidcAuthProviders

`func (o *SSFProviderRequest) HasOidcAuthProviders() bool`

HasOidcAuthProviders returns a boolean if a field has been set.

### GetEventRetention

`func (o *SSFProviderRequest) GetEventRetention() string`

GetEventRetention returns the EventRetention field if non-nil, zero value otherwise.

### GetEventRetentionOk

`func (o *SSFProviderRequest) GetEventRetentionOk() (*string, bool)`

GetEventRetentionOk returns a tuple with the EventRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRetention

`func (o *SSFProviderRequest) SetEventRetention(v string)`

SetEventRetention sets EventRetention field to given value.

### HasEventRetention

`func (o *SSFProviderRequest) HasEventRetention() bool`

HasEventRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


