/*
authentik

Making authentication simple.

API version: 2023.10.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// TokenModel Serializer for BaseGrantModel and RefreshToken
type TokenModel struct {
	Pk       int32          `json:"pk"`
	Provider OAuth2Provider `json:"provider"`
	User     User           `json:"user"`
	// Check if token is expired yet.
	IsExpired bool       `json:"is_expired"`
	Expires   *time.Time `json:"expires,omitempty"`
	Scope     []string   `json:"scope"`
	// Get the token's id_token as JSON String
	IdToken string `json:"id_token"`
	Revoked *bool  `json:"revoked,omitempty"`
}

// NewTokenModel instantiates a new TokenModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenModel(pk int32, provider OAuth2Provider, user User, isExpired bool, scope []string, idToken string) *TokenModel {
	this := TokenModel{}
	this.Pk = pk
	this.Provider = provider
	this.User = user
	this.IsExpired = isExpired
	this.Scope = scope
	this.IdToken = idToken
	return &this
}

// NewTokenModelWithDefaults instantiates a new TokenModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenModelWithDefaults() *TokenModel {
	this := TokenModel{}
	return &this
}

// GetPk returns the Pk field value
func (o *TokenModel) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *TokenModel) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *TokenModel) SetPk(v int32) {
	o.Pk = v
}

// GetProvider returns the Provider field value
func (o *TokenModel) GetProvider() OAuth2Provider {
	if o == nil {
		var ret OAuth2Provider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *TokenModel) GetProviderOk() (*OAuth2Provider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *TokenModel) SetProvider(v OAuth2Provider) {
	o.Provider = v
}

// GetUser returns the User field value
func (o *TokenModel) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TokenModel) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TokenModel) SetUser(v User) {
	o.User = v
}

// GetIsExpired returns the IsExpired field value
func (o *TokenModel) GetIsExpired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsExpired
}

// GetIsExpiredOk returns a tuple with the IsExpired field value
// and a boolean to check if the value has been set.
func (o *TokenModel) GetIsExpiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsExpired, true
}

// SetIsExpired sets field value
func (o *TokenModel) SetIsExpired(v bool) {
	o.IsExpired = v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *TokenModel) GetExpires() time.Time {
	if o == nil || o.Expires == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenModel) GetExpiresOk() (*time.Time, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *TokenModel) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *TokenModel) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetScope returns the Scope field value
func (o *TokenModel) GetScope() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *TokenModel) GetScopeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope, true
}

// SetScope sets field value
func (o *TokenModel) SetScope(v []string) {
	o.Scope = v
}

// GetIdToken returns the IdToken field value
func (o *TokenModel) GetIdToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value
// and a boolean to check if the value has been set.
func (o *TokenModel) GetIdTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdToken, true
}

// SetIdToken sets field value
func (o *TokenModel) SetIdToken(v string) {
	o.IdToken = v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise.
func (o *TokenModel) GetRevoked() bool {
	if o == nil || o.Revoked == nil {
		var ret bool
		return ret
	}
	return *o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenModel) GetRevokedOk() (*bool, bool) {
	if o == nil || o.Revoked == nil {
		return nil, false
	}
	return o.Revoked, true
}

// HasRevoked returns a boolean if a field has been set.
func (o *TokenModel) HasRevoked() bool {
	if o != nil && o.Revoked != nil {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given bool and assigns it to the Revoked field.
func (o *TokenModel) SetRevoked(v bool) {
	o.Revoked = &v
}

func (o TokenModel) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["id_token"] = o.IdToken
	}
	if o.Revoked != nil {
		toSerialize["revoked"] = o.Revoked
	}
	return json.Marshal(toSerialize)
}

type NullableTokenModel struct {
	value *TokenModel
	isSet bool
}

func (v NullableTokenModel) Get() *TokenModel {
	return v.value
}

func (v *NullableTokenModel) Set(val *TokenModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenModel(val *TokenModel) *NullableTokenModel {
	return &NullableTokenModel{value: val, isSet: true}
}

func (v NullableTokenModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
