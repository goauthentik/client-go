/*
authentik

Making authentication simple.

API version: 2024.4.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedSourceStageRequest SourceStage Serializer
type PatchedSourceStageRequest struct {
	Name    *string          `json:"name,omitempty"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	Source  *string          `json:"source,omitempty"`
	// Amount of time a user can take to return from the source to continue the flow (Format: hours=-1;minutes=-2;seconds=-3)
	ResumeTimeout *string `json:"resume_timeout,omitempty"`
}

// NewPatchedSourceStageRequest instantiates a new PatchedSourceStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedSourceStageRequest() *PatchedSourceStageRequest {
	this := PatchedSourceStageRequest{}
	return &this
}

// NewPatchedSourceStageRequestWithDefaults instantiates a new PatchedSourceStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedSourceStageRequestWithDefaults() *PatchedSourceStageRequest {
	this := PatchedSourceStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedSourceStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSourceStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedSourceStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedSourceStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedSourceStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSourceStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedSourceStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedSourceStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PatchedSourceStageRequest) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSourceStageRequest) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PatchedSourceStageRequest) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *PatchedSourceStageRequest) SetSource(v string) {
	o.Source = &v
}

// GetResumeTimeout returns the ResumeTimeout field value if set, zero value otherwise.
func (o *PatchedSourceStageRequest) GetResumeTimeout() string {
	if o == nil || o.ResumeTimeout == nil {
		var ret string
		return ret
	}
	return *o.ResumeTimeout
}

// GetResumeTimeoutOk returns a tuple with the ResumeTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSourceStageRequest) GetResumeTimeoutOk() (*string, bool) {
	if o == nil || o.ResumeTimeout == nil {
		return nil, false
	}
	return o.ResumeTimeout, true
}

// HasResumeTimeout returns a boolean if a field has been set.
func (o *PatchedSourceStageRequest) HasResumeTimeout() bool {
	if o != nil && o.ResumeTimeout != nil {
		return true
	}

	return false
}

// SetResumeTimeout gets a reference to the given string and assigns it to the ResumeTimeout field.
func (o *PatchedSourceStageRequest) SetResumeTimeout(v string) {
	o.ResumeTimeout = &v
}

func (o PatchedSourceStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.ResumeTimeout != nil {
		toSerialize["resume_timeout"] = o.ResumeTimeout
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedSourceStageRequest struct {
	value *PatchedSourceStageRequest
	isSet bool
}

func (v NullablePatchedSourceStageRequest) Get() *PatchedSourceStageRequest {
	return v.value
}

func (v *NullablePatchedSourceStageRequest) Set(val *PatchedSourceStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedSourceStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedSourceStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedSourceStageRequest(val *PatchedSourceStageRequest) *NullablePatchedSourceStageRequest {
	return &NullablePatchedSourceStageRequest{value: val, isSet: true}
}

func (v NullablePatchedSourceStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedSourceStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
