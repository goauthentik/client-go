/*
authentik

Making authentication simple.

API version: 2023.3.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the FlowSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowSet{}

// FlowSet Stripped down flow serializer
type FlowSet struct {
	Pk                      string `json:"pk"`
	PolicybindingmodelPtrId string `json:"policybindingmodel_ptr_id"`
	Name                    string `json:"name"`
	// Visible in the URL.
	Slug string `json:"slug"`
	// Shown as the Title in Flow pages.
	Title       string              `json:"title"`
	Designation FlowDesignationEnum `json:"designation"`
	// Get the URL to the background image. If the name is /static or starts with http it is returned as-is
	Background       string            `json:"background"`
	PolicyEngineMode *PolicyEngineMode `json:"policy_engine_mode,omitempty"`
	// Enable compatibility mode, increases compatibility with password managers on mobile devices.
	CompatibilityMode *bool `json:"compatibility_mode,omitempty"`
	// Get export URL for flow
	ExportUrl    string            `json:"export_url"`
	Layout       *LayoutEnum       `json:"layout,omitempty"`
	DeniedAction *DeniedActionEnum `json:"denied_action,omitempty"`
}

// NewFlowSet instantiates a new FlowSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowSet(pk string, policybindingmodelPtrId string, name string, slug string, title string, designation FlowDesignationEnum, background string, exportUrl string) *FlowSet {
	this := FlowSet{}
	this.Pk = pk
	this.PolicybindingmodelPtrId = policybindingmodelPtrId
	this.Name = name
	this.Slug = slug
	this.Title = title
	this.Designation = designation
	this.Background = background
	this.ExportUrl = exportUrl
	return &this
}

// NewFlowSetWithDefaults instantiates a new FlowSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowSetWithDefaults() *FlowSet {
	this := FlowSet{}
	return &this
}

// GetPk returns the Pk field value
func (o *FlowSet) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *FlowSet) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *FlowSet) SetPk(v string) {
	o.Pk = v
}

// GetPolicybindingmodelPtrId returns the PolicybindingmodelPtrId field value
func (o *FlowSet) GetPolicybindingmodelPtrId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicybindingmodelPtrId
}

// GetPolicybindingmodelPtrIdOk returns a tuple with the PolicybindingmodelPtrId field value
// and a boolean to check if the value has been set.
func (o *FlowSet) GetPolicybindingmodelPtrIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicybindingmodelPtrId, true
}

// SetPolicybindingmodelPtrId sets field value
func (o *FlowSet) SetPolicybindingmodelPtrId(v string) {
	o.PolicybindingmodelPtrId = v
}

// GetName returns the Name field value
func (o *FlowSet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FlowSet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FlowSet) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *FlowSet) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *FlowSet) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *FlowSet) SetSlug(v string) {
	o.Slug = v
}

// GetTitle returns the Title field value
func (o *FlowSet) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *FlowSet) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *FlowSet) SetTitle(v string) {
	o.Title = v
}

// GetDesignation returns the Designation field value
func (o *FlowSet) GetDesignation() FlowDesignationEnum {
	if o == nil {
		var ret FlowDesignationEnum
		return ret
	}

	return o.Designation
}

// GetDesignationOk returns a tuple with the Designation field value
// and a boolean to check if the value has been set.
func (o *FlowSet) GetDesignationOk() (*FlowDesignationEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Designation, true
}

// SetDesignation sets field value
func (o *FlowSet) SetDesignation(v FlowDesignationEnum) {
	o.Designation = v
}

// GetBackground returns the Background field value
func (o *FlowSet) GetBackground() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Background
}

// GetBackgroundOk returns a tuple with the Background field value
// and a boolean to check if the value has been set.
func (o *FlowSet) GetBackgroundOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Background, true
}

// SetBackground sets field value
func (o *FlowSet) SetBackground(v string) {
	o.Background = v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *FlowSet) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || IsNil(o.PolicyEngineMode) {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowSet) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || IsNil(o.PolicyEngineMode) {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *FlowSet) HasPolicyEngineMode() bool {
	if o != nil && !IsNil(o.PolicyEngineMode) {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *FlowSet) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetCompatibilityMode returns the CompatibilityMode field value if set, zero value otherwise.
func (o *FlowSet) GetCompatibilityMode() bool {
	if o == nil || IsNil(o.CompatibilityMode) {
		var ret bool
		return ret
	}
	return *o.CompatibilityMode
}

// GetCompatibilityModeOk returns a tuple with the CompatibilityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowSet) GetCompatibilityModeOk() (*bool, bool) {
	if o == nil || IsNil(o.CompatibilityMode) {
		return nil, false
	}
	return o.CompatibilityMode, true
}

// HasCompatibilityMode returns a boolean if a field has been set.
func (o *FlowSet) HasCompatibilityMode() bool {
	if o != nil && !IsNil(o.CompatibilityMode) {
		return true
	}

	return false
}

// SetCompatibilityMode gets a reference to the given bool and assigns it to the CompatibilityMode field.
func (o *FlowSet) SetCompatibilityMode(v bool) {
	o.CompatibilityMode = &v
}

// GetExportUrl returns the ExportUrl field value
func (o *FlowSet) GetExportUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExportUrl
}

// GetExportUrlOk returns a tuple with the ExportUrl field value
// and a boolean to check if the value has been set.
func (o *FlowSet) GetExportUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportUrl, true
}

// SetExportUrl sets field value
func (o *FlowSet) SetExportUrl(v string) {
	o.ExportUrl = v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *FlowSet) GetLayout() LayoutEnum {
	if o == nil || IsNil(o.Layout) {
		var ret LayoutEnum
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowSet) GetLayoutOk() (*LayoutEnum, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *FlowSet) HasLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given LayoutEnum and assigns it to the Layout field.
func (o *FlowSet) SetLayout(v LayoutEnum) {
	o.Layout = &v
}

// GetDeniedAction returns the DeniedAction field value if set, zero value otherwise.
func (o *FlowSet) GetDeniedAction() DeniedActionEnum {
	if o == nil || IsNil(o.DeniedAction) {
		var ret DeniedActionEnum
		return ret
	}
	return *o.DeniedAction
}

// GetDeniedActionOk returns a tuple with the DeniedAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowSet) GetDeniedActionOk() (*DeniedActionEnum, bool) {
	if o == nil || IsNil(o.DeniedAction) {
		return nil, false
	}
	return o.DeniedAction, true
}

// HasDeniedAction returns a boolean if a field has been set.
func (o *FlowSet) HasDeniedAction() bool {
	if o != nil && !IsNil(o.DeniedAction) {
		return true
	}

	return false
}

// SetDeniedAction gets a reference to the given DeniedActionEnum and assigns it to the DeniedAction field.
func (o *FlowSet) SetDeniedAction(v DeniedActionEnum) {
	o.DeniedAction = &v
}

func (o FlowSet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pk is readOnly
	// skip: policybindingmodel_ptr_id is readOnly
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	toSerialize["title"] = o.Title
	toSerialize["designation"] = o.Designation
	// skip: background is readOnly
	if !IsNil(o.PolicyEngineMode) {
		toSerialize["policy_engine_mode"] = o.PolicyEngineMode
	}
	if !IsNil(o.CompatibilityMode) {
		toSerialize["compatibility_mode"] = o.CompatibilityMode
	}
	// skip: export_url is readOnly
	if !IsNil(o.Layout) {
		toSerialize["layout"] = o.Layout
	}
	if !IsNil(o.DeniedAction) {
		toSerialize["denied_action"] = o.DeniedAction
	}
	return toSerialize, nil
}

type NullableFlowSet struct {
	value *FlowSet
	isSet bool
}

func (v NullableFlowSet) Get() *FlowSet {
	return v.value
}

func (v *NullableFlowSet) Set(val *FlowSet) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowSet) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowSet(val *FlowSet) *NullableFlowSet {
	return &NullableFlowSet{value: val, isSet: true}
}

func (v NullableFlowSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
