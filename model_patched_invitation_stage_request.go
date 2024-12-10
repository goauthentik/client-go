/*
authentik

Making authentication simple.

API version: 2024.10.5
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedInvitationStageRequest InvitationStage Serializer
type PatchedInvitationStageRequest struct {
	Name    *string          `json:"name,omitempty"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// If this flag is set, this Stage will jump to the next Stage when no Invitation is given. By default this Stage will cancel the Flow when no invitation is given.
	ContinueFlowWithoutInvitation *bool `json:"continue_flow_without_invitation,omitempty"`
}

// NewPatchedInvitationStageRequest instantiates a new PatchedInvitationStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedInvitationStageRequest() *PatchedInvitationStageRequest {
	this := PatchedInvitationStageRequest{}
	return &this
}

// NewPatchedInvitationStageRequestWithDefaults instantiates a new PatchedInvitationStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedInvitationStageRequestWithDefaults() *PatchedInvitationStageRequest {
	this := PatchedInvitationStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedInvitationStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInvitationStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedInvitationStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedInvitationStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedInvitationStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInvitationStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedInvitationStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedInvitationStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetContinueFlowWithoutInvitation returns the ContinueFlowWithoutInvitation field value if set, zero value otherwise.
func (o *PatchedInvitationStageRequest) GetContinueFlowWithoutInvitation() bool {
	if o == nil || o.ContinueFlowWithoutInvitation == nil {
		var ret bool
		return ret
	}
	return *o.ContinueFlowWithoutInvitation
}

// GetContinueFlowWithoutInvitationOk returns a tuple with the ContinueFlowWithoutInvitation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInvitationStageRequest) GetContinueFlowWithoutInvitationOk() (*bool, bool) {
	if o == nil || o.ContinueFlowWithoutInvitation == nil {
		return nil, false
	}
	return o.ContinueFlowWithoutInvitation, true
}

// HasContinueFlowWithoutInvitation returns a boolean if a field has been set.
func (o *PatchedInvitationStageRequest) HasContinueFlowWithoutInvitation() bool {
	if o != nil && o.ContinueFlowWithoutInvitation != nil {
		return true
	}

	return false
}

// SetContinueFlowWithoutInvitation gets a reference to the given bool and assigns it to the ContinueFlowWithoutInvitation field.
func (o *PatchedInvitationStageRequest) SetContinueFlowWithoutInvitation(v bool) {
	o.ContinueFlowWithoutInvitation = &v
}

func (o PatchedInvitationStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.ContinueFlowWithoutInvitation != nil {
		toSerialize["continue_flow_without_invitation"] = o.ContinueFlowWithoutInvitation
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedInvitationStageRequest struct {
	value *PatchedInvitationStageRequest
	isSet bool
}

func (v NullablePatchedInvitationStageRequest) Get() *PatchedInvitationStageRequest {
	return v.value
}

func (v *NullablePatchedInvitationStageRequest) Set(val *PatchedInvitationStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedInvitationStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedInvitationStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedInvitationStageRequest(val *PatchedInvitationStageRequest) *NullablePatchedInvitationStageRequest {
	return &NullablePatchedInvitationStageRequest{value: val, isSet: true}
}

func (v NullablePatchedInvitationStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedInvitationStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
