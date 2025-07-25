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

// PatchedNotificationWebhookMappingRequest NotificationWebhookMapping Serializer
type PatchedNotificationWebhookMappingRequest struct {
	Name       *string `json:"name,omitempty"`
	Expression *string `json:"expression,omitempty"`
}

// NewPatchedNotificationWebhookMappingRequest instantiates a new PatchedNotificationWebhookMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedNotificationWebhookMappingRequest() *PatchedNotificationWebhookMappingRequest {
	this := PatchedNotificationWebhookMappingRequest{}
	return &this
}

// NewPatchedNotificationWebhookMappingRequestWithDefaults instantiates a new PatchedNotificationWebhookMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedNotificationWebhookMappingRequestWithDefaults() *PatchedNotificationWebhookMappingRequest {
	this := PatchedNotificationWebhookMappingRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedNotificationWebhookMappingRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationWebhookMappingRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedNotificationWebhookMappingRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedNotificationWebhookMappingRequest) SetName(v string) {
	o.Name = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *PatchedNotificationWebhookMappingRequest) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationWebhookMappingRequest) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *PatchedNotificationWebhookMappingRequest) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *PatchedNotificationWebhookMappingRequest) SetExpression(v string) {
	o.Expression = &v
}

func (o PatchedNotificationWebhookMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedNotificationWebhookMappingRequest struct {
	value *PatchedNotificationWebhookMappingRequest
	isSet bool
}

func (v NullablePatchedNotificationWebhookMappingRequest) Get() *PatchedNotificationWebhookMappingRequest {
	return v.value
}

func (v *NullablePatchedNotificationWebhookMappingRequest) Set(val *PatchedNotificationWebhookMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedNotificationWebhookMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedNotificationWebhookMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedNotificationWebhookMappingRequest(val *PatchedNotificationWebhookMappingRequest) *NullablePatchedNotificationWebhookMappingRequest {
	return &NullablePatchedNotificationWebhookMappingRequest{value: val, isSet: true}
}

func (v NullablePatchedNotificationWebhookMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedNotificationWebhookMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
