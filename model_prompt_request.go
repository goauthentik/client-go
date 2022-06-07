/*
authentik

Making authentication simple.

API version: 2022.6.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PromptRequest Prompt Serializer
type PromptRequest struct {
	// Name of the form field, also used to store the value
	FieldKey              string         `json:"field_key"`
	Label                 string         `json:"label"`
	Type                  PromptTypeEnum `json:"type"`
	Required              *bool          `json:"required,omitempty"`
	Placeholder           *string        `json:"placeholder,omitempty"`
	Order                 *int32         `json:"order,omitempty"`
	PromptstageSet        []StageRequest `json:"promptstage_set,omitempty"`
	SubText               *string        `json:"sub_text,omitempty"`
	PlaceholderExpression *bool          `json:"placeholder_expression,omitempty"`
}

// NewPromptRequest instantiates a new PromptRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptRequest(fieldKey string, label string, type_ PromptTypeEnum) *PromptRequest {
	this := PromptRequest{}
	this.FieldKey = fieldKey
	this.Label = label
	this.Type = type_
	return &this
}

// NewPromptRequestWithDefaults instantiates a new PromptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptRequestWithDefaults() *PromptRequest {
	this := PromptRequest{}
	return &this
}

// GetFieldKey returns the FieldKey field value
func (o *PromptRequest) GetFieldKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldKey
}

// GetFieldKeyOk returns a tuple with the FieldKey field value
// and a boolean to check if the value has been set.
func (o *PromptRequest) GetFieldKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldKey, true
}

// SetFieldKey sets field value
func (o *PromptRequest) SetFieldKey(v string) {
	o.FieldKey = v
}

// GetLabel returns the Label field value
func (o *PromptRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *PromptRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *PromptRequest) SetLabel(v string) {
	o.Label = v
}

// GetType returns the Type field value
func (o *PromptRequest) GetType() PromptTypeEnum {
	if o == nil {
		var ret PromptTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PromptRequest) GetTypeOk() (*PromptTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PromptRequest) SetType(v PromptTypeEnum) {
	o.Type = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *PromptRequest) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptRequest) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *PromptRequest) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *PromptRequest) SetRequired(v bool) {
	o.Required = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *PromptRequest) GetPlaceholder() string {
	if o == nil || o.Placeholder == nil {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptRequest) GetPlaceholderOk() (*string, bool) {
	if o == nil || o.Placeholder == nil {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *PromptRequest) HasPlaceholder() bool {
	if o != nil && o.Placeholder != nil {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *PromptRequest) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *PromptRequest) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptRequest) GetOrderOk() (*int32, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *PromptRequest) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *PromptRequest) SetOrder(v int32) {
	o.Order = &v
}

// GetPromptstageSet returns the PromptstageSet field value if set, zero value otherwise.
func (o *PromptRequest) GetPromptstageSet() []StageRequest {
	if o == nil || o.PromptstageSet == nil {
		var ret []StageRequest
		return ret
	}
	return o.PromptstageSet
}

// GetPromptstageSetOk returns a tuple with the PromptstageSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptRequest) GetPromptstageSetOk() ([]StageRequest, bool) {
	if o == nil || o.PromptstageSet == nil {
		return nil, false
	}
	return o.PromptstageSet, true
}

// HasPromptstageSet returns a boolean if a field has been set.
func (o *PromptRequest) HasPromptstageSet() bool {
	if o != nil && o.PromptstageSet != nil {
		return true
	}

	return false
}

// SetPromptstageSet gets a reference to the given []StageRequest and assigns it to the PromptstageSet field.
func (o *PromptRequest) SetPromptstageSet(v []StageRequest) {
	o.PromptstageSet = v
}

// GetSubText returns the SubText field value if set, zero value otherwise.
func (o *PromptRequest) GetSubText() string {
	if o == nil || o.SubText == nil {
		var ret string
		return ret
	}
	return *o.SubText
}

// GetSubTextOk returns a tuple with the SubText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptRequest) GetSubTextOk() (*string, bool) {
	if o == nil || o.SubText == nil {
		return nil, false
	}
	return o.SubText, true
}

// HasSubText returns a boolean if a field has been set.
func (o *PromptRequest) HasSubText() bool {
	if o != nil && o.SubText != nil {
		return true
	}

	return false
}

// SetSubText gets a reference to the given string and assigns it to the SubText field.
func (o *PromptRequest) SetSubText(v string) {
	o.SubText = &v
}

// GetPlaceholderExpression returns the PlaceholderExpression field value if set, zero value otherwise.
func (o *PromptRequest) GetPlaceholderExpression() bool {
	if o == nil || o.PlaceholderExpression == nil {
		var ret bool
		return ret
	}
	return *o.PlaceholderExpression
}

// GetPlaceholderExpressionOk returns a tuple with the PlaceholderExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptRequest) GetPlaceholderExpressionOk() (*bool, bool) {
	if o == nil || o.PlaceholderExpression == nil {
		return nil, false
	}
	return o.PlaceholderExpression, true
}

// HasPlaceholderExpression returns a boolean if a field has been set.
func (o *PromptRequest) HasPlaceholderExpression() bool {
	if o != nil && o.PlaceholderExpression != nil {
		return true
	}

	return false
}

// SetPlaceholderExpression gets a reference to the given bool and assigns it to the PlaceholderExpression field.
func (o *PromptRequest) SetPlaceholderExpression(v bool) {
	o.PlaceholderExpression = &v
}

func (o PromptRequest) MarshalJSON() ([]byte, error) {
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

type NullablePromptRequest struct {
	value *PromptRequest
	isSet bool
}

func (v NullablePromptRequest) Get() *PromptRequest {
	return v.value
}

func (v *NullablePromptRequest) Set(val *PromptRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptRequest(val *PromptRequest) *NullablePromptRequest {
	return &NullablePromptRequest{value: val, isSet: true}
}

func (v NullablePromptRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
