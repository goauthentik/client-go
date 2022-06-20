/*
authentik

Making authentication simple.

API version: 2022.6.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// Reputation Reputation Serializer
type Reputation struct {
	Pk         *string                `json:"pk,omitempty"`
	Identifier string                 `json:"identifier"`
	Ip         string                 `json:"ip"`
	IpGeoData  map[string]interface{} `json:"ip_geo_data,omitempty"`
	Score      *int64                 `json:"score,omitempty"`
	Updated    time.Time              `json:"updated"`
}

// NewReputation instantiates a new Reputation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReputation(identifier string, ip string, updated time.Time) *Reputation {
	this := Reputation{}
	this.Identifier = identifier
	this.Ip = ip
	this.Updated = updated
	return &this
}

// NewReputationWithDefaults instantiates a new Reputation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReputationWithDefaults() *Reputation {
	this := Reputation{}
	return &this
}

// GetPk returns the Pk field value if set, zero value otherwise.
func (o *Reputation) GetPk() string {
	if o == nil || o.Pk == nil {
		var ret string
		return ret
	}
	return *o.Pk
}

// GetPkOk returns a tuple with the Pk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reputation) GetPkOk() (*string, bool) {
	if o == nil || o.Pk == nil {
		return nil, false
	}
	return o.Pk, true
}

// HasPk returns a boolean if a field has been set.
func (o *Reputation) HasPk() bool {
	if o != nil && o.Pk != nil {
		return true
	}

	return false
}

// SetPk gets a reference to the given string and assigns it to the Pk field.
func (o *Reputation) SetPk(v string) {
	o.Pk = &v
}

// GetIdentifier returns the Identifier field value
func (o *Reputation) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *Reputation) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *Reputation) SetIdentifier(v string) {
	o.Identifier = v
}

// GetIp returns the Ip field value
func (o *Reputation) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *Reputation) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *Reputation) SetIp(v string) {
	o.Ip = v
}

// GetIpGeoData returns the IpGeoData field value if set, zero value otherwise.
func (o *Reputation) GetIpGeoData() map[string]interface{} {
	if o == nil || o.IpGeoData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.IpGeoData
}

// GetIpGeoDataOk returns a tuple with the IpGeoData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reputation) GetIpGeoDataOk() (map[string]interface{}, bool) {
	if o == nil || o.IpGeoData == nil {
		return nil, false
	}
	return o.IpGeoData, true
}

// HasIpGeoData returns a boolean if a field has been set.
func (o *Reputation) HasIpGeoData() bool {
	if o != nil && o.IpGeoData != nil {
		return true
	}

	return false
}

// SetIpGeoData gets a reference to the given map[string]interface{} and assigns it to the IpGeoData field.
func (o *Reputation) SetIpGeoData(v map[string]interface{}) {
	o.IpGeoData = v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *Reputation) GetScore() int64 {
	if o == nil || o.Score == nil {
		var ret int64
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reputation) GetScoreOk() (*int64, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *Reputation) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given int64 and assigns it to the Score field.
func (o *Reputation) SetScore(v int64) {
	o.Score = &v
}

// GetUpdated returns the Updated field value
func (o *Reputation) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *Reputation) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *Reputation) SetUpdated(v time.Time) {
	o.Updated = v
}

func (o Reputation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pk != nil {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["ip"] = o.Ip
	}
	if o.IpGeoData != nil {
		toSerialize["ip_geo_data"] = o.IpGeoData
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if true {
		toSerialize["updated"] = o.Updated
	}
	return json.Marshal(toSerialize)
}

type NullableReputation struct {
	value *Reputation
	isSet bool
}

func (v NullableReputation) Get() *Reputation {
	return v.value
}

func (v *NullableReputation) Set(val *Reputation) {
	v.value = val
	v.isSet = true
}

func (v NullableReputation) IsSet() bool {
	return v.isSet
}

func (v *NullableReputation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReputation(val *Reputation) *NullableReputation {
	return &NullableReputation{value: val, isSet: true}
}

func (v NullableReputation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReputation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
