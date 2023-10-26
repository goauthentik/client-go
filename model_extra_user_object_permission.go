/*
authentik

Making authentication simple.

API version: 2023.10.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ExtraUserObjectPermission User permission with additional object-related data
type ExtraUserObjectPermission struct {
	Id       int32  `json:"id"`
	Codename string `json:"codename"`
	Model    string `json:"model"`
	AppLabel string `json:"app_label"`
	ObjectPk string `json:"object_pk"`
	Name     string `json:"name"`
	// Get app label from permission's model
	AppLabelVerbose string `json:"app_label_verbose"`
	// Get model label from permission's model
	ModelVerbose string `json:"model_verbose"`
	// Get model description from attached model. This operation takes at least one additional query, and the description is only shown if the user/role has the view_ permission on the object
	ObjectDescription NullableString `json:"object_description"`
}

// NewExtraUserObjectPermission instantiates a new ExtraUserObjectPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtraUserObjectPermission(id int32, codename string, model string, appLabel string, objectPk string, name string, appLabelVerbose string, modelVerbose string, objectDescription NullableString) *ExtraUserObjectPermission {
	this := ExtraUserObjectPermission{}
	this.Id = id
	this.Codename = codename
	this.Model = model
	this.AppLabel = appLabel
	this.ObjectPk = objectPk
	this.Name = name
	this.AppLabelVerbose = appLabelVerbose
	this.ModelVerbose = modelVerbose
	this.ObjectDescription = objectDescription
	return &this
}

// NewExtraUserObjectPermissionWithDefaults instantiates a new ExtraUserObjectPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtraUserObjectPermissionWithDefaults() *ExtraUserObjectPermission {
	this := ExtraUserObjectPermission{}
	return &this
}

// GetId returns the Id field value
func (o *ExtraUserObjectPermission) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExtraUserObjectPermission) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExtraUserObjectPermission) SetId(v int32) {
	o.Id = v
}

// GetCodename returns the Codename field value
func (o *ExtraUserObjectPermission) GetCodename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Codename
}

// GetCodenameOk returns a tuple with the Codename field value
// and a boolean to check if the value has been set.
func (o *ExtraUserObjectPermission) GetCodenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Codename, true
}

// SetCodename sets field value
func (o *ExtraUserObjectPermission) SetCodename(v string) {
	o.Codename = v
}

// GetModel returns the Model field value
func (o *ExtraUserObjectPermission) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *ExtraUserObjectPermission) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *ExtraUserObjectPermission) SetModel(v string) {
	o.Model = v
}

// GetAppLabel returns the AppLabel field value
func (o *ExtraUserObjectPermission) GetAppLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppLabel
}

// GetAppLabelOk returns a tuple with the AppLabel field value
// and a boolean to check if the value has been set.
func (o *ExtraUserObjectPermission) GetAppLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppLabel, true
}

// SetAppLabel sets field value
func (o *ExtraUserObjectPermission) SetAppLabel(v string) {
	o.AppLabel = v
}

// GetObjectPk returns the ObjectPk field value
func (o *ExtraUserObjectPermission) GetObjectPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectPk
}

// GetObjectPkOk returns a tuple with the ObjectPk field value
// and a boolean to check if the value has been set.
func (o *ExtraUserObjectPermission) GetObjectPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectPk, true
}

// SetObjectPk sets field value
func (o *ExtraUserObjectPermission) SetObjectPk(v string) {
	o.ObjectPk = v
}

// GetName returns the Name field value
func (o *ExtraUserObjectPermission) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExtraUserObjectPermission) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExtraUserObjectPermission) SetName(v string) {
	o.Name = v
}

// GetAppLabelVerbose returns the AppLabelVerbose field value
func (o *ExtraUserObjectPermission) GetAppLabelVerbose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppLabelVerbose
}

// GetAppLabelVerboseOk returns a tuple with the AppLabelVerbose field value
// and a boolean to check if the value has been set.
func (o *ExtraUserObjectPermission) GetAppLabelVerboseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppLabelVerbose, true
}

// SetAppLabelVerbose sets field value
func (o *ExtraUserObjectPermission) SetAppLabelVerbose(v string) {
	o.AppLabelVerbose = v
}

// GetModelVerbose returns the ModelVerbose field value
func (o *ExtraUserObjectPermission) GetModelVerbose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelVerbose
}

// GetModelVerboseOk returns a tuple with the ModelVerbose field value
// and a boolean to check if the value has been set.
func (o *ExtraUserObjectPermission) GetModelVerboseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelVerbose, true
}

// SetModelVerbose sets field value
func (o *ExtraUserObjectPermission) SetModelVerbose(v string) {
	o.ModelVerbose = v
}

// GetObjectDescription returns the ObjectDescription field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ExtraUserObjectPermission) GetObjectDescription() string {
	if o == nil || o.ObjectDescription.Get() == nil {
		var ret string
		return ret
	}

	return *o.ObjectDescription.Get()
}

// GetObjectDescriptionOk returns a tuple with the ObjectDescription field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtraUserObjectPermission) GetObjectDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectDescription.Get(), o.ObjectDescription.IsSet()
}

// SetObjectDescription sets field value
func (o *ExtraUserObjectPermission) SetObjectDescription(v string) {
	o.ObjectDescription.Set(&v)
}

func (o ExtraUserObjectPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
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
		toSerialize["object_pk"] = o.ObjectPk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["app_label_verbose"] = o.AppLabelVerbose
	}
	if true {
		toSerialize["model_verbose"] = o.ModelVerbose
	}
	if true {
		toSerialize["object_description"] = o.ObjectDescription.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableExtraUserObjectPermission struct {
	value *ExtraUserObjectPermission
	isSet bool
}

func (v NullableExtraUserObjectPermission) Get() *ExtraUserObjectPermission {
	return v.value
}

func (v *NullableExtraUserObjectPermission) Set(val *ExtraUserObjectPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableExtraUserObjectPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableExtraUserObjectPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtraUserObjectPermission(val *ExtraUserObjectPermission) *NullableExtraUserObjectPermission {
	return &NullableExtraUserObjectPermission{value: val, isSet: true}
}

func (v NullableExtraUserObjectPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtraUserObjectPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
