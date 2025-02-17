# GeoIPPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**Component** | **string** | Get object component so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**BoundTo** | **int32** | Return objects policy is bound to | [readonly] 
**Asns** | Pointer to **[]int32** |  | [optional] 
**Countries** | [**[]CountryCodeEnum**](CountryCodeEnum.md) |  | 
**CountriesObj** | [**[]DetailedCountryField**](DetailedCountryField.md) |  | [readonly] 
**CheckHistoryDistance** | Pointer to **bool** |  | [optional] 
**HistoryMaxDistanceKm** | Pointer to **int64** |  | [optional] 
**DistanceToleranceKm** | Pointer to **int32** |  | [optional] 
**HistoryLoginCount** | Pointer to **int32** |  | [optional] 
**CheckImpossibleTravel** | Pointer to **bool** |  | [optional] 
**ImpossibleToleranceKm** | Pointer to **int32** |  | [optional] 

## Methods

### NewGeoIPPolicy

`func NewGeoIPPolicy(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, boundTo int32, countries []CountryCodeEnum, countriesObj []DetailedCountryField, ) *GeoIPPolicy`

NewGeoIPPolicy instantiates a new GeoIPPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoIPPolicyWithDefaults

`func NewGeoIPPolicyWithDefaults() *GeoIPPolicy`

NewGeoIPPolicyWithDefaults instantiates a new GeoIPPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *GeoIPPolicy) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *GeoIPPolicy) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *GeoIPPolicy) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *GeoIPPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GeoIPPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GeoIPPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetExecutionLogging

`func (o *GeoIPPolicy) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *GeoIPPolicy) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *GeoIPPolicy) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *GeoIPPolicy) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetComponent

`func (o *GeoIPPolicy) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *GeoIPPolicy) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *GeoIPPolicy) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *GeoIPPolicy) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *GeoIPPolicy) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *GeoIPPolicy) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *GeoIPPolicy) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *GeoIPPolicy) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *GeoIPPolicy) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *GeoIPPolicy) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *GeoIPPolicy) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *GeoIPPolicy) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetBoundTo

`func (o *GeoIPPolicy) GetBoundTo() int32`

GetBoundTo returns the BoundTo field if non-nil, zero value otherwise.

### GetBoundToOk

`func (o *GeoIPPolicy) GetBoundToOk() (*int32, bool)`

GetBoundToOk returns a tuple with the BoundTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundTo

`func (o *GeoIPPolicy) SetBoundTo(v int32)`

SetBoundTo sets BoundTo field to given value.


### GetAsns

`func (o *GeoIPPolicy) GetAsns() []int32`

GetAsns returns the Asns field if non-nil, zero value otherwise.

### GetAsnsOk

`func (o *GeoIPPolicy) GetAsnsOk() (*[]int32, bool)`

GetAsnsOk returns a tuple with the Asns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsns

`func (o *GeoIPPolicy) SetAsns(v []int32)`

SetAsns sets Asns field to given value.

### HasAsns

`func (o *GeoIPPolicy) HasAsns() bool`

HasAsns returns a boolean if a field has been set.

### GetCountries

`func (o *GeoIPPolicy) GetCountries() []CountryCodeEnum`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *GeoIPPolicy) GetCountriesOk() (*[]CountryCodeEnum, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *GeoIPPolicy) SetCountries(v []CountryCodeEnum)`

SetCountries sets Countries field to given value.


### GetCountriesObj

`func (o *GeoIPPolicy) GetCountriesObj() []DetailedCountryField`

GetCountriesObj returns the CountriesObj field if non-nil, zero value otherwise.

### GetCountriesObjOk

`func (o *GeoIPPolicy) GetCountriesObjOk() (*[]DetailedCountryField, bool)`

GetCountriesObjOk returns a tuple with the CountriesObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountriesObj

`func (o *GeoIPPolicy) SetCountriesObj(v []DetailedCountryField)`

SetCountriesObj sets CountriesObj field to given value.


### GetCheckHistoryDistance

`func (o *GeoIPPolicy) GetCheckHistoryDistance() bool`

GetCheckHistoryDistance returns the CheckHistoryDistance field if non-nil, zero value otherwise.

### GetCheckHistoryDistanceOk

`func (o *GeoIPPolicy) GetCheckHistoryDistanceOk() (*bool, bool)`

GetCheckHistoryDistanceOk returns a tuple with the CheckHistoryDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckHistoryDistance

`func (o *GeoIPPolicy) SetCheckHistoryDistance(v bool)`

SetCheckHistoryDistance sets CheckHistoryDistance field to given value.

### HasCheckHistoryDistance

`func (o *GeoIPPolicy) HasCheckHistoryDistance() bool`

HasCheckHistoryDistance returns a boolean if a field has been set.

### GetHistoryMaxDistanceKm

`func (o *GeoIPPolicy) GetHistoryMaxDistanceKm() int64`

GetHistoryMaxDistanceKm returns the HistoryMaxDistanceKm field if non-nil, zero value otherwise.

### GetHistoryMaxDistanceKmOk

`func (o *GeoIPPolicy) GetHistoryMaxDistanceKmOk() (*int64, bool)`

GetHistoryMaxDistanceKmOk returns a tuple with the HistoryMaxDistanceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryMaxDistanceKm

`func (o *GeoIPPolicy) SetHistoryMaxDistanceKm(v int64)`

SetHistoryMaxDistanceKm sets HistoryMaxDistanceKm field to given value.

### HasHistoryMaxDistanceKm

`func (o *GeoIPPolicy) HasHistoryMaxDistanceKm() bool`

HasHistoryMaxDistanceKm returns a boolean if a field has been set.

### GetDistanceToleranceKm

`func (o *GeoIPPolicy) GetDistanceToleranceKm() int32`

GetDistanceToleranceKm returns the DistanceToleranceKm field if non-nil, zero value otherwise.

### GetDistanceToleranceKmOk

`func (o *GeoIPPolicy) GetDistanceToleranceKmOk() (*int32, bool)`

GetDistanceToleranceKmOk returns a tuple with the DistanceToleranceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceToleranceKm

`func (o *GeoIPPolicy) SetDistanceToleranceKm(v int32)`

SetDistanceToleranceKm sets DistanceToleranceKm field to given value.

### HasDistanceToleranceKm

`func (o *GeoIPPolicy) HasDistanceToleranceKm() bool`

HasDistanceToleranceKm returns a boolean if a field has been set.

### GetHistoryLoginCount

`func (o *GeoIPPolicy) GetHistoryLoginCount() int32`

GetHistoryLoginCount returns the HistoryLoginCount field if non-nil, zero value otherwise.

### GetHistoryLoginCountOk

`func (o *GeoIPPolicy) GetHistoryLoginCountOk() (*int32, bool)`

GetHistoryLoginCountOk returns a tuple with the HistoryLoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryLoginCount

`func (o *GeoIPPolicy) SetHistoryLoginCount(v int32)`

SetHistoryLoginCount sets HistoryLoginCount field to given value.

### HasHistoryLoginCount

`func (o *GeoIPPolicy) HasHistoryLoginCount() bool`

HasHistoryLoginCount returns a boolean if a field has been set.

### GetCheckImpossibleTravel

`func (o *GeoIPPolicy) GetCheckImpossibleTravel() bool`

GetCheckImpossibleTravel returns the CheckImpossibleTravel field if non-nil, zero value otherwise.

### GetCheckImpossibleTravelOk

`func (o *GeoIPPolicy) GetCheckImpossibleTravelOk() (*bool, bool)`

GetCheckImpossibleTravelOk returns a tuple with the CheckImpossibleTravel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckImpossibleTravel

`func (o *GeoIPPolicy) SetCheckImpossibleTravel(v bool)`

SetCheckImpossibleTravel sets CheckImpossibleTravel field to given value.

### HasCheckImpossibleTravel

`func (o *GeoIPPolicy) HasCheckImpossibleTravel() bool`

HasCheckImpossibleTravel returns a boolean if a field has been set.

### GetImpossibleToleranceKm

`func (o *GeoIPPolicy) GetImpossibleToleranceKm() int32`

GetImpossibleToleranceKm returns the ImpossibleToleranceKm field if non-nil, zero value otherwise.

### GetImpossibleToleranceKmOk

`func (o *GeoIPPolicy) GetImpossibleToleranceKmOk() (*int32, bool)`

GetImpossibleToleranceKmOk returns a tuple with the ImpossibleToleranceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpossibleToleranceKm

`func (o *GeoIPPolicy) SetImpossibleToleranceKm(v int32)`

SetImpossibleToleranceKm sets ImpossibleToleranceKm field to given value.

### HasImpossibleToleranceKm

`func (o *GeoIPPolicy) HasImpossibleToleranceKm() bool`

HasImpossibleToleranceKm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


