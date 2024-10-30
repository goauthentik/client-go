/*
authentik

Making authentication simple.

API version: 2024.10.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// EventMatcherPolicyRequest Event Matcher Policy Serializer
type EventMatcherPolicyRequest struct {
	Name string `json:"name"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	// Match created events with this action type. When left empty, all action types will be matched.
	Action NullableEventActions `json:"action,omitempty"`
	// Matches Event's Client IP (strict matching, for network matching use an Expression Policy)
	ClientIp NullableString `json:"client_ip,omitempty"`
	// Match events created by selected application. When left empty, all applications are matched.
	App NullableAppEnum `json:"app,omitempty"`
	// Match events created by selected model. When left empty, all models are matched. When an app is selected, all the application's models are matched.
	Model NullableModelEnum `json:"model,omitempty"`
}

// NewEventMatcherPolicyRequest instantiates a new EventMatcherPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventMatcherPolicyRequest(name string) *EventMatcherPolicyRequest {
	this := EventMatcherPolicyRequest{}
	this.Name = name
	return &this
}

// NewEventMatcherPolicyRequestWithDefaults instantiates a new EventMatcherPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventMatcherPolicyRequestWithDefaults() *EventMatcherPolicyRequest {
	this := EventMatcherPolicyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *EventMatcherPolicyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventMatcherPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventMatcherPolicyRequest) SetName(v string) {
	o.Name = v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *EventMatcherPolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMatcherPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *EventMatcherPolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *EventMatcherPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventMatcherPolicyRequest) GetAction() EventActions {
	if o == nil || o.Action.Get() == nil {
		var ret EventActions
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventMatcherPolicyRequest) GetActionOk() (*EventActions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *EventMatcherPolicyRequest) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableEventActions and assigns it to the Action field.
func (o *EventMatcherPolicyRequest) SetAction(v EventActions) {
	o.Action.Set(&v)
}

// SetActionNil sets the value for Action to be an explicit nil
func (o *EventMatcherPolicyRequest) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *EventMatcherPolicyRequest) UnsetAction() {
	o.Action.Unset()
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventMatcherPolicyRequest) GetClientIp() string {
	if o == nil || o.ClientIp.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientIp.Get()
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventMatcherPolicyRequest) GetClientIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientIp.Get(), o.ClientIp.IsSet()
}

// HasClientIp returns a boolean if a field has been set.
func (o *EventMatcherPolicyRequest) HasClientIp() bool {
	if o != nil && o.ClientIp.IsSet() {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given NullableString and assigns it to the ClientIp field.
func (o *EventMatcherPolicyRequest) SetClientIp(v string) {
	o.ClientIp.Set(&v)
}

// SetClientIpNil sets the value for ClientIp to be an explicit nil
func (o *EventMatcherPolicyRequest) SetClientIpNil() {
	o.ClientIp.Set(nil)
}

// UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
func (o *EventMatcherPolicyRequest) UnsetClientIp() {
	o.ClientIp.Unset()
}

// GetApp returns the App field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventMatcherPolicyRequest) GetApp() AppEnum {
	if o == nil || o.App.Get() == nil {
		var ret AppEnum
		return ret
	}
	return *o.App.Get()
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventMatcherPolicyRequest) GetAppOk() (*AppEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.App.Get(), o.App.IsSet()
}

// HasApp returns a boolean if a field has been set.
func (o *EventMatcherPolicyRequest) HasApp() bool {
	if o != nil && o.App.IsSet() {
		return true
	}

	return false
}

// SetApp gets a reference to the given NullableAppEnum and assigns it to the App field.
func (o *EventMatcherPolicyRequest) SetApp(v AppEnum) {
	o.App.Set(&v)
}

// SetAppNil sets the value for App to be an explicit nil
func (o *EventMatcherPolicyRequest) SetAppNil() {
	o.App.Set(nil)
}

// UnsetApp ensures that no value is present for App, not even an explicit nil
func (o *EventMatcherPolicyRequest) UnsetApp() {
	o.App.Unset()
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventMatcherPolicyRequest) GetModel() ModelEnum {
	if o == nil || o.Model.Get() == nil {
		var ret ModelEnum
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventMatcherPolicyRequest) GetModelOk() (*ModelEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *EventMatcherPolicyRequest) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableModelEnum and assigns it to the Model field.
func (o *EventMatcherPolicyRequest) SetModel(v ModelEnum) {
	o.Model.Set(&v)
}

// SetModelNil sets the value for Model to be an explicit nil
func (o *EventMatcherPolicyRequest) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *EventMatcherPolicyRequest) UnsetModel() {
	o.Model.Unset()
}

func (o EventMatcherPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.ClientIp.IsSet() {
		toSerialize["client_ip"] = o.ClientIp.Get()
	}
	if o.App.IsSet() {
		toSerialize["app"] = o.App.Get()
	}
	if o.Model.IsSet() {
		toSerialize["model"] = o.Model.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEventMatcherPolicyRequest struct {
	value *EventMatcherPolicyRequest
	isSet bool
}

func (v NullableEventMatcherPolicyRequest) Get() *EventMatcherPolicyRequest {
	return v.value
}

func (v *NullableEventMatcherPolicyRequest) Set(val *EventMatcherPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEventMatcherPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEventMatcherPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventMatcherPolicyRequest(val *EventMatcherPolicyRequest) *NullableEventMatcherPolicyRequest {
	return &NullableEventMatcherPolicyRequest{value: val, isSet: true}
}

func (v NullableEventMatcherPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventMatcherPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
