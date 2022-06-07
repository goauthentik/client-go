/*
authentik

Making authentication simple.

API version: 2022.6.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedNotificationRuleRequest NotificationRule Serializer
type PatchedNotificationRuleRequest struct {
	Name *string `json:"name,omitempty"`
	// Select which transports should be used to notify the user. If none are selected, the notification will only be shown in the authentik UI.
	Transports []string `json:"transports,omitempty"`
	// Controls which severity level the created notifications will have.
	Severity NullableSeverityEnum `json:"severity,omitempty"`
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
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationRuleRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedNotificationRuleRequest) HasName() bool {
	if o != nil && o.Name != nil {
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
	if o == nil || o.Transports == nil {
		var ret []string
		return ret
	}
	return o.Transports
}

// GetTransportsOk returns a tuple with the Transports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationRuleRequest) GetTransportsOk() ([]string, bool) {
	if o == nil || o.Transports == nil {
		return nil, false
	}
	return o.Transports, true
}

// HasTransports returns a boolean if a field has been set.
func (o *PatchedNotificationRuleRequest) HasTransports() bool {
	if o != nil && o.Transports != nil {
		return true
	}

	return false
}

// SetTransports gets a reference to the given []string and assigns it to the Transports field.
func (o *PatchedNotificationRuleRequest) SetTransports(v []string) {
	o.Transports = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedNotificationRuleRequest) GetSeverity() SeverityEnum {
	if o == nil || o.Severity.Get() == nil {
		var ret SeverityEnum
		return ret
	}
	return *o.Severity.Get()
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedNotificationRuleRequest) GetSeverityOk() (*SeverityEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Severity.Get(), o.Severity.IsSet()
}

// HasSeverity returns a boolean if a field has been set.
func (o *PatchedNotificationRuleRequest) HasSeverity() bool {
	if o != nil && o.Severity.IsSet() {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given NullableSeverityEnum and assigns it to the Severity field.
func (o *PatchedNotificationRuleRequest) SetSeverity(v SeverityEnum) {
	o.Severity.Set(&v)
}

// SetSeverityNil sets the value for Severity to be an explicit nil
func (o *PatchedNotificationRuleRequest) SetSeverityNil() {
	o.Severity.Set(nil)
}

// UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
func (o *PatchedNotificationRuleRequest) UnsetSeverity() {
	o.Severity.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedNotificationRuleRequest) GetGroup() string {
	if o == nil || o.Group.Get() == nil {
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
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Transports != nil {
		toSerialize["transports"] = o.Transports
	}
	if o.Severity.IsSet() {
		toSerialize["severity"] = o.Severity.Get()
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	return json.Marshal(toSerialize)
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
