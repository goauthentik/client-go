# LicenseUserCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveInternalUsers** | **int32** |  | 
**ActiveExternalUsers** | **int32** |  | 
**Ranges** | [**[]LicenseUserCountRange**](LicenseUserCountRange.md) |  | 

## Methods

### NewLicenseUserCounts

`func NewLicenseUserCounts(activeInternalUsers int32, activeExternalUsers int32, ranges []LicenseUserCountRange, ) *LicenseUserCounts`

NewLicenseUserCounts instantiates a new LicenseUserCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseUserCountsWithDefaults

`func NewLicenseUserCountsWithDefaults() *LicenseUserCounts`

NewLicenseUserCountsWithDefaults instantiates a new LicenseUserCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveInternalUsers

`func (o *LicenseUserCounts) GetActiveInternalUsers() int32`

GetActiveInternalUsers returns the ActiveInternalUsers field if non-nil, zero value otherwise.

### GetActiveInternalUsersOk

`func (o *LicenseUserCounts) GetActiveInternalUsersOk() (*int32, bool)`

GetActiveInternalUsersOk returns a tuple with the ActiveInternalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveInternalUsers

`func (o *LicenseUserCounts) SetActiveInternalUsers(v int32)`

SetActiveInternalUsers sets ActiveInternalUsers field to given value.


### GetActiveExternalUsers

`func (o *LicenseUserCounts) GetActiveExternalUsers() int32`

GetActiveExternalUsers returns the ActiveExternalUsers field if non-nil, zero value otherwise.

### GetActiveExternalUsersOk

`func (o *LicenseUserCounts) GetActiveExternalUsersOk() (*int32, bool)`

GetActiveExternalUsersOk returns a tuple with the ActiveExternalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveExternalUsers

`func (o *LicenseUserCounts) SetActiveExternalUsers(v int32)`

SetActiveExternalUsers sets ActiveExternalUsers field to given value.


### GetRanges

`func (o *LicenseUserCounts) GetRanges() []LicenseUserCountRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *LicenseUserCounts) GetRangesOk() (*[]LicenseUserCountRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *LicenseUserCounts) SetRanges(v []LicenseUserCountRange)`

SetRanges sets Ranges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


