# UserSwitchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**UserSwitchActionEnum**](UserSwitchActionEnum.md) |  | [optional] [default to USERSWITCHACTIONENUM_SWITCH]
**UserPk** | Pointer to **int32** |  | [optional] 

## Methods

### NewUserSwitchRequest

`func NewUserSwitchRequest() *UserSwitchRequest`

NewUserSwitchRequest instantiates a new UserSwitchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSwitchRequestWithDefaults

`func NewUserSwitchRequestWithDefaults() *UserSwitchRequest`

NewUserSwitchRequestWithDefaults instantiates a new UserSwitchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UserSwitchRequest) GetAction() UserSwitchActionEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UserSwitchRequest) GetActionOk() (*UserSwitchActionEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UserSwitchRequest) SetAction(v UserSwitchActionEnum)`

SetAction sets Action field to given value.

### HasAction

`func (o *UserSwitchRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetUserPk

`func (o *UserSwitchRequest) GetUserPk() int32`

GetUserPk returns the UserPk field if non-nil, zero value otherwise.

### GetUserPkOk

`func (o *UserSwitchRequest) GetUserPkOk() (*int32, bool)`

GetUserPkOk returns a tuple with the UserPk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPk

`func (o *UserSwitchRequest) SetUserPk(v int32)`

SetUserPk sets UserPk field to given value.

### HasUserPk

`func (o *UserSwitchRequest) HasUserPk() bool`

HasUserPk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


