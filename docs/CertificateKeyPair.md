# CertificateKeyPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**FingerprintSha256** | **NullableString** | SHA256 fingerprint of the certificate | [readonly] 
**FingerprintSha1** | **NullableString** | SHA1 fingerprint of the certificate | [readonly] 
**CertExpiry** | **NullableTime** | Certificate expiry date | [readonly] 
**CertSubject** | **NullableString** | Certificate subject as RFC4514 string | [readonly] 
**PrivateKeyAvailable** | **bool** | Show if this keypair has a private key configured or not | [readonly] 
**KeyType** | [**NullableKeyTypeEnum**](KeyTypeEnum.md) | Key algorithm type detected from the certificate&#39;s public key | [readonly] 
**CertificateDownloadUrl** | **string** | Get URL to download certificate | [readonly] 
**PrivateKeyDownloadUrl** | **string** | Get URL to download private key | [readonly] 
**Managed** | **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [readonly] 

## Methods

### NewCertificateKeyPair

`func NewCertificateKeyPair(pk string, name string, fingerprintSha256 NullableString, fingerprintSha1 NullableString, certExpiry NullableTime, certSubject NullableString, privateKeyAvailable bool, keyType NullableKeyTypeEnum, certificateDownloadUrl string, privateKeyDownloadUrl string, managed NullableString, ) *CertificateKeyPair`

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


### SetFingerprintSha256Nil

`func (o *CertificateKeyPair) SetFingerprintSha256Nil(b bool)`

 SetFingerprintSha256Nil sets the value for FingerprintSha256 to be an explicit nil

### UnsetFingerprintSha256
`func (o *CertificateKeyPair) UnsetFingerprintSha256()`

UnsetFingerprintSha256 ensures that no value is present for FingerprintSha256, not even an explicit nil
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


### SetFingerprintSha1Nil

`func (o *CertificateKeyPair) SetFingerprintSha1Nil(b bool)`

 SetFingerprintSha1Nil sets the value for FingerprintSha1 to be an explicit nil

### UnsetFingerprintSha1
`func (o *CertificateKeyPair) UnsetFingerprintSha1()`

UnsetFingerprintSha1 ensures that no value is present for FingerprintSha1, not even an explicit nil
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


### SetCertExpiryNil

`func (o *CertificateKeyPair) SetCertExpiryNil(b bool)`

 SetCertExpiryNil sets the value for CertExpiry to be an explicit nil

### UnsetCertExpiry
`func (o *CertificateKeyPair) UnsetCertExpiry()`

UnsetCertExpiry ensures that no value is present for CertExpiry, not even an explicit nil
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


### SetCertSubjectNil

`func (o *CertificateKeyPair) SetCertSubjectNil(b bool)`

 SetCertSubjectNil sets the value for CertSubject to be an explicit nil

### UnsetCertSubject
`func (o *CertificateKeyPair) UnsetCertSubject()`

UnsetCertSubject ensures that no value is present for CertSubject, not even an explicit nil
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


### GetKeyType

`func (o *CertificateKeyPair) GetKeyType() KeyTypeEnum`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CertificateKeyPair) GetKeyTypeOk() (*KeyTypeEnum, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CertificateKeyPair) SetKeyType(v KeyTypeEnum)`

SetKeyType sets KeyType field to given value.


### SetKeyTypeNil

`func (o *CertificateKeyPair) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *CertificateKeyPair) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
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


### SetManagedNil

`func (o *CertificateKeyPair) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *CertificateKeyPair) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


