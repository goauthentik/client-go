/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserServiceAccountResponse struct for UserServiceAccountResponse
type UserServiceAccountResponse struct {
	Username string  `json:"username"`
	Token    string  `json:"token"`
	UserUid  string  `json:"user_uid"`
	UserPk   int32   `json:"user_pk"`
	GroupPk  *string `json:"group_pk,omitempty"`
}

// NewUserServiceAccountResponse instantiates a new UserServiceAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserServiceAccountResponse(username string, token string, userUid string, userPk int32) *UserServiceAccountResponse {
	this := UserServiceAccountResponse{}
	this.Username = username
	this.Token = token
	this.UserUid = userUid
	this.UserPk = userPk
	return &this
}

// NewUserServiceAccountResponseWithDefaults instantiates a new UserServiceAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserServiceAccountResponseWithDefaults() *UserServiceAccountResponse {
	this := UserServiceAccountResponse{}
	return &this
}

// GetUsername returns the Username field value
func (o *UserServiceAccountResponse) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserServiceAccountResponse) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserServiceAccountResponse) SetUsername(v string) {
	o.Username = v
}

// GetToken returns the Token field value
func (o *UserServiceAccountResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UserServiceAccountResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UserServiceAccountResponse) SetToken(v string) {
	o.Token = v
}

// GetUserUid returns the UserUid field value
func (o *UserServiceAccountResponse) GetUserUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserUid
}

// GetUserUidOk returns a tuple with the UserUid field value
// and a boolean to check if the value has been set.
func (o *UserServiceAccountResponse) GetUserUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserUid, true
}

// SetUserUid sets field value
func (o *UserServiceAccountResponse) SetUserUid(v string) {
	o.UserUid = v
}

// GetUserPk returns the UserPk field value
func (o *UserServiceAccountResponse) GetUserPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserPk
}

// GetUserPkOk returns a tuple with the UserPk field value
// and a boolean to check if the value has been set.
func (o *UserServiceAccountResponse) GetUserPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserPk, true
}

// SetUserPk sets field value
func (o *UserServiceAccountResponse) SetUserPk(v int32) {
	o.UserPk = v
}

// GetGroupPk returns the GroupPk field value if set, zero value otherwise.
func (o *UserServiceAccountResponse) GetGroupPk() string {
	if o == nil || o.GroupPk == nil {
		var ret string
		return ret
	}
	return *o.GroupPk
}

// GetGroupPkOk returns a tuple with the GroupPk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserServiceAccountResponse) GetGroupPkOk() (*string, bool) {
	if o == nil || o.GroupPk == nil {
		return nil, false
	}
	return o.GroupPk, true
}

// HasGroupPk returns a boolean if a field has been set.
func (o *UserServiceAccountResponse) HasGroupPk() bool {
	if o != nil && o.GroupPk != nil {
		return true
	}

	return false
}

// SetGroupPk gets a reference to the given string and assigns it to the GroupPk field.
func (o *UserServiceAccountResponse) SetGroupPk(v string) {
	o.GroupPk = &v
}

func (o UserServiceAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["user_uid"] = o.UserUid
	}
	if true {
		toSerialize["user_pk"] = o.UserPk
	}
	if o.GroupPk != nil {
		toSerialize["group_pk"] = o.GroupPk
	}
	return json.Marshal(toSerialize)
}

type NullableUserServiceAccountResponse struct {
	value *UserServiceAccountResponse
	isSet bool
}

func (v NullableUserServiceAccountResponse) Get() *UserServiceAccountResponse {
	return v.value
}

func (v *NullableUserServiceAccountResponse) Set(val *UserServiceAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserServiceAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserServiceAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserServiceAccountResponse(val *UserServiceAccountResponse) *NullableUserServiceAccountResponse {
	return &NullableUserServiceAccountResponse{value: val, isSet: true}
}

func (v NullableUserServiceAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserServiceAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
