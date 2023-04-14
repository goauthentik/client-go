/*
authentik

Making authentication simple.

API version: 2023.4.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Prompt Prompt Serializer
type Prompt struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// Name of the form field, also used to store the value
	FieldKey string         `json:"field_key"`
	Label    string         `json:"label"`
	Type     PromptTypeEnum `json:"type"`
	Required *bool          `json:"required,omitempty"`
	// When creating a Radio Button Group or Dropdown, enable interpreting as expression and return a list to return multiple choices.
	Placeholder           *string `json:"placeholder,omitempty"`
	Order                 *int32  `json:"order,omitempty"`
	PromptstageSet        []Stage `json:"promptstage_set,omitempty"`
	SubText               *string `json:"sub_text,omitempty"`
	PlaceholderExpression *bool   `json:"placeholder_expression,omitempty"`
}

// NewPrompt instantiates a new Prompt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrompt(pk string, name string, fieldKey string, label string, type_ PromptTypeEnum) *Prompt {
	this := Prompt{}
	this.Pk = pk
	this.Name = name
	this.FieldKey = fieldKey
	this.Label = label
	this.Type = type_
	return &this
}

// NewPromptWithDefaults instantiates a new Prompt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptWithDefaults() *Prompt {
	this := Prompt{}
	return &this
}

// GetPk returns the Pk field value
func (o *Prompt) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *Prompt) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *Prompt) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *Prompt) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Prompt) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Prompt) SetName(v string) {
	o.Name = v
}

// GetFieldKey returns the FieldKey field value
func (o *Prompt) GetFieldKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldKey
}

// GetFieldKeyOk returns a tuple with the FieldKey field value
// and a boolean to check if the value has been set.
func (o *Prompt) GetFieldKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldKey, true
}

// SetFieldKey sets field value
func (o *Prompt) SetFieldKey(v string) {
	o.FieldKey = v
}

// GetLabel returns the Label field value
func (o *Prompt) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Prompt) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Prompt) SetLabel(v string) {
	o.Label = v
}

// GetType returns the Type field value
func (o *Prompt) GetType() PromptTypeEnum {
	if o == nil {
		var ret PromptTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Prompt) GetTypeOk() (*PromptTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Prompt) SetType(v PromptTypeEnum) {
	o.Type = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *Prompt) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prompt) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *Prompt) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *Prompt) SetRequired(v bool) {
	o.Required = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *Prompt) GetPlaceholder() string {
	if o == nil || o.Placeholder == nil {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prompt) GetPlaceholderOk() (*string, bool) {
	if o == nil || o.Placeholder == nil {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *Prompt) HasPlaceholder() bool {
	if o != nil && o.Placeholder != nil {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *Prompt) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *Prompt) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prompt) GetOrderOk() (*int32, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *Prompt) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *Prompt) SetOrder(v int32) {
	o.Order = &v
}

// GetPromptstageSet returns the PromptstageSet field value if set, zero value otherwise.
func (o *Prompt) GetPromptstageSet() []Stage {
	if o == nil || o.PromptstageSet == nil {
		var ret []Stage
		return ret
	}
	return o.PromptstageSet
}

// GetPromptstageSetOk returns a tuple with the PromptstageSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prompt) GetPromptstageSetOk() ([]Stage, bool) {
	if o == nil || o.PromptstageSet == nil {
		return nil, false
	}
	return o.PromptstageSet, true
}

// HasPromptstageSet returns a boolean if a field has been set.
func (o *Prompt) HasPromptstageSet() bool {
	if o != nil && o.PromptstageSet != nil {
		return true
	}

	return false
}

// SetPromptstageSet gets a reference to the given []Stage and assigns it to the PromptstageSet field.
func (o *Prompt) SetPromptstageSet(v []Stage) {
	o.PromptstageSet = v
}

// GetSubText returns the SubText field value if set, zero value otherwise.
func (o *Prompt) GetSubText() string {
	if o == nil || o.SubText == nil {
		var ret string
		return ret
	}
	return *o.SubText
}

// GetSubTextOk returns a tuple with the SubText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prompt) GetSubTextOk() (*string, bool) {
	if o == nil || o.SubText == nil {
		return nil, false
	}
	return o.SubText, true
}

// HasSubText returns a boolean if a field has been set.
func (o *Prompt) HasSubText() bool {
	if o != nil && o.SubText != nil {
		return true
	}

	return false
}

// SetSubText gets a reference to the given string and assigns it to the SubText field.
func (o *Prompt) SetSubText(v string) {
	o.SubText = &v
}

// GetPlaceholderExpression returns the PlaceholderExpression field value if set, zero value otherwise.
func (o *Prompt) GetPlaceholderExpression() bool {
	if o == nil || o.PlaceholderExpression == nil {
		var ret bool
		return ret
	}
	return *o.PlaceholderExpression
}

// GetPlaceholderExpressionOk returns a tuple with the PlaceholderExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prompt) GetPlaceholderExpressionOk() (*bool, bool) {
	if o == nil || o.PlaceholderExpression == nil {
		return nil, false
	}
	return o.PlaceholderExpression, true
}

// HasPlaceholderExpression returns a boolean if a field has been set.
func (o *Prompt) HasPlaceholderExpression() bool {
	if o != nil && o.PlaceholderExpression != nil {
		return true
	}

	return false
}

// SetPlaceholderExpression gets a reference to the given bool and assigns it to the PlaceholderExpression field.
func (o *Prompt) SetPlaceholderExpression(v bool) {
	o.PlaceholderExpression = &v
}

func (o Prompt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["field_key"] = o.FieldKey
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Placeholder != nil {
		toSerialize["placeholder"] = o.Placeholder
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.PromptstageSet != nil {
		toSerialize["promptstage_set"] = o.PromptstageSet
	}
	if o.SubText != nil {
		toSerialize["sub_text"] = o.SubText
	}
	if o.PlaceholderExpression != nil {
		toSerialize["placeholder_expression"] = o.PlaceholderExpression
	}
	return json.Marshal(toSerialize)
}

type NullablePrompt struct {
	value *Prompt
	isSet bool
}

func (v NullablePrompt) Get() *Prompt {
	return v.value
}

func (v *NullablePrompt) Set(val *Prompt) {
	v.value = val
	v.isSet = true
}

func (v NullablePrompt) IsSet() bool {
	return v.isSet
}

func (v *NullablePrompt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrompt(val *Prompt) *NullablePrompt {
	return &NullablePrompt{value: val, isSet: true}
}

func (v NullablePrompt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrompt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
