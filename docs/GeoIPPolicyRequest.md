# GeoIPPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**Asns** | Pointer to **[]int32** |  | [optional] 
**Countries** | [**[]CountryCodeEnum**](CountryCodeEnum.md) |  | 
**CheckHistoryDistance** | Pointer to **bool** |  | [optional] 
**HistoryMaxDistanceKm** | Pointer to **int64** |  | [optional] 
**DistanceToleranceKm** | Pointer to **int32** |  | [optional] 
**HistoryLoginCount** | Pointer to **int32** |  | [optional] 
**CheckImpossibleTravel** | Pointer to **bool** |  | [optional] 
**ImpossibleToleranceKm** | Pointer to **int32** |  | [optional] 

## Methods

### NewGeoIPPolicyRequest

`func NewGeoIPPolicyRequest(name string, countries []CountryCodeEnum, ) *GeoIPPolicyRequest`

NewGeoIPPolicyRequest instantiates a new GeoIPPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoIPPolicyRequestWithDefaults

`func NewGeoIPPolicyRequestWithDefaults() *GeoIPPolicyRequest`

NewGeoIPPolicyRequestWithDefaults instantiates a new GeoIPPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GeoIPPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GeoIPPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GeoIPPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExecutionLogging

`func (o *GeoIPPolicyRequest) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *GeoIPPolicyRequest) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *GeoIPPolicyRequest) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *GeoIPPolicyRequest) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetAsns

`func (o *GeoIPPolicyRequest) GetAsns() []int32`

GetAsns returns the Asns field if non-nil, zero value otherwise.

### GetAsnsOk

`func (o *GeoIPPolicyRequest) GetAsnsOk() (*[]int32, bool)`

GetAsnsOk returns a tuple with the Asns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsns

`func (o *GeoIPPolicyRequest) SetAsns(v []int32)`

SetAsns sets Asns field to given value.

### HasAsns

`func (o *GeoIPPolicyRequest) HasAsns() bool`

HasAsns returns a boolean if a field has been set.

### GetCountries

`func (o *GeoIPPolicyRequest) GetCountries() []CountryCodeEnum`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *GeoIPPolicyRequest) GetCountriesOk() (*[]CountryCodeEnum, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *GeoIPPolicyRequest) SetCountries(v []CountryCodeEnum)`

SetCountries sets Countries field to given value.


### GetCheckHistoryDistance

`func (o *GeoIPPolicyRequest) GetCheckHistoryDistance() bool`

GetCheckHistoryDistance returns the CheckHistoryDistance field if non-nil, zero value otherwise.

### GetCheckHistoryDistanceOk

`func (o *GeoIPPolicyRequest) GetCheckHistoryDistanceOk() (*bool, bool)`

GetCheckHistoryDistanceOk returns a tuple with the CheckHistoryDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckHistoryDistance

`func (o *GeoIPPolicyRequest) SetCheckHistoryDistance(v bool)`

SetCheckHistoryDistance sets CheckHistoryDistance field to given value.

### HasCheckHistoryDistance

`func (o *GeoIPPolicyRequest) HasCheckHistoryDistance() bool`

HasCheckHistoryDistance returns a boolean if a field has been set.

### GetHistoryMaxDistanceKm

`func (o *GeoIPPolicyRequest) GetHistoryMaxDistanceKm() int64`

GetHistoryMaxDistanceKm returns the HistoryMaxDistanceKm field if non-nil, zero value otherwise.

### GetHistoryMaxDistanceKmOk

`func (o *GeoIPPolicyRequest) GetHistoryMaxDistanceKmOk() (*int64, bool)`

GetHistoryMaxDistanceKmOk returns a tuple with the HistoryMaxDistanceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryMaxDistanceKm

`func (o *GeoIPPolicyRequest) SetHistoryMaxDistanceKm(v int64)`

SetHistoryMaxDistanceKm sets HistoryMaxDistanceKm field to given value.

### HasHistoryMaxDistanceKm

`func (o *GeoIPPolicyRequest) HasHistoryMaxDistanceKm() bool`

HasHistoryMaxDistanceKm returns a boolean if a field has been set.

### GetDistanceToleranceKm

`func (o *GeoIPPolicyRequest) GetDistanceToleranceKm() int32`

GetDistanceToleranceKm returns the DistanceToleranceKm field if non-nil, zero value otherwise.

### GetDistanceToleranceKmOk

`func (o *GeoIPPolicyRequest) GetDistanceToleranceKmOk() (*int32, bool)`

GetDistanceToleranceKmOk returns a tuple with the DistanceToleranceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceToleranceKm

`func (o *GeoIPPolicyRequest) SetDistanceToleranceKm(v int32)`

SetDistanceToleranceKm sets DistanceToleranceKm field to given value.

### HasDistanceToleranceKm

`func (o *GeoIPPolicyRequest) HasDistanceToleranceKm() bool`

HasDistanceToleranceKm returns a boolean if a field has been set.

### GetHistoryLoginCount

`func (o *GeoIPPolicyRequest) GetHistoryLoginCount() int32`

GetHistoryLoginCount returns the HistoryLoginCount field if non-nil, zero value otherwise.

### GetHistoryLoginCountOk

`func (o *GeoIPPolicyRequest) GetHistoryLoginCountOk() (*int32, bool)`

GetHistoryLoginCountOk returns a tuple with the HistoryLoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryLoginCount

`func (o *GeoIPPolicyRequest) SetHistoryLoginCount(v int32)`

SetHistoryLoginCount sets HistoryLoginCount field to given value.

### HasHistoryLoginCount

`func (o *GeoIPPolicyRequest) HasHistoryLoginCount() bool`

HasHistoryLoginCount returns a boolean if a field has been set.

### GetCheckImpossibleTravel

`func (o *GeoIPPolicyRequest) GetCheckImpossibleTravel() bool`

GetCheckImpossibleTravel returns the CheckImpossibleTravel field if non-nil, zero value otherwise.

### GetCheckImpossibleTravelOk

`func (o *GeoIPPolicyRequest) GetCheckImpossibleTravelOk() (*bool, bool)`

GetCheckImpossibleTravelOk returns a tuple with the CheckImpossibleTravel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckImpossibleTravel

`func (o *GeoIPPolicyRequest) SetCheckImpossibleTravel(v bool)`

SetCheckImpossibleTravel sets CheckImpossibleTravel field to given value.

### HasCheckImpossibleTravel

`func (o *GeoIPPolicyRequest) HasCheckImpossibleTravel() bool`

HasCheckImpossibleTravel returns a boolean if a field has been set.

### GetImpossibleToleranceKm

`func (o *GeoIPPolicyRequest) GetImpossibleToleranceKm() int32`

GetImpossibleToleranceKm returns the ImpossibleToleranceKm field if non-nil, zero value otherwise.

### GetImpossibleToleranceKmOk

`func (o *GeoIPPolicyRequest) GetImpossibleToleranceKmOk() (*int32, bool)`

GetImpossibleToleranceKmOk returns a tuple with the ImpossibleToleranceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpossibleToleranceKm

`func (o *GeoIPPolicyRequest) SetImpossibleToleranceKm(v int32)`

SetImpossibleToleranceKm sets ImpossibleToleranceKm field to given value.

### HasImpossibleToleranceKm

`func (o *GeoIPPolicyRequest) HasImpossibleToleranceKm() bool`

HasImpossibleToleranceKm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


