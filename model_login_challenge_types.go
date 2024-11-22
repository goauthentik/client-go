/*
authentik

Making authentication simple.

API version: 2024.10.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// LoginChallengeTypes - struct for LoginChallengeTypes
type LoginChallengeTypes struct {
	AppleLoginChallenge         *AppleLoginChallenge
	PlexAuthenticationChallenge *PlexAuthenticationChallenge
	RedirectChallenge           *RedirectChallenge
}

// AppleLoginChallengeAsLoginChallengeTypes is a convenience function that returns AppleLoginChallenge wrapped in LoginChallengeTypes
func AppleLoginChallengeAsLoginChallengeTypes(v *AppleLoginChallenge) LoginChallengeTypes {
	return LoginChallengeTypes{
		AppleLoginChallenge: v,
	}
}

// PlexAuthenticationChallengeAsLoginChallengeTypes is a convenience function that returns PlexAuthenticationChallenge wrapped in LoginChallengeTypes
func PlexAuthenticationChallengeAsLoginChallengeTypes(v *PlexAuthenticationChallenge) LoginChallengeTypes {
	return LoginChallengeTypes{
		PlexAuthenticationChallenge: v,
	}
}

// RedirectChallengeAsLoginChallengeTypes is a convenience function that returns RedirectChallenge wrapped in LoginChallengeTypes
func RedirectChallengeAsLoginChallengeTypes(v *RedirectChallenge) LoginChallengeTypes {
	return LoginChallengeTypes{
		RedirectChallenge: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *LoginChallengeTypes) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'AppleLoginChallenge'
	if jsonDict["component"] == "AppleLoginChallenge" {
		// try to unmarshal JSON data into AppleLoginChallenge
		err = json.Unmarshal(data, &dst.AppleLoginChallenge)
		if err == nil {
			return nil // data stored in dst.AppleLoginChallenge, return on the first match
		} else {
			dst.AppleLoginChallenge = nil
			return fmt.Errorf("Failed to unmarshal LoginChallengeTypes as AppleLoginChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PlexAuthenticationChallenge'
	if jsonDict["component"] == "PlexAuthenticationChallenge" {
		// try to unmarshal JSON data into PlexAuthenticationChallenge
		err = json.Unmarshal(data, &dst.PlexAuthenticationChallenge)
		if err == nil {
			return nil // data stored in dst.PlexAuthenticationChallenge, return on the first match
		} else {
			dst.PlexAuthenticationChallenge = nil
			return fmt.Errorf("Failed to unmarshal LoginChallengeTypes as PlexAuthenticationChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RedirectChallenge'
	if jsonDict["component"] == "RedirectChallenge" {
		// try to unmarshal JSON data into RedirectChallenge
		err = json.Unmarshal(data, &dst.RedirectChallenge)
		if err == nil {
			return nil // data stored in dst.RedirectChallenge, return on the first match
		} else {
			dst.RedirectChallenge = nil
			return fmt.Errorf("Failed to unmarshal LoginChallengeTypes as RedirectChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-source-oauth-apple'
	if jsonDict["component"] == "ak-source-oauth-apple" {
		// try to unmarshal JSON data into AppleLoginChallenge
		err = json.Unmarshal(data, &dst.AppleLoginChallenge)
		if err == nil {
			return nil // data stored in dst.AppleLoginChallenge, return on the first match
		} else {
			dst.AppleLoginChallenge = nil
			return fmt.Errorf("Failed to unmarshal LoginChallengeTypes as AppleLoginChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ak-source-plex'
	if jsonDict["component"] == "ak-source-plex" {
		// try to unmarshal JSON data into PlexAuthenticationChallenge
		err = json.Unmarshal(data, &dst.PlexAuthenticationChallenge)
		if err == nil {
			return nil // data stored in dst.PlexAuthenticationChallenge, return on the first match
		} else {
			dst.PlexAuthenticationChallenge = nil
			return fmt.Errorf("Failed to unmarshal LoginChallengeTypes as PlexAuthenticationChallenge: %s", err.Error())
		}
	}

	// check if the discriminator value is 'xak-flow-redirect'
	if jsonDict["component"] == "xak-flow-redirect" {
		// try to unmarshal JSON data into RedirectChallenge
		err = json.Unmarshal(data, &dst.RedirectChallenge)
		if err == nil {
			return nil // data stored in dst.RedirectChallenge, return on the first match
		} else {
			dst.RedirectChallenge = nil
			return fmt.Errorf("Failed to unmarshal LoginChallengeTypes as RedirectChallenge: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LoginChallengeTypes) MarshalJSON() ([]byte, error) {
	if src.AppleLoginChallenge != nil {
		return json.Marshal(&src.AppleLoginChallenge)
	}

	if src.PlexAuthenticationChallenge != nil {
		return json.Marshal(&src.PlexAuthenticationChallenge)
	}

	if src.RedirectChallenge != nil {
		return json.Marshal(&src.RedirectChallenge)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LoginChallengeTypes) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AppleLoginChallenge != nil {
		return obj.AppleLoginChallenge
	}

	if obj.PlexAuthenticationChallenge != nil {
		return obj.PlexAuthenticationChallenge
	}

	if obj.RedirectChallenge != nil {
		return obj.RedirectChallenge
	}

	// all schemas are nil
	return nil
}

type NullableLoginChallengeTypes struct {
	value *LoginChallengeTypes
	isSet bool
}

func (v NullableLoginChallengeTypes) Get() *LoginChallengeTypes {
	return v.value
}

func (v *NullableLoginChallengeTypes) Set(val *LoginChallengeTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginChallengeTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginChallengeTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginChallengeTypes(val *LoginChallengeTypes) *NullableLoginChallengeTypes {
	return &NullableLoginChallengeTypes{value: val, isSet: true}
}

func (v NullableLoginChallengeTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginChallengeTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
