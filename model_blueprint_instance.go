/*
authentik

Making authentication simple.

API version: 2023.5.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// BlueprintInstance Info about a single blueprint instance file
type BlueprintInstance struct {
	Pk              string                      `json:"pk"`
	Name            string                      `json:"name"`
	Path            *string                     `json:"path,omitempty"`
	Context         map[string]interface{}      `json:"context,omitempty"`
	LastApplied     time.Time                   `json:"last_applied"`
	LastAppliedHash string                      `json:"last_applied_hash"`
	Status          BlueprintInstanceStatusEnum `json:"status"`
	Enabled         *bool                       `json:"enabled,omitempty"`
	ManagedModels   []string                    `json:"managed_models"`
	Metadata        map[string]interface{}      `json:"metadata"`
	Content         *string                     `json:"content,omitempty"`
}

// NewBlueprintInstance instantiates a new BlueprintInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintInstance(pk string, name string, lastApplied time.Time, lastAppliedHash string, status BlueprintInstanceStatusEnum, managedModels []string, metadata map[string]interface{}) *BlueprintInstance {
	this := BlueprintInstance{}
	this.Pk = pk
	this.Name = name
	var path string = ""
	this.Path = &path
	this.LastApplied = lastApplied
	this.LastAppliedHash = lastAppliedHash
	this.Status = status
	this.ManagedModels = managedModels
	this.Metadata = metadata
	return &this
}

// NewBlueprintInstanceWithDefaults instantiates a new BlueprintInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintInstanceWithDefaults() *BlueprintInstance {
	this := BlueprintInstance{}
	var path string = ""
	this.Path = &path
	return &this
}

// GetPk returns the Pk field value
func (o *BlueprintInstance) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *BlueprintInstance) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *BlueprintInstance) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *BlueprintInstance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BlueprintInstance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BlueprintInstance) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BlueprintInstance) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintInstance) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BlueprintInstance) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *BlueprintInstance) SetPath(v string) {
	o.Path = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *BlueprintInstance) GetContext() map[string]interface{} {
	if o == nil || o.Context == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintInstance) GetContextOk() (map[string]interface{}, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *BlueprintInstance) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *BlueprintInstance) SetContext(v map[string]interface{}) {
	o.Context = v
}

// GetLastApplied returns the LastApplied field value
func (o *BlueprintInstance) GetLastApplied() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastApplied
}

// GetLastAppliedOk returns a tuple with the LastApplied field value
// and a boolean to check if the value has been set.
func (o *BlueprintInstance) GetLastAppliedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastApplied, true
}

// SetLastApplied sets field value
func (o *BlueprintInstance) SetLastApplied(v time.Time) {
	o.LastApplied = v
}

// GetLastAppliedHash returns the LastAppliedHash field value
func (o *BlueprintInstance) GetLastAppliedHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastAppliedHash
}

// GetLastAppliedHashOk returns a tuple with the LastAppliedHash field value
// and a boolean to check if the value has been set.
func (o *BlueprintInstance) GetLastAppliedHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastAppliedHash, true
}

// SetLastAppliedHash sets field value
func (o *BlueprintInstance) SetLastAppliedHash(v string) {
	o.LastAppliedHash = v
}

// GetStatus returns the Status field value
func (o *BlueprintInstance) GetStatus() BlueprintInstanceStatusEnum {
	if o == nil {
		var ret BlueprintInstanceStatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BlueprintInstance) GetStatusOk() (*BlueprintInstanceStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BlueprintInstance) SetStatus(v BlueprintInstanceStatusEnum) {
	o.Status = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BlueprintInstance) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintInstance) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BlueprintInstance) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BlueprintInstance) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetManagedModels returns the ManagedModels field value
func (o *BlueprintInstance) GetManagedModels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ManagedModels
}

// GetManagedModelsOk returns a tuple with the ManagedModels field value
// and a boolean to check if the value has been set.
func (o *BlueprintInstance) GetManagedModelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManagedModels, true
}

// SetManagedModels sets field value
func (o *BlueprintInstance) SetManagedModels(v []string) {
	o.ManagedModels = v
}

// GetMetadata returns the Metadata field value
func (o *BlueprintInstance) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *BlueprintInstance) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *BlueprintInstance) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *BlueprintInstance) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintInstance) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *BlueprintInstance) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *BlueprintInstance) SetContent(v string) {
	o.Content = &v
}

func (o BlueprintInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["last_applied"] = o.LastApplied
	}
	if true {
		toSerialize["last_applied_hash"] = o.LastAppliedHash
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["managed_models"] = o.ManagedModels
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableBlueprintInstance struct {
	value *BlueprintInstance
	isSet bool
}

func (v NullableBlueprintInstance) Get() *BlueprintInstance {
	return v.value
}

func (v *NullableBlueprintInstance) Set(val *BlueprintInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintInstance(val *BlueprintInstance) *NullableBlueprintInstance {
	return &NullableBlueprintInstance{value: val, isSet: true}
}

func (v NullableBlueprintInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
