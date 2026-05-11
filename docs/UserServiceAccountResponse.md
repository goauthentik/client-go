# UserServiceAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Token** | **string** |  | 
**UserUid** | **string** |  | 
**UserPk** | **int32** |  | 
**GroupPk** | Pointer to **string** |  | [optional] 

## Methods

### NewUserServiceAccountResponse

`func NewUserServiceAccountResponse(username string, token string, userUid string, userPk int32, ) *UserServiceAccountResponse`

NewUserServiceAccountResponse instantiates a new UserServiceAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserServiceAccountResponseWithDefaults

`func NewUserServiceAccountResponseWithDefaults() *UserServiceAccountResponse`

NewUserServiceAccountResponseWithDefaults instantiates a new UserServiceAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserServiceAccountResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserServiceAccountResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserServiceAccountResponse) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetToken

`func (o *UserServiceAccountResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UserServiceAccountResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UserServiceAccountResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserUid

`func (o *UserServiceAccountResponse) GetUserUid() string`

GetUserUid returns the UserUid field if non-nil, zero value otherwise.

### GetUserUidOk

`func (o *UserServiceAccountResponse) GetUserUidOk() (*string, bool)`

GetUserUidOk returns a tuple with the UserUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUid

`func (o *UserServiceAccountResponse) SetUserUid(v string)`

SetUserUid sets UserUid field to given value.


### GetUserPk

`func (o *UserServiceAccountResponse) GetUserPk() int32`

GetUserPk returns the UserPk field if non-nil, zero value otherwise.

### GetUserPkOk

`func (o *UserServiceAccountResponse) GetUserPkOk() (*int32, bool)`

GetUserPkOk returns a tuple with the UserPk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPk

`func (o *UserServiceAccountResponse) SetUserPk(v int32)`

SetUserPk sets UserPk field to given value.


### GetGroupPk

`func (o *UserServiceAccountResponse) GetGroupPk() string`

GetGroupPk returns the GroupPk field if non-nil, zero value otherwise.

### GetGroupPkOk

`func (o *UserServiceAccountResponse) GetGroupPkOk() (*string, bool)`

GetGroupPkOk returns a tuple with the GroupPk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPk

`func (o *UserServiceAccountResponse) SetGroupPk(v string)`

SetGroupPk sets GroupPk field to given value.

### HasGroupPk

`func (o *UserServiceAccountResponse) HasGroupPk() bool`

HasGroupPk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


