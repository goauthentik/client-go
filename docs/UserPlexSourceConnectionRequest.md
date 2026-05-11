# UserPlexSourceConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **int32** |  | 
**Source** | **string** |  | 
**Identifier** | **string** |  | 
**PlexToken** | **string** |  | 

## Methods

### NewUserPlexSourceConnectionRequest

`func NewUserPlexSourceConnectionRequest(user int32, source string, identifier string, plexToken string, ) *UserPlexSourceConnectionRequest`

NewUserPlexSourceConnectionRequest instantiates a new UserPlexSourceConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPlexSourceConnectionRequestWithDefaults

`func NewUserPlexSourceConnectionRequestWithDefaults() *UserPlexSourceConnectionRequest`

NewUserPlexSourceConnectionRequestWithDefaults instantiates a new UserPlexSourceConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserPlexSourceConnectionRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserPlexSourceConnectionRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserPlexSourceConnectionRequest) SetUser(v int32)`

SetUser sets User field to given value.


### GetSource

`func (o *UserPlexSourceConnectionRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UserPlexSourceConnectionRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UserPlexSourceConnectionRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetIdentifier

`func (o *UserPlexSourceConnectionRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *UserPlexSourceConnectionRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *UserPlexSourceConnectionRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetPlexToken

`func (o *UserPlexSourceConnectionRequest) GetPlexToken() string`

GetPlexToken returns the PlexToken field if non-nil, zero value otherwise.

### GetPlexTokenOk

`func (o *UserPlexSourceConnectionRequest) GetPlexTokenOk() (*string, bool)`

GetPlexTokenOk returns a tuple with the PlexToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlexToken

`func (o *UserPlexSourceConnectionRequest) SetPlexToken(v string)`

SetPlexToken sets PlexToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


