/*
authentik

Making authentication simple.

API version: 2025.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SSFStream SSFStream Serializer
type SSFStream struct {
	Pk              string                `json:"pk"`
	Provider        int32                 `json:"provider"`
	ProviderObj     SSFProvider           `json:"provider_obj"`
	DeliveryMethod  DeliveryMethodEnum    `json:"delivery_method"`
	EndpointUrl     NullableString        `json:"endpoint_url,omitempty"`
	EventsRequested []EventsRequestedEnum `json:"events_requested,omitempty"`
	Format          string                `json:"format"`
	Aud             []string              `json:"aud,omitempty"`
	Iss             string                `json:"iss"`
}

// NewSSFStream instantiates a new SSFStream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSFStream(pk string, provider int32, providerObj SSFProvider, deliveryMethod DeliveryMethodEnum, format string, iss string) *SSFStream {
	this := SSFStream{}
	this.Pk = pk
	this.Provider = provider
	this.ProviderObj = providerObj
	this.DeliveryMethod = deliveryMethod
	this.Format = format
	this.Iss = iss
	return &this
}

// NewSSFStreamWithDefaults instantiates a new SSFStream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSFStreamWithDefaults() *SSFStream {
	this := SSFStream{}
	return &this
}

// GetPk returns the Pk field value
func (o *SSFStream) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *SSFStream) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *SSFStream) SetPk(v string) {
	o.Pk = v
}

// GetProvider returns the Provider field value
func (o *SSFStream) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *SSFStream) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *SSFStream) SetProvider(v int32) {
	o.Provider = v
}

// GetProviderObj returns the ProviderObj field value
func (o *SSFStream) GetProviderObj() SSFProvider {
	if o == nil {
		var ret SSFProvider
		return ret
	}

	return o.ProviderObj
}

// GetProviderObjOk returns a tuple with the ProviderObj field value
// and a boolean to check if the value has been set.
func (o *SSFStream) GetProviderObjOk() (*SSFProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderObj, true
}

// SetProviderObj sets field value
func (o *SSFStream) SetProviderObj(v SSFProvider) {
	o.ProviderObj = v
}

// GetDeliveryMethod returns the DeliveryMethod field value
func (o *SSFStream) GetDeliveryMethod() DeliveryMethodEnum {
	if o == nil {
		var ret DeliveryMethodEnum
		return ret
	}

	return o.DeliveryMethod
}

// GetDeliveryMethodOk returns a tuple with the DeliveryMethod field value
// and a boolean to check if the value has been set.
func (o *SSFStream) GetDeliveryMethodOk() (*DeliveryMethodEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryMethod, true
}

// SetDeliveryMethod sets field value
func (o *SSFStream) SetDeliveryMethod(v DeliveryMethodEnum) {
	o.DeliveryMethod = v
}

// GetEndpointUrl returns the EndpointUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SSFStream) GetEndpointUrl() string {
	if o == nil || o.EndpointUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.EndpointUrl.Get()
}

// GetEndpointUrlOk returns a tuple with the EndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SSFStream) GetEndpointUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndpointUrl.Get(), o.EndpointUrl.IsSet()
}

// HasEndpointUrl returns a boolean if a field has been set.
func (o *SSFStream) HasEndpointUrl() bool {
	if o != nil && o.EndpointUrl.IsSet() {
		return true
	}

	return false
}

// SetEndpointUrl gets a reference to the given NullableString and assigns it to the EndpointUrl field.
func (o *SSFStream) SetEndpointUrl(v string) {
	o.EndpointUrl.Set(&v)
}

// SetEndpointUrlNil sets the value for EndpointUrl to be an explicit nil
func (o *SSFStream) SetEndpointUrlNil() {
	o.EndpointUrl.Set(nil)
}

// UnsetEndpointUrl ensures that no value is present for EndpointUrl, not even an explicit nil
func (o *SSFStream) UnsetEndpointUrl() {
	o.EndpointUrl.Unset()
}

// GetEventsRequested returns the EventsRequested field value if set, zero value otherwise.
func (o *SSFStream) GetEventsRequested() []EventsRequestedEnum {
	if o == nil || o.EventsRequested == nil {
		var ret []EventsRequestedEnum
		return ret
	}
	return o.EventsRequested
}

// GetEventsRequestedOk returns a tuple with the EventsRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSFStream) GetEventsRequestedOk() ([]EventsRequestedEnum, bool) {
	if o == nil || o.EventsRequested == nil {
		return nil, false
	}
	return o.EventsRequested, true
}

// HasEventsRequested returns a boolean if a field has been set.
func (o *SSFStream) HasEventsRequested() bool {
	if o != nil && o.EventsRequested != nil {
		return true
	}

	return false
}

// SetEventsRequested gets a reference to the given []EventsRequestedEnum and assigns it to the EventsRequested field.
func (o *SSFStream) SetEventsRequested(v []EventsRequestedEnum) {
	o.EventsRequested = v
}

// GetFormat returns the Format field value
func (o *SSFStream) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *SSFStream) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *SSFStream) SetFormat(v string) {
	o.Format = v
}

// GetAud returns the Aud field value if set, zero value otherwise.
func (o *SSFStream) GetAud() []string {
	if o == nil || o.Aud == nil {
		var ret []string
		return ret
	}
	return o.Aud
}

// GetAudOk returns a tuple with the Aud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSFStream) GetAudOk() ([]string, bool) {
	if o == nil || o.Aud == nil {
		return nil, false
	}
	return o.Aud, true
}

// HasAud returns a boolean if a field has been set.
func (o *SSFStream) HasAud() bool {
	if o != nil && o.Aud != nil {
		return true
	}

	return false
}

// SetAud gets a reference to the given []string and assigns it to the Aud field.
func (o *SSFStream) SetAud(v []string) {
	o.Aud = v
}

// GetIss returns the Iss field value
func (o *SSFStream) GetIss() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iss
}

// GetIssOk returns a tuple with the Iss field value
// and a boolean to check if the value has been set.
func (o *SSFStream) GetIssOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iss, true
}

// SetIss sets field value
func (o *SSFStream) SetIss(v string) {
	o.Iss = v
}

func (o SSFStream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["provider_obj"] = o.ProviderObj
	}
	if true {
		toSerialize["delivery_method"] = o.DeliveryMethod
	}
	if o.EndpointUrl.IsSet() {
		toSerialize["endpoint_url"] = o.EndpointUrl.Get()
	}
	if o.EventsRequested != nil {
		toSerialize["events_requested"] = o.EventsRequested
	}
	if true {
		toSerialize["format"] = o.Format
	}
	if o.Aud != nil {
		toSerialize["aud"] = o.Aud
	}
	if true {
		toSerialize["iss"] = o.Iss
	}
	return json.Marshal(toSerialize)
}

type NullableSSFStream struct {
	value *SSFStream
	isSet bool
}

func (v NullableSSFStream) Get() *SSFStream {
	return v.value
}

func (v *NullableSSFStream) Set(val *SSFStream) {
	v.value = val
	v.isSet = true
}

func (v NullableSSFStream) IsSet() bool {
	return v.isSet
}

func (v *NullableSSFStream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSFStream(val *SSFStream) *NullableSSFStream {
	return &NullableSSFStream{value: val, isSet: true}
}

func (v NullableSSFStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSFStream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
