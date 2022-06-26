# UserConsent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**User** | [**User**](User.md) |  | 
**Application** | [**Application**](Application.md) |  | 
**Permissions** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewUserConsent

`func NewUserConsent(pk int32, user User, application Application, ) *UserConsent`

NewUserConsent instantiates a new UserConsent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserConsentWithDefaults

`func NewUserConsentWithDefaults() *UserConsent`

NewUserConsentWithDefaults instantiates a new UserConsent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *UserConsent) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *UserConsent) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *UserConsent) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetExpires

`func (o *UserConsent) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *UserConsent) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *UserConsent) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *UserConsent) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetUser

`func (o *UserConsent) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserConsent) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserConsent) SetUser(v User)`

SetUser sets User field to given value.


### GetApplication

`func (o *UserConsent) GetApplication() Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *UserConsent) GetApplicationOk() (*Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *UserConsent) SetApplication(v Application)`

SetApplication sets Application field to given value.


### GetPermissions

`func (o *UserConsent) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserConsent) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserConsent) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UserConsent) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


