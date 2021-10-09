# CertificateKeyPairRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CertificateData** | **string** | PEM-encoded Certificate data | 
**KeyData** | Pointer to **string** | Optional Private Key. If this is set, you can use this keypair for encryption. | [optional] 
**Managed** | Pointer to **NullableString** | Objects which are managed by authentik. These objects are created and updated automatically. This is flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [optional] 

## Methods

### NewCertificateKeyPairRequest

`func NewCertificateKeyPairRequest(name string, certificateData string, ) *CertificateKeyPairRequest`

NewCertificateKeyPairRequest instantiates a new CertificateKeyPairRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateKeyPairRequestWithDefaults

`func NewCertificateKeyPairRequestWithDefaults() *CertificateKeyPairRequest`

NewCertificateKeyPairRequestWithDefaults instantiates a new CertificateKeyPairRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertificateKeyPairRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateKeyPairRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateKeyPairRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCertificateData

`func (o *CertificateKeyPairRequest) GetCertificateData() string`

GetCertificateData returns the CertificateData field if non-nil, zero value otherwise.

### GetCertificateDataOk

`func (o *CertificateKeyPairRequest) GetCertificateDataOk() (*string, bool)`

GetCertificateDataOk returns a tuple with the CertificateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateData

`func (o *CertificateKeyPairRequest) SetCertificateData(v string)`

SetCertificateData sets CertificateData field to given value.


### GetKeyData

`func (o *CertificateKeyPairRequest) GetKeyData() string`

GetKeyData returns the KeyData field if non-nil, zero value otherwise.

### GetKeyDataOk

`func (o *CertificateKeyPairRequest) GetKeyDataOk() (*string, bool)`

GetKeyDataOk returns a tuple with the KeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyData

`func (o *CertificateKeyPairRequest) SetKeyData(v string)`

SetKeyData sets KeyData field to given value.

### HasKeyData

`func (o *CertificateKeyPairRequest) HasKeyData() bool`

HasKeyData returns a boolean if a field has been set.

### GetManaged

`func (o *CertificateKeyPairRequest) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *CertificateKeyPairRequest) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *CertificateKeyPairRequest) SetManaged(v string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *CertificateKeyPairRequest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *CertificateKeyPairRequest) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *CertificateKeyPairRequest) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


