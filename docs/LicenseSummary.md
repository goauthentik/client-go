# LicenseSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalUsers** | **int32** |  | 
**ExternalUsers** | **int32** |  | 
**Status** | [**LicenseSummaryStatusEnum**](LicenseSummaryStatusEnum.md) |  | 
**LatestValid** | **time.Time** |  | 

## Methods

### NewLicenseSummary

`func NewLicenseSummary(internalUsers int32, externalUsers int32, status LicenseSummaryStatusEnum, latestValid time.Time, ) *LicenseSummary`

NewLicenseSummary instantiates a new LicenseSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseSummaryWithDefaults

`func NewLicenseSummaryWithDefaults() *LicenseSummary`

NewLicenseSummaryWithDefaults instantiates a new LicenseSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternalUsers

`func (o *LicenseSummary) GetInternalUsers() int32`

GetInternalUsers returns the InternalUsers field if non-nil, zero value otherwise.

### GetInternalUsersOk

`func (o *LicenseSummary) GetInternalUsersOk() (*int32, bool)`

GetInternalUsersOk returns a tuple with the InternalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalUsers

`func (o *LicenseSummary) SetInternalUsers(v int32)`

SetInternalUsers sets InternalUsers field to given value.


### GetExternalUsers

`func (o *LicenseSummary) GetExternalUsers() int32`

GetExternalUsers returns the ExternalUsers field if non-nil, zero value otherwise.

### GetExternalUsersOk

`func (o *LicenseSummary) GetExternalUsersOk() (*int32, bool)`

GetExternalUsersOk returns a tuple with the ExternalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUsers

`func (o *LicenseSummary) SetExternalUsers(v int32)`

SetExternalUsers sets ExternalUsers field to given value.


### GetStatus

`func (o *LicenseSummary) GetStatus() LicenseSummaryStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LicenseSummary) GetStatusOk() (*LicenseSummaryStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LicenseSummary) SetStatus(v LicenseSummaryStatusEnum)`

SetStatus sets Status field to given value.


### GetLatestValid

`func (o *LicenseSummary) GetLatestValid() time.Time`

GetLatestValid returns the LatestValid field if non-nil, zero value otherwise.

### GetLatestValidOk

`func (o *LicenseSummary) GetLatestValidOk() (*time.Time, bool)`

GetLatestValidOk returns a tuple with the LatestValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestValid

`func (o *LicenseSummary) SetLatestValid(v time.Time)`

SetLatestValid sets LatestValid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


