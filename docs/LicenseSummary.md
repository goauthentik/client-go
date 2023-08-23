# LicenseSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalUsers** | **int32** |  | 
**ExternalUsers** | **int32** |  | 
**Valid** | **bool** |  | 
**ShowAdminWarning** | **bool** |  | 
**ShowUserWarning** | **bool** |  | 
**ReadOnly** | **bool** |  | 
**LatestValid** | **time.Time** |  | 
**HasLicense** | **bool** |  | 

## Methods

### NewLicenseSummary

`func NewLicenseSummary(internalUsers int32, externalUsers int32, valid bool, showAdminWarning bool, showUserWarning bool, readOnly bool, latestValid time.Time, hasLicense bool, ) *LicenseSummary`

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


### GetValid

`func (o *LicenseSummary) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *LicenseSummary) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *LicenseSummary) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetShowAdminWarning

`func (o *LicenseSummary) GetShowAdminWarning() bool`

GetShowAdminWarning returns the ShowAdminWarning field if non-nil, zero value otherwise.

### GetShowAdminWarningOk

`func (o *LicenseSummary) GetShowAdminWarningOk() (*bool, bool)`

GetShowAdminWarningOk returns a tuple with the ShowAdminWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAdminWarning

`func (o *LicenseSummary) SetShowAdminWarning(v bool)`

SetShowAdminWarning sets ShowAdminWarning field to given value.


### GetShowUserWarning

`func (o *LicenseSummary) GetShowUserWarning() bool`

GetShowUserWarning returns the ShowUserWarning field if non-nil, zero value otherwise.

### GetShowUserWarningOk

`func (o *LicenseSummary) GetShowUserWarningOk() (*bool, bool)`

GetShowUserWarningOk returns a tuple with the ShowUserWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowUserWarning

`func (o *LicenseSummary) SetShowUserWarning(v bool)`

SetShowUserWarning sets ShowUserWarning field to given value.


### GetReadOnly

`func (o *LicenseSummary) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *LicenseSummary) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *LicenseSummary) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.


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


### GetHasLicense

`func (o *LicenseSummary) GetHasLicense() bool`

GetHasLicense returns the HasLicense field if non-nil, zero value otherwise.

### GetHasLicenseOk

`func (o *LicenseSummary) GetHasLicenseOk() (*bool, bool)`

GetHasLicenseOk returns a tuple with the HasLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLicense

`func (o *LicenseSummary) SetHasLicense(v bool)`

SetHasLicense sets HasLicense field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


