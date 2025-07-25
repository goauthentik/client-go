/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// FlowSetRequest Stripped down flow serializer
type FlowSetRequest struct {
	Name string `json:"name"`
	// Visible in the URL.
	Slug string `json:"slug"`
	// Shown as the Title in Flow pages.
	Title string `json:"title"`
	// Decides what this Flow is used for. For example, the Authentication flow is redirect to when an un-authenticated user visits authentik.
	Designation      FlowDesignationEnum `json:"designation"`
	PolicyEngineMode *PolicyEngineMode   `json:"policy_engine_mode,omitempty"`
	// Enable compatibility mode, increases compatibility with password managers on mobile devices.
	CompatibilityMode *bool           `json:"compatibility_mode,omitempty"`
	Layout            *FlowLayoutEnum `json:"layout,omitempty"`
	// Configure what should happen when a flow denies access to a user.
	DeniedAction *DeniedActionEnum `json:"denied_action,omitempty"`
}

// NewFlowSetRequest instantiates a new FlowSetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowSetRequest(name string, slug string, title string, designation FlowDesignationEnum) *FlowSetRequest {
	this := FlowSetRequest{}
	this.Name = name
	this.Slug = slug
	this.Title = title
	this.Designation = designation
	return &this
}

// NewFlowSetRequestWithDefaults instantiates a new FlowSetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowSetRequestWithDefaults() *FlowSetRequest {
	this := FlowSetRequest{}
	return &this
}

// GetName returns the Name field value
func (o *FlowSetRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FlowSetRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FlowSetRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *FlowSetRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *FlowSetRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *FlowSetRequest) SetSlug(v string) {
	o.Slug = v
}

// GetTitle returns the Title field value
func (o *FlowSetRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *FlowSetRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *FlowSetRequest) SetTitle(v string) {
	o.Title = v
}

// GetDesignation returns the Designation field value
func (o *FlowSetRequest) GetDesignation() FlowDesignationEnum {
	if o == nil {
		var ret FlowDesignationEnum
		return ret
	}

	return o.Designation
}

// GetDesignationOk returns a tuple with the Designation field value
// and a boolean to check if the value has been set.
func (o *FlowSetRequest) GetDesignationOk() (*FlowDesignationEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Designation, true
}

// SetDesignation sets field value
func (o *FlowSetRequest) SetDesignation(v FlowDesignationEnum) {
	o.Designation = v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *FlowSetRequest) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowSetRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *FlowSetRequest) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *FlowSetRequest) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetCompatibilityMode returns the CompatibilityMode field value if set, zero value otherwise.
func (o *FlowSetRequest) GetCompatibilityMode() bool {
	if o == nil || o.CompatibilityMode == nil {
		var ret bool
		return ret
	}
	return *o.CompatibilityMode
}

// GetCompatibilityModeOk returns a tuple with the CompatibilityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowSetRequest) GetCompatibilityModeOk() (*bool, bool) {
	if o == nil || o.CompatibilityMode == nil {
		return nil, false
	}
	return o.CompatibilityMode, true
}

// HasCompatibilityMode returns a boolean if a field has been set.
func (o *FlowSetRequest) HasCompatibilityMode() bool {
	if o != nil && o.CompatibilityMode != nil {
		return true
	}

	return false
}

// SetCompatibilityMode gets a reference to the given bool and assigns it to the CompatibilityMode field.
func (o *FlowSetRequest) SetCompatibilityMode(v bool) {
	o.CompatibilityMode = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *FlowSetRequest) GetLayout() FlowLayoutEnum {
	if o == nil || o.Layout == nil {
		var ret FlowLayoutEnum
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowSetRequest) GetLayoutOk() (*FlowLayoutEnum, bool) {
	if o == nil || o.Layout == nil {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *FlowSetRequest) HasLayout() bool {
	if o != nil && o.Layout != nil {
		return true
	}

	return false
}

// SetLayout gets a reference to the given FlowLayoutEnum and assigns it to the Layout field.
func (o *FlowSetRequest) SetLayout(v FlowLayoutEnum) {
	o.Layout = &v
}

// GetDeniedAction returns the DeniedAction field value if set, zero value otherwise.
func (o *FlowSetRequest) GetDeniedAction() DeniedActionEnum {
	if o == nil || o.DeniedAction == nil {
		var ret DeniedActionEnum
		return ret
	}
	return *o.DeniedAction
}

// GetDeniedActionOk returns a tuple with the DeniedAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowSetRequest) GetDeniedActionOk() (*DeniedActionEnum, bool) {
	if o == nil || o.DeniedAction == nil {
		return nil, false
	}
	return o.DeniedAction, true
}

// HasDeniedAction returns a boolean if a field has been set.
func (o *FlowSetRequest) HasDeniedAction() bool {
	if o != nil && o.DeniedAction != nil {
		return true
	}

	return false
}

// SetDeniedAction gets a reference to the given DeniedActionEnum and assigns it to the DeniedAction field.
func (o *FlowSetRequest) SetDeniedAction(v DeniedActionEnum) {
	o.DeniedAction = &v
}

func (o FlowSetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
		toSerialize["designation"] = o.Designation
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
	if o.DeniedAction != nil {
		toSerialize["denied_action"] = o.DeniedAction
	}
	return json.Marshal(toSerialize)
}

type NullableFlowSetRequest struct {
	value *FlowSetRequest
	isSet bool
}

func (v NullableFlowSetRequest) Get() *FlowSetRequest {
	return v.value
}

func (v *NullableFlowSetRequest) Set(val *FlowSetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowSetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowSetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowSetRequest(val *FlowSetRequest) *NullableFlowSetRequest {
	return &NullableFlowSetRequest{value: val, isSet: true}
}

func (v NullableFlowSetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowSetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
