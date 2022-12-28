/*
authentik

Making authentication simple.

API version: 2022.12.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// RefreshTokenModel Serializer for BaseGrantModel and RefreshToken
type RefreshTokenModel struct {
	Pk        int32          `json:"pk"`
	Provider  OAuth2Provider `json:"provider"`
	User      User           `json:"user"`
	IsExpired bool           `json:"is_expired"`
	Expires   *time.Time     `json:"expires,omitempty"`
	Scope     []string       `json:"scope"`
	IdToken   string         `json:"id_token"`
	Revoked   *bool          `json:"revoked,omitempty"`
}

// NewRefreshTokenModel instantiates a new RefreshTokenModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshTokenModel(pk int32, provider OAuth2Provider, user User, isExpired bool, scope []string, idToken string) *RefreshTokenModel {
	this := RefreshTokenModel{}
	this.Pk = pk
	this.Provider = provider
	this.User = user
	this.IsExpired = isExpired
	this.Scope = scope
	this.IdToken = idToken
	return &this
}

// NewRefreshTokenModelWithDefaults instantiates a new RefreshTokenModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshTokenModelWithDefaults() *RefreshTokenModel {
	this := RefreshTokenModel{}
	return &this
}

// GetPk returns the Pk field value
func (o *RefreshTokenModel) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *RefreshTokenModel) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *RefreshTokenModel) SetPk(v int32) {
	o.Pk = v
}

// GetProvider returns the Provider field value
func (o *RefreshTokenModel) GetProvider() OAuth2Provider {
	if o == nil {
		var ret OAuth2Provider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *RefreshTokenModel) GetProviderOk() (*OAuth2Provider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *RefreshTokenModel) SetProvider(v OAuth2Provider) {
	o.Provider = v
}

// GetUser returns the User field value
func (o *RefreshTokenModel) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *RefreshTokenModel) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *RefreshTokenModel) SetUser(v User) {
	o.User = v
}

// GetIsExpired returns the IsExpired field value
func (o *RefreshTokenModel) GetIsExpired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsExpired
}

// GetIsExpiredOk returns a tuple with the IsExpired field value
// and a boolean to check if the value has been set.
func (o *RefreshTokenModel) GetIsExpiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsExpired, true
}

// SetIsExpired sets field value
func (o *RefreshTokenModel) SetIsExpired(v bool) {
	o.IsExpired = v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *RefreshTokenModel) GetExpires() time.Time {
	if o == nil || o.Expires == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenModel) GetExpiresOk() (*time.Time, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *RefreshTokenModel) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *RefreshTokenModel) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetScope returns the Scope field value
func (o *RefreshTokenModel) GetScope() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *RefreshTokenModel) GetScopeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope, true
}

// SetScope sets field value
func (o *RefreshTokenModel) SetScope(v []string) {
	o.Scope = v
}

// GetIdToken returns the IdToken field value
func (o *RefreshTokenModel) GetIdToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value
// and a boolean to check if the value has been set.
func (o *RefreshTokenModel) GetIdTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdToken, true
}

// SetIdToken sets field value
func (o *RefreshTokenModel) SetIdToken(v string) {
	o.IdToken = v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise.
func (o *RefreshTokenModel) GetRevoked() bool {
	if o == nil || o.Revoked == nil {
		var ret bool
		return ret
	}
	return *o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenModel) GetRevokedOk() (*bool, bool) {
	if o == nil || o.Revoked == nil {
		return nil, false
	}
	return o.Revoked, true
}

// HasRevoked returns a boolean if a field has been set.
func (o *RefreshTokenModel) HasRevoked() bool {
	if o != nil && o.Revoked != nil {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given bool and assigns it to the Revoked field.
func (o *RefreshTokenModel) SetRevoked(v bool) {
	o.Revoked = &v
}

func (o RefreshTokenModel) MarshalJSON() ([]byte, error) {
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

type NullableRefreshTokenModel struct {
	value *RefreshTokenModel
	isSet bool
}

func (v NullableRefreshTokenModel) Get() *RefreshTokenModel {
	return v.value
}

func (v *NullableRefreshTokenModel) Set(val *RefreshTokenModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshTokenModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshTokenModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshTokenModel(val *RefreshTokenModel) *NullableRefreshTokenModel {
	return &NullableRefreshTokenModel{value: val, isSet: true}
}

func (v NullableRefreshTokenModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshTokenModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
