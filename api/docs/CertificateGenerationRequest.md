# CertificateGenerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonName** | **string** |  | 
**SubjectAltName** | Pointer to **string** |  | [optional] 
**ValidityDays** | **int32** |  | 

## Methods

### NewCertificateGenerationRequest

`func NewCertificateGenerationRequest(commonName string, validityDays int32, ) *CertificateGenerationRequest`

NewCertificateGenerationRequest instantiates a new CertificateGenerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateGenerationRequestWithDefaults

`func NewCertificateGenerationRequestWithDefaults() *CertificateGenerationRequest`

NewCertificateGenerationRequestWithDefaults instantiates a new CertificateGenerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonName

`func (o *CertificateGenerationRequest) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CertificateGenerationRequest) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CertificateGenerationRequest) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetSubjectAltName

`func (o *CertificateGenerationRequest) GetSubjectAltName() string`

GetSubjectAltName returns the SubjectAltName field if non-nil, zero value otherwise.

### GetSubjectAltNameOk

`func (o *CertificateGenerationRequest) GetSubjectAltNameOk() (*string, bool)`

GetSubjectAltNameOk returns a tuple with the SubjectAltName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAltName

`func (o *CertificateGenerationRequest) SetSubjectAltName(v string)`

SetSubjectAltName sets SubjectAltName field to given value.

### HasSubjectAltName

`func (o *CertificateGenerationRequest) HasSubjectAltName() bool`

HasSubjectAltName returns a boolean if a field has been set.

### GetValidityDays

`func (o *CertificateGenerationRequest) GetValidityDays() int32`

GetValidityDays returns the ValidityDays field if non-nil, zero value otherwise.

### GetValidityDaysOk

`func (o *CertificateGenerationRequest) GetValidityDaysOk() (*int32, bool)`

GetValidityDaysOk returns a tuple with the ValidityDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDays

`func (o *CertificateGenerationRequest) SetValidityDays(v int32)`

SetValidityDays sets ValidityDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


