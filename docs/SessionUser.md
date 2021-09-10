# SessionUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**UserSelf**](UserSelf.md) |  | 
**Original** | Pointer to [**UserSelf**](UserSelf.md) |  | [optional] 

## Methods

### NewSessionUser

`func NewSessionUser(user UserSelf, ) *SessionUser`

NewSessionUser instantiates a new SessionUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionUserWithDefaults

`func NewSessionUserWithDefaults() *SessionUser`

NewSessionUserWithDefaults instantiates a new SessionUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *SessionUser) GetUser() UserSelf`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SessionUser) GetUserOk() (*UserSelf, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SessionUser) SetUser(v UserSelf)`

SetUser sets User field to given value.


### GetOriginal

`func (o *SessionUser) GetOriginal() UserSelf`

GetOriginal returns the Original field if non-nil, zero value otherwise.

### GetOriginalOk

`func (o *SessionUser) GetOriginalOk() (*UserSelf, bool)`

GetOriginalOk returns a tuple with the Original field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginal

`func (o *SessionUser) SetOriginal(v UserSelf)`

SetOriginal sets Original field to given value.

### HasOriginal

`func (o *SessionUser) HasOriginal() bool`

HasOriginal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


