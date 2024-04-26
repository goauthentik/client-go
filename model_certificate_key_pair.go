/*
authentik

Making authentication simple.

API version: 2024.4.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// CertificateKeyPair CertificateKeyPair Serializer
type CertificateKeyPair struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// Get certificate Hash (SHA256)
	FingerprintSha256 NullableString `json:"fingerprint_sha256"`
	// Get certificate Hash (SHA1)
	FingerprintSha1 NullableString `json:"fingerprint_sha1"`
	// Get certificate expiry
	CertExpiry NullableTime `json:"cert_expiry"`
	// Get certificate subject as full rfc4514
	CertSubject NullableString `json:"cert_subject"`
	// Show if this keypair has a private key configured or not
	PrivateKeyAvailable bool `json:"private_key_available"`
	// Get the private key's type, if set
	PrivateKeyType NullableString `json:"private_key_type"`
	// Get URL to download certificate
	CertificateDownloadUrl string `json:"certificate_download_url"`
	// Get URL to download private key
	PrivateKeyDownloadUrl string `json:"private_key_download_url"`
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed NullableString `json:"managed"`
}

// NewCertificateKeyPair instantiates a new CertificateKeyPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateKeyPair(pk string, name string, fingerprintSha256 NullableString, fingerprintSha1 NullableString, certExpiry NullableTime, certSubject NullableString, privateKeyAvailable bool, privateKeyType NullableString, certificateDownloadUrl string, privateKeyDownloadUrl string, managed NullableString) *CertificateKeyPair {
	this := CertificateKeyPair{}
	this.Pk = pk
	this.Name = name
	this.FingerprintSha256 = fingerprintSha256
	this.FingerprintSha1 = fingerprintSha1
	this.CertExpiry = certExpiry
	this.CertSubject = certSubject
	this.PrivateKeyAvailable = privateKeyAvailable
	this.PrivateKeyType = privateKeyType
	this.CertificateDownloadUrl = certificateDownloadUrl
	this.PrivateKeyDownloadUrl = privateKeyDownloadUrl
	this.Managed = managed
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
// If the value is explicit nil, the zero value for string will be returned
func (o *CertificateKeyPair) GetFingerprintSha256() string {
	if o == nil || o.FingerprintSha256.Get() == nil {
		var ret string
		return ret
	}

	return *o.FingerprintSha256.Get()
}

// GetFingerprintSha256Ok returns a tuple with the FingerprintSha256 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateKeyPair) GetFingerprintSha256Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FingerprintSha256.Get(), o.FingerprintSha256.IsSet()
}

// SetFingerprintSha256 sets field value
func (o *CertificateKeyPair) SetFingerprintSha256(v string) {
	o.FingerprintSha256.Set(&v)
}

// GetFingerprintSha1 returns the FingerprintSha1 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CertificateKeyPair) GetFingerprintSha1() string {
	if o == nil || o.FingerprintSha1.Get() == nil {
		var ret string
		return ret
	}

	return *o.FingerprintSha1.Get()
}

// GetFingerprintSha1Ok returns a tuple with the FingerprintSha1 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateKeyPair) GetFingerprintSha1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FingerprintSha1.Get(), o.FingerprintSha1.IsSet()
}

// SetFingerprintSha1 sets field value
func (o *CertificateKeyPair) SetFingerprintSha1(v string) {
	o.FingerprintSha1.Set(&v)
}

// GetCertExpiry returns the CertExpiry field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CertificateKeyPair) GetCertExpiry() time.Time {
	if o == nil || o.CertExpiry.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CertExpiry.Get()
}

// GetCertExpiryOk returns a tuple with the CertExpiry field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateKeyPair) GetCertExpiryOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertExpiry.Get(), o.CertExpiry.IsSet()
}

// SetCertExpiry sets field value
func (o *CertificateKeyPair) SetCertExpiry(v time.Time) {
	o.CertExpiry.Set(&v)
}

// GetCertSubject returns the CertSubject field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CertificateKeyPair) GetCertSubject() string {
	if o == nil || o.CertSubject.Get() == nil {
		var ret string
		return ret
	}

	return *o.CertSubject.Get()
}

// GetCertSubjectOk returns a tuple with the CertSubject field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateKeyPair) GetCertSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertSubject.Get(), o.CertSubject.IsSet()
}

// SetCertSubject sets field value
func (o *CertificateKeyPair) SetCertSubject(v string) {
	o.CertSubject.Set(&v)
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

// GetPrivateKeyType returns the PrivateKeyType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CertificateKeyPair) GetPrivateKeyType() string {
	if o == nil || o.PrivateKeyType.Get() == nil {
		var ret string
		return ret
	}

	return *o.PrivateKeyType.Get()
}

// GetPrivateKeyTypeOk returns a tuple with the PrivateKeyType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateKeyPair) GetPrivateKeyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateKeyType.Get(), o.PrivateKeyType.IsSet()
}

// SetPrivateKeyType sets field value
func (o *CertificateKeyPair) SetPrivateKeyType(v string) {
	o.PrivateKeyType.Set(&v)
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

// GetManaged returns the Managed field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CertificateKeyPair) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}

	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateKeyPair) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// SetManaged sets field value
func (o *CertificateKeyPair) SetManaged(v string) {
	o.Managed.Set(&v)
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
		toSerialize["fingerprint_sha256"] = o.FingerprintSha256.Get()
	}
	if true {
		toSerialize["fingerprint_sha1"] = o.FingerprintSha1.Get()
	}
	if true {
		toSerialize["cert_expiry"] = o.CertExpiry.Get()
	}
	if true {
		toSerialize["cert_subject"] = o.CertSubject.Get()
	}
	if true {
		toSerialize["private_key_available"] = o.PrivateKeyAvailable
	}
	if true {
		toSerialize["private_key_type"] = o.PrivateKeyType.Get()
	}
	if true {
		toSerialize["certificate_download_url"] = o.CertificateDownloadUrl
	}
	if true {
		toSerialize["private_key_download_url"] = o.PrivateKeyDownloadUrl
	}
	if true {
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
