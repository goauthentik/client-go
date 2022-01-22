/*
authentik

Making authentication simple.

API version: 2022.1.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedUserLoginStageRequest UserLoginStage Serializer
type PatchedUserLoginStageRequest struct {
	Name    *string        `json:"name,omitempty"`
	FlowSet *[]FlowRequest `json:"flow_set,omitempty"`
	// Determines how long a session lasts. Default of 0 means that the sessions lasts until the browser is closed. (Format: hours=-1;minutes=-2;seconds=-3)
	SessionDuration *string `json:"session_duration,omitempty"`
}

// NewPatchedUserLoginStageRequest instantiates a new PatchedUserLoginStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedUserLoginStageRequest() *PatchedUserLoginStageRequest {
	this := PatchedUserLoginStageRequest{}
	return &this
}

// NewPatchedUserLoginStageRequestWithDefaults instantiates a new PatchedUserLoginStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedUserLoginStageRequestWithDefaults() *PatchedUserLoginStageRequest {
	this := PatchedUserLoginStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedUserLoginStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserLoginStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedUserLoginStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedUserLoginStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedUserLoginStageRequest) GetFlowSet() []FlowRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowRequest
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserLoginStageRequest) GetFlowSetOk() (*[]FlowRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedUserLoginStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowRequest and assigns it to the FlowSet field.
func (o *PatchedUserLoginStageRequest) SetFlowSet(v []FlowRequest) {
	o.FlowSet = &v
}

// GetSessionDuration returns the SessionDuration field value if set, zero value otherwise.
func (o *PatchedUserLoginStageRequest) GetSessionDuration() string {
	if o == nil || o.SessionDuration == nil {
		var ret string
		return ret
	}
	return *o.SessionDuration
}

// GetSessionDurationOk returns a tuple with the SessionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserLoginStageRequest) GetSessionDurationOk() (*string, bool) {
	if o == nil || o.SessionDuration == nil {
		return nil, false
	}
	return o.SessionDuration, true
}

// HasSessionDuration returns a boolean if a field has been set.
func (o *PatchedUserLoginStageRequest) HasSessionDuration() bool {
	if o != nil && o.SessionDuration != nil {
		return true
	}

	return false
}

// SetSessionDuration gets a reference to the given string and assigns it to the SessionDuration field.
func (o *PatchedUserLoginStageRequest) SetSessionDuration(v string) {
	o.SessionDuration = &v
}

func (o PatchedUserLoginStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.SessionDuration != nil {
		toSerialize["session_duration"] = o.SessionDuration
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedUserLoginStageRequest struct {
	value *PatchedUserLoginStageRequest
	isSet bool
}

func (v NullablePatchedUserLoginStageRequest) Get() *PatchedUserLoginStageRequest {
	return v.value
}

func (v *NullablePatchedUserLoginStageRequest) Set(val *PatchedUserLoginStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedUserLoginStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedUserLoginStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedUserLoginStageRequest(val *PatchedUserLoginStageRequest) *NullablePatchedUserLoginStageRequest {
	return &NullablePatchedUserLoginStageRequest{value: val, isSet: true}
}

func (v NullablePatchedUserLoginStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedUserLoginStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
