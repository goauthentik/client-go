# TenantRecoveryKeyRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **string** |  | 
**DurationDays** | **int32** |  | 

## Methods

### NewTenantRecoveryKeyRequestRequest

`func NewTenantRecoveryKeyRequestRequest(user string, durationDays int32, ) *TenantRecoveryKeyRequestRequest`

NewTenantRecoveryKeyRequestRequest instantiates a new TenantRecoveryKeyRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantRecoveryKeyRequestRequestWithDefaults

`func NewTenantRecoveryKeyRequestRequestWithDefaults() *TenantRecoveryKeyRequestRequest`

NewTenantRecoveryKeyRequestRequestWithDefaults instantiates a new TenantRecoveryKeyRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *TenantRecoveryKeyRequestRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TenantRecoveryKeyRequestRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TenantRecoveryKeyRequestRequest) SetUser(v string)`

SetUser sets User field to given value.


### GetDurationDays

`func (o *TenantRecoveryKeyRequestRequest) GetDurationDays() int32`

GetDurationDays returns the DurationDays field if non-nil, zero value otherwise.

### GetDurationDaysOk

`func (o *TenantRecoveryKeyRequestRequest) GetDurationDaysOk() (*int32, bool)`

GetDurationDaysOk returns a tuple with the DurationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationDays

`func (o *TenantRecoveryKeyRequestRequest) SetDurationDays(v int32)`

SetDurationDays sets DurationDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


