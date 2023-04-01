# UserOAuthSourceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**User** | **int32** |  | 
**Source** | [**Source**](Source.md) |  | [readonly] 
**Identifier** | **string** |  | 

## Methods

### NewUserOAuthSourceConnection

`func NewUserOAuthSourceConnection(pk int32, user int32, source Source, identifier string, ) *UserOAuthSourceConnection`

NewUserOAuthSourceConnection instantiates a new UserOAuthSourceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOAuthSourceConnectionWithDefaults

`func NewUserOAuthSourceConnectionWithDefaults() *UserOAuthSourceConnection`

NewUserOAuthSourceConnectionWithDefaults instantiates a new UserOAuthSourceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *UserOAuthSourceConnection) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *UserOAuthSourceConnection) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *UserOAuthSourceConnection) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetUser

`func (o *UserOAuthSourceConnection) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserOAuthSourceConnection) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserOAuthSourceConnection) SetUser(v int32)`

SetUser sets User field to given value.


### GetSource

`func (o *UserOAuthSourceConnection) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UserOAuthSourceConnection) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UserOAuthSourceConnection) SetSource(v Source)`

SetSource sets Source field to given value.


### GetIdentifier

`func (o *UserOAuthSourceConnection) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *UserOAuthSourceConnection) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *UserOAuthSourceConnection) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


