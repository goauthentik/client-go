# PatchedGeoIPPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**Asns** | Pointer to **[]int32** |  | [optional] 
**Countries** | Pointer to [**[]CountryCodeEnum**](CountryCodeEnum.md) |  | [optional] 

## Methods

### NewPatchedGeoIPPolicyRequest

`func NewPatchedGeoIPPolicyRequest() *PatchedGeoIPPolicyRequest`

NewPatchedGeoIPPolicyRequest instantiates a new PatchedGeoIPPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedGeoIPPolicyRequestWithDefaults

`func NewPatchedGeoIPPolicyRequestWithDefaults() *PatchedGeoIPPolicyRequest`

NewPatchedGeoIPPolicyRequestWithDefaults instantiates a new PatchedGeoIPPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedGeoIPPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedGeoIPPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedGeoIPPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedGeoIPPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExecutionLogging

`func (o *PatchedGeoIPPolicyRequest) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *PatchedGeoIPPolicyRequest) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *PatchedGeoIPPolicyRequest) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *PatchedGeoIPPolicyRequest) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetAsns

`func (o *PatchedGeoIPPolicyRequest) GetAsns() []int32`

GetAsns returns the Asns field if non-nil, zero value otherwise.

### GetAsnsOk

`func (o *PatchedGeoIPPolicyRequest) GetAsnsOk() (*[]int32, bool)`

GetAsnsOk returns a tuple with the Asns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsns

`func (o *PatchedGeoIPPolicyRequest) SetAsns(v []int32)`

SetAsns sets Asns field to given value.

### HasAsns

`func (o *PatchedGeoIPPolicyRequest) HasAsns() bool`

HasAsns returns a boolean if a field has been set.

### GetCountries

`func (o *PatchedGeoIPPolicyRequest) GetCountries() []CountryCodeEnum`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *PatchedGeoIPPolicyRequest) GetCountriesOk() (*[]CountryCodeEnum, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *PatchedGeoIPPolicyRequest) SetCountries(v []CountryCodeEnum)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *PatchedGeoIPPolicyRequest) HasCountries() bool`

HasCountries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


