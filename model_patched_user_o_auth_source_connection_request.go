/*
authentik

Making authentication simple.

API version: 2023.8.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedUserOAuthSourceConnectionRequest OAuth Source Serializer
type PatchedUserOAuthSourceConnectionRequest struct {
	User        *int32         `json:"user,omitempty"`
	Identifier  *string        `json:"identifier,omitempty"`
	AccessToken NullableString `json:"access_token,omitempty"`
}

// NewPatchedUserOAuthSourceConnectionRequest instantiates a new PatchedUserOAuthSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedUserOAuthSourceConnectionRequest() *PatchedUserOAuthSourceConnectionRequest {
	this := PatchedUserOAuthSourceConnectionRequest{}
	return &this
}

// NewPatchedUserOAuthSourceConnectionRequestWithDefaults instantiates a new PatchedUserOAuthSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedUserOAuthSourceConnectionRequestWithDefaults() *PatchedUserOAuthSourceConnectionRequest {
	this := PatchedUserOAuthSourceConnectionRequest{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *PatchedUserOAuthSourceConnectionRequest) GetUser() int32 {
	if o == nil || o.User == nil {
		var ret int32
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserOAuthSourceConnectionRequest) GetUserOk() (*int32, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedUserOAuthSourceConnectionRequest) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given int32 and assigns it to the User field.
func (o *PatchedUserOAuthSourceConnectionRequest) SetUser(v int32) {
	o.User = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *PatchedUserOAuthSourceConnectionRequest) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserOAuthSourceConnectionRequest) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *PatchedUserOAuthSourceConnectionRequest) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *PatchedUserOAuthSourceConnectionRequest) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedUserOAuthSourceConnectionRequest) GetAccessToken() string {
	if o == nil || o.AccessToken.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedUserOAuthSourceConnectionRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *PatchedUserOAuthSourceConnectionRequest) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *PatchedUserOAuthSourceConnectionRequest) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}

// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *PatchedUserOAuthSourceConnectionRequest) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *PatchedUserOAuthSourceConnectionRequest) UnsetAccessToken() {
	o.AccessToken.Unset()
}

func (o PatchedUserOAuthSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.AccessToken.IsSet() {
		toSerialize["access_token"] = o.AccessToken.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedUserOAuthSourceConnectionRequest struct {
	value *PatchedUserOAuthSourceConnectionRequest
	isSet bool
}

func (v NullablePatchedUserOAuthSourceConnectionRequest) Get() *PatchedUserOAuthSourceConnectionRequest {
	return v.value
}

func (v *NullablePatchedUserOAuthSourceConnectionRequest) Set(val *PatchedUserOAuthSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedUserOAuthSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedUserOAuthSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedUserOAuthSourceConnectionRequest(val *PatchedUserOAuthSourceConnectionRequest) *NullablePatchedUserOAuthSourceConnectionRequest {
	return &NullablePatchedUserOAuthSourceConnectionRequest{value: val, isSet: true}
}

func (v NullablePatchedUserOAuthSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedUserOAuthSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
