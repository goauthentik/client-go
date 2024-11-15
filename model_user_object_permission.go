/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserObjectPermission User-bound object level permission
type UserObjectPermission struct {
	Id       int32  `json:"id"`
	Codename string `json:"codename"`
	Model    string `json:"model"`
	AppLabel string `json:"app_label"`
	ObjectPk string `json:"object_pk"`
	Name     string `json:"name"`
}

// NewUserObjectPermission instantiates a new UserObjectPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserObjectPermission(id int32, codename string, model string, appLabel string, objectPk string, name string) *UserObjectPermission {
	this := UserObjectPermission{}
	this.Id = id
	this.Codename = codename
	this.Model = model
	this.AppLabel = appLabel
	this.ObjectPk = objectPk
	this.Name = name
	return &this
}

// NewUserObjectPermissionWithDefaults instantiates a new UserObjectPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserObjectPermissionWithDefaults() *UserObjectPermission {
	this := UserObjectPermission{}
	return &this
}

// GetId returns the Id field value
func (o *UserObjectPermission) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserObjectPermission) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserObjectPermission) SetId(v int32) {
	o.Id = v
}

// GetCodename returns the Codename field value
func (o *UserObjectPermission) GetCodename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Codename
}

// GetCodenameOk returns a tuple with the Codename field value
// and a boolean to check if the value has been set.
func (o *UserObjectPermission) GetCodenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Codename, true
}

// SetCodename sets field value
func (o *UserObjectPermission) SetCodename(v string) {
	o.Codename = v
}

// GetModel returns the Model field value
func (o *UserObjectPermission) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *UserObjectPermission) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *UserObjectPermission) SetModel(v string) {
	o.Model = v
}

// GetAppLabel returns the AppLabel field value
func (o *UserObjectPermission) GetAppLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppLabel
}

// GetAppLabelOk returns a tuple with the AppLabel field value
// and a boolean to check if the value has been set.
func (o *UserObjectPermission) GetAppLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppLabel, true
}

// SetAppLabel sets field value
func (o *UserObjectPermission) SetAppLabel(v string) {
	o.AppLabel = v
}

// GetObjectPk returns the ObjectPk field value
func (o *UserObjectPermission) GetObjectPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectPk
}

// GetObjectPkOk returns a tuple with the ObjectPk field value
// and a boolean to check if the value has been set.
func (o *UserObjectPermission) GetObjectPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectPk, true
}

// SetObjectPk sets field value
func (o *UserObjectPermission) SetObjectPk(v string) {
	o.ObjectPk = v
}

// GetName returns the Name field value
func (o *UserObjectPermission) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserObjectPermission) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserObjectPermission) SetName(v string) {
	o.Name = v
}

func (o UserObjectPermission) MarshalJSON() ([]byte, error) {
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

type NullableUserObjectPermission struct {
	value *UserObjectPermission
	isSet bool
}

func (v NullableUserObjectPermission) Get() *UserObjectPermission {
	return v.value
}

func (v *NullableUserObjectPermission) Set(val *UserObjectPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableUserObjectPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableUserObjectPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserObjectPermission(val *UserObjectPermission) *NullableUserObjectPermission {
	return &NullableUserObjectPermission{value: val, isSet: true}
}

func (v NullableUserObjectPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserObjectPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
