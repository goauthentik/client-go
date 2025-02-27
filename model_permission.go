/*
authentik

Making authentication simple.

API version: 2025.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Permission Global permission
type Permission struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Codename string `json:"codename"`
	Model    string `json:"model"`
	AppLabel string `json:"app_label"`
	// Human-readable app label
	AppLabelVerbose string `json:"app_label_verbose"`
	// Human-readable model name
	ModelVerbose string `json:"model_verbose"`
}

// NewPermission instantiates a new Permission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermission(id int32, name string, codename string, model string, appLabel string, appLabelVerbose string, modelVerbose string) *Permission {
	this := Permission{}
	this.Id = id
	this.Name = name
	this.Codename = codename
	this.Model = model
	this.AppLabel = appLabel
	this.AppLabelVerbose = appLabelVerbose
	this.ModelVerbose = modelVerbose
	return &this
}

// NewPermissionWithDefaults instantiates a new Permission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionWithDefaults() *Permission {
	this := Permission{}
	return &this
}

// GetId returns the Id field value
func (o *Permission) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Permission) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Permission) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Permission) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Permission) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Permission) SetName(v string) {
	o.Name = v
}

// GetCodename returns the Codename field value
func (o *Permission) GetCodename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Codename
}

// GetCodenameOk returns a tuple with the Codename field value
// and a boolean to check if the value has been set.
func (o *Permission) GetCodenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Codename, true
}

// SetCodename sets field value
func (o *Permission) SetCodename(v string) {
	o.Codename = v
}

// GetModel returns the Model field value
func (o *Permission) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *Permission) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *Permission) SetModel(v string) {
	o.Model = v
}

// GetAppLabel returns the AppLabel field value
func (o *Permission) GetAppLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppLabel
}

// GetAppLabelOk returns a tuple with the AppLabel field value
// and a boolean to check if the value has been set.
func (o *Permission) GetAppLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppLabel, true
}

// SetAppLabel sets field value
func (o *Permission) SetAppLabel(v string) {
	o.AppLabel = v
}

// GetAppLabelVerbose returns the AppLabelVerbose field value
func (o *Permission) GetAppLabelVerbose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppLabelVerbose
}

// GetAppLabelVerboseOk returns a tuple with the AppLabelVerbose field value
// and a boolean to check if the value has been set.
func (o *Permission) GetAppLabelVerboseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppLabelVerbose, true
}

// SetAppLabelVerbose sets field value
func (o *Permission) SetAppLabelVerbose(v string) {
	o.AppLabelVerbose = v
}

// GetModelVerbose returns the ModelVerbose field value
func (o *Permission) GetModelVerbose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelVerbose
}

// GetModelVerboseOk returns a tuple with the ModelVerbose field value
// and a boolean to check if the value has been set.
func (o *Permission) GetModelVerboseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelVerbose, true
}

// SetModelVerbose sets field value
func (o *Permission) SetModelVerbose(v string) {
	o.ModelVerbose = v
}

func (o Permission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["codename"] = o.Codename
	}
	if true {
		toSerialize["model"] = o.Model
	}
	if true {
		toSerialize["app_label"] = o.AppLabel
	}
	if true {
		toSerialize["app_label_verbose"] = o.AppLabelVerbose
	}
	if true {
		toSerialize["model_verbose"] = o.ModelVerbose
	}
	return json.Marshal(toSerialize)
}

type NullablePermission struct {
	value *Permission
	isSet bool
}

func (v NullablePermission) Get() *Permission {
	return v.value
}

func (v *NullablePermission) Set(val *Permission) {
	v.value = val
	v.isSet = true
}

func (v NullablePermission) IsSet() bool {
	return v.isSet
}

func (v *NullablePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermission(val *Permission) *NullablePermission {
	return &NullablePermission{value: val, isSet: true}
}

func (v NullablePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
