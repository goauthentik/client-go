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

// PatchedIdentificationStageRequest IdentificationStage Serializer
type PatchedIdentificationStageRequest struct {
	Name    *string          `json:"name,omitempty"`
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

// NewPatchedIdentificationStageRequest instantiates a new PatchedIdentificationStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedIdentificationStageRequest() *PatchedIdentificationStageRequest {
	this := PatchedIdentificationStageRequest{}
	return &this
}

// NewPatchedIdentificationStageRequestWithDefaults instantiates a new PatchedIdentificationStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedIdentificationStageRequestWithDefaults() *PatchedIdentificationStageRequest {
	this := PatchedIdentificationStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedIdentificationStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedIdentificationStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedIdentificationStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedIdentificationStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedIdentificationStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedIdentificationStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedIdentificationStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedIdentificationStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetUserFields returns the UserFields field value if set, zero value otherwise.
func (o *PatchedIdentificationStageRequest) GetUserFields() []UserFieldsEnum {
	if o == nil || o.UserFields == nil {
		var ret []UserFieldsEnum
		return ret
	}
	return o.UserFields
}

// GetUserFieldsOk returns a tuple with the UserFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedIdentificationStageRequest) GetUserFieldsOk() ([]UserFieldsEnum, bool) {
	if o == nil || o.UserFields == nil {
		return nil, false
	}
	return o.UserFields, true
}

// HasUserFields returns a boolean if a field has been set.
func (o *PatchedIdentificationStageRequest) HasUserFields() bool {
	if o != nil && o.UserFields != nil {
		return true
	}

	return false
}

// SetUserFields gets a reference to the given []UserFieldsEnum and assigns it to the UserFields field.
func (o *PatchedIdentificationStageRequest) SetUserFields(v []UserFieldsEnum) {
	o.UserFields = v
}

// GetPasswordStage returns the PasswordStage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedIdentificationStageRequest) GetPasswordStage() string {
	if o == nil || o.PasswordStage.Get() == nil {
		var ret string
		return ret
	}
	return *o.PasswordStage.Get()
}

// GetPasswordStageOk returns a tuple with the PasswordStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedIdentificationStageRequest) GetPasswordStageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordStage.Get(), o.PasswordStage.IsSet()
}

// HasPasswordStage returns a boolean if a field has been set.
func (o *PatchedIdentificationStageRequest) HasPasswordStage() bool {
	if o != nil && o.PasswordStage.IsSet() {
		return true
	}

	return false
}

// SetPasswordStage gets a reference to the given NullableString and assigns it to the PasswordStage field.
func (o *PatchedIdentificationStageRequest) SetPasswordStage(v string) {
	o.PasswordStage.Set(&v)
}

// SetPasswordStageNil sets the value for PasswordStage to be an explicit nil
func (o *PatchedIdentificationStageRequest) SetPasswordStageNil() {
	o.PasswordStage.Set(nil)
}

// UnsetPasswordStage ensures that no value is present for PasswordStage, not even an explicit nil
func (o *PatchedIdentificationStageRequest) UnsetPasswordStage() {
	o.PasswordStage.Unset()
}

// GetCaptchaStage returns the CaptchaStage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedIdentificationStageRequest) GetCaptchaStage() string {
	if o == nil || o.CaptchaStage.Get() == nil {
		var ret string
		return ret
	}
	return *o.CaptchaStage.Get()
}

// GetCaptchaStageOk returns a tuple with the CaptchaStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedIdentificationStageRequest) GetCaptchaStageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaptchaStage.Get(), o.CaptchaStage.IsSet()
}

// HasCaptchaStage returns a boolean if a field has been set.
func (o *PatchedIdentificationStageRequest) HasCaptchaStage() bool {
	if o != nil && o.CaptchaStage.IsSet() {
		return true
	}

	return false
}

// SetCaptchaStage gets a reference to the given NullableString and assigns it to the CaptchaStage field.
func (o *PatchedIdentificationStageRequest) SetCaptchaStage(v string) {
	o.CaptchaStage.Set(&v)
}

// SetCaptchaStageNil sets the value for CaptchaStage to be an explicit nil
func (o *PatchedIdentificationStageRequest) SetCaptchaStageNil() {
	o.CaptchaStage.Set(nil)
}

// UnsetCaptchaStage ensures that no value is present for CaptchaStage, not even an explicit nil
func (o *PatchedIdentificationStageRequest) UnsetCaptchaStage() {
	o.CaptchaStage.Unset()
}

// GetCaseInsensitiveMatching returns the CaseInsensitiveMatching field value if set, zero value otherwise.
func (o *PatchedIdentificationStageRequest) GetCaseInsensitiveMatching() bool {
	if o == nil || o.CaseInsensitiveMatching == nil {
		var ret bool
		return ret
	}
	return *o.CaseInsensitiveMatching
}

// GetCaseInsensitiveMatchingOk returns a tuple with the CaseInsensitiveMatching field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedIdentificationStageRequest) GetCaseInsensitiveMatchingOk() (*bool, bool) {
	if o == nil || o.CaseInsensitiveMatching == nil {
		return nil, false
	}
	return o.CaseInsensitiveMatching, true
}

// HasCaseInsensitiveMatching returns a boolean if a field has been set.
func (o *PatchedIdentificationStageRequest) HasCaseInsensitiveMatching() bool {
	if o != nil && o.CaseInsensitiveMatching != nil {
		return true
	}

	return false
}

// SetCaseInsensitiveMatching gets a reference to the given bool and assigns it to the CaseInsensitiveMatching field.
func (o *PatchedIdentificationStageRequest) SetCaseInsensitiveMatching(v bool) {
	o.CaseInsensitiveMatching = &v
}

// GetShowMatchedUser returns the ShowMatchedUser field value if set, zero value otherwise.
func (o *PatchedIdentificationStageRequest) GetShowMatchedUser() bool {
	if o == nil || o.ShowMatchedUser == nil {
		var ret bool
		return ret
	}
	return *o.ShowMatchedUser
}

// GetShowMatchedUserOk returns a tuple with the ShowMatchedUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedIdentificationStageRequest) GetShowMatchedUserOk() (*bool, bool) {
	if o == nil || o.ShowMatchedUser == nil {
		return nil, false
	}
	return o.ShowMatchedUser, true
}

// HasShowMatchedUser returns a boolean if a field has been set.
func (o *PatchedIdentificationStageRequest) HasShowMatchedUser() bool {
	if o != nil && o.ShowMatchedUser != nil {
		return true
	}

	return false
}

// SetShowMatchedUser gets a reference to the given bool and assigns it to the ShowMatchedUser field.
func (o *PatchedIdentificationStageRequest) SetShowMatchedUser(v bool) {
	o.ShowMatchedUser = &v
}

// GetEnrollmentFlow returns the EnrollmentFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedIdentificationStageRequest) GetEnrollmentFlow() string {
	if o == nil || o.EnrollmentFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentFlow.Get()
}

// GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedIdentificationStageRequest) GetEnrollmentFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentFlow.Get(), o.EnrollmentFlow.IsSet()
}

// HasEnrollmentFlow returns a boolean if a field has been set.
func (o *PatchedIdentificationStageRequest) HasEnrollmentFlow() bool {
	if o != nil && o.EnrollmentFlow.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentFlow gets a reference to the given NullableString and assigns it to the EnrollmentFlow field.
func (o *PatchedIdentificationStageRequest) SetEnrollmentFlow(v string) {
	o.EnrollmentFlow.Set(&v)
}

// SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil
func (o *PatchedIdentificationStageRequest) SetEnrollmentFlowNil() {
	o.EnrollmentFlow.Set(nil)
}

// UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
func (o *PatchedIdentificationStageRequest) UnsetEnrollmentFlow() {
	o.EnrollmentFlow.Unset()
}

// GetRecoveryFlow returns the RecoveryFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedIdentificationStageRequest) GetRecoveryFlow() string {
	if o == nil || o.RecoveryFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.RecoveryFlow.Get()
}

// GetRecoveryFlowOk returns a tuple with the RecoveryFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedIdentificationStageRequest) GetRecoveryFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecoveryFlow.Get(), o.RecoveryFlow.IsSet()
}

// HasRecoveryFlow returns a boolean if a field has been set.
func (o *PatchedIdentificationStageRequest) HasRecoveryFlow() bool {
	if o != nil && o.RecoveryFlow.IsSet() {
		return true
	}

	return false
}

// SetRecoveryFlow gets a reference to the given NullableString and assigns it to the RecoveryFlow field.
func (o *PatchedIdentificationStageRequest) SetRecoveryFlow(v string) {
	o.RecoveryFlow.Set(&v)
}

// SetRecoveryFlowNil sets the value for RecoveryFlow to be an explicit nil
func (o *PatchedIdentificationStageRequest) SetRecoveryFlowNil() {
	o.RecoveryFlow.Set(nil)
}

// UnsetRecoveryFlow ensures that no value is present for RecoveryFlow, not even an explicit nil
func (o *PatchedIdentificationStageRequest) UnsetRecoveryFlow() {
	o.RecoveryFlow.Unset()
}

// GetPasswordlessFlow returns the PasswordlessFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedIdentificationStageRequest) GetPasswordlessFlow() string {
	if o == nil || o.PasswordlessFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.PasswordlessFlow.Get()
}

// GetPasswordlessFlowOk returns a tuple with the PasswordlessFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedIdentificationStageRequest) GetPasswordlessFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordlessFlow.Get(), o.PasswordlessFlow.IsSet()
}

// HasPasswordlessFlow returns a boolean if a field has been set.
func (o *PatchedIdentificationStageRequest) HasPasswordlessFlow() bool {
	if o != nil && o.PasswordlessFlow.IsSet() {
		return true
	}

	return false
}

// SetPasswordlessFlow gets a reference to the given NullableString and assigns it to the PasswordlessFlow field.
func (o *PatchedIdentificationStageRequest) SetPasswordlessFlow(v string) {
	o.PasswordlessFlow.Set(&v)
}

// SetPasswordlessFlowNil sets the value for PasswordlessFlow to be an explicit nil
func (o *PatchedIdentificationStageRequest) SetPasswordlessFlowNil() {
	o.PasswordlessFlow.Set(nil)
}

// UnsetPasswordlessFlow ensures that no value is present for PasswordlessFlow, not even an explicit nil
func (o *PatchedIdentificationStageRequest) UnsetPasswordlessFlow() {
	o.PasswordlessFlow.Unset()
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *PatchedIdentificationStageRequest) GetSources() []string {
	if o == nil || o.Sources == nil {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedIdentificationStageRequest) GetSourcesOk() ([]string, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *PatchedIdentificationStageRequest) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given []string and assigns it to the Sources field.
func (o *PatchedIdentificationStageRequest) SetSources(v []string) {
	o.Sources = v
}

// GetShowSourceLabels returns the ShowSourceLabels field value if set, zero value otherwise.
func (o *PatchedIdentificationStageRequest) GetShowSourceLabels() bool {
	if o == nil || o.ShowSourceLabels == nil {
		var ret bool
		return ret
	}
	return *o.ShowSourceLabels
}

// GetShowSourceLabelsOk returns a tuple with the ShowSourceLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedIdentificationStageRequest) GetShowSourceLabelsOk() (*bool, bool) {
	if o == nil || o.ShowSourceLabels == nil {
		return nil, false
	}
	return o.ShowSourceLabels, true
}

// HasShowSourceLabels returns a boolean if a field has been set.
func (o *PatchedIdentificationStageRequest) HasShowSourceLabels() bool {
	if o != nil && o.ShowSourceLabels != nil {
		return true
	}

	return false
}

// SetShowSourceLabels gets a reference to the given bool and assigns it to the ShowSourceLabels field.
func (o *PatchedIdentificationStageRequest) SetShowSourceLabels(v bool) {
	o.ShowSourceLabels = &v
}

// GetPretendUserExists returns the PretendUserExists field value if set, zero value otherwise.
func (o *PatchedIdentificationStageRequest) GetPretendUserExists() bool {
	if o == nil || o.PretendUserExists == nil {
		var ret bool
		return ret
	}
	return *o.PretendUserExists
}

// GetPretendUserExistsOk returns a tuple with the PretendUserExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedIdentificationStageRequest) GetPretendUserExistsOk() (*bool, bool) {
	if o == nil || o.PretendUserExists == nil {
		return nil, false
	}
	return o.PretendUserExists, true
}

// HasPretendUserExists returns a boolean if a field has been set.
func (o *PatchedIdentificationStageRequest) HasPretendUserExists() bool {
	if o != nil && o.PretendUserExists != nil {
		return true
	}

	return false
}

// SetPretendUserExists gets a reference to the given bool and assigns it to the PretendUserExists field.
func (o *PatchedIdentificationStageRequest) SetPretendUserExists(v bool) {
	o.PretendUserExists = &v
}

func (o PatchedIdentificationStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.UserFields != nil {
		toSerialize["user_fields"] = o.UserFields
	}
	if o.PasswordStage.IsSet() {
		toSerialize["password_stage"] = o.PasswordStage.Get()
	}
	if o.CaptchaStage.IsSet() {
		toSerialize["captcha_stage"] = o.CaptchaStage.Get()
	}
	if o.CaseInsensitiveMatching != nil {
		toSerialize["case_insensitive_matching"] = o.CaseInsensitiveMatching
	}
	if o.ShowMatchedUser != nil {
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
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	if o.ShowSourceLabels != nil {
		toSerialize["show_source_labels"] = o.ShowSourceLabels
	}
	if o.PretendUserExists != nil {
		toSerialize["pretend_user_exists"] = o.PretendUserExists
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedIdentificationStageRequest struct {
	value *PatchedIdentificationStageRequest
	isSet bool
}

func (v NullablePatchedIdentificationStageRequest) Get() *PatchedIdentificationStageRequest {
	return v.value
}

func (v *NullablePatchedIdentificationStageRequest) Set(val *PatchedIdentificationStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedIdentificationStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedIdentificationStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedIdentificationStageRequest(val *PatchedIdentificationStageRequest) *NullablePatchedIdentificationStageRequest {
	return &NullablePatchedIdentificationStageRequest{value: val, isSet: true}
}

func (v NullablePatchedIdentificationStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedIdentificationStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
