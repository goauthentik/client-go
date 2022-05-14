# UserOAuthSourceConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **int32** |  | 
**Identifier** | **string** |  | 
**AccessToken** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserOAuthSourceConnectionRequest

`func NewUserOAuthSourceConnectionRequest(user int32, identifier string, ) *UserOAuthSourceConnectionRequest`

NewUserOAuthSourceConnectionRequest instantiates a new UserOAuthSourceConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOAuthSourceConnectionRequestWithDefaults

`func NewUserOAuthSourceConnectionRequestWithDefaults() *UserOAuthSourceConnectionRequest`

NewUserOAuthSourceConnectionRequestWithDefaults instantiates a new UserOAuthSourceConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserOAuthSourceConnectionRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserOAuthSourceConnectionRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserOAuthSourceConnectionRequest) SetUser(v int32)`

SetUser sets User field to given value.


### GetIdentifier

`func (o *UserOAuthSourceConnectionRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *UserOAuthSourceConnectionRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *UserOAuthSourceConnectionRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetAccessToken

`func (o *UserOAuthSourceConnectionRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *UserOAuthSourceConnectionRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *UserOAuthSourceConnectionRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *UserOAuthSourceConnectionRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *UserOAuthSourceConnectionRequest) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *UserOAuthSourceConnectionRequest) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


