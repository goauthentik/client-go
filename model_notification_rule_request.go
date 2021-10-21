/*
authentik

Making authentication simple.

API version: 2021.10.1-rc2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// NotificationRuleRequest NotificationRule Serializer
type NotificationRuleRequest struct {
	Name string `json:"name"`
	// Select which transports should be used to notify the user. If none are selected, the notification will only be shown in the authentik UI.
	Transports []string `json:"transports"`
	// Controls which severity level the created notifications will have.
	Severity *SeverityEnum `json:"severity,omitempty"`
	// Define which group of users this notification should be sent and shown to. If left empty, Notification won't ben sent.
	Group NullableString `json:"group,omitempty"`
}

// NewNotificationRuleRequest instantiates a new NotificationRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationRuleRequest(name string, transports []string) *NotificationRuleRequest {
	this := NotificationRuleRequest{}
	this.Name = name
	this.Transports = transports
	return &this
}

// NewNotificationRuleRequestWithDefaults instantiates a new NotificationRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationRuleRequestWithDefaults() *NotificationRuleRequest {
	this := NotificationRuleRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NotificationRuleRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotificationRuleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NotificationRuleRequest) SetName(v string) {
	o.Name = v
}

// GetTransports returns the Transports field value
func (o *NotificationRuleRequest) GetTransports() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Transports
}

// GetTransportsOk returns a tuple with the Transports field value
// and a boolean to check if the value has been set.
func (o *NotificationRuleRequest) GetTransportsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transports, true
}

// SetTransports sets field value
func (o *NotificationRuleRequest) SetTransports(v []string) {
	o.Transports = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *NotificationRuleRequest) GetSeverity() SeverityEnum {
	if o == nil || o.Severity == nil {
		var ret SeverityEnum
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRuleRequest) GetSeverityOk() (*SeverityEnum, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *NotificationRuleRequest) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given SeverityEnum and assigns it to the Severity field.
func (o *NotificationRuleRequest) SetSeverity(v SeverityEnum) {
	o.Severity = &v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationRuleRequest) GetGroup() string {
	if o == nil || o.Group.Get() == nil {
		var ret string
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationRuleRequest) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *NotificationRuleRequest) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableString and assigns it to the Group field.
func (o *NotificationRuleRequest) SetGroup(v string) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil
func (o *NotificationRuleRequest) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *NotificationRuleRequest) UnsetGroup() {
	o.Group.Unset()
}

func (o NotificationRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["transports"] = o.Transports
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationRuleRequest struct {
	value *NotificationRuleRequest
	isSet bool
}

func (v NullableNotificationRuleRequest) Get() *NotificationRuleRequest {
	return v.value
}

func (v *NullableNotificationRuleRequest) Set(val *NotificationRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationRuleRequest(val *NotificationRuleRequest) *NullableNotificationRuleRequest {
	return &NullableNotificationRuleRequest{value: val, isSet: true}
}

func (v NullableNotificationRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
