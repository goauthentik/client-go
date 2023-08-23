# License

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseUuid** | **string** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**Key** | **string** |  | 
**Expiry** | **time.Time** |  | [readonly] 
**InternalUsers** | **int32** |  | [readonly] 
**ExternalUsers** | **int32** |  | [readonly] 

## Methods

### NewLicense

`func NewLicense(licenseUuid string, name string, key string, expiry time.Time, internalUsers int32, externalUsers int32, ) *License`

NewLicense instantiates a new License object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseWithDefaults

`func NewLicenseWithDefaults() *License`

NewLicenseWithDefaults instantiates a new License object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseUuid

`func (o *License) GetLicenseUuid() string`

GetLicenseUuid returns the LicenseUuid field if non-nil, zero value otherwise.

### GetLicenseUuidOk

`func (o *License) GetLicenseUuidOk() (*string, bool)`

GetLicenseUuidOk returns a tuple with the LicenseUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUuid

`func (o *License) SetLicenseUuid(v string)`

SetLicenseUuid sets LicenseUuid field to given value.


### GetName

`func (o *License) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *License) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *License) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *License) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *License) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *License) SetKey(v string)`

SetKey sets Key field to given value.


### GetExpiry

`func (o *License) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *License) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *License) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.


### GetInternalUsers

`func (o *License) GetInternalUsers() int32`

GetInternalUsers returns the InternalUsers field if non-nil, zero value otherwise.

### GetInternalUsersOk

`func (o *License) GetInternalUsersOk() (*int32, bool)`

GetInternalUsersOk returns a tuple with the InternalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalUsers

`func (o *License) SetInternalUsers(v int32)`

SetInternalUsers sets InternalUsers field to given value.


### GetExternalUsers

`func (o *License) GetExternalUsers() int32`

GetExternalUsers returns the ExternalUsers field if non-nil, zero value otherwise.

### GetExternalUsersOk

`func (o *License) GetExternalUsersOk() (*int32, bool)`

GetExternalUsersOk returns a tuple with the ExternalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUsers

`func (o *License) SetExternalUsers(v int32)`

SetExternalUsers sets ExternalUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


