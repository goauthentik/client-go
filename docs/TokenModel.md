# TokenModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Provider** | [**OAuth2Provider**](OAuth2Provider.md) |  | 
**User** | [**User**](User.md) |  | 
**IsExpired** | **bool** | Check if token is expired yet. | [readonly] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**Scope** | **[]string** |  | 
**IdToken** | **string** | Get the token&#39;s id_token as JSON String | [readonly] 
**Revoked** | Pointer to **bool** |  | [optional] 

## Methods

### NewTokenModel

`func NewTokenModel(pk int32, provider OAuth2Provider, user User, isExpired bool, scope []string, idToken string, ) *TokenModel`

NewTokenModel instantiates a new TokenModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenModelWithDefaults

`func NewTokenModelWithDefaults() *TokenModel`

NewTokenModelWithDefaults instantiates a new TokenModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *TokenModel) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *TokenModel) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *TokenModel) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetProvider

`func (o *TokenModel) GetProvider() OAuth2Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *TokenModel) GetProviderOk() (*OAuth2Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *TokenModel) SetProvider(v OAuth2Provider)`

SetProvider sets Provider field to given value.


### GetUser

`func (o *TokenModel) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TokenModel) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TokenModel) SetUser(v User)`

SetUser sets User field to given value.


### GetIsExpired

`func (o *TokenModel) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *TokenModel) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *TokenModel) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.


### GetExpires

`func (o *TokenModel) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TokenModel) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TokenModel) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *TokenModel) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetScope

`func (o *TokenModel) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *TokenModel) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *TokenModel) SetScope(v []string)`

SetScope sets Scope field to given value.


### GetIdToken

`func (o *TokenModel) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *TokenModel) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *TokenModel) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.


### GetRevoked

`func (o *TokenModel) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *TokenModel) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *TokenModel) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *TokenModel) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


