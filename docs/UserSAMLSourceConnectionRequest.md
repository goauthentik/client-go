# UserSAMLSourceConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **int32** |  | 
**Source** | **string** |  | 
**Identifier** | **string** |  | 

## Methods

### NewUserSAMLSourceConnectionRequest

`func NewUserSAMLSourceConnectionRequest(user int32, source string, identifier string, ) *UserSAMLSourceConnectionRequest`

NewUserSAMLSourceConnectionRequest instantiates a new UserSAMLSourceConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSAMLSourceConnectionRequestWithDefaults

`func NewUserSAMLSourceConnectionRequestWithDefaults() *UserSAMLSourceConnectionRequest`

NewUserSAMLSourceConnectionRequestWithDefaults instantiates a new UserSAMLSourceConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserSAMLSourceConnectionRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserSAMLSourceConnectionRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserSAMLSourceConnectionRequest) SetUser(v int32)`

SetUser sets User field to given value.


### GetSource

`func (o *UserSAMLSourceConnectionRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UserSAMLSourceConnectionRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UserSAMLSourceConnectionRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetIdentifier

`func (o *UserSAMLSourceConnectionRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *UserSAMLSourceConnectionRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *UserSAMLSourceConnectionRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


