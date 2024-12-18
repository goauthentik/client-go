# GroupPlexSourceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Group** | **string** |  | 
**Source** | **string** |  | 
**SourceObj** | [**Source**](Source.md) |  | [readonly] 
**Identifier** | **string** |  | 
**Created** | **time.Time** |  | [readonly] 

## Methods

### NewGroupPlexSourceConnection

`func NewGroupPlexSourceConnection(pk int32, group string, source string, sourceObj Source, identifier string, created time.Time, ) *GroupPlexSourceConnection`

NewGroupPlexSourceConnection instantiates a new GroupPlexSourceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPlexSourceConnectionWithDefaults

`func NewGroupPlexSourceConnectionWithDefaults() *GroupPlexSourceConnection`

NewGroupPlexSourceConnectionWithDefaults instantiates a new GroupPlexSourceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *GroupPlexSourceConnection) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *GroupPlexSourceConnection) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *GroupPlexSourceConnection) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetGroup

`func (o *GroupPlexSourceConnection) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GroupPlexSourceConnection) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GroupPlexSourceConnection) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetSource

`func (o *GroupPlexSourceConnection) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GroupPlexSourceConnection) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GroupPlexSourceConnection) SetSource(v string)`

SetSource sets Source field to given value.


### GetSourceObj

`func (o *GroupPlexSourceConnection) GetSourceObj() Source`

GetSourceObj returns the SourceObj field if non-nil, zero value otherwise.

### GetSourceObjOk

`func (o *GroupPlexSourceConnection) GetSourceObjOk() (*Source, bool)`

GetSourceObjOk returns a tuple with the SourceObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObj

`func (o *GroupPlexSourceConnection) SetSourceObj(v Source)`

SetSourceObj sets SourceObj field to given value.


### GetIdentifier

`func (o *GroupPlexSourceConnection) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *GroupPlexSourceConnection) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *GroupPlexSourceConnection) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetCreated

`func (o *GroupPlexSourceConnection) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GroupPlexSourceConnection) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GroupPlexSourceConnection) SetCreated(v time.Time)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


