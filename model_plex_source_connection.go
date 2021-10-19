/*
authentik

Making authentication simple.

API version: 2021.10.1-rc1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PlexSourceConnection Plex Source connection Serializer
type PlexSourceConnection struct {
	Pk         int32  `json:"pk"`
	User       int32  `json:"user"`
	Source     string `json:"source"`
	Identifier string `json:"identifier"`
	PlexToken  string `json:"plex_token"`
}

// NewPlexSourceConnection instantiates a new PlexSourceConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlexSourceConnection(pk int32, user int32, source string, identifier string, plexToken string) *PlexSourceConnection {
	this := PlexSourceConnection{}
	this.Pk = pk
	this.User = user
	this.Source = source
	this.Identifier = identifier
	this.PlexToken = plexToken
	return &this
}

// NewPlexSourceConnectionWithDefaults instantiates a new PlexSourceConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlexSourceConnectionWithDefaults() *PlexSourceConnection {
	this := PlexSourceConnection{}
	return &this
}

// GetPk returns the Pk field value
func (o *PlexSourceConnection) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *PlexSourceConnection) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *PlexSourceConnection) SetPk(v int32) {
	o.Pk = v
}

// GetUser returns the User field value
func (o *PlexSourceConnection) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *PlexSourceConnection) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *PlexSourceConnection) SetUser(v int32) {
	o.User = v
}

// GetSource returns the Source field value
func (o *PlexSourceConnection) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *PlexSourceConnection) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *PlexSourceConnection) SetSource(v string) {
	o.Source = v
}

// GetIdentifier returns the Identifier field value
func (o *PlexSourceConnection) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *PlexSourceConnection) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *PlexSourceConnection) SetIdentifier(v string) {
	o.Identifier = v
}

// GetPlexToken returns the PlexToken field value
func (o *PlexSourceConnection) GetPlexToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlexToken
}

// GetPlexTokenOk returns a tuple with the PlexToken field value
// and a boolean to check if the value has been set.
func (o *PlexSourceConnection) GetPlexTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlexToken, true
}

// SetPlexToken sets field value
func (o *PlexSourceConnection) SetPlexToken(v string) {
	o.PlexToken = v
}

func (o PlexSourceConnection) MarshalJSON() ([]byte, error) {
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
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["plex_token"] = o.PlexToken
	}
	return json.Marshal(toSerialize)
}

type NullablePlexSourceConnection struct {
	value *PlexSourceConnection
	isSet bool
}

func (v NullablePlexSourceConnection) Get() *PlexSourceConnection {
	return v.value
}

func (v *NullablePlexSourceConnection) Set(val *PlexSourceConnection) {
	v.value = val
	v.isSet = true
}

func (v NullablePlexSourceConnection) IsSet() bool {
	return v.isSet
}

func (v *NullablePlexSourceConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlexSourceConnection(val *PlexSourceConnection) *NullablePlexSourceConnection {
	return &NullablePlexSourceConnection{value: val, isSet: true}
}

func (v NullablePlexSourceConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlexSourceConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
