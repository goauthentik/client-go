/*
authentik

Making authentication simple.

API version: 2024.10.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SyncObjectResult Result of a single object sync
type SyncObjectResult struct {
	Messages []LogEvent `json:"messages"`
}

// NewSyncObjectResult instantiates a new SyncObjectResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncObjectResult(messages []LogEvent) *SyncObjectResult {
	this := SyncObjectResult{}
	this.Messages = messages
	return &this
}

// NewSyncObjectResultWithDefaults instantiates a new SyncObjectResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncObjectResultWithDefaults() *SyncObjectResult {
	this := SyncObjectResult{}
	return &this
}

// GetMessages returns the Messages field value
func (o *SyncObjectResult) GetMessages() []LogEvent {
	if o == nil {
		var ret []LogEvent
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *SyncObjectResult) GetMessagesOk() ([]LogEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *SyncObjectResult) SetMessages(v []LogEvent) {
	o.Messages = v
}

func (o SyncObjectResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["messages"] = o.Messages
	}
	return json.Marshal(toSerialize)
}

type NullableSyncObjectResult struct {
	value *SyncObjectResult
	isSet bool
}

func (v NullableSyncObjectResult) Get() *SyncObjectResult {
	return v.value
}

func (v *NullableSyncObjectResult) Set(val *SyncObjectResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncObjectResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncObjectResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncObjectResult(val *SyncObjectResult) *NullableSyncObjectResult {
	return &NullableSyncObjectResult{value: val, isSet: true}
}

func (v NullableSyncObjectResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncObjectResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
