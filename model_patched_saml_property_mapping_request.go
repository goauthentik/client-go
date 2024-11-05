/*
authentik

Making authentication simple.

API version: 2024.10.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedSAMLPropertyMappingRequest SAMLPropertyMapping Serializer
type PatchedSAMLPropertyMappingRequest struct {
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed      NullableString `json:"managed,omitempty"`
	Name         *string        `json:"name,omitempty"`
	Expression   *string        `json:"expression,omitempty"`
	SamlName     *string        `json:"saml_name,omitempty"`
	FriendlyName NullableString `json:"friendly_name,omitempty"`
}

// NewPatchedSAMLPropertyMappingRequest instantiates a new PatchedSAMLPropertyMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedSAMLPropertyMappingRequest() *PatchedSAMLPropertyMappingRequest {
	this := PatchedSAMLPropertyMappingRequest{}
	return &this
}

// NewPatchedSAMLPropertyMappingRequestWithDefaults instantiates a new PatchedSAMLPropertyMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedSAMLPropertyMappingRequestWithDefaults() *PatchedSAMLPropertyMappingRequest {
	this := PatchedSAMLPropertyMappingRequest{}
	return &this
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedSAMLPropertyMappingRequest) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedSAMLPropertyMappingRequest) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *PatchedSAMLPropertyMappingRequest) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *PatchedSAMLPropertyMappingRequest) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *PatchedSAMLPropertyMappingRequest) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *PatchedSAMLPropertyMappingRequest) UnsetManaged() {
	o.Managed.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedSAMLPropertyMappingRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLPropertyMappingRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedSAMLPropertyMappingRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedSAMLPropertyMappingRequest) SetName(v string) {
	o.Name = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *PatchedSAMLPropertyMappingRequest) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLPropertyMappingRequest) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *PatchedSAMLPropertyMappingRequest) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *PatchedSAMLPropertyMappingRequest) SetExpression(v string) {
	o.Expression = &v
}

// GetSamlName returns the SamlName field value if set, zero value otherwise.
func (o *PatchedSAMLPropertyMappingRequest) GetSamlName() string {
	if o == nil || o.SamlName == nil {
		var ret string
		return ret
	}
	return *o.SamlName
}

// GetSamlNameOk returns a tuple with the SamlName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSAMLPropertyMappingRequest) GetSamlNameOk() (*string, bool) {
	if o == nil || o.SamlName == nil {
		return nil, false
	}
	return o.SamlName, true
}

// HasSamlName returns a boolean if a field has been set.
func (o *PatchedSAMLPropertyMappingRequest) HasSamlName() bool {
	if o != nil && o.SamlName != nil {
		return true
	}

	return false
}

// SetSamlName gets a reference to the given string and assigns it to the SamlName field.
func (o *PatchedSAMLPropertyMappingRequest) SetSamlName(v string) {
	o.SamlName = &v
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedSAMLPropertyMappingRequest) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedSAMLPropertyMappingRequest) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *PatchedSAMLPropertyMappingRequest) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *PatchedSAMLPropertyMappingRequest) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}

// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *PatchedSAMLPropertyMappingRequest) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *PatchedSAMLPropertyMappingRequest) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

func (o PatchedSAMLPropertyMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Managed.IsSet() {
		toSerialize["managed"] = o.Managed.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.SamlName != nil {
		toSerialize["saml_name"] = o.SamlName
	}
	if o.FriendlyName.IsSet() {
		toSerialize["friendly_name"] = o.FriendlyName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedSAMLPropertyMappingRequest struct {
	value *PatchedSAMLPropertyMappingRequest
	isSet bool
}

func (v NullablePatchedSAMLPropertyMappingRequest) Get() *PatchedSAMLPropertyMappingRequest {
	return v.value
}

func (v *NullablePatchedSAMLPropertyMappingRequest) Set(val *PatchedSAMLPropertyMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedSAMLPropertyMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedSAMLPropertyMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedSAMLPropertyMappingRequest(val *PatchedSAMLPropertyMappingRequest) *NullablePatchedSAMLPropertyMappingRequest {
	return &NullablePatchedSAMLPropertyMappingRequest{value: val, isSet: true}
}

func (v NullablePatchedSAMLPropertyMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedSAMLPropertyMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
