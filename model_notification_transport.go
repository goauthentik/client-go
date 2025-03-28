/*
authentik

Making authentication simple.

API version: 2025.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// NotificationTransport NotificationTransport Serializer
type NotificationTransport struct {
	Pk   string                         `json:"pk"`
	Name string                         `json:"name"`
	Mode *NotificationTransportModeEnum `json:"mode,omitempty"`
	// Return selected mode with a UI Label
	ModeVerbose string  `json:"mode_verbose"`
	WebhookUrl  *string `json:"webhook_url,omitempty"`
	// Customize the body of the request. Mapping should return data that is JSON-serializable.
	WebhookMappingBody NullableString `json:"webhook_mapping_body,omitempty"`
	// Configure additional headers to be sent. Mapping should return a dictionary of key-value pairs
	WebhookMappingHeaders NullableString `json:"webhook_mapping_headers,omitempty"`
	// Only send notification once, for example when sending a webhook into a chat channel.
	SendOnce *bool `json:"send_once,omitempty"`
}

// NewNotificationTransport instantiates a new NotificationTransport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationTransport(pk string, name string, modeVerbose string) *NotificationTransport {
	this := NotificationTransport{}
	this.Pk = pk
	this.Name = name
	this.ModeVerbose = modeVerbose
	return &this
}

// NewNotificationTransportWithDefaults instantiates a new NotificationTransport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationTransportWithDefaults() *NotificationTransport {
	this := NotificationTransport{}
	return &this
}

// GetPk returns the Pk field value
func (o *NotificationTransport) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *NotificationTransport) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *NotificationTransport) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *NotificationTransport) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotificationTransport) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NotificationTransport) SetName(v string) {
	o.Name = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *NotificationTransport) GetMode() NotificationTransportModeEnum {
	if o == nil || o.Mode == nil {
		var ret NotificationTransportModeEnum
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTransport) GetModeOk() (*NotificationTransportModeEnum, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *NotificationTransport) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given NotificationTransportModeEnum and assigns it to the Mode field.
func (o *NotificationTransport) SetMode(v NotificationTransportModeEnum) {
	o.Mode = &v
}

// GetModeVerbose returns the ModeVerbose field value
func (o *NotificationTransport) GetModeVerbose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModeVerbose
}

// GetModeVerboseOk returns a tuple with the ModeVerbose field value
// and a boolean to check if the value has been set.
func (o *NotificationTransport) GetModeVerboseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModeVerbose, true
}

// SetModeVerbose sets field value
func (o *NotificationTransport) SetModeVerbose(v string) {
	o.ModeVerbose = v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *NotificationTransport) GetWebhookUrl() string {
	if o == nil || o.WebhookUrl == nil {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTransport) GetWebhookUrlOk() (*string, bool) {
	if o == nil || o.WebhookUrl == nil {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *NotificationTransport) HasWebhookUrl() bool {
	if o != nil && o.WebhookUrl != nil {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *NotificationTransport) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

// GetWebhookMappingBody returns the WebhookMappingBody field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationTransport) GetWebhookMappingBody() string {
	if o == nil || o.WebhookMappingBody.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebhookMappingBody.Get()
}

// GetWebhookMappingBodyOk returns a tuple with the WebhookMappingBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationTransport) GetWebhookMappingBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebhookMappingBody.Get(), o.WebhookMappingBody.IsSet()
}

// HasWebhookMappingBody returns a boolean if a field has been set.
func (o *NotificationTransport) HasWebhookMappingBody() bool {
	if o != nil && o.WebhookMappingBody.IsSet() {
		return true
	}

	return false
}

// SetWebhookMappingBody gets a reference to the given NullableString and assigns it to the WebhookMappingBody field.
func (o *NotificationTransport) SetWebhookMappingBody(v string) {
	o.WebhookMappingBody.Set(&v)
}

// SetWebhookMappingBodyNil sets the value for WebhookMappingBody to be an explicit nil
func (o *NotificationTransport) SetWebhookMappingBodyNil() {
	o.WebhookMappingBody.Set(nil)
}

// UnsetWebhookMappingBody ensures that no value is present for WebhookMappingBody, not even an explicit nil
func (o *NotificationTransport) UnsetWebhookMappingBody() {
	o.WebhookMappingBody.Unset()
}

// GetWebhookMappingHeaders returns the WebhookMappingHeaders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationTransport) GetWebhookMappingHeaders() string {
	if o == nil || o.WebhookMappingHeaders.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebhookMappingHeaders.Get()
}

// GetWebhookMappingHeadersOk returns a tuple with the WebhookMappingHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationTransport) GetWebhookMappingHeadersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebhookMappingHeaders.Get(), o.WebhookMappingHeaders.IsSet()
}

// HasWebhookMappingHeaders returns a boolean if a field has been set.
func (o *NotificationTransport) HasWebhookMappingHeaders() bool {
	if o != nil && o.WebhookMappingHeaders.IsSet() {
		return true
	}

	return false
}

// SetWebhookMappingHeaders gets a reference to the given NullableString and assigns it to the WebhookMappingHeaders field.
func (o *NotificationTransport) SetWebhookMappingHeaders(v string) {
	o.WebhookMappingHeaders.Set(&v)
}

// SetWebhookMappingHeadersNil sets the value for WebhookMappingHeaders to be an explicit nil
func (o *NotificationTransport) SetWebhookMappingHeadersNil() {
	o.WebhookMappingHeaders.Set(nil)
}

// UnsetWebhookMappingHeaders ensures that no value is present for WebhookMappingHeaders, not even an explicit nil
func (o *NotificationTransport) UnsetWebhookMappingHeaders() {
	o.WebhookMappingHeaders.Unset()
}

// GetSendOnce returns the SendOnce field value if set, zero value otherwise.
func (o *NotificationTransport) GetSendOnce() bool {
	if o == nil || o.SendOnce == nil {
		var ret bool
		return ret
	}
	return *o.SendOnce
}

// GetSendOnceOk returns a tuple with the SendOnce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationTransport) GetSendOnceOk() (*bool, bool) {
	if o == nil || o.SendOnce == nil {
		return nil, false
	}
	return o.SendOnce, true
}

// HasSendOnce returns a boolean if a field has been set.
func (o *NotificationTransport) HasSendOnce() bool {
	if o != nil && o.SendOnce != nil {
		return true
	}

	return false
}

// SetSendOnce gets a reference to the given bool and assigns it to the SendOnce field.
func (o *NotificationTransport) SetSendOnce(v bool) {
	o.SendOnce = &v
}

func (o NotificationTransport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if true {
		toSerialize["mode_verbose"] = o.ModeVerbose
	}
	if o.WebhookUrl != nil {
		toSerialize["webhook_url"] = o.WebhookUrl
	}
	if o.WebhookMappingBody.IsSet() {
		toSerialize["webhook_mapping_body"] = o.WebhookMappingBody.Get()
	}
	if o.WebhookMappingHeaders.IsSet() {
		toSerialize["webhook_mapping_headers"] = o.WebhookMappingHeaders.Get()
	}
	if o.SendOnce != nil {
		toSerialize["send_once"] = o.SendOnce
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationTransport struct {
	value *NotificationTransport
	isSet bool
}

func (v NullableNotificationTransport) Get() *NotificationTransport {
	return v.value
}

func (v *NullableNotificationTransport) Set(val *NotificationTransport) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationTransport) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationTransport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationTransport(val *NotificationTransport) *NullableNotificationTransport {
	return &NullableNotificationTransport{value: val, isSet: true}
}

func (v NullableNotificationTransport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationTransport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
