/*
authentik

Making authentication simple.

API version: 2023.5.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedPromptRequest Prompt Serializer
type PatchedPromptRequest struct {
	Name *string `json:"name,omitempty"`
	// Name of the form field, also used to store the value
	FieldKey *string         `json:"field_key,omitempty"`
	Label    *string         `json:"label,omitempty"`
	Type     *PromptTypeEnum `json:"type,omitempty"`
	Required *bool           `json:"required,omitempty"`
	// Optionally provide a short hint that describes the expected input value. When creating a fixed choice field, enable interpreting as expression and return a list to return multiple choices.
	Placeholder *string `json:"placeholder,omitempty"`
	// Optionally pre-fill the input with an initial value. When creating a fixed choice field, enable interpreting as expression and return a list to return multiple default choices.
	InitialValue           *string        `json:"initial_value,omitempty"`
	Order                  *int32         `json:"order,omitempty"`
	PromptstageSet         []StageRequest `json:"promptstage_set,omitempty"`
	SubText                *string        `json:"sub_text,omitempty"`
	PlaceholderExpression  *bool          `json:"placeholder_expression,omitempty"`
	InitialValueExpression *bool          `json:"initial_value_expression,omitempty"`
}

// NewPatchedPromptRequest instantiates a new PatchedPromptRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedPromptRequest() *PatchedPromptRequest {
	this := PatchedPromptRequest{}
	return &this
}

// NewPatchedPromptRequestWithDefaults instantiates a new PatchedPromptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedPromptRequestWithDefaults() *PatchedPromptRequest {
	this := PatchedPromptRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedPromptRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPromptRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedPromptRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedPromptRequest) SetName(v string) {
	o.Name = &v
}

// GetFieldKey returns the FieldKey field value if set, zero value otherwise.
func (o *PatchedPromptRequest) GetFieldKey() string {
	if o == nil || o.FieldKey == nil {
		var ret string
		return ret
	}
	return *o.FieldKey
}

// GetFieldKeyOk returns a tuple with the FieldKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPromptRequest) GetFieldKeyOk() (*string, bool) {
	if o == nil || o.FieldKey == nil {
		return nil, false
	}
	return o.FieldKey, true
}

// HasFieldKey returns a boolean if a field has been set.
func (o *PatchedPromptRequest) HasFieldKey() bool {
	if o != nil && o.FieldKey != nil {
		return true
	}

	return false
}

// SetFieldKey gets a reference to the given string and assigns it to the FieldKey field.
func (o *PatchedPromptRequest) SetFieldKey(v string) {
	o.FieldKey = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedPromptRequest) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPromptRequest) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedPromptRequest) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedPromptRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedPromptRequest) GetType() PromptTypeEnum {
	if o == nil || o.Type == nil {
		var ret PromptTypeEnum
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPromptRequest) GetTypeOk() (*PromptTypeEnum, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedPromptRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given PromptTypeEnum and assigns it to the Type field.
func (o *PatchedPromptRequest) SetType(v PromptTypeEnum) {
	o.Type = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *PatchedPromptRequest) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPromptRequest) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *PatchedPromptRequest) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *PatchedPromptRequest) SetRequired(v bool) {
	o.Required = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *PatchedPromptRequest) GetPlaceholder() string {
	if o == nil || o.Placeholder == nil {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPromptRequest) GetPlaceholderOk() (*string, bool) {
	if o == nil || o.Placeholder == nil {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *PatchedPromptRequest) HasPlaceholder() bool {
	if o != nil && o.Placeholder != nil {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *PatchedPromptRequest) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetInitialValue returns the InitialValue field value if set, zero value otherwise.
func (o *PatchedPromptRequest) GetInitialValue() string {
	if o == nil || o.InitialValue == nil {
		var ret string
		return ret
	}
	return *o.InitialValue
}

// GetInitialValueOk returns a tuple with the InitialValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPromptRequest) GetInitialValueOk() (*string, bool) {
	if o == nil || o.InitialValue == nil {
		return nil, false
	}
	return o.InitialValue, true
}

// HasInitialValue returns a boolean if a field has been set.
func (o *PatchedPromptRequest) HasInitialValue() bool {
	if o != nil && o.InitialValue != nil {
		return true
	}

	return false
}

// SetInitialValue gets a reference to the given string and assigns it to the InitialValue field.
func (o *PatchedPromptRequest) SetInitialValue(v string) {
	o.InitialValue = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *PatchedPromptRequest) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPromptRequest) GetOrderOk() (*int32, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *PatchedPromptRequest) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *PatchedPromptRequest) SetOrder(v int32) {
	o.Order = &v
}

// GetPromptstageSet returns the PromptstageSet field value if set, zero value otherwise.
func (o *PatchedPromptRequest) GetPromptstageSet() []StageRequest {
	if o == nil || o.PromptstageSet == nil {
		var ret []StageRequest
		return ret
	}
	return o.PromptstageSet
}

// GetPromptstageSetOk returns a tuple with the PromptstageSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPromptRequest) GetPromptstageSetOk() ([]StageRequest, bool) {
	if o == nil || o.PromptstageSet == nil {
		return nil, false
	}
	return o.PromptstageSet, true
}

// HasPromptstageSet returns a boolean if a field has been set.
func (o *PatchedPromptRequest) HasPromptstageSet() bool {
	if o != nil && o.PromptstageSet != nil {
		return true
	}

	return false
}

// SetPromptstageSet gets a reference to the given []StageRequest and assigns it to the PromptstageSet field.
func (o *PatchedPromptRequest) SetPromptstageSet(v []StageRequest) {
	o.PromptstageSet = v
}

// GetSubText returns the SubText field value if set, zero value otherwise.
func (o *PatchedPromptRequest) GetSubText() string {
	if o == nil || o.SubText == nil {
		var ret string
		return ret
	}
	return *o.SubText
}

// GetSubTextOk returns a tuple with the SubText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPromptRequest) GetSubTextOk() (*string, bool) {
	if o == nil || o.SubText == nil {
		return nil, false
	}
	return o.SubText, true
}

// HasSubText returns a boolean if a field has been set.
func (o *PatchedPromptRequest) HasSubText() bool {
	if o != nil && o.SubText != nil {
		return true
	}

	return false
}

// SetSubText gets a reference to the given string and assigns it to the SubText field.
func (o *PatchedPromptRequest) SetSubText(v string) {
	o.SubText = &v
}

// GetPlaceholderExpression returns the PlaceholderExpression field value if set, zero value otherwise.
func (o *PatchedPromptRequest) GetPlaceholderExpression() bool {
	if o == nil || o.PlaceholderExpression == nil {
		var ret bool
		return ret
	}
	return *o.PlaceholderExpression
}

// GetPlaceholderExpressionOk returns a tuple with the PlaceholderExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPromptRequest) GetPlaceholderExpressionOk() (*bool, bool) {
	if o == nil || o.PlaceholderExpression == nil {
		return nil, false
	}
	return o.PlaceholderExpression, true
}

// HasPlaceholderExpression returns a boolean if a field has been set.
func (o *PatchedPromptRequest) HasPlaceholderExpression() bool {
	if o != nil && o.PlaceholderExpression != nil {
		return true
	}

	return false
}

// SetPlaceholderExpression gets a reference to the given bool and assigns it to the PlaceholderExpression field.
func (o *PatchedPromptRequest) SetPlaceholderExpression(v bool) {
	o.PlaceholderExpression = &v
}

// GetInitialValueExpression returns the InitialValueExpression field value if set, zero value otherwise.
func (o *PatchedPromptRequest) GetInitialValueExpression() bool {
	if o == nil || o.InitialValueExpression == nil {
		var ret bool
		return ret
	}
	return *o.InitialValueExpression
}

// GetInitialValueExpressionOk returns a tuple with the InitialValueExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPromptRequest) GetInitialValueExpressionOk() (*bool, bool) {
	if o == nil || o.InitialValueExpression == nil {
		return nil, false
	}
	return o.InitialValueExpression, true
}

// HasInitialValueExpression returns a boolean if a field has been set.
func (o *PatchedPromptRequest) HasInitialValueExpression() bool {
	if o != nil && o.InitialValueExpression != nil {
		return true
	}

	return false
}

// SetInitialValueExpression gets a reference to the given bool and assigns it to the InitialValueExpression field.
func (o *PatchedPromptRequest) SetInitialValueExpression(v bool) {
	o.InitialValueExpression = &v
}

func (o PatchedPromptRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FieldKey != nil {
		toSerialize["field_key"] = o.FieldKey
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Placeholder != nil {
		toSerialize["placeholder"] = o.Placeholder
	}
	if o.InitialValue != nil {
		toSerialize["initial_value"] = o.InitialValue
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
	if o.InitialValueExpression != nil {
		toSerialize["initial_value_expression"] = o.InitialValueExpression
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedPromptRequest struct {
	value *PatchedPromptRequest
	isSet bool
}

func (v NullablePatchedPromptRequest) Get() *PatchedPromptRequest {
	return v.value
}

func (v *NullablePatchedPromptRequest) Set(val *PatchedPromptRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedPromptRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedPromptRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedPromptRequest(val *PatchedPromptRequest) *NullablePatchedPromptRequest {
	return &NullablePatchedPromptRequest{value: val, isSet: true}
}

func (v NullablePatchedPromptRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedPromptRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
