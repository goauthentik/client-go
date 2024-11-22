/*
authentik

Making authentication simple.

API version: 2024.10.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SCIMSourceUserRequest SCIMSourceUser Serializer
type SCIMSourceUserRequest struct {
	Id         string      `json:"id"`
	User       int32       `json:"user"`
	Source     string      `json:"source"`
	Attributes interface{} `json:"attributes,omitempty"`
}

// NewSCIMSourceUserRequest instantiates a new SCIMSourceUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSCIMSourceUserRequest(id string, user int32, source string) *SCIMSourceUserRequest {
	this := SCIMSourceUserRequest{}
	this.Id = id
	this.User = user
	this.Source = source
	return &this
}

// NewSCIMSourceUserRequestWithDefaults instantiates a new SCIMSourceUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSCIMSourceUserRequestWithDefaults() *SCIMSourceUserRequest {
	this := SCIMSourceUserRequest{}
	return &this
}

// GetId returns the Id field value
func (o *SCIMSourceUserRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SCIMSourceUserRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SCIMSourceUserRequest) SetId(v string) {
	o.Id = v
}

// GetUser returns the User field value
func (o *SCIMSourceUserRequest) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *SCIMSourceUserRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *SCIMSourceUserRequest) SetUser(v int32) {
	o.User = v
}

// GetSource returns the Source field value
func (o *SCIMSourceUserRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *SCIMSourceUserRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *SCIMSourceUserRequest) SetSource(v string) {
	o.Source = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SCIMSourceUserRequest) GetAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SCIMSourceUserRequest) GetAttributesOk() (*interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SCIMSourceUserRequest) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *SCIMSourceUserRequest) SetAttributes(v interface{}) {
	o.Attributes = v
}

func (o SCIMSourceUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableSCIMSourceUserRequest struct {
	value *SCIMSourceUserRequest
	isSet bool
}

func (v NullableSCIMSourceUserRequest) Get() *SCIMSourceUserRequest {
	return v.value
}

func (v *NullableSCIMSourceUserRequest) Set(val *SCIMSourceUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSCIMSourceUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSCIMSourceUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSCIMSourceUserRequest(val *SCIMSourceUserRequest) *NullableSCIMSourceUserRequest {
	return &NullableSCIMSourceUserRequest{value: val, isSet: true}
}

func (v NullableSCIMSourceUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSCIMSourceUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
