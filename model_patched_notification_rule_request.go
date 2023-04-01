/*
authentik

Making authentication simple.

API version: 2023.3.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PatchedNotificationRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedNotificationRuleRequest{}

// PatchedNotificationRuleRequest NotificationRule Serializer
type PatchedNotificationRuleRequest struct {
	Name *string `json:"name,omitempty"`
	// Select which transports should be used to notify the user. If none are selected, the notification will only be shown in the authentik UI.
	Transports []string      `json:"transports,omitempty"`
	Severity   *SeverityEnum `json:"severity,omitempty"`
	// Define which group of users this notification should be sent and shown to. If left empty, Notification won't ben sent.
	Group NullableString `json:"group,omitempty"`
}

// NewPatchedNotificationRuleRequest instantiates a new PatchedNotificationRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedNotificationRuleRequest() *PatchedNotificationRuleRequest {
	this := PatchedNotificationRuleRequest{}
	return &this
}

// NewPatchedNotificationRuleRequestWithDefaults instantiates a new PatchedNotificationRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedNotificationRuleRequestWithDefaults() *PatchedNotificationRuleRequest {
	this := PatchedNotificationRuleRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedNotificationRuleRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationRuleRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedNotificationRuleRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedNotificationRuleRequest) SetName(v string) {
	o.Name = &v
}

// GetTransports returns the Transports field value if set, zero value otherwise.
func (o *PatchedNotificationRuleRequest) GetTransports() []string {
	if o == nil || IsNil(o.Transports) {
		var ret []string
		return ret
	}
	return o.Transports
}

// GetTransportsOk returns a tuple with the Transports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationRuleRequest) GetTransportsOk() ([]string, bool) {
	if o == nil || IsNil(o.Transports) {
		return nil, false
	}
	return o.Transports, true
}

// HasTransports returns a boolean if a field has been set.
func (o *PatchedNotificationRuleRequest) HasTransports() bool {
	if o != nil && !IsNil(o.Transports) {
		return true
	}

	return false
}

// SetTransports gets a reference to the given []string and assigns it to the Transports field.
func (o *PatchedNotificationRuleRequest) SetTransports(v []string) {
	o.Transports = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *PatchedNotificationRuleRequest) GetSeverity() SeverityEnum {
	if o == nil || IsNil(o.Severity) {
		var ret SeverityEnum
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationRuleRequest) GetSeverityOk() (*SeverityEnum, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *PatchedNotificationRuleRequest) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given SeverityEnum and assigns it to the Severity field.
func (o *PatchedNotificationRuleRequest) SetSeverity(v SeverityEnum) {
	o.Severity = &v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedNotificationRuleRequest) GetGroup() string {
	if o == nil || IsNil(o.Group.Get()) {
		var ret string
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedNotificationRuleRequest) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *PatchedNotificationRuleRequest) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableString and assigns it to the Group field.
func (o *PatchedNotificationRuleRequest) SetGroup(v string) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil
func (o *PatchedNotificationRuleRequest) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *PatchedNotificationRuleRequest) UnsetGroup() {
	o.Group.Unset()
}

func (o PatchedNotificationRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedNotificationRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Transports) {
		toSerialize["transports"] = o.Transports
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	return toSerialize, nil
}

type NullablePatchedNotificationRuleRequest struct {
	value *PatchedNotificationRuleRequest
	isSet bool
}

func (v NullablePatchedNotificationRuleRequest) Get() *PatchedNotificationRuleRequest {
	return v.value
}

func (v *NullablePatchedNotificationRuleRequest) Set(val *PatchedNotificationRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedNotificationRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedNotificationRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedNotificationRuleRequest(val *PatchedNotificationRuleRequest) *NullablePatchedNotificationRuleRequest {
	return &NullablePatchedNotificationRuleRequest{value: val, isSet: true}
}

func (v NullablePatchedNotificationRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedNotificationRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
