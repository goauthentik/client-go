/*
authentik

Making authentication simple.

API version: 2022.4.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedCaptchaStageRequest CaptchaStage Serializer
type PatchedCaptchaStageRequest struct {
	Name    *string        `json:"name,omitempty"`
	FlowSet *[]FlowRequest `json:"flow_set,omitempty"`
	// Public key, acquired from https://www.google.com/recaptcha/intro/v3.html
	PublicKey *string `json:"public_key,omitempty"`
	// Private key, acquired from https://www.google.com/recaptcha/intro/v3.html
	PrivateKey *string `json:"private_key,omitempty"`
}

// NewPatchedCaptchaStageRequest instantiates a new PatchedCaptchaStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedCaptchaStageRequest() *PatchedCaptchaStageRequest {
	this := PatchedCaptchaStageRequest{}
	return &this
}

// NewPatchedCaptchaStageRequestWithDefaults instantiates a new PatchedCaptchaStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedCaptchaStageRequestWithDefaults() *PatchedCaptchaStageRequest {
	this := PatchedCaptchaStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedCaptchaStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCaptchaStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedCaptchaStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedCaptchaStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedCaptchaStageRequest) GetFlowSet() []FlowRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowRequest
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCaptchaStageRequest) GetFlowSetOk() (*[]FlowRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedCaptchaStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowRequest and assigns it to the FlowSet field.
func (o *PatchedCaptchaStageRequest) SetFlowSet(v []FlowRequest) {
	o.FlowSet = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *PatchedCaptchaStageRequest) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCaptchaStageRequest) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *PatchedCaptchaStageRequest) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *PatchedCaptchaStageRequest) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *PatchedCaptchaStageRequest) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCaptchaStageRequest) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *PatchedCaptchaStageRequest) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *PatchedCaptchaStageRequest) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

func (o PatchedCaptchaStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
	if o.PrivateKey != nil {
		toSerialize["private_key"] = o.PrivateKey
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedCaptchaStageRequest struct {
	value *PatchedCaptchaStageRequest
	isSet bool
}

func (v NullablePatchedCaptchaStageRequest) Get() *PatchedCaptchaStageRequest {
	return v.value
}

func (v *NullablePatchedCaptchaStageRequest) Set(val *PatchedCaptchaStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedCaptchaStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedCaptchaStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedCaptchaStageRequest(val *PatchedCaptchaStageRequest) *NullablePatchedCaptchaStageRequest {
	return &NullablePatchedCaptchaStageRequest{value: val, isSet: true}
}

func (v NullablePatchedCaptchaStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedCaptchaStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
