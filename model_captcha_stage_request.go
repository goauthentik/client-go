/*
authentik

Making authentication simple.

API version: 2024.10.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CaptchaStageRequest CaptchaStage Serializer
type CaptchaStageRequest struct {
	Name    string           `json:"name"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Public key, acquired your captcha Provider.
	PublicKey string `json:"public_key"`
	// Private key, acquired your captcha Provider.
	PrivateKey        string   `json:"private_key"`
	JsUrl             *string  `json:"js_url,omitempty"`
	ApiUrl            *string  `json:"api_url,omitempty"`
	ScoreMinThreshold *float64 `json:"score_min_threshold,omitempty"`
	ScoreMaxThreshold *float64 `json:"score_max_threshold,omitempty"`
	// When enabled and the received captcha score is outside of the given threshold, the stage will show an error message. When not enabled, the flow will continue, but the data from the captcha will be available in the context for policy decisions
	ErrorOnInvalidScore *bool `json:"error_on_invalid_score,omitempty"`
}

// NewCaptchaStageRequest instantiates a new CaptchaStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptchaStageRequest(name string, publicKey string, privateKey string) *CaptchaStageRequest {
	this := CaptchaStageRequest{}
	this.Name = name
	this.PublicKey = publicKey
	this.PrivateKey = privateKey
	return &this
}

// NewCaptchaStageRequestWithDefaults instantiates a new CaptchaStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptchaStageRequestWithDefaults() *CaptchaStageRequest {
	this := CaptchaStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CaptchaStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CaptchaStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CaptchaStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *CaptchaStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *CaptchaStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *CaptchaStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetPublicKey returns the PublicKey field value
func (o *CaptchaStageRequest) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *CaptchaStageRequest) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *CaptchaStageRequest) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetPrivateKey returns the PrivateKey field value
func (o *CaptchaStageRequest) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *CaptchaStageRequest) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *CaptchaStageRequest) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetJsUrl returns the JsUrl field value if set, zero value otherwise.
func (o *CaptchaStageRequest) GetJsUrl() string {
	if o == nil || o.JsUrl == nil {
		var ret string
		return ret
	}
	return *o.JsUrl
}

// GetJsUrlOk returns a tuple with the JsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaStageRequest) GetJsUrlOk() (*string, bool) {
	if o == nil || o.JsUrl == nil {
		return nil, false
	}
	return o.JsUrl, true
}

// HasJsUrl returns a boolean if a field has been set.
func (o *CaptchaStageRequest) HasJsUrl() bool {
	if o != nil && o.JsUrl != nil {
		return true
	}

	return false
}

// SetJsUrl gets a reference to the given string and assigns it to the JsUrl field.
func (o *CaptchaStageRequest) SetJsUrl(v string) {
	o.JsUrl = &v
}

// GetApiUrl returns the ApiUrl field value if set, zero value otherwise.
func (o *CaptchaStageRequest) GetApiUrl() string {
	if o == nil || o.ApiUrl == nil {
		var ret string
		return ret
	}
	return *o.ApiUrl
}

// GetApiUrlOk returns a tuple with the ApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaStageRequest) GetApiUrlOk() (*string, bool) {
	if o == nil || o.ApiUrl == nil {
		return nil, false
	}
	return o.ApiUrl, true
}

// HasApiUrl returns a boolean if a field has been set.
func (o *CaptchaStageRequest) HasApiUrl() bool {
	if o != nil && o.ApiUrl != nil {
		return true
	}

	return false
}

// SetApiUrl gets a reference to the given string and assigns it to the ApiUrl field.
func (o *CaptchaStageRequest) SetApiUrl(v string) {
	o.ApiUrl = &v
}

// GetScoreMinThreshold returns the ScoreMinThreshold field value if set, zero value otherwise.
func (o *CaptchaStageRequest) GetScoreMinThreshold() float64 {
	if o == nil || o.ScoreMinThreshold == nil {
		var ret float64
		return ret
	}
	return *o.ScoreMinThreshold
}

// GetScoreMinThresholdOk returns a tuple with the ScoreMinThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaStageRequest) GetScoreMinThresholdOk() (*float64, bool) {
	if o == nil || o.ScoreMinThreshold == nil {
		return nil, false
	}
	return o.ScoreMinThreshold, true
}

// HasScoreMinThreshold returns a boolean if a field has been set.
func (o *CaptchaStageRequest) HasScoreMinThreshold() bool {
	if o != nil && o.ScoreMinThreshold != nil {
		return true
	}

	return false
}

// SetScoreMinThreshold gets a reference to the given float64 and assigns it to the ScoreMinThreshold field.
func (o *CaptchaStageRequest) SetScoreMinThreshold(v float64) {
	o.ScoreMinThreshold = &v
}

// GetScoreMaxThreshold returns the ScoreMaxThreshold field value if set, zero value otherwise.
func (o *CaptchaStageRequest) GetScoreMaxThreshold() float64 {
	if o == nil || o.ScoreMaxThreshold == nil {
		var ret float64
		return ret
	}
	return *o.ScoreMaxThreshold
}

// GetScoreMaxThresholdOk returns a tuple with the ScoreMaxThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaStageRequest) GetScoreMaxThresholdOk() (*float64, bool) {
	if o == nil || o.ScoreMaxThreshold == nil {
		return nil, false
	}
	return o.ScoreMaxThreshold, true
}

// HasScoreMaxThreshold returns a boolean if a field has been set.
func (o *CaptchaStageRequest) HasScoreMaxThreshold() bool {
	if o != nil && o.ScoreMaxThreshold != nil {
		return true
	}

	return false
}

// SetScoreMaxThreshold gets a reference to the given float64 and assigns it to the ScoreMaxThreshold field.
func (o *CaptchaStageRequest) SetScoreMaxThreshold(v float64) {
	o.ScoreMaxThreshold = &v
}

// GetErrorOnInvalidScore returns the ErrorOnInvalidScore field value if set, zero value otherwise.
func (o *CaptchaStageRequest) GetErrorOnInvalidScore() bool {
	if o == nil || o.ErrorOnInvalidScore == nil {
		var ret bool
		return ret
	}
	return *o.ErrorOnInvalidScore
}

// GetErrorOnInvalidScoreOk returns a tuple with the ErrorOnInvalidScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaStageRequest) GetErrorOnInvalidScoreOk() (*bool, bool) {
	if o == nil || o.ErrorOnInvalidScore == nil {
		return nil, false
	}
	return o.ErrorOnInvalidScore, true
}

// HasErrorOnInvalidScore returns a boolean if a field has been set.
func (o *CaptchaStageRequest) HasErrorOnInvalidScore() bool {
	if o != nil && o.ErrorOnInvalidScore != nil {
		return true
	}

	return false
}

// SetErrorOnInvalidScore gets a reference to the given bool and assigns it to the ErrorOnInvalidScore field.
func (o *CaptchaStageRequest) SetErrorOnInvalidScore(v bool) {
	o.ErrorOnInvalidScore = &v
}

func (o CaptchaStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if true {
		toSerialize["public_key"] = o.PublicKey
	}
	if true {
		toSerialize["private_key"] = o.PrivateKey
	}
	if o.JsUrl != nil {
		toSerialize["js_url"] = o.JsUrl
	}
	if o.ApiUrl != nil {
		toSerialize["api_url"] = o.ApiUrl
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

type NullableCaptchaStageRequest struct {
	value *CaptchaStageRequest
	isSet bool
}

func (v NullableCaptchaStageRequest) Get() *CaptchaStageRequest {
	return v.value
}

func (v *NullableCaptchaStageRequest) Set(val *CaptchaStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptchaStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptchaStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptchaStageRequest(val *CaptchaStageRequest) *NullableCaptchaStageRequest {
	return &NullableCaptchaStageRequest{value: val, isSet: true}
}

func (v NullableCaptchaStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptchaStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
