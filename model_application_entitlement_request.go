/*
authentik

Making authentication simple.

API version: 2024.12.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ApplicationEntitlementRequest ApplicationEntitlement Serializer
type ApplicationEntitlementRequest struct {
	Name       string      `json:"name"`
	App        string      `json:"app"`
	Attributes interface{} `json:"attributes,omitempty"`
}

// NewApplicationEntitlementRequest instantiates a new ApplicationEntitlementRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationEntitlementRequest(name string, app string) *ApplicationEntitlementRequest {
	this := ApplicationEntitlementRequest{}
	this.Name = name
	this.App = app
	return &this
}

// NewApplicationEntitlementRequestWithDefaults instantiates a new ApplicationEntitlementRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationEntitlementRequestWithDefaults() *ApplicationEntitlementRequest {
	this := ApplicationEntitlementRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ApplicationEntitlementRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationEntitlementRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationEntitlementRequest) SetName(v string) {
	o.Name = v
}

// GetApp returns the App field value
func (o *ApplicationEntitlementRequest) GetApp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *ApplicationEntitlementRequest) GetAppOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *ApplicationEntitlementRequest) SetApp(v string) {
	o.App = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationEntitlementRequest) GetAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationEntitlementRequest) GetAttributesOk() (*interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ApplicationEntitlementRequest) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *ApplicationEntitlementRequest) SetAttributes(v interface{}) {
	o.Attributes = v
}

func (o ApplicationEntitlementRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["app"] = o.App
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationEntitlementRequest struct {
	value *ApplicationEntitlementRequest
	isSet bool
}

func (v NullableApplicationEntitlementRequest) Get() *ApplicationEntitlementRequest {
	return v.value
}

func (v *NullableApplicationEntitlementRequest) Set(val *ApplicationEntitlementRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationEntitlementRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationEntitlementRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationEntitlementRequest(val *ApplicationEntitlementRequest) *NullableApplicationEntitlementRequest {
	return &NullableApplicationEntitlementRequest{value: val, isSet: true}
}

func (v NullableApplicationEntitlementRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationEntitlementRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
