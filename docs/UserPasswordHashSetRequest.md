# UserPasswordHashSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** |  | 
**Override** | Pointer to **bool** | Import a valid password hash even when its parameters do not match authentik&#39;s current password hashing settings. | [optional] [default to false]

## Methods

### NewUserPasswordHashSetRequest

`func NewUserPasswordHashSetRequest(password string, ) *UserPasswordHashSetRequest`

NewUserPasswordHashSetRequest instantiates a new UserPasswordHashSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPasswordHashSetRequestWithDefaults

`func NewUserPasswordHashSetRequestWithDefaults() *UserPasswordHashSetRequest`

NewUserPasswordHashSetRequestWithDefaults instantiates a new UserPasswordHashSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UserPasswordHashSetRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserPasswordHashSetRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserPasswordHashSetRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetOverride

`func (o *UserPasswordHashSetRequest) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *UserPasswordHashSetRequest) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *UserPasswordHashSetRequest) SetOverride(v bool)`

SetOverride sets Override field to given value.

### HasOverride

`func (o *UserPasswordHashSetRequest) HasOverride() bool`

HasOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


