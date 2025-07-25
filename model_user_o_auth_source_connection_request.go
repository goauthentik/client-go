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
)

// UserOAuthSourceConnectionRequest User source connection
type UserOAuthSourceConnectionRequest struct {
	User        int32          `json:"user"`
	Source      string         `json:"source"`
	Identifier  string         `json:"identifier"`
	AccessToken NullableString `json:"access_token,omitempty"`
}

// NewUserOAuthSourceConnectionRequest instantiates a new UserOAuthSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserOAuthSourceConnectionRequest(user int32, source string, identifier string) *UserOAuthSourceConnectionRequest {
	this := UserOAuthSourceConnectionRequest{}
	this.User = user
	this.Source = source
	this.Identifier = identifier
	return &this
}

// NewUserOAuthSourceConnectionRequestWithDefaults instantiates a new UserOAuthSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserOAuthSourceConnectionRequestWithDefaults() *UserOAuthSourceConnectionRequest {
	this := UserOAuthSourceConnectionRequest{}
	return &this
}

// GetUser returns the User field value
func (o *UserOAuthSourceConnectionRequest) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserOAuthSourceConnectionRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserOAuthSourceConnectionRequest) SetUser(v int32) {
	o.User = v
}

// GetSource returns the Source field value
func (o *UserOAuthSourceConnectionRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *UserOAuthSourceConnectionRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *UserOAuthSourceConnectionRequest) SetSource(v string) {
	o.Source = v
}

// GetIdentifier returns the Identifier field value
func (o *UserOAuthSourceConnectionRequest) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *UserOAuthSourceConnectionRequest) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *UserOAuthSourceConnectionRequest) SetIdentifier(v string) {
	o.Identifier = v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserOAuthSourceConnectionRequest) GetAccessToken() string {
	if o == nil || o.AccessToken.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserOAuthSourceConnectionRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *UserOAuthSourceConnectionRequest) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *UserOAuthSourceConnectionRequest) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}

// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *UserOAuthSourceConnectionRequest) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *UserOAuthSourceConnectionRequest) UnsetAccessToken() {
	o.AccessToken.Unset()
}

func (o UserOAuthSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if o.AccessToken.IsSet() {
		toSerialize["access_token"] = o.AccessToken.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUserOAuthSourceConnectionRequest struct {
	value *UserOAuthSourceConnectionRequest
	isSet bool
}

func (v NullableUserOAuthSourceConnectionRequest) Get() *UserOAuthSourceConnectionRequest {
	return v.value
}

func (v *NullableUserOAuthSourceConnectionRequest) Set(val *UserOAuthSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserOAuthSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserOAuthSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserOAuthSourceConnectionRequest(val *UserOAuthSourceConnectionRequest) *NullableUserOAuthSourceConnectionRequest {
	return &NullableUserOAuthSourceConnectionRequest{value: val, isSet: true}
}

func (v NullableUserOAuthSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserOAuthSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
