# UserAccountLockdownRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **NullableInt32** | User to lock. If omitted, locks the current user (self-service). | [optional] 

## Methods

### NewUserAccountLockdownRequest

`func NewUserAccountLockdownRequest() *UserAccountLockdownRequest`

NewUserAccountLockdownRequest instantiates a new UserAccountLockdownRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccountLockdownRequestWithDefaults

`func NewUserAccountLockdownRequestWithDefaults() *UserAccountLockdownRequest`

NewUserAccountLockdownRequestWithDefaults instantiates a new UserAccountLockdownRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserAccountLockdownRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserAccountLockdownRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserAccountLockdownRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *UserAccountLockdownRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *UserAccountLockdownRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *UserAccountLockdownRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


