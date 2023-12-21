/*
authentik

Making authentication simple.

API version: 2023.10.5
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// FlowErrorChallenge Challenge class when an unhandled error occurs during a stage. Normal users are shown an error message, superusers are shown a full stacktrace.
type FlowErrorChallenge struct {
	Type           *string                   `json:"type,omitempty"`
	FlowInfo       *ContextualFlowInfo       `json:"flow_info,omitempty"`
	Component      *string                   `json:"component,omitempty"`
	ResponseErrors *map[string][]ErrorDetail `json:"response_errors,omitempty"`
	RequestId      string                    `json:"request_id"`
	Error          *string                   `json:"error,omitempty"`
	Traceback      *string                   `json:"traceback,omitempty"`
}

// NewFlowErrorChallenge instantiates a new FlowErrorChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowErrorChallenge(requestId string) *FlowErrorChallenge {
	this := FlowErrorChallenge{}
	var type_ string = "native"
	this.Type = &type_
	var component string = "ak-stage-flow-error"
	this.Component = &component
	this.RequestId = requestId
	return &this
}

// NewFlowErrorChallengeWithDefaults instantiates a new FlowErrorChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowErrorChallengeWithDefaults() *FlowErrorChallenge {
	this := FlowErrorChallenge{}
	var type_ string = "native"
	this.Type = &type_
	var component string = "ak-stage-flow-error"
	this.Component = &component
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FlowErrorChallenge) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowErrorChallenge) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FlowErrorChallenge) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FlowErrorChallenge) SetType(v string) {
	o.Type = &v
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *FlowErrorChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || o.FlowInfo == nil {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowErrorChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || o.FlowInfo == nil {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *FlowErrorChallenge) HasFlowInfo() bool {
	if o != nil && o.FlowInfo != nil {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *FlowErrorChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *FlowErrorChallenge) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowErrorChallenge) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *FlowErrorChallenge) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *FlowErrorChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *FlowErrorChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || o.ResponseErrors == nil {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowErrorChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || o.ResponseErrors == nil {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *FlowErrorChallenge) HasResponseErrors() bool {
	if o != nil && o.ResponseErrors != nil {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *FlowErrorChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

// GetRequestId returns the RequestId field value
func (o *FlowErrorChallenge) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *FlowErrorChallenge) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *FlowErrorChallenge) SetRequestId(v string) {
	o.RequestId = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *FlowErrorChallenge) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowErrorChallenge) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *FlowErrorChallenge) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *FlowErrorChallenge) SetError(v string) {
	o.Error = &v
}

// GetTraceback returns the Traceback field value if set, zero value otherwise.
func (o *FlowErrorChallenge) GetTraceback() string {
	if o == nil || o.Traceback == nil {
		var ret string
		return ret
	}
	return *o.Traceback
}

// GetTracebackOk returns a tuple with the Traceback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowErrorChallenge) GetTracebackOk() (*string, bool) {
	if o == nil || o.Traceback == nil {
		return nil, false
	}
	return o.Traceback, true
}

// HasTraceback returns a boolean if a field has been set.
func (o *FlowErrorChallenge) HasTraceback() bool {
	if o != nil && o.Traceback != nil {
		return true
	}

	return false
}

// SetTraceback gets a reference to the given string and assigns it to the Traceback field.
func (o *FlowErrorChallenge) SetTraceback(v string) {
	o.Traceback = &v
}

func (o FlowErrorChallenge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.FlowInfo != nil {
		toSerialize["flow_info"] = o.FlowInfo
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.ResponseErrors != nil {
		toSerialize["response_errors"] = o.ResponseErrors
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Traceback != nil {
		toSerialize["traceback"] = o.Traceback
	}
	return json.Marshal(toSerialize)
}

type NullableFlowErrorChallenge struct {
	value *FlowErrorChallenge
	isSet bool
}

func (v NullableFlowErrorChallenge) Get() *FlowErrorChallenge {
	return v.value
}

func (v *NullableFlowErrorChallenge) Set(val *FlowErrorChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowErrorChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowErrorChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowErrorChallenge(val *FlowErrorChallenge) *NullableFlowErrorChallenge {
	return &NullableFlowErrorChallenge{value: val, isSet: true}
}

func (v NullableFlowErrorChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowErrorChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
