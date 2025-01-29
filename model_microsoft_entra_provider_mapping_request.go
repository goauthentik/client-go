/*
authentik

Making authentication simple.

API version: 2024.12.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// MicrosoftEntraProviderMappingRequest MicrosoftEntraProviderMapping Serializer
type MicrosoftEntraProviderMappingRequest struct {
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed    NullableString `json:"managed,omitempty"`
	Name       string         `json:"name"`
	Expression string         `json:"expression"`
}

// NewMicrosoftEntraProviderMappingRequest instantiates a new MicrosoftEntraProviderMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftEntraProviderMappingRequest(name string, expression string) *MicrosoftEntraProviderMappingRequest {
	this := MicrosoftEntraProviderMappingRequest{}
	this.Name = name
	this.Expression = expression
	return &this
}

// NewMicrosoftEntraProviderMappingRequestWithDefaults instantiates a new MicrosoftEntraProviderMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftEntraProviderMappingRequestWithDefaults() *MicrosoftEntraProviderMappingRequest {
	this := MicrosoftEntraProviderMappingRequest{}
	return &this
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftEntraProviderMappingRequest) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftEntraProviderMappingRequest) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *MicrosoftEntraProviderMappingRequest) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *MicrosoftEntraProviderMappingRequest) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *MicrosoftEntraProviderMappingRequest) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *MicrosoftEntraProviderMappingRequest) UnsetManaged() {
	o.Managed.Unset()
}

// GetName returns the Name field value
func (o *MicrosoftEntraProviderMappingRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderMappingRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MicrosoftEntraProviderMappingRequest) SetName(v string) {
	o.Name = v
}

// GetExpression returns the Expression field value
func (o *MicrosoftEntraProviderMappingRequest) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderMappingRequest) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *MicrosoftEntraProviderMappingRequest) SetExpression(v string) {
	o.Expression = v
}

func (o MicrosoftEntraProviderMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Managed.IsSet() {
		toSerialize["managed"] = o.Managed.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["expression"] = o.Expression
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftEntraProviderMappingRequest struct {
	value *MicrosoftEntraProviderMappingRequest
	isSet bool
}

func (v NullableMicrosoftEntraProviderMappingRequest) Get() *MicrosoftEntraProviderMappingRequest {
	return v.value
}

func (v *NullableMicrosoftEntraProviderMappingRequest) Set(val *MicrosoftEntraProviderMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftEntraProviderMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftEntraProviderMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftEntraProviderMappingRequest(val *MicrosoftEntraProviderMappingRequest) *NullableMicrosoftEntraProviderMappingRequest {
	return &NullableMicrosoftEntraProviderMappingRequest{value: val, isSet: true}
}

func (v NullableMicrosoftEntraProviderMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftEntraProviderMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
