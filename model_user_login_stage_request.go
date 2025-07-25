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

// UserLoginStageRequest UserLoginStage Serializer
type UserLoginStageRequest struct {
	Name    string           `json:"name"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Determines how long a session lasts. Default of 0 means that the sessions lasts until the browser is closed. (Format: hours=-1;minutes=-2;seconds=-3)
	SessionDuration *string `json:"session_duration,omitempty"`
	// Terminate all other sessions of the user logging in.
	TerminateOtherSessions *bool `json:"terminate_other_sessions,omitempty"`
	// Offset the session will be extended by when the user picks the remember me option. Default of 0 means that the remember me option will not be shown. (Format: hours=-1;minutes=-2;seconds=-3)
	RememberMeOffset *string `json:"remember_me_offset,omitempty"`
	// Bind sessions created by this stage to the configured network
	NetworkBinding *NetworkBindingEnum `json:"network_binding,omitempty"`
	// Bind sessions created by this stage to the configured GeoIP location
	GeoipBinding *GeoipBindingEnum `json:"geoip_binding,omitempty"`
	// When set to a non-zero value, authentik will save a cookie with a longer expiry,to remember the device the user is logging in from. (Format: hours=-1;minutes=-2;seconds=-3)
	RememberDevice *string `json:"remember_device,omitempty"`
}

// NewUserLoginStageRequest instantiates a new UserLoginStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLoginStageRequest(name string) *UserLoginStageRequest {
	this := UserLoginStageRequest{}
	this.Name = name
	return &this
}

// NewUserLoginStageRequestWithDefaults instantiates a new UserLoginStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLoginStageRequestWithDefaults() *UserLoginStageRequest {
	this := UserLoginStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UserLoginStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserLoginStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserLoginStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *UserLoginStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *UserLoginStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *UserLoginStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetSessionDuration returns the SessionDuration field value if set, zero value otherwise.
func (o *UserLoginStageRequest) GetSessionDuration() string {
	if o == nil || o.SessionDuration == nil {
		var ret string
		return ret
	}
	return *o.SessionDuration
}

// GetSessionDurationOk returns a tuple with the SessionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginStageRequest) GetSessionDurationOk() (*string, bool) {
	if o == nil || o.SessionDuration == nil {
		return nil, false
	}
	return o.SessionDuration, true
}

// HasSessionDuration returns a boolean if a field has been set.
func (o *UserLoginStageRequest) HasSessionDuration() bool {
	if o != nil && o.SessionDuration != nil {
		return true
	}

	return false
}

// SetSessionDuration gets a reference to the given string and assigns it to the SessionDuration field.
func (o *UserLoginStageRequest) SetSessionDuration(v string) {
	o.SessionDuration = &v
}

// GetTerminateOtherSessions returns the TerminateOtherSessions field value if set, zero value otherwise.
func (o *UserLoginStageRequest) GetTerminateOtherSessions() bool {
	if o == nil || o.TerminateOtherSessions == nil {
		var ret bool
		return ret
	}
	return *o.TerminateOtherSessions
}

// GetTerminateOtherSessionsOk returns a tuple with the TerminateOtherSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginStageRequest) GetTerminateOtherSessionsOk() (*bool, bool) {
	if o == nil || o.TerminateOtherSessions == nil {
		return nil, false
	}
	return o.TerminateOtherSessions, true
}

// HasTerminateOtherSessions returns a boolean if a field has been set.
func (o *UserLoginStageRequest) HasTerminateOtherSessions() bool {
	if o != nil && o.TerminateOtherSessions != nil {
		return true
	}

	return false
}

// SetTerminateOtherSessions gets a reference to the given bool and assigns it to the TerminateOtherSessions field.
func (o *UserLoginStageRequest) SetTerminateOtherSessions(v bool) {
	o.TerminateOtherSessions = &v
}

// GetRememberMeOffset returns the RememberMeOffset field value if set, zero value otherwise.
func (o *UserLoginStageRequest) GetRememberMeOffset() string {
	if o == nil || o.RememberMeOffset == nil {
		var ret string
		return ret
	}
	return *o.RememberMeOffset
}

// GetRememberMeOffsetOk returns a tuple with the RememberMeOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginStageRequest) GetRememberMeOffsetOk() (*string, bool) {
	if o == nil || o.RememberMeOffset == nil {
		return nil, false
	}
	return o.RememberMeOffset, true
}

// HasRememberMeOffset returns a boolean if a field has been set.
func (o *UserLoginStageRequest) HasRememberMeOffset() bool {
	if o != nil && o.RememberMeOffset != nil {
		return true
	}

	return false
}

// SetRememberMeOffset gets a reference to the given string and assigns it to the RememberMeOffset field.
func (o *UserLoginStageRequest) SetRememberMeOffset(v string) {
	o.RememberMeOffset = &v
}

// GetNetworkBinding returns the NetworkBinding field value if set, zero value otherwise.
func (o *UserLoginStageRequest) GetNetworkBinding() NetworkBindingEnum {
	if o == nil || o.NetworkBinding == nil {
		var ret NetworkBindingEnum
		return ret
	}
	return *o.NetworkBinding
}

// GetNetworkBindingOk returns a tuple with the NetworkBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginStageRequest) GetNetworkBindingOk() (*NetworkBindingEnum, bool) {
	if o == nil || o.NetworkBinding == nil {
		return nil, false
	}
	return o.NetworkBinding, true
}

// HasNetworkBinding returns a boolean if a field has been set.
func (o *UserLoginStageRequest) HasNetworkBinding() bool {
	if o != nil && o.NetworkBinding != nil {
		return true
	}

	return false
}

// SetNetworkBinding gets a reference to the given NetworkBindingEnum and assigns it to the NetworkBinding field.
func (o *UserLoginStageRequest) SetNetworkBinding(v NetworkBindingEnum) {
	o.NetworkBinding = &v
}

// GetGeoipBinding returns the GeoipBinding field value if set, zero value otherwise.
func (o *UserLoginStageRequest) GetGeoipBinding() GeoipBindingEnum {
	if o == nil || o.GeoipBinding == nil {
		var ret GeoipBindingEnum
		return ret
	}
	return *o.GeoipBinding
}

// GetGeoipBindingOk returns a tuple with the GeoipBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginStageRequest) GetGeoipBindingOk() (*GeoipBindingEnum, bool) {
	if o == nil || o.GeoipBinding == nil {
		return nil, false
	}
	return o.GeoipBinding, true
}

// HasGeoipBinding returns a boolean if a field has been set.
func (o *UserLoginStageRequest) HasGeoipBinding() bool {
	if o != nil && o.GeoipBinding != nil {
		return true
	}

	return false
}

// SetGeoipBinding gets a reference to the given GeoipBindingEnum and assigns it to the GeoipBinding field.
func (o *UserLoginStageRequest) SetGeoipBinding(v GeoipBindingEnum) {
	o.GeoipBinding = &v
}

// GetRememberDevice returns the RememberDevice field value if set, zero value otherwise.
func (o *UserLoginStageRequest) GetRememberDevice() string {
	if o == nil || o.RememberDevice == nil {
		var ret string
		return ret
	}
	return *o.RememberDevice
}

// GetRememberDeviceOk returns a tuple with the RememberDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginStageRequest) GetRememberDeviceOk() (*string, bool) {
	if o == nil || o.RememberDevice == nil {
		return nil, false
	}
	return o.RememberDevice, true
}

// HasRememberDevice returns a boolean if a field has been set.
func (o *UserLoginStageRequest) HasRememberDevice() bool {
	if o != nil && o.RememberDevice != nil {
		return true
	}

	return false
}

// SetRememberDevice gets a reference to the given string and assigns it to the RememberDevice field.
func (o *UserLoginStageRequest) SetRememberDevice(v string) {
	o.RememberDevice = &v
}

func (o UserLoginStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
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
	if o.RememberDevice != nil {
		toSerialize["remember_device"] = o.RememberDevice
	}
	return json.Marshal(toSerialize)
}

type NullableUserLoginStageRequest struct {
	value *UserLoginStageRequest
	isSet bool
}

func (v NullableUserLoginStageRequest) Get() *UserLoginStageRequest {
	return v.value
}

func (v *NullableUserLoginStageRequest) Set(val *UserLoginStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLoginStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLoginStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLoginStageRequest(val *UserLoginStageRequest) *NullableUserLoginStageRequest {
	return &NullableUserLoginStageRequest{value: val, isSet: true}
}

func (v NullableUserLoginStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLoginStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
