# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionCurrent** | **string** | Get current version | [readonly] 
**VersionLatest** | **string** | Get latest version from cache | [readonly] 
**VersionLatestValid** | **bool** | Check if latest version is valid | [readonly] 
**BuildHash** | **string** | Get build hash, if version is not latest or released | [readonly] 
**Outdated** | **bool** | Check if we&#39;re running the latest version | [readonly] 
**OutpostOutdated** | **bool** | Check if any outpost is outdated/has a version mismatch | [readonly] 

## Methods

### NewVersion

`func NewVersion(versionCurrent string, versionLatest string, versionLatestValid bool, buildHash string, outdated bool, outpostOutdated bool, ) *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionCurrent

`func (o *Version) GetVersionCurrent() string`

GetVersionCurrent returns the VersionCurrent field if non-nil, zero value otherwise.

### GetVersionCurrentOk

`func (o *Version) GetVersionCurrentOk() (*string, bool)`

GetVersionCurrentOk returns a tuple with the VersionCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCurrent

`func (o *Version) SetVersionCurrent(v string)`

SetVersionCurrent sets VersionCurrent field to given value.


### GetVersionLatest

`func (o *Version) GetVersionLatest() string`

GetVersionLatest returns the VersionLatest field if non-nil, zero value otherwise.

### GetVersionLatestOk

`func (o *Version) GetVersionLatestOk() (*string, bool)`

GetVersionLatestOk returns a tuple with the VersionLatest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionLatest

`func (o *Version) SetVersionLatest(v string)`

SetVersionLatest sets VersionLatest field to given value.


### GetVersionLatestValid

`func (o *Version) GetVersionLatestValid() bool`

GetVersionLatestValid returns the VersionLatestValid field if non-nil, zero value otherwise.

### GetVersionLatestValidOk

`func (o *Version) GetVersionLatestValidOk() (*bool, bool)`

GetVersionLatestValidOk returns a tuple with the VersionLatestValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionLatestValid

`func (o *Version) SetVersionLatestValid(v bool)`

SetVersionLatestValid sets VersionLatestValid field to given value.


### GetBuildHash

`func (o *Version) GetBuildHash() string`

GetBuildHash returns the BuildHash field if non-nil, zero value otherwise.

### GetBuildHashOk

`func (o *Version) GetBuildHashOk() (*string, bool)`

GetBuildHashOk returns a tuple with the BuildHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildHash

`func (o *Version) SetBuildHash(v string)`

SetBuildHash sets BuildHash field to given value.


### GetOutdated

`func (o *Version) GetOutdated() bool`

GetOutdated returns the Outdated field if non-nil, zero value otherwise.

### GetOutdatedOk

`func (o *Version) GetOutdatedOk() (*bool, bool)`

GetOutdatedOk returns a tuple with the Outdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutdated

`func (o *Version) SetOutdated(v bool)`

SetOutdated sets Outdated field to given value.


### GetOutpostOutdated

`func (o *Version) GetOutpostOutdated() bool`

GetOutpostOutdated returns the OutpostOutdated field if non-nil, zero value otherwise.

### GetOutpostOutdatedOk

`func (o *Version) GetOutpostOutdatedOk() (*bool, bool)`

GetOutpostOutdatedOk returns a tuple with the OutpostOutdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutpostOutdated

`func (o *Version) SetOutpostOutdated(v bool)`

SetOutpostOutdated sets OutpostOutdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


