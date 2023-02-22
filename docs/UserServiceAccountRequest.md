# UserServiceAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CreateGroup** | Pointer to **bool** |  | [optional] [default to false]
**Expiring** | Pointer to **bool** |  | [optional] [default to true]
**Expires** | Pointer to **time.Time** | If not provided, valid for 360 days | [optional] 

## Methods

### NewUserServiceAccountRequest

`func NewUserServiceAccountRequest(name string, ) *UserServiceAccountRequest`

NewUserServiceAccountRequest instantiates a new UserServiceAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserServiceAccountRequestWithDefaults

`func NewUserServiceAccountRequestWithDefaults() *UserServiceAccountRequest`

NewUserServiceAccountRequestWithDefaults instantiates a new UserServiceAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserServiceAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserServiceAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserServiceAccountRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCreateGroup

`func (o *UserServiceAccountRequest) GetCreateGroup() bool`

GetCreateGroup returns the CreateGroup field if non-nil, zero value otherwise.

### GetCreateGroupOk

`func (o *UserServiceAccountRequest) GetCreateGroupOk() (*bool, bool)`

GetCreateGroupOk returns a tuple with the CreateGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateGroup

`func (o *UserServiceAccountRequest) SetCreateGroup(v bool)`

SetCreateGroup sets CreateGroup field to given value.

### HasCreateGroup

`func (o *UserServiceAccountRequest) HasCreateGroup() bool`

HasCreateGroup returns a boolean if a field has been set.

### GetExpiring

`func (o *UserServiceAccountRequest) GetExpiring() bool`

GetExpiring returns the Expiring field if non-nil, zero value otherwise.

### GetExpiringOk

`func (o *UserServiceAccountRequest) GetExpiringOk() (*bool, bool)`

GetExpiringOk returns a tuple with the Expiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiring

`func (o *UserServiceAccountRequest) SetExpiring(v bool)`

SetExpiring sets Expiring field to given value.

### HasExpiring

`func (o *UserServiceAccountRequest) HasExpiring() bool`

HasExpiring returns a boolean if a field has been set.

### GetExpires

`func (o *UserServiceAccountRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *UserServiceAccountRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *UserServiceAccountRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *UserServiceAccountRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


