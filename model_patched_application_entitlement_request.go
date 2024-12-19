/*
authentik

Making authentication simple.

API version: 2024.12.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedApplicationEntitlementRequest ApplicationEntitlement Serializer
type PatchedApplicationEntitlementRequest struct {
	Name       *string     `json:"name,omitempty"`
	App        *string     `json:"app,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}

// NewPatchedApplicationEntitlementRequest instantiates a new PatchedApplicationEntitlementRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedApplicationEntitlementRequest() *PatchedApplicationEntitlementRequest {
	this := PatchedApplicationEntitlementRequest{}
	return &this
}

// NewPatchedApplicationEntitlementRequestWithDefaults instantiates a new PatchedApplicationEntitlementRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedApplicationEntitlementRequestWithDefaults() *PatchedApplicationEntitlementRequest {
	this := PatchedApplicationEntitlementRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedApplicationEntitlementRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedApplicationEntitlementRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedApplicationEntitlementRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedApplicationEntitlementRequest) SetName(v string) {
	o.Name = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *PatchedApplicationEntitlementRequest) GetApp() string {
	if o == nil || o.App == nil {
		var ret string
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedApplicationEntitlementRequest) GetAppOk() (*string, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *PatchedApplicationEntitlementRequest) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given string and assigns it to the App field.
func (o *PatchedApplicationEntitlementRequest) SetApp(v string) {
	o.App = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedApplicationEntitlementRequest) GetAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedApplicationEntitlementRequest) GetAttributesOk() (*interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PatchedApplicationEntitlementRequest) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *PatchedApplicationEntitlementRequest) SetAttributes(v interface{}) {
	o.Attributes = v
}

func (o PatchedApplicationEntitlementRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.App != nil {
		toSerialize["app"] = o.App
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedApplicationEntitlementRequest struct {
	value *PatchedApplicationEntitlementRequest
	isSet bool
}

func (v NullablePatchedApplicationEntitlementRequest) Get() *PatchedApplicationEntitlementRequest {
	return v.value
}

func (v *NullablePatchedApplicationEntitlementRequest) Set(val *PatchedApplicationEntitlementRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedApplicationEntitlementRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedApplicationEntitlementRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedApplicationEntitlementRequest(val *PatchedApplicationEntitlementRequest) *NullablePatchedApplicationEntitlementRequest {
	return &NullablePatchedApplicationEntitlementRequest{value: val, isSet: true}
}

func (v NullablePatchedApplicationEntitlementRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedApplicationEntitlementRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
