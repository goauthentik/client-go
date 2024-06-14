# OutpostHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** |  | [readonly] 
**LastSeen** | **time.Time** |  | [readonly] 
**Version** | **string** |  | [readonly] 
**GolangVersion** | **string** |  | [readonly] 
**OpensslEnabled** | **bool** |  | [readonly] 
**OpensslVersion** | **string** |  | [readonly] 
**FipsEnabled** | **NullableBool** | Get FIPS enabled | [readonly] 
**VersionShould** | **string** |  | [readonly] 
**VersionOutdated** | **bool** |  | [readonly] 
**BuildHash** | **string** |  | [readonly] 
**BuildHashShould** | **string** |  | [readonly] 
**Hostname** | **string** |  | [readonly] 

## Methods

### NewOutpostHealth

`func NewOutpostHealth(uid string, lastSeen time.Time, version string, golangVersion string, opensslEnabled bool, opensslVersion string, fipsEnabled NullableBool, versionShould string, versionOutdated bool, buildHash string, buildHashShould string, hostname string, ) *OutpostHealth`

NewOutpostHealth instantiates a new OutpostHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutpostHealthWithDefaults

`func NewOutpostHealthWithDefaults() *OutpostHealth`

NewOutpostHealthWithDefaults instantiates a new OutpostHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *OutpostHealth) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *OutpostHealth) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *OutpostHealth) SetUid(v string)`

SetUid sets Uid field to given value.


### GetLastSeen

`func (o *OutpostHealth) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *OutpostHealth) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *OutpostHealth) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.


### GetVersion

`func (o *OutpostHealth) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OutpostHealth) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OutpostHealth) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetGolangVersion

`func (o *OutpostHealth) GetGolangVersion() string`

GetGolangVersion returns the GolangVersion field if non-nil, zero value otherwise.

### GetGolangVersionOk

`func (o *OutpostHealth) GetGolangVersionOk() (*string, bool)`

GetGolangVersionOk returns a tuple with the GolangVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGolangVersion

`func (o *OutpostHealth) SetGolangVersion(v string)`

SetGolangVersion sets GolangVersion field to given value.


### GetOpensslEnabled

`func (o *OutpostHealth) GetOpensslEnabled() bool`

GetOpensslEnabled returns the OpensslEnabled field if non-nil, zero value otherwise.

### GetOpensslEnabledOk

`func (o *OutpostHealth) GetOpensslEnabledOk() (*bool, bool)`

GetOpensslEnabledOk returns a tuple with the OpensslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpensslEnabled

`func (o *OutpostHealth) SetOpensslEnabled(v bool)`

SetOpensslEnabled sets OpensslEnabled field to given value.


### GetOpensslVersion

`func (o *OutpostHealth) GetOpensslVersion() string`

GetOpensslVersion returns the OpensslVersion field if non-nil, zero value otherwise.

### GetOpensslVersionOk

`func (o *OutpostHealth) GetOpensslVersionOk() (*string, bool)`

GetOpensslVersionOk returns a tuple with the OpensslVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpensslVersion

`func (o *OutpostHealth) SetOpensslVersion(v string)`

SetOpensslVersion sets OpensslVersion field to given value.


### GetFipsEnabled

`func (o *OutpostHealth) GetFipsEnabled() bool`

GetFipsEnabled returns the FipsEnabled field if non-nil, zero value otherwise.

### GetFipsEnabledOk

`func (o *OutpostHealth) GetFipsEnabledOk() (*bool, bool)`

GetFipsEnabledOk returns a tuple with the FipsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsEnabled

`func (o *OutpostHealth) SetFipsEnabled(v bool)`

SetFipsEnabled sets FipsEnabled field to given value.


### SetFipsEnabledNil

`func (o *OutpostHealth) SetFipsEnabledNil(b bool)`

 SetFipsEnabledNil sets the value for FipsEnabled to be an explicit nil

### UnsetFipsEnabled
`func (o *OutpostHealth) UnsetFipsEnabled()`

UnsetFipsEnabled ensures that no value is present for FipsEnabled, not even an explicit nil
### GetVersionShould

`func (o *OutpostHealth) GetVersionShould() string`

GetVersionShould returns the VersionShould field if non-nil, zero value otherwise.

### GetVersionShouldOk

`func (o *OutpostHealth) GetVersionShouldOk() (*string, bool)`

GetVersionShouldOk returns a tuple with the VersionShould field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionShould

`func (o *OutpostHealth) SetVersionShould(v string)`

SetVersionShould sets VersionShould field to given value.


### GetVersionOutdated

`func (o *OutpostHealth) GetVersionOutdated() bool`

GetVersionOutdated returns the VersionOutdated field if non-nil, zero value otherwise.

### GetVersionOutdatedOk

`func (o *OutpostHealth) GetVersionOutdatedOk() (*bool, bool)`

GetVersionOutdatedOk returns a tuple with the VersionOutdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionOutdated

`func (o *OutpostHealth) SetVersionOutdated(v bool)`

SetVersionOutdated sets VersionOutdated field to given value.


### GetBuildHash

`func (o *OutpostHealth) GetBuildHash() string`

GetBuildHash returns the BuildHash field if non-nil, zero value otherwise.

### GetBuildHashOk

`func (o *OutpostHealth) GetBuildHashOk() (*string, bool)`

GetBuildHashOk returns a tuple with the BuildHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildHash

`func (o *OutpostHealth) SetBuildHash(v string)`

SetBuildHash sets BuildHash field to given value.


### GetBuildHashShould

`func (o *OutpostHealth) GetBuildHashShould() string`

GetBuildHashShould returns the BuildHashShould field if non-nil, zero value otherwise.

### GetBuildHashShouldOk

`func (o *OutpostHealth) GetBuildHashShouldOk() (*string, bool)`

GetBuildHashShouldOk returns a tuple with the BuildHashShould field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildHashShould

`func (o *OutpostHealth) SetBuildHashShould(v string)`

SetBuildHashShould sets BuildHashShould field to given value.


### GetHostname

`func (o *OutpostHealth) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *OutpostHealth) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *OutpostHealth) SetHostname(v string)`

SetHostname sets Hostname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


