# UpdateVersionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**Changelog** | **string** |  | 
**ChangelogUrl** | **string** |  | 
**Reason** | [**UpdateVersionReason**](UpdateVersionReason.md) |  | 
**Valid** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewUpdateVersionInfo

`func NewUpdateVersionInfo(version string, changelog string, changelogUrl string, reason UpdateVersionReason, ) *UpdateVersionInfo`

NewUpdateVersionInfo instantiates a new UpdateVersionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVersionInfoWithDefaults

`func NewUpdateVersionInfoWithDefaults() *UpdateVersionInfo`

NewUpdateVersionInfoWithDefaults instantiates a new UpdateVersionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *UpdateVersionInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateVersionInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateVersionInfo) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetChangelog

`func (o *UpdateVersionInfo) GetChangelog() string`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *UpdateVersionInfo) GetChangelogOk() (*string, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *UpdateVersionInfo) SetChangelog(v string)`

SetChangelog sets Changelog field to given value.


### GetChangelogUrl

`func (o *UpdateVersionInfo) GetChangelogUrl() string`

GetChangelogUrl returns the ChangelogUrl field if non-nil, zero value otherwise.

### GetChangelogUrlOk

`func (o *UpdateVersionInfo) GetChangelogUrlOk() (*string, bool)`

GetChangelogUrlOk returns a tuple with the ChangelogUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogUrl

`func (o *UpdateVersionInfo) SetChangelogUrl(v string)`

SetChangelogUrl sets ChangelogUrl field to given value.


### GetReason

`func (o *UpdateVersionInfo) GetReason() UpdateVersionReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UpdateVersionInfo) GetReasonOk() (*UpdateVersionReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UpdateVersionInfo) SetReason(v UpdateVersionReason)`

SetReason sets Reason field to given value.


### GetValid

`func (o *UpdateVersionInfo) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *UpdateVersionInfo) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *UpdateVersionInfo) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *UpdateVersionInfo) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


