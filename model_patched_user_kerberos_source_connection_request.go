/*
authentik

Making authentication simple.

API version: 2024.10.5
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedUserKerberosSourceConnectionRequest Kerberos Source Serializer
type PatchedUserKerberosSourceConnectionRequest struct {
	User       *int32  `json:"user,omitempty"`
	Source     *string `json:"source,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
}

// NewPatchedUserKerberosSourceConnectionRequest instantiates a new PatchedUserKerberosSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedUserKerberosSourceConnectionRequest() *PatchedUserKerberosSourceConnectionRequest {
	this := PatchedUserKerberosSourceConnectionRequest{}
	return &this
}

// NewPatchedUserKerberosSourceConnectionRequestWithDefaults instantiates a new PatchedUserKerberosSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedUserKerberosSourceConnectionRequestWithDefaults() *PatchedUserKerberosSourceConnectionRequest {
	this := PatchedUserKerberosSourceConnectionRequest{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *PatchedUserKerberosSourceConnectionRequest) GetUser() int32 {
	if o == nil || o.User == nil {
		var ret int32
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserKerberosSourceConnectionRequest) GetUserOk() (*int32, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedUserKerberosSourceConnectionRequest) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given int32 and assigns it to the User field.
func (o *PatchedUserKerberosSourceConnectionRequest) SetUser(v int32) {
	o.User = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PatchedUserKerberosSourceConnectionRequest) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserKerberosSourceConnectionRequest) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PatchedUserKerberosSourceConnectionRequest) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *PatchedUserKerberosSourceConnectionRequest) SetSource(v string) {
	o.Source = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *PatchedUserKerberosSourceConnectionRequest) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserKerberosSourceConnectionRequest) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *PatchedUserKerberosSourceConnectionRequest) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *PatchedUserKerberosSourceConnectionRequest) SetIdentifier(v string) {
	o.Identifier = &v
}

func (o PatchedUserKerberosSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedUserKerberosSourceConnectionRequest struct {
	value *PatchedUserKerberosSourceConnectionRequest
	isSet bool
}

func (v NullablePatchedUserKerberosSourceConnectionRequest) Get() *PatchedUserKerberosSourceConnectionRequest {
	return v.value
}

func (v *NullablePatchedUserKerberosSourceConnectionRequest) Set(val *PatchedUserKerberosSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedUserKerberosSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedUserKerberosSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedUserKerberosSourceConnectionRequest(val *PatchedUserKerberosSourceConnectionRequest) *NullablePatchedUserKerberosSourceConnectionRequest {
	return &NullablePatchedUserKerberosSourceConnectionRequest{value: val, isSet: true}
}

func (v NullablePatchedUserKerberosSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedUserKerberosSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
