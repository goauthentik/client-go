# UserSAMLSourceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**User** | **int32** |  | 
**Source** | [**PlexSourceConnectionSource**](PlexSourceConnectionSource.md) |  | 
**Identifier** | **string** |  | 

## Methods

### NewUserSAMLSourceConnection

`func NewUserSAMLSourceConnection(pk int32, user int32, source PlexSourceConnectionSource, identifier string, ) *UserSAMLSourceConnection`

NewUserSAMLSourceConnection instantiates a new UserSAMLSourceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSAMLSourceConnectionWithDefaults

`func NewUserSAMLSourceConnectionWithDefaults() *UserSAMLSourceConnection`

NewUserSAMLSourceConnectionWithDefaults instantiates a new UserSAMLSourceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *UserSAMLSourceConnection) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *UserSAMLSourceConnection) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *UserSAMLSourceConnection) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetUser

`func (o *UserSAMLSourceConnection) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserSAMLSourceConnection) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserSAMLSourceConnection) SetUser(v int32)`

SetUser sets User field to given value.


### GetSource

`func (o *UserSAMLSourceConnection) GetSource() PlexSourceConnectionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UserSAMLSourceConnection) GetSourceOk() (*PlexSourceConnectionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UserSAMLSourceConnection) SetSource(v PlexSourceConnectionSource)`

SetSource sets Source field to given value.


### GetIdentifier

`func (o *UserSAMLSourceConnection) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *UserSAMLSourceConnection) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *UserSAMLSourceConnection) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


