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

// GoogleWorkspaceProviderMappingRequest GoogleWorkspaceProviderMapping Serializer
type GoogleWorkspaceProviderMappingRequest struct {
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed    NullableString `json:"managed,omitempty"`
	Name       string         `json:"name"`
	Expression string         `json:"expression"`
}

// NewGoogleWorkspaceProviderMappingRequest instantiates a new GoogleWorkspaceProviderMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleWorkspaceProviderMappingRequest(name string, expression string) *GoogleWorkspaceProviderMappingRequest {
	this := GoogleWorkspaceProviderMappingRequest{}
	this.Name = name
	this.Expression = expression
	return &this
}

// NewGoogleWorkspaceProviderMappingRequestWithDefaults instantiates a new GoogleWorkspaceProviderMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleWorkspaceProviderMappingRequestWithDefaults() *GoogleWorkspaceProviderMappingRequest {
	this := GoogleWorkspaceProviderMappingRequest{}
	return &this
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoogleWorkspaceProviderMappingRequest) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleWorkspaceProviderMappingRequest) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *GoogleWorkspaceProviderMappingRequest) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *GoogleWorkspaceProviderMappingRequest) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *GoogleWorkspaceProviderMappingRequest) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *GoogleWorkspaceProviderMappingRequest) UnsetManaged() {
	o.Managed.Unset()
}

// GetName returns the Name field value
func (o *GoogleWorkspaceProviderMappingRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderMappingRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GoogleWorkspaceProviderMappingRequest) SetName(v string) {
	o.Name = v
}

// GetExpression returns the Expression field value
func (o *GoogleWorkspaceProviderMappingRequest) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderMappingRequest) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *GoogleWorkspaceProviderMappingRequest) SetExpression(v string) {
	o.Expression = v
}

func (o GoogleWorkspaceProviderMappingRequest) MarshalJSON() ([]byte, error) {
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

type NullableGoogleWorkspaceProviderMappingRequest struct {
	value *GoogleWorkspaceProviderMappingRequest
	isSet bool
}

func (v NullableGoogleWorkspaceProviderMappingRequest) Get() *GoogleWorkspaceProviderMappingRequest {
	return v.value
}

func (v *NullableGoogleWorkspaceProviderMappingRequest) Set(val *GoogleWorkspaceProviderMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleWorkspaceProviderMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleWorkspaceProviderMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleWorkspaceProviderMappingRequest(val *GoogleWorkspaceProviderMappingRequest) *NullableGoogleWorkspaceProviderMappingRequest {
	return &NullableGoogleWorkspaceProviderMappingRequest{value: val, isSet: true}
}

func (v NullableGoogleWorkspaceProviderMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleWorkspaceProviderMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
