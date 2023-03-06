# AuthenticatedSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Current** | **bool** | Check if session is currently active session | [readonly] 
**UserAgent** | [**AuthenticatedSessionUserAgent**](AuthenticatedSessionUserAgent.md) |  | 
**GeoIp** | [**NullableAuthenticatedSessionGeoIp**](AuthenticatedSessionGeoIp.md) |  | 
**User** | **int32** |  | 
**LastIp** | **string** |  | 
**LastUserAgent** | Pointer to **string** |  | [optional] 
**LastUsed** | **time.Time** |  | [readonly] 
**Expires** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAuthenticatedSession

`func NewAuthenticatedSession(current bool, userAgent AuthenticatedSessionUserAgent, geoIp NullableAuthenticatedSessionGeoIp, user int32, lastIp string, lastUsed time.Time, ) *AuthenticatedSession`

NewAuthenticatedSession instantiates a new AuthenticatedSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatedSessionWithDefaults

`func NewAuthenticatedSessionWithDefaults() *AuthenticatedSession`

NewAuthenticatedSessionWithDefaults instantiates a new AuthenticatedSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *AuthenticatedSession) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AuthenticatedSession) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AuthenticatedSession) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AuthenticatedSession) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCurrent

`func (o *AuthenticatedSession) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *AuthenticatedSession) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *AuthenticatedSession) SetCurrent(v bool)`

SetCurrent sets Current field to given value.


### GetUserAgent

`func (o *AuthenticatedSession) GetUserAgent() AuthenticatedSessionUserAgent`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *AuthenticatedSession) GetUserAgentOk() (*AuthenticatedSessionUserAgent, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *AuthenticatedSession) SetUserAgent(v AuthenticatedSessionUserAgent)`

SetUserAgent sets UserAgent field to given value.


### GetGeoIp

`func (o *AuthenticatedSession) GetGeoIp() AuthenticatedSessionGeoIp`

GetGeoIp returns the GeoIp field if non-nil, zero value otherwise.

### GetGeoIpOk

`func (o *AuthenticatedSession) GetGeoIpOk() (*AuthenticatedSessionGeoIp, bool)`

GetGeoIpOk returns a tuple with the GeoIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoIp

`func (o *AuthenticatedSession) SetGeoIp(v AuthenticatedSessionGeoIp)`

SetGeoIp sets GeoIp field to given value.


### SetGeoIpNil

`func (o *AuthenticatedSession) SetGeoIpNil(b bool)`

 SetGeoIpNil sets the value for GeoIp to be an explicit nil

### UnsetGeoIp
`func (o *AuthenticatedSession) UnsetGeoIp()`

UnsetGeoIp ensures that no value is present for GeoIp, not even an explicit nil
### GetUser

`func (o *AuthenticatedSession) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuthenticatedSession) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuthenticatedSession) SetUser(v int32)`

SetUser sets User field to given value.


### GetLastIp

`func (o *AuthenticatedSession) GetLastIp() string`

GetLastIp returns the LastIp field if non-nil, zero value otherwise.

### GetLastIpOk

`func (o *AuthenticatedSession) GetLastIpOk() (*string, bool)`

GetLastIpOk returns a tuple with the LastIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIp

`func (o *AuthenticatedSession) SetLastIp(v string)`

SetLastIp sets LastIp field to given value.


### GetLastUserAgent

`func (o *AuthenticatedSession) GetLastUserAgent() string`

GetLastUserAgent returns the LastUserAgent field if non-nil, zero value otherwise.

### GetLastUserAgentOk

`func (o *AuthenticatedSession) GetLastUserAgentOk() (*string, bool)`

GetLastUserAgentOk returns a tuple with the LastUserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUserAgent

`func (o *AuthenticatedSession) SetLastUserAgent(v string)`

SetLastUserAgent sets LastUserAgent field to given value.

### HasLastUserAgent

`func (o *AuthenticatedSession) HasLastUserAgent() bool`

HasLastUserAgent returns a boolean if a field has been set.

### GetLastUsed

`func (o *AuthenticatedSession) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *AuthenticatedSession) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *AuthenticatedSession) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.


### GetExpires

`func (o *AuthenticatedSession) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *AuthenticatedSession) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *AuthenticatedSession) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *AuthenticatedSession) HasExpires() bool`

HasExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


