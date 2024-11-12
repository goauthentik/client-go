# GroupKerberosSourceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Group** | **string** |  | [readonly] 
**Source** | [**Source**](Source.md) |  | [readonly] 
**Identifier** | **string** |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 

## Methods

### NewGroupKerberosSourceConnection

`func NewGroupKerberosSourceConnection(pk int32, group string, source Source, identifier string, created time.Time, ) *GroupKerberosSourceConnection`

NewGroupKerberosSourceConnection instantiates a new GroupKerberosSourceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupKerberosSourceConnectionWithDefaults

`func NewGroupKerberosSourceConnectionWithDefaults() *GroupKerberosSourceConnection`

NewGroupKerberosSourceConnectionWithDefaults instantiates a new GroupKerberosSourceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *GroupKerberosSourceConnection) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *GroupKerberosSourceConnection) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *GroupKerberosSourceConnection) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetGroup

`func (o *GroupKerberosSourceConnection) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GroupKerberosSourceConnection) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GroupKerberosSourceConnection) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetSource

`func (o *GroupKerberosSourceConnection) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GroupKerberosSourceConnection) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GroupKerberosSourceConnection) SetSource(v Source)`

SetSource sets Source field to given value.


### GetIdentifier

`func (o *GroupKerberosSourceConnection) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *GroupKerberosSourceConnection) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *GroupKerberosSourceConnection) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetCreated

`func (o *GroupKerberosSourceConnection) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GroupKerberosSourceConnection) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GroupKerberosSourceConnection) SetCreated(v time.Time)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

