/*
authentik

Making authentication simple.

API version: 2022.6.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// StagePrompt Serializer for a single Prompt field
type StagePrompt struct {
	FieldKey    string         `json:"field_key"`
	Label       string         `json:"label"`
	Type        PromptTypeEnum `json:"type"`
	Required    bool           `json:"required"`
	Placeholder string         `json:"placeholder"`
	Order       int32          `json:"order"`
	SubText     string         `json:"sub_text"`
}

// NewStagePrompt instantiates a new StagePrompt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStagePrompt(fieldKey string, label string, type_ PromptTypeEnum, required bool, placeholder string, order int32, subText string) *StagePrompt {
	this := StagePrompt{}
	this.FieldKey = fieldKey
	this.Label = label
	this.Type = type_
	this.Required = required
	this.Placeholder = placeholder
	this.Order = order
	this.SubText = subText
	return &this
}

// NewStagePromptWithDefaults instantiates a new StagePrompt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStagePromptWithDefaults() *StagePrompt {
	this := StagePrompt{}
	return &this
}

// GetFieldKey returns the FieldKey field value
func (o *StagePrompt) GetFieldKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldKey
}

// GetFieldKeyOk returns a tuple with the FieldKey field value
// and a boolean to check if the value has been set.
func (o *StagePrompt) GetFieldKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldKey, true
}

// SetFieldKey sets field value
func (o *StagePrompt) SetFieldKey(v string) {
	o.FieldKey = v
}

// GetLabel returns the Label field value
func (o *StagePrompt) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *StagePrompt) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *StagePrompt) SetLabel(v string) {
	o.Label = v
}

// GetType returns the Type field value
func (o *StagePrompt) GetType() PromptTypeEnum {
	if o == nil {
		var ret PromptTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StagePrompt) GetTypeOk() (*PromptTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StagePrompt) SetType(v PromptTypeEnum) {
	o.Type = v
}

// GetRequired returns the Required field value
func (o *StagePrompt) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *StagePrompt) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *StagePrompt) SetRequired(v bool) {
	o.Required = v
}

// GetPlaceholder returns the Placeholder field value
func (o *StagePrompt) GetPlaceholder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value
// and a boolean to check if the value has been set.
func (o *StagePrompt) GetPlaceholderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Placeholder, true
}

// SetPlaceholder sets field value
func (o *StagePrompt) SetPlaceholder(v string) {
	o.Placeholder = v
}

// GetOrder returns the Order field value
func (o *StagePrompt) GetOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *StagePrompt) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *StagePrompt) SetOrder(v int32) {
	o.Order = v
}

// GetSubText returns the SubText field value
func (o *StagePrompt) GetSubText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubText
}

// GetSubTextOk returns a tuple with the SubText field value
// and a boolean to check if the value has been set.
func (o *StagePrompt) GetSubTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubText, true
}

// SetSubText sets field value
func (o *StagePrompt) SetSubText(v string) {
	o.SubText = v
}

func (o StagePrompt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["field_key"] = o.FieldKey
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["required"] = o.Required
	}
	if true {
		toSerialize["placeholder"] = o.Placeholder
	}
	if true {
		toSerialize["order"] = o.Order
	}
	if true {
		toSerialize["sub_text"] = o.SubText
	}
	return json.Marshal(toSerialize)
}

type NullableStagePrompt struct {
	value *StagePrompt
	isSet bool
}

func (v NullableStagePrompt) Get() *StagePrompt {
	return v.value
}

func (v *NullableStagePrompt) Set(val *StagePrompt) {
	v.value = val
	v.isSet = true
}

func (v NullableStagePrompt) IsSet() bool {
	return v.isSet
}

func (v *NullableStagePrompt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStagePrompt(val *StagePrompt) *NullableStagePrompt {
	return &NullableStagePrompt{value: val, isSet: true}
}

func (v NullableStagePrompt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStagePrompt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
