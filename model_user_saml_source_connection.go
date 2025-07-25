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
	"time"
)

// UserSAMLSourceConnection User source connection
type UserSAMLSourceConnection struct {
	Pk          int32     `json:"pk"`
	User        int32     `json:"user"`
	Source      string    `json:"source"`
	SourceObj   Source    `json:"source_obj"`
	Identifier  string    `json:"identifier"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
}

// NewUserSAMLSourceConnection instantiates a new UserSAMLSourceConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSAMLSourceConnection(pk int32, user int32, source string, sourceObj Source, identifier string, created time.Time, lastUpdated time.Time) *UserSAMLSourceConnection {
	this := UserSAMLSourceConnection{}
	this.Pk = pk
	this.User = user
	this.Source = source
	this.SourceObj = sourceObj
	this.Identifier = identifier
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewUserSAMLSourceConnectionWithDefaults instantiates a new UserSAMLSourceConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSAMLSourceConnectionWithDefaults() *UserSAMLSourceConnection {
	this := UserSAMLSourceConnection{}
	return &this
}

// GetPk returns the Pk field value
func (o *UserSAMLSourceConnection) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *UserSAMLSourceConnection) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *UserSAMLSourceConnection) SetPk(v int32) {
	o.Pk = v
}

// GetUser returns the User field value
func (o *UserSAMLSourceConnection) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserSAMLSourceConnection) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserSAMLSourceConnection) SetUser(v int32) {
	o.User = v
}

// GetSource returns the Source field value
func (o *UserSAMLSourceConnection) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *UserSAMLSourceConnection) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *UserSAMLSourceConnection) SetSource(v string) {
	o.Source = v
}

// GetSourceObj returns the SourceObj field value
func (o *UserSAMLSourceConnection) GetSourceObj() Source {
	if o == nil {
		var ret Source
		return ret
	}

	return o.SourceObj
}

// GetSourceObjOk returns a tuple with the SourceObj field value
// and a boolean to check if the value has been set.
func (o *UserSAMLSourceConnection) GetSourceObjOk() (*Source, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceObj, true
}

// SetSourceObj sets field value
func (o *UserSAMLSourceConnection) SetSourceObj(v Source) {
	o.SourceObj = v
}

// GetIdentifier returns the Identifier field value
func (o *UserSAMLSourceConnection) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *UserSAMLSourceConnection) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *UserSAMLSourceConnection) SetIdentifier(v string) {
	o.Identifier = v
}

// GetCreated returns the Created field value
func (o *UserSAMLSourceConnection) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *UserSAMLSourceConnection) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *UserSAMLSourceConnection) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *UserSAMLSourceConnection) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *UserSAMLSourceConnection) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *UserSAMLSourceConnection) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

func (o UserSAMLSourceConnection) MarshalJSON() ([]byte, error) {
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
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableUserSAMLSourceConnection struct {
	value *UserSAMLSourceConnection
	isSet bool
}

func (v NullableUserSAMLSourceConnection) Get() *UserSAMLSourceConnection {
	return v.value
}

func (v *NullableUserSAMLSourceConnection) Set(val *UserSAMLSourceConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSAMLSourceConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSAMLSourceConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSAMLSourceConnection(val *UserSAMLSourceConnection) *NullableUserSAMLSourceConnection {
	return &NullableUserSAMLSourceConnection{value: val, isSet: true}
}

func (v NullableUserSAMLSourceConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSAMLSourceConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
