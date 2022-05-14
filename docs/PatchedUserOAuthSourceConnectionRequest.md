# PatchedUserOAuthSourceConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **int32** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**AccessToken** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPatchedUserOAuthSourceConnectionRequest

`func NewPatchedUserOAuthSourceConnectionRequest() *PatchedUserOAuthSourceConnectionRequest`

NewPatchedUserOAuthSourceConnectionRequest instantiates a new PatchedUserOAuthSourceConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedUserOAuthSourceConnectionRequestWithDefaults

`func NewPatchedUserOAuthSourceConnectionRequestWithDefaults() *PatchedUserOAuthSourceConnectionRequest`

NewPatchedUserOAuthSourceConnectionRequestWithDefaults instantiates a new PatchedUserOAuthSourceConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *PatchedUserOAuthSourceConnectionRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedUserOAuthSourceConnectionRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedUserOAuthSourceConnectionRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedUserOAuthSourceConnectionRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetIdentifier

`func (o *PatchedUserOAuthSourceConnectionRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PatchedUserOAuthSourceConnectionRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PatchedUserOAuthSourceConnectionRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PatchedUserOAuthSourceConnectionRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetAccessToken

`func (o *PatchedUserOAuthSourceConnectionRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *PatchedUserOAuthSourceConnectionRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *PatchedUserOAuthSourceConnectionRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *PatchedUserOAuthSourceConnectionRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *PatchedUserOAuthSourceConnectionRequest) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *PatchedUserOAuthSourceConnectionRequest) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


