/*
authentik

Making authentication simple.

API version: 2025.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SourceStageRequest SourceStage Serializer
type SourceStageRequest struct {
	Name    string           `json:"name"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	Source  string           `json:"source"`
	// Amount of time a user can take to return from the source to continue the flow (Format: hours=-1;minutes=-2;seconds=-3)
	ResumeTimeout *string `json:"resume_timeout,omitempty"`
}

// NewSourceStageRequest instantiates a new SourceStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceStageRequest(name string, source string) *SourceStageRequest {
	this := SourceStageRequest{}
	this.Name = name
	this.Source = source
	return &this
}

// NewSourceStageRequestWithDefaults instantiates a new SourceStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceStageRequestWithDefaults() *SourceStageRequest {
	this := SourceStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *SourceStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SourceStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SourceStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *SourceStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *SourceStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *SourceStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetSource returns the Source field value
func (o *SourceStageRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *SourceStageRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *SourceStageRequest) SetSource(v string) {
	o.Source = v
}

// GetResumeTimeout returns the ResumeTimeout field value if set, zero value otherwise.
func (o *SourceStageRequest) GetResumeTimeout() string {
	if o == nil || o.ResumeTimeout == nil {
		var ret string
		return ret
	}
	return *o.ResumeTimeout
}

// GetResumeTimeoutOk returns a tuple with the ResumeTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceStageRequest) GetResumeTimeoutOk() (*string, bool) {
	if o == nil || o.ResumeTimeout == nil {
		return nil, false
	}
	return o.ResumeTimeout, true
}

// HasResumeTimeout returns a boolean if a field has been set.
func (o *SourceStageRequest) HasResumeTimeout() bool {
	if o != nil && o.ResumeTimeout != nil {
		return true
	}

	return false
}

// SetResumeTimeout gets a reference to the given string and assigns it to the ResumeTimeout field.
func (o *SourceStageRequest) SetResumeTimeout(v string) {
	o.ResumeTimeout = &v
}

func (o SourceStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if o.ResumeTimeout != nil {
		toSerialize["resume_timeout"] = o.ResumeTimeout
	}
	return json.Marshal(toSerialize)
}

type NullableSourceStageRequest struct {
	value *SourceStageRequest
	isSet bool
}

func (v NullableSourceStageRequest) Get() *SourceStageRequest {
	return v.value
}

func (v *NullableSourceStageRequest) Set(val *SourceStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceStageRequest(val *SourceStageRequest) *NullableSourceStageRequest {
	return &NullableSourceStageRequest{value: val, isSet: true}
}

func (v NullableSourceStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
