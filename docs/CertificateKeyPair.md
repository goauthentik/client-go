# CertificateKeyPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**FingerprintSha256** | **string** |  | [readonly] 
**FingerprintSha1** | **string** |  | [readonly] 
**CertExpiry** | **time.Time** |  | [readonly] 
**CertSubject** | **string** |  | [readonly] 
**PrivateKeyAvailable** | **bool** |  | [readonly] 
**PrivateKeyType** | **NullableString** |  | [readonly] 
**CertificateDownloadUrl** | **string** |  | [readonly] 
**PrivateKeyDownloadUrl** | **string** |  | [readonly] 
**Managed** | Pointer to **NullableString** | Objects which are managed by authentik. These objects are created and updated automatically. This is flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 

## Methods

### NewCertificateKeyPair

`func NewCertificateKeyPair(pk string, name string, fingerprintSha256 string, fingerprintSha1 string, certExpiry time.Time, certSubject string, privateKeyAvailable bool, privateKeyType NullableString, certificateDownloadUrl string, privateKeyDownloadUrl string, ) *CertificateKeyPair`

NewCertificateKeyPair instantiates a new CertificateKeyPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateKeyPairWithDefaults

`func NewCertificateKeyPairWithDefaults() *CertificateKeyPair`

NewCertificateKeyPairWithDefaults instantiates a new CertificateKeyPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *CertificateKeyPair) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *CertificateKeyPair) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *CertificateKeyPair) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *CertificateKeyPair) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateKeyPair) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateKeyPair) SetName(v string)`

SetName sets Name field to given value.


### GetFingerprintSha256

`func (o *CertificateKeyPair) GetFingerprintSha256() string`

GetFingerprintSha256 returns the FingerprintSha256 field if non-nil, zero value otherwise.

### GetFingerprintSha256Ok

`func (o *CertificateKeyPair) GetFingerprintSha256Ok() (*string, bool)`

GetFingerprintSha256Ok returns a tuple with the FingerprintSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintSha256

`func (o *CertificateKeyPair) SetFingerprintSha256(v string)`

SetFingerprintSha256 sets FingerprintSha256 field to given value.


### GetFingerprintSha1

`func (o *CertificateKeyPair) GetFingerprintSha1() string`

GetFingerprintSha1 returns the FingerprintSha1 field if non-nil, zero value otherwise.

### GetFingerprintSha1Ok

`func (o *CertificateKeyPair) GetFingerprintSha1Ok() (*string, bool)`

GetFingerprintSha1Ok returns a tuple with the FingerprintSha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintSha1

`func (o *CertificateKeyPair) SetFingerprintSha1(v string)`

SetFingerprintSha1 sets FingerprintSha1 field to given value.


### GetCertExpiry

`func (o *CertificateKeyPair) GetCertExpiry() time.Time`

GetCertExpiry returns the CertExpiry field if non-nil, zero value otherwise.

### GetCertExpiryOk

`func (o *CertificateKeyPair) GetCertExpiryOk() (*time.Time, bool)`

GetCertExpiryOk returns a tuple with the CertExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertExpiry

`func (o *CertificateKeyPair) SetCertExpiry(v time.Time)`

SetCertExpiry sets CertExpiry field to given value.


### GetCertSubject

`func (o *CertificateKeyPair) GetCertSubject() string`

GetCertSubject returns the CertSubject field if non-nil, zero value otherwise.

### GetCertSubjectOk

`func (o *CertificateKeyPair) GetCertSubjectOk() (*string, bool)`

GetCertSubjectOk returns a tuple with the CertSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertSubject

`func (o *CertificateKeyPair) SetCertSubject(v string)`

SetCertSubject sets CertSubject field to given value.


### GetPrivateKeyAvailable

`func (o *CertificateKeyPair) GetPrivateKeyAvailable() bool`

GetPrivateKeyAvailable returns the PrivateKeyAvailable field if non-nil, zero value otherwise.

### GetPrivateKeyAvailableOk

`func (o *CertificateKeyPair) GetPrivateKeyAvailableOk() (*bool, bool)`

GetPrivateKeyAvailableOk returns a tuple with the PrivateKeyAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyAvailable

`func (o *CertificateKeyPair) SetPrivateKeyAvailable(v bool)`

SetPrivateKeyAvailable sets PrivateKeyAvailable field to given value.


### GetPrivateKeyType

`func (o *CertificateKeyPair) GetPrivateKeyType() string`

GetPrivateKeyType returns the PrivateKeyType field if non-nil, zero value otherwise.

### GetPrivateKeyTypeOk

`func (o *CertificateKeyPair) GetPrivateKeyTypeOk() (*string, bool)`

GetPrivateKeyTypeOk returns a tuple with the PrivateKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyType

`func (o *CertificateKeyPair) SetPrivateKeyType(v string)`

SetPrivateKeyType sets PrivateKeyType field to given value.


### SetPrivateKeyTypeNil

`func (o *CertificateKeyPair) SetPrivateKeyTypeNil(b bool)`

 SetPrivateKeyTypeNil sets the value for PrivateKeyType to be an explicit nil

### UnsetPrivateKeyType
`func (o *CertificateKeyPair) UnsetPrivateKeyType()`

UnsetPrivateKeyType ensures that no value is present for PrivateKeyType, not even an explicit nil
### GetCertificateDownloadUrl

`func (o *CertificateKeyPair) GetCertificateDownloadUrl() string`

GetCertificateDownloadUrl returns the CertificateDownloadUrl field if non-nil, zero value otherwise.

### GetCertificateDownloadUrlOk

`func (o *CertificateKeyPair) GetCertificateDownloadUrlOk() (*string, bool)`

GetCertificateDownloadUrlOk returns a tuple with the CertificateDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateDownloadUrl

`func (o *CertificateKeyPair) SetCertificateDownloadUrl(v string)`

SetCertificateDownloadUrl sets CertificateDownloadUrl field to given value.


### GetPrivateKeyDownloadUrl

`func (o *CertificateKeyPair) GetPrivateKeyDownloadUrl() string`

GetPrivateKeyDownloadUrl returns the PrivateKeyDownloadUrl field if non-nil, zero value otherwise.

### GetPrivateKeyDownloadUrlOk

`func (o *CertificateKeyPair) GetPrivateKeyDownloadUrlOk() (*string, bool)`

GetPrivateKeyDownloadUrlOk returns a tuple with the PrivateKeyDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyDownloadUrl

`func (o *CertificateKeyPair) SetPrivateKeyDownloadUrl(v string)`

SetPrivateKeyDownloadUrl sets PrivateKeyDownloadUrl field to given value.


### GetManaged

`func (o *CertificateKeyPair) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *CertificateKeyPair) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *CertificateKeyPair) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *CertificateKeyPair) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *CertificateKeyPair) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *CertificateKeyPair) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


