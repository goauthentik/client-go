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
	"time"
)

// Device Serializer for Duo authenticator devices
type Device struct {
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName string `json:"meta_model_name"`
	Pk            string `json:"pk"`
	Name          string `json:"name"`
	// Get type of device
	Type        string       `json:"type"`
	Confirmed   bool         `json:"confirmed"`
	Created     time.Time    `json:"created"`
	LastUpdated time.Time    `json:"last_updated"`
	LastUsed    NullableTime `json:"last_used"`
	// Get extra description
	ExtraDescription string `json:"extra_description"`
}

// NewDevice instantiates a new Device object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevice(verboseName string, verboseNamePlural string, metaModelName string, pk string, name string, type_ string, confirmed bool, created time.Time, lastUpdated time.Time, lastUsed NullableTime, extraDescription string) *Device {
	this := Device{}
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.Pk = pk
	this.Name = name
	this.Type = type_
	this.Confirmed = confirmed
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUsed = lastUsed
	this.ExtraDescription = extraDescription
	return &this
}

// NewDeviceWithDefaults instantiates a new Device object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceWithDefaults() *Device {
	this := Device{}
	return &this
}

// GetVerboseName returns the VerboseName field value
func (o *Device) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *Device) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *Device) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *Device) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *Device) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *Device) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *Device) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *Device) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *Device) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetPk returns the Pk field value
func (o *Device) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *Device) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *Device) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *Device) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Device) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Device) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *Device) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Device) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Device) SetType(v string) {
	o.Type = v
}

// GetConfirmed returns the Confirmed field value
func (o *Device) GetConfirmed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value
// and a boolean to check if the value has been set.
func (o *Device) GetConfirmedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confirmed, true
}

// SetConfirmed sets field value
func (o *Device) SetConfirmed(v bool) {
	o.Confirmed = v
}

// GetCreated returns the Created field value
func (o *Device) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Device) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Device) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *Device) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *Device) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *Device) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUsed returns the LastUsed field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Device) GetLastUsed() time.Time {
	if o == nil || o.LastUsed.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUsed.Get()
}

// GetLastUsedOk returns a tuple with the LastUsed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetLastUsedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUsed.Get(), o.LastUsed.IsSet()
}

// SetLastUsed sets field value
func (o *Device) SetLastUsed(v time.Time) {
	o.LastUsed.Set(&v)
}

// GetExtraDescription returns the ExtraDescription field value
func (o *Device) GetExtraDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtraDescription
}

// GetExtraDescriptionOk returns a tuple with the ExtraDescription field value
// and a boolean to check if the value has been set.
func (o *Device) GetExtraDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtraDescription, true
}

// SetExtraDescription sets field value
func (o *Device) SetExtraDescription(v string) {
	o.ExtraDescription = v
}

func (o Device) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["confirmed"] = o.Confirmed
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if true {
		toSerialize["last_used"] = o.LastUsed.Get()
	}
	if true {
		toSerialize["extra_description"] = o.ExtraDescription
	}
	return json.Marshal(toSerialize)
}

type NullableDevice struct {
	value *Device
	isSet bool
}

func (v NullableDevice) Get() *Device {
	return v.value
}

func (v *NullableDevice) Set(val *Device) {
	v.value = val
	v.isSet = true
}

func (v NullableDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevice(val *Device) *NullableDevice {
	return &NullableDevice{value: val, isSet: true}
}

func (v NullableDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
