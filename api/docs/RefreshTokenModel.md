# RefreshTokenModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Provider** | [**OAuth2Provider**](OAuth2Provider.md) |  | 
**User** | [**User**](User.md) |  | 
**IsExpired** | **bool** |  | [readonly] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**Scope** | **[]string** |  | 
**IdToken** | **string** |  | [readonly] 
**Revoked** | Pointer to **bool** |  | [optional] 

## Methods

### NewRefreshTokenModel

`func NewRefreshTokenModel(pk int32, provider OAuth2Provider, user User, isExpired bool, scope []string, idToken string, ) *RefreshTokenModel`

NewRefreshTokenModel instantiates a new RefreshTokenModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshTokenModelWithDefaults

`func NewRefreshTokenModelWithDefaults() *RefreshTokenModel`

NewRefreshTokenModelWithDefaults instantiates a new RefreshTokenModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *RefreshTokenModel) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *RefreshTokenModel) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *RefreshTokenModel) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetProvider

`func (o *RefreshTokenModel) GetProvider() OAuth2Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *RefreshTokenModel) GetProviderOk() (*OAuth2Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *RefreshTokenModel) SetProvider(v OAuth2Provider)`

SetProvider sets Provider field to given value.


### GetUser

`func (o *RefreshTokenModel) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RefreshTokenModel) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RefreshTokenModel) SetUser(v User)`

SetUser sets User field to given value.


### GetIsExpired

`func (o *RefreshTokenModel) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *RefreshTokenModel) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *RefreshTokenModel) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.


### GetExpires

`func (o *RefreshTokenModel) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *RefreshTokenModel) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *RefreshTokenModel) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *RefreshTokenModel) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetScope

`func (o *RefreshTokenModel) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RefreshTokenModel) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RefreshTokenModel) SetScope(v []string)`

SetScope sets Scope field to given value.


### GetIdToken

`func (o *RefreshTokenModel) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *RefreshTokenModel) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *RefreshTokenModel) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.


### GetRevoked

`func (o *RefreshTokenModel) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *RefreshTokenModel) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *RefreshTokenModel) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *RefreshTokenModel) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


