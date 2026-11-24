# LicenseUserCountRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **time.Time** |  | 
**End** | **time.Time** |  | 
**Interval** | **NullableString** |  | 
**InternalUsersAdded** | **int32** |  | 
**ExternalUsersAdded** | **int32** |  | 

## Methods

### NewLicenseUserCountRange

`func NewLicenseUserCountRange(start time.Time, end time.Time, interval NullableString, internalUsersAdded int32, externalUsersAdded int32, ) *LicenseUserCountRange`

NewLicenseUserCountRange instantiates a new LicenseUserCountRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseUserCountRangeWithDefaults

`func NewLicenseUserCountRangeWithDefaults() *LicenseUserCountRange`

NewLicenseUserCountRangeWithDefaults instantiates a new LicenseUserCountRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *LicenseUserCountRange) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *LicenseUserCountRange) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *LicenseUserCountRange) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *LicenseUserCountRange) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *LicenseUserCountRange) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *LicenseUserCountRange) SetEnd(v time.Time)`

SetEnd sets End field to given value.


### GetInterval

`func (o *LicenseUserCountRange) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *LicenseUserCountRange) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *LicenseUserCountRange) SetInterval(v string)`

SetInterval sets Interval field to given value.


### SetIntervalNil

`func (o *LicenseUserCountRange) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *LicenseUserCountRange) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetInternalUsersAdded

`func (o *LicenseUserCountRange) GetInternalUsersAdded() int32`

GetInternalUsersAdded returns the InternalUsersAdded field if non-nil, zero value otherwise.

### GetInternalUsersAddedOk

`func (o *LicenseUserCountRange) GetInternalUsersAddedOk() (*int32, bool)`

GetInternalUsersAddedOk returns a tuple with the InternalUsersAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalUsersAdded

`func (o *LicenseUserCountRange) SetInternalUsersAdded(v int32)`

SetInternalUsersAdded sets InternalUsersAdded field to given value.


### GetExternalUsersAdded

`func (o *LicenseUserCountRange) GetExternalUsersAdded() int32`

GetExternalUsersAdded returns the ExternalUsersAdded field if non-nil, zero value otherwise.

### GetExternalUsersAddedOk

`func (o *LicenseUserCountRange) GetExternalUsersAddedOk() (*int32, bool)`

GetExternalUsersAddedOk returns a tuple with the ExternalUsersAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUsersAdded

`func (o *LicenseUserCountRange) SetExternalUsersAdded(v int32)`

SetExternalUsersAdded sets ExternalUsersAdded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


