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

// checks if the EventMatcherPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventMatcherPolicy{}

// EventMatcherPolicy Event Matcher Policy Serializer
type EventMatcherPolicy struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	// Get object component so that we know how to edit the object
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName string `json:"meta_model_name"`
	// Return objects policy is bound to
	BoundTo int32         `json:"bound_to"`
	Action  *EventActions `json:"action,omitempty"`
	// Matches Event's Client IP (strict matching, for network matching use an Expression Policy)
	ClientIp *string  `json:"client_ip,omitempty"`
	App      *AppEnum `json:"app,omitempty"`
}

// NewEventMatcherPolicy instantiates a new EventMatcherPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventMatcherPolicy(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, boundTo int32) *EventMatcherPolicy {
	this := EventMatcherPolicy{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.BoundTo = boundTo
	return &this
}

// NewEventMatcherPolicyWithDefaults instantiates a new EventMatcherPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventMatcherPolicyWithDefaults() *EventMatcherPolicy {
	this := EventMatcherPolicy{}
	return &this
}

// GetPk returns the Pk field value
func (o *EventMatcherPolicy) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *EventMatcherPolicy) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *EventMatcherPolicy) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *EventMatcherPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventMatcherPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventMatcherPolicy) SetName(v string) {
	o.Name = v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *EventMatcherPolicy) GetExecutionLogging() bool {
	if o == nil || IsNil(o.ExecutionLogging) {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMatcherPolicy) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || IsNil(o.ExecutionLogging) {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *EventMatcherPolicy) HasExecutionLogging() bool {
	if o != nil && !IsNil(o.ExecutionLogging) {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *EventMatcherPolicy) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetComponent returns the Component field value
func (o *EventMatcherPolicy) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *EventMatcherPolicy) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *EventMatcherPolicy) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *EventMatcherPolicy) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *EventMatcherPolicy) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *EventMatcherPolicy) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *EventMatcherPolicy) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *EventMatcherPolicy) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *EventMatcherPolicy) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *EventMatcherPolicy) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *EventMatcherPolicy) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *EventMatcherPolicy) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetBoundTo returns the BoundTo field value
func (o *EventMatcherPolicy) GetBoundTo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BoundTo
}

// GetBoundToOk returns a tuple with the BoundTo field value
// and a boolean to check if the value has been set.
func (o *EventMatcherPolicy) GetBoundToOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoundTo, true
}

// SetBoundTo sets field value
func (o *EventMatcherPolicy) SetBoundTo(v int32) {
	o.BoundTo = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *EventMatcherPolicy) GetAction() EventActions {
	if o == nil || IsNil(o.Action) {
		var ret EventActions
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMatcherPolicy) GetActionOk() (*EventActions, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *EventMatcherPolicy) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given EventActions and assigns it to the Action field.
func (o *EventMatcherPolicy) SetAction(v EventActions) {
	o.Action = &v
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *EventMatcherPolicy) GetClientIp() string {
	if o == nil || IsNil(o.ClientIp) {
		var ret string
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMatcherPolicy) GetClientIpOk() (*string, bool) {
	if o == nil || IsNil(o.ClientIp) {
		return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *EventMatcherPolicy) HasClientIp() bool {
	if o != nil && !IsNil(o.ClientIp) {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given string and assigns it to the ClientIp field.
func (o *EventMatcherPolicy) SetClientIp(v string) {
	o.ClientIp = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *EventMatcherPolicy) GetApp() AppEnum {
	if o == nil || IsNil(o.App) {
		var ret AppEnum
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMatcherPolicy) GetAppOk() (*AppEnum, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *EventMatcherPolicy) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppEnum and assigns it to the App field.
func (o *EventMatcherPolicy) SetApp(v AppEnum) {
	o.App = &v
}

func (o EventMatcherPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventMatcherPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pk is readOnly
	toSerialize["name"] = o.Name
	if !IsNil(o.ExecutionLogging) {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	// skip: component is readOnly
	// skip: verbose_name is readOnly
	// skip: verbose_name_plural is readOnly
	// skip: meta_model_name is readOnly
	// skip: bound_to is readOnly
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.ClientIp) {
		toSerialize["client_ip"] = o.ClientIp
	}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	return toSerialize, nil
}

type NullableEventMatcherPolicy struct {
	value *EventMatcherPolicy
	isSet bool
}

func (v NullableEventMatcherPolicy) Get() *EventMatcherPolicy {
	return v.value
}

func (v *NullableEventMatcherPolicy) Set(val *EventMatcherPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableEventMatcherPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableEventMatcherPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventMatcherPolicy(val *EventMatcherPolicy) *NullableEventMatcherPolicy {
	return &NullableEventMatcherPolicy{value: val, isSet: true}
}

func (v NullableEventMatcherPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventMatcherPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
