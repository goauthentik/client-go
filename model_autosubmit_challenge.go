/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AutosubmitChallenge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutosubmitChallenge{}

// AutosubmitChallenge Autosubmit challenge used to send and navigate a POST request
type AutosubmitChallenge struct {
	FlowInfo       *ContextualFlowInfo       `json:"flow_info,omitempty"`
	Component      *string                   `json:"component,omitempty"`
	ResponseErrors *map[string][]ErrorDetail `json:"response_errors,omitempty"`
	Url            string                    `json:"url"`
	Attrs          map[string]string         `json:"attrs"`
	Title          *string                   `json:"title,omitempty"`
}

type _AutosubmitChallenge AutosubmitChallenge

// NewAutosubmitChallenge instantiates a new AutosubmitChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutosubmitChallenge(url string, attrs map[string]string) *AutosubmitChallenge {
	this := AutosubmitChallenge{}
	var component string = "ak-stage-autosubmit"
	this.Component = &component
	this.Url = url
	this.Attrs = attrs
	return &this
}

// NewAutosubmitChallengeWithDefaults instantiates a new AutosubmitChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutosubmitChallengeWithDefaults() *AutosubmitChallenge {
	this := AutosubmitChallenge{}
	var component string = "ak-stage-autosubmit"
	this.Component = &component
	return &this
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *AutosubmitChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || IsNil(o.FlowInfo) {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutosubmitChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || IsNil(o.FlowInfo) {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *AutosubmitChallenge) HasFlowInfo() bool {
	if o != nil && !IsNil(o.FlowInfo) {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *AutosubmitChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AutosubmitChallenge) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutosubmitChallenge) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AutosubmitChallenge) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AutosubmitChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *AutosubmitChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || IsNil(o.ResponseErrors) {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutosubmitChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || IsNil(o.ResponseErrors) {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *AutosubmitChallenge) HasResponseErrors() bool {
	if o != nil && !IsNil(o.ResponseErrors) {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *AutosubmitChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

// GetUrl returns the Url field value
func (o *AutosubmitChallenge) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AutosubmitChallenge) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AutosubmitChallenge) SetUrl(v string) {
	o.Url = v
}

// GetAttrs returns the Attrs field value
func (o *AutosubmitChallenge) GetAttrs() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Attrs
}

// GetAttrsOk returns a tuple with the Attrs field value
// and a boolean to check if the value has been set.
func (o *AutosubmitChallenge) GetAttrsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attrs, true
}

// SetAttrs sets field value
func (o *AutosubmitChallenge) SetAttrs(v map[string]string) {
	o.Attrs = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AutosubmitChallenge) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutosubmitChallenge) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AutosubmitChallenge) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AutosubmitChallenge) SetTitle(v string) {
	o.Title = &v
}

func (o AutosubmitChallenge) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutosubmitChallenge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FlowInfo) {
		toSerialize["flow_info"] = o.FlowInfo
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.ResponseErrors) {
		toSerialize["response_errors"] = o.ResponseErrors
	}
	toSerialize["url"] = o.Url
	toSerialize["attrs"] = o.Attrs
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

func (o *AutosubmitChallenge) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"attrs",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAutosubmitChallenge := _AutosubmitChallenge{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAutosubmitChallenge)

	if err != nil {
		return err
	}

	*o = AutosubmitChallenge(varAutosubmitChallenge)

	return err
}

type NullableAutosubmitChallenge struct {
	value *AutosubmitChallenge
	isSet bool
}

func (v NullableAutosubmitChallenge) Get() *AutosubmitChallenge {
	return v.value
}

func (v *NullableAutosubmitChallenge) Set(val *AutosubmitChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableAutosubmitChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableAutosubmitChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutosubmitChallenge(val *AutosubmitChallenge) *NullableAutosubmitChallenge {
	return &NullableAutosubmitChallenge{value: val, isSet: true}
}

func (v NullableAutosubmitChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutosubmitChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
