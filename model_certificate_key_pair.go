/*
authentik

Making authentication simple.

API version: 2021.10.1-rc1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// CertificateKeyPair CertificateKeyPair Serializer
type CertificateKeyPair struct {
	Pk                     string    `json:"pk"`
	Name                   string    `json:"name"`
	FingerprintSha256      string    `json:"fingerprint_sha256"`
	FingerprintSha1        string    `json:"fingerprint_sha1"`
	CertExpiry             time.Time `json:"cert_expiry"`
	CertSubject            string    `json:"cert_subject"`
	PrivateKeyAvailable    bool      `json:"private_key_available"`
	CertificateDownloadUrl string    `json:"certificate_download_url"`
	PrivateKeyDownloadUrl  string    `json:"private_key_download_url"`
	// Objects which are managed by authentik. These objects are created and updated automatically. This is flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed NullableString `json:"managed,omitempty"`
}

// NewCertificateKeyPair instantiates a new CertificateKeyPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateKeyPair(pk string, name string, fingerprintSha256 string, fingerprintSha1 string, certExpiry time.Time, certSubject string, privateKeyAvailable bool, certificateDownloadUrl string, privateKeyDownloadUrl string) *CertificateKeyPair {
	this := CertificateKeyPair{}
	this.Pk = pk
	this.Name = name
	this.FingerprintSha256 = fingerprintSha256
	this.FingerprintSha1 = fingerprintSha1
	this.CertExpiry = certExpiry
	this.CertSubject = certSubject
	this.PrivateKeyAvailable = privateKeyAvailable
	this.CertificateDownloadUrl = certificateDownloadUrl
	this.PrivateKeyDownloadUrl = privateKeyDownloadUrl
	return &this
}

// NewCertificateKeyPairWithDefaults instantiates a new CertificateKeyPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateKeyPairWithDefaults() *CertificateKeyPair {
	this := CertificateKeyPair{}
	return &this
}

// GetPk returns the Pk field value
func (o *CertificateKeyPair) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *CertificateKeyPair) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *CertificateKeyPair) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *CertificateKeyPair) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CertificateKeyPair) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CertificateKeyPair) SetName(v string) {
	o.Name = v
}

// GetFingerprintSha256 returns the FingerprintSha256 field value
func (o *CertificateKeyPair) GetFingerprintSha256() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FingerprintSha256
}

// GetFingerprintSha256Ok returns a tuple with the FingerprintSha256 field value
// and a boolean to check if the value has been set.
func (o *CertificateKeyPair) GetFingerprintSha256Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FingerprintSha256, true
}

// SetFingerprintSha256 sets field value
func (o *CertificateKeyPair) SetFingerprintSha256(v string) {
	o.FingerprintSha256 = v
}

// GetFingerprintSha1 returns the FingerprintSha1 field value
func (o *CertificateKeyPair) GetFingerprintSha1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FingerprintSha1
}

// GetFingerprintSha1Ok returns a tuple with the FingerprintSha1 field value
// and a boolean to check if the value has been set.
func (o *CertificateKeyPair) GetFingerprintSha1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FingerprintSha1, true
}

// SetFingerprintSha1 sets field value
func (o *CertificateKeyPair) SetFingerprintSha1(v string) {
	o.FingerprintSha1 = v
}

// GetCertExpiry returns the CertExpiry field value
func (o *CertificateKeyPair) GetCertExpiry() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CertExpiry
}

// GetCertExpiryOk returns a tuple with the CertExpiry field value
// and a boolean to check if the value has been set.
func (o *CertificateKeyPair) GetCertExpiryOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertExpiry, true
}

// SetCertExpiry sets field value
func (o *CertificateKeyPair) SetCertExpiry(v time.Time) {
	o.CertExpiry = v
}

// GetCertSubject returns the CertSubject field value
func (o *CertificateKeyPair) GetCertSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertSubject
}

// GetCertSubjectOk returns a tuple with the CertSubject field value
// and a boolean to check if the value has been set.
func (o *CertificateKeyPair) GetCertSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertSubject, true
}

// SetCertSubject sets field value
func (o *CertificateKeyPair) SetCertSubject(v string) {
	o.CertSubject = v
}

// GetPrivateKeyAvailable returns the PrivateKeyAvailable field value
func (o *CertificateKeyPair) GetPrivateKeyAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PrivateKeyAvailable
}

// GetPrivateKeyAvailableOk returns a tuple with the PrivateKeyAvailable field value
// and a boolean to check if the value has been set.
func (o *CertificateKeyPair) GetPrivateKeyAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKeyAvailable, true
}

// SetPrivateKeyAvailable sets field value
func (o *CertificateKeyPair) SetPrivateKeyAvailable(v bool) {
	o.PrivateKeyAvailable = v
}

// GetCertificateDownloadUrl returns the CertificateDownloadUrl field value
func (o *CertificateKeyPair) GetCertificateDownloadUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertificateDownloadUrl
}

// GetCertificateDownloadUrlOk returns a tuple with the CertificateDownloadUrl field value
// and a boolean to check if the value has been set.
func (o *CertificateKeyPair) GetCertificateDownloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificateDownloadUrl, true
}

// SetCertificateDownloadUrl sets field value
func (o *CertificateKeyPair) SetCertificateDownloadUrl(v string) {
	o.CertificateDownloadUrl = v
}

// GetPrivateKeyDownloadUrl returns the PrivateKeyDownloadUrl field value
func (o *CertificateKeyPair) GetPrivateKeyDownloadUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKeyDownloadUrl
}

// GetPrivateKeyDownloadUrlOk returns a tuple with the PrivateKeyDownloadUrl field value
// and a boolean to check if the value has been set.
func (o *CertificateKeyPair) GetPrivateKeyDownloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKeyDownloadUrl, true
}

// SetPrivateKeyDownloadUrl sets field value
func (o *CertificateKeyPair) SetPrivateKeyDownloadUrl(v string) {
	o.PrivateKeyDownloadUrl = v
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateKeyPair) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateKeyPair) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *CertificateKeyPair) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *CertificateKeyPair) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *CertificateKeyPair) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *CertificateKeyPair) UnsetManaged() {
	o.Managed.Unset()
}

func (o CertificateKeyPair) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["fingerprint_sha256"] = o.FingerprintSha256
	}
	if true {
		toSerialize["fingerprint_sha1"] = o.FingerprintSha1
	}
	if true {
		toSerialize["cert_expiry"] = o.CertExpiry
	}
	if true {
		toSerialize["cert_subject"] = o.CertSubject
	}
	if true {
		toSerialize["private_key_available"] = o.PrivateKeyAvailable
	}
	if true {
		toSerialize["certificate_download_url"] = o.CertificateDownloadUrl
	}
	if true {
		toSerialize["private_key_download_url"] = o.PrivateKeyDownloadUrl
	}
	if o.Managed.IsSet() {
		toSerialize["managed"] = o.Managed.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateKeyPair struct {
	value *CertificateKeyPair
	isSet bool
}

func (v NullableCertificateKeyPair) Get() *CertificateKeyPair {
	return v.value
}

func (v *NullableCertificateKeyPair) Set(val *CertificateKeyPair) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateKeyPair) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateKeyPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateKeyPair(val *CertificateKeyPair) *NullableCertificateKeyPair {
	return &NullableCertificateKeyPair{value: val, isSet: true}
}

func (v NullableCertificateKeyPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateKeyPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
