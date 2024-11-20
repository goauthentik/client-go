/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the IdentificationStageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentificationStageRequest{}

// IdentificationStageRequest IdentificationStage Serializer
type IdentificationStageRequest struct {
	Name    string           `json:"name"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Fields of the user object to match against. (Hold shift to select multiple options)
	UserFields []UserFieldsEnum `json:"user_fields,omitempty"`
	// When set, shows a password field, instead of showing the password field as separate step.
	PasswordStage NullableString `json:"password_stage,omitempty"`
	// When set, adds functionality exactly like a Captcha stage, but baked into the Identification stage.
	CaptchaStage NullableString `json:"captcha_stage,omitempty"`
	// When enabled, user fields are matched regardless of their casing.
	CaseInsensitiveMatching *bool `json:"case_insensitive_matching,omitempty"`
	// When a valid username/email has been entered, and this option is enabled, the user's username and avatar will be shown. Otherwise, the text that the user entered will be shown
	ShowMatchedUser *bool `json:"show_matched_user,omitempty"`
	// Optional enrollment flow, which is linked at the bottom of the page.
	EnrollmentFlow NullableString `json:"enrollment_flow,omitempty"`
	// Optional recovery flow, which is linked at the bottom of the page.
	RecoveryFlow NullableString `json:"recovery_flow,omitempty"`
	// Optional passwordless flow, which is linked at the bottom of the page.
	PasswordlessFlow NullableString `json:"passwordless_flow,omitempty"`
	// Specify which sources should be shown.
	Sources          []string `json:"sources,omitempty"`
	ShowSourceLabels *bool    `json:"show_source_labels,omitempty"`
	// When enabled, the stage will succeed and continue even when incorrect user info is entered.
	PretendUserExists *bool `json:"pretend_user_exists,omitempty"`
}

type _IdentificationStageRequest IdentificationStageRequest

// NewIdentificationStageRequest instantiates a new IdentificationStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentificationStageRequest(name string) *IdentificationStageRequest {
	this := IdentificationStageRequest{}
	this.Name = name
	return &this
}

// NewIdentificationStageRequestWithDefaults instantiates a new IdentificationStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentificationStageRequestWithDefaults() *IdentificationStageRequest {
	this := IdentificationStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *IdentificationStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IdentificationStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IdentificationStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *IdentificationStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || IsNil(o.FlowSet) {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || IsNil(o.FlowSet) {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *IdentificationStageRequest) HasFlowSet() bool {
	if o != nil && !IsNil(o.FlowSet) {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *IdentificationStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetUserFields returns the UserFields field value if set, zero value otherwise.
func (o *IdentificationStageRequest) GetUserFields() []UserFieldsEnum {
	if o == nil || IsNil(o.UserFields) {
		var ret []UserFieldsEnum
		return ret
	}
	return o.UserFields
}

// GetUserFieldsOk returns a tuple with the UserFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationStageRequest) GetUserFieldsOk() ([]UserFieldsEnum, bool) {
	if o == nil || IsNil(o.UserFields) {
		return nil, false
	}
	return o.UserFields, true
}

// HasUserFields returns a boolean if a field has been set.
func (o *IdentificationStageRequest) HasUserFields() bool {
	if o != nil && !IsNil(o.UserFields) {
		return true
	}

	return false
}

// SetUserFields gets a reference to the given []UserFieldsEnum and assigns it to the UserFields field.
func (o *IdentificationStageRequest) SetUserFields(v []UserFieldsEnum) {
	o.UserFields = v
}

// GetPasswordStage returns the PasswordStage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentificationStageRequest) GetPasswordStage() string {
	if o == nil || IsNil(o.PasswordStage.Get()) {
		var ret string
		return ret
	}
	return *o.PasswordStage.Get()
}

// GetPasswordStageOk returns a tuple with the PasswordStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentificationStageRequest) GetPasswordStageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordStage.Get(), o.PasswordStage.IsSet()
}

// HasPasswordStage returns a boolean if a field has been set.
func (o *IdentificationStageRequest) HasPasswordStage() bool {
	if o != nil && o.PasswordStage.IsSet() {
		return true
	}

	return false
}

// SetPasswordStage gets a reference to the given NullableString and assigns it to the PasswordStage field.
func (o *IdentificationStageRequest) SetPasswordStage(v string) {
	o.PasswordStage.Set(&v)
}

// SetPasswordStageNil sets the value for PasswordStage to be an explicit nil
func (o *IdentificationStageRequest) SetPasswordStageNil() {
	o.PasswordStage.Set(nil)
}

// UnsetPasswordStage ensures that no value is present for PasswordStage, not even an explicit nil
func (o *IdentificationStageRequest) UnsetPasswordStage() {
	o.PasswordStage.Unset()
}

// GetCaptchaStage returns the CaptchaStage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentificationStageRequest) GetCaptchaStage() string {
	if o == nil || IsNil(o.CaptchaStage.Get()) {
		var ret string
		return ret
	}
	return *o.CaptchaStage.Get()
}

// GetCaptchaStageOk returns a tuple with the CaptchaStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentificationStageRequest) GetCaptchaStageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaptchaStage.Get(), o.CaptchaStage.IsSet()
}

// HasCaptchaStage returns a boolean if a field has been set.
func (o *IdentificationStageRequest) HasCaptchaStage() bool {
	if o != nil && o.CaptchaStage.IsSet() {
		return true
	}

	return false
}

// SetCaptchaStage gets a reference to the given NullableString and assigns it to the CaptchaStage field.
func (o *IdentificationStageRequest) SetCaptchaStage(v string) {
	o.CaptchaStage.Set(&v)
}

// SetCaptchaStageNil sets the value for CaptchaStage to be an explicit nil
func (o *IdentificationStageRequest) SetCaptchaStageNil() {
	o.CaptchaStage.Set(nil)
}

// UnsetCaptchaStage ensures that no value is present for CaptchaStage, not even an explicit nil
func (o *IdentificationStageRequest) UnsetCaptchaStage() {
	o.CaptchaStage.Unset()
}

// GetCaseInsensitiveMatching returns the CaseInsensitiveMatching field value if set, zero value otherwise.
func (o *IdentificationStageRequest) GetCaseInsensitiveMatching() bool {
	if o == nil || IsNil(o.CaseInsensitiveMatching) {
		var ret bool
		return ret
	}
	return *o.CaseInsensitiveMatching
}

// GetCaseInsensitiveMatchingOk returns a tuple with the CaseInsensitiveMatching field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationStageRequest) GetCaseInsensitiveMatchingOk() (*bool, bool) {
	if o == nil || IsNil(o.CaseInsensitiveMatching) {
		return nil, false
	}
	return o.CaseInsensitiveMatching, true
}

// HasCaseInsensitiveMatching returns a boolean if a field has been set.
func (o *IdentificationStageRequest) HasCaseInsensitiveMatching() bool {
	if o != nil && !IsNil(o.CaseInsensitiveMatching) {
		return true
	}

	return false
}

// SetCaseInsensitiveMatching gets a reference to the given bool and assigns it to the CaseInsensitiveMatching field.
func (o *IdentificationStageRequest) SetCaseInsensitiveMatching(v bool) {
	o.CaseInsensitiveMatching = &v
}

// GetShowMatchedUser returns the ShowMatchedUser field value if set, zero value otherwise.
func (o *IdentificationStageRequest) GetShowMatchedUser() bool {
	if o == nil || IsNil(o.ShowMatchedUser) {
		var ret bool
		return ret
	}
	return *o.ShowMatchedUser
}

// GetShowMatchedUserOk returns a tuple with the ShowMatchedUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationStageRequest) GetShowMatchedUserOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowMatchedUser) {
		return nil, false
	}
	return o.ShowMatchedUser, true
}

// HasShowMatchedUser returns a boolean if a field has been set.
func (o *IdentificationStageRequest) HasShowMatchedUser() bool {
	if o != nil && !IsNil(o.ShowMatchedUser) {
		return true
	}

	return false
}

// SetShowMatchedUser gets a reference to the given bool and assigns it to the ShowMatchedUser field.
func (o *IdentificationStageRequest) SetShowMatchedUser(v bool) {
	o.ShowMatchedUser = &v
}

// GetEnrollmentFlow returns the EnrollmentFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentificationStageRequest) GetEnrollmentFlow() string {
	if o == nil || IsNil(o.EnrollmentFlow.Get()) {
		var ret string
		return ret
	}
	return *o.EnrollmentFlow.Get()
}

// GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentificationStageRequest) GetEnrollmentFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentFlow.Get(), o.EnrollmentFlow.IsSet()
}

// HasEnrollmentFlow returns a boolean if a field has been set.
func (o *IdentificationStageRequest) HasEnrollmentFlow() bool {
	if o != nil && o.EnrollmentFlow.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentFlow gets a reference to the given NullableString and assigns it to the EnrollmentFlow field.
func (o *IdentificationStageRequest) SetEnrollmentFlow(v string) {
	o.EnrollmentFlow.Set(&v)
}

// SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil
func (o *IdentificationStageRequest) SetEnrollmentFlowNil() {
	o.EnrollmentFlow.Set(nil)
}

// UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
func (o *IdentificationStageRequest) UnsetEnrollmentFlow() {
	o.EnrollmentFlow.Unset()
}

// GetRecoveryFlow returns the RecoveryFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentificationStageRequest) GetRecoveryFlow() string {
	if o == nil || IsNil(o.RecoveryFlow.Get()) {
		var ret string
		return ret
	}
	return *o.RecoveryFlow.Get()
}

// GetRecoveryFlowOk returns a tuple with the RecoveryFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentificationStageRequest) GetRecoveryFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecoveryFlow.Get(), o.RecoveryFlow.IsSet()
}

// HasRecoveryFlow returns a boolean if a field has been set.
func (o *IdentificationStageRequest) HasRecoveryFlow() bool {
	if o != nil && o.RecoveryFlow.IsSet() {
		return true
	}

	return false
}

// SetRecoveryFlow gets a reference to the given NullableString and assigns it to the RecoveryFlow field.
func (o *IdentificationStageRequest) SetRecoveryFlow(v string) {
	o.RecoveryFlow.Set(&v)
}

// SetRecoveryFlowNil sets the value for RecoveryFlow to be an explicit nil
func (o *IdentificationStageRequest) SetRecoveryFlowNil() {
	o.RecoveryFlow.Set(nil)
}

// UnsetRecoveryFlow ensures that no value is present for RecoveryFlow, not even an explicit nil
func (o *IdentificationStageRequest) UnsetRecoveryFlow() {
	o.RecoveryFlow.Unset()
}

// GetPasswordlessFlow returns the PasswordlessFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentificationStageRequest) GetPasswordlessFlow() string {
	if o == nil || IsNil(o.PasswordlessFlow.Get()) {
		var ret string
		return ret
	}
	return *o.PasswordlessFlow.Get()
}

// GetPasswordlessFlowOk returns a tuple with the PasswordlessFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentificationStageRequest) GetPasswordlessFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordlessFlow.Get(), o.PasswordlessFlow.IsSet()
}

// HasPasswordlessFlow returns a boolean if a field has been set.
func (o *IdentificationStageRequest) HasPasswordlessFlow() bool {
	if o != nil && o.PasswordlessFlow.IsSet() {
		return true
	}

	return false
}

// SetPasswordlessFlow gets a reference to the given NullableString and assigns it to the PasswordlessFlow field.
func (o *IdentificationStageRequest) SetPasswordlessFlow(v string) {
	o.PasswordlessFlow.Set(&v)
}

// SetPasswordlessFlowNil sets the value for PasswordlessFlow to be an explicit nil
func (o *IdentificationStageRequest) SetPasswordlessFlowNil() {
	o.PasswordlessFlow.Set(nil)
}

// UnsetPasswordlessFlow ensures that no value is present for PasswordlessFlow, not even an explicit nil
func (o *IdentificationStageRequest) UnsetPasswordlessFlow() {
	o.PasswordlessFlow.Unset()
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *IdentificationStageRequest) GetSources() []string {
	if o == nil || IsNil(o.Sources) {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationStageRequest) GetSourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *IdentificationStageRequest) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []string and assigns it to the Sources field.
func (o *IdentificationStageRequest) SetSources(v []string) {
	o.Sources = v
}

// GetShowSourceLabels returns the ShowSourceLabels field value if set, zero value otherwise.
func (o *IdentificationStageRequest) GetShowSourceLabels() bool {
	if o == nil || IsNil(o.ShowSourceLabels) {
		var ret bool
		return ret
	}
	return *o.ShowSourceLabels
}

// GetShowSourceLabelsOk returns a tuple with the ShowSourceLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationStageRequest) GetShowSourceLabelsOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowSourceLabels) {
		return nil, false
	}
	return o.ShowSourceLabels, true
}

// HasShowSourceLabels returns a boolean if a field has been set.
func (o *IdentificationStageRequest) HasShowSourceLabels() bool {
	if o != nil && !IsNil(o.ShowSourceLabels) {
		return true
	}

	return false
}

// SetShowSourceLabels gets a reference to the given bool and assigns it to the ShowSourceLabels field.
func (o *IdentificationStageRequest) SetShowSourceLabels(v bool) {
	o.ShowSourceLabels = &v
}

// GetPretendUserExists returns the PretendUserExists field value if set, zero value otherwise.
func (o *IdentificationStageRequest) GetPretendUserExists() bool {
	if o == nil || IsNil(o.PretendUserExists) {
		var ret bool
		return ret
	}
	return *o.PretendUserExists
}

// GetPretendUserExistsOk returns a tuple with the PretendUserExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationStageRequest) GetPretendUserExistsOk() (*bool, bool) {
	if o == nil || IsNil(o.PretendUserExists) {
		return nil, false
	}
	return o.PretendUserExists, true
}

// HasPretendUserExists returns a boolean if a field has been set.
func (o *IdentificationStageRequest) HasPretendUserExists() bool {
	if o != nil && !IsNil(o.PretendUserExists) {
		return true
	}

	return false
}

// SetPretendUserExists gets a reference to the given bool and assigns it to the PretendUserExists field.
func (o *IdentificationStageRequest) SetPretendUserExists(v bool) {
	o.PretendUserExists = &v
}

func (o IdentificationStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentificationStageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.FlowSet) {
		toSerialize["flow_set"] = o.FlowSet
	}
	if !IsNil(o.UserFields) {
		toSerialize["user_fields"] = o.UserFields
	}
	if o.PasswordStage.IsSet() {
		toSerialize["password_stage"] = o.PasswordStage.Get()
	}
	if o.CaptchaStage.IsSet() {
		toSerialize["captcha_stage"] = o.CaptchaStage.Get()
	}
	if !IsNil(o.CaseInsensitiveMatching) {
		toSerialize["case_insensitive_matching"] = o.CaseInsensitiveMatching
	}
	if !IsNil(o.ShowMatchedUser) {
		toSerialize["show_matched_user"] = o.ShowMatchedUser
	}
	if o.EnrollmentFlow.IsSet() {
		toSerialize["enrollment_flow"] = o.EnrollmentFlow.Get()
	}
	if o.RecoveryFlow.IsSet() {
		toSerialize["recovery_flow"] = o.RecoveryFlow.Get()
	}
	if o.PasswordlessFlow.IsSet() {
		toSerialize["passwordless_flow"] = o.PasswordlessFlow.Get()
	}
	if !IsNil(o.Sources) {
		toSerialize["sources"] = o.Sources
	}
	if !IsNil(o.ShowSourceLabels) {
		toSerialize["show_source_labels"] = o.ShowSourceLabels
	}
	if !IsNil(o.PretendUserExists) {
		toSerialize["pretend_user_exists"] = o.PretendUserExists
	}
	return toSerialize, nil
}

func (o *IdentificationStageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varIdentificationStageRequest := _IdentificationStageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIdentificationStageRequest)

	if err != nil {
		return err
	}

	*o = IdentificationStageRequest(varIdentificationStageRequest)

	return err
}

type NullableIdentificationStageRequest struct {
	value *IdentificationStageRequest
	isSet bool
}

func (v NullableIdentificationStageRequest) Get() *IdentificationStageRequest {
	return v.value
}

func (v *NullableIdentificationStageRequest) Set(val *IdentificationStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentificationStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentificationStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentificationStageRequest(val *IdentificationStageRequest) *NullableIdentificationStageRequest {
	return &NullableIdentificationStageRequest{value: val, isSet: true}
}

func (v NullableIdentificationStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentificationStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
