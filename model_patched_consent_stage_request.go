/*
authentik

Making authentication simple.

API version: 2025.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedConsentStageRequest ConsentStage Serializer
type PatchedConsentStageRequest struct {
	Name    *string               `json:"name,omitempty"`
	FlowSet []FlowSetRequest      `json:"flow_set,omitempty"`
	Mode    *ConsentStageModeEnum `json:"mode,omitempty"`
	// Offset after which consent expires. (Format: hours=1;minutes=2;seconds=3).
	ConsentExpireIn *string `json:"consent_expire_in,omitempty"`
}

// NewPatchedConsentStageRequest instantiates a new PatchedConsentStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedConsentStageRequest() *PatchedConsentStageRequest {
	this := PatchedConsentStageRequest{}
	return &this
}

// NewPatchedConsentStageRequestWithDefaults instantiates a new PatchedConsentStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedConsentStageRequestWithDefaults() *PatchedConsentStageRequest {
	this := PatchedConsentStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedConsentStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedConsentStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedConsentStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedConsentStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedConsentStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedConsentStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedConsentStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedConsentStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *PatchedConsentStageRequest) GetMode() ConsentStageModeEnum {
	if o == nil || o.Mode == nil {
		var ret ConsentStageModeEnum
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedConsentStageRequest) GetModeOk() (*ConsentStageModeEnum, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *PatchedConsentStageRequest) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given ConsentStageModeEnum and assigns it to the Mode field.
func (o *PatchedConsentStageRequest) SetMode(v ConsentStageModeEnum) {
	o.Mode = &v
}

// GetConsentExpireIn returns the ConsentExpireIn field value if set, zero value otherwise.
func (o *PatchedConsentStageRequest) GetConsentExpireIn() string {
	if o == nil || o.ConsentExpireIn == nil {
		var ret string
		return ret
	}
	return *o.ConsentExpireIn
}

// GetConsentExpireInOk returns a tuple with the ConsentExpireIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedConsentStageRequest) GetConsentExpireInOk() (*string, bool) {
	if o == nil || o.ConsentExpireIn == nil {
		return nil, false
	}
	return o.ConsentExpireIn, true
}

// HasConsentExpireIn returns a boolean if a field has been set.
func (o *PatchedConsentStageRequest) HasConsentExpireIn() bool {
	if o != nil && o.ConsentExpireIn != nil {
		return true
	}

	return false
}

// SetConsentExpireIn gets a reference to the given string and assigns it to the ConsentExpireIn field.
func (o *PatchedConsentStageRequest) SetConsentExpireIn(v string) {
	o.ConsentExpireIn = &v
}

func (o PatchedConsentStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.ConsentExpireIn != nil {
		toSerialize["consent_expire_in"] = o.ConsentExpireIn
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedConsentStageRequest struct {
	value *PatchedConsentStageRequest
	isSet bool
}

func (v NullablePatchedConsentStageRequest) Get() *PatchedConsentStageRequest {
	return v.value
}

func (v *NullablePatchedConsentStageRequest) Set(val *PatchedConsentStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedConsentStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedConsentStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedConsentStageRequest(val *PatchedConsentStageRequest) *NullablePatchedConsentStageRequest {
	return &NullablePatchedConsentStageRequest{value: val, isSet: true}
}

func (v NullablePatchedConsentStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedConsentStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
