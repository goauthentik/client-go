# UserKerberosSourceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**User** | **int32** |  | 
**Source** | **string** |  | 
**SourceObj** | [**Source**](Source.md) |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Identifier** | **string** |  | 

## Methods

### NewUserKerberosSourceConnection

`func NewUserKerberosSourceConnection(pk int32, user int32, source string, sourceObj Source, created time.Time, identifier string, ) *UserKerberosSourceConnection`

NewUserKerberosSourceConnection instantiates a new UserKerberosSourceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserKerberosSourceConnectionWithDefaults

`func NewUserKerberosSourceConnectionWithDefaults() *UserKerberosSourceConnection`

NewUserKerberosSourceConnectionWithDefaults instantiates a new UserKerberosSourceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *UserKerberosSourceConnection) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *UserKerberosSourceConnection) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *UserKerberosSourceConnection) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetUser

`func (o *UserKerberosSourceConnection) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserKerberosSourceConnection) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserKerberosSourceConnection) SetUser(v int32)`

SetUser sets User field to given value.


### GetSource

`func (o *UserKerberosSourceConnection) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UserKerberosSourceConnection) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UserKerberosSourceConnection) SetSource(v string)`

SetSource sets Source field to given value.


### GetSourceObj

`func (o *UserKerberosSourceConnection) GetSourceObj() Source`

GetSourceObj returns the SourceObj field if non-nil, zero value otherwise.

### GetSourceObjOk

`func (o *UserKerberosSourceConnection) GetSourceObjOk() (*Source, bool)`

GetSourceObjOk returns a tuple with the SourceObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObj

`func (o *UserKerberosSourceConnection) SetSourceObj(v Source)`

SetSourceObj sets SourceObj field to given value.


### GetCreated

`func (o *UserKerberosSourceConnection) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserKerberosSourceConnection) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserKerberosSourceConnection) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetIdentifier

`func (o *UserKerberosSourceConnection) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *UserKerberosSourceConnection) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *UserKerberosSourceConnection) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


