/*
authentik

Making authentication simple.

API version: 2024.6.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// AuthenticatedSession AuthenticatedSession Serializer
type AuthenticatedSession struct {
	Uuid *string `json:"uuid,omitempty"`
	// Check if session is currently active session
	Current       bool                              `json:"current"`
	UserAgent     AuthenticatedSessionUserAgent     `json:"user_agent"`
	GeoIp         NullableAuthenticatedSessionGeoIp `json:"geo_ip"`
	Asn           NullableAuthenticatedSessionAsn   `json:"asn"`
	User          int32                             `json:"user"`
	LastIp        string                            `json:"last_ip"`
	LastUserAgent *string                           `json:"last_user_agent,omitempty"`
	LastUsed      time.Time                         `json:"last_used"`
	Expires       NullableTime                      `json:"expires,omitempty"`
}

// NewAuthenticatedSession instantiates a new AuthenticatedSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatedSession(current bool, userAgent AuthenticatedSessionUserAgent, geoIp NullableAuthenticatedSessionGeoIp, asn NullableAuthenticatedSessionAsn, user int32, lastIp string, lastUsed time.Time) *AuthenticatedSession {
	this := AuthenticatedSession{}
	this.Current = current
	this.UserAgent = userAgent
	this.GeoIp = geoIp
	this.Asn = asn
	this.User = user
	this.LastIp = lastIp
	this.LastUsed = lastUsed
	return &this
}

// NewAuthenticatedSessionWithDefaults instantiates a new AuthenticatedSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatedSessionWithDefaults() *AuthenticatedSession {
	this := AuthenticatedSession{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *AuthenticatedSession) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSession) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *AuthenticatedSession) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *AuthenticatedSession) SetUuid(v string) {
	o.Uuid = &v
}

// GetCurrent returns the Current field value
func (o *AuthenticatedSession) GetCurrent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSession) GetCurrentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *AuthenticatedSession) SetCurrent(v bool) {
	o.Current = v
}

// GetUserAgent returns the UserAgent field value
func (o *AuthenticatedSession) GetUserAgent() AuthenticatedSessionUserAgent {
	if o == nil {
		var ret AuthenticatedSessionUserAgent
		return ret
	}

	return o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSession) GetUserAgentOk() (*AuthenticatedSessionUserAgent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAgent, true
}

// SetUserAgent sets field value
func (o *AuthenticatedSession) SetUserAgent(v AuthenticatedSessionUserAgent) {
	o.UserAgent = v
}

// GetGeoIp returns the GeoIp field value
// If the value is explicit nil, the zero value for AuthenticatedSessionGeoIp will be returned
func (o *AuthenticatedSession) GetGeoIp() AuthenticatedSessionGeoIp {
	if o == nil || o.GeoIp.Get() == nil {
		var ret AuthenticatedSessionGeoIp
		return ret
	}

	return *o.GeoIp.Get()
}

// GetGeoIpOk returns a tuple with the GeoIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatedSession) GetGeoIpOk() (*AuthenticatedSessionGeoIp, bool) {
	if o == nil {
		return nil, false
	}
	return o.GeoIp.Get(), o.GeoIp.IsSet()
}

// SetGeoIp sets field value
func (o *AuthenticatedSession) SetGeoIp(v AuthenticatedSessionGeoIp) {
	o.GeoIp.Set(&v)
}

// GetAsn returns the Asn field value
// If the value is explicit nil, the zero value for AuthenticatedSessionAsn will be returned
func (o *AuthenticatedSession) GetAsn() AuthenticatedSessionAsn {
	if o == nil || o.Asn.Get() == nil {
		var ret AuthenticatedSessionAsn
		return ret
	}

	return *o.Asn.Get()
}

// GetAsnOk returns a tuple with the Asn field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatedSession) GetAsnOk() (*AuthenticatedSessionAsn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Asn.Get(), o.Asn.IsSet()
}

// SetAsn sets field value
func (o *AuthenticatedSession) SetAsn(v AuthenticatedSessionAsn) {
	o.Asn.Set(&v)
}

// GetUser returns the User field value
func (o *AuthenticatedSession) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSession) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *AuthenticatedSession) SetUser(v int32) {
	o.User = v
}

// GetLastIp returns the LastIp field value
func (o *AuthenticatedSession) GetLastIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastIp
}

// GetLastIpOk returns a tuple with the LastIp field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSession) GetLastIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastIp, true
}

// SetLastIp sets field value
func (o *AuthenticatedSession) SetLastIp(v string) {
	o.LastIp = v
}

// GetLastUserAgent returns the LastUserAgent field value if set, zero value otherwise.
func (o *AuthenticatedSession) GetLastUserAgent() string {
	if o == nil || o.LastUserAgent == nil {
		var ret string
		return ret
	}
	return *o.LastUserAgent
}

// GetLastUserAgentOk returns a tuple with the LastUserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSession) GetLastUserAgentOk() (*string, bool) {
	if o == nil || o.LastUserAgent == nil {
		return nil, false
	}
	return o.LastUserAgent, true
}

// HasLastUserAgent returns a boolean if a field has been set.
func (o *AuthenticatedSession) HasLastUserAgent() bool {
	if o != nil && o.LastUserAgent != nil {
		return true
	}

	return false
}

// SetLastUserAgent gets a reference to the given string and assigns it to the LastUserAgent field.
func (o *AuthenticatedSession) SetLastUserAgent(v string) {
	o.LastUserAgent = &v
}

// GetLastUsed returns the LastUsed field value
func (o *AuthenticatedSession) GetLastUsed() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUsed
}

// GetLastUsedOk returns a tuple with the LastUsed field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSession) GetLastUsedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUsed, true
}

// SetLastUsed sets field value
func (o *AuthenticatedSession) SetLastUsed(v time.Time) {
	o.LastUsed = v
}

// GetExpires returns the Expires field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatedSession) GetExpires() time.Time {
	if o == nil || o.Expires.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatedSession) GetExpiresOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// HasExpires returns a boolean if a field has been set.
func (o *AuthenticatedSession) HasExpires() bool {
	if o != nil && o.Expires.IsSet() {
		return true
	}

	return false
}

// SetExpires gets a reference to the given NullableTime and assigns it to the Expires field.
func (o *AuthenticatedSession) SetExpires(v time.Time) {
	o.Expires.Set(&v)
}

// SetExpiresNil sets the value for Expires to be an explicit nil
func (o *AuthenticatedSession) SetExpiresNil() {
	o.Expires.Set(nil)
}

// UnsetExpires ensures that no value is present for Expires, not even an explicit nil
func (o *AuthenticatedSession) UnsetExpires() {
	o.Expires.Unset()
}

func (o AuthenticatedSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["current"] = o.Current
	}
	if true {
		toSerialize["user_agent"] = o.UserAgent
	}
	if true {
		toSerialize["geo_ip"] = o.GeoIp.Get()
	}
	if true {
		toSerialize["asn"] = o.Asn.Get()
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["last_ip"] = o.LastIp
	}
	if o.LastUserAgent != nil {
		toSerialize["last_user_agent"] = o.LastUserAgent
	}
	if true {
		toSerialize["last_used"] = o.LastUsed
	}
	if o.Expires.IsSet() {
		toSerialize["expires"] = o.Expires.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatedSession struct {
	value *AuthenticatedSession
	isSet bool
}

func (v NullableAuthenticatedSession) Get() *AuthenticatedSession {
	return v.value
}

func (v *NullableAuthenticatedSession) Set(val *AuthenticatedSession) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatedSession) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatedSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatedSession(val *AuthenticatedSession) *NullableAuthenticatedSession {
	return &NullableAuthenticatedSession{value: val, isSet: true}
}

func (v NullableAuthenticatedSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatedSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
