/*
authentik

Making authentication simple.

API version: 2024.4.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// GeoipBindingEnum the model 'GeoipBindingEnum'
type GeoipBindingEnum string

// List of GeoipBindingEnum
const (
	GEOIPBINDINGENUM_NO_BINDING                  GeoipBindingEnum = "no_binding"
	GEOIPBINDINGENUM_BIND_CONTINENT              GeoipBindingEnum = "bind_continent"
	GEOIPBINDINGENUM_BIND_CONTINENT_COUNTRY      GeoipBindingEnum = "bind_continent_country"
	GEOIPBINDINGENUM_BIND_CONTINENT_COUNTRY_CITY GeoipBindingEnum = "bind_continent_country_city"
)

// All allowed values of GeoipBindingEnum enum
var AllowedGeoipBindingEnumEnumValues = []GeoipBindingEnum{
	"no_binding",
	"bind_continent",
	"bind_continent_country",
	"bind_continent_country_city",
}

func (v *GeoipBindingEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GeoipBindingEnum(value)
	for _, existing := range AllowedGeoipBindingEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GeoipBindingEnum", value)
}

// NewGeoipBindingEnumFromValue returns a pointer to a valid GeoipBindingEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGeoipBindingEnumFromValue(v string) (*GeoipBindingEnum, error) {
	ev := GeoipBindingEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GeoipBindingEnum: valid values are %v", v, AllowedGeoipBindingEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GeoipBindingEnum) IsValid() bool {
	for _, existing := range AllowedGeoipBindingEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeoipBindingEnum value
func (v GeoipBindingEnum) Ptr() *GeoipBindingEnum {
	return &v
}

type NullableGeoipBindingEnum struct {
	value *GeoipBindingEnum
	isSet bool
}

func (v NullableGeoipBindingEnum) Get() *GeoipBindingEnum {
	return v.value
}

func (v *NullableGeoipBindingEnum) Set(val *GeoipBindingEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoipBindingEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoipBindingEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoipBindingEnum(val *GeoipBindingEnum) *NullableGeoipBindingEnum {
	return &NullableGeoipBindingEnum{value: val, isSet: true}
}

func (v NullableGeoipBindingEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoipBindingEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
