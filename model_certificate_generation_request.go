/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CertificateGenerationRequest Certificate generation parameters
type CertificateGenerationRequest struct {
	CommonName     string   `json:"common_name"`
	SubjectAltName *string  `json:"subject_alt_name,omitempty"`
	ValidityDays   int32    `json:"validity_days"`
	Alg            *AlgEnum `json:"alg,omitempty"`
}

// NewCertificateGenerationRequest instantiates a new CertificateGenerationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateGenerationRequest(commonName string, validityDays int32) *CertificateGenerationRequest {
	this := CertificateGenerationRequest{}
	this.CommonName = commonName
	this.ValidityDays = validityDays
	return &this
}

// NewCertificateGenerationRequestWithDefaults instantiates a new CertificateGenerationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateGenerationRequestWithDefaults() *CertificateGenerationRequest {
	this := CertificateGenerationRequest{}
	return &this
}

// GetCommonName returns the CommonName field value
func (o *CertificateGenerationRequest) GetCommonName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value
// and a boolean to check if the value has been set.
func (o *CertificateGenerationRequest) GetCommonNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommonName, true
}

// SetCommonName sets field value
func (o *CertificateGenerationRequest) SetCommonName(v string) {
	o.CommonName = v
}

// GetSubjectAltName returns the SubjectAltName field value if set, zero value otherwise.
func (o *CertificateGenerationRequest) GetSubjectAltName() string {
	if o == nil || o.SubjectAltName == nil {
		var ret string
		return ret
	}
	return *o.SubjectAltName
}

// GetSubjectAltNameOk returns a tuple with the SubjectAltName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateGenerationRequest) GetSubjectAltNameOk() (*string, bool) {
	if o == nil || o.SubjectAltName == nil {
		return nil, false
	}
	return o.SubjectAltName, true
}

// HasSubjectAltName returns a boolean if a field has been set.
func (o *CertificateGenerationRequest) HasSubjectAltName() bool {
	if o != nil && o.SubjectAltName != nil {
		return true
	}

	return false
}

// SetSubjectAltName gets a reference to the given string and assigns it to the SubjectAltName field.
func (o *CertificateGenerationRequest) SetSubjectAltName(v string) {
	o.SubjectAltName = &v
}

// GetValidityDays returns the ValidityDays field value
func (o *CertificateGenerationRequest) GetValidityDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ValidityDays
}

// GetValidityDaysOk returns a tuple with the ValidityDays field value
// and a boolean to check if the value has been set.
func (o *CertificateGenerationRequest) GetValidityDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidityDays, true
}

// SetValidityDays sets field value
func (o *CertificateGenerationRequest) SetValidityDays(v int32) {
	o.ValidityDays = v
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *CertificateGenerationRequest) GetAlg() AlgEnum {
	if o == nil || o.Alg == nil {
		var ret AlgEnum
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateGenerationRequest) GetAlgOk() (*AlgEnum, bool) {
	if o == nil || o.Alg == nil {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *CertificateGenerationRequest) HasAlg() bool {
	if o != nil && o.Alg != nil {
		return true
	}

	return false
}

// SetAlg gets a reference to the given AlgEnum and assigns it to the Alg field.
func (o *CertificateGenerationRequest) SetAlg(v AlgEnum) {
	o.Alg = &v
}

func (o CertificateGenerationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["common_name"] = o.CommonName
	}
	if o.SubjectAltName != nil {
		toSerialize["subject_alt_name"] = o.SubjectAltName
	}
	if true {
		toSerialize["validity_days"] = o.ValidityDays
	}
	if o.Alg != nil {
		toSerialize["alg"] = o.Alg
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateGenerationRequest struct {
	value *CertificateGenerationRequest
	isSet bool
}

func (v NullableCertificateGenerationRequest) Get() *CertificateGenerationRequest {
	return v.value
}

func (v *NullableCertificateGenerationRequest) Set(val *CertificateGenerationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateGenerationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateGenerationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateGenerationRequest(val *CertificateGenerationRequest) *NullableCertificateGenerationRequest {
	return &NullableCertificateGenerationRequest{value: val, isSet: true}
}

func (v NullableCertificateGenerationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateGenerationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
