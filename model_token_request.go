/*
authentik

Making authentication simple.

API version: 2024.6.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// TokenRequest Token Serializer
type TokenRequest struct {
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed     NullableString `json:"managed,omitempty"`
	Identifier  string         `json:"identifier"`
	Intent      *IntentEnum    `json:"intent,omitempty"`
	User        *int32         `json:"user,omitempty"`
	Description *string        `json:"description,omitempty"`
	Expires     NullableTime   `json:"expires,omitempty"`
	Expiring    *bool          `json:"expiring,omitempty"`
}

// NewTokenRequest instantiates a new TokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRequest(identifier string) *TokenRequest {
	this := TokenRequest{}
	this.Identifier = identifier
	return &this
}

// NewTokenRequestWithDefaults instantiates a new TokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRequestWithDefaults() *TokenRequest {
	this := TokenRequest{}
	return &this
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenRequest) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenRequest) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *TokenRequest) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *TokenRequest) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *TokenRequest) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *TokenRequest) UnsetManaged() {
	o.Managed.Unset()
}

// GetIdentifier returns the Identifier field value
func (o *TokenRequest) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *TokenRequest) SetIdentifier(v string) {
	o.Identifier = v
}

// GetIntent returns the Intent field value if set, zero value otherwise.
func (o *TokenRequest) GetIntent() IntentEnum {
	if o == nil || o.Intent == nil {
		var ret IntentEnum
		return ret
	}
	return *o.Intent
}

// GetIntentOk returns a tuple with the Intent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetIntentOk() (*IntentEnum, bool) {
	if o == nil || o.Intent == nil {
		return nil, false
	}
	return o.Intent, true
}

// HasIntent returns a boolean if a field has been set.
func (o *TokenRequest) HasIntent() bool {
	if o != nil && o.Intent != nil {
		return true
	}

	return false
}

// SetIntent gets a reference to the given IntentEnum and assigns it to the Intent field.
func (o *TokenRequest) SetIntent(v IntentEnum) {
	o.Intent = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *TokenRequest) GetUser() int32 {
	if o == nil || o.User == nil {
		var ret int32
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetUserOk() (*int32, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *TokenRequest) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given int32 and assigns it to the User field.
func (o *TokenRequest) SetUser(v int32) {
	o.User = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TokenRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TokenRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TokenRequest) SetDescription(v string) {
	o.Description = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TokenRequest) GetExpires() time.Time {
	if o == nil || o.Expires.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TokenRequest) GetExpiresOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// HasExpires returns a boolean if a field has been set.
func (o *TokenRequest) HasExpires() bool {
	if o != nil && o.Expires.IsSet() {
		return true
	}

	return false
}

// SetExpires gets a reference to the given NullableTime and assigns it to the Expires field.
func (o *TokenRequest) SetExpires(v time.Time) {
	o.Expires.Set(&v)
}

// SetExpiresNil sets the value for Expires to be an explicit nil
func (o *TokenRequest) SetExpiresNil() {
	o.Expires.Set(nil)
}

// UnsetExpires ensures that no value is present for Expires, not even an explicit nil
func (o *TokenRequest) UnsetExpires() {
	o.Expires.Unset()
}

// GetExpiring returns the Expiring field value if set, zero value otherwise.
func (o *TokenRequest) GetExpiring() bool {
	if o == nil || o.Expiring == nil {
		var ret bool
		return ret
	}
	return *o.Expiring
}

// GetExpiringOk returns a tuple with the Expiring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetExpiringOk() (*bool, bool) {
	if o == nil || o.Expiring == nil {
		return nil, false
	}
	return o.Expiring, true
}

// HasExpiring returns a boolean if a field has been set.
func (o *TokenRequest) HasExpiring() bool {
	if o != nil && o.Expiring != nil {
		return true
	}

	return false
}

// SetExpiring gets a reference to the given bool and assigns it to the Expiring field.
func (o *TokenRequest) SetExpiring(v bool) {
	o.Expiring = &v
}

func (o TokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Managed.IsSet() {
		toSerialize["managed"] = o.Managed.Get()
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if o.Intent != nil {
		toSerialize["intent"] = o.Intent
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Expires.IsSet() {
		toSerialize["expires"] = o.Expires.Get()
	}
	if o.Expiring != nil {
		toSerialize["expiring"] = o.Expiring
	}
	return json.Marshal(toSerialize)
}

type NullableTokenRequest struct {
	value *TokenRequest
	isSet bool
}

func (v NullableTokenRequest) Get() *TokenRequest {
	return v.value
}

func (v *NullableTokenRequest) Set(val *TokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRequest(val *TokenRequest) *NullableTokenRequest {
	return &NullableTokenRequest{value: val, isSet: true}
}

func (v NullableTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
