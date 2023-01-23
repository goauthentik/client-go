/*
authentik

Making authentication simple.

API version: 2023.1.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// InvitationFlowObj struct for InvitationFlowObj
type InvitationFlowObj struct {
	Pk                      string `json:"pk"`
	PolicybindingmodelPtrId string `json:"policybindingmodel_ptr_id"`
	Name                    string `json:"name"`
	// Visible in the URL.
	Slug string `json:"slug"`
	// Shown as the Title in Flow pages.
	Title string `json:"title"`
	// Decides what this Flow is used for. For example, the Authentication flow is redirect to when an un-authenticated user visits authentik.
	Designation      NullableFlowDesignationEnum `json:"designation"`
	Background       string                      `json:"background"`
	Stages           []string                    `json:"stages"`
	Policies         []string                    `json:"policies"`
	CacheCount       int32                       `json:"cache_count"`
	PolicyEngineMode *PolicyEngineMode           `json:"policy_engine_mode,omitempty"`
	// Enable compatibility mode, increases compatibility with password managers on mobile devices.
	CompatibilityMode *bool       `json:"compatibility_mode,omitempty"`
	ExportUrl         string      `json:"export_url"`
	Layout            *LayoutEnum `json:"layout,omitempty"`
	// Configure what should happen when a flow denies access to a user.
	DeniedAction NullableDeniedActionEnum `json:"denied_action,omitempty"`
	// Required level of authentication and authorization to access a flow.
	Authentication NullableAuthenticationEnum `json:"authentication,omitempty"`
}

// NewInvitationFlowObj instantiates a new InvitationFlowObj object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationFlowObj(pk string, policybindingmodelPtrId string, name string, slug string, title string, designation NullableFlowDesignationEnum, background string, stages []string, policies []string, cacheCount int32, exportUrl string) *InvitationFlowObj {
	this := InvitationFlowObj{}
	this.Pk = pk
	this.PolicybindingmodelPtrId = policybindingmodelPtrId
	this.Name = name
	this.Slug = slug
	this.Title = title
	this.Designation = designation
	this.Background = background
	this.Stages = stages
	this.Policies = policies
	this.CacheCount = cacheCount
	this.ExportUrl = exportUrl
	return &this
}

// NewInvitationFlowObjWithDefaults instantiates a new InvitationFlowObj object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationFlowObjWithDefaults() *InvitationFlowObj {
	this := InvitationFlowObj{}
	return &this
}

// GetPk returns the Pk field value
func (o *InvitationFlowObj) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *InvitationFlowObj) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *InvitationFlowObj) SetPk(v string) {
	o.Pk = v
}

// GetPolicybindingmodelPtrId returns the PolicybindingmodelPtrId field value
func (o *InvitationFlowObj) GetPolicybindingmodelPtrId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicybindingmodelPtrId
}

// GetPolicybindingmodelPtrIdOk returns a tuple with the PolicybindingmodelPtrId field value
// and a boolean to check if the value has been set.
func (o *InvitationFlowObj) GetPolicybindingmodelPtrIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicybindingmodelPtrId, true
}

// SetPolicybindingmodelPtrId sets field value
func (o *InvitationFlowObj) SetPolicybindingmodelPtrId(v string) {
	o.PolicybindingmodelPtrId = v
}

// GetName returns the Name field value
func (o *InvitationFlowObj) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InvitationFlowObj) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InvitationFlowObj) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *InvitationFlowObj) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *InvitationFlowObj) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *InvitationFlowObj) SetSlug(v string) {
	o.Slug = v
}

// GetTitle returns the Title field value
func (o *InvitationFlowObj) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InvitationFlowObj) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *InvitationFlowObj) SetTitle(v string) {
	o.Title = v
}

// GetDesignation returns the Designation field value
// If the value is explicit nil, the zero value for FlowDesignationEnum will be returned
func (o *InvitationFlowObj) GetDesignation() FlowDesignationEnum {
	if o == nil || o.Designation.Get() == nil {
		var ret FlowDesignationEnum
		return ret
	}

	return *o.Designation.Get()
}

// GetDesignationOk returns a tuple with the Designation field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitationFlowObj) GetDesignationOk() (*FlowDesignationEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Designation.Get(), o.Designation.IsSet()
}

// SetDesignation sets field value
func (o *InvitationFlowObj) SetDesignation(v FlowDesignationEnum) {
	o.Designation.Set(&v)
}

// GetBackground returns the Background field value
func (o *InvitationFlowObj) GetBackground() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Background
}

// GetBackgroundOk returns a tuple with the Background field value
// and a boolean to check if the value has been set.
func (o *InvitationFlowObj) GetBackgroundOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Background, true
}

// SetBackground sets field value
func (o *InvitationFlowObj) SetBackground(v string) {
	o.Background = v
}

// GetStages returns the Stages field value
func (o *InvitationFlowObj) GetStages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value
// and a boolean to check if the value has been set.
func (o *InvitationFlowObj) GetStagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stages, true
}

// SetStages sets field value
func (o *InvitationFlowObj) SetStages(v []string) {
	o.Stages = v
}

// GetPolicies returns the Policies field value
func (o *InvitationFlowObj) GetPolicies() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value
// and a boolean to check if the value has been set.
func (o *InvitationFlowObj) GetPoliciesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Policies, true
}

// SetPolicies sets field value
func (o *InvitationFlowObj) SetPolicies(v []string) {
	o.Policies = v
}

// GetCacheCount returns the CacheCount field value
func (o *InvitationFlowObj) GetCacheCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CacheCount
}

// GetCacheCountOk returns a tuple with the CacheCount field value
// and a boolean to check if the value has been set.
func (o *InvitationFlowObj) GetCacheCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheCount, true
}

// SetCacheCount sets field value
func (o *InvitationFlowObj) SetCacheCount(v int32) {
	o.CacheCount = v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *InvitationFlowObj) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationFlowObj) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *InvitationFlowObj) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *InvitationFlowObj) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetCompatibilityMode returns the CompatibilityMode field value if set, zero value otherwise.
func (o *InvitationFlowObj) GetCompatibilityMode() bool {
	if o == nil || o.CompatibilityMode == nil {
		var ret bool
		return ret
	}
	return *o.CompatibilityMode
}

// GetCompatibilityModeOk returns a tuple with the CompatibilityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationFlowObj) GetCompatibilityModeOk() (*bool, bool) {
	if o == nil || o.CompatibilityMode == nil {
		return nil, false
	}
	return o.CompatibilityMode, true
}

// HasCompatibilityMode returns a boolean if a field has been set.
func (o *InvitationFlowObj) HasCompatibilityMode() bool {
	if o != nil && o.CompatibilityMode != nil {
		return true
	}

	return false
}

// SetCompatibilityMode gets a reference to the given bool and assigns it to the CompatibilityMode field.
func (o *InvitationFlowObj) SetCompatibilityMode(v bool) {
	o.CompatibilityMode = &v
}

// GetExportUrl returns the ExportUrl field value
func (o *InvitationFlowObj) GetExportUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExportUrl
}

// GetExportUrlOk returns a tuple with the ExportUrl field value
// and a boolean to check if the value has been set.
func (o *InvitationFlowObj) GetExportUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportUrl, true
}

// SetExportUrl sets field value
func (o *InvitationFlowObj) SetExportUrl(v string) {
	o.ExportUrl = v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *InvitationFlowObj) GetLayout() LayoutEnum {
	if o == nil || o.Layout == nil {
		var ret LayoutEnum
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationFlowObj) GetLayoutOk() (*LayoutEnum, bool) {
	if o == nil || o.Layout == nil {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *InvitationFlowObj) HasLayout() bool {
	if o != nil && o.Layout != nil {
		return true
	}

	return false
}

// SetLayout gets a reference to the given LayoutEnum and assigns it to the Layout field.
func (o *InvitationFlowObj) SetLayout(v LayoutEnum) {
	o.Layout = &v
}

// GetDeniedAction returns the DeniedAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitationFlowObj) GetDeniedAction() DeniedActionEnum {
	if o == nil || o.DeniedAction.Get() == nil {
		var ret DeniedActionEnum
		return ret
	}
	return *o.DeniedAction.Get()
}

// GetDeniedActionOk returns a tuple with the DeniedAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitationFlowObj) GetDeniedActionOk() (*DeniedActionEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeniedAction.Get(), o.DeniedAction.IsSet()
}

// HasDeniedAction returns a boolean if a field has been set.
func (o *InvitationFlowObj) HasDeniedAction() bool {
	if o != nil && o.DeniedAction.IsSet() {
		return true
	}

	return false
}

// SetDeniedAction gets a reference to the given NullableDeniedActionEnum and assigns it to the DeniedAction field.
func (o *InvitationFlowObj) SetDeniedAction(v DeniedActionEnum) {
	o.DeniedAction.Set(&v)
}

// SetDeniedActionNil sets the value for DeniedAction to be an explicit nil
func (o *InvitationFlowObj) SetDeniedActionNil() {
	o.DeniedAction.Set(nil)
}

// UnsetDeniedAction ensures that no value is present for DeniedAction, not even an explicit nil
func (o *InvitationFlowObj) UnsetDeniedAction() {
	o.DeniedAction.Unset()
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitationFlowObj) GetAuthentication() AuthenticationEnum {
	if o == nil || o.Authentication.Get() == nil {
		var ret AuthenticationEnum
		return ret
	}
	return *o.Authentication.Get()
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitationFlowObj) GetAuthenticationOk() (*AuthenticationEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Authentication.Get(), o.Authentication.IsSet()
}

// HasAuthentication returns a boolean if a field has been set.
func (o *InvitationFlowObj) HasAuthentication() bool {
	if o != nil && o.Authentication.IsSet() {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given NullableAuthenticationEnum and assigns it to the Authentication field.
func (o *InvitationFlowObj) SetAuthentication(v AuthenticationEnum) {
	o.Authentication.Set(&v)
}

// SetAuthenticationNil sets the value for Authentication to be an explicit nil
func (o *InvitationFlowObj) SetAuthenticationNil() {
	o.Authentication.Set(nil)
}

// UnsetAuthentication ensures that no value is present for Authentication, not even an explicit nil
func (o *InvitationFlowObj) UnsetAuthentication() {
	o.Authentication.Unset()
}

func (o InvitationFlowObj) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["policybindingmodel_ptr_id"] = o.PolicybindingmodelPtrId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["designation"] = o.Designation.Get()
	}
	if true {
		toSerialize["background"] = o.Background
	}
	if true {
		toSerialize["stages"] = o.Stages
	}
	if true {
		toSerialize["policies"] = o.Policies
	}
	if true {
		toSerialize["cache_count"] = o.CacheCount
	}
	if o.PolicyEngineMode != nil {
		toSerialize["policy_engine_mode"] = o.PolicyEngineMode
	}
	if o.CompatibilityMode != nil {
		toSerialize["compatibility_mode"] = o.CompatibilityMode
	}
	if true {
		toSerialize["export_url"] = o.ExportUrl
	}
	if o.Layout != nil {
		toSerialize["layout"] = o.Layout
	}
	if o.DeniedAction.IsSet() {
		toSerialize["denied_action"] = o.DeniedAction.Get()
	}
	if o.Authentication.IsSet() {
		toSerialize["authentication"] = o.Authentication.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInvitationFlowObj struct {
	value *InvitationFlowObj
	isSet bool
}

func (v NullableInvitationFlowObj) Get() *InvitationFlowObj {
	return v.value
}

func (v *NullableInvitationFlowObj) Set(val *InvitationFlowObj) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitationFlowObj) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitationFlowObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitationFlowObj(val *InvitationFlowObj) *NullableInvitationFlowObj {
	return &NullableInvitationFlowObj{value: val, isSet: true}
}

func (v NullableInvitationFlowObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitationFlowObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
