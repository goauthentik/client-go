/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SyncObjectRequest Sync object serializer
type SyncObjectRequest struct {
	SyncObjectModel SyncObjectModelEnum `json:"sync_object_model"`
	SyncObjectId    string              `json:"sync_object_id"`
	OverrideDryRun  *bool               `json:"override_dry_run,omitempty"`
}

// NewSyncObjectRequest instantiates a new SyncObjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncObjectRequest(syncObjectModel SyncObjectModelEnum, syncObjectId string) *SyncObjectRequest {
	this := SyncObjectRequest{}
	this.SyncObjectModel = syncObjectModel
	this.SyncObjectId = syncObjectId
	var overrideDryRun bool = false
	this.OverrideDryRun = &overrideDryRun
	return &this
}

// NewSyncObjectRequestWithDefaults instantiates a new SyncObjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncObjectRequestWithDefaults() *SyncObjectRequest {
	this := SyncObjectRequest{}
	var overrideDryRun bool = false
	this.OverrideDryRun = &overrideDryRun
	return &this
}

// GetSyncObjectModel returns the SyncObjectModel field value
func (o *SyncObjectRequest) GetSyncObjectModel() SyncObjectModelEnum {
	if o == nil {
		var ret SyncObjectModelEnum
		return ret
	}

	return o.SyncObjectModel
}

// GetSyncObjectModelOk returns a tuple with the SyncObjectModel field value
// and a boolean to check if the value has been set.
func (o *SyncObjectRequest) GetSyncObjectModelOk() (*SyncObjectModelEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyncObjectModel, true
}

// SetSyncObjectModel sets field value
func (o *SyncObjectRequest) SetSyncObjectModel(v SyncObjectModelEnum) {
	o.SyncObjectModel = v
}

// GetSyncObjectId returns the SyncObjectId field value
func (o *SyncObjectRequest) GetSyncObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SyncObjectId
}

// GetSyncObjectIdOk returns a tuple with the SyncObjectId field value
// and a boolean to check if the value has been set.
func (o *SyncObjectRequest) GetSyncObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyncObjectId, true
}

// SetSyncObjectId sets field value
func (o *SyncObjectRequest) SetSyncObjectId(v string) {
	o.SyncObjectId = v
}

// GetOverrideDryRun returns the OverrideDryRun field value if set, zero value otherwise.
func (o *SyncObjectRequest) GetOverrideDryRun() bool {
	if o == nil || o.OverrideDryRun == nil {
		var ret bool
		return ret
	}
	return *o.OverrideDryRun
}

// GetOverrideDryRunOk returns a tuple with the OverrideDryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncObjectRequest) GetOverrideDryRunOk() (*bool, bool) {
	if o == nil || o.OverrideDryRun == nil {
		return nil, false
	}
	return o.OverrideDryRun, true
}

// HasOverrideDryRun returns a boolean if a field has been set.
func (o *SyncObjectRequest) HasOverrideDryRun() bool {
	if o != nil && o.OverrideDryRun != nil {
		return true
	}

	return false
}

// SetOverrideDryRun gets a reference to the given bool and assigns it to the OverrideDryRun field.
func (o *SyncObjectRequest) SetOverrideDryRun(v bool) {
	o.OverrideDryRun = &v
}

func (o SyncObjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sync_object_model"] = o.SyncObjectModel
	}
	if true {
		toSerialize["sync_object_id"] = o.SyncObjectId
	}
	if o.OverrideDryRun != nil {
		toSerialize["override_dry_run"] = o.OverrideDryRun
	}
	return json.Marshal(toSerialize)
}

type NullableSyncObjectRequest struct {
	value *SyncObjectRequest
	isSet bool
}

func (v NullableSyncObjectRequest) Get() *SyncObjectRequest {
	return v.value
}

func (v *NullableSyncObjectRequest) Set(val *SyncObjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncObjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncObjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncObjectRequest(val *SyncObjectRequest) *NullableSyncObjectRequest {
	return &NullableSyncObjectRequest{value: val, isSet: true}
}

func (v NullableSyncObjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncObjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
