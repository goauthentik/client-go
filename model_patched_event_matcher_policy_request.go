/*
authentik

Making authentication simple.

API version: 2022.9.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedEventMatcherPolicyRequest Event Matcher Policy Serializer
type PatchedEventMatcherPolicyRequest struct {
	Name NullableString `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	// Match created events with this action type. When left empty, all action types will be matched.
	Action NullableEventActions `json:"action,omitempty"`
	// Matches Event's Client IP (strict matching, for network matching use an Expression Policy)
	ClientIp *string `json:"client_ip,omitempty"`
	// Match events created by selected application. When left empty, all applications are matched.
	App NullableAppEnum `json:"app,omitempty"`
}

// NewPatchedEventMatcherPolicyRequest instantiates a new PatchedEventMatcherPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedEventMatcherPolicyRequest() *PatchedEventMatcherPolicyRequest {
	this := PatchedEventMatcherPolicyRequest{}
	return &this
}

// NewPatchedEventMatcherPolicyRequestWithDefaults instantiates a new PatchedEventMatcherPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedEventMatcherPolicyRequestWithDefaults() *PatchedEventMatcherPolicyRequest {
	this := PatchedEventMatcherPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEventMatcherPolicyRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEventMatcherPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PatchedEventMatcherPolicyRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PatchedEventMatcherPolicyRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PatchedEventMatcherPolicyRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PatchedEventMatcherPolicyRequest) UnsetName() {
	o.Name.Unset()
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *PatchedEventMatcherPolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEventMatcherPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *PatchedEventMatcherPolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *PatchedEventMatcherPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEventMatcherPolicyRequest) GetAction() EventActions {
	if o == nil || o.Action.Get() == nil {
		var ret EventActions
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEventMatcherPolicyRequest) GetActionOk() (*EventActions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *PatchedEventMatcherPolicyRequest) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableEventActions and assigns it to the Action field.
func (o *PatchedEventMatcherPolicyRequest) SetAction(v EventActions) {
	o.Action.Set(&v)
}

// SetActionNil sets the value for Action to be an explicit nil
func (o *PatchedEventMatcherPolicyRequest) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *PatchedEventMatcherPolicyRequest) UnsetAction() {
	o.Action.Unset()
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *PatchedEventMatcherPolicyRequest) GetClientIp() string {
	if o == nil || o.ClientIp == nil {
		var ret string
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEventMatcherPolicyRequest) GetClientIpOk() (*string, bool) {
	if o == nil || o.ClientIp == nil {
		return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *PatchedEventMatcherPolicyRequest) HasClientIp() bool {
	if o != nil && o.ClientIp != nil {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given string and assigns it to the ClientIp field.
func (o *PatchedEventMatcherPolicyRequest) SetClientIp(v string) {
	o.ClientIp = &v
}

// GetApp returns the App field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEventMatcherPolicyRequest) GetApp() AppEnum {
	if o == nil || o.App.Get() == nil {
		var ret AppEnum
		return ret
	}
	return *o.App.Get()
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEventMatcherPolicyRequest) GetAppOk() (*AppEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.App.Get(), o.App.IsSet()
}

// HasApp returns a boolean if a field has been set.
func (o *PatchedEventMatcherPolicyRequest) HasApp() bool {
	if o != nil && o.App.IsSet() {
		return true
	}

	return false
}

// SetApp gets a reference to the given NullableAppEnum and assigns it to the App field.
func (o *PatchedEventMatcherPolicyRequest) SetApp(v AppEnum) {
	o.App.Set(&v)
}

// SetAppNil sets the value for App to be an explicit nil
func (o *PatchedEventMatcherPolicyRequest) SetAppNil() {
	o.App.Set(nil)
}

// UnsetApp ensures that no value is present for App, not even an explicit nil
func (o *PatchedEventMatcherPolicyRequest) UnsetApp() {
	o.App.Unset()
}

func (o PatchedEventMatcherPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.ClientIp != nil {
		toSerialize["client_ip"] = o.ClientIp
	}
	if o.App.IsSet() {
		toSerialize["app"] = o.App.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedEventMatcherPolicyRequest struct {
	value *PatchedEventMatcherPolicyRequest
	isSet bool
}

func (v NullablePatchedEventMatcherPolicyRequest) Get() *PatchedEventMatcherPolicyRequest {
	return v.value
}

func (v *NullablePatchedEventMatcherPolicyRequest) Set(val *PatchedEventMatcherPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedEventMatcherPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedEventMatcherPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedEventMatcherPolicyRequest(val *PatchedEventMatcherPolicyRequest) *NullablePatchedEventMatcherPolicyRequest {
	return &NullablePatchedEventMatcherPolicyRequest{value: val, isSet: true}
}

func (v NullablePatchedEventMatcherPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedEventMatcherPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
