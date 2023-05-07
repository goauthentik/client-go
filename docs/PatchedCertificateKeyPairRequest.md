# PatchedCertificateKeyPairRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**CertificateData** | Pointer to **string** | PEM-encoded Certificate data | [optional] 
**KeyData** | Pointer to **string** | Optional Private Key. If this is set, you can use this keypair for encryption. | [optional] 

## Methods

### NewPatchedCertificateKeyPairRequest

`func NewPatchedCertificateKeyPairRequest() *PatchedCertificateKeyPairRequest`

NewPatchedCertificateKeyPairRequest instantiates a new PatchedCertificateKeyPairRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCertificateKeyPairRequestWithDefaults

`func NewPatchedCertificateKeyPairRequestWithDefaults() *PatchedCertificateKeyPairRequest`

NewPatchedCertificateKeyPairRequestWithDefaults instantiates a new PatchedCertificateKeyPairRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedCertificateKeyPairRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedCertificateKeyPairRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedCertificateKeyPairRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedCertificateKeyPairRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCertificateData

`func (o *PatchedCertificateKeyPairRequest) GetCertificateData() string`

GetCertificateData returns the CertificateData field if non-nil, zero value otherwise.

### GetCertificateDataOk

`func (o *PatchedCertificateKeyPairRequest) GetCertificateDataOk() (*string, bool)`

GetCertificateDataOk returns a tuple with the CertificateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateData

`func (o *PatchedCertificateKeyPairRequest) SetCertificateData(v string)`

SetCertificateData sets CertificateData field to given value.

### HasCertificateData

`func (o *PatchedCertificateKeyPairRequest) HasCertificateData() bool`

HasCertificateData returns a boolean if a field has been set.

### GetKeyData

`func (o *PatchedCertificateKeyPairRequest) GetKeyData() string`

GetKeyData returns the KeyData field if non-nil, zero value otherwise.

### GetKeyDataOk

`func (o *PatchedCertificateKeyPairRequest) GetKeyDataOk() (*string, bool)`

GetKeyDataOk returns a tuple with the KeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyData

`func (o *PatchedCertificateKeyPairRequest) SetKeyData(v string)`

SetKeyData sets KeyData field to given value.

### HasKeyData

`func (o *PatchedCertificateKeyPairRequest) HasKeyData() bool`

HasKeyData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


