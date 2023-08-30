/*
authentik

Making authentication simple.

API version: 2023.8.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// OutpostHealth Outpost health status
type OutpostHealth struct {
	Uid             string    `json:"uid"`
	LastSeen        time.Time `json:"last_seen"`
	Version         string    `json:"version"`
	VersionShould   string    `json:"version_should"`
	VersionOutdated bool      `json:"version_outdated"`
	BuildHash       string    `json:"build_hash"`
	BuildHashShould string    `json:"build_hash_should"`
	Hostname        string    `json:"hostname"`
}

// NewOutpostHealth instantiates a new OutpostHealth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutpostHealth(uid string, lastSeen time.Time, version string, versionShould string, versionOutdated bool, buildHash string, buildHashShould string, hostname string) *OutpostHealth {
	this := OutpostHealth{}
	this.Uid = uid
	this.LastSeen = lastSeen
	this.Version = version
	this.VersionShould = versionShould
	this.VersionOutdated = versionOutdated
	this.BuildHash = buildHash
	this.BuildHashShould = buildHashShould
	this.Hostname = hostname
	return &this
}

// NewOutpostHealthWithDefaults instantiates a new OutpostHealth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutpostHealthWithDefaults() *OutpostHealth {
	this := OutpostHealth{}
	return &this
}

// GetUid returns the Uid field value
func (o *OutpostHealth) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *OutpostHealth) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *OutpostHealth) SetUid(v string) {
	o.Uid = v
}

// GetLastSeen returns the LastSeen field value
func (o *OutpostHealth) GetLastSeen() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value
// and a boolean to check if the value has been set.
func (o *OutpostHealth) GetLastSeenOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSeen, true
}

// SetLastSeen sets field value
func (o *OutpostHealth) SetLastSeen(v time.Time) {
	o.LastSeen = v
}

// GetVersion returns the Version field value
func (o *OutpostHealth) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *OutpostHealth) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *OutpostHealth) SetVersion(v string) {
	o.Version = v
}

// GetVersionShould returns the VersionShould field value
func (o *OutpostHealth) GetVersionShould() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionShould
}

// GetVersionShouldOk returns a tuple with the VersionShould field value
// and a boolean to check if the value has been set.
func (o *OutpostHealth) GetVersionShouldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionShould, true
}

// SetVersionShould sets field value
func (o *OutpostHealth) SetVersionShould(v string) {
	o.VersionShould = v
}

// GetVersionOutdated returns the VersionOutdated field value
func (o *OutpostHealth) GetVersionOutdated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VersionOutdated
}

// GetVersionOutdatedOk returns a tuple with the VersionOutdated field value
// and a boolean to check if the value has been set.
func (o *OutpostHealth) GetVersionOutdatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionOutdated, true
}

// SetVersionOutdated sets field value
func (o *OutpostHealth) SetVersionOutdated(v bool) {
	o.VersionOutdated = v
}

// GetBuildHash returns the BuildHash field value
func (o *OutpostHealth) GetBuildHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildHash
}

// GetBuildHashOk returns a tuple with the BuildHash field value
// and a boolean to check if the value has been set.
func (o *OutpostHealth) GetBuildHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildHash, true
}

// SetBuildHash sets field value
func (o *OutpostHealth) SetBuildHash(v string) {
	o.BuildHash = v
}

// GetBuildHashShould returns the BuildHashShould field value
func (o *OutpostHealth) GetBuildHashShould() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildHashShould
}

// GetBuildHashShouldOk returns a tuple with the BuildHashShould field value
// and a boolean to check if the value has been set.
func (o *OutpostHealth) GetBuildHashShouldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildHashShould, true
}

// SetBuildHashShould sets field value
func (o *OutpostHealth) SetBuildHashShould(v string) {
	o.BuildHashShould = v
}

// GetHostname returns the Hostname field value
func (o *OutpostHealth) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *OutpostHealth) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *OutpostHealth) SetHostname(v string) {
	o.Hostname = v
}

func (o OutpostHealth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uid"] = o.Uid
	}
	if true {
		toSerialize["last_seen"] = o.LastSeen
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["version_should"] = o.VersionShould
	}
	if true {
		toSerialize["version_outdated"] = o.VersionOutdated
	}
	if true {
		toSerialize["build_hash"] = o.BuildHash
	}
	if true {
		toSerialize["build_hash_should"] = o.BuildHashShould
	}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	return json.Marshal(toSerialize)
}

type NullableOutpostHealth struct {
	value *OutpostHealth
	isSet bool
}

func (v NullableOutpostHealth) Get() *OutpostHealth {
	return v.value
}

func (v *NullableOutpostHealth) Set(val *OutpostHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableOutpostHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableOutpostHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutpostHealth(val *OutpostHealth) *NullableOutpostHealth {
	return &NullableOutpostHealth{value: val, isSet: true}
}

func (v NullableOutpostHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutpostHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
