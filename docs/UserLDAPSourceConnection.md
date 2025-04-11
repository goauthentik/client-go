# UserLDAPSourceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**User** | **int32** |  | 
**Source** | **string** |  | 
**SourceObj** | [**Source**](Source.md) |  | [readonly] 
**Identifier** | **string** |  | 
**Created** | **time.Time** |  | [readonly] 
**LastUpdated** | **time.Time** |  | [readonly] 

## Methods

### NewUserLDAPSourceConnection

`func NewUserLDAPSourceConnection(pk int32, user int32, source string, sourceObj Source, identifier string, created time.Time, lastUpdated time.Time, ) *UserLDAPSourceConnection`

NewUserLDAPSourceConnection instantiates a new UserLDAPSourceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLDAPSourceConnectionWithDefaults

`func NewUserLDAPSourceConnectionWithDefaults() *UserLDAPSourceConnection`

NewUserLDAPSourceConnectionWithDefaults instantiates a new UserLDAPSourceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *UserLDAPSourceConnection) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *UserLDAPSourceConnection) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *UserLDAPSourceConnection) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetUser

`func (o *UserLDAPSourceConnection) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserLDAPSourceConnection) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserLDAPSourceConnection) SetUser(v int32)`

SetUser sets User field to given value.


### GetSource

`func (o *UserLDAPSourceConnection) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UserLDAPSourceConnection) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UserLDAPSourceConnection) SetSource(v string)`

SetSource sets Source field to given value.


### GetSourceObj

`func (o *UserLDAPSourceConnection) GetSourceObj() Source`

GetSourceObj returns the SourceObj field if non-nil, zero value otherwise.

### GetSourceObjOk

`func (o *UserLDAPSourceConnection) GetSourceObjOk() (*Source, bool)`

GetSourceObjOk returns a tuple with the SourceObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObj

`func (o *UserLDAPSourceConnection) SetSourceObj(v Source)`

SetSourceObj sets SourceObj field to given value.


### GetIdentifier

`func (o *UserLDAPSourceConnection) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *UserLDAPSourceConnection) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *UserLDAPSourceConnection) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetCreated

`func (o *UserLDAPSourceConnection) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserLDAPSourceConnection) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserLDAPSourceConnection) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *UserLDAPSourceConnection) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UserLDAPSourceConnection) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UserLDAPSourceConnection) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


