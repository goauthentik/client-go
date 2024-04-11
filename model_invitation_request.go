/*
authentik

Making authentication simple.

API version: 2024.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// InvitationRequest Invitation Serializer
type InvitationRequest struct {
	Name      string                 `json:"name"`
	Expires   NullableTime           `json:"expires,omitempty"`
	FixedData map[string]interface{} `json:"fixed_data,omitempty"`
	// When enabled, the invitation will be deleted after usage.
	SingleUse *bool `json:"single_use,omitempty"`
	// When set, only the configured flow can use this invitation.
	Flow NullableString `json:"flow,omitempty"`
}

// NewInvitationRequest instantiates a new InvitationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationRequest(name string) *InvitationRequest {
	this := InvitationRequest{}
	this.Name = name
	return &this
}

// NewInvitationRequestWithDefaults instantiates a new InvitationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationRequestWithDefaults() *InvitationRequest {
	this := InvitationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *InvitationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InvitationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InvitationRequest) SetName(v string) {
	o.Name = v
}

// GetExpires returns the Expires field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitationRequest) GetExpires() time.Time {
	if o == nil || o.Expires.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitationRequest) GetExpiresOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// HasExpires returns a boolean if a field has been set.
func (o *InvitationRequest) HasExpires() bool {
	if o != nil && o.Expires.IsSet() {
		return true
	}

	return false
}

// SetExpires gets a reference to the given NullableTime and assigns it to the Expires field.
func (o *InvitationRequest) SetExpires(v time.Time) {
	o.Expires.Set(&v)
}

// SetExpiresNil sets the value for Expires to be an explicit nil
func (o *InvitationRequest) SetExpiresNil() {
	o.Expires.Set(nil)
}

// UnsetExpires ensures that no value is present for Expires, not even an explicit nil
func (o *InvitationRequest) UnsetExpires() {
	o.Expires.Unset()
}

// GetFixedData returns the FixedData field value if set, zero value otherwise.
func (o *InvitationRequest) GetFixedData() map[string]interface{} {
	if o == nil || o.FixedData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.FixedData
}

// GetFixedDataOk returns a tuple with the FixedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationRequest) GetFixedDataOk() (map[string]interface{}, bool) {
	if o == nil || o.FixedData == nil {
		return nil, false
	}
	return o.FixedData, true
}

// HasFixedData returns a boolean if a field has been set.
func (o *InvitationRequest) HasFixedData() bool {
	if o != nil && o.FixedData != nil {
		return true
	}

	return false
}

// SetFixedData gets a reference to the given map[string]interface{} and assigns it to the FixedData field.
func (o *InvitationRequest) SetFixedData(v map[string]interface{}) {
	o.FixedData = v
}

// GetSingleUse returns the SingleUse field value if set, zero value otherwise.
func (o *InvitationRequest) GetSingleUse() bool {
	if o == nil || o.SingleUse == nil {
		var ret bool
		return ret
	}
	return *o.SingleUse
}

// GetSingleUseOk returns a tuple with the SingleUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationRequest) GetSingleUseOk() (*bool, bool) {
	if o == nil || o.SingleUse == nil {
		return nil, false
	}
	return o.SingleUse, true
}

// HasSingleUse returns a boolean if a field has been set.
func (o *InvitationRequest) HasSingleUse() bool {
	if o != nil && o.SingleUse != nil {
		return true
	}

	return false
}

// SetSingleUse gets a reference to the given bool and assigns it to the SingleUse field.
func (o *InvitationRequest) SetSingleUse(v bool) {
	o.SingleUse = &v
}

// GetFlow returns the Flow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitationRequest) GetFlow() string {
	if o == nil || o.Flow.Get() == nil {
		var ret string
		return ret
	}
	return *o.Flow.Get()
}

// GetFlowOk returns a tuple with the Flow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitationRequest) GetFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Flow.Get(), o.Flow.IsSet()
}

// HasFlow returns a boolean if a field has been set.
func (o *InvitationRequest) HasFlow() bool {
	if o != nil && o.Flow.IsSet() {
		return true
	}

	return false
}

// SetFlow gets a reference to the given NullableString and assigns it to the Flow field.
func (o *InvitationRequest) SetFlow(v string) {
	o.Flow.Set(&v)
}

// SetFlowNil sets the value for Flow to be an explicit nil
func (o *InvitationRequest) SetFlowNil() {
	o.Flow.Set(nil)
}

// UnsetFlow ensures that no value is present for Flow, not even an explicit nil
func (o *InvitationRequest) UnsetFlow() {
	o.Flow.Unset()
}

func (o InvitationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Expires.IsSet() {
		toSerialize["expires"] = o.Expires.Get()
	}
	if o.FixedData != nil {
		toSerialize["fixed_data"] = o.FixedData
	}
	if o.SingleUse != nil {
		toSerialize["single_use"] = o.SingleUse
	}
	if o.Flow.IsSet() {
		toSerialize["flow"] = o.Flow.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInvitationRequest struct {
	value *InvitationRequest
	isSet bool
}

func (v NullableInvitationRequest) Get() *InvitationRequest {
	return v.value
}

func (v *NullableInvitationRequest) Set(val *InvitationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitationRequest(val *InvitationRequest) *NullableInvitationRequest {
	return &NullableInvitationRequest{value: val, isSet: true}
}

func (v NullableInvitationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
