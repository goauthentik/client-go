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

// RoleObjectPermission Role-bound object level permission
type RoleObjectPermission struct {
	Id       int32  `json:"id"`
	Codename string `json:"codename"`
	Model    string `json:"model"`
	AppLabel string `json:"app_label"`
	ObjectPk string `json:"object_pk"`
	Name     string `json:"name"`
}

// NewRoleObjectPermission instantiates a new RoleObjectPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleObjectPermission(id int32, codename string, model string, appLabel string, objectPk string, name string) *RoleObjectPermission {
	this := RoleObjectPermission{}
	this.Id = id
	this.Codename = codename
	this.Model = model
	this.AppLabel = appLabel
	this.ObjectPk = objectPk
	this.Name = name
	return &this
}

// NewRoleObjectPermissionWithDefaults instantiates a new RoleObjectPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleObjectPermissionWithDefaults() *RoleObjectPermission {
	this := RoleObjectPermission{}
	return &this
}

// GetId returns the Id field value
func (o *RoleObjectPermission) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RoleObjectPermission) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RoleObjectPermission) SetId(v int32) {
	o.Id = v
}

// GetCodename returns the Codename field value
func (o *RoleObjectPermission) GetCodename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Codename
}

// GetCodenameOk returns a tuple with the Codename field value
// and a boolean to check if the value has been set.
func (o *RoleObjectPermission) GetCodenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Codename, true
}

// SetCodename sets field value
func (o *RoleObjectPermission) SetCodename(v string) {
	o.Codename = v
}

// GetModel returns the Model field value
func (o *RoleObjectPermission) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *RoleObjectPermission) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *RoleObjectPermission) SetModel(v string) {
	o.Model = v
}

// GetAppLabel returns the AppLabel field value
func (o *RoleObjectPermission) GetAppLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppLabel
}

// GetAppLabelOk returns a tuple with the AppLabel field value
// and a boolean to check if the value has been set.
func (o *RoleObjectPermission) GetAppLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppLabel, true
}

// SetAppLabel sets field value
func (o *RoleObjectPermission) SetAppLabel(v string) {
	o.AppLabel = v
}

// GetObjectPk returns the ObjectPk field value
func (o *RoleObjectPermission) GetObjectPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectPk
}

// GetObjectPkOk returns a tuple with the ObjectPk field value
// and a boolean to check if the value has been set.
func (o *RoleObjectPermission) GetObjectPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectPk, true
}

// SetObjectPk sets field value
func (o *RoleObjectPermission) SetObjectPk(v string) {
	o.ObjectPk = v
}

// GetName returns the Name field value
func (o *RoleObjectPermission) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoleObjectPermission) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoleObjectPermission) SetName(v string) {
	o.Name = v
}

func (o RoleObjectPermission) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableRoleObjectPermission struct {
	value *RoleObjectPermission
	isSet bool
}

func (v NullableRoleObjectPermission) Get() *RoleObjectPermission {
	return v.value
}

func (v *NullableRoleObjectPermission) Set(val *RoleObjectPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleObjectPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleObjectPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleObjectPermission(val *RoleObjectPermission) *NullableRoleObjectPermission {
	return &NullableRoleObjectPermission{value: val, isSet: true}
}

func (v NullableRoleObjectPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleObjectPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
