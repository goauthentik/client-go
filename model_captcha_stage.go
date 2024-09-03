/*
authentik

Making authentication simple.

API version: 2024.8.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CaptchaStage CaptchaStage Serializer
type CaptchaStage struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// Get object type so that we know how to edit the object
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName string    `json:"meta_model_name"`
	FlowSet       []FlowSet `json:"flow_set,omitempty"`
	// Public key, acquired your captcha Provider.
	PublicKey         string   `json:"public_key"`
	JsUrl             *string  `json:"js_url,omitempty"`
	ApiUrl            *string  `json:"api_url,omitempty"`
	ScoreMinThreshold *float64 `json:"score_min_threshold,omitempty"`
	ScoreMaxThreshold *float64 `json:"score_max_threshold,omitempty"`
	// When enabled and the received captcha score is outside of the given threshold, the stage will show an error message. When not enabled, the flow will continue, but the data from the captcha will be available in the context for policy decisions
	ErrorOnInvalidScore *bool `json:"error_on_invalid_score,omitempty"`
}

// NewCaptchaStage instantiates a new CaptchaStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptchaStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, publicKey string) *CaptchaStage {
	this := CaptchaStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.PublicKey = publicKey
	return &this
}

// NewCaptchaStageWithDefaults instantiates a new CaptchaStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptchaStageWithDefaults() *CaptchaStage {
	this := CaptchaStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *CaptchaStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *CaptchaStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *CaptchaStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *CaptchaStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CaptchaStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CaptchaStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *CaptchaStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *CaptchaStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *CaptchaStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *CaptchaStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *CaptchaStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *CaptchaStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *CaptchaStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *CaptchaStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *CaptchaStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *CaptchaStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *CaptchaStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *CaptchaStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *CaptchaStage) GetFlowSet() []FlowSet {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSet
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaStage) GetFlowSetOk() ([]FlowSet, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *CaptchaStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSet and assigns it to the FlowSet field.
func (o *CaptchaStage) SetFlowSet(v []FlowSet) {
	o.FlowSet = v
}

// GetPublicKey returns the PublicKey field value
func (o *CaptchaStage) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *CaptchaStage) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *CaptchaStage) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetJsUrl returns the JsUrl field value if set, zero value otherwise.
func (o *CaptchaStage) GetJsUrl() string {
	if o == nil || o.JsUrl == nil {
		var ret string
		return ret
	}
	return *o.JsUrl
}

// GetJsUrlOk returns a tuple with the JsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaStage) GetJsUrlOk() (*string, bool) {
	if o == nil || o.JsUrl == nil {
		return nil, false
	}
	return o.JsUrl, true
}

// HasJsUrl returns a boolean if a field has been set.
func (o *CaptchaStage) HasJsUrl() bool {
	if o != nil && o.JsUrl != nil {
		return true
	}

	return false
}

// SetJsUrl gets a reference to the given string and assigns it to the JsUrl field.
func (o *CaptchaStage) SetJsUrl(v string) {
	o.JsUrl = &v
}

// GetApiUrl returns the ApiUrl field value if set, zero value otherwise.
func (o *CaptchaStage) GetApiUrl() string {
	if o == nil || o.ApiUrl == nil {
		var ret string
		return ret
	}
	return *o.ApiUrl
}

// GetApiUrlOk returns a tuple with the ApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaStage) GetApiUrlOk() (*string, bool) {
	if o == nil || o.ApiUrl == nil {
		return nil, false
	}
	return o.ApiUrl, true
}

// HasApiUrl returns a boolean if a field has been set.
func (o *CaptchaStage) HasApiUrl() bool {
	if o != nil && o.ApiUrl != nil {
		return true
	}

	return false
}

// SetApiUrl gets a reference to the given string and assigns it to the ApiUrl field.
func (o *CaptchaStage) SetApiUrl(v string) {
	o.ApiUrl = &v
}

// GetScoreMinThreshold returns the ScoreMinThreshold field value if set, zero value otherwise.
func (o *CaptchaStage) GetScoreMinThreshold() float64 {
	if o == nil || o.ScoreMinThreshold == nil {
		var ret float64
		return ret
	}
	return *o.ScoreMinThreshold
}

// GetScoreMinThresholdOk returns a tuple with the ScoreMinThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaStage) GetScoreMinThresholdOk() (*float64, bool) {
	if o == nil || o.ScoreMinThreshold == nil {
		return nil, false
	}
	return o.ScoreMinThreshold, true
}

// HasScoreMinThreshold returns a boolean if a field has been set.
func (o *CaptchaStage) HasScoreMinThreshold() bool {
	if o != nil && o.ScoreMinThreshold != nil {
		return true
	}

	return false
}

// SetScoreMinThreshold gets a reference to the given float64 and assigns it to the ScoreMinThreshold field.
func (o *CaptchaStage) SetScoreMinThreshold(v float64) {
	o.ScoreMinThreshold = &v
}

// GetScoreMaxThreshold returns the ScoreMaxThreshold field value if set, zero value otherwise.
func (o *CaptchaStage) GetScoreMaxThreshold() float64 {
	if o == nil || o.ScoreMaxThreshold == nil {
		var ret float64
		return ret
	}
	return *o.ScoreMaxThreshold
}

// GetScoreMaxThresholdOk returns a tuple with the ScoreMaxThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaStage) GetScoreMaxThresholdOk() (*float64, bool) {
	if o == nil || o.ScoreMaxThreshold == nil {
		return nil, false
	}
	return o.ScoreMaxThreshold, true
}

// HasScoreMaxThreshold returns a boolean if a field has been set.
func (o *CaptchaStage) HasScoreMaxThreshold() bool {
	if o != nil && o.ScoreMaxThreshold != nil {
		return true
	}

	return false
}

// SetScoreMaxThreshold gets a reference to the given float64 and assigns it to the ScoreMaxThreshold field.
func (o *CaptchaStage) SetScoreMaxThreshold(v float64) {
	o.ScoreMaxThreshold = &v
}

// GetErrorOnInvalidScore returns the ErrorOnInvalidScore field value if set, zero value otherwise.
func (o *CaptchaStage) GetErrorOnInvalidScore() bool {
	if o == nil || o.ErrorOnInvalidScore == nil {
		var ret bool
		return ret
	}
	return *o.ErrorOnInvalidScore
}

// GetErrorOnInvalidScoreOk returns a tuple with the ErrorOnInvalidScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaStage) GetErrorOnInvalidScoreOk() (*bool, bool) {
	if o == nil || o.ErrorOnInvalidScore == nil {
		return nil, false
	}
	return o.ErrorOnInvalidScore, true
}

// HasErrorOnInvalidScore returns a boolean if a field has been set.
func (o *CaptchaStage) HasErrorOnInvalidScore() bool {
	if o != nil && o.ErrorOnInvalidScore != nil {
		return true
	}

	return false
}

// SetErrorOnInvalidScore gets a reference to the given bool and assigns it to the ErrorOnInvalidScore field.
func (o *CaptchaStage) SetErrorOnInvalidScore(v bool) {
	o.ErrorOnInvalidScore = &v
}

func (o CaptchaStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["verbose_name"] = o.VerboseName
	}
	if true {
		toSerialize["verbose_name_plural"] = o.VerboseNamePlural
	}
	if true {
		toSerialize["meta_model_name"] = o.MetaModelName
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if true {
		toSerialize["public_key"] = o.PublicKey
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

type NullableCaptchaStage struct {
	value *CaptchaStage
	isSet bool
}

func (v NullableCaptchaStage) Get() *CaptchaStage {
	return v.value
}

func (v *NullableCaptchaStage) Set(val *CaptchaStage) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptchaStage) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptchaStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptchaStage(val *CaptchaStage) *NullableCaptchaStage {
	return &NullableCaptchaStage{value: val, isSet: true}
}

func (v NullableCaptchaStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptchaStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
