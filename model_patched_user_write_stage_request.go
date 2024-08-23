/*
authentik

Making authentication simple.

API version: 2024.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedUserWriteStageRequest UserWriteStage Serializer
type PatchedUserWriteStageRequest struct {
	Name             *string               `json:"name,omitempty"`
	FlowSet          []FlowSetRequest      `json:"flow_set,omitempty"`
	UserCreationMode *UserCreationModeEnum `json:"user_creation_mode,omitempty"`
	// When set, newly created users are inactive and cannot login.
	CreateUsersAsInactive *bool `json:"create_users_as_inactive,omitempty"`
	// Optionally add newly created users to this group.
	CreateUsersGroup NullableString `json:"create_users_group,omitempty"`
	UserType         *UserTypeEnum  `json:"user_type,omitempty"`
	UserPathTemplate *string        `json:"user_path_template,omitempty"`
}

// NewPatchedUserWriteStageRequest instantiates a new PatchedUserWriteStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedUserWriteStageRequest() *PatchedUserWriteStageRequest {
	this := PatchedUserWriteStageRequest{}
	return &this
}

// NewPatchedUserWriteStageRequestWithDefaults instantiates a new PatchedUserWriteStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedUserWriteStageRequestWithDefaults() *PatchedUserWriteStageRequest {
	this := PatchedUserWriteStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedUserWriteStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserWriteStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedUserWriteStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedUserWriteStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedUserWriteStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserWriteStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedUserWriteStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedUserWriteStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetUserCreationMode returns the UserCreationMode field value if set, zero value otherwise.
func (o *PatchedUserWriteStageRequest) GetUserCreationMode() UserCreationModeEnum {
	if o == nil || o.UserCreationMode == nil {
		var ret UserCreationModeEnum
		return ret
	}
	return *o.UserCreationMode
}

// GetUserCreationModeOk returns a tuple with the UserCreationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserWriteStageRequest) GetUserCreationModeOk() (*UserCreationModeEnum, bool) {
	if o == nil || o.UserCreationMode == nil {
		return nil, false
	}
	return o.UserCreationMode, true
}

// HasUserCreationMode returns a boolean if a field has been set.
func (o *PatchedUserWriteStageRequest) HasUserCreationMode() bool {
	if o != nil && o.UserCreationMode != nil {
		return true
	}

	return false
}

// SetUserCreationMode gets a reference to the given UserCreationModeEnum and assigns it to the UserCreationMode field.
func (o *PatchedUserWriteStageRequest) SetUserCreationMode(v UserCreationModeEnum) {
	o.UserCreationMode = &v
}

// GetCreateUsersAsInactive returns the CreateUsersAsInactive field value if set, zero value otherwise.
func (o *PatchedUserWriteStageRequest) GetCreateUsersAsInactive() bool {
	if o == nil || o.CreateUsersAsInactive == nil {
		var ret bool
		return ret
	}
	return *o.CreateUsersAsInactive
}

// GetCreateUsersAsInactiveOk returns a tuple with the CreateUsersAsInactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserWriteStageRequest) GetCreateUsersAsInactiveOk() (*bool, bool) {
	if o == nil || o.CreateUsersAsInactive == nil {
		return nil, false
	}
	return o.CreateUsersAsInactive, true
}

// HasCreateUsersAsInactive returns a boolean if a field has been set.
func (o *PatchedUserWriteStageRequest) HasCreateUsersAsInactive() bool {
	if o != nil && o.CreateUsersAsInactive != nil {
		return true
	}

	return false
}

// SetCreateUsersAsInactive gets a reference to the given bool and assigns it to the CreateUsersAsInactive field.
func (o *PatchedUserWriteStageRequest) SetCreateUsersAsInactive(v bool) {
	o.CreateUsersAsInactive = &v
}

// GetCreateUsersGroup returns the CreateUsersGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedUserWriteStageRequest) GetCreateUsersGroup() string {
	if o == nil || o.CreateUsersGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreateUsersGroup.Get()
}

// GetCreateUsersGroupOk returns a tuple with the CreateUsersGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedUserWriteStageRequest) GetCreateUsersGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreateUsersGroup.Get(), o.CreateUsersGroup.IsSet()
}

// HasCreateUsersGroup returns a boolean if a field has been set.
func (o *PatchedUserWriteStageRequest) HasCreateUsersGroup() bool {
	if o != nil && o.CreateUsersGroup.IsSet() {
		return true
	}

	return false
}

// SetCreateUsersGroup gets a reference to the given NullableString and assigns it to the CreateUsersGroup field.
func (o *PatchedUserWriteStageRequest) SetCreateUsersGroup(v string) {
	o.CreateUsersGroup.Set(&v)
}

// SetCreateUsersGroupNil sets the value for CreateUsersGroup to be an explicit nil
func (o *PatchedUserWriteStageRequest) SetCreateUsersGroupNil() {
	o.CreateUsersGroup.Set(nil)
}

// UnsetCreateUsersGroup ensures that no value is present for CreateUsersGroup, not even an explicit nil
func (o *PatchedUserWriteStageRequest) UnsetCreateUsersGroup() {
	o.CreateUsersGroup.Unset()
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *PatchedUserWriteStageRequest) GetUserType() UserTypeEnum {
	if o == nil || o.UserType == nil {
		var ret UserTypeEnum
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserWriteStageRequest) GetUserTypeOk() (*UserTypeEnum, bool) {
	if o == nil || o.UserType == nil {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *PatchedUserWriteStageRequest) HasUserType() bool {
	if o != nil && o.UserType != nil {
		return true
	}

	return false
}

// SetUserType gets a reference to the given UserTypeEnum and assigns it to the UserType field.
func (o *PatchedUserWriteStageRequest) SetUserType(v UserTypeEnum) {
	o.UserType = &v
}

// GetUserPathTemplate returns the UserPathTemplate field value if set, zero value otherwise.
func (o *PatchedUserWriteStageRequest) GetUserPathTemplate() string {
	if o == nil || o.UserPathTemplate == nil {
		var ret string
		return ret
	}
	return *o.UserPathTemplate
}

// GetUserPathTemplateOk returns a tuple with the UserPathTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserWriteStageRequest) GetUserPathTemplateOk() (*string, bool) {
	if o == nil || o.UserPathTemplate == nil {
		return nil, false
	}
	return o.UserPathTemplate, true
}

// HasUserPathTemplate returns a boolean if a field has been set.
func (o *PatchedUserWriteStageRequest) HasUserPathTemplate() bool {
	if o != nil && o.UserPathTemplate != nil {
		return true
	}

	return false
}

// SetUserPathTemplate gets a reference to the given string and assigns it to the UserPathTemplate field.
func (o *PatchedUserWriteStageRequest) SetUserPathTemplate(v string) {
	o.UserPathTemplate = &v
}

func (o PatchedUserWriteStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.UserCreationMode != nil {
		toSerialize["user_creation_mode"] = o.UserCreationMode
	}
	if o.CreateUsersAsInactive != nil {
		toSerialize["create_users_as_inactive"] = o.CreateUsersAsInactive
	}
	if o.CreateUsersGroup.IsSet() {
		toSerialize["create_users_group"] = o.CreateUsersGroup.Get()
	}
	if o.UserType != nil {
		toSerialize["user_type"] = o.UserType
	}
	if o.UserPathTemplate != nil {
		toSerialize["user_path_template"] = o.UserPathTemplate
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedUserWriteStageRequest struct {
	value *PatchedUserWriteStageRequest
	isSet bool
}

func (v NullablePatchedUserWriteStageRequest) Get() *PatchedUserWriteStageRequest {
	return v.value
}

func (v *NullablePatchedUserWriteStageRequest) Set(val *PatchedUserWriteStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedUserWriteStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedUserWriteStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedUserWriteStageRequest(val *PatchedUserWriteStageRequest) *NullablePatchedUserWriteStageRequest {
	return &NullablePatchedUserWriteStageRequest{value: val, isSet: true}
}

func (v NullablePatchedUserWriteStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedUserWriteStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
