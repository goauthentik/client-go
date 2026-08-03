# CertificateKeyPairRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CertificateData** | **string** | PEM-encoded Certificate data | 
**KeyData** | Pointer to **string** | Optional Private Key. If this is set, you can use this keypair for encryption. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


