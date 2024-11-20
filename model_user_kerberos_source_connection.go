/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the UserKerberosSourceConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserKerberosSourceConnection{}

// UserKerberosSourceConnection Kerberos Source Serializer
type UserKerberosSourceConnection struct {
	Pk         int32     `json:"pk"`
	User       int32     `json:"user"`
	Source     Source    `json:"source"`
	Created    time.Time `json:"created"`
	Identifier string    `json:"identifier"`
}

type _UserKerberosSourceConnection UserKerberosSourceConnection

// NewUserKerberosSourceConnection instantiates a new UserKerberosSourceConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserKerberosSourceConnection(pk int32, user int32, source Source, created time.Time, identifier string) *UserKerberosSourceConnection {
	this := UserKerberosSourceConnection{}
	this.Pk = pk
	this.User = user
	this.Source = source
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
func (o *UserKerberosSourceConnection) GetSource() Source {
	if o == nil {
		var ret Source
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *UserKerberosSourceConnection) GetSourceOk() (*Source, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *UserKerberosSourceConnection) SetSource(v Source) {
	o.Source = v
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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserKerberosSourceConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pk"] = o.Pk
	toSerialize["user"] = o.User
	toSerialize["source"] = o.Source
	toSerialize["created"] = o.Created
	toSerialize["identifier"] = o.Identifier
	return toSerialize, nil
}

func (o *UserKerberosSourceConnection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pk",
		"user",
		"source",
		"created",
		"identifier",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUserKerberosSourceConnection := _UserKerberosSourceConnection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserKerberosSourceConnection)

	if err != nil {
		return err
	}

	*o = UserKerberosSourceConnection(varUserKerberosSourceConnection)

	return err
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
