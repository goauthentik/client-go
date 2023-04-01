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

// checks if the BlueprintInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlueprintInstanceRequest{}

// BlueprintInstanceRequest Info about a single blueprint instance file
type BlueprintInstanceRequest struct {
	Name    string                 `json:"name"`
	Path    *string                `json:"path,omitempty"`
	Context map[string]interface{} `json:"context,omitempty"`
	Enabled *bool                  `json:"enabled,omitempty"`
	Content *string                `json:"content,omitempty"`
}

// NewBlueprintInstanceRequest instantiates a new BlueprintInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintInstanceRequest(name string) *BlueprintInstanceRequest {
	this := BlueprintInstanceRequest{}
	this.Name = name
	var path string = ""
	this.Path = &path
	return &this
}

// NewBlueprintInstanceRequestWithDefaults instantiates a new BlueprintInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintInstanceRequestWithDefaults() *BlueprintInstanceRequest {
	this := BlueprintInstanceRequest{}
	var path string = ""
	this.Path = &path
	return &this
}

// GetName returns the Name field value
func (o *BlueprintInstanceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BlueprintInstanceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BlueprintInstanceRequest) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BlueprintInstanceRequest) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintInstanceRequest) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BlueprintInstanceRequest) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *BlueprintInstanceRequest) SetPath(v string) {
	o.Path = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *BlueprintInstanceRequest) GetContext() map[string]interface{} {
	if o == nil || IsNil(o.Context) {
		var ret map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintInstanceRequest) GetContextOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Context) {
		return map[string]interface{}{}, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *BlueprintInstanceRequest) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *BlueprintInstanceRequest) SetContext(v map[string]interface{}) {
	o.Context = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BlueprintInstanceRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintInstanceRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BlueprintInstanceRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BlueprintInstanceRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *BlueprintInstanceRequest) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintInstanceRequest) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *BlueprintInstanceRequest) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *BlueprintInstanceRequest) SetContent(v string) {
	o.Content = &v
}

func (o BlueprintInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlueprintInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

type NullableBlueprintInstanceRequest struct {
	value *BlueprintInstanceRequest
	isSet bool
}

func (v NullableBlueprintInstanceRequest) Get() *BlueprintInstanceRequest {
	return v.value
}

func (v *NullableBlueprintInstanceRequest) Set(val *BlueprintInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintInstanceRequest(val *BlueprintInstanceRequest) *NullableBlueprintInstanceRequest {
	return &NullableBlueprintInstanceRequest{value: val, isSet: true}
}

func (v NullableBlueprintInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
