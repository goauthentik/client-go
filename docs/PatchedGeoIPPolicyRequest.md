# PatchedGeoIPPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**Asns** | Pointer to **[]int32** |  | [optional] 
**Countries** | Pointer to [**[]CountryCodeEnum**](CountryCodeEnum.md) |  | [optional] 
**CheckHistoryDistance** | Pointer to **bool** |  | [optional] 
**HistoryMaxDistanceKm** | Pointer to **int64** |  | [optional] 
**DistanceToleranceKm** | Pointer to **int32** |  | [optional] 
**HistoryLoginCount** | Pointer to **int32** |  | [optional] 
**CheckImpossibleTravel** | Pointer to **bool** |  | [optional] 
**ImpossibleToleranceKm** | Pointer to **int32** |  | [optional] 

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

### GetCheckHistoryDistance

`func (o *PatchedGeoIPPolicyRequest) GetCheckHistoryDistance() bool`

GetCheckHistoryDistance returns the CheckHistoryDistance field if non-nil, zero value otherwise.

### GetCheckHistoryDistanceOk

`func (o *PatchedGeoIPPolicyRequest) GetCheckHistoryDistanceOk() (*bool, bool)`

GetCheckHistoryDistanceOk returns a tuple with the CheckHistoryDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckHistoryDistance

`func (o *PatchedGeoIPPolicyRequest) SetCheckHistoryDistance(v bool)`

SetCheckHistoryDistance sets CheckHistoryDistance field to given value.

### HasCheckHistoryDistance

`func (o *PatchedGeoIPPolicyRequest) HasCheckHistoryDistance() bool`

HasCheckHistoryDistance returns a boolean if a field has been set.

### GetHistoryMaxDistanceKm

`func (o *PatchedGeoIPPolicyRequest) GetHistoryMaxDistanceKm() int64`

GetHistoryMaxDistanceKm returns the HistoryMaxDistanceKm field if non-nil, zero value otherwise.

### GetHistoryMaxDistanceKmOk

`func (o *PatchedGeoIPPolicyRequest) GetHistoryMaxDistanceKmOk() (*int64, bool)`

GetHistoryMaxDistanceKmOk returns a tuple with the HistoryMaxDistanceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryMaxDistanceKm

`func (o *PatchedGeoIPPolicyRequest) SetHistoryMaxDistanceKm(v int64)`

SetHistoryMaxDistanceKm sets HistoryMaxDistanceKm field to given value.

### HasHistoryMaxDistanceKm

`func (o *PatchedGeoIPPolicyRequest) HasHistoryMaxDistanceKm() bool`

HasHistoryMaxDistanceKm returns a boolean if a field has been set.

### GetDistanceToleranceKm

`func (o *PatchedGeoIPPolicyRequest) GetDistanceToleranceKm() int32`

GetDistanceToleranceKm returns the DistanceToleranceKm field if non-nil, zero value otherwise.

### GetDistanceToleranceKmOk

`func (o *PatchedGeoIPPolicyRequest) GetDistanceToleranceKmOk() (*int32, bool)`

GetDistanceToleranceKmOk returns a tuple with the DistanceToleranceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceToleranceKm

`func (o *PatchedGeoIPPolicyRequest) SetDistanceToleranceKm(v int32)`

SetDistanceToleranceKm sets DistanceToleranceKm field to given value.

### HasDistanceToleranceKm

`func (o *PatchedGeoIPPolicyRequest) HasDistanceToleranceKm() bool`

HasDistanceToleranceKm returns a boolean if a field has been set.

### GetHistoryLoginCount

`func (o *PatchedGeoIPPolicyRequest) GetHistoryLoginCount() int32`

GetHistoryLoginCount returns the HistoryLoginCount field if non-nil, zero value otherwise.

### GetHistoryLoginCountOk

`func (o *PatchedGeoIPPolicyRequest) GetHistoryLoginCountOk() (*int32, bool)`

GetHistoryLoginCountOk returns a tuple with the HistoryLoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryLoginCount

`func (o *PatchedGeoIPPolicyRequest) SetHistoryLoginCount(v int32)`

SetHistoryLoginCount sets HistoryLoginCount field to given value.

### HasHistoryLoginCount

`func (o *PatchedGeoIPPolicyRequest) HasHistoryLoginCount() bool`

HasHistoryLoginCount returns a boolean if a field has been set.

### GetCheckImpossibleTravel

`func (o *PatchedGeoIPPolicyRequest) GetCheckImpossibleTravel() bool`

GetCheckImpossibleTravel returns the CheckImpossibleTravel field if non-nil, zero value otherwise.

### GetCheckImpossibleTravelOk

`func (o *PatchedGeoIPPolicyRequest) GetCheckImpossibleTravelOk() (*bool, bool)`

GetCheckImpossibleTravelOk returns a tuple with the CheckImpossibleTravel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckImpossibleTravel

`func (o *PatchedGeoIPPolicyRequest) SetCheckImpossibleTravel(v bool)`

SetCheckImpossibleTravel sets CheckImpossibleTravel field to given value.

### HasCheckImpossibleTravel

`func (o *PatchedGeoIPPolicyRequest) HasCheckImpossibleTravel() bool`

HasCheckImpossibleTravel returns a boolean if a field has been set.

### GetImpossibleToleranceKm

`func (o *PatchedGeoIPPolicyRequest) GetImpossibleToleranceKm() int32`

GetImpossibleToleranceKm returns the ImpossibleToleranceKm field if non-nil, zero value otherwise.

### GetImpossibleToleranceKmOk

`func (o *PatchedGeoIPPolicyRequest) GetImpossibleToleranceKmOk() (*int32, bool)`

GetImpossibleToleranceKmOk returns a tuple with the ImpossibleToleranceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpossibleToleranceKm

`func (o *PatchedGeoIPPolicyRequest) SetImpossibleToleranceKm(v int32)`

SetImpossibleToleranceKm sets ImpossibleToleranceKm field to given value.

### HasImpossibleToleranceKm

`func (o *PatchedGeoIPPolicyRequest) HasImpossibleToleranceKm() bool`

HasImpossibleToleranceKm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


