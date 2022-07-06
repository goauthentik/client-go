/*
authentik

Making authentication simple.

API version: 2022.7.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// NotificationRule NotificationRule Serializer
type NotificationRule struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// Select which transports should be used to notify the user. If none are selected, the notification will only be shown in the authentik UI.
	Transports []string `json:"transports,omitempty"`
	// Controls which severity level the created notifications will have.
	Severity NullableSeverityEnum `json:"severity,omitempty"`
	// Define which group of users this notification should be sent and shown to. If left empty, Notification won't ben sent.
	Group    NullableString           `json:"group,omitempty"`
	GroupObj NotificationRuleGroupObj `json:"group_obj"`
}

// NewNotificationRule instantiates a new NotificationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationRule(pk string, name string, groupObj NotificationRuleGroupObj) *NotificationRule {
	this := NotificationRule{}
	this.Pk = pk
	this.Name = name
	this.GroupObj = groupObj
	return &this
}

// NewNotificationRuleWithDefaults instantiates a new NotificationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationRuleWithDefaults() *NotificationRule {
	this := NotificationRule{}
	return &this
}

// GetPk returns the Pk field value
func (o *NotificationRule) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *NotificationRule) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *NotificationRule) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *NotificationRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotificationRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NotificationRule) SetName(v string) {
	o.Name = v
}

// GetTransports returns the Transports field value if set, zero value otherwise.
func (o *NotificationRule) GetTransports() []string {
	if o == nil || o.Transports == nil {
		var ret []string
		return ret
	}
	return o.Transports
}

// GetTransportsOk returns a tuple with the Transports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRule) GetTransportsOk() ([]string, bool) {
	if o == nil || o.Transports == nil {
		return nil, false
	}
	return o.Transports, true
}

// HasTransports returns a boolean if a field has been set.
func (o *NotificationRule) HasTransports() bool {
	if o != nil && o.Transports != nil {
		return true
	}

	return false
}

// SetTransports gets a reference to the given []string and assigns it to the Transports field.
func (o *NotificationRule) SetTransports(v []string) {
	o.Transports = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationRule) GetSeverity() SeverityEnum {
	if o == nil || o.Severity.Get() == nil {
		var ret SeverityEnum
		return ret
	}
	return *o.Severity.Get()
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationRule) GetSeverityOk() (*SeverityEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Severity.Get(), o.Severity.IsSet()
}

// HasSeverity returns a boolean if a field has been set.
func (o *NotificationRule) HasSeverity() bool {
	if o != nil && o.Severity.IsSet() {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given NullableSeverityEnum and assigns it to the Severity field.
func (o *NotificationRule) SetSeverity(v SeverityEnum) {
	o.Severity.Set(&v)
}

// SetSeverityNil sets the value for Severity to be an explicit nil
func (o *NotificationRule) SetSeverityNil() {
	o.Severity.Set(nil)
}

// UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
func (o *NotificationRule) UnsetSeverity() {
	o.Severity.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationRule) GetGroup() string {
	if o == nil || o.Group.Get() == nil {
		var ret string
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationRule) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *NotificationRule) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableString and assigns it to the Group field.
func (o *NotificationRule) SetGroup(v string) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil
func (o *NotificationRule) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *NotificationRule) UnsetGroup() {
	o.Group.Unset()
}

// GetGroupObj returns the GroupObj field value
func (o *NotificationRule) GetGroupObj() NotificationRuleGroupObj {
	if o == nil {
		var ret NotificationRuleGroupObj
		return ret
	}

	return o.GroupObj
}

// GetGroupObjOk returns a tuple with the GroupObj field value
// and a boolean to check if the value has been set.
func (o *NotificationRule) GetGroupObjOk() (*NotificationRuleGroupObj, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupObj, true
}

// SetGroupObj sets field value
func (o *NotificationRule) SetGroupObj(v NotificationRuleGroupObj) {
	o.GroupObj = v
}

func (o NotificationRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
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
	if true {
		toSerialize["group_obj"] = o.GroupObj
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationRule struct {
	value *NotificationRule
	isSet bool
}

func (v NullableNotificationRule) Get() *NotificationRule {
	return v.value
}

func (v *NullableNotificationRule) Set(val *NotificationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationRule(val *NotificationRule) *NullableNotificationRule {
	return &NullableNotificationRule{value: val, isSet: true}
}

func (v NullableNotificationRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
