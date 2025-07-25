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

// PatchedOutpostRequest Outpost Serializer
type PatchedOutpostRequest struct {
	Name      *string          `json:"name,omitempty"`
	Type      *OutpostTypeEnum `json:"type,omitempty"`
	Providers []int32          `json:"providers,omitempty"`
	// Select Service-Connection authentik should use to manage this outpost. Leave empty if authentik should not handle the deployment.
	ServiceConnection NullableString         `json:"service_connection,omitempty"`
	Config            map[string]interface{} `json:"config,omitempty"`
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed NullableString `json:"managed,omitempty"`
}

// NewPatchedOutpostRequest instantiates a new PatchedOutpostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedOutpostRequest() *PatchedOutpostRequest {
	this := PatchedOutpostRequest{}
	return &this
}

// NewPatchedOutpostRequestWithDefaults instantiates a new PatchedOutpostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedOutpostRequestWithDefaults() *PatchedOutpostRequest {
	this := PatchedOutpostRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedOutpostRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOutpostRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedOutpostRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedOutpostRequest) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedOutpostRequest) GetType() OutpostTypeEnum {
	if o == nil || o.Type == nil {
		var ret OutpostTypeEnum
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOutpostRequest) GetTypeOk() (*OutpostTypeEnum, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedOutpostRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given OutpostTypeEnum and assigns it to the Type field.
func (o *PatchedOutpostRequest) SetType(v OutpostTypeEnum) {
	o.Type = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *PatchedOutpostRequest) GetProviders() []int32 {
	if o == nil || o.Providers == nil {
		var ret []int32
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOutpostRequest) GetProvidersOk() ([]int32, bool) {
	if o == nil || o.Providers == nil {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *PatchedOutpostRequest) HasProviders() bool {
	if o != nil && o.Providers != nil {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []int32 and assigns it to the Providers field.
func (o *PatchedOutpostRequest) SetProviders(v []int32) {
	o.Providers = v
}

// GetServiceConnection returns the ServiceConnection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOutpostRequest) GetServiceConnection() string {
	if o == nil || o.ServiceConnection.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceConnection.Get()
}

// GetServiceConnectionOk returns a tuple with the ServiceConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOutpostRequest) GetServiceConnectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceConnection.Get(), o.ServiceConnection.IsSet()
}

// HasServiceConnection returns a boolean if a field has been set.
func (o *PatchedOutpostRequest) HasServiceConnection() bool {
	if o != nil && o.ServiceConnection.IsSet() {
		return true
	}

	return false
}

// SetServiceConnection gets a reference to the given NullableString and assigns it to the ServiceConnection field.
func (o *PatchedOutpostRequest) SetServiceConnection(v string) {
	o.ServiceConnection.Set(&v)
}

// SetServiceConnectionNil sets the value for ServiceConnection to be an explicit nil
func (o *PatchedOutpostRequest) SetServiceConnectionNil() {
	o.ServiceConnection.Set(nil)
}

// UnsetServiceConnection ensures that no value is present for ServiceConnection, not even an explicit nil
func (o *PatchedOutpostRequest) UnsetServiceConnection() {
	o.ServiceConnection.Unset()
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *PatchedOutpostRequest) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOutpostRequest) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *PatchedOutpostRequest) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *PatchedOutpostRequest) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOutpostRequest) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOutpostRequest) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *PatchedOutpostRequest) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *PatchedOutpostRequest) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *PatchedOutpostRequest) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *PatchedOutpostRequest) UnsetManaged() {
	o.Managed.Unset()
}

func (o PatchedOutpostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Providers != nil {
		toSerialize["providers"] = o.Providers
	}
	if o.ServiceConnection.IsSet() {
		toSerialize["service_connection"] = o.ServiceConnection.Get()
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Managed.IsSet() {
		toSerialize["managed"] = o.Managed.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedOutpostRequest struct {
	value *PatchedOutpostRequest
	isSet bool
}

func (v NullablePatchedOutpostRequest) Get() *PatchedOutpostRequest {
	return v.value
}

func (v *NullablePatchedOutpostRequest) Set(val *PatchedOutpostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedOutpostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedOutpostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedOutpostRequest(val *PatchedOutpostRequest) *NullablePatchedOutpostRequest {
	return &NullablePatchedOutpostRequest{value: val, isSet: true}
}

func (v NullablePatchedOutpostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedOutpostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
