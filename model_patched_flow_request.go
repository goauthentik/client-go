/*
authentik

Making authentication simple.

API version: 2022.11.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedFlowRequest Flow Serializer
type PatchedFlowRequest struct {
	Name *string `json:"name,omitempty"`
	// Visible in the URL.
	Slug *string `json:"slug,omitempty"`
	// Shown as the Title in Flow pages.
	Title *string `json:"title,omitempty"`
	// Decides what this Flow is used for. For example, the Authentication flow is redirect to when an un-authenticated user visits authentik.
	Designation      NullableFlowDesignationEnum `json:"designation,omitempty"`
	PolicyEngineMode *PolicyEngineMode           `json:"policy_engine_mode,omitempty"`
	// Enable compatibility mode, increases compatibility with password managers on mobile devices.
	CompatibilityMode *bool       `json:"compatibility_mode,omitempty"`
	Layout            *LayoutEnum `json:"layout,omitempty"`
	// Configure what should happen when a flow denies access to a user.
	DeniedAction NullableDeniedActionEnum `json:"denied_action,omitempty"`
}

// NewPatchedFlowRequest instantiates a new PatchedFlowRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedFlowRequest() *PatchedFlowRequest {
	this := PatchedFlowRequest{}
	return &this
}

// NewPatchedFlowRequestWithDefaults instantiates a new PatchedFlowRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedFlowRequestWithDefaults() *PatchedFlowRequest {
	this := PatchedFlowRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedFlowRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFlowRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedFlowRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedFlowRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PatchedFlowRequest) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFlowRequest) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PatchedFlowRequest) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PatchedFlowRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PatchedFlowRequest) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFlowRequest) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PatchedFlowRequest) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PatchedFlowRequest) SetTitle(v string) {
	o.Title = &v
}

// GetDesignation returns the Designation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedFlowRequest) GetDesignation() FlowDesignationEnum {
	if o == nil || o.Designation.Get() == nil {
		var ret FlowDesignationEnum
		return ret
	}
	return *o.Designation.Get()
}

// GetDesignationOk returns a tuple with the Designation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedFlowRequest) GetDesignationOk() (*FlowDesignationEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Designation.Get(), o.Designation.IsSet()
}

// HasDesignation returns a boolean if a field has been set.
func (o *PatchedFlowRequest) HasDesignation() bool {
	if o != nil && o.Designation.IsSet() {
		return true
	}

	return false
}

// SetDesignation gets a reference to the given NullableFlowDesignationEnum and assigns it to the Designation field.
func (o *PatchedFlowRequest) SetDesignation(v FlowDesignationEnum) {
	o.Designation.Set(&v)
}

// SetDesignationNil sets the value for Designation to be an explicit nil
func (o *PatchedFlowRequest) SetDesignationNil() {
	o.Designation.Set(nil)
}

// UnsetDesignation ensures that no value is present for Designation, not even an explicit nil
func (o *PatchedFlowRequest) UnsetDesignation() {
	o.Designation.Unset()
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *PatchedFlowRequest) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFlowRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *PatchedFlowRequest) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *PatchedFlowRequest) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetCompatibilityMode returns the CompatibilityMode field value if set, zero value otherwise.
func (o *PatchedFlowRequest) GetCompatibilityMode() bool {
	if o == nil || o.CompatibilityMode == nil {
		var ret bool
		return ret
	}
	return *o.CompatibilityMode
}

// GetCompatibilityModeOk returns a tuple with the CompatibilityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFlowRequest) GetCompatibilityModeOk() (*bool, bool) {
	if o == nil || o.CompatibilityMode == nil {
		return nil, false
	}
	return o.CompatibilityMode, true
}

// HasCompatibilityMode returns a boolean if a field has been set.
func (o *PatchedFlowRequest) HasCompatibilityMode() bool {
	if o != nil && o.CompatibilityMode != nil {
		return true
	}

	return false
}

// SetCompatibilityMode gets a reference to the given bool and assigns it to the CompatibilityMode field.
func (o *PatchedFlowRequest) SetCompatibilityMode(v bool) {
	o.CompatibilityMode = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *PatchedFlowRequest) GetLayout() LayoutEnum {
	if o == nil || o.Layout == nil {
		var ret LayoutEnum
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFlowRequest) GetLayoutOk() (*LayoutEnum, bool) {
	if o == nil || o.Layout == nil {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *PatchedFlowRequest) HasLayout() bool {
	if o != nil && o.Layout != nil {
		return true
	}

	return false
}

// SetLayout gets a reference to the given LayoutEnum and assigns it to the Layout field.
func (o *PatchedFlowRequest) SetLayout(v LayoutEnum) {
	o.Layout = &v
}

// GetDeniedAction returns the DeniedAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedFlowRequest) GetDeniedAction() DeniedActionEnum {
	if o == nil || o.DeniedAction.Get() == nil {
		var ret DeniedActionEnum
		return ret
	}
	return *o.DeniedAction.Get()
}

// GetDeniedActionOk returns a tuple with the DeniedAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedFlowRequest) GetDeniedActionOk() (*DeniedActionEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeniedAction.Get(), o.DeniedAction.IsSet()
}

// HasDeniedAction returns a boolean if a field has been set.
func (o *PatchedFlowRequest) HasDeniedAction() bool {
	if o != nil && o.DeniedAction.IsSet() {
		return true
	}

	return false
}

// SetDeniedAction gets a reference to the given NullableDeniedActionEnum and assigns it to the DeniedAction field.
func (o *PatchedFlowRequest) SetDeniedAction(v DeniedActionEnum) {
	o.DeniedAction.Set(&v)
}

// SetDeniedActionNil sets the value for DeniedAction to be an explicit nil
func (o *PatchedFlowRequest) SetDeniedActionNil() {
	o.DeniedAction.Set(nil)
}

// UnsetDeniedAction ensures that no value is present for DeniedAction, not even an explicit nil
func (o *PatchedFlowRequest) UnsetDeniedAction() {
	o.DeniedAction.Unset()
}

func (o PatchedFlowRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Designation.IsSet() {
		toSerialize["designation"] = o.Designation.Get()
	}
	if o.PolicyEngineMode != nil {
		toSerialize["policy_engine_mode"] = o.PolicyEngineMode
	}
	if o.CompatibilityMode != nil {
		toSerialize["compatibility_mode"] = o.CompatibilityMode
	}
	if o.Layout != nil {
		toSerialize["layout"] = o.Layout
	}
	if o.DeniedAction.IsSet() {
		toSerialize["denied_action"] = o.DeniedAction.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedFlowRequest struct {
	value *PatchedFlowRequest
	isSet bool
}

func (v NullablePatchedFlowRequest) Get() *PatchedFlowRequest {
	return v.value
}

func (v *NullablePatchedFlowRequest) Set(val *PatchedFlowRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedFlowRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedFlowRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedFlowRequest(val *PatchedFlowRequest) *NullablePatchedFlowRequest {
	return &NullablePatchedFlowRequest{value: val, isSet: true}
}

func (v NullablePatchedFlowRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedFlowRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
