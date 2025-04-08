/*
authentik

Making authentication simple.

API version: 2025.2.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// LicenseSummary Serializer for license status
type LicenseSummary struct {
	InternalUsers int32                    `json:"internal_users"`
	ExternalUsers int32                    `json:"external_users"`
	Status        LicenseSummaryStatusEnum `json:"status"`
	LatestValid   time.Time                `json:"latest_valid"`
	LicenseFlags  []LicenseFlagsEnum       `json:"license_flags"`
}

// NewLicenseSummary instantiates a new LicenseSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseSummary(internalUsers int32, externalUsers int32, status LicenseSummaryStatusEnum, latestValid time.Time, licenseFlags []LicenseFlagsEnum) *LicenseSummary {
	this := LicenseSummary{}
	this.InternalUsers = internalUsers
	this.ExternalUsers = externalUsers
	this.Status = status
	this.LatestValid = latestValid
	this.LicenseFlags = licenseFlags
	return &this
}

// NewLicenseSummaryWithDefaults instantiates a new LicenseSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseSummaryWithDefaults() *LicenseSummary {
	this := LicenseSummary{}
	return &this
}

// GetInternalUsers returns the InternalUsers field value
func (o *LicenseSummary) GetInternalUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InternalUsers
}

// GetInternalUsersOk returns a tuple with the InternalUsers field value
// and a boolean to check if the value has been set.
func (o *LicenseSummary) GetInternalUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InternalUsers, true
}

// SetInternalUsers sets field value
func (o *LicenseSummary) SetInternalUsers(v int32) {
	o.InternalUsers = v
}

// GetExternalUsers returns the ExternalUsers field value
func (o *LicenseSummary) GetExternalUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExternalUsers
}

// GetExternalUsersOk returns a tuple with the ExternalUsers field value
// and a boolean to check if the value has been set.
func (o *LicenseSummary) GetExternalUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUsers, true
}

// SetExternalUsers sets field value
func (o *LicenseSummary) SetExternalUsers(v int32) {
	o.ExternalUsers = v
}

// GetStatus returns the Status field value
func (o *LicenseSummary) GetStatus() LicenseSummaryStatusEnum {
	if o == nil {
		var ret LicenseSummaryStatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LicenseSummary) GetStatusOk() (*LicenseSummaryStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LicenseSummary) SetStatus(v LicenseSummaryStatusEnum) {
	o.Status = v
}

// GetLatestValid returns the LatestValid field value
func (o *LicenseSummary) GetLatestValid() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LatestValid
}

// GetLatestValidOk returns a tuple with the LatestValid field value
// and a boolean to check if the value has been set.
func (o *LicenseSummary) GetLatestValidOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LatestValid, true
}

// SetLatestValid sets field value
func (o *LicenseSummary) SetLatestValid(v time.Time) {
	o.LatestValid = v
}

// GetLicenseFlags returns the LicenseFlags field value
func (o *LicenseSummary) GetLicenseFlags() []LicenseFlagsEnum {
	if o == nil {
		var ret []LicenseFlagsEnum
		return ret
	}

	return o.LicenseFlags
}

// GetLicenseFlagsOk returns a tuple with the LicenseFlags field value
// and a boolean to check if the value has been set.
func (o *LicenseSummary) GetLicenseFlagsOk() ([]LicenseFlagsEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.LicenseFlags, true
}

// SetLicenseFlags sets field value
func (o *LicenseSummary) SetLicenseFlags(v []LicenseFlagsEnum) {
	o.LicenseFlags = v
}

func (o LicenseSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["internal_users"] = o.InternalUsers
	}
	if true {
		toSerialize["external_users"] = o.ExternalUsers
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["latest_valid"] = o.LatestValid
	}
	if true {
		toSerialize["license_flags"] = o.LicenseFlags
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseSummary struct {
	value *LicenseSummary
	isSet bool
}

func (v NullableLicenseSummary) Get() *LicenseSummary {
	return v.value
}

func (v *NullableLicenseSummary) Set(val *LicenseSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseSummary(val *LicenseSummary) *NullableLicenseSummary {
	return &NullableLicenseSummary{value: val, isSet: true}
}

func (v NullableLicenseSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
