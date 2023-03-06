# ExpiringBaseGrantModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Provider** | [**OAuth2Provider**](OAuth2Provider.md) |  | 
**User** | [**User**](User.md) |  | 
**IsExpired** | **bool** | Check if token is expired yet. | [readonly] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**Scope** | **[]string** |  | 

## Methods

### NewExpiringBaseGrantModel

`func NewExpiringBaseGrantModel(pk int32, provider OAuth2Provider, user User, isExpired bool, scope []string, ) *ExpiringBaseGrantModel`

NewExpiringBaseGrantModel instantiates a new ExpiringBaseGrantModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiringBaseGrantModelWithDefaults

`func NewExpiringBaseGrantModelWithDefaults() *ExpiringBaseGrantModel`

NewExpiringBaseGrantModelWithDefaults instantiates a new ExpiringBaseGrantModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *ExpiringBaseGrantModel) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *ExpiringBaseGrantModel) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *ExpiringBaseGrantModel) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetProvider

`func (o *ExpiringBaseGrantModel) GetProvider() OAuth2Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ExpiringBaseGrantModel) GetProviderOk() (*OAuth2Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ExpiringBaseGrantModel) SetProvider(v OAuth2Provider)`

SetProvider sets Provider field to given value.


### GetUser

`func (o *ExpiringBaseGrantModel) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ExpiringBaseGrantModel) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ExpiringBaseGrantModel) SetUser(v User)`

SetUser sets User field to given value.


### GetIsExpired

`func (o *ExpiringBaseGrantModel) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *ExpiringBaseGrantModel) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *ExpiringBaseGrantModel) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.


### GetExpires

`func (o *ExpiringBaseGrantModel) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ExpiringBaseGrantModel) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ExpiringBaseGrantModel) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *ExpiringBaseGrantModel) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetScope

`func (o *ExpiringBaseGrantModel) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ExpiringBaseGrantModel) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ExpiringBaseGrantModel) SetScope(v []string)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


