/*
authentik

Making authentication simple.

API version: 2024.12.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PropertyMappingTestRequest Test property mapping execution for a user/group with context
type PropertyMappingTestRequest struct {
	User    NullableInt32          `json:"user,omitempty"`
	Context map[string]interface{} `json:"context,omitempty"`
	Group   NullableString         `json:"group,omitempty"`
}

// NewPropertyMappingTestRequest instantiates a new PropertyMappingTestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyMappingTestRequest() *PropertyMappingTestRequest {
	this := PropertyMappingTestRequest{}
	return &this
}

// NewPropertyMappingTestRequestWithDefaults instantiates a new PropertyMappingTestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyMappingTestRequestWithDefaults() *PropertyMappingTestRequest {
	this := PropertyMappingTestRequest{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PropertyMappingTestRequest) GetUser() int32 {
	if o == nil || o.User.Get() == nil {
		var ret int32
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PropertyMappingTestRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *PropertyMappingTestRequest) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableInt32 and assigns it to the User field.
func (o *PropertyMappingTestRequest) SetUser(v int32) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *PropertyMappingTestRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *PropertyMappingTestRequest) UnsetUser() {
	o.User.Unset()
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *PropertyMappingTestRequest) GetContext() map[string]interface{} {
	if o == nil || o.Context == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyMappingTestRequest) GetContextOk() (map[string]interface{}, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *PropertyMappingTestRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *PropertyMappingTestRequest) SetContext(v map[string]interface{}) {
	o.Context = v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PropertyMappingTestRequest) GetGroup() string {
	if o == nil || o.Group.Get() == nil {
		var ret string
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PropertyMappingTestRequest) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *PropertyMappingTestRequest) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableString and assigns it to the Group field.
func (o *PropertyMappingTestRequest) SetGroup(v string) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil
func (o *PropertyMappingTestRequest) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *PropertyMappingTestRequest) UnsetGroup() {
	o.Group.Unset()
}

func (o PropertyMappingTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyMappingTestRequest struct {
	value *PropertyMappingTestRequest
	isSet bool
}

func (v NullablePropertyMappingTestRequest) Get() *PropertyMappingTestRequest {
	return v.value
}

func (v *NullablePropertyMappingTestRequest) Set(val *PropertyMappingTestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyMappingTestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyMappingTestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyMappingTestRequest(val *PropertyMappingTestRequest) *NullablePropertyMappingTestRequest {
	return &NullablePropertyMappingTestRequest{value: val, isSet: true}
}

func (v NullablePropertyMappingTestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyMappingTestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
