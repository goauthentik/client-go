/*
authentik

Making authentication simple.

API version: 2022.11.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserWriteStageRequest UserWriteStage Serializer
type UserWriteStageRequest struct {
	Name    string           `json:"name"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// When set, newly created users are inactive and cannot login.
	CreateUsersAsInactive *bool `json:"create_users_as_inactive,omitempty"`
	// Optionally add newly created users to this group.
	CreateUsersGroup NullableString `json:"create_users_group,omitempty"`
	// When set, this stage can create users. If not enabled and no user is available, stage will fail.
	CanCreateUsers   *bool   `json:"can_create_users,omitempty"`
	UserPathTemplate *string `json:"user_path_template,omitempty"`
}

// NewUserWriteStageRequest instantiates a new UserWriteStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserWriteStageRequest(name string) *UserWriteStageRequest {
	this := UserWriteStageRequest{}
	this.Name = name
	return &this
}

// NewUserWriteStageRequestWithDefaults instantiates a new UserWriteStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWriteStageRequestWithDefaults() *UserWriteStageRequest {
	this := UserWriteStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UserWriteStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserWriteStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserWriteStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *UserWriteStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWriteStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *UserWriteStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *UserWriteStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetCreateUsersAsInactive returns the CreateUsersAsInactive field value if set, zero value otherwise.
func (o *UserWriteStageRequest) GetCreateUsersAsInactive() bool {
	if o == nil || o.CreateUsersAsInactive == nil {
		var ret bool
		return ret
	}
	return *o.CreateUsersAsInactive
}

// GetCreateUsersAsInactiveOk returns a tuple with the CreateUsersAsInactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWriteStageRequest) GetCreateUsersAsInactiveOk() (*bool, bool) {
	if o == nil || o.CreateUsersAsInactive == nil {
		return nil, false
	}
	return o.CreateUsersAsInactive, true
}

// HasCreateUsersAsInactive returns a boolean if a field has been set.
func (o *UserWriteStageRequest) HasCreateUsersAsInactive() bool {
	if o != nil && o.CreateUsersAsInactive != nil {
		return true
	}

	return false
}

// SetCreateUsersAsInactive gets a reference to the given bool and assigns it to the CreateUsersAsInactive field.
func (o *UserWriteStageRequest) SetCreateUsersAsInactive(v bool) {
	o.CreateUsersAsInactive = &v
}

// GetCreateUsersGroup returns the CreateUsersGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserWriteStageRequest) GetCreateUsersGroup() string {
	if o == nil || o.CreateUsersGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreateUsersGroup.Get()
}

// GetCreateUsersGroupOk returns a tuple with the CreateUsersGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserWriteStageRequest) GetCreateUsersGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreateUsersGroup.Get(), o.CreateUsersGroup.IsSet()
}

// HasCreateUsersGroup returns a boolean if a field has been set.
func (o *UserWriteStageRequest) HasCreateUsersGroup() bool {
	if o != nil && o.CreateUsersGroup.IsSet() {
		return true
	}

	return false
}

// SetCreateUsersGroup gets a reference to the given NullableString and assigns it to the CreateUsersGroup field.
func (o *UserWriteStageRequest) SetCreateUsersGroup(v string) {
	o.CreateUsersGroup.Set(&v)
}

// SetCreateUsersGroupNil sets the value for CreateUsersGroup to be an explicit nil
func (o *UserWriteStageRequest) SetCreateUsersGroupNil() {
	o.CreateUsersGroup.Set(nil)
}

// UnsetCreateUsersGroup ensures that no value is present for CreateUsersGroup, not even an explicit nil
func (o *UserWriteStageRequest) UnsetCreateUsersGroup() {
	o.CreateUsersGroup.Unset()
}

// GetCanCreateUsers returns the CanCreateUsers field value if set, zero value otherwise.
func (o *UserWriteStageRequest) GetCanCreateUsers() bool {
	if o == nil || o.CanCreateUsers == nil {
		var ret bool
		return ret
	}
	return *o.CanCreateUsers
}

// GetCanCreateUsersOk returns a tuple with the CanCreateUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWriteStageRequest) GetCanCreateUsersOk() (*bool, bool) {
	if o == nil || o.CanCreateUsers == nil {
		return nil, false
	}
	return o.CanCreateUsers, true
}

// HasCanCreateUsers returns a boolean if a field has been set.
func (o *UserWriteStageRequest) HasCanCreateUsers() bool {
	if o != nil && o.CanCreateUsers != nil {
		return true
	}

	return false
}

// SetCanCreateUsers gets a reference to the given bool and assigns it to the CanCreateUsers field.
func (o *UserWriteStageRequest) SetCanCreateUsers(v bool) {
	o.CanCreateUsers = &v
}

// GetUserPathTemplate returns the UserPathTemplate field value if set, zero value otherwise.
func (o *UserWriteStageRequest) GetUserPathTemplate() string {
	if o == nil || o.UserPathTemplate == nil {
		var ret string
		return ret
	}
	return *o.UserPathTemplate
}

// GetUserPathTemplateOk returns a tuple with the UserPathTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWriteStageRequest) GetUserPathTemplateOk() (*string, bool) {
	if o == nil || o.UserPathTemplate == nil {
		return nil, false
	}
	return o.UserPathTemplate, true
}

// HasUserPathTemplate returns a boolean if a field has been set.
func (o *UserWriteStageRequest) HasUserPathTemplate() bool {
	if o != nil && o.UserPathTemplate != nil {
		return true
	}

	return false
}

// SetUserPathTemplate gets a reference to the given string and assigns it to the UserPathTemplate field.
func (o *UserWriteStageRequest) SetUserPathTemplate(v string) {
	o.UserPathTemplate = &v
}

func (o UserWriteStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.CreateUsersAsInactive != nil {
		toSerialize["create_users_as_inactive"] = o.CreateUsersAsInactive
	}
	if o.CreateUsersGroup.IsSet() {
		toSerialize["create_users_group"] = o.CreateUsersGroup.Get()
	}
	if o.CanCreateUsers != nil {
		toSerialize["can_create_users"] = o.CanCreateUsers
	}
	if o.UserPathTemplate != nil {
		toSerialize["user_path_template"] = o.UserPathTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableUserWriteStageRequest struct {
	value *UserWriteStageRequest
	isSet bool
}

func (v NullableUserWriteStageRequest) Get() *UserWriteStageRequest {
	return v.value
}

func (v *NullableUserWriteStageRequest) Set(val *UserWriteStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserWriteStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserWriteStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserWriteStageRequest(val *UserWriteStageRequest) *NullableUserWriteStageRequest {
	return &NullableUserWriteStageRequest{value: val, isSet: true}
}

func (v NullableUserWriteStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserWriteStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
