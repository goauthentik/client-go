/*
authentik

Making authentication simple.

API version: 2024.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ConsentStageRequest ConsentStage Serializer
type ConsentStageRequest struct {
	Name    string                `json:"name"`
	FlowSet []FlowSetRequest      `json:"flow_set,omitempty"`
	Mode    *ConsentStageModeEnum `json:"mode,omitempty"`
	// Offset after which consent expires. (Format: hours=1;minutes=2;seconds=3).
	ConsentExpireIn *string `json:"consent_expire_in,omitempty"`
}

// NewConsentStageRequest instantiates a new ConsentStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsentStageRequest(name string) *ConsentStageRequest {
	this := ConsentStageRequest{}
	this.Name = name
	return &this
}

// NewConsentStageRequestWithDefaults instantiates a new ConsentStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsentStageRequestWithDefaults() *ConsentStageRequest {
	this := ConsentStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ConsentStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConsentStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConsentStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *ConsentStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *ConsentStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *ConsentStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ConsentStageRequest) GetMode() ConsentStageModeEnum {
	if o == nil || o.Mode == nil {
		var ret ConsentStageModeEnum
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentStageRequest) GetModeOk() (*ConsentStageModeEnum, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ConsentStageRequest) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given ConsentStageModeEnum and assigns it to the Mode field.
func (o *ConsentStageRequest) SetMode(v ConsentStageModeEnum) {
	o.Mode = &v
}

// GetConsentExpireIn returns the ConsentExpireIn field value if set, zero value otherwise.
func (o *ConsentStageRequest) GetConsentExpireIn() string {
	if o == nil || o.ConsentExpireIn == nil {
		var ret string
		return ret
	}
	return *o.ConsentExpireIn
}

// GetConsentExpireInOk returns a tuple with the ConsentExpireIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentStageRequest) GetConsentExpireInOk() (*string, bool) {
	if o == nil || o.ConsentExpireIn == nil {
		return nil, false
	}
	return o.ConsentExpireIn, true
}

// HasConsentExpireIn returns a boolean if a field has been set.
func (o *ConsentStageRequest) HasConsentExpireIn() bool {
	if o != nil && o.ConsentExpireIn != nil {
		return true
	}

	return false
}

// SetConsentExpireIn gets a reference to the given string and assigns it to the ConsentExpireIn field.
func (o *ConsentStageRequest) SetConsentExpireIn(v string) {
	o.ConsentExpireIn = &v
}

func (o ConsentStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
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

type NullableConsentStageRequest struct {
	value *ConsentStageRequest
	isSet bool
}

func (v NullableConsentStageRequest) Get() *ConsentStageRequest {
	return v.value
}

func (v *NullableConsentStageRequest) Set(val *ConsentStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentStageRequest(val *ConsentStageRequest) *NullableConsentStageRequest {
	return &NullableConsentStageRequest{value: val, isSet: true}
}

func (v NullableConsentStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
