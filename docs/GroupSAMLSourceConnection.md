# GroupSAMLSourceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Group** | **string** |  | 
**Source** | **string** |  | 
**SourceObj** | [**Source**](Source.md) |  | [readonly] 
**Identifier** | **string** |  | 
**Created** | **time.Time** |  | [readonly] 
**LastUpdated** | **time.Time** |  | [readonly] 

## Methods

### NewGroupSAMLSourceConnection

`func NewGroupSAMLSourceConnection(pk int32, group string, source string, sourceObj Source, identifier string, created time.Time, lastUpdated time.Time, ) *GroupSAMLSourceConnection`

NewGroupSAMLSourceConnection instantiates a new GroupSAMLSourceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSAMLSourceConnectionWithDefaults

`func NewGroupSAMLSourceConnectionWithDefaults() *GroupSAMLSourceConnection`

NewGroupSAMLSourceConnectionWithDefaults instantiates a new GroupSAMLSourceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *GroupSAMLSourceConnection) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *GroupSAMLSourceConnection) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *GroupSAMLSourceConnection) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetGroup

`func (o *GroupSAMLSourceConnection) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GroupSAMLSourceConnection) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GroupSAMLSourceConnection) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetSource

`func (o *GroupSAMLSourceConnection) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GroupSAMLSourceConnection) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GroupSAMLSourceConnection) SetSource(v string)`

SetSource sets Source field to given value.


### GetSourceObj

`func (o *GroupSAMLSourceConnection) GetSourceObj() Source`

GetSourceObj returns the SourceObj field if non-nil, zero value otherwise.

### GetSourceObjOk

`func (o *GroupSAMLSourceConnection) GetSourceObjOk() (*Source, bool)`

GetSourceObjOk returns a tuple with the SourceObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObj

`func (o *GroupSAMLSourceConnection) SetSourceObj(v Source)`

SetSourceObj sets SourceObj field to given value.


### GetIdentifier

`func (o *GroupSAMLSourceConnection) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *GroupSAMLSourceConnection) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *GroupSAMLSourceConnection) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetCreated

`func (o *GroupSAMLSourceConnection) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GroupSAMLSourceConnection) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GroupSAMLSourceConnection) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *GroupSAMLSourceConnection) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GroupSAMLSourceConnection) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GroupSAMLSourceConnection) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


