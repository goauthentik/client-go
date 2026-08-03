# VersionHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Timestamp** | **time.Time** |  | 
**Version** | **string** |  | 
**Build** | **string** |  | 

## Methods

### NewVersionHistory

`func NewVersionHistory(id int32, timestamp time.Time, version string, build string, ) *VersionHistory`

NewVersionHistory instantiates a new VersionHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionHistoryWithDefaults

`func NewVersionHistoryWithDefaults() *VersionHistory`

NewVersionHistoryWithDefaults instantiates a new VersionHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VersionHistory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VersionHistory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VersionHistory) SetId(v int32)`

SetId sets Id field to given value.


### GetTimestamp

`func (o *VersionHistory) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VersionHistory) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VersionHistory) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetVersion

`func (o *VersionHistory) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VersionHistory) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VersionHistory) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetBuild

`func (o *VersionHistory) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *VersionHistory) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *VersionHistory) SetBuild(v string)`

SetBuild sets Build field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


