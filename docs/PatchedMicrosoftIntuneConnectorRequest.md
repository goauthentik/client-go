# PatchedMicrosoftIntuneConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorUuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**CertificateAuthorities** | Pointer to **[]string** | Certificate authorities which issue device certificates via Intune (Cloud PKI, SCEP or PKCS profiles), used by the endpoint stage to validate certificates. This option has a higher priority than the &#x60;client_certificate&#x60; option on &#x60;Brand&#x60;. | [optional] 

## Methods

### NewPatchedMicrosoftIntuneConnectorRequest

`func NewPatchedMicrosoftIntuneConnectorRequest() *PatchedMicrosoftIntuneConnectorRequest`

NewPatchedMicrosoftIntuneConnectorRequest instantiates a new PatchedMicrosoftIntuneConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedMicrosoftIntuneConnectorRequestWithDefaults

`func NewPatchedMicrosoftIntuneConnectorRequestWithDefaults() *PatchedMicrosoftIntuneConnectorRequest`

NewPatchedMicrosoftIntuneConnectorRequestWithDefaults instantiates a new PatchedMicrosoftIntuneConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorUuid

`func (o *PatchedMicrosoftIntuneConnectorRequest) GetConnectorUuid() string`

GetConnectorUuid returns the ConnectorUuid field if non-nil, zero value otherwise.

### GetConnectorUuidOk

`func (o *PatchedMicrosoftIntuneConnectorRequest) GetConnectorUuidOk() (*string, bool)`

GetConnectorUuidOk returns a tuple with the ConnectorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorUuid

`func (o *PatchedMicrosoftIntuneConnectorRequest) SetConnectorUuid(v string)`

SetConnectorUuid sets ConnectorUuid field to given value.

### HasConnectorUuid

`func (o *PatchedMicrosoftIntuneConnectorRequest) HasConnectorUuid() bool`

HasConnectorUuid returns a boolean if a field has been set.

### GetName

`func (o *PatchedMicrosoftIntuneConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedMicrosoftIntuneConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedMicrosoftIntuneConnectorRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedMicrosoftIntuneConnectorRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedMicrosoftIntuneConnectorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedMicrosoftIntuneConnectorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedMicrosoftIntuneConnectorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedMicrosoftIntuneConnectorRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetClientId

`func (o *PatchedMicrosoftIntuneConnectorRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PatchedMicrosoftIntuneConnectorRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PatchedMicrosoftIntuneConnectorRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PatchedMicrosoftIntuneConnectorRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *PatchedMicrosoftIntuneConnectorRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PatchedMicrosoftIntuneConnectorRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PatchedMicrosoftIntuneConnectorRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PatchedMicrosoftIntuneConnectorRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetTenantId

`func (o *PatchedMicrosoftIntuneConnectorRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PatchedMicrosoftIntuneConnectorRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PatchedMicrosoftIntuneConnectorRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PatchedMicrosoftIntuneConnectorRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetCertificateAuthorities

`func (o *PatchedMicrosoftIntuneConnectorRequest) GetCertificateAuthorities() []string`

GetCertificateAuthorities returns the CertificateAuthorities field if non-nil, zero value otherwise.

### GetCertificateAuthoritiesOk

`func (o *PatchedMicrosoftIntuneConnectorRequest) GetCertificateAuthoritiesOk() (*[]string, bool)`

GetCertificateAuthoritiesOk returns a tuple with the CertificateAuthorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorities

`func (o *PatchedMicrosoftIntuneConnectorRequest) SetCertificateAuthorities(v []string)`

SetCertificateAuthorities sets CertificateAuthorities field to given value.

### HasCertificateAuthorities

`func (o *PatchedMicrosoftIntuneConnectorRequest) HasCertificateAuthorities() bool`

HasCertificateAuthorities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


