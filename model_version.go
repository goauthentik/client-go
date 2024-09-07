/*
authentik

Making authentication simple.

API version: 2024.8.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Version Get running and latest version.
type Version struct {
	// Get current version
	VersionCurrent string `json:"version_current"`
	// Get latest version from cache
	VersionLatest string `json:"version_latest"`
	// Check if latest version is valid
	VersionLatestValid bool `json:"version_latest_valid"`
	// Get build hash, if version is not latest or released
	BuildHash string `json:"build_hash"`
	// Check if we're running the latest version
	Outdated bool `json:"outdated"`
	// Check if any outpost is outdated/has a version mismatch
	OutpostOutdated bool `json:"outpost_outdated"`
}

// NewVersion instantiates a new Version object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersion(versionCurrent string, versionLatest string, versionLatestValid bool, buildHash string, outdated bool, outpostOutdated bool) *Version {
	this := Version{}
	this.VersionCurrent = versionCurrent
	this.VersionLatest = versionLatest
	this.VersionLatestValid = versionLatestValid
	this.BuildHash = buildHash
	this.Outdated = outdated
	this.OutpostOutdated = outpostOutdated
	return &this
}

// NewVersionWithDefaults instantiates a new Version object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionWithDefaults() *Version {
	this := Version{}
	return &this
}

// GetVersionCurrent returns the VersionCurrent field value
func (o *Version) GetVersionCurrent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionCurrent
}

// GetVersionCurrentOk returns a tuple with the VersionCurrent field value
// and a boolean to check if the value has been set.
func (o *Version) GetVersionCurrentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionCurrent, true
}

// SetVersionCurrent sets field value
func (o *Version) SetVersionCurrent(v string) {
	o.VersionCurrent = v
}

// GetVersionLatest returns the VersionLatest field value
func (o *Version) GetVersionLatest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionLatest
}

// GetVersionLatestOk returns a tuple with the VersionLatest field value
// and a boolean to check if the value has been set.
func (o *Version) GetVersionLatestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionLatest, true
}

// SetVersionLatest sets field value
func (o *Version) SetVersionLatest(v string) {
	o.VersionLatest = v
}

// GetVersionLatestValid returns the VersionLatestValid field value
func (o *Version) GetVersionLatestValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VersionLatestValid
}

// GetVersionLatestValidOk returns a tuple with the VersionLatestValid field value
// and a boolean to check if the value has been set.
func (o *Version) GetVersionLatestValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionLatestValid, true
}

// SetVersionLatestValid sets field value
func (o *Version) SetVersionLatestValid(v bool) {
	o.VersionLatestValid = v
}

// GetBuildHash returns the BuildHash field value
func (o *Version) GetBuildHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildHash
}

// GetBuildHashOk returns a tuple with the BuildHash field value
// and a boolean to check if the value has been set.
func (o *Version) GetBuildHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildHash, true
}

// SetBuildHash sets field value
func (o *Version) SetBuildHash(v string) {
	o.BuildHash = v
}

// GetOutdated returns the Outdated field value
func (o *Version) GetOutdated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Outdated
}

// GetOutdatedOk returns a tuple with the Outdated field value
// and a boolean to check if the value has been set.
func (o *Version) GetOutdatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outdated, true
}

// SetOutdated sets field value
func (o *Version) SetOutdated(v bool) {
	o.Outdated = v
}

// GetOutpostOutdated returns the OutpostOutdated field value
func (o *Version) GetOutpostOutdated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OutpostOutdated
}

// GetOutpostOutdatedOk returns a tuple with the OutpostOutdated field value
// and a boolean to check if the value has been set.
func (o *Version) GetOutpostOutdatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutpostOutdated, true
}

// SetOutpostOutdated sets field value
func (o *Version) SetOutpostOutdated(v bool) {
	o.OutpostOutdated = v
}

func (o Version) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["version_current"] = o.VersionCurrent
	}
	if true {
		toSerialize["version_latest"] = o.VersionLatest
	}
	if true {
		toSerialize["version_latest_valid"] = o.VersionLatestValid
	}
	if true {
		toSerialize["build_hash"] = o.BuildHash
	}
	if true {
		toSerialize["outdated"] = o.Outdated
	}
	if true {
		toSerialize["outpost_outdated"] = o.OutpostOutdated
	}
	return json.Marshal(toSerialize)
}

type NullableVersion struct {
	value *Version
	isSet bool
}

func (v NullableVersion) Get() *Version {
	return v.value
}

func (v *NullableVersion) Set(val *Version) {
	v.value = val
	v.isSet = true
}

func (v NullableVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersion(val *Version) *NullableVersion {
	return &NullableVersion{value: val, isSet: true}
}

func (v NullableVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
