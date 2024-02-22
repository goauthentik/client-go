/*
authentik

Making authentication simple.

API version: 2024.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedBlueprintInstanceRequest Info about a single blueprint instance file
type PatchedBlueprintInstanceRequest struct {
	Name    *string     `json:"name,omitempty"`
	Path    *string     `json:"path,omitempty"`
	Context interface{} `json:"context,omitempty"`
	Enabled *bool       `json:"enabled,omitempty"`
	Content *string     `json:"content,omitempty"`
}

// NewPatchedBlueprintInstanceRequest instantiates a new PatchedBlueprintInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBlueprintInstanceRequest() *PatchedBlueprintInstanceRequest {
	this := PatchedBlueprintInstanceRequest{}
	var path string = ""
	this.Path = &path
	return &this
}

// NewPatchedBlueprintInstanceRequestWithDefaults instantiates a new PatchedBlueprintInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBlueprintInstanceRequestWithDefaults() *PatchedBlueprintInstanceRequest {
	this := PatchedBlueprintInstanceRequest{}
	var path string = ""
	this.Path = &path
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedBlueprintInstanceRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBlueprintInstanceRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedBlueprintInstanceRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedBlueprintInstanceRequest) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PatchedBlueprintInstanceRequest) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBlueprintInstanceRequest) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PatchedBlueprintInstanceRequest) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *PatchedBlueprintInstanceRequest) SetPath(v string) {
	o.Path = &v
}

// GetContext returns the Context field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBlueprintInstanceRequest) GetContext() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBlueprintInstanceRequest) GetContextOk() (*interface{}, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return &o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *PatchedBlueprintInstanceRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given interface{} and assigns it to the Context field.
func (o *PatchedBlueprintInstanceRequest) SetContext(v interface{}) {
	o.Context = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchedBlueprintInstanceRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBlueprintInstanceRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchedBlueprintInstanceRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchedBlueprintInstanceRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *PatchedBlueprintInstanceRequest) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBlueprintInstanceRequest) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *PatchedBlueprintInstanceRequest) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *PatchedBlueprintInstanceRequest) SetContent(v string) {
	o.Content = &v
}

func (o PatchedBlueprintInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedBlueprintInstanceRequest struct {
	value *PatchedBlueprintInstanceRequest
	isSet bool
}

func (v NullablePatchedBlueprintInstanceRequest) Get() *PatchedBlueprintInstanceRequest {
	return v.value
}

func (v *NullablePatchedBlueprintInstanceRequest) Set(val *PatchedBlueprintInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBlueprintInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBlueprintInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBlueprintInstanceRequest(val *PatchedBlueprintInstanceRequest) *NullablePatchedBlueprintInstanceRequest {
	return &NullablePatchedBlueprintInstanceRequest{value: val, isSet: true}
}

func (v NullablePatchedBlueprintInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBlueprintInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
