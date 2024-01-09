/*
authentik

Making authentication simple.

API version: 2023.10.6
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserLoginStage UserLoginStage Serializer
type UserLoginStage struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// Get object type so that we know how to edit the object
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName string    `json:"meta_model_name"`
	FlowSet       []FlowSet `json:"flow_set,omitempty"`
	// Determines how long a session lasts. Default of 0 means that the sessions lasts until the browser is closed. (Format: hours=-1;minutes=-2;seconds=-3)
	SessionDuration *string `json:"session_duration,omitempty"`
	// Terminate all other sessions of the user logging in.
	TerminateOtherSessions *bool `json:"terminate_other_sessions,omitempty"`
	// Offset the session will be extended by when the user picks the remember me option. Default of 0 means that the remember me option will not be shown. (Format: hours=-1;minutes=-2;seconds=-3)
	RememberMeOffset *string `json:"remember_me_offset,omitempty"`
	// Bind sessions created by this stage to the configured network  * `no_binding` - No Binding * `bind_asn` - Bind Asn * `bind_asn_network` - Bind Asn Network * `bind_asn_network_ip` - Bind Asn Network Ip
	NetworkBinding *NetworkBindingEnum `json:"network_binding,omitempty"`
	// Bind sessions created by this stage to the configured GeoIP location  * `no_binding` - No Binding * `bind_continent` - Bind Continent * `bind_continent_country` - Bind Continent Country * `bind_continent_country_city` - Bind Continent Country City
	GeoipBinding *GeoipBindingEnum `json:"geoip_binding,omitempty"`
}

// NewUserLoginStage instantiates a new UserLoginStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLoginStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string) *UserLoginStage {
	this := UserLoginStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	return &this
}

// NewUserLoginStageWithDefaults instantiates a new UserLoginStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLoginStageWithDefaults() *UserLoginStage {
	this := UserLoginStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *UserLoginStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *UserLoginStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *UserLoginStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserLoginStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *UserLoginStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *UserLoginStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *UserLoginStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *UserLoginStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *UserLoginStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *UserLoginStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *UserLoginStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *UserLoginStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *UserLoginStage) GetFlowSet() []FlowSet {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSet
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetFlowSetOk() ([]FlowSet, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *UserLoginStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSet and assigns it to the FlowSet field.
func (o *UserLoginStage) SetFlowSet(v []FlowSet) {
	o.FlowSet = v
}

// GetSessionDuration returns the SessionDuration field value if set, zero value otherwise.
func (o *UserLoginStage) GetSessionDuration() string {
	if o == nil || o.SessionDuration == nil {
		var ret string
		return ret
	}
	return *o.SessionDuration
}

// GetSessionDurationOk returns a tuple with the SessionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetSessionDurationOk() (*string, bool) {
	if o == nil || o.SessionDuration == nil {
		return nil, false
	}
	return o.SessionDuration, true
}

// HasSessionDuration returns a boolean if a field has been set.
func (o *UserLoginStage) HasSessionDuration() bool {
	if o != nil && o.SessionDuration != nil {
		return true
	}

	return false
}

// SetSessionDuration gets a reference to the given string and assigns it to the SessionDuration field.
func (o *UserLoginStage) SetSessionDuration(v string) {
	o.SessionDuration = &v
}

// GetTerminateOtherSessions returns the TerminateOtherSessions field value if set, zero value otherwise.
func (o *UserLoginStage) GetTerminateOtherSessions() bool {
	if o == nil || o.TerminateOtherSessions == nil {
		var ret bool
		return ret
	}
	return *o.TerminateOtherSessions
}

// GetTerminateOtherSessionsOk returns a tuple with the TerminateOtherSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetTerminateOtherSessionsOk() (*bool, bool) {
	if o == nil || o.TerminateOtherSessions == nil {
		return nil, false
	}
	return o.TerminateOtherSessions, true
}

// HasTerminateOtherSessions returns a boolean if a field has been set.
func (o *UserLoginStage) HasTerminateOtherSessions() bool {
	if o != nil && o.TerminateOtherSessions != nil {
		return true
	}

	return false
}

// SetTerminateOtherSessions gets a reference to the given bool and assigns it to the TerminateOtherSessions field.
func (o *UserLoginStage) SetTerminateOtherSessions(v bool) {
	o.TerminateOtherSessions = &v
}

// GetRememberMeOffset returns the RememberMeOffset field value if set, zero value otherwise.
func (o *UserLoginStage) GetRememberMeOffset() string {
	if o == nil || o.RememberMeOffset == nil {
		var ret string
		return ret
	}
	return *o.RememberMeOffset
}

// GetRememberMeOffsetOk returns a tuple with the RememberMeOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetRememberMeOffsetOk() (*string, bool) {
	if o == nil || o.RememberMeOffset == nil {
		return nil, false
	}
	return o.RememberMeOffset, true
}

// HasRememberMeOffset returns a boolean if a field has been set.
func (o *UserLoginStage) HasRememberMeOffset() bool {
	if o != nil && o.RememberMeOffset != nil {
		return true
	}

	return false
}

// SetRememberMeOffset gets a reference to the given string and assigns it to the RememberMeOffset field.
func (o *UserLoginStage) SetRememberMeOffset(v string) {
	o.RememberMeOffset = &v
}

// GetNetworkBinding returns the NetworkBinding field value if set, zero value otherwise.
func (o *UserLoginStage) GetNetworkBinding() NetworkBindingEnum {
	if o == nil || o.NetworkBinding == nil {
		var ret NetworkBindingEnum
		return ret
	}
	return *o.NetworkBinding
}

// GetNetworkBindingOk returns a tuple with the NetworkBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetNetworkBindingOk() (*NetworkBindingEnum, bool) {
	if o == nil || o.NetworkBinding == nil {
		return nil, false
	}
	return o.NetworkBinding, true
}

// HasNetworkBinding returns a boolean if a field has been set.
func (o *UserLoginStage) HasNetworkBinding() bool {
	if o != nil && o.NetworkBinding != nil {
		return true
	}

	return false
}

// SetNetworkBinding gets a reference to the given NetworkBindingEnum and assigns it to the NetworkBinding field.
func (o *UserLoginStage) SetNetworkBinding(v NetworkBindingEnum) {
	o.NetworkBinding = &v
}

// GetGeoipBinding returns the GeoipBinding field value if set, zero value otherwise.
func (o *UserLoginStage) GetGeoipBinding() GeoipBindingEnum {
	if o == nil || o.GeoipBinding == nil {
		var ret GeoipBindingEnum
		return ret
	}
	return *o.GeoipBinding
}

// GetGeoipBindingOk returns a tuple with the GeoipBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetGeoipBindingOk() (*GeoipBindingEnum, bool) {
	if o == nil || o.GeoipBinding == nil {
		return nil, false
	}
	return o.GeoipBinding, true
}

// HasGeoipBinding returns a boolean if a field has been set.
func (o *UserLoginStage) HasGeoipBinding() bool {
	if o != nil && o.GeoipBinding != nil {
		return true
	}

	return false
}

// SetGeoipBinding gets a reference to the given GeoipBindingEnum and assigns it to the GeoipBinding field.
func (o *UserLoginStage) SetGeoipBinding(v GeoipBindingEnum) {
	o.GeoipBinding = &v
}

func (o UserLoginStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["verbose_name"] = o.VerboseName
	}
	if true {
		toSerialize["verbose_name_plural"] = o.VerboseNamePlural
	}
	if true {
		toSerialize["meta_model_name"] = o.MetaModelName
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.SessionDuration != nil {
		toSerialize["session_duration"] = o.SessionDuration
	}
	if o.TerminateOtherSessions != nil {
		toSerialize["terminate_other_sessions"] = o.TerminateOtherSessions
	}
	if o.RememberMeOffset != nil {
		toSerialize["remember_me_offset"] = o.RememberMeOffset
	}
	if o.NetworkBinding != nil {
		toSerialize["network_binding"] = o.NetworkBinding
	}
	if o.GeoipBinding != nil {
		toSerialize["geoip_binding"] = o.GeoipBinding
	}
	return json.Marshal(toSerialize)
}

type NullableUserLoginStage struct {
	value *UserLoginStage
	isSet bool
}

func (v NullableUserLoginStage) Get() *UserLoginStage {
	return v.value
}

func (v *NullableUserLoginStage) Set(val *UserLoginStage) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLoginStage) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLoginStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLoginStage(val *UserLoginStage) *NullableUserLoginStage {
	return &NullableUserLoginStage{value: val, isSet: true}
}

func (v NullableUserLoginStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLoginStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
