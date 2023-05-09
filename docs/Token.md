# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Managed** | Pointer to **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 
**Identifier** | **string** |  | 
**Intent** | Pointer to [**IntentEnum**](IntentEnum.md) |  | [optional] 
**User** | Pointer to **int32** |  | [optional] 
**UserObj** | [**User**](User.md) |  | [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **time.Time** |  | [optional] 
**Expiring** | Pointer to **bool** |  | [optional] 

## Methods

### NewToken

`func NewToken(pk string, identifier string, userObj User, ) *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *Token) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *Token) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *Token) SetPk(v string)`

SetPk sets Pk field to given value.


### GetManaged

`func (o *Token) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *Token) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *Token) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *Token) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *Token) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *Token) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetIdentifier

`func (o *Token) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Token) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Token) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetIntent

`func (o *Token) GetIntent() IntentEnum`

GetIntent returns the Intent field if non-nil, zero value otherwise.

### GetIntentOk

`func (o *Token) GetIntentOk() (*IntentEnum, bool)`

GetIntentOk returns a tuple with the Intent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntent

`func (o *Token) SetIntent(v IntentEnum)`

SetIntent sets Intent field to given value.

### HasIntent

`func (o *Token) HasIntent() bool`

HasIntent returns a boolean if a field has been set.

### GetUser

`func (o *Token) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Token) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Token) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *Token) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserObj

`func (o *Token) GetUserObj() User`

GetUserObj returns the UserObj field if non-nil, zero value otherwise.

### GetUserObjOk

`func (o *Token) GetUserObjOk() (*User, bool)`

GetUserObjOk returns a tuple with the UserObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObj

`func (o *Token) SetUserObj(v User)`

SetUserObj sets UserObj field to given value.


### GetDescription

`func (o *Token) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Token) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Token) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Token) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpires

`func (o *Token) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *Token) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *Token) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *Token) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetExpiring

`func (o *Token) GetExpiring() bool`

GetExpiring returns the Expiring field if non-nil, zero value otherwise.

### GetExpiringOk

`func (o *Token) GetExpiringOk() (*bool, bool)`

GetExpiringOk returns a tuple with the Expiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiring

`func (o *Token) SetExpiring(v bool)`

SetExpiring sets Expiring field to given value.

### HasExpiring

`func (o *Token) HasExpiring() bool`

HasExpiring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


