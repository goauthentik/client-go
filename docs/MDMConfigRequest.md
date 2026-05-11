# MDMConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | [**DeviceFactsOSFamily**](DeviceFactsOSFamily.md) |  | 
**EnrollmentToken** | **string** |  | 

## Methods

### NewMDMConfigRequest

`func NewMDMConfigRequest(platform DeviceFactsOSFamily, enrollmentToken string, ) *MDMConfigRequest`

NewMDMConfigRequest instantiates a new MDMConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMDMConfigRequestWithDefaults

`func NewMDMConfigRequestWithDefaults() *MDMConfigRequest`

NewMDMConfigRequestWithDefaults instantiates a new MDMConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *MDMConfigRequest) GetPlatform() DeviceFactsOSFamily`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *MDMConfigRequest) GetPlatformOk() (*DeviceFactsOSFamily, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *MDMConfigRequest) SetPlatform(v DeviceFactsOSFamily)`

SetPlatform sets Platform field to given value.


### GetEnrollmentToken

`func (o *MDMConfigRequest) GetEnrollmentToken() string`

GetEnrollmentToken returns the EnrollmentToken field if non-nil, zero value otherwise.

### GetEnrollmentTokenOk

`func (o *MDMConfigRequest) GetEnrollmentTokenOk() (*string, bool)`

GetEnrollmentTokenOk returns a tuple with the EnrollmentToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentToken

`func (o *MDMConfigRequest) SetEnrollmentToken(v string)`

SetEnrollmentToken sets EnrollmentToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


