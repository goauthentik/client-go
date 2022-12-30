/*
authentik

Making authentication simple.

API version: 2022.12.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// ExpiringBaseGrantModel Serializer for BaseGrantModel and ExpiringBaseGrant
type ExpiringBaseGrantModel struct {
	Pk        int32          `json:"pk"`
	Provider  OAuth2Provider `json:"provider"`
	User      User           `json:"user"`
	IsExpired bool           `json:"is_expired"`
	Expires   *time.Time     `json:"expires,omitempty"`
	Scope     []string       `json:"scope"`
}

// NewExpiringBaseGrantModel instantiates a new ExpiringBaseGrantModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpiringBaseGrantModel(pk int32, provider OAuth2Provider, user User, isExpired bool, scope []string) *ExpiringBaseGrantModel {
	this := ExpiringBaseGrantModel{}
	this.Pk = pk
	this.Provider = provider
	this.User = user
	this.IsExpired = isExpired
	this.Scope = scope
	return &this
}

// NewExpiringBaseGrantModelWithDefaults instantiates a new ExpiringBaseGrantModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpiringBaseGrantModelWithDefaults() *ExpiringBaseGrantModel {
	this := ExpiringBaseGrantModel{}
	return &this
}

// GetPk returns the Pk field value
func (o *ExpiringBaseGrantModel) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *ExpiringBaseGrantModel) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *ExpiringBaseGrantModel) SetPk(v int32) {
	o.Pk = v
}

// GetProvider returns the Provider field value
func (o *ExpiringBaseGrantModel) GetProvider() OAuth2Provider {
	if o == nil {
		var ret OAuth2Provider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *ExpiringBaseGrantModel) GetProviderOk() (*OAuth2Provider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *ExpiringBaseGrantModel) SetProvider(v OAuth2Provider) {
	o.Provider = v
}

// GetUser returns the User field value
func (o *ExpiringBaseGrantModel) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ExpiringBaseGrantModel) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ExpiringBaseGrantModel) SetUser(v User) {
	o.User = v
}

// GetIsExpired returns the IsExpired field value
func (o *ExpiringBaseGrantModel) GetIsExpired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsExpired
}

// GetIsExpiredOk returns a tuple with the IsExpired field value
// and a boolean to check if the value has been set.
func (o *ExpiringBaseGrantModel) GetIsExpiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsExpired, true
}

// SetIsExpired sets field value
func (o *ExpiringBaseGrantModel) SetIsExpired(v bool) {
	o.IsExpired = v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *ExpiringBaseGrantModel) GetExpires() time.Time {
	if o == nil || o.Expires == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpiringBaseGrantModel) GetExpiresOk() (*time.Time, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *ExpiringBaseGrantModel) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *ExpiringBaseGrantModel) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetScope returns the Scope field value
func (o *ExpiringBaseGrantModel) GetScope() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *ExpiringBaseGrantModel) GetScopeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope, true
}

// SetScope sets field value
func (o *ExpiringBaseGrantModel) SetScope(v []string) {
	o.Scope = v
}

func (o ExpiringBaseGrantModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["is_expired"] = o.IsExpired
	}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableExpiringBaseGrantModel struct {
	value *ExpiringBaseGrantModel
	isSet bool
}

func (v NullableExpiringBaseGrantModel) Get() *ExpiringBaseGrantModel {
	return v.value
}

func (v *NullableExpiringBaseGrantModel) Set(val *ExpiringBaseGrantModel) {
	v.value = val
	v.isSet = true
}

func (v NullableExpiringBaseGrantModel) IsSet() bool {
	return v.isSet
}

func (v *NullableExpiringBaseGrantModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpiringBaseGrantModel(val *ExpiringBaseGrantModel) *NullableExpiringBaseGrantModel {
	return &NullableExpiringBaseGrantModel{value: val, isSet: true}
}

func (v NullableExpiringBaseGrantModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpiringBaseGrantModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
