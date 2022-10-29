/*
authentik

Making authentication simple.

API version: 2022.10.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// NotificationTransportRequest NotificationTransport Serializer
type NotificationTransportRequest struct {
	Name           string                         `json:"name"`
	Mode           *NotificationTransportModeEnum `json:"mode,omitempty"`
	WebhookUrl     *string                        `json:"webhook_url,omitempty"`
	WebhookMapping NullableString                 `json:"webhook_mapping,omitempty"`
	// Only send notification once, for example when sending a webhook into a chat channel.
	SendOnce *bool `json:"send_once,omitempty"`
}

// NewNotificationTransportRequest instantiates a new NotificationTransportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationTransportRequest(name string) *NotificationTransportRequest {
	this := NotificationTransportRequest{}
	this.Name = name
	return &this
}

// NewNotificationTransportRequestWithDefaults instantiates a new NotificationTransportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationTransportRequestWithDefaults() *NotificationTransportRequest {
	this := NotificationTransportRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NotificationTransportRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotificationTransportRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NotificationTransportRequest) SetName(v string) {
	o.Name = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *NotificationTransportRequest) GetMode() NotificationTransportModeEnum {
	if o == nil || o.Mode == nil {
		var ret NotificationTransportModeEnum
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTransportRequest) GetModeOk() (*NotificationTransportModeEnum, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *NotificationTransportRequest) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given NotificationTransportModeEnum and assigns it to the Mode field.
func (o *NotificationTransportRequest) SetMode(v NotificationTransportModeEnum) {
	o.Mode = &v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *NotificationTransportRequest) GetWebhookUrl() string {
	if o == nil || o.WebhookUrl == nil {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTransportRequest) GetWebhookUrlOk() (*string, bool) {
	if o == nil || o.WebhookUrl == nil {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *NotificationTransportRequest) HasWebhookUrl() bool {
	if o != nil && o.WebhookUrl != nil {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *NotificationTransportRequest) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

// GetWebhookMapping returns the WebhookMapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationTransportRequest) GetWebhookMapping() string {
	if o == nil || o.WebhookMapping.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebhookMapping.Get()
}

// GetWebhookMappingOk returns a tuple with the WebhookMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationTransportRequest) GetWebhookMappingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebhookMapping.Get(), o.WebhookMapping.IsSet()
}

// HasWebhookMapping returns a boolean if a field has been set.
func (o *NotificationTransportRequest) HasWebhookMapping() bool {
	if o != nil && o.WebhookMapping.IsSet() {
		return true
	}

	return false
}

// SetWebhookMapping gets a reference to the given NullableString and assigns it to the WebhookMapping field.
func (o *NotificationTransportRequest) SetWebhookMapping(v string) {
	o.WebhookMapping.Set(&v)
}

// SetWebhookMappingNil sets the value for WebhookMapping to be an explicit nil
func (o *NotificationTransportRequest) SetWebhookMappingNil() {
	o.WebhookMapping.Set(nil)
}

// UnsetWebhookMapping ensures that no value is present for WebhookMapping, not even an explicit nil
func (o *NotificationTransportRequest) UnsetWebhookMapping() {
	o.WebhookMapping.Unset()
}

// GetSendOnce returns the SendOnce field value if set, zero value otherwise.
func (o *NotificationTransportRequest) GetSendOnce() bool {
	if o == nil || o.SendOnce == nil {
		var ret bool
		return ret
	}
	return *o.SendOnce
}

// GetSendOnceOk returns a tuple with the SendOnce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTransportRequest) GetSendOnceOk() (*bool, bool) {
	if o == nil || o.SendOnce == nil {
		return nil, false
	}
	return o.SendOnce, true
}

// HasSendOnce returns a boolean if a field has been set.
func (o *NotificationTransportRequest) HasSendOnce() bool {
	if o != nil && o.SendOnce != nil {
		return true
	}

	return false
}

// SetSendOnce gets a reference to the given bool and assigns it to the SendOnce field.
func (o *NotificationTransportRequest) SetSendOnce(v bool) {
	o.SendOnce = &v
}

func (o NotificationTransportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.WebhookUrl != nil {
		toSerialize["webhook_url"] = o.WebhookUrl
	}
	if o.WebhookMapping.IsSet() {
		toSerialize["webhook_mapping"] = o.WebhookMapping.Get()
	}
	if o.SendOnce != nil {
		toSerialize["send_once"] = o.SendOnce
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationTransportRequest struct {
	value *NotificationTransportRequest
	isSet bool
}

func (v NullableNotificationTransportRequest) Get() *NotificationTransportRequest {
	return v.value
}

func (v *NullableNotificationTransportRequest) Set(val *NotificationTransportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationTransportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationTransportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationTransportRequest(val *NotificationTransportRequest) *NullableNotificationTransportRequest {
	return &NullableNotificationTransportRequest{value: val, isSet: true}
}

func (v NullableNotificationTransportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationTransportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
