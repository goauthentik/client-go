/*
authentik

Making authentication simple.

API version: 2025.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GeoIPPolicy GeoIP Policy Serializer
type GeoIPPolicy struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	// Get object component so that we know how to edit the object
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName string `json:"meta_model_name"`
	// Return objects policy is bound to
	BoundTo               int32                  `json:"bound_to"`
	Asns                  []int32                `json:"asns,omitempty"`
	Countries             []CountryCodeEnum      `json:"countries"`
	CountriesObj          []DetailedCountryField `json:"countries_obj"`
	CheckHistoryDistance  *bool                  `json:"check_history_distance,omitempty"`
	HistoryMaxDistanceKm  *int64                 `json:"history_max_distance_km,omitempty"`
	DistanceToleranceKm   *int32                 `json:"distance_tolerance_km,omitempty"`
	HistoryLoginCount     *int32                 `json:"history_login_count,omitempty"`
	CheckImpossibleTravel *bool                  `json:"check_impossible_travel,omitempty"`
	ImpossibleToleranceKm *int32                 `json:"impossible_tolerance_km,omitempty"`
}

// NewGeoIPPolicy instantiates a new GeoIPPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoIPPolicy(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, boundTo int32, countries []CountryCodeEnum, countriesObj []DetailedCountryField) *GeoIPPolicy {
	this := GeoIPPolicy{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.BoundTo = boundTo
	this.Countries = countries
	this.CountriesObj = countriesObj
	return &this
}

// NewGeoIPPolicyWithDefaults instantiates a new GeoIPPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoIPPolicyWithDefaults() *GeoIPPolicy {
	this := GeoIPPolicy{}
	return &this
}

// GetPk returns the Pk field value
func (o *GeoIPPolicy) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *GeoIPPolicy) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *GeoIPPolicy) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *GeoIPPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GeoIPPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GeoIPPolicy) SetName(v string) {
	o.Name = v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *GeoIPPolicy) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoIPPolicy) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *GeoIPPolicy) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *GeoIPPolicy) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetComponent returns the Component field value
func (o *GeoIPPolicy) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *GeoIPPolicy) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *GeoIPPolicy) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *GeoIPPolicy) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *GeoIPPolicy) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *GeoIPPolicy) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *GeoIPPolicy) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *GeoIPPolicy) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *GeoIPPolicy) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *GeoIPPolicy) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *GeoIPPolicy) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *GeoIPPolicy) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetBoundTo returns the BoundTo field value
func (o *GeoIPPolicy) GetBoundTo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BoundTo
}

// GetBoundToOk returns a tuple with the BoundTo field value
// and a boolean to check if the value has been set.
func (o *GeoIPPolicy) GetBoundToOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoundTo, true
}

// SetBoundTo sets field value
func (o *GeoIPPolicy) SetBoundTo(v int32) {
	o.BoundTo = v
}

// GetAsns returns the Asns field value if set, zero value otherwise.
func (o *GeoIPPolicy) GetAsns() []int32 {
	if o == nil || o.Asns == nil {
		var ret []int32
		return ret
	}
	return o.Asns
}

// GetAsnsOk returns a tuple with the Asns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoIPPolicy) GetAsnsOk() ([]int32, bool) {
	if o == nil || o.Asns == nil {
		return nil, false
	}
	return o.Asns, true
}

// HasAsns returns a boolean if a field has been set.
func (o *GeoIPPolicy) HasAsns() bool {
	if o != nil && o.Asns != nil {
		return true
	}

	return false
}

// SetAsns gets a reference to the given []int32 and assigns it to the Asns field.
func (o *GeoIPPolicy) SetAsns(v []int32) {
	o.Asns = v
}

// GetCountries returns the Countries field value
func (o *GeoIPPolicy) GetCountries() []CountryCodeEnum {
	if o == nil {
		var ret []CountryCodeEnum
		return ret
	}

	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value
// and a boolean to check if the value has been set.
func (o *GeoIPPolicy) GetCountriesOk() ([]CountryCodeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Countries, true
}

// SetCountries sets field value
func (o *GeoIPPolicy) SetCountries(v []CountryCodeEnum) {
	o.Countries = v
}

// GetCountriesObj returns the CountriesObj field value
func (o *GeoIPPolicy) GetCountriesObj() []DetailedCountryField {
	if o == nil {
		var ret []DetailedCountryField
		return ret
	}

	return o.CountriesObj
}

// GetCountriesObjOk returns a tuple with the CountriesObj field value
// and a boolean to check if the value has been set.
func (o *GeoIPPolicy) GetCountriesObjOk() ([]DetailedCountryField, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountriesObj, true
}

// SetCountriesObj sets field value
func (o *GeoIPPolicy) SetCountriesObj(v []DetailedCountryField) {
	o.CountriesObj = v
}

// GetCheckHistoryDistance returns the CheckHistoryDistance field value if set, zero value otherwise.
func (o *GeoIPPolicy) GetCheckHistoryDistance() bool {
	if o == nil || o.CheckHistoryDistance == nil {
		var ret bool
		return ret
	}
	return *o.CheckHistoryDistance
}

// GetCheckHistoryDistanceOk returns a tuple with the CheckHistoryDistance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoIPPolicy) GetCheckHistoryDistanceOk() (*bool, bool) {
	if o == nil || o.CheckHistoryDistance == nil {
		return nil, false
	}
	return o.CheckHistoryDistance, true
}

// HasCheckHistoryDistance returns a boolean if a field has been set.
func (o *GeoIPPolicy) HasCheckHistoryDistance() bool {
	if o != nil && o.CheckHistoryDistance != nil {
		return true
	}

	return false
}

// SetCheckHistoryDistance gets a reference to the given bool and assigns it to the CheckHistoryDistance field.
func (o *GeoIPPolicy) SetCheckHistoryDistance(v bool) {
	o.CheckHistoryDistance = &v
}

// GetHistoryMaxDistanceKm returns the HistoryMaxDistanceKm field value if set, zero value otherwise.
func (o *GeoIPPolicy) GetHistoryMaxDistanceKm() int64 {
	if o == nil || o.HistoryMaxDistanceKm == nil {
		var ret int64
		return ret
	}
	return *o.HistoryMaxDistanceKm
}

// GetHistoryMaxDistanceKmOk returns a tuple with the HistoryMaxDistanceKm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoIPPolicy) GetHistoryMaxDistanceKmOk() (*int64, bool) {
	if o == nil || o.HistoryMaxDistanceKm == nil {
		return nil, false
	}
	return o.HistoryMaxDistanceKm, true
}

// HasHistoryMaxDistanceKm returns a boolean if a field has been set.
func (o *GeoIPPolicy) HasHistoryMaxDistanceKm() bool {
	if o != nil && o.HistoryMaxDistanceKm != nil {
		return true
	}

	return false
}

// SetHistoryMaxDistanceKm gets a reference to the given int64 and assigns it to the HistoryMaxDistanceKm field.
func (o *GeoIPPolicy) SetHistoryMaxDistanceKm(v int64) {
	o.HistoryMaxDistanceKm = &v
}

// GetDistanceToleranceKm returns the DistanceToleranceKm field value if set, zero value otherwise.
func (o *GeoIPPolicy) GetDistanceToleranceKm() int32 {
	if o == nil || o.DistanceToleranceKm == nil {
		var ret int32
		return ret
	}
	return *o.DistanceToleranceKm
}

// GetDistanceToleranceKmOk returns a tuple with the DistanceToleranceKm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoIPPolicy) GetDistanceToleranceKmOk() (*int32, bool) {
	if o == nil || o.DistanceToleranceKm == nil {
		return nil, false
	}
	return o.DistanceToleranceKm, true
}

// HasDistanceToleranceKm returns a boolean if a field has been set.
func (o *GeoIPPolicy) HasDistanceToleranceKm() bool {
	if o != nil && o.DistanceToleranceKm != nil {
		return true
	}

	return false
}

// SetDistanceToleranceKm gets a reference to the given int32 and assigns it to the DistanceToleranceKm field.
func (o *GeoIPPolicy) SetDistanceToleranceKm(v int32) {
	o.DistanceToleranceKm = &v
}

// GetHistoryLoginCount returns the HistoryLoginCount field value if set, zero value otherwise.
func (o *GeoIPPolicy) GetHistoryLoginCount() int32 {
	if o == nil || o.HistoryLoginCount == nil {
		var ret int32
		return ret
	}
	return *o.HistoryLoginCount
}

// GetHistoryLoginCountOk returns a tuple with the HistoryLoginCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoIPPolicy) GetHistoryLoginCountOk() (*int32, bool) {
	if o == nil || o.HistoryLoginCount == nil {
		return nil, false
	}
	return o.HistoryLoginCount, true
}

// HasHistoryLoginCount returns a boolean if a field has been set.
func (o *GeoIPPolicy) HasHistoryLoginCount() bool {
	if o != nil && o.HistoryLoginCount != nil {
		return true
	}

	return false
}

// SetHistoryLoginCount gets a reference to the given int32 and assigns it to the HistoryLoginCount field.
func (o *GeoIPPolicy) SetHistoryLoginCount(v int32) {
	o.HistoryLoginCount = &v
}

// GetCheckImpossibleTravel returns the CheckImpossibleTravel field value if set, zero value otherwise.
func (o *GeoIPPolicy) GetCheckImpossibleTravel() bool {
	if o == nil || o.CheckImpossibleTravel == nil {
		var ret bool
		return ret
	}
	return *o.CheckImpossibleTravel
}

// GetCheckImpossibleTravelOk returns a tuple with the CheckImpossibleTravel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoIPPolicy) GetCheckImpossibleTravelOk() (*bool, bool) {
	if o == nil || o.CheckImpossibleTravel == nil {
		return nil, false
	}
	return o.CheckImpossibleTravel, true
}

// HasCheckImpossibleTravel returns a boolean if a field has been set.
func (o *GeoIPPolicy) HasCheckImpossibleTravel() bool {
	if o != nil && o.CheckImpossibleTravel != nil {
		return true
	}

	return false
}

// SetCheckImpossibleTravel gets a reference to the given bool and assigns it to the CheckImpossibleTravel field.
func (o *GeoIPPolicy) SetCheckImpossibleTravel(v bool) {
	o.CheckImpossibleTravel = &v
}

// GetImpossibleToleranceKm returns the ImpossibleToleranceKm field value if set, zero value otherwise.
func (o *GeoIPPolicy) GetImpossibleToleranceKm() int32 {
	if o == nil || o.ImpossibleToleranceKm == nil {
		var ret int32
		return ret
	}
	return *o.ImpossibleToleranceKm
}

// GetImpossibleToleranceKmOk returns a tuple with the ImpossibleToleranceKm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoIPPolicy) GetImpossibleToleranceKmOk() (*int32, bool) {
	if o == nil || o.ImpossibleToleranceKm == nil {
		return nil, false
	}
	return o.ImpossibleToleranceKm, true
}

// HasImpossibleToleranceKm returns a boolean if a field has been set.
func (o *GeoIPPolicy) HasImpossibleToleranceKm() bool {
	if o != nil && o.ImpossibleToleranceKm != nil {
		return true
	}

	return false
}

// SetImpossibleToleranceKm gets a reference to the given int32 and assigns it to the ImpossibleToleranceKm field.
func (o *GeoIPPolicy) SetImpossibleToleranceKm(v int32) {
	o.ImpossibleToleranceKm = &v
}

func (o GeoIPPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["verbose_name"] = o.VerboseName
	}
	if true {
		toSerialize["verbose_name_plural"] = o.VerboseNamePlural
	}
	if true {
		toSerialize["meta_model_name"] = o.MetaModelName
	}
	if true {
		toSerialize["bound_to"] = o.BoundTo
	}
	if o.Asns != nil {
		toSerialize["asns"] = o.Asns
	}
	if true {
		toSerialize["countries"] = o.Countries
	}
	if true {
		toSerialize["countries_obj"] = o.CountriesObj
	}
	if o.CheckHistoryDistance != nil {
		toSerialize["check_history_distance"] = o.CheckHistoryDistance
	}
	if o.HistoryMaxDistanceKm != nil {
		toSerialize["history_max_distance_km"] = o.HistoryMaxDistanceKm
	}
	if o.DistanceToleranceKm != nil {
		toSerialize["distance_tolerance_km"] = o.DistanceToleranceKm
	}
	if o.HistoryLoginCount != nil {
		toSerialize["history_login_count"] = o.HistoryLoginCount
	}
	if o.CheckImpossibleTravel != nil {
		toSerialize["check_impossible_travel"] = o.CheckImpossibleTravel
	}
	if o.ImpossibleToleranceKm != nil {
		toSerialize["impossible_tolerance_km"] = o.ImpossibleToleranceKm
	}
	return json.Marshal(toSerialize)
}

type NullableGeoIPPolicy struct {
	value *GeoIPPolicy
	isSet bool
}

func (v NullableGeoIPPolicy) Get() *GeoIPPolicy {
	return v.value
}

func (v *NullableGeoIPPolicy) Set(val *GeoIPPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoIPPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoIPPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoIPPolicy(val *GeoIPPolicy) *NullableGeoIPPolicy {
	return &NullableGeoIPPolicy{value: val, isSet: true}
}

func (v NullableGeoIPPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoIPPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
