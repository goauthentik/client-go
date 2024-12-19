/*
authentik

Making authentication simple.

API version: 2024.12.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// UserKerberosSourceConnection Kerberos Source Serializer
type UserKerberosSourceConnection struct {
	Pk         int32     `json:"pk"`
	User       int32     `json:"user"`
	Source     string    `json:"source"`
	SourceObj  Source    `json:"source_obj"`
	Created    time.Time `json:"created"`
	Identifier string    `json:"identifier"`
}

// NewUserKerberosSourceConnection instantiates a new UserKerberosSourceConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserKerberosSourceConnection(pk int32, user int32, source string, sourceObj Source, created time.Time, identifier string) *UserKerberosSourceConnection {
	this := UserKerberosSourceConnection{}
	this.Pk = pk
	this.User = user
	this.Source = source
	this.SourceObj = sourceObj
	this.Created = created
	this.Identifier = identifier
	return &this
}

// NewUserKerberosSourceConnectionWithDefaults instantiates a new UserKerberosSourceConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserKerberosSourceConnectionWithDefaults() *UserKerberosSourceConnection {
	this := UserKerberosSourceConnection{}
	return &this
}

// GetPk returns the Pk field value
func (o *UserKerberosSourceConnection) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *UserKerberosSourceConnection) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *UserKerberosSourceConnection) SetPk(v int32) {
	o.Pk = v
}

// GetUser returns the User field value
func (o *UserKerberosSourceConnection) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserKerberosSourceConnection) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserKerberosSourceConnection) SetUser(v int32) {
	o.User = v
}

// GetSource returns the Source field value
func (o *UserKerberosSourceConnection) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *UserKerberosSourceConnection) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *UserKerberosSourceConnection) SetSource(v string) {
	o.Source = v
}

// GetSourceObj returns the SourceObj field value
func (o *UserKerberosSourceConnection) GetSourceObj() Source {
	if o == nil {
		var ret Source
		return ret
	}

	return o.SourceObj
}

// GetSourceObjOk returns a tuple with the SourceObj field value
// and a boolean to check if the value has been set.
func (o *UserKerberosSourceConnection) GetSourceObjOk() (*Source, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceObj, true
}

// SetSourceObj sets field value
func (o *UserKerberosSourceConnection) SetSourceObj(v Source) {
	o.SourceObj = v
}

// GetCreated returns the Created field value
func (o *UserKerberosSourceConnection) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *UserKerberosSourceConnection) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *UserKerberosSourceConnection) SetCreated(v time.Time) {
	o.Created = v
}

// GetIdentifier returns the Identifier field value
func (o *UserKerberosSourceConnection) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *UserKerberosSourceConnection) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *UserKerberosSourceConnection) SetIdentifier(v string) {
	o.Identifier = v
}

func (o UserKerberosSourceConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["source_obj"] = o.SourceObj
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	return json.Marshal(toSerialize)
}

type NullableUserKerberosSourceConnection struct {
	value *UserKerberosSourceConnection
	isSet bool
}

func (v NullableUserKerberosSourceConnection) Get() *UserKerberosSourceConnection {
	return v.value
}

func (v *NullableUserKerberosSourceConnection) Set(val *UserKerberosSourceConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableUserKerberosSourceConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableUserKerberosSourceConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserKerberosSourceConnection(val *UserKerberosSourceConnection) *NullableUserKerberosSourceConnection {
	return &NullableUserKerberosSourceConnection{value: val, isSet: true}
}

func (v NullableUserKerberosSourceConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserKerberosSourceConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
