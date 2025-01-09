/*
authentik

Making authentication simple.

API version: 2024.12.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedCaptchaStageRequest CaptchaStage Serializer
type PatchedCaptchaStageRequest struct {
	Name    *string          `json:"name,omitempty"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Public key, acquired your captcha Provider.
	PublicKey *string `json:"public_key,omitempty"`
	// Private key, acquired your captcha Provider.
	PrivateKey        *string  `json:"private_key,omitempty"`
	JsUrl             *string  `json:"js_url,omitempty"`
	ApiUrl            *string  `json:"api_url,omitempty"`
	Interactive       *bool    `json:"interactive,omitempty"`
	ScoreMinThreshold *float64 `json:"score_min_threshold,omitempty"`
	ScoreMaxThreshold *float64 `json:"score_max_threshold,omitempty"`
	// When enabled and the received captcha score is outside of the given threshold, the stage will show an error message. When not enabled, the flow will continue, but the data from the captcha will be available in the context for policy decisions
	ErrorOnInvalidScore *bool `json:"error_on_invalid_score,omitempty"`
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
func (o *PatchedCaptchaStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCaptchaStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
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

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedCaptchaStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
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

// GetJsUrl returns the JsUrl field value if set, zero value otherwise.
func (o *PatchedCaptchaStageRequest) GetJsUrl() string {
	if o == nil || o.JsUrl == nil {
		var ret string
		return ret
	}
	return *o.JsUrl
}

// GetJsUrlOk returns a tuple with the JsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCaptchaStageRequest) GetJsUrlOk() (*string, bool) {
	if o == nil || o.JsUrl == nil {
		return nil, false
	}
	return o.JsUrl, true
}

// HasJsUrl returns a boolean if a field has been set.
func (o *PatchedCaptchaStageRequest) HasJsUrl() bool {
	if o != nil && o.JsUrl != nil {
		return true
	}

	return false
}

// SetJsUrl gets a reference to the given string and assigns it to the JsUrl field.
func (o *PatchedCaptchaStageRequest) SetJsUrl(v string) {
	o.JsUrl = &v
}

// GetApiUrl returns the ApiUrl field value if set, zero value otherwise.
func (o *PatchedCaptchaStageRequest) GetApiUrl() string {
	if o == nil || o.ApiUrl == nil {
		var ret string
		return ret
	}
	return *o.ApiUrl
}

// GetApiUrlOk returns a tuple with the ApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCaptchaStageRequest) GetApiUrlOk() (*string, bool) {
	if o == nil || o.ApiUrl == nil {
		return nil, false
	}
	return o.ApiUrl, true
}

// HasApiUrl returns a boolean if a field has been set.
func (o *PatchedCaptchaStageRequest) HasApiUrl() bool {
	if o != nil && o.ApiUrl != nil {
		return true
	}

	return false
}

// SetApiUrl gets a reference to the given string and assigns it to the ApiUrl field.
func (o *PatchedCaptchaStageRequest) SetApiUrl(v string) {
	o.ApiUrl = &v
}

// GetInteractive returns the Interactive field value if set, zero value otherwise.
func (o *PatchedCaptchaStageRequest) GetInteractive() bool {
	if o == nil || o.Interactive == nil {
		var ret bool
		return ret
	}
	return *o.Interactive
}

// GetInteractiveOk returns a tuple with the Interactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCaptchaStageRequest) GetInteractiveOk() (*bool, bool) {
	if o == nil || o.Interactive == nil {
		return nil, false
	}
	return o.Interactive, true
}

// HasInteractive returns a boolean if a field has been set.
func (o *PatchedCaptchaStageRequest) HasInteractive() bool {
	if o != nil && o.Interactive != nil {
		return true
	}

	return false
}

// SetInteractive gets a reference to the given bool and assigns it to the Interactive field.
func (o *PatchedCaptchaStageRequest) SetInteractive(v bool) {
	o.Interactive = &v
}

// GetScoreMinThreshold returns the ScoreMinThreshold field value if set, zero value otherwise.
func (o *PatchedCaptchaStageRequest) GetScoreMinThreshold() float64 {
	if o == nil || o.ScoreMinThreshold == nil {
		var ret float64
		return ret
	}
	return *o.ScoreMinThreshold
}

// GetScoreMinThresholdOk returns a tuple with the ScoreMinThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCaptchaStageRequest) GetScoreMinThresholdOk() (*float64, bool) {
	if o == nil || o.ScoreMinThreshold == nil {
		return nil, false
	}
	return o.ScoreMinThreshold, true
}

// HasScoreMinThreshold returns a boolean if a field has been set.
func (o *PatchedCaptchaStageRequest) HasScoreMinThreshold() bool {
	if o != nil && o.ScoreMinThreshold != nil {
		return true
	}

	return false
}

// SetScoreMinThreshold gets a reference to the given float64 and assigns it to the ScoreMinThreshold field.
func (o *PatchedCaptchaStageRequest) SetScoreMinThreshold(v float64) {
	o.ScoreMinThreshold = &v
}

// GetScoreMaxThreshold returns the ScoreMaxThreshold field value if set, zero value otherwise.
func (o *PatchedCaptchaStageRequest) GetScoreMaxThreshold() float64 {
	if o == nil || o.ScoreMaxThreshold == nil {
		var ret float64
		return ret
	}
	return *o.ScoreMaxThreshold
}

// GetScoreMaxThresholdOk returns a tuple with the ScoreMaxThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCaptchaStageRequest) GetScoreMaxThresholdOk() (*float64, bool) {
	if o == nil || o.ScoreMaxThreshold == nil {
		return nil, false
	}
	return o.ScoreMaxThreshold, true
}

// HasScoreMaxThreshold returns a boolean if a field has been set.
func (o *PatchedCaptchaStageRequest) HasScoreMaxThreshold() bool {
	if o != nil && o.ScoreMaxThreshold != nil {
		return true
	}

	return false
}

// SetScoreMaxThreshold gets a reference to the given float64 and assigns it to the ScoreMaxThreshold field.
func (o *PatchedCaptchaStageRequest) SetScoreMaxThreshold(v float64) {
	o.ScoreMaxThreshold = &v
}

// GetErrorOnInvalidScore returns the ErrorOnInvalidScore field value if set, zero value otherwise.
func (o *PatchedCaptchaStageRequest) GetErrorOnInvalidScore() bool {
	if o == nil || o.ErrorOnInvalidScore == nil {
		var ret bool
		return ret
	}
	return *o.ErrorOnInvalidScore
}

// GetErrorOnInvalidScoreOk returns a tuple with the ErrorOnInvalidScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCaptchaStageRequest) GetErrorOnInvalidScoreOk() (*bool, bool) {
	if o == nil || o.ErrorOnInvalidScore == nil {
		return nil, false
	}
	return o.ErrorOnInvalidScore, true
}

// HasErrorOnInvalidScore returns a boolean if a field has been set.
func (o *PatchedCaptchaStageRequest) HasErrorOnInvalidScore() bool {
	if o != nil && o.ErrorOnInvalidScore != nil {
		return true
	}

	return false
}

// SetErrorOnInvalidScore gets a reference to the given bool and assigns it to the ErrorOnInvalidScore field.
func (o *PatchedCaptchaStageRequest) SetErrorOnInvalidScore(v bool) {
	o.ErrorOnInvalidScore = &v
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
	if o.JsUrl != nil {
		toSerialize["js_url"] = o.JsUrl
	}
	if o.ApiUrl != nil {
		toSerialize["api_url"] = o.ApiUrl
	}
	if o.Interactive != nil {
		toSerialize["interactive"] = o.Interactive
	}
	if o.ScoreMinThreshold != nil {
		toSerialize["score_min_threshold"] = o.ScoreMinThreshold
	}
	if o.ScoreMaxThreshold != nil {
		toSerialize["score_max_threshold"] = o.ScoreMaxThreshold
	}
	if o.ErrorOnInvalidScore != nil {
		toSerialize["error_on_invalid_score"] = o.ErrorOnInvalidScore
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
