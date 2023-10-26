/*
authentik

Making authentication simple.

API version: 2023.10.1
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
	InternalUsers    int32     `json:"internal_users"`
	ExternalUsers    int32     `json:"external_users"`
	Valid            bool      `json:"valid"`
	ShowAdminWarning bool      `json:"show_admin_warning"`
	ShowUserWarning  bool      `json:"show_user_warning"`
	ReadOnly         bool      `json:"read_only"`
	LatestValid      time.Time `json:"latest_valid"`
	HasLicense       bool      `json:"has_license"`
}

// NewLicenseSummary instantiates a new LicenseSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseSummary(internalUsers int32, externalUsers int32, valid bool, showAdminWarning bool, showUserWarning bool, readOnly bool, latestValid time.Time, hasLicense bool) *LicenseSummary {
	this := LicenseSummary{}
	this.InternalUsers = internalUsers
	this.ExternalUsers = externalUsers
	this.Valid = valid
	this.ShowAdminWarning = showAdminWarning
	this.ShowUserWarning = showUserWarning
	this.ReadOnly = readOnly
	this.LatestValid = latestValid
	this.HasLicense = hasLicense
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

// GetValid returns the Valid field value
func (o *LicenseSummary) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *LicenseSummary) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *LicenseSummary) SetValid(v bool) {
	o.Valid = v
}

// GetShowAdminWarning returns the ShowAdminWarning field value
func (o *LicenseSummary) GetShowAdminWarning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowAdminWarning
}

// GetShowAdminWarningOk returns a tuple with the ShowAdminWarning field value
// and a boolean to check if the value has been set.
func (o *LicenseSummary) GetShowAdminWarningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowAdminWarning, true
}

// SetShowAdminWarning sets field value
func (o *LicenseSummary) SetShowAdminWarning(v bool) {
	o.ShowAdminWarning = v
}

// GetShowUserWarning returns the ShowUserWarning field value
func (o *LicenseSummary) GetShowUserWarning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowUserWarning
}

// GetShowUserWarningOk returns a tuple with the ShowUserWarning field value
// and a boolean to check if the value has been set.
func (o *LicenseSummary) GetShowUserWarningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowUserWarning, true
}

// SetShowUserWarning sets field value
func (o *LicenseSummary) SetShowUserWarning(v bool) {
	o.ShowUserWarning = v
}

// GetReadOnly returns the ReadOnly field value
func (o *LicenseSummary) GetReadOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value
// and a boolean to check if the value has been set.
func (o *LicenseSummary) GetReadOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadOnly, true
}

// SetReadOnly sets field value
func (o *LicenseSummary) SetReadOnly(v bool) {
	o.ReadOnly = v
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

// GetHasLicense returns the HasLicense field value
func (o *LicenseSummary) GetHasLicense() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasLicense
}

// GetHasLicenseOk returns a tuple with the HasLicense field value
// and a boolean to check if the value has been set.
func (o *LicenseSummary) GetHasLicenseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasLicense, true
}

// SetHasLicense sets field value
func (o *LicenseSummary) SetHasLicense(v bool) {
	o.HasLicense = v
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
		toSerialize["valid"] = o.Valid
	}
	if true {
		toSerialize["show_admin_warning"] = o.ShowAdminWarning
	}
	if true {
		toSerialize["show_user_warning"] = o.ShowUserWarning
	}
	if true {
		toSerialize["read_only"] = o.ReadOnly
	}
	if true {
		toSerialize["latest_valid"] = o.LatestValid
	}
	if true {
		toSerialize["has_license"] = o.HasLicense
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
