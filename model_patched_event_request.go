/*
authentik

Making authentication simple.

API version: 2024.12.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// PatchedEventRequest Event Serializer
type PatchedEventRequest struct {
	User     interface{}    `json:"user,omitempty"`
	Action   *EventActions  `json:"action,omitempty"`
	App      *string        `json:"app,omitempty"`
	Context  interface{}    `json:"context,omitempty"`
	ClientIp NullableString `json:"client_ip,omitempty"`
	Expires  *time.Time     `json:"expires,omitempty"`
	Brand    interface{}    `json:"brand,omitempty"`
}

// NewPatchedEventRequest instantiates a new PatchedEventRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedEventRequest() *PatchedEventRequest {
	this := PatchedEventRequest{}
	return &this
}

// NewPatchedEventRequestWithDefaults instantiates a new PatchedEventRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedEventRequestWithDefaults() *PatchedEventRequest {
	this := PatchedEventRequest{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEventRequest) GetUser() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEventRequest) GetUserOk() (*interface{}, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return &o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedEventRequest) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given interface{} and assigns it to the User field.
func (o *PatchedEventRequest) SetUser(v interface{}) {
	o.User = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PatchedEventRequest) GetAction() EventActions {
	if o == nil || o.Action == nil {
		var ret EventActions
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEventRequest) GetActionOk() (*EventActions, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PatchedEventRequest) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given EventActions and assigns it to the Action field.
func (o *PatchedEventRequest) SetAction(v EventActions) {
	o.Action = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *PatchedEventRequest) GetApp() string {
	if o == nil || o.App == nil {
		var ret string
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEventRequest) GetAppOk() (*string, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *PatchedEventRequest) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given string and assigns it to the App field.
func (o *PatchedEventRequest) SetApp(v string) {
	o.App = &v
}

// GetContext returns the Context field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEventRequest) GetContext() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEventRequest) GetContextOk() (*interface{}, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return &o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *PatchedEventRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given interface{} and assigns it to the Context field.
func (o *PatchedEventRequest) SetContext(v interface{}) {
	o.Context = v
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEventRequest) GetClientIp() string {
	if o == nil || o.ClientIp.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientIp.Get()
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEventRequest) GetClientIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientIp.Get(), o.ClientIp.IsSet()
}

// HasClientIp returns a boolean if a field has been set.
func (o *PatchedEventRequest) HasClientIp() bool {
	if o != nil && o.ClientIp.IsSet() {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given NullableString and assigns it to the ClientIp field.
func (o *PatchedEventRequest) SetClientIp(v string) {
	o.ClientIp.Set(&v)
}

// SetClientIpNil sets the value for ClientIp to be an explicit nil
func (o *PatchedEventRequest) SetClientIpNil() {
	o.ClientIp.Set(nil)
}

// UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
func (o *PatchedEventRequest) UnsetClientIp() {
	o.ClientIp.Unset()
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *PatchedEventRequest) GetExpires() time.Time {
	if o == nil || o.Expires == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEventRequest) GetExpiresOk() (*time.Time, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *PatchedEventRequest) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *PatchedEventRequest) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetBrand returns the Brand field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEventRequest) GetBrand() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEventRequest) GetBrandOk() (*interface{}, bool) {
	if o == nil || o.Brand == nil {
		return nil, false
	}
	return &o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *PatchedEventRequest) HasBrand() bool {
	if o != nil && o.Brand != nil {
		return true
	}

	return false
}

// SetBrand gets a reference to the given interface{} and assigns it to the Brand field.
func (o *PatchedEventRequest) SetBrand(v interface{}) {
	o.Brand = v
}

func (o PatchedEventRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.App != nil {
		toSerialize["app"] = o.App
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.ClientIp.IsSet() {
		toSerialize["client_ip"] = o.ClientIp.Get()
	}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	if o.Brand != nil {
		toSerialize["brand"] = o.Brand
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedEventRequest struct {
	value *PatchedEventRequest
	isSet bool
}

func (v NullablePatchedEventRequest) Get() *PatchedEventRequest {
	return v.value
}

func (v *NullablePatchedEventRequest) Set(val *PatchedEventRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedEventRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedEventRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedEventRequest(val *PatchedEventRequest) *NullablePatchedEventRequest {
	return &NullablePatchedEventRequest{value: val, isSet: true}
}

func (v NullablePatchedEventRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedEventRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
