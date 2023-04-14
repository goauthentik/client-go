/*
authentik

Making authentication simple.

API version: 2023.4.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// FlowRequest Flow Serializer
type FlowRequest struct {
	Name string `json:"name"`
	// Visible in the URL.
	Slug string `json:"slug"`
	// Shown as the Title in Flow pages.
	Title string `json:"title"`
	// Decides what this Flow is used for. For example, the Authentication flow is redirect to when an un-authenticated user visits authentik.  * `authentication` - Authentication * `authorization` - Authorization * `invalidation` - Invalidation * `enrollment` - Enrollment * `unenrollment` - Unrenollment * `recovery` - Recovery * `stage_configuration` - Stage Configuration
	Designation      FlowDesignationEnum `json:"designation"`
	PolicyEngineMode *PolicyEngineMode   `json:"policy_engine_mode,omitempty"`
	// Enable compatibility mode, increases compatibility with password managers on mobile devices.
	CompatibilityMode *bool       `json:"compatibility_mode,omitempty"`
	Layout            *LayoutEnum `json:"layout,omitempty"`
	// Configure what should happen when a flow denies access to a user.  * `message_continue` - Message Continue * `message` - Message * `continue` - Continue
	DeniedAction *DeniedActionEnum `json:"denied_action,omitempty"`
	// Required level of authentication and authorization to access a flow.  * `none` - None * `require_authenticated` - Require Authenticated * `require_unauthenticated` - Require Unauthenticated * `require_superuser` - Require Superuser
	Authentication *AuthenticationEnum `json:"authentication,omitempty"`
}

// NewFlowRequest instantiates a new FlowRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowRequest(name string, slug string, title string, designation FlowDesignationEnum) *FlowRequest {
	this := FlowRequest{}
	this.Name = name
	this.Slug = slug
	this.Title = title
	this.Designation = designation
	return &this
}

// NewFlowRequestWithDefaults instantiates a new FlowRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowRequestWithDefaults() *FlowRequest {
	this := FlowRequest{}
	return &this
}

// GetName returns the Name field value
func (o *FlowRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FlowRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FlowRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *FlowRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *FlowRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *FlowRequest) SetSlug(v string) {
	o.Slug = v
}

// GetTitle returns the Title field value
func (o *FlowRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *FlowRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *FlowRequest) SetTitle(v string) {
	o.Title = v
}

// GetDesignation returns the Designation field value
func (o *FlowRequest) GetDesignation() FlowDesignationEnum {
	if o == nil {
		var ret FlowDesignationEnum
		return ret
	}

	return o.Designation
}

// GetDesignationOk returns a tuple with the Designation field value
// and a boolean to check if the value has been set.
func (o *FlowRequest) GetDesignationOk() (*FlowDesignationEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Designation, true
}

// SetDesignation sets field value
func (o *FlowRequest) SetDesignation(v FlowDesignationEnum) {
	o.Designation = v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *FlowRequest) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *FlowRequest) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *FlowRequest) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetCompatibilityMode returns the CompatibilityMode field value if set, zero value otherwise.
func (o *FlowRequest) GetCompatibilityMode() bool {
	if o == nil || o.CompatibilityMode == nil {
		var ret bool
		return ret
	}
	return *o.CompatibilityMode
}

// GetCompatibilityModeOk returns a tuple with the CompatibilityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowRequest) GetCompatibilityModeOk() (*bool, bool) {
	if o == nil || o.CompatibilityMode == nil {
		return nil, false
	}
	return o.CompatibilityMode, true
}

// HasCompatibilityMode returns a boolean if a field has been set.
func (o *FlowRequest) HasCompatibilityMode() bool {
	if o != nil && o.CompatibilityMode != nil {
		return true
	}

	return false
}

// SetCompatibilityMode gets a reference to the given bool and assigns it to the CompatibilityMode field.
func (o *FlowRequest) SetCompatibilityMode(v bool) {
	o.CompatibilityMode = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *FlowRequest) GetLayout() LayoutEnum {
	if o == nil || o.Layout == nil {
		var ret LayoutEnum
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowRequest) GetLayoutOk() (*LayoutEnum, bool) {
	if o == nil || o.Layout == nil {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *FlowRequest) HasLayout() bool {
	if o != nil && o.Layout != nil {
		return true
	}

	return false
}

// SetLayout gets a reference to the given LayoutEnum and assigns it to the Layout field.
func (o *FlowRequest) SetLayout(v LayoutEnum) {
	o.Layout = &v
}

// GetDeniedAction returns the DeniedAction field value if set, zero value otherwise.
func (o *FlowRequest) GetDeniedAction() DeniedActionEnum {
	if o == nil || o.DeniedAction == nil {
		var ret DeniedActionEnum
		return ret
	}
	return *o.DeniedAction
}

// GetDeniedActionOk returns a tuple with the DeniedAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowRequest) GetDeniedActionOk() (*DeniedActionEnum, bool) {
	if o == nil || o.DeniedAction == nil {
		return nil, false
	}
	return o.DeniedAction, true
}

// HasDeniedAction returns a boolean if a field has been set.
func (o *FlowRequest) HasDeniedAction() bool {
	if o != nil && o.DeniedAction != nil {
		return true
	}

	return false
}

// SetDeniedAction gets a reference to the given DeniedActionEnum and assigns it to the DeniedAction field.
func (o *FlowRequest) SetDeniedAction(v DeniedActionEnum) {
	o.DeniedAction = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *FlowRequest) GetAuthentication() AuthenticationEnum {
	if o == nil || o.Authentication == nil {
		var ret AuthenticationEnum
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowRequest) GetAuthenticationOk() (*AuthenticationEnum, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *FlowRequest) HasAuthentication() bool {
	if o != nil && o.Authentication != nil {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given AuthenticationEnum and assigns it to the Authentication field.
func (o *FlowRequest) SetAuthentication(v AuthenticationEnum) {
	o.Authentication = &v
}

func (o FlowRequest) MarshalJSON() ([]byte, error) {
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
	if o.Authentication != nil {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableFlowRequest struct {
	value *FlowRequest
	isSet bool
}

func (v NullableFlowRequest) Get() *FlowRequest {
	return v.value
}

func (v *NullableFlowRequest) Set(val *FlowRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowRequest(val *FlowRequest) *NullableFlowRequest {
	return &NullableFlowRequest{value: val, isSet: true}
}

func (v NullableFlowRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
