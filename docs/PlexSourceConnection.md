# PlexSourceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**User** | **int32** |  | [readonly] 
**Source** | [**Source**](Source.md) |  | [readonly] 
**Identifier** | **string** |  | 
**PlexToken** | **string** |  | 

## Methods

### NewPlexSourceConnection

`func NewPlexSourceConnection(pk int32, user int32, source Source, identifier string, plexToken string, ) *PlexSourceConnection`

NewPlexSourceConnection instantiates a new PlexSourceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlexSourceConnectionWithDefaults

`func NewPlexSourceConnectionWithDefaults() *PlexSourceConnection`

NewPlexSourceConnectionWithDefaults instantiates a new PlexSourceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *PlexSourceConnection) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *PlexSourceConnection) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *PlexSourceConnection) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetUser

`func (o *PlexSourceConnection) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PlexSourceConnection) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PlexSourceConnection) SetUser(v int32)`

SetUser sets User field to given value.


### GetSource

`func (o *PlexSourceConnection) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PlexSourceConnection) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PlexSourceConnection) SetSource(v Source)`

SetSource sets Source field to given value.


### GetIdentifier

`func (o *PlexSourceConnection) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PlexSourceConnection) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PlexSourceConnection) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetPlexToken

`func (o *PlexSourceConnection) GetPlexToken() string`

GetPlexToken returns the PlexToken field if non-nil, zero value otherwise.

### GetPlexTokenOk

`func (o *PlexSourceConnection) GetPlexTokenOk() (*string, bool)`

GetPlexTokenOk returns a tuple with the PlexToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlexToken

`func (o *PlexSourceConnection) SetPlexToken(v string)`

SetPlexToken sets PlexToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


