# OutpostHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSeen** | **time.Time** |  | [readonly] 
**Version** | **string** |  | [readonly] 
**VersionShould** | **string** |  | [readonly] 
**VersionOutdated** | **bool** |  | [readonly] 

## Methods

### NewOutpostHealth

`func NewOutpostHealth(lastSeen time.Time, version string, versionShould string, versionOutdated bool, ) *OutpostHealth`

NewOutpostHealth instantiates a new OutpostHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutpostHealthWithDefaults

`func NewOutpostHealthWithDefaults() *OutpostHealth`

NewOutpostHealthWithDefaults instantiates a new OutpostHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


