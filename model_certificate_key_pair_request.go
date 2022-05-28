/*
authentik

Making authentication simple.

API version: 2022.5.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CertificateKeyPairRequest CertificateKeyPair Serializer
type CertificateKeyPairRequest struct {
	Name string `json:"name"`
	// PEM-encoded Certificate data
	CertificateData string `json:"certificate_data"`
	// Optional Private Key. If this is set, you can use this keypair for encryption.
	KeyData *string `json:"key_data,omitempty"`
	// Objects which are managed by authentik. These objects are created and updated automatically. This is flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed NullableString `json:"managed,omitempty"`
}

// NewCertificateKeyPairRequest instantiates a new CertificateKeyPairRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateKeyPairRequest(name string, certificateData string) *CertificateKeyPairRequest {
	this := CertificateKeyPairRequest{}
	this.Name = name
	this.CertificateData = certificateData
	return &this
}

// NewCertificateKeyPairRequestWithDefaults instantiates a new CertificateKeyPairRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateKeyPairRequestWithDefaults() *CertificateKeyPairRequest {
	this := CertificateKeyPairRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CertificateKeyPairRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CertificateKeyPairRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CertificateKeyPairRequest) SetName(v string) {
	o.Name = v
}

// GetCertificateData returns the CertificateData field value
func (o *CertificateKeyPairRequest) GetCertificateData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertificateData
}

// GetCertificateDataOk returns a tuple with the CertificateData field value
// and a boolean to check if the value has been set.
func (o *CertificateKeyPairRequest) GetCertificateDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificateData, true
}

// SetCertificateData sets field value
func (o *CertificateKeyPairRequest) SetCertificateData(v string) {
	o.CertificateData = v
}

// GetKeyData returns the KeyData field value if set, zero value otherwise.
func (o *CertificateKeyPairRequest) GetKeyData() string {
	if o == nil || o.KeyData == nil {
		var ret string
		return ret
	}
	return *o.KeyData
}

// GetKeyDataOk returns a tuple with the KeyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateKeyPairRequest) GetKeyDataOk() (*string, bool) {
	if o == nil || o.KeyData == nil {
		return nil, false
	}
	return o.KeyData, true
}

// HasKeyData returns a boolean if a field has been set.
func (o *CertificateKeyPairRequest) HasKeyData() bool {
	if o != nil && o.KeyData != nil {
		return true
	}

	return false
}

// SetKeyData gets a reference to the given string and assigns it to the KeyData field.
func (o *CertificateKeyPairRequest) SetKeyData(v string) {
	o.KeyData = &v
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateKeyPairRequest) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateKeyPairRequest) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *CertificateKeyPairRequest) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *CertificateKeyPairRequest) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *CertificateKeyPairRequest) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *CertificateKeyPairRequest) UnsetManaged() {
	o.Managed.Unset()
}

func (o CertificateKeyPairRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
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

type NullableCertificateKeyPairRequest struct {
	value *CertificateKeyPairRequest
	isSet bool
}

func (v NullableCertificateKeyPairRequest) Get() *CertificateKeyPairRequest {
	return v.value
}

func (v *NullableCertificateKeyPairRequest) Set(val *CertificateKeyPairRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateKeyPairRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateKeyPairRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateKeyPairRequest(val *CertificateKeyPairRequest) *NullableCertificateKeyPairRequest {
	return &NullableCertificateKeyPairRequest{value: val, isSet: true}
}

func (v NullableCertificateKeyPairRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateKeyPairRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
