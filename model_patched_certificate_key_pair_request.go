/*
authentik

Making authentication simple.

API version: 2022.1.2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedCertificateKeyPairRequest CertificateKeyPair Serializer
type PatchedCertificateKeyPairRequest struct {
	Name *string `json:"name,omitempty"`
	// PEM-encoded Certificate data
	CertificateData *string `json:"certificate_data,omitempty"`
	// Optional Private Key. If this is set, you can use this keypair for encryption.
	KeyData *string `json:"key_data,omitempty"`
	// Objects which are managed by authentik. These objects are created and updated automatically. This is flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed NullableString `json:"managed,omitempty"`
}

// NewPatchedCertificateKeyPairRequest instantiates a new PatchedCertificateKeyPairRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedCertificateKeyPairRequest() *PatchedCertificateKeyPairRequest {
	this := PatchedCertificateKeyPairRequest{}
	return &this
}

// NewPatchedCertificateKeyPairRequestWithDefaults instantiates a new PatchedCertificateKeyPairRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedCertificateKeyPairRequestWithDefaults() *PatchedCertificateKeyPairRequest {
	this := PatchedCertificateKeyPairRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedCertificateKeyPairRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCertificateKeyPairRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedCertificateKeyPairRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedCertificateKeyPairRequest) SetName(v string) {
	o.Name = &v
}

// GetCertificateData returns the CertificateData field value if set, zero value otherwise.
func (o *PatchedCertificateKeyPairRequest) GetCertificateData() string {
	if o == nil || o.CertificateData == nil {
		var ret string
		return ret
	}
	return *o.CertificateData
}

// GetCertificateDataOk returns a tuple with the CertificateData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCertificateKeyPairRequest) GetCertificateDataOk() (*string, bool) {
	if o == nil || o.CertificateData == nil {
		return nil, false
	}
	return o.CertificateData, true
}

// HasCertificateData returns a boolean if a field has been set.
func (o *PatchedCertificateKeyPairRequest) HasCertificateData() bool {
	if o != nil && o.CertificateData != nil {
		return true
	}

	return false
}

// SetCertificateData gets a reference to the given string and assigns it to the CertificateData field.
func (o *PatchedCertificateKeyPairRequest) SetCertificateData(v string) {
	o.CertificateData = &v
}

// GetKeyData returns the KeyData field value if set, zero value otherwise.
func (o *PatchedCertificateKeyPairRequest) GetKeyData() string {
	if o == nil || o.KeyData == nil {
		var ret string
		return ret
	}
	return *o.KeyData
}

// GetKeyDataOk returns a tuple with the KeyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCertificateKeyPairRequest) GetKeyDataOk() (*string, bool) {
	if o == nil || o.KeyData == nil {
		return nil, false
	}
	return o.KeyData, true
}

// HasKeyData returns a boolean if a field has been set.
func (o *PatchedCertificateKeyPairRequest) HasKeyData() bool {
	if o != nil && o.KeyData != nil {
		return true
	}

	return false
}

// SetKeyData gets a reference to the given string and assigns it to the KeyData field.
func (o *PatchedCertificateKeyPairRequest) SetKeyData(v string) {
	o.KeyData = &v
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCertificateKeyPairRequest) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCertificateKeyPairRequest) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *PatchedCertificateKeyPairRequest) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *PatchedCertificateKeyPairRequest) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *PatchedCertificateKeyPairRequest) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *PatchedCertificateKeyPairRequest) UnsetManaged() {
	o.Managed.Unset()
}

func (o PatchedCertificateKeyPairRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CertificateData != nil {
		toSerialize["certificate_data"] = o.CertificateData
	}
	if o.KeyData != nil {
		toSerialize["key_data"] = o.KeyData
	}
	if o.Managed.IsSet() {
		toSerialize["managed"] = o.Managed.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedCertificateKeyPairRequest struct {
	value *PatchedCertificateKeyPairRequest
	isSet bool
}

func (v NullablePatchedCertificateKeyPairRequest) Get() *PatchedCertificateKeyPairRequest {
	return v.value
}

func (v *NullablePatchedCertificateKeyPairRequest) Set(val *PatchedCertificateKeyPairRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedCertificateKeyPairRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedCertificateKeyPairRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedCertificateKeyPairRequest(val *PatchedCertificateKeyPairRequest) *NullablePatchedCertificateKeyPairRequest {
	return &NullablePatchedCertificateKeyPairRequest{value: val, isSet: true}
}

func (v NullablePatchedCertificateKeyPairRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedCertificateKeyPairRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
